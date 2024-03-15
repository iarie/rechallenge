output "id" {
  description = "Primary VPC id"
  value       = aws_vpc.primary.id
}

output "public_subnet_ids" {
  description = "Public Subnets ids"
  value       = [for v in aws_subnet.public : v.id]
}

output "security_group_id" {
  description = "Security Group id"
  value       = aws_security_group.primary.id
}
