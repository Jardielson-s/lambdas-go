data "aws_caller_identity" "aws_identity" {

}
locals {
  aws_account_id = data.aws_caller_identity.aws_identity.account_id
}
resource "aws_iam_role" "lambda_role" {
  name = "lambda_execution_role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy" "lambda_policy" {
  name       = "lambda_policy"
  role       = aws_iam_role.lambda_role.id
  depends_on = [aws_iam_role.lambda_role]
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        Effect   = "Allow"
        Resource = "*"
      },
      {
        Action   = "lambda:*",
        Effect   = "Allow"
        Resource = "*"
      },
      {
        Action = [
          "s3:GetObject"
        ],
        Effect   = "Allow"
        Resource = "arn:aws:s3:::${var.bucket_name}-${var.env}/*"
      },
      {
        Action = [
          "sqs:SendMessage",
          "sqs:GetQueueUrl",
          "sqs:ReceiveMessage",
          "sqs:DeleteMessage",
          "sqs:GetQueueAttributes",
          "sqs:ChangeMessageVisibility"
        ],
        Effect   = "Allow"
        Resource = "arn:aws:sqs:${var.region}:${local.aws_account_id}:${var.queue_name}-${var.env}"
      },
    ]
  })
}
