# ConfigureObjectConfigureOperations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identificador de tipo de Operación, Ej Sale | [optional] 
**Description** | Pointer to **string** | Descripción de la Operación | [optional] 
**Request** | Pointer to **map[string]interface{}** | Especificación del requerimiento | [optional] 
**Answer** | Pointer to **map[string]interface{}** | Especificación de la respuesta | [optional] 
**DefaultTimeout** | Pointer to **int32** | Timeout por default para dicha operación expresado en milisegúndos | [optional] 
**Version** | Pointer to **string** | Versión de la transacción | [optional] 
**EffectiveFrom** | Pointer to **time.Time** | Fecha y hora de a partir de la cual este archivo entra en vigencia - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**EffectiveTo** | Pointer to **time.Time** | Fecha y hora de a partir de la cual este archivo entra en vigencia - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 

## Methods

### NewConfigureObjectConfigureOperations

`func NewConfigureObjectConfigureOperations() *ConfigureObjectConfigureOperations`

NewConfigureObjectConfigureOperations instantiates a new ConfigureObjectConfigureOperations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigureObjectConfigureOperationsWithDefaults

`func NewConfigureObjectConfigureOperationsWithDefaults() *ConfigureObjectConfigureOperations`

NewConfigureObjectConfigureOperationsWithDefaults instantiates a new ConfigureObjectConfigureOperations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigureObjectConfigureOperations) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigureObjectConfigureOperations) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigureObjectConfigureOperations) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigureObjectConfigureOperations) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigureObjectConfigureOperations) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigureObjectConfigureOperations) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigureObjectConfigureOperations) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigureObjectConfigureOperations) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequest

`func (o *ConfigureObjectConfigureOperations) GetRequest() map[string]interface{}`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ConfigureObjectConfigureOperations) GetRequestOk() (*map[string]interface{}, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ConfigureObjectConfigureOperations) SetRequest(v map[string]interface{})`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ConfigureObjectConfigureOperations) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetAnswer

`func (o *ConfigureObjectConfigureOperations) GetAnswer() map[string]interface{}`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *ConfigureObjectConfigureOperations) GetAnswerOk() (*map[string]interface{}, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *ConfigureObjectConfigureOperations) SetAnswer(v map[string]interface{})`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *ConfigureObjectConfigureOperations) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetDefaultTimeout

`func (o *ConfigureObjectConfigureOperations) GetDefaultTimeout() int32`

GetDefaultTimeout returns the DefaultTimeout field if non-nil, zero value otherwise.

### GetDefaultTimeoutOk

`func (o *ConfigureObjectConfigureOperations) GetDefaultTimeoutOk() (*int32, bool)`

GetDefaultTimeoutOk returns a tuple with the DefaultTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeout

`func (o *ConfigureObjectConfigureOperations) SetDefaultTimeout(v int32)`

SetDefaultTimeout sets DefaultTimeout field to given value.

### HasDefaultTimeout

`func (o *ConfigureObjectConfigureOperations) HasDefaultTimeout() bool`

HasDefaultTimeout returns a boolean if a field has been set.

### GetVersion

`func (o *ConfigureObjectConfigureOperations) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConfigureObjectConfigureOperations) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConfigureObjectConfigureOperations) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConfigureObjectConfigureOperations) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEffectiveFrom

`func (o *ConfigureObjectConfigureOperations) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *ConfigureObjectConfigureOperations) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *ConfigureObjectConfigureOperations) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *ConfigureObjectConfigureOperations) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetEffectiveTo

`func (o *ConfigureObjectConfigureOperations) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *ConfigureObjectConfigureOperations) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *ConfigureObjectConfigureOperations) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *ConfigureObjectConfigureOperations) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


