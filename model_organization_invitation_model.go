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
	"time"
)

// checks if the OrganizationInvitationModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationInvitationModel{}

// OrganizationInvitationModel struct for OrganizationInvitationModel
type OrganizationInvitationModel struct {
	// The identifier of the Invitation.
	InvitationId *string `json:"invitationId,omitempty"`
	// The invited user's email address.
	Email NullableString `json:"email,omitempty"`
	// The identifier of the Product the user was invited to.
	ProductId *string `json:"productId,omitempty"`
	// The name of the Product the user was invited to.
	ProductName NullableString `json:"productName,omitempty"`
	// The identifier of the Permission Group the user was invited to.
	PermissionGroupId *int64 `json:"permissionGroupId,omitempty"`
	// Creation time of the Invitation.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Determines whether the Invitation is expired.
	Expired *bool `json:"expired,omitempty"`
	// Expiration time of the Invitation.
	Expires *time.Time `json:"expires,omitempty"`
}

// NewOrganizationInvitationModel instantiates a new OrganizationInvitationModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInvitationModel() *OrganizationInvitationModel {
	this := OrganizationInvitationModel{}
	return &this
}

// NewOrganizationInvitationModelWithDefaults instantiates a new OrganizationInvitationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInvitationModelWithDefaults() *OrganizationInvitationModel {
	this := OrganizationInvitationModel{}
	return &this
}

// GetInvitationId returns the InvitationId field value if set, zero value otherwise.
func (o *OrganizationInvitationModel) GetInvitationId() string {
	if o == nil || IsNil(o.InvitationId) {
		var ret string
		return ret
	}
	return *o.InvitationId
}

// GetInvitationIdOk returns a tuple with the InvitationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationModel) GetInvitationIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvitationId) {
		return nil, false
	}
	return o.InvitationId, true
}

// HasInvitationId returns a boolean if a field has been set.
func (o *OrganizationInvitationModel) HasInvitationId() bool {
	if o != nil && !IsNil(o.InvitationId) {
		return true
	}

	return false
}

// SetInvitationId gets a reference to the given string and assigns it to the InvitationId field.
func (o *OrganizationInvitationModel) SetInvitationId(v string) {
	o.InvitationId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationInvitationModel) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationInvitationModel) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *OrganizationInvitationModel) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *OrganizationInvitationModel) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *OrganizationInvitationModel) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *OrganizationInvitationModel) UnsetEmail() {
	o.Email.Unset()
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *OrganizationInvitationModel) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationModel) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *OrganizationInvitationModel) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *OrganizationInvitationModel) SetProductId(v string) {
	o.ProductId = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationInvitationModel) GetProductName() string {
	if o == nil || IsNil(o.ProductName.Get()) {
		var ret string
		return ret
	}
	return *o.ProductName.Get()
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationInvitationModel) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductName.Get(), o.ProductName.IsSet()
}

// HasProductName returns a boolean if a field has been set.
func (o *OrganizationInvitationModel) HasProductName() bool {
	if o != nil && o.ProductName.IsSet() {
		return true
	}

	return false
}

// SetProductName gets a reference to the given NullableString and assigns it to the ProductName field.
func (o *OrganizationInvitationModel) SetProductName(v string) {
	o.ProductName.Set(&v)
}
// SetProductNameNil sets the value for ProductName to be an explicit nil
func (o *OrganizationInvitationModel) SetProductNameNil() {
	o.ProductName.Set(nil)
}

// UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
func (o *OrganizationInvitationModel) UnsetProductName() {
	o.ProductName.Unset()
}

// GetPermissionGroupId returns the PermissionGroupId field value if set, zero value otherwise.
func (o *OrganizationInvitationModel) GetPermissionGroupId() int64 {
	if o == nil || IsNil(o.PermissionGroupId) {
		var ret int64
		return ret
	}
	return *o.PermissionGroupId
}

// GetPermissionGroupIdOk returns a tuple with the PermissionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationModel) GetPermissionGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PermissionGroupId) {
		return nil, false
	}
	return o.PermissionGroupId, true
}

// HasPermissionGroupId returns a boolean if a field has been set.
func (o *OrganizationInvitationModel) HasPermissionGroupId() bool {
	if o != nil && !IsNil(o.PermissionGroupId) {
		return true
	}

	return false
}

// SetPermissionGroupId gets a reference to the given int64 and assigns it to the PermissionGroupId field.
func (o *OrganizationInvitationModel) SetPermissionGroupId(v int64) {
	o.PermissionGroupId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrganizationInvitationModel) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrganizationInvitationModel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OrganizationInvitationModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *OrganizationInvitationModel) GetExpired() bool {
	if o == nil || IsNil(o.Expired) {
		var ret bool
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationModel) GetExpiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Expired) {
		return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *OrganizationInvitationModel) HasExpired() bool {
	if o != nil && !IsNil(o.Expired) {
		return true
	}

	return false
}

// SetExpired gets a reference to the given bool and assigns it to the Expired field.
func (o *OrganizationInvitationModel) SetExpired(v bool) {
	o.Expired = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *OrganizationInvitationModel) GetExpires() time.Time {
	if o == nil || IsNil(o.Expires) {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationModel) GetExpiresOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *OrganizationInvitationModel) HasExpires() bool {
	if o != nil && !IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *OrganizationInvitationModel) SetExpires(v time.Time) {
	o.Expires = &v
}

func (o OrganizationInvitationModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationInvitationModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvitationId) {
		toSerialize["invitationId"] = o.InvitationId
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if o.ProductName.IsSet() {
		toSerialize["productName"] = o.ProductName.Get()
	}
	if !IsNil(o.PermissionGroupId) {
		toSerialize["permissionGroupId"] = o.PermissionGroupId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Expired) {
		toSerialize["expired"] = o.Expired
	}
	if !IsNil(o.Expires) {
		toSerialize["expires"] = o.Expires
	}
	return toSerialize, nil
}

type NullableOrganizationInvitationModel struct {
	value *OrganizationInvitationModel
	isSet bool
}

func (v NullableOrganizationInvitationModel) Get() *OrganizationInvitationModel {
	return v.value
}

func (v *NullableOrganizationInvitationModel) Set(val *OrganizationInvitationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationInvitationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationInvitationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationInvitationModel(val *OrganizationInvitationModel) *NullableOrganizationInvitationModel {
	return &NullableOrganizationInvitationModel{value: val, isSet: true}
}

func (v NullableOrganizationInvitationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationInvitationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

