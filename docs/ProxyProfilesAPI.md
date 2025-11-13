# \ProxyProfilesAPI

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProxyProfile**](ProxyProfilesAPI.md#CreateProxyProfile) | **Post** /v1/organizations/{organizationId}/proxy-profiles | Create Proxy Profile
[**DeleteProxyProfile**](ProxyProfilesAPI.md#DeleteProxyProfile) | **Delete** /v1/proxy-profiles/{proxyProfileId} | Delete Proxy Profile
[**DeselectProxyProfileSdkKeys**](ProxyProfilesAPI.md#DeselectProxyProfileSdkKeys) | **Post** /v1/proxy-profiles/{proxyProfileId}/sdk-keys/deselect | Deselect SDK keys
[**GenerateProxyProfileSecret**](ProxyProfilesAPI.md#GenerateProxyProfileSecret) | **Post** /v1/proxy-profiles/{proxyProfileId}/secret | Generate Secret
[**GetProxyProfile**](ProxyProfilesAPI.md#GetProxyProfile) | **Get** /v1/proxy-profiles/{proxyProfileId} | Get Proxy Profile
[**GetProxyProfileSdkKeys**](ProxyProfilesAPI.md#GetProxyProfileSdkKeys) | **Get** /v1/proxy-profiles/{proxyProfileId}/sdk-keys | Get selected SDK keys
[**GetProxyProfiles**](ProxyProfilesAPI.md#GetProxyProfiles) | **Get** /v1/organizations/{organizationId}/proxy-profiles | List Proxy Profiles
[**ReplaceProxyProfile**](ProxyProfilesAPI.md#ReplaceProxyProfile) | **Put** /v1/proxy-profiles/{proxyProfileId} | Replace Proxy Profile
[**SelectProxyProfileSdkKeys**](ProxyProfilesAPI.md#SelectProxyProfileSdkKeys) | **Post** /v1/proxy-profiles/{proxyProfileId}/sdk-keys/select | Select SDK keys
[**UpdateProxyProfile**](ProxyProfilesAPI.md#UpdateProxyProfile) | **Patch** /v1/proxy-profiles/{proxyProfileId} | Update Proxy Profile



## CreateProxyProfile

> ProxyProfileModel CreateProxyProfile(ctx, organizationId).CreateOrUpdateProxyProfileRequest(createOrUpdateProxyProfileRequest).Execute()

Create Proxy Profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Organization.
	createOrUpdateProxyProfileRequest := *openapiclient.NewCreateOrUpdateProxyProfileRequest("Name_example") // CreateOrUpdateProxyProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyProfilesAPI.CreateProxyProfile(context.Background(), organizationId).CreateOrUpdateProxyProfileRequest(createOrUpdateProxyProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyProfilesAPI.CreateProxyProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProxyProfile`: ProxyProfileModel
	fmt.Fprintf(os.Stdout, "Response from `ProxyProfilesAPI.CreateProxyProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The identifier of the Organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProxyProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrUpdateProxyProfileRequest** | [**CreateOrUpdateProxyProfileRequest**](CreateOrUpdateProxyProfileRequest.md) |  | 

### Return type

[**ProxyProfileModel**](ProxyProfileModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProxyProfile

> DeleteProxyProfile(ctx, proxyProfileId).Execute()

Delete Proxy Profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	proxyProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Proxy Profile.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProxyProfilesAPI.DeleteProxyProfile(context.Background(), proxyProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyProfilesAPI.DeleteProxyProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proxyProfileId** | **string** | The identifier of the Proxy Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProxyProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeselectProxyProfileSdkKeys

> ProxyProfileSdkKeysListModel DeselectProxyProfileSdkKeys(ctx, proxyProfileId).ProxyProfileSdkKeysRequest(proxyProfileSdkKeysRequest).Execute()

Deselect SDK keys



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	proxyProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Proxy Profile.
	proxyProfileSdkKeysRequest := *openapiclient.NewProxyProfileSdkKeysRequest([]openapiclient.ProxyProfileSdkKeyRequestItem{*openapiclient.NewProxyProfileSdkKeyRequestItem()}) // ProxyProfileSdkKeysRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyProfilesAPI.DeselectProxyProfileSdkKeys(context.Background(), proxyProfileId).ProxyProfileSdkKeysRequest(proxyProfileSdkKeysRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyProfilesAPI.DeselectProxyProfileSdkKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeselectProxyProfileSdkKeys`: ProxyProfileSdkKeysListModel
	fmt.Fprintf(os.Stdout, "Response from `ProxyProfilesAPI.DeselectProxyProfileSdkKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proxyProfileId** | **string** | The identifier of the Proxy Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeselectProxyProfileSdkKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **proxyProfileSdkKeysRequest** | [**ProxyProfileSdkKeysRequest**](ProxyProfileSdkKeysRequest.md) |  | 

### Return type

[**ProxyProfileSdkKeysListModel**](ProxyProfileSdkKeysListModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateProxyProfileSecret

> ProxyProfileSecretModel GenerateProxyProfileSecret(ctx, proxyProfileId).Execute()

Generate Secret



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	proxyProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Proxy Profile.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyProfilesAPI.GenerateProxyProfileSecret(context.Background(), proxyProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyProfilesAPI.GenerateProxyProfileSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateProxyProfileSecret`: ProxyProfileSecretModel
	fmt.Fprintf(os.Stdout, "Response from `ProxyProfilesAPI.GenerateProxyProfileSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proxyProfileId** | **string** | The identifier of the Proxy Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateProxyProfileSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProxyProfileSecretModel**](ProxyProfileSecretModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProxyProfile

> ProxyProfileModel GetProxyProfile(ctx, proxyProfileId).Execute()

Get Proxy Profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	proxyProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Proxy Profile.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyProfilesAPI.GetProxyProfile(context.Background(), proxyProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyProfilesAPI.GetProxyProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProxyProfile`: ProxyProfileModel
	fmt.Fprintf(os.Stdout, "Response from `ProxyProfilesAPI.GetProxyProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proxyProfileId** | **string** | The identifier of the Proxy Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProxyProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProxyProfileModel**](ProxyProfileModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProxyProfileSdkKeys

> ProxyProfileSdkKeysListModel GetProxyProfileSdkKeys(ctx, proxyProfileId).Execute()

Get selected SDK keys



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	proxyProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Proxy Profile.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyProfilesAPI.GetProxyProfileSdkKeys(context.Background(), proxyProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyProfilesAPI.GetProxyProfileSdkKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProxyProfileSdkKeys`: ProxyProfileSdkKeysListModel
	fmt.Fprintf(os.Stdout, "Response from `ProxyProfilesAPI.GetProxyProfileSdkKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proxyProfileId** | **string** | The identifier of the Proxy Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProxyProfileSdkKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProxyProfileSdkKeysListModel**](ProxyProfileSdkKeysListModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProxyProfiles

> ProxyProfileListModel GetProxyProfiles(ctx, organizationId).Execute()

List Proxy Profiles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Organization.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyProfilesAPI.GetProxyProfiles(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyProfilesAPI.GetProxyProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProxyProfiles`: ProxyProfileListModel
	fmt.Fprintf(os.Stdout, "Response from `ProxyProfilesAPI.GetProxyProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The identifier of the Organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProxyProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProxyProfileListModel**](ProxyProfileListModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceProxyProfile

> ProxyProfileModel ReplaceProxyProfile(ctx, proxyProfileId).CreateOrUpdateProxyProfileRequest(createOrUpdateProxyProfileRequest).Execute()

Replace Proxy Profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	proxyProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Proxy Profile.
	createOrUpdateProxyProfileRequest := *openapiclient.NewCreateOrUpdateProxyProfileRequest("Name_example") // CreateOrUpdateProxyProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyProfilesAPI.ReplaceProxyProfile(context.Background(), proxyProfileId).CreateOrUpdateProxyProfileRequest(createOrUpdateProxyProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyProfilesAPI.ReplaceProxyProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceProxyProfile`: ProxyProfileModel
	fmt.Fprintf(os.Stdout, "Response from `ProxyProfilesAPI.ReplaceProxyProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proxyProfileId** | **string** | The identifier of the Proxy Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceProxyProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrUpdateProxyProfileRequest** | [**CreateOrUpdateProxyProfileRequest**](CreateOrUpdateProxyProfileRequest.md) |  | 

### Return type

[**ProxyProfileModel**](ProxyProfileModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectProxyProfileSdkKeys

> ProxyProfileSdkKeysListModel SelectProxyProfileSdkKeys(ctx, proxyProfileId).ProxyProfileSdkKeysRequest(proxyProfileSdkKeysRequest).Execute()

Select SDK keys



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	proxyProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Proxy Profile.
	proxyProfileSdkKeysRequest := *openapiclient.NewProxyProfileSdkKeysRequest([]openapiclient.ProxyProfileSdkKeyRequestItem{*openapiclient.NewProxyProfileSdkKeyRequestItem()}) // ProxyProfileSdkKeysRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyProfilesAPI.SelectProxyProfileSdkKeys(context.Background(), proxyProfileId).ProxyProfileSdkKeysRequest(proxyProfileSdkKeysRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyProfilesAPI.SelectProxyProfileSdkKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectProxyProfileSdkKeys`: ProxyProfileSdkKeysListModel
	fmt.Fprintf(os.Stdout, "Response from `ProxyProfilesAPI.SelectProxyProfileSdkKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proxyProfileId** | **string** | The identifier of the Proxy Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSelectProxyProfileSdkKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **proxyProfileSdkKeysRequest** | [**ProxyProfileSdkKeysRequest**](ProxyProfileSdkKeysRequest.md) |  | 

### Return type

[**ProxyProfileSdkKeysListModel**](ProxyProfileSdkKeysListModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProxyProfile

> ProxyProfileModel UpdateProxyProfile(ctx, proxyProfileId).JsonPatchOperation(jsonPatchOperation).Execute()

Update Proxy Profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	proxyProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Proxy Profile.
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation(openapiclient.OperationType("unknown"), "Path_example")} // []JsonPatchOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyProfilesAPI.UpdateProxyProfile(context.Background(), proxyProfileId).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyProfilesAPI.UpdateProxyProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProxyProfile`: ProxyProfileModel
	fmt.Fprintf(os.Stdout, "Response from `ProxyProfilesAPI.UpdateProxyProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proxyProfileId** | **string** | The identifier of the Proxy Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProxyProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 

### Return type

[**ProxyProfileModel**](ProxyProfileModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

