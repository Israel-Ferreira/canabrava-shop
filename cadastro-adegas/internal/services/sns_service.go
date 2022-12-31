package services

import (
	"encoding/json"

	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/internal/config"
	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/pkg/responses"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func SendMsgToSNS(msg responses.AdegaCreatedMsg) error {

	sess, err := session.NewSession(&aws.Config{
		Region:      &config.SnsConf.AwsRegion,
		Credentials: credentials.NewStaticCredentials(config.SnsConf.AwsAccessKey, config.SnsConf.AwsSecretKey, ""),
	})

	if err != nil {
		return err
	}

	svc := sns.New(sess)

	topic, err := svc.CreateTopic(&sns.CreateTopicInput{
		Name: &config.SnsConf.SnsQueueName,
	})

	if err != nil {
		return err
	}

	data, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	content := string(data)

	if _, err := svc.Publish(&sns.PublishInput{TopicArn: topic.TopicArn, Message: &content}); err != nil {
		return err
	}

	return nil

}
