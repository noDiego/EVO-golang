# QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Nombre del dato a solicitar | [optional] 
**Type** | Pointer to **string** | Tipo de dato a solicitar | [optional] 
**InterpretFor** | Pointer to **string** | Este elemento se le solicita al POS o el dispositivo de venta/ReadingDevice.  En caso de que se requiera RequieredInformation y no esté definida esta propiedad  se debe asumir que es para el POS. | [optional] 
**ValidationExpressionType** | Pointer to **string** | Indica el tipo de Expresión de Validación. | [optional] 
**ValidationExpression** | Pointer to **string** | Expresión de Validación. | [optional] 
**UIType** | Pointer to **string** | Tipo de elemento a utilizar para su representación en la Interfaz de Usuario ( Button, etc ) | [optional] 
**UIAttributes** | Pointer to **string** | Atributos de Forma, font, color, etc necesarios por el aplicativo | [optional] 
**Value** | Pointer to **string** | Valor del Elemento Adicional Solicitado | [optional] 
**Label** | Pointer to **string** | Leyenda a mostrar en pantalla para el dato a solicitar | [optional] 
**MinLength** | Pointer to **int32** | Longitud mínima del dato a solicitar | [optional] 
**MaxLength** | Pointer to **int32** | Longitud máxima del dato a solicitar | [optional] 
**Mandatory** | Pointer to **bool** | Indicador de obligatoriedad de ingreso de este dato | [optional] 

## Methods

### NewQueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation

`func NewQueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation() *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation`

NewQueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation instantiates a new QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformationWithDefaults

`func NewQueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformationWithDefaults() *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation`

NewQueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformationWithDefaults instantiates a new QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInterpretFor

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetInterpretFor() string`

GetInterpretFor returns the InterpretFor field if non-nil, zero value otherwise.

### GetInterpretForOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetInterpretForOk() (*string, bool)`

GetInterpretForOk returns a tuple with the InterpretFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpretFor

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) SetInterpretFor(v string)`

SetInterpretFor sets InterpretFor field to given value.

### HasInterpretFor

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) HasInterpretFor() bool`

HasInterpretFor returns a boolean if a field has been set.

### GetValidationExpressionType

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetValidationExpressionType() string`

GetValidationExpressionType returns the ValidationExpressionType field if non-nil, zero value otherwise.

### GetValidationExpressionTypeOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetValidationExpressionTypeOk() (*string, bool)`

GetValidationExpressionTypeOk returns a tuple with the ValidationExpressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationExpressionType

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) SetValidationExpressionType(v string)`

SetValidationExpressionType sets ValidationExpressionType field to given value.

### HasValidationExpressionType

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) HasValidationExpressionType() bool`

HasValidationExpressionType returns a boolean if a field has been set.

### GetValidationExpression

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetValidationExpression() string`

GetValidationExpression returns the ValidationExpression field if non-nil, zero value otherwise.

### GetValidationExpressionOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetValidationExpressionOk() (*string, bool)`

GetValidationExpressionOk returns a tuple with the ValidationExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationExpression

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) SetValidationExpression(v string)`

SetValidationExpression sets ValidationExpression field to given value.

### HasValidationExpression

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) HasValidationExpression() bool`

HasValidationExpression returns a boolean if a field has been set.

### GetUIType

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetUIType() string`

GetUIType returns the UIType field if non-nil, zero value otherwise.

### GetUITypeOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetUITypeOk() (*string, bool)`

GetUITypeOk returns a tuple with the UIType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUIType

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) SetUIType(v string)`

SetUIType sets UIType field to given value.

### HasUIType

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) HasUIType() bool`

HasUIType returns a boolean if a field has been set.

### GetUIAttributes

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetUIAttributes() string`

GetUIAttributes returns the UIAttributes field if non-nil, zero value otherwise.

### GetUIAttributesOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetUIAttributesOk() (*string, bool)`

GetUIAttributesOk returns a tuple with the UIAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUIAttributes

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) SetUIAttributes(v string)`

SetUIAttributes sets UIAttributes field to given value.

### HasUIAttributes

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) HasUIAttributes() bool`

HasUIAttributes returns a boolean if a field has been set.

### GetValue

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMinLength

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMandatory

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


