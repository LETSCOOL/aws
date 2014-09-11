package dynamodb

import (
	"github.com/letscool/aws/common"
	"fmt"
	"encoding/json"
)

type DynamoDBReq struct {
	common.AWSRequest

}

// ==========================  ListTablesReq
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


// ======================================= DescribeTableReq
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

type AttributeValue struct {
	B		[]byte				`json:",omitempty"`
	BS		[][]byte			`json:",omitempty"`
	N		string				`json:",omitempty"`
	NS		[]string			`json:",omitempty"`
	S		string				`json:",omitempty"`
	SS		[]string			`json:",omitempty"`
}

// ref: http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ExpectedAttributeValue.html
type ExpectedAttributeValue struct {
	AttributeValueList		[]AttributeValue		`json:",omitempty"`
	ComparisonOperator		string					// EQ | NE | LE | LT | GE | GT | NOT_NULL | NULL | CONTAINS | NOT_CONTAINS | BEGINS_WITH | IN | BETWEEN
	// We don't support following deprecated variables, it is more complex to work together.
	//Exists				bool					`json:",omitempty"`	// deprecated, use AttributeValueList and ComparisonOperator instead
	//Value					AttributeValue			`json:",omitempty"`	// deprecated, use AttributeValueList and ComparisonOperator instead
}

// ref: http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html#DDB-PutItem-request-Expected
type PutItemReq struct {
	DynamoDBReq							`json:"-"`

	TableName				string
	Attributes				map[string]AttributeValue			`json:"Item"`
	ConditionalOperator		string								`json:",omitempty"`				// "AND", "OR"
	Expected				map[string]ExpectedAttributeValue	`json:",omitempty"`
	ReturnConsumedCapacity	string								`json:",omitempty"`				// INDEXES | TOTAL | NONE
	ReturnItemCollectionMetrics string							`json:",omitempty"`				// "SIZE", "NONE"
	ReturnValues			string								`json:",omitempty"`				// NONE | ALL_OLD | UPDATED_OLD | ALL_NEW | UPDATED_NEW
}

func (this *PutItemReq) Init() (*PutItemReq) {
	if this.DynamoDBReq.Init() == nil {
		return nil
	}

	this.Method = "POST"

	this.Headers["X-Amz-Target"] = "DynamoDB_"+DynamoDB_API_VERSION+"."+"PutItem"
	this.Headers["Content-Type"] = "application/x-amz-json-1.0"

	return this
}


func (this *PutItemReq) generatePayload() {
	this.Payload.Truncate(0)
	marshal, _ := json.Marshal(this)
	this.Payload.WriteString(string(marshal))

	if common.DEBUG_VERBOSE!=0 {
		fmt.Printf("Payload: %s\n", string(this.Payload.Bytes()))
	}
}


//==========================  UpdateItemReq

type AttributeValueUpdate struct {
	Action				string					// ADD | PUT | DELETE
	Value				AttributeValue
}

type UpdateItemReq struct {
	DynamoDBReq							`json:"-"`

	TableName				string
	Key						map[string]AttributeValue
	AttributeUpdates		map[string]AttributeValueUpdate

	ConditionalOperator		string								`json:",omitempty"`				// "AND", "OR"
	Expected				map[string]ExpectedAttributeValue	`json:",omitempty"`
	ReturnConsumedCapacity	string								`json:",omitempty"`				// INDEXES | TOTAL | NONE
	ReturnItemCollectionMetrics string							`json:",omitempty"`				// "SIZE", "NONE"
	ReturnValues			string								`json:",omitempty"`				// NONE | ALL_OLD | UPDATED_OLD | ALL_NEW | UPDATED_NEW

}

func (this *UpdateItemReq) Init() (*UpdateItemReq) {
	if this.DynamoDBReq.Init() == nil {
		return nil
	}

	this.Method = "POST"

	this.Headers["X-Amz-Target"] = "DynamoDB_"+DynamoDB_API_VERSION+"."+"UpdateItem"
	this.Headers["Content-Type"] = "application/x-amz-json-1.0"

	return this
}


func (this *UpdateItemReq) generatePayload() {
	this.Payload.Truncate(0)
	marshal, _ := json.Marshal(this)
	this.Payload.WriteString(string(marshal))

	if common.DEBUG_VERBOSE!=0 {
		fmt.Printf("Payload: %s\n", string(this.Payload.Bytes()))
	}
}

// =================================== DeleteItemReq

type DeleteItemReq struct {
	DynamoDBReq							`json:"-"`

	TableName				string
	Key						map[string]AttributeValue
	ConditionalOperator		string								`json:",omitempty"`				// "AND", "OR"
	Expected				map[string]ExpectedAttributeValue	`json:",omitempty"`
	ReturnConsumedCapacity	string								`json:",omitempty"`				// INDEXES | TOTAL | NONE
	ReturnItemCollectionMetrics string							`json:",omitempty"`				// "SIZE", "NONE"
	ReturnValues			string								`json:",omitempty"`				// NONE | ALL_OLD | UPDATED_OLD | ALL_NEW | UPDATED_NEW
}

