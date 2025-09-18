# WellKnownURIObjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Representation** | Pointer to **map[string]interface{}** | The well-known URI content in JSON format | [optional] 
**Links** | Pointer to [**WellKnownURIArrayResponseLinks**](WellKnownURIArrayResponseLinks.md) |  | [optional] 

## Methods

### NewWellKnownURIObjectResponse

`func NewWellKnownURIObjectResponse() *WellKnownURIObjectResponse`

NewWellKnownURIObjectResponse instantiates a new WellKnownURIObjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWellKnownURIObjectResponseWithDefaults

`func NewWellKnownURIObjectResponseWithDefaults() *WellKnownURIObjectResponse`

NewWellKnownURIObjectResponseWithDefaults instantiates a new WellKnownURIObjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepresentation

`func (o *WellKnownURIObjectResponse) GetRepresentation() map[string]interface{}`

GetRepresentation returns the Representation field if non-nil, zero value otherwise.

### GetRepresentationOk

`func (o *WellKnownURIObjectResponse) GetRepresentationOk() (*map[string]interface{}, bool)`

GetRepresentationOk returns a tuple with the Representation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentation

`func (o *WellKnownURIObjectResponse) SetRepresentation(v map[string]interface{})`

SetRepresentation sets Representation field to given value.

### HasRepresentation

`func (o *WellKnownURIObjectResponse) HasRepresentation() bool`

HasRepresentation returns a boolean if a field has been set.

### GetLinks

`func (o *WellKnownURIObjectResponse) GetLinks() WellKnownURIArrayResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WellKnownURIObjectResponse) GetLinksOk() (*WellKnownURIArrayResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WellKnownURIObjectResponse) SetLinks(v WellKnownURIArrayResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WellKnownURIObjectResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


