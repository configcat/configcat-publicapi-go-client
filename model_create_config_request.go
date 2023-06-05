/*
ConfigCat Public Management API

**Base API URL**: https://api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://api.configcat.com/swagger).  The purpose of this API is to access the ConfigCat platform programmatically.  You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.   The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  and JSON+HAL format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available here: [swagger.json](v1/swagger.json).  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
)

// checks if the CreateConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateConfigRequest{}

// CreateConfigRequest struct for CreateConfigRequest
type CreateConfigRequest struct {
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	EvaluationVersion *EvaluationVersion `json:"evaluationVersion,omitempty"`
}

// NewCreateConfigRequest instantiates a new CreateConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConfigRequest(name string) *CreateConfigRequest {
	this := CreateConfigRequest{}
	this.Name = name
	return &this
}

// NewCreateConfigRequestWithDefaults instantiates a new CreateConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConfigRequestWithDefaults() *CreateConfigRequest {
	this := CreateConfigRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateConfigRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateConfigRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateConfigRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateConfigRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateConfigRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateConfigRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateConfigRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateConfigRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateConfigRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetEvaluationVersion returns the EvaluationVersion field value if set, zero value otherwise.
func (o *CreateConfigRequest) GetEvaluationVersion() EvaluationVersion {
	if o == nil || IsNil(o.EvaluationVersion) {
		var ret EvaluationVersion
		return ret
	}
	return *o.EvaluationVersion
}

// GetEvaluationVersionOk returns a tuple with the EvaluationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfigRequest) GetEvaluationVersionOk() (*EvaluationVersion, bool) {
	if o == nil || IsNil(o.EvaluationVersion) {
		return nil, false
	}
	return o.EvaluationVersion, true
}

// HasEvaluationVersion returns a boolean if a field has been set.
func (o *CreateConfigRequest) HasEvaluationVersion() bool {
	if o != nil && !IsNil(o.EvaluationVersion) {
		return true
	}

	return false
}

// SetEvaluationVersion gets a reference to the given EvaluationVersion and assigns it to the EvaluationVersion field.
func (o *CreateConfigRequest) SetEvaluationVersion(v EvaluationVersion) {
	o.EvaluationVersion = &v
}

func (o CreateConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.EvaluationVersion) {
		toSerialize["evaluationVersion"] = o.EvaluationVersion
	}
	return toSerialize, nil
}

type NullableCreateConfigRequest struct {
	value *CreateConfigRequest
	isSet bool
}

func (v NullableCreateConfigRequest) Get() *CreateConfigRequest {
	return v.value
}

func (v *NullableCreateConfigRequest) Set(val *CreateConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConfigRequest(val *CreateConfigRequest) *NullableCreateConfigRequest {
	return &NullableCreateConfigRequest{value: val, isSet: true}
}

func (v NullableCreateConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


