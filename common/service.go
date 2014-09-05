package common

import (
	"errors"
	"fmt"
	"sort"
	"bytes"
	"strings"
	"time"
	"net/http"
	"net/textproto"
	//"net/url"
)

type AWSService struct {
	// ref: http://docs.aws.amazon.com/general/latest/gr/rande.html
	// TO-DO: convert endpoint to region & service.
	Region            string 	//  "us-west-1", etc.
	Service           string 	//  "ec2", "s3", etc.
	Endpoint          string	// Note: S3 doesn't support hosted style, use path style always.
	Credentials    *Credentials
}

func (this *AWSService) Init(cred *Credentials) (*AWSService,error) {
	if cred == nil || cred.Expired() {
		return nil, errors.New("Credentials incorrect")
	}
	this.Credentials = cred

	return this,nil
}


func (this *AWSService) Sign4(req *AWSRequest, useAuthorizationHeader bool) (*http.Request, error) {

	// doc: http://docs.aws.amazon.com/general/latest/gr/sigv4_signing.html
	// s3 doc: http://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html
	//// clone

	pars := CloneHeader(req.Parameters, func(k string, v string) (string, string) {
			k2 := textproto.CanonicalMIMEHeaderKey(k)
			if strings.HasPrefix(k2, "X-Amz-") {
				return k2, v
			}
			return k, v
		})
	headers := CloneHeader(req.Headers, func(k string, v string) (string, string) {
			return textproto.CanonicalMIMEHeaderKey(k), strings.TrimSpace(v)
		})
	if _, ok := headers["Host"]; !ok {
		headers["Host"] = this.Endpoint
	}

	var Scope string
	var X_Amz_Date string

	var HashedPayload string            /// == X-Amz-Content-Sha256
	//var Content
	if req.Payload.Len() == 0 {
		HashedPayload = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	} else {
		HashedPayload = HashSHA256(req.Payload.Bytes())
		if _, ok := headers["Content-Type"]; !ok {
			return nil, errors.New("Content-Type should be added.")
		}
	}

	if useAuthorizationHeader {
		if _, ok := headers["X-Amz-Date"]; !ok {
			headers["X-Amz-Date"] = time.Now().UTC().Format("20060102T150405Z")
		}
		X_Amz_Date = headers["X-Amz-Date"]
		Scope = Concat("/", X_Amz_Date[:8], this.Region, this.Service, "aws4_request")
		if _, ok := headers["X-Amz-Expires"]; !ok {
			headers["X-Amz-Expires"] = "604800"
		}
		headers["X-Amz-Content-Sha256"] = HashedPayload
		if this.Credentials.SecurityToken != "" {
			headers["X-Amz-Security-Token"] = strings.TrimSpace(this.Credentials.SecurityToken)
		}
	} else {
		// X-Amz-Credential, X-Amz-Algorithm, X-Amz-Date, X-Amz-Expires, X-Amz-SignedHeaders, X-Amz-Signature
		if _, ok := pars["X-Amz-Date"]; !ok {
			pars["X-Amz-Date"] = time.Now().UTC().Format("20060102T150405Z")
		}
		X_Amz_Date = pars["X-Amz-Date"]
		Scope = Concat("/", X_Amz_Date[:8], this.Region, this.Service, "aws4_request")
		pars["X-Amz-Algorithm"] = "AWS4-HMAC-SHA256"
		pars["X-Amz-Credential"] = this.Credentials.AccessKeyID+"/"+Scope
		if _, ok := pars["X-Amz-Expires"]; !ok {
			pars["X-Amz-Expires"] = "604800"
		}
		//pars["X-Amz-Content-Sha256"] = HashedPayload		// amazon services seems not like this in query url
		if this.Credentials.SecurityToken != "" {
			pars["X-Amz-Security-Token"] = strings.TrimSpace(this.Credentials.SecurityToken)
		}
		//
	}


	CanonicalHeaders, SignedHeaders := EncodeHeaderString(headers)

	if useAuthorizationHeader {

	} else {
		pars["X-Amz-SignedHeaders"] = SignedHeaders
	}


	// Task 1: Create a Canonical Request
	HTTPMethod := req.Method
	CanonicalURI := EncodeUri(req.Path, false)
	CanonicalQueryString := EncodeQueryString(pars)

	CanonicalRequest := Concat("\n", HTTPMethod, CanonicalURI, CanonicalQueryString, CanonicalHeaders, SignedHeaders, HashedPayload)

	// Task 2: Create a String to Sign

	StringToSign := Concat("\n", "AWS4-HMAC-SHA256", X_Amz_Date, Scope, HashSHA256([]byte(CanonicalRequest)))


	// Task 3: Calculate Signature
	DateKey := HmacSHA256([]byte("AWS4"+this.Credentials.SecretAccessKey), X_Amz_Date[:8])
	DateRegionKey := HmacSHA256(DateKey, this.Region)
	DateRegionServiceKey := HmacSHA256(DateRegionKey, this.Service)
	SigningKey := HmacSHA256(DateRegionServiceKey, "aws4_request")

	Signature := fmt.Sprintf("%x", HmacSHA256(SigningKey, StringToSign))


	//
	if useAuthorizationHeader {
		headers["Authorization"] = "AWS4-HMAC-SHA256"+" "+"Credential="+this.Credentials.AccessKeyID+"/"+Scope +
				",SignedHeaders="+SignedHeaders+",Signature="+Signature

		httpReq, err := http.NewRequest(req.Method, "https://"+this.Endpoint+req.Path+EncodeURLString(pars), &(req.Payload))
		for k, v := range (headers) {
			httpReq.Header.Set(k, v)
		}

		if DEBUG_VERBOSE!=0 {
			fmt.Printf("HTTP REQ: %s\n", httpReq)
		}
		return httpReq, err
	} else {
		// generate query string
		if req.Payload.Len() > 0 {
			return nil, errors.New("Payload is not nil, and should not use query string")
		}

		pars["X-Amz-Signature"] = Signature

		httpReq, err := http.NewRequest(req.Method, "https://"+this.Endpoint+req.Path+EncodeURLString(pars), nil)
		for k, v := range (headers) {
			httpReq.Header.Set(k, v)
		}

		if DEBUG_VERBOSE!=0 {
			fmt.Printf("HTTP REQ: %s\n", httpReq)
		}
		return httpReq, err
	}


	return nil, errors.New("Not implement")

}

