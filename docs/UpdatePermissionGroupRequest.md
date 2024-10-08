# UpdatePermissionGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the Permission Group. | [optional] 
**CanManageMembers** | Pointer to **NullableBool** | Group members can manage team members. | [optional] 
**CanCreateOrUpdateConfig** | Pointer to **NullableBool** | Group members can create/update Configs. | [optional] 
**CanDeleteConfig** | Pointer to **NullableBool** | Group members can delete Configs. | [optional] 
**CanCreateOrUpdateEnvironment** | Pointer to **NullableBool** | Group members can create/update Environments. | [optional] 
**CanDeleteEnvironment** | Pointer to **NullableBool** | Group members can delete Environments. | [optional] 
**CanCreateOrUpdateSetting** | Pointer to **NullableBool** | Group members can create/update Feature Flags and Settings. | [optional] 
**CanTagSetting** | Pointer to **NullableBool** | Group members can attach/detach Tags to Feature Flags and Settings. | [optional] 
**CanDeleteSetting** | Pointer to **NullableBool** | Group members can delete Feature Flags and Settings. | [optional] 
**CanCreateOrUpdateTag** | Pointer to **NullableBool** | Group members can create/update Tags. | [optional] 
**CanDeleteTag** | Pointer to **NullableBool** | Group members can delete Tags. | [optional] 
**CanManageWebhook** | Pointer to **NullableBool** | Group members can create/update/delete Webhooks. | [optional] 
**CanUseExportImport** | Pointer to **NullableBool** | Group members can use the export/import feature. | [optional] 
**CanManageProductPreferences** | Pointer to **NullableBool** | Group members can update Product preferences. | [optional] 
**CanManageIntegrations** | Pointer to **NullableBool** | Group members can add and configure integrations. | [optional] 
**CanViewSdkKey** | Pointer to **NullableBool** | Group members has access to SDK keys. | [optional] 
**CanRotateSdkKey** | Pointer to **NullableBool** | Group members can rotate SDK keys. | [optional] 
**CanCreateOrUpdateSegments** | Pointer to **NullableBool** | Group members can create/update Segments. | [optional] 
**CanDeleteSegments** | Pointer to **NullableBool** | Group members can delete Segments. | [optional] 
**CanViewProductAuditLog** | Pointer to **NullableBool** | Group members has access to audit logs. | [optional] 
**CanViewProductStatistics** | Pointer to **NullableBool** | Group members has access to product statistics. | [optional] 
**CanDisable2FA** | Pointer to **NullableBool** | Group members can disable two-factor authentication for other members. | [optional] 
**AccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**NewEnvironmentAccessType** | Pointer to [**EnvironmentAccessType**](EnvironmentAccessType.md) |  | [optional] 
**EnvironmentAccesses** | Pointer to [**[]CreateOrUpdateEnvironmentAccessModel**](CreateOrUpdateEnvironmentAccessModel.md) | List of environment specific permissions. | [optional] 

## Methods

### NewUpdatePermissionGroupRequest

`func NewUpdatePermissionGroupRequest() *UpdatePermissionGroupRequest`

NewUpdatePermissionGroupRequest instantiates a new UpdatePermissionGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePermissionGroupRequestWithDefaults

`func NewUpdatePermissionGroupRequestWithDefaults() *UpdatePermissionGroupRequest`

NewUpdatePermissionGroupRequestWithDefaults instantiates a new UpdatePermissionGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePermissionGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePermissionGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePermissionGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePermissionGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdatePermissionGroupRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdatePermissionGroupRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCanManageMembers

`func (o *UpdatePermissionGroupRequest) GetCanManageMembers() bool`

GetCanManageMembers returns the CanManageMembers field if non-nil, zero value otherwise.

### GetCanManageMembersOk

`func (o *UpdatePermissionGroupRequest) GetCanManageMembersOk() (*bool, bool)`

GetCanManageMembersOk returns a tuple with the CanManageMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageMembers

`func (o *UpdatePermissionGroupRequest) SetCanManageMembers(v bool)`

SetCanManageMembers sets CanManageMembers field to given value.

### HasCanManageMembers

`func (o *UpdatePermissionGroupRequest) HasCanManageMembers() bool`

HasCanManageMembers returns a boolean if a field has been set.

### SetCanManageMembersNil

`func (o *UpdatePermissionGroupRequest) SetCanManageMembersNil(b bool)`

 SetCanManageMembersNil sets the value for CanManageMembers to be an explicit nil

### UnsetCanManageMembers
`func (o *UpdatePermissionGroupRequest) UnsetCanManageMembers()`

UnsetCanManageMembers ensures that no value is present for CanManageMembers, not even an explicit nil
### GetCanCreateOrUpdateConfig

`func (o *UpdatePermissionGroupRequest) GetCanCreateOrUpdateConfig() bool`

GetCanCreateOrUpdateConfig returns the CanCreateOrUpdateConfig field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateConfigOk

`func (o *UpdatePermissionGroupRequest) GetCanCreateOrUpdateConfigOk() (*bool, bool)`

GetCanCreateOrUpdateConfigOk returns a tuple with the CanCreateOrUpdateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateConfig

`func (o *UpdatePermissionGroupRequest) SetCanCreateOrUpdateConfig(v bool)`

SetCanCreateOrUpdateConfig sets CanCreateOrUpdateConfig field to given value.

### HasCanCreateOrUpdateConfig

`func (o *UpdatePermissionGroupRequest) HasCanCreateOrUpdateConfig() bool`

HasCanCreateOrUpdateConfig returns a boolean if a field has been set.

### SetCanCreateOrUpdateConfigNil

`func (o *UpdatePermissionGroupRequest) SetCanCreateOrUpdateConfigNil(b bool)`

 SetCanCreateOrUpdateConfigNil sets the value for CanCreateOrUpdateConfig to be an explicit nil

### UnsetCanCreateOrUpdateConfig
`func (o *UpdatePermissionGroupRequest) UnsetCanCreateOrUpdateConfig()`

UnsetCanCreateOrUpdateConfig ensures that no value is present for CanCreateOrUpdateConfig, not even an explicit nil
### GetCanDeleteConfig

`func (o *UpdatePermissionGroupRequest) GetCanDeleteConfig() bool`

GetCanDeleteConfig returns the CanDeleteConfig field if non-nil, zero value otherwise.

### GetCanDeleteConfigOk

`func (o *UpdatePermissionGroupRequest) GetCanDeleteConfigOk() (*bool, bool)`

GetCanDeleteConfigOk returns a tuple with the CanDeleteConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteConfig

`func (o *UpdatePermissionGroupRequest) SetCanDeleteConfig(v bool)`

SetCanDeleteConfig sets CanDeleteConfig field to given value.

### HasCanDeleteConfig

`func (o *UpdatePermissionGroupRequest) HasCanDeleteConfig() bool`

HasCanDeleteConfig returns a boolean if a field has been set.

### SetCanDeleteConfigNil

`func (o *UpdatePermissionGroupRequest) SetCanDeleteConfigNil(b bool)`

 SetCanDeleteConfigNil sets the value for CanDeleteConfig to be an explicit nil

### UnsetCanDeleteConfig
`func (o *UpdatePermissionGroupRequest) UnsetCanDeleteConfig()`

UnsetCanDeleteConfig ensures that no value is present for CanDeleteConfig, not even an explicit nil
### GetCanCreateOrUpdateEnvironment

`func (o *UpdatePermissionGroupRequest) GetCanCreateOrUpdateEnvironment() bool`

GetCanCreateOrUpdateEnvironment returns the CanCreateOrUpdateEnvironment field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateEnvironmentOk

`func (o *UpdatePermissionGroupRequest) GetCanCreateOrUpdateEnvironmentOk() (*bool, bool)`

GetCanCreateOrUpdateEnvironmentOk returns a tuple with the CanCreateOrUpdateEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateEnvironment

`func (o *UpdatePermissionGroupRequest) SetCanCreateOrUpdateEnvironment(v bool)`

SetCanCreateOrUpdateEnvironment sets CanCreateOrUpdateEnvironment field to given value.

### HasCanCreateOrUpdateEnvironment

`func (o *UpdatePermissionGroupRequest) HasCanCreateOrUpdateEnvironment() bool`

HasCanCreateOrUpdateEnvironment returns a boolean if a field has been set.

### SetCanCreateOrUpdateEnvironmentNil

`func (o *UpdatePermissionGroupRequest) SetCanCreateOrUpdateEnvironmentNil(b bool)`

 SetCanCreateOrUpdateEnvironmentNil sets the value for CanCreateOrUpdateEnvironment to be an explicit nil

### UnsetCanCreateOrUpdateEnvironment
`func (o *UpdatePermissionGroupRequest) UnsetCanCreateOrUpdateEnvironment()`

UnsetCanCreateOrUpdateEnvironment ensures that no value is present for CanCreateOrUpdateEnvironment, not even an explicit nil
### GetCanDeleteEnvironment

`func (o *UpdatePermissionGroupRequest) GetCanDeleteEnvironment() bool`

GetCanDeleteEnvironment returns the CanDeleteEnvironment field if non-nil, zero value otherwise.

### GetCanDeleteEnvironmentOk

`func (o *UpdatePermissionGroupRequest) GetCanDeleteEnvironmentOk() (*bool, bool)`

GetCanDeleteEnvironmentOk returns a tuple with the CanDeleteEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteEnvironment

`func (o *UpdatePermissionGroupRequest) SetCanDeleteEnvironment(v bool)`

SetCanDeleteEnvironment sets CanDeleteEnvironment field to given value.

### HasCanDeleteEnvironment

`func (o *UpdatePermissionGroupRequest) HasCanDeleteEnvironment() bool`

HasCanDeleteEnvironment returns a boolean if a field has been set.

### SetCanDeleteEnvironmentNil

`func (o *UpdatePermissionGroupRequest) SetCanDeleteEnvironmentNil(b bool)`

 SetCanDeleteEnvironmentNil sets the value for CanDeleteEnvironment to be an explicit nil

### UnsetCanDeleteEnvironment
`func (o *UpdatePermissionGroupRequest) UnsetCanDeleteEnvironment()`

UnsetCanDeleteEnvironment ensures that no value is present for CanDeleteEnvironment, not even an explicit nil
### GetCanCreateOrUpdateSetting

`func (o *UpdatePermissionGroupRequest) GetCanCreateOrUpdateSetting() bool`

GetCanCreateOrUpdateSetting returns the CanCreateOrUpdateSetting field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateSettingOk

`func (o *UpdatePermissionGroupRequest) GetCanCreateOrUpdateSettingOk() (*bool, bool)`

GetCanCreateOrUpdateSettingOk returns a tuple with the CanCreateOrUpdateSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateSetting

`func (o *UpdatePermissionGroupRequest) SetCanCreateOrUpdateSetting(v bool)`

SetCanCreateOrUpdateSetting sets CanCreateOrUpdateSetting field to given value.

### HasCanCreateOrUpdateSetting

`func (o *UpdatePermissionGroupRequest) HasCanCreateOrUpdateSetting() bool`

HasCanCreateOrUpdateSetting returns a boolean if a field has been set.

### SetCanCreateOrUpdateSettingNil

`func (o *UpdatePermissionGroupRequest) SetCanCreateOrUpdateSettingNil(b bool)`

 SetCanCreateOrUpdateSettingNil sets the value for CanCreateOrUpdateSetting to be an explicit nil

### UnsetCanCreateOrUpdateSetting
`func (o *UpdatePermissionGroupRequest) UnsetCanCreateOrUpdateSetting()`

UnsetCanCreateOrUpdateSetting ensures that no value is present for CanCreateOrUpdateSetting, not even an explicit nil
### GetCanTagSetting

`func (o *UpdatePermissionGroupRequest) GetCanTagSetting() bool`

GetCanTagSetting returns the CanTagSetting field if non-nil, zero value otherwise.

### GetCanTagSettingOk

`func (o *UpdatePermissionGroupRequest) GetCanTagSettingOk() (*bool, bool)`

GetCanTagSettingOk returns a tuple with the CanTagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTagSetting

`func (o *UpdatePermissionGroupRequest) SetCanTagSetting(v bool)`

SetCanTagSetting sets CanTagSetting field to given value.

### HasCanTagSetting

`func (o *UpdatePermissionGroupRequest) HasCanTagSetting() bool`

HasCanTagSetting returns a boolean if a field has been set.

### SetCanTagSettingNil

`func (o *UpdatePermissionGroupRequest) SetCanTagSettingNil(b bool)`

 SetCanTagSettingNil sets the value for CanTagSetting to be an explicit nil

### UnsetCanTagSetting
`func (o *UpdatePermissionGroupRequest) UnsetCanTagSetting()`

UnsetCanTagSetting ensures that no value is present for CanTagSetting, not even an explicit nil
### GetCanDeleteSetting

`func (o *UpdatePermissionGroupRequest) GetCanDeleteSetting() bool`

GetCanDeleteSetting returns the CanDeleteSetting field if non-nil, zero value otherwise.

### GetCanDeleteSettingOk

`func (o *UpdatePermissionGroupRequest) GetCanDeleteSettingOk() (*bool, bool)`

GetCanDeleteSettingOk returns a tuple with the CanDeleteSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteSetting

`func (o *UpdatePermissionGroupRequest) SetCanDeleteSetting(v bool)`

SetCanDeleteSetting sets CanDeleteSetting field to given value.

### HasCanDeleteSetting

`func (o *UpdatePermissionGroupRequest) HasCanDeleteSetting() bool`

HasCanDeleteSetting returns a boolean if a field has been set.

### SetCanDeleteSettingNil

`func (o *UpdatePermissionGroupRequest) SetCanDeleteSettingNil(b bool)`

 SetCanDeleteSettingNil sets the value for CanDeleteSetting to be an explicit nil

### UnsetCanDeleteSetting
`func (o *UpdatePermissionGroupRequest) UnsetCanDeleteSetting()`

UnsetCanDeleteSetting ensures that no value is present for CanDeleteSetting, not even an explicit nil
### GetCanCreateOrUpdateTag

`func (o *UpdatePermissionGroupRequest) GetCanCreateOrUpdateTag() bool`

GetCanCreateOrUpdateTag returns the CanCreateOrUpdateTag field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateTagOk

`func (o *UpdatePermissionGroupRequest) GetCanCreateOrUpdateTagOk() (*bool, bool)`

GetCanCreateOrUpdateTagOk returns a tuple with the CanCreateOrUpdateTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateTag

`func (o *UpdatePermissionGroupRequest) SetCanCreateOrUpdateTag(v bool)`

SetCanCreateOrUpdateTag sets CanCreateOrUpdateTag field to given value.

### HasCanCreateOrUpdateTag

`func (o *UpdatePermissionGroupRequest) HasCanCreateOrUpdateTag() bool`

HasCanCreateOrUpdateTag returns a boolean if a field has been set.

### SetCanCreateOrUpdateTagNil

`func (o *UpdatePermissionGroupRequest) SetCanCreateOrUpdateTagNil(b bool)`

 SetCanCreateOrUpdateTagNil sets the value for CanCreateOrUpdateTag to be an explicit nil

### UnsetCanCreateOrUpdateTag
`func (o *UpdatePermissionGroupRequest) UnsetCanCreateOrUpdateTag()`

UnsetCanCreateOrUpdateTag ensures that no value is present for CanCreateOrUpdateTag, not even an explicit nil
### GetCanDeleteTag

`func (o *UpdatePermissionGroupRequest) GetCanDeleteTag() bool`

GetCanDeleteTag returns the CanDeleteTag field if non-nil, zero value otherwise.

### GetCanDeleteTagOk

`func (o *UpdatePermissionGroupRequest) GetCanDeleteTagOk() (*bool, bool)`

GetCanDeleteTagOk returns a tuple with the CanDeleteTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteTag

`func (o *UpdatePermissionGroupRequest) SetCanDeleteTag(v bool)`

SetCanDeleteTag sets CanDeleteTag field to given value.

### HasCanDeleteTag

`func (o *UpdatePermissionGroupRequest) HasCanDeleteTag() bool`

HasCanDeleteTag returns a boolean if a field has been set.

### SetCanDeleteTagNil

`func (o *UpdatePermissionGroupRequest) SetCanDeleteTagNil(b bool)`

 SetCanDeleteTagNil sets the value for CanDeleteTag to be an explicit nil

### UnsetCanDeleteTag
`func (o *UpdatePermissionGroupRequest) UnsetCanDeleteTag()`

UnsetCanDeleteTag ensures that no value is present for CanDeleteTag, not even an explicit nil
### GetCanManageWebhook

`func (o *UpdatePermissionGroupRequest) GetCanManageWebhook() bool`

GetCanManageWebhook returns the CanManageWebhook field if non-nil, zero value otherwise.

### GetCanManageWebhookOk

`func (o *UpdatePermissionGroupRequest) GetCanManageWebhookOk() (*bool, bool)`

GetCanManageWebhookOk returns a tuple with the CanManageWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageWebhook

`func (o *UpdatePermissionGroupRequest) SetCanManageWebhook(v bool)`

SetCanManageWebhook sets CanManageWebhook field to given value.

### HasCanManageWebhook

`func (o *UpdatePermissionGroupRequest) HasCanManageWebhook() bool`

HasCanManageWebhook returns a boolean if a field has been set.

### SetCanManageWebhookNil

`func (o *UpdatePermissionGroupRequest) SetCanManageWebhookNil(b bool)`

 SetCanManageWebhookNil sets the value for CanManageWebhook to be an explicit nil

### UnsetCanManageWebhook
`func (o *UpdatePermissionGroupRequest) UnsetCanManageWebhook()`

UnsetCanManageWebhook ensures that no value is present for CanManageWebhook, not even an explicit nil
### GetCanUseExportImport

`func (o *UpdatePermissionGroupRequest) GetCanUseExportImport() bool`

GetCanUseExportImport returns the CanUseExportImport field if non-nil, zero value otherwise.

### GetCanUseExportImportOk

`func (o *UpdatePermissionGroupRequest) GetCanUseExportImportOk() (*bool, bool)`

GetCanUseExportImportOk returns a tuple with the CanUseExportImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUseExportImport

`func (o *UpdatePermissionGroupRequest) SetCanUseExportImport(v bool)`

SetCanUseExportImport sets CanUseExportImport field to given value.

### HasCanUseExportImport

`func (o *UpdatePermissionGroupRequest) HasCanUseExportImport() bool`

HasCanUseExportImport returns a boolean if a field has been set.

### SetCanUseExportImportNil

`func (o *UpdatePermissionGroupRequest) SetCanUseExportImportNil(b bool)`

 SetCanUseExportImportNil sets the value for CanUseExportImport to be an explicit nil

### UnsetCanUseExportImport
`func (o *UpdatePermissionGroupRequest) UnsetCanUseExportImport()`

UnsetCanUseExportImport ensures that no value is present for CanUseExportImport, not even an explicit nil
### GetCanManageProductPreferences

`func (o *UpdatePermissionGroupRequest) GetCanManageProductPreferences() bool`

GetCanManageProductPreferences returns the CanManageProductPreferences field if non-nil, zero value otherwise.

### GetCanManageProductPreferencesOk

`func (o *UpdatePermissionGroupRequest) GetCanManageProductPreferencesOk() (*bool, bool)`

GetCanManageProductPreferencesOk returns a tuple with the CanManageProductPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageProductPreferences

`func (o *UpdatePermissionGroupRequest) SetCanManageProductPreferences(v bool)`

SetCanManageProductPreferences sets CanManageProductPreferences field to given value.

### HasCanManageProductPreferences

`func (o *UpdatePermissionGroupRequest) HasCanManageProductPreferences() bool`

HasCanManageProductPreferences returns a boolean if a field has been set.

### SetCanManageProductPreferencesNil

`func (o *UpdatePermissionGroupRequest) SetCanManageProductPreferencesNil(b bool)`

 SetCanManageProductPreferencesNil sets the value for CanManageProductPreferences to be an explicit nil

### UnsetCanManageProductPreferences
`func (o *UpdatePermissionGroupRequest) UnsetCanManageProductPreferences()`

UnsetCanManageProductPreferences ensures that no value is present for CanManageProductPreferences, not even an explicit nil
### GetCanManageIntegrations

`func (o *UpdatePermissionGroupRequest) GetCanManageIntegrations() bool`

GetCanManageIntegrations returns the CanManageIntegrations field if non-nil, zero value otherwise.

### GetCanManageIntegrationsOk

`func (o *UpdatePermissionGroupRequest) GetCanManageIntegrationsOk() (*bool, bool)`

GetCanManageIntegrationsOk returns a tuple with the CanManageIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageIntegrations

`func (o *UpdatePermissionGroupRequest) SetCanManageIntegrations(v bool)`

SetCanManageIntegrations sets CanManageIntegrations field to given value.

### HasCanManageIntegrations

`func (o *UpdatePermissionGroupRequest) HasCanManageIntegrations() bool`

HasCanManageIntegrations returns a boolean if a field has been set.

### SetCanManageIntegrationsNil

`func (o *UpdatePermissionGroupRequest) SetCanManageIntegrationsNil(b bool)`

 SetCanManageIntegrationsNil sets the value for CanManageIntegrations to be an explicit nil

### UnsetCanManageIntegrations
`func (o *UpdatePermissionGroupRequest) UnsetCanManageIntegrations()`

UnsetCanManageIntegrations ensures that no value is present for CanManageIntegrations, not even an explicit nil
### GetCanViewSdkKey

`func (o *UpdatePermissionGroupRequest) GetCanViewSdkKey() bool`

GetCanViewSdkKey returns the CanViewSdkKey field if non-nil, zero value otherwise.

### GetCanViewSdkKeyOk

`func (o *UpdatePermissionGroupRequest) GetCanViewSdkKeyOk() (*bool, bool)`

GetCanViewSdkKeyOk returns a tuple with the CanViewSdkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewSdkKey

`func (o *UpdatePermissionGroupRequest) SetCanViewSdkKey(v bool)`

SetCanViewSdkKey sets CanViewSdkKey field to given value.

### HasCanViewSdkKey

`func (o *UpdatePermissionGroupRequest) HasCanViewSdkKey() bool`

HasCanViewSdkKey returns a boolean if a field has been set.

### SetCanViewSdkKeyNil

`func (o *UpdatePermissionGroupRequest) SetCanViewSdkKeyNil(b bool)`

 SetCanViewSdkKeyNil sets the value for CanViewSdkKey to be an explicit nil

### UnsetCanViewSdkKey
`func (o *UpdatePermissionGroupRequest) UnsetCanViewSdkKey()`

UnsetCanViewSdkKey ensures that no value is present for CanViewSdkKey, not even an explicit nil
### GetCanRotateSdkKey

`func (o *UpdatePermissionGroupRequest) GetCanRotateSdkKey() bool`

GetCanRotateSdkKey returns the CanRotateSdkKey field if non-nil, zero value otherwise.

### GetCanRotateSdkKeyOk

`func (o *UpdatePermissionGroupRequest) GetCanRotateSdkKeyOk() (*bool, bool)`

GetCanRotateSdkKeyOk returns a tuple with the CanRotateSdkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRotateSdkKey

`func (o *UpdatePermissionGroupRequest) SetCanRotateSdkKey(v bool)`

SetCanRotateSdkKey sets CanRotateSdkKey field to given value.

### HasCanRotateSdkKey

`func (o *UpdatePermissionGroupRequest) HasCanRotateSdkKey() bool`

HasCanRotateSdkKey returns a boolean if a field has been set.

### SetCanRotateSdkKeyNil

`func (o *UpdatePermissionGroupRequest) SetCanRotateSdkKeyNil(b bool)`

 SetCanRotateSdkKeyNil sets the value for CanRotateSdkKey to be an explicit nil

### UnsetCanRotateSdkKey
`func (o *UpdatePermissionGroupRequest) UnsetCanRotateSdkKey()`

UnsetCanRotateSdkKey ensures that no value is present for CanRotateSdkKey, not even an explicit nil
### GetCanCreateOrUpdateSegments

`func (o *UpdatePermissionGroupRequest) GetCanCreateOrUpdateSegments() bool`

GetCanCreateOrUpdateSegments returns the CanCreateOrUpdateSegments field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateSegmentsOk

`func (o *UpdatePermissionGroupRequest) GetCanCreateOrUpdateSegmentsOk() (*bool, bool)`

GetCanCreateOrUpdateSegmentsOk returns a tuple with the CanCreateOrUpdateSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateSegments

`func (o *UpdatePermissionGroupRequest) SetCanCreateOrUpdateSegments(v bool)`

SetCanCreateOrUpdateSegments sets CanCreateOrUpdateSegments field to given value.

### HasCanCreateOrUpdateSegments

`func (o *UpdatePermissionGroupRequest) HasCanCreateOrUpdateSegments() bool`

HasCanCreateOrUpdateSegments returns a boolean if a field has been set.

### SetCanCreateOrUpdateSegmentsNil

`func (o *UpdatePermissionGroupRequest) SetCanCreateOrUpdateSegmentsNil(b bool)`

 SetCanCreateOrUpdateSegmentsNil sets the value for CanCreateOrUpdateSegments to be an explicit nil

### UnsetCanCreateOrUpdateSegments
`func (o *UpdatePermissionGroupRequest) UnsetCanCreateOrUpdateSegments()`

UnsetCanCreateOrUpdateSegments ensures that no value is present for CanCreateOrUpdateSegments, not even an explicit nil
### GetCanDeleteSegments

`func (o *UpdatePermissionGroupRequest) GetCanDeleteSegments() bool`

GetCanDeleteSegments returns the CanDeleteSegments field if non-nil, zero value otherwise.

### GetCanDeleteSegmentsOk

`func (o *UpdatePermissionGroupRequest) GetCanDeleteSegmentsOk() (*bool, bool)`

GetCanDeleteSegmentsOk returns a tuple with the CanDeleteSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteSegments

`func (o *UpdatePermissionGroupRequest) SetCanDeleteSegments(v bool)`

SetCanDeleteSegments sets CanDeleteSegments field to given value.

### HasCanDeleteSegments

`func (o *UpdatePermissionGroupRequest) HasCanDeleteSegments() bool`

HasCanDeleteSegments returns a boolean if a field has been set.

### SetCanDeleteSegmentsNil

`func (o *UpdatePermissionGroupRequest) SetCanDeleteSegmentsNil(b bool)`

 SetCanDeleteSegmentsNil sets the value for CanDeleteSegments to be an explicit nil

### UnsetCanDeleteSegments
`func (o *UpdatePermissionGroupRequest) UnsetCanDeleteSegments()`

UnsetCanDeleteSegments ensures that no value is present for CanDeleteSegments, not even an explicit nil
### GetCanViewProductAuditLog

`func (o *UpdatePermissionGroupRequest) GetCanViewProductAuditLog() bool`

GetCanViewProductAuditLog returns the CanViewProductAuditLog field if non-nil, zero value otherwise.

### GetCanViewProductAuditLogOk

`func (o *UpdatePermissionGroupRequest) GetCanViewProductAuditLogOk() (*bool, bool)`

GetCanViewProductAuditLogOk returns a tuple with the CanViewProductAuditLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewProductAuditLog

`func (o *UpdatePermissionGroupRequest) SetCanViewProductAuditLog(v bool)`

SetCanViewProductAuditLog sets CanViewProductAuditLog field to given value.

### HasCanViewProductAuditLog

`func (o *UpdatePermissionGroupRequest) HasCanViewProductAuditLog() bool`

HasCanViewProductAuditLog returns a boolean if a field has been set.

### SetCanViewProductAuditLogNil

`func (o *UpdatePermissionGroupRequest) SetCanViewProductAuditLogNil(b bool)`

 SetCanViewProductAuditLogNil sets the value for CanViewProductAuditLog to be an explicit nil

### UnsetCanViewProductAuditLog
`func (o *UpdatePermissionGroupRequest) UnsetCanViewProductAuditLog()`

UnsetCanViewProductAuditLog ensures that no value is present for CanViewProductAuditLog, not even an explicit nil
### GetCanViewProductStatistics

`func (o *UpdatePermissionGroupRequest) GetCanViewProductStatistics() bool`

GetCanViewProductStatistics returns the CanViewProductStatistics field if non-nil, zero value otherwise.

### GetCanViewProductStatisticsOk

`func (o *UpdatePermissionGroupRequest) GetCanViewProductStatisticsOk() (*bool, bool)`

GetCanViewProductStatisticsOk returns a tuple with the CanViewProductStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewProductStatistics

`func (o *UpdatePermissionGroupRequest) SetCanViewProductStatistics(v bool)`

SetCanViewProductStatistics sets CanViewProductStatistics field to given value.

### HasCanViewProductStatistics

`func (o *UpdatePermissionGroupRequest) HasCanViewProductStatistics() bool`

HasCanViewProductStatistics returns a boolean if a field has been set.

### SetCanViewProductStatisticsNil

`func (o *UpdatePermissionGroupRequest) SetCanViewProductStatisticsNil(b bool)`

 SetCanViewProductStatisticsNil sets the value for CanViewProductStatistics to be an explicit nil

### UnsetCanViewProductStatistics
`func (o *UpdatePermissionGroupRequest) UnsetCanViewProductStatistics()`

UnsetCanViewProductStatistics ensures that no value is present for CanViewProductStatistics, not even an explicit nil
### GetCanDisable2FA

`func (o *UpdatePermissionGroupRequest) GetCanDisable2FA() bool`

GetCanDisable2FA returns the CanDisable2FA field if non-nil, zero value otherwise.

### GetCanDisable2FAOk

`func (o *UpdatePermissionGroupRequest) GetCanDisable2FAOk() (*bool, bool)`

GetCanDisable2FAOk returns a tuple with the CanDisable2FA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDisable2FA

`func (o *UpdatePermissionGroupRequest) SetCanDisable2FA(v bool)`

SetCanDisable2FA sets CanDisable2FA field to given value.

### HasCanDisable2FA

`func (o *UpdatePermissionGroupRequest) HasCanDisable2FA() bool`

HasCanDisable2FA returns a boolean if a field has been set.

### SetCanDisable2FANil

`func (o *UpdatePermissionGroupRequest) SetCanDisable2FANil(b bool)`

 SetCanDisable2FANil sets the value for CanDisable2FA to be an explicit nil

### UnsetCanDisable2FA
`func (o *UpdatePermissionGroupRequest) UnsetCanDisable2FA()`

UnsetCanDisable2FA ensures that no value is present for CanDisable2FA, not even an explicit nil
### GetAccessType

`func (o *UpdatePermissionGroupRequest) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *UpdatePermissionGroupRequest) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *UpdatePermissionGroupRequest) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *UpdatePermissionGroupRequest) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetNewEnvironmentAccessType

`func (o *UpdatePermissionGroupRequest) GetNewEnvironmentAccessType() EnvironmentAccessType`

GetNewEnvironmentAccessType returns the NewEnvironmentAccessType field if non-nil, zero value otherwise.

### GetNewEnvironmentAccessTypeOk

`func (o *UpdatePermissionGroupRequest) GetNewEnvironmentAccessTypeOk() (*EnvironmentAccessType, bool)`

GetNewEnvironmentAccessTypeOk returns a tuple with the NewEnvironmentAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEnvironmentAccessType

`func (o *UpdatePermissionGroupRequest) SetNewEnvironmentAccessType(v EnvironmentAccessType)`

SetNewEnvironmentAccessType sets NewEnvironmentAccessType field to given value.

### HasNewEnvironmentAccessType

`func (o *UpdatePermissionGroupRequest) HasNewEnvironmentAccessType() bool`

HasNewEnvironmentAccessType returns a boolean if a field has been set.

### GetEnvironmentAccesses

`func (o *UpdatePermissionGroupRequest) GetEnvironmentAccesses() []CreateOrUpdateEnvironmentAccessModel`

GetEnvironmentAccesses returns the EnvironmentAccesses field if non-nil, zero value otherwise.

### GetEnvironmentAccessesOk

`func (o *UpdatePermissionGroupRequest) GetEnvironmentAccessesOk() (*[]CreateOrUpdateEnvironmentAccessModel, bool)`

GetEnvironmentAccessesOk returns a tuple with the EnvironmentAccesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentAccesses

`func (o *UpdatePermissionGroupRequest) SetEnvironmentAccesses(v []CreateOrUpdateEnvironmentAccessModel)`

SetEnvironmentAccesses sets EnvironmentAccesses field to given value.

### HasEnvironmentAccesses

`func (o *UpdatePermissionGroupRequest) HasEnvironmentAccesses() bool`

HasEnvironmentAccesses returns a boolean if a field has been set.

### SetEnvironmentAccessesNil

`func (o *UpdatePermissionGroupRequest) SetEnvironmentAccessesNil(b bool)`

 SetEnvironmentAccessesNil sets the value for EnvironmentAccesses to be an explicit nil

### UnsetEnvironmentAccesses
`func (o *UpdatePermissionGroupRequest) UnsetEnvironmentAccesses()`

UnsetEnvironmentAccesses ensures that no value is present for EnvironmentAccesses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


