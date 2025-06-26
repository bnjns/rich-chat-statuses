package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/smithy-go/ptr"
)

type secretFetcher interface {
	FetchFromAWSSecretsManager(ctx context.Context, secretArn string) (string, error)
	FetchFromAWSSSMParameterStore(ctx context.Context, parameterArn string) (string, error)
}

type secretClients struct {
	secrets *secretsmanager.Client
	ssm     *ssm.Client
}

func newSecretClients(ctx context.Context) (*secretClients, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &secretClients{
		secrets: secretsmanager.NewFromConfig(cfg),
		ssm:     ssm.NewFromConfig(cfg),
	}, nil
}

func (c *secretClients) FetchFromAWSSecretsManager(ctx context.Context, secretArn string) (string, error) {
	out, err := c.secrets.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: &secretArn,
	})
	if err != nil {
		return "", fmt.Errorf("error fetching secret %s: %w", secretArn, err)
	}

	return *out.SecretString, nil
}

func (c *secretClients) FetchFromAWSSSMParameterStore(ctx context.Context, parameterArn string) (string, error) {
	out, err := c.ssm.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           &parameterArn,
		WithDecryption: ptr.Bool(true),
	})
	if err != nil {
		return "", fmt.Errorf("error fetching SSM parameter %s: %w", parameterArn, err)
	}

	return *out.Parameter.Value, nil
}
