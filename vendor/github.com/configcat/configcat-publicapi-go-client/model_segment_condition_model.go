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

// checks if the SegmentConditionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentConditionModel{}

// SegmentConditionModel Describes a condition that is based on a segment.
type SegmentConditionModel struct {
	// The segment's identifier.
	SegmentId string `json:"segmentId"`
	Comparator SegmentComparator `json:"comparator"`
}

// NewSegmentConditionModel instantiates a new SegmentConditionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentConditionModel(segmentId string, comparator SegmentComparator) *SegmentConditionModel {
	this := SegmentConditionModel{}
	this.SegmentId = segmentId
	this.Comparator = comparator
	return &this
}

// NewSegmentConditionModelWithDefaults instantiates a new SegmentConditionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentConditionModelWithDefaults() *SegmentConditionModel {
	this := SegmentConditionModel{}
	return &this
}

// GetSegmentId returns the SegmentId field value
func (o *SegmentConditionModel) GetSegmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SegmentId
}

// GetSegmentIdOk returns a tuple with the SegmentId field value
// and a boolean to check if the value has been set.
func (o *SegmentConditionModel) GetSegmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SegmentId, true
}

// SetSegmentId sets field value
func (o *SegmentConditionModel) SetSegmentId(v string) {
	o.SegmentId = v
}

// GetComparator returns the Comparator field value
func (o *SegmentConditionModel) GetComparator() SegmentComparator {
	if o == nil {
		var ret SegmentComparator
		return ret
	}

	return o.Comparator
}

// GetComparatorOk returns a tuple with the Comparator field value
// and a boolean to check if the value has been set.
func (o *SegmentConditionModel) GetComparatorOk() (*SegmentComparator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comparator, true
}

// SetComparator sets field value
func (o *SegmentConditionModel) SetComparator(v SegmentComparator) {
	o.Comparator = v
}

func (o SegmentConditionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentConditionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["segmentId"] = o.SegmentId
	toSerialize["comparator"] = o.Comparator
	return toSerialize, nil
}

type NullableSegmentConditionModel struct {
	value *SegmentConditionModel
	isSet bool
}

func (v NullableSegmentConditionModel) Get() *SegmentConditionModel {
	return v.value
}

func (v *NullableSegmentConditionModel) Set(val *SegmentConditionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentConditionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentConditionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentConditionModel(val *SegmentConditionModel) *NullableSegmentConditionModel {
	return &NullableSegmentConditionModel{value: val, isSet: true}
}

func (v NullableSegmentConditionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentConditionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


