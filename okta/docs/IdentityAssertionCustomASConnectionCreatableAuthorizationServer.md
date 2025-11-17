# IdentityAssertionCustomASConnectionCreatableAuthorizationServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orn** | **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the authorization server | 
**ResourceIndicator** | Pointer to **string** | Optional resource indicator (audience) used when requesting tokens | [optional] 

## Methods

### NewIdentityAssertionCustomASConnectionCreatableAuthorizationServer

`func NewIdentityAssertionCustomASConnectionCreatableAuthorizationServer(orn string, ) *IdentityAssertionCustomASConnectionCreatableAuthorizationServer`

NewIdentityAssertionCustomASConnectionCreatableAuthorizationServer instantiates a new IdentityAssertionCustomASConnectionCreatableAuthorizationServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAssertionCustomASConnectionCreatableAuthorizationServerWithDefaults

`func NewIdentityAssertionCustomASConnectionCreatableAuthorizationServerWithDefaults() *IdentityAssertionCustomASConnectionCreatableAuthorizationServer`

NewIdentityAssertionCustomASConnectionCreatableAuthorizationServerWithDefaults instantiates a new IdentityAssertionCustomASConnectionCreatableAuthorizationServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrn

`func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetResourceIndicator

`func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) GetResourceIndicator() string`

GetResourceIndicator returns the ResourceIndicator field if non-nil, zero value otherwise.

### GetResourceIndicatorOk

`func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) GetResourceIndicatorOk() (*string, bool)`

GetResourceIndicatorOk returns a tuple with the ResourceIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIndicator

`func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) SetResourceIndicator(v string)`

SetResourceIndicator sets ResourceIndicator field to given value.

### HasResourceIndicator

`func (o *IdentityAssertionCustomASConnectionCreatableAuthorizationServer) HasResourceIndicator() bool`

HasResourceIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


