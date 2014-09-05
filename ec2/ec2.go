package ec2

import (
	"net/http"
	"github.com/letscool/aws/common"
)

// ref: http://docs.aws.amazon.com/AWSEC2/latest/APIReference/OperationList-query.html

const (
	EC2_API_VERSION = "2014-06-15"
)

type EC2 struct {
	common.AWSService
}

func NewEC2(cred *common.Credentials) (*EC2, error) {
	return new(EC2).Init(cred)
}

func (this *EC2) Init(cred *common.Credentials) (*EC2, error) {
	if _, err := this.AWSService.Init(cred); err!=nil {
		return nil, err
	}

	this.Region = "us-east-1"
	this.Service = "ec2"
	this.Endpoint = "ec2.us-east-1.amazonaws.com"

	return this, nil
}

func (this *EC2) DescribeInstances(req *DescribeInstancesReq) (resp *DescribeInstancesResp, err error) {

	httpReq, err := this.Sign4(&req.AWSRequest, false)
	if err != nil {
		return nil, err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return new(DescribeInstancesResp).Init(req, httpResp)
}













