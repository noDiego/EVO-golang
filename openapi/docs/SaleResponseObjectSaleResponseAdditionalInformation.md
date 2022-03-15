# SaleResponseObjectSaleResponseAdditionalInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Nombre del dato a solicitar | [optional] 
**Type** | Pointer to **string** | Tipo de dato a solicitar | [optional] 
**InterpretFor** | Pointer to **string** | Este elemento se le solicita al POS o el dispositivo de venta/ReadingDevice.  En caso de que se requiera RequieredInformation y no esté definida esta propiedad  se debe asumir que es para el POS. | [optional] 
**UIType** | Pointer to **string** | Tipo de elemento a utilizar para su representación en la Interfaz de Usuario ( Button, etc ) | [optional] 
**UIAttributes** | Pointer to **string** | Atributos de Forma, font, color, etc necesarios por el aplicativo | [optional] 
**Value** | Pointer to **string** | Valor del Elemento Adicional Solicitado | [optional] 
**Label** | Pointer to **string** | Leyenda a mostrar en pantalla para el dato a solicitar | [optional] 
**MinLength** | Pointer to **int32** | Longitud mínima del dato a solicitar | [optional] 
**MaxLength** | Pointer to **int32** | Longitud máxima del dato a solicitar | [optional] 

## Methods

### NewSaleResponseObjectSaleResponseAdditionalInformation

`func NewSaleResponseObjectSaleResponseAdditionalInformation() *SaleResponseObjectSaleResponseAdditionalInformation`

NewSaleResponseObjectSaleResponseAdditionalInformation instantiates a new SaleResponseObjectSaleResponseAdditionalInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleResponseObjectSaleResponseAdditionalInformationWithDefaults

`func NewSaleResponseObjectSaleResponseAdditionalInformationWithDefaults() *SaleResponseObjectSaleResponseAdditionalInformation`

NewSaleResponseObjectSaleResponseAdditionalInformationWithDefaults instantiates a new SaleResponseObjectSaleResponseAdditionalInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInterpretFor

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetInterpretFor() string`

GetInterpretFor returns the InterpretFor field if non-nil, zero value otherwise.

### GetInterpretForOk

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetInterpretForOk() (*string, bool)`

GetInterpretForOk returns a tuple with the InterpretFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpretFor

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) SetInterpretFor(v string)`

SetInterpretFor sets InterpretFor field to given value.

### HasInterpretFor

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) HasInterpretFor() bool`

HasInterpretFor returns a boolean if a field has been set.

### GetUIType

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetUIType() string`

GetUIType returns the UIType field if non-nil, zero value otherwise.

### GetUITypeOk

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetUITypeOk() (*string, bool)`

GetUITypeOk returns a tuple with the UIType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUIType

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) SetUIType(v string)`

SetUIType sets UIType field to given value.

### HasUIType

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) HasUIType() bool`

HasUIType returns a boolean if a field has been set.

### GetUIAttributes

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetUIAttributes() string`

GetUIAttributes returns the UIAttributes field if non-nil, zero value otherwise.

### GetUIAttributesOk

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetUIAttributesOk() (*string, bool)`

GetUIAttributesOk returns a tuple with the UIAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUIAttributes

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) SetUIAttributes(v string)`

SetUIAttributes sets UIAttributes field to given value.

### HasUIAttributes

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) HasUIAttributes() bool`

HasUIAttributes returns a boolean if a field has been set.

### GetValue

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMinLength

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *SaleResponseObjectSaleResponseAdditionalInformation) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


