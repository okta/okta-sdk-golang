# EmailDefaultContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | The email&#39;s HTML body. May contain [variable references](https://velocity.apache.org/engine/1.7/user-guide.html#references). | 
**Subject** | **string** | The email&#39;s subject. May contain [variable references](https://velocity.apache.org/engine/1.7/user-guide.html#references). | 
**Links** | Pointer to [**EmailDefaultContentAllOfLinks**](EmailDefaultContentAllOfLinks.md) |  | [optional] 

## Methods

### NewEmailDefaultContent

`func NewEmailDefaultContent(body string, subject string, ) *EmailDefaultContent`

NewEmailDefaultContent instantiates a new EmailDefaultContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDefaultContentWithDefaults

`func NewEmailDefaultContentWithDefaults() *EmailDefaultContent`

NewEmailDefaultContentWithDefaults instantiates a new EmailDefaultContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *EmailDefaultContent) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EmailDefaultContent) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EmailDefaultContent) SetBody(v string)`

SetBody sets Body field to given value.


### GetSubject

`func (o *EmailDefaultContent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailDefaultContent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailDefaultContent) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetLinks

`func (o *EmailDefaultContent) GetLinks() EmailDefaultContentAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EmailDefaultContent) GetLinksOk() (*EmailDefaultContentAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EmailDefaultContent) SetLinks(v EmailDefaultContentAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EmailDefaultContent) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


