

resource "aws_lb_target_group" "adminv2_lb" {
  name        = "admin-mock-api"
  target_type = "instance"
  port        = 8080
  protocol    = "HTTP"
  vpc_id      = data.aws_vpc.dev_vpc.id
  health_check {
    healthy_threshold   = var.health_check["healthy_threshold"]
    interval            = var.health_check["interval"]
    unhealthy_threshold = var.health_check["unhealthy_threshold"]
    timeout             = var.health_check["timeout"]
    path                = var.health_check["path"]
    port                = var.health_check["port"]
  }
}

resource "aws_ecs_service" "adminmockapi" {
  name                               = "admin-mock-api"
  cluster                            = data.aws_ecs_cluster.ecs-qa.id
  launch_type                        = "EC2"
  task_definition                    = data.aws_ecs_task_definition.admin_task.arn
  desired_count                      = 1
  iam_role                           = data.aws_iam_role.ecs_role.arn
  scheduling_strategy                = "REPLICA"
  deployment_maximum_percent         = 100
  deployment_minimum_healthy_percent = 0

  ordered_placement_strategy {
    type  = "spread"
    field = "instanceId"
  }

  deployment_circuit_breaker {
    enable   = true
    rollback = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.adminv2_lb.arn
    container_name   = "admin-mock-api"
    container_port   = 8080
  }

  depends_on = [
    data.aws_iam_role.ecs_role,
    aws_lb_target_group.adminv2_lb
  ]
}