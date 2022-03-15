# DebtInquiryObjectDebtInquiry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyIdentification** | **string** | ID que identifica la companía desde donde proviene la petición. | 
**SystemIdentification** | **string** | ID que identifica el sistema desde donde proviene la petición. | 
**BranchIdentification** | Pointer to **string** | ID que identifica la sucursal desde donde proviene la petición. Esta sucursal pertenece a una determinada companía. | [optional] 
**POSIdentification** | Pointer to **string** | ID que identifica la caja o punto de venta desde donde proviene la petición. Este punto de venta pertenece a una determinada sucursal y companía. | [optional] 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 
**Block** | Pointer to **string** | ID que identifica a un grupo de transacciones que serán confirmadas o canceladas | [optional] 
**ReasonSequenceBreak** | Pointer to **string** | Motivo por el cual se requiere romper la secuencia. | [optional] 
**ReadingDeviceType** | Pointer to **string** | Tipo de dispositivo utilizado para ingresar los datos de la tarjeta. En función al dispositivo usado, serán realizadas ciertas verificaciones, por lo que ciertos datos serán requeridos. CustomerKeyboard, utilizado para ingreso manual de tarjeta a través de un portal web, por ejemplo - Keyboard, utilizado para ingreso manual de la tarjeta por parte del vendedor - MagneticStripReader, lector de banda de tarjetas por emulación de teclado, u otro valor configurado  que indentifique el dispositivo que se esta utilizando. | [optional] 
**ReadingDeviceOperatingFrom** | Pointer to **time.Time** | Indica desde cuando se encuentra operativo o encendido el dispositivo | [optional] 
**CardAppName** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAppIdentifier** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAppLabel** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAuthRequestCryptogram** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAuthResponseCryptogram** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardReadMode** | Pointer to **string** | Modo de ingreso de los datos de la tarjeta. Los posibles valores significan: C - EMV Chip / B - Banda magnética / L - Contactless Chip / S - Contactless Banda / M - Manual (Tarjeta Presente) / T - Digitada (Tarjeta no Presente) / E - ECOMMERCE (Ventas por Internet)  / F - FALLBACK (Banda por falla en Chip) / K - TOKEN / R - Recurring ( Pagos Recurrentes ) | [optional] 
**CardGetMode** | Pointer to **string** | Indican por cada elemento que contiene los datos sensibles, si están encriptados  y también el algoritmo usado. En caso de no estar especificado se asume PLAIN. | [optional] 
**CardNumber** | Pointer to **string** | Número de Tarjeta, en el caso de las respuestas el mismo estará enmascarado. | [optional] 
**CardNumberMasked** | Pointer to **string** | Número de tarjeta enmascarado, según indica la parametrización en la base de datos. Se utilizará para imprimir en el cupón.   | [optional] 
**CardNumberEncrypted** | Pointer to **string** | Número de tarjeta encriptado. | [optional] 
**CardExp** | Pointer to **string** | Fecha de vencimiento de la tarjeta. Este dato sera necesario si el modo de ingreso fue manual/digitada. | [optional] 
**Track1** | Pointer to **string** | Banda Número 1 de la tarjeta. Este dato sera necesario si el modo de ingreso fue por banda (deslizando la banda por el lector). | [optional] 
**Track2** | Pointer to **string** | Banda Número 2 de la tarjeta. Este dato sera necesario si el modo de ingreso fue por banda (deslizando la banda por el lector). | [optional] 
**InputTokens** | Pointer to [**[]SaleObjectSaleInputTokens**](SaleObjectSaleInputTokens.md) | Tokens. | [optional] 
**SecurityCode** | Pointer to **string** | Código de seguridad de la tarjeta. | [optional] 
**Pin** | Pointer to **string** | PIN block | [optional] 
**CardLastFourDigits** | Pointer to **string** | últimos 4 digitos de la tarjeta, ingresado a partir de lo que el vendedor verifica en la superficie de la tarjeta. Este dato solo sera requerido para ciertos tipos de dispositivos de capturas, como por ejemplo lectores de tarjetas por emulación de teclado, y siempre y cuando la parametrización de la marca de la tarjeta asi lo indique. | [optional] 
**CredentialToken** | Pointer to **string** | Token asociado a la Credencial Enrolada | [optional] 
**CredentialIssuerToken** | Pointer to **string** | Emisor del Token asociado a la credencial enrolada | [optional] 
**Payer** | Pointer to [**SaleObjectSalePayer**](SaleObjectSalePayer.md) |  | [optional] 
**Customer** | Pointer to [**SaleObjectSaleCustomer**](SaleObjectSaleCustomer.md) |  | [optional] 
**PaymentFacilitatorID** | Pointer to **string** | Identificador de facilitador de pagos o Payfac. | [optional] 
**MerchantID** | Pointer to **string** | Número de comercio utilizado para realizar la transacción. Este Número es asignado por el host, y parametrizado en la BD, relacionado a cada uno de los planes disponibles. | [optional] 
**TerminalID** | Pointer to **string** | Identificador de Terminal por el cual se envía la Transacción al Host. | [optional] 
**TerminalTrace** | Pointer to **int32** | Número de Trace/Secuencia que genera la plataforma para la transacción asociado al TerminalID. | [optional] 
**SettlementBatchNumber** | Pointer to **int32** | Para aquellos host que exista el concepto de lote, es el número de lote al cual pertenece la transacción. | [optional] 

