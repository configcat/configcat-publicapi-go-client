/*
ConfigCat Public Management API

**Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The purpose of this API is to access the ConfigCat platform programmatically.  You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.   The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  and JSON+HAL format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available here: [swagger.json](v1/swagger.json).  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
)

// checks if the CreateOrUpdateEnvironmentAccessModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateEnvironmentAccessModel{}

// CreateOrUpdateEnvironmentAccessModel struct for CreateOrUpdateEnvironmentAccessModel
type CreateOrUpdateEnvironmentAccessModel struct {
	EnvironmentId *string `json:"environmentId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Color NullableString `json:"color,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Order *int32 `json:"order,omitempty"`
	ReasonRequired *bool `json:"reasonRequired,omitempty"`
	EnvironmentAccessType *EnvironmentAccessType `json:"environmentAccessType,omitempty"`
}

// NewCreateOrUpdateEnvironmentAccessModel instantiates a new CreateOrUpdateEnvironmentAccessModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateEnvironmentAccessModel() *CreateOrUpdateEnvironmentAccessModel {
	this := CreateOrUpdateEnvironmentAccessModel{}
	return &this
}

// NewCreateOrUpdateEnvironmentAccessModelWithDefaults instantiates a new CreateOrUpdateEnvironmentAccessModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateEnvironmentAccessModelWithDefaults() *CreateOrUpdateEnvironmentAccessModel {
	this := CreateOrUpdateEnvironmentAccessModel{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *CreateOrUpdateEnvironmentAccessModel) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateEnvironmentAccessModel) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *CreateOrUpdateEnvironmentAccessModel) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *CreateOrUpdateEnvironmentAccessModel) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateEnvironmentAccessModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateEnvironmentAccessModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrUpdateEnvironmentAccessModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateOrUpdateEnvironmentAccessModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateOrUpdateEnvironmentAccessModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateOrUpdateEnvironmentAccessModel) UnsetName() {
	o.Name.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateEnvironmentAccessModel) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateEnvironmentAccessModel) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *CreateOrUpdateEnvironmentAccessModel) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *CreateOrUpdateEnvironmentAccessModel) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *CreateOrUpdateEnvironmentAccessModel) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *CreateOrUpdateEnvironmentAccessModel) UnsetColor() {
	o.Color.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateEnvironmentAccessModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateEnvironmentAccessModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrUpdateEnvironmentAccessModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateOrUpdateEnvironmentAccessModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateOrUpdateEnvironmentAccessModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateOrUpdateEnvironmentAccessModel) UnsetDescription() {
	o.Description.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *CreateOrUpdateEnvironmentAccessModel) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateEnvironmentAccessModel) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *CreateOrUpdateEnvironmentAccessModel) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *CreateOrUpdateEnvironmentAccessModel) SetOrder(v int32) {
	o.Order = &v
}

// GetReasonRequired returns the ReasonRequired field value if set, zero value otherwise.
func (o *CreateOrUpdateEnvironmentAccessModel) GetReasonRequired() bool {
	if o == nil || IsNil(o.ReasonRequired) {
		var ret bool
		return ret
	}
	return *o.ReasonRequired
}

// GetReasonRequiredOk returns a tuple with the ReasonRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateEnvironmentAccessModel) GetReasonRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ReasonRequired) {
		return nil, false
	}
	return o.ReasonRequired, true
}

// HasReasonRequired returns a boolean if a field has been set.
func (o *CreateOrUpdateEnvironmentAccessModel) HasReasonRequired() bool {
	if o != nil && !IsNil(o.ReasonRequired) {
		return true
	}

	return false
}

// SetReasonRequired gets a reference to the given bool and assigns it to the ReasonRequired field.
func (o *CreateOrUpdateEnvironmentAccessModel) SetReasonRequired(v bool) {
	o.ReasonRequired = &v
}

// GetEnvironmentAccessType returns the EnvironmentAccessType field value if set, zero value otherwise.
func (o *CreateOrUpdateEnvironmentAccessModel) GetEnvironmentAccessType() EnvironmentAccessType {
	if o == nil || IsNil(o.EnvironmentAccessType) {
		var ret EnvironmentAccessType
		return ret
	}
	return *o.EnvironmentAccessType
}

// GetEnvironmentAccessTypeOk returns a tuple with the EnvironmentAccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateEnvironmentAccessModel) GetEnvironmentAccessTypeOk() (*EnvironmentAccessType, bool) {
	if o == nil || IsNil(o.EnvironmentAccessType) {
		return nil, false
	}
	return o.EnvironmentAccessType, true
}

// HasEnvironmentAccessType returns a boolean if a field has been set.
func (o *CreateOrUpdateEnvironmentAccessModel) HasEnvironmentAccessType() bool {
	if o != nil && !IsNil(o.EnvironmentAccessType) {
		return true
	}

	return false
}

// SetEnvironmentAccessType gets a reference to the given EnvironmentAccessType and assigns it to the EnvironmentAccessType field.
func (o *CreateOrUpdateEnvironmentAccessModel) SetEnvironmentAccessType(v EnvironmentAccessType) {
	o.EnvironmentAccessType = &v
}

func (o CreateOrUpdateEnvironmentAccessModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateEnvironmentAccessModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.ReasonRequired) {
		toSerialize["reasonRequired"] = o.ReasonRequired
	}
	if !IsNil(o.EnvironmentAccessType) {
		toSerialize["environmentAccessType"] = o.EnvironmentAccessType
	}
	return toSerialize, nil
}

type NullableCreateOrUpdateEnvironmentAccessModel struct {
	value *CreateOrUpdateEnvironmentAccessModel
	isSet bool
}

func (v NullableCreateOrUpdateEnvironmentAccessModel) Get() *CreateOrUpdateEnvironmentAccessModel {
	return v.value
}

func (v *NullableCreateOrUpdateEnvironmentAccessModel) Set(val *CreateOrUpdateEnvironmentAccessModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateEnvironmentAccessModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateEnvironmentAccessModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateEnvironmentAccessModel(val *CreateOrUpdateEnvironmentAccessModel) *NullableCreateOrUpdateEnvironmentAccessModel {
	return &NullableCreateOrUpdateEnvironmentAccessModel{value: val, isSet: true}
}

func (v NullableCreateOrUpdateEnvironmentAccessModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateEnvironmentAccessModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


