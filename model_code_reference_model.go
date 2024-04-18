/*
ConfigCat Public Management API

The purpose of this API is to access the ConfigCat platform programmatically. You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.  **Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://test-api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://test-api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
	"time"
)

// checks if the CodeReferenceModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CodeReferenceModel{}

// CodeReferenceModel struct for CodeReferenceModel
type CodeReferenceModel struct {
	// The source control branch on where the scan was performed. (Source of the branch selector on the ConfigCat Dashboard)
	Branch NullableString `json:"branch,omitempty"`
	// The actual references to the given Feature Flag or Setting.
	References []ReferenceLines `json:"references,omitempty"`
	// The related commit's URL.
	CommitUrl NullableString `json:"commitUrl,omitempty"`
	// The related commit's hash.
	CommitHash NullableString `json:"commitHash,omitempty"`
	// The date and time when the reference report was uploaded.
	SyncedAt *time.Time `json:"syncedAt,omitempty"`
	// The source control repository that contains the scanned code.
	Repository NullableString `json:"repository,omitempty"`
	// The identifier of the reference report.
	CodeReferenceId *string `json:"codeReferenceId,omitempty"`
	// The code reference scanning tool's name.
	Uploader NullableString `json:"uploader,omitempty"`
}

// NewCodeReferenceModel instantiates a new CodeReferenceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeReferenceModel() *CodeReferenceModel {
	this := CodeReferenceModel{}
	return &this
}

// NewCodeReferenceModelWithDefaults instantiates a new CodeReferenceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeReferenceModelWithDefaults() *CodeReferenceModel {
	this := CodeReferenceModel{}
	return &this
}

// GetBranch returns the Branch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeReferenceModel) GetBranch() string {
	if o == nil || IsNil(o.Branch.Get()) {
		var ret string
		return ret
	}
	return *o.Branch.Get()
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeReferenceModel) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Branch.Get(), o.Branch.IsSet()
}

// HasBranch returns a boolean if a field has been set.
func (o *CodeReferenceModel) HasBranch() bool {
	if o != nil && o.Branch.IsSet() {
		return true
	}

	return false
}

// SetBranch gets a reference to the given NullableString and assigns it to the Branch field.
func (o *CodeReferenceModel) SetBranch(v string) {
	o.Branch.Set(&v)
}
// SetBranchNil sets the value for Branch to be an explicit nil
func (o *CodeReferenceModel) SetBranchNil() {
	o.Branch.Set(nil)
}

// UnsetBranch ensures that no value is present for Branch, not even an explicit nil
func (o *CodeReferenceModel) UnsetBranch() {
	o.Branch.Unset()
}

// GetReferences returns the References field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeReferenceModel) GetReferences() []ReferenceLines {
	if o == nil {
		var ret []ReferenceLines
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeReferenceModel) GetReferencesOk() ([]ReferenceLines, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *CodeReferenceModel) HasReferences() bool {
	if o != nil && IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []ReferenceLines and assigns it to the References field.
func (o *CodeReferenceModel) SetReferences(v []ReferenceLines) {
	o.References = v
}

// GetCommitUrl returns the CommitUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeReferenceModel) GetCommitUrl() string {
	if o == nil || IsNil(o.CommitUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CommitUrl.Get()
}

// GetCommitUrlOk returns a tuple with the CommitUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeReferenceModel) GetCommitUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitUrl.Get(), o.CommitUrl.IsSet()
}

// HasCommitUrl returns a boolean if a field has been set.
func (o *CodeReferenceModel) HasCommitUrl() bool {
	if o != nil && o.CommitUrl.IsSet() {
		return true
	}

	return false
}

// SetCommitUrl gets a reference to the given NullableString and assigns it to the CommitUrl field.
func (o *CodeReferenceModel) SetCommitUrl(v string) {
	o.CommitUrl.Set(&v)
}
// SetCommitUrlNil sets the value for CommitUrl to be an explicit nil
func (o *CodeReferenceModel) SetCommitUrlNil() {
	o.CommitUrl.Set(nil)
}

// UnsetCommitUrl ensures that no value is present for CommitUrl, not even an explicit nil
func (o *CodeReferenceModel) UnsetCommitUrl() {
	o.CommitUrl.Unset()
}

// GetCommitHash returns the CommitHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeReferenceModel) GetCommitHash() string {
	if o == nil || IsNil(o.CommitHash.Get()) {
		var ret string
		return ret
	}
	return *o.CommitHash.Get()
}

// GetCommitHashOk returns a tuple with the CommitHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeReferenceModel) GetCommitHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitHash.Get(), o.CommitHash.IsSet()
}

// HasCommitHash returns a boolean if a field has been set.
func (o *CodeReferenceModel) HasCommitHash() bool {
	if o != nil && o.CommitHash.IsSet() {
		return true
	}

	return false
}

// SetCommitHash gets a reference to the given NullableString and assigns it to the CommitHash field.
func (o *CodeReferenceModel) SetCommitHash(v string) {
	o.CommitHash.Set(&v)
}
// SetCommitHashNil sets the value for CommitHash to be an explicit nil
func (o *CodeReferenceModel) SetCommitHashNil() {
	o.CommitHash.Set(nil)
}

// UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
func (o *CodeReferenceModel) UnsetCommitHash() {
	o.CommitHash.Unset()
}

// GetSyncedAt returns the SyncedAt field value if set, zero value otherwise.
func (o *CodeReferenceModel) GetSyncedAt() time.Time {
	if o == nil || IsNil(o.SyncedAt) {
		var ret time.Time
		return ret
	}
	return *o.SyncedAt
}

// GetSyncedAtOk returns a tuple with the SyncedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeReferenceModel) GetSyncedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SyncedAt) {
		return nil, false
	}
	return o.SyncedAt, true
}

// HasSyncedAt returns a boolean if a field has been set.
func (o *CodeReferenceModel) HasSyncedAt() bool {
	if o != nil && !IsNil(o.SyncedAt) {
		return true
	}

	return false
}

// SetSyncedAt gets a reference to the given time.Time and assigns it to the SyncedAt field.
func (o *CodeReferenceModel) SetSyncedAt(v time.Time) {
	o.SyncedAt = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeReferenceModel) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeReferenceModel) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *CodeReferenceModel) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *CodeReferenceModel) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *CodeReferenceModel) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *CodeReferenceModel) UnsetRepository() {
	o.Repository.Unset()
}

// GetCodeReferenceId returns the CodeReferenceId field value if set, zero value otherwise.
func (o *CodeReferenceModel) GetCodeReferenceId() string {
	if o == nil || IsNil(o.CodeReferenceId) {
		var ret string
		return ret
	}
	return *o.CodeReferenceId
}

// GetCodeReferenceIdOk returns a tuple with the CodeReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeReferenceModel) GetCodeReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.CodeReferenceId) {
		return nil, false
	}
	return o.CodeReferenceId, true
}

// HasCodeReferenceId returns a boolean if a field has been set.
func (o *CodeReferenceModel) HasCodeReferenceId() bool {
	if o != nil && !IsNil(o.CodeReferenceId) {
		return true
	}

	return false
}

// SetCodeReferenceId gets a reference to the given string and assigns it to the CodeReferenceId field.
func (o *CodeReferenceModel) SetCodeReferenceId(v string) {
	o.CodeReferenceId = &v
}

// GetUploader returns the Uploader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeReferenceModel) GetUploader() string {
	if o == nil || IsNil(o.Uploader.Get()) {
		var ret string
		return ret
	}
	return *o.Uploader.Get()
}

// GetUploaderOk returns a tuple with the Uploader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeReferenceModel) GetUploaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uploader.Get(), o.Uploader.IsSet()
}

// HasUploader returns a boolean if a field has been set.
func (o *CodeReferenceModel) HasUploader() bool {
	if o != nil && o.Uploader.IsSet() {
		return true
	}

	return false
}

// SetUploader gets a reference to the given NullableString and assigns it to the Uploader field.
func (o *CodeReferenceModel) SetUploader(v string) {
	o.Uploader.Set(&v)
}
// SetUploaderNil sets the value for Uploader to be an explicit nil
func (o *CodeReferenceModel) SetUploaderNil() {
	o.Uploader.Set(nil)
}

// UnsetUploader ensures that no value is present for Uploader, not even an explicit nil
func (o *CodeReferenceModel) UnsetUploader() {
	o.Uploader.Unset()
}

func (o CodeReferenceModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodeReferenceModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Branch.IsSet() {
		toSerialize["branch"] = o.Branch.Get()
	}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	if o.CommitUrl.IsSet() {
		toSerialize["commitUrl"] = o.CommitUrl.Get()
	}
	if o.CommitHash.IsSet() {
		toSerialize["commitHash"] = o.CommitHash.Get()
	}
	if !IsNil(o.SyncedAt) {
		toSerialize["syncedAt"] = o.SyncedAt
	}
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if !IsNil(o.CodeReferenceId) {
		toSerialize["codeReferenceId"] = o.CodeReferenceId
	}
	if o.Uploader.IsSet() {
		toSerialize["uploader"] = o.Uploader.Get()
	}
	return toSerialize, nil
}

type NullableCodeReferenceModel struct {
	value *CodeReferenceModel
	isSet bool
}

func (v NullableCodeReferenceModel) Get() *CodeReferenceModel {
	return v.value
}

func (v *NullableCodeReferenceModel) Set(val *CodeReferenceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeReferenceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeReferenceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeReferenceModel(val *CodeReferenceModel) *NullableCodeReferenceModel {
	return &NullableCodeReferenceModel{value: val, isSet: true}
}

func (v NullableCodeReferenceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeReferenceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


