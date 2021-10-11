variable "name" {
  type        = string
  description = "name of key pair, e.g. 'key-1' or 'key-2'"
}

variable "tags" {
  type        = map(string)
  description = "tags (e.g. `map('BusinessUnit','XYZ')`"
  default     = {}
}

variable "public_key" {
  type        = string
  description = "public key"
}
