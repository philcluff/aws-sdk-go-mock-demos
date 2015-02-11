package app

import (
	"fmt"
	"github.com/GeneticGenesis/aws-sdk-go-mock-demos/mocks"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
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

	m := make(map[string]dynamodb.Condition)
	var alist []dynamodb.AttributeValue

	alist = append(alist, dynamodb.AttributeValue{
		S: aws.String("1"),
	})

	m["Id"] = dynamodb.Condition{
		AttributeValueList: alist,
		ComparisonOperator: aws.String(dynamodb.ComparisonOperatorEq),
	}

	foo := &dynamodb.QueryInput{
		KeyConditions: m,
		TableName:     aws.String("foobar2"),
	}

	client := new(mocks.DynamoDBer)
	client.On("Query", foo).Return(&dynamodb.QueryOutput{Count: aws.Integer(0)}, nil)

	helper := &DynamoHelper{
		client,
	}
	output, err := helper.HitDynamo("1")
	fmt.Printf("Found: %v entries\n", *output.Count)
	fmt.Println(output, err)
}
