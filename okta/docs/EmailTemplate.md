# EmailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this email template. | [optional] [readonly] 
**Embedded** | Pointer to [**EmailTemplateEmbedded**](EmailTemplateEmbedded.md) |  | [optional] 
**Links** | Pointer to [**EmailTemplateLinks**](EmailTemplateLinks.md) |  | [optional] 

## Methods

### NewEmailTemplate

`func NewEmailTemplate() *EmailTemplate`

NewEmailTemplate instantiates a new EmailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTemplateWithDefaults

`func NewEmailTemplateWithDefaults() *EmailTemplate`

NewEmailTemplateWithDefaults instantiates a new EmailTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmbedded

`func (o *EmailTemplate) GetEmbedded() EmailTemplateEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *EmailTemplate) GetEmbeddedOk() (*EmailTemplateEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *EmailTemplate) SetEmbedded(v EmailTemplateEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *EmailTemplate) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *EmailTemplate) GetLinks() EmailTemplateLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EmailTemplate) GetLinksOk() (*EmailTemplateLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EmailTemplate) SetLinks(v EmailTemplateLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EmailTemplate) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


