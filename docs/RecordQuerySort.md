# RecordQuerySort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Sortable field path such as &#x60;created_at&#x60; or &#x60;payload.record_scope&#x60;. | 
**Direction** | **string** | Sort direction. | 

## Methods

### NewRecordQuerySort

`func NewRecordQuerySort(field string, direction string, ) *RecordQuerySort`

NewRecordQuerySort instantiates a new RecordQuerySort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordQuerySortWithDefaults

`func NewRecordQuerySortWithDefaults() *RecordQuerySort`

NewRecordQuerySortWithDefaults instantiates a new RecordQuerySort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *RecordQuerySort) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *RecordQuerySort) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *RecordQuerySort) SetField(v string)`

SetField sets Field field to given value.


### GetDirection

`func (o *RecordQuerySort) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *RecordQuerySort) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *RecordQuerySort) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


