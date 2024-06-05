#!/bin/bash

# Create SQS Queue
awslocal sqs create-queue \
  --region ap-northeast-1 \
  --queue-name default
