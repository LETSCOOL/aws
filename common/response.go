package common

import (
	"net/http"
)


type AWSResponse struct {
	//Request *AWSRequest
}

func (this *AWSResponse) Init(req *AWSRequest, resp *http.Response) (*AWSResponse, error) {
	//this.Request = req

	return this, nil
}











