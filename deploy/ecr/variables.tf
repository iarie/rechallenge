variable "ecr_repo_name" {
  type        = string
  nullable    = false
  description = "An ESR Repository name"
}

variable "ecr_image_tag" {
  type        = string
  nullable    = false
  description = "Tag associated with image"
}

variable "region" {
  type        = string
  nullable    = false
  description = "AWS Region name"
}
