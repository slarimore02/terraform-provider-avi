package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceIpamDnsProviderProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSIpamDnsProviderProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_ipamdnsproviderprofile.testIpamDnsProviderProfile", "name", "test-ipam-abc-abc"),
					resource.TestCheckResourceAttr(
						"avi_ipamdnsproviderprofile.testIpamDnsProviderProfile", "allocate_ip_in_vrf", "false"),
				),
			},
		},
	})

}

const testAccAVIDSIpamDnsProviderProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_ipamdnsproviderprofile" "testIpamDnsProviderProfile" {
	name = "test-ipam-abc-abc"
	allocate_ip_in_vrf = false
	internal_profile {
		ttl = "31"
	}
	type = "IPAMDNS_TYPE_INTERNAL"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_ipamdnsproviderprofile" "testIpamDnsProviderProfile" {
    name= "${avi_ipamdnsproviderprofile.testIpamDnsProviderProfile.name}"
}
`
