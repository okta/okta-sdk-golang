# GoogleApplicationSettingsApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Your Google company domain | 
**RpId** | Pointer to **string** | RPID | [optional] 

## Methods

### NewGoogleApplicationSettingsApplication

`func NewGoogleApplicationSettingsApplication(domain string, ) *GoogleApplicationSettingsApplication`

NewGoogleApplicationSettingsApplication instantiates a new GoogleApplicationSettingsApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleApplicationSettingsApplicationWithDefaults

`func NewGoogleApplicationSettingsApplicationWithDefaults() *GoogleApplicationSettingsApplication`

NewGoogleApplicationSettingsApplicationWithDefaults instantiates a new GoogleApplicationSettingsApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *GoogleApplicationSettingsApplication) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GoogleApplicationSettingsApplication) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GoogleApplicationSettingsApplication) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetRpId

`func (o *GoogleApplicationSettingsApplication) GetRpId() string`

GetRpId returns the RpId field if non-nil, zero value otherwise.

### GetRpIdOk

`func (o *GoogleApplicationSettingsApplication) GetRpIdOk() (*string, bool)`

GetRpIdOk returns a tuple with the RpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpId

`func (o *GoogleApplicationSettingsApplication) SetRpId(v string)`

SetRpId sets RpId field to given value.

### HasRpId

`func (o *GoogleApplicationSettingsApplication) HasRpId() bool`

HasRpId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


