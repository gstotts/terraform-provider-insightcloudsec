package insightcloudsec

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestCloudDataSource_NameLookup(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.insightcloudsec_cloud.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testPreCheck(t) },
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testICSCloudConfigBasic(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "test123_cloud"),
				),
			},
		},
	})

}

func testICSCloudConfigBasic(randomstr string) string {
	return fmt.Sprintf(`
data "insightcloudsec_cloud" "%[1]s" {
	name = "test123_cloud"
}
`, randomstr)
}
