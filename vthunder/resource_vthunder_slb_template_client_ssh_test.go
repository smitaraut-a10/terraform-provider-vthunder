package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_TEMPLATE_CLIENT_SSH_RESOURCE = `
resource "vthunder_TemplateClientSSH" "testname" {
	name = "testssh"
	user_tag = "test_tag"
	forward_proxy_enable = 1
	forward_proxy_hostkey = "test"
	passphrase = "testing123"
}
`

//Acceptance test
func TestTemplateClientSSH_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_CLIENT_SSH_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_TemplateClientSSH.testname", "name", "testssh"),
					resource.TestCheckResourceAttr("vthunder_TemplateClientSSH.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_TemplateClientSSH.testname", "forward_proxy_enable", 1),
					resource.TestCheckResourceAttr("vthunder_TemplateClientSSH.testname", "forward_proxy_hostkey", "test"),
					resource.TestCheckResourceAttr("vthunder_TemplateClientSSH.testname", "passphrase", "testing123"),
				),
			},
		},
	})
}
