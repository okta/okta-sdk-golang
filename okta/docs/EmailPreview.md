# EmailPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | The email&#39;s HTML body. | [optional] [readonly] 
**Subject** | Pointer to **string** | The email&#39;s subject. | [optional] [readonly] 
**Links** | Pointer to [**EmailPreviewLinks**](EmailPreviewLinks.md) |  | [optional] 

## Methods

### NewEmailPreview

`func NewEmailPreview() *EmailPreview`

NewEmailPreview instantiates a new EmailPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailPreviewWithDefaults

`func NewEmailPreviewWithDefaults() *EmailPreview`

NewEmailPreviewWithDefaults instantiates a new EmailPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *EmailPreview) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EmailPreview) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EmailPreview) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *EmailPreview) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetSubject

`func (o *EmailPreview) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailPreview) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailPreview) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailPreview) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetLinks

`func (o *EmailPreview) GetLinks() EmailPreviewLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EmailPreview) GetLinksOk() (*EmailPreviewLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EmailPreview) SetLinks(v EmailPreviewLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EmailPreview) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


