package dynamodb

import (
	"net/http"
	"io/ioutil"
	"github.com/letscool/aws/common"
	"encoding/json"
	"strings"
	"time"
)


/**
	DynamoDB error handle: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ErrorHandling.html
 */
type Error struct {
	Type		string			`json:"__type"`
	Message		string			`json`
	Exception	string			`json:"-"`			// parse from type
}

type DynamoDBResp struct {
	common.AWSResponse
	Error	*Error			// response error code from server
}

func (this *DynamoDBResp) Init(req *DynamoDBReq, resp *http.Response) (*DynamoDBResp, error) {
	if _, err := this.AWSResponse.Init(&req.AWSRequest, resp); err != nil {
		return nil, err
	}

	this.Error = nil

	if resp.StatusCode >= 300 {
		this.Error = new(Error)

		if strings.ToLower(resp.Header.Get("Content-Type"))== "application/x-amz-json-1.0" {
			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(body, this.Error)
			if err!=nil {
				return nil, err
			}

			this.Error.Exception = this.Error.Type[strings.IndexRune(this.Error.Type, rune('#'))+1:]

		} else {
			// impossible, and no ideal how to deal.
		}

	}

	return this, nil
}

// ================================= ListTablesResp
type ListTablesResp struct {
	DynamoDBResp			`json:"-"`
	TableNames				[]string
	LastEvaluatedTableName	string
}

func (this *ListTablesResp) Init(req *ListTablesReq, resp *http.Response) (*ListTablesResp, error) {
	if _, err := this.DynamoDBResp.Init(&req.DynamoDBReq, resp); err != nil {
		return nil, err
	}

	if this.Error!=nil {
		return this, nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, this); err!=nil {
		return nil, err
	}

	return this, nil
}

// ================================= DescribeTableResp
// ref: http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TableDescription.html
type TableDescription struct {
	TableName				string			// 3 <= Len(TableName) <= 255
	TableSizeBytes			int64
	TableStatus				string			// "CREATING", "UPDATING", "DELETING", "ACTIVE"
	ItemCount				int64
	CreationDateTime		float64
	AttributeDefinitions	[]AttributeDefinition
	KeySchema				[]KeySchemaElement
	LocalSecondaryIndexes	[]LocalSecondaryIndexDescription
	ProvisionedThroughput	ProvisionedThroughputDescription
	GlobalSecondaryIndexes	[]GlobalSecondaryIndexDescription
}

func (this *TableDescription) CreationTime() time.Time {
	return UnixDateTimeToTime(this.CreationDateTime)
}

/// Unix date time to systme time
func UnixDateTimeToTime(dateTime float64) time.Time {
	return time.Unix(int64(dateTime), int64((dateTime-float64(int64(dateTime)))*float64(time.Second)))
}

// ref: http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GlobalSecondaryIndexDescription.html
type GlobalSecondaryIndexDescription struct {
	IndexName				string						// 3 <= Len(IndexName) <= 255
	IndexSizeBytes			int64
	IndexStatus				string						// "CREATING", "UPDATING", "DELETING", "ACTIVE"
	ItemCount				int64
	KeySchema				[]KeySchemaElement
	Projection				Projection
	ProvisionedThroughput	ProvisionedThroughputDescription
}

// ref: http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughputDescription.html
type ProvisionedThroughputDescription struct {
	LastDecreaseDateTime		float64
	LastIncreaseDateTime		float64
	NumberOfDecreasesToday		int64
	ReadCapacityUnits			int64
	WriteCapacityUnits			int64
}

func (this *ProvisionedThroughputDescription) LastDecreaseTime() time.Time {
	return UnixDateTimeToTime(this.LastDecreaseDateTime)
}

func (this *ProvisionedThroughputDescription) LastIncreaseTime() time.Time {
	return UnixDateTimeToTime(this.LastIncreaseDateTime)
}

// http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_LocalSecondaryIndexDescription.html
type LocalSecondaryIndexDescription struct {
	IndexName				string			// 3 <= Len(IndexName) <= 255
	IndexSizeBytes			int64
	ItemCount				int64
	KeySchema				[]KeySchemaElement
	Projection				Projection
}

type DescribeTableResp struct {
	DynamoDBResp			`json:"-"`
	TableDescription	TableDescription		`json:"Table"`
}


