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
	"fmt"
)

// OperationType the model 'OperationType'
type OperationType string

// List of OperationType
const (
	OPERATIONTYPE_UNKNOWN OperationType = "unknown"
	OPERATIONTYPE_ADD OperationType = "add"
	OPERATIONTYPE_REMOVE OperationType = "remove"
	OPERATIONTYPE_REPLACE OperationType = "replace"
	OPERATIONTYPE_MOVE OperationType = "move"
	OPERATIONTYPE_COPY OperationType = "copy"
	OPERATIONTYPE_TEST OperationType = "test"
)

// All allowed values of OperationType enum
var AllowedOperationTypeEnumValues = []OperationType{
	"unknown",
	"add",
	"remove",
	"replace",
	"move",
	"copy",
	"test",
}

func (v *OperationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OperationType(value)
	for _, existing := range AllowedOperationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperationType", value)
}

// NewOperationTypeFromValue returns a pointer to a valid OperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperationTypeFromValue(v string) (*OperationType, error) {
	ev := OperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperationType: valid values are %v", v, AllowedOperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperationType) IsValid() bool {
	for _, existing := range AllowedOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OperationType value
func (v OperationType) Ptr() *OperationType {
	return &v
}

type NullableOperationType struct {
	value *OperationType
	isSet bool
}

func (v NullableOperationType) Get() *OperationType {
	return v.value
}

func (v *NullableOperationType) Set(val *OperationType) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationType) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationType(val *OperationType) *NullableOperationType {
	return &NullableOperationType{value: val, isSet: true}
}

func (v NullableOperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

