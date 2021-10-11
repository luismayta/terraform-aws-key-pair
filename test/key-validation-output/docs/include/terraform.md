<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.13 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_main"></a> [main](#module\_main) | ../.. | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_name"></a> [name](#input\_name) | name of key pair, e.g. 'key-1' or 'key-2' | `string` | n/a | yes |
| <a name="input_public_key"></a> [public\_key](#input\_public\_key) | public key | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | tags (e.g. `map('BusinessUnit','XYZ')` | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_instance"></a> [instance](#output\_instance) | key\_pair of instance. |
| <a name="output_name"></a> [name](#output\_name) | name of key\_pair. |
<!-- END_TF_DOCS -->