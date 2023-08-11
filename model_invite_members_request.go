/*
ConfigCat Public Management API

**Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The purpose of this API is to access the ConfigCat platform programmatically.  You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.   The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  and JSON+HAL format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://test-api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://test-api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
)

// checks if the InviteMembersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteMembersRequest{}

// InviteMembersRequest struct for InviteMembersRequest
type InviteMembersRequest struct {
	// List of email addresses to invite.
	Emails []string `json:"emails"`
	// Identifier of the Permission Group to where the invited users should be added.
	PermissionGroupId int64 `json:"permissionGroupId"`
}

// NewInviteMembersRequest instantiates a new InviteMembersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteMembersRequest(emails []string, permissionGroupId int64) *InviteMembersRequest {
	this := InviteMembersRequest{}
	this.Emails = emails
	this.PermissionGroupId = permissionGroupId
	return &this
}

// NewInviteMembersRequestWithDefaults instantiates a new InviteMembersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteMembersRequestWithDefaults() *InviteMembersRequest {
	this := InviteMembersRequest{}
	return &this
}

// GetEmails returns the Emails field value
func (o *InviteMembersRequest) GetEmails() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *InviteMembersRequest) GetEmailsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *InviteMembersRequest) SetEmails(v []string) {
	o.Emails = v
}

// GetPermissionGroupId returns the PermissionGroupId field value
func (o *InviteMembersRequest) GetPermissionGroupId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PermissionGroupId
}

// GetPermissionGroupIdOk returns a tuple with the PermissionGroupId field value
// and a boolean to check if the value has been set.
func (o *InviteMembersRequest) GetPermissionGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionGroupId, true
}

// SetPermissionGroupId sets field value
func (o *InviteMembersRequest) SetPermissionGroupId(v int64) {
	o.PermissionGroupId = v
}

func (o InviteMembersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteMembersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emails"] = o.Emails
	toSerialize["permissionGroupId"] = o.PermissionGroupId
	return toSerialize, nil
}

type NullableInviteMembersRequest struct {
	value *InviteMembersRequest
	isSet bool
}

func (v NullableInviteMembersRequest) Get() *InviteMembersRequest {
	return v.value
}

func (v *NullableInviteMembersRequest) Set(val *InviteMembersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteMembersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteMembersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteMembersRequest(val *InviteMembersRequest) *NullableInviteMembersRequest {
	return &NullableInviteMembersRequest{value: val, isSet: true}
}

func (v NullableInviteMembersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteMembersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


