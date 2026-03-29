package tfutils

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure TFUtilsProvider satisfies various provider interfaces.
var _ provider.Provider = &TFUtilsProvider{}
var _ provider.ProviderWithFunctions = &TFUtilsProvider{}

// TFUtilsProvider defines the provider implementation.
type TFUtilsProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func (p *TFUtilsProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "tfutils"
	resp.Version = p.version
}

func (p *TFUtilsProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A provider for various utility functions in Terraform/OpenTofu.",
		Attributes:  map[string]schema.Attribute{},
	}
}

func (p *TFUtilsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// No configuration needed for this provider
}

func (p *TFUtilsProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *TFUtilsProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewJsonFormatDataSource,
		NewJsonDeepMergeDataSource,
		NewYamlDeepMergeDataSource,
	}
}

func (p *TFUtilsProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewJsonFormatFunction,
		NewJsonDeepMergeFunction,
		NewYamlDeepMergeFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &TFUtilsProvider{
			version: version,
		}
	}
}
