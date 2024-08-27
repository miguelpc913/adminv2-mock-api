resource "aws_cloudwatch_log_group" "ContainerLogGroup" {
  name              = "ECS-CW-Containers"
  retention_in_days = var.retention_in_days
}

resource "aws_cloudwatch_log_stream" "log_stream" {
  name           = "qa-adminv2-mock"
  log_group_name = aws_cloudwatch_log_group.ContainerLogGroup.name
}