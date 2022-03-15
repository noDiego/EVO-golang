# WalletsResponseObjectWalletsResponseWallets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identification** | Pointer to **string** | Identificador o código de la Billetera o Wallet | [optional] 
**Name** | Pointer to **string** | Nombre del Wallet | [optional] 
**Label** | Pointer to **string** | Etiqueta a Mostrar para el Wallet | [optional] 
**Image** | Pointer to **string** | Nombre de la Imagen o Logo o la imagen codificado en base64 | [optional] 
**LargeImage** | Pointer to **string** | Nombre de la Imagen o Logo o la imagen codificado en base64 | [optional] 
**SmallImage** | Pointer to **string** | Nombre de la Imagen o Logo o la imagen codificado en base64 | [optional] 
**ForeignIdentifier** | Pointer to **string** | ID externo de este metodo de pago, utilizado por el punto de venta para reconocer al medio de pago en su base de datos | [optional] 
**AutoConfirm** | Pointer to **bool** | Indica de las Operaciones Financieras son AutoConfirmadas para este Wallet | [optional] 
**SupportRequestCancel** | Pointer to **bool** | Indica si dicho Wallet Soporta la cancelación del Requerimiento | [optional] 
**SupportValidityOfTheRequest** | Pointer to **bool** | Indica si dicho Wallet Soporta la en el envio de Requerimiento de Pago el tiempo de Vida del Requerimiento de Pagos usando el Elemento TransactionTimeout | [optional] 
**VoidSupport** | Pointer to **string** | Especifica el tipo las Canlelaciones/Anulaciones son soportadas | [optional] 
**WalletUseInVoidTransaction** | Pointer to **bool** | Indica si para Cancelar/Anular una transacción es requerida nuevamente la interacción contra el Wallet | [optional] 
**ReturnSupport** | Pointer to **string** | Especifica el tipo las Devoluciones son soportadas y de que tipo | [optional] 
**WalletUseInReturnTransaction** | Pointer to **bool** | Indica si para Devolver una transacción es requerida nuevamente la interacción contra el Wallet | [optional] 
**AmountRequired** | Pointer to **bool** | Indica Si el importe a cobrar debe enviarse en la Operación WalletRequest | [optional] 
**TokenType** | Pointer to **string** | Indica el Tipo de Token | [optional] 
**DefaultRetryTime** | Pointer to **int32** | Tiempo por omision en caso de que no sea retornaro en la respuesta de la operación WalletRequest expresado en milisegúndos | [optional] 

## Methods

### NewWalletsResponseObjectWalletsResponseWallets

`func NewWalletsResponseObjectWalletsResponseWallets() *WalletsResponseObjectWalletsResponseWallets`

NewWalletsResponseObjectWalletsResponseWallets instantiates a new WalletsResponseObjectWalletsResponseWallets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletsResponseObjectWalletsResponseWalletsWithDefaults

`func NewWalletsResponseObjectWalletsResponseWalletsWithDefaults() *WalletsResponseObjectWalletsResponseWallets`

NewWalletsResponseObjectWalletsResponseWalletsWithDefaults instantiates a new WalletsResponseObjectWalletsResponseWallets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentification

`func (o *WalletsResponseObjectWalletsResponseWallets) GetIdentification() string`

GetIdentification returns the Identification field if non-nil, zero value otherwise.

### GetIdentificationOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetIdentificationOk() (*string, bool)`

GetIdentificationOk returns a tuple with the Identification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentification

`func (o *WalletsResponseObjectWalletsResponseWallets) SetIdentification(v string)`

SetIdentification sets Identification field to given value.

### HasIdentification

`func (o *WalletsResponseObjectWalletsResponseWallets) HasIdentification() bool`

HasIdentification returns a boolean if a field has been set.

### GetName

`func (o *WalletsResponseObjectWalletsResponseWallets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletsResponseObjectWalletsResponseWallets) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WalletsResponseObjectWalletsResponseWallets) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *WalletsResponseObjectWalletsResponseWallets) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WalletsResponseObjectWalletsResponseWallets) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WalletsResponseObjectWalletsResponseWallets) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetImage

`func (o *WalletsResponseObjectWalletsResponseWallets) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WalletsResponseObjectWalletsResponseWallets) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *WalletsResponseObjectWalletsResponseWallets) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLargeImage

`func (o *WalletsResponseObjectWalletsResponseWallets) GetLargeImage() string`

GetLargeImage returns the LargeImage field if non-nil, zero value otherwise.

### GetLargeImageOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetLargeImageOk() (*string, bool)`

GetLargeImageOk returns a tuple with the LargeImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeImage

`func (o *WalletsResponseObjectWalletsResponseWallets) SetLargeImage(v string)`

SetLargeImage sets LargeImage field to given value.

### HasLargeImage

`func (o *WalletsResponseObjectWalletsResponseWallets) HasLargeImage() bool`

HasLargeImage returns a boolean if a field has been set.

### GetSmallImage

`func (o *WalletsResponseObjectWalletsResponseWallets) GetSmallImage() string`

GetSmallImage returns the SmallImage field if non-nil, zero value otherwise.

### GetSmallImageOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetSmallImageOk() (*string, bool)`

GetSmallImageOk returns a tuple with the SmallImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallImage

`func (o *WalletsResponseObjectWalletsResponseWallets) SetSmallImage(v string)`

SetSmallImage sets SmallImage field to given value.

### HasSmallImage

`func (o *WalletsResponseObjectWalletsResponseWallets) HasSmallImage() bool`

HasSmallImage returns a boolean if a field has been set.

### GetForeignIdentifier

`func (o *WalletsResponseObjectWalletsResponseWallets) GetForeignIdentifier() string`

GetForeignIdentifier returns the ForeignIdentifier field if non-nil, zero value otherwise.

### GetForeignIdentifierOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetForeignIdentifierOk() (*string, bool)`

GetForeignIdentifierOk returns a tuple with the ForeignIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignIdentifier

`func (o *WalletsResponseObjectWalletsResponseWallets) SetForeignIdentifier(v string)`

SetForeignIdentifier sets ForeignIdentifier field to given value.

### HasForeignIdentifier

`func (o *WalletsResponseObjectWalletsResponseWallets) HasForeignIdentifier() bool`

HasForeignIdentifier returns a boolean if a field has been set.

### GetAutoConfirm

`func (o *WalletsResponseObjectWalletsResponseWallets) GetAutoConfirm() bool`

GetAutoConfirm returns the AutoConfirm field if non-nil, zero value otherwise.

### GetAutoConfirmOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetAutoConfirmOk() (*bool, bool)`

GetAutoConfirmOk returns a tuple with the AutoConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConfirm

`func (o *WalletsResponseObjectWalletsResponseWallets) SetAutoConfirm(v bool)`

SetAutoConfirm sets AutoConfirm field to given value.

### HasAutoConfirm

`func (o *WalletsResponseObjectWalletsResponseWallets) HasAutoConfirm() bool`

HasAutoConfirm returns a boolean if a field has been set.

### GetSupportRequestCancel

`func (o *WalletsResponseObjectWalletsResponseWallets) GetSupportRequestCancel() bool`

GetSupportRequestCancel returns the SupportRequestCancel field if non-nil, zero value otherwise.

### GetSupportRequestCancelOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetSupportRequestCancelOk() (*bool, bool)`

GetSupportRequestCancelOk returns a tuple with the SupportRequestCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRequestCancel

`func (o *WalletsResponseObjectWalletsResponseWallets) SetSupportRequestCancel(v bool)`

SetSupportRequestCancel sets SupportRequestCancel field to given value.

### HasSupportRequestCancel

`func (o *WalletsResponseObjectWalletsResponseWallets) HasSupportRequestCancel() bool`

HasSupportRequestCancel returns a boolean if a field has been set.

### GetSupportValidityOfTheRequest

`func (o *WalletsResponseObjectWalletsResponseWallets) GetSupportValidityOfTheRequest() bool`

GetSupportValidityOfTheRequest returns the SupportValidityOfTheRequest field if non-nil, zero value otherwise.

### GetSupportValidityOfTheRequestOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetSupportValidityOfTheRequestOk() (*bool, bool)`

GetSupportValidityOfTheRequestOk returns a tuple with the SupportValidityOfTheRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportValidityOfTheRequest

`func (o *WalletsResponseObjectWalletsResponseWallets) SetSupportValidityOfTheRequest(v bool)`

SetSupportValidityOfTheRequest sets SupportValidityOfTheRequest field to given value.

### HasSupportValidityOfTheRequest

`func (o *WalletsResponseObjectWalletsResponseWallets) HasSupportValidityOfTheRequest() bool`

HasSupportValidityOfTheRequest returns a boolean if a field has been set.

### GetVoidSupport

`func (o *WalletsResponseObjectWalletsResponseWallets) GetVoidSupport() string`

GetVoidSupport returns the VoidSupport field if non-nil, zero value otherwise.

### GetVoidSupportOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetVoidSupportOk() (*string, bool)`

GetVoidSupportOk returns a tuple with the VoidSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidSupport

`func (o *WalletsResponseObjectWalletsResponseWallets) SetVoidSupport(v string)`

