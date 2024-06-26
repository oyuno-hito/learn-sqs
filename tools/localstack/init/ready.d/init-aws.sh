#!/bin/bash

# Create SQS Queue
awslocal sqs create-queue \
  --region ap-northeast-1 \
  --queue-name default

# Create Lambda Function
awslocal lambda create-function \
  --region ap-northeast-1 \
  --role arn:aws:iam::000000000000:role/lambda-role\
  --function-name handler \
  --runtime provided.al2023 \
  --handler handler \
  --zip-file fileb:///opt/files/lambda.zip

# Update lambda environment variables
awslocal lambda update-function-configuration --region ap-northeast-1 --function-name handler --environment '{"Variables":{"MYSQL_USER":"user","MYSQL_PASSWORD":"password","MYSQL_HOST":"db","MYSQL_PORT":"3306","MYSQL_DATABASE":"learn_sqs"}}'

# Attach SQS Queue to Lambda Function
awslocal lambda create-event-source-mapping \
  --region ap-northeast-1 \
  --function-name handler \
  --batch-size 1 \
  --event-source arn:aws:sqs:ap-northeast-1:000000000000:default
