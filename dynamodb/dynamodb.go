package dynamodb

import (
	"github.com/letscool/aws/common"
	"net/http"
)

const (
	DynamoDB_API_VERSION = "20120810"
)


/**
	ref: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/API.html
 */
type DynamoDB struct {
	common.AWSService
}

func NewDynamoDB(cred *common.Credentials) (*DynamoDB,error) {
	return new(DynamoDB).Init(cred)
}

func (this *DynamoDB) Init(cred *common.Credentials) (*DynamoDB,error) {
	if _, err := this.AWSService.Init(cred); err!=nil {
		return nil,err
	}

	this.Region = "us-east-1"
	this.Service = "dynamodb"
	this.Endpoint = "dynamodb.us-east-1.amazonaws.com"

	return this,nil
}


func (this *DynamoDB) ListTables(req *ListTablesReq) (*ListTablesResp, error) {
	req.generatePayload()

	httpReq, err := this.Sign4(&req.AWSRequest, true)
	if err != nil {
		return nil, err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return new(ListTablesResp).Init(req, httpResp)
}


func (this *DynamoDB) DescribeTable(req *DescribeTableReq) (*DescribeTableResp, error) {
	req.generatePayload()

	httpReq, err := this.Sign4(&req.AWSRequest, true)
	if err != nil {
		return nil, err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return new(DescribeTableResp).Init(req, httpResp)
}
