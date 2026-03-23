# IntegrationCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the integration capability | [optional] 
**HelpUrl** | Pointer to **string** | URL to the help documentation for the integration capability | [optional] 
**Id** | Pointer to **string** | Unique identifier for the integration capability | [optional] 
**Name** | Pointer to **string** | Name of the integration capability | [optional] 
**ReleaseDate** | Pointer to **time.Time** | The date when the integration capability was released | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the integration capability | [optional] 

## Methods

### NewIntegrationCapability

`func NewIntegrationCapability() *IntegrationCapability`

NewIntegrationCapability instantiates a new IntegrationCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationCapabilityWithDefaults

`func NewIntegrationCapabilityWithDefaults() *IntegrationCapability`

NewIntegrationCapabilityWithDefaults instantiates a new IntegrationCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IntegrationCapability) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationCapability) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationCapability) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntegrationCapability) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHelpUrl

`func (o *IntegrationCapability) GetHelpUrl() string`

GetHelpUrl returns the HelpUrl field if non-nil, zero value otherwise.

### GetHelpUrlOk

`func (o *IntegrationCapability) GetHelpUrlOk() (*string, bool)`

GetHelpUrlOk returns a tuple with the HelpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpUrl

`func (o *IntegrationCapability) SetHelpUrl(v string)`

SetHelpUrl sets HelpUrl field to given value.

### HasHelpUrl

`func (o *IntegrationCapability) HasHelpUrl() bool`

HasHelpUrl returns a boolean if a field has been set.

### GetId

`func (o *IntegrationCapability) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationCapability) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationCapability) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationCapability) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IntegrationCapability) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationCapability) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationCapability) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationCapability) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReleaseDate

`func (o *IntegrationCapability) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *IntegrationCapability) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *IntegrationCapability) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *IntegrationCapability) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetStatus

`func (o *IntegrationCapability) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntegrationCapability) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntegrationCapability) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IntegrationCapability) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


