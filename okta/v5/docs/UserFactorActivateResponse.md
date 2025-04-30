# UserFactorActivateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorType** | Pointer to **string** | Type of the Factor | [optional] 
**Links** | Pointer to [**UserFactorActivateResponseLinks**](UserFactorActivateResponseLinks.md) |  | [optional] 

## Methods

### NewUserFactorActivateResponse

`func NewUserFactorActivateResponse() *UserFactorActivateResponse`

NewUserFactorActivateResponse instantiates a new UserFactorActivateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorActivateResponseWithDefaults

`func NewUserFactorActivateResponseWithDefaults() *UserFactorActivateResponse`

NewUserFactorActivateResponseWithDefaults instantiates a new UserFactorActivateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorType

`func (o *UserFactorActivateResponse) GetFactorType() string`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *UserFactorActivateResponse) GetFactorTypeOk() (*string, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *UserFactorActivateResponse) SetFactorType(v string)`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *UserFactorActivateResponse) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### GetLinks

`func (o *UserFactorActivateResponse) GetLinks() UserFactorActivateResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserFactorActivateResponse) GetLinksOk() (*UserFactorActivateResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserFactorActivateResponse) SetLinks(v UserFactorActivateResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserFactorActivateResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


