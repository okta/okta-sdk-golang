# IdpPolicyRuleActionProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | IdP types of &#x60;OKTA&#x60;, &#x60;AgentlessDSSO&#x60;, and &#x60;IWA&#x60; don&#39;t require an ID. | [optional] 
**Name** | Pointer to **string** | Provider &#x60;name&#x60; in Okta. Optional. Supported in &#x60;IDENTITY ENGINE&#x60;. | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewIdpPolicyRuleActionProvider

`func NewIdpPolicyRuleActionProvider() *IdpPolicyRuleActionProvider`

NewIdpPolicyRuleActionProvider instantiates a new IdpPolicyRuleActionProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpPolicyRuleActionProviderWithDefaults

`func NewIdpPolicyRuleActionProviderWithDefaults() *IdpPolicyRuleActionProvider`

NewIdpPolicyRuleActionProviderWithDefaults instantiates a new IdpPolicyRuleActionProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdpPolicyRuleActionProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdpPolicyRuleActionProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdpPolicyRuleActionProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdpPolicyRuleActionProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdpPolicyRuleActionProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpPolicyRuleActionProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpPolicyRuleActionProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdpPolicyRuleActionProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IdpPolicyRuleActionProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdpPolicyRuleActionProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdpPolicyRuleActionProvider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdpPolicyRuleActionProvider) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


