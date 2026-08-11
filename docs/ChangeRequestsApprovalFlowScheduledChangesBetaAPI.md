# \ChangeRequestsApprovalFlowScheduledChangesBetaAPI

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddChangeRequestComment**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#AddChangeRequestComment) | **Post** /v2/change-requests/{changeRequestId}/comments | Add Comment
[**ApplyChangeRequest**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#ApplyChangeRequest) | **Post** /v2/change-requests/{changeRequestId}/apply | Apply Change Request
[**ApproveChangeRequest**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#ApproveChangeRequest) | **Post** /v2/change-requests/{changeRequestId}/approve | Approve Change Request
[**ClaimChangeRequestOwnership**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#ClaimChangeRequestOwnership) | **Post** /v2/change-requests/{changeRequestId}/claim-ownership | Claim Ownership
[**CloseChangeRequest**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#CloseChangeRequest) | **Post** /v2/change-requests/{changeRequestId}/close | Close Change Request
[**CreateChangeRequest**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#CreateChangeRequest) | **Post** /v2/configs/{configId}/environments/{environmentId}/change-requests | Create Change Request
[**DeleteChangeRequestComment**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#DeleteChangeRequestComment) | **Delete** /v2/change-request-comments/{commentId} | Delete Comment
[**DeleteChangeRequestProposedChange**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#DeleteChangeRequestProposedChange) | **Delete** /v2/change-requests/{changeRequestId}/proposed-changes/{settingId} | Delete Setting from Change Request
[**GetChangeRequest**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#GetChangeRequest) | **Get** /v2/change-requests/{changeRequestId} | Get Change Request
[**GetChangeRequestProposedChanges**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#GetChangeRequestProposedChanges) | **Get** /v2/change-requests/{changeRequestId}/proposed-changes | Get Settings included in Change Request
[**GetChangeRequests**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#GetChangeRequests) | **Get** /v2/products/{productId}/change-requests | List Change Requests
[**RemoveChangeRequestApproval**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#RemoveChangeRequestApproval) | **Post** /v2/change-requests/{changeRequestId}/remove-approval | Remove Approval
[**ResolveChangeRequestSettingConflicts**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#ResolveChangeRequestSettingConflicts) | **Post** /v2/change-requests/{changeRequestId}/proposed-changes/{settingId}/resolve-conflicts | Resolve Setting Conflicts
[**UpdateChangeRequest**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#UpdateChangeRequest) | **Put** /v2/change-requests/{changeRequestId} | Update Change Request
[**UpdateChangeRequestComment**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#UpdateChangeRequestComment) | **Put** /v2/change-request-comments/{commentId} | Update Comment
[**UpdateChangeRequestProposedChanges**](ChangeRequestsApprovalFlowScheduledChangesBetaAPI.md#UpdateChangeRequestProposedChanges) | **Put** /v2/change-requests/{changeRequestId}/proposed-changes | Update Settings included in Change Request



## AddChangeRequestComment

> ChangeRequestCommentModel AddChangeRequestComment(ctx, changeRequestId).AddChangeRequestCommentModel(addChangeRequestCommentModel).Execute()

Add Comment



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
	changeRequestId := int64(789) // int64 | The identifier of the Change Request.
	addChangeRequestCommentModel := *openapiclient.NewAddChangeRequestCommentModel("Body_example") // AddChangeRequestCommentModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.AddChangeRequestComment(context.Background(), changeRequestId).AddChangeRequestCommentModel(addChangeRequestCommentModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.AddChangeRequestComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddChangeRequestComment`: ChangeRequestCommentModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.AddChangeRequestComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int64** | The identifier of the Change Request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddChangeRequestCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addChangeRequestCommentModel** | [**AddChangeRequestCommentModel**](AddChangeRequestCommentModel.md) |  | 

### Return type

[**ChangeRequestCommentModel**](ChangeRequestCommentModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyChangeRequest

> ChangeRequestModel ApplyChangeRequest(ctx, changeRequestId).Execute()

Apply Change Request



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
	changeRequestId := int64(789) // int64 | The identifier of the Change Request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.ApplyChangeRequest(context.Background(), changeRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.ApplyChangeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyChangeRequest`: ChangeRequestModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.ApplyChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int64** | The identifier of the Change Request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChangeRequestModel**](ChangeRequestModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApproveChangeRequest

> ChangeRequestModel ApproveChangeRequest(ctx, changeRequestId).Execute()

Approve Change Request



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
	changeRequestId := int64(789) // int64 | The identifier of the Change Request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.ApproveChangeRequest(context.Background(), changeRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.ApproveChangeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveChangeRequest`: ChangeRequestModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.ApproveChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int64** | The identifier of the Change Request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChangeRequestModel**](ChangeRequestModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimChangeRequestOwnership

> ChangeRequestModel ClaimChangeRequestOwnership(ctx, changeRequestId).Execute()

Claim Ownership



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
	changeRequestId := int64(789) // int64 | The identifier of the Change Request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.ClaimChangeRequestOwnership(context.Background(), changeRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.ClaimChangeRequestOwnership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimChangeRequestOwnership`: ChangeRequestModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.ClaimChangeRequestOwnership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int64** | The identifier of the Change Request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimChangeRequestOwnershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChangeRequestModel**](ChangeRequestModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloseChangeRequest

> ChangeRequestModel CloseChangeRequest(ctx, changeRequestId).Execute()

Close Change Request



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
	changeRequestId := int64(789) // int64 | The identifier of the Change Request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.CloseChangeRequest(context.Background(), changeRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.CloseChangeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloseChangeRequest`: ChangeRequestModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.CloseChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int64** | The identifier of the Change Request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChangeRequestModel**](ChangeRequestModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChangeRequest

> ChangeRequestModel CreateChangeRequest(ctx, configId, environmentId).CreateChangeRequestModel(createChangeRequestModel).Execute()

Create Change Request



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
	configId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Config.
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment.
	createChangeRequestModel := *openapiclient.NewCreateChangeRequestModel("Title_example") // CreateChangeRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.CreateChangeRequest(context.Background(), configId, environmentId).CreateChangeRequestModel(createChangeRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.CreateChangeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChangeRequest`: ChangeRequestModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.CreateChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | The identifier of the Config. | 
**environmentId** | **string** | The identifier of the Environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createChangeRequestModel** | [**CreateChangeRequestModel**](CreateChangeRequestModel.md) |  | 

### Return type

[**ChangeRequestModel**](ChangeRequestModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChangeRequestComment

> DeleteChangeRequestComment(ctx, commentId).Execute()

Delete Comment



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
	commentId := int64(789) // int64 | The identifier of the Change Request comment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.DeleteChangeRequestComment(context.Background(), commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.DeleteChangeRequestComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | The identifier of the Change Request comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChangeRequestCommentRequest struct via the builder pattern


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


## DeleteChangeRequestProposedChange

> ChangeRequestModel DeleteChangeRequestProposedChange(ctx, changeRequestId, settingId).Execute()

Delete Setting from Change Request



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
	changeRequestId := int64(789) // int64 | The identifier of the Change Request.
	settingId := int32(56) // int32 | The identifier of the Setting.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.DeleteChangeRequestProposedChange(context.Background(), changeRequestId, settingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.DeleteChangeRequestProposedChange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteChangeRequestProposedChange`: ChangeRequestModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.DeleteChangeRequestProposedChange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int64** | The identifier of the Change Request. | 
**settingId** | **int32** | The identifier of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChangeRequestProposedChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ChangeRequestModel**](ChangeRequestModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChangeRequest

> ChangeRequestModel GetChangeRequest(ctx, changeRequestId).Execute()

Get Change Request



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
	changeRequestId := int64(789) // int64 | The identifier of the Change Request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.GetChangeRequest(context.Background(), changeRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.GetChangeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChangeRequest`: ChangeRequestModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.GetChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int64** | The identifier of the Change Request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChangeRequestModel**](ChangeRequestModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChangeRequestProposedChanges

> ChangeRequestProposedChangesModel GetChangeRequestProposedChanges(ctx, changeRequestId).SettingId(settingId).Execute()

Get Settings included in Change Request



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
	changeRequestId := int64(789) // int64 | The identifier of the Change Request.
	settingId := int32(56) // int32 | The optional identifier of the Setting. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.GetChangeRequestProposedChanges(context.Background(), changeRequestId).SettingId(settingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.GetChangeRequestProposedChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChangeRequestProposedChanges`: ChangeRequestProposedChangesModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.GetChangeRequestProposedChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int64** | The identifier of the Change Request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChangeRequestProposedChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settingId** | **int32** | The optional identifier of the Setting. | 

### Return type

[**ChangeRequestProposedChangesModel**](ChangeRequestProposedChangesModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChangeRequests

> ChangeRequestsModel GetChangeRequests(ctx, productId).ConfigId(configId).EnvironmentId(environmentId).SettingId(settingId).ChangeRequestStatusFilter(changeRequestStatusFilter).ScheduleFilter(scheduleFilter).ApproveRequiredFilter(approveRequiredFilter).NeedsAttentionFilter(needsAttentionFilter).PageNumber(pageNumber).PageSize(pageSize).Execute()

List Change Requests



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
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Product.
	configId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter Change Requests by Config identifier. (optional)
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter Change Requests by Environment identifier. (optional)
	settingId := int32(56) // int32 | Filter Change Requests by Setting identifier. (optional)
	changeRequestStatusFilter := []openapiclient.ChangeRequestStatus{openapiclient.ChangeRequestStatus("open")} // []ChangeRequestStatus | Filter Change Requests by status values. (optional)
	scheduleFilter := openapiclient.ChangeRequestScheduleFilter("nonScheduled") // ChangeRequestScheduleFilter | Filter Change Requests by schedule state. (optional)
	approveRequiredFilter := openapiclient.ChangeRequestApproveRequiredFilter("approveNotRequired") // ChangeRequestApproveRequiredFilter | Filter Change Requests by approval requirement. (optional)
	needsAttentionFilter := openapiclient.NeedsAttentionFilter("notNeedsAttention") // NeedsAttentionFilter | Filter Change Requests by whether they need attention. (optional)
	pageNumber := int32(56) // int32 | Page number (min: 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Page size (min: 1, max: 100). (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.GetChangeRequests(context.Background(), productId).ConfigId(configId).EnvironmentId(environmentId).SettingId(settingId).ChangeRequestStatusFilter(changeRequestStatusFilter).ScheduleFilter(scheduleFilter).ApproveRequiredFilter(approveRequiredFilter).NeedsAttentionFilter(needsAttentionFilter).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.GetChangeRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChangeRequests`: ChangeRequestsModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.GetChangeRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | The identifier of the Product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChangeRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configId** | **string** | Filter Change Requests by Config identifier. | 
 **environmentId** | **string** | Filter Change Requests by Environment identifier. | 
 **settingId** | **int32** | Filter Change Requests by Setting identifier. | 
 **changeRequestStatusFilter** | [**[]ChangeRequestStatus**](ChangeRequestStatus.md) | Filter Change Requests by status values. | 
 **scheduleFilter** | [**ChangeRequestScheduleFilter**](ChangeRequestScheduleFilter.md) | Filter Change Requests by schedule state. | 
 **approveRequiredFilter** | [**ChangeRequestApproveRequiredFilter**](ChangeRequestApproveRequiredFilter.md) | Filter Change Requests by approval requirement. | 
 **needsAttentionFilter** | [**NeedsAttentionFilter**](NeedsAttentionFilter.md) | Filter Change Requests by whether they need attention. | 
 **pageNumber** | **int32** | Page number (min: 1). | [default to 1]
 **pageSize** | **int32** | Page size (min: 1, max: 100). | [default to 25]

### Return type

[**ChangeRequestsModel**](ChangeRequestsModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveChangeRequestApproval

> ChangeRequestModel RemoveChangeRequestApproval(ctx, changeRequestId).Execute()

Remove Approval



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
	changeRequestId := int64(789) // int64 | The identifier of the Change Request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.RemoveChangeRequestApproval(context.Background(), changeRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.RemoveChangeRequestApproval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveChangeRequestApproval`: ChangeRequestModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.RemoveChangeRequestApproval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int64** | The identifier of the Change Request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveChangeRequestApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChangeRequestModel**](ChangeRequestModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveChangeRequestSettingConflicts

> ChangeRequestModel ResolveChangeRequestSettingConflicts(ctx, changeRequestId, settingId).ResolveChangeRequestSettingConflictsModel(resolveChangeRequestSettingConflictsModel).Execute()

Resolve Setting Conflicts



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
	changeRequestId := int64(789) // int64 | The identifier of the Change Request.
	settingId := int32(56) // int32 | The identifier of the Setting.
	resolveChangeRequestSettingConflictsModel := *openapiclient.NewResolveChangeRequestSettingConflictsModel(*openapiclient.NewUpdateEvaluationFormulaWithLatestVersionModel(*openapiclient.NewUpdateValueModel())) // ResolveChangeRequestSettingConflictsModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.ResolveChangeRequestSettingConflicts(context.Background(), changeRequestId, settingId).ResolveChangeRequestSettingConflictsModel(resolveChangeRequestSettingConflictsModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.ResolveChangeRequestSettingConflicts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveChangeRequestSettingConflicts`: ChangeRequestModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.ResolveChangeRequestSettingConflicts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int64** | The identifier of the Change Request. | 
**settingId** | **int32** | The identifier of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveChangeRequestSettingConflictsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resolveChangeRequestSettingConflictsModel** | [**ResolveChangeRequestSettingConflictsModel**](ResolveChangeRequestSettingConflictsModel.md) |  | 

### Return type

[**ChangeRequestModel**](ChangeRequestModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChangeRequest

> ChangeRequestModel UpdateChangeRequest(ctx, changeRequestId).UpdateChangeRequestModel(updateChangeRequestModel).Execute()

Update Change Request



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
	changeRequestId := int64(789) // int64 | The identifier of the Change Request.
	updateChangeRequestModel := *openapiclient.NewUpdateChangeRequestModel("Title_example") // UpdateChangeRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.UpdateChangeRequest(context.Background(), changeRequestId).UpdateChangeRequestModel(updateChangeRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.UpdateChangeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChangeRequest`: ChangeRequestModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.UpdateChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int64** | The identifier of the Change Request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateChangeRequestModel** | [**UpdateChangeRequestModel**](UpdateChangeRequestModel.md) |  | 

### Return type

[**ChangeRequestModel**](ChangeRequestModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChangeRequestComment

> ChangeRequestCommentModel UpdateChangeRequestComment(ctx, commentId).UpdateChangeRequestCommentModel(updateChangeRequestCommentModel).Execute()

Update Comment



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
	commentId := int64(789) // int64 | The identifier of the Change Request comment.
	updateChangeRequestCommentModel := *openapiclient.NewUpdateChangeRequestCommentModel("Body_example") // UpdateChangeRequestCommentModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.UpdateChangeRequestComment(context.Background(), commentId).UpdateChangeRequestCommentModel(updateChangeRequestCommentModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.UpdateChangeRequestComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChangeRequestComment`: ChangeRequestCommentModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.UpdateChangeRequestComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | The identifier of the Change Request comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChangeRequestCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateChangeRequestCommentModel** | [**UpdateChangeRequestCommentModel**](UpdateChangeRequestCommentModel.md) |  | 

### Return type

[**ChangeRequestCommentModel**](ChangeRequestCommentModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChangeRequestProposedChanges

> ChangeRequestModel UpdateChangeRequestProposedChanges(ctx, changeRequestId).UpdateChangeRequestProposedChangesModel(updateChangeRequestProposedChangesModel).SettingId(settingId).Execute()

Update Settings included in Change Request



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
	changeRequestId := int64(789) // int64 | The identifier of the Change Request.
	updateChangeRequestProposedChangesModel := *openapiclient.NewUpdateChangeRequestProposedChangesModel() // UpdateChangeRequestProposedChangesModel | 
	settingId := int32(56) // int32 | The optional identifier of the Setting. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsApprovalFlowScheduledChangesBetaAPI.UpdateChangeRequestProposedChanges(context.Background(), changeRequestId).UpdateChangeRequestProposedChangesModel(updateChangeRequestProposedChangesModel).SettingId(settingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.UpdateChangeRequestProposedChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChangeRequestProposedChanges`: ChangeRequestModel
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsApprovalFlowScheduledChangesBetaAPI.UpdateChangeRequestProposedChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int64** | The identifier of the Change Request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChangeRequestProposedChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateChangeRequestProposedChangesModel** | [**UpdateChangeRequestProposedChangesModel**](UpdateChangeRequestProposedChangesModel.md) |  | 
 **settingId** | **int32** | The optional identifier of the Setting. | 

### Return type

[**ChangeRequestModel**](ChangeRequestModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

