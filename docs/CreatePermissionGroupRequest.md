# CreatePermissionGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
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
**CanViewProductStatistics** | Pointer to **bool** |  | [optional] 
**CanViewProductAuditLog** | Pointer to **bool** |  | [optional] 
**CanCreateOrUpdateSegments** | Pointer to **bool** |  | [optional] 
**CanDeleteSegments** | Pointer to **bool** |  | [optional] 
**AccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**NewEnvironmentAccessType** | Pointer to [**EnvironmentAccessType**](EnvironmentAccessType.md) |  | [optional] 
**EnvironmentAccesses** | Pointer to [**[]CreateOrUpdateEnvironmentAccessModel**](CreateOrUpdateEnvironmentAccessModel.md) |  | [optional] 

## Methods

### NewCreatePermissionGroupRequest

`func NewCreatePermissionGroupRequest(name string, ) *CreatePermissionGroupRequest`

NewCreatePermissionGroupRequest instantiates a new CreatePermissionGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePermissionGroupRequestWithDefaults

`func NewCreatePermissionGroupRequestWithDefaults() *CreatePermissionGroupRequest`

NewCreatePermissionGroupRequestWithDefaults instantiates a new CreatePermissionGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePermissionGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePermissionGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePermissionGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCanManageMembers

`func (o *CreatePermissionGroupRequest) GetCanManageMembers() bool`

GetCanManageMembers returns the CanManageMembers field if non-nil, zero value otherwise.

### GetCanManageMembersOk

`func (o *CreatePermissionGroupRequest) GetCanManageMembersOk() (*bool, bool)`

GetCanManageMembersOk returns a tuple with the CanManageMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageMembers

`func (o *CreatePermissionGroupRequest) SetCanManageMembers(v bool)`

SetCanManageMembers sets CanManageMembers field to given value.

### HasCanManageMembers

`func (o *CreatePermissionGroupRequest) HasCanManageMembers() bool`

HasCanManageMembers returns a boolean if a field has been set.

### GetCanCreateOrUpdateConfig

`func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateConfig() bool`

GetCanCreateOrUpdateConfig returns the CanCreateOrUpdateConfig field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateConfigOk

`func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateConfigOk() (*bool, bool)`

GetCanCreateOrUpdateConfigOk returns a tuple with the CanCreateOrUpdateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateConfig

`func (o *CreatePermissionGroupRequest) SetCanCreateOrUpdateConfig(v bool)`

SetCanCreateOrUpdateConfig sets CanCreateOrUpdateConfig field to given value.

### HasCanCreateOrUpdateConfig

`func (o *CreatePermissionGroupRequest) HasCanCreateOrUpdateConfig() bool`

HasCanCreateOrUpdateConfig returns a boolean if a field has been set.

### GetCanDeleteConfig

`func (o *CreatePermissionGroupRequest) GetCanDeleteConfig() bool`

GetCanDeleteConfig returns the CanDeleteConfig field if non-nil, zero value otherwise.

### GetCanDeleteConfigOk

`func (o *CreatePermissionGroupRequest) GetCanDeleteConfigOk() (*bool, bool)`

GetCanDeleteConfigOk returns a tuple with the CanDeleteConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteConfig

`func (o *CreatePermissionGroupRequest) SetCanDeleteConfig(v bool)`

SetCanDeleteConfig sets CanDeleteConfig field to given value.

### HasCanDeleteConfig

`func (o *CreatePermissionGroupRequest) HasCanDeleteConfig() bool`

HasCanDeleteConfig returns a boolean if a field has been set.

### GetCanCreateOrUpdateEnvironment

`func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateEnvironment() bool`

GetCanCreateOrUpdateEnvironment returns the CanCreateOrUpdateEnvironment field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateEnvironmentOk

`func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateEnvironmentOk() (*bool, bool)`

GetCanCreateOrUpdateEnvironmentOk returns a tuple with the CanCreateOrUpdateEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateEnvironment

`func (o *CreatePermissionGroupRequest) SetCanCreateOrUpdateEnvironment(v bool)`

SetCanCreateOrUpdateEnvironment sets CanCreateOrUpdateEnvironment field to given value.

### HasCanCreateOrUpdateEnvironment

`func (o *CreatePermissionGroupRequest) HasCanCreateOrUpdateEnvironment() bool`

HasCanCreateOrUpdateEnvironment returns a boolean if a field has been set.

### GetCanDeleteEnvironment

`func (o *CreatePermissionGroupRequest) GetCanDeleteEnvironment() bool`

GetCanDeleteEnvironment returns the CanDeleteEnvironment field if non-nil, zero value otherwise.

### GetCanDeleteEnvironmentOk

`func (o *CreatePermissionGroupRequest) GetCanDeleteEnvironmentOk() (*bool, bool)`

GetCanDeleteEnvironmentOk returns a tuple with the CanDeleteEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteEnvironment

`func (o *CreatePermissionGroupRequest) SetCanDeleteEnvironment(v bool)`

SetCanDeleteEnvironment sets CanDeleteEnvironment field to given value.

### HasCanDeleteEnvironment

`func (o *CreatePermissionGroupRequest) HasCanDeleteEnvironment() bool`

HasCanDeleteEnvironment returns a boolean if a field has been set.

### GetCanCreateOrUpdateSetting

`func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateSetting() bool`

GetCanCreateOrUpdateSetting returns the CanCreateOrUpdateSetting field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateSettingOk

`func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateSettingOk() (*bool, bool)`

GetCanCreateOrUpdateSettingOk returns a tuple with the CanCreateOrUpdateSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateSetting

`func (o *CreatePermissionGroupRequest) SetCanCreateOrUpdateSetting(v bool)`

SetCanCreateOrUpdateSetting sets CanCreateOrUpdateSetting field to given value.

### HasCanCreateOrUpdateSetting

`func (o *CreatePermissionGroupRequest) HasCanCreateOrUpdateSetting() bool`

HasCanCreateOrUpdateSetting returns a boolean if a field has been set.

### GetCanTagSetting

`func (o *CreatePermissionGroupRequest) GetCanTagSetting() bool`

GetCanTagSetting returns the CanTagSetting field if non-nil, zero value otherwise.

### GetCanTagSettingOk

`func (o *CreatePermissionGroupRequest) GetCanTagSettingOk() (*bool, bool)`

GetCanTagSettingOk returns a tuple with the CanTagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTagSetting

`func (o *CreatePermissionGroupRequest) SetCanTagSetting(v bool)`

SetCanTagSetting sets CanTagSetting field to given value.

### HasCanTagSetting

`func (o *CreatePermissionGroupRequest) HasCanTagSetting() bool`

HasCanTagSetting returns a boolean if a field has been set.

### GetCanDeleteSetting

`func (o *CreatePermissionGroupRequest) GetCanDeleteSetting() bool`

GetCanDeleteSetting returns the CanDeleteSetting field if non-nil, zero value otherwise.

### GetCanDeleteSettingOk

`func (o *CreatePermissionGroupRequest) GetCanDeleteSettingOk() (*bool, bool)`

GetCanDeleteSettingOk returns a tuple with the CanDeleteSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteSetting

`func (o *CreatePermissionGroupRequest) SetCanDeleteSetting(v bool)`

SetCanDeleteSetting sets CanDeleteSetting field to given value.

### HasCanDeleteSetting

`func (o *CreatePermissionGroupRequest) HasCanDeleteSetting() bool`

HasCanDeleteSetting returns a boolean if a field has been set.

### GetCanCreateOrUpdateTag

`func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateTag() bool`

GetCanCreateOrUpdateTag returns the CanCreateOrUpdateTag field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateTagOk

`func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateTagOk() (*bool, bool)`

GetCanCreateOrUpdateTagOk returns a tuple with the CanCreateOrUpdateTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateTag

`func (o *CreatePermissionGroupRequest) SetCanCreateOrUpdateTag(v bool)`

SetCanCreateOrUpdateTag sets CanCreateOrUpdateTag field to given value.

### HasCanCreateOrUpdateTag

`func (o *CreatePermissionGroupRequest) HasCanCreateOrUpdateTag() bool`

HasCanCreateOrUpdateTag returns a boolean if a field has been set.

### GetCanDeleteTag

`func (o *CreatePermissionGroupRequest) GetCanDeleteTag() bool`

GetCanDeleteTag returns the CanDeleteTag field if non-nil, zero value otherwise.

### GetCanDeleteTagOk

`func (o *CreatePermissionGroupRequest) GetCanDeleteTagOk() (*bool, bool)`

GetCanDeleteTagOk returns a tuple with the CanDeleteTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteTag

`func (o *CreatePermissionGroupRequest) SetCanDeleteTag(v bool)`

SetCanDeleteTag sets CanDeleteTag field to given value.

### HasCanDeleteTag

`func (o *CreatePermissionGroupRequest) HasCanDeleteTag() bool`

HasCanDeleteTag returns a boolean if a field has been set.

### GetCanManageWebhook

`func (o *CreatePermissionGroupRequest) GetCanManageWebhook() bool`

GetCanManageWebhook returns the CanManageWebhook field if non-nil, zero value otherwise.

### GetCanManageWebhookOk

`func (o *CreatePermissionGroupRequest) GetCanManageWebhookOk() (*bool, bool)`

GetCanManageWebhookOk returns a tuple with the CanManageWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageWebhook

`func (o *CreatePermissionGroupRequest) SetCanManageWebhook(v bool)`

SetCanManageWebhook sets CanManageWebhook field to given value.

### HasCanManageWebhook

`func (o *CreatePermissionGroupRequest) HasCanManageWebhook() bool`

HasCanManageWebhook returns a boolean if a field has been set.

### GetCanUseExportImport

`func (o *CreatePermissionGroupRequest) GetCanUseExportImport() bool`

GetCanUseExportImport returns the CanUseExportImport field if non-nil, zero value otherwise.

### GetCanUseExportImportOk

`func (o *CreatePermissionGroupRequest) GetCanUseExportImportOk() (*bool, bool)`

GetCanUseExportImportOk returns a tuple with the CanUseExportImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUseExportImport

`func (o *CreatePermissionGroupRequest) SetCanUseExportImport(v bool)`

SetCanUseExportImport sets CanUseExportImport field to given value.

### HasCanUseExportImport

`func (o *CreatePermissionGroupRequest) HasCanUseExportImport() bool`

HasCanUseExportImport returns a boolean if a field has been set.

### GetCanManageProductPreferences

`func (o *CreatePermissionGroupRequest) GetCanManageProductPreferences() bool`

GetCanManageProductPreferences returns the CanManageProductPreferences field if non-nil, zero value otherwise.

### GetCanManageProductPreferencesOk

`func (o *CreatePermissionGroupRequest) GetCanManageProductPreferencesOk() (*bool, bool)`

GetCanManageProductPreferencesOk returns a tuple with the CanManageProductPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageProductPreferences

`func (o *CreatePermissionGroupRequest) SetCanManageProductPreferences(v bool)`

SetCanManageProductPreferences sets CanManageProductPreferences field to given value.

### HasCanManageProductPreferences

`func (o *CreatePermissionGroupRequest) HasCanManageProductPreferences() bool`

HasCanManageProductPreferences returns a boolean if a field has been set.

### GetCanManageIntegrations

`func (o *CreatePermissionGroupRequest) GetCanManageIntegrations() bool`

GetCanManageIntegrations returns the CanManageIntegrations field if non-nil, zero value otherwise.

### GetCanManageIntegrationsOk

`func (o *CreatePermissionGroupRequest) GetCanManageIntegrationsOk() (*bool, bool)`

GetCanManageIntegrationsOk returns a tuple with the CanManageIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageIntegrations

