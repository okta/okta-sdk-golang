# AuthenticatorBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the Authenticator was created | [optional] [readonly] 
**Id** | Pointer to **string** | A unique identifier for the Authenticator | [optional] [readonly] 
**Key** | Pointer to **string** | A human-readable string that identifies the Authenticator | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the Authenticator was last modified | [optional] [readonly] 
**Name** | Pointer to **string** | Display name of the Authenticator | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The type of Authenticator | [optional] 
**Links** | Pointer to [**AuthenticatorLinks**](AuthenticatorLinks.md) |  | [optional] 

## Methods

### NewAuthenticatorBase

`func NewAuthenticatorBase() *AuthenticatorBase`

NewAuthenticatorBase instantiates a new AuthenticatorBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorBaseWithDefaults

`func NewAuthenticatorBaseWithDefaults() *AuthenticatorBase`

NewAuthenticatorBaseWithDefaults instantiates a new AuthenticatorBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AuthenticatorBase) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuthenticatorBase) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuthenticatorBase) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AuthenticatorBase) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *AuthenticatorBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticatorBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticatorBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticatorBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *AuthenticatorBase) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AuthenticatorBase) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AuthenticatorBase) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AuthenticatorBase) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AuthenticatorBase) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AuthenticatorBase) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AuthenticatorBase) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AuthenticatorBase) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *AuthenticatorBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticatorBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticatorBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthenticatorBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *AuthenticatorBase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthenticatorBase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthenticatorBase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthenticatorBase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *AuthenticatorBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticatorBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticatorBase) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthenticatorBase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *AuthenticatorBase) GetLinks() AuthenticatorLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthenticatorBase) GetLinksOk() (*AuthenticatorLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthenticatorBase) SetLinks(v AuthenticatorLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthenticatorBase) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


