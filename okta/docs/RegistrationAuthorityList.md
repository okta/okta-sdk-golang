# RegistrationAuthorityList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RegistrationAuthority**](RegistrationAuthority.md) | The configurations, grouped by the certificate they&#39;re bound to, newest certificate first | [optional] 

## Methods

### NewRegistrationAuthorityList

`func NewRegistrationAuthorityList() *RegistrationAuthorityList`

NewRegistrationAuthorityList instantiates a new RegistrationAuthorityList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationAuthorityListWithDefaults

`func NewRegistrationAuthorityListWithDefaults() *RegistrationAuthorityList`

NewRegistrationAuthorityListWithDefaults instantiates a new RegistrationAuthorityList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RegistrationAuthorityList) GetData() []RegistrationAuthority`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RegistrationAuthorityList) GetDataOk() (*[]RegistrationAuthority, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RegistrationAuthorityList) SetData(v []RegistrationAuthority)`

SetData sets Data field to given value.

### HasData

`func (o *RegistrationAuthorityList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


