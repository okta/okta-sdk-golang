# IdentitySourceUserProfileForUpsertRequired

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address of the user | 
**FirstName** | Pointer to **NullableString** | First name of the user | [optional] 
**HomeAddress** | Pointer to **NullableString** | Home address of the user | [optional] 
**LastName** | Pointer to **NullableString** | Last name of the user | [optional] 
**MobilePhone** | Pointer to **NullableString** | Mobile phone number of the user | [optional] 
**SecondEmail** | Pointer to **string** | Alternative email address of the user | [optional] 
**UserName** | **string** | Username of the user | 

## Methods

### NewIdentitySourceUserProfileForUpsertRequired

`func NewIdentitySourceUserProfileForUpsertRequired(email string, userName string, ) *IdentitySourceUserProfileForUpsertRequired`

NewIdentitySourceUserProfileForUpsertRequired instantiates a new IdentitySourceUserProfileForUpsertRequired object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitySourceUserProfileForUpsertRequiredWithDefaults

`func NewIdentitySourceUserProfileForUpsertRequiredWithDefaults() *IdentitySourceUserProfileForUpsertRequired`

NewIdentitySourceUserProfileForUpsertRequiredWithDefaults instantiates a new IdentitySourceUserProfileForUpsertRequired object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *IdentitySourceUserProfileForUpsertRequired) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdentitySourceUserProfileForUpsertRequired) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdentitySourceUserProfileForUpsertRequired) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *IdentitySourceUserProfileForUpsertRequired) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IdentitySourceUserProfileForUpsertRequired) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IdentitySourceUserProfileForUpsertRequired) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *IdentitySourceUserProfileForUpsertRequired) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *IdentitySourceUserProfileForUpsertRequired) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *IdentitySourceUserProfileForUpsertRequired) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetHomeAddress

`func (o *IdentitySourceUserProfileForUpsertRequired) GetHomeAddress() string`

GetHomeAddress returns the HomeAddress field if non-nil, zero value otherwise.

### GetHomeAddressOk

`func (o *IdentitySourceUserProfileForUpsertRequired) GetHomeAddressOk() (*string, bool)`

GetHomeAddressOk returns a tuple with the HomeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeAddress

`func (o *IdentitySourceUserProfileForUpsertRequired) SetHomeAddress(v string)`

SetHomeAddress sets HomeAddress field to given value.

### HasHomeAddress

`func (o *IdentitySourceUserProfileForUpsertRequired) HasHomeAddress() bool`

HasHomeAddress returns a boolean if a field has been set.

### SetHomeAddressNil

`func (o *IdentitySourceUserProfileForUpsertRequired) SetHomeAddressNil(b bool)`

 SetHomeAddressNil sets the value for HomeAddress to be an explicit nil

### UnsetHomeAddress
`func (o *IdentitySourceUserProfileForUpsertRequired) UnsetHomeAddress()`

UnsetHomeAddress ensures that no value is present for HomeAddress, not even an explicit nil
### GetLastName

`func (o *IdentitySourceUserProfileForUpsertRequired) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IdentitySourceUserProfileForUpsertRequired) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IdentitySourceUserProfileForUpsertRequired) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *IdentitySourceUserProfileForUpsertRequired) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *IdentitySourceUserProfileForUpsertRequired) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *IdentitySourceUserProfileForUpsertRequired) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMobilePhone

`func (o *IdentitySourceUserProfileForUpsertRequired) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *IdentitySourceUserProfileForUpsertRequired) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *IdentitySourceUserProfileForUpsertRequired) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *IdentitySourceUserProfileForUpsertRequired) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *IdentitySourceUserProfileForUpsertRequired) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *IdentitySourceUserProfileForUpsertRequired) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetSecondEmail

`func (o *IdentitySourceUserProfileForUpsertRequired) GetSecondEmail() string`

GetSecondEmail returns the SecondEmail field if non-nil, zero value otherwise.

### GetSecondEmailOk

`func (o *IdentitySourceUserProfileForUpsertRequired) GetSecondEmailOk() (*string, bool)`

GetSecondEmailOk returns a tuple with the SecondEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondEmail

`func (o *IdentitySourceUserProfileForUpsertRequired) SetSecondEmail(v string)`

SetSecondEmail sets SecondEmail field to given value.

### HasSecondEmail

`func (o *IdentitySourceUserProfileForUpsertRequired) HasSecondEmail() bool`

HasSecondEmail returns a boolean if a field has been set.

### GetUserName

`func (o *IdentitySourceUserProfileForUpsertRequired) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *IdentitySourceUserProfileForUpsertRequired) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *IdentitySourceUserProfileForUpsertRequired) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


