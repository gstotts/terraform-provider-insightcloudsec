package insightcloudsec

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceUsers(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.insightcloudsec_users.%s", rnd)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUsersConfig(rnd),
				Check: resource.ComposeTestCheckFunc(
					testResourceID("Users", name),
				),
			},
		},
	})
}

func testAccDataSourceUsersConfig(name string) string {
	return fmt.Sprintf(`data "insightcloudsec_users" "%[1]s" {}`, name)
}
