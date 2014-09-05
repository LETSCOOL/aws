package ec2

import (
	"github.com/letscool/aws/common"
)

type EC2Req struct {
	common.AWSRequest
}


type DescribeInstancesReq struct {
	EC2Req
}

func (this *DescribeInstancesReq) Init() (*DescribeInstancesReq) {
	if this.EC2Req.Init() == nil {
		return nil
	}

	this.Method = "GET"

	this.Parameters["Version"] = EC2_API_VERSION
	this.Parameters["Action"] = "DescribeInstances"

	//this.Headers["Content-Type"] = "application/json"

	return this
}







