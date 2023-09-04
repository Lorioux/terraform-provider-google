// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceComputeNetworkAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNetworkAttachmentCreate,
		Read:   resourceComputeNetworkAttachmentRead,
		Delete: resourceComputeNetworkAttachmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeNetworkAttachmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"connection_preference": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ACCEPT_AUTOMATIC", "ACCEPT_MANUAL", "INVALID"}),
				Description:  `The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules. Possible values: ["ACCEPT_AUTOMATIC", "ACCEPT_MANUAL", "INVALID"]`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `URL of the region where the network attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.`,
			},
			"subnetworks": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `An array of URLs where each entry is the URL of a subnet provided by the service consumer to use for endpoints in the producers that connect to this network attachment.`,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource. Provide this property when you create the resource.`,
			},
			"producer_accept_lists": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Projects that are allowed to connect to this network attachment. The project can be specified using its id or number.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"producer_reject_lists": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Projects that are not allowed to connect to this network attachment. The project can be specified using its id or number.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"connection_endpoints": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `An array of connections for all the producers connected to this network attachment.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The IPv4 address assigned to the producer instance network interface. This value will be a range in case of Serverless.`,
						},
						"project_id_or_num": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The project id or number of the interface to which the IP was assigned.`,
						},
						"secondary_ip_cidr_ranges": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Alias IP ranges from the same subnetwork.`,
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The status of a connected endpoint to this network attachment.`,
						},
						"subnetwork": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The subnetwork used to assign the IP to the producer instance network interface.`,
						},
					},
				},
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Fingerprint of this resource. A hash of the contents stored in this object. This
field is used in optimistic locking. An up-to-date fingerprint must be provided in order to patch.`,
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The unique identifier for the resource type. The server generates this identifier.`,
			},
			"kind": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Type of the resource.`,
			},
			"network": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The URL of the network which the Network Attachment belongs to. Practically it is inferred by fetching the network of the first subnetwork associated.
