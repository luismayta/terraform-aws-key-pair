locals {
  defaults = {
    tags = {
      "management" = "terraform"
    }
  }

  input = {
    tags = try(var.tags, local.defaults.tags)
  }

  generated = {
    name = var.name
    tags = local.input.tags
  }

  outputs = {
    name = local.generated.name
    tags = local.generated.tags
  }

}

resource "aws_key_pair" "this" {
  key_name   = local.outputs.name
  public_key = file(var.public_key)
  tags       = local.generated.tags

  lifecycle {
    create_before_destroy = true
  }
}
