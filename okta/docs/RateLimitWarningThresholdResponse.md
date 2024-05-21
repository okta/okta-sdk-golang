# RateLimitWarningThresholdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarningThreshold** | Pointer to **int32** | The threshold value (percentage) of a rate limit that, when exceeded, triggers a warning notification. By default, this value is 90 for Workforce orgs and 60 for CIAM orgs. | [optional] 

## Methods

### NewRateLimitWarningThresholdResponse

`func NewRateLimitWarningThresholdResponse() *RateLimitWarningThresholdResponse`

NewRateLimitWarningThresholdResponse instantiates a new RateLimitWarningThresholdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitWarningThresholdResponseWithDefaults

`func NewRateLimitWarningThresholdResponseWithDefaults() *RateLimitWarningThresholdResponse`

NewRateLimitWarningThresholdResponseWithDefaults instantiates a new RateLimitWarningThresholdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarningThreshold

`func (o *RateLimitWarningThresholdResponse) GetWarningThreshold() int32`

GetWarningThreshold returns the WarningThreshold field if non-nil, zero value otherwise.

### GetWarningThresholdOk

`func (o *RateLimitWarningThresholdResponse) GetWarningThresholdOk() (*int32, bool)`

GetWarningThresholdOk returns a tuple with the WarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThreshold

`func (o *RateLimitWarningThresholdResponse) SetWarningThreshold(v int32)`

SetWarningThreshold sets WarningThreshold field to given value.

### HasWarningThreshold

`func (o *RateLimitWarningThresholdResponse) HasWarningThreshold() bool`

HasWarningThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


