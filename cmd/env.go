package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"os"
)

func loadEnv(ctx context.Context, secrets secretFetcher, name string) (string, error) {
	envValue, ok := os.LookupEnv(name)
	if !ok {
		return "", fmt.Errorf("environment variable %s not found", name)
	}

	if arn.IsARN(envValue) {
		parsedArn, err := arn.Parse(envValue)
		if err != nil {
			return "", fmt.Errorf("error parsing arn: %w", err)
		}

		switch parsedArn.Service {
		case "secretsmanager":
			return secrets.FetchFromAWSSecretsManager(ctx, envValue)
		case "ssm":
			return secrets.FetchFromAWSSSMParameterStore(ctx, envValue)
		default:
			return "", fmt.Errorf("environment variable %s set to invalid ARN (unsupported service %s)", name, parsedArn.Service)
		}
	}

	return envValue, nil
}
