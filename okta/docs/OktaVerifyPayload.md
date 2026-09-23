# OktaVerifyPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientVersion** | Pointer to **string** | The version of the Okta Verify client | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the client instance was last updated | [optional] [readonly] 

## Methods

### NewOktaVerifyPayload

`func NewOktaVerifyPayload() *OktaVerifyPayload`

NewOktaVerifyPayload instantiates a new OktaVerifyPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaVerifyPayloadWithDefaults

`func NewOktaVerifyPayloadWithDefaults() *OktaVerifyPayload`

NewOktaVerifyPayloadWithDefaults instantiates a new OktaVerifyPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientVersion

`func (o *OktaVerifyPayload) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *OktaVerifyPayload) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *OktaVerifyPayload) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.

### HasClientVersion

`func (o *OktaVerifyPayload) HasClientVersion() bool`

HasClientVersion returns a boolean if a field has been set.

### GetLastUpdated

`func (o *OktaVerifyPayload) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *OktaVerifyPayload) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *OktaVerifyPayload) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *OktaVerifyPayload) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


