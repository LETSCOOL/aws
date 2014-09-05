package s3

import (
	"github.com/letscool/aws/common"
	"net/http"
	"encoding/xml"
	"io/ioutil"
	"errors"
	"strings"
)

/**
	S3 error handle: http://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html
 */
type Error struct {
	Code	string			`xml`
	Message	string			`xml`
	RequestId	string		`xml`
	Resource	string		`xml`
}

/**
	S3 base type, handle error code.
 */
type S3Resp struct {
	common.AWSResponse
	Error	*Error			// response error code from server
}

func (this *S3Resp) Init(req *S3Req, resp *http.Response) (*S3Resp, error) {
	if _, err := this.AWSResponse.Init(&req.AWSRequest, resp); err != nil {
		return nil, err
	}

	this.Error = nil

	if resp.StatusCode >= 300 {
		this.Error = new(Error)

		if strings.ToLower(resp.Header.Get("Content-Type"))== "application/xml" {
			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				return nil, err
			}

			err = xml.Unmarshal(body, this.Error)
			if err!=nil {
				return nil, err
			}
		} else {
			// impossible, and no ideal how to deal.
		}
	}
	
	return this, nil
}


type Owner struct {
	ID	string
	DisplayName	string
}

type Buckets struct {
	Buckets	[]Bucket	`xml:"Bucket"`
}

type Bucket struct {
	Name	string
	CreationDate	string
}

type GetServiceResp struct {
	S3Resp		`xml:"-"`

	Owner	Owner			`xml`
	Buckets	Buckets			`xml`
}

func (this *GetServiceResp) Init(req *GetServiceReq, resp *http.Response) (*GetServiceResp, error) {
	if _, err := this.S3Resp.Init(&req.S3Req, resp); err != nil {
		return nil, err
	}

	if this.Error!=nil {
		return this, nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	//fmt.Printf("%s\n", string(body))

	err = xml.Unmarshal(body, this)
	if err!=nil {
		return nil, err
	}


	return this, nil
}


type NewBucketResp struct {
	S3Resp
}

func (this *NewBucketResp) Init(req *NewBucketReq, resp *http.Response) (*NewBucketResp, error) {
	if _, err := this.S3Resp.Init(&req.S3Req, resp); err != nil {
		return nil, err
	}

	if this.Error!=nil {
		return this, nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if (resp.StatusCode/100) != 2 {
		return nil, errors.New(string(body))
	}

	return this, nil
}


type ExistBucketResp struct {
	S3Resp
	Exists	bool
}

func (this *ExistBucketResp) Init(req *ExistBucketReq, resp *http.Response) (*ExistBucketResp, error) {
	if _, err := this.S3Resp.Init(&req.S3Req, resp); err != nil {
		return nil, err
	}

	this.Exists	= (this.Error==nil)

	return this, nil
}


type DeleteBucketResp struct {
	S3Resp
}

func (this *DeleteBucketResp) Init(req *DeleteBucketReq, resp *http.Response) (*DeleteBucketResp, error) {
	if _, err := this.S3Resp.Init(&req.S3Req, resp); err != nil {
		return nil, err
	}

	/*if this.Error!=nil {
		return this, nil
	}*/

	return this, nil
}
