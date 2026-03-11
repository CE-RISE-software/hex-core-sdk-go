# RecordQueryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Where** | [**[]RecordQueryCondition**](RecordQueryCondition.md) | AND-only list of query predicates. | 
**Sort** | Pointer to [**[]RecordQuerySort**](RecordQuerySort.md) | Optional list of sort directives applied in order. | [optional] 
**Limit** | Pointer to **int64** | Maximum number of records to return. | [optional] 
**Offset** | Pointer to **int64** | Zero-based row offset for pagination. | [optional] 

## Methods

### NewRecordQueryFilter

`func NewRecordQueryFilter(where []RecordQueryCondition, ) *RecordQueryFilter`

NewRecordQueryFilter instantiates a new RecordQueryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordQueryFilterWithDefaults

`func NewRecordQueryFilterWithDefaults() *RecordQueryFilter`

NewRecordQueryFilterWithDefaults instantiates a new RecordQueryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhere

`func (o *RecordQueryFilter) GetWhere() []RecordQueryCondition`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *RecordQueryFilter) GetWhereOk() (*[]RecordQueryCondition, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *RecordQueryFilter) SetWhere(v []RecordQueryCondition)`

SetWhere sets Where field to given value.


### GetSort

`func (o *RecordQueryFilter) GetSort() []RecordQuerySort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RecordQueryFilter) GetSortOk() (*[]RecordQuerySort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RecordQueryFilter) SetSort(v []RecordQuerySort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RecordQueryFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetLimit

`func (o *RecordQueryFilter) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RecordQueryFilter) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RecordQueryFilter) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RecordQueryFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *RecordQueryFilter) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RecordQueryFilter) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RecordQueryFilter) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RecordQueryFilter) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


