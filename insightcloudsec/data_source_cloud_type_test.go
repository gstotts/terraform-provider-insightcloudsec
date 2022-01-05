package insightcloudsec

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCloudTypes(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.insightcloudsec_cloud_types.%s", rnd)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCloudTypesConfig(rnd),
				Check: resource.ComposeTestCheckFunc(
					testResourceID("Cloud Types", name),
					resource.TestCheckResourceAttr(name, "clouds.#", "10"),
				),
			},
		},
	})
}

func testAccDataSourceCloudTypesConfig(name string) string {
	return fmt.Sprintf(`data "insightcloudsec_cloud_types" "%[1]s" {}`, name)
}
