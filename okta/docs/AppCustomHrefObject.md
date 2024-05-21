# AppCustomHrefObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hints** | Pointer to [**AppCustomHrefObjectHints**](AppCustomHrefObjectHints.md) |  | [optional] 
**Href** | **string** | Link URI | 
**Title** | Pointer to **string** | Link name | [optional] 
**Type** | Pointer to **string** | The media type of the link. If omitted, it is implicitly &#x60;application/json&#x60;. | [optional] 

## Methods

### NewAppCustomHrefObject

`func NewAppCustomHrefObject(href string, ) *AppCustomHrefObject`

NewAppCustomHrefObject instantiates a new AppCustomHrefObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCustomHrefObjectWithDefaults

`func NewAppCustomHrefObjectWithDefaults() *AppCustomHrefObject`

NewAppCustomHrefObjectWithDefaults instantiates a new AppCustomHrefObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHints

`func (o *AppCustomHrefObject) GetHints() AppCustomHrefObjectHints`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *AppCustomHrefObject) GetHintsOk() (*AppCustomHrefObjectHints, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *AppCustomHrefObject) SetHints(v AppCustomHrefObjectHints)`

SetHints sets Hints field to given value.

### HasHints

`func (o *AppCustomHrefObject) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetHref

`func (o *AppCustomHrefObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppCustomHrefObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppCustomHrefObject) SetHref(v string)`

SetHref sets Href field to given value.


### GetTitle

`func (o *AppCustomHrefObject) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AppCustomHrefObject) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AppCustomHrefObject) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AppCustomHrefObject) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *AppCustomHrefObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppCustomHrefObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppCustomHrefObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppCustomHrefObject) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


