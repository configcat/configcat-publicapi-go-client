/*
ConfigCat Public Management API

The purpose of this API is to access the ConfigCat platform programmatically. You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.  **Base API URL**: https://api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://api.configcat.com/swagger).  The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  format.   **Important:** Do not use this API for accessing and evaluating feature flag values. Use the [SDKs](https://configcat.com/docs/sdk-reference/overview) or the [ConfigCat Proxy](https://configcat.com/docs/advanced/proxy/proxy-overview/) instead.  # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
)

// checks if the ReferenceLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferenceLine{}

// ReferenceLine Determines a code reference line.
type ReferenceLine struct {
	// The content of the reference line.
	LineText NullableString `json:"lineText,omitempty"`
	// The line number.
	LineNumber int32 `json:"lineNumber"`
}

// NewReferenceLine instantiates a new ReferenceLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceLine(lineNumber int32) *ReferenceLine {
	this := ReferenceLine{}
	this.LineNumber = lineNumber
	return &this
}

// NewReferenceLineWithDefaults instantiates a new ReferenceLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceLineWithDefaults() *ReferenceLine {
	this := ReferenceLine{}
	return &this
}

// GetLineText returns the LineText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReferenceLine) GetLineText() string {
	if o == nil || IsNil(o.LineText.Get()) {
		var ret string
		return ret
	}
	return *o.LineText.Get()
}

// GetLineTextOk returns a tuple with the LineText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReferenceLine) GetLineTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineText.Get(), o.LineText.IsSet()
}

// HasLineText returns a boolean if a field has been set.
func (o *ReferenceLine) HasLineText() bool {
	if o != nil && o.LineText.IsSet() {
		return true
	}

	return false
}

// SetLineText gets a reference to the given NullableString and assigns it to the LineText field.
func (o *ReferenceLine) SetLineText(v string) {
	o.LineText.Set(&v)
}
// SetLineTextNil sets the value for LineText to be an explicit nil
func (o *ReferenceLine) SetLineTextNil() {
	o.LineText.Set(nil)
}

// UnsetLineText ensures that no value is present for LineText, not even an explicit nil
func (o *ReferenceLine) UnsetLineText() {
	o.LineText.Unset()
}

// GetLineNumber returns the LineNumber field value
func (o *ReferenceLine) GetLineNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LineNumber
}

// GetLineNumberOk returns a tuple with the LineNumber field value
// and a boolean to check if the value has been set.
func (o *ReferenceLine) GetLineNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LineNumber, true
}

// SetLineNumber sets field value
func (o *ReferenceLine) SetLineNumber(v int32) {
	o.LineNumber = v
}

func (o ReferenceLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferenceLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LineText.IsSet() {
		toSerialize["lineText"] = o.LineText.Get()
	}
	toSerialize["lineNumber"] = o.LineNumber
	return toSerialize, nil
}

type NullableReferenceLine struct {
	value *ReferenceLine
	isSet bool
}

func (v NullableReferenceLine) Get() *ReferenceLine {
	return v.value
}

func (v *NullableReferenceLine) Set(val *ReferenceLine) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceLine) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceLine(val *ReferenceLine) *NullableReferenceLine {
	return &NullableReferenceLine{value: val, isSet: true}
}

func (v NullableReferenceLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


