package common

import (
	"bytes"
)

type AWSRequest struct {
	Method        string // "GET", "POST", "DELETE", etc.
	Path          string // absolute path, starts with '/'
	Headers        map[string]string
	Parameters     map[string]string
	Payload        bytes.Buffer
}

func (this *AWSRequest) Init() *AWSRequest {
	this.Headers = make(map[string]string)
	this.Parameters = make(map[string]string)
	this.Path = "/"

	return this
}


func (this *AWSRequest) SetParameter(key string, value string) {
	this.Parameters[key] = value
}

func (this *AWSRequest) DeleteParameter(key string) {
	delete(this.Parameters, key)
}


func (this *AWSRequest) Parameter(key string) string {
	value, ok := this.Parameters[key]
	if ok {
		return value
	}
	return ""
}

/*func (this *AWSRequest) GeneratePayload() {
	// empty
}*/
















