// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviImage() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviImageRead,
		Schema: map[string]*schema.Schema{
			"controller_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePackageDetailsSchema(),
			},
			"controller_patch_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"controller_patch_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"migrations": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSupportedMigrationsSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePackageDetailsSchema(),
			},
			"se_patch_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_patch_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uber_bundle": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
