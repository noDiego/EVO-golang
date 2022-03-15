# GetTransactionResponseObjectGetTransactionResponseTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseActions** | Pointer to **[]string** | Acciones a realizar por parte del POS y/o PINPAD en base al resultado de la operación que ha sido procesada. Cada uno de estos actions o acciones están concatenadas por comas. Los posibles actions son OK, OK, NotExist, ConfigurationError, SystemError, ResourceError, ProcessError, Configure. | [optional] 
**ResponseMessage** | Pointer to **string** | Descripción del resultado del proceso del requerimiento recibido. Esta descripción es generada por la plataforma, no por el Host que termine resolviendo la transacción. | [optional] 
**ForeignResponseCode** | Pointer to **string** | Código de respuesta para el sistema externo, es decir, para la aplicación cliente que se comunica con el TEF. | [optional] 
**ResponseCode** | Pointer to **int32** | Código de Respuesta Interno de la plataforma, el POS debe actuar por lo que indican las acciones especificadas en ResponseActions y no por el código de respuesta informado en este campo o elemento, pero es una buena práctica que sea persistido por el mismo. | [optional] 
**TaxFinancialCostAmount** | Pointer to **float32** | Monto del recargo impositivo aplicado al costo financiero que la transacción tiene | [optional] 
**TaxFinancialCostPercentage** | Pointer to **float32** | Porcentaje de recargo impositivo a aplicar sobre el monto del costo financiero | [optional] 
**FinancialCostAmount** | Pointer to **float32** | Monto del Costo financiero total generado en base al plan elegido | [optional] 
**FinancialCostPercentage** | Pointer to **float32** | Porcentaje del costo financiero a aplicar al monto de la transacción | [optional] 
**RequestAmount** | Pointer to **float32** | Monto libre de costos financerios e impuestos por el que la venta se realizó. El monto cobrado realmente no es este, dado que no incluye las tasas e impuestos | [optional] 
**Amount** | Pointer to **float32** | Importe o Monto de la Transacción. | [optional] 
**AlternativeAmount** | Pointer to **float32** | Monto con la que se realizó transacción. Si este valor es recibido, la búsqueda de los planes será limitada con este criterio. | [optional] 
**HostResultCode** | Pointer to **int32** | Código de Resultado retornado por el Host Adquirente. | [optional] 
**HostMessage** | Pointer to **string** | Mensaje Retornado por el Host Adquirente, normalmente asociado al valor de HostResultCode | [optional] 
**HostCode** | Pointer to **string** | código de autorización retornado por el Host que resuelve la transacción. | [optional] 
**HostDateTime** | Pointer to **time.Time** | Fecha y Hora de la transacción retornada por el Host que resuelve la Transacción - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**TransmitionDateTime** | Pointer to **time.Time** | Fecha y hora de transmision de la operación hacia el host - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**AuthTicket** | Pointer to **int32** | Número Ticket  o Voucher Generado para la Plataforma. | [optional] 
**AuthRRN** | Pointer to **string** | Número de identificación de la transacción, utilizado por la mayoría de los hosts para realizar anulaciones y devoluciones | [optional] 
**IdentifierForTheAdquirer** | Pointer to **string** | Identificador que genera el Host Adquirente para la Transacción en algunos podrá ser igual al AuthRRN | [optional] 
**PaymentFacilitatorID** | Pointer to **string** | Identificador de facilitador de pagos o Payfac. | [optional] 
**MerchantID** | Pointer to **string** | Número de comercio utilizado para realizar la transacción. Este Número es asignado por el host, y parametrizado en la BD, relaciónado a cada uno de los planes disponibles. | [optional] 
**TerminalID** | Pointer to **int32** | Identificador de Terminal por el cual se envía la Transacción al Host. | [optional] 
**TerminalTrace** | Pointer to **int32** | Número de Trace/Secuencia que genera la plataforma para la transacción asociado al TerminalID. | [optional] 
**SettlementBatchNumber** | Pointer to **int32** | Para aquellos host que exista el concepto de Lote, es el número de lote al cual pertenece la transacción. | [optional] 
**HostID** | Pointer to **int32** | Número de identificación del host al cual fue enviada la petición, y por el cual fue finalmente procesada | [optional] 
**IssuerName** | Pointer to **string** | Nombre del Emisor de la Credencial o Tarjeta que se usó en la transacción. | [optional] 
**CardDescription** | Pointer to **string** | Nombre de la Tarjeta que se usó en la transacción, usado para la impresión del voucher. | [optional] 
**PaymentMethodDescription** | Pointer to **string** | Descripción o nombre de la marca con la cual la tarjeta fue identificada | [optional] 
**PlanDescription** | Pointer to **string** | Descripción del plan utilizado para para realizar la operación | [optional] 
**CardReadMode** | Pointer to **string** | Modo de ingreso de los datos de la tarjeta. Los posibles valores significan: C - EMV Chip / B - Banda magnética / L - Contactless Chip / S - Contactless Banda / M - Manual (Tarjeta Presente) / T - Digitada (Tarjeta no Presente) / E - ECOMMERCE (Ventas por Internet)  / F - FALLBACK (Banda por falla en Chip) / K - TOKEN / R - Recurring ( Pagos Recurrentes ) | [optional] 
**CardNumber** | Pointer to **string** | Número de Tarjeta, en el caso de las respuestas el mismo estará enmascarado. | [optional] 
**CardNumberMasked** | Pointer to **string** | Número de tarjeta enmascarado, según indica la parametrización en la base de datos. Se utilizará para imprimir en el cupón | [optional] 
**CardHashing** | Pointer to **string** | Hash de la tarjeta generado por la plataforma. | [optional] 
**CurrencyDescription** | Pointer to **string** | Descripción del tipo de cambio utilizado en la transacción | [optional] 
**CurrencySymbol** | Pointer to **string** | Simbolo monetario del tipo de cambio utilizado en la transacción | [optional] 
**CardCryptogramResponse** | Pointer to **string** | Tags EMV en format TLV recibidos desde el Host. | [optional] 
**CardAppName** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAppIdentifier** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAppLabel** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 

