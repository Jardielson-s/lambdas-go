resource "aws_s3_bucket" "bucket_backup" {
  bucket        = "${var.bucket_name}-${var.env}"
  force_destroy = true
  tags = {
    Environment = var.env
  }
}

resource "aws_s3_bucket_notification" "aws_lambda_trigger" {
  bucket     = aws_s3_bucket.bucket_backup.id
  depends_on = [aws_lambda_function.terraform_lambda_func, aws_s3_bucket.bucket_backup]
  dynamic "lambda_function" {
    for_each = { for k, v in aws_lambda_function.terraform_lambda_func : k => v if k == "process_csv" }
    content {
      lambda_function_arn = lambda_function.value.arn
      events              = ["s3:ObjectCreated:*"]
    }
  }
}

resource "aws_lambda_permission" "s3_aws_lambda_permission" {
  for_each      = { for k, v in aws_lambda_function.terraform_lambda_func : k => v if k == "process_csv" }
  statement_id  = "AllowS3Invoke"
  action        = "lambda:InvokeFunction"
  function_name = each.value.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = "arn:aws:s3:::${aws_s3_bucket.bucket_backup.id}"
}
