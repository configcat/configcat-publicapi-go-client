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
	"time"
)

// checks if the SegmentModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentModel{}

// SegmentModel struct for SegmentModel
type SegmentModel struct {
	Product *ProductModel `json:"product,omitempty"`
	// Identifier of the Segment.
	SegmentId *string `json:"segmentId,omitempty"`
	// Name of the Segment.
	Name NullableString `json:"name,omitempty"`
	// Description of the Segment.
	Description NullableString `json:"description,omitempty"`
	// The email of the user who created the Segment.
	CreatorEmail NullableString `json:"creatorEmail,omitempty"`
	// The name of the user who created the Segment.
	CreatorFullName NullableString `json:"creatorFullName,omitempty"`
	// The date and time when the Segment was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The email of the user who last updated the Segment.
	LastUpdaterEmail NullableString `json:"lastUpdaterEmail,omitempty"`
	// The name of the user who last updated the Segment.
	LastUpdaterFullName NullableString `json:"lastUpdaterFullName,omitempty"`
	// The date and time when the Segment was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// The user's attribute the evaluation process must take into account.
	ComparisonAttribute NullableString `json:"comparisonAttribute,omitempty"`
	Comparator *RolloutRuleComparator `json:"comparator,omitempty"`
	// The value to compare with the given user attribute's value.
	ComparisonValue NullableString `json:"comparisonValue,omitempty"`
}

// NewSegmentModel instantiates a new SegmentModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentModel() *SegmentModel {
	this := SegmentModel{}
	return &this
}

// NewSegmentModelWithDefaults instantiates a new SegmentModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentModelWithDefaults() *SegmentModel {
	this := SegmentModel{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *SegmentModel) GetProduct() ProductModel {
	if o == nil || IsNil(o.Product) {
		var ret ProductModel
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentModel) GetProductOk() (*ProductModel, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *SegmentModel) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given ProductModel and assigns it to the Product field.
func (o *SegmentModel) SetProduct(v ProductModel) {
	o.Product = &v
}

// GetSegmentId returns the SegmentId field value if set, zero value otherwise.
func (o *SegmentModel) GetSegmentId() string {
	if o == nil || IsNil(o.SegmentId) {
		var ret string
		return ret
	}
	return *o.SegmentId
}

// GetSegmentIdOk returns a tuple with the SegmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentModel) GetSegmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.SegmentId) {
		return nil, false
	}
	return o.SegmentId, true
}

// HasSegmentId returns a boolean if a field has been set.
func (o *SegmentModel) HasSegmentId() bool {
	if o != nil && !IsNil(o.SegmentId) {
		return true
	}

	return false
}

// SetSegmentId gets a reference to the given string and assigns it to the SegmentId field.
func (o *SegmentModel) SetSegmentId(v string) {
	o.SegmentId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SegmentModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SegmentModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SegmentModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SegmentModel) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SegmentModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SegmentModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SegmentModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SegmentModel) UnsetDescription() {
	o.Description.Unset()
}

// GetCreatorEmail returns the CreatorEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentModel) GetCreatorEmail() string {
	if o == nil || IsNil(o.CreatorEmail.Get()) {
		var ret string
		return ret
	}
	return *o.CreatorEmail.Get()
}

// GetCreatorEmailOk returns a tuple with the CreatorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentModel) GetCreatorEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorEmail.Get(), o.CreatorEmail.IsSet()
}

// HasCreatorEmail returns a boolean if a field has been set.
func (o *SegmentModel) HasCreatorEmail() bool {
	if o != nil && o.CreatorEmail.IsSet() {
		return true
	}

	return false
}

// SetCreatorEmail gets a reference to the given NullableString and assigns it to the CreatorEmail field.
func (o *SegmentModel) SetCreatorEmail(v string) {
	o.CreatorEmail.Set(&v)
}
// SetCreatorEmailNil sets the value for CreatorEmail to be an explicit nil
func (o *SegmentModel) SetCreatorEmailNil() {
	o.CreatorEmail.Set(nil)
}

// UnsetCreatorEmail ensures that no value is present for CreatorEmail, not even an explicit nil
func (o *SegmentModel) UnsetCreatorEmail() {
	o.CreatorEmail.Unset()
}

// GetCreatorFullName returns the CreatorFullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentModel) GetCreatorFullName() string {
	if o == nil || IsNil(o.CreatorFullName.Get()) {
		var ret string
		return ret
	}
	return *o.CreatorFullName.Get()
}

