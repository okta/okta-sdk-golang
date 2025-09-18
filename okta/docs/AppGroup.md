# AppGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The external ID of the app group whose members might be privileged app users | 
**Name** | **string** | The name of the app group whose members might be privileged app users | 

## Methods

### NewAppGroup

`func NewAppGroup(externalId string, name string, ) *AppGroup`

NewAppGroup instantiates a new AppGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppGroupWithDefaults

`func NewAppGroupWithDefaults() *AppGroup`

NewAppGroupWithDefaults instantiates a new AppGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *AppGroup) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AppGroup) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AppGroup) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetName

`func (o *AppGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppGroup) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


