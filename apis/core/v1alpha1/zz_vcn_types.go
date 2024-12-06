// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type Byoipv6CidrDetailsInitParameters struct {

	// The OCID of the ByoipRange resource to which the CIDR block belongs.
	Byoipv6RangeID *string `json:"byoipv6rangeId,omitempty" tf:"byoipv6range_id,omitempty"`

	// An IPv6 prefix required to create a VCN with a BYOIP prefix. It could be the whole prefix identified in byoipv6RangeId, or a subrange. Example: 2001:0db8:0123::/48
	Ipv6CidrBlock *string `json:"ipv6cidrBlock,omitempty" tf:"ipv6cidr_block,omitempty"`
}

type Byoipv6CidrDetailsObservation struct {

	// The OCID of the ByoipRange resource to which the CIDR block belongs.
	Byoipv6RangeID *string `json:"byoipv6rangeId,omitempty" tf:"byoipv6range_id,omitempty"`

	// An IPv6 prefix required to create a VCN with a BYOIP prefix. It could be the whole prefix identified in byoipv6RangeId, or a subrange. Example: 2001:0db8:0123::/48
	Ipv6CidrBlock *string `json:"ipv6cidrBlock,omitempty" tf:"ipv6cidr_block,omitempty"`
}

type Byoipv6CidrDetailsParameters struct {

	// The OCID of the ByoipRange resource to which the CIDR block belongs.
	// +kubebuilder:validation:Optional
	Byoipv6RangeID *string `json:"byoipv6rangeId" tf:"byoipv6range_id,omitempty"`

	// An IPv6 prefix required to create a VCN with a BYOIP prefix. It could be the whole prefix identified in byoipv6RangeId, or a subrange. Example: 2001:0db8:0123::/48
	// +kubebuilder:validation:Optional
	Ipv6CidrBlock *string `json:"ipv6cidrBlock" tf:"ipv6cidr_block,omitempty"`
}

