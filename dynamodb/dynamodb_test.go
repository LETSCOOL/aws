package dynamodb

import (
	"testing"
	"time"
	"github.com/letscool/aws/common"
)

var credentials        *common.Credentials

func TestDynamoDBConnect(t *testing.T) {
	cred, err := common.NewCredentialsFromCSV(common.TEST_CREDENTIALS_FILE)
	if err != nil {
		t.Errorf("Create credentials fail (%s)\n", err)
		return
	}

	ddb, err := NewDynamoDB(cred)
	if err!=nil {
		t.Errorf("NewDynamoDB fail (%s)\n", err)
		return
	}

	//ddb.Endpoint = "dynamodb.us-west-2.amazonaws.com"
	//ddb.Region = "us-west-2"

	if resp, err := ddb.ListTables(new(ListTablesReq).Init()); err != nil {
		t.Errorf("DynamoDB ListTables fail (%s)\n", err)
		return
	} else {
		t.Log(resp.TableNames)
	}


	credentials = cred

}

func TestDynamoDBCreateTable(t *testing.T) {
	ddb, err := NewDynamoDB(credentials)
	if err!=nil {
		t.Errorf("NewDynamoDB fail (%s)\n", err)
		return
	}

	//ddb.Endpoint = "dynamodb.us-west-2.amazonaws.com"
	//ddb.Region = "us-west-2"

	req := (&CreateTableReq{
		TableName:"Abcdefg__99999",
		KeySchema: []KeySchemaElement{KeySchemaElement{"HASH_KEY", "HASH"}, KeySchemaElement{"RANGE_KEY", "RANGE"}},
		AttributeDefinitions: []AttributeDefinition{AttributeDefinition{"HASH_KEY", "S"}, AttributeDefinition{"RANGE_KEY", "N"}, AttributeDefinition{"ATTR_1", "B"}, AttributeDefinition{"ATTR_2", "S"}},
		GlobalSecondaryIndexes: []GlobalSecondaryIndex {
			GlobalSecondaryIndex{
				IndexName:"Global_attr_1_2",
				KeySchema:[]KeySchemaElement{KeySchemaElement{"ATTR_1", "HASH"}, KeySchemaElement{"ATTR_2", "RANGE"}},
				Projection:Projection{ProjectionType:"ALL"},
				ProvisionedThroughput: ProvisionedThroughput{2, 2},
			},
		},
		LocalSecondaryIndexes: []LocalSecondaryIndex {
			LocalSecondaryIndex{
				IndexName:"Local_attr_1",
				KeySchema:[]KeySchemaElement{KeySchemaElement{"HASH_KEY", "HASH"}, KeySchemaElement{"ATTR_1", "RANGE"}},
				Projection:Projection{ProjectionType:"KEYS_ONLY"},
			},
		},
		ProvisionedThroughput: ProvisionedThroughput{2, 2},
	}).Init()
	if resp, err := ddb.CreateTable(req); err!=nil {
		t.Errorf("DynamoDB CreateTable fail (%s)\n", err)
	} else {
		if resp.Error!=nil {
			t.Errorf("DynamoDB CreateTable fail (%s, %s)\n", resp.Error.Exception, resp.Error.Message)
		} else {
			t.Logf("Create table table: %s(%s, %s) \n%s\n", resp.TableDescription.TableName, resp.TableDescription.TableStatus, resp.TableDescription.CreationTime(), resp.TableDescription)
		}
	}
}

func TestDynamoDBDescribeTable(t *testing.T) {
	ddb, err := NewDynamoDB(credentials)
	if err!=nil {
		t.Errorf("NewDynamoDB fail (%s)\n", err)
		return
	}

	//ddb.Endpoint = "dynamodb.us-west-2.amazonaws.com"
	//ddb.Region = "us-west-2"

	req := new(DescribeTableReq).Init()
	req.TableName = "Abcdefg__99999"
	Repeat_again:
	if resp, err := ddb.DescribeTable(req); err!=nil {
		t.Errorf("DynamoDB DescribeTable fail (%s)\n", err)
	} else {
		if resp.Error!=nil {
			t.Errorf("DynamoDB DescribeTable fail (%s, %s)\n", resp.Error.Exception, resp.Error.Message)
		} else {
			t.Logf("Describe table: %s(%s, %s) \n%s\n", resp.TableDescription.TableName, resp.TableDescription.TableStatus, resp.TableDescription.CreationTime(), resp.TableDescription)
			switch resp.TableDescription.TableStatus {
			case "CREATING", "UPDATING":
				t.Log("Waiting 'ACTIVE' state\n")
				time.Sleep(time.Second*2)
				goto Repeat_again
			case "DELETING":
				// strange, but it will fail in next test unit
			case "ACTIVE":
				// do nothing

			}
		}
	}
}


func TestDynamoDBDeleteTable(t *testing.T) {
	ddb, err := NewDynamoDB(credentials)
	if err!=nil {
		t.Errorf("NewDynamoDB fail (%s)\n", err)
		return
	}

	//ddb.Endpoint = "dynamodb.us-west-2.amazonaws.com"
	//ddb.Region = "us-west-2"

	req := new(DeleteTableReq).Init()
	req.TableName = "Abcdefg__99999"
	if resp, err := ddb.DeleteTable(req); err!=nil {
		t.Errorf("DynamoDB DeleteTable fail (%s)\n", err)
	} else {
		if resp.Error!=nil {
			t.Errorf("DynamoDB DeleteTable fail (%s, %s)\n", resp.Error.Exception, resp.Error.Message)
		} else {
			t.Logf("Delete table table: %s(%s, %s) \n%s\n", resp.TableDescription.TableName, resp.TableDescription.TableStatus, resp.TableDescription.CreationTime(), resp.TableDescription)
		}
	}
}

