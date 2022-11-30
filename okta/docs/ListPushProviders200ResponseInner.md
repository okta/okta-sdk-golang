# ListPushProviders200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdatedDate** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | Display name of the push provider | [optional] 
**ProviderType** | Pointer to [**ProviderType**](ProviderType.md) |  | [optional] 
**Links** | Pointer to [**ApiTokenLink**](ApiTokenLink.md) |  | [optional] 
**Configuration** | Pointer to [**FCMConfiguration**](FCMConfiguration.md) |  | [optional] 

## Methods

### NewListPushProviders200ResponseInner

`func NewListPushProviders200ResponseInner() *ListPushProviders200ResponseInner`

NewListPushProviders200ResponseInner instantiates a new ListPushProviders200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPushProviders200ResponseInnerWithDefaults

`func NewListPushProviders200ResponseInnerWithDefaults() *ListPushProviders200ResponseInner`

NewListPushProviders200ResponseInnerWithDefaults instantiates a new ListPushProviders200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListPushProviders200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListPushProviders200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListPushProviders200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListPushProviders200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *ListPushProviders200ResponseInner) GetLastUpdatedDate() string`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *ListPushProviders200ResponseInner) GetLastUpdatedDateOk() (*string, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *ListPushProviders200ResponseInner) SetLastUpdatedDate(v string)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *ListPushProviders200ResponseInner) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetName

`func (o *ListPushProviders200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListPushProviders200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListPushProviders200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListPushProviders200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProviderType

`func (o *ListPushProviders200ResponseInner) GetProviderType() ProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *ListPushProviders200ResponseInner) GetProviderTypeOk() (*ProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *ListPushProviders200ResponseInner) SetProviderType(v ProviderType)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *ListPushProviders200ResponseInner) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetLinks

`func (o *ListPushProviders200ResponseInner) GetLinks() ApiTokenLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListPushProviders200ResponseInner) GetLinksOk() (*ApiTokenLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListPushProviders200ResponseInner) SetLinks(v ApiTokenLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListPushProviders200ResponseInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetConfiguration

`func (o *ListPushProviders200ResponseInner) GetConfiguration() FCMConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ListPushProviders200ResponseInner) GetConfigurationOk() (*FCMConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ListPushProviders200ResponseInner) SetConfiguration(v FCMConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ListPushProviders200ResponseInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


