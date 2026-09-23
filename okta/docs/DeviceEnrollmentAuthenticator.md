# DeviceEnrollmentAuthenticator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp when the enrollment authenticator was created | [readonly] 
**Id** | **string** | The unique identifier of the enrollment authenticator | [readonly] 
**Key** | **string** | The authenticator key. For example: &#x60;okta_verify&#x60;. | 
**LastUpdated** | **time.Time** | Timestamp when the enrollment authenticator was last updated | [readonly] 
**LastVerified** | Pointer to **time.Time** | Timestamp when the enrollment authenticator was last verified | [optional] [readonly] 
**Name** | **string** | The authenticator display name | 
**Nickname** | Pointer to **string** | A user-friendly name for the enrollment authenticator | [optional] 
**Profile** | Pointer to [**DeviceEnrollmentAuthenticatorProfile**](DeviceEnrollmentAuthenticatorProfile.md) |  | [optional] 
**Status** | **string** | Status of the enrollment | 
**Type** | **string** | The type of authenticator | 
**Links** | [**DeviceEnrollmentAuthenticatorLinks**](DeviceEnrollmentAuthenticatorLinks.md) |  | 

## Methods

### NewDeviceEnrollmentAuthenticator

`func NewDeviceEnrollmentAuthenticator(created time.Time, id string, key string, lastUpdated time.Time, name string, status string, type_ string, links DeviceEnrollmentAuthenticatorLinks, ) *DeviceEnrollmentAuthenticator`

NewDeviceEnrollmentAuthenticator instantiates a new DeviceEnrollmentAuthenticator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceEnrollmentAuthenticatorWithDefaults

`func NewDeviceEnrollmentAuthenticatorWithDefaults() *DeviceEnrollmentAuthenticator`

NewDeviceEnrollmentAuthenticatorWithDefaults instantiates a new DeviceEnrollmentAuthenticator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *DeviceEnrollmentAuthenticator) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeviceEnrollmentAuthenticator) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeviceEnrollmentAuthenticator) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *DeviceEnrollmentAuthenticator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceEnrollmentAuthenticator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceEnrollmentAuthenticator) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *DeviceEnrollmentAuthenticator) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DeviceEnrollmentAuthenticator) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DeviceEnrollmentAuthenticator) SetKey(v string)`

SetKey sets Key field to given value.


### GetLastUpdated

`func (o *DeviceEnrollmentAuthenticator) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeviceEnrollmentAuthenticator) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeviceEnrollmentAuthenticator) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastVerified

`func (o *DeviceEnrollmentAuthenticator) GetLastVerified() time.Time`

GetLastVerified returns the LastVerified field if non-nil, zero value otherwise.

### GetLastVerifiedOk

`func (o *DeviceEnrollmentAuthenticator) GetLastVerifiedOk() (*time.Time, bool)`

GetLastVerifiedOk returns a tuple with the LastVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerified

`func (o *DeviceEnrollmentAuthenticator) SetLastVerified(v time.Time)`

SetLastVerified sets LastVerified field to given value.

### HasLastVerified

`func (o *DeviceEnrollmentAuthenticator) HasLastVerified() bool`

HasLastVerified returns a boolean if a field has been set.

### GetName

`func (o *DeviceEnrollmentAuthenticator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceEnrollmentAuthenticator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceEnrollmentAuthenticator) SetName(v string)`

SetName sets Name field to given value.


### GetNickname

`func (o *DeviceEnrollmentAuthenticator) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *DeviceEnrollmentAuthenticator) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *DeviceEnrollmentAuthenticator) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *DeviceEnrollmentAuthenticator) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetProfile

`func (o *DeviceEnrollmentAuthenticator) GetProfile() DeviceEnrollmentAuthenticatorProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *DeviceEnrollmentAuthenticator) GetProfileOk() (*DeviceEnrollmentAuthenticatorProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *DeviceEnrollmentAuthenticator) SetProfile(v DeviceEnrollmentAuthenticatorProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *DeviceEnrollmentAuthenticator) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceEnrollmentAuthenticator) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceEnrollmentAuthenticator) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceEnrollmentAuthenticator) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *DeviceEnrollmentAuthenticator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceEnrollmentAuthenticator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceEnrollmentAuthenticator) SetType(v string)`

SetType sets Type field to given value.


### GetLinks

`func (o *DeviceEnrollmentAuthenticator) GetLinks() DeviceEnrollmentAuthenticatorLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceEnrollmentAuthenticator) GetLinksOk() (*DeviceEnrollmentAuthenticatorLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceEnrollmentAuthenticator) SetLinks(v DeviceEnrollmentAuthenticatorLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