func EncodeUri(value string, encodeSlash bool) string {
	var buf bytes.Buffer

	for _, ch := range (value) {
		if ((ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') || ch == '_' || ch == '-' || ch == '~' || ch == '.' || (ch == '/' && !encodeSlash)) {
			buf.WriteString(string(ch))
		} else {
			buf.WriteString(GenUTF8Hex(ch));
		}
	}
	return buf.String()
}


func GenUTF8Hex(ch rune) string {
	if (ch&0x7f000000) != 0 || ch < 0 {
		return fmt.Sprintf("%%%02X%%%02X%%%02X%%%02X", (ch>>24)&0x0ff, (ch>>16)&0x0ff, (ch>>8)&0x0ff, (ch>>0)&0x0ff)
	} else if (ch&0x0ff0000) != 0 {
		return fmt.Sprintf("%%%02X%%%02X%%%02X", (ch>>16)&0x0ff, (ch>>8)&0x0ff, (ch>>0)&0x0ff)
	} else if (ch&0x0ff00) != 0 {
		return fmt.Sprintf("%%%02X%%%02X", (ch>>8)&0x0ff, (ch>>0)&0x0ff)
	} else {
		return fmt.Sprintf("%%%02X", (ch>>0)&0x0ff)
	}
}

func EncodeQueryString(pars map[string]string) string {
	var buf bytes.Buffer

	keys := make([]string, 0, len(pars))
	keysvalues := make(map[string]string)

	for k, v := range pars {
		k = EncodeUri(k, true)
		v = EncodeUri(v, true)

		keys = append(keys, k)
		keysvalues[k] = v
	}

	sort.Strings(keys)
	for _, k := range keys {
		v := keysvalues[k]

		if buf.Len() > 0 {
			buf.WriteRune(rune('&'))
		}

		buf.WriteString(k)
		buf.WriteRune(rune('='))
		buf.WriteString(v)
	}

	return buf.String()
}

func EncodeURLString(pars map[string]string) string {
	if len(pars) == 0 {
		return ""
	}

	keys := make([]string, 0, len(pars))
	keysvalues := make(map[string]string)

	for k, v := range pars {
		k = EncodeUri(k, false)
		v = EncodeUri(v, false)

		keys = append(keys, k)
		keysvalues[k] = v
	}

	var buf bytes.Buffer

	buf.WriteRune(rune('?'))

	sort.Strings(keys)
	for _, k := range keys {
		v := keysvalues[k]

		if buf.Len() > 1 {
			buf.WriteRune(rune('&'))
		}

		buf.WriteString(k)
		buf.WriteRune(rune('='))
		buf.WriteString(v)
	}

	return buf.String()
}


func CloneHeader(h map[string]string, enc func(string, string) (string, string)) map[string]string {
	h2 := make(map[string]string, len(h))
	for k, v := range h {
		k2, v2 := enc(k, v)
		//vv2 := make([]string, len(vv))
		//copy(vv2, vv)
		h2[k2] = v2
	}
	return h2
}


func EncodeHeaderString(h map[string]string) (string, string) {
	var buf1 bytes.Buffer        // CanonicalHeaders
	var buf2 bytes.Buffer        // SignedHeaders

	keys := make([]string, 0, len(h))

	for k := range (h) {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range (keys) {
		v := h[k]
		k2 := strings.ToLower(k)

		buf1.WriteString(k2)
		buf1.WriteRune(rune(':'))
		buf1.WriteString(v)
		buf1.WriteRune(rune('\n'))

		if buf2.Len() > 0 {
			buf2.WriteRune(rune(';'))
			buf2.WriteString(k2)
		} else {
			buf2.WriteString(k2)
		}
	}

	return buf1.String(), buf2.String()
}




