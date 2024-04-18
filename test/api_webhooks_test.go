/*
ConfigCat Public Management API

Testing WebhooksApiService

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

func Test_configcatpublicapi_WebhooksApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WebhooksApiService CreateWebhook", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var configId string
		var environmentId string

		resp, httpRes, err := apiClient.WebhooksApi.CreateWebhook(context.Background(), configId, environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService DeleteWebhook", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var webhookId int32

		httpRes, err := apiClient.WebhooksApi.DeleteWebhook(context.Background(), webhookId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetWebhook", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var webhookId int32

		resp, httpRes, err := apiClient.WebhooksApi.GetWebhook(context.Background(), webhookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetWebhookSigningKeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var webhookId int32

		resp, httpRes, err := apiClient.WebhooksApi.GetWebhookSigningKeys(context.Background(), webhookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetWebhooks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var productId string

		resp, httpRes, err := apiClient.WebhooksApi.GetWebhooks(context.Background(), productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService UpdateWebhook", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var webhookId int32

		resp, httpRes, err := apiClient.WebhooksApi.UpdateWebhook(context.Background(), webhookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}