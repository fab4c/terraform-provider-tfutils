package tfutils

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = JsonDeepMergeFunction{}
)

func NewJsonDeepMergeFunction() function.Function {
	return JsonDeepMergeFunction{}
}

type JsonDeepMergeFunction struct{}

func (f JsonDeepMergeFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "json_deepmerge"
}

func (f JsonDeepMergeFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "JSON Deep Merge function",
		MarkdownDescription: "Deeply merges two or more JSON strings into a single JSON string. Inputs are merged left-to-right; later inputs take precedence for scalar values and objects. Arrays are replaced by later inputs. Output is indented with 2 spaces.",
		VariadicParameter: function.StringParameter{
			Name:                "inputs",
			MarkdownDescription: "JSON strings to deep merge, ordered from lowest to highest precedence.",
		},
		Return: function.StringReturn{},
	}
}

func (f JsonDeepMergeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var inputs []string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &inputs))
	if resp.Error != nil {
		return
	}

	result, err := deepMergeJSONStrings(inputs, false, false, 2)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}
