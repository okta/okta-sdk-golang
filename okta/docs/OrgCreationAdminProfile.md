# OrgCreationAdminProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **NullableString** | Given name of the User (&#x60;givenName&#x60;) | 
**LastName** | **NullableString** | The family name of the User (&#x60;familyName&#x60;) | 
**Email** | **string** | The primary email address of the User. For validation, see [RFC 5322 Section 3.2.3](https://datatracker.ietf.org/doc/html/rfc5322#section-3.2.3). | 
**Login** | **string** | The unique identifier for the User (&#x60;username&#x60;) | 

## Methods

### NewOrgCreationAdminProfile

`func NewOrgCreationAdminProfile(firstName NullableString, lastName NullableString, email string, login string, ) *OrgCreationAdminProfile`

NewOrgCreationAdminProfile instantiates a new OrgCreationAdminProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgCreationAdminProfileWithDefaults

`func NewOrgCreationAdminProfileWithDefaults() *OrgCreationAdminProfile`

NewOrgCreationAdminProfileWithDefaults instantiates a new OrgCreationAdminProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *OrgCreationAdminProfile) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OrgCreationAdminProfile) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OrgCreationAdminProfile) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *OrgCreationAdminProfile) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *OrgCreationAdminProfile) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *OrgCreationAdminProfile) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *OrgCreationAdminProfile) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *OrgCreationAdminProfile) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *OrgCreationAdminProfile) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *OrgCreationAdminProfile) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetEmail

`func (o *OrgCreationAdminProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrgCreationAdminProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrgCreationAdminProfile) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLogin

`func (o *OrgCreationAdminProfile) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *OrgCreationAdminProfile) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *OrgCreationAdminProfile) SetLogin(v string)`

SetLogin sets Login field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


