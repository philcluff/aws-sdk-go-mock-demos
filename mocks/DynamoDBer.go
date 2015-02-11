package mocks

import "github.com/stretchr/testify/mock"

import "github.com/awslabs/aws-sdk-go/service/dynamodb"

type DynamoDBer struct {
	mock.Mock
}

func (m *DynamoDBer) Query(_a0 *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	ret := m.Called(_a0)

	r0 := ret.Get(0).(*dynamodb.QueryOutput)
	r1 := ret.Error(1)

	return r0, r1
}
