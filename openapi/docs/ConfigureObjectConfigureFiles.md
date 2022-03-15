# ConfigureObjectConfigureFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identificador del archivo | [optional] 
**Description** | Pointer to **string** | Decripción | [optional] 
**Type** | Pointer to **string** | Tipo de archivo | [optional] 
**Path** | Pointer to **string** | Lugar donde el archivo se encuentra o sebe alojarse | [optional] 
**Version** | Pointer to **string** | Version del archivo | [optional] 
**Attributes** | Pointer to **string** | Atributos del archivo | [optional] 
**Data** | Pointer to **map[string]interface{}** | contenido del archivo, codificado en Base64 | [optional] 
**EffectiveFrom** | Pointer to **time.Time** | Fecha y hora de a partir de la cual este archivo entra en vigencia - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**EffectiveTo** | Pointer to **time.Time** | Fecha y hora de a partir de la cual este archivo entra en vigencia - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**TimeOfLife** | Pointer to **int32** | Timeout de vida  para dicha tabla expresado en milisegúndos | [optional] 
**TimeOfLifeOffline** | Pointer to **int32** | Timeout de vida  para dicha tabla expresado en milisegúndos para el caso de que el POS o dispositivo se encuentre fuera de línea de la plataforma | [optional] 

## Methods

### NewConfigureObjectConfigureFiles

`func NewConfigureObjectConfigureFiles() *ConfigureObjectConfigureFiles`

NewConfigureObjectConfigureFiles instantiates a new ConfigureObjectConfigureFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigureObjectConfigureFilesWithDefaults

`func NewConfigureObjectConfigureFilesWithDefaults() *ConfigureObjectConfigureFiles`

NewConfigureObjectConfigureFilesWithDefaults instantiates a new ConfigureObjectConfigureFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigureObjectConfigureFiles) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigureObjectConfigureFiles) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigureObjectConfigureFiles) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigureObjectConfigureFiles) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigureObjectConfigureFiles) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigureObjectConfigureFiles) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigureObjectConfigureFiles) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigureObjectConfigureFiles) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ConfigureObjectConfigureFiles) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigureObjectConfigureFiles) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigureObjectConfigureFiles) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConfigureObjectConfigureFiles) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPath

`func (o *ConfigureObjectConfigureFiles) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ConfigureObjectConfigureFiles) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ConfigureObjectConfigureFiles) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ConfigureObjectConfigureFiles) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetVersion

`func (o *ConfigureObjectConfigureFiles) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConfigureObjectConfigureFiles) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConfigureObjectConfigureFiles) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConfigureObjectConfigureFiles) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAttributes

`func (o *ConfigureObjectConfigureFiles) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConfigureObjectConfigureFiles) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConfigureObjectConfigureFiles) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ConfigureObjectConfigureFiles) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetData

`func (o *ConfigureObjectConfigureFiles) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConfigureObjectConfigureFiles) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConfigureObjectConfigureFiles) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ConfigureObjectConfigureFiles) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEffectiveFrom

`func (o *ConfigureObjectConfigureFiles) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *ConfigureObjectConfigureFiles) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *ConfigureObjectConfigureFiles) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *ConfigureObjectConfigureFiles) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetEffectiveTo

`func (o *ConfigureObjectConfigureFiles) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *ConfigureObjectConfigureFiles) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *ConfigureObjectConfigureFiles) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *ConfigureObjectConfigureFiles) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### GetTimeOfLife

`func (o *ConfigureObjectConfigureFiles) GetTimeOfLife() int32`

GetTimeOfLife returns the TimeOfLife field if non-nil, zero value otherwise.

### GetTimeOfLifeOk

`func (o *ConfigureObjectConfigureFiles) GetTimeOfLifeOk() (*int32, bool)`

GetTimeOfLifeOk returns a tuple with the TimeOfLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfLife

`func (o *ConfigureObjectConfigureFiles) SetTimeOfLife(v int32)`

SetTimeOfLife sets TimeOfLife field to given value.

### HasTimeOfLife

`func (o *ConfigureObjectConfigureFiles) HasTimeOfLife() bool`

HasTimeOfLife returns a boolean if a field has been set.

### GetTimeOfLifeOffline

`func (o *ConfigureObjectConfigureFiles) GetTimeOfLifeOffline() int32`

GetTimeOfLifeOffline returns the TimeOfLifeOffline field if non-nil, zero value otherwise.

### GetTimeOfLifeOfflineOk

`func (o *ConfigureObjectConfigureFiles) GetTimeOfLifeOfflineOk() (*int32, bool)`

GetTimeOfLifeOfflineOk returns a tuple with the TimeOfLifeOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfLifeOffline

`func (o *ConfigureObjectConfigureFiles) SetTimeOfLifeOffline(v int32)`

SetTimeOfLifeOffline sets TimeOfLifeOffline field to given value.

### HasTimeOfLifeOffline

`func (o *ConfigureObjectConfigureFiles) HasTimeOfLifeOffline() bool`

HasTimeOfLifeOffline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


