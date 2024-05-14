package google

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const LogBucketAssetType string = "logging.googleapis.com/LogBucket"

func resourceLogBucket() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: LogBucketAssetType,
		Convert:   GetLogBucketCaiObject,
	}
}

func GetLogBucketCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//logging.googleapis.com/projects/{{project}}/locations/{{location}}/buckets/{{bucket}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetLogBucketApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: LogBucketAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/logging/v2/rest",
				DiscoveryName:        "LogBucket",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetLogBucketApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandLogBucketName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	analyticsEnabledProp, err := expandandAnalyticsEnabled(d.Get("analytics_enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("analytics_enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(analyticsEnabledProp)) && (ok || !reflect.DeepEqual(v, analyticsEnabledProp)) {
		obj["analyticsEnabled"] = analyticsEnabledProp
	}

	return obj, nil
}

func expandLogBucketName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandandAnalyticsEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}