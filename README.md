# Terraform Provider TFUtils

A Terraform/OpenTofu provider that provides utility functions for common operations.

> Attribution. This provider is based on the original work published at
> https://github.com/TheNicholi/terraform-provider-json-formatter and
> https://github.com/Yantrio/terraform-provider-tfutils

## Features

This provider supports both **data sources** and **provider functions** (Terraform 1.8+/OpenTofu 1.7+), allowing you to choose the approach that best fits your needs.

### Available Utilities

- **json_format**: Formats/beautifies a JSON string with configurable indentation

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.8 (for provider functions) or >= 1.0 (for data sources only)
- [OpenTofu](https://opentofu.org) >= 1.7
- [Go](https://golang.org/doc/install) >= 1.21

## Usage

### Using Provider Functions (Terraform 1.8+ / OpenTofu 1.7+)

Provider functions offer a more concise syntax:

```hcl
terraform {
  required_providers {
    tfutils = {
      source = "fab4c/tfutils"
    }
  }
}

provider "tfutils" {}

output "formatted_json" {
  value = provider::tfutils::json_format("{\"a\":\"b\",\"myObj\":{\"prop1\":1}}", 4)
}
```

**Output**

```console
formatted_json = {
    "a": "b",
    "myObj": {
        "prop1": 1
    }
}
```

### Using Data Sources (All Versions)

Data sources work with all Terraform/OpenTofu versions:

```hcl
terraform {
  required_providers {
    tfutils = {
      source = "fab4c/tfutils"
    }
  }
}

provider "tfutils" {}

data "tfutils_json_format" "example" {
  json   = "{\"a\":\"b\",\"myObj\":{\"prop1\":1}}"
  indent = 4
}

output "formatted_json" {
  value = data.tfutils_json_format.example.result
}
```

**Output**

```console
formatted_json = {
    "a": "b",
    "myObj": {
        "prop1": 1
    }
}
```

## License

This provider is distributed under the MPL-2.0 license.
