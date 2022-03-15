# SaleResponseObjectSaleResponseRequiredInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Nombre del dato a solicitar. | [optional] 
**Type** | Pointer to **string** | Tipo de dato a solicitar. | [optional] 
**InterpretFor** | Pointer to **string** | Este elemento se le solicita al POS o el dispositivo de venta/ReadingDevice.  En caso de que se requiera RequieredInformation y no esté definida esta propiedad  se debe asumir que es para el POS. | [optional] 
**ItIsDefined** | Pointer to **bool** | Indicador si el elemento esta documentado/definido en el contrato, en caso de true, entonces solo se informará el Name y opcionalmente Mandatory, ya que todos los otros atributos del mismo, pero si fuese false entonces estarán presentes los otros atributos. Con respecto a cómo deben ser enviados, si está ya definido en el contrato deberá ser enviado como el contrato lo especifica. En caso contrario deberá ser enviado como un objeto compuesto del array de objetos en el elemento RequiredInformation. | [optional] 
**UIType** | Pointer to **string** | Tipo de elemento a utilizar para su representación en la Interfaz de Usuario ( Button, etc ). | [optional] 
**UIAttributes** | Pointer to **string** | Atributos de Forma, font, color, etc necesarios por el aplicativo. | [optional] 
**Value** | Pointer to **string** | Valor del Elemento Adicional Solicitado. | [optional] 
**Label** | Pointer to **string** | Leyenda a mostrar en pantalla para el dato a solicitar. | [optional] 
**MinLength** | Pointer to **int32** | Longitud mínima del dato a solicitar. | [optional] 
**MaxLength** | Pointer to **int32** | Longitud máxima del dato a solicitar. | [optional] 
**ValidationExpressionType** | Pointer to **string** | Indica el tipo de Expresión de Validación. | [optional] 
**ValidationExpression** | Pointer to **string** | Expresión de Validación. | [optional] 
**Mandatory** | Pointer to **bool** | Indicador de obligatoriedad de ingreso de este dato. | [optional] 
**AdditionalInformation** | Pointer to [**[]SaleResponseObjectSaleResponseAdditionalInformation**](SaleResponseObjectSaleResponseAdditionalInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 

## Methods

### NewSaleResponseObjectSaleResponseRequiredInformation

`func NewSaleResponseObjectSaleResponseRequiredInformation() *SaleResponseObjectSaleResponseRequiredInformation`

NewSaleResponseObjectSaleResponseRequiredInformation instantiates a new SaleResponseObjectSaleResponseRequiredInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleResponseObjectSaleResponseRequiredInformationWithDefaults

`func NewSaleResponseObjectSaleResponseRequiredInformationWithDefaults() *SaleResponseObjectSaleResponseRequiredInformation`

NewSaleResponseObjectSaleResponseRequiredInformationWithDefaults instantiates a new SaleResponseObjectSaleResponseRequiredInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaleResponseObjectSaleResponseRequiredInformation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SaleResponseObjectSaleResponseRequiredInformation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SaleResponseObjectSaleResponseRequiredInformation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SaleResponseObjectSaleResponseRequiredInformation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInterpretFor

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetInterpretFor() string`

GetInterpretFor returns the InterpretFor field if non-nil, zero value otherwise.

### GetInterpretForOk

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetInterpretForOk() (*string, bool)`

GetInterpretForOk returns a tuple with the InterpretFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpretFor

`func (o *SaleResponseObjectSaleResponseRequiredInformation) SetInterpretFor(v string)`

SetInterpretFor sets InterpretFor field to given value.

### HasInterpretFor

`func (o *SaleResponseObjectSaleResponseRequiredInformation) HasInterpretFor() bool`

HasInterpretFor returns a boolean if a field has been set.

### GetItIsDefined

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetItIsDefined() bool`

GetItIsDefined returns the ItIsDefined field if non-nil, zero value otherwise.

### GetItIsDefinedOk

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetItIsDefinedOk() (*bool, bool)`

GetItIsDefinedOk returns a tuple with the ItIsDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItIsDefined

`func (o *SaleResponseObjectSaleResponseRequiredInformation) SetItIsDefined(v bool)`

SetItIsDefined sets ItIsDefined field to given value.

### HasItIsDefined

`func (o *SaleResponseObjectSaleResponseRequiredInformation) HasItIsDefined() bool`

HasItIsDefined returns a boolean if a field has been set.

### GetUIType

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetUIType() string`

GetUIType returns the UIType field if non-nil, zero value otherwise.

### GetUITypeOk

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetUITypeOk() (*string, bool)`

GetUITypeOk returns a tuple with the UIType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUIType

`func (o *SaleResponseObjectSaleResponseRequiredInformation) SetUIType(v string)`

SetUIType sets UIType field to given value.

### HasUIType

`func (o *SaleResponseObjectSaleResponseRequiredInformation) HasUIType() bool`

HasUIType returns a boolean if a field has been set.

### GetUIAttributes

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetUIAttributes() string`

GetUIAttributes returns the UIAttributes field if non-nil, zero value otherwise.

### GetUIAttributesOk

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetUIAttributesOk() (*string, bool)`

GetUIAttributesOk returns a tuple with the UIAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUIAttributes

`func (o *SaleResponseObjectSaleResponseRequiredInformation) SetUIAttributes(v string)`

SetUIAttributes sets UIAttributes field to given value.

### HasUIAttributes

`func (o *SaleResponseObjectSaleResponseRequiredInformation) HasUIAttributes() bool`

HasUIAttributes returns a boolean if a field has been set.

### GetValue

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SaleResponseObjectSaleResponseRequiredInformation) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SaleResponseObjectSaleResponseRequiredInformation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SaleResponseObjectSaleResponseRequiredInformation) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SaleResponseObjectSaleResponseRequiredInformation) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMinLength

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *SaleResponseObjectSaleResponseRequiredInformation) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *SaleResponseObjectSaleResponseRequiredInformation) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *SaleResponseObjectSaleResponseRequiredInformation) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *SaleResponseObjectSaleResponseRequiredInformation) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetValidationExpressionType

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetValidationExpressionType() string`

GetValidationExpressionType returns the ValidationExpressionType field if non-nil, zero value otherwise.

### GetValidationExpressionTypeOk

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetValidationExpressionTypeOk() (*string, bool)`

