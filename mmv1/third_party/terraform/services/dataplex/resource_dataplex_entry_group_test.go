package dataplex_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccDataplexEntryGroup_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataplexEntryGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexEntryGroup_full(context),
			},
			{
				ResourceName:            "google_dataplex_entry_group.test_entry_group_full",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "entry_group_id", "labels", "terraform_labels"},
			},
			{
				Config: testAccDataplexEntryGroup_update(context),
			},
			{
				ResourceName:            "google_dataplex_entry_group.test_entry_group_basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "entry_group_id", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccDataplexEntryGroup_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_entry_group" "test_entry_group_full" {
  entry_group_id = "tf-test-entry-group-full%{random_suffix}"
  project = "%{project_name}"
  location = "us-central1"

  labels = { "tag": "test-tf" }
  display_name = "terraform entry group"
  description = "entry group created by Terraform"
}
`, context)
}

func testAccDataplexEntryGroup_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_entry_group" "test_entry_group_basic" {
  entry_group_id = "tf-test-entry-group-basic%{random_suffix}"
  project = "%{project_name}"
  location = "us-central1"
}
`, context)
}