func (this *DeleteItemReq) Init() (*DeleteItemReq) {
	if this.DynamoDBReq.Init() == nil {
		return nil
	}

	this.Method = "POST"

	this.Headers["X-Amz-Target"] = "DynamoDB_"+DynamoDB_API_VERSION+"."+"DeleteItem"
	this.Headers["Content-Type"] = "application/x-amz-json-1.0"

	return this
}


func (this *DeleteItemReq) generatePayload() {
	this.Payload.Truncate(0)
	marshal, _ := json.Marshal(this)
	this.Payload.WriteString(string(marshal))

	if common.DEBUG_VERBOSE!=0 {
		fmt.Printf("Payload: %s\n", string(this.Payload.Bytes()))
	}
}

// ================================= GetItemReq

// doc: http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html
type GetItemReq struct {
	DynamoDBReq							`json:"-"`

	TableName			string							`json:"TableName"`
	Key					map[string]AttributeValue		`json:"Key"`
	AttributesToGet		[]string		`json:",omitempty"`
	ConsistentRead		bool			`json:",omitempty"`
	ReturnConsumedCapacity	string		`json:",omitempty"`				// INDEXES | TOTAL | NONE
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


// ================================ ScanReq

type Condition struct {
	ComparisonOperator	string			// EQ | NE | LE | LT | GE | GT | NOT_NULL | NULL | CONTAINS | NOT_CONTAINS | BEGINS_WITH | IN | BETWEEN
	AttributeValueList	[]AttributeValue
}

type ScanReq struct {
	DynamoDBReq									`json:"-"`

	TableName			string					`json:"TableName"`
	Select				string					`json:",omitempty"`				// ALL_ATTRIBUTES | ALL_PROJECTED_ATTRIBUTES | SPECIFIC_ATTRIBUTES | COUNT
	AttributesToGet		[]string				`json:",omitempty"`
	ExclusiveStartKey	map[string]AttributeValue	`json:",omitempty"`
	Limit				int						`json:",omitempty"`

	ConditionalOperator	string					`json:",omitempty"`				// "AND", "OR"
	ScanFilter			map[string]Condition	`json:",omitempty"`

	TotalSegments		int						`json:",omitempty"`
	Segment				int						`json:",omitempty"`

	ReturnConsumedCapacity	string				`json:",omitempty"`				// INDEXES | TOTAL | NONE
}


func (this *ScanReq) Init() (*ScanReq) {
	if this.DynamoDBReq.Init() == nil {
		return nil
	}

	this.Method = "POST"

	this.Headers["X-Amz-Target"] = "DynamoDB_"+DynamoDB_API_VERSION+"."+"Scan"
	this.Headers["Content-Type"] = "application/x-amz-json-1.0"

	return this
}


func (this *ScanReq) generatePayload() {
	this.Payload.Truncate(0)
	marshal, _ := json.Marshal(this)
	this.Payload.WriteString(string(marshal))

	if common.DEBUG_VERBOSE!=0 {
		fmt.Printf("Payload: %s\n", string(this.Payload.Bytes()))
	}
}


//==================================== QueryReq
// ref: http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Query.html
type QueryReq struct {
	DynamoDBReq

	TableName			string						`json:"TableName"`
	KeyConditions		map[string]Condition
	Select				string						`json:",omitempty"`				// ALL_ATTRIBUTES | ALL_PROJECTED_ATTRIBUTES | SPECIFIC_ATTRIBUTES | COUNT
	AttributesToGet		[]string					`json:",omitempty"`
	IndexName			string						`json:",omitempty"`
	ConsistentRead		bool
	ScanIndexForward	bool

	ConditionalOperator	string						`json:",omitempty"`				// "AND", "OR"  // A logical operator to apply to the conditions in the QueryFilter map
	QueryFilter			map[string]Condition		`json:",omitempty"`

	Limit				int							`json:",omitempty"`
	ExclusiveStartKey	map[string]AttributeValue	`json:",omitempty"`


	ReturnConsumedCapacity	string					`json:",omitempty"`				// INDEXES | TOTAL | NONE
}

func (this *QueryReq) Init() (*QueryReq) {
	if this.DynamoDBReq.Init() == nil {
		return nil
	}

	this.Method = "POST"

	this.Headers["X-Amz-Target"] = "DynamoDB_"+DynamoDB_API_VERSION+"."+"Query"
	this.Headers["Content-Type"] = "application/x-amz-json-1.0"

	this.ScanIndexForward = true				// force to set as default situation in aws doc

	return this
}


func (this *QueryReq) generatePayload() {
	this.Payload.Truncate(0)
	marshal, _ := json.Marshal(this)
	this.Payload.WriteString(string(marshal))

	if common.DEBUG_VERBOSE!=0 {
		fmt.Printf("Payload: %s\n", string(this.Payload.Bytes()))
	}
}


