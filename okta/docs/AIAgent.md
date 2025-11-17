# AIAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | The ID of the connected app for the AI agent | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when the AI agent was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique ID for the AI agent | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the AI agent was updated | [optional] [readonly] 
**Profile** | Pointer to [**AIAgentProfile**](AIAgentProfile.md) |  | [optional] 
**Status** | Pointer to **string** | When an AI agent is created, it&#39;s in the STAGED status.  After credentials and owners are associated with the agent, it can be set to the ACTIVE status. | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewAIAgent

`func NewAIAgent() *AIAgent`

NewAIAgent instantiates a new AIAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIAgentWithDefaults

`func NewAIAgentWithDefaults() *AIAgent`

NewAIAgentWithDefaults instantiates a new AIAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *AIAgent) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AIAgent) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AIAgent) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AIAgent) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetCreated

`func (o *AIAgent) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AIAgent) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AIAgent) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AIAgent) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *AIAgent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AIAgent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AIAgent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AIAgent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AIAgent) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AIAgent) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AIAgent) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AIAgent) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProfile

`func (o *AIAgent) GetProfile() AIAgentProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AIAgent) GetProfileOk() (*AIAgentProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AIAgent) SetProfile(v AIAgentProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AIAgent) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetStatus

`func (o *AIAgent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AIAgent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AIAgent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AIAgent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *AIAgent) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AIAgent) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AIAgent) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AIAgent) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


