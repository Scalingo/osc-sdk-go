/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// Nic Information about the NIC.
type Nic struct {
	// The account ID of the owner of the NIC.
	AccountId string `json:"AccountId,omitempty"`
	// The description of the NIC.
	Description string `json:"Description,omitempty"`
	// (Net only) If `true`, the source/destination check is enabled. If `false`, it is disabled. This value must be `false` for a NAT VM to perform network address translation (NAT) in a Net.
	IsSourceDestChecked bool `json:"IsSourceDestChecked,omitempty"`
	LinkNic LinkNic `json:"LinkNic,omitempty"`
	LinkPublicIp LinkPublicIp `json:"LinkPublicIp,omitempty"`
	// The Media Access Control (MAC) address of the NIC.
	MacAddress string `json:"MacAddress,omitempty"`
	// The ID of the Net for the NIC.
	NetId string `json:"NetId,omitempty"`
	// The ID of the NIC.
	NicId string `json:"NicId,omitempty"`
	// The name of the private DNS.
	PrivateDnsName string `json:"PrivateDnsName,omitempty"`
	// The private IP addresses of the NIC.
	PrivateIps []PrivateIp `json:"PrivateIps,omitempty"`
	// One or more IDs of security groups for the NIC.
	SecurityGroups []SecurityGroupLight `json:"SecurityGroups,omitempty"`
	// The state of the NIC (`available` \\| `attaching` \\| `in-use` \\| `detaching`).
	State string `json:"State,omitempty"`
	// The ID of the Subnet.
	SubnetId string `json:"SubnetId,omitempty"`
	// The Subregion in which the NIC is located.
	SubregionName string `json:"SubregionName,omitempty"`
	// One or more tags associated with the NIC.
	Tags []ResourceTag `json:"Tags,omitempty"`
}