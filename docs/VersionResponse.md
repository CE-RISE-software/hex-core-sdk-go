# VersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | **string** | Stable service identifier. | 
**ServiceVersion** | **string** | Running service build version from crate metadata. | 
**OpenapiVersion** | **string** | Bundled OpenAPI document version from info.version. | 

## Methods

### NewVersionResponse

`func NewVersionResponse(service string, serviceVersion string, openapiVersion string, ) *VersionResponse`

NewVersionResponse instantiates a new VersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionResponseWithDefaults

`func NewVersionResponseWithDefaults() *VersionResponse`

NewVersionResponseWithDefaults instantiates a new VersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *VersionResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *VersionResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *VersionResponse) SetService(v string)`

SetService sets Service field to given value.


### GetServiceVersion

`func (o *VersionResponse) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *VersionResponse) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *VersionResponse) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.


### GetOpenapiVersion

`func (o *VersionResponse) GetOpenapiVersion() string`

GetOpenapiVersion returns the OpenapiVersion field if non-nil, zero value otherwise.

### GetOpenapiVersionOk

`func (o *VersionResponse) GetOpenapiVersionOk() (*string, bool)`

GetOpenapiVersionOk returns a tuple with the OpenapiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapiVersion

`func (o *VersionResponse) SetOpenapiVersion(v string)`

SetOpenapiVersion sets OpenapiVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


