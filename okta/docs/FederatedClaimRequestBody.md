# FederatedClaimRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **string** | The Okta Expression Language expression to be evaluated at runtime | [optional] 
**Name** | Pointer to **string** | The name of the claim to be used in the produced token | [optional] 

## Methods

### NewFederatedClaimRequestBody

`func NewFederatedClaimRequestBody() *FederatedClaimRequestBody`

NewFederatedClaimRequestBody instantiates a new FederatedClaimRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederatedClaimRequestBodyWithDefaults

`func NewFederatedClaimRequestBodyWithDefaults() *FederatedClaimRequestBody`

NewFederatedClaimRequestBodyWithDefaults instantiates a new FederatedClaimRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *FederatedClaimRequestBody) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *FederatedClaimRequestBody) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *FederatedClaimRequestBody) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *FederatedClaimRequestBody) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetName

`func (o *FederatedClaimRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FederatedClaimRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FederatedClaimRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FederatedClaimRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