func (this *DescribeTableResp) Init(req *DescribeTableReq, resp *http.Response) (*DescribeTableResp, error) {
	if _, err := this.DynamoDBResp.Init(&req.DynamoDBReq, resp); err != nil {
		return nil, err
	}

	if this.Error!=nil {
		return this, nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, this); err!=nil {
		return nil, err
	}

	return this, nil
}

// ==================================== CreateTableResp
type CreateTableResp struct {
	DynamoDBResp			`json:"-"`
	TableDescription		TableDescription
}

func (this *CreateTableResp) Init(req *CreateTableReq, resp *http.Response) (*CreateTableResp, error) {
	if _, err := this.DynamoDBResp.Init(&req.DynamoDBReq, resp); err != nil {
		return nil, err
	}

	if this.Error!=nil {
		return this, nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, this); err!=nil {
		return nil, err
	}

	return this, nil
}

// ==================================== DeleteTableResp
type DeleteTableResp struct {
	DynamoDBResp			`json:"-"`
	TableDescription		TableDescription
}

func (this *DeleteTableResp) Init(req *DeleteTableReq, resp *http.Response) (*DeleteTableResp, error) {
	if _, err := this.DynamoDBResp.Init(&req.DynamoDBReq, resp); err != nil {
		return nil, err
	}

	if this.Error!=nil {
		return this, nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, this); err!=nil {
		return nil, err
	}

	return this, nil
}

// ==================================== PutItemResp

// ref: http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ConsumedCapacity.html
type ConsumedCapacity struct {
	CapacityUnits		float64
	GlobalSecondaryIndexes	map[string]Capacity
	LocalSecondaryIndexes	map[string]Capacity
	Table				Capacity
	TableName			string
}

type Capacity struct {
	CapacityUnits		float64
}

type ItemCollectionMetrics struct {
	ItemCollectionKey		map[string]AttributeValue
	SizeEstimateRangeGB		[]float64
}

type PutItemResp struct {
	DynamoDBResp			`json:"-"`

	Attributes				map[string]AttributeValue
	ConsumedCapacity		ConsumedCapacity
	ItemCollectionMetrics	ItemCollectionMetrics
}

func (this *PutItemResp) Init(req *PutItemReq, resp *http.Response) (*PutItemResp, error) {
	if _, err := this.DynamoDBResp.Init(&req.DynamoDBReq, resp); err != nil {
		return nil, err
	}

	if this.Error!=nil {
		return this, nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, this); err!=nil {
		return nil, err
	}

	return this, nil
}

// ==================================== DeleteItemResp
type DeleteItemResp struct {
	DynamoDBResp			`json:"-"`

	Attributes				map[string]AttributeValue
	ConsumedCapacity		ConsumedCapacity
	ItemCollectionMetrics	ItemCollectionMetrics
}

func (this *DeleteItemResp) Init(req *DeleteItemReq, resp *http.Response) (*DeleteItemResp, error) {
	if _, err := this.DynamoDBResp.Init(&req.DynamoDBReq, resp); err != nil {
		return nil, err
	}

	if this.Error!=nil {
		return this, nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, this); err!=nil {
		return nil, err
	}

	return this, nil
}

// ==================================== GetItemResp
type GetItemResp struct {
	DynamoDBResp			`json:"-"`

	Attributes				map[string]AttributeValue			`json:"Item"`
	ConsumedCapacity		ConsumedCapacity
}

func (this *GetItemResp) Init(req *GetItemReq, resp *http.Response) (*GetItemResp, error) {
	if _, err := this.DynamoDBResp.Init(&req.DynamoDBReq, resp); err != nil {
		return nil, err
	}

	if this.Error!=nil {
		return this, nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, this); err!=nil {
		return nil, err
	}

	return this, nil
}

// ================================== ScanResp
type ScanResp struct {
	DynamoDBResp			`json:"-"`
	Count					int
	ScannedCount			int
	Items					[]map[string]AttributeValue
	LastEvaluatedKey		map[string]AttributeValue

	ConsumedCapacity		ConsumedCapacity
}

func (this *ScanResp) Init(req *ScanReq, resp *http.Response) (*ScanResp, error) {
	if _, err := this.DynamoDBResp.Init(&req.DynamoDBReq, resp); err != nil {
		return nil, err
	}

	if this.Error!=nil {
		return this, nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, this); err!=nil {
		return nil, err
	}

	return this, nil
}


