# TacAuthenticatorEnrollment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the authenticator enrollment was created | [optional] 
**Id** | Pointer to **string** | A unique identifier of the authenticator enrollment | [optional] 
**Key** | Pointer to **string** | A human-readable string that identifies the authenticator | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the authenticator enrollment was last updated | [optional] 
**Name** | Pointer to **string** | The authenticator display name | [optional] 
**Nickname** | Pointer to **string** | A user-friendly name for the authenticator enrollment | [optional] 
**Profile** | Pointer to [**AuthenticatorProfileTacResponsePost**](AuthenticatorProfileTacResponsePost.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the enrollment | [optional] 
**Type** | Pointer to **string** | The type of authenticator | [optional] 
**Links** | Pointer to [**AuthenticatorEnrollmentLinks**](AuthenticatorEnrollmentLinks.md) |  | [optional] 

## Methods

### NewTacAuthenticatorEnrollment

`func NewTacAuthenticatorEnrollment() *TacAuthenticatorEnrollment`

NewTacAuthenticatorEnrollment instantiates a new TacAuthenticatorEnrollment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTacAuthenticatorEnrollmentWithDefaults

`func NewTacAuthenticatorEnrollmentWithDefaults() *TacAuthenticatorEnrollment`

NewTacAuthenticatorEnrollmentWithDefaults instantiates a new TacAuthenticatorEnrollment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *TacAuthenticatorEnrollment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TacAuthenticatorEnrollment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TacAuthenticatorEnrollment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TacAuthenticatorEnrollment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *TacAuthenticatorEnrollment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TacAuthenticatorEnrollment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TacAuthenticatorEnrollment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TacAuthenticatorEnrollment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *TacAuthenticatorEnrollment) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TacAuthenticatorEnrollment) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TacAuthenticatorEnrollment) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TacAuthenticatorEnrollment) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastUpdated

`func (o *TacAuthenticatorEnrollment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *TacAuthenticatorEnrollment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *TacAuthenticatorEnrollment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *TacAuthenticatorEnrollment) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *TacAuthenticatorEnrollment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TacAuthenticatorEnrollment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TacAuthenticatorEnrollment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TacAuthenticatorEnrollment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *TacAuthenticatorEnrollment) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *TacAuthenticatorEnrollment) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *TacAuthenticatorEnrollment) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *TacAuthenticatorEnrollment) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetProfile

`func (o *TacAuthenticatorEnrollment) GetProfile() AuthenticatorProfileTacResponsePost`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *TacAuthenticatorEnrollment) GetProfileOk() (*AuthenticatorProfileTacResponsePost, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *TacAuthenticatorEnrollment) SetProfile(v AuthenticatorProfileTacResponsePost)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *TacAuthenticatorEnrollment) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetStatus

`func (o *TacAuthenticatorEnrollment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TacAuthenticatorEnrollment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TacAuthenticatorEnrollment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TacAuthenticatorEnrollment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *TacAuthenticatorEnrollment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TacAuthenticatorEnrollment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TacAuthenticatorEnrollment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TacAuthenticatorEnrollment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *TacAuthenticatorEnrollment) GetLinks() AuthenticatorEnrollmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TacAuthenticatorEnrollment) GetLinksOk() (*AuthenticatorEnrollmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TacAuthenticatorEnrollment) SetLinks(v AuthenticatorEnrollmentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TacAuthenticatorEnrollment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