## Methods

### NewDebtInquiryObjectDebtInquiry

`func NewDebtInquiryObjectDebtInquiry(companyIdentification string, systemIdentification string, ) *DebtInquiryObjectDebtInquiry`

NewDebtInquiryObjectDebtInquiry instantiates a new DebtInquiryObjectDebtInquiry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebtInquiryObjectDebtInquiryWithDefaults

`func NewDebtInquiryObjectDebtInquiryWithDefaults() *DebtInquiryObjectDebtInquiry`

NewDebtInquiryObjectDebtInquiryWithDefaults instantiates a new DebtInquiryObjectDebtInquiry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyIdentification

`func (o *DebtInquiryObjectDebtInquiry) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *DebtInquiryObjectDebtInquiry) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *DebtInquiryObjectDebtInquiry) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.


### GetSystemIdentification

`func (o *DebtInquiryObjectDebtInquiry) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *DebtInquiryObjectDebtInquiry) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *DebtInquiryObjectDebtInquiry) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.


### GetBranchIdentification

`func (o *DebtInquiryObjectDebtInquiry) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *DebtInquiryObjectDebtInquiry) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *DebtInquiryObjectDebtInquiry) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *DebtInquiryObjectDebtInquiry) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *DebtInquiryObjectDebtInquiry) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *DebtInquiryObjectDebtInquiry) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *DebtInquiryObjectDebtInquiry) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *DebtInquiryObjectDebtInquiry) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetServiceVersion

`func (o *DebtInquiryObjectDebtInquiry) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *DebtInquiryObjectDebtInquiry) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *DebtInquiryObjectDebtInquiry) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *DebtInquiryObjectDebtInquiry) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *DebtInquiryObjectDebtInquiry) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *DebtInquiryObjectDebtInquiry) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *DebtInquiryObjectDebtInquiry) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *DebtInquiryObjectDebtInquiry) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *DebtInquiryObjectDebtInquiry) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *DebtInquiryObjectDebtInquiry) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *DebtInquiryObjectDebtInquiry) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *DebtInquiryObjectDebtInquiry) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetBlock

`func (o *DebtInquiryObjectDebtInquiry) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *DebtInquiryObjectDebtInquiry) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *DebtInquiryObjectDebtInquiry) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *DebtInquiryObjectDebtInquiry) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetReasonSequenceBreak

`func (o *DebtInquiryObjectDebtInquiry) GetReasonSequenceBreak() string`

GetReasonSequenceBreak returns the ReasonSequenceBreak field if non-nil, zero value otherwise.

### GetReasonSequenceBreakOk

`func (o *DebtInquiryObjectDebtInquiry) GetReasonSequenceBreakOk() (*string, bool)`

GetReasonSequenceBreakOk returns a tuple with the ReasonSequenceBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonSequenceBreak

`func (o *DebtInquiryObjectDebtInquiry) SetReasonSequenceBreak(v string)`

SetReasonSequenceBreak sets ReasonSequenceBreak field to given value.

### HasReasonSequenceBreak

`func (o *DebtInquiryObjectDebtInquiry) HasReasonSequenceBreak() bool`

HasReasonSequenceBreak returns a boolean if a field has been set.

### GetReadingDeviceType

`func (o *DebtInquiryObjectDebtInquiry) GetReadingDeviceType() string`

GetReadingDeviceType returns the ReadingDeviceType field if non-nil, zero value otherwise.

### GetReadingDeviceTypeOk

`func (o *DebtInquiryObjectDebtInquiry) GetReadingDeviceTypeOk() (*string, bool)`

GetReadingDeviceTypeOk returns a tuple with the ReadingDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceType

`func (o *DebtInquiryObjectDebtInquiry) SetReadingDeviceType(v string)`

SetReadingDeviceType sets ReadingDeviceType field to given value.

### HasReadingDeviceType

`func (o *DebtInquiryObjectDebtInquiry) HasReadingDeviceType() bool`

HasReadingDeviceType returns a boolean if a field has been set.

### GetReadingDeviceOperatingFrom

`func (o *DebtInquiryObjectDebtInquiry) GetReadingDeviceOperatingFrom() time.Time`

GetReadingDeviceOperatingFrom returns the ReadingDeviceOperatingFrom field if non-nil, zero value otherwise.

### GetReadingDeviceOperatingFromOk

`func (o *DebtInquiryObjectDebtInquiry) GetReadingDeviceOperatingFromOk() (*time.Time, bool)`

GetReadingDeviceOperatingFromOk returns a tuple with the ReadingDeviceOperatingFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceOperatingFrom

`func (o *DebtInquiryObjectDebtInquiry) SetReadingDeviceOperatingFrom(v time.Time)`

SetReadingDeviceOperatingFrom sets ReadingDeviceOperatingFrom field to given value.

### HasReadingDeviceOperatingFrom

`func (o *DebtInquiryObjectDebtInquiry) HasReadingDeviceOperatingFrom() bool`

HasReadingDeviceOperatingFrom returns a boolean if a field has been set.

### GetCardAppName

`func (o *DebtInquiryObjectDebtInquiry) GetCardAppName() string`

GetCardAppName returns the CardAppName field if non-nil, zero value otherwise.

### GetCardAppNameOk

`func (o *DebtInquiryObjectDebtInquiry) GetCardAppNameOk() (*string, bool)`

GetCardAppNameOk returns a tuple with the CardAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppName

`func (o *DebtInquiryObjectDebtInquiry) SetCardAppName(v string)`

SetCardAppName sets CardAppName field to given value.

### HasCardAppName

`func (o *DebtInquiryObjectDebtInquiry) HasCardAppName() bool`

HasCardAppName returns a boolean if a field has been set.

### GetCardAppIdentifier

`func (o *DebtInquiryObjectDebtInquiry) GetCardAppIdentifier() string`

GetCardAppIdentifier returns the CardAppIdentifier field if non-nil, zero value otherwise.

### GetCardAppIdentifierOk

`func (o *DebtInquiryObjectDebtInquiry) GetCardAppIdentifierOk() (*string, bool)`

GetCardAppIdentifierOk returns a tuple with the CardAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppIdentifier

`func (o *DebtInquiryObjectDebtInquiry) SetCardAppIdentifier(v string)`

SetCardAppIdentifier sets CardAppIdentifier field to given value.

### HasCardAppIdentifier

`func (o *DebtInquiryObjectDebtInquiry) HasCardAppIdentifier() bool`

HasCardAppIdentifier returns a boolean if a field has been set.

### GetCardAppLabel

`func (o *DebtInquiryObjectDebtInquiry) GetCardAppLabel() string`

GetCardAppLabel returns the CardAppLabel field if non-nil, zero value otherwise.

### GetCardAppLabelOk

`func (o *DebtInquiryObjectDebtInquiry) GetCardAppLabelOk() (*string, bool)`

GetCardAppLabelOk returns a tuple with the CardAppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppLabel

`func (o *DebtInquiryObjectDebtInquiry) SetCardAppLabel(v string)`

SetCardAppLabel sets CardAppLabel field to given value.

### HasCardAppLabel

`func (o *DebtInquiryObjectDebtInquiry) HasCardAppLabel() bool`

HasCardAppLabel returns a boolean if a field has been set.

### GetCardAuthRequestCryptogram

`func (o *DebtInquiryObjectDebtInquiry) GetCardAuthRequestCryptogram() string`

GetCardAuthRequestCryptogram returns the CardAuthRequestCryptogram field if non-nil, zero value otherwise.

### GetCardAuthRequestCryptogramOk

`func (o *DebtInquiryObjectDebtInquiry) GetCardAuthRequestCryptogramOk() (*string, bool)`

GetCardAuthRequestCryptogramOk returns a tuple with the CardAuthRequestCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthRequestCryptogram

`func (o *DebtInquiryObjectDebtInquiry) SetCardAuthRequestCryptogram(v string)`

SetCardAuthRequestCryptogram sets CardAuthRequestCryptogram field to given value.

### HasCardAuthRequestCryptogram

`func (o *DebtInquiryObjectDebtInquiry) HasCardAuthRequestCryptogram() bool`

HasCardAuthRequestCryptogram returns a boolean if a field has been set.

### GetCardAuthResponseCryptogram

`func (o *DebtInquiryObjectDebtInquiry) GetCardAuthResponseCryptogram() string`

GetCardAuthResponseCryptogram returns the CardAuthResponseCryptogram field if non-nil, zero value otherwise.

### GetCardAuthResponseCryptogramOk

`func (o *DebtInquiryObjectDebtInquiry) GetCardAuthResponseCryptogramOk() (*string, bool)`

GetCardAuthResponseCryptogramOk returns a tuple with the CardAuthResponseCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthResponseCryptogram

`func (o *DebtInquiryObjectDebtInquiry) SetCardAuthResponseCryptogram(v string)`

SetCardAuthResponseCryptogram sets CardAuthResponseCryptogram field to given value.

### HasCardAuthResponseCryptogram

`func (o *DebtInquiryObjectDebtInquiry) HasCardAuthResponseCryptogram() bool`

HasCardAuthResponseCryptogram returns a boolean if a field has been set.

### GetCardReadMode

`func (o *DebtInquiryObjectDebtInquiry) GetCardReadMode() string`

GetCardReadMode returns the CardReadMode field if non-nil, zero value otherwise.

### GetCardReadModeOk

`func (o *DebtInquiryObjectDebtInquiry) GetCardReadModeOk() (*string, bool)`

GetCardReadModeOk returns a tuple with the CardReadMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardReadMode

`func (o *DebtInquiryObjectDebtInquiry) SetCardReadMode(v string)`

SetCardReadMode sets CardReadMode field to given value.

### HasCardReadMode

`func (o *DebtInquiryObjectDebtInquiry) HasCardReadMode() bool`

HasCardReadMode returns a boolean if a field has been set.

### GetCardGetMode

`func (o *DebtInquiryObjectDebtInquiry) GetCardGetMode() string`

GetCardGetMode returns the CardGetMode field if non-nil, zero value otherwise.

### GetCardGetModeOk

`func (o *DebtInquiryObjectDebtInquiry) GetCardGetModeOk() (*string, bool)`

GetCardGetModeOk returns a tuple with the CardGetMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardGetMode

`func (o *DebtInquiryObjectDebtInquiry) SetCardGetMode(v string)`

SetCardGetMode sets CardGetMode field to given value.

### HasCardGetMode

`func (o *DebtInquiryObjectDebtInquiry) HasCardGetMode() bool`

HasCardGetMode returns a boolean if a field has been set.

### GetCardNumber

`func (o *DebtInquiryObjectDebtInquiry) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *DebtInquiryObjectDebtInquiry) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *DebtInquiryObjectDebtInquiry) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *DebtInquiryObjectDebtInquiry) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNumberMasked

`func (o *DebtInquiryObjectDebtInquiry) GetCardNumberMasked() string`

GetCardNumberMasked returns the CardNumberMasked field if non-nil, zero value otherwise.

### GetCardNumberMaskedOk

`func (o *DebtInquiryObjectDebtInquiry) GetCardNumberMaskedOk() (*string, bool)`

GetCardNumberMaskedOk returns a tuple with the CardNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberMasked

`func (o *DebtInquiryObjectDebtInquiry) SetCardNumberMasked(v string)`

SetCardNumberMasked sets CardNumberMasked field to given value.

### HasCardNumberMasked

`func (o *DebtInquiryObjectDebtInquiry) HasCardNumberMasked() bool`

HasCardNumberMasked returns a boolean if a field has been set.

### GetCardNumberEncrypted

`func (o *DebtInquiryObjectDebtInquiry) GetCardNumberEncrypted() string`

GetCardNumberEncrypted returns the CardNumberEncrypted field if non-nil, zero value otherwise.

### GetCardNumberEncryptedOk

`func (o *DebtInquiryObjectDebtInquiry) GetCardNumberEncryptedOk() (*string, bool)`

GetCardNumberEncryptedOk returns a tuple with the CardNumberEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberEncrypted

`func (o *DebtInquiryObjectDebtInquiry) SetCardNumberEncrypted(v string)`