SetVoidSupport sets VoidSupport field to given value.

### HasVoidSupport

`func (o *WalletsResponseObjectWalletsResponseWallets) HasVoidSupport() bool`

HasVoidSupport returns a boolean if a field has been set.

### GetWalletUseInVoidTransaction

`func (o *WalletsResponseObjectWalletsResponseWallets) GetWalletUseInVoidTransaction() bool`

GetWalletUseInVoidTransaction returns the WalletUseInVoidTransaction field if non-nil, zero value otherwise.

### GetWalletUseInVoidTransactionOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetWalletUseInVoidTransactionOk() (*bool, bool)`

GetWalletUseInVoidTransactionOk returns a tuple with the WalletUseInVoidTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletUseInVoidTransaction

`func (o *WalletsResponseObjectWalletsResponseWallets) SetWalletUseInVoidTransaction(v bool)`

SetWalletUseInVoidTransaction sets WalletUseInVoidTransaction field to given value.

### HasWalletUseInVoidTransaction

`func (o *WalletsResponseObjectWalletsResponseWallets) HasWalletUseInVoidTransaction() bool`

HasWalletUseInVoidTransaction returns a boolean if a field has been set.

### GetReturnSupport

`func (o *WalletsResponseObjectWalletsResponseWallets) GetReturnSupport() string`

GetReturnSupport returns the ReturnSupport field if non-nil, zero value otherwise.

### GetReturnSupportOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetReturnSupportOk() (*string, bool)`

GetReturnSupportOk returns a tuple with the ReturnSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnSupport

`func (o *WalletsResponseObjectWalletsResponseWallets) SetReturnSupport(v string)`

SetReturnSupport sets ReturnSupport field to given value.

### HasReturnSupport

`func (o *WalletsResponseObjectWalletsResponseWallets) HasReturnSupport() bool`

HasReturnSupport returns a boolean if a field has been set.

### GetWalletUseInReturnTransaction

`func (o *WalletsResponseObjectWalletsResponseWallets) GetWalletUseInReturnTransaction() bool`

GetWalletUseInReturnTransaction returns the WalletUseInReturnTransaction field if non-nil, zero value otherwise.

### GetWalletUseInReturnTransactionOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetWalletUseInReturnTransactionOk() (*bool, bool)`

GetWalletUseInReturnTransactionOk returns a tuple with the WalletUseInReturnTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletUseInReturnTransaction

`func (o *WalletsResponseObjectWalletsResponseWallets) SetWalletUseInReturnTransaction(v bool)`

SetWalletUseInReturnTransaction sets WalletUseInReturnTransaction field to given value.

### HasWalletUseInReturnTransaction

`func (o *WalletsResponseObjectWalletsResponseWallets) HasWalletUseInReturnTransaction() bool`

HasWalletUseInReturnTransaction returns a boolean if a field has been set.

### GetAmountRequired

`func (o *WalletsResponseObjectWalletsResponseWallets) GetAmountRequired() bool`

GetAmountRequired returns the AmountRequired field if non-nil, zero value otherwise.

### GetAmountRequiredOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetAmountRequiredOk() (*bool, bool)`

GetAmountRequiredOk returns a tuple with the AmountRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRequired

`func (o *WalletsResponseObjectWalletsResponseWallets) SetAmountRequired(v bool)`

SetAmountRequired sets AmountRequired field to given value.

### HasAmountRequired

`func (o *WalletsResponseObjectWalletsResponseWallets) HasAmountRequired() bool`

HasAmountRequired returns a boolean if a field has been set.

### GetTokenType

`func (o *WalletsResponseObjectWalletsResponseWallets) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *WalletsResponseObjectWalletsResponseWallets) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *WalletsResponseObjectWalletsResponseWallets) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetDefaultRetryTime

`func (o *WalletsResponseObjectWalletsResponseWallets) GetDefaultRetryTime() int32`

GetDefaultRetryTime returns the DefaultRetryTime field if non-nil, zero value otherwise.

### GetDefaultRetryTimeOk

`func (o *WalletsResponseObjectWalletsResponseWallets) GetDefaultRetryTimeOk() (*int32, bool)`

GetDefaultRetryTimeOk returns a tuple with the DefaultRetryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRetryTime

`func (o *WalletsResponseObjectWalletsResponseWallets) SetDefaultRetryTime(v int32)`

SetDefaultRetryTime sets DefaultRetryTime field to given value.

### HasDefaultRetryTime

`func (o *WalletsResponseObjectWalletsResponseWallets) HasDefaultRetryTime() bool`

HasDefaultRetryTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


