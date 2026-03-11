# RecordQueryCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Field path such as &#x60;id&#x60; or &#x60;payload.record_scope&#x60;. | 
**Op** | **string** | Comparison operator. | 
**Value** | **interface{}** | Operand value. For &#x60;in&#x60; this must be an array. For &#x60;exists&#x60; this should be boolean. | 

## Methods

### NewRecordQueryCondition

`func NewRecordQueryCondition(field string, op string, value interface{}, ) *RecordQueryCondition`

NewRecordQueryCondition instantiates a new RecordQueryCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordQueryConditionWithDefaults

`func NewRecordQueryConditionWithDefaults() *RecordQueryCondition`

NewRecordQueryConditionWithDefaults instantiates a new RecordQueryCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *RecordQueryCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *RecordQueryCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *RecordQueryCondition) SetField(v string)`

SetField sets Field field to given value.


### GetOp

`func (o *RecordQueryCondition) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RecordQueryCondition) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RecordQueryCondition) SetOp(v string)`

SetOp sets Op field to given value.


### GetValue

`func (o *RecordQueryCondition) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RecordQueryCondition) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RecordQueryCondition) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *RecordQueryCondition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *RecordQueryCondition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


