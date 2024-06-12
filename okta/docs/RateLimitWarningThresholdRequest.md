# RateLimitWarningThresholdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarningThreshold** | **int32** | The threshold value (percentage) of a rate limit that, when exceeded, triggers a warning notification. By default, this value is 90 for Workforce orgs and 60 for CIAM orgs. | 

## Methods

### NewRateLimitWarningThresholdRequest

`func NewRateLimitWarningThresholdRequest(warningThreshold int32, ) *RateLimitWarningThresholdRequest`

NewRateLimitWarningThresholdRequest instantiates a new RateLimitWarningThresholdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitWarningThresholdRequestWithDefaults

`func NewRateLimitWarningThresholdRequestWithDefaults() *RateLimitWarningThresholdRequest`

NewRateLimitWarningThresholdRequestWithDefaults instantiates a new RateLimitWarningThresholdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarningThreshold

`func (o *RateLimitWarningThresholdRequest) GetWarningThreshold() int32`

GetWarningThreshold returns the WarningThreshold field if non-nil, zero value otherwise.

### GetWarningThresholdOk

`func (o *RateLimitWarningThresholdRequest) GetWarningThresholdOk() (*int32, bool)`

GetWarningThresholdOk returns a tuple with the WarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThreshold

`func (o *RateLimitWarningThresholdRequest) SetWarningThreshold(v int32)`

SetWarningThreshold sets WarningThreshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


