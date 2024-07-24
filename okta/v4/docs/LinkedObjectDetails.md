# LinkedObjectDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the &#x60;primary&#x60; or the &#x60;associated&#x60; relationship | [optional] 
**Name** | **string** | API name of the &#x60;primary&#x60; or the &#x60;associated&#x60; link | 
**Title** | **string** | Display name of the &#x60;primary&#x60; or the &#x60;associated&#x60; link | 
**Type** | **string** | The object type for this relationship | 

## Methods

### NewLinkedObjectDetails

`func NewLinkedObjectDetails(name string, title string, type_ string, ) *LinkedObjectDetails`

NewLinkedObjectDetails instantiates a new LinkedObjectDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedObjectDetailsWithDefaults

`func NewLinkedObjectDetailsWithDefaults() *LinkedObjectDetails`

NewLinkedObjectDetailsWithDefaults instantiates a new LinkedObjectDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LinkedObjectDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LinkedObjectDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LinkedObjectDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LinkedObjectDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *LinkedObjectDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinkedObjectDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinkedObjectDetails) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *LinkedObjectDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LinkedObjectDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LinkedObjectDetails) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *LinkedObjectDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedObjectDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedObjectDetails) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