SetCardNumberEncrypted sets CardNumberEncrypted field to given value.

### HasCardNumberEncrypted

`func (o *DebtInquiryObjectDebtInquiry) HasCardNumberEncrypted() bool`

HasCardNumberEncrypted returns a boolean if a field has been set.

### GetCardExp

`func (o *DebtInquiryObjectDebtInquiry) GetCardExp() string`

GetCardExp returns the CardExp field if non-nil, zero value otherwise.

### GetCardExpOk

`func (o *DebtInquiryObjectDebtInquiry) GetCardExpOk() (*string, bool)`

GetCardExpOk returns a tuple with the CardExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExp

`func (o *DebtInquiryObjectDebtInquiry) SetCardExp(v string)`

SetCardExp sets CardExp field to given value.

### HasCardExp

`func (o *DebtInquiryObjectDebtInquiry) HasCardExp() bool`

HasCardExp returns a boolean if a field has been set.

### GetTrack1

`func (o *DebtInquiryObjectDebtInquiry) GetTrack1() string`

GetTrack1 returns the Track1 field if non-nil, zero value otherwise.

### GetTrack1Ok

`func (o *DebtInquiryObjectDebtInquiry) GetTrack1Ok() (*string, bool)`

GetTrack1Ok returns a tuple with the Track1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack1

`func (o *DebtInquiryObjectDebtInquiry) SetTrack1(v string)`

SetTrack1 sets Track1 field to given value.

### HasTrack1

`func (o *DebtInquiryObjectDebtInquiry) HasTrack1() bool`

HasTrack1 returns a boolean if a field has been set.

### GetTrack2

`func (o *DebtInquiryObjectDebtInquiry) GetTrack2() string`

GetTrack2 returns the Track2 field if non-nil, zero value otherwise.

### GetTrack2Ok

`func (o *DebtInquiryObjectDebtInquiry) GetTrack2Ok() (*string, bool)`