// GetCreatorFullNameOk returns a tuple with the CreatorFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentModel) GetCreatorFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorFullName.Get(), o.CreatorFullName.IsSet()
}

// HasCreatorFullName returns a boolean if a field has been set.
func (o *SegmentModel) HasCreatorFullName() bool {
	if o != nil && o.CreatorFullName.IsSet() {
		return true
	}

	return false
}

// SetCreatorFullName gets a reference to the given NullableString and assigns it to the CreatorFullName field.
func (o *SegmentModel) SetCreatorFullName(v string) {
	o.CreatorFullName.Set(&v)
}
// SetCreatorFullNameNil sets the value for CreatorFullName to be an explicit nil
func (o *SegmentModel) SetCreatorFullNameNil() {
	o.CreatorFullName.Set(nil)
}

// UnsetCreatorFullName ensures that no value is present for CreatorFullName, not even an explicit nil
func (o *SegmentModel) UnsetCreatorFullName() {
	o.CreatorFullName.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SegmentModel) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SegmentModel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SegmentModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdaterEmail returns the LastUpdaterEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentModel) GetLastUpdaterEmail() string {
	if o == nil || IsNil(o.LastUpdaterEmail.Get()) {
		var ret string
		return ret
	}
	return *o.LastUpdaterEmail.Get()
}

// GetLastUpdaterEmailOk returns a tuple with the LastUpdaterEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentModel) GetLastUpdaterEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdaterEmail.Get(), o.LastUpdaterEmail.IsSet()
}

// HasLastUpdaterEmail returns a boolean if a field has been set.
func (o *SegmentModel) HasLastUpdaterEmail() bool {
	if o != nil && o.LastUpdaterEmail.IsSet() {
		return true
	}

	return false
}

// SetLastUpdaterEmail gets a reference to the given NullableString and assigns it to the LastUpdaterEmail field.
func (o *SegmentModel) SetLastUpdaterEmail(v string) {
	o.LastUpdaterEmail.Set(&v)
}
// SetLastUpdaterEmailNil sets the value for LastUpdaterEmail to be an explicit nil
func (o *SegmentModel) SetLastUpdaterEmailNil() {
	o.LastUpdaterEmail.Set(nil)
}

// UnsetLastUpdaterEmail ensures that no value is present for LastUpdaterEmail, not even an explicit nil
func (o *SegmentModel) UnsetLastUpdaterEmail() {
	o.LastUpdaterEmail.Unset()
}

// GetLastUpdaterFullName returns the LastUpdaterFullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentModel) GetLastUpdaterFullName() string {
	if o == nil || IsNil(o.LastUpdaterFullName.Get()) {
		var ret string
		return ret
	}
	return *o.LastUpdaterFullName.Get()
}

// GetLastUpdaterFullNameOk returns a tuple with the LastUpdaterFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentModel) GetLastUpdaterFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdaterFullName.Get(), o.LastUpdaterFullName.IsSet()
}

// HasLastUpdaterFullName returns a boolean if a field has been set.
func (o *SegmentModel) HasLastUpdaterFullName() bool {
	if o != nil && o.LastUpdaterFullName.IsSet() {
		return true
	}

	return false
}

// SetLastUpdaterFullName gets a reference to the given NullableString and assigns it to the LastUpdaterFullName field.
func (o *SegmentModel) SetLastUpdaterFullName(v string) {
	o.LastUpdaterFullName.Set(&v)
}
// SetLastUpdaterFullNameNil sets the value for LastUpdaterFullName to be an explicit nil
func (o *SegmentModel) SetLastUpdaterFullNameNil() {
	o.LastUpdaterFullName.Set(nil)
}

// UnsetLastUpdaterFullName ensures that no value is present for LastUpdaterFullName, not even an explicit nil
func (o *SegmentModel) UnsetLastUpdaterFullName() {
	o.LastUpdaterFullName.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SegmentModel) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentModel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SegmentModel) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SegmentModel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetComparisonAttribute returns the ComparisonAttribute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentModel) GetComparisonAttribute() string {
	if o == nil || IsNil(o.ComparisonAttribute.Get()) {
		var ret string
		return ret
	}
	return *o.ComparisonAttribute.Get()
}

// GetComparisonAttributeOk returns a tuple with the ComparisonAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentModel) GetComparisonAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComparisonAttribute.Get(), o.ComparisonAttribute.IsSet()
}

