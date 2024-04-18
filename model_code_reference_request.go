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
)

// checks if the CodeReferenceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CodeReferenceRequest{}

// CodeReferenceRequest struct for CodeReferenceRequest
type CodeReferenceRequest struct {
	// The Config's identifier the scanning was performed against.
	ConfigId string `json:"configId"`
	// The source control repository that contains the scanned code. (Source of the repository selector on the ConfigCat Dashboard)
	Repository string `json:"repository"`
	// The source control branch on where the scan was performed. (Source of the branch selector on the ConfigCat Dashboard)
	Branch string `json:"branch"`
	// The related commit's URL. (Appears on the ConfigCat Dashboard)
	CommitUrl NullableString `json:"commitUrl,omitempty"`
	// The related commit's hash. (Appears on the ConfigCat Dashboard)
	CommitHash NullableString `json:"commitHash,omitempty"`
	// The scanning tool's name. (Appears on the ConfigCat Dashboard)
	Uploader NullableString `json:"uploader,omitempty"`
	// The currently active branches of the repository. Each previously uploaded report that belongs to a non-reported active branch is being deleted.
	ActiveBranches []string `json:"activeBranches,omitempty"`
	// The actual code reference collection.
	FlagReferences []FlagReference `json:"flagReferences,omitempty"`
}

// NewCodeReferenceRequest instantiates a new CodeReferenceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeReferenceRequest(configId string, repository string, branch string) *CodeReferenceRequest {
	this := CodeReferenceRequest{}
	this.ConfigId = configId
	this.Repository = repository
	this.Branch = branch
	return &this
}

// NewCodeReferenceRequestWithDefaults instantiates a new CodeReferenceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeReferenceRequestWithDefaults() *CodeReferenceRequest {
	this := CodeReferenceRequest{}
	return &this
}

// GetConfigId returns the ConfigId field value
func (o *CodeReferenceRequest) GetConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value
// and a boolean to check if the value has been set.
func (o *CodeReferenceRequest) GetConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigId, true
}

// SetConfigId sets field value
func (o *CodeReferenceRequest) SetConfigId(v string) {
	o.ConfigId = v
}

// GetRepository returns the Repository field value
func (o *CodeReferenceRequest) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *CodeReferenceRequest) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *CodeReferenceRequest) SetRepository(v string) {
	o.Repository = v
}

// GetBranch returns the Branch field value
func (o *CodeReferenceRequest) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *CodeReferenceRequest) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *CodeReferenceRequest) SetBranch(v string) {
	o.Branch = v
}

// GetCommitUrl returns the CommitUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeReferenceRequest) GetCommitUrl() string {
	if o == nil || IsNil(o.CommitUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CommitUrl.Get()
}

// GetCommitUrlOk returns a tuple with the CommitUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeReferenceRequest) GetCommitUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitUrl.Get(), o.CommitUrl.IsSet()
}

// HasCommitUrl returns a boolean if a field has been set.
func (o *CodeReferenceRequest) HasCommitUrl() bool {
	if o != nil && o.CommitUrl.IsSet() {
		return true
	}

	return false
}

// SetCommitUrl gets a reference to the given NullableString and assigns it to the CommitUrl field.
func (o *CodeReferenceRequest) SetCommitUrl(v string) {
	o.CommitUrl.Set(&v)
}
// SetCommitUrlNil sets the value for CommitUrl to be an explicit nil
func (o *CodeReferenceRequest) SetCommitUrlNil() {
	o.CommitUrl.Set(nil)
}

// UnsetCommitUrl ensures that no value is present for CommitUrl, not even an explicit nil
func (o *CodeReferenceRequest) UnsetCommitUrl() {
	o.CommitUrl.Unset()
}

// GetCommitHash returns the CommitHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeReferenceRequest) GetCommitHash() string {
	if o == nil || IsNil(o.CommitHash.Get()) {
		var ret string
		return ret
	}
	return *o.CommitHash.Get()
}

// GetCommitHashOk returns a tuple with the CommitHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeReferenceRequest) GetCommitHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitHash.Get(), o.CommitHash.IsSet()
}

// HasCommitHash returns a boolean if a field has been set.
func (o *CodeReferenceRequest) HasCommitHash() bool {
	if o != nil && o.CommitHash.IsSet() {
		return true
	}

	return false
}

