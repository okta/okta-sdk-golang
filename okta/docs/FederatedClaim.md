# FederatedClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | Timestamp when the federated claim was created | [optional] [readonly] 
**Expression** | Pointer to **string** | The Okta Expression Language expression to be evaluated at runtime | [optional] 
**Id** | Pointer to **string** | The unique ID of the federated claim | [optional] [readonly] 
**LastUpdated** | Pointer to **string** | Timestamp when the federated claim was updated | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the claim to be used in the produced token | [optional] 

## Methods

### NewFederatedClaim

`func NewFederatedClaim() *FederatedClaim`

NewFederatedClaim instantiates a new FederatedClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederatedClaimWithDefaults

`func NewFederatedClaimWithDefaults() *FederatedClaim`

NewFederatedClaimWithDefaults instantiates a new FederatedClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *FederatedClaim) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FederatedClaim) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FederatedClaim) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FederatedClaim) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpression

`func (o *FederatedClaim) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *FederatedClaim) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *FederatedClaim) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *FederatedClaim) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetId

`func (o *FederatedClaim) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FederatedClaim) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FederatedClaim) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FederatedClaim) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *FederatedClaim) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *FederatedClaim) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *FederatedClaim) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *FederatedClaim) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *FederatedClaim) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FederatedClaim) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FederatedClaim) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FederatedClaim) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


