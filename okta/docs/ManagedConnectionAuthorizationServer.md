# ManagedConnectionAuthorizationServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuerUrl** | **string** | Issuer URL for the authorization server | 
**ResourceIndicator** | Pointer to **string** | Audience value to use when requesting tokens | [optional] 

## Methods

### NewManagedConnectionAuthorizationServer

`func NewManagedConnectionAuthorizationServer(issuerUrl string, ) *ManagedConnectionAuthorizationServer`

NewManagedConnectionAuthorizationServer instantiates a new ManagedConnectionAuthorizationServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedConnectionAuthorizationServerWithDefaults

`func NewManagedConnectionAuthorizationServerWithDefaults() *ManagedConnectionAuthorizationServer`

NewManagedConnectionAuthorizationServerWithDefaults instantiates a new ManagedConnectionAuthorizationServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuerUrl

`func (o *ManagedConnectionAuthorizationServer) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *ManagedConnectionAuthorizationServer) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *ManagedConnectionAuthorizationServer) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.


### GetResourceIndicator

`func (o *ManagedConnectionAuthorizationServer) GetResourceIndicator() string`

GetResourceIndicator returns the ResourceIndicator field if non-nil, zero value otherwise.

### GetResourceIndicatorOk

`func (o *ManagedConnectionAuthorizationServer) GetResourceIndicatorOk() (*string, bool)`

GetResourceIndicatorOk returns a tuple with the ResourceIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIndicator

`func (o *ManagedConnectionAuthorizationServer) SetResourceIndicator(v string)`

SetResourceIndicator sets ResourceIndicator field to given value.

### HasResourceIndicator

`func (o *ManagedConnectionAuthorizationServer) HasResourceIndicator() bool`

HasResourceIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


