package app

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
)

type DynamoHelper struct {
	client DynamoDBer
}

func (h *DynamoHelper) HitDynamo(query string) (*dynamodb.QueryOutput, error) {

	m := make(map[string]dynamodb.Condition)

	var alist []dynamodb.AttributeValue

	alist = append(alist, dynamodb.AttributeValue{
		S: aws.String(query),
	})

	m["Id"] = dynamodb.Condition{
		AttributeValueList: alist,
		ComparisonOperator: aws.String(dynamodb.ComparisonOperatorEq),
	}

	output, err := h.client.Query(&dynamodb.QueryInput{
		KeyConditions: m,
		TableName:     aws.String("foobar2"),
	})

	return output, err

}
