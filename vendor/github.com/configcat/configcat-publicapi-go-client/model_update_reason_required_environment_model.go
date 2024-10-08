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

// checks if the UpdateReasonRequiredEnvironmentModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateReasonRequiredEnvironmentModel{}

// UpdateReasonRequiredEnvironmentModel struct for UpdateReasonRequiredEnvironmentModel
type UpdateReasonRequiredEnvironmentModel struct {
	// Identifier of the Environment.
	EnvironmentId *string `json:"environmentId,omitempty"`
	// Indicates that a mandatory note is required in this Environment for saving and publishing.
	ReasonRequired *bool `json:"reasonRequired,omitempty"`
}

// NewUpdateReasonRequiredEnvironmentModel instantiates a new UpdateReasonRequiredEnvironmentModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReasonRequiredEnvironmentModel() *UpdateReasonRequiredEnvironmentModel {
	this := UpdateReasonRequiredEnvironmentModel{}
	return &this
}

// NewUpdateReasonRequiredEnvironmentModelWithDefaults instantiates a new UpdateReasonRequiredEnvironmentModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReasonRequiredEnvironmentModelWithDefaults() *UpdateReasonRequiredEnvironmentModel {
	this := UpdateReasonRequiredEnvironmentModel{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *UpdateReasonRequiredEnvironmentModel) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReasonRequiredEnvironmentModel) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *UpdateReasonRequiredEnvironmentModel) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *UpdateReasonRequiredEnvironmentModel) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetReasonRequired returns the ReasonRequired field value if set, zero value otherwise.
func (o *UpdateReasonRequiredEnvironmentModel) GetReasonRequired() bool {
	if o == nil || IsNil(o.ReasonRequired) {
		var ret bool
		return ret
	}
	return *o.ReasonRequired
}

// GetReasonRequiredOk returns a tuple with the ReasonRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReasonRequiredEnvironmentModel) GetReasonRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ReasonRequired) {
		return nil, false
	}
	return o.ReasonRequired, true
}

// HasReasonRequired returns a boolean if a field has been set.
func (o *UpdateReasonRequiredEnvironmentModel) HasReasonRequired() bool {
	if o != nil && !IsNil(o.ReasonRequired) {
		return true
	}

	return false
}

// SetReasonRequired gets a reference to the given bool and assigns it to the ReasonRequired field.
func (o *UpdateReasonRequiredEnvironmentModel) SetReasonRequired(v bool) {
	o.ReasonRequired = &v
}

func (o UpdateReasonRequiredEnvironmentModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateReasonRequiredEnvironmentModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	if !IsNil(o.ReasonRequired) {
		toSerialize["reasonRequired"] = o.ReasonRequired
	}
	return toSerialize, nil
}

type NullableUpdateReasonRequiredEnvironmentModel struct {
	value *UpdateReasonRequiredEnvironmentModel
	isSet bool
}

func (v NullableUpdateReasonRequiredEnvironmentModel) Get() *UpdateReasonRequiredEnvironmentModel {
	return v.value
}

func (v *NullableUpdateReasonRequiredEnvironmentModel) Set(val *UpdateReasonRequiredEnvironmentModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateReasonRequiredEnvironmentModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateReasonRequiredEnvironmentModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateReasonRequiredEnvironmentModel(val *UpdateReasonRequiredEnvironmentModel) *NullableUpdateReasonRequiredEnvironmentModel {
	return &NullableUpdateReasonRequiredEnvironmentModel{value: val, isSet: true}
}

func (v NullableUpdateReasonRequiredEnvironmentModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateReasonRequiredEnvironmentModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


