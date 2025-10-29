# ApplicationUniversalLogout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityStack** | Pointer to **string** | Indicates whether the app uses a shared identity stack that may cause the user to sign out of other apps by the same company | [optional] 
**Protocol** | Pointer to **string** | The protocol used for Universal Logout | [optional] 
**Status** | Pointer to **string** | Universal Logout status for the app instance | [optional] 
**SupportType** | Pointer to **string** | Indicates whether the app supports full or partial Universal Logout (UL). | [optional] 

## Methods

### NewApplicationUniversalLogout

`func NewApplicationUniversalLogout() *ApplicationUniversalLogout`

NewApplicationUniversalLogout instantiates a new ApplicationUniversalLogout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationUniversalLogoutWithDefaults

`func NewApplicationUniversalLogoutWithDefaults() *ApplicationUniversalLogout`

NewApplicationUniversalLogoutWithDefaults instantiates a new ApplicationUniversalLogout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityStack

`func (o *ApplicationUniversalLogout) GetIdentityStack() string`

GetIdentityStack returns the IdentityStack field if non-nil, zero value otherwise.

### GetIdentityStackOk

`func (o *ApplicationUniversalLogout) GetIdentityStackOk() (*string, bool)`

GetIdentityStackOk returns a tuple with the IdentityStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStack

`func (o *ApplicationUniversalLogout) SetIdentityStack(v string)`

SetIdentityStack sets IdentityStack field to given value.

### HasIdentityStack

`func (o *ApplicationUniversalLogout) HasIdentityStack() bool`

HasIdentityStack returns a boolean if a field has been set.

### GetProtocol

`func (o *ApplicationUniversalLogout) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApplicationUniversalLogout) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApplicationUniversalLogout) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ApplicationUniversalLogout) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetStatus

`func (o *ApplicationUniversalLogout) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationUniversalLogout) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationUniversalLogout) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplicationUniversalLogout) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportType

`func (o *ApplicationUniversalLogout) GetSupportType() string`

GetSupportType returns the SupportType field if non-nil, zero value otherwise.

### GetSupportTypeOk

`func (o *ApplicationUniversalLogout) GetSupportTypeOk() (*string, bool)`

GetSupportTypeOk returns a tuple with the SupportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportType

`func (o *ApplicationUniversalLogout) SetSupportType(v string)`

SetSupportType sets SupportType field to given value.

### HasSupportType

`func (o *ApplicationUniversalLogout) HasSupportType() bool`

HasSupportType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


