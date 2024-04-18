/*
ConfigCat Public Management API

The purpose of this API is to access the ConfigCat platform programmatically. You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.  **Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://test-api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://test-api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
)

// checks if the OrganizationPermissionGroupModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationPermissionGroupModel{}

// OrganizationPermissionGroupModel Describes the Member's Permission Group within a Product.
type OrganizationPermissionGroupModel struct {
	// Identifier of the Member's Permission Group.
	PermissionGroupId *int64 `json:"permissionGroupId,omitempty"`
	// Name of the Member's Permission Group.
	Name NullableString `json:"name,omitempty"`
}

// NewOrganizationPermissionGroupModel instantiates a new OrganizationPermissionGroupModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationPermissionGroupModel() *OrganizationPermissionGroupModel {
	this := OrganizationPermissionGroupModel{}
	return &this
}

// NewOrganizationPermissionGroupModelWithDefaults instantiates a new OrganizationPermissionGroupModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationPermissionGroupModelWithDefaults() *OrganizationPermissionGroupModel {
	this := OrganizationPermissionGroupModel{}
	return &this
}

// GetPermissionGroupId returns the PermissionGroupId field value if set, zero value otherwise.
func (o *OrganizationPermissionGroupModel) GetPermissionGroupId() int64 {
	if o == nil || IsNil(o.PermissionGroupId) {
		var ret int64
		return ret
	}
	return *o.PermissionGroupId
}

// GetPermissionGroupIdOk returns a tuple with the PermissionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationPermissionGroupModel) GetPermissionGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PermissionGroupId) {
		return nil, false
	}
	return o.PermissionGroupId, true
}

// HasPermissionGroupId returns a boolean if a field has been set.
func (o *OrganizationPermissionGroupModel) HasPermissionGroupId() bool {
	if o != nil && !IsNil(o.PermissionGroupId) {
		return true
	}

	return false
}

// SetPermissionGroupId gets a reference to the given int64 and assigns it to the PermissionGroupId field.
func (o *OrganizationPermissionGroupModel) SetPermissionGroupId(v int64) {
	o.PermissionGroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationPermissionGroupModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationPermissionGroupModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationPermissionGroupModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OrganizationPermissionGroupModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *OrganizationPermissionGroupModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OrganizationPermissionGroupModel) UnsetName() {
	o.Name.Unset()
}

func (o OrganizationPermissionGroupModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationPermissionGroupModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PermissionGroupId) {
		toSerialize["permissionGroupId"] = o.PermissionGroupId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableOrganizationPermissionGroupModel struct {
	value *OrganizationPermissionGroupModel
	isSet bool
}

func (v NullableOrganizationPermissionGroupModel) Get() *OrganizationPermissionGroupModel {
	return v.value
}

func (v *NullableOrganizationPermissionGroupModel) Set(val *OrganizationPermissionGroupModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationPermissionGroupModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationPermissionGroupModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationPermissionGroupModel(val *OrganizationPermissionGroupModel) *NullableOrganizationPermissionGroupModel {
	return &NullableOrganizationPermissionGroupModel{value: val, isSet: true}
}

func (v NullableOrganizationPermissionGroupModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationPermissionGroupModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


