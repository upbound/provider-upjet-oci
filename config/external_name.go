/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// NoForkExternalNameConfigs contains all external name configurations
// belonging to Terraform resources to be reconciled under the no-fork
// architecture for this provider.
var NoForkExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"oci_core_vcn": config.IdentifierFromProvider,
}

var CLIReconciledExternalNameConfigs = map[string]config.ExternalName{}

// ResourceConfigurator applies all external name configs
// listed in the table NoForkExternalNameConfigs and
// CLIReconciledExternalNameConfigs and sets the version
// of those resources to v1beta1. For those resource in
// NoForkExternalNameConfigs, it also sets
// config.Resource.UseNoForkClient to `true`.
func ResourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		// if configured both for the no-fork and CLI based architectures,
		// no-fork configuration prevails
		e, configured := NoForkExternalNameConfigs[r.Name]
		if !configured {
			e, configured = CLIReconciledExternalNameConfigs[r.Name]
		}
		if !configured {
			return
		}
		r.Version = "v1beta1"
		r.ExternalName = e
	}
}
