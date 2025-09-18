# RegistrationResponseErrorErrorCausesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorSummary** | Pointer to **string** | Human-readable summary of the error. | [optional] 
**Reason** | Pointer to **string** | A brief, enum-like string that indicates the nature of the error. For example, &#x60;UNIQUE_CONSTRAINT&#x60; for a property uniqueness violation. | [optional] 
**LocationType** | Pointer to **string** | Where in the request the error was found (&#x60;body&#x60;, &#x60;header&#x60;, &#x60;url&#x60;, or &#x60;query&#x60;). | [optional] 
**Location** | Pointer to **string** | The valid JSON path to the location of the error. For example, if there was an error in the user&#39;s &#x60;login&#x60; field, the &#x60;location&#x60; might be &#x60;data.userProfile.login&#x60;. | [optional] 
**Domain** | Pointer to **string** | Indicates the source of the error. If the error was in the user&#39;s profile, for example, you might use &#x60;end-user&#x60;. If the error occurred in the external service, you might use &#x60;external-service&#x60;. | [optional] 

## Methods

### NewRegistrationResponseErrorErrorCausesInner

`func NewRegistrationResponseErrorErrorCausesInner() *RegistrationResponseErrorErrorCausesInner`

NewRegistrationResponseErrorErrorCausesInner instantiates a new RegistrationResponseErrorErrorCausesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationResponseErrorErrorCausesInnerWithDefaults

`func NewRegistrationResponseErrorErrorCausesInnerWithDefaults() *RegistrationResponseErrorErrorCausesInner`

NewRegistrationResponseErrorErrorCausesInnerWithDefaults instantiates a new RegistrationResponseErrorErrorCausesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorSummary

`func (o *RegistrationResponseErrorErrorCausesInner) GetErrorSummary() string`

GetErrorSummary returns the ErrorSummary field if non-nil, zero value otherwise.

### GetErrorSummaryOk

`func (o *RegistrationResponseErrorErrorCausesInner) GetErrorSummaryOk() (*string, bool)`

GetErrorSummaryOk returns a tuple with the ErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSummary

`func (o *RegistrationResponseErrorErrorCausesInner) SetErrorSummary(v string)`

SetErrorSummary sets ErrorSummary field to given value.

### HasErrorSummary

`func (o *RegistrationResponseErrorErrorCausesInner) HasErrorSummary() bool`

HasErrorSummary returns a boolean if a field has been set.

### GetReason

`func (o *RegistrationResponseErrorErrorCausesInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RegistrationResponseErrorErrorCausesInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RegistrationResponseErrorErrorCausesInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RegistrationResponseErrorErrorCausesInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetLocationType

`func (o *RegistrationResponseErrorErrorCausesInner) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *RegistrationResponseErrorErrorCausesInner) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *RegistrationResponseErrorErrorCausesInner) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *RegistrationResponseErrorErrorCausesInner) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetLocation

`func (o *RegistrationResponseErrorErrorCausesInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *RegistrationResponseErrorErrorCausesInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *RegistrationResponseErrorErrorCausesInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *RegistrationResponseErrorErrorCausesInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDomain

`func (o *RegistrationResponseErrorErrorCausesInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RegistrationResponseErrorErrorCausesInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RegistrationResponseErrorErrorCausesInner) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *RegistrationResponseErrorErrorCausesInner) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


