# Office365Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** | The domain for your Office 365 account | [optional] 

## Methods

### NewOffice365Domain

`func NewOffice365Domain() *Office365Domain`

NewOffice365Domain instantiates a new Office365Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365DomainWithDefaults

`func NewOffice365DomainWithDefaults() *Office365Domain`

NewOffice365DomainWithDefaults instantiates a new Office365Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Office365Domain) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Office365Domain) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Office365Domain) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Office365Domain) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Office365Domain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Office365Domain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Office365Domain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Office365Domain) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


