# RiskProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**RiskProviderAction**](RiskProviderAction.md) |  | [optional] [default to RISKPROVIDERACTION_LOG_ONLY]
**ClientId** | **string** | The ID of the [OAuth service app](https://developer.okta.com/docs/guides/implement-oauth-for-okta-serviceapp/main/#create-a-service-app-and-grant-scopes) that is used to send risk events to Okta | 
**Created** | Pointer to **time.Time** | Timestamp when the risk provider was created | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the risk provider | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the risk provider was last updated | [optional] [readonly] 
**Name** | **string** | Name of the risk provider | 
**Links** | Pointer to [**ApiTokenLink**](ApiTokenLink.md) |  | [optional] 

## Methods

### NewRiskProvider

`func NewRiskProvider(clientId string, name string, ) *RiskProvider`

NewRiskProvider instantiates a new RiskProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskProviderWithDefaults

`func NewRiskProviderWithDefaults() *RiskProvider`

NewRiskProviderWithDefaults instantiates a new RiskProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RiskProvider) GetAction() RiskProviderAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RiskProvider) GetActionOk() (*RiskProviderAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RiskProvider) SetAction(v RiskProviderAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *RiskProvider) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetClientId

`func (o *RiskProvider) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *RiskProvider) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *RiskProvider) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCreated

`func (o *RiskProvider) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RiskProvider) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RiskProvider) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RiskProvider) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *RiskProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *RiskProvider) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RiskProvider) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RiskProvider) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *RiskProvider) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *RiskProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskProvider) SetName(v string)`

SetName sets Name field to given value.


### GetLinks

`func (o *RiskProvider) GetLinks() ApiTokenLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RiskProvider) GetLinksOk() (*ApiTokenLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RiskProvider) SetLinks(v ApiTokenLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RiskProvider) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


