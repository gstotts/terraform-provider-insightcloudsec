package insightcloudsec

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
					testAccDataSourceCloudTypesId(name),
					resource.TestCheckResourceAttr(name, "clouds.#", "10"),
				),
			},
		},
	})
}

func testAccDataSourceCloudTypesConfig(name string) string {
	return fmt.Sprintf(`data "insightcloudsec_cloud_types" "%[1]s" {}`, name)
}

func testAccDataSourceCloudTypesId(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		all := s.RootModule().Resources
		rs, ok := all[n]

		if !ok {
			return fmt.Errorf("can't find Cloud Types data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Snapshot Cloud Types source ID not set")
		}
		return nil
	}
}
