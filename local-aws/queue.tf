resource "aws_sqs_queue" "aws_queue" {
  name                      = "${var.queue_name}-${var.env}"
  message_retention_seconds = 86400

  tags = {
    Environment = var.env
  }

}
