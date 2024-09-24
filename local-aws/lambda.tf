resource "random_id" "unique_suffix" {
  byte_length = 2
}

locals {
  binary_name = "app"
  app_id      = "test-dev-${random_id.unique_suffix.hex}"
}

resource "null_resource" "function_binary" {
  for_each = var.lambdas
  triggers = {
    sequence = each.key
  }
  provisioner "local-exec" {
    command = <<EOT
      mkdir ${each.key}
      cd ${each.key}
      GOOS=linux CGO_ENABLED=0 go build -ldflags '-s -w' -o bootstrap ../src/infra/functions/${each.key}/${each.key}.go
      cd ..
    EOT
    //"mkdir ${each.key} &&  GOOS=linux CGO_ENABLED=0 go build -ldflags '-s -w' -o bootstrap ../src/infra/functions/${each.key}/${each.key}.go && cd .."
  }
}

output "null_resource" {
  description = "null_resource"
  value       = null_resource.function_binary
}
data "archive_file" "lambda_zip" {
  for_each = var.lambdas
  type     = "zip"
  #   source_dir  = "../"
  output_path = "${each.key}/bootstrap.zip"
  source_file = "${each.key}/bootstrap"
  #   excludes    = ["lambda.zip", "environment", ".gitignore", ".env", "README.md", ".git", "requirements.txt", "terraform", "images"]
  depends_on = [null_resource.function_binary]
}


resource "aws_lambda_function" "terraform_lambda_func" {
  for_each         = var.lambdas
  filename         = data.archive_file.lambda_zip[each.key].output_path
  function_name    = "${each.key}_${var.env}"
  role             = aws_iam_role.lambda_role.arn
  handler          = each.key
  source_code_hash = base64sha256(data.archive_file.lambda_zip[each.key].source_file)
  runtime          = "provided.al2"
  timeout          = 20
  depends_on       = [aws_cloudwatch_log_group.lambda_log_group]
}