## Methods

### NewGetTransactionResponseObjectGetTransactionResponseTransaction

`func NewGetTransactionResponseObjectGetTransactionResponseTransaction() *GetTransactionResponseObjectGetTransactionResponseTransaction`

NewGetTransactionResponseObjectGetTransactionResponseTransaction instantiates a new GetTransactionResponseObjectGetTransactionResponseTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionResponseObjectGetTransactionResponseTransactionWithDefaults

`func NewGetTransactionResponseObjectGetTransactionResponseTransactionWithDefaults() *GetTransactionResponseObjectGetTransactionResponseTransaction`

NewGetTransactionResponseObjectGetTransactionResponseTransactionWithDefaults instantiates a new GetTransactionResponseObjectGetTransactionResponseTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseActions

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetResponseActions() []string`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetResponseActionsOk() (*[]string, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetResponseActions(v []string)`

SetResponseActions sets ResponseActions field to given value.

### HasResponseActions

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasResponseActions() bool`

HasResponseActions returns a boolean if a field has been set.

### GetResponseMessage

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.

### HasResponseMessage

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasResponseMessage() bool`

HasResponseMessage returns a boolean if a field has been set.

### GetForeignResponseCode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetForeignResponseCode() string`

GetForeignResponseCode returns the ForeignResponseCode field if non-nil, zero value otherwise.

### GetForeignResponseCodeOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetForeignResponseCodeOk() (*string, bool)`

GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignResponseCode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetForeignResponseCode(v string)`

SetForeignResponseCode sets ForeignResponseCode field to given value.

### HasForeignResponseCode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasForeignResponseCode() bool`

HasForeignResponseCode returns a boolean if a field has been set.

### GetResponseCode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### GetTaxFinancialCostAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTaxFinancialCostAmount() float32`

GetTaxFinancialCostAmount returns the TaxFinancialCostAmount field if non-nil, zero value otherwise.

### GetTaxFinancialCostAmountOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTaxFinancialCostAmountOk() (*float32, bool)`

GetTaxFinancialCostAmountOk returns a tuple with the TaxFinancialCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFinancialCostAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetTaxFinancialCostAmount(v float32)`

SetTaxFinancialCostAmount sets TaxFinancialCostAmount field to given value.

### HasTaxFinancialCostAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasTaxFinancialCostAmount() bool`

HasTaxFinancialCostAmount returns a boolean if a field has been set.

### GetTaxFinancialCostPercentage

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTaxFinancialCostPercentage() float32`

GetTaxFinancialCostPercentage returns the TaxFinancialCostPercentage field if non-nil, zero value otherwise.

### GetTaxFinancialCostPercentageOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTaxFinancialCostPercentageOk() (*float32, bool)`

GetTaxFinancialCostPercentageOk returns a tuple with the TaxFinancialCostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFinancialCostPercentage

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetTaxFinancialCostPercentage(v float32)`

SetTaxFinancialCostPercentage sets TaxFinancialCostPercentage field to given value.

### HasTaxFinancialCostPercentage

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasTaxFinancialCostPercentage() bool`

HasTaxFinancialCostPercentage returns a boolean if a field has been set.

### GetFinancialCostAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetFinancialCostAmount() float32`

GetFinancialCostAmount returns the FinancialCostAmount field if non-nil, zero value otherwise.

### GetFinancialCostAmountOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetFinancialCostAmountOk() (*float32, bool)`

GetFinancialCostAmountOk returns a tuple with the FinancialCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialCostAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetFinancialCostAmount(v float32)`

SetFinancialCostAmount sets FinancialCostAmount field to given value.

### HasFinancialCostAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasFinancialCostAmount() bool`

HasFinancialCostAmount returns a boolean if a field has been set.

### GetFinancialCostPercentage

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetFinancialCostPercentage() float32`

GetFinancialCostPercentage returns the FinancialCostPercentage field if non-nil, zero value otherwise.

### GetFinancialCostPercentageOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetFinancialCostPercentageOk() (*float32, bool)`

GetFinancialCostPercentageOk returns a tuple with the FinancialCostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialCostPercentage

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetFinancialCostPercentage(v float32)`

SetFinancialCostPercentage sets FinancialCostPercentage field to given value.

### HasFinancialCostPercentage

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasFinancialCostPercentage() bool`

HasFinancialCostPercentage returns a boolean if a field has been set.

### GetRequestAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetRequestAmount() float32`

GetRequestAmount returns the RequestAmount field if non-nil, zero value otherwise.

### GetRequestAmountOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetRequestAmountOk() (*float32, bool)`

GetRequestAmountOk returns a tuple with the RequestAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetRequestAmount(v float32)`

SetRequestAmount sets RequestAmount field to given value.

### HasRequestAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasRequestAmount() bool`

HasRequestAmount returns a boolean if a field has been set.

### GetAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAlternativeAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAlternativeAmount() float32`

GetAlternativeAmount returns the AlternativeAmount field if non-nil, zero value otherwise.

### GetAlternativeAmountOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAlternativeAmountOk() (*float32, bool)`

GetAlternativeAmountOk returns a tuple with the AlternativeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetAlternativeAmount(v float32)`

SetAlternativeAmount sets AlternativeAmount field to given value.

### HasAlternativeAmount

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasAlternativeAmount() bool`

HasAlternativeAmount returns a boolean if a field has been set.

### GetHostResultCode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostResultCode() int32`

GetHostResultCode returns the HostResultCode field if non-nil, zero value otherwise.

### GetHostResultCodeOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostResultCodeOk() (*int32, bool)`

GetHostResultCodeOk returns a tuple with the HostResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostResultCode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetHostResultCode(v int32)`

SetHostResultCode sets HostResultCode field to given value.

### HasHostResultCode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasHostResultCode() bool`

HasHostResultCode returns a boolean if a field has been set.

### GetHostMessage

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostMessage() string`

GetHostMessage returns the HostMessage field if non-nil, zero value otherwise.

### GetHostMessageOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostMessageOk() (*string, bool)`

