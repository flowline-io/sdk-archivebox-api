# CLICommandResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Errors** | **[]string** |  | 
**Result** | [**NullableResult**](Result.md) |  | 
**Stdout** | **string** |  | 
**Stderr** | **string** |  | 

## Methods

### NewCLICommandResponseSchema

`func NewCLICommandResponseSchema(success bool, errors []string, result NullableResult, stdout string, stderr string, ) *CLICommandResponseSchema`

NewCLICommandResponseSchema instantiates a new CLICommandResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCLICommandResponseSchemaWithDefaults

`func NewCLICommandResponseSchemaWithDefaults() *CLICommandResponseSchema`

NewCLICommandResponseSchemaWithDefaults instantiates a new CLICommandResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CLICommandResponseSchema) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CLICommandResponseSchema) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CLICommandResponseSchema) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrors

`func (o *CLICommandResponseSchema) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CLICommandResponseSchema) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CLICommandResponseSchema) SetErrors(v []string)`

SetErrors sets Errors field to given value.


### GetResult

`func (o *CLICommandResponseSchema) GetResult() Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CLICommandResponseSchema) GetResultOk() (*Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CLICommandResponseSchema) SetResult(v Result)`

SetResult sets Result field to given value.


### SetResultNil

`func (o *CLICommandResponseSchema) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *CLICommandResponseSchema) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetStdout

`func (o *CLICommandResponseSchema) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *CLICommandResponseSchema) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *CLICommandResponseSchema) SetStdout(v string)`

SetStdout sets Stdout field to given value.


### GetStderr

`func (o *CLICommandResponseSchema) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *CLICommandResponseSchema) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *CLICommandResponseSchema) SetStderr(v string)`

SetStderr sets Stderr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


