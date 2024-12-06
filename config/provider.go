/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	"context"
	_ "embed"

	tjconfig "github.com/crossplane/upjet/pkg/config"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	provider "github.com/oracle/terraform-provider-oci/oci"
	"github.com/pkg/errors"
	"github.com/upbound/provider-oci/config/core"
	"github.com/upbound/provider-oci/hack"
)

const (
	resourcePrefix = "oci"
	modulePath     = "github.com/upbound/provider-oci"
)

//go:embed schema.json
var providerSchema string

func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider(_ context.Context, generationProvider bool) (*tjconfig.Provider, error) {
	var p *schema.Provider
	var err error
	if generationProvider {
		p, err = getProviderSchema(providerSchema)
	} else {
		p = provider.Provider()
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}

	pc := tjconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		tjconfig.WithRootGroup("oci.upbound.io"),
		tjconfig.WithIncludeList(resourceList(cliReconciledExternalNameConfigs)),
		tjconfig.WithNoForkIncludeList(resourceList(noForkExternalNameConfigs)),
		tjconfig.WithFeaturesPackage("internal/features"),
		tjconfig.WithMainTemplate(hack.MainTemplate),
		tjconfig.WithTerraformProvider(p),
		tjconfig.WithDefaultResourceOptions(
			resourceConfigurator(),
		))

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		core.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

// resourceList returns the list of resources that have external
// name configured in the specified table.
func resourceList(t map[string]tjconfig.ExternalName) []string {
	l := make([]string, len(t))
	i := 0
	for n := range t {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = n + "$"
		i++
	}
	return l
}
