package tfutils

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// deepMergeYAMLStrings parses each YAML string, deep-merges them left to right
// (later entries take precedence for scalar values and maps), and returns the
// result serialised as a YAML string.
func deepMergeYAMLStrings(inputs []string, appendList, deepCopyList bool) (string, error) {
	var maps []map[string]interface{}

	for _, input := range inputs {
		if strings.TrimSpace(input) == "" {
			continue
		}
		m := make(map[string]interface{})
		if err := yaml.Unmarshal([]byte(input), &m); err != nil {
			return "", fmt.Errorf("failed to parse YAML: %w", err)
		}
		maps = append(maps, m)
	}

	if len(maps) == 0 {
		return "", nil
	}

	result := maps[0]
	for _, m := range maps[1:] {
		result = yamlMergeMaps(result, m, appendList, deepCopyList)
	}

	out, err := yaml.Marshal(result)
	if err != nil {
		return "", fmt.Errorf("failed to marshal YAML: %w", err)
	}
	return string(out), nil
}

// yamlMergeMaps deep-merges src into dst, returning a new map.
func yamlMergeMaps(dst, src map[string]interface{}, appendList, deepCopyList bool) map[string]interface{} {
	result := make(map[string]interface{}, len(dst))
	for k, v := range dst {
		result[k] = v
	}
	for k, srcVal := range src {
		if dstVal, exists := result[k]; exists {
			result[k] = yamlMergeValues(dstVal, srcVal, appendList, deepCopyList)
		} else {
			result[k] = srcVal
		}
	}
	return result
}

// yamlMergeValues merges two values. Maps are recursed; slices are appended or
// merged element-by-element depending on the options; scalars are replaced by src.
func yamlMergeValues(dst, src interface{}, appendList, deepCopyList bool) interface{} {
	dstMap, dstIsMap := dst.(map[string]interface{})
	srcMap, srcIsMap := src.(map[string]interface{})
	if dstIsMap && srcIsMap {
		return yamlMergeMaps(dstMap, srcMap, appendList, deepCopyList)
	}

	dstSlice, dstIsSlice := dst.([]interface{})
	srcSlice, srcIsSlice := src.([]interface{})
	if dstIsSlice && srcIsSlice {
		if appendList {
			combined := make([]interface{}, len(dstSlice)+len(srcSlice))
			copy(combined, dstSlice)
			copy(combined[len(dstSlice):], srcSlice)
			return combined
		}
		if deepCopyList {
			return yamlMergeSlices(dstSlice, srcSlice, appendList, deepCopyList)
		}
	}

	// Default: src overrides dst.
	return src
}

// yamlMergeSlices merges two slices element-by-element.
func yamlMergeSlices(dst, src []interface{}, appendList, deepCopyList bool) []interface{} {
	maxLen := len(dst)
	if len(src) > maxLen {
		maxLen = len(src)
	}
	result := make([]interface{}, maxLen)
	for i := 0; i < maxLen; i++ {
		switch {
		case i >= len(dst):
			result[i] = src[i]
		case i >= len(src):
			result[i] = dst[i]
		default:
			result[i] = yamlMergeValues(dst[i], src[i], appendList, deepCopyList)
		}
	}
	return result
}
