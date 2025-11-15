# UpdateAIAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | The ID of the connected app for the AI agent | [optional] 
**Profile** | Pointer to [**AIAgentProfile**](AIAgentProfile.md) |  | [optional] 

## Methods

### NewUpdateAIAgentRequest

`func NewUpdateAIAgentRequest() *UpdateAIAgentRequest`

NewUpdateAIAgentRequest instantiates a new UpdateAIAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAIAgentRequestWithDefaults

`func NewUpdateAIAgentRequestWithDefaults() *UpdateAIAgentRequest`

NewUpdateAIAgentRequestWithDefaults instantiates a new UpdateAIAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *UpdateAIAgentRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *UpdateAIAgentRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *UpdateAIAgentRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *UpdateAIAgentRequest) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetProfile

`func (o *UpdateAIAgentRequest) GetProfile() AIAgentProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UpdateAIAgentRequest) GetProfileOk() (*AIAgentProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UpdateAIAgentRequest) SetProfile(v AIAgentProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UpdateAIAgentRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


