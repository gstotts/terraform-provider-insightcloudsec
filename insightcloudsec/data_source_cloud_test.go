package insightcloudsec

package insightcloudsec

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceCloud(t *testing.T) {
	rnd := generateRandomResourceName()
	name := 
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCloudConfig(rnd),
				Check: resource.ComposeTestCheckFunc(
					,
				),
			},
		},
	})
}

func testAccDataSourceCloudConfig(name string) string {
	return fmt.Sprintf(`data "insightcloudsec_cloud" "%[1]s" {
		name = ""
	}`, name)
}