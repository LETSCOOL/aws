package s3

import (
	"testing"
	"github.com/letscool/aws/common"
)

var credentials        *common.Credentials

func TestS3Connect(t *testing.T) {
	cred, err := common.NewCredentialsFromCSV(common.TEST_CREDENTIALS_FILE)
	if err != nil {
		t.Errorf("Create credentials fail (%s)\n", err)
		return
	}

	s3, err := NewS3(cred)
	if err!=nil {
		t.Errorf("NewS3 fail (%s)\n", err)
		return
	}

	if resp, err := s3.GetService(new(GetServiceReq).Init()); err != nil {
		t.Errorf("S3 GetService fail (%s)\n", err)
		return
	} else {
		t.Logf("Owner: %s, Buckets: %s\n", resp.Owner, resp.Buckets)
	}

	credentials = cred

}

func TestS3NewBucket(t *testing.T) {

	s3, err := NewS3(credentials)
	if err!=nil {
		t.Errorf("NewS3 fail (%s)\n", err)
		return
	}

	nb := new(NewBucketReq).Init()
	nb.BucketName = "SS1234567"
	if _, err := s3.NewBucket(nb); err!=nil {
		t.Errorf("S3 NewBucket fail (%s)\n", err)
		return
	} else {
		t.Logf("NewBucket done\n")
	}

}

func TestS3ExistBucket(t *testing.T) {
	s3, err := NewS3(credentials)
	if err!=nil {
		t.Errorf("NewS3 fail (%s)\n", err)
		return
	}

	req := new(ExistBucketReq).Init()
	req.BucketName = "SS1234567"
	if resp, err := s3.ExistBucket(req); err!=nil {
		t.Errorf("S3 ExistBucket fail (%s)\n", err)
		return
	} else {
		if resp.Exists {
			t.Logf("Bucket Exists\n")
		} else {
			t.Errorf("S3 Bucket NOT Exist\n")
		}
	}
}

func TestS3DeleteBucket(t *testing.T) {
	s3, err := NewS3(credentials)
	if err!=nil {
		t.Errorf("NewS3 fail (%s)\n", err)
		return
	}

	req := new(DeleteBucketReq).Init()
	req.BucketName = "SS1234567"
	if resp, err := s3.DeleteBucket(req); err!=nil {
		t.Errorf("S3 DeleteBucket fail (%s)\n", err)
		return
	} else {
		if resp.Error==nil {
			t.Logf("Bucket Delete\n")
		} else {
			t.Errorf("S3 Bucket NOT Delete (%s)\n", resp.Error.Message)
		}
	}
}


