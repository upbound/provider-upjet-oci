/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// terraformPluginSDKExternalNameConfigs contains all external name configurations
// belonging to Terraform resources to be reconciled under the no-fork
// architecture for this provider.
var terraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{
	// core
	//
	// Vcns can be imported using the id
	"oci_core_vcn": config.IdentifierFromProvider,

	// objectstorage
	//
	// Buckets can be imported using the id, "n/{namespaceName}/b/{bucketName}"
	"oci_objectstorage_bucket": config.TemplatedStringAsIdentifier("name", "n/{{ .parameters.namespace }}/b/{{ .external_name }}"),
}

var CLIReconciledExternalNameConfigs = map[string]config.ExternalName{}

// ResourceConfigurator applies all external name configs
// listed in the table terraformPluginSDKExternalNameConfigs and
// CLIReconciledExternalNameConfigs and sets the version
// of those resources to v1beta1. For those resource in
// terraformPluginSDKExternalNameConfigs, it also sets
// config.Resource.UseNoForkClient to `true`.
func ResourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		// if configured both for the no-fork and CLI based architectures,
		// no-fork configuration prevails
		e, configured := terraformPluginSDKExternalNameConfigs[r.Name]
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
