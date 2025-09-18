# AccessPolicyAllOfEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **string** | The resource that this policy controls. For the [Okta account management policy](https://developer.okta.com/docs/guides/okta-account-management-policy/main/#example-response), &#x60;END_USER_ACCOUNT_MANAGEMENT&#x60; is returned. For other policies, &#x60;APP&#x60; is returned. | [optional] 

## Methods

### NewAccessPolicyAllOfEmbedded

`func NewAccessPolicyAllOfEmbedded() *AccessPolicyAllOfEmbedded`

NewAccessPolicyAllOfEmbedded instantiates a new AccessPolicyAllOfEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyAllOfEmbeddedWithDefaults

`func NewAccessPolicyAllOfEmbeddedWithDefaults() *AccessPolicyAllOfEmbedded`

NewAccessPolicyAllOfEmbeddedWithDefaults instantiates a new AccessPolicyAllOfEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *AccessPolicyAllOfEmbedded) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AccessPolicyAllOfEmbedded) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AccessPolicyAllOfEmbedded) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AccessPolicyAllOfEmbedded) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


