package insightcloudsec

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	ProviderNameInsightCloudSec = "insightcloudsec"
)

var (
	testAccProvider  *schema.Provider
	testAccProviders map[string]*schema.Provider
)

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		ProviderNameInsightCloudSec: testAccProvider,
	}
}

func TestInsightCloudSec_Provider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestInsightCloudSec_Provider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testPreCheck(t *testing.T) {
	testPreCheckBaseUrl(t)
	testPreCheckApiKey(t)
}

func testPreCheckBaseUrl(t *testing.T) {
	if v := os.Getenv("INSIGHTCLOUDSEC_BASE_URL"); v == "" {
		t.Fatal("INSIGHTCLOUDSEC_BASE_URL must be set for acceptance testing")
	}
}

func testPreCheckApiKey(t *testing.T) {
	if v := os.Getenv("INSIGHTCLOUDSEC_API_KEY"); v == "" {
		t.Fatal("INSIGHTCLOUDSEC_API_KEY must be set for acceptance testing")
	}
}

func generateRandomResourceName() string {
	return acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
}

func testDataSourceID(r string, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		all := s.RootModule().Resources
		rs, ok := all[n]

		if !ok {
			return fmt.Errorf("can't find %s data source: %s", r, n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("%s source ID not set", r)
		}
		return nil
	}
}
