# StatisticsV2Model

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasConnectedApplication** | **bool** | Indicates whether the Organization has a connected application. | 
**MillionRequestLimitPerMonth** | **int32** | The monthly request quota limit in millions. | 
**NetworkTrafficGigaByteLimitPerMonth** | **int32** | The monthly network traffic quota limit in gigabytes. | 
**OrganizationStatistics** | [**[]OrganizationMonthlyStatisticV2Model**](OrganizationMonthlyStatisticV2Model.md) | The aggregated monthly statistics for the Organization. | 
**ProductStatistics** | [**[]ProductMonthlyStatisticV2Model**](ProductMonthlyStatisticV2Model.md) | The aggregated monthly statistics for Products within the scope. | 
**DetailedStatistics** | [**[]DetailedStatisticV2Model**](DetailedStatisticV2Model.md) | The detailed per-day, per-Config, per-Environment usage statistics. | 

## Methods

### NewStatisticsV2Model

`func NewStatisticsV2Model(hasConnectedApplication bool, millionRequestLimitPerMonth int32, networkTrafficGigaByteLimitPerMonth int32, organizationStatistics []OrganizationMonthlyStatisticV2Model, productStatistics []ProductMonthlyStatisticV2Model, detailedStatistics []DetailedStatisticV2Model, ) *StatisticsV2Model`

NewStatisticsV2Model instantiates a new StatisticsV2Model object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsV2ModelWithDefaults

`func NewStatisticsV2ModelWithDefaults() *StatisticsV2Model`

NewStatisticsV2ModelWithDefaults instantiates a new StatisticsV2Model object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasConnectedApplication

`func (o *StatisticsV2Model) GetHasConnectedApplication() bool`

GetHasConnectedApplication returns the HasConnectedApplication field if non-nil, zero value otherwise.

### GetHasConnectedApplicationOk

`func (o *StatisticsV2Model) GetHasConnectedApplicationOk() (*bool, bool)`

GetHasConnectedApplicationOk returns a tuple with the HasConnectedApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConnectedApplication

`func (o *StatisticsV2Model) SetHasConnectedApplication(v bool)`

SetHasConnectedApplication sets HasConnectedApplication field to given value.


### GetMillionRequestLimitPerMonth

`func (o *StatisticsV2Model) GetMillionRequestLimitPerMonth() int32`

GetMillionRequestLimitPerMonth returns the MillionRequestLimitPerMonth field if non-nil, zero value otherwise.

### GetMillionRequestLimitPerMonthOk

`func (o *StatisticsV2Model) GetMillionRequestLimitPerMonthOk() (*int32, bool)`

GetMillionRequestLimitPerMonthOk returns a tuple with the MillionRequestLimitPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillionRequestLimitPerMonth

`func (o *StatisticsV2Model) SetMillionRequestLimitPerMonth(v int32)`

SetMillionRequestLimitPerMonth sets MillionRequestLimitPerMonth field to given value.


### GetNetworkTrafficGigaByteLimitPerMonth

`func (o *StatisticsV2Model) GetNetworkTrafficGigaByteLimitPerMonth() int32`

GetNetworkTrafficGigaByteLimitPerMonth returns the NetworkTrafficGigaByteLimitPerMonth field if non-nil, zero value otherwise.

### GetNetworkTrafficGigaByteLimitPerMonthOk

`func (o *StatisticsV2Model) GetNetworkTrafficGigaByteLimitPerMonthOk() (*int32, bool)`

GetNetworkTrafficGigaByteLimitPerMonthOk returns a tuple with the NetworkTrafficGigaByteLimitPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTrafficGigaByteLimitPerMonth

`func (o *StatisticsV2Model) SetNetworkTrafficGigaByteLimitPerMonth(v int32)`

SetNetworkTrafficGigaByteLimitPerMonth sets NetworkTrafficGigaByteLimitPerMonth field to given value.


### GetOrganizationStatistics

`func (o *StatisticsV2Model) GetOrganizationStatistics() []OrganizationMonthlyStatisticV2Model`

GetOrganizationStatistics returns the OrganizationStatistics field if non-nil, zero value otherwise.

### GetOrganizationStatisticsOk

`func (o *StatisticsV2Model) GetOrganizationStatisticsOk() (*[]OrganizationMonthlyStatisticV2Model, bool)`

GetOrganizationStatisticsOk returns a tuple with the OrganizationStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationStatistics

`func (o *StatisticsV2Model) SetOrganizationStatistics(v []OrganizationMonthlyStatisticV2Model)`

SetOrganizationStatistics sets OrganizationStatistics field to given value.


### GetProductStatistics

`func (o *StatisticsV2Model) GetProductStatistics() []ProductMonthlyStatisticV2Model`

GetProductStatistics returns the ProductStatistics field if non-nil, zero value otherwise.

### GetProductStatisticsOk

`func (o *StatisticsV2Model) GetProductStatisticsOk() (*[]ProductMonthlyStatisticV2Model, bool)`

GetProductStatisticsOk returns a tuple with the ProductStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductStatistics

`func (o *StatisticsV2Model) SetProductStatistics(v []ProductMonthlyStatisticV2Model)`

SetProductStatistics sets ProductStatistics field to given value.


### GetDetailedStatistics

`func (o *StatisticsV2Model) GetDetailedStatistics() []DetailedStatisticV2Model`

GetDetailedStatistics returns the DetailedStatistics field if non-nil, zero value otherwise.

### GetDetailedStatisticsOk

`func (o *StatisticsV2Model) GetDetailedStatisticsOk() (*[]DetailedStatisticV2Model, bool)`

GetDetailedStatisticsOk returns a tuple with the DetailedStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedStatistics

`func (o *StatisticsV2Model) SetDetailedStatistics(v []DetailedStatisticV2Model)`

SetDetailedStatistics sets DetailedStatistics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


