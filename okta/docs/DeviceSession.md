# DeviceSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Factors** | [**[]DeviceSessionFactor**](DeviceSessionFactor.md) | Authenticator methods that are verified in this session, with the most recently verified authenticator method first. This array always has at least one item. | 
**Id** | **string** | The session&#39;s external identifier. The same value is recorded in the System Log as &#x60;authenticationContext.deviceSessionId&#x60;. Use that event to correlate this session with its events. | 
**UserId** | **string** | The unique identifier of the user that this session belongs to | 
**Links** | Pointer to [**DeviceSessionLinks**](DeviceSessionLinks.md) |  | [optional] 

## Methods

### NewDeviceSession

`func NewDeviceSession(factors []DeviceSessionFactor, id string, userId string, ) *DeviceSession`

NewDeviceSession instantiates a new DeviceSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceSessionWithDefaults

`func NewDeviceSessionWithDefaults() *DeviceSession`

NewDeviceSessionWithDefaults instantiates a new DeviceSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactors

`func (o *DeviceSession) GetFactors() []DeviceSessionFactor`

GetFactors returns the Factors field if non-nil, zero value otherwise.

### GetFactorsOk

`func (o *DeviceSession) GetFactorsOk() (*[]DeviceSessionFactor, bool)`

GetFactorsOk returns a tuple with the Factors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactors

`func (o *DeviceSession) SetFactors(v []DeviceSessionFactor)`

SetFactors sets Factors field to given value.


### GetId

`func (o *DeviceSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceSession) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *DeviceSession) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DeviceSession) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DeviceSession) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetLinks

`func (o *DeviceSession) GetLinks() DeviceSessionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceSession) GetLinksOk() (*DeviceSessionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceSession) SetLinks(v DeviceSessionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeviceSession) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


