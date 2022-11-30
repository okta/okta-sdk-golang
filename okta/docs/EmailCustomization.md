# EmailCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | The email&#39;s HTML body. May contain [variable references](https://velocity.apache.org/engine/1.7/user-guide.html#references). | 
**Subject** | **string** | The email&#39;s subject. May contain [variable references](https://velocity.apache.org/engine/1.7/user-guide.html#references). | 
**Created** | Pointer to **time.Time** | The UTC time at which this email customization was created. | [optional] [readonly] 
**Id** | Pointer to **string** | A unique identifier for this email customization. | [optional] [readonly] 
**IsDefault** | Pointer to **bool** | Whether this is the default customization for the email template. Each customized email template must have exactly one default customization. Defaults to &#x60;true&#x60; for the first customization and &#x60;false&#x60; thereafter. | [optional] 
**Language** | **string** | The language specified as an [IETF BCP 47 language tag](https://datatracker.ietf.org/doc/html/rfc5646). | 
**LastUpdated** | Pointer to **time.Time** | The UTC time at which this email customization was last updated. | [optional] [readonly] 
**Links** | Pointer to [**EmailCustomizationAllOfLinks**](EmailCustomizationAllOfLinks.md) |  | [optional] 

## Methods

### NewEmailCustomization

`func NewEmailCustomization(body string, subject string, language string, ) *EmailCustomization`

NewEmailCustomization instantiates a new EmailCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailCustomizationWithDefaults

`func NewEmailCustomizationWithDefaults() *EmailCustomization`

NewEmailCustomizationWithDefaults instantiates a new EmailCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *EmailCustomization) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EmailCustomization) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EmailCustomization) SetBody(v string)`

SetBody sets Body field to given value.


### GetSubject

`func (o *EmailCustomization) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailCustomization) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailCustomization) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetCreated

`func (o *EmailCustomization) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EmailCustomization) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EmailCustomization) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EmailCustomization) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *EmailCustomization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailCustomization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailCustomization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailCustomization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *EmailCustomization) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *EmailCustomization) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *EmailCustomization) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *EmailCustomization) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLanguage

`func (o *EmailCustomization) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *EmailCustomization) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *EmailCustomization) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLastUpdated

`func (o *EmailCustomization) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EmailCustomization) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EmailCustomization) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *EmailCustomization) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *EmailCustomization) GetLinks() EmailCustomizationAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EmailCustomization) GetLinksOk() (*EmailCustomizationAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EmailCustomization) SetLinks(v EmailCustomizationAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EmailCustomization) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


