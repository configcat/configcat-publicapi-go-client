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

// checks if the ConfigModelHaljsonEmbeddedProductLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigModelHaljsonEmbeddedProductLinks{}

// ConfigModelHaljsonEmbeddedProductLinks struct for ConfigModelHaljsonEmbeddedProductLinks
type ConfigModelHaljsonEmbeddedProductLinks struct {
	Self *string `json:"self,omitempty"`
	Configs *string `json:"configs,omitempty"`
	Environments *string `json:"environments,omitempty"`
	Tags *string `json:"tags,omitempty"`
	PermissionGroups *string `json:"permission-groups,omitempty"`
	Members *string `json:"members,omitempty"`
	Segments *string `json:"segments,omitempty"`
}

// NewConfigModelHaljsonEmbeddedProductLinks instantiates a new ConfigModelHaljsonEmbeddedProductLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigModelHaljsonEmbeddedProductLinks() *ConfigModelHaljsonEmbeddedProductLinks {
	this := ConfigModelHaljsonEmbeddedProductLinks{}
	return &this
}

// NewConfigModelHaljsonEmbeddedProductLinksWithDefaults instantiates a new ConfigModelHaljsonEmbeddedProductLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigModelHaljsonEmbeddedProductLinksWithDefaults() *ConfigModelHaljsonEmbeddedProductLinks {
	this := ConfigModelHaljsonEmbeddedProductLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ConfigModelHaljsonEmbeddedProductLinks) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModelHaljsonEmbeddedProductLinks) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ConfigModelHaljsonEmbeddedProductLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *ConfigModelHaljsonEmbeddedProductLinks) SetSelf(v string) {
	o.Self = &v
}

// GetConfigs returns the Configs field value if set, zero value otherwise.
func (o *ConfigModelHaljsonEmbeddedProductLinks) GetConfigs() string {
	if o == nil || IsNil(o.Configs) {
		var ret string
		return ret
	}
	return *o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModelHaljsonEmbeddedProductLinks) GetConfigsOk() (*string, bool) {
	if o == nil || IsNil(o.Configs) {
		return nil, false
	}
	return o.Configs, true
}

// HasConfigs returns a boolean if a field has been set.
func (o *ConfigModelHaljsonEmbeddedProductLinks) HasConfigs() bool {
	if o != nil && !IsNil(o.Configs) {
		return true
	}

	return false
}

// SetConfigs gets a reference to the given string and assigns it to the Configs field.
func (o *ConfigModelHaljsonEmbeddedProductLinks) SetConfigs(v string) {
	o.Configs = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *ConfigModelHaljsonEmbeddedProductLinks) GetEnvironments() string {
	if o == nil || IsNil(o.Environments) {
		var ret string
		return ret
	}
	return *o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModelHaljsonEmbeddedProductLinks) GetEnvironmentsOk() (*string, bool) {
	if o == nil || IsNil(o.Environments) {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *ConfigModelHaljsonEmbeddedProductLinks) HasEnvironments() bool {
	if o != nil && !IsNil(o.Environments) {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given string and assigns it to the Environments field.
func (o *ConfigModelHaljsonEmbeddedProductLinks) SetEnvironments(v string) {
	o.Environments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConfigModelHaljsonEmbeddedProductLinks) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModelHaljsonEmbeddedProductLinks) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConfigModelHaljsonEmbeddedProductLinks) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *ConfigModelHaljsonEmbeddedProductLinks) SetTags(v string) {
	o.Tags = &v
}

// GetPermissionGroups returns the PermissionGroups field value if set, zero value otherwise.
func (o *ConfigModelHaljsonEmbeddedProductLinks) GetPermissionGroups() string {
	if o == nil || IsNil(o.PermissionGroups) {
		var ret string
		return ret
	}
	return *o.PermissionGroups
}

// GetPermissionGroupsOk returns a tuple with the PermissionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModelHaljsonEmbeddedProductLinks) GetPermissionGroupsOk() (*string, bool) {
	if o == nil || IsNil(o.PermissionGroups) {
		return nil, false
	}
	return o.PermissionGroups, true
}

// HasPermissionGroups returns a boolean if a field has been set.
func (o *ConfigModelHaljsonEmbeddedProductLinks) HasPermissionGroups() bool {
	if o != nil && !IsNil(o.PermissionGroups) {
		return true
	}

	return false
}

// SetPermissionGroups gets a reference to the given string and assigns it to the PermissionGroups field.
func (o *ConfigModelHaljsonEmbeddedProductLinks) SetPermissionGroups(v string) {
	o.PermissionGroups = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ConfigModelHaljsonEmbeddedProductLinks) GetMembers() string {
	if o == nil || IsNil(o.Members) {
		var ret string
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModelHaljsonEmbeddedProductLinks) GetMembersOk() (*string, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ConfigModelHaljsonEmbeddedProductLinks) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given string and assigns it to the Members field.
func (o *ConfigModelHaljsonEmbeddedProductLinks) SetMembers(v string) {
	o.Members = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *ConfigModelHaljsonEmbeddedProductLinks) GetSegments() string {
	if o == nil || IsNil(o.Segments) {
		var ret string
		return ret
	}
	return *o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModelHaljsonEmbeddedProductLinks) GetSegmentsOk() (*string, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *ConfigModelHaljsonEmbeddedProductLinks) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given string and assigns it to the Segments field.
func (o *ConfigModelHaljsonEmbeddedProductLinks) SetSegments(v string) {
	o.Segments = &v
}

func (o ConfigModelHaljsonEmbeddedProductLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigModelHaljsonEmbeddedProductLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Configs) {
		toSerialize["configs"] = o.Configs
	}
	if !IsNil(o.Environments) {
		toSerialize["environments"] = o.Environments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.PermissionGroups) {
		toSerialize["permission-groups"] = o.PermissionGroups
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	return toSerialize, nil
}

type NullableConfigModelHaljsonEmbeddedProductLinks struct {
	value *ConfigModelHaljsonEmbeddedProductLinks
	isSet bool
}

func (v NullableConfigModelHaljsonEmbeddedProductLinks) Get() *ConfigModelHaljsonEmbeddedProductLinks {
	return v.value
}

func (v *NullableConfigModelHaljsonEmbeddedProductLinks) Set(val *ConfigModelHaljsonEmbeddedProductLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigModelHaljsonEmbeddedProductLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigModelHaljsonEmbeddedProductLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigModelHaljsonEmbeddedProductLinks(val *ConfigModelHaljsonEmbeddedProductLinks) *NullableConfigModelHaljsonEmbeddedProductLinks {
	return &NullableConfigModelHaljsonEmbeddedProductLinks{value: val, isSet: true}
}

func (v NullableConfigModelHaljsonEmbeddedProductLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigModelHaljsonEmbeddedProductLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


