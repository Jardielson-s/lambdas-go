resource "aws_cloudwatch_log_group" "lambda_log_group" {
  for_each          = { for k, v in var.lambdas : k => v if k == "process_csv" }
  name              = "/aws/lambda/${each.value}"
  retention_in_days = 7
  lifecycle {
    prevent_destroy = false
  }
}