type VcnInitParameters struct {

	// The list of BYOIPv6 OCIDs and BYOIPv6 prefixes required to create a VCN that uses BYOIPv6 address ranges.
	Byoipv6CidrDetails []Byoipv6CidrDetailsInitParameters `json:"byoipv6cidrDetails,omitempty" tf:"byoipv6cidr_details,omitempty"`

	// Deprecated. Do not set this value. Use cidr_blocks instead. Example: 10.0.0.0/16
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// (Updatable) The list of one or more IPv4 CIDR blocks for the VCN that meet the following criteria:
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	// (Updatable) The OCID of the compartment to contain the VCN.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// A DNS label for the VCN, used in conjunction with the VNIC's hostname and subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, bminstance1.subnet123.vcn1.oraclevcn.com). Not required to be unique, but it's a best practice to set unique DNS labels for VCNs in your tenancy. Must be an alphanumeric string that begins with a letter. The value cannot be changed.
	// You must set this value if you want instances to be able to use hostnames to resolve other instances in the VCN. Otherwise the Internet and VCN Resolver will not work.
	DNSLabel *string `json:"dnsLabel,omitempty" tf:"dns_label,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The list of one or more ULA or Private IPv6 CIDR blocks for the vcn that meets the following criteria:
	Ipv6PrivateCidrBlocks []*string `json:"ipv6privateCidrBlocks,omitempty" tf:"ipv6private_cidr_blocks,omitempty"`

	// Whether IPv6 is enabled for the VCN. Default is false. If enabled, Oracle will assign the VCN a IPv6 /56 CIDR block. You may skip having Oracle allocate the VCN a IPv6 /56 CIDR block by setting isOracleGuaAllocationEnabled to false. For important details about IPv6 addressing in a VCN, see IPv6 Addresses.  Example: true
	IsIpv6Enabled *bool `json:"isIpv6Enabled,omitempty" tf:"is_ipv6enabled,omitempty"`

	// Specifies whether to skip Oracle allocated IPv6 GUA. By default, Oracle will allocate one GUA of /56 size for an IPv6 enabled VCN.
	IsOracleGuaAllocationEnabled *bool `json:"isOracleGuaAllocationEnabled,omitempty" tf:"is_oracle_gua_allocation_enabled,omitempty"`

	// (Updatable) Security Attributes for this resource. This is unique to ZPR, and helps identify which resources are allowed to be accessed by what permission controls.  Example: {"Oracle-DataSecurity-ZPR.MaxEgressCount.value": "42", "Oracle-DataSecurity-ZPR.MaxEgressCount.mode": "audit"}
	// +mapType=granular
	SecurityAttributes map[string]*string `json:"securityAttributes,omitempty" tf:"security_attributes,omitempty"`
}

type VcnObservation struct {

	// The list of BYOIPv6 prefixes required to create a VCN that uses BYOIPv6 ranges.
	Byoipv6CidrBlocks []*string `json:"byoipv6cidrBlocks,omitempty" tf:"byoipv6cidr_blocks,omitempty"`

	// The list of BYOIPv6 OCIDs and BYOIPv6 prefixes required to create a VCN that uses BYOIPv6 address ranges.
	Byoipv6CidrDetails []Byoipv6CidrDetailsObservation `json:"byoipv6cidrDetails,omitempty" tf:"byoipv6cidr_details,omitempty"`

	// Deprecated. Do not set this value. Use cidr_blocks instead. Example: 10.0.0.0/16
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// (Updatable) The list of one or more IPv4 CIDR blocks for the VCN that meet the following criteria:
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	// (Updatable) The OCID of the compartment to contain the VCN.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// A DNS label for the VCN, used in conjunction with the VNIC's hostname and subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, bminstance1.subnet123.vcn1.oraclevcn.com). Not required to be unique, but it's a best practice to set unique DNS labels for VCNs in your tenancy. Must be an alphanumeric string that begins with a letter. The value cannot be changed.
	// You must set this value if you want instances to be able to use hostnames to resolve other instances in the VCN. Otherwise the Internet and VCN Resolver will not work.
	DNSLabel *string `json:"dnsLabel,omitempty" tf:"dns_label,omitempty"`

	// The OCID for the VCN's default set of DHCP options.
	DefaultDHCPOptionsID *string `json:"defaultDhcpOptionsId,omitempty" tf:"default_dhcp_options_id,omitempty"`

	// The OCID for the VCN's default route table.
	DefaultRouteTableID *string `json:"defaultRouteTableId,omitempty" tf:"default_route_table_id,omitempty"`

	// The OCID for the VCN's default security list.
	DefaultSecurityListID *string `json:"defaultSecurityListId,omitempty" tf:"default_security_list_id,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The VCN's Oracle ID (OCID).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// For an IPv6-enabled VCN, this is the list of IPv6 prefixes for the VCN's IP address space. The prefixes are provided by Oracle and the sizes are always /56.
	Ipv6CidrBlocks []*string `json:"ipv6cidrBlocks,omitempty" tf:"ipv6cidr_blocks,omitempty"`

	// The list of one or more ULA or Private IPv6 CIDR blocks for the vcn that meets the following criteria:
	Ipv6PrivateCidrBlocks []*string `json:"ipv6privateCidrBlocks,omitempty" tf:"ipv6private_cidr_blocks,omitempty"`

	// Whether IPv6 is enabled for the VCN. Default is false. If enabled, Oracle will assign the VCN a IPv6 /56 CIDR block. You may skip having Oracle allocate the VCN a IPv6 /56 CIDR block by setting isOracleGuaAllocationEnabled to false. For important details about IPv6 addressing in a VCN, see IPv6 Addresses.  Example: true
	IsIpv6Enabled *bool `json:"isIpv6Enabled,omitempty" tf:"is_ipv6enabled,omitempty"`

	// Specifies whether to skip Oracle allocated IPv6 GUA. By default, Oracle will allocate one GUA of /56 size for an IPv6 enabled VCN.
	IsOracleGuaAllocationEnabled *bool `json:"isOracleGuaAllocationEnabled,omitempty" tf:"is_oracle_gua_allocation_enabled,omitempty"`

	// (Updatable) Security Attributes for this resource. This is unique to ZPR, and helps identify which resources are allowed to be accessed by what permission controls.  Example: {"Oracle-DataSecurity-ZPR.MaxEgressCount.value": "42", "Oracle-DataSecurity-ZPR.MaxEgressCount.mode": "audit"}
	// +mapType=granular
	SecurityAttributes map[string]*string `json:"securityAttributes,omitempty" tf:"security_attributes,omitempty"`

	// The VCN's current state.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time the VCN was created, in the format defined by RFC3339.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The VCN's domain name, which consists of the VCN's DNS label, and the oraclevcn.com domain.
	VcnDomainName *string `json:"vcnDomainName,omitempty" tf:"vcn_domain_name,omitempty"`
}

