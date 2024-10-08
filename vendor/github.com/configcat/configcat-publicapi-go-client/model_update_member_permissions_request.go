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

// checks if the UpdateMemberPermissionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMemberPermissionsRequest{}

// UpdateMemberPermissionsRequest struct for UpdateMemberPermissionsRequest
type UpdateMemberPermissionsRequest struct {
	// List of Permission Group identifiers to where the Member should be added.
	PermissionGroupIds []int64 `json:"permissionGroupIds,omitempty"`
	// Indicates that the member must be Organization Admin.
	IsAdmin NullableBool `json:"isAdmin,omitempty"`
	// Indicates that the member must be Billing Manager.
	IsBillingManager NullableBool `json:"isBillingManager,omitempty"`
	// When `true`, the member will be removed from those Permission Groups that are not listed in the `permissionGroupIds` field.
	RemoveFromPermissionGroupsWhereIdNotSet *bool `json:"removeFromPermissionGroupsWhereIdNotSet,omitempty"`
}

// NewUpdateMemberPermissionsRequest instantiates a new UpdateMemberPermissionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMemberPermissionsRequest() *UpdateMemberPermissionsRequest {
	this := UpdateMemberPermissionsRequest{}
	return &this
}

// NewUpdateMemberPermissionsRequestWithDefaults instantiates a new UpdateMemberPermissionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMemberPermissionsRequestWithDefaults() *UpdateMemberPermissionsRequest {
	this := UpdateMemberPermissionsRequest{}
	return &this
}

// GetPermissionGroupIds returns the PermissionGroupIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberPermissionsRequest) GetPermissionGroupIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.PermissionGroupIds
}

// GetPermissionGroupIdsOk returns a tuple with the PermissionGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberPermissionsRequest) GetPermissionGroupIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.PermissionGroupIds) {
		return nil, false
	}
	return o.PermissionGroupIds, true
}

// HasPermissionGroupIds returns a boolean if a field has been set.
func (o *UpdateMemberPermissionsRequest) HasPermissionGroupIds() bool {
	if o != nil && IsNil(o.PermissionGroupIds) {
		return true
	}

	return false
}

// SetPermissionGroupIds gets a reference to the given []int64 and assigns it to the PermissionGroupIds field.
func (o *UpdateMemberPermissionsRequest) SetPermissionGroupIds(v []int64) {
	o.PermissionGroupIds = v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberPermissionsRequest) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAdmin.Get()
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberPermissionsRequest) GetIsAdminOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAdmin.Get(), o.IsAdmin.IsSet()
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *UpdateMemberPermissionsRequest) HasIsAdmin() bool {
	if o != nil && o.IsAdmin.IsSet() {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given NullableBool and assigns it to the IsAdmin field.
func (o *UpdateMemberPermissionsRequest) SetIsAdmin(v bool) {
	o.IsAdmin.Set(&v)
}
// SetIsAdminNil sets the value for IsAdmin to be an explicit nil
func (o *UpdateMemberPermissionsRequest) SetIsAdminNil() {
	o.IsAdmin.Set(nil)
}

// UnsetIsAdmin ensures that no value is present for IsAdmin, not even an explicit nil
func (o *UpdateMemberPermissionsRequest) UnsetIsAdmin() {
	o.IsAdmin.Unset()
}

// GetIsBillingManager returns the IsBillingManager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberPermissionsRequest) GetIsBillingManager() bool {
	if o == nil || IsNil(o.IsBillingManager.Get()) {
		var ret bool
		return ret
	}
	return *o.IsBillingManager.Get()
}

// GetIsBillingManagerOk returns a tuple with the IsBillingManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberPermissionsRequest) GetIsBillingManagerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsBillingManager.Get(), o.IsBillingManager.IsSet()
}

