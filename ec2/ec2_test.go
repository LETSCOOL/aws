package ec2

import (
	"testing"
	"github.com/letscool/aws/common"
)

var credentials        *common.Credentials

func TestEC2Connect(t *testing.T) {
	cred, err := common.NewCredentialsFromCSV(common.TEST_CREDENTIALS_FILE)
	if err != nil {
		t.Errorf("Create credentials fail (%s)\n", err)
		return
	}

	ec2, err := NewEC2(cred)
	if err!=nil {
		t.Errorf("NewEC2 fail (%s)\n", err)
		return
	}

	if resp, err := ec2.DescribeInstances(new(DescribeInstancesReq).Init()); err != nil {
		t.Errorf("EC2 DescribeInstances fail (%s)\n", err)
		return
	} else {
		if resp.Errors!=nil {
			t.Errorf("EC2 DescribeInstances fail: %s\n", resp.Errors.Error[0].Message)

		} else {
			t.Logf("RequestId: %s, NextToken: %s\n", resp.RequestId, resp.NextToken)
		}
	}

	credentials = cred
}