GetHostMessageOk returns a tuple with the HostMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMessage

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetHostMessage(v string)`

SetHostMessage sets HostMessage field to given value.

### HasHostMessage

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasHostMessage() bool`

HasHostMessage returns a boolean if a field has been set.

### GetHostCode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostCode() string`

GetHostCode returns the HostCode field if non-nil, zero value otherwise.

### GetHostCodeOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostCodeOk() (*string, bool)`

GetHostCodeOk returns a tuple with the HostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetHostCode(v string)`

SetHostCode sets HostCode field to given value.

### HasHostCode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasHostCode() bool`

HasHostCode returns a boolean if a field has been set.

### GetHostDateTime

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostDateTime() time.Time`

GetHostDateTime returns the HostDateTime field if non-nil, zero value otherwise.

### GetHostDateTimeOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostDateTimeOk() (*time.Time, bool)`

GetHostDateTimeOk returns a tuple with the HostDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDateTime

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetHostDateTime(v time.Time)`

SetHostDateTime sets HostDateTime field to given value.

### HasHostDateTime

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasHostDateTime() bool`

HasHostDateTime returns a boolean if a field has been set.

### GetTransmitionDateTime

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTransmitionDateTime() time.Time`

GetTransmitionDateTime returns the TransmitionDateTime field if non-nil, zero value otherwise.

### GetTransmitionDateTimeOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTransmitionDateTimeOk() (*time.Time, bool)`

GetTransmitionDateTimeOk returns a tuple with the TransmitionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitionDateTime

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetTransmitionDateTime(v time.Time)`

SetTransmitionDateTime sets TransmitionDateTime field to given value.

### HasTransmitionDateTime

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasTransmitionDateTime() bool`

HasTransmitionDateTime returns a boolean if a field has been set.

### GetAuthTicket

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAuthTicket() int32`

GetAuthTicket returns the AuthTicket field if non-nil, zero value otherwise.

### GetAuthTicketOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAuthTicketOk() (*int32, bool)`

GetAuthTicketOk returns a tuple with the AuthTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTicket

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetAuthTicket(v int32)`

SetAuthTicket sets AuthTicket field to given value.

### HasAuthTicket

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasAuthTicket() bool`

HasAuthTicket returns a boolean if a field has been set.

### GetAuthRRN

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAuthRRN() string`

GetAuthRRN returns the AuthRRN field if non-nil, zero value otherwise.

### GetAuthRRNOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAuthRRNOk() (*string, bool)`

GetAuthRRNOk returns a tuple with the AuthRRN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRRN

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetAuthRRN(v string)`

SetAuthRRN sets AuthRRN field to given value.

### HasAuthRRN

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasAuthRRN() bool`

HasAuthRRN returns a boolean if a field has been set.

### GetIdentifierForTheAdquirer

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetIdentifierForTheAdquirer() string`

GetIdentifierForTheAdquirer returns the IdentifierForTheAdquirer field if non-nil, zero value otherwise.

### GetIdentifierForTheAdquirerOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetIdentifierForTheAdquirerOk() (*string, bool)`

GetIdentifierForTheAdquirerOk returns a tuple with the IdentifierForTheAdquirer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierForTheAdquirer

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetIdentifierForTheAdquirer(v string)`

SetIdentifierForTheAdquirer sets IdentifierForTheAdquirer field to given value.

### HasIdentifierForTheAdquirer

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasIdentifierForTheAdquirer() bool`

HasIdentifierForTheAdquirer returns a boolean if a field has been set.

### GetPaymentFacilitatorID

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetPaymentFacilitatorID() string`

GetPaymentFacilitatorID returns the PaymentFacilitatorID field if non-nil, zero value otherwise.

### GetPaymentFacilitatorIDOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetPaymentFacilitatorIDOk() (*string, bool)`

