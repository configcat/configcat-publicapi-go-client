/*
ConfigCat Public Management API

The purpose of this API is to access the ConfigCat platform programmatically. You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.  **Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  format.   **Important:** Do not use this API for accessing and evaluating feature flag values. Use the [SDKs](https://configcat.com/docs/sdk-reference/overview) or the [ConfigCat Proxy](https://configcat.com/docs/advanced/proxy/proxy-overview/) instead.  # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://test-api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://test-api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
)

// checks if the PermissionGroupModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionGroupModel{}

// PermissionGroupModel struct for PermissionGroupModel
type PermissionGroupModel struct {
	// Identifier of the Permission Group.
	PermissionGroupId *int64 `json:"permissionGroupId,omitempty"`
	// Name of the Permission Group.
	Name NullableString `json:"name,omitempty"`
	// Group members can manage team members.
	CanManageMembers *bool `json:"canManageMembers,omitempty"`
	// Group members can create/update Configs.
	CanCreateOrUpdateConfig *bool `json:"canCreateOrUpdateConfig,omitempty"`
	// Group members can delete Configs.
	CanDeleteConfig *bool `json:"canDeleteConfig,omitempty"`
	// Group members can create/update Environments.
	CanCreateOrUpdateEnvironment *bool `json:"canCreateOrUpdateEnvironment,omitempty"`
	// Group members can delete Environments.
	CanDeleteEnvironment *bool `json:"canDeleteEnvironment,omitempty"`
	// Group members can create/update Feature Flags and Settings.
	CanCreateOrUpdateSetting *bool `json:"canCreateOrUpdateSetting,omitempty"`
	// Group members can attach/detach Tags to Feature Flags and Settings.
	CanTagSetting *bool `json:"canTagSetting,omitempty"`
	// Group members can delete Feature Flags and Settings.
	CanDeleteSetting *bool `json:"canDeleteSetting,omitempty"`
	// Group members can create/update Tags.
	CanCreateOrUpdateTag *bool `json:"canCreateOrUpdateTag,omitempty"`
	// Group members can delete Tags.
	CanDeleteTag *bool `json:"canDeleteTag,omitempty"`
	// Group members can create/update/delete Webhooks.
	CanManageWebhook *bool `json:"canManageWebhook,omitempty"`
	// Group members can use the export/import feature.
	CanUseExportImport *bool `json:"canUseExportImport,omitempty"`
	// Group members can update Product preferences.
	CanManageProductPreferences *bool `json:"canManageProductPreferences,omitempty"`
	// Group members can add and configure integrations.
	CanManageIntegrations *bool `json:"canManageIntegrations,omitempty"`
	// Group members has access to SDK keys.
	CanViewSdkKey *bool `json:"canViewSdkKey,omitempty"`
	// Group members can rotate SDK keys.
	CanRotateSdkKey *bool `json:"canRotateSdkKey,omitempty"`
	// Group members can create/update Segments.
	CanCreateOrUpdateSegments *bool `json:"canCreateOrUpdateSegments,omitempty"`
	// Group members can delete Segments.
	CanDeleteSegments *bool `json:"canDeleteSegments,omitempty"`
	// Group members has access to audit logs.
	CanViewProductAuditLog *bool `json:"canViewProductAuditLog,omitempty"`
	// Group members has access to product statistics.
	CanViewProductStatistics *bool `json:"canViewProductStatistics,omitempty"`
	AccessType *AccessType `json:"accessType,omitempty"`
	NewEnvironmentAccessType *EnvironmentAccessType `json:"newEnvironmentAccessType,omitempty"`
	// List of environment specific permissions.
	EnvironmentAccesses []EnvironmentAccessModel `json:"environmentAccesses,omitempty"`
	Product *ProductModel `json:"product,omitempty"`
}

// NewPermissionGroupModel instantiates a new PermissionGroupModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionGroupModel() *PermissionGroupModel {
	this := PermissionGroupModel{}
	return &this
}

// NewPermissionGroupModelWithDefaults instantiates a new PermissionGroupModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionGroupModelWithDefaults() *PermissionGroupModel {
	this := PermissionGroupModel{}
	return &this
}

// GetPermissionGroupId returns the PermissionGroupId field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetPermissionGroupId() int64 {
	if o == nil || IsNil(o.PermissionGroupId) {
		var ret int64
		return ret
	}
	return *o.PermissionGroupId
}

// GetPermissionGroupIdOk returns a tuple with the PermissionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetPermissionGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PermissionGroupId) {
		return nil, false
	}
	return o.PermissionGroupId, true
}

// HasPermissionGroupId returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasPermissionGroupId() bool {
	if o != nil && !IsNil(o.PermissionGroupId) {
		return true
	}

	return false
}

// SetPermissionGroupId gets a reference to the given int64 and assigns it to the PermissionGroupId field.
func (o *PermissionGroupModel) SetPermissionGroupId(v int64) {
	o.PermissionGroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PermissionGroupModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PermissionGroupModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PermissionGroupModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PermissionGroupModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PermissionGroupModel) UnsetName() {
	o.Name.Unset()
}

// GetCanManageMembers returns the CanManageMembers field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanManageMembers() bool {
	if o == nil || IsNil(o.CanManageMembers) {
		var ret bool
		return ret
	}
	return *o.CanManageMembers
}

// GetCanManageMembersOk returns a tuple with the CanManageMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanManageMembersOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManageMembers) {
		return nil, false
	}
	return o.CanManageMembers, true
}

// HasCanManageMembers returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanManageMembers() bool {
	if o != nil && !IsNil(o.CanManageMembers) {
		return true
	}

	return false
}

// SetCanManageMembers gets a reference to the given bool and assigns it to the CanManageMembers field.
func (o *PermissionGroupModel) SetCanManageMembers(v bool) {
	o.CanManageMembers = &v
}

// GetCanCreateOrUpdateConfig returns the CanCreateOrUpdateConfig field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanCreateOrUpdateConfig() bool {
	if o == nil || IsNil(o.CanCreateOrUpdateConfig) {
		var ret bool
		return ret
	}
	return *o.CanCreateOrUpdateConfig
}

// GetCanCreateOrUpdateConfigOk returns a tuple with the CanCreateOrUpdateConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanCreateOrUpdateConfigOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCreateOrUpdateConfig) {
		return nil, false
	}
	return o.CanCreateOrUpdateConfig, true
}

// HasCanCreateOrUpdateConfig returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanCreateOrUpdateConfig() bool {
	if o != nil && !IsNil(o.CanCreateOrUpdateConfig) {
		return true
	}

	return false
}

// SetCanCreateOrUpdateConfig gets a reference to the given bool and assigns it to the CanCreateOrUpdateConfig field.
func (o *PermissionGroupModel) SetCanCreateOrUpdateConfig(v bool) {
	o.CanCreateOrUpdateConfig = &v
}

// GetCanDeleteConfig returns the CanDeleteConfig field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanDeleteConfig() bool {
	if o == nil || IsNil(o.CanDeleteConfig) {
		var ret bool
		return ret
	}
	return *o.CanDeleteConfig
}

// GetCanDeleteConfigOk returns a tuple with the CanDeleteConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanDeleteConfigOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeleteConfig) {
		return nil, false
	}
	return o.CanDeleteConfig, true
}

// HasCanDeleteConfig returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanDeleteConfig() bool {
	if o != nil && !IsNil(o.CanDeleteConfig) {
		return true
	}

	return false
}

// SetCanDeleteConfig gets a reference to the given bool and assigns it to the CanDeleteConfig field.
func (o *PermissionGroupModel) SetCanDeleteConfig(v bool) {
	o.CanDeleteConfig = &v
}

// GetCanCreateOrUpdateEnvironment returns the CanCreateOrUpdateEnvironment field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanCreateOrUpdateEnvironment() bool {
	if o == nil || IsNil(o.CanCreateOrUpdateEnvironment) {
		var ret bool
		return ret
	}
	return *o.CanCreateOrUpdateEnvironment
}

// GetCanCreateOrUpdateEnvironmentOk returns a tuple with the CanCreateOrUpdateEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanCreateOrUpdateEnvironmentOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCreateOrUpdateEnvironment) {
		return nil, false
	}
	return o.CanCreateOrUpdateEnvironment, true
}

// HasCanCreateOrUpdateEnvironment returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanCreateOrUpdateEnvironment() bool {
	if o != nil && !IsNil(o.CanCreateOrUpdateEnvironment) {
		return true
	}

	return false
}

// SetCanCreateOrUpdateEnvironment gets a reference to the given bool and assigns it to the CanCreateOrUpdateEnvironment field.
func (o *PermissionGroupModel) SetCanCreateOrUpdateEnvironment(v bool) {
	o.CanCreateOrUpdateEnvironment = &v
}

// GetCanDeleteEnvironment returns the CanDeleteEnvironment field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanDeleteEnvironment() bool {
	if o == nil || IsNil(o.CanDeleteEnvironment) {
		var ret bool
		return ret
	}
	return *o.CanDeleteEnvironment
}

// GetCanDeleteEnvironmentOk returns a tuple with the CanDeleteEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanDeleteEnvironmentOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeleteEnvironment) {
		return nil, false
	}
	return o.CanDeleteEnvironment, true
}

// HasCanDeleteEnvironment returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanDeleteEnvironment() bool {
	if o != nil && !IsNil(o.CanDeleteEnvironment) {
		return true
	}

	return false
}

// SetCanDeleteEnvironment gets a reference to the given bool and assigns it to the CanDeleteEnvironment field.
func (o *PermissionGroupModel) SetCanDeleteEnvironment(v bool) {
	o.CanDeleteEnvironment = &v
}

// GetCanCreateOrUpdateSetting returns the CanCreateOrUpdateSetting field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanCreateOrUpdateSetting() bool {
	if o == nil || IsNil(o.CanCreateOrUpdateSetting) {
		var ret bool
		return ret
	}
	return *o.CanCreateOrUpdateSetting
}

// GetCanCreateOrUpdateSettingOk returns a tuple with the CanCreateOrUpdateSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanCreateOrUpdateSettingOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCreateOrUpdateSetting) {
		return nil, false
	}
	return o.CanCreateOrUpdateSetting, true
}

// HasCanCreateOrUpdateSetting returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanCreateOrUpdateSetting() bool {
	if o != nil && !IsNil(o.CanCreateOrUpdateSetting) {
		return true
	}

	return false
}

// SetCanCreateOrUpdateSetting gets a reference to the given bool and assigns it to the CanCreateOrUpdateSetting field.
func (o *PermissionGroupModel) SetCanCreateOrUpdateSetting(v bool) {
	o.CanCreateOrUpdateSetting = &v
}

// GetCanTagSetting returns the CanTagSetting field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanTagSetting() bool {
	if o == nil || IsNil(o.CanTagSetting) {
		var ret bool
		return ret
	}
	return *o.CanTagSetting
}

// GetCanTagSettingOk returns a tuple with the CanTagSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanTagSettingOk() (*bool, bool) {
	if o == nil || IsNil(o.CanTagSetting) {
		return nil, false
	}
	return o.CanTagSetting, true
}

// HasCanTagSetting returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanTagSetting() bool {
	if o != nil && !IsNil(o.CanTagSetting) {
		return true
	}

	return false
}

// SetCanTagSetting gets a reference to the given bool and assigns it to the CanTagSetting field.
func (o *PermissionGroupModel) SetCanTagSetting(v bool) {
	o.CanTagSetting = &v
}

// GetCanDeleteSetting returns the CanDeleteSetting field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanDeleteSetting() bool {
	if o == nil || IsNil(o.CanDeleteSetting) {
		var ret bool
		return ret
	}
	return *o.CanDeleteSetting
}

// GetCanDeleteSettingOk returns a tuple with the CanDeleteSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanDeleteSettingOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeleteSetting) {
		return nil, false
	}
	return o.CanDeleteSetting, true
}

// HasCanDeleteSetting returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanDeleteSetting() bool {
	if o != nil && !IsNil(o.CanDeleteSetting) {
		return true
	}

	return false
}

// SetCanDeleteSetting gets a reference to the given bool and assigns it to the CanDeleteSetting field.
func (o *PermissionGroupModel) SetCanDeleteSetting(v bool) {
	o.CanDeleteSetting = &v
}

// GetCanCreateOrUpdateTag returns the CanCreateOrUpdateTag field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanCreateOrUpdateTag() bool {
	if o == nil || IsNil(o.CanCreateOrUpdateTag) {
		var ret bool
		return ret
	}
	return *o.CanCreateOrUpdateTag
}

// GetCanCreateOrUpdateTagOk returns a tuple with the CanCreateOrUpdateTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanCreateOrUpdateTagOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCreateOrUpdateTag) {
		return nil, false
	}
	return o.CanCreateOrUpdateTag, true
}

// HasCanCreateOrUpdateTag returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanCreateOrUpdateTag() bool {
	if o != nil && !IsNil(o.CanCreateOrUpdateTag) {
		return true
	}

	return false
}

// SetCanCreateOrUpdateTag gets a reference to the given bool and assigns it to the CanCreateOrUpdateTag field.
func (o *PermissionGroupModel) SetCanCreateOrUpdateTag(v bool) {
	o.CanCreateOrUpdateTag = &v
}

// GetCanDeleteTag returns the CanDeleteTag field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanDeleteTag() bool {
	if o == nil || IsNil(o.CanDeleteTag) {
		var ret bool
		return ret
	}
	return *o.CanDeleteTag
}

// GetCanDeleteTagOk returns a tuple with the CanDeleteTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanDeleteTagOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeleteTag) {
		return nil, false
	}
	return o.CanDeleteTag, true
}

// HasCanDeleteTag returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanDeleteTag() bool {
	if o != nil && !IsNil(o.CanDeleteTag) {
		return true
	}

	return false
}

// SetCanDeleteTag gets a reference to the given bool and assigns it to the CanDeleteTag field.
func (o *PermissionGroupModel) SetCanDeleteTag(v bool) {
	o.CanDeleteTag = &v
}

// GetCanManageWebhook returns the CanManageWebhook field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanManageWebhook() bool {
	if o == nil || IsNil(o.CanManageWebhook) {
		var ret bool
		return ret
	}
	return *o.CanManageWebhook
}

// GetCanManageWebhookOk returns a tuple with the CanManageWebhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanManageWebhookOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManageWebhook) {
		return nil, false
	}
	return o.CanManageWebhook, true
}

// HasCanManageWebhook returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanManageWebhook() bool {
	if o != nil && !IsNil(o.CanManageWebhook) {
		return true
	}

	return false
}

// SetCanManageWebhook gets a reference to the given bool and assigns it to the CanManageWebhook field.
func (o *PermissionGroupModel) SetCanManageWebhook(v bool) {
	o.CanManageWebhook = &v
}

// GetCanUseExportImport returns the CanUseExportImport field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanUseExportImport() bool {
	if o == nil || IsNil(o.CanUseExportImport) {
		var ret bool
		return ret
	}
	return *o.CanUseExportImport
}

// GetCanUseExportImportOk returns a tuple with the CanUseExportImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanUseExportImportOk() (*bool, bool) {
	if o == nil || IsNil(o.CanUseExportImport) {
		return nil, false
	}
	return o.CanUseExportImport, true
}

// HasCanUseExportImport returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanUseExportImport() bool {
	if o != nil && !IsNil(o.CanUseExportImport) {
		return true
	}

	return false
}

// SetCanUseExportImport gets a reference to the given bool and assigns it to the CanUseExportImport field.
func (o *PermissionGroupModel) SetCanUseExportImport(v bool) {
	o.CanUseExportImport = &v
}

// GetCanManageProductPreferences returns the CanManageProductPreferences field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanManageProductPreferences() bool {
	if o == nil || IsNil(o.CanManageProductPreferences) {
		var ret bool
		return ret
	}
	return *o.CanManageProductPreferences
}

// GetCanManageProductPreferencesOk returns a tuple with the CanManageProductPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanManageProductPreferencesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManageProductPreferences) {
		return nil, false
	}
	return o.CanManageProductPreferences, true
}

// HasCanManageProductPreferences returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanManageProductPreferences() bool {
	if o != nil && !IsNil(o.CanManageProductPreferences) {
		return true
	}

	return false
}

// SetCanManageProductPreferences gets a reference to the given bool and assigns it to the CanManageProductPreferences field.
func (o *PermissionGroupModel) SetCanManageProductPreferences(v bool) {
	o.CanManageProductPreferences = &v
}

// GetCanManageIntegrations returns the CanManageIntegrations field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanManageIntegrations() bool {
	if o == nil || IsNil(o.CanManageIntegrations) {
		var ret bool
		return ret
	}
	return *o.CanManageIntegrations
}

// GetCanManageIntegrationsOk returns a tuple with the CanManageIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanManageIntegrationsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManageIntegrations) {
		return nil, false
	}
	return o.CanManageIntegrations, true
}

// HasCanManageIntegrations returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanManageIntegrations() bool {
	if o != nil && !IsNil(o.CanManageIntegrations) {
		return true
	}

	return false
}

// SetCanManageIntegrations gets a reference to the given bool and assigns it to the CanManageIntegrations field.
func (o *PermissionGroupModel) SetCanManageIntegrations(v bool) {
	o.CanManageIntegrations = &v
}

// GetCanViewSdkKey returns the CanViewSdkKey field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanViewSdkKey() bool {
	if o == nil || IsNil(o.CanViewSdkKey) {
		var ret bool
		return ret
	}
	return *o.CanViewSdkKey
}

// GetCanViewSdkKeyOk returns a tuple with the CanViewSdkKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanViewSdkKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.CanViewSdkKey) {
		return nil, false
	}
	return o.CanViewSdkKey, true
}

// HasCanViewSdkKey returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanViewSdkKey() bool {
	if o != nil && !IsNil(o.CanViewSdkKey) {
		return true
	}

	return false
}

// SetCanViewSdkKey gets a reference to the given bool and assigns it to the CanViewSdkKey field.
func (o *PermissionGroupModel) SetCanViewSdkKey(v bool) {
	o.CanViewSdkKey = &v
}

// GetCanRotateSdkKey returns the CanRotateSdkKey field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanRotateSdkKey() bool {
	if o == nil || IsNil(o.CanRotateSdkKey) {
		var ret bool
		return ret
	}
	return *o.CanRotateSdkKey
}

// GetCanRotateSdkKeyOk returns a tuple with the CanRotateSdkKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanRotateSdkKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.CanRotateSdkKey) {
		return nil, false
	}
	return o.CanRotateSdkKey, true
}

// HasCanRotateSdkKey returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanRotateSdkKey() bool {
	if o != nil && !IsNil(o.CanRotateSdkKey) {
		return true
	}

	return false
}

// SetCanRotateSdkKey gets a reference to the given bool and assigns it to the CanRotateSdkKey field.
func (o *PermissionGroupModel) SetCanRotateSdkKey(v bool) {
	o.CanRotateSdkKey = &v
}

// GetCanCreateOrUpdateSegments returns the CanCreateOrUpdateSegments field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanCreateOrUpdateSegments() bool {
	if o == nil || IsNil(o.CanCreateOrUpdateSegments) {
		var ret bool
		return ret
	}
	return *o.CanCreateOrUpdateSegments
}

// GetCanCreateOrUpdateSegmentsOk returns a tuple with the CanCreateOrUpdateSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanCreateOrUpdateSegmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCreateOrUpdateSegments) {
		return nil, false
	}
	return o.CanCreateOrUpdateSegments, true
}

// HasCanCreateOrUpdateSegments returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanCreateOrUpdateSegments() bool {
	if o != nil && !IsNil(o.CanCreateOrUpdateSegments) {
		return true
	}

	return false
}

// SetCanCreateOrUpdateSegments gets a reference to the given bool and assigns it to the CanCreateOrUpdateSegments field.
func (o *PermissionGroupModel) SetCanCreateOrUpdateSegments(v bool) {
	o.CanCreateOrUpdateSegments = &v
}

// GetCanDeleteSegments returns the CanDeleteSegments field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanDeleteSegments() bool {
	if o == nil || IsNil(o.CanDeleteSegments) {
		var ret bool
		return ret
	}
	return *o.CanDeleteSegments
}

// GetCanDeleteSegmentsOk returns a tuple with the CanDeleteSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanDeleteSegmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeleteSegments) {
		return nil, false
	}
	return o.CanDeleteSegments, true
}

// HasCanDeleteSegments returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanDeleteSegments() bool {
	if o != nil && !IsNil(o.CanDeleteSegments) {
		return true
	}

	return false
}

// SetCanDeleteSegments gets a reference to the given bool and assigns it to the CanDeleteSegments field.
func (o *PermissionGroupModel) SetCanDeleteSegments(v bool) {
	o.CanDeleteSegments = &v
}

// GetCanViewProductAuditLog returns the CanViewProductAuditLog field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanViewProductAuditLog() bool {
	if o == nil || IsNil(o.CanViewProductAuditLog) {
		var ret bool
		return ret
	}
	return *o.CanViewProductAuditLog
}

// GetCanViewProductAuditLogOk returns a tuple with the CanViewProductAuditLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanViewProductAuditLogOk() (*bool, bool) {
	if o == nil || IsNil(o.CanViewProductAuditLog) {
		return nil, false
	}
	return o.CanViewProductAuditLog, true
}

// HasCanViewProductAuditLog returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanViewProductAuditLog() bool {
	if o != nil && !IsNil(o.CanViewProductAuditLog) {
		return true
	}

	return false
}

// SetCanViewProductAuditLog gets a reference to the given bool and assigns it to the CanViewProductAuditLog field.
func (o *PermissionGroupModel) SetCanViewProductAuditLog(v bool) {
	o.CanViewProductAuditLog = &v
}

// GetCanViewProductStatistics returns the CanViewProductStatistics field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetCanViewProductStatistics() bool {
	if o == nil || IsNil(o.CanViewProductStatistics) {
		var ret bool
		return ret
	}
	return *o.CanViewProductStatistics
}

// GetCanViewProductStatisticsOk returns a tuple with the CanViewProductStatistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetCanViewProductStatisticsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanViewProductStatistics) {
		return nil, false
	}
	return o.CanViewProductStatistics, true
}

// HasCanViewProductStatistics returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasCanViewProductStatistics() bool {
	if o != nil && !IsNil(o.CanViewProductStatistics) {
		return true
	}

	return false
}

// SetCanViewProductStatistics gets a reference to the given bool and assigns it to the CanViewProductStatistics field.
func (o *PermissionGroupModel) SetCanViewProductStatistics(v bool) {
	o.CanViewProductStatistics = &v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetAccessType() AccessType {
	if o == nil || IsNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *PermissionGroupModel) SetAccessType(v AccessType) {
	o.AccessType = &v
}

// GetNewEnvironmentAccessType returns the NewEnvironmentAccessType field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetNewEnvironmentAccessType() EnvironmentAccessType {
	if o == nil || IsNil(o.NewEnvironmentAccessType) {
		var ret EnvironmentAccessType
		return ret
	}
	return *o.NewEnvironmentAccessType
}

// GetNewEnvironmentAccessTypeOk returns a tuple with the NewEnvironmentAccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetNewEnvironmentAccessTypeOk() (*EnvironmentAccessType, bool) {
	if o == nil || IsNil(o.NewEnvironmentAccessType) {
		return nil, false
	}
	return o.NewEnvironmentAccessType, true
}

// HasNewEnvironmentAccessType returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasNewEnvironmentAccessType() bool {
	if o != nil && !IsNil(o.NewEnvironmentAccessType) {
		return true
	}

	return false
}

// SetNewEnvironmentAccessType gets a reference to the given EnvironmentAccessType and assigns it to the NewEnvironmentAccessType field.
func (o *PermissionGroupModel) SetNewEnvironmentAccessType(v EnvironmentAccessType) {
	o.NewEnvironmentAccessType = &v
}

// GetEnvironmentAccesses returns the EnvironmentAccesses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PermissionGroupModel) GetEnvironmentAccesses() []EnvironmentAccessModel {
	if o == nil {
		var ret []EnvironmentAccessModel
		return ret
	}
	return o.EnvironmentAccesses
}

// GetEnvironmentAccessesOk returns a tuple with the EnvironmentAccesses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PermissionGroupModel) GetEnvironmentAccessesOk() ([]EnvironmentAccessModel, bool) {
	if o == nil || IsNil(o.EnvironmentAccesses) {
		return nil, false
	}
	return o.EnvironmentAccesses, true
}

// HasEnvironmentAccesses returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasEnvironmentAccesses() bool {
	if o != nil && !IsNil(o.EnvironmentAccesses) {
		return true
	}

	return false
}

// SetEnvironmentAccesses gets a reference to the given []EnvironmentAccessModel and assigns it to the EnvironmentAccesses field.
func (o *PermissionGroupModel) SetEnvironmentAccesses(v []EnvironmentAccessModel) {
	o.EnvironmentAccesses = v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *PermissionGroupModel) GetProduct() ProductModel {
	if o == nil || IsNil(o.Product) {
		var ret ProductModel
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupModel) GetProductOk() (*ProductModel, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *PermissionGroupModel) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given ProductModel and assigns it to the Product field.
func (o *PermissionGroupModel) SetProduct(v ProductModel) {
	o.Product = &v
}

func (o PermissionGroupModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionGroupModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PermissionGroupId) {
		toSerialize["permissionGroupId"] = o.PermissionGroupId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.CanManageMembers) {
		toSerialize["canManageMembers"] = o.CanManageMembers
	}
	if !IsNil(o.CanCreateOrUpdateConfig) {
		toSerialize["canCreateOrUpdateConfig"] = o.CanCreateOrUpdateConfig
	}
	if !IsNil(o.CanDeleteConfig) {
		toSerialize["canDeleteConfig"] = o.CanDeleteConfig
	}
	if !IsNil(o.CanCreateOrUpdateEnvironment) {
		toSerialize["canCreateOrUpdateEnvironment"] = o.CanCreateOrUpdateEnvironment
	}
	if !IsNil(o.CanDeleteEnvironment) {
		toSerialize["canDeleteEnvironment"] = o.CanDeleteEnvironment
	}
	if !IsNil(o.CanCreateOrUpdateSetting) {
		toSerialize["canCreateOrUpdateSetting"] = o.CanCreateOrUpdateSetting
	}
	if !IsNil(o.CanTagSetting) {
		toSerialize["canTagSetting"] = o.CanTagSetting
	}
	if !IsNil(o.CanDeleteSetting) {
		toSerialize["canDeleteSetting"] = o.CanDeleteSetting
	}
	if !IsNil(o.CanCreateOrUpdateTag) {
		toSerialize["canCreateOrUpdateTag"] = o.CanCreateOrUpdateTag
	}
	if !IsNil(o.CanDeleteTag) {
		toSerialize["canDeleteTag"] = o.CanDeleteTag
	}
	if !IsNil(o.CanManageWebhook) {
		toSerialize["canManageWebhook"] = o.CanManageWebhook
	}
	if !IsNil(o.CanUseExportImport) {
		toSerialize["canUseExportImport"] = o.CanUseExportImport
	}
	if !IsNil(o.CanManageProductPreferences) {
		toSerialize["canManageProductPreferences"] = o.CanManageProductPreferences
	}
	if !IsNil(o.CanManageIntegrations) {
		toSerialize["canManageIntegrations"] = o.CanManageIntegrations
	}
	if !IsNil(o.CanViewSdkKey) {
		toSerialize["canViewSdkKey"] = o.CanViewSdkKey
	}
	if !IsNil(o.CanRotateSdkKey) {
		toSerialize["canRotateSdkKey"] = o.CanRotateSdkKey
	}
	if !IsNil(o.CanCreateOrUpdateSegments) {
		toSerialize["canCreateOrUpdateSegments"] = o.CanCreateOrUpdateSegments
	}
	if !IsNil(o.CanDeleteSegments) {
		toSerialize["canDeleteSegments"] = o.CanDeleteSegments
	}
	if !IsNil(o.CanViewProductAuditLog) {
		toSerialize["canViewProductAuditLog"] = o.CanViewProductAuditLog
	}
	if !IsNil(o.CanViewProductStatistics) {
		toSerialize["canViewProductStatistics"] = o.CanViewProductStatistics
	}
	if !IsNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !IsNil(o.NewEnvironmentAccessType) {
		toSerialize["newEnvironmentAccessType"] = o.NewEnvironmentAccessType
	}
	if o.EnvironmentAccesses != nil {
		toSerialize["environmentAccesses"] = o.EnvironmentAccesses
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	return toSerialize, nil
}

type NullablePermissionGroupModel struct {
	value *PermissionGroupModel
	isSet bool
}

func (v NullablePermissionGroupModel) Get() *PermissionGroupModel {
	return v.value
}

func (v *NullablePermissionGroupModel) Set(val *PermissionGroupModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionGroupModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionGroupModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionGroupModel(val *PermissionGroupModel) *NullablePermissionGroupModel {
	return &NullablePermissionGroupModel{value: val, isSet: true}
}

func (v NullablePermissionGroupModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionGroupModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


