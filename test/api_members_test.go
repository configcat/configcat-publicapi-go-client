/*
ConfigCat Public Management API

Testing MembersApiService

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

func Test_configcatpublicapi_MembersApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MembersApiService AddMemberToGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var userId string

		httpRes, err := apiClient.MembersApi.AddMemberToGroup(context.Background(), organizationId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MembersApiService DeleteOrganizationMember", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var userId string

		httpRes, err := apiClient.MembersApi.DeleteOrganizationMember(context.Background(), organizationId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MembersApiService DeleteProductMember", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var productId string
		var userId string

		httpRes, err := apiClient.MembersApi.DeleteProductMember(context.Background(), productId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MembersApiService GetOrganizationMembers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.MembersApi.GetOrganizationMembers(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MembersApiService GetProductMembers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var productId string

		resp, httpRes, err := apiClient.MembersApi.GetProductMembers(context.Background(), productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MembersApiService InviteMember", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var productId string

		httpRes, err := apiClient.MembersApi.InviteMember(context.Background(), productId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
