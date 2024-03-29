/*
ConfigCat Public Management API

Testing CodeReferencesApiService

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

func Test_configcatpublicapi_CodeReferencesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CodeReferencesApiService V1CodeReferencesDeleteReportsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CodeReferencesApi.V1CodeReferencesDeleteReportsPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CodeReferencesApiService V1CodeReferencesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CodeReferencesApi.V1CodeReferencesPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CodeReferencesApiService V1ConfigsConfigIdCodeReferencesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var configId string

		resp, httpRes, err := apiClient.CodeReferencesApi.V1ConfigsConfigIdCodeReferencesGet(context.Background(), configId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CodeReferencesApiService V1SettingsSettingIdCodeReferencesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var settingId int32

		resp, httpRes, err := apiClient.CodeReferencesApi.V1SettingsSettingIdCodeReferencesGet(context.Background(), settingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
