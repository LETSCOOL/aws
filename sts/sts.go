package sts

import (
	"github.com/letscool/aws/common"
	"net/http"
)

const (
	STS_API_VERSION = "2011-06-15"
)

/**
	ref: http://docs.aws.amazon.com/STS/latest/APIReference/Welcome.html
 */
type STS struct {
	common.AWSService
}

func NewSTS(cred *common.Credentials) (*STS, error) {
	return new(STS).Init(cred)
}

func (this *STS) Init(cred *common.Credentials) (*STS, error) {
	if _, err := this.AWSService.Init(cred); err != nil {
		return nil, err
	}

	this.Region = "us-east-1"
	this.Service = "sts"
	this.Endpoint = "sts.amazonaws.com"

	return this, nil
}


func (this *STS) GetFederationToken(req *GetFederationTokenReq) (resp *GetFederationTokenResp, err error) {
	req.generatePayload()

	httpReq, err := this.Sign4(&req.AWSRequest, false)
	if err != nil {
		return nil, err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return new(GetFederationTokenResp).Init(req, httpResp)
}

