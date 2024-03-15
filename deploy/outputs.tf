output "alb_dns" {
  description = "Application's front entry"
  value       = module.alb.dns
}
