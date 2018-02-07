package telegram

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"os"
	"strconv"
)

type dynamoStore struct {
	db *dynamodb.DynamoDB
}

func NewDynamoState(db *dynamodb.DynamoDB) Store {
	return &dynamoStore{db: db}
}

func (s *dynamoStore) SetState(user int, state string, field []string) error {

	r := State{State: state, Userid: user, Field: field}

	av, err := dynamodbattribute.MarshalMap(r)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(getStateTableName()),
	}

	_, err = s.db.PutItem(input)
	return err
}
func (s *dynamoStore) getState(user int) State {
	record, err := s.getUser(user)
	if err != nil {
		log.Println(err.Error())
		return State{}
	}
	return *record
}

func (s *dynamoStore) getUser(user int) (*State, error) {
	query := &dynamodb.GetItemInput{
		TableName: aws.String(getStateTableName()),
		Key: map[string]*dynamodb.AttributeValue{
			"Userid":
			{

				N: aws.String(strconv.Itoa(user)),
			},
		},
	}

	resp, err := s.db.GetItem(query)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return &State{}, nil
	}

	r := State{}
	err = dynamodbattribute.UnmarshalMap(resp.Item, &r)
	return &r, err
}

func getStateTableName() string {
	return os.Getenv("STATE_TABLE")
}
