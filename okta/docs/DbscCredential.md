# DbscCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | **string** | Cookie attributes | [readonly] 
**Name** | **string** | Cookie name | [readonly] 
**Type** | **string** | Credential type | [readonly] 

## Methods

### NewDbscCredential

`func NewDbscCredential(attributes string, name string, type_ string, ) *DbscCredential`

NewDbscCredential instantiates a new DbscCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbscCredentialWithDefaults

`func NewDbscCredentialWithDefaults() *DbscCredential`

NewDbscCredentialWithDefaults instantiates a new DbscCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *DbscCredential) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DbscCredential) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DbscCredential) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.


### GetName

`func (o *DbscCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbscCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbscCredential) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DbscCredential) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DbscCredential) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DbscCredential) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