// SetCommitHash gets a reference to the given NullableString and assigns it to the CommitHash field.
func (o *CodeReferenceRequest) SetCommitHash(v string) {
	o.CommitHash.Set(&v)
}
// SetCommitHashNil sets the value for CommitHash to be an explicit nil
func (o *CodeReferenceRequest) SetCommitHashNil() {
	o.CommitHash.Set(nil)
}

// UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
func (o *CodeReferenceRequest) UnsetCommitHash() {
	o.CommitHash.Unset()
}

// GetUploader returns the Uploader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeReferenceRequest) GetUploader() string {
	if o == nil || IsNil(o.Uploader.Get()) {
		var ret string
		return ret
	}
	return *o.Uploader.Get()
}

// GetUploaderOk returns a tuple with the Uploader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeReferenceRequest) GetUploaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uploader.Get(), o.Uploader.IsSet()
}

// HasUploader returns a boolean if a field has been set.
func (o *CodeReferenceRequest) HasUploader() bool {
	if o != nil && o.Uploader.IsSet() {
		return true
	}

	return false
}

// SetUploader gets a reference to the given NullableString and assigns it to the Uploader field.
func (o *CodeReferenceRequest) SetUploader(v string) {
	o.Uploader.Set(&v)
}
// SetUploaderNil sets the value for Uploader to be an explicit nil
func (o *CodeReferenceRequest) SetUploaderNil() {
	o.Uploader.Set(nil)
}

// UnsetUploader ensures that no value is present for Uploader, not even an explicit nil
func (o *CodeReferenceRequest) UnsetUploader() {
	o.Uploader.Unset()
}

// GetActiveBranches returns the ActiveBranches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeReferenceRequest) GetActiveBranches() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ActiveBranches
}

// GetActiveBranchesOk returns a tuple with the ActiveBranches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeReferenceRequest) GetActiveBranchesOk() ([]string, bool) {
	if o == nil || IsNil(o.ActiveBranches) {
		return nil, false
	}
	return o.ActiveBranches, true
}

// HasActiveBranches returns a boolean if a field has been set.
func (o *CodeReferenceRequest) HasActiveBranches() bool {
	if o != nil && IsNil(o.ActiveBranches) {
		return true
	}

	return false
}

// SetActiveBranches gets a reference to the given []string and assigns it to the ActiveBranches field.
func (o *CodeReferenceRequest) SetActiveBranches(v []string) {
	o.ActiveBranches = v
}

// GetFlagReferences returns the FlagReferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeReferenceRequest) GetFlagReferences() []FlagReference {
	if o == nil {
		var ret []FlagReference
		return ret
	}
	return o.FlagReferences
}

// GetFlagReferencesOk returns a tuple with the FlagReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeReferenceRequest) GetFlagReferencesOk() ([]FlagReference, bool) {
	if o == nil || IsNil(o.FlagReferences) {
		return nil, false
	}
	return o.FlagReferences, true
}

// HasFlagReferences returns a boolean if a field has been set.
func (o *CodeReferenceRequest) HasFlagReferences() bool {
	if o != nil && IsNil(o.FlagReferences) {
		return true
	}

	return false
}

// SetFlagReferences gets a reference to the given []FlagReference and assigns it to the FlagReferences field.
func (o *CodeReferenceRequest) SetFlagReferences(v []FlagReference) {
	o.FlagReferences = v
}

func (o CodeReferenceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodeReferenceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configId"] = o.ConfigId
	toSerialize["repository"] = o.Repository
	toSerialize["branch"] = o.Branch
	if o.CommitUrl.IsSet() {
		toSerialize["commitUrl"] = o.CommitUrl.Get()
	}
	if o.CommitHash.IsSet() {
		toSerialize["commitHash"] = o.CommitHash.Get()
	}
	if o.Uploader.IsSet() {
		toSerialize["uploader"] = o.Uploader.Get()
	}
	if o.ActiveBranches != nil {
		toSerialize["activeBranches"] = o.ActiveBranches
	}
	if o.FlagReferences != nil {
		toSerialize["flagReferences"] = o.FlagReferences
	}
	return toSerialize, nil
}

type NullableCodeReferenceRequest struct {
	value *CodeReferenceRequest
	isSet bool
}

func (v NullableCodeReferenceRequest) Get() *CodeReferenceRequest {
	return v.value
}

func (v *NullableCodeReferenceRequest) Set(val *CodeReferenceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeReferenceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeReferenceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeReferenceRequest(val *CodeReferenceRequest) *NullableCodeReferenceRequest {
	return &NullableCodeReferenceRequest{value: val, isSet: true}
}

func (v NullableCodeReferenceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeReferenceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


