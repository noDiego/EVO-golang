# EnrollmentObjectEnrollmentRecurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntervalType** | Pointer to **string** | Tipo de intervalo. 0 - SEMANAL; 1 - MENSUAL; 2 - ANUAL | [optional] 
**IntervalValue** | Pointer to **int32** | Cantidad de pagos a realizar. | [optional] 
**NumberOfRecurrences** | Pointer to **string** | Cantidad de recurrencias. | [optional] 

## Methods

### NewEnrollmentObjectEnrollmentRecurrence

`func NewEnrollmentObjectEnrollmentRecurrence() *EnrollmentObjectEnrollmentRecurrence`

NewEnrollmentObjectEnrollmentRecurrence instantiates a new EnrollmentObjectEnrollmentRecurrence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentObjectEnrollmentRecurrenceWithDefaults

`func NewEnrollmentObjectEnrollmentRecurrenceWithDefaults() *EnrollmentObjectEnrollmentRecurrence`

NewEnrollmentObjectEnrollmentRecurrenceWithDefaults instantiates a new EnrollmentObjectEnrollmentRecurrence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervalType

`func (o *EnrollmentObjectEnrollmentRecurrence) GetIntervalType() string`

GetIntervalType returns the IntervalType field if non-nil, zero value otherwise.

### GetIntervalTypeOk

`func (o *EnrollmentObjectEnrollmentRecurrence) GetIntervalTypeOk() (*string, bool)`

GetIntervalTypeOk returns a tuple with the IntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalType

`func (o *EnrollmentObjectEnrollmentRecurrence) SetIntervalType(v string)`

SetIntervalType sets IntervalType field to given value.

### HasIntervalType

`func (o *EnrollmentObjectEnrollmentRecurrence) HasIntervalType() bool`

HasIntervalType returns a boolean if a field has been set.

### GetIntervalValue

`func (o *EnrollmentObjectEnrollmentRecurrence) GetIntervalValue() int32`

GetIntervalValue returns the IntervalValue field if non-nil, zero value otherwise.

### GetIntervalValueOk

`func (o *EnrollmentObjectEnrollmentRecurrence) GetIntervalValueOk() (*int32, bool)`

GetIntervalValueOk returns a tuple with the IntervalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalValue

`func (o *EnrollmentObjectEnrollmentRecurrence) SetIntervalValue(v int32)`

SetIntervalValue sets IntervalValue field to given value.

### HasIntervalValue

`func (o *EnrollmentObjectEnrollmentRecurrence) HasIntervalValue() bool`

HasIntervalValue returns a boolean if a field has been set.

### GetNumberOfRecurrences

`func (o *EnrollmentObjectEnrollmentRecurrence) GetNumberOfRecurrences() string`

GetNumberOfRecurrences returns the NumberOfRecurrences field if non-nil, zero value otherwise.

### GetNumberOfRecurrencesOk

`func (o *EnrollmentObjectEnrollmentRecurrence) GetNumberOfRecurrencesOk() (*string, bool)`

GetNumberOfRecurrencesOk returns a tuple with the NumberOfRecurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRecurrences

`func (o *EnrollmentObjectEnrollmentRecurrence) SetNumberOfRecurrences(v string)`

SetNumberOfRecurrences sets NumberOfRecurrences field to given value.

### HasNumberOfRecurrences

`func (o *EnrollmentObjectEnrollmentRecurrence) HasNumberOfRecurrences() bool`

HasNumberOfRecurrences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


