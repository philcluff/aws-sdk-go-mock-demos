package app

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
)

func HitDynamo() (*dynamodb.QueryOutput, error) {

	m := make(map[string]dynamodb.Condition)

	var alist []dynamodb.AttributeValue

	alist = append(alist, dynamodb.AttributeValue{
		S: aws.String("somevalue"),
	})

	m["Id"] = dynamodb.Condition{
		AttributeValueList: alist,
		ComparisonOperator: aws.String(dynamodb.ComparisonOperatorEq),
	}

	client := dynamodb.New(nil)

	output, err := client.Query(&dynamodb.QueryInput{
		KeyConditions: m,
		TableName:     aws.String("foobar2"),
	})

	return output, err

}
