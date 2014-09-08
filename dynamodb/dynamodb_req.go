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

// ================================= CreateTableReq

type AttributeDefinition struct {
	AttributeName	string
	AttributeType	string				// Hash/Range Key: S | N | B, Attribute: S | N | B | SS | NS | BS
}

type KeySchemaElement struct {
	AttributeName	string				// 1 <= Len(AttributeName) <= 255
	KeyType			string				// "HASH" or "RANGE"
}

type Projection struct {
	NonKeyAttributes		[]string			`json:",omitempty"` // one at least, no more than 20 attributes
	ProjectionType			string				`json:",omitempty"` // "KEYS_ONLY", "INCLUDE", "ALL"
}

type ProvisionedThroughput struct {
	ReadCapacityUnits		int					//`json:",string"`
	WriteCapacityUnits		int					//`json:",string"`
}

type GlobalSecondaryIndex struct {
	IndexName				string
	KeySchema				[]KeySchemaElement
	Projection				Projection
	ProvisionedThroughput	ProvisionedThroughput
}

type LocalSecondaryIndex struct {
	IndexName				string
	KeySchema				[]KeySchemaElement
	Projection				Projection
}

type CreateTableReq struct {
	DynamoDBReq							`json:"-"`

	TableName				string
	KeySchema				[]KeySchemaElement
	AttributeDefinitions	[]AttributeDefinition
	GlobalSecondaryIndexes	[]GlobalSecondaryIndex		`json:",omitempty"`
	LocalSecondaryIndexes	[]LocalSecondaryIndex		`json:",omitempty"`
	ProvisionedThroughput	ProvisionedThroughput

}

func (this *CreateTableReq) Init() (*CreateTableReq) {
	if this.DynamoDBReq.Init() == nil {
		return nil
	}

	this.Method = "POST"

	this.Headers["X-Amz-Target"] = "DynamoDB_"+DynamoDB_API_VERSION+"."+"CreateTable"
	this.Headers["Content-Type"] = "application/x-amz-json-1.0"

	return this
}


func (this *CreateTableReq) generatePayload() {
	this.Payload.Truncate(0)
	marshal, _ := json.Marshal(this)
	this.Payload.WriteString(string(marshal))

	if common.DEBUG_VERBOSE!=0 {
		fmt.Printf("Payload: %s\n", string(this.Payload.Bytes()))
	}
}


// ================================= DeleteTableReq

type DeleteTableReq struct {
	DynamoDBReq							`json:"-"`

	TableName				string
}

func (this *DeleteTableReq) Init() (*DeleteTableReq) {
	if this.DynamoDBReq.Init() == nil {
		return nil
	}

	this.Method = "POST"

	this.Headers["X-Amz-Target"] = "DynamoDB_"+DynamoDB_API_VERSION+"."+"DeleteTable"
	this.Headers["Content-Type"] = "application/x-amz-json-1.0"

	return this
}


func (this *DeleteTableReq) generatePayload() {
	this.Payload.Truncate(0)
	marshal, _ := json.Marshal(this)
	this.Payload.WriteString(string(marshal))

	if common.DEBUG_VERBOSE!=0 {
		fmt.Printf("Payload: %s\n", string(this.Payload.Bytes()))
	}
}

// ================================== PutItemReq

type PutItemReq struct {
	DynamoDBReq							`json:"-"`
}


// =================================== DeleteItemReq

type DeleteItemReq struct {
	DynamoDBReq							`json:"-"`
}


// ================================= GetItemReq

type AttributeValue struct {
	B		[]byte				`json:",omitempty"`
	BS		[][]byte			`json:",omitempty"`
	N		string				`json:",omitempty"`
	NS		[]string			`json:",omitempty"`
	S		string				`json:",omitempty"`
	SS		[]string			`json:",omitempty"`
}

// doc: http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html
type GetItemReq struct {
	DynamoDBReq							`json:"-"`

	TableName			string			`json:"TableName"`
	AttributesToGet		[]string		`json:",omitempty"`
	ConsistentRead		bool			`json:",omitempty"`
	ReturnConsumedCapacity	string		`json:",omitempty"`				// INDEXES | TOTAL | NONE
	Key					map[string]AttributeValue		`json:"Key"`
}

func (this *GetItemReq) Init() (*GetItemReq) {
	if this.DynamoDBReq.Init() == nil {
		return nil
	}

	this.Method = "POST"

	this.Headers["X-Amz-Target"] = "DynamoDB_"+DynamoDB_API_VERSION+"."+"GetItem"
	this.Headers["Content-Type"] = "application/x-amz-json-1.0"

	return this
}


func (this *GetItemReq) generatePayload() {
	this.Payload.Truncate(0)
	marshal, _ := json.Marshal(this)
	this.Payload.WriteString(string(marshal))

	if common.DEBUG_VERBOSE!=0 {
		fmt.Printf("Payload: %s\n", string(this.Payload.Bytes()))
	}
}
