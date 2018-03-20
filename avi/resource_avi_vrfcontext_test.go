package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIVrfContextBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIVrfContextDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIVrfContextConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIVrfContextExists("avi_vrfcontext.testVrfContext"),
					resource.TestCheckResourceAttr(
						"avi_vrfcontext.testVrfContext", "name", "testglobal")),
			},
			{
				Config: testAccAVIVrfContextupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIVrfContextExists("avi_vrfcontext.testVrfContext"),
					resource.TestCheckResourceAttr(
						"avi_vrfcontext.testVrfContext", "name", "testglobal-abc")),
			},
		},
	})

}

func testAccCheckAVIVrfContextExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI VrfContext ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIVrfContextDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_vrfcontext" {
			continue
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI VrfContext still exists")
		}
	}
	return nil
}

const testAccAVIVrfContextConfig = `
data "avi_tenant" "default_tenant"{
        name= "admin"
}
data "avi_cloud" "default_cloud" {
        name= "Default-Cloud"
}
resource "avi_vrfcontext" "testVrfContext" {
"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
"cloud_ref" = "${data.avi_cloud.default_cloud.id}"
"name" = "testglobal"
}
`

const testAccAVIVrfContextupdatedConfig = `
data "avi_tenant" "default_tenant"{
        name= "admin"
}
data "avi_cloud" "default_cloud" {
        name= "Default-Cloud"
}
resource "avi_vrfcontext" "testVrfContext" {
"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
"cloud_ref" = "${data.avi_cloud.default_cloud.id}"
"name" = "testglobal-abc"
}
`
