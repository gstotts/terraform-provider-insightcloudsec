package insightcloudsec

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccInsightCloudSec_DataSource_CloudTypes(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.insightcloudsec_cloud_types.%s", rnd)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightCloudSec_DataSource_CloudTypesConfig(rnd),
				Check: resource.ComposeTestCheckFunc(
					testDataSourceID("Cloud Types", name),
					resource.TestCheckResourceAttr(name, "clouds.#", "10"),
				),
			},
		},
	})
}

func testAccInsightCloudSec_DataSource_CloudTypesConfig(name string) string {
	return fmt.Sprintf(`data "insightcloudsec_cloud_types" "%[1]s" {}`, name)
}
