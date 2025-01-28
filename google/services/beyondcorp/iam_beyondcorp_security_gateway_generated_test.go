// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package beyondcorp_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
)

func TestAccBeyondcorpSecurityGatewayIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpSecurityGatewayIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_beyondcorp_security_gateway_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s roles/beyondcorp.securityGatewayUser", envvar.GetTestProjectFromEnv(), "global", fmt.Sprintf("default%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccBeyondcorpSecurityGatewayIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_beyondcorp_security_gateway_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s roles/beyondcorp.securityGatewayUser", envvar.GetTestProjectFromEnv(), "global", fmt.Sprintf("default%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBeyondcorpSecurityGatewayIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccBeyondcorpSecurityGatewayIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_beyondcorp_security_gateway_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s roles/beyondcorp.securityGatewayUser user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), "global", fmt.Sprintf("default%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBeyondcorpSecurityGatewayIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpSecurityGatewayIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_beyondcorp_security_gateway_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_beyondcorp_security_gateway_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s", envvar.GetTestProjectFromEnv(), "global", fmt.Sprintf("default%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBeyondcorpSecurityGatewayIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_beyondcorp_security_gateway_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s", envvar.GetTestProjectFromEnv(), "global", fmt.Sprintf("default%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBeyondcorpSecurityGatewayIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpSecurityGatewayIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_beyondcorp_security_gateway_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s roles/beyondcorp.securityGatewayUser %s", envvar.GetTestProjectFromEnv(), "global", fmt.Sprintf("default%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBeyondcorpSecurityGatewayIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpSecurityGatewayIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_beyondcorp_security_gateway_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s roles/beyondcorp.securityGatewayUser", envvar.GetTestProjectFromEnv(), "global", fmt.Sprintf("default%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_beyondcorp_security_gateway_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s roles/beyondcorp.securityGatewayUser %s", envvar.GetTestProjectFromEnv(), "global", fmt.Sprintf("default%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_beyondcorp_security_gateway_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s roles/beyondcorp.securityGatewayUser %s", envvar.GetTestProjectFromEnv(), "global", fmt.Sprintf("default%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBeyondcorpSecurityGatewayIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpSecurityGatewayIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_beyondcorp_security_gateway_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s roles/beyondcorp.securityGatewayUser user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), "global", fmt.Sprintf("default%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBeyondcorpSecurityGatewayIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpSecurityGatewayIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_beyondcorp_security_gateway_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s roles/beyondcorp.securityGatewayUser user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), "global", fmt.Sprintf("default%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_beyondcorp_security_gateway_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s roles/beyondcorp.securityGatewayUser user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), "global", fmt.Sprintf("default%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_beyondcorp_security_gateway_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s roles/beyondcorp.securityGatewayUser user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), "global", fmt.Sprintf("default%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBeyondcorpSecurityGatewayIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/beyondcorp.securityGatewayUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := acctest.Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpSecurityGatewayIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO(SarahFrench) - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_beyondcorp_security_gateway_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", tpgresource.CheckGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_beyondcorp_security_gateway_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s", envvar.GetTestProjectFromEnv(), "global", fmt.Sprintf("default%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBeyondcorpSecurityGatewayIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "example" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_security_gateway_iam_member" "foo" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccBeyondcorpSecurityGatewayIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "example" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_beyondcorp_security_gateway_iam_policy" "foo" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_beyondcorp_security_gateway_iam_policy" "foo" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  depends_on = [
    google_beyondcorp_security_gateway_iam_policy.foo
  ]
}
`, context)
}

func testAccBeyondcorpSecurityGatewayIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "example" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

data "google_iam_policy" "foo" {
}

resource "google_beyondcorp_security_gateway_iam_policy" "foo" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccBeyondcorpSecurityGatewayIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "example" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_security_gateway_iam_binding" "foo" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccBeyondcorpSecurityGatewayIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "example" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_security_gateway_iam_binding" "foo" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccBeyondcorpSecurityGatewayIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "example" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_security_gateway_iam_binding" "foo" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccBeyondcorpSecurityGatewayIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "example" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_security_gateway_iam_binding" "foo" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_beyondcorp_security_gateway_iam_binding" "foo2" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_beyondcorp_security_gateway_iam_binding" "foo3" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccBeyondcorpSecurityGatewayIamMember_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "example" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_security_gateway_iam_member" "foo" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccBeyondcorpSecurityGatewayIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "example" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_security_gateway_iam_member" "foo" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_beyondcorp_security_gateway_iam_member" "foo2" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_beyondcorp_security_gateway_iam_member" "foo3" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccBeyondcorpSecurityGatewayIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "example" {
  security_gateway_id = "default%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      # Check that lack of description doesn't cause any issues
      # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
      title       = "%{condition_title_no_desc}"
      expression  = "%{condition_expr_no_desc}"
    }
  }
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "%{condition_desc}"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_beyondcorp_security_gateway_iam_policy" "foo" {
  project = google_beyondcorp_security_gateway.example.project
  location = google_beyondcorp_security_gateway.example.location
  security_gateway_id = google_beyondcorp_security_gateway.example.security_gateway_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