GetValidationExpressionTypeOk returns a tuple with the ValidationExpressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationExpressionType

`func (o *SaleResponseObjectSaleResponseRequiredInformation) SetValidationExpressionType(v string)`

SetValidationExpressionType sets ValidationExpressionType field to given value.

### HasValidationExpressionType

`func (o *SaleResponseObjectSaleResponseRequiredInformation) HasValidationExpressionType() bool`

HasValidationExpressionType returns a boolean if a field has been set.

### GetValidationExpression

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetValidationExpression() string`

GetValidationExpression returns the ValidationExpression field if non-nil, zero value otherwise.

### GetValidationExpressionOk

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetValidationExpressionOk() (*string, bool)`

GetValidationExpressionOk returns a tuple with the ValidationExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationExpression

`func (o *SaleResponseObjectSaleResponseRequiredInformation) SetValidationExpression(v string)`

SetValidationExpression sets ValidationExpression field to given value.

### HasValidationExpression

`func (o *SaleResponseObjectSaleResponseRequiredInformation) HasValidationExpression() bool`

HasValidationExpression returns a boolean if a field has been set.

### GetMandatory

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *SaleResponseObjectSaleResponseRequiredInformation) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *SaleResponseObjectSaleResponseRequiredInformation) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *SaleResponseObjectSaleResponseRequiredInformation) GetAdditionalInformationOk() (*[]SaleResponseObjectSaleResponseAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *SaleResponseObjectSaleResponseRequiredInformation) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *SaleResponseObjectSaleResponseRequiredInformation) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


