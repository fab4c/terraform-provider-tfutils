package tfutils

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = YamlDeepMergeFunction{}
)

func NewYamlDeepMergeFunction() function.Function {
	return YamlDeepMergeFunction{}
}

type YamlDeepMergeFunction struct{}

func (f YamlDeepMergeFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "yaml_deepmerge"
}

func (f YamlDeepMergeFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "YAML Deep Merge function",
		MarkdownDescription: "Deeply merges two or more YAML strings into a single YAML string. Inputs are merged left-to-right; later inputs take precedence for scalar values and maps. Lists are replaced by later inputs.",
		VariadicParameter: function.StringParameter{
			Name:                "inputs",
			MarkdownDescription: "YAML strings to deep merge, ordered from lowest to highest precedence.",
		},
		Return: function.StringReturn{},
	}
}

func (f YamlDeepMergeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var inputs []string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &inputs))
	if resp.Error != nil {
		return
	}

	result, err := deepMergeYAMLStrings(inputs, false, false)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}
