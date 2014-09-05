package ec2

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"strings"
	"github.com/letscool/aws/common"
)

/*
	EC2 error handle: http://docs.aws.amazon.com/AWSEC2/latest/APIReference/api-error-codes.html
 */
type Errors struct {
	RequestID	string		`xml:"RequestID"`
	Error	[]Error			`xml:"Errors>Error"`
}

type Error struct {
	Code	string
	Message	string
}

type EC2Resp struct {
	common.AWSResponse
	Errors	*Errors
}

func (this *EC2Resp) Init(req *EC2Req, resp *http.Response) (*EC2Resp, error) {
	if _, err := this.AWSResponse.Init(&req.AWSRequest, resp); err != nil {
		return nil, err
	}

	this.Errors = nil

	if resp.StatusCode >= 300 {
		this.Errors = new(Errors)

		if strings.ToLower(resp.Header.Get("Content-Type"))== "application/xml" {
			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				return nil, err
			}

			err = xml.Unmarshal(body, this.Errors)
			if err!=nil {
				return nil, err
			}
		} else {
			// impossible, and no ideal how to deal.
		}
	}

	return this, nil
}


type ReservationSet struct {

}

type DescribeInstancesResp struct {
	EC2Resp				`xml:"-"`

	RequestId	string	`xml:"requestId"`
	ReservationSet	ReservationSet	`xml:"reservationSet"`
	NextToken	string	`xml:"nextToken"`
}

func (this *DescribeInstancesResp) Init(req *DescribeInstancesReq, resp *http.Response) (*DescribeInstancesResp, error) {
	if _, err := this.EC2Resp.Init(&req.EC2Req, resp); err != nil {
		return nil, err
	}

	if this.Errors!=nil {
		return this, nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, this)
	if err!=nil {
		return nil, err
	}

	return this, nil
}





