terraform {
  required_providers {
    tfutils = {
      source = "fab4c/tfutils"
    }
  }
}

provider "tfutils" {}

# Example 1: Using the provider function with default indent (Terraform 1.8+ / OpenTofu 1.7+)
output "function_example_default" {
  description = "JSON formatted using provider function with default 2-space indent"
  value       = provider::tfutils::json_format("{\"name\":\"test\",\"nested\":{\"key\":\"value\"}}")
}

# Example 1b: Using the provider function with custom indent
output "function_example_custom" {
  description = "JSON formatted using provider function with custom 4-space indent"
  value       = provider::tfutils::json_format("{\"name\":\"test\",\"nested\":{\"key\":\"value\"}}", 4)
}

# Example 2: Using the data source (works with all versions)
data "tfutils_json_format" "example" {
  json   = "{\"a\":\"b\",\"myObj\":{\"prop1\":1}}"
  indent = 4
}

output "data_source_example" {
  description = "JSON formatted using data source"
  value       = data.tfutils_json_format.example.result
}

# Example 3: Custom indentation with function
output "custom_indent_function" {
  description = "JSON with 4-space indentation"
  value       = provider::tfutils::json_format("{\"items\":[1,2,3]}", 4)
}

# Example 4: transcoded from YAML with default indent via data source
data "tfutils_json_format" "transcoded_yaml_output_func" {
  json = jsonencode(yamldecode(file("../.goreleaser.yml")))
}

output "transcoded_yaml_output_data" {
  value = data.tfutils_json_format.transcoded_yaml_output_func.result
}

# Example 5: transcoded from YAML with default indent via function
locals {
  yaml_file_content = yamldecode(file("../.goreleaser.yml"))
}

output "transcoded_yaml_output_func" {
  description = "YAML file transcoded to formatted JSON using function"
  value       = provider::tfutils::json_format(jsonencode(local.yaml_file_content))
}