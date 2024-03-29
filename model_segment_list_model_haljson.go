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
	"time"
)

// checks if the SegmentListModelHaljson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentListModelHaljson{}

// SegmentListModelHaljson struct for SegmentListModelHaljson
type SegmentListModelHaljson struct {
	Embedded *ConfigModelHaljsonEmbedded `json:"_embedded,omitempty"`
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
	// Determines how many Feature Flags and Settings are using the Segment.
	Usage *int32 `json:"usage,omitempty"`
	Links *ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks `json:"_links,omitempty"`
}

// NewSegmentListModelHaljson instantiates a new SegmentListModelHaljson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentListModelHaljson() *SegmentListModelHaljson {
	this := SegmentListModelHaljson{}
	return &this
}

// NewSegmentListModelHaljsonWithDefaults instantiates a new SegmentListModelHaljson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentListModelHaljsonWithDefaults() *SegmentListModelHaljson {
	this := SegmentListModelHaljson{}
	return &this
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *SegmentListModelHaljson) GetEmbedded() ConfigModelHaljsonEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret ConfigModelHaljsonEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentListModelHaljson) GetEmbeddedOk() (*ConfigModelHaljsonEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *SegmentListModelHaljson) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given ConfigModelHaljsonEmbedded and assigns it to the Embedded field.
func (o *SegmentListModelHaljson) SetEmbedded(v ConfigModelHaljsonEmbedded) {
	o.Embedded = &v
}

// GetSegmentId returns the SegmentId field value if set, zero value otherwise.
func (o *SegmentListModelHaljson) GetSegmentId() string {
	if o == nil || IsNil(o.SegmentId) {
		var ret string
		return ret
	}
	return *o.SegmentId
}

// GetSegmentIdOk returns a tuple with the SegmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentListModelHaljson) GetSegmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.SegmentId) {
		return nil, false
	}
	return o.SegmentId, true
}

// HasSegmentId returns a boolean if a field has been set.
func (o *SegmentListModelHaljson) HasSegmentId() bool {
	if o != nil && !IsNil(o.SegmentId) {
		return true
	}

	return false
}

// SetSegmentId gets a reference to the given string and assigns it to the SegmentId field.
func (o *SegmentListModelHaljson) SetSegmentId(v string) {
	o.SegmentId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentListModelHaljson) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentListModelHaljson) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SegmentListModelHaljson) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SegmentListModelHaljson) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SegmentListModelHaljson) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SegmentListModelHaljson) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentListModelHaljson) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentListModelHaljson) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SegmentListModelHaljson) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SegmentListModelHaljson) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SegmentListModelHaljson) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SegmentListModelHaljson) UnsetDescription() {
	o.Description.Unset()
}

// GetCreatorEmail returns the CreatorEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentListModelHaljson) GetCreatorEmail() string {
	if o == nil || IsNil(o.CreatorEmail.Get()) {
		var ret string
		return ret
	}
	return *o.CreatorEmail.Get()
}

// GetCreatorEmailOk returns a tuple with the CreatorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentListModelHaljson) GetCreatorEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorEmail.Get(), o.CreatorEmail.IsSet()
}

// HasCreatorEmail returns a boolean if a field has been set.
func (o *SegmentListModelHaljson) HasCreatorEmail() bool {
	if o != nil && o.CreatorEmail.IsSet() {
		return true
	}

	return false
}

// SetCreatorEmail gets a reference to the given NullableString and assigns it to the CreatorEmail field.
func (o *SegmentListModelHaljson) SetCreatorEmail(v string) {
	o.CreatorEmail.Set(&v)
}
// SetCreatorEmailNil sets the value for CreatorEmail to be an explicit nil
func (o *SegmentListModelHaljson) SetCreatorEmailNil() {
	o.CreatorEmail.Set(nil)
}

// UnsetCreatorEmail ensures that no value is present for CreatorEmail, not even an explicit nil
func (o *SegmentListModelHaljson) UnsetCreatorEmail() {
	o.CreatorEmail.Unset()
}

// GetCreatorFullName returns the CreatorFullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentListModelHaljson) GetCreatorFullName() string {
	if o == nil || IsNil(o.CreatorFullName.Get()) {
		var ret string
		return ret
	}
	return *o.CreatorFullName.Get()
}

// GetCreatorFullNameOk returns a tuple with the CreatorFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentListModelHaljson) GetCreatorFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorFullName.Get(), o.CreatorFullName.IsSet()
}

// HasCreatorFullName returns a boolean if a field has been set.
func (o *SegmentListModelHaljson) HasCreatorFullName() bool {
	if o != nil && o.CreatorFullName.IsSet() {
		return true
	}

	return false
}

// SetCreatorFullName gets a reference to the given NullableString and assigns it to the CreatorFullName field.
func (o *SegmentListModelHaljson) SetCreatorFullName(v string) {
	o.CreatorFullName.Set(&v)
}
// SetCreatorFullNameNil sets the value for CreatorFullName to be an explicit nil
func (o *SegmentListModelHaljson) SetCreatorFullNameNil() {
	o.CreatorFullName.Set(nil)
}

