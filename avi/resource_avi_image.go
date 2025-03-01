// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceImageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_info_values": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceImageCloudDataSchema(),
		},
		"controller_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcePackageDetailsSchema(),
		},
		"controller_patch_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"controller_patch_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"events": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceImageEventMapSchema(),
		},
		"img_state": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceImageUploadOpsStatusSchema(),
		},
		"migrations": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSupportedMigrationsSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"progress": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"se_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcePackageDetailsSchema(),
		},
		"se_patch_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"se_patch_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tasks_completed": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"total_tasks": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uber_bundle": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviImageCreate,
		Read:   ResourceAviImageRead,
		Update: resourceAviImageUpdate,
		Delete: resourceAviImageDelete,
		Schema: ResourceImageSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceImageImporter,
		},
	}
}

func ResourceImageImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceImageSchema()
	return ResourceImporter(d, m, "image", s)
}

func ResourceAviImageRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceImageSchema()
	err := APIRead(d, meta, "image", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviImageCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceImageSchema()
	err := APICreateOrUpdate(d, meta, "image", s)
	if err == nil {
		err = ResourceAviImageRead(d, meta)
	}
	return err
}

func resourceAviImageUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceImageSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "image", s)
	if err == nil {
		err = ResourceAviImageRead(d, meta)
	}
	return err
}

func resourceAviImageDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "image"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviImageDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
