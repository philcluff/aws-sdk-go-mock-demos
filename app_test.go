package app

import (
	"fmt"
	"github.com/GeneticGenesis/aws-sdk-go-mock-demos/mocks"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AppSuite struct {
	suite.Suite
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(AppSuite))
}

func (s *AppSuite) SetupTest() {
	// Test Setup
}

func (s *AppSuite) SetupSuite() {
	// Suite Setup
}

func (s *AppSuite) TestHitDynamo() {
	client := new(mocks.DynamoDBer)
	client.On("Query", mock.AnythingOfType("*dynamodb.QueryInput")).Return(&dynamodb.QueryOutput{Count: aws.Integer(0)}, nil)

	helper := &DynamoHelper{
		client,
	}
	output, err := helper.HitDynamo("1")
	fmt.Printf("Found: %v entries\n", *output.Count)
	fmt.Println(output, err)
}
