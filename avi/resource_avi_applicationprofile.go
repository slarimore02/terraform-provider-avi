// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceApplicationProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_config_cksum": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dns_service_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceDnsServiceApplicationProfileSchema(),
		},
		"dos_rl_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceDosRateLimitProfileSchema(),
		},
		"http_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHTTPApplicationProfileSchema(),
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"preserve_client_ip": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"preserve_client_port": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"preserve_dest_ip_port": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"sip_service_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSipServiceApplicationProfileSchema(),
		},
		"tcp_app_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTCPApplicationProfileSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviApplicationProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviApplicationProfileCreate,
		Read:   ResourceAviApplicationProfileRead,
		Update: resourceAviApplicationProfileUpdate,
		Delete: resourceAviApplicationProfileDelete,
		Schema: ResourceApplicationProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceApplicationProfileImporter,
		},
	}
}

func ResourceApplicationProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceApplicationProfileSchema()
	return ResourceImporter(d, m, "applicationprofile", s)
}

func ResourceAviApplicationProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationProfileSchema()
	err := APIRead(d, meta, "applicationprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviApplicationProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationProfileSchema()
	err := APICreateOrUpdate(d, meta, "applicationprofile", s)
	if err == nil {
		err = ResourceAviApplicationProfileRead(d, meta)
	}
	return err
}

func resourceAviApplicationProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "applicationprofile", s)
	if err == nil {
		err = ResourceAviApplicationProfileRead(d, meta)
	}
	return err
}

func resourceAviApplicationProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "applicationprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviApplicationProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
