package s3

import (
	"github.com/letscool/aws/common"
	"net/http"
)

const (
	S3_API_VERSION = "2006-03-01"
)

/**
	S3 doesn't support virtual hosted style, use path style always.
	Ref: http://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region
 */
type S3 struct {
	common.AWSService
}

func NewS3(cred *common.Credentials) (*S3, error) {
	return new(S3).Init(cred)
}

func (this *S3) Init(cred *common.Credentials) (*S3, error) {
	if _, err := this.AWSService.Init(cred); err!=nil {
		return nil, err
	}

	this.Region = "us-east-1"
	this.Service = "s3"
	this.Endpoint = "s3.amazonaws.com"

	return this, nil
}


func (this *S3) GetService(req *GetServiceReq) (*GetServiceResp, error) {
	httpReq, err := this.Sign4(&req.AWSRequest, true)
	if err != nil {
		return nil, err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return new(GetServiceResp).Init(req, httpResp)
}

func (this *S3) NewBucket(req *NewBucketReq) (*NewBucketResp, error) {
	req.Path = "/" + req.BucketName

	httpReq, err := this.Sign4(&req.AWSRequest, true)
	if err != nil {
		return nil, err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return new(NewBucketResp).Init(req, httpResp)
}

func (this *S3) ExistBucket(req *ExistBucketReq) (*ExistBucketResp, error) {
	req.Path = "/" + req.BucketName

	httpReq, err := this.Sign4(&req.AWSRequest, true)
	if err != nil {
		return nil, err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return new(ExistBucketResp).Init(req, httpResp)
}

func (this *S3) DeleteBucket(req *DeleteBucketReq) (*DeleteBucketResp, error) {
	req.Path = "/" + req.BucketName

	httpReq, err := this.Sign4(&req.AWSRequest, true)
	if err != nil {
		return nil, err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return new(DeleteBucketResp).Init(req, httpResp)
}


