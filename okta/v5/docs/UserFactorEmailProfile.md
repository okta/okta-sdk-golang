# UserFactorEmailProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email address of the user. Must be either the primary or secondary email address associated with the Okta user account. | [optional] 

## Methods

### NewUserFactorEmailProfile

`func NewUserFactorEmailProfile() *UserFactorEmailProfile`

NewUserFactorEmailProfile instantiates a new UserFactorEmailProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorEmailProfileWithDefaults

`func NewUserFactorEmailProfileWithDefaults() *UserFactorEmailProfile`

NewUserFactorEmailProfileWithDefaults instantiates a new UserFactorEmailProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserFactorEmailProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserFactorEmailProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserFactorEmailProfile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserFactorEmailProfile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


