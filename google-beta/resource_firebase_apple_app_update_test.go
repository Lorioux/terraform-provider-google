package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFirebaseAppleApp_update(t *testing.T) {
	t.Parallel()
	context := map[string]interface{}{
		"project_id":    acctest.GetTestProjectFromEnv(),
		"bundle_id":     "apple.app.12345",
		"random_suffix": RandString(t, 10),
		"display_name":  "tf-test Display Name N",
	}
	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAppleApp(context, 12345, "1"),
			},
			{
				Config: testAccFirebaseAppleApp(context, 67890, "2"),
			},
		},
	})
}

func testAccFirebaseAppleApp(context map[string]interface{}, appStoreId int, delta string) string {
	context["display_name"] = context["display_name"].(string) + delta
	context["app_store_id"] = appStoreId
	context["team_id"] = "123456789" + delta
	return Nprintf(`
resource "google_firebase_apple_app" "update" {
        provider = google-beta
        project = "%{project_id}"
        bundle_id = "%{bundle_id}"
        display_name = "%{display_name} %{random_suffix}"
        app_store_id = "%{app_store_id}"
        team_id = "%{team_id}"
}
`, context)
}
