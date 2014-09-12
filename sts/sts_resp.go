package sts

import (
	"github.com/letscool/aws/common"
	"net/http"
	"encoding/xml"
	"io/ioutil"
	//"errors"
	//"strings"
	//"fmt"
)

type Error struct {
	Type       string
	Code       string
	Message    string
}

type ErrorResponse struct {
	Err		     Error			`xml:"Error"`
	RequestId    string
}

func (this *ErrorResponse) Error() string {
	return this.Err.Code + ":" + this.Err.Message
}

type STSResp struct {
	common.AWSResponse

	ErrorResponse    *ErrorResponse
}

func (this *STSResp) Init(req *STSReq, resp *http.Response) (*STSResp, error) {
	if _, err := this.AWSResponse.Init(&req.AWSRequest, resp); err != nil {
		return this, err
	}

	this.ErrorResponse = nil
	if resp.StatusCode >= 300 {
		this.ErrorResponse = new(ErrorResponse)

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return this, err
		}

		//fmt.Printf("%s\n", string(body))

		err = xml.Unmarshal(body, this.ErrorResponse)
		if err != nil {
			return this, err
		}
		return this, this.ErrorResponse
	}

	return this, nil
}

type GetFederationTokenResult struct {
	Credentials            Credentials
	FederatedUser          FederatedUser
	PackedPolicySize       int
}

type ResponseMetadata struct {
	RequestId        string
}

type Credentials struct {
	SessionToken       string
	SecretAccessKey    string
	Expiration         string
	AccessKeyId        string
}

type FederatedUser struct {
	Arn                string
	FederatedUserId    string
}

type GetFederationTokenResp struct {
	STSResp            `xml:"-"`

	GetFederationTokenResult    GetFederationTokenResult
	ResponseMetadata            ResponseMetadata
}

func (this *GetFederationTokenResp) Init(req *GetFederationTokenReq, resp *http.Response) (*GetFederationTokenResp, error) {
	if _, err := this.STSResp.Init(&req.STSReq, resp); err != nil {
		return this, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return this, err
	}

	//fmt.Printf("%s\n", string(body))

	err = xml.Unmarshal(body, this)
	if err != nil {
		return this, err
	}

	return this, nil
}




