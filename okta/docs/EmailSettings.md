# EmailSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipients** | **string** |  | 

## Methods

### NewEmailSettings

`func NewEmailSettings(recipients string, ) *EmailSettings`

NewEmailSettings instantiates a new EmailSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSettingsWithDefaults

`func NewEmailSettingsWithDefaults() *EmailSettings`

NewEmailSettingsWithDefaults instantiates a new EmailSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipients

`func (o *EmailSettings) GetRecipients() string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *EmailSettings) GetRecipientsOk() (*string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *EmailSettings) SetRecipients(v string)`

SetRecipients sets Recipients field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


