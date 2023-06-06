/*
ConfigCat Public Management API

Testing FeatureFlagSettingValuesUsingSDKKeyApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package configcatpublicapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/configcat/configcat-publicapi-go-client"
)

func Test_configcatpublicapi_FeatureFlagSettingValuesUsingSDKKeyApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FeatureFlagSettingValuesUsingSDKKeyApiService GetSettingValueBySdkkey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var settingKeyOrId string

		resp, httpRes, err := apiClient.FeatureFlagSettingValuesUsingSDKKeyApi.GetSettingValueBySdkkey(context.Background(), settingKeyOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeatureFlagSettingValuesUsingSDKKeyApiService ReplaceSettingValueBySdkkey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var settingKeyOrId string

		resp, httpRes, err := apiClient.FeatureFlagSettingValuesUsingSDKKeyApi.ReplaceSettingValueBySdkkey(context.Background(), settingKeyOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeatureFlagSettingValuesUsingSDKKeyApiService UpdateSettingValueBySdkkey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var settingKeyOrId string

		resp, httpRes, err := apiClient.FeatureFlagSettingValuesUsingSDKKeyApi.UpdateSettingValueBySdkkey(context.Background(), settingKeyOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}