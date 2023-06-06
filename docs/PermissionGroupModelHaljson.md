# PermissionGroupModelHaljson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionGroupId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CanManageMembers** | Pointer to **bool** |  | [optional] 
**CanCreateOrUpdateConfig** | Pointer to **bool** |  | [optional] 
**CanDeleteConfig** | Pointer to **bool** |  | [optional] 
**CanCreateOrUpdateEnvironment** | Pointer to **bool** |  | [optional] 
**CanDeleteEnvironment** | Pointer to **bool** |  | [optional] 
**CanCreateOrUpdateSetting** | Pointer to **bool** |  | [optional] 
**CanTagSetting** | Pointer to **bool** |  | [optional] 
**CanDeleteSetting** | Pointer to **bool** |  | [optional] 
**CanCreateOrUpdateTag** | Pointer to **bool** |  | [optional] 
**CanDeleteTag** | Pointer to **bool** |  | [optional] 
**CanManageWebhook** | Pointer to **bool** |  | [optional] 
**CanUseExportImport** | Pointer to **bool** |  | [optional] 
**CanManageProductPreferences** | Pointer to **bool** |  | [optional] 
**CanManageIntegrations** | Pointer to **bool** |  | [optional] 
**CanViewSdkKey** | Pointer to **bool** |  | [optional] 
**CanRotateSdkKey** | Pointer to **bool** |  | [optional] 
**CanCreateOrUpdateSegments** | Pointer to **bool** |  | [optional] 
**CanDeleteSegments** | Pointer to **bool** |  | [optional] 
**CanViewProductAuditLog** | Pointer to **bool** |  | [optional] 
**CanViewProductStatistics** | Pointer to **bool** |  | [optional] 
**AccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**NewEnvironmentAccessType** | Pointer to [**EnvironmentAccessType**](EnvironmentAccessType.md) |  | [optional] 
**EnvironmentAccesses** | Pointer to [**[]EnvironmentAccessModel**](EnvironmentAccessModel.md) |  | [optional] 
**Embedded** | Pointer to [**ConfigModelHaljsonEmbedded**](ConfigModelHaljsonEmbedded.md) |  | [optional] 
**Links** | Pointer to [**EnvironmentModelHaljsonLinks**](EnvironmentModelHaljsonLinks.md) |  | [optional] 

## Methods

### NewPermissionGroupModelHaljson

`func NewPermissionGroupModelHaljson() *PermissionGroupModelHaljson`

NewPermissionGroupModelHaljson instantiates a new PermissionGroupModelHaljson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionGroupModelHaljsonWithDefaults

`func NewPermissionGroupModelHaljsonWithDefaults() *PermissionGroupModelHaljson`

NewPermissionGroupModelHaljsonWithDefaults instantiates a new PermissionGroupModelHaljson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionGroupId

`func (o *PermissionGroupModelHaljson) GetPermissionGroupId() int64`

GetPermissionGroupId returns the PermissionGroupId field if non-nil, zero value otherwise.

### GetPermissionGroupIdOk

`func (o *PermissionGroupModelHaljson) GetPermissionGroupIdOk() (*int64, bool)`

GetPermissionGroupIdOk returns a tuple with the PermissionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroupId

`func (o *PermissionGroupModelHaljson) SetPermissionGroupId(v int64)`

SetPermissionGroupId sets PermissionGroupId field to given value.

### HasPermissionGroupId

`func (o *PermissionGroupModelHaljson) HasPermissionGroupId() bool`

HasPermissionGroupId returns a boolean if a field has been set.

### GetName

`func (o *PermissionGroupModelHaljson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionGroupModelHaljson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionGroupModelHaljson) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PermissionGroupModelHaljson) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PermissionGroupModelHaljson) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PermissionGroupModelHaljson) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCanManageMembers

`func (o *PermissionGroupModelHaljson) GetCanManageMembers() bool`

GetCanManageMembers returns the CanManageMembers field if non-nil, zero value otherwise.

### GetCanManageMembersOk

`func (o *PermissionGroupModelHaljson) GetCanManageMembersOk() (*bool, bool)`

GetCanManageMembersOk returns a tuple with the CanManageMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageMembers

`func (o *PermissionGroupModelHaljson) SetCanManageMembers(v bool)`

SetCanManageMembers sets CanManageMembers field to given value.

### HasCanManageMembers

`func (o *PermissionGroupModelHaljson) HasCanManageMembers() bool`

HasCanManageMembers returns a boolean if a field has been set.

### GetCanCreateOrUpdateConfig

`func (o *PermissionGroupModelHaljson) GetCanCreateOrUpdateConfig() bool`

GetCanCreateOrUpdateConfig returns the CanCreateOrUpdateConfig field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateConfigOk

`func (o *PermissionGroupModelHaljson) GetCanCreateOrUpdateConfigOk() (*bool, bool)`

GetCanCreateOrUpdateConfigOk returns a tuple with the CanCreateOrUpdateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateConfig

`func (o *PermissionGroupModelHaljson) SetCanCreateOrUpdateConfig(v bool)`

SetCanCreateOrUpdateConfig sets CanCreateOrUpdateConfig field to given value.

### HasCanCreateOrUpdateConfig

`func (o *PermissionGroupModelHaljson) HasCanCreateOrUpdateConfig() bool`

HasCanCreateOrUpdateConfig returns a boolean if a field has been set.

### GetCanDeleteConfig

`func (o *PermissionGroupModelHaljson) GetCanDeleteConfig() bool`

GetCanDeleteConfig returns the CanDeleteConfig field if non-nil, zero value otherwise.

### GetCanDeleteConfigOk

`func (o *PermissionGroupModelHaljson) GetCanDeleteConfigOk() (*bool, bool)`

GetCanDeleteConfigOk returns a tuple with the CanDeleteConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteConfig

`func (o *PermissionGroupModelHaljson) SetCanDeleteConfig(v bool)`

SetCanDeleteConfig sets CanDeleteConfig field to given value.

### HasCanDeleteConfig

`func (o *PermissionGroupModelHaljson) HasCanDeleteConfig() bool`

HasCanDeleteConfig returns a boolean if a field has been set.

### GetCanCreateOrUpdateEnvironment

`func (o *PermissionGroupModelHaljson) GetCanCreateOrUpdateEnvironment() bool`

GetCanCreateOrUpdateEnvironment returns the CanCreateOrUpdateEnvironment field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateEnvironmentOk

`func (o *PermissionGroupModelHaljson) GetCanCreateOrUpdateEnvironmentOk() (*bool, bool)`

GetCanCreateOrUpdateEnvironmentOk returns a tuple with the CanCreateOrUpdateEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateEnvironment

`func (o *PermissionGroupModelHaljson) SetCanCreateOrUpdateEnvironment(v bool)`

SetCanCreateOrUpdateEnvironment sets CanCreateOrUpdateEnvironment field to given value.

### HasCanCreateOrUpdateEnvironment

`func (o *PermissionGroupModelHaljson) HasCanCreateOrUpdateEnvironment() bool`

HasCanCreateOrUpdateEnvironment returns a boolean if a field has been set.

### GetCanDeleteEnvironment

`func (o *PermissionGroupModelHaljson) GetCanDeleteEnvironment() bool`

GetCanDeleteEnvironment returns the CanDeleteEnvironment field if non-nil, zero value otherwise.

### GetCanDeleteEnvironmentOk

`func (o *PermissionGroupModelHaljson) GetCanDeleteEnvironmentOk() (*bool, bool)`

GetCanDeleteEnvironmentOk returns a tuple with the CanDeleteEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteEnvironment

`func (o *PermissionGroupModelHaljson) SetCanDeleteEnvironment(v bool)`

SetCanDeleteEnvironment sets CanDeleteEnvironment field to given value.

### HasCanDeleteEnvironment

`func (o *PermissionGroupModelHaljson) HasCanDeleteEnvironment() bool`

HasCanDeleteEnvironment returns a boolean if a field has been set.

### GetCanCreateOrUpdateSetting

`func (o *PermissionGroupModelHaljson) GetCanCreateOrUpdateSetting() bool`

GetCanCreateOrUpdateSetting returns the CanCreateOrUpdateSetting field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateSettingOk

`func (o *PermissionGroupModelHaljson) GetCanCreateOrUpdateSettingOk() (*bool, bool)`

GetCanCreateOrUpdateSettingOk returns a tuple with the CanCreateOrUpdateSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateSetting

`func (o *PermissionGroupModelHaljson) SetCanCreateOrUpdateSetting(v bool)`

SetCanCreateOrUpdateSetting sets CanCreateOrUpdateSetting field to given value.

### HasCanCreateOrUpdateSetting

`func (o *PermissionGroupModelHaljson) HasCanCreateOrUpdateSetting() bool`

HasCanCreateOrUpdateSetting returns a boolean if a field has been set.

### GetCanTagSetting

`func (o *PermissionGroupModelHaljson) GetCanTagSetting() bool`

GetCanTagSetting returns the CanTagSetting field if non-nil, zero value otherwise.

### GetCanTagSettingOk

`func (o *PermissionGroupModelHaljson) GetCanTagSettingOk() (*bool, bool)`

GetCanTagSettingOk returns a tuple with the CanTagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTagSetting

`func (o *PermissionGroupModelHaljson) SetCanTagSetting(v bool)`

SetCanTagSetting sets CanTagSetting field to given value.

### HasCanTagSetting

`func (o *PermissionGroupModelHaljson) HasCanTagSetting() bool`

HasCanTagSetting returns a boolean if a field has been set.

### GetCanDeleteSetting

`func (o *PermissionGroupModelHaljson) GetCanDeleteSetting() bool`

GetCanDeleteSetting returns the CanDeleteSetting field if non-nil, zero value otherwise.

### GetCanDeleteSettingOk

`func (o *PermissionGroupModelHaljson) GetCanDeleteSettingOk() (*bool, bool)`

GetCanDeleteSettingOk returns a tuple with the CanDeleteSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteSetting

`func (o *PermissionGroupModelHaljson) SetCanDeleteSetting(v bool)`

SetCanDeleteSetting sets CanDeleteSetting field to given value.

### HasCanDeleteSetting

`func (o *PermissionGroupModelHaljson) HasCanDeleteSetting() bool`

HasCanDeleteSetting returns a boolean if a field has been set.

### GetCanCreateOrUpdateTag

`func (o *PermissionGroupModelHaljson) GetCanCreateOrUpdateTag() bool`

GetCanCreateOrUpdateTag returns the CanCreateOrUpdateTag field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateTagOk

`func (o *PermissionGroupModelHaljson) GetCanCreateOrUpdateTagOk() (*bool, bool)`

GetCanCreateOrUpdateTagOk returns a tuple with the CanCreateOrUpdateTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateTag

`func (o *PermissionGroupModelHaljson) SetCanCreateOrUpdateTag(v bool)`

SetCanCreateOrUpdateTag sets CanCreateOrUpdateTag field to given value.

### HasCanCreateOrUpdateTag

`func (o *PermissionGroupModelHaljson) HasCanCreateOrUpdateTag() bool`

HasCanCreateOrUpdateTag returns a boolean if a field has been set.

### GetCanDeleteTag

`func (o *PermissionGroupModelHaljson) GetCanDeleteTag() bool`

GetCanDeleteTag returns the CanDeleteTag field if non-nil, zero value otherwise.

### GetCanDeleteTagOk

`func (o *PermissionGroupModelHaljson) GetCanDeleteTagOk() (*bool, bool)`

GetCanDeleteTagOk returns a tuple with the CanDeleteTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteTag

`func (o *PermissionGroupModelHaljson) SetCanDeleteTag(v bool)`

SetCanDeleteTag sets CanDeleteTag field to given value.

### HasCanDeleteTag

`func (o *PermissionGroupModelHaljson) HasCanDeleteTag() bool`

HasCanDeleteTag returns a boolean if a field has been set.

### GetCanManageWebhook

`func (o *PermissionGroupModelHaljson) GetCanManageWebhook() bool`

GetCanManageWebhook returns the CanManageWebhook field if non-nil, zero value otherwise.

### GetCanManageWebhookOk

`func (o *PermissionGroupModelHaljson) GetCanManageWebhookOk() (*bool, bool)`

GetCanManageWebhookOk returns a tuple with the CanManageWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageWebhook

`func (o *PermissionGroupModelHaljson) SetCanManageWebhook(v bool)`

SetCanManageWebhook sets CanManageWebhook field to given value.

### HasCanManageWebhook

`func (o *PermissionGroupModelHaljson) HasCanManageWebhook() bool`

HasCanManageWebhook returns a boolean if a field has been set.

### GetCanUseExportImport

`func (o *PermissionGroupModelHaljson) GetCanUseExportImport() bool`

GetCanUseExportImport returns the CanUseExportImport field if non-nil, zero value otherwise.

### GetCanUseExportImportOk

`func (o *PermissionGroupModelHaljson) GetCanUseExportImportOk() (*bool, bool)`

GetCanUseExportImportOk returns a tuple with the CanUseExportImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUseExportImport

`func (o *PermissionGroupModelHaljson) SetCanUseExportImport(v bool)`

SetCanUseExportImport sets CanUseExportImport field to given value.

### HasCanUseExportImport

`func (o *PermissionGroupModelHaljson) HasCanUseExportImport() bool`

HasCanUseExportImport returns a boolean if a field has been set.

### GetCanManageProductPreferences

`func (o *PermissionGroupModelHaljson) GetCanManageProductPreferences() bool`

GetCanManageProductPreferences returns the CanManageProductPreferences field if non-nil, zero value otherwise.

### GetCanManageProductPreferencesOk

`func (o *PermissionGroupModelHaljson) GetCanManageProductPreferencesOk() (*bool, bool)`

GetCanManageProductPreferencesOk returns a tuple with the CanManageProductPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageProductPreferences

`func (o *PermissionGroupModelHaljson) SetCanManageProductPreferences(v bool)`

SetCanManageProductPreferences sets CanManageProductPreferences field to given value.

### HasCanManageProductPreferences

`func (o *PermissionGroupModelHaljson) HasCanManageProductPreferences() bool`

HasCanManageProductPreferences returns a boolean if a field has been set.

### GetCanManageIntegrations

`func (o *PermissionGroupModelHaljson) GetCanManageIntegrations() bool`

GetCanManageIntegrations returns the CanManageIntegrations field if non-nil, zero value otherwise.

### GetCanManageIntegrationsOk

`func (o *PermissionGroupModelHaljson) GetCanManageIntegrationsOk() (*bool, bool)`

GetCanManageIntegrationsOk returns a tuple with the CanManageIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageIntegrations

`func (o *PermissionGroupModelHaljson) SetCanManageIntegrations(v bool)`

SetCanManageIntegrations sets CanManageIntegrations field to given value.

### HasCanManageIntegrations

`func (o *PermissionGroupModelHaljson) HasCanManageIntegrations() bool`

HasCanManageIntegrations returns a boolean if a field has been set.

### GetCanViewSdkKey

`func (o *PermissionGroupModelHaljson) GetCanViewSdkKey() bool`

GetCanViewSdkKey returns the CanViewSdkKey field if non-nil, zero value otherwise.

### GetCanViewSdkKeyOk

`func (o *PermissionGroupModelHaljson) GetCanViewSdkKeyOk() (*bool, bool)`

GetCanViewSdkKeyOk returns a tuple with the CanViewSdkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewSdkKey

`func (o *PermissionGroupModelHaljson) SetCanViewSdkKey(v bool)`

SetCanViewSdkKey sets CanViewSdkKey field to given value.

### HasCanViewSdkKey

`func (o *PermissionGroupModelHaljson) HasCanViewSdkKey() bool`

HasCanViewSdkKey returns a boolean if a field has been set.

### GetCanRotateSdkKey

`func (o *PermissionGroupModelHaljson) GetCanRotateSdkKey() bool`

GetCanRotateSdkKey returns the CanRotateSdkKey field if non-nil, zero value otherwise.

### GetCanRotateSdkKeyOk

`func (o *PermissionGroupModelHaljson) GetCanRotateSdkKeyOk() (*bool, bool)`

GetCanRotateSdkKeyOk returns a tuple with the CanRotateSdkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRotateSdkKey

`func (o *PermissionGroupModelHaljson) SetCanRotateSdkKey(v bool)`

SetCanRotateSdkKey sets CanRotateSdkKey field to given value.

### HasCanRotateSdkKey

`func (o *PermissionGroupModelHaljson) HasCanRotateSdkKey() bool`

HasCanRotateSdkKey returns a boolean if a field has been set.

### GetCanCreateOrUpdateSegments

`func (o *PermissionGroupModelHaljson) GetCanCreateOrUpdateSegments() bool`

GetCanCreateOrUpdateSegments returns the CanCreateOrUpdateSegments field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateSegmentsOk

`func (o *PermissionGroupModelHaljson) GetCanCreateOrUpdateSegmentsOk() (*bool, bool)`

GetCanCreateOrUpdateSegmentsOk returns a tuple with the CanCreateOrUpdateSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateSegments

`func (o *PermissionGroupModelHaljson) SetCanCreateOrUpdateSegments(v bool)`

SetCanCreateOrUpdateSegments sets CanCreateOrUpdateSegments field to given value.

### HasCanCreateOrUpdateSegments

`func (o *PermissionGroupModelHaljson) HasCanCreateOrUpdateSegments() bool`

HasCanCreateOrUpdateSegments returns a boolean if a field has been set.

### GetCanDeleteSegments

`func (o *PermissionGroupModelHaljson) GetCanDeleteSegments() bool`

GetCanDeleteSegments returns the CanDeleteSegments field if non-nil, zero value otherwise.

### GetCanDeleteSegmentsOk

`func (o *PermissionGroupModelHaljson) GetCanDeleteSegmentsOk() (*bool, bool)`

GetCanDeleteSegmentsOk returns a tuple with the CanDeleteSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteSegments

`func (o *PermissionGroupModelHaljson) SetCanDeleteSegments(v bool)`

SetCanDeleteSegments sets CanDeleteSegments field to given value.

### HasCanDeleteSegments

`func (o *PermissionGroupModelHaljson) HasCanDeleteSegments() bool`

HasCanDeleteSegments returns a boolean if a field has been set.

### GetCanViewProductAuditLog

`func (o *PermissionGroupModelHaljson) GetCanViewProductAuditLog() bool`

GetCanViewProductAuditLog returns the CanViewProductAuditLog field if non-nil, zero value otherwise.

### GetCanViewProductAuditLogOk

`func (o *PermissionGroupModelHaljson) GetCanViewProductAuditLogOk() (*bool, bool)`

GetCanViewProductAuditLogOk returns a tuple with the CanViewProductAuditLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewProductAuditLog

`func (o *PermissionGroupModelHaljson) SetCanViewProductAuditLog(v bool)`

SetCanViewProductAuditLog sets CanViewProductAuditLog field to given value.

### HasCanViewProductAuditLog

`func (o *PermissionGroupModelHaljson) HasCanViewProductAuditLog() bool`

HasCanViewProductAuditLog returns a boolean if a field has been set.

### GetCanViewProductStatistics

`func (o *PermissionGroupModelHaljson) GetCanViewProductStatistics() bool`

GetCanViewProductStatistics returns the CanViewProductStatistics field if non-nil, zero value otherwise.

### GetCanViewProductStatisticsOk

`func (o *PermissionGroupModelHaljson) GetCanViewProductStatisticsOk() (*bool, bool)`

GetCanViewProductStatisticsOk returns a tuple with the CanViewProductStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewProductStatistics

`func (o *PermissionGroupModelHaljson) SetCanViewProductStatistics(v bool)`

SetCanViewProductStatistics sets CanViewProductStatistics field to given value.

### HasCanViewProductStatistics

`func (o *PermissionGroupModelHaljson) HasCanViewProductStatistics() bool`

HasCanViewProductStatistics returns a boolean if a field has been set.

### GetAccessType

`func (o *PermissionGroupModelHaljson) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *PermissionGroupModelHaljson) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *PermissionGroupModelHaljson) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *PermissionGroupModelHaljson) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetNewEnvironmentAccessType

`func (o *PermissionGroupModelHaljson) GetNewEnvironmentAccessType() EnvironmentAccessType`

GetNewEnvironmentAccessType returns the NewEnvironmentAccessType field if non-nil, zero value otherwise.

### GetNewEnvironmentAccessTypeOk

`func (o *PermissionGroupModelHaljson) GetNewEnvironmentAccessTypeOk() (*EnvironmentAccessType, bool)`

GetNewEnvironmentAccessTypeOk returns a tuple with the NewEnvironmentAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEnvironmentAccessType

`func (o *PermissionGroupModelHaljson) SetNewEnvironmentAccessType(v EnvironmentAccessType)`

SetNewEnvironmentAccessType sets NewEnvironmentAccessType field to given value.

### HasNewEnvironmentAccessType

`func (o *PermissionGroupModelHaljson) HasNewEnvironmentAccessType() bool`

HasNewEnvironmentAccessType returns a boolean if a field has been set.

### GetEnvironmentAccesses

`func (o *PermissionGroupModelHaljson) GetEnvironmentAccesses() []EnvironmentAccessModel`

GetEnvironmentAccesses returns the EnvironmentAccesses field if non-nil, zero value otherwise.

### GetEnvironmentAccessesOk

`func (o *PermissionGroupModelHaljson) GetEnvironmentAccessesOk() (*[]EnvironmentAccessModel, bool)`

GetEnvironmentAccessesOk returns a tuple with the EnvironmentAccesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentAccesses

`func (o *PermissionGroupModelHaljson) SetEnvironmentAccesses(v []EnvironmentAccessModel)`

SetEnvironmentAccesses sets EnvironmentAccesses field to given value.

### HasEnvironmentAccesses

`func (o *PermissionGroupModelHaljson) HasEnvironmentAccesses() bool`

HasEnvironmentAccesses returns a boolean if a field has been set.

### SetEnvironmentAccessesNil

`func (o *PermissionGroupModelHaljson) SetEnvironmentAccessesNil(b bool)`

 SetEnvironmentAccessesNil sets the value for EnvironmentAccesses to be an explicit nil

### UnsetEnvironmentAccesses
`func (o *PermissionGroupModelHaljson) UnsetEnvironmentAccesses()`

UnsetEnvironmentAccesses ensures that no value is present for EnvironmentAccesses, not even an explicit nil
### GetEmbedded

`func (o *PermissionGroupModelHaljson) GetEmbedded() ConfigModelHaljsonEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *PermissionGroupModelHaljson) GetEmbeddedOk() (*ConfigModelHaljsonEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *PermissionGroupModelHaljson) SetEmbedded(v ConfigModelHaljsonEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *PermissionGroupModelHaljson) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *PermissionGroupModelHaljson) GetLinks() EnvironmentModelHaljsonLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PermissionGroupModelHaljson) GetLinksOk() (*EnvironmentModelHaljsonLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PermissionGroupModelHaljson) SetLinks(v EnvironmentModelHaljsonLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PermissionGroupModelHaljson) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


