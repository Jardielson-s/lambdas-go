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
    ]
  })
}
