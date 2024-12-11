package core

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_core_vcn", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "core"
		// This custom TerraformDiff function removes the `byoipv6cidr_details.#` attribute
		// from the diff process to prevent the detection of unnecessary differences (diffs).
		// However, since this makes the changes to this attribute insensitive, it should be
		// applied with caution. To maintain compatibility with future Terraform version updates
		// or API changes, such customizations should be carefully reviewed.
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil && diff.Attributes != nil {
				delete(diff.Attributes, "byoipv6cidr_details.#")
			}
			return diff, nil
		}
	})
}
