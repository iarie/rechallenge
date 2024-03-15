resource "aws_lb" "primary" {
  load_balancer_type = "application"
  name               = var.app_name
  security_groups    = var.security_groups_ids
  subnets            = var.subnet_ids
  internal           = var.internal
}

resource "aws_lb_target_group" "primary" {
  name        = "${var.app_name}-primary-tg"
  port        = 80
  protocol    = "HTTP"
  vpc_id      = var.vpc_id
  target_type = "ip"
}

resource "aws_lb_listener" "entrypoint" {
  load_balancer_arn = aws_lb.primary.arn
  port              = 80
  protocol          = "HTTP"
  depends_on        = [aws_lb_target_group.primary]

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.primary.arn
  }
}
