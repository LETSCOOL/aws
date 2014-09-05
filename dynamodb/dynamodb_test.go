package dynamodb

import (
	"testing"
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


func TestDynamoDBDescribeTable(t *testing.T) {
	ddb, err := NewDynamoDB(credentials)
	if err!=nil {
		t.Errorf("NewDynamoDB fail (%s)\n", err)
		return
	}

	//ddb.Endpoint = "dynamodb.us-west-2.amazonaws.com"
	//ddb.Region = "us-west-2"

	req := new(DescribeTableReq).Init()
	req.TableName = "Abc"
	if resp, err := ddb.DescribeTable(req); err!=nil {
		t.Errorf("DynamoDB ListTables fail (%s)\n", err)
	} else {
		if resp.Error!=nil {
			t.Errorf("DynamoDB ListTables fail (%s, %s)\n", resp.Error.Exception, resp.Error.Message)
		} else {
			t.Logf("Describe table: %s(%s) \n", resp.TableDescription.TableName, resp.TableDescription.TableStatus)
		}
	}
}