type VcnParameters struct {

	// The list of BYOIPv6 OCIDs and BYOIPv6 prefixes required to create a VCN that uses BYOIPv6 address ranges.
	// +kubebuilder:validation:Optional
	Byoipv6CidrDetails []Byoipv6CidrDetailsParameters `json:"byoipv6cidrDetails,omitempty" tf:"byoipv6cidr_details,omitempty"`

	// Deprecated. Do not set this value. Use cidr_blocks instead. Example: 10.0.0.0/16
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// (Updatable) The list of one or more IPv4 CIDR blocks for the VCN that meet the following criteria:
	// +kubebuilder:validation:Optional
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	// (Updatable) The OCID of the compartment to contain the VCN.
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// A DNS label for the VCN, used in conjunction with the VNIC's hostname and subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, bminstance1.subnet123.vcn1.oraclevcn.com). Not required to be unique, but it's a best practice to set unique DNS labels for VCNs in your tenancy. Must be an alphanumeric string that begins with a letter. The value cannot be changed.
	// You must set this value if you want instances to be able to use hostnames to resolve other instances in the VCN. Otherwise the Internet and VCN Resolver will not work.
	// +kubebuilder:validation:Optional
	DNSLabel *string `json:"dnsLabel,omitempty" tf:"dns_label,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The list of one or more ULA or Private IPv6 CIDR blocks for the vcn that meets the following criteria:
	// +kubebuilder:validation:Optional
	Ipv6PrivateCidrBlocks []*string `json:"ipv6privateCidrBlocks,omitempty" tf:"ipv6private_cidr_blocks,omitempty"`

	// Whether IPv6 is enabled for the VCN. Default is false. If enabled, Oracle will assign the VCN a IPv6 /56 CIDR block. You may skip having Oracle allocate the VCN a IPv6 /56 CIDR block by setting isOracleGuaAllocationEnabled to false. For important details about IPv6 addressing in a VCN, see IPv6 Addresses.  Example: true
	// +kubebuilder:validation:Optional
	IsIpv6Enabled *bool `json:"isIpv6Enabled,omitempty" tf:"is_ipv6enabled,omitempty"`

	// Specifies whether to skip Oracle allocated IPv6 GUA. By default, Oracle will allocate one GUA of /56 size for an IPv6 enabled VCN.
	// +kubebuilder:validation:Optional
	IsOracleGuaAllocationEnabled *bool `json:"isOracleGuaAllocationEnabled,omitempty" tf:"is_oracle_gua_allocation_enabled,omitempty"`

	// (Updatable) Security Attributes for this resource. This is unique to ZPR, and helps identify which resources are allowed to be accessed by what permission controls.  Example: {"Oracle-DataSecurity-ZPR.MaxEgressCount.value": "42", "Oracle-DataSecurity-ZPR.MaxEgressCount.mode": "audit"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	SecurityAttributes map[string]*string `json:"securityAttributes,omitempty" tf:"security_attributes,omitempty"`
}

// VcnSpec defines the desired state of Vcn
type VcnSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VcnParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider VcnInitParameters `json:"initProvider,omitempty"`
}

// VcnStatus defines the observed state of Vcn.
type VcnStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VcnObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Vcn is the Schema for the Vcns API. Provides the Vcn resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type Vcn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.compartmentId) || (has(self.initProvider) && has(self.initProvider.compartmentId))",message="spec.forProvider.compartmentId is a required parameter"
	Spec   VcnSpec   `json:"spec"`
	Status VcnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VcnList contains a list of Vcns
type VcnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vcn `json:"items"`
}

// Repository type metadata.
var (
	Vcn_Kind             = "Vcn"
	Vcn_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vcn_Kind}.String()
	Vcn_KindAPIVersion   = Vcn_Kind + "." + CRDGroupVersion.String()
	Vcn_GroupVersionKind = CRDGroupVersion.WithKind(Vcn_Kind)
)

func init() {
	SchemeBuilder.Register(&Vcn{}, &VcnList{})
}