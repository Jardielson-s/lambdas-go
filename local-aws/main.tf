terraform {
  required_version = ">= 1.2.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws",
      version = "~> 4.16"
    }
  }
}

provider "aws" {
  region = "us-east-1"
  # access_key                  = "teste"
  # secret_key                  = "teste"
  # region                      = "us-east-1"
  # s3_use_path_style           = true
  # skip_credentials_validation = true
  # skip_requesting_account_id  = true
  # skip_metadata_api_check     = true
  # endpoints {
  #   acm                      = "http://localhost:4566"
  #   apigateway               = "http://localhost:4566"
  #   cloudformation           = "http://localhost:4566"
  #   cloudwatch               = "http://localhost:4566"
  #   config                   = "http://localhost:4566"
  #   dynamodb                 = "http://localhost:4566"
  #   dynamodbstreams          = "http://localhost:4566"
  #   ec2                      = "http://localhost:4566"
  #   es                       = "http://localhost:4566"
  #   events                   = "http://localhost:4566"
  #   firehose                 = "http://localhost:4566"
  #   iam                      = "http://localhost:4566"
  #   kinesis                  = "http://localhost:4566"
  #   kms                      = "http://localhost:4566"
  #   lambda                   = "http://localhost:4566"
  #   logs                     = "http://localhost:4566"
  #   opensearch               = "http://localhost:4566"
  #   redshift                 = "http://localhost:4566"
  #   resourcegroupstaggingapi = "http://localhost:4566"
  #   route53                  = "http://localhost:4566"
  #   route53resolver          = "http://localhost:4566"
  #   s3                       = "http://localhost:4566"
  #   s3control                = "http://localhost:4566"
  #   secretsmanager           = "http://localhost:4566"
  #   ses                      = "http://localhost:4566"
  #   sns                      = "http://localhost:4566"
  #   sqs                      = "http://localhost:4566"
  #   ssm                      = "http://localhost:4566"
  #   stepfunctions            = "http://localhost:4566"
  #   sts                      = "http://localhost:4566"
  #   support                  = "http://localhost:4566"
  #   swf                      = "http://localhost:4566"
  #   transcribe               = "http://localhost:4566"
  # }
}

