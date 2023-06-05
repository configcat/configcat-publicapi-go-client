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
	"fmt"
)

// IntegrationLinkType the model 'IntegrationLinkType'
type IntegrationLinkType string

// List of IntegrationLinkType
const (
	INTEGRATIONLINKTYPE_TRELLO IntegrationLinkType = "trello"
	INTEGRATIONLINKTYPE_JIRA IntegrationLinkType = "jira"
	INTEGRATIONLINKTYPE_MONDAY IntegrationLinkType = "monday"
)

// All allowed values of IntegrationLinkType enum
var AllowedIntegrationLinkTypeEnumValues = []IntegrationLinkType{
	"trello",
	"jira",
	"monday",
}

func (v *IntegrationLinkType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IntegrationLinkType(value)
	for _, existing := range AllowedIntegrationLinkTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IntegrationLinkType", value)
}

// NewIntegrationLinkTypeFromValue returns a pointer to a valid IntegrationLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIntegrationLinkTypeFromValue(v string) (*IntegrationLinkType, error) {
	ev := IntegrationLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IntegrationLinkType: valid values are %v", v, AllowedIntegrationLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IntegrationLinkType) IsValid() bool {
	for _, existing := range AllowedIntegrationLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationLinkType value
func (v IntegrationLinkType) Ptr() *IntegrationLinkType {
	return &v
}

type NullableIntegrationLinkType struct {
	value *IntegrationLinkType
	isSet bool
}

func (v NullableIntegrationLinkType) Get() *IntegrationLinkType {
	return v.value
}

func (v *NullableIntegrationLinkType) Set(val *IntegrationLinkType) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationLinkType) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationLinkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationLinkType(val *IntegrationLinkType) *NullableIntegrationLinkType {
	return &NullableIntegrationLinkType{value: val, isSet: true}
}

func (v NullableIntegrationLinkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationLinkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

