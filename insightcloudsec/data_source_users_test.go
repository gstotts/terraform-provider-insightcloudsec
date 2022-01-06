package insightcloudsec

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccInsightCloudSec_DataSource_Users(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.insightcloudsec_users.%s", rnd)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightCloudSec_DataSource_UsersConfig(rnd),
				Check: resource.ComposeTestCheckFunc(
					testDataSourceID("Users", name),
				),
			},
		},
	})
}

func testAccInsightCloudSec_DataSource_UsersConfig(name string) string {
	return fmt.Sprintf(`data "insightcloudsec_users" "%[1]s" {}`, name)
}
