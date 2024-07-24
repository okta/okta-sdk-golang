# TelephonyRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**TelephonyRequestDataContext**](TelephonyRequestDataContext.md) |  | [optional] 
**MessageProfile** | Pointer to [**TelephonyRequestDataMessageProfile**](TelephonyRequestDataMessageProfile.md) |  | [optional] 
**UserProfile** | Pointer to [**TelephonyRequestDataUserProfile**](TelephonyRequestDataUserProfile.md) |  | [optional] 

## Methods

### NewTelephonyRequestData

`func NewTelephonyRequestData() *TelephonyRequestData`

NewTelephonyRequestData instantiates a new TelephonyRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelephonyRequestDataWithDefaults

`func NewTelephonyRequestDataWithDefaults() *TelephonyRequestData`

NewTelephonyRequestDataWithDefaults instantiates a new TelephonyRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *TelephonyRequestData) GetContext() TelephonyRequestDataContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TelephonyRequestData) GetContextOk() (*TelephonyRequestDataContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TelephonyRequestData) SetContext(v TelephonyRequestDataContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TelephonyRequestData) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetMessageProfile

`func (o *TelephonyRequestData) GetMessageProfile() TelephonyRequestDataMessageProfile`

GetMessageProfile returns the MessageProfile field if non-nil, zero value otherwise.

### GetMessageProfileOk

`func (o *TelephonyRequestData) GetMessageProfileOk() (*TelephonyRequestDataMessageProfile, bool)`

GetMessageProfileOk returns a tuple with the MessageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageProfile

`func (o *TelephonyRequestData) SetMessageProfile(v TelephonyRequestDataMessageProfile)`

SetMessageProfile sets MessageProfile field to given value.

### HasMessageProfile

`func (o *TelephonyRequestData) HasMessageProfile() bool`

HasMessageProfile returns a boolean if a field has been set.

### GetUserProfile

`func (o *TelephonyRequestData) GetUserProfile() TelephonyRequestDataUserProfile`

GetUserProfile returns the UserProfile field if non-nil, zero value otherwise.

### GetUserProfileOk

`func (o *TelephonyRequestData) GetUserProfileOk() (*TelephonyRequestDataUserProfile, bool)`

GetUserProfileOk returns a tuple with the UserProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfile

`func (o *TelephonyRequestData) SetUserProfile(v TelephonyRequestDataUserProfile)`

SetUserProfile sets UserProfile field to given value.

### HasUserProfile

`func (o *TelephonyRequestData) HasUserProfile() bool`

HasUserProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


