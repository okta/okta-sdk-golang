# AuthenticatorEnrollment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the authenticator enrollment was created | [optional] 
**Id** | Pointer to **string** | The unique identifier of the authenticator enrollment | [optional] 
**Key** | Pointer to **string** | A human-readable string that identifies the authenticator | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the authenticator enrollment was last updated | [optional] 
**Name** | Pointer to **string** | The authenticator display name | [optional] 
**Profile** | Pointer to [**AuthenticatorProfile**](AuthenticatorProfile.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the enrollment | [optional] 
**Type** | Pointer to **string** | The type of authenticator | [optional] 
**Links** | Pointer to [**AuthenticatorEnrollmentLinks**](AuthenticatorEnrollmentLinks.md) |  | [optional] 

## Methods

### NewAuthenticatorEnrollment

`func NewAuthenticatorEnrollment() *AuthenticatorEnrollment`

NewAuthenticatorEnrollment instantiates a new AuthenticatorEnrollment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorEnrollmentWithDefaults

`func NewAuthenticatorEnrollmentWithDefaults() *AuthenticatorEnrollment`

NewAuthenticatorEnrollmentWithDefaults instantiates a new AuthenticatorEnrollment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AuthenticatorEnrollment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuthenticatorEnrollment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuthenticatorEnrollment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AuthenticatorEnrollment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *AuthenticatorEnrollment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticatorEnrollment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticatorEnrollment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticatorEnrollment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *AuthenticatorEnrollment) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AuthenticatorEnrollment) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AuthenticatorEnrollment) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AuthenticatorEnrollment) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AuthenticatorEnrollment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AuthenticatorEnrollment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AuthenticatorEnrollment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AuthenticatorEnrollment) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *AuthenticatorEnrollment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticatorEnrollment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticatorEnrollment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthenticatorEnrollment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfile

`func (o *AuthenticatorEnrollment) GetProfile() AuthenticatorProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AuthenticatorEnrollment) GetProfileOk() (*AuthenticatorProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AuthenticatorEnrollment) SetProfile(v AuthenticatorProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AuthenticatorEnrollment) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetStatus

`func (o *AuthenticatorEnrollment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthenticatorEnrollment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthenticatorEnrollment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthenticatorEnrollment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *AuthenticatorEnrollment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticatorEnrollment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticatorEnrollment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthenticatorEnrollment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *AuthenticatorEnrollment) GetLinks() AuthenticatorEnrollmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthenticatorEnrollment) GetLinksOk() (*AuthenticatorEnrollmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthenticatorEnrollment) SetLinks(v AuthenticatorEnrollmentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthenticatorEnrollment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


