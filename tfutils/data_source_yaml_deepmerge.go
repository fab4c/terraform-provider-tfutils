package tfutils

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &yamlDeepMergeDataSource{}
	_ datasource.DataSourceWithConfigure = &yamlDeepMergeDataSource{}
)

func NewYamlDeepMergeDataSource() datasource.DataSource {
	return &yamlDeepMergeDataSource{}
}

type yamlDeepMergeDataSource struct{}

type yamlDeepMergeDataSourceModel struct {
	Input        types.List   `tfsdk:"input"`
	AppendList   types.Bool   `tfsdk:"append_list"`
	DeepCopyList types.Bool   `tfsdk:"deep_copy_list"`
	Output       types.String `tfsdk:"output"`
	ID           types.String `tfsdk:"id"`
}

func (d *yamlDeepMergeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_yaml_deepmerge"
}

func (d *yamlDeepMergeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Accepts a list of YAML strings and deep-merges them into a single YAML string. Later inputs take precedence over earlier ones for scalar values and maps. Lists are replaced by default.",
		Attributes: map[string]schema.Attribute{
			"input": schema.ListAttribute{
				Description: "An ordered list of YAML strings to deep merge. Later entries take precedence.",
				Required:    true,
				ElementType: types.StringType,
			},
			"append_list": schema.BoolAttribute{
				Description: "When true, lists are appended rather than replaced by later inputs. Defaults to false.",
				Optional:    true,
			},
			"deep_copy_list": schema.BoolAttribute{
				Description: "When true, list elements are merged one-by-one: element N from each input is deep-merged together. Defaults to false.",
				Optional:    true,
			},
			"output": schema.StringAttribute{
				Description: "The deep-merged YAML string.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "Placeholder identifier attribute.",
				Computed:    true,
			},
		},
	}
}

func (d *yamlDeepMergeDataSource) Configure(_ context.Context, _ datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
}

func (d *yamlDeepMergeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data yamlDeepMergeDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var inputs []string
	resp.Diagnostics.Append(data.Input.ElementsAs(ctx, &inputs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	appendList := !data.AppendList.IsNull() && !data.AppendList.IsUnknown() && data.AppendList.ValueBool()
	deepCopyList := !data.DeepCopyList.IsNull() && !data.DeepCopyList.IsUnknown() && data.DeepCopyList.ValueBool()

	result, err := deepMergeYAMLStrings(inputs, appendList, deepCopyList)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deep-merging YAML",
			fmt.Sprintf("Could not deep merge YAML inputs: %s", err.Error()),
		)
		return
	}

	data.Output = types.StringValue(result)
	data.ID = types.StringValue("static")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