Because it is required that all the subnetworks must be from the same network, it is assured that the Network Attachment belongs to the same network as all the subnetworks.`,
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL for the resource.`,
			},
			"self_link_with_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL for this resource's resource id.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeNetworkAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeNetworkAttachmentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	connectionPreferenceProp, err := expandComputeNetworkAttachmentConnectionPreference(d.Get("connection_preference"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("connection_preference"); !tpgresource.IsEmptyValue(reflect.ValueOf(connectionPreferenceProp)) && (ok || !reflect.DeepEqual(v, connectionPreferenceProp)) {
		obj["connectionPreference"] = connectionPreferenceProp
	}
	subnetworksProp, err := expandComputeNetworkAttachmentSubnetworks(d.Get("subnetworks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("subnetworks"); !tpgresource.IsEmptyValue(reflect.ValueOf(subnetworksProp)) && (ok || !reflect.DeepEqual(v, subnetworksProp)) {
		obj["subnetworks"] = subnetworksProp
	}
	producerRejectListsProp, err := expandComputeNetworkAttachmentProducerRejectLists(d.Get("producer_reject_lists"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("producer_reject_lists"); !tpgresource.IsEmptyValue(reflect.ValueOf(producerRejectListsProp)) && (ok || !reflect.DeepEqual(v, producerRejectListsProp)) {
		obj["producerRejectLists"] = producerRejectListsProp
	}
	producerAcceptListsProp, err := expandComputeNetworkAttachmentProducerAcceptLists(d.Get("producer_accept_lists"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("producer_accept_lists"); !tpgresource.IsEmptyValue(reflect.ValueOf(producerAcceptListsProp)) && (ok || !reflect.DeepEqual(v, producerAcceptListsProp)) {
		obj["producerAcceptLists"] = producerAcceptListsProp
	}
	fingerprintProp, err := expandComputeNetworkAttachmentFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	nameProp, err := expandComputeNetworkAttachmentName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	regionProp, err := expandComputeNetworkAttachmentRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkAttachments")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NetworkAttachment: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkAttachment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating NetworkAttachment: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/networkAttachments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating NetworkAttachment", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create NetworkAttachment: %s", err)
	}

	log.Printf("[DEBUG] Finished creating NetworkAttachment %q: %#v", d.Id(), res)

	return resourceComputeNetworkAttachmentRead(d, meta)
}

func resourceComputeNetworkAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkAttachments/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkAttachment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeNetworkAttachment %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}

	if err := d.Set("kind", flattenComputeNetworkAttachmentKind(res["kind"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}
	if err := d.Set("id", flattenComputeNetworkAttachmentId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeNetworkAttachmentCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}
	if err := d.Set("description", flattenComputeNetworkAttachmentDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}
	if err := d.Set("self_link", flattenComputeNetworkAttachmentSelfLink(res["selfLink"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}
	if err := d.Set("self_link_with_id", flattenComputeNetworkAttachmentSelfLinkWithId(res["selfLinkWithId"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}
	if err := d.Set("connection_preference", flattenComputeNetworkAttachmentConnectionPreference(res["connectionPreference"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}
	if err := d.Set("connection_endpoints", flattenComputeNetworkAttachmentConnectionEndpoints(res["connectionEndpoints"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}
	if err := d.Set("subnetworks", flattenComputeNetworkAttachmentSubnetworks(res["subnetworks"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}
	if err := d.Set("producer_reject_lists", flattenComputeNetworkAttachmentProducerRejectLists(res["producerRejectLists"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}
	if err := d.Set("producer_accept_lists", flattenComputeNetworkAttachmentProducerAcceptLists(res["producerAcceptLists"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}
	if err := d.Set("fingerprint", flattenComputeNetworkAttachmentFingerprint(res["fingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}
	if err := d.Set("network", flattenComputeNetworkAttachmentNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}
	if err := d.Set("name", flattenComputeNetworkAttachmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}
	if err := d.Set("region", flattenComputeNetworkAttachmentRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkAttachment: %s", err)
	}

	return nil
}

func resourceComputeNetworkAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkAttachment: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkAttachments/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting NetworkAttachment %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "NetworkAttachment")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting NetworkAttachment", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting NetworkAttachment %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeNetworkAttachmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/networkAttachments/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/networkAttachments/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeNetworkAttachmentKind(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentSelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentSelfLinkWithId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentConnectionPreference(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentConnectionEndpoints(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"status":                   flattenComputeNetworkAttachmentConnectionEndpointsStatus(original["status"], d, config),
			"project_id_or_num":        flattenComputeNetworkAttachmentConnectionEndpointsProjectIdOrNum(original["projectIdOrNum"], d, config),
			"subnetwork":               flattenComputeNetworkAttachmentConnectionEndpointsSubnetwork(original["subnetwork"], d, config),
			"ip_address":               flattenComputeNetworkAttachmentConnectionEndpointsIpAddress(original["ipAddress"], d, config),
			"secondary_ip_cidr_ranges": flattenComputeNetworkAttachmentConnectionEndpointsSecondaryIpCidrRanges(original["secondaryIpCidrRanges"], d, config),
		})
	}
	return transformed
}
func flattenComputeNetworkAttachmentConnectionEndpointsStatus(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentConnectionEndpointsProjectIdOrNum(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentConnectionEndpointsSubnetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentConnectionEndpointsIpAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentConnectionEndpointsSecondaryIpCidrRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentSubnetworks(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)
}

func flattenComputeNetworkAttachmentProducerRejectLists(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentProducerAcceptLists(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentFingerprint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAttachmentRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func expandComputeNetworkAttachmentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAttachmentConnectionPreference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAttachmentSubnetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for subnetworks: nil")
		}
		f, err := tpgresource.ParseRegionalFieldValue("subnetworks", raw.(string), "project", "region", "zone", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for subnetworks: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeNetworkAttachmentProducerRejectLists(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAttachmentProducerAcceptLists(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAttachmentFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAttachmentName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAttachmentRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