GetPaymentFacilitatorIDOk returns a tuple with the PaymentFacilitatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFacilitatorID

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetPaymentFacilitatorID(v string)`

SetPaymentFacilitatorID sets PaymentFacilitatorID field to given value.

### HasPaymentFacilitatorID

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasPaymentFacilitatorID() bool`

HasPaymentFacilitatorID returns a boolean if a field has been set.

### GetMerchantID

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetTerminalID

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTerminalID() int32`

GetTerminalID returns the TerminalID field if non-nil, zero value otherwise.

### GetTerminalIDOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTerminalIDOk() (*int32, bool)`

GetTerminalIDOk returns a tuple with the TerminalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalID

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetTerminalID(v int32)`

SetTerminalID sets TerminalID field to given value.

### HasTerminalID

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasTerminalID() bool`

HasTerminalID returns a boolean if a field has been set.

### GetTerminalTrace

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTerminalTrace() int32`

GetTerminalTrace returns the TerminalTrace field if non-nil, zero value otherwise.

### GetTerminalTraceOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTerminalTraceOk() (*int32, bool)`

GetTerminalTraceOk returns a tuple with the TerminalTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalTrace

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetTerminalTrace(v int32)`

SetTerminalTrace sets TerminalTrace field to given value.

### HasTerminalTrace

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasTerminalTrace() bool`

HasTerminalTrace returns a boolean if a field has been set.

### GetSettlementBatchNumber

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetSettlementBatchNumber() int32`

GetSettlementBatchNumber returns the SettlementBatchNumber field if non-nil, zero value otherwise.

### GetSettlementBatchNumberOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetSettlementBatchNumberOk() (*int32, bool)`

GetSettlementBatchNumberOk returns a tuple with the SettlementBatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementBatchNumber

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetSettlementBatchNumber(v int32)`

SetSettlementBatchNumber sets SettlementBatchNumber field to given value.

### HasSettlementBatchNumber

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasSettlementBatchNumber() bool`

HasSettlementBatchNumber returns a boolean if a field has been set.

### GetHostID

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostID() int32`

GetHostID returns the HostID field if non-nil, zero value otherwise.

### GetHostIDOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostIDOk() (*int32, bool)`

GetHostIDOk returns a tuple with the HostID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostID

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetHostID(v int32)`

SetHostID sets HostID field to given value.

### HasHostID

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasHostID() bool`

HasHostID returns a boolean if a field has been set.

### GetIssuerName

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### GetCardDescription

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardDescription() string`

GetCardDescription returns the CardDescription field if non-nil, zero value otherwise.

### GetCardDescriptionOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardDescriptionOk() (*string, bool)`

GetCardDescriptionOk returns a tuple with the CardDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDescription

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardDescription(v string)`

SetCardDescription sets CardDescription field to given value.

### HasCardDescription

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardDescription() bool`

HasCardDescription returns a boolean if a field has been set.

### GetPaymentMethodDescription

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetPaymentMethodDescription() string`

GetPaymentMethodDescription returns the PaymentMethodDescription field if non-nil, zero value otherwise.

### GetPaymentMethodDescriptionOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetPaymentMethodDescriptionOk() (*string, bool)`

GetPaymentMethodDescriptionOk returns a tuple with the PaymentMethodDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodDescription

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetPaymentMethodDescription(v string)`

SetPaymentMethodDescription sets PaymentMethodDescription field to given value.

### HasPaymentMethodDescription

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasPaymentMethodDescription() bool`

HasPaymentMethodDescription returns a boolean if a field has been set.

### GetPlanDescription

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetPlanDescription() string`

GetPlanDescription returns the PlanDescription field if non-nil, zero value otherwise.

### GetPlanDescriptionOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetPlanDescriptionOk() (*string, bool)`

GetPlanDescriptionOk returns a tuple with the PlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDescription

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetPlanDescription(v string)`

SetPlanDescription sets PlanDescription field to given value.

### HasPlanDescription

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasPlanDescription() bool`

HasPlanDescription returns a boolean if a field has been set.

### GetCardReadMode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardReadMode() string`

GetCardReadMode returns the CardReadMode field if non-nil, zero value otherwise.

### GetCardReadModeOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardReadModeOk() (*string, bool)`

GetCardReadModeOk returns a tuple with the CardReadMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardReadMode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardReadMode(v string)`

SetCardReadMode sets CardReadMode field to given value.

### HasCardReadMode

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardReadMode() bool`

HasCardReadMode returns a boolean if a field has been set.

### GetCardNumber

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNumberMasked

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardNumberMasked() string`

GetCardNumberMasked returns the CardNumberMasked field if non-nil, zero value otherwise.

### GetCardNumberMaskedOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardNumberMaskedOk() (*string, bool)`

GetCardNumberMaskedOk returns a tuple with the CardNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberMasked

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardNumberMasked(v string)`

SetCardNumberMasked sets CardNumberMasked field to given value.

### HasCardNumberMasked

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardNumberMasked() bool`

HasCardNumberMasked returns a boolean if a field has been set.

### GetCardHashing

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardHashing() string`

GetCardHashing returns the CardHashing field if non-nil, zero value otherwise.

### GetCardHashingOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardHashingOk() (*string, bool)`

GetCardHashingOk returns a tuple with the CardHashing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHashing

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardHashing(v string)`

SetCardHashing sets CardHashing field to given value.

### HasCardHashing

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardHashing() bool`

HasCardHashing returns a boolean if a field has been set.

### GetCurrencyDescription

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCurrencyDescription() string`

GetCurrencyDescription returns the CurrencyDescription field if non-nil, zero value otherwise.

### GetCurrencyDescriptionOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCurrencyDescriptionOk() (*string, bool)`

GetCurrencyDescriptionOk returns a tuple with the CurrencyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDescription

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCurrencyDescription(v string)`

SetCurrencyDescription sets CurrencyDescription field to given value.

### HasCurrencyDescription

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCurrencyDescription() bool`

HasCurrencyDescription returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCardCryptogramResponse

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardCryptogramResponse() string`

GetCardCryptogramResponse returns the CardCryptogramResponse field if non-nil, zero value otherwise.

### GetCardCryptogramResponseOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardCryptogramResponseOk() (*string, bool)`

GetCardCryptogramResponseOk returns a tuple with the CardCryptogramResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCryptogramResponse

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardCryptogramResponse(v string)`

SetCardCryptogramResponse sets CardCryptogramResponse field to given value.

### HasCardCryptogramResponse

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardCryptogramResponse() bool`

HasCardCryptogramResponse returns a boolean if a field has been set.

### GetCardAppName

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardAppName() string`

GetCardAppName returns the CardAppName field if non-nil, zero value otherwise.

### GetCardAppNameOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardAppNameOk() (*string, bool)`

GetCardAppNameOk returns a tuple with the CardAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppName

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardAppName(v string)`

SetCardAppName sets CardAppName field to given value.

### HasCardAppName

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardAppName() bool`

HasCardAppName returns a boolean if a field has been set.

### GetCardAppIdentifier

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardAppIdentifier() string`

GetCardAppIdentifier returns the CardAppIdentifier field if non-nil, zero value otherwise.

### GetCardAppIdentifierOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardAppIdentifierOk() (*string, bool)`

GetCardAppIdentifierOk returns a tuple with the CardAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppIdentifier

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardAppIdentifier(v string)`

SetCardAppIdentifier sets CardAppIdentifier field to given value.

### HasCardAppIdentifier

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardAppIdentifier() bool`

HasCardAppIdentifier returns a boolean if a field has been set.

### GetCardAppLabel

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardAppLabel() string`

GetCardAppLabel returns the CardAppLabel field if non-nil, zero value otherwise.

### GetCardAppLabelOk

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardAppLabelOk() (*string, bool)`

GetCardAppLabelOk returns a tuple with the CardAppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppLabel

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardAppLabel(v string)`

SetCardAppLabel sets CardAppLabel field to given value.

### HasCardAppLabel

`func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardAppLabel() bool`

HasCardAppLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


