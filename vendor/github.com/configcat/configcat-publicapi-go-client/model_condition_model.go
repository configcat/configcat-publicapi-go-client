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

// checks if the ConditionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionModel{}

// ConditionModel struct for ConditionModel
type ConditionModel struct {
	UserCondition *UserConditionModel `json:"userCondition,omitempty"`
	SegmentCondition *SegmentConditionModel `json:"segmentCondition,omitempty"`
	PrerequisiteFlagCondition *PrerequisiteFlagConditionModel `json:"prerequisiteFlagCondition,omitempty"`
}

// NewConditionModel instantiates a new ConditionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionModel() *ConditionModel {
	this := ConditionModel{}
	return &this
}

// NewConditionModelWithDefaults instantiates a new ConditionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionModelWithDefaults() *ConditionModel {
	this := ConditionModel{}
	return &this
}

// GetUserCondition returns the UserCondition field value if set, zero value otherwise.
func (o *ConditionModel) GetUserCondition() UserConditionModel {
	if o == nil || IsNil(o.UserCondition) {
		var ret UserConditionModel
		return ret
	}
	return *o.UserCondition
}

// GetUserConditionOk returns a tuple with the UserCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionModel) GetUserConditionOk() (*UserConditionModel, bool) {
	if o == nil || IsNil(o.UserCondition) {
		return nil, false
	}
	return o.UserCondition, true
}

// HasUserCondition returns a boolean if a field has been set.
func (o *ConditionModel) HasUserCondition() bool {
	if o != nil && !IsNil(o.UserCondition) {
		return true
	}

	return false
}

// SetUserCondition gets a reference to the given UserConditionModel and assigns it to the UserCondition field.
func (o *ConditionModel) SetUserCondition(v UserConditionModel) {
	o.UserCondition = &v
}

// GetSegmentCondition returns the SegmentCondition field value if set, zero value otherwise.
func (o *ConditionModel) GetSegmentCondition() SegmentConditionModel {
	if o == nil || IsNil(o.SegmentCondition) {
		var ret SegmentConditionModel
		return ret
	}
	return *o.SegmentCondition
}

// GetSegmentConditionOk returns a tuple with the SegmentCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionModel) GetSegmentConditionOk() (*SegmentConditionModel, bool) {
	if o == nil || IsNil(o.SegmentCondition) {
		return nil, false
	}
	return o.SegmentCondition, true
}

// HasSegmentCondition returns a boolean if a field has been set.
func (o *ConditionModel) HasSegmentCondition() bool {
	if o != nil && !IsNil(o.SegmentCondition) {
		return true
	}

	return false
}

// SetSegmentCondition gets a reference to the given SegmentConditionModel and assigns it to the SegmentCondition field.
func (o *ConditionModel) SetSegmentCondition(v SegmentConditionModel) {
	o.SegmentCondition = &v
}

// GetPrerequisiteFlagCondition returns the PrerequisiteFlagCondition field value if set, zero value otherwise.
func (o *ConditionModel) GetPrerequisiteFlagCondition() PrerequisiteFlagConditionModel {
	if o == nil || IsNil(o.PrerequisiteFlagCondition) {
		var ret PrerequisiteFlagConditionModel
		return ret
	}
	return *o.PrerequisiteFlagCondition
}

// GetPrerequisiteFlagConditionOk returns a tuple with the PrerequisiteFlagCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionModel) GetPrerequisiteFlagConditionOk() (*PrerequisiteFlagConditionModel, bool) {
	if o == nil || IsNil(o.PrerequisiteFlagCondition) {
		return nil, false
	}
	return o.PrerequisiteFlagCondition, true
}

// HasPrerequisiteFlagCondition returns a boolean if a field has been set.
func (o *ConditionModel) HasPrerequisiteFlagCondition() bool {
	if o != nil && !IsNil(o.PrerequisiteFlagCondition) {
		return true
	}

	return false
}

// SetPrerequisiteFlagCondition gets a reference to the given PrerequisiteFlagConditionModel and assigns it to the PrerequisiteFlagCondition field.
func (o *ConditionModel) SetPrerequisiteFlagCondition(v PrerequisiteFlagConditionModel) {
	o.PrerequisiteFlagCondition = &v
}

func (o ConditionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserCondition) {
		toSerialize["userCondition"] = o.UserCondition
	}
	if !IsNil(o.SegmentCondition) {
		toSerialize["segmentCondition"] = o.SegmentCondition
	}
	if !IsNil(o.PrerequisiteFlagCondition) {
		toSerialize["prerequisiteFlagCondition"] = o.PrerequisiteFlagCondition
	}
	return toSerialize, nil
}

type NullableConditionModel struct {
	value *ConditionModel
	isSet bool
}

func (v NullableConditionModel) Get() *ConditionModel {
	return v.value
}

func (v *NullableConditionModel) Set(val *ConditionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionModel(val *ConditionModel) *NullableConditionModel {
	return &NullableConditionModel{value: val, isSet: true}
}

func (v NullableConditionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


