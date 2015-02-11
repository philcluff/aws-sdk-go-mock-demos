# AWS SDK Go Mock Demos

A demonstration of approaches to mocking the behaviour of the new aws-sdk-go.

So far this approach demonstrates a usage of:

* Testify: https://github.com/stretchr/testify/
* Mockery: https://github.com/vektra/mockery

The test case is mocking out a DynamoDB query operation.

My intention is to expand this repository to also show gomock usage, as well as the suggested Amazon approach, using handlers.

## Running

Dependencies are vendor'd. As such, you should use godep:

    godep go test

## Regerating the mocks

You'll need mockery insatlled and on your path:

    mockery -name DynamoDBer