`func (o *CreatePermissionGroupRequest) SetCanManageIntegrations(v bool)`

SetCanManageIntegrations sets CanManageIntegrations field to given value.

### HasCanManageIntegrations

`func (o *CreatePermissionGroupRequest) HasCanManageIntegrations() bool`

HasCanManageIntegrations returns a boolean if a field has been set.

### GetCanViewSdkKey

`func (o *CreatePermissionGroupRequest) GetCanViewSdkKey() bool`

GetCanViewSdkKey returns the CanViewSdkKey field if non-nil, zero value otherwise.

### GetCanViewSdkKeyOk

`func (o *CreatePermissionGroupRequest) GetCanViewSdkKeyOk() (*bool, bool)`

GetCanViewSdkKeyOk returns a tuple with the CanViewSdkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewSdkKey

`func (o *CreatePermissionGroupRequest) SetCanViewSdkKey(v bool)`

SetCanViewSdkKey sets CanViewSdkKey field to given value.

### HasCanViewSdkKey

`func (o *CreatePermissionGroupRequest) HasCanViewSdkKey() bool`

HasCanViewSdkKey returns a boolean if a field has been set.

### GetCanRotateSdkKey

`func (o *CreatePermissionGroupRequest) GetCanRotateSdkKey() bool`

GetCanRotateSdkKey returns the CanRotateSdkKey field if non-nil, zero value otherwise.

### GetCanRotateSdkKeyOk

`func (o *CreatePermissionGroupRequest) GetCanRotateSdkKeyOk() (*bool, bool)`

GetCanRotateSdkKeyOk returns a tuple with the CanRotateSdkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRotateSdkKey

`func (o *CreatePermissionGroupRequest) SetCanRotateSdkKey(v bool)`

SetCanRotateSdkKey sets CanRotateSdkKey field to given value.

### HasCanRotateSdkKey

`func (o *CreatePermissionGroupRequest) HasCanRotateSdkKey() bool`

HasCanRotateSdkKey returns a boolean if a field has been set.

### GetCanViewProductStatistics

`func (o *CreatePermissionGroupRequest) GetCanViewProductStatistics() bool`

GetCanViewProductStatistics returns the CanViewProductStatistics field if non-nil, zero value otherwise.

### GetCanViewProductStatisticsOk

`func (o *CreatePermissionGroupRequest) GetCanViewProductStatisticsOk() (*bool, bool)`

GetCanViewProductStatisticsOk returns a tuple with the CanViewProductStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewProductStatistics

`func (o *CreatePermissionGroupRequest) SetCanViewProductStatistics(v bool)`

SetCanViewProductStatistics sets CanViewProductStatistics field to given value.

### HasCanViewProductStatistics

`func (o *CreatePermissionGroupRequest) HasCanViewProductStatistics() bool`

HasCanViewProductStatistics returns a boolean if a field has been set.

### GetCanViewProductAuditLog

`func (o *CreatePermissionGroupRequest) GetCanViewProductAuditLog() bool`

GetCanViewProductAuditLog returns the CanViewProductAuditLog field if non-nil, zero value otherwise.

### GetCanViewProductAuditLogOk

`func (o *CreatePermissionGroupRequest) GetCanViewProductAuditLogOk() (*bool, bool)`

GetCanViewProductAuditLogOk returns a tuple with the CanViewProductAuditLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewProductAuditLog

`func (o *CreatePermissionGroupRequest) SetCanViewProductAuditLog(v bool)`

SetCanViewProductAuditLog sets CanViewProductAuditLog field to given value.

### HasCanViewProductAuditLog

`func (o *CreatePermissionGroupRequest) HasCanViewProductAuditLog() bool`

HasCanViewProductAuditLog returns a boolean if a field has been set.

### GetCanCreateOrUpdateSegments

`func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateSegments() bool`

GetCanCreateOrUpdateSegments returns the CanCreateOrUpdateSegments field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateSegmentsOk

`func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateSegmentsOk() (*bool, bool)`

