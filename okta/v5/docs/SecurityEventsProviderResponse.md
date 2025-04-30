# SecurityEventsProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of this instance | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Security Events Provider instance | [optional] 
**Settings** | Pointer to [**SecurityEventsProviderSettingsResponse**](SecurityEventsProviderSettingsResponse.md) |  | [optional] 
**Status** | Pointer to **string** | Indicates whether the Security Events Provider is active or not | [optional] [readonly] 
**Type** | Pointer to **string** | The application type of the Security Events Provider | [optional] 
**Links** | Pointer to [**LinksSelfAndLifecycle**](LinksSelfAndLifecycle.md) |  | [optional] 

## Methods

### NewSecurityEventsProviderResponse

`func NewSecurityEventsProviderResponse() *SecurityEventsProviderResponse`

NewSecurityEventsProviderResponse instantiates a new SecurityEventsProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityEventsProviderResponseWithDefaults

`func NewSecurityEventsProviderResponseWithDefaults() *SecurityEventsProviderResponse`

NewSecurityEventsProviderResponseWithDefaults instantiates a new SecurityEventsProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityEventsProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityEventsProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityEventsProviderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityEventsProviderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SecurityEventsProviderResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityEventsProviderResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityEventsProviderResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityEventsProviderResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *SecurityEventsProviderResponse) GetSettings() SecurityEventsProviderSettingsResponse`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SecurityEventsProviderResponse) GetSettingsOk() (*SecurityEventsProviderSettingsResponse, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SecurityEventsProviderResponse) SetSettings(v SecurityEventsProviderSettingsResponse)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SecurityEventsProviderResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityEventsProviderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityEventsProviderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityEventsProviderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityEventsProviderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *SecurityEventsProviderResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityEventsProviderResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityEventsProviderResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecurityEventsProviderResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *SecurityEventsProviderResponse) GetLinks() LinksSelfAndLifecycle`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecurityEventsProviderResponse) GetLinksOk() (*LinksSelfAndLifecycle, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecurityEventsProviderResponse) SetLinks(v LinksSelfAndLifecycle)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SecurityEventsProviderResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


