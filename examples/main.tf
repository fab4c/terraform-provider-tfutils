terraform {
  required_providers {
    tfutils = {
      source = "fab4c/tfutils"
    }
  }
}

data "tfutils_json_format" "example" {
    json = "{\"a\":\"b\",\"myObj\":{\"prop1\":1}}"
    indent = 4
}

output "example_output" {
    value = data.tfutils_json_format.example.result
}