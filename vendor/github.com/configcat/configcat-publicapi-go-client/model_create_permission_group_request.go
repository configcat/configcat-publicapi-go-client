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

// checks if the CreatePermissionGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePermissionGroupRequest{}

// CreatePermissionGroupRequest struct for CreatePermissionGroupRequest
type CreatePermissionGroupRequest struct {
	// Name of the Permission Group.
	Name string `json:"name"`
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
	EnvironmentAccesses []CreateOrUpdateEnvironmentAccessModel `json:"environmentAccesses,omitempty"`
}

// NewCreatePermissionGroupRequest instantiates a new CreatePermissionGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePermissionGroupRequest(name string) *CreatePermissionGroupRequest {
	this := CreatePermissionGroupRequest{}
	this.Name = name
	return &this
}

// NewCreatePermissionGroupRequestWithDefaults instantiates a new CreatePermissionGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePermissionGroupRequestWithDefaults() *CreatePermissionGroupRequest {
	this := CreatePermissionGroupRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreatePermissionGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatePermissionGroupRequest) SetName(v string) {
	o.Name = v
}

// GetCanManageMembers returns the CanManageMembers field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanManageMembers() bool {
	if o == nil || IsNil(o.CanManageMembers) {
		var ret bool
		return ret
	}
	return *o.CanManageMembers
}

// GetCanManageMembersOk returns a tuple with the CanManageMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanManageMembersOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManageMembers) {
		return nil, false
	}
	return o.CanManageMembers, true
}

// HasCanManageMembers returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanManageMembers() bool {
	if o != nil && !IsNil(o.CanManageMembers) {
		return true
	}

	return false
}

// SetCanManageMembers gets a reference to the given bool and assigns it to the CanManageMembers field.
func (o *CreatePermissionGroupRequest) SetCanManageMembers(v bool) {
	o.CanManageMembers = &v
}

// GetCanCreateOrUpdateConfig returns the CanCreateOrUpdateConfig field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateConfig() bool {
	if o == nil || IsNil(o.CanCreateOrUpdateConfig) {
		var ret bool
		return ret
	}
	return *o.CanCreateOrUpdateConfig
}

// GetCanCreateOrUpdateConfigOk returns a tuple with the CanCreateOrUpdateConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateConfigOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCreateOrUpdateConfig) {
		return nil, false
	}
	return o.CanCreateOrUpdateConfig, true
}

// HasCanCreateOrUpdateConfig returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanCreateOrUpdateConfig() bool {
	if o != nil && !IsNil(o.CanCreateOrUpdateConfig) {
		return true
	}

	return false
}

// SetCanCreateOrUpdateConfig gets a reference to the given bool and assigns it to the CanCreateOrUpdateConfig field.
func (o *CreatePermissionGroupRequest) SetCanCreateOrUpdateConfig(v bool) {
	o.CanCreateOrUpdateConfig = &v
}

// GetCanDeleteConfig returns the CanDeleteConfig field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanDeleteConfig() bool {
	if o == nil || IsNil(o.CanDeleteConfig) {
		var ret bool
		return ret
	}
	return *o.CanDeleteConfig
}

// GetCanDeleteConfigOk returns a tuple with the CanDeleteConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanDeleteConfigOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeleteConfig) {
		return nil, false
	}
	return o.CanDeleteConfig, true
}

// HasCanDeleteConfig returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanDeleteConfig() bool {
	if o != nil && !IsNil(o.CanDeleteConfig) {
		return true
	}

	return false
}

// SetCanDeleteConfig gets a reference to the given bool and assigns it to the CanDeleteConfig field.
func (o *CreatePermissionGroupRequest) SetCanDeleteConfig(v bool) {
	o.CanDeleteConfig = &v
}

// GetCanCreateOrUpdateEnvironment returns the CanCreateOrUpdateEnvironment field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateEnvironment() bool {
	if o == nil || IsNil(o.CanCreateOrUpdateEnvironment) {
		var ret bool
		return ret
	}
	return *o.CanCreateOrUpdateEnvironment
}

// GetCanCreateOrUpdateEnvironmentOk returns a tuple with the CanCreateOrUpdateEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateEnvironmentOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCreateOrUpdateEnvironment) {
		return nil, false
	}
	return o.CanCreateOrUpdateEnvironment, true
}

// HasCanCreateOrUpdateEnvironment returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanCreateOrUpdateEnvironment() bool {
	if o != nil && !IsNil(o.CanCreateOrUpdateEnvironment) {
		return true
	}

	return false
}

