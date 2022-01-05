package insightcloudsec

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCloud(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.insightcloudsec_cloud.%s", rnd)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Requires a "Test Cloud" be set in the instance used for testing
			{
				Config: testAccDataSourceCloudConfig(rnd),
				Check: resource.ComposeTestCheckFunc(
					testResourceID("Cloud", name),
				),
			},
		},
	})
}

func testAccDataSourceCloudConfig(dataSourceName string) string {
	return fmt.Sprintf(`
data "insightcloudsec_cloud" "%[1]s" {
	name = "Test Cloud"
}`, dataSourceName)
}
