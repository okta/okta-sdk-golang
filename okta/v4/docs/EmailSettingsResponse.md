# EmailSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipients** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**EmailSettingsResponseLinks**](EmailSettingsResponseLinks.md) |  | [optional] 

## Methods

### NewEmailSettingsResponse

`func NewEmailSettingsResponse() *EmailSettingsResponse`

NewEmailSettingsResponse instantiates a new EmailSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSettingsResponseWithDefaults

`func NewEmailSettingsResponseWithDefaults() *EmailSettingsResponse`

NewEmailSettingsResponseWithDefaults instantiates a new EmailSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipients

`func (o *EmailSettingsResponse) GetRecipients() string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *EmailSettingsResponse) GetRecipientsOk() (*string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *EmailSettingsResponse) SetRecipients(v string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *EmailSettingsResponse) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetLinks

`func (o *EmailSettingsResponse) GetLinks() EmailSettingsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EmailSettingsResponse) GetLinksOk() (*EmailSettingsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EmailSettingsResponse) SetLinks(v EmailSettingsResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EmailSettingsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


