# ConfigCat Public Management API Client for Go

**Base API URL**: https://api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://api.configcat.com/swagger).  

The purpose of this client is to access the ConfigCat Public Management API programmatically.  You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.   

The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  and JSON+HAL format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   

## OpenAPI Specification  
The complete specification is publicly available here: [swagger.json](v1/swagger.json).  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  

## Authentication 
This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).  

## Throttling and rate limits 
All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  
| Header | Description | 
| :- | :- | 
| X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | 
| X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://configcat.com](https://configcat.com)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./configcatpublicapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.configcat.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuditLogsApi* | [**GetAuditlogs**](docs/AuditLogsApi.md#getauditlogs) | **Get** /v1/products/{productId}/auditlogs | List Audit log items for Product
*AuditLogsApi* | [**GetDeletedSettings**](docs/AuditLogsApi.md#getdeletedsettings) | **Get** /v1/configs/{configId}/deleted-settings | List Deleted Settings
*AuditLogsApi* | [**GetOrganizationAuditlogs**](docs/AuditLogsApi.md#getorganizationauditlogs) | **Get** /v1/organizations/{organizationId}/auditlogs | List Audit log items for Organization
*ConfigsApi* | [**CreateConfig**](docs/ConfigsApi.md#createconfig) | **Post** /v1/products/{productId}/configs | Create Config
*ConfigsApi* | [**DeleteConfig**](docs/ConfigsApi.md#deleteconfig) | **Delete** /v1/configs/{configId} | Delete Config
*ConfigsApi* | [**GetConfig**](docs/ConfigsApi.md#getconfig) | **Get** /v1/configs/{configId} | Get Config
*ConfigsApi* | [**GetConfigs**](docs/ConfigsApi.md#getconfigs) | **Get** /v1/products/{productId}/configs | List Configs
*ConfigsApi* | [**UpdateConfig**](docs/ConfigsApi.md#updateconfig) | **Put** /v1/configs/{configId} | Update Config
*EnvironmentsApi* | [**CreateEnvironment**](docs/EnvironmentsApi.md#createenvironment) | **Post** /v1/products/{productId}/environments | Create Environment
*EnvironmentsApi* | [**DeleteEnvironment**](docs/EnvironmentsApi.md#deleteenvironment) | **Delete** /v1/environments/{environmentId} | Delete Environment
*EnvironmentsApi* | [**GetEnvironment**](docs/EnvironmentsApi.md#getenvironment) | **Get** /v1/environments/{environmentId} | Get Environment
*EnvironmentsApi* | [**GetEnvironments**](docs/EnvironmentsApi.md#getenvironments) | **Get** /v1/products/{productId}/environments | List Environments
*EnvironmentsApi* | [**UpdateEnvironment**](docs/EnvironmentsApi.md#updateenvironment) | **Put** /v1/environments/{environmentId} | Update Environment
*FeatureFlagSettingValuesApi* | [**GetSettingValue**](docs/FeatureFlagSettingValuesApi.md#getsettingvalue) | **Get** /v1/environments/{environmentId}/settings/{settingId}/value | Get value
*FeatureFlagSettingValuesApi* | [**GetSettingValues**](docs/FeatureFlagSettingValuesApi.md#getsettingvalues) | **Get** /v1/configs/{configId}/environments/{environmentId}/values | Get values
*FeatureFlagSettingValuesApi* | [**ReplaceSettingValue**](docs/FeatureFlagSettingValuesApi.md#replacesettingvalue) | **Put** /v1/environments/{environmentId}/settings/{settingId}/value | Replace value
*FeatureFlagSettingValuesApi* | [**UpdateSettingValue**](docs/FeatureFlagSettingValuesApi.md#updatesettingvalue) | **Patch** /v1/environments/{environmentId}/settings/{settingId}/value | Update value
*FeatureFlagSettingValuesUsingSDKKeyApi* | [**GetSettingValueBySdkkey**](docs/FeatureFlagSettingValuesUsingSDKKeyApi.md#getsettingvaluebysdkkey) | **Get** /v1/settings/{settingKeyOrId}/value | Get value
*FeatureFlagSettingValuesUsingSDKKeyApi* | [**ReplaceSettingValueBySdkkey**](docs/FeatureFlagSettingValuesUsingSDKKeyApi.md#replacesettingvaluebysdkkey) | **Put** /v1/settings/{settingKeyOrId}/value | Replace value
*FeatureFlagSettingValuesUsingSDKKeyApi* | [**UpdateSettingValueBySdkkey**](docs/FeatureFlagSettingValuesUsingSDKKeyApi.md#updatesettingvaluebysdkkey) | **Patch** /v1/settings/{settingKeyOrId}/value | Update value
*FeatureFlagsSettingsApi* | [**CreateSetting**](docs/FeatureFlagsSettingsApi.md#createsetting) | **Post** /v1/configs/{configId}/settings | Create Flag
*FeatureFlagsSettingsApi* | [**DeleteSetting**](docs/FeatureFlagsSettingsApi.md#deletesetting) | **Delete** /v1/settings/{settingId} | Delete Flag
*FeatureFlagsSettingsApi* | [**GetSetting**](docs/FeatureFlagsSettingsApi.md#getsetting) | **Get** /v1/settings/{settingId} | Get Flag
*FeatureFlagsSettingsApi* | [**GetSettings**](docs/FeatureFlagsSettingsApi.md#getsettings) | **Get** /v1/configs/{configId}/settings | List Flags
*FeatureFlagsSettingsApi* | [**UpdateSetting**](docs/FeatureFlagsSettingsApi.md#updatesetting) | **Patch** /v1/settings/{settingId} | Update Flag
*IntegrationLinksApi* | [**AddOrUpdateIntegrationLink**](docs/IntegrationLinksApi.md#addorupdateintegrationlink) | **Post** /v1/environments/{environmentId}/settings/{settingId}/integrationLinks/{integrationLinkType}/{key} | Add or update Integration link
*IntegrationLinksApi* | [**DeleteIntegrationLink**](docs/IntegrationLinksApi.md#deleteintegrationlink) | **Delete** /v1/environments/{environmentId}/settings/{settingId}/integrationLinks/{integrationLinkType}/{key} | Delete Integration link
*IntegrationLinksApi* | [**GetIntegrationLinkDetails**](docs/IntegrationLinksApi.md#getintegrationlinkdetails) | **Get** /v1/integrationLink/{integrationLinkType}/{key}/details | Get Integration link
*MeApi* | [**GetMe**](docs/MeApi.md#getme) | **Get** /v1/me | Get authenticated user details
*MembersApi* | [**DeleteOrganizationMember**](docs/MembersApi.md#deleteorganizationmember) | **Delete** /v1/organizations/{organizationId}/members/{userId} | Delete Member from Organization
*MembersApi* | [**DeleteProductMember**](docs/MembersApi.md#deleteproductmember) | **Delete** /v1/products/{productId}/members/{userId} | Delete Member from Product
*MembersApi* | [**GetOrganizationMembers**](docs/MembersApi.md#getorganizationmembers) | **Get** /v1/organizations/{organizationId}/members | List Organization Members
*MembersApi* | [**GetProductMembers**](docs/MembersApi.md#getproductmembers) | **Get** /v1/products/{productId}/members | List Product Members
*MembersApi* | [**InviteMember**](docs/MembersApi.md#invitemember) | **Post** /v1/products/{productId}/members/invite | Invite Member
*OrganizationsApi* | [**GetOrganizations**](docs/OrganizationsApi.md#getorganizations) | **Get** /v1/organizations | List Organizations
*PermissionGroupsApi* | [**CreatePermissionGroup**](docs/PermissionGroupsApi.md#createpermissiongroup) | **Post** /v1/products/{productId}/permissions | Create Permission Group
*PermissionGroupsApi* | [**DeletePermissionGroup**](docs/PermissionGroupsApi.md#deletepermissiongroup) | **Delete** /v1/permissions/{permissionGroupId} | Delete Permission Group
*PermissionGroupsApi* | [**GetPermissionGroup**](docs/PermissionGroupsApi.md#getpermissiongroup) | **Get** /v1/permissions/{permissionGroupId} | Get Permission Group
*PermissionGroupsApi* | [**GetPermissionGroups**](docs/PermissionGroupsApi.md#getpermissiongroups) | **Get** /v1/products/{productId}/permissions | List Permission Groups
*PermissionGroupsApi* | [**UpdatePermissionGroup**](docs/PermissionGroupsApi.md#updatepermissiongroup) | **Put** /v1/permissions/{permissionGroupId} | Update Permission Group
*ProductsApi* | [**CreateProduct**](docs/ProductsApi.md#createproduct) | **Post** /v1/organizations/{organizationId}/products | Create Product
*ProductsApi* | [**DeleteProduct**](docs/ProductsApi.md#deleteproduct) | **Delete** /v1/products/{productId} | Delete Product
*ProductsApi* | [**GetProduct**](docs/ProductsApi.md#getproduct) | **Get** /v1/products/{productId} | Get Product
*ProductsApi* | [**GetProducts**](docs/ProductsApi.md#getproducts) | **Get** /v1/products | List Products
*ProductsApi* | [**UpdateProduct**](docs/ProductsApi.md#updateproduct) | **Put** /v1/products/{productId} | Update Product
*SDKKeysApi* | [**GetSdkKeys**](docs/SDKKeysApi.md#getsdkkeys) | **Get** /v1/configs/{configId}/environments/{environmentId} | Get SDK Key
*TagsApi* | [**CreateTag**](docs/TagsApi.md#createtag) | **Post** /v1/products/{productId}/tags | Create Tag
*TagsApi* | [**DeleteTag**](docs/TagsApi.md#deletetag) | **Delete** /v1/tags/{tagId} | Delete Tag
*TagsApi* | [**GetSettingsByTag**](docs/TagsApi.md#getsettingsbytag) | **Get** /v1/tags/{tagId}/settings | List Settings by Tag
*TagsApi* | [**GetTag**](docs/TagsApi.md#gettag) | **Get** /v1/tags/{tagId} | Get Tag
*TagsApi* | [**GetTags**](docs/TagsApi.md#gettags) | **Get** /v1/products/{productId}/tags | List Tags
*TagsApi* | [**UpdateTag**](docs/TagsApi.md#updatetag) | **Put** /v1/tags/{tagId} | Update Tag

## Documentation For Models

 - [AccessType](docs/AccessType.md)
 - [AddOrUpdateIntegrationLinkModel](docs/AddOrUpdateIntegrationLinkModel.md)
 - [AuditLogItemModel](docs/AuditLogItemModel.md)
 - [AuditLogType](docs/AuditLogType.md)
 - [AuditLogType1](docs/AuditLogType1.md)
 - [ConfigModel](docs/ConfigModel.md)
 - [ConfigModelHaljson](docs/ConfigModelHaljson.md)
 - [ConfigModelhaljsonEmbedded](docs/ConfigModelhaljsonEmbedded.md)
 - [ConfigModelhaljsonEmbeddedProduct](docs/ConfigModelhaljsonEmbeddedProduct.md)
 - [ConfigModelhaljsonEmbeddedProductEmbedded](docs/ConfigModelhaljsonEmbeddedProductEmbedded.md)
 - [ConfigModelhaljsonEmbeddedProductEmbeddedOrganization](docs/ConfigModelhaljsonEmbeddedProductEmbeddedOrganization.md)
 - [ConfigModelhaljsonEmbeddedProductEmbeddedOrganizationLinks](docs/ConfigModelhaljsonEmbeddedProductEmbeddedOrganizationLinks.md)
 - [ConfigModelhaljsonEmbeddedProductLinks](docs/ConfigModelhaljsonEmbeddedProductLinks.md)
 - [ConfigModelhaljsonLinks](docs/ConfigModelhaljsonLinks.md)
 - [ConfigSettingValueModel](docs/ConfigSettingValueModel.md)
 - [ConfigSettingValuesModel](docs/ConfigSettingValuesModel.md)
 - [CreateConfigRequest](docs/CreateConfigRequest.md)
 - [CreateEnvironmentModel](docs/CreateEnvironmentModel.md)
 - [CreatePermissionGroupRequest](docs/CreatePermissionGroupRequest.md)
 - [CreateProductRequest](docs/CreateProductRequest.md)
 - [CreateSettingModel](docs/CreateSettingModel.md)
 - [CreateTagModel](docs/CreateTagModel.md)
 - [DeleteIntegrationLinkModel](docs/DeleteIntegrationLinkModel.md)
 - [EnvironmentAccessModel](docs/EnvironmentAccessModel.md)
 - [EnvironmentAccessType](docs/EnvironmentAccessType.md)
 - [EnvironmentModel](docs/EnvironmentModel.md)
 - [EnvironmentModelHaljson](docs/EnvironmentModelHaljson.md)
 - [EnvironmentModelhaljsonLinks](docs/EnvironmentModelhaljsonLinks.md)
 - [IntegrationLinkDetail](docs/IntegrationLinkDetail.md)
 - [IntegrationLinkDetailsModel](docs/IntegrationLinkDetailsModel.md)
 - [IntegrationLinkModel](docs/IntegrationLinkModel.md)
 - [IntegrationLinkType](docs/IntegrationLinkType.md)
 - [InviteMembersRequest](docs/InviteMembersRequest.md)
 - [MeModel](docs/MeModel.md)
 - [MemberModel](docs/MemberModel.md)
 - [Operation](docs/Operation.md)
 - [OrganizationModel](docs/OrganizationModel.md)
 - [OrganizationModelHaljson](docs/OrganizationModelHaljson.md)
 - [PermissionGroupModel](docs/PermissionGroupModel.md)
 - [ProductModel](docs/ProductModel.md)
 - [ProductModelHaljson](docs/ProductModelHaljson.md)
 - [RolloutPercentageItemModel](docs/RolloutPercentageItemModel.md)
 - [RolloutRuleComparator](docs/RolloutRuleComparator.md)
 - [RolloutRuleModel](docs/RolloutRuleModel.md)
 - [SdkKeysModel](docs/SdkKeysModel.md)
 - [SettingDataModel](docs/SettingDataModel.md)
 - [SettingModel](docs/SettingModel.md)
 - [SettingModelHaljson](docs/SettingModelHaljson.md)
 - [SettingModelhaljsonEmbedded](docs/SettingModelhaljsonEmbedded.md)
 - [SettingModelhaljsonEmbeddedTags](docs/SettingModelhaljsonEmbeddedTags.md)
 - [SettingTagModel](docs/SettingTagModel.md)
 - [SettingType](docs/SettingType.md)
 - [SettingValueModel](docs/SettingValueModel.md)
 - [SettingValueModelHaljson](docs/SettingValueModelHaljson.md)
 - [SettingValueModelhaljsonEmbedded](docs/SettingValueModelhaljsonEmbedded.md)
 - [SettingValueModelhaljsonEmbeddedConfig](docs/SettingValueModelhaljsonEmbeddedConfig.md)
 - [SettingValueModelhaljsonEmbeddedEnvironment](docs/SettingValueModelhaljsonEmbeddedEnvironment.md)
 - [SettingValueModelhaljsonEmbeddedIntegrationLinks](docs/SettingValueModelhaljsonEmbeddedIntegrationLinks.md)
 - [SettingValueModelhaljsonEmbeddedSetting](docs/SettingValueModelhaljsonEmbeddedSetting.md)
 - [SettingValueModelhaljsonEmbeddedSettingTags](docs/SettingValueModelhaljsonEmbeddedSettingTags.md)
 - [TagModel](docs/TagModel.md)
 - [TagModelHaljson](docs/TagModelHaljson.md)
 - [UpdateConfigRequest](docs/UpdateConfigRequest.md)
 - [UpdateEnvironmentModel](docs/UpdateEnvironmentModel.md)
 - [UpdatePermissionGroupRequest](docs/UpdatePermissionGroupRequest.md)
 - [UpdateProductRequest](docs/UpdateProductRequest.md)
 - [UpdateSettingValueModel](docs/UpdateSettingValueModel.md)
 - [UpdateTagModel](docs/UpdateTagModel.md)
 - [UserModel](docs/UserModel.md)

## Documentation For Authorization

## Basic
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

## Author

support@configcat.com
