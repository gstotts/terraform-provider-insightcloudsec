package insightcloudsec

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccInsightCloudSec_DataSource_Cloud(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.insightcloudsec_cloud.%s", rnd)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Requires a "Test Cloud" be set in the instance used for testing
			{
				Config: testAccInsightCloudSec_DataSource_CloudConfig(rnd),
				Check: resource.ComposeTestCheckFunc(
					testDataSourceID("Cloud", name),
				),
			},
		},
	})
}

func testAccInsightCloudSec_DataSource_CloudConfig(dataSourceName string) string {
	return fmt.Sprintf(`
data "insightcloudsec_cloud" "%[1]s" {
	name = "Test Cloud"
}`, dataSourceName)
}
