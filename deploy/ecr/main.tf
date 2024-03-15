data "aws_ecr_repository" "repo" {
  name = var.ecr_repo_name
}

data "aws_ecr_image" "app" {
  repository_name = data.aws_ecr_repository.repo.name
  image_tag       = var.ecr_image_tag
}

data "aws_caller_identity" "current" {}
