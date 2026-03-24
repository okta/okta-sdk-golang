# WindowsOSAccountProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryJoinStatus** | Pointer to **string** | Active Directory join status | [optional] 
**DownLevelUsername** | Pointer to **string** | Domain\\username format (down-level logon name) | [optional] 
**FullName** | Pointer to **string** | Full name of the account user | [optional] 
**GUID** | Pointer to **string** | Globally Unique Identifier for the account | [optional] 
**QualifiedUsername** | Pointer to **string** | Fully qualified username | [optional] 
**SecurityId** | Pointer to **string** | Windows Security Identifier (SID) | [optional] 
**Upn** | Pointer to **string** | User principal name | [optional] 

## Methods

### NewWindowsOSAccountProfile

`func NewWindowsOSAccountProfile() *WindowsOSAccountProfile`

NewWindowsOSAccountProfile instantiates a new WindowsOSAccountProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsOSAccountProfileWithDefaults

`func NewWindowsOSAccountProfileWithDefaults() *WindowsOSAccountProfile`

NewWindowsOSAccountProfileWithDefaults instantiates a new WindowsOSAccountProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryJoinStatus

`func (o *WindowsOSAccountProfile) GetDirectoryJoinStatus() string`

GetDirectoryJoinStatus returns the DirectoryJoinStatus field if non-nil, zero value otherwise.

### GetDirectoryJoinStatusOk

`func (o *WindowsOSAccountProfile) GetDirectoryJoinStatusOk() (*string, bool)`

GetDirectoryJoinStatusOk returns a tuple with the DirectoryJoinStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryJoinStatus

`func (o *WindowsOSAccountProfile) SetDirectoryJoinStatus(v string)`

SetDirectoryJoinStatus sets DirectoryJoinStatus field to given value.

### HasDirectoryJoinStatus

`func (o *WindowsOSAccountProfile) HasDirectoryJoinStatus() bool`

HasDirectoryJoinStatus returns a boolean if a field has been set.

### GetDownLevelUsername

`func (o *WindowsOSAccountProfile) GetDownLevelUsername() string`

GetDownLevelUsername returns the DownLevelUsername field if non-nil, zero value otherwise.

### GetDownLevelUsernameOk

`func (o *WindowsOSAccountProfile) GetDownLevelUsernameOk() (*string, bool)`

GetDownLevelUsernameOk returns a tuple with the DownLevelUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLevelUsername

`func (o *WindowsOSAccountProfile) SetDownLevelUsername(v string)`

SetDownLevelUsername sets DownLevelUsername field to given value.

### HasDownLevelUsername

`func (o *WindowsOSAccountProfile) HasDownLevelUsername() bool`

HasDownLevelUsername returns a boolean if a field has been set.

### GetFullName

`func (o *WindowsOSAccountProfile) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *WindowsOSAccountProfile) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *WindowsOSAccountProfile) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *WindowsOSAccountProfile) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetGUID

`func (o *WindowsOSAccountProfile) GetGUID() string`

GetGUID returns the GUID field if non-nil, zero value otherwise.

### GetGUIDOk

`func (o *WindowsOSAccountProfile) GetGUIDOk() (*string, bool)`

GetGUIDOk returns a tuple with the GUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGUID

`func (o *WindowsOSAccountProfile) SetGUID(v string)`

SetGUID sets GUID field to given value.

### HasGUID

`func (o *WindowsOSAccountProfile) HasGUID() bool`

HasGUID returns a boolean if a field has been set.

### GetQualifiedUsername

`func (o *WindowsOSAccountProfile) GetQualifiedUsername() string`

GetQualifiedUsername returns the QualifiedUsername field if non-nil, zero value otherwise.

### GetQualifiedUsernameOk

`func (o *WindowsOSAccountProfile) GetQualifiedUsernameOk() (*string, bool)`

GetQualifiedUsernameOk returns a tuple with the QualifiedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedUsername

`func (o *WindowsOSAccountProfile) SetQualifiedUsername(v string)`

SetQualifiedUsername sets QualifiedUsername field to given value.

### HasQualifiedUsername

`func (o *WindowsOSAccountProfile) HasQualifiedUsername() bool`

HasQualifiedUsername returns a boolean if a field has been set.

### GetSecurityId

`func (o *WindowsOSAccountProfile) GetSecurityId() string`

GetSecurityId returns the SecurityId field if non-nil, zero value otherwise.

### GetSecurityIdOk

`func (o *WindowsOSAccountProfile) GetSecurityIdOk() (*string, bool)`

GetSecurityIdOk returns a tuple with the SecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityId

`func (o *WindowsOSAccountProfile) SetSecurityId(v string)`

SetSecurityId sets SecurityId field to given value.

### HasSecurityId

`func (o *WindowsOSAccountProfile) HasSecurityId() bool`

HasSecurityId returns a boolean if a field has been set.

### GetUpn

`func (o *WindowsOSAccountProfile) GetUpn() string`

GetUpn returns the Upn field if non-nil, zero value otherwise.

### GetUpnOk

`func (o *WindowsOSAccountProfile) GetUpnOk() (*string, bool)`

GetUpnOk returns a tuple with the Upn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpn

`func (o *WindowsOSAccountProfile) SetUpn(v string)`

SetUpn sets Upn field to given value.

### HasUpn

`func (o *WindowsOSAccountProfile) HasUpn() bool`

HasUpn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


