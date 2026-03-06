# ValidationReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passed** | **bool** | True when no blocking validation issues were found. | 
**Results** | [**[]ValidationResult**](ValidationResult.md) |  | 

## Methods

### NewValidationReport

`func NewValidationReport(passed bool, results []ValidationResult, ) *ValidationReport`

NewValidationReport instantiates a new ValidationReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationReportWithDefaults

`func NewValidationReportWithDefaults() *ValidationReport`

NewValidationReportWithDefaults instantiates a new ValidationReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassed

`func (o *ValidationReport) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *ValidationReport) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *ValidationReport) SetPassed(v bool)`

SetPassed sets Passed field to given value.


### GetResults

`func (o *ValidationReport) GetResults() []ValidationResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ValidationReport) GetResultsOk() (*[]ValidationResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ValidationReport) SetResults(v []ValidationResult)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


