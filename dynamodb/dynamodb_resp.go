package dynamodb

import (
	"net/http"
	"io/ioutil"
	"github.com/letscool/aws/common"
	"encoding/json"
	"strings"
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

type TableDescription struct {
	TableName		string
	TableSizeBytes	int64
	TableStatus		string			// "CREATING", "UPDATING", "DELETING", "ACTIVE"
	ItemCount		int64
	CreationDateTime	float64
	//AttributeDefinitions
	//KeySchema
	//LocalSecondaryIndexes
	//ProvisionedThroughput
	//GlobalSecondaryIndexes
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
