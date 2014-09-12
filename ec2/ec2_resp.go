package ec2

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	//"strings"
	"github.com/letscool/aws/common"
	//"fmt"
)

/*
	EC2 error handle: http://docs.aws.amazon.com/AWSEC2/latest/APIReference/api-error-codes.html
 */
type Errors struct {
	RequestID	string			`xml:"RequestID"`
	Errors		[]Error			`xml:"Errors>Error"`
}

func (this *Errors) Error() string {
	var str string
	for _, err := range(this.Errors) {
		if len(str)>0 {
			str = str + "\n"
		}
		str = str + err.Code + ":" + err.Message
	}
	return str
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
		return this, err
	}

	this.Errors = nil

	if resp.StatusCode >= 300 {
		this.Errors = new(Errors)
		//fmt.Println(resp.Header)

		//if strings.ToLower(resp.Header.Get("Content-Type"))== "application/xml" {
			body, err := ioutil.ReadAll(resp.Body)
			//fmt.Println(string(body))

			if err != nil {
				return this, err
			}

			err = xml.Unmarshal(body, this.Errors)
			if err!=nil {
				return this, err
			} else {
				return this, this.Errors
			}
		//} else {
			// impossible, and no ideal how to deal.
		//}
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
		return this, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return this, err
	}

	err = xml.Unmarshal(body, this)
	if err!=nil {
		return this, err
	}

	return this, nil
}





