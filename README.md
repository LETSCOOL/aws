When I want to use Amazon's AWS in Go environment, I found no full library in Amazon's office site.
And I found two projects in github.com, but these projects focus on AWS Signature. 
If you need only AWS Signature API, you can refer these two projects.

https://github.com/smartystreets/go-aws-auth

https://github.com/bmizerany/aws4

Because of these reasons, I try to write a AWS library for Go language.

INSTALLATION
------------

You should set global parameter 'GOPATH' first, than download this package

`go get github.com/letscool/aws`

You can install each package separately, go tool will also generate package file(s) (.a or .dll) for you. Note: Package `common` always need to be installed first.

`go get github.com/letscool/aws/common`

`go get github.com/letscool/aws/dynamodb`

`go get github.com/letscool/aws/ec2`

`go get github.com/letscool/aws/s3`

`go get github.com/letscool/aws/sts`

RUN TEST
--------

There are some test units in each packages, you can run them as following command.
Set the global parameter 'TEST_CREDENTIALS_FILE' to a credentials.csv file location.
The csv file was downloaded when you created an IAM user in Amazon AWS site.

`export TEST_CREDENTIALS_FILE=/your/folder/credentials.csv`

`go test github.com/letscool/aws/common`

`go test github.com/letscool/aws/dynamodb`

`go test github.com/letscool/aws/ec2`

`go test github.com/letscool/aws/s3`

`go test github.com/letscool/aws/sts`

Use `go test -v github.com/....` to see more information.

USAGE
-----

`import "github.com/letscool/aws/ec2"`

`import "github.com/letscool/aws/dynamodb"`

`import "github.com/letscool/aws/s3"`

`import "github.com/letscool/aws/sts"`

SAMPLES
-------

You can refer to all unit test files (filename ends with '_test').


API
---

Only support some simple API(s):

* DynamoDB: ListTables, CreateTable, DescribeTable, DeleteTable, PutItem, GetItem, DeleteItem

* EC2: DescribeInstances

* S3: GetService, NewBucket, ExistBucket, DeleteBucket

* STS: GetFederationToken


