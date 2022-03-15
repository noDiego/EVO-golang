# PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID interno en la base de datos para este metodo de pago | [optional] 
**ForeignIdentifier** | Pointer to **string** | ID externo de este metodo de pago, utilizado por el punto de venta para reconocer al medio de pago en su base de datos | [optional] 
**Image** | Pointer to **string** | Nombre de la Imagen o Logo o la imagen codificado en base64 | [optional] 
**LargeImage** | Pointer to **string** | Nombre de la Imagen o Logo o la imagen codificado en base64 | [optional] 
**SmallImage** | Pointer to **string** | Nombre de la Imagen o Logo o la imagen codificado en base64 | [optional] 
**CardNumberMaxLength** | Pointer to **int32** | Longitud máxima del Número de tarjeta de este medio de pago | [optional] 
**CardNumberMinLength** | Pointer to **int32** | Longitud mínima del Número de tarjeta de este medio de pago | [optional] 
**SecurityCodeMaxLength** | Pointer to **int32** | Longitud máxima del código de seguridad de tarjeta de este medio de pago | [optional] 
**SecurityCodeMinLength** | Pointer to **int32** | Longitud mínima del código de seguridad de tarjeta de este medio de pago | [optional] 
**Description** | Pointer to **string** | LIena que describe al metodo de pago | [optional] 
**Type** | Pointer to [**PaymentMethodsResponseObjectPaymentMethodsResponseType**](PaymentMethodsResponseObjectPaymentMethodsResponseType.md) |  | [optional] 
**Category** | Pointer to [**PaymentMethodsResponseObjectPaymentMethodsResponseCategory**](PaymentMethodsResponseObjectPaymentMethodsResponseCategory.md) |  | [optional] 

## Methods

### NewPaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods

`func NewPaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods() *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods`

NewPaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods instantiates a new PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethodsWithDefaults

`func NewPaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethodsWithDefaults() *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods`

NewPaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethodsWithDefaults instantiates a new PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) HasId() bool`

HasId returns a boolean if a field has been set.

### GetForeignIdentifier

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetForeignIdentifier() string`

GetForeignIdentifier returns the ForeignIdentifier field if non-nil, zero value otherwise.

### GetForeignIdentifierOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetForeignIdentifierOk() (*string, bool)`

GetForeignIdentifierOk returns a tuple with the ForeignIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignIdentifier

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) SetForeignIdentifier(v string)`

SetForeignIdentifier sets ForeignIdentifier field to given value.

### HasForeignIdentifier

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) HasForeignIdentifier() bool`

HasForeignIdentifier returns a boolean if a field has been set.

### GetImage

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLargeImage

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetLargeImage() string`

GetLargeImage returns the LargeImage field if non-nil, zero value otherwise.

### GetLargeImageOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetLargeImageOk() (*string, bool)`

GetLargeImageOk returns a tuple with the LargeImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeImage

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) SetLargeImage(v string)`

SetLargeImage sets LargeImage field to given value.

### HasLargeImage

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) HasLargeImage() bool`

HasLargeImage returns a boolean if a field has been set.

### GetSmallImage

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetSmallImage() string`

GetSmallImage returns the SmallImage field if non-nil, zero value otherwise.

### GetSmallImageOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetSmallImageOk() (*string, bool)`

GetSmallImageOk returns a tuple with the SmallImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallImage

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) SetSmallImage(v string)`

SetSmallImage sets SmallImage field to given value.

### HasSmallImage

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) HasSmallImage() bool`

HasSmallImage returns a boolean if a field has been set.

### GetCardNumberMaxLength

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetCardNumberMaxLength() int32`

GetCardNumberMaxLength returns the CardNumberMaxLength field if non-nil, zero value otherwise.

### GetCardNumberMaxLengthOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetCardNumberMaxLengthOk() (*int32, bool)`

GetCardNumberMaxLengthOk returns a tuple with the CardNumberMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberMaxLength

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) SetCardNumberMaxLength(v int32)`

SetCardNumberMaxLength sets CardNumberMaxLength field to given value.

### HasCardNumberMaxLength

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) HasCardNumberMaxLength() bool`

HasCardNumberMaxLength returns a boolean if a field has been set.

### GetCardNumberMinLength

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetCardNumberMinLength() int32`

GetCardNumberMinLength returns the CardNumberMinLength field if non-nil, zero value otherwise.

### GetCardNumberMinLengthOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetCardNumberMinLengthOk() (*int32, bool)`

GetCardNumberMinLengthOk returns a tuple with the CardNumberMinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberMinLength

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) SetCardNumberMinLength(v int32)`

SetCardNumberMinLength sets CardNumberMinLength field to given value.

### HasCardNumberMinLength

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) HasCardNumberMinLength() bool`

HasCardNumberMinLength returns a boolean if a field has been set.

### GetSecurityCodeMaxLength

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetSecurityCodeMaxLength() int32`

GetSecurityCodeMaxLength returns the SecurityCodeMaxLength field if non-nil, zero value otherwise.

### GetSecurityCodeMaxLengthOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetSecurityCodeMaxLengthOk() (*int32, bool)`

GetSecurityCodeMaxLengthOk returns a tuple with the SecurityCodeMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCodeMaxLength

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) SetSecurityCodeMaxLength(v int32)`

SetSecurityCodeMaxLength sets SecurityCodeMaxLength field to given value.

### HasSecurityCodeMaxLength

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) HasSecurityCodeMaxLength() bool`

HasSecurityCodeMaxLength returns a boolean if a field has been set.

### GetSecurityCodeMinLength

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetSecurityCodeMinLength() int32`

GetSecurityCodeMinLength returns the SecurityCodeMinLength field if non-nil, zero value otherwise.

### GetSecurityCodeMinLengthOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetSecurityCodeMinLengthOk() (*int32, bool)`

GetSecurityCodeMinLengthOk returns a tuple with the SecurityCodeMinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCodeMinLength

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) SetSecurityCodeMinLength(v int32)`

SetSecurityCodeMinLength sets SecurityCodeMinLength field to given value.

### HasSecurityCodeMinLength

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) HasSecurityCodeMinLength() bool`

HasSecurityCodeMinLength returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetType() PaymentMethodsResponseObjectPaymentMethodsResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetTypeOk() (*PaymentMethodsResponseObjectPaymentMethodsResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) SetType(v PaymentMethodsResponseObjectPaymentMethodsResponseType)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCategory

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetCategory() PaymentMethodsResponseObjectPaymentMethodsResponseCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) GetCategoryOk() (*PaymentMethodsResponseObjectPaymentMethodsResponseCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) SetCategory(v PaymentMethodsResponseObjectPaymentMethodsResponseCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponsePaymentMethods) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


