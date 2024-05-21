# SlackApplicationSettingsApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The Slack app domain name | 
**UserEmailValue** | Pointer to **string** | The &#x60;User.Email&#x60; attribute value | [optional] 

## Methods

### NewSlackApplicationSettingsApplication

`func NewSlackApplicationSettingsApplication(domain string, ) *SlackApplicationSettingsApplication`

NewSlackApplicationSettingsApplication instantiates a new SlackApplicationSettingsApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackApplicationSettingsApplicationWithDefaults

`func NewSlackApplicationSettingsApplicationWithDefaults() *SlackApplicationSettingsApplication`

NewSlackApplicationSettingsApplicationWithDefaults instantiates a new SlackApplicationSettingsApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *SlackApplicationSettingsApplication) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SlackApplicationSettingsApplication) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SlackApplicationSettingsApplication) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetUserEmailValue

`func (o *SlackApplicationSettingsApplication) GetUserEmailValue() string`

GetUserEmailValue returns the UserEmailValue field if non-nil, zero value otherwise.

### GetUserEmailValueOk

`func (o *SlackApplicationSettingsApplication) GetUserEmailValueOk() (*string, bool)`

GetUserEmailValueOk returns a tuple with the UserEmailValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmailValue

`func (o *SlackApplicationSettingsApplication) SetUserEmailValue(v string)`

SetUserEmailValue sets UserEmailValue field to given value.

### HasUserEmailValue

`func (o *SlackApplicationSettingsApplication) HasUserEmailValue() bool`

HasUserEmailValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


