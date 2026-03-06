# ValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | JSON path or pointer of the failing field when available. | [optional] 
**Message** | Pointer to **string** | Human-readable validation message. | [optional] 
**Severity** | Pointer to **string** | Validation severity label when provided by validator. | [optional] 

## Methods

### NewValidationResult

`func NewValidationResult() *ValidationResult`

NewValidationResult instantiates a new ValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationResultWithDefaults

`func NewValidationResultWithDefaults() *ValidationResult`

NewValidationResultWithDefaults instantiates a new ValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ValidationResult) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ValidationResult) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ValidationResult) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ValidationResult) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMessage

`func (o *ValidationResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ValidationResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSeverity

`func (o *ValidationResult) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ValidationResult) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ValidationResult) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ValidationResult) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


