# ConnectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientKey** | **string** |  | 
**JiraJwtToken** | **string** |  | 

## Methods

### NewConnectRequest

`func NewConnectRequest(clientKey string, jiraJwtToken string, ) *ConnectRequest`

NewConnectRequest instantiates a new ConnectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectRequestWithDefaults

`func NewConnectRequestWithDefaults() *ConnectRequest`

NewConnectRequestWithDefaults instantiates a new ConnectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientKey

`func (o *ConnectRequest) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *ConnectRequest) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *ConnectRequest) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.


### GetJiraJwtToken

`func (o *ConnectRequest) GetJiraJwtToken() string`

GetJiraJwtToken returns the JiraJwtToken field if non-nil, zero value otherwise.

### GetJiraJwtTokenOk

`func (o *ConnectRequest) GetJiraJwtTokenOk() (*string, bool)`

GetJiraJwtTokenOk returns a tuple with the JiraJwtToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraJwtToken

`func (o *ConnectRequest) SetJiraJwtToken(v string)`

SetJiraJwtToken sets JiraJwtToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


