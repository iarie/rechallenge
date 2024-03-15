output "task_role_arn" {
  description = "ECS Task Role Arn"
  value       = aws_iam_role.fargate_task_role.arn
}

output "execution_role_arn" {
  description = "ECS Execution Role Arn"
  value       = aws_iam_role.fargate_execution_role.arn
}
