variable "app_name" {
  type        = string
  nullable    = false
  description = "Desired application name"
}

variable "security_groups_ids" {
  type        = set(string)
  nullable    = false
  description = "VPC Security groups id assigned to load balancer"
}

variable "subnet_ids" {
  type        = set(string)
  nullable    = false
  description = "Set of subnet ids"
}

variable "internal" {
  type        = bool
  description = "If true, the LB will be internal."
}

variable "vpc_id" {
  type        = string
  nullable    = false
  description = "a VPC id for aws_lb_target_group"
}

