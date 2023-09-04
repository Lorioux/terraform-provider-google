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

package monitoring_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccMonitoringMetricDescriptor_monitoringMetricDescriptorBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMonitoringMetricDescriptorDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringMetricDescriptor_monitoringMetricDescriptorBasicExample(context),
			},
			{
				ResourceName:            "google_monitoring_metric_descriptor.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata", "launch_stage"},
			},
		},
	})
}

func testAccMonitoringMetricDescriptor_monitoringMetricDescriptorBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_monitoring_metric_descriptor" "basic" {
  description = "Daily sales records from all branch stores."
  display_name = "tf-test-metric-descriptor%{random_suffix}"
  type = "custom.googleapis.com/stores/tf_test_daily_sales%{random_suffix}"
  metric_kind = "GAUGE"
  value_type = "DOUBLE"
  unit = "{USD}"
  labels {
      key = "store_id"
      value_type = "STRING"
      description = "The ID of the store."
  }
  launch_stage = "BETA"
  metadata {
    sample_period = "60s"
    ingest_delay = "30s"
  }
}
`, context)
}

func TestAccMonitoringMetricDescriptor_monitoringMetricDescriptorAlertExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMonitoringMetricDescriptorDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringMetricDescriptor_monitoringMetricDescriptorAlertExample(context),
			},
			{
				ResourceName:            "google_monitoring_metric_descriptor.with_alert",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata", "launch_stage"},
			},
		},
	})
}

func testAccMonitoringMetricDescriptor_monitoringMetricDescriptorAlertExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_monitoring_metric_descriptor" "with_alert" {
  description = "Daily sales records from all branch stores."
  display_name = "tf-test-metric-descriptor%{random_suffix}"
  type = "custom.googleapis.com/stores/tf_test_daily_sales%{random_suffix}"
  metric_kind = "GAUGE"
  value_type = "DOUBLE"
  unit = "{USD}"
}

resource "google_monitoring_alert_policy" "alert_policy" {
  display_name = "tf-test-metric-descriptor%{random_suffix}"
  combiner     = "OR"
  conditions {
    display_name = "test condition"
    condition_threshold {
      filter     = "metric.type=\"${google_monitoring_metric_descriptor.with_alert.type}\" AND resource.type=\"gce_instance\""
      duration   = "60s"
      comparison = "COMPARISON_GT"
    }
  }
}
`, context)
}

func testAccCheckMonitoringMetricDescriptorDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_monitoring_metric_descriptor" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{MonitoringBasePath}}v3/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:               config,
				Method:               "GET",
				Project:              billingProject,
				RawURL:               url,
				UserAgent:            config.UserAgent,
				ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsMonitoringConcurrentEditError},
			})
			if err == nil {
				return fmt.Errorf("MonitoringMetricDescriptor still exists at %s", url)
			}
		}

		return nil
	}
}