// SetCanCreateOrUpdateEnvironment gets a reference to the given bool and assigns it to the CanCreateOrUpdateEnvironment field.
func (o *CreatePermissionGroupRequest) SetCanCreateOrUpdateEnvironment(v bool) {
	o.CanCreateOrUpdateEnvironment = &v
}

// GetCanDeleteEnvironment returns the CanDeleteEnvironment field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanDeleteEnvironment() bool {
	if o == nil || IsNil(o.CanDeleteEnvironment) {
		var ret bool
		return ret
	}
	return *o.CanDeleteEnvironment
}

// GetCanDeleteEnvironmentOk returns a tuple with the CanDeleteEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanDeleteEnvironmentOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeleteEnvironment) {
		return nil, false
	}
	return o.CanDeleteEnvironment, true
}

// HasCanDeleteEnvironment returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanDeleteEnvironment() bool {
	if o != nil && !IsNil(o.CanDeleteEnvironment) {
		return true
	}

	return false
}

// SetCanDeleteEnvironment gets a reference to the given bool and assigns it to the CanDeleteEnvironment field.
func (o *CreatePermissionGroupRequest) SetCanDeleteEnvironment(v bool) {
	o.CanDeleteEnvironment = &v
}

// GetCanCreateOrUpdateSetting returns the CanCreateOrUpdateSetting field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateSetting() bool {
	if o == nil || IsNil(o.CanCreateOrUpdateSetting) {
		var ret bool
		return ret
	}
	return *o.CanCreateOrUpdateSetting
}

// GetCanCreateOrUpdateSettingOk returns a tuple with the CanCreateOrUpdateSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateSettingOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCreateOrUpdateSetting) {
		return nil, false
	}
	return o.CanCreateOrUpdateSetting, true
}

// HasCanCreateOrUpdateSetting returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanCreateOrUpdateSetting() bool {
	if o != nil && !IsNil(o.CanCreateOrUpdateSetting) {
		return true
	}

	return false
}

// SetCanCreateOrUpdateSetting gets a reference to the given bool and assigns it to the CanCreateOrUpdateSetting field.
func (o *CreatePermissionGroupRequest) SetCanCreateOrUpdateSetting(v bool) {
	o.CanCreateOrUpdateSetting = &v
}

// GetCanTagSetting returns the CanTagSetting field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanTagSetting() bool {
	if o == nil || IsNil(o.CanTagSetting) {
		var ret bool
		return ret
	}
	return *o.CanTagSetting
}

// GetCanTagSettingOk returns a tuple with the CanTagSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanTagSettingOk() (*bool, bool) {
	if o == nil || IsNil(o.CanTagSetting) {
		return nil, false
	}
	return o.CanTagSetting, true
}

// HasCanTagSetting returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanTagSetting() bool {
	if o != nil && !IsNil(o.CanTagSetting) {
		return true
	}

	return false
}

// SetCanTagSetting gets a reference to the given bool and assigns it to the CanTagSetting field.
func (o *CreatePermissionGroupRequest) SetCanTagSetting(v bool) {
	o.CanTagSetting = &v
}

// GetCanDeleteSetting returns the CanDeleteSetting field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanDeleteSetting() bool {
	if o == nil || IsNil(o.CanDeleteSetting) {
		var ret bool
		return ret
	}
	return *o.CanDeleteSetting
}

// GetCanDeleteSettingOk returns a tuple with the CanDeleteSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanDeleteSettingOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeleteSetting) {
		return nil, false
	}
	return o.CanDeleteSetting, true
}

// HasCanDeleteSetting returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanDeleteSetting() bool {
	if o != nil && !IsNil(o.CanDeleteSetting) {
		return true
	}

	return false
}

// SetCanDeleteSetting gets a reference to the given bool and assigns it to the CanDeleteSetting field.
func (o *CreatePermissionGroupRequest) SetCanDeleteSetting(v bool) {
	o.CanDeleteSetting = &v
}

// GetCanCreateOrUpdateTag returns the CanCreateOrUpdateTag field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateTag() bool {
	if o == nil || IsNil(o.CanCreateOrUpdateTag) {
		var ret bool
		return ret
	}
	return *o.CanCreateOrUpdateTag
}

// GetCanCreateOrUpdateTagOk returns a tuple with the CanCreateOrUpdateTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateTagOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCreateOrUpdateTag) {
		return nil, false
	}
	return o.CanCreateOrUpdateTag, true
}

