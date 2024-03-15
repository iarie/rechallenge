output "image_tag" {
  description = "App docker image"
  value       = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${var.region}.amazonaws.com/${data.aws_ecr_image.app.repository_name}:${data.aws_ecr_image.app.image_tag}"
}
