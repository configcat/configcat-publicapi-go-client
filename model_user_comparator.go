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
	"fmt"
)

// UserComparator The comparison operator the evaluation process must use when it compares the given user attribute's value with the comparison value.
type UserComparator string

// List of UserComparator
const (
	USERCOMPARATOR_CONTAINS_ANY_OF UserComparator = "containsAnyOf"
	USERCOMPARATOR_DOES_NOT_CONTAIN_ANY_OF UserComparator = "doesNotContainAnyOf"
	USERCOMPARATOR_SEM_VER_IS_ONE_OF UserComparator = "semVerIsOneOf"
	USERCOMPARATOR_SEM_VER_IS_NOT_ONE_OF UserComparator = "semVerIsNotOneOf"
	USERCOMPARATOR_SEM_VER_LESS UserComparator = "semVerLess"
	USERCOMPARATOR_SEM_VER_LESS_OR_EQUALS UserComparator = "semVerLessOrEquals"
	USERCOMPARATOR_SEM_VER_GREATER UserComparator = "semVerGreater"
	USERCOMPARATOR_SEM_VER_GREATER_OR_EQUALS UserComparator = "semVerGreaterOrEquals"
	USERCOMPARATOR_NUMBER_EQUALS UserComparator = "numberEquals"
	USERCOMPARATOR_NUMBER_DOES_NOT_EQUAL UserComparator = "numberDoesNotEqual"
	USERCOMPARATOR_NUMBER_LESS UserComparator = "numberLess"
	USERCOMPARATOR_NUMBER_LESS_OR_EQUALS UserComparator = "numberLessOrEquals"
	USERCOMPARATOR_NUMBER_GREATER UserComparator = "numberGreater"
	USERCOMPARATOR_NUMBER_GREATER_OR_EQUALS UserComparator = "numberGreaterOrEquals"
	USERCOMPARATOR_SENSITIVE_IS_ONE_OF UserComparator = "sensitiveIsOneOf"
	USERCOMPARATOR_SENSITIVE_IS_NOT_ONE_OF UserComparator = "sensitiveIsNotOneOf"
	USERCOMPARATOR_DATE_TIME_BEFORE UserComparator = "dateTimeBefore"
	USERCOMPARATOR_DATE_TIME_AFTER UserComparator = "dateTimeAfter"
	USERCOMPARATOR_SENSITIVE_TEXT_EQUALS UserComparator = "sensitiveTextEquals"
	USERCOMPARATOR_SENSITIVE_TEXT_DOES_NOT_EQUAL UserComparator = "sensitiveTextDoesNotEqual"
	USERCOMPARATOR_SENSITIVE_TEXT_STARTS_WITH_ANY_OF UserComparator = "sensitiveTextStartsWithAnyOf"
	USERCOMPARATOR_SENSITIVE_TEXT_NOT_STARTS_WITH_ANY_OF UserComparator = "sensitiveTextNotStartsWithAnyOf"
	USERCOMPARATOR_SENSITIVE_TEXT_ENDS_WITH_ANY_OF UserComparator = "sensitiveTextEndsWithAnyOf"
	USERCOMPARATOR_SENSITIVE_TEXT_NOT_ENDS_WITH_ANY_OF UserComparator = "sensitiveTextNotEndsWithAnyOf"
	USERCOMPARATOR_ARRAY_CONTAINS UserComparator = "arrayContains"
	USERCOMPARATOR_ARRAY_DOES_NOT_CONTAIN UserComparator = "arrayDoesNotContain"
)

// All allowed values of UserComparator enum
var AllowedUserComparatorEnumValues = []UserComparator{
	"containsAnyOf",
	"doesNotContainAnyOf",
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
	"dateTimeBefore",
	"dateTimeAfter",
	"sensitiveTextEquals",
	"sensitiveTextDoesNotEqual",
	"sensitiveTextStartsWithAnyOf",
	"sensitiveTextNotStartsWithAnyOf",
	"sensitiveTextEndsWithAnyOf",
	"sensitiveTextNotEndsWithAnyOf",
	"arrayContains",
	"arrayDoesNotContain",
}

func (v *UserComparator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserComparator(value)
	for _, existing := range AllowedUserComparatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserComparator", value)
}

// NewUserComparatorFromValue returns a pointer to a valid UserComparator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserComparatorFromValue(v string) (*UserComparator, error) {
	ev := UserComparator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserComparator: valid values are %v", v, AllowedUserComparatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserComparator) IsValid() bool {
	for _, existing := range AllowedUserComparatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserComparator value
func (v UserComparator) Ptr() *UserComparator {
	return &v
}

type NullableUserComparator struct {
	value *UserComparator
	isSet bool
}

func (v NullableUserComparator) Get() *UserComparator {
	return v.value
}

func (v *NullableUserComparator) Set(val *UserComparator) {
	v.value = val
	v.isSet = true
}

func (v NullableUserComparator) IsSet() bool {
	return v.isSet
}

func (v *NullableUserComparator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserComparator(val *UserComparator) *NullableUserComparator {
	return &NullableUserComparator{value: val, isSet: true}
}

func (v NullableUserComparator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserComparator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