GetTrack2Ok returns a tuple with the Track2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack2

`func (o *DebtInquiryObjectDebtInquiry) SetTrack2(v string)`

SetTrack2 sets Track2 field to given value.

### HasTrack2

`func (o *DebtInquiryObjectDebtInquiry) HasTrack2() bool`

HasTrack2 returns a boolean if a field has been set.

### GetInputTokens

`func (o *DebtInquiryObjectDebtInquiry) GetInputTokens() []SaleObjectSaleInputTokens`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *DebtInquiryObjectDebtInquiry) GetInputTokensOk() (*[]SaleObjectSaleInputTokens, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *DebtInquiryObjectDebtInquiry) SetInputTokens(v []SaleObjectSaleInputTokens)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *DebtInquiryObjectDebtInquiry) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetSecurityCode

`func (o *DebtInquiryObjectDebtInquiry) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *DebtInquiryObjectDebtInquiry) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *DebtInquiryObjectDebtInquiry) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *DebtInquiryObjectDebtInquiry) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### GetPin

`func (o *DebtInquiryObjectDebtInquiry) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *DebtInquiryObjectDebtInquiry) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *DebtInquiryObjectDebtInquiry) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *DebtInquiryObjectDebtInquiry) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetCardLastFourDigits

`func (o *DebtInquiryObjectDebtInquiry) GetCardLastFourDigits() string`

GetCardLastFourDigits returns the CardLastFourDigits field if non-nil, zero value otherwise.

### GetCardLastFourDigitsOk

`func (o *DebtInquiryObjectDebtInquiry) GetCardLastFourDigitsOk() (*string, bool)`

GetCardLastFourDigitsOk returns a tuple with the CardLastFourDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLastFourDigits

`func (o *DebtInquiryObjectDebtInquiry) SetCardLastFourDigits(v string)`

SetCardLastFourDigits sets CardLastFourDigits field to given value.

### HasCardLastFourDigits

`func (o *DebtInquiryObjectDebtInquiry) HasCardLastFourDigits() bool`

HasCardLastFourDigits returns a boolean if a field has been set.

### GetCredentialToken

`func (o *DebtInquiryObjectDebtInquiry) GetCredentialToken() string`

GetCredentialToken returns the CredentialToken field if non-nil, zero value otherwise.

### GetCredentialTokenOk

`func (o *DebtInquiryObjectDebtInquiry) GetCredentialTokenOk() (*string, bool)`

GetCredentialTokenOk returns a tuple with the CredentialToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialToken

`func (o *DebtInquiryObjectDebtInquiry) SetCredentialToken(v string)`

SetCredentialToken sets CredentialToken field to given value.

### HasCredentialToken

`func (o *DebtInquiryObjectDebtInquiry) HasCredentialToken() bool`

HasCredentialToken returns a boolean if a field has been set.

### GetCredentialIssuerToken

`func (o *DebtInquiryObjectDebtInquiry) GetCredentialIssuerToken() string`

GetCredentialIssuerToken returns the CredentialIssuerToken field if non-nil, zero value otherwise.

### GetCredentialIssuerTokenOk

`func (o *DebtInquiryObjectDebtInquiry) GetCredentialIssuerTokenOk() (*string, bool)`

GetCredentialIssuerTokenOk returns a tuple with the CredentialIssuerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIssuerToken

`func (o *DebtInquiryObjectDebtInquiry) SetCredentialIssuerToken(v string)`

SetCredentialIssuerToken sets CredentialIssuerToken field to given value.

### HasCredentialIssuerToken

`func (o *DebtInquiryObjectDebtInquiry) HasCredentialIssuerToken() bool`

HasCredentialIssuerToken returns a boolean if a field has been set.

### GetPayer

`func (o *DebtInquiryObjectDebtInquiry) GetPayer() SaleObjectSalePayer`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *DebtInquiryObjectDebtInquiry) GetPayerOk() (*SaleObjectSalePayer, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *DebtInquiryObjectDebtInquiry) SetPayer(v SaleObjectSalePayer)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *DebtInquiryObjectDebtInquiry) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetCustomer

`func (o *DebtInquiryObjectDebtInquiry) GetCustomer() SaleObjectSaleCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *DebtInquiryObjectDebtInquiry) GetCustomerOk() (*SaleObjectSaleCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *DebtInquiryObjectDebtInquiry) SetCustomer(v SaleObjectSaleCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *DebtInquiryObjectDebtInquiry) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetPaymentFacilitatorID

`func (o *DebtInquiryObjectDebtInquiry) GetPaymentFacilitatorID() string`

GetPaymentFacilitatorID returns the PaymentFacilitatorID field if non-nil, zero value otherwise.

### GetPaymentFacilitatorIDOk

`func (o *DebtInquiryObjectDebtInquiry) GetPaymentFacilitatorIDOk() (*string, bool)`

GetPaymentFacilitatorIDOk returns a tuple with the PaymentFacilitatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFacilitatorID

`func (o *DebtInquiryObjectDebtInquiry) SetPaymentFacilitatorID(v string)`

SetPaymentFacilitatorID sets PaymentFacilitatorID field to given value.

### HasPaymentFacilitatorID

`func (o *DebtInquiryObjectDebtInquiry) HasPaymentFacilitatorID() bool`

HasPaymentFacilitatorID returns a boolean if a field has been set.

### GetMerchantID

`func (o *DebtInquiryObjectDebtInquiry) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *DebtInquiryObjectDebtInquiry) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *DebtInquiryObjectDebtInquiry) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *DebtInquiryObjectDebtInquiry) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetTerminalID

`func (o *DebtInquiryObjectDebtInquiry) GetTerminalID() string`

GetTerminalID returns the TerminalID field if non-nil, zero value otherwise.

### GetTerminalIDOk

`func (o *DebtInquiryObjectDebtInquiry) GetTerminalIDOk() (*string, bool)`

GetTerminalIDOk returns a tuple with the TerminalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalID

`func (o *DebtInquiryObjectDebtInquiry) SetTerminalID(v string)`

SetTerminalID sets TerminalID field to given value.

### HasTerminalID

`func (o *DebtInquiryObjectDebtInquiry) HasTerminalID() bool`

HasTerminalID returns a boolean if a field has been set.

### GetTerminalTrace

`func (o *DebtInquiryObjectDebtInquiry) GetTerminalTrace() int32`

GetTerminalTrace returns the TerminalTrace field if non-nil, zero value otherwise.

### GetTerminalTraceOk

`func (o *DebtInquiryObjectDebtInquiry) GetTerminalTraceOk() (*int32, bool)`

GetTerminalTraceOk returns a tuple with the TerminalTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalTrace

`func (o *DebtInquiryObjectDebtInquiry) SetTerminalTrace(v int32)`

SetTerminalTrace sets TerminalTrace field to given value.

### HasTerminalTrace

`func (o *DebtInquiryObjectDebtInquiry) HasTerminalTrace() bool`

HasTerminalTrace returns a boolean if a field has been set.

### GetSettlementBatchNumber

`func (o *DebtInquiryObjectDebtInquiry) GetSettlementBatchNumber() int32`

GetSettlementBatchNumber returns the SettlementBatchNumber field if non-nil, zero value otherwise.

### GetSettlementBatchNumberOk

`func (o *DebtInquiryObjectDebtInquiry) GetSettlementBatchNumberOk() (*int32, bool)`

GetSettlementBatchNumberOk returns a tuple with the SettlementBatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementBatchNumber

`func (o *DebtInquiryObjectDebtInquiry) SetSettlementBatchNumber(v int32)`

SetSettlementBatchNumber sets SettlementBatchNumber field to given value.

### HasSettlementBatchNumber

`func (o *DebtInquiryObjectDebtInquiry) HasSettlementBatchNumber() bool`

HasSettlementBatchNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


