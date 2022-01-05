package insightcloudsec

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceUsers(t *testing.T) {
	rnd := generateRandomResourceName()
	name := 
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUsersConfig(rnd),
				Check: resource.ComposeTestCheckFunc(
					,
				),
			},
		},
	})
}

func testAccDataSourceUsersConfig(name string) string {
	return fmt.Sprintf(`data "insightcloudsec_users" "%[1]s" {}`, name)
}