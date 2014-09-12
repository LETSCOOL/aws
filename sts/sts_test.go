package sts

import (
	"testing"
	"github.com/letscool/aws/common"
)

var credentials        *common.Credentials

func TestSTSConnect(t *testing.T) {
	cred, err := common.NewCredentialsFromCSV(common.TEST_CREDENTIALS_FILE)
	if err != nil {
		t.Errorf("Create credentials fail (%s)\n", err)
		return
	}

	_, err = NewSTS(cred)
	if err != nil {
		t.Errorf("NewS3 fail (%s)\n", err)
		return
	}

	credentials = cred

}


func TestSTSGetFederationToken(t *testing.T) {

	sts, err := NewSTS(credentials)
	if err != nil {
		t.Errorf("NewS3 fail (%s)\n", err)
		return
	}

	req := new(GetFederationTokenReq).Init()
	req.Name = "ABCDEFG"
	//req.DurationSeconds = 100
	/*req.Policy = `{"Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "sts:GetFederationToken",
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": "dynamodb:*",
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": "sqs:*",
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": "s3:*",
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": "sns:*",
      "Resource": "*"
    }
  ]
}`*/

	if resp, err := sts.GetFederationToken(req); err != nil {
		t.Errorf("STS GetFederationToken fail (%s)\n", err)
	} else {
		if resp.ErrorResponse != nil {
			t.Errorf("STS GetFederationToken fail (%s)\n", resp.ErrorResponse.Err.Message)
		} else {
			t.Logf("GetFederationToken done (%s)\n", resp.GetFederationTokenResult.FederatedUser.FederatedUserId)
		}
	}

}
