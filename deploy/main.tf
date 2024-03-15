variable "aws_profile" {
    default = "default"
}

variable "aws_region" {
    default = "eu-west-2"
}

variable "app_name" {
    default = "rechallenge"
}

provider "aws" {
  profile = var.aws_profile
  region  = var.aws_region
}

terraform {
  backend "local" {
    path = "terraform.tfstate"
  }
}

module "iam" {
  source = "./iam"
}

module "ecr" {
  source        = "./ecr"
  region        = var.aws_region
  ecr_repo_name = "redocker"
  ecr_image_tag = "latest"
}

module "vpc" {
  source             = "./vpc"
  availability_zones = tolist(["euw2-az2", "euw2-az3"])
}


module "alb" {
  source              = "./alb"
  app_name            = var.app_name
  security_groups_ids = [module.vpc.security_group_id]
  subnet_ids          = module.vpc.public_subnet_ids
  internal            = false
  vpc_id              = module.vpc.id
}


resource "aws_ecs_cluster" "re_cluster" {
  name = "re-cluster"
}

resource "aws_ecs_service" "re_service" {
  name            = "${var.app_name}-service"
  cluster         = aws_ecs_cluster.re_cluster.arn
  task_definition = aws_ecs_task_definition.re_task.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    security_groups  = [module.vpc.security_group_id]
    subnets          = module.vpc.public_subnet_ids
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = module.alb.target_group_arn
    container_name   = "re-container"
    container_port   = 80
  }
}

resource "aws_ecs_task_definition" "re_task" {
  family                   = "my-task"
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]

  task_role_arn      = module.iam.task_role_arn
  execution_role_arn = module.iam.execution_role_arn

  container_definitions = jsonencode([
    {
      name      = "re-container"
      image     = module.ecr.image_tag
      cpu       = 256
      memory    = 512
      essential = true
      portMappings = [
        {
          containerPort = 80
        }
      ]
      logConfiguration = {
        logDriver = "awslogs",
        options = {
          awslogs-group         = var.app_name,
          awslogs-region        = var.aws_region,
          awslogs-stream-prefix = var.app_name
        }
      },
    }
  ])
}

resource "aws_cloudwatch_log_group" "primary" {
  name = var.app_name
}

resource "aws_cloudwatch_log_stream" "primary" {
  name           = var.app_name
  log_group_name = aws_cloudwatch_log_group.primary.name
}

