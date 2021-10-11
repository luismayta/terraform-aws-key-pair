output "instance" {
  description = "key_pair of instance."
  value       = aws_key_pair.this
}

output "name" {
  description = "name of key_pair."
  value       = aws_key_pair.this.key_name
}