// HasCanCreateOrUpdateTag returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanCreateOrUpdateTag() bool {
	if o != nil && !IsNil(o.CanCreateOrUpdateTag) {
		return true
	}

	return false
}

// SetCanCreateOrUpdateTag gets a reference to the given bool and assigns it to the CanCreateOrUpdateTag field.
func (o *CreatePermissionGroupRequest) SetCanCreateOrUpdateTag(v bool) {
	o.CanCreateOrUpdateTag = &v
}

// GetCanDeleteTag returns the CanDeleteTag field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanDeleteTag() bool {
	if o == nil || IsNil(o.CanDeleteTag) {
		var ret bool
		return ret
	}
	return *o.CanDeleteTag
}

// GetCanDeleteTagOk returns a tuple with the CanDeleteTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanDeleteTagOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeleteTag) {
		return nil, false
	}
	return o.CanDeleteTag, true
}

// HasCanDeleteTag returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanDeleteTag() bool {
	if o != nil && !IsNil(o.CanDeleteTag) {
		return true
	}

	return false
}

// SetCanDeleteTag gets a reference to the given bool and assigns it to the CanDeleteTag field.
func (o *CreatePermissionGroupRequest) SetCanDeleteTag(v bool) {
	o.CanDeleteTag = &v
}

// GetCanManageWebhook returns the CanManageWebhook field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanManageWebhook() bool {
	if o == nil || IsNil(o.CanManageWebhook) {
		var ret bool
		return ret
	}
	return *o.CanManageWebhook
}

// GetCanManageWebhookOk returns a tuple with the CanManageWebhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanManageWebhookOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManageWebhook) {
		return nil, false
	}
	return o.CanManageWebhook, true
}

// HasCanManageWebhook returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanManageWebhook() bool {
	if o != nil && !IsNil(o.CanManageWebhook) {
		return true
	}

	return false
}

// SetCanManageWebhook gets a reference to the given bool and assigns it to the CanManageWebhook field.
func (o *CreatePermissionGroupRequest) SetCanManageWebhook(v bool) {
	o.CanManageWebhook = &v
}

// GetCanUseExportImport returns the CanUseExportImport field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanUseExportImport() bool {
	if o == nil || IsNil(o.CanUseExportImport) {
		var ret bool
		return ret
	}
	return *o.CanUseExportImport
}

// GetCanUseExportImportOk returns a tuple with the CanUseExportImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanUseExportImportOk() (*bool, bool) {
	if o == nil || IsNil(o.CanUseExportImport) {
		return nil, false
	}
	return o.CanUseExportImport, true
}

// HasCanUseExportImport returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanUseExportImport() bool {
	if o != nil && !IsNil(o.CanUseExportImport) {
		return true
	}

	return false
}

// SetCanUseExportImport gets a reference to the given bool and assigns it to the CanUseExportImport field.
func (o *CreatePermissionGroupRequest) SetCanUseExportImport(v bool) {
	o.CanUseExportImport = &v
}

// GetCanManageProductPreferences returns the CanManageProductPreferences field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanManageProductPreferences() bool {
	if o == nil || IsNil(o.CanManageProductPreferences) {
		var ret bool
		return ret
	}
	return *o.CanManageProductPreferences
}

// GetCanManageProductPreferencesOk returns a tuple with the CanManageProductPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanManageProductPreferencesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManageProductPreferences) {
		return nil, false
	}
	return o.CanManageProductPreferences, true
}

// HasCanManageProductPreferences returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanManageProductPreferences() bool {
	if o != nil && !IsNil(o.CanManageProductPreferences) {
		return true
	}

	return false
}

// SetCanManageProductPreferences gets a reference to the given bool and assigns it to the CanManageProductPreferences field.
func (o *CreatePermissionGroupRequest) SetCanManageProductPreferences(v bool) {
	o.CanManageProductPreferences = &v
}

// GetCanManageIntegrations returns the CanManageIntegrations field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanManageIntegrations() bool {
	if o == nil || IsNil(o.CanManageIntegrations) {
		var ret bool
		return ret
	}
	return *o.CanManageIntegrations
}

// GetCanManageIntegrationsOk returns a tuple with the CanManageIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanManageIntegrationsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManageIntegrations) {
		return nil, false
	}
	return o.CanManageIntegrations, true
}

// HasCanManageIntegrations returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanManageIntegrations() bool {
	if o != nil && !IsNil(o.CanManageIntegrations) {
		return true
	}

	return false
}

// SetCanManageIntegrations gets a reference to the given bool and assigns it to the CanManageIntegrations field.
func (o *CreatePermissionGroupRequest) SetCanManageIntegrations(v bool) {
	o.CanManageIntegrations = &v
}

