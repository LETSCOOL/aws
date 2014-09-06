package sts

import (
	"fmt"
	"github.com/letscool/aws/common"
)

type STSReq struct {
	common.AWSRequest
}

type GetFederationTokenReq struct {
	STSReq

	DurationSeconds    int    // 900 seconds(15 minutes) ~ 129600 seconds(36 hours)
	Name               string // 2 <= Length(Name) <= 12
	Policy             string // json format, 0, 1 <= Length(Policy) <= 2048
}

func (this *GetFederationTokenReq) Init() (*GetFederationTokenReq) {
	if this.STSReq.Init() == nil {
		return nil
	}

	this.Method = "GET"
	this.DurationSeconds = 43200

	return this
}

func (this *GetFederationTokenReq) generatePayload() {
	this.Parameters["DurationSeconds"] = fmt.Sprintf("%d", this.DurationSeconds)
	this.Parameters["Name"] = this.Name
	this.Parameters["Action"] = "GetFederationToken"
	this.Parameters["Version"] = STS_API_VERSION
	if this.Policy != "" {
		this.Parameters["Policy"] = this.Policy
	}
}



