---
page_title: "yaml_deepmerge function - tfutils"
subcategory: ""
description: |-
  YAML Deep Merge function
---

# function: yaml_deepmerge

Deeply merges two or more YAML strings into a single YAML string. Inputs are processed left-to-right; later inputs take precedence for scalar values and maps. Lists are replaced by later inputs.

## Signature

```text
yaml_deepmerge(inputs ...string) string
```

## Arguments

1. `inputs` (String, Variadic) Two or more YAML strings to deep merge, ordered from lowest to highest precedence.

If you need more customization when merging objects you may prefer to use the `tfutils_yaml_deepmerge` Data Source instead that allows for customization when working with nested arrays

## Example Usage

```terraform
# Merge two YAML strings - later input wins on conflicts
output "merged_yaml" {
  value = provider::tfutils::yaml_deepmerge(
    "region: us-east-1\ntags:\n  env: dev\n  team: platform\n",
    "tags:\n  env: prod\n  cost_center: '1234'\n"
  )
  # region: us-east-1
  # tags:
  #   env: prod
  #   team: platform
  #   cost_center: "1234"
}

# Merge three YAML strings
output "triple_merge" {
  value = provider::tfutils::yaml_deepmerge(
    "a: 1\nb: 2\n",
    "b: 3\nc: 4\n",
    "c: 5\nd: 6\n"
  )
  # a: 1
  # b: 3
  # c: 5
  # d: 6
}
```
