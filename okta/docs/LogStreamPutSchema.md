# LogStreamPutSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name for the Log Stream object | 
**Type** | **string** | Specifies the streaming provider used  Supported providers:   * &#x60;aws_eventbridge&#x60; ([AWS EventBridge](https://aws.amazon.com/eventbridge))   * &#x60;splunk_cloud_logstreaming&#x60; ([Splunk Cloud](https://www.splunk.com/en_us/software/splunk-cloud-platform.html))  Select the provider type to see provider-specific configurations in the &#x60;settings&#x60; property: | 

## Methods

### NewLogStreamPutSchema

`func NewLogStreamPutSchema(name string, type_ string, ) *LogStreamPutSchema`

NewLogStreamPutSchema instantiates a new LogStreamPutSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogStreamPutSchemaWithDefaults

`func NewLogStreamPutSchemaWithDefaults() *LogStreamPutSchema`

NewLogStreamPutSchemaWithDefaults instantiates a new LogStreamPutSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LogStreamPutSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogStreamPutSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogStreamPutSchema) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *LogStreamPutSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogStreamPutSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogStreamPutSchema) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


