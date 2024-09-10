/*
ConfigCat Public Management API

The purpose of this API is to access the ConfigCat platform programmatically. You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.  **Base API URL**: https://api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://api.configcat.com/swagger).  The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  format.   **Important:** Do not use this API for accessing and evaluating feature flag values. Use the [SDKs](https://configcat.com/docs/sdk-reference/overview) or the [ConfigCat Proxy](https://configcat.com/docs/advanced/proxy/proxy-overview/) instead.  # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
	"fmt"
)

// EvaluationVersion Determines the evaluation version of a Config.  Using `v2` enables the new features of Config V2 (https://configcat.com/docs/advanced/config-v2).
type EvaluationVersion string

// List of EvaluationVersion
const (
	EVALUATIONVERSION_V1 EvaluationVersion = "v1"
	EVALUATIONVERSION_V2 EvaluationVersion = "v2"
)

// All allowed values of EvaluationVersion enum
var AllowedEvaluationVersionEnumValues = []EvaluationVersion{
	"v1",
	"v2",
}

func (v *EvaluationVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EvaluationVersion(value)
	for _, existing := range AllowedEvaluationVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EvaluationVersion", value)
}

// NewEvaluationVersionFromValue returns a pointer to a valid EvaluationVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEvaluationVersionFromValue(v string) (*EvaluationVersion, error) {
	ev := EvaluationVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EvaluationVersion: valid values are %v", v, AllowedEvaluationVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EvaluationVersion) IsValid() bool {
	for _, existing := range AllowedEvaluationVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EvaluationVersion value
func (v EvaluationVersion) Ptr() *EvaluationVersion {
	return &v
}

type NullableEvaluationVersion struct {
	value *EvaluationVersion
	isSet bool
}

func (v NullableEvaluationVersion) Get() *EvaluationVersion {
	return v.value
}

func (v *NullableEvaluationVersion) Set(val *EvaluationVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableEvaluationVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableEvaluationVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvaluationVersion(val *EvaluationVersion) *NullableEvaluationVersion {
	return &NullableEvaluationVersion{value: val, isSet: true}
}

func (v NullableEvaluationVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvaluationVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

