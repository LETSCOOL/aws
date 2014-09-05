package dynamodb

import (
	"github.com/letscool/aws/common"
	"fmt"
	"encoding/json"
)

type DynamoDBReq struct {
	common.AWSRequest

}

type ListTablesReq struct {
	DynamoDBReq                				`json:"-"`
	ExclusiveStartTableName    string       `json:",omitempty"`
	Limit                      uint         `json:",string,omitempty"`
}

func (this *ListTablesReq) Init() (*ListTablesReq) {
	if this.DynamoDBReq.Init() == nil {
		return nil
	}

	this.Method = "POST"

	this.Headers["X-Amz-Target"] = "DynamoDB_"+DynamoDB_API_VERSION+"."+"ListTables"
	this.Headers["Content-Type"] = "application/x-amz-json-1.0"

	return this
}

func (this *ListTablesReq) WithExclusiveStartTableName(exclusive string) (*ListTablesReq) {
	this.ExclusiveStartTableName = exclusive
	return this
}

func (this *ListTablesReq) generatePayload() {
	this.Payload.Truncate(0)
	marshal, _ := json.Marshal(this)
	this.Payload.WriteString(string(marshal))

	if common.DEBUG_VERBOSE!=0 {
		fmt.Printf("Payload: %s\n", string(this.Payload.Bytes()))
	}
}



type DescribeTableReq struct {
	DynamoDBReq					`json:"-"`
	TableName		string		`json:"TableName"`
}

func (this *DescribeTableReq) Init() (*DescribeTableReq) {
	if this.DynamoDBReq.Init() == nil {
		return nil
	}

	this.Method = "POST"

	this.Headers["X-Amz-Target"] = "DynamoDB_"+DynamoDB_API_VERSION+"."+"DescribeTable"
	this.Headers["Content-Type"] = "application/x-amz-json-1.0"

	return this
}


func (this *DescribeTableReq) generatePayload() {
	this.Payload.Truncate(0)
	marshal, _ := json.Marshal(this)
	this.Payload.WriteString(string(marshal))

	if common.DEBUG_VERBOSE!=0 {
		fmt.Printf("Payload: %s\n", string(this.Payload.Bytes()))
	}
}
