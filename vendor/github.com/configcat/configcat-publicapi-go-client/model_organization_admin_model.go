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

// checks if the OrganizationAdminModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationAdminModel{}

// OrganizationAdminModel Describes an Organization Admin.
type OrganizationAdminModel struct {
	// Identifier of the Organization Admin.
	UserId NullableString `json:"userId,omitempty"`
	// Name of the Organization Admin.
	FullName NullableString `json:"fullName,omitempty"`
	// Email of the OrganizationAdmin.
	Email NullableString `json:"email,omitempty"`
	// Determines whether 2FA is enabled for the Organization Admin.
	TwoFactorEnabled *bool `json:"twoFactorEnabled,omitempty"`
}

// NewOrganizationAdminModel instantiates a new OrganizationAdminModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationAdminModel() *OrganizationAdminModel {
	this := OrganizationAdminModel{}
	return &this
}

// NewOrganizationAdminModelWithDefaults instantiates a new OrganizationAdminModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationAdminModelWithDefaults() *OrganizationAdminModel {
	this := OrganizationAdminModel{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationAdminModel) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationAdminModel) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *OrganizationAdminModel) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *OrganizationAdminModel) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *OrganizationAdminModel) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *OrganizationAdminModel) UnsetUserId() {
	o.UserId.Unset()
}

// GetFullName returns the FullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationAdminModel) GetFullName() string {
	if o == nil || IsNil(o.FullName.Get()) {
		var ret string
		return ret
	}
	return *o.FullName.Get()
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationAdminModel) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullName.Get(), o.FullName.IsSet()
}

// HasFullName returns a boolean if a field has been set.
func (o *OrganizationAdminModel) HasFullName() bool {
	if o != nil && o.FullName.IsSet() {
		return true
	}

	return false
}

// SetFullName gets a reference to the given NullableString and assigns it to the FullName field.
func (o *OrganizationAdminModel) SetFullName(v string) {
	o.FullName.Set(&v)
}
// SetFullNameNil sets the value for FullName to be an explicit nil
func (o *OrganizationAdminModel) SetFullNameNil() {
	o.FullName.Set(nil)
}

// UnsetFullName ensures that no value is present for FullName, not even an explicit nil
func (o *OrganizationAdminModel) UnsetFullName() {
	o.FullName.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationAdminModel) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationAdminModel) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *OrganizationAdminModel) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *OrganizationAdminModel) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *OrganizationAdminModel) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *OrganizationAdminModel) UnsetEmail() {
	o.Email.Unset()
}

// GetTwoFactorEnabled returns the TwoFactorEnabled field value if set, zero value otherwise.
func (o *OrganizationAdminModel) GetTwoFactorEnabled() bool {
	if o == nil || IsNil(o.TwoFactorEnabled) {
		var ret bool
		return ret
	}
	return *o.TwoFactorEnabled
}

// GetTwoFactorEnabledOk returns a tuple with the TwoFactorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAdminModel) GetTwoFactorEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TwoFactorEnabled) {
		return nil, false
	}
	return o.TwoFactorEnabled, true
}

// HasTwoFactorEnabled returns a boolean if a field has been set.
func (o *OrganizationAdminModel) HasTwoFactorEnabled() bool {
	if o != nil && !IsNil(o.TwoFactorEnabled) {
		return true
	}

	return false
}

// SetTwoFactorEnabled gets a reference to the given bool and assigns it to the TwoFactorEnabled field.
func (o *OrganizationAdminModel) SetTwoFactorEnabled(v bool) {
	o.TwoFactorEnabled = &v
}

func (o OrganizationAdminModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationAdminModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.FullName.IsSet() {
		toSerialize["fullName"] = o.FullName.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if !IsNil(o.TwoFactorEnabled) {
		toSerialize["twoFactorEnabled"] = o.TwoFactorEnabled
	}
	return toSerialize, nil
}

type NullableOrganizationAdminModel struct {
	value *OrganizationAdminModel
	isSet bool
}

func (v NullableOrganizationAdminModel) Get() *OrganizationAdminModel {
	return v.value
}

func (v *NullableOrganizationAdminModel) Set(val *OrganizationAdminModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationAdminModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationAdminModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationAdminModel(val *OrganizationAdminModel) *NullableOrganizationAdminModel {
	return &NullableOrganizationAdminModel{value: val, isSet: true}
}

func (v NullableOrganizationAdminModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationAdminModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


