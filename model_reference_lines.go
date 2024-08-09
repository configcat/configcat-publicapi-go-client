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

// checks if the ReferenceLines type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferenceLines{}

// ReferenceLines struct for ReferenceLines
type ReferenceLines struct {
	// The file's name in where the code reference has been found. (Appears on the ConfigCat Dashboard)
	File string `json:"file"`
	// The file's url. (Used to point to the file on the repository's website)
	FileUrl NullableString `json:"fileUrl,omitempty"`
	// The lines before the actual reference line.
	PreLines []ReferenceLine `json:"preLines,omitempty"`
	// The lines after the actual reference line.
	PostLines []ReferenceLine `json:"postLines,omitempty"`
	ReferenceLine ReferenceLine `json:"referenceLine"`
}

// NewReferenceLines instantiates a new ReferenceLines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceLines(file string, referenceLine ReferenceLine) *ReferenceLines {
	this := ReferenceLines{}
	this.File = file
	this.ReferenceLine = referenceLine
	return &this
}

// NewReferenceLinesWithDefaults instantiates a new ReferenceLines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceLinesWithDefaults() *ReferenceLines {
	this := ReferenceLines{}
	return &this
}

// GetFile returns the File field value
func (o *ReferenceLines) GetFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *ReferenceLines) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *ReferenceLines) SetFile(v string) {
	o.File = v
}

// GetFileUrl returns the FileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReferenceLines) GetFileUrl() string {
	if o == nil || IsNil(o.FileUrl.Get()) {
		var ret string
		return ret
	}
	return *o.FileUrl.Get()
}

// GetFileUrlOk returns a tuple with the FileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReferenceLines) GetFileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileUrl.Get(), o.FileUrl.IsSet()
}

// HasFileUrl returns a boolean if a field has been set.
func (o *ReferenceLines) HasFileUrl() bool {
	if o != nil && o.FileUrl.IsSet() {
		return true
	}

	return false
}

// SetFileUrl gets a reference to the given NullableString and assigns it to the FileUrl field.
func (o *ReferenceLines) SetFileUrl(v string) {
	o.FileUrl.Set(&v)
}
// SetFileUrlNil sets the value for FileUrl to be an explicit nil
func (o *ReferenceLines) SetFileUrlNil() {
	o.FileUrl.Set(nil)
}

// UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
func (o *ReferenceLines) UnsetFileUrl() {
	o.FileUrl.Unset()
}

// GetPreLines returns the PreLines field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReferenceLines) GetPreLines() []ReferenceLine {
	if o == nil {
		var ret []ReferenceLine
		return ret
	}
	return o.PreLines
}

// GetPreLinesOk returns a tuple with the PreLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReferenceLines) GetPreLinesOk() ([]ReferenceLine, bool) {
	if o == nil || IsNil(o.PreLines) {
		return nil, false
	}
	return o.PreLines, true
}

// HasPreLines returns a boolean if a field has been set.
func (o *ReferenceLines) HasPreLines() bool {
	if o != nil && IsNil(o.PreLines) {
		return true
	}

	return false
}

// SetPreLines gets a reference to the given []ReferenceLine and assigns it to the PreLines field.
func (o *ReferenceLines) SetPreLines(v []ReferenceLine) {
	o.PreLines = v
}

// GetPostLines returns the PostLines field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReferenceLines) GetPostLines() []ReferenceLine {
	if o == nil {
		var ret []ReferenceLine
		return ret
	}
	return o.PostLines
}

// GetPostLinesOk returns a tuple with the PostLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReferenceLines) GetPostLinesOk() ([]ReferenceLine, bool) {
	if o == nil || IsNil(o.PostLines) {
		return nil, false
	}
	return o.PostLines, true
}

// HasPostLines returns a boolean if a field has been set.
func (o *ReferenceLines) HasPostLines() bool {
	if o != nil && IsNil(o.PostLines) {
		return true
	}

	return false
}

// SetPostLines gets a reference to the given []ReferenceLine and assigns it to the PostLines field.
func (o *ReferenceLines) SetPostLines(v []ReferenceLine) {
	o.PostLines = v
}

// GetReferenceLine returns the ReferenceLine field value
func (o *ReferenceLines) GetReferenceLine() ReferenceLine {
	if o == nil {
		var ret ReferenceLine
		return ret
	}

	return o.ReferenceLine
}

// GetReferenceLineOk returns a tuple with the ReferenceLine field value
// and a boolean to check if the value has been set.
func (o *ReferenceLines) GetReferenceLineOk() (*ReferenceLine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceLine, true
}

// SetReferenceLine sets field value
func (o *ReferenceLines) SetReferenceLine(v ReferenceLine) {
	o.ReferenceLine = v
}

func (o ReferenceLines) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferenceLines) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file"] = o.File
	if o.FileUrl.IsSet() {
		toSerialize["fileUrl"] = o.FileUrl.Get()
	}
	if o.PreLines != nil {
		toSerialize["preLines"] = o.PreLines
	}
	if o.PostLines != nil {
		toSerialize["postLines"] = o.PostLines
	}
	toSerialize["referenceLine"] = o.ReferenceLine
	return toSerialize, nil
}

type NullableReferenceLines struct {
	value *ReferenceLines
	isSet bool
}

func (v NullableReferenceLines) Get() *ReferenceLines {
	return v.value
}

func (v *NullableReferenceLines) Set(val *ReferenceLines) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceLines) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceLines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceLines(val *ReferenceLines) *NullableReferenceLines {
	return &NullableReferenceLines{value: val, isSet: true}
}

func (v NullableReferenceLines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceLines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