// UnsetCreatorFullName ensures that no value is present for CreatorFullName, not even an explicit nil
func (o *SegmentListModelHaljson) UnsetCreatorFullName() {
	o.CreatorFullName.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SegmentListModelHaljson) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentListModelHaljson) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SegmentListModelHaljson) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SegmentListModelHaljson) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdaterEmail returns the LastUpdaterEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentListModelHaljson) GetLastUpdaterEmail() string {
	if o == nil || IsNil(o.LastUpdaterEmail.Get()) {
		var ret string
		return ret
	}
	return *o.LastUpdaterEmail.Get()
}

// GetLastUpdaterEmailOk returns a tuple with the LastUpdaterEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentListModelHaljson) GetLastUpdaterEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdaterEmail.Get(), o.LastUpdaterEmail.IsSet()
}

// HasLastUpdaterEmail returns a boolean if a field has been set.
func (o *SegmentListModelHaljson) HasLastUpdaterEmail() bool {
	if o != nil && o.LastUpdaterEmail.IsSet() {
		return true
	}

	return false
}

// SetLastUpdaterEmail gets a reference to the given NullableString and assigns it to the LastUpdaterEmail field.
func (o *SegmentListModelHaljson) SetLastUpdaterEmail(v string) {
	o.LastUpdaterEmail.Set(&v)
}
// SetLastUpdaterEmailNil sets the value for LastUpdaterEmail to be an explicit nil
func (o *SegmentListModelHaljson) SetLastUpdaterEmailNil() {
	o.LastUpdaterEmail.Set(nil)
}

// UnsetLastUpdaterEmail ensures that no value is present for LastUpdaterEmail, not even an explicit nil
func (o *SegmentListModelHaljson) UnsetLastUpdaterEmail() {
	o.LastUpdaterEmail.Unset()
}

// GetLastUpdaterFullName returns the LastUpdaterFullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentListModelHaljson) GetLastUpdaterFullName() string {
	if o == nil || IsNil(o.LastUpdaterFullName.Get()) {
		var ret string
		return ret
	}
	return *o.LastUpdaterFullName.Get()
}

// GetLastUpdaterFullNameOk returns a tuple with the LastUpdaterFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentListModelHaljson) GetLastUpdaterFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdaterFullName.Get(), o.LastUpdaterFullName.IsSet()
}

// HasLastUpdaterFullName returns a boolean if a field has been set.
func (o *SegmentListModelHaljson) HasLastUpdaterFullName() bool {
	if o != nil && o.LastUpdaterFullName.IsSet() {
		return true
	}

	return false
}

// SetLastUpdaterFullName gets a reference to the given NullableString and assigns it to the LastUpdaterFullName field.
func (o *SegmentListModelHaljson) SetLastUpdaterFullName(v string) {
	o.LastUpdaterFullName.Set(&v)
}
// SetLastUpdaterFullNameNil sets the value for LastUpdaterFullName to be an explicit nil
func (o *SegmentListModelHaljson) SetLastUpdaterFullNameNil() {
	o.LastUpdaterFullName.Set(nil)
}

// UnsetLastUpdaterFullName ensures that no value is present for LastUpdaterFullName, not even an explicit nil
func (o *SegmentListModelHaljson) UnsetLastUpdaterFullName() {
	o.LastUpdaterFullName.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SegmentListModelHaljson) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentListModelHaljson) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SegmentListModelHaljson) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SegmentListModelHaljson) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *SegmentListModelHaljson) GetUsage() int32 {
	if o == nil || IsNil(o.Usage) {
		var ret int32
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentListModelHaljson) GetUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *SegmentListModelHaljson) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given int32 and assigns it to the Usage field.
func (o *SegmentListModelHaljson) SetUsage(v int32) {
	o.Usage = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SegmentListModelHaljson) GetLinks() ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks {
	if o == nil || IsNil(o.Links) {
		var ret ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentListModelHaljson) GetLinksOk() (*ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SegmentListModelHaljson) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks and assigns it to the Links field.
func (o *SegmentListModelHaljson) SetLinks(v ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks) {
	o.Links = &v
}

func (o SegmentListModelHaljson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentListModelHaljson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
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
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableSegmentListModelHaljson struct {
	value *SegmentListModelHaljson
	isSet bool
}

func (v NullableSegmentListModelHaljson) Get() *SegmentListModelHaljson {
	return v.value
}

func (v *NullableSegmentListModelHaljson) Set(val *SegmentListModelHaljson) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentListModelHaljson) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentListModelHaljson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentListModelHaljson(val *SegmentListModelHaljson) *NullableSegmentListModelHaljson {
	return &NullableSegmentListModelHaljson{value: val, isSet: true}
}

func (v NullableSegmentListModelHaljson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentListModelHaljson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


