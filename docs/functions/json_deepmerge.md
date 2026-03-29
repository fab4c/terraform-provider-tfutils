---
page_title: "json_deepmerge function - tfutils"
subcategory: ""
description: |-
  JSON Deep Merge function
---

# function: json_deepmerge

Deeply merges two or more JSON strings into a single JSON string. Inputs are processed left-to-right; later inputs take precedence for scalar values and objects. Arrays are replaced by later inputs. Output is indented with 2 spaces.

## Signature

```text
json_deepmerge(inputs ...string) string
```

## Arguments

1. `inputs` (String, Variadic) Two or more JSON strings to deep merge, ordered from lowest to highest precedence.

If you need more customization when merging objects you may prefer to use the `tfutils_json_deepmerge` Data Source instead that allows for customization when working with nested arrays

## Example Usage

```terraform
# Merge two JSON objects - later input wins on conflicts
output "merged_json" {
  value = provider::tfutils::json_deepmerge(
    "{\"region\":\"us-east-1\",\"tags\":{\"env\":\"dev\",\"team\":\"platform\"}}",
    "{\"tags\":{\"env\":\"prod\",\"cost_center\":\"1234\"}}"
  )
  # {
  #   "region": "us-east-1",
  #   "tags": {
  #     "cost_center": "1234",
  #     "env": "prod",
  #     "team": "platform"
  #   }
  # }
}

# Merge three JSON objects
output "triple_merge" {
  value = provider::tfutils::json_deepmerge(
    "{\"a\":1,\"b\":2}",
    "{\"b\":3,\"c\":4}",
    "{\"c\":5,\"d\":6}"
  )
  # {
  #   "a": 1,
  #   "b": 3,
  #   "c": 5,
  #   "d": 6
  # }
}
```