// HasIsBillingManager returns a boolean if a field has been set.
func (o *UpdateMemberPermissionsRequest) HasIsBillingManager() bool {
	if o != nil && o.IsBillingManager.IsSet() {
		return true
	}

	return false
}

// SetIsBillingManager gets a reference to the given NullableBool and assigns it to the IsBillingManager field.
func (o *UpdateMemberPermissionsRequest) SetIsBillingManager(v bool) {
	o.IsBillingManager.Set(&v)
}
// SetIsBillingManagerNil sets the value for IsBillingManager to be an explicit nil
func (o *UpdateMemberPermissionsRequest) SetIsBillingManagerNil() {
	o.IsBillingManager.Set(nil)
}

// UnsetIsBillingManager ensures that no value is present for IsBillingManager, not even an explicit nil
func (o *UpdateMemberPermissionsRequest) UnsetIsBillingManager() {
	o.IsBillingManager.Unset()
}

// GetRemoveFromPermissionGroupsWhereIdNotSet returns the RemoveFromPermissionGroupsWhereIdNotSet field value if set, zero value otherwise.
func (o *UpdateMemberPermissionsRequest) GetRemoveFromPermissionGroupsWhereIdNotSet() bool {
	if o == nil || IsNil(o.RemoveFromPermissionGroupsWhereIdNotSet) {
		var ret bool
		return ret
	}
	return *o.RemoveFromPermissionGroupsWhereIdNotSet
}

// GetRemoveFromPermissionGroupsWhereIdNotSetOk returns a tuple with the RemoveFromPermissionGroupsWhereIdNotSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMemberPermissionsRequest) GetRemoveFromPermissionGroupsWhereIdNotSetOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoveFromPermissionGroupsWhereIdNotSet) {
		return nil, false
	}
	return o.RemoveFromPermissionGroupsWhereIdNotSet, true
}

// HasRemoveFromPermissionGroupsWhereIdNotSet returns a boolean if a field has been set.
func (o *UpdateMemberPermissionsRequest) HasRemoveFromPermissionGroupsWhereIdNotSet() bool {
	if o != nil && !IsNil(o.RemoveFromPermissionGroupsWhereIdNotSet) {
		return true
	}

	return false
}

// SetRemoveFromPermissionGroupsWhereIdNotSet gets a reference to the given bool and assigns it to the RemoveFromPermissionGroupsWhereIdNotSet field.
func (o *UpdateMemberPermissionsRequest) SetRemoveFromPermissionGroupsWhereIdNotSet(v bool) {
	o.RemoveFromPermissionGroupsWhereIdNotSet = &v
}

func (o UpdateMemberPermissionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMemberPermissionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PermissionGroupIds != nil {
		toSerialize["permissionGroupIds"] = o.PermissionGroupIds
	}
	if o.IsAdmin.IsSet() {
		toSerialize["isAdmin"] = o.IsAdmin.Get()
	}
	if o.IsBillingManager.IsSet() {
		toSerialize["isBillingManager"] = o.IsBillingManager.Get()
	}
	if !IsNil(o.RemoveFromPermissionGroupsWhereIdNotSet) {
		toSerialize["removeFromPermissionGroupsWhereIdNotSet"] = o.RemoveFromPermissionGroupsWhereIdNotSet
	}
	return toSerialize, nil
}

type NullableUpdateMemberPermissionsRequest struct {
	value *UpdateMemberPermissionsRequest
	isSet bool
}

func (v NullableUpdateMemberPermissionsRequest) Get() *UpdateMemberPermissionsRequest {
	return v.value
}

func (v *NullableUpdateMemberPermissionsRequest) Set(val *UpdateMemberPermissionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMemberPermissionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMemberPermissionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMemberPermissionsRequest(val *UpdateMemberPermissionsRequest) *NullableUpdateMemberPermissionsRequest {
	return &NullableUpdateMemberPermissionsRequest{value: val, isSet: true}
}

func (v NullableUpdateMemberPermissionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMemberPermissionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


