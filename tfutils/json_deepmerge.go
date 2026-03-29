package tfutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// deepMergeJSONStrings parses each JSON string, deep-merges them left to right
// (later entries take precedence for scalar values and maps), and returns the
// result serialised as an indented JSON string.
func deepMergeJSONStrings(inputs []string, appendList, deepCopyList bool, indent int) (string, error) {
	var maps []map[string]interface{}

	for _, input := range inputs {
		if strings.TrimSpace(input) == "" {
			continue
		}
		m := make(map[string]interface{})
		if err := json.Unmarshal([]byte(input), &m); err != nil {
			return "", fmt.Errorf("failed to parse JSON: %w", err)
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

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetIndent("", strings.Repeat(" ", indent))
	enc.SetEscapeHTML(false)
	if err := enc.Encode(result); err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}
	// json.Encoder.Encode appends a trailing newline; trim it for consistency.
	return strings.TrimRight(buf.String(), "\n"), nil
}
