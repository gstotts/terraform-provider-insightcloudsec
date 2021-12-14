package insightcloudsec

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ProviderNameInsightCloudSec = "insightcloudsec"
)

var (
	testProvider  *schema.Provider
	testProviders map[string]*schema.Provider
)

func init() {
	testProvider = Provider()
	testProviders = map[string]*schema.Provider{
		ProviderNameInsightCloudSec: testProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
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