GetCanCreateOrUpdateSegmentsOk returns a tuple with the CanCreateOrUpdateSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateSegments

`func (o *CreatePermissionGroupRequest) SetCanCreateOrUpdateSegments(v bool)`

SetCanCreateOrUpdateSegments sets CanCreateOrUpdateSegments field to given value.

### HasCanCreateOrUpdateSegments

`func (o *CreatePermissionGroupRequest) HasCanCreateOrUpdateSegments() bool`

HasCanCreateOrUpdateSegments returns a boolean if a field has been set.

### GetCanDeleteSegments

`func (o *CreatePermissionGroupRequest) GetCanDeleteSegments() bool`

GetCanDeleteSegments returns the CanDeleteSegments field if non-nil, zero value otherwise.

### GetCanDeleteSegmentsOk

`func (o *CreatePermissionGroupRequest) GetCanDeleteSegmentsOk() (*bool, bool)`

GetCanDeleteSegmentsOk returns a tuple with the CanDeleteSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteSegments

`func (o *CreatePermissionGroupRequest) SetCanDeleteSegments(v bool)`

SetCanDeleteSegments sets CanDeleteSegments field to given value.

### HasCanDeleteSegments

`func (o *CreatePermissionGroupRequest) HasCanDeleteSegments() bool`

HasCanDeleteSegments returns a boolean if a field has been set.

### GetAccessType

`func (o *CreatePermissionGroupRequest) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CreatePermissionGroupRequest) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CreatePermissionGroupRequest) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *CreatePermissionGroupRequest) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetNewEnvironmentAccessType

`func (o *CreatePermissionGroupRequest) GetNewEnvironmentAccessType() EnvironmentAccessType`

GetNewEnvironmentAccessType returns the NewEnvironmentAccessType field if non-nil, zero value otherwise.

### GetNewEnvironmentAccessTypeOk

`func (o *CreatePermissionGroupRequest) GetNewEnvironmentAccessTypeOk() (*EnvironmentAccessType, bool)`

GetNewEnvironmentAccessTypeOk returns a tuple with the NewEnvironmentAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEnvironmentAccessType

`func (o *CreatePermissionGroupRequest) SetNewEnvironmentAccessType(v EnvironmentAccessType)`

SetNewEnvironmentAccessType sets NewEnvironmentAccessType field to given value.

### HasNewEnvironmentAccessType

`func (o *CreatePermissionGroupRequest) HasNewEnvironmentAccessType() bool`

HasNewEnvironmentAccessType returns a boolean if a field has been set.

### GetEnvironmentAccesses

`func (o *CreatePermissionGroupRequest) GetEnvironmentAccesses() []CreateOrUpdateEnvironmentAccessModel`

GetEnvironmentAccesses returns the EnvironmentAccesses field if non-nil, zero value otherwise.

### GetEnvironmentAccessesOk

`func (o *CreatePermissionGroupRequest) GetEnvironmentAccessesOk() (*[]CreateOrUpdateEnvironmentAccessModel, bool)`

GetEnvironmentAccessesOk returns a tuple with the EnvironmentAccesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentAccesses

`func (o *CreatePermissionGroupRequest) SetEnvironmentAccesses(v []CreateOrUpdateEnvironmentAccessModel)`

SetEnvironmentAccesses sets EnvironmentAccesses field to given value.

### HasEnvironmentAccesses

`func (o *CreatePermissionGroupRequest) HasEnvironmentAccesses() bool`

HasEnvironmentAccesses returns a boolean if a field has been set.

### SetEnvironmentAccessesNil

`func (o *CreatePermissionGroupRequest) SetEnvironmentAccessesNil(b bool)`

 SetEnvironmentAccessesNil sets the value for EnvironmentAccesses to be an explicit nil

### UnsetEnvironmentAccesses
`func (o *CreatePermissionGroupRequest) UnsetEnvironmentAccesses()`

UnsetEnvironmentAccesses ensures that no value is present for EnvironmentAccesses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


