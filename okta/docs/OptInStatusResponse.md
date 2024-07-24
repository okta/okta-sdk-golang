# OptInStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptInStatus** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**OptInStatusResponseLinks**](OptInStatusResponseLinks.md) |  | [optional] 

## Methods

### NewOptInStatusResponse

`func NewOptInStatusResponse() *OptInStatusResponse`

NewOptInStatusResponse instantiates a new OptInStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptInStatusResponseWithDefaults

`func NewOptInStatusResponseWithDefaults() *OptInStatusResponse`

NewOptInStatusResponseWithDefaults instantiates a new OptInStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptInStatus

`func (o *OptInStatusResponse) GetOptInStatus() string`

GetOptInStatus returns the OptInStatus field if non-nil, zero value otherwise.

### GetOptInStatusOk

`func (o *OptInStatusResponse) GetOptInStatusOk() (*string, bool)`

GetOptInStatusOk returns a tuple with the OptInStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptInStatus

`func (o *OptInStatusResponse) SetOptInStatus(v string)`

SetOptInStatus sets OptInStatus field to given value.

### HasOptInStatus

`func (o *OptInStatusResponse) HasOptInStatus() bool`

HasOptInStatus returns a boolean if a field has been set.

### GetLinks

`func (o *OptInStatusResponse) GetLinks() OptInStatusResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OptInStatusResponse) GetLinksOk() (*OptInStatusResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OptInStatusResponse) SetLinks(v OptInStatusResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OptInStatusResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


