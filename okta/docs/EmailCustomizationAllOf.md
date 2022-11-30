# EmailCustomizationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The UTC time at which this email customization was created. | [optional] [readonly] 
**Id** | Pointer to **string** | A unique identifier for this email customization. | [optional] [readonly] 
**IsDefault** | Pointer to **bool** | Whether this is the default customization for the email template. Each customized email template must have exactly one default customization. Defaults to &#x60;true&#x60; for the first customization and &#x60;false&#x60; thereafter. | [optional] 
**Language** | **string** | The language specified as an [IETF BCP 47 language tag](https://datatracker.ietf.org/doc/html/rfc5646). | 
**LastUpdated** | Pointer to **time.Time** | The UTC time at which this email customization was last updated. | [optional] [readonly] 
**Links** | Pointer to [**EmailCustomizationAllOfLinks**](EmailCustomizationAllOfLinks.md) |  | [optional] 

## Methods

### NewEmailCustomizationAllOf

`func NewEmailCustomizationAllOf(language string, ) *EmailCustomizationAllOf`

NewEmailCustomizationAllOf instantiates a new EmailCustomizationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailCustomizationAllOfWithDefaults

`func NewEmailCustomizationAllOfWithDefaults() *EmailCustomizationAllOf`

NewEmailCustomizationAllOfWithDefaults instantiates a new EmailCustomizationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *EmailCustomizationAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EmailCustomizationAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EmailCustomizationAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EmailCustomizationAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *EmailCustomizationAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailCustomizationAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailCustomizationAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailCustomizationAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *EmailCustomizationAllOf) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *EmailCustomizationAllOf) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *EmailCustomizationAllOf) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *EmailCustomizationAllOf) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLanguage

`func (o *EmailCustomizationAllOf) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *EmailCustomizationAllOf) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *EmailCustomizationAllOf) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLastUpdated

`func (o *EmailCustomizationAllOf) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EmailCustomizationAllOf) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EmailCustomizationAllOf) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *EmailCustomizationAllOf) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *EmailCustomizationAllOf) GetLinks() EmailCustomizationAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EmailCustomizationAllOf) GetLinksOk() (*EmailCustomizationAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EmailCustomizationAllOf) SetLinks(v EmailCustomizationAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EmailCustomizationAllOf) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


