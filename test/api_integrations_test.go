/*
ConfigCat Public Management API

Testing IntegrationsAPIService

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

func Test_configcatpublicapi_IntegrationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IntegrationsAPIService CreateIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var productId string

		resp, httpRes, err := apiClient.IntegrationsAPI.CreateIntegration(context.Background(), productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IntegrationsAPIService DeleteIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var integrationId string

		httpRes, err := apiClient.IntegrationsAPI.DeleteIntegration(context.Background(), integrationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IntegrationsAPIService GetIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var integrationId string

		resp, httpRes, err := apiClient.IntegrationsAPI.GetIntegration(context.Background(), integrationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IntegrationsAPIService GetIntegrations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var productId string

		resp, httpRes, err := apiClient.IntegrationsAPI.GetIntegrations(context.Background(), productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IntegrationsAPIService UpdateIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var integrationId string

		resp, httpRes, err := apiClient.IntegrationsAPI.UpdateIntegration(context.Background(), integrationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
