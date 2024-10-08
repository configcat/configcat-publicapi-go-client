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

// checks if the IntegrationLinkDetailsModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationLinkDetailsModel{}

// IntegrationLinkDetailsModel struct for IntegrationLinkDetailsModel
type IntegrationLinkDetailsModel struct {
	Details []IntegrationLinkDetail `json:"details,omitempty"`
	AllIntegrationLinkCount *int32 `json:"allIntegrationLinkCount,omitempty"`
}

// NewIntegrationLinkDetailsModel instantiates a new IntegrationLinkDetailsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationLinkDetailsModel() *IntegrationLinkDetailsModel {
	this := IntegrationLinkDetailsModel{}
	return &this
}

// NewIntegrationLinkDetailsModelWithDefaults instantiates a new IntegrationLinkDetailsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationLinkDetailsModelWithDefaults() *IntegrationLinkDetailsModel {
	this := IntegrationLinkDetailsModel{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationLinkDetailsModel) GetDetails() []IntegrationLinkDetail {
	if o == nil {
		var ret []IntegrationLinkDetail
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationLinkDetailsModel) GetDetailsOk() ([]IntegrationLinkDetail, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *IntegrationLinkDetailsModel) HasDetails() bool {
	if o != nil && IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []IntegrationLinkDetail and assigns it to the Details field.
func (o *IntegrationLinkDetailsModel) SetDetails(v []IntegrationLinkDetail) {
	o.Details = v
}

// GetAllIntegrationLinkCount returns the AllIntegrationLinkCount field value if set, zero value otherwise.
func (o *IntegrationLinkDetailsModel) GetAllIntegrationLinkCount() int32 {
	if o == nil || IsNil(o.AllIntegrationLinkCount) {
		var ret int32
		return ret
	}
	return *o.AllIntegrationLinkCount
}

// GetAllIntegrationLinkCountOk returns a tuple with the AllIntegrationLinkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationLinkDetailsModel) GetAllIntegrationLinkCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AllIntegrationLinkCount) {
		return nil, false
	}
	return o.AllIntegrationLinkCount, true
}

// HasAllIntegrationLinkCount returns a boolean if a field has been set.
func (o *IntegrationLinkDetailsModel) HasAllIntegrationLinkCount() bool {
	if o != nil && !IsNil(o.AllIntegrationLinkCount) {
		return true
	}

	return false
}

// SetAllIntegrationLinkCount gets a reference to the given int32 and assigns it to the AllIntegrationLinkCount field.
func (o *IntegrationLinkDetailsModel) SetAllIntegrationLinkCount(v int32) {
	o.AllIntegrationLinkCount = &v
}

func (o IntegrationLinkDetailsModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationLinkDetailsModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.AllIntegrationLinkCount) {
		toSerialize["allIntegrationLinkCount"] = o.AllIntegrationLinkCount
	}
	return toSerialize, nil
}

type NullableIntegrationLinkDetailsModel struct {
	value *IntegrationLinkDetailsModel
	isSet bool
}

func (v NullableIntegrationLinkDetailsModel) Get() *IntegrationLinkDetailsModel {
	return v.value
}

func (v *NullableIntegrationLinkDetailsModel) Set(val *IntegrationLinkDetailsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationLinkDetailsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationLinkDetailsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationLinkDetailsModel(val *IntegrationLinkDetailsModel) *NullableIntegrationLinkDetailsModel {
	return &NullableIntegrationLinkDetailsModel{value: val, isSet: true}
}

func (v NullableIntegrationLinkDetailsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationLinkDetailsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

