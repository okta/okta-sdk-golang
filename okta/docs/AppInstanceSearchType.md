# AppInstanceSearchType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Number of app instances of this type in the org. For admins with scoped app access, the count only includes apps they can view. | 
**Type** | **string** | App Instance type refers to the integration type of the instance. | 

## Methods

### NewAppInstanceSearchType

`func NewAppInstanceSearchType(count int32, type_ string, ) *AppInstanceSearchType`

NewAppInstanceSearchType instantiates a new AppInstanceSearchType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInstanceSearchTypeWithDefaults

`func NewAppInstanceSearchTypeWithDefaults() *AppInstanceSearchType`

NewAppInstanceSearchTypeWithDefaults instantiates a new AppInstanceSearchType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AppInstanceSearchType) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AppInstanceSearchType) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AppInstanceSearchType) SetCount(v int32)`

SetCount sets Count field to given value.


### GetType

`func (o *AppInstanceSearchType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppInstanceSearchType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppInstanceSearchType) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