// GetCanViewSdkKey returns the CanViewSdkKey field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanViewSdkKey() bool {
	if o == nil || IsNil(o.CanViewSdkKey) {
		var ret bool
		return ret
	}
	return *o.CanViewSdkKey
}

// GetCanViewSdkKeyOk returns a tuple with the CanViewSdkKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanViewSdkKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.CanViewSdkKey) {
		return nil, false
	}
	return o.CanViewSdkKey, true
}

// HasCanViewSdkKey returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanViewSdkKey() bool {
	if o != nil && !IsNil(o.CanViewSdkKey) {
		return true
	}

	return false
}

// SetCanViewSdkKey gets a reference to the given bool and assigns it to the CanViewSdkKey field.
func (o *CreatePermissionGroupRequest) SetCanViewSdkKey(v bool) {
	o.CanViewSdkKey = &v
}

// GetCanRotateSdkKey returns the CanRotateSdkKey field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanRotateSdkKey() bool {
	if o == nil || IsNil(o.CanRotateSdkKey) {
		var ret bool
		return ret
	}
	return *o.CanRotateSdkKey
}

// GetCanRotateSdkKeyOk returns a tuple with the CanRotateSdkKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanRotateSdkKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.CanRotateSdkKey) {
		return nil, false
	}
	return o.CanRotateSdkKey, true
}

// HasCanRotateSdkKey returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanRotateSdkKey() bool {
	if o != nil && !IsNil(o.CanRotateSdkKey) {
		return true
	}

	return false
}

// SetCanRotateSdkKey gets a reference to the given bool and assigns it to the CanRotateSdkKey field.
func (o *CreatePermissionGroupRequest) SetCanRotateSdkKey(v bool) {
	o.CanRotateSdkKey = &v
}

// GetCanCreateOrUpdateSegments returns the CanCreateOrUpdateSegments field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateSegments() bool {
	if o == nil || IsNil(o.CanCreateOrUpdateSegments) {
		var ret bool
		return ret
	}
	return *o.CanCreateOrUpdateSegments
}

// GetCanCreateOrUpdateSegmentsOk returns a tuple with the CanCreateOrUpdateSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanCreateOrUpdateSegmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCreateOrUpdateSegments) {
		return nil, false
	}
	return o.CanCreateOrUpdateSegments, true
}

// HasCanCreateOrUpdateSegments returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanCreateOrUpdateSegments() bool {
	if o != nil && !IsNil(o.CanCreateOrUpdateSegments) {
		return true
	}

	return false
}

// SetCanCreateOrUpdateSegments gets a reference to the given bool and assigns it to the CanCreateOrUpdateSegments field.
func (o *CreatePermissionGroupRequest) SetCanCreateOrUpdateSegments(v bool) {
	o.CanCreateOrUpdateSegments = &v
}

// GetCanDeleteSegments returns the CanDeleteSegments field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanDeleteSegments() bool {
	if o == nil || IsNil(o.CanDeleteSegments) {
		var ret bool
		return ret
	}
	return *o.CanDeleteSegments
}

// GetCanDeleteSegmentsOk returns a tuple with the CanDeleteSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanDeleteSegmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeleteSegments) {
		return nil, false
	}
	return o.CanDeleteSegments, true
}

// HasCanDeleteSegments returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanDeleteSegments() bool {
	if o != nil && !IsNil(o.CanDeleteSegments) {
		return true
	}

	return false
}

// SetCanDeleteSegments gets a reference to the given bool and assigns it to the CanDeleteSegments field.
func (o *CreatePermissionGroupRequest) SetCanDeleteSegments(v bool) {
	o.CanDeleteSegments = &v
}

// GetCanViewProductAuditLog returns the CanViewProductAuditLog field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanViewProductAuditLog() bool {
	if o == nil || IsNil(o.CanViewProductAuditLog) {
		var ret bool
		return ret
	}
	return *o.CanViewProductAuditLog
}

// GetCanViewProductAuditLogOk returns a tuple with the CanViewProductAuditLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanViewProductAuditLogOk() (*bool, bool) {
	if o == nil || IsNil(o.CanViewProductAuditLog) {
		return nil, false
	}
	return o.CanViewProductAuditLog, true
}

// HasCanViewProductAuditLog returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanViewProductAuditLog() bool {
	if o != nil && !IsNil(o.CanViewProductAuditLog) {
		return true
	}

	return false
}

