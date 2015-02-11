# AWS SDK Go Mock Demos

A demonstration of approaches to mocking the behaviour of the new aws-sdk-go.

So far this approach demonstrates a usage of:

* Testify: https://github.com/stretchr/testify/
* Mockery: https://github.com/vektra/mockery

My intention is to expand it to also show gomock usage.

## Running

Dependencies are vendor'd. As such, you should use godep:

    godep go test
