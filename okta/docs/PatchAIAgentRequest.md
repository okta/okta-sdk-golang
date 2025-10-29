# PatchAIAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **NullableString** | The ID of the connected app for the AI Agent | [optional] 
**Profile** | Pointer to [**PatchAIAgentProfile**](PatchAIAgentProfile.md) |  | [optional] 

## Methods

### NewPatchAIAgentRequest

`func NewPatchAIAgentRequest() *PatchAIAgentRequest`

NewPatchAIAgentRequest instantiates a new PatchAIAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchAIAgentRequestWithDefaults

`func NewPatchAIAgentRequestWithDefaults() *PatchAIAgentRequest`

NewPatchAIAgentRequestWithDefaults instantiates a new PatchAIAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *PatchAIAgentRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *PatchAIAgentRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *PatchAIAgentRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *PatchAIAgentRequest) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *PatchAIAgentRequest) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *PatchAIAgentRequest) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetProfile

`func (o *PatchAIAgentRequest) GetProfile() PatchAIAgentProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PatchAIAgentRequest) GetProfileOk() (*PatchAIAgentProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PatchAIAgentRequest) SetProfile(v PatchAIAgentProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PatchAIAgentRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


