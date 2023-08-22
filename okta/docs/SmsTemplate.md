# SmsTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**Translations** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | Pointer to [**SmsTemplateType**](SmsTemplateType.md) |  | [optional] 

## Methods

### NewSmsTemplate

`func NewSmsTemplate() *SmsTemplate`

NewSmsTemplate instantiates a new SmsTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsTemplateWithDefaults

`func NewSmsTemplateWithDefaults() *SmsTemplate`

NewSmsTemplateWithDefaults instantiates a new SmsTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SmsTemplate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SmsTemplate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SmsTemplate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SmsTemplate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *SmsTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmsTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmsTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SmsTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SmsTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SmsTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SmsTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SmsTemplate) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *SmsTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmsTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmsTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SmsTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplate

`func (o *SmsTemplate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SmsTemplate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SmsTemplate) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *SmsTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTranslations

`func (o *SmsTemplate) GetTranslations() map[string]interface{}`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *SmsTemplate) GetTranslationsOk() (*map[string]interface{}, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *SmsTemplate) SetTranslations(v map[string]interface{})`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *SmsTemplate) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.

### GetType

`func (o *SmsTemplate) GetType() SmsTemplateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SmsTemplate) GetTypeOk() (*SmsTemplateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SmsTemplate) SetType(v SmsTemplateType)`

SetType sets Type field to given value.

### HasType

`func (o *SmsTemplate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


