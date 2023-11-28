# SloParticipate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingType** | Pointer to **string** | Request binding type | [optional] 
**Enabled** | Pointer to **bool** | Allows the app to participate in front-channel single logout. | [optional] 
**LogoutRequestUrl** | Pointer to **string** | URL where Okta sends the logout request. | [optional] 
**SessionIndexRequired** | Pointer to **bool** | Include user session details. | [optional] 

## Methods

### NewSloParticipate

`func NewSloParticipate() *SloParticipate`

NewSloParticipate instantiates a new SloParticipate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloParticipateWithDefaults

`func NewSloParticipateWithDefaults() *SloParticipate`

NewSloParticipateWithDefaults instantiates a new SloParticipate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingType

`func (o *SloParticipate) GetBindingType() string`

GetBindingType returns the BindingType field if non-nil, zero value otherwise.

### GetBindingTypeOk

`func (o *SloParticipate) GetBindingTypeOk() (*string, bool)`

GetBindingTypeOk returns a tuple with the BindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingType

`func (o *SloParticipate) SetBindingType(v string)`

SetBindingType sets BindingType field to given value.

### HasBindingType

`func (o *SloParticipate) HasBindingType() bool`

HasBindingType returns a boolean if a field has been set.

### GetEnabled

`func (o *SloParticipate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SloParticipate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SloParticipate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SloParticipate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLogoutRequestUrl

`func (o *SloParticipate) GetLogoutRequestUrl() string`

GetLogoutRequestUrl returns the LogoutRequestUrl field if non-nil, zero value otherwise.

### GetLogoutRequestUrlOk

`func (o *SloParticipate) GetLogoutRequestUrlOk() (*string, bool)`

GetLogoutRequestUrlOk returns a tuple with the LogoutRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutRequestUrl

`func (o *SloParticipate) SetLogoutRequestUrl(v string)`

SetLogoutRequestUrl sets LogoutRequestUrl field to given value.

### HasLogoutRequestUrl

`func (o *SloParticipate) HasLogoutRequestUrl() bool`

HasLogoutRequestUrl returns a boolean if a field has been set.

### GetSessionIndexRequired

`func (o *SloParticipate) GetSessionIndexRequired() bool`

GetSessionIndexRequired returns the SessionIndexRequired field if non-nil, zero value otherwise.

### GetSessionIndexRequiredOk

`func (o *SloParticipate) GetSessionIndexRequiredOk() (*bool, bool)`

GetSessionIndexRequiredOk returns a tuple with the SessionIndexRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIndexRequired

`func (o *SloParticipate) SetSessionIndexRequired(v bool)`

SetSessionIndexRequired sets SessionIndexRequired field to given value.

### HasSessionIndexRequired

`func (o *SloParticipate) HasSessionIndexRequired() bool`

HasSessionIndexRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


