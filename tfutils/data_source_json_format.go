// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfutils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &jsonFormatDataSource{}
	_ datasource.DataSourceWithConfigure = &jsonFormatDataSource{}
)

func NewJsonFormatDataSource() datasource.DataSource {
	return &jsonFormatDataSource{}
}

type jsonFormatDataSource struct{}

type jsonFormatDataSourceModel struct {
	JSON   types.String `tfsdk:"json"`
	Indent types.Int64  `tfsdk:"indent"`
	Result types.String `tfsdk:"result"`
	ID     types.String `tfsdk:"id"`
}

func (d *jsonFormatDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_json_format"
}

func (d *jsonFormatDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Formats / beautifies a given JSON string.",
		Attributes: map[string]schema.Attribute{
			"json": schema.StringAttribute{
				Description: "A JSON string to format / beautify.",
				Required:    true,
			},
			"indent": schema.Int64Attribute{
				Description: "The number of spaces to use for each indent. The default is two spaces.",
				Optional:    true,
			},
			"result": schema.StringAttribute{
				Description: "A formatted JSON string.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "Placeholder identifier attribute.",
				Computed:    true,
			},
		},
	}
}

func (d *jsonFormatDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// No provider-level configuration needed
}

func (d *jsonFormatDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data jsonFormatDataSourceModel

	// Read configuration from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Default indent to 2 if not provided
	indent := int64(2)
	if !data.Indent.IsNull() && !data.Indent.IsUnknown() {
		indent = data.Indent.ValueInt64()
	}

	// Format the JSON
	result := &bytes.Buffer{}
	spaces := strings.Repeat(" ", int(indent))

	if err := json.Indent(result, []byte(data.JSON.ValueString()), "", spaces); err != nil {
		resp.Diagnostics.AddError(
			"Error formatting JSON",
			fmt.Sprintf("Could not format JSON: %s", err.Error()),
		)
		return
	}

	// Set the result
	data.Result = types.StringValue(result.String())
	data.ID = types.StringValue("static")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
