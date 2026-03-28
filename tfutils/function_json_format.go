// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfutils

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = JsonFormatFunction{}
)

func NewJsonFormatFunction() function.Function {
	return JsonFormatFunction{}
}

type JsonFormatFunction struct{}

func (f JsonFormatFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "json_format"
}

func (f JsonFormatFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "JSON Format function",
		MarkdownDescription: "Formats / beautifies a given JSON string",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "json",
				MarkdownDescription: "A JSON string to format / beautify",
			},
		},
		VariadicParameter: function.Int64Parameter{
			Name:                "indent",
			MarkdownDescription: "The number of spaces to use for each indent (default: 2)",
		},
		Return: function.StringReturn{},
	}
}

func (f JsonFormatFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var jsonStr string
	var indentValues []int64

	// Get the required json argument and optional indent variadic argument
	err := req.Arguments.Get(ctx, &jsonStr, &indentValues)
	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	// Default indent to 2 if not provided
	indent := int64(2)
	if len(indentValues) > 0 {
		indent = indentValues[0]
		// Ensure indent is at least 1
		if indent < 1 {
			indent = 2
		}
	}

	// Format the JSON
	result := &bytes.Buffer{}
	spaces := strings.Repeat(" ", int(indent))

	if err := json.Indent(result, []byte(jsonStr), "", spaces); err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	// Set the result
	err = resp.Result.Set(ctx, result.String())
	resp.Error = function.ConcatFuncErrors(err)
}
