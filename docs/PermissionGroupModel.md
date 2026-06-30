# PermissionGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionGroupId** | **int64** | Identifier of the Permission Group. | 
**Name** | **string** | Name of the Permission Group. | 
**CanManageMembers** | **bool** | Group members can manage team members. | 
**CanCreateOrUpdateConfig** | **bool** | Group members can create/update Configs. | 
**CanDeleteConfig** | **bool** | Group members can delete Configs. | 
**CanCreateOrUpdateEnvironment** | **bool** | Group members can create/update Environments. | 
**CanDeleteEnvironment** | **bool** | Group members can delete Environments. | 
**CanCreateOrUpdateSetting** | **bool** | Group members can create/update Feature Flags and Settings. | 
**CanTagSetting** | **bool** | Group members can attach/detach Tags to Feature Flags and Settings. | 
**CanDeleteSetting** | **bool** | Group members can delete Feature Flags and Settings. | 
**CanCreateOrUpdateTag** | **bool** | Group members can create/update Tags. | 
**CanDeleteTag** | **bool** | Group members can delete Tags. | 
**CanManageWebhook** | **bool** | Group members can create/update/delete Webhooks. | 
**CanUseExportImport** | **bool** | Group members can use the export/import feature. | 
**CanManageProductPreferences** | **bool** | Group members can update Product preferences. | 
**CanManageIntegrations** | **bool** | Group members can add and configure integrations. | 
**CanViewSdkKey** | **bool** | Group members has access to SDK keys. | 
**CanRotateSdkKey** | **bool** | Group members can rotate SDK keys. | 
**CanCreateOrUpdateSegments** | **bool** | Group members can create/update Segments. | 
**CanDeleteSegments** | **bool** | Group members can delete Segments. | 
**CanViewProductAuditLog** | **bool** | Group members has access to audit logs. | 
**CanViewProductStatistics** | **bool** | Group members has access to product statistics. | 
**CanDisable2FA** | **bool** | Group members can disable two-factor authentication for other members. | 
**AccessType** | [**AccessType**](AccessType.md) |  | 
**NewEnvironmentAccessType** | [**EnvironmentAccessType**](EnvironmentAccessType.md) |  | 
**EnvironmentAccesses** | [**[]PermissionGroupEnvironmentAccessModel**](PermissionGroupEnvironmentAccessModel.md) | List of environment specific permissions. | 
**Product** | [**ProductModel**](ProductModel.md) |  | 
**ApprovalPermissionType** | [**ApprovalPermissionType**](ApprovalPermissionType.md) |  | 
**NewEnvironmentApprovalPermissionType** | [**EnvironmentApprovalPermissionType**](EnvironmentApprovalPermissionType.md) |  | 
**EnvironmentApprovalPermissions** | [**[]PermissionGroupEnvironmentApprovalPermissionModel**](PermissionGroupEnvironmentApprovalPermissionModel.md) | List of environment specific approval permissions. | 

## Methods

### NewPermissionGroupModel

`func NewPermissionGroupModel(permissionGroupId int64, name string, canManageMembers bool, canCreateOrUpdateConfig bool, canDeleteConfig bool, canCreateOrUpdateEnvironment bool, canDeleteEnvironment bool, canCreateOrUpdateSetting bool, canTagSetting bool, canDeleteSetting bool, canCreateOrUpdateTag bool, canDeleteTag bool, canManageWebhook bool, canUseExportImport bool, canManageProductPreferences bool, canManageIntegrations bool, canViewSdkKey bool, canRotateSdkKey bool, canCreateOrUpdateSegments bool, canDeleteSegments bool, canViewProductAuditLog bool, canViewProductStatistics bool, canDisable2FA bool, accessType AccessType, newEnvironmentAccessType EnvironmentAccessType, environmentAccesses []PermissionGroupEnvironmentAccessModel, product ProductModel, approvalPermissionType ApprovalPermissionType, newEnvironmentApprovalPermissionType EnvironmentApprovalPermissionType, environmentApprovalPermissions []PermissionGroupEnvironmentApprovalPermissionModel, ) *PermissionGroupModel`

NewPermissionGroupModel instantiates a new PermissionGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionGroupModelWithDefaults

`func NewPermissionGroupModelWithDefaults() *PermissionGroupModel`

NewPermissionGroupModelWithDefaults instantiates a new PermissionGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionGroupId

`func (o *PermissionGroupModel) GetPermissionGroupId() int64`

GetPermissionGroupId returns the PermissionGroupId field if non-nil, zero value otherwise.

### GetPermissionGroupIdOk

`func (o *PermissionGroupModel) GetPermissionGroupIdOk() (*int64, bool)`

GetPermissionGroupIdOk returns a tuple with the PermissionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroupId

`func (o *PermissionGroupModel) SetPermissionGroupId(v int64)`

SetPermissionGroupId sets PermissionGroupId field to given value.


### GetName

`func (o *PermissionGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionGroupModel) SetName(v string)`

SetName sets Name field to given value.


### GetCanManageMembers

`func (o *PermissionGroupModel) GetCanManageMembers() bool`

GetCanManageMembers returns the CanManageMembers field if non-nil, zero value otherwise.

### GetCanManageMembersOk

`func (o *PermissionGroupModel) GetCanManageMembersOk() (*bool, bool)`

GetCanManageMembersOk returns a tuple with the CanManageMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageMembers

`func (o *PermissionGroupModel) SetCanManageMembers(v bool)`

SetCanManageMembers sets CanManageMembers field to given value.


### GetCanCreateOrUpdateConfig

`func (o *PermissionGroupModel) GetCanCreateOrUpdateConfig() bool`

GetCanCreateOrUpdateConfig returns the CanCreateOrUpdateConfig field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateConfigOk

`func (o *PermissionGroupModel) GetCanCreateOrUpdateConfigOk() (*bool, bool)`

GetCanCreateOrUpdateConfigOk returns a tuple with the CanCreateOrUpdateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateConfig

`func (o *PermissionGroupModel) SetCanCreateOrUpdateConfig(v bool)`

SetCanCreateOrUpdateConfig sets CanCreateOrUpdateConfig field to given value.


### GetCanDeleteConfig

`func (o *PermissionGroupModel) GetCanDeleteConfig() bool`

GetCanDeleteConfig returns the CanDeleteConfig field if non-nil, zero value otherwise.

### GetCanDeleteConfigOk

`func (o *PermissionGroupModel) GetCanDeleteConfigOk() (*bool, bool)`

GetCanDeleteConfigOk returns a tuple with the CanDeleteConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteConfig

`func (o *PermissionGroupModel) SetCanDeleteConfig(v bool)`

SetCanDeleteConfig sets CanDeleteConfig field to given value.


### GetCanCreateOrUpdateEnvironment

`func (o *PermissionGroupModel) GetCanCreateOrUpdateEnvironment() bool`

GetCanCreateOrUpdateEnvironment returns the CanCreateOrUpdateEnvironment field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateEnvironmentOk

`func (o *PermissionGroupModel) GetCanCreateOrUpdateEnvironmentOk() (*bool, bool)`

GetCanCreateOrUpdateEnvironmentOk returns a tuple with the CanCreateOrUpdateEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateEnvironment

`func (o *PermissionGroupModel) SetCanCreateOrUpdateEnvironment(v bool)`

SetCanCreateOrUpdateEnvironment sets CanCreateOrUpdateEnvironment field to given value.


### GetCanDeleteEnvironment

`func (o *PermissionGroupModel) GetCanDeleteEnvironment() bool`

GetCanDeleteEnvironment returns the CanDeleteEnvironment field if non-nil, zero value otherwise.

### GetCanDeleteEnvironmentOk

`func (o *PermissionGroupModel) GetCanDeleteEnvironmentOk() (*bool, bool)`

GetCanDeleteEnvironmentOk returns a tuple with the CanDeleteEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteEnvironment

`func (o *PermissionGroupModel) SetCanDeleteEnvironment(v bool)`

SetCanDeleteEnvironment sets CanDeleteEnvironment field to given value.


### GetCanCreateOrUpdateSetting

`func (o *PermissionGroupModel) GetCanCreateOrUpdateSetting() bool`

GetCanCreateOrUpdateSetting returns the CanCreateOrUpdateSetting field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateSettingOk

`func (o *PermissionGroupModel) GetCanCreateOrUpdateSettingOk() (*bool, bool)`

GetCanCreateOrUpdateSettingOk returns a tuple with the CanCreateOrUpdateSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateSetting

`func (o *PermissionGroupModel) SetCanCreateOrUpdateSetting(v bool)`

SetCanCreateOrUpdateSetting sets CanCreateOrUpdateSetting field to given value.


### GetCanTagSetting

`func (o *PermissionGroupModel) GetCanTagSetting() bool`

GetCanTagSetting returns the CanTagSetting field if non-nil, zero value otherwise.

### GetCanTagSettingOk

`func (o *PermissionGroupModel) GetCanTagSettingOk() (*bool, bool)`

GetCanTagSettingOk returns a tuple with the CanTagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTagSetting

`func (o *PermissionGroupModel) SetCanTagSetting(v bool)`

SetCanTagSetting sets CanTagSetting field to given value.


### GetCanDeleteSetting

`func (o *PermissionGroupModel) GetCanDeleteSetting() bool`

GetCanDeleteSetting returns the CanDeleteSetting field if non-nil, zero value otherwise.

### GetCanDeleteSettingOk

`func (o *PermissionGroupModel) GetCanDeleteSettingOk() (*bool, bool)`

GetCanDeleteSettingOk returns a tuple with the CanDeleteSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteSetting

`func (o *PermissionGroupModel) SetCanDeleteSetting(v bool)`

SetCanDeleteSetting sets CanDeleteSetting field to given value.


### GetCanCreateOrUpdateTag

`func (o *PermissionGroupModel) GetCanCreateOrUpdateTag() bool`

GetCanCreateOrUpdateTag returns the CanCreateOrUpdateTag field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateTagOk

`func (o *PermissionGroupModel) GetCanCreateOrUpdateTagOk() (*bool, bool)`

GetCanCreateOrUpdateTagOk returns a tuple with the CanCreateOrUpdateTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateTag

`func (o *PermissionGroupModel) SetCanCreateOrUpdateTag(v bool)`

SetCanCreateOrUpdateTag sets CanCreateOrUpdateTag field to given value.


### GetCanDeleteTag

`func (o *PermissionGroupModel) GetCanDeleteTag() bool`

GetCanDeleteTag returns the CanDeleteTag field if non-nil, zero value otherwise.

### GetCanDeleteTagOk

`func (o *PermissionGroupModel) GetCanDeleteTagOk() (*bool, bool)`

GetCanDeleteTagOk returns a tuple with the CanDeleteTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteTag

`func (o *PermissionGroupModel) SetCanDeleteTag(v bool)`

SetCanDeleteTag sets CanDeleteTag field to given value.


### GetCanManageWebhook

`func (o *PermissionGroupModel) GetCanManageWebhook() bool`

GetCanManageWebhook returns the CanManageWebhook field if non-nil, zero value otherwise.

### GetCanManageWebhookOk

`func (o *PermissionGroupModel) GetCanManageWebhookOk() (*bool, bool)`

GetCanManageWebhookOk returns a tuple with the CanManageWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageWebhook

`func (o *PermissionGroupModel) SetCanManageWebhook(v bool)`

SetCanManageWebhook sets CanManageWebhook field to given value.


### GetCanUseExportImport

`func (o *PermissionGroupModel) GetCanUseExportImport() bool`

GetCanUseExportImport returns the CanUseExportImport field if non-nil, zero value otherwise.

### GetCanUseExportImportOk

`func (o *PermissionGroupModel) GetCanUseExportImportOk() (*bool, bool)`

GetCanUseExportImportOk returns a tuple with the CanUseExportImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUseExportImport

`func (o *PermissionGroupModel) SetCanUseExportImport(v bool)`

SetCanUseExportImport sets CanUseExportImport field to given value.


### GetCanManageProductPreferences

`func (o *PermissionGroupModel) GetCanManageProductPreferences() bool`

GetCanManageProductPreferences returns the CanManageProductPreferences field if non-nil, zero value otherwise.

### GetCanManageProductPreferencesOk

`func (o *PermissionGroupModel) GetCanManageProductPreferencesOk() (*bool, bool)`

GetCanManageProductPreferencesOk returns a tuple with the CanManageProductPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageProductPreferences

`func (o *PermissionGroupModel) SetCanManageProductPreferences(v bool)`

SetCanManageProductPreferences sets CanManageProductPreferences field to given value.


### GetCanManageIntegrations

`func (o *PermissionGroupModel) GetCanManageIntegrations() bool`

GetCanManageIntegrations returns the CanManageIntegrations field if non-nil, zero value otherwise.

### GetCanManageIntegrationsOk

`func (o *PermissionGroupModel) GetCanManageIntegrationsOk() (*bool, bool)`

GetCanManageIntegrationsOk returns a tuple with the CanManageIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageIntegrations

`func (o *PermissionGroupModel) SetCanManageIntegrations(v bool)`

SetCanManageIntegrations sets CanManageIntegrations field to given value.


### GetCanViewSdkKey

`func (o *PermissionGroupModel) GetCanViewSdkKey() bool`

GetCanViewSdkKey returns the CanViewSdkKey field if non-nil, zero value otherwise.

### GetCanViewSdkKeyOk

`func (o *PermissionGroupModel) GetCanViewSdkKeyOk() (*bool, bool)`

GetCanViewSdkKeyOk returns a tuple with the CanViewSdkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewSdkKey

`func (o *PermissionGroupModel) SetCanViewSdkKey(v bool)`

SetCanViewSdkKey sets CanViewSdkKey field to given value.


### GetCanRotateSdkKey

`func (o *PermissionGroupModel) GetCanRotateSdkKey() bool`

GetCanRotateSdkKey returns the CanRotateSdkKey field if non-nil, zero value otherwise.

### GetCanRotateSdkKeyOk

`func (o *PermissionGroupModel) GetCanRotateSdkKeyOk() (*bool, bool)`

GetCanRotateSdkKeyOk returns a tuple with the CanRotateSdkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRotateSdkKey

`func (o *PermissionGroupModel) SetCanRotateSdkKey(v bool)`

SetCanRotateSdkKey sets CanRotateSdkKey field to given value.


### GetCanCreateOrUpdateSegments

`func (o *PermissionGroupModel) GetCanCreateOrUpdateSegments() bool`

GetCanCreateOrUpdateSegments returns the CanCreateOrUpdateSegments field if non-nil, zero value otherwise.

### GetCanCreateOrUpdateSegmentsOk

`func (o *PermissionGroupModel) GetCanCreateOrUpdateSegmentsOk() (*bool, bool)`

GetCanCreateOrUpdateSegmentsOk returns a tuple with the CanCreateOrUpdateSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateOrUpdateSegments

`func (o *PermissionGroupModel) SetCanCreateOrUpdateSegments(v bool)`

SetCanCreateOrUpdateSegments sets CanCreateOrUpdateSegments field to given value.


### GetCanDeleteSegments

`func (o *PermissionGroupModel) GetCanDeleteSegments() bool`

GetCanDeleteSegments returns the CanDeleteSegments field if non-nil, zero value otherwise.

### GetCanDeleteSegmentsOk

`func (o *PermissionGroupModel) GetCanDeleteSegmentsOk() (*bool, bool)`

GetCanDeleteSegmentsOk returns a tuple with the CanDeleteSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteSegments

`func (o *PermissionGroupModel) SetCanDeleteSegments(v bool)`

SetCanDeleteSegments sets CanDeleteSegments field to given value.


### GetCanViewProductAuditLog

`func (o *PermissionGroupModel) GetCanViewProductAuditLog() bool`

GetCanViewProductAuditLog returns the CanViewProductAuditLog field if non-nil, zero value otherwise.

### GetCanViewProductAuditLogOk

`func (o *PermissionGroupModel) GetCanViewProductAuditLogOk() (*bool, bool)`

GetCanViewProductAuditLogOk returns a tuple with the CanViewProductAuditLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewProductAuditLog

`func (o *PermissionGroupModel) SetCanViewProductAuditLog(v bool)`

SetCanViewProductAuditLog sets CanViewProductAuditLog field to given value.


### GetCanViewProductStatistics

`func (o *PermissionGroupModel) GetCanViewProductStatistics() bool`

GetCanViewProductStatistics returns the CanViewProductStatistics field if non-nil, zero value otherwise.

### GetCanViewProductStatisticsOk

`func (o *PermissionGroupModel) GetCanViewProductStatisticsOk() (*bool, bool)`

GetCanViewProductStatisticsOk returns a tuple with the CanViewProductStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewProductStatistics

`func (o *PermissionGroupModel) SetCanViewProductStatistics(v bool)`

SetCanViewProductStatistics sets CanViewProductStatistics field to given value.


### GetCanDisable2FA

`func (o *PermissionGroupModel) GetCanDisable2FA() bool`

GetCanDisable2FA returns the CanDisable2FA field if non-nil, zero value otherwise.

### GetCanDisable2FAOk

`func (o *PermissionGroupModel) GetCanDisable2FAOk() (*bool, bool)`

GetCanDisable2FAOk returns a tuple with the CanDisable2FA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDisable2FA

`func (o *PermissionGroupModel) SetCanDisable2FA(v bool)`

SetCanDisable2FA sets CanDisable2FA field to given value.


### GetAccessType

`func (o *PermissionGroupModel) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *PermissionGroupModel) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *PermissionGroupModel) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.


### GetNewEnvironmentAccessType

`func (o *PermissionGroupModel) GetNewEnvironmentAccessType() EnvironmentAccessType`

GetNewEnvironmentAccessType returns the NewEnvironmentAccessType field if non-nil, zero value otherwise.

### GetNewEnvironmentAccessTypeOk

`func (o *PermissionGroupModel) GetNewEnvironmentAccessTypeOk() (*EnvironmentAccessType, bool)`

GetNewEnvironmentAccessTypeOk returns a tuple with the NewEnvironmentAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEnvironmentAccessType

`func (o *PermissionGroupModel) SetNewEnvironmentAccessType(v EnvironmentAccessType)`

SetNewEnvironmentAccessType sets NewEnvironmentAccessType field to given value.


### GetEnvironmentAccesses

`func (o *PermissionGroupModel) GetEnvironmentAccesses() []PermissionGroupEnvironmentAccessModel`

GetEnvironmentAccesses returns the EnvironmentAccesses field if non-nil, zero value otherwise.

### GetEnvironmentAccessesOk

`func (o *PermissionGroupModel) GetEnvironmentAccessesOk() (*[]PermissionGroupEnvironmentAccessModel, bool)`

GetEnvironmentAccessesOk returns a tuple with the EnvironmentAccesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentAccesses

`func (o *PermissionGroupModel) SetEnvironmentAccesses(v []PermissionGroupEnvironmentAccessModel)`

SetEnvironmentAccesses sets EnvironmentAccesses field to given value.


### GetProduct

`func (o *PermissionGroupModel) GetProduct() ProductModel`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PermissionGroupModel) GetProductOk() (*ProductModel, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PermissionGroupModel) SetProduct(v ProductModel)`

SetProduct sets Product field to given value.


### GetApprovalPermissionType

`func (o *PermissionGroupModel) GetApprovalPermissionType() ApprovalPermissionType`

GetApprovalPermissionType returns the ApprovalPermissionType field if non-nil, zero value otherwise.

### GetApprovalPermissionTypeOk

`func (o *PermissionGroupModel) GetApprovalPermissionTypeOk() (*ApprovalPermissionType, bool)`

GetApprovalPermissionTypeOk returns a tuple with the ApprovalPermissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalPermissionType

`func (o *PermissionGroupModel) SetApprovalPermissionType(v ApprovalPermissionType)`

SetApprovalPermissionType sets ApprovalPermissionType field to given value.


### GetNewEnvironmentApprovalPermissionType

`func (o *PermissionGroupModel) GetNewEnvironmentApprovalPermissionType() EnvironmentApprovalPermissionType`

GetNewEnvironmentApprovalPermissionType returns the NewEnvironmentApprovalPermissionType field if non-nil, zero value otherwise.

### GetNewEnvironmentApprovalPermissionTypeOk

`func (o *PermissionGroupModel) GetNewEnvironmentApprovalPermissionTypeOk() (*EnvironmentApprovalPermissionType, bool)`

GetNewEnvironmentApprovalPermissionTypeOk returns a tuple with the NewEnvironmentApprovalPermissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEnvironmentApprovalPermissionType

`func (o *PermissionGroupModel) SetNewEnvironmentApprovalPermissionType(v EnvironmentApprovalPermissionType)`

SetNewEnvironmentApprovalPermissionType sets NewEnvironmentApprovalPermissionType field to given value.


### GetEnvironmentApprovalPermissions

`func (o *PermissionGroupModel) GetEnvironmentApprovalPermissions() []PermissionGroupEnvironmentApprovalPermissionModel`

GetEnvironmentApprovalPermissions returns the EnvironmentApprovalPermissions field if non-nil, zero value otherwise.

### GetEnvironmentApprovalPermissionsOk

`func (o *PermissionGroupModel) GetEnvironmentApprovalPermissionsOk() (*[]PermissionGroupEnvironmentApprovalPermissionModel, bool)`

GetEnvironmentApprovalPermissionsOk returns a tuple with the EnvironmentApprovalPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentApprovalPermissions

`func (o *PermissionGroupModel) SetEnvironmentApprovalPermissions(v []PermissionGroupEnvironmentApprovalPermissionModel)`

SetEnvironmentApprovalPermissions sets EnvironmentApprovalPermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


