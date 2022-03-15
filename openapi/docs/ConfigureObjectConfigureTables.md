# ConfigureObjectConfigureTables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identificador de tabla | [optional] 
**Description** | Pointer to **string** | Descripción de la Tabla | [optional] 
**Version** | Pointer to **string** | Versión de la Tabla | [optional] 
**Data** | Pointer to **map[string]interface{}** | Representación de la tabla, codificado en Base64 en caso de no ser JSON el mismo | [optional] 
**EffectiveFrom** | Pointer to **time.Time** | Fecha y hora de a partir de la cual este archivo entra en vigencia - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**EffectiveTo** | Pointer to **time.Time** | Fecha y hora de a partir de la cual este archivo entra en vigencia - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**TimeOfLife** | Pointer to **int32** | Timeout de vida  para dicha tabla expresado en milisegúndos | [optional] 
**TimeOfLifeOffline** | Pointer to **int32** | Timeout de vida  para dicha tabla expresado en milisegúndos para el caso de que el POS o dispositivo se encuentre fuera de línea de la plataforma | [optional] 

## Methods

### NewConfigureObjectConfigureTables

`func NewConfigureObjectConfigureTables() *ConfigureObjectConfigureTables`

NewConfigureObjectConfigureTables instantiates a new ConfigureObjectConfigureTables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigureObjectConfigureTablesWithDefaults

`func NewConfigureObjectConfigureTablesWithDefaults() *ConfigureObjectConfigureTables`

NewConfigureObjectConfigureTablesWithDefaults instantiates a new ConfigureObjectConfigureTables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigureObjectConfigureTables) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigureObjectConfigureTables) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigureObjectConfigureTables) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigureObjectConfigureTables) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigureObjectConfigureTables) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigureObjectConfigureTables) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigureObjectConfigureTables) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigureObjectConfigureTables) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *ConfigureObjectConfigureTables) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConfigureObjectConfigureTables) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConfigureObjectConfigureTables) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConfigureObjectConfigureTables) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetData

`func (o *ConfigureObjectConfigureTables) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConfigureObjectConfigureTables) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConfigureObjectConfigureTables) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ConfigureObjectConfigureTables) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEffectiveFrom

`func (o *ConfigureObjectConfigureTables) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *ConfigureObjectConfigureTables) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *ConfigureObjectConfigureTables) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *ConfigureObjectConfigureTables) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetEffectiveTo

`func (o *ConfigureObjectConfigureTables) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *ConfigureObjectConfigureTables) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *ConfigureObjectConfigureTables) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *ConfigureObjectConfigureTables) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### GetTimeOfLife

`func (o *ConfigureObjectConfigureTables) GetTimeOfLife() int32`

GetTimeOfLife returns the TimeOfLife field if non-nil, zero value otherwise.

### GetTimeOfLifeOk

`func (o *ConfigureObjectConfigureTables) GetTimeOfLifeOk() (*int32, bool)`

GetTimeOfLifeOk returns a tuple with the TimeOfLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfLife

`func (o *ConfigureObjectConfigureTables) SetTimeOfLife(v int32)`

SetTimeOfLife sets TimeOfLife field to given value.

### HasTimeOfLife

`func (o *ConfigureObjectConfigureTables) HasTimeOfLife() bool`

HasTimeOfLife returns a boolean if a field has been set.

### GetTimeOfLifeOffline

`func (o *ConfigureObjectConfigureTables) GetTimeOfLifeOffline() int32`

GetTimeOfLifeOffline returns the TimeOfLifeOffline field if non-nil, zero value otherwise.

### GetTimeOfLifeOfflineOk

`func (o *ConfigureObjectConfigureTables) GetTimeOfLifeOfflineOk() (*int32, bool)`

GetTimeOfLifeOfflineOk returns a tuple with the TimeOfLifeOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfLifeOffline

`func (o *ConfigureObjectConfigureTables) SetTimeOfLifeOffline(v int32)`

SetTimeOfLifeOffline sets TimeOfLifeOffline field to given value.

### HasTimeOfLifeOffline

`func (o *ConfigureObjectConfigureTables) HasTimeOfLifeOffline() bool`

HasTimeOfLifeOffline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