// SetCanViewProductAuditLog gets a reference to the given bool and assigns it to the CanViewProductAuditLog field.
func (o *CreatePermissionGroupRequest) SetCanViewProductAuditLog(v bool) {
	o.CanViewProductAuditLog = &v
}

// GetCanViewProductStatistics returns the CanViewProductStatistics field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetCanViewProductStatistics() bool {
	if o == nil || IsNil(o.CanViewProductStatistics) {
		var ret bool
		return ret
	}
	return *o.CanViewProductStatistics
}

// GetCanViewProductStatisticsOk returns a tuple with the CanViewProductStatistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetCanViewProductStatisticsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanViewProductStatistics) {
		return nil, false
	}
	return o.CanViewProductStatistics, true
}

// HasCanViewProductStatistics returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasCanViewProductStatistics() bool {
	if o != nil && !IsNil(o.CanViewProductStatistics) {
		return true
	}

	return false
}

// SetCanViewProductStatistics gets a reference to the given bool and assigns it to the CanViewProductStatistics field.
func (o *CreatePermissionGroupRequest) SetCanViewProductStatistics(v bool) {
	o.CanViewProductStatistics = &v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetAccessType() AccessType {
	if o == nil || IsNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *CreatePermissionGroupRequest) SetAccessType(v AccessType) {
	o.AccessType = &v
}

// GetNewEnvironmentAccessType returns the NewEnvironmentAccessType field value if set, zero value otherwise.
func (o *CreatePermissionGroupRequest) GetNewEnvironmentAccessType() EnvironmentAccessType {
	if o == nil || IsNil(o.NewEnvironmentAccessType) {
		var ret EnvironmentAccessType
		return ret
	}
	return *o.NewEnvironmentAccessType
}

// GetNewEnvironmentAccessTypeOk returns a tuple with the NewEnvironmentAccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePermissionGroupRequest) GetNewEnvironmentAccessTypeOk() (*EnvironmentAccessType, bool) {
	if o == nil || IsNil(o.NewEnvironmentAccessType) {
		return nil, false
	}
	return o.NewEnvironmentAccessType, true
}

// HasNewEnvironmentAccessType returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasNewEnvironmentAccessType() bool {
	if o != nil && !IsNil(o.NewEnvironmentAccessType) {
		return true
	}

	return false
}

// SetNewEnvironmentAccessType gets a reference to the given EnvironmentAccessType and assigns it to the NewEnvironmentAccessType field.
func (o *CreatePermissionGroupRequest) SetNewEnvironmentAccessType(v EnvironmentAccessType) {
	o.NewEnvironmentAccessType = &v
}

// GetEnvironmentAccesses returns the EnvironmentAccesses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePermissionGroupRequest) GetEnvironmentAccesses() []CreateOrUpdateEnvironmentAccessModel {
	if o == nil {
		var ret []CreateOrUpdateEnvironmentAccessModel
		return ret
	}
	return o.EnvironmentAccesses
}

// GetEnvironmentAccessesOk returns a tuple with the EnvironmentAccesses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePermissionGroupRequest) GetEnvironmentAccessesOk() ([]CreateOrUpdateEnvironmentAccessModel, bool) {
	if o == nil || IsNil(o.EnvironmentAccesses) {
		return nil, false
	}
	return o.EnvironmentAccesses, true
}

// HasEnvironmentAccesses returns a boolean if a field has been set.
func (o *CreatePermissionGroupRequest) HasEnvironmentAccesses() bool {
	if o != nil && IsNil(o.EnvironmentAccesses) {
		return true
	}

	return false
}

// SetEnvironmentAccesses gets a reference to the given []CreateOrUpdateEnvironmentAccessModel and assigns it to the EnvironmentAccesses field.
func (o *CreatePermissionGroupRequest) SetEnvironmentAccesses(v []CreateOrUpdateEnvironmentAccessModel) {
	o.EnvironmentAccesses = v
}

func (o CreatePermissionGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePermissionGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
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
	return toSerialize, nil
}

type NullableCreatePermissionGroupRequest struct {
	value *CreatePermissionGroupRequest
	isSet bool
}

func (v NullableCreatePermissionGroupRequest) Get() *CreatePermissionGroupRequest {
	return v.value
}

func (v *NullableCreatePermissionGroupRequest) Set(val *CreatePermissionGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePermissionGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePermissionGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePermissionGroupRequest(val *CreatePermissionGroupRequest) *NullableCreatePermissionGroupRequest {
	return &NullableCreatePermissionGroupRequest{value: val, isSet: true}
}

func (v NullableCreatePermissionGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePermissionGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