// HasComparisonAttribute returns a boolean if a field has been set.
func (o *SegmentModel) HasComparisonAttribute() bool {
	if o != nil && o.ComparisonAttribute.IsSet() {
		return true
	}

	return false
}

// SetComparisonAttribute gets a reference to the given NullableString and assigns it to the ComparisonAttribute field.
func (o *SegmentModel) SetComparisonAttribute(v string) {
	o.ComparisonAttribute.Set(&v)
}
// SetComparisonAttributeNil sets the value for ComparisonAttribute to be an explicit nil
func (o *SegmentModel) SetComparisonAttributeNil() {
	o.ComparisonAttribute.Set(nil)
}

// UnsetComparisonAttribute ensures that no value is present for ComparisonAttribute, not even an explicit nil
func (o *SegmentModel) UnsetComparisonAttribute() {
	o.ComparisonAttribute.Unset()
}

// GetComparator returns the Comparator field value if set, zero value otherwise.
func (o *SegmentModel) GetComparator() RolloutRuleComparator {
	if o == nil || IsNil(o.Comparator) {
		var ret RolloutRuleComparator
		return ret
	}
	return *o.Comparator
}

// GetComparatorOk returns a tuple with the Comparator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentModel) GetComparatorOk() (*RolloutRuleComparator, bool) {
	if o == nil || IsNil(o.Comparator) {
		return nil, false
	}
	return o.Comparator, true
}

// HasComparator returns a boolean if a field has been set.
func (o *SegmentModel) HasComparator() bool {
	if o != nil && !IsNil(o.Comparator) {
		return true
	}

	return false
}

// SetComparator gets a reference to the given RolloutRuleComparator and assigns it to the Comparator field.
func (o *SegmentModel) SetComparator(v RolloutRuleComparator) {
	o.Comparator = &v
}

// GetComparisonValue returns the ComparisonValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentModel) GetComparisonValue() string {
	if o == nil || IsNil(o.ComparisonValue.Get()) {
		var ret string
		return ret
	}
	return *o.ComparisonValue.Get()
}

// GetComparisonValueOk returns a tuple with the ComparisonValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentModel) GetComparisonValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComparisonValue.Get(), o.ComparisonValue.IsSet()
}

// HasComparisonValue returns a boolean if a field has been set.
func (o *SegmentModel) HasComparisonValue() bool {
	if o != nil && o.ComparisonValue.IsSet() {
		return true
	}

	return false
}

// SetComparisonValue gets a reference to the given NullableString and assigns it to the ComparisonValue field.
func (o *SegmentModel) SetComparisonValue(v string) {
	o.ComparisonValue.Set(&v)
}
// SetComparisonValueNil sets the value for ComparisonValue to be an explicit nil
func (o *SegmentModel) SetComparisonValueNil() {
	o.ComparisonValue.Set(nil)
}

// UnsetComparisonValue ensures that no value is present for ComparisonValue, not even an explicit nil
func (o *SegmentModel) UnsetComparisonValue() {
	o.ComparisonValue.Unset()
}

func (o SegmentModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.SegmentId) {
		toSerialize["segmentId"] = o.SegmentId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.CreatorEmail.IsSet() {
		toSerialize["creatorEmail"] = o.CreatorEmail.Get()
	}
	if o.CreatorFullName.IsSet() {
		toSerialize["creatorFullName"] = o.CreatorFullName.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.LastUpdaterEmail.IsSet() {
		toSerialize["lastUpdaterEmail"] = o.LastUpdaterEmail.Get()
	}
	if o.LastUpdaterFullName.IsSet() {
		toSerialize["lastUpdaterFullName"] = o.LastUpdaterFullName.Get()
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.ComparisonAttribute.IsSet() {
		toSerialize["comparisonAttribute"] = o.ComparisonAttribute.Get()
	}
	if !IsNil(o.Comparator) {
		toSerialize["comparator"] = o.Comparator
	}
	if o.ComparisonValue.IsSet() {
		toSerialize["comparisonValue"] = o.ComparisonValue.Get()
	}
	return toSerialize, nil
}

type NullableSegmentModel struct {
	value *SegmentModel
	isSet bool
}

func (v NullableSegmentModel) Get() *SegmentModel {
	return v.value
}

func (v *NullableSegmentModel) Set(val *SegmentModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentModel(val *SegmentModel) *NullableSegmentModel {
	return &NullableSegmentModel{value: val, isSet: true}
}

func (v NullableSegmentModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


