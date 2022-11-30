# HostedPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**HostedPageType**](HostedPageType.md) |  | 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewHostedPage

`func NewHostedPage(type_ HostedPageType, ) *HostedPage`

NewHostedPage instantiates a new HostedPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedPageWithDefaults

`func NewHostedPageWithDefaults() *HostedPage`

NewHostedPageWithDefaults instantiates a new HostedPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HostedPage) GetType() HostedPageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HostedPage) GetTypeOk() (*HostedPageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HostedPage) SetType(v HostedPageType)`

SetType sets Type field to given value.


### GetUrl

`func (o *HostedPage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HostedPage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HostedPage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HostedPage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


