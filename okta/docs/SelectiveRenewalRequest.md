# SelectiveRenewalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParallelConfigs** | Pointer to [**[]ParallelScepConfig**](ParallelScepConfig.md) | The configurations to clone against the renewed certificate. Each source configuration stays on the certificate it&#39;s already bound to, so both certificates can enroll devices during the migration. | [optional] 
**RolloverConfigIds** | Pointer to **[]string** | The IDs of the configurations to move onto the renewed certificate authority. Each keeps its ID and SCEP enrollment URL. | [optional] 

## Methods

### NewSelectiveRenewalRequest

`func NewSelectiveRenewalRequest() *SelectiveRenewalRequest`

NewSelectiveRenewalRequest instantiates a new SelectiveRenewalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectiveRenewalRequestWithDefaults

`func NewSelectiveRenewalRequestWithDefaults() *SelectiveRenewalRequest`

NewSelectiveRenewalRequestWithDefaults instantiates a new SelectiveRenewalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParallelConfigs

`func (o *SelectiveRenewalRequest) GetParallelConfigs() []ParallelScepConfig`

GetParallelConfigs returns the ParallelConfigs field if non-nil, zero value otherwise.

### GetParallelConfigsOk

`func (o *SelectiveRenewalRequest) GetParallelConfigsOk() (*[]ParallelScepConfig, bool)`

GetParallelConfigsOk returns a tuple with the ParallelConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelConfigs

`func (o *SelectiveRenewalRequest) SetParallelConfigs(v []ParallelScepConfig)`

SetParallelConfigs sets ParallelConfigs field to given value.

### HasParallelConfigs

`func (o *SelectiveRenewalRequest) HasParallelConfigs() bool`

HasParallelConfigs returns a boolean if a field has been set.

### GetRolloverConfigIds

`func (o *SelectiveRenewalRequest) GetRolloverConfigIds() []string`

GetRolloverConfigIds returns the RolloverConfigIds field if non-nil, zero value otherwise.

### GetRolloverConfigIdsOk

`func (o *SelectiveRenewalRequest) GetRolloverConfigIdsOk() (*[]string, bool)`

GetRolloverConfigIdsOk returns a tuple with the RolloverConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverConfigIds

`func (o *SelectiveRenewalRequest) SetRolloverConfigIds(v []string)`

SetRolloverConfigIds sets RolloverConfigIds field to given value.

### HasRolloverConfigIds

`func (o *SelectiveRenewalRequest) HasRolloverConfigIds() bool`

HasRolloverConfigIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


