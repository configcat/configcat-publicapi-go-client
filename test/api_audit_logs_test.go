/*
ConfigCat Public Management API

Testing AuditLogsApiService

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

func Test_configcatpublicapi_AuditLogsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuditLogsApiService GetAuditlogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var productId string

		resp, httpRes, err := apiClient.AuditLogsApi.GetAuditlogs(context.Background(), productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditLogsApiService GetDeletedSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var configId string

		resp, httpRes, err := apiClient.AuditLogsApi.GetDeletedSettings(context.Background(), configId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditLogsApiService GetOrganizationAuditlogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.AuditLogsApi.GetOrganizationAuditlogs(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
