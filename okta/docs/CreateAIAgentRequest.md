# CreateAIAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | The ID of the connected app for the AI agent | [optional] 
**Profile** | Pointer to [**AIAgentProfile**](AIAgentProfile.md) |  | [optional] 

## Methods

### NewCreateAIAgentRequest

`func NewCreateAIAgentRequest() *CreateAIAgentRequest`

NewCreateAIAgentRequest instantiates a new CreateAIAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAIAgentRequestWithDefaults

`func NewCreateAIAgentRequestWithDefaults() *CreateAIAgentRequest`

NewCreateAIAgentRequestWithDefaults instantiates a new CreateAIAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *CreateAIAgentRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CreateAIAgentRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CreateAIAgentRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *CreateAIAgentRequest) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetProfile

`func (o *CreateAIAgentRequest) GetProfile() AIAgentProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CreateAIAgentRequest) GetProfileOk() (*AIAgentProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CreateAIAgentRequest) SetProfile(v AIAgentProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *CreateAIAgentRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


