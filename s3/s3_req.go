package s3

import (
	"github.com/letscool/aws/common"
)

type S3Req struct {
	common.AWSRequest
}

// doc: http://docs.aws.amazon.com/AmazonS3/latest/API/RESTServiceGET.html
type GetServiceReq struct {
	S3Req
}

func (this *GetServiceReq) Init() (*GetServiceReq) {
	if this.S3Req.Init() == nil {
		return nil
	}

	this.Method = "GET"

	return this
}




// doc: http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketOps.html


// =================================================== NewBucketReq
// ref: http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUT.html
// PUT Bucket
type NewBucketReq struct {
	S3Req
	BucketName	string
}

func (this *NewBucketReq) Init() (*NewBucketReq) {
	if this.S3Req.Init() == nil {
		return nil
	}

	this.Method = "PUT"

	return this
}



// HEAD Bucket
type ExistBucketReq struct {
	S3Req
	BucketName	string
}

func (this *ExistBucketReq) Init() (*ExistBucketReq) {
	if this.S3Req.Init() == nil {
		return nil
	}

	this.Method = "HEAD"

	return this
}


// DELETE Bucket
type DeleteBucketReq struct {
	S3Req
	BucketName	string
}

func (this *DeleteBucketReq) Init() (*DeleteBucketReq) {
	if this.S3Req.Init() == nil {
		return nil
	}

	this.Method = "DELETE"

	return this
}


