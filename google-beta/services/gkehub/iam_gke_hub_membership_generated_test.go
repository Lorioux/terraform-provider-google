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

package gkehub_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccGKEHubMembershipIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEHubMembershipIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_gke_hub_membership_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/memberships/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("basic%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccGKEHubMembershipIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_gke_hub_membership_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/memberships/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("basic%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGKEHubMembershipIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccGKEHubMembershipIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_gke_hub_membership_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/memberships/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("basic%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGKEHubMembershipIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEHubMembershipIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_gke_hub_membership_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_gke_hub_membership_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/memberships/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("basic%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGKEHubMembershipIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_gke_hub_membership_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/memberships/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("basic%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccGKEHubMembershipIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "basiccluster%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
}

resource "google_gke_hub_membership" "membership" {
  membership_id = "basic%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
}

resource "google_gke_hub_membership_iam_member" "foo" {
  project = google_gke_hub_membership.membership.project
  membership_id = google_gke_hub_membership.membership.membership_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccGKEHubMembershipIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "basiccluster%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
}

resource "google_gke_hub_membership" "membership" {
  membership_id = "basic%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_gke_hub_membership_iam_policy" "foo" {
  project = google_gke_hub_membership.membership.project
  membership_id = google_gke_hub_membership.membership.membership_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_gke_hub_membership_iam_policy" "foo" {
  project = google_gke_hub_membership.membership.project
  membership_id = google_gke_hub_membership.membership.membership_id
  depends_on = [
    google_gke_hub_membership_iam_policy.foo
  ]
}
`, context)
}

func testAccGKEHubMembershipIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "basiccluster%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
}

resource "google_gke_hub_membership" "membership" {
  membership_id = "basic%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
}

data "google_iam_policy" "foo" {
}

resource "google_gke_hub_membership_iam_policy" "foo" {
  project = google_gke_hub_membership.membership.project
  membership_id = google_gke_hub_membership.membership.membership_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccGKEHubMembershipIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "basiccluster%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
}

resource "google_gke_hub_membership" "membership" {
  membership_id = "basic%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
}

resource "google_gke_hub_membership_iam_binding" "foo" {
  project = google_gke_hub_membership.membership.project
  membership_id = google_gke_hub_membership.membership.membership_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccGKEHubMembershipIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "basiccluster%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
}

resource "google_gke_hub_membership" "membership" {
  membership_id = "basic%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
}

resource "google_gke_hub_membership_iam_binding" "foo" {
  project = google_gke_hub_membership.membership.project
  membership_id = google_gke_hub_membership.membership.membership_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
