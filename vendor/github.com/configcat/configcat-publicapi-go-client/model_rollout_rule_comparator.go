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
	"fmt"
)

// RolloutRuleComparator The comparison operator the evaluation process must use when it compares the given user attribute's value with the comparison value.
type RolloutRuleComparator string

// List of RolloutRuleComparator
const (
	ROLLOUTRULECOMPARATOR_IS_ONE_OF RolloutRuleComparator = "isOneOf"
	ROLLOUTRULECOMPARATOR_IS_NOT_ONE_OF RolloutRuleComparator = "isNotOneOf"
	ROLLOUTRULECOMPARATOR_CONTAINS RolloutRuleComparator = "contains"
	ROLLOUTRULECOMPARATOR_DOES_NOT_CONTAIN RolloutRuleComparator = "doesNotContain"
	ROLLOUTRULECOMPARATOR_SEM_VER_IS_ONE_OF RolloutRuleComparator = "semVerIsOneOf"
	ROLLOUTRULECOMPARATOR_SEM_VER_IS_NOT_ONE_OF RolloutRuleComparator = "semVerIsNotOneOf"
	ROLLOUTRULECOMPARATOR_SEM_VER_LESS RolloutRuleComparator = "semVerLess"
	ROLLOUTRULECOMPARATOR_SEM_VER_LESS_OR_EQUALS RolloutRuleComparator = "semVerLessOrEquals"
	ROLLOUTRULECOMPARATOR_SEM_VER_GREATER RolloutRuleComparator = "semVerGreater"
	ROLLOUTRULECOMPARATOR_SEM_VER_GREATER_OR_EQUALS RolloutRuleComparator = "semVerGreaterOrEquals"
	ROLLOUTRULECOMPARATOR_NUMBER_EQUALS RolloutRuleComparator = "numberEquals"
	ROLLOUTRULECOMPARATOR_NUMBER_DOES_NOT_EQUAL RolloutRuleComparator = "numberDoesNotEqual"
	ROLLOUTRULECOMPARATOR_NUMBER_LESS RolloutRuleComparator = "numberLess"
	ROLLOUTRULECOMPARATOR_NUMBER_LESS_OR_EQUALS RolloutRuleComparator = "numberLessOrEquals"
	ROLLOUTRULECOMPARATOR_NUMBER_GREATER RolloutRuleComparator = "numberGreater"
	ROLLOUTRULECOMPARATOR_NUMBER_GREATER_OR_EQUALS RolloutRuleComparator = "numberGreaterOrEquals"
	ROLLOUTRULECOMPARATOR_SENSITIVE_IS_ONE_OF RolloutRuleComparator = "sensitiveIsOneOf"
	ROLLOUTRULECOMPARATOR_SENSITIVE_IS_NOT_ONE_OF RolloutRuleComparator = "sensitiveIsNotOneOf"
)

// All allowed values of RolloutRuleComparator enum
var AllowedRolloutRuleComparatorEnumValues = []RolloutRuleComparator{
	"isOneOf",
	"isNotOneOf",
	"contains",
	"doesNotContain",
	"semVerIsOneOf",
	"semVerIsNotOneOf",
	"semVerLess",
	"semVerLessOrEquals",
	"semVerGreater",
	"semVerGreaterOrEquals",
	"numberEquals",
	"numberDoesNotEqual",
	"numberLess",
	"numberLessOrEquals",
	"numberGreater",
	"numberGreaterOrEquals",
	"sensitiveIsOneOf",
	"sensitiveIsNotOneOf",
}

func (v *RolloutRuleComparator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RolloutRuleComparator(value)
	for _, existing := range AllowedRolloutRuleComparatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RolloutRuleComparator", value)
}

// NewRolloutRuleComparatorFromValue returns a pointer to a valid RolloutRuleComparator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRolloutRuleComparatorFromValue(v string) (*RolloutRuleComparator, error) {
	ev := RolloutRuleComparator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RolloutRuleComparator: valid values are %v", v, AllowedRolloutRuleComparatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RolloutRuleComparator) IsValid() bool {
	for _, existing := range AllowedRolloutRuleComparatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RolloutRuleComparator value
func (v RolloutRuleComparator) Ptr() *RolloutRuleComparator {
	return &v
}

type NullableRolloutRuleComparator struct {
	value *RolloutRuleComparator
	isSet bool
}

func (v NullableRolloutRuleComparator) Get() *RolloutRuleComparator {
	return v.value
}

func (v *NullableRolloutRuleComparator) Set(val *RolloutRuleComparator) {
	v.value = val
	v.isSet = true
}

func (v NullableRolloutRuleComparator) IsSet() bool {
	return v.isSet
}

func (v *NullableRolloutRuleComparator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolloutRuleComparator(val *RolloutRuleComparator) *NullableRolloutRuleComparator {
	return &NullableRolloutRuleComparator{value: val, isSet: true}
}

func (v NullableRolloutRuleComparator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolloutRuleComparator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

