# EmailTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this email template | [optional] [readonly] 
**Embedded** | Pointer to [**EmailTemplateResponseEmbedded**](EmailTemplateResponseEmbedded.md) |  | [optional] 
**Links** | Pointer to [**EmailTemplateResponseLinks**](EmailTemplateResponseLinks.md) |  | [optional] 

## Methods

### NewEmailTemplateResponse

`func NewEmailTemplateResponse() *EmailTemplateResponse`

NewEmailTemplateResponse instantiates a new EmailTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTemplateResponseWithDefaults

`func NewEmailTemplateResponseWithDefaults() *EmailTemplateResponse`

NewEmailTemplateResponseWithDefaults instantiates a new EmailTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailTemplateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailTemplateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmbedded

`func (o *EmailTemplateResponse) GetEmbedded() EmailTemplateResponseEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *EmailTemplateResponse) GetEmbeddedOk() (*EmailTemplateResponseEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *EmailTemplateResponse) SetEmbedded(v EmailTemplateResponseEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *EmailTemplateResponse) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *EmailTemplateResponse) GetLinks() EmailTemplateResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EmailTemplateResponse) GetLinksOk() (*EmailTemplateResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EmailTemplateResponse) SetLinks(v EmailTemplateResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EmailTemplateResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


