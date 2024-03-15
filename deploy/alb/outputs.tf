output "dns" {
  description = "Application's front entry"
  value       = aws_lb.primary.dns_name
}

output "target_group_arn" {
  description = "Target group arn"
  value       = aws_lb_target_group.primary.arn
}
