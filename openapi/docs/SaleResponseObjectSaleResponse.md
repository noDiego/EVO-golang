# SaleResponseObjectSaleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **int32** | Código de Respuesta Interno de la plataforma, el POS debe actuar por lo que indican las acciones especificadas en ResponseActions y no por el código de respuesta informado en este campo o elemento, pero es una buena práctica que sea persistido por el mismo. | 
**ResponseActions** | **[]string** | Acciones a realizar por parte del POS y/o PINPAD en base al resultado de la operación que ha sido procesada. Cada uno de estos actions o acciones están concatenadas por comas. Los posibles actions son OK, Approve, Refuse, IssuerCall, Tickets, WithHold, GetCard, UseTerminalToAuthorize, ConfigurationError, SystemError, ResourceError, ProcessError, Completed, Configure, Display, EnableService y Print. | 
**ResponseMessage** | **string** | Descripción del resultado del proceso del requerimiento recibido. Esta descripción es generada por la plataforma, no por el Host que termine resolviendo la transacción. | 
**CompanyIdentification** | Pointer to **string** | ID que identifica la companía desde donde proviene la petición. | [optional] 
**SystemIdentification** | Pointer to **string** | ID que identifica el sistema desde donde proviene la petición. | [optional] 
**BranchIdentification** | Pointer to **string** | ID que identifica la sucursal desde donde proviene la petición. Esta sucursal pertenece a una determinada companía. | [optional] 
**POSIdentification** | Pointer to **string** | ID que identifica la caja o punto de venta desde donde proviene la petición. Este punto de venta pertenece a una determinada sucursal y companía. | [optional] 
**ForeignResponseCode** | Pointer to **string** | Código de respuesta para el sistema externo, es decir, para la aplicación cliente que se comunica con el TEF. | [optional] 
**RequestType** | Pointer to **string** | Tipo de Operación que se requirió, contendrá el mismo valor que se recibió en el requerimiento, sobre formatos que no soportan elementos complejos o compuestos. | [optional] 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 
**Notification** | Pointer to [**SaleResponseObjectSaleResponseNotification**](SaleResponseObjectSaleResponseNotification.md) |  | [optional] 
**Block** | Pointer to **string** | ID que identifica a un grupo de transacciones que serán confirmadas o canceladas | [optional] 
**RequestKey** | Pointer to **string** | ID generado para la identificación por parte del Plataforma de la información generada en la ejecución de un GetCard o un Payment Method. Será necesario para que un mensaje de Sale, Void, Payment Method o Enrollment identifique el contexto generado y lo utilice para esa operación. | [optional] 
**WorkingKeys** | Pointer to [**[]SaleResponseObjectSaleResponseWorkingKeys**](SaleResponseObjectSaleResponseWorkingKeys.md) | Nueva forma de enviar las llaves disponibles para esta caja. | [optional] 
**RequiredInformation** | Pointer to [**[]SaleResponseObjectSaleResponseRequiredInformation**](SaleResponseObjectSaleResponseRequiredInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**AnswerType** | Pointer to **string** | Tipo de Operación que se está requiriendo, solo necesario sobre formatos que no soportan elementos complejos o compuestos. | [optional] 
**AnswerKey** | Pointer to **string** | Código de identificación, generado por Plataforma, de la operación realizada&#39;. | [optional] 
**PublicAnswerKey** | Pointer to **string** | Código público de identificación, generado por Plataforma, de la operación realizada&#39;. | [optional] 
**WasReversePrevious** | Pointer to **int32** | Flag indicador de generación de reversa para la última operación reversable. | [optional] 
**ReversedAnswerKey** | Pointer to **string** | ID que identifica a la operación que acaba de ser reversada. | [optional] 
**ReversedSequence** | Pointer to **string** | Secuencia de la transacción que fue reversada. | [optional] 
**CommittedBlock** | Pointer to **string** | ID del bloque de transacciones que ha sido confirmado de forma automática (es decir, sin recibir un requerimiento de BlockClose). Este escenario se presentará si el Plataforma así se ha configurado para actuar bajo esa circunstancia. | [optional] 
**ReversedBlock** | Pointer to **string** | ID del bloque de transacciones que ha sido cancelado de forma automática (es decir, sin recibir un requerimiento de BlockClose). Este escenario se presentara si el Plataforma así se ha configurado para actuar bajo esa circunstancia. | [optional] 
**MessageID** | Pointer to **string** | Identificador Unívoco del Mensaje ( UUID v5 ). | [optional] 
**ServerAddress** | Pointer to **string** | Dirección IP del Server que atiende el requerimiento. | [optional] 
**ServerInstance** | Pointer to **string** | Instancia de Server que atiende el requerimiento. | [optional] 
**ServerNodeName** | Pointer to **string** | Nombre del Nodo que atendió el requerimiento. | [optional] 
**AdapterInputVersion** | Pointer to **string** | Versión del  Adaptador de Protocolo Entrante que atiende el Requerimiento. | [optional] 
**AdapterInputAddress** | Pointer to **string** | Dirección IP del Adaptador de Protocolo Entrante que atiende el requerimiento. | [optional] 
**AdapterInputNodeName** | Pointer to **string** | Nombre del Nodo del Adaptador de Protocolo Entrante que atiende el requerimiento. | [optional] 
**AdapterOutputVersion** | Pointer to **string** | Versión del  Adaptador de Protocolo Saliente que atiende el Requerimiento. | [optional] 
**AdapterOutputAddress** | Pointer to **string** | Dirección IP  del  Adaptador de Protocolo Saliente que atiende el Requerimiento. | [optional] 
**AdapterOutputNodeName** | Pointer to **string** | Nombre del Nodo  del  Adaptador de Protocolo Saliente que atiende el Requerimiento. | [optional] 
**AdditionalInformation** | Pointer to [**[]SaleResponseObjectSaleResponseAdditionalInformation**](SaleResponseObjectSaleResponseAdditionalInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**PrinterResponseMessage** | Pointer to **[]string** | Información adicional/Mensaje promocional/Leyenda de respuesta a imprimir en el ticket de la operación. Cada línea de este mensaje sera un elemento dentro del array. | [optional] 
**DisplayResponseMessage** | Pointer to **[]string** | Información adicional/Mensaje promocional/Leyenda de respuesta a mostrar en pantalla en el ticket de la operación. Cada línea de este mensaje será un elemento dentro del array. | [optional] 
**CurrencyCode** | Pointer to **string** | código de Moneda - ISO 4217 &lt;https://en.wikipedia.org/wiki/ISO_4217 Se puede utilizar la Codificación Alfabética o Numérica &lt;br /&gt;   * Num   - Alpha - Description &lt;br /&gt;   * &#39;032&#39; - &#39;ARS&#39; - Pesos Argentinos &lt;br /&gt;   * &#39;152&#39; - &#39;CLP&#39; - Pesos Chilenos &lt;br/&gt;   * &#39;484&#39; - &#39;MXN&#39; - Pesos Mexicanos &lt;br/&gt;   * &#39;840&#39; - &#39;USD&#39; - dólares Americanos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;858&#39; - &#39;UYU&#39; - Pesos Uruguayos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;986&#39; - &#39;BRL&#39; - Real Brasileño | [optional] 
**CurrencyDescription** | Pointer to **string** | Descripción del tipo de cambio utilizado en la transacción. | [optional] 
**TransactionDescription** | Pointer to **string** | Tipo de transacción | [optional] 
**TransactionResolutionMode** | Pointer to **string** | Modo en que fue realizada la transacción | [optional] 
**CurrencySymbol** | Pointer to **string** | Símbolo monetario del tipo de cambio utilizado en la transacción. | [optional] 
**FacilityPayments** | Pointer to **int32** | Cantidad de cuotas utilizadas para realizar la operación | [optional] 
**FacilityType** | Pointer to **int32** | Tipo de plan utilizado para para realizar la operación | [optional] 
**ValueFacilityPayments** | Pointer to **float32** | Monto final a pagar en cada una de las cuotas en las que se divida la compra | [optional] 
**Amount** | Pointer to **float32** | Importe o Monto de la Transacción. | [optional] 
**AlternativeAmount** | Pointer to **float32** | Monto con la que se realizó transacción. Si este valor es recibido, la búsqueda de los planes será limitada con este criterio. | [optional] 
**AmountCharged** | Pointer to **float32** | Importe o Monto de la Transacción que efectivamente se cobró.  Si se envía en Void o Return en lugar de Amount, se genera un Ajuste si el Host lo soporta. | [optional] 
**AmountToApply** | Pointer to **float32** | Importe o Monto de la Transacción a aplicar. | [optional] 
**CashbackAmount** | Pointer to **float32** | Monto del dinero en efectivo (cashback). | [optional] 
**TipAmount** | Pointer to **float32** | Importe o Monto de la Propina. | [optional] 
**RequestAmount** | Pointer to **float32** | Monto libre de costos financerios e impuestos por el que la venta se realizó. El monto cobrado realmente no es este, dado que no incluye las tasas e impuestos | [optional] 
**PromotedAmount** | Pointer to **float32** | Monto Sujeto a Promoción monto de la operación | [optional] 
**CredentialToken** | Pointer to **string** | Token asociado a la Credencial Enrolada | [optional] 
**CredentialIssuerToken** | Pointer to **string** | Emisor del Token asociado a la credencial enrolada | [optional] 
**InputTokens** | Pointer to [**[]SaleObjectSaleInputTokens**](SaleObjectSaleInputTokens.md) | Tokens. | [optional] 
**OutputTokens** | Pointer to [**[]SaleObjectSaleInputTokens**](SaleObjectSaleInputTokens.md) | Tokens. | [optional] 
**TaxFinancialCostAmount** | Pointer to **float32** | Monto del recargo impositivo aplicado al costo financiero que la transacción tiene | [optional] 
**TaxFinancialCostPercentage** | Pointer to **float32** | Porcentaje de recargo impositivo a aplicar sobre el monto del costo financiero | [optional] 
**FinancialCostAmount** | Pointer to **float32** | Monto del Costo financiero total generado en base al plan elegido | [optional] 
**FinancialCostPercentage** | Pointer to **float32** | Porcentaje del costo financiero a aplicar al monto de la transacción | [optional] 
**HostResultCode** | Pointer to **int32** | Código de Resultado retornado por el Host Adquirente. | [optional] 
**HostMessage** | Pointer to **string** | Mensaje Retornado por el Host Adquirente, normalmente asociado al valor de HostResultCode | [optional] 
**HostCode** | Pointer to **string** | Código de autorización retornado por el Host que resuelve la transacción. | [optional] 
**HostID** | Pointer to **int32** | Número de identificación del host al cual fue enviada la petición, y por el cual fue finalmente procesada. | [optional] 
**UserID** | Pointer to **string** | Identificador del usuario que está realizando la Transacción. | [optional] 
**IssuerName** | Pointer to **string** | Nombre del Emisor de la Credencial o Tarjeta que se usó en la transacción. | [optional] 
**HostDateTime** | Pointer to **time.Time** | Fecha y Hora de la transacción retornada por el Host que resuelve la Transacción - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**TransmitionDateTime** | Pointer to **time.Time** | Fecha y hora de transmisión de la operación hacia el host - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**AuthResultCode** | Pointer to **string** | Código de Resultado retornado por el Host Adquirente. | [optional] 
**AuthHostProcessCode** | Pointer to **string** | Código de procesamiento de la operación enviada al host. Elemento 3 del protocolo de comunicaciones ISO 8583, formato de mensajes utilizado por los hosts para realizar operaciones financieras. | [optional] 
**AuthHostMsgType** | Pointer to **string** | Elemento 0 del protocolo de comunicaciones ISO 8583, formato de mensajes utilizado por los hosts para realizar operaciones financieras. El valor de este campo es el que devuelve el host en una respuesta a una petición. | [optional] 
**AuthHostMessage** | Pointer to **string** | Mensaje Retornado por el Host Adquirente, normalmente asociado al valor de AuthResultCode | [optional] 
**HostMsgType** | Pointer to **string** | Elemento 0 del protocolo de comunicaciones ISO 8583, formato de mensajes utilizado por los hosts para realizar operaciones financieras. El valor de este campo es el que se envio al host en el envio de una petición. | [optional] 
**AuthCode** | Pointer to **string** | Código de autorización retornado por el Host que resuelve la transacción. | [optional] 
**AuthDateTime** | Pointer to **time.Time** | Fecha y Hora de la transacción retornada por el Host que resuelve la Transacción - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**AuthTicket** | Pointer to **int32** | Número Ticket  o Voucher Generado para la Plataforma. | [optional] 
**AuthRRN** | Pointer to **string** | Número de identificación de la transacción, utilizado por la mayoría de los hosts para realizar anulaciones y devoluciones. | [optional] 
**TransactionAuthenticationType** | Pointer to **string** | Tipo de autenticación | [optional] 
**IdentifierForTheAdquirer** | Pointer to **string** | Identificador que genera el Host Adquirente para la Transacción. En algunos podrá ser igual al AuthRRN. | [optional] 
**IdentifierForTheResolutor** | Pointer to **string** | Identificador que genera el Host o Plataforma que resuelve la transacción. | [optional] 
**PaymentFacilitatorID** | Pointer to **string** | Identificador de facilitador de pagos o Payfac. | [optional] 
**MerchantID** | Pointer to **string** | Número de comercio utilizado para realizar la transacción. Este Número es asignado por el host, y parametrizado en la BD, relacionado a cada uno de los planes disponibles. | [optional] 
**TerminalID** | Pointer to **string** | Identificador de Terminal por el cual se envía la Transacción al Host. | [optional] 
**TerminalTrace** | Pointer to **int32** | Número de Trace/Secuencia que genera la plataforma para la transacción asociado al TerminalID. | [optional] 
**SettlementBatchNumber** | Pointer to **int32** | Para aquellos host que exista el concepto de lote, es el número de lote al cual pertenece la transacción. | [optional] 
**CardReadMode** | Pointer to **string** | Modo de ingreso de los datos de la tarjeta. Los posibles valores significan: C - EMV Chip / B - Banda magnética / L - Contactless Chip / S - Contactless Banda / M - Manual (Tarjeta Presente) / T - Digitada (Tarjeta no Presente) / E - ECOMMERCE (Ventas por Internet)  / F - FALLBACK (Banda por falla en Chip) / K - TOKEN / R - Recurring ( Pagos Recurrentes ) | [optional] 
**CardReadModeDescription** | Pointer to **string** | Descripción del modo de ingreso con el que fue leída la tarjeta | [optional] 
**CardDescription** | Pointer to **string** | Nombre de la Tarjeta que se usó en la transacción, usado para la impresión del voucher. | [optional] 
**CardTypeDescription** | Pointer to **string** | Descripción del modo de ingreso utilizado para capturar los datos de la tarjeta. | [optional] 
**CardCategory** | Pointer to [**SaleResponseObjectSaleResponseCardCategory**](SaleResponseObjectSaleResponseCardCategory.md) |  | [optional] 
**CardNumber** | Pointer to **string** | Número de Tarjeta. En el caso de las respuestas el mismo estará enmascarado. | [optional] 
**CardNumberMasked** | Pointer to **string** | Número de tarjeta enmascarado, según indica la parametrización en la base de datos. Se utilizará para imprimir en el cupón. | [optional] 
**CardHashing** | Pointer to **string** | Hash de la tarjeta generado por la plataforma. | [optional] 
**CardExp** | Pointer to **string** | Fecha de vencimiento de la tarjeta. Este dato sera necesario si el modo de ingreso fue manual/digitada. | [optional] 
**CardCryptogramResponse** | Pointer to **string** | Tags EMV en format TLV recibidos desde el Host. | [optional] 
**CardAppName** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAppIdentifier** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAppLabel** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAuthRequestCryptogram** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAuthResponseCryptogram** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**Payer** | Pointer to [**SaleObjectSalePayer**](SaleObjectSalePayer.md) |  | [optional] 
**Customer** | Pointer to [**SaleObjectSaleCustomer**](SaleObjectSaleCustomer.md) |  | [optional] 
**MerchantCategory** | Pointer to [**SaleResponseObjectSaleResponseMerchantCategory**](SaleResponseObjectSaleResponseMerchantCategory.md) |  | [optional] 
**Products** | Pointer to [**[]SaleResponseObjectSaleResponseProducts**](SaleResponseObjectSaleResponseProducts.md) | Detalle de Productos de la Operación. | [optional] 
**PaymentMethod** | Pointer to [**SaleResponseObjectSaleResponsePaymentMethod**](SaleResponseObjectSaleResponsePaymentMethod.md) |  | [optional] 
**Plans** | Pointer to [**SaleResponseObjectSaleResponsePlans**](SaleResponseObjectSaleResponsePlans.md) |  | [optional] 
**PlanDescription** | Pointer to **string** | Descripción del plan utilizado para para realizar la operación | [optional] 
**PlanConfigVersion** | Pointer to **int32** | Identificador de Versión utilizada por Plataforma en la evaluación del Plan | [optional] 
**Tickets** | Pointer to [**[]SaleResponseObjectSaleResponseTickets**](SaleResponseObjectSaleResponseTickets.md) | Elemento Compuesto que indica qué Tickets intervienen en la transacción y si deben ser digitalizados o impresos. | [optional] 
**Configuration** | Pointer to [**SaleResponseObjectSaleResponseConfiguration**](SaleResponseObjectSaleResponseConfiguration.md) |  | [optional] 
**PaymentMethodId** | Pointer to **int32** | ID de la marca con la cual la tarjeta fue identificada | [optional] 
**PaymentMethodDescription** | Pointer to **string** | Descripción o nombre de la marca con la cual la tarjeta fue identificada | [optional] 

## Methods

### NewSaleResponseObjectSaleResponse

`func NewSaleResponseObjectSaleResponse(responseCode int32, responseActions []string, responseMessage string, ) *SaleResponseObjectSaleResponse`

NewSaleResponseObjectSaleResponse instantiates a new SaleResponseObjectSaleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleResponseObjectSaleResponseWithDefaults

`func NewSaleResponseObjectSaleResponseWithDefaults() *SaleResponseObjectSaleResponse`

NewSaleResponseObjectSaleResponseWithDefaults instantiates a new SaleResponseObjectSaleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *SaleResponseObjectSaleResponse) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *SaleResponseObjectSaleResponse) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *SaleResponseObjectSaleResponse) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseActions

`func (o *SaleResponseObjectSaleResponse) GetResponseActions() []string`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *SaleResponseObjectSaleResponse) GetResponseActionsOk() (*[]string, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *SaleResponseObjectSaleResponse) SetResponseActions(v []string)`

SetResponseActions sets ResponseActions field to given value.


### GetResponseMessage

`func (o *SaleResponseObjectSaleResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *SaleResponseObjectSaleResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *SaleResponseObjectSaleResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetCompanyIdentification

`func (o *SaleResponseObjectSaleResponse) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *SaleResponseObjectSaleResponse) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *SaleResponseObjectSaleResponse) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *SaleResponseObjectSaleResponse) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetSystemIdentification

`func (o *SaleResponseObjectSaleResponse) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *SaleResponseObjectSaleResponse) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *SaleResponseObjectSaleResponse) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.

### HasSystemIdentification

`func (o *SaleResponseObjectSaleResponse) HasSystemIdentification() bool`

HasSystemIdentification returns a boolean if a field has been set.

### GetBranchIdentification

`func (o *SaleResponseObjectSaleResponse) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *SaleResponseObjectSaleResponse) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *SaleResponseObjectSaleResponse) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *SaleResponseObjectSaleResponse) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *SaleResponseObjectSaleResponse) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *SaleResponseObjectSaleResponse) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *SaleResponseObjectSaleResponse) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *SaleResponseObjectSaleResponse) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetForeignResponseCode

`func (o *SaleResponseObjectSaleResponse) GetForeignResponseCode() string`

GetForeignResponseCode returns the ForeignResponseCode field if non-nil, zero value otherwise.

### GetForeignResponseCodeOk

`func (o *SaleResponseObjectSaleResponse) GetForeignResponseCodeOk() (*string, bool)`

GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignResponseCode

`func (o *SaleResponseObjectSaleResponse) SetForeignResponseCode(v string)`

SetForeignResponseCode sets ForeignResponseCode field to given value.

### HasForeignResponseCode

`func (o *SaleResponseObjectSaleResponse) HasForeignResponseCode() bool`

HasForeignResponseCode returns a boolean if a field has been set.

### GetRequestType

`func (o *SaleResponseObjectSaleResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *SaleResponseObjectSaleResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *SaleResponseObjectSaleResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *SaleResponseObjectSaleResponse) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetServiceVersion

`func (o *SaleResponseObjectSaleResponse) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *SaleResponseObjectSaleResponse) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *SaleResponseObjectSaleResponse) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *SaleResponseObjectSaleResponse) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *SaleResponseObjectSaleResponse) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *SaleResponseObjectSaleResponse) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *SaleResponseObjectSaleResponse) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *SaleResponseObjectSaleResponse) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *SaleResponseObjectSaleResponse) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *SaleResponseObjectSaleResponse) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *SaleResponseObjectSaleResponse) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *SaleResponseObjectSaleResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetNotification

`func (o *SaleResponseObjectSaleResponse) GetNotification() SaleResponseObjectSaleResponseNotification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *SaleResponseObjectSaleResponse) GetNotificationOk() (*SaleResponseObjectSaleResponseNotification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *SaleResponseObjectSaleResponse) SetNotification(v SaleResponseObjectSaleResponseNotification)`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *SaleResponseObjectSaleResponse) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetBlock

`func (o *SaleResponseObjectSaleResponse) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *SaleResponseObjectSaleResponse) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *SaleResponseObjectSaleResponse) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *SaleResponseObjectSaleResponse) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetRequestKey

`func (o *SaleResponseObjectSaleResponse) GetRequestKey() string`

GetRequestKey returns the RequestKey field if non-nil, zero value otherwise.

### GetRequestKeyOk

`func (o *SaleResponseObjectSaleResponse) GetRequestKeyOk() (*string, bool)`

GetRequestKeyOk returns a tuple with the RequestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestKey

`func (o *SaleResponseObjectSaleResponse) SetRequestKey(v string)`

SetRequestKey sets RequestKey field to given value.

### HasRequestKey

`func (o *SaleResponseObjectSaleResponse) HasRequestKey() bool`

HasRequestKey returns a boolean if a field has been set.

### GetWorkingKeys

`func (o *SaleResponseObjectSaleResponse) GetWorkingKeys() []SaleResponseObjectSaleResponseWorkingKeys`

GetWorkingKeys returns the WorkingKeys field if non-nil, zero value otherwise.

### GetWorkingKeysOk

`func (o *SaleResponseObjectSaleResponse) GetWorkingKeysOk() (*[]SaleResponseObjectSaleResponseWorkingKeys, bool)`

GetWorkingKeysOk returns a tuple with the WorkingKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingKeys

`func (o *SaleResponseObjectSaleResponse) SetWorkingKeys(v []SaleResponseObjectSaleResponseWorkingKeys)`

SetWorkingKeys sets WorkingKeys field to given value.

### HasWorkingKeys

`func (o *SaleResponseObjectSaleResponse) HasWorkingKeys() bool`

HasWorkingKeys returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *SaleResponseObjectSaleResponse) GetRequiredInformation() []SaleResponseObjectSaleResponseRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *SaleResponseObjectSaleResponse) GetRequiredInformationOk() (*[]SaleResponseObjectSaleResponseRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *SaleResponseObjectSaleResponse) SetRequiredInformation(v []SaleResponseObjectSaleResponseRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *SaleResponseObjectSaleResponse) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetAnswerType

`func (o *SaleResponseObjectSaleResponse) GetAnswerType() string`

GetAnswerType returns the AnswerType field if non-nil, zero value otherwise.

### GetAnswerTypeOk

`func (o *SaleResponseObjectSaleResponse) GetAnswerTypeOk() (*string, bool)`

GetAnswerTypeOk returns a tuple with the AnswerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerType

`func (o *SaleResponseObjectSaleResponse) SetAnswerType(v string)`

SetAnswerType sets AnswerType field to given value.

### HasAnswerType

`func (o *SaleResponseObjectSaleResponse) HasAnswerType() bool`

HasAnswerType returns a boolean if a field has been set.

### GetAnswerKey

`func (o *SaleResponseObjectSaleResponse) GetAnswerKey() string`

GetAnswerKey returns the AnswerKey field if non-nil, zero value otherwise.

### GetAnswerKeyOk

`func (o *SaleResponseObjectSaleResponse) GetAnswerKeyOk() (*string, bool)`

GetAnswerKeyOk returns a tuple with the AnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerKey

`func (o *SaleResponseObjectSaleResponse) SetAnswerKey(v string)`

SetAnswerKey sets AnswerKey field to given value.

### HasAnswerKey

`func (o *SaleResponseObjectSaleResponse) HasAnswerKey() bool`

HasAnswerKey returns a boolean if a field has been set.

### GetPublicAnswerKey

`func (o *SaleResponseObjectSaleResponse) GetPublicAnswerKey() string`

GetPublicAnswerKey returns the PublicAnswerKey field if non-nil, zero value otherwise.

### GetPublicAnswerKeyOk

`func (o *SaleResponseObjectSaleResponse) GetPublicAnswerKeyOk() (*string, bool)`

GetPublicAnswerKeyOk returns a tuple with the PublicAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAnswerKey

`func (o *SaleResponseObjectSaleResponse) SetPublicAnswerKey(v string)`

SetPublicAnswerKey sets PublicAnswerKey field to given value.

### HasPublicAnswerKey

`func (o *SaleResponseObjectSaleResponse) HasPublicAnswerKey() bool`

HasPublicAnswerKey returns a boolean if a field has been set.

### GetWasReversePrevious

`func (o *SaleResponseObjectSaleResponse) GetWasReversePrevious() int32`

GetWasReversePrevious returns the WasReversePrevious field if non-nil, zero value otherwise.

### GetWasReversePreviousOk

`func (o *SaleResponseObjectSaleResponse) GetWasReversePreviousOk() (*int32, bool)`

GetWasReversePreviousOk returns a tuple with the WasReversePrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasReversePrevious

`func (o *SaleResponseObjectSaleResponse) SetWasReversePrevious(v int32)`

SetWasReversePrevious sets WasReversePrevious field to given value.

### HasWasReversePrevious

`func (o *SaleResponseObjectSaleResponse) HasWasReversePrevious() bool`

HasWasReversePrevious returns a boolean if a field has been set.

### GetReversedAnswerKey

`func (o *SaleResponseObjectSaleResponse) GetReversedAnswerKey() string`

GetReversedAnswerKey returns the ReversedAnswerKey field if non-nil, zero value otherwise.

### GetReversedAnswerKeyOk

`func (o *SaleResponseObjectSaleResponse) GetReversedAnswerKeyOk() (*string, bool)`

GetReversedAnswerKeyOk returns a tuple with the ReversedAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedAnswerKey

`func (o *SaleResponseObjectSaleResponse) SetReversedAnswerKey(v string)`

SetReversedAnswerKey sets ReversedAnswerKey field to given value.

### HasReversedAnswerKey

`func (o *SaleResponseObjectSaleResponse) HasReversedAnswerKey() bool`

HasReversedAnswerKey returns a boolean if a field has been set.

### GetReversedSequence

`func (o *SaleResponseObjectSaleResponse) GetReversedSequence() string`

GetReversedSequence returns the ReversedSequence field if non-nil, zero value otherwise.

### GetReversedSequenceOk

`func (o *SaleResponseObjectSaleResponse) GetReversedSequenceOk() (*string, bool)`

GetReversedSequenceOk returns a tuple with the ReversedSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedSequence

`func (o *SaleResponseObjectSaleResponse) SetReversedSequence(v string)`

SetReversedSequence sets ReversedSequence field to given value.

### HasReversedSequence

`func (o *SaleResponseObjectSaleResponse) HasReversedSequence() bool`

HasReversedSequence returns a boolean if a field has been set.

### GetCommittedBlock

`func (o *SaleResponseObjectSaleResponse) GetCommittedBlock() string`

GetCommittedBlock returns the CommittedBlock field if non-nil, zero value otherwise.

### GetCommittedBlockOk

`func (o *SaleResponseObjectSaleResponse) GetCommittedBlockOk() (*string, bool)`

GetCommittedBlockOk returns a tuple with the CommittedBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedBlock

`func (o *SaleResponseObjectSaleResponse) SetCommittedBlock(v string)`

SetCommittedBlock sets CommittedBlock field to given value.

### HasCommittedBlock

`func (o *SaleResponseObjectSaleResponse) HasCommittedBlock() bool`

HasCommittedBlock returns a boolean if a field has been set.

### GetReversedBlock

`func (o *SaleResponseObjectSaleResponse) GetReversedBlock() string`

GetReversedBlock returns the ReversedBlock field if non-nil, zero value otherwise.

### GetReversedBlockOk

`func (o *SaleResponseObjectSaleResponse) GetReversedBlockOk() (*string, bool)`

GetReversedBlockOk returns a tuple with the ReversedBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedBlock

`func (o *SaleResponseObjectSaleResponse) SetReversedBlock(v string)`

SetReversedBlock sets ReversedBlock field to given value.

### HasReversedBlock

`func (o *SaleResponseObjectSaleResponse) HasReversedBlock() bool`

HasReversedBlock returns a boolean if a field has been set.

### GetMessageID

`func (o *SaleResponseObjectSaleResponse) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *SaleResponseObjectSaleResponse) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *SaleResponseObjectSaleResponse) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.

### HasMessageID

`func (o *SaleResponseObjectSaleResponse) HasMessageID() bool`

HasMessageID returns a boolean if a field has been set.

### GetServerAddress

`func (o *SaleResponseObjectSaleResponse) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *SaleResponseObjectSaleResponse) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *SaleResponseObjectSaleResponse) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *SaleResponseObjectSaleResponse) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.

### GetServerInstance

`func (o *SaleResponseObjectSaleResponse) GetServerInstance() string`

GetServerInstance returns the ServerInstance field if non-nil, zero value otherwise.

### GetServerInstanceOk

`func (o *SaleResponseObjectSaleResponse) GetServerInstanceOk() (*string, bool)`

GetServerInstanceOk returns a tuple with the ServerInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstance

`func (o *SaleResponseObjectSaleResponse) SetServerInstance(v string)`

SetServerInstance sets ServerInstance field to given value.

### HasServerInstance

`func (o *SaleResponseObjectSaleResponse) HasServerInstance() bool`

HasServerInstance returns a boolean if a field has been set.

### GetServerNodeName

`func (o *SaleResponseObjectSaleResponse) GetServerNodeName() string`

GetServerNodeName returns the ServerNodeName field if non-nil, zero value otherwise.

### GetServerNodeNameOk

`func (o *SaleResponseObjectSaleResponse) GetServerNodeNameOk() (*string, bool)`

GetServerNodeNameOk returns a tuple with the ServerNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNodeName

`func (o *SaleResponseObjectSaleResponse) SetServerNodeName(v string)`

SetServerNodeName sets ServerNodeName field to given value.

### HasServerNodeName

`func (o *SaleResponseObjectSaleResponse) HasServerNodeName() bool`

HasServerNodeName returns a boolean if a field has been set.

### GetAdapterInputVersion

`func (o *SaleResponseObjectSaleResponse) GetAdapterInputVersion() string`

GetAdapterInputVersion returns the AdapterInputVersion field if non-nil, zero value otherwise.

### GetAdapterInputVersionOk

`func (o *SaleResponseObjectSaleResponse) GetAdapterInputVersionOk() (*string, bool)`

GetAdapterInputVersionOk returns a tuple with the AdapterInputVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputVersion

`func (o *SaleResponseObjectSaleResponse) SetAdapterInputVersion(v string)`

SetAdapterInputVersion sets AdapterInputVersion field to given value.

### HasAdapterInputVersion

`func (o *SaleResponseObjectSaleResponse) HasAdapterInputVersion() bool`

HasAdapterInputVersion returns a boolean if a field has been set.

### GetAdapterInputAddress

`func (o *SaleResponseObjectSaleResponse) GetAdapterInputAddress() string`

GetAdapterInputAddress returns the AdapterInputAddress field if non-nil, zero value otherwise.

### GetAdapterInputAddressOk

`func (o *SaleResponseObjectSaleResponse) GetAdapterInputAddressOk() (*string, bool)`

GetAdapterInputAddressOk returns a tuple with the AdapterInputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputAddress

`func (o *SaleResponseObjectSaleResponse) SetAdapterInputAddress(v string)`

SetAdapterInputAddress sets AdapterInputAddress field to given value.

### HasAdapterInputAddress

`func (o *SaleResponseObjectSaleResponse) HasAdapterInputAddress() bool`

HasAdapterInputAddress returns a boolean if a field has been set.

### GetAdapterInputNodeName

`func (o *SaleResponseObjectSaleResponse) GetAdapterInputNodeName() string`

GetAdapterInputNodeName returns the AdapterInputNodeName field if non-nil, zero value otherwise.

### GetAdapterInputNodeNameOk

`func (o *SaleResponseObjectSaleResponse) GetAdapterInputNodeNameOk() (*string, bool)`

GetAdapterInputNodeNameOk returns a tuple with the AdapterInputNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputNodeName

`func (o *SaleResponseObjectSaleResponse) SetAdapterInputNodeName(v string)`

SetAdapterInputNodeName sets AdapterInputNodeName field to given value.

### HasAdapterInputNodeName

`func (o *SaleResponseObjectSaleResponse) HasAdapterInputNodeName() bool`

HasAdapterInputNodeName returns a boolean if a field has been set.

### GetAdapterOutputVersion

`func (o *SaleResponseObjectSaleResponse) GetAdapterOutputVersion() string`

GetAdapterOutputVersion returns the AdapterOutputVersion field if non-nil, zero value otherwise.

### GetAdapterOutputVersionOk

`func (o *SaleResponseObjectSaleResponse) GetAdapterOutputVersionOk() (*string, bool)`

GetAdapterOutputVersionOk returns a tuple with the AdapterOutputVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputVersion

`func (o *SaleResponseObjectSaleResponse) SetAdapterOutputVersion(v string)`

SetAdapterOutputVersion sets AdapterOutputVersion field to given value.

### HasAdapterOutputVersion

`func (o *SaleResponseObjectSaleResponse) HasAdapterOutputVersion() bool`

HasAdapterOutputVersion returns a boolean if a field has been set.

### GetAdapterOutputAddress

`func (o *SaleResponseObjectSaleResponse) GetAdapterOutputAddress() string`

GetAdapterOutputAddress returns the AdapterOutputAddress field if non-nil, zero value otherwise.

### GetAdapterOutputAddressOk

`func (o *SaleResponseObjectSaleResponse) GetAdapterOutputAddressOk() (*string, bool)`

GetAdapterOutputAddressOk returns a tuple with the AdapterOutputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputAddress

`func (o *SaleResponseObjectSaleResponse) SetAdapterOutputAddress(v string)`

SetAdapterOutputAddress sets AdapterOutputAddress field to given value.

### HasAdapterOutputAddress

`func (o *SaleResponseObjectSaleResponse) HasAdapterOutputAddress() bool`

HasAdapterOutputAddress returns a boolean if a field has been set.

### GetAdapterOutputNodeName

`func (o *SaleResponseObjectSaleResponse) GetAdapterOutputNodeName() string`

GetAdapterOutputNodeName returns the AdapterOutputNodeName field if non-nil, zero value otherwise.

### GetAdapterOutputNodeNameOk

`func (o *SaleResponseObjectSaleResponse) GetAdapterOutputNodeNameOk() (*string, bool)`

GetAdapterOutputNodeNameOk returns a tuple with the AdapterOutputNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputNodeName

`func (o *SaleResponseObjectSaleResponse) SetAdapterOutputNodeName(v string)`

SetAdapterOutputNodeName sets AdapterOutputNodeName field to given value.

### HasAdapterOutputNodeName

`func (o *SaleResponseObjectSaleResponse) HasAdapterOutputNodeName() bool`

HasAdapterOutputNodeName returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *SaleResponseObjectSaleResponse) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *SaleResponseObjectSaleResponse) GetAdditionalInformationOk() (*[]SaleResponseObjectSaleResponseAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *SaleResponseObjectSaleResponse) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *SaleResponseObjectSaleResponse) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetPrinterResponseMessage

`func (o *SaleResponseObjectSaleResponse) GetPrinterResponseMessage() []string`

GetPrinterResponseMessage returns the PrinterResponseMessage field if non-nil, zero value otherwise.

### GetPrinterResponseMessageOk

`func (o *SaleResponseObjectSaleResponse) GetPrinterResponseMessageOk() (*[]string, bool)`

GetPrinterResponseMessageOk returns a tuple with the PrinterResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinterResponseMessage

`func (o *SaleResponseObjectSaleResponse) SetPrinterResponseMessage(v []string)`

SetPrinterResponseMessage sets PrinterResponseMessage field to given value.

### HasPrinterResponseMessage

`func (o *SaleResponseObjectSaleResponse) HasPrinterResponseMessage() bool`

HasPrinterResponseMessage returns a boolean if a field has been set.

### GetDisplayResponseMessage

`func (o *SaleResponseObjectSaleResponse) GetDisplayResponseMessage() []string`

GetDisplayResponseMessage returns the DisplayResponseMessage field if non-nil, zero value otherwise.

### GetDisplayResponseMessageOk

`func (o *SaleResponseObjectSaleResponse) GetDisplayResponseMessageOk() (*[]string, bool)`

GetDisplayResponseMessageOk returns a tuple with the DisplayResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayResponseMessage

`func (o *SaleResponseObjectSaleResponse) SetDisplayResponseMessage(v []string)`

SetDisplayResponseMessage sets DisplayResponseMessage field to given value.

### HasDisplayResponseMessage

`func (o *SaleResponseObjectSaleResponse) HasDisplayResponseMessage() bool`

HasDisplayResponseMessage returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *SaleResponseObjectSaleResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *SaleResponseObjectSaleResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *SaleResponseObjectSaleResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *SaleResponseObjectSaleResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencyDescription

`func (o *SaleResponseObjectSaleResponse) GetCurrencyDescription() string`

GetCurrencyDescription returns the CurrencyDescription field if non-nil, zero value otherwise.

### GetCurrencyDescriptionOk

`func (o *SaleResponseObjectSaleResponse) GetCurrencyDescriptionOk() (*string, bool)`

GetCurrencyDescriptionOk returns a tuple with the CurrencyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDescription

`func (o *SaleResponseObjectSaleResponse) SetCurrencyDescription(v string)`

SetCurrencyDescription sets CurrencyDescription field to given value.

### HasCurrencyDescription

`func (o *SaleResponseObjectSaleResponse) HasCurrencyDescription() bool`

HasCurrencyDescription returns a boolean if a field has been set.

### GetTransactionDescription

`func (o *SaleResponseObjectSaleResponse) GetTransactionDescription() string`

GetTransactionDescription returns the TransactionDescription field if non-nil, zero value otherwise.

### GetTransactionDescriptionOk

`func (o *SaleResponseObjectSaleResponse) GetTransactionDescriptionOk() (*string, bool)`

GetTransactionDescriptionOk returns a tuple with the TransactionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDescription

`func (o *SaleResponseObjectSaleResponse) SetTransactionDescription(v string)`

SetTransactionDescription sets TransactionDescription field to given value.

### HasTransactionDescription

`func (o *SaleResponseObjectSaleResponse) HasTransactionDescription() bool`

HasTransactionDescription returns a boolean if a field has been set.

### GetTransactionResolutionMode

`func (o *SaleResponseObjectSaleResponse) GetTransactionResolutionMode() string`

GetTransactionResolutionMode returns the TransactionResolutionMode field if non-nil, zero value otherwise.

### GetTransactionResolutionModeOk

`func (o *SaleResponseObjectSaleResponse) GetTransactionResolutionModeOk() (*string, bool)`

GetTransactionResolutionModeOk returns a tuple with the TransactionResolutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionResolutionMode

`func (o *SaleResponseObjectSaleResponse) SetTransactionResolutionMode(v string)`

SetTransactionResolutionMode sets TransactionResolutionMode field to given value.

### HasTransactionResolutionMode

`func (o *SaleResponseObjectSaleResponse) HasTransactionResolutionMode() bool`

HasTransactionResolutionMode returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *SaleResponseObjectSaleResponse) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *SaleResponseObjectSaleResponse) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *SaleResponseObjectSaleResponse) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *SaleResponseObjectSaleResponse) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetFacilityPayments

`func (o *SaleResponseObjectSaleResponse) GetFacilityPayments() int32`

GetFacilityPayments returns the FacilityPayments field if non-nil, zero value otherwise.

### GetFacilityPaymentsOk

`func (o *SaleResponseObjectSaleResponse) GetFacilityPaymentsOk() (*int32, bool)`

GetFacilityPaymentsOk returns a tuple with the FacilityPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityPayments

`func (o *SaleResponseObjectSaleResponse) SetFacilityPayments(v int32)`

SetFacilityPayments sets FacilityPayments field to given value.

### HasFacilityPayments

`func (o *SaleResponseObjectSaleResponse) HasFacilityPayments() bool`

HasFacilityPayments returns a boolean if a field has been set.

### GetFacilityType

`func (o *SaleResponseObjectSaleResponse) GetFacilityType() int32`

GetFacilityType returns the FacilityType field if non-nil, zero value otherwise.

### GetFacilityTypeOk

`func (o *SaleResponseObjectSaleResponse) GetFacilityTypeOk() (*int32, bool)`

GetFacilityTypeOk returns a tuple with the FacilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityType

`func (o *SaleResponseObjectSaleResponse) SetFacilityType(v int32)`

SetFacilityType sets FacilityType field to given value.

### HasFacilityType

`func (o *SaleResponseObjectSaleResponse) HasFacilityType() bool`

HasFacilityType returns a boolean if a field has been set.

### GetValueFacilityPayments

`func (o *SaleResponseObjectSaleResponse) GetValueFacilityPayments() float32`

GetValueFacilityPayments returns the ValueFacilityPayments field if non-nil, zero value otherwise.

### GetValueFacilityPaymentsOk

`func (o *SaleResponseObjectSaleResponse) GetValueFacilityPaymentsOk() (*float32, bool)`

GetValueFacilityPaymentsOk returns a tuple with the ValueFacilityPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFacilityPayments

`func (o *SaleResponseObjectSaleResponse) SetValueFacilityPayments(v float32)`

SetValueFacilityPayments sets ValueFacilityPayments field to given value.

### HasValueFacilityPayments

`func (o *SaleResponseObjectSaleResponse) HasValueFacilityPayments() bool`

HasValueFacilityPayments returns a boolean if a field has been set.

### GetAmount

`func (o *SaleResponseObjectSaleResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SaleResponseObjectSaleResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SaleResponseObjectSaleResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SaleResponseObjectSaleResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAlternativeAmount

`func (o *SaleResponseObjectSaleResponse) GetAlternativeAmount() float32`

GetAlternativeAmount returns the AlternativeAmount field if non-nil, zero value otherwise.

### GetAlternativeAmountOk

`func (o *SaleResponseObjectSaleResponse) GetAlternativeAmountOk() (*float32, bool)`

GetAlternativeAmountOk returns a tuple with the AlternativeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeAmount

`func (o *SaleResponseObjectSaleResponse) SetAlternativeAmount(v float32)`

SetAlternativeAmount sets AlternativeAmount field to given value.

### HasAlternativeAmount

`func (o *SaleResponseObjectSaleResponse) HasAlternativeAmount() bool`

HasAlternativeAmount returns a boolean if a field has been set.

### GetAmountCharged

`func (o *SaleResponseObjectSaleResponse) GetAmountCharged() float32`

GetAmountCharged returns the AmountCharged field if non-nil, zero value otherwise.

### GetAmountChargedOk

`func (o *SaleResponseObjectSaleResponse) GetAmountChargedOk() (*float32, bool)`

GetAmountChargedOk returns a tuple with the AmountCharged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCharged

`func (o *SaleResponseObjectSaleResponse) SetAmountCharged(v float32)`

SetAmountCharged sets AmountCharged field to given value.

### HasAmountCharged

`func (o *SaleResponseObjectSaleResponse) HasAmountCharged() bool`

HasAmountCharged returns a boolean if a field has been set.

### GetAmountToApply

`func (o *SaleResponseObjectSaleResponse) GetAmountToApply() float32`

GetAmountToApply returns the AmountToApply field if non-nil, zero value otherwise.

### GetAmountToApplyOk

`func (o *SaleResponseObjectSaleResponse) GetAmountToApplyOk() (*float32, bool)`

GetAmountToApplyOk returns a tuple with the AmountToApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountToApply

`func (o *SaleResponseObjectSaleResponse) SetAmountToApply(v float32)`

SetAmountToApply sets AmountToApply field to given value.

### HasAmountToApply

`func (o *SaleResponseObjectSaleResponse) HasAmountToApply() bool`

HasAmountToApply returns a boolean if a field has been set.

### GetCashbackAmount

`func (o *SaleResponseObjectSaleResponse) GetCashbackAmount() float32`

GetCashbackAmount returns the CashbackAmount field if non-nil, zero value otherwise.

### GetCashbackAmountOk

`func (o *SaleResponseObjectSaleResponse) GetCashbackAmountOk() (*float32, bool)`

GetCashbackAmountOk returns a tuple with the CashbackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackAmount

`func (o *SaleResponseObjectSaleResponse) SetCashbackAmount(v float32)`

SetCashbackAmount sets CashbackAmount field to given value.

### HasCashbackAmount

`func (o *SaleResponseObjectSaleResponse) HasCashbackAmount() bool`

HasCashbackAmount returns a boolean if a field has been set.

### GetTipAmount

`func (o *SaleResponseObjectSaleResponse) GetTipAmount() float32`

GetTipAmount returns the TipAmount field if non-nil, zero value otherwise.

### GetTipAmountOk

`func (o *SaleResponseObjectSaleResponse) GetTipAmountOk() (*float32, bool)`

GetTipAmountOk returns a tuple with the TipAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipAmount

`func (o *SaleResponseObjectSaleResponse) SetTipAmount(v float32)`

SetTipAmount sets TipAmount field to given value.

### HasTipAmount

`func (o *SaleResponseObjectSaleResponse) HasTipAmount() bool`

HasTipAmount returns a boolean if a field has been set.

### GetRequestAmount

`func (o *SaleResponseObjectSaleResponse) GetRequestAmount() float32`

GetRequestAmount returns the RequestAmount field if non-nil, zero value otherwise.

### GetRequestAmountOk

`func (o *SaleResponseObjectSaleResponse) GetRequestAmountOk() (*float32, bool)`

GetRequestAmountOk returns a tuple with the RequestAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAmount

`func (o *SaleResponseObjectSaleResponse) SetRequestAmount(v float32)`

SetRequestAmount sets RequestAmount field to given value.

### HasRequestAmount

`func (o *SaleResponseObjectSaleResponse) HasRequestAmount() bool`

HasRequestAmount returns a boolean if a field has been set.

### GetPromotedAmount

`func (o *SaleResponseObjectSaleResponse) GetPromotedAmount() float32`

GetPromotedAmount returns the PromotedAmount field if non-nil, zero value otherwise.

### GetPromotedAmountOk

`func (o *SaleResponseObjectSaleResponse) GetPromotedAmountOk() (*float32, bool)`

GetPromotedAmountOk returns a tuple with the PromotedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotedAmount

`func (o *SaleResponseObjectSaleResponse) SetPromotedAmount(v float32)`

SetPromotedAmount sets PromotedAmount field to given value.

### HasPromotedAmount

`func (o *SaleResponseObjectSaleResponse) HasPromotedAmount() bool`

HasPromotedAmount returns a boolean if a field has been set.

### GetCredentialToken

`func (o *SaleResponseObjectSaleResponse) GetCredentialToken() string`

GetCredentialToken returns the CredentialToken field if non-nil, zero value otherwise.

### GetCredentialTokenOk

`func (o *SaleResponseObjectSaleResponse) GetCredentialTokenOk() (*string, bool)`

GetCredentialTokenOk returns a tuple with the CredentialToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialToken

`func (o *SaleResponseObjectSaleResponse) SetCredentialToken(v string)`

SetCredentialToken sets CredentialToken field to given value.

### HasCredentialToken

`func (o *SaleResponseObjectSaleResponse) HasCredentialToken() bool`

HasCredentialToken returns a boolean if a field has been set.

### GetCredentialIssuerToken

`func (o *SaleResponseObjectSaleResponse) GetCredentialIssuerToken() string`

GetCredentialIssuerToken returns the CredentialIssuerToken field if non-nil, zero value otherwise.

### GetCredentialIssuerTokenOk

`func (o *SaleResponseObjectSaleResponse) GetCredentialIssuerTokenOk() (*string, bool)`

GetCredentialIssuerTokenOk returns a tuple with the CredentialIssuerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIssuerToken

`func (o *SaleResponseObjectSaleResponse) SetCredentialIssuerToken(v string)`

SetCredentialIssuerToken sets CredentialIssuerToken field to given value.

### HasCredentialIssuerToken

`func (o *SaleResponseObjectSaleResponse) HasCredentialIssuerToken() bool`

HasCredentialIssuerToken returns a boolean if a field has been set.

### GetInputTokens

`func (o *SaleResponseObjectSaleResponse) GetInputTokens() []SaleObjectSaleInputTokens`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *SaleResponseObjectSaleResponse) GetInputTokensOk() (*[]SaleObjectSaleInputTokens, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *SaleResponseObjectSaleResponse) SetInputTokens(v []SaleObjectSaleInputTokens)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *SaleResponseObjectSaleResponse) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetOutputTokens

`func (o *SaleResponseObjectSaleResponse) GetOutputTokens() []SaleObjectSaleInputTokens`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *SaleResponseObjectSaleResponse) GetOutputTokensOk() (*[]SaleObjectSaleInputTokens, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *SaleResponseObjectSaleResponse) SetOutputTokens(v []SaleObjectSaleInputTokens)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *SaleResponseObjectSaleResponse) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetTaxFinancialCostAmount

`func (o *SaleResponseObjectSaleResponse) GetTaxFinancialCostAmount() float32`

GetTaxFinancialCostAmount returns the TaxFinancialCostAmount field if non-nil, zero value otherwise.

### GetTaxFinancialCostAmountOk

`func (o *SaleResponseObjectSaleResponse) GetTaxFinancialCostAmountOk() (*float32, bool)`

GetTaxFinancialCostAmountOk returns a tuple with the TaxFinancialCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFinancialCostAmount

`func (o *SaleResponseObjectSaleResponse) SetTaxFinancialCostAmount(v float32)`

SetTaxFinancialCostAmount sets TaxFinancialCostAmount field to given value.

### HasTaxFinancialCostAmount

`func (o *SaleResponseObjectSaleResponse) HasTaxFinancialCostAmount() bool`

HasTaxFinancialCostAmount returns a boolean if a field has been set.

### GetTaxFinancialCostPercentage

`func (o *SaleResponseObjectSaleResponse) GetTaxFinancialCostPercentage() float32`

GetTaxFinancialCostPercentage returns the TaxFinancialCostPercentage field if non-nil, zero value otherwise.

### GetTaxFinancialCostPercentageOk

`func (o *SaleResponseObjectSaleResponse) GetTaxFinancialCostPercentageOk() (*float32, bool)`

GetTaxFinancialCostPercentageOk returns a tuple with the TaxFinancialCostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFinancialCostPercentage

`func (o *SaleResponseObjectSaleResponse) SetTaxFinancialCostPercentage(v float32)`

SetTaxFinancialCostPercentage sets TaxFinancialCostPercentage field to given value.

### HasTaxFinancialCostPercentage

`func (o *SaleResponseObjectSaleResponse) HasTaxFinancialCostPercentage() bool`

HasTaxFinancialCostPercentage returns a boolean if a field has been set.

### GetFinancialCostAmount

`func (o *SaleResponseObjectSaleResponse) GetFinancialCostAmount() float32`

GetFinancialCostAmount returns the FinancialCostAmount field if non-nil, zero value otherwise.

### GetFinancialCostAmountOk

`func (o *SaleResponseObjectSaleResponse) GetFinancialCostAmountOk() (*float32, bool)`

GetFinancialCostAmountOk returns a tuple with the FinancialCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialCostAmount

`func (o *SaleResponseObjectSaleResponse) SetFinancialCostAmount(v float32)`

SetFinancialCostAmount sets FinancialCostAmount field to given value.

### HasFinancialCostAmount

`func (o *SaleResponseObjectSaleResponse) HasFinancialCostAmount() bool`

HasFinancialCostAmount returns a boolean if a field has been set.

### GetFinancialCostPercentage

`func (o *SaleResponseObjectSaleResponse) GetFinancialCostPercentage() float32`

GetFinancialCostPercentage returns the FinancialCostPercentage field if non-nil, zero value otherwise.

### GetFinancialCostPercentageOk

`func (o *SaleResponseObjectSaleResponse) GetFinancialCostPercentageOk() (*float32, bool)`

GetFinancialCostPercentageOk returns a tuple with the FinancialCostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialCostPercentage

`func (o *SaleResponseObjectSaleResponse) SetFinancialCostPercentage(v float32)`

SetFinancialCostPercentage sets FinancialCostPercentage field to given value.

### HasFinancialCostPercentage

`func (o *SaleResponseObjectSaleResponse) HasFinancialCostPercentage() bool`

HasFinancialCostPercentage returns a boolean if a field has been set.

### GetHostResultCode

`func (o *SaleResponseObjectSaleResponse) GetHostResultCode() int32`

GetHostResultCode returns the HostResultCode field if non-nil, zero value otherwise.

### GetHostResultCodeOk

`func (o *SaleResponseObjectSaleResponse) GetHostResultCodeOk() (*int32, bool)`

GetHostResultCodeOk returns a tuple with the HostResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostResultCode

`func (o *SaleResponseObjectSaleResponse) SetHostResultCode(v int32)`

SetHostResultCode sets HostResultCode field to given value.

### HasHostResultCode

`func (o *SaleResponseObjectSaleResponse) HasHostResultCode() bool`

HasHostResultCode returns a boolean if a field has been set.

### GetHostMessage

`func (o *SaleResponseObjectSaleResponse) GetHostMessage() string`

GetHostMessage returns the HostMessage field if non-nil, zero value otherwise.

### GetHostMessageOk

`func (o *SaleResponseObjectSaleResponse) GetHostMessageOk() (*string, bool)`

GetHostMessageOk returns a tuple with the HostMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMessage

`func (o *SaleResponseObjectSaleResponse) SetHostMessage(v string)`

SetHostMessage sets HostMessage field to given value.

### HasHostMessage

`func (o *SaleResponseObjectSaleResponse) HasHostMessage() bool`

HasHostMessage returns a boolean if a field has been set.

### GetHostCode

`func (o *SaleResponseObjectSaleResponse) GetHostCode() string`

GetHostCode returns the HostCode field if non-nil, zero value otherwise.

### GetHostCodeOk

`func (o *SaleResponseObjectSaleResponse) GetHostCodeOk() (*string, bool)`

GetHostCodeOk returns a tuple with the HostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCode

`func (o *SaleResponseObjectSaleResponse) SetHostCode(v string)`

SetHostCode sets HostCode field to given value.

### HasHostCode

`func (o *SaleResponseObjectSaleResponse) HasHostCode() bool`

HasHostCode returns a boolean if a field has been set.

### GetHostID

`func (o *SaleResponseObjectSaleResponse) GetHostID() int32`

GetHostID returns the HostID field if non-nil, zero value otherwise.

### GetHostIDOk

`func (o *SaleResponseObjectSaleResponse) GetHostIDOk() (*int32, bool)`

GetHostIDOk returns a tuple with the HostID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostID

`func (o *SaleResponseObjectSaleResponse) SetHostID(v int32)`

SetHostID sets HostID field to given value.

### HasHostID

`func (o *SaleResponseObjectSaleResponse) HasHostID() bool`

HasHostID returns a boolean if a field has been set.

### GetUserID

`func (o *SaleResponseObjectSaleResponse) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *SaleResponseObjectSaleResponse) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *SaleResponseObjectSaleResponse) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *SaleResponseObjectSaleResponse) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetIssuerName

`func (o *SaleResponseObjectSaleResponse) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *SaleResponseObjectSaleResponse) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *SaleResponseObjectSaleResponse) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *SaleResponseObjectSaleResponse) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### GetHostDateTime

`func (o *SaleResponseObjectSaleResponse) GetHostDateTime() time.Time`

GetHostDateTime returns the HostDateTime field if non-nil, zero value otherwise.

### GetHostDateTimeOk

`func (o *SaleResponseObjectSaleResponse) GetHostDateTimeOk() (*time.Time, bool)`

GetHostDateTimeOk returns a tuple with the HostDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDateTime

`func (o *SaleResponseObjectSaleResponse) SetHostDateTime(v time.Time)`

SetHostDateTime sets HostDateTime field to given value.

### HasHostDateTime

`func (o *SaleResponseObjectSaleResponse) HasHostDateTime() bool`

HasHostDateTime returns a boolean if a field has been set.

### GetTransmitionDateTime

`func (o *SaleResponseObjectSaleResponse) GetTransmitionDateTime() time.Time`

GetTransmitionDateTime returns the TransmitionDateTime field if non-nil, zero value otherwise.

### GetTransmitionDateTimeOk

`func (o *SaleResponseObjectSaleResponse) GetTransmitionDateTimeOk() (*time.Time, bool)`

GetTransmitionDateTimeOk returns a tuple with the TransmitionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitionDateTime

`func (o *SaleResponseObjectSaleResponse) SetTransmitionDateTime(v time.Time)`

SetTransmitionDateTime sets TransmitionDateTime field to given value.

### HasTransmitionDateTime

`func (o *SaleResponseObjectSaleResponse) HasTransmitionDateTime() bool`

HasTransmitionDateTime returns a boolean if a field has been set.

### GetAuthResultCode

`func (o *SaleResponseObjectSaleResponse) GetAuthResultCode() string`

GetAuthResultCode returns the AuthResultCode field if non-nil, zero value otherwise.

### GetAuthResultCodeOk

`func (o *SaleResponseObjectSaleResponse) GetAuthResultCodeOk() (*string, bool)`

GetAuthResultCodeOk returns a tuple with the AuthResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResultCode

`func (o *SaleResponseObjectSaleResponse) SetAuthResultCode(v string)`

SetAuthResultCode sets AuthResultCode field to given value.

### HasAuthResultCode

`func (o *SaleResponseObjectSaleResponse) HasAuthResultCode() bool`

HasAuthResultCode returns a boolean if a field has been set.

### GetAuthHostProcessCode

`func (o *SaleResponseObjectSaleResponse) GetAuthHostProcessCode() string`

GetAuthHostProcessCode returns the AuthHostProcessCode field if non-nil, zero value otherwise.

### GetAuthHostProcessCodeOk

`func (o *SaleResponseObjectSaleResponse) GetAuthHostProcessCodeOk() (*string, bool)`

GetAuthHostProcessCodeOk returns a tuple with the AuthHostProcessCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHostProcessCode

`func (o *SaleResponseObjectSaleResponse) SetAuthHostProcessCode(v string)`

SetAuthHostProcessCode sets AuthHostProcessCode field to given value.

### HasAuthHostProcessCode

`func (o *SaleResponseObjectSaleResponse) HasAuthHostProcessCode() bool`

HasAuthHostProcessCode returns a boolean if a field has been set.

### GetAuthHostMsgType

`func (o *SaleResponseObjectSaleResponse) GetAuthHostMsgType() string`

GetAuthHostMsgType returns the AuthHostMsgType field if non-nil, zero value otherwise.

### GetAuthHostMsgTypeOk

`func (o *SaleResponseObjectSaleResponse) GetAuthHostMsgTypeOk() (*string, bool)`

GetAuthHostMsgTypeOk returns a tuple with the AuthHostMsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHostMsgType

`func (o *SaleResponseObjectSaleResponse) SetAuthHostMsgType(v string)`

SetAuthHostMsgType sets AuthHostMsgType field to given value.

### HasAuthHostMsgType

`func (o *SaleResponseObjectSaleResponse) HasAuthHostMsgType() bool`

HasAuthHostMsgType returns a boolean if a field has been set.

### GetAuthHostMessage

`func (o *SaleResponseObjectSaleResponse) GetAuthHostMessage() string`

GetAuthHostMessage returns the AuthHostMessage field if non-nil, zero value otherwise.

### GetAuthHostMessageOk

`func (o *SaleResponseObjectSaleResponse) GetAuthHostMessageOk() (*string, bool)`

GetAuthHostMessageOk returns a tuple with the AuthHostMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHostMessage

`func (o *SaleResponseObjectSaleResponse) SetAuthHostMessage(v string)`

SetAuthHostMessage sets AuthHostMessage field to given value.

### HasAuthHostMessage

`func (o *SaleResponseObjectSaleResponse) HasAuthHostMessage() bool`

HasAuthHostMessage returns a boolean if a field has been set.

### GetHostMsgType

`func (o *SaleResponseObjectSaleResponse) GetHostMsgType() string`

GetHostMsgType returns the HostMsgType field if non-nil, zero value otherwise.

### GetHostMsgTypeOk

`func (o *SaleResponseObjectSaleResponse) GetHostMsgTypeOk() (*string, bool)`

GetHostMsgTypeOk returns a tuple with the HostMsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMsgType

`func (o *SaleResponseObjectSaleResponse) SetHostMsgType(v string)`

SetHostMsgType sets HostMsgType field to given value.

### HasHostMsgType

`func (o *SaleResponseObjectSaleResponse) HasHostMsgType() bool`

HasHostMsgType returns a boolean if a field has been set.

### GetAuthCode

`func (o *SaleResponseObjectSaleResponse) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *SaleResponseObjectSaleResponse) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *SaleResponseObjectSaleResponse) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *SaleResponseObjectSaleResponse) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetAuthDateTime

`func (o *SaleResponseObjectSaleResponse) GetAuthDateTime() time.Time`

GetAuthDateTime returns the AuthDateTime field if non-nil, zero value otherwise.

### GetAuthDateTimeOk

`func (o *SaleResponseObjectSaleResponse) GetAuthDateTimeOk() (*time.Time, bool)`

GetAuthDateTimeOk returns a tuple with the AuthDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDateTime

`func (o *SaleResponseObjectSaleResponse) SetAuthDateTime(v time.Time)`

SetAuthDateTime sets AuthDateTime field to given value.

### HasAuthDateTime

`func (o *SaleResponseObjectSaleResponse) HasAuthDateTime() bool`

HasAuthDateTime returns a boolean if a field has been set.

### GetAuthTicket

`func (o *SaleResponseObjectSaleResponse) GetAuthTicket() int32`

GetAuthTicket returns the AuthTicket field if non-nil, zero value otherwise.

### GetAuthTicketOk

`func (o *SaleResponseObjectSaleResponse) GetAuthTicketOk() (*int32, bool)`

GetAuthTicketOk returns a tuple with the AuthTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTicket

`func (o *SaleResponseObjectSaleResponse) SetAuthTicket(v int32)`

SetAuthTicket sets AuthTicket field to given value.

### HasAuthTicket

`func (o *SaleResponseObjectSaleResponse) HasAuthTicket() bool`

HasAuthTicket returns a boolean if a field has been set.

### GetAuthRRN

`func (o *SaleResponseObjectSaleResponse) GetAuthRRN() string`

GetAuthRRN returns the AuthRRN field if non-nil, zero value otherwise.

### GetAuthRRNOk

`func (o *SaleResponseObjectSaleResponse) GetAuthRRNOk() (*string, bool)`

GetAuthRRNOk returns a tuple with the AuthRRN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRRN

`func (o *SaleResponseObjectSaleResponse) SetAuthRRN(v string)`

SetAuthRRN sets AuthRRN field to given value.

### HasAuthRRN

`func (o *SaleResponseObjectSaleResponse) HasAuthRRN() bool`

HasAuthRRN returns a boolean if a field has been set.

### GetTransactionAuthenticationType

`func (o *SaleResponseObjectSaleResponse) GetTransactionAuthenticationType() string`

GetTransactionAuthenticationType returns the TransactionAuthenticationType field if non-nil, zero value otherwise.

### GetTransactionAuthenticationTypeOk

`func (o *SaleResponseObjectSaleResponse) GetTransactionAuthenticationTypeOk() (*string, bool)`

GetTransactionAuthenticationTypeOk returns a tuple with the TransactionAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAuthenticationType

`func (o *SaleResponseObjectSaleResponse) SetTransactionAuthenticationType(v string)`

SetTransactionAuthenticationType sets TransactionAuthenticationType field to given value.

### HasTransactionAuthenticationType

`func (o *SaleResponseObjectSaleResponse) HasTransactionAuthenticationType() bool`

HasTransactionAuthenticationType returns a boolean if a field has been set.

### GetIdentifierForTheAdquirer

`func (o *SaleResponseObjectSaleResponse) GetIdentifierForTheAdquirer() string`

GetIdentifierForTheAdquirer returns the IdentifierForTheAdquirer field if non-nil, zero value otherwise.

### GetIdentifierForTheAdquirerOk

`func (o *SaleResponseObjectSaleResponse) GetIdentifierForTheAdquirerOk() (*string, bool)`

GetIdentifierForTheAdquirerOk returns a tuple with the IdentifierForTheAdquirer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierForTheAdquirer

`func (o *SaleResponseObjectSaleResponse) SetIdentifierForTheAdquirer(v string)`

SetIdentifierForTheAdquirer sets IdentifierForTheAdquirer field to given value.

### HasIdentifierForTheAdquirer

`func (o *SaleResponseObjectSaleResponse) HasIdentifierForTheAdquirer() bool`

HasIdentifierForTheAdquirer returns a boolean if a field has been set.

### GetIdentifierForTheResolutor

`func (o *SaleResponseObjectSaleResponse) GetIdentifierForTheResolutor() string`

GetIdentifierForTheResolutor returns the IdentifierForTheResolutor field if non-nil, zero value otherwise.

### GetIdentifierForTheResolutorOk

`func (o *SaleResponseObjectSaleResponse) GetIdentifierForTheResolutorOk() (*string, bool)`

GetIdentifierForTheResolutorOk returns a tuple with the IdentifierForTheResolutor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierForTheResolutor

`func (o *SaleResponseObjectSaleResponse) SetIdentifierForTheResolutor(v string)`

SetIdentifierForTheResolutor sets IdentifierForTheResolutor field to given value.

### HasIdentifierForTheResolutor

`func (o *SaleResponseObjectSaleResponse) HasIdentifierForTheResolutor() bool`

HasIdentifierForTheResolutor returns a boolean if a field has been set.

### GetPaymentFacilitatorID

`func (o *SaleResponseObjectSaleResponse) GetPaymentFacilitatorID() string`

GetPaymentFacilitatorID returns the PaymentFacilitatorID field if non-nil, zero value otherwise.

### GetPaymentFacilitatorIDOk

`func (o *SaleResponseObjectSaleResponse) GetPaymentFacilitatorIDOk() (*string, bool)`

GetPaymentFacilitatorIDOk returns a tuple with the PaymentFacilitatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFacilitatorID

`func (o *SaleResponseObjectSaleResponse) SetPaymentFacilitatorID(v string)`

SetPaymentFacilitatorID sets PaymentFacilitatorID field to given value.

### HasPaymentFacilitatorID

`func (o *SaleResponseObjectSaleResponse) HasPaymentFacilitatorID() bool`

HasPaymentFacilitatorID returns a boolean if a field has been set.

### GetMerchantID

`func (o *SaleResponseObjectSaleResponse) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *SaleResponseObjectSaleResponse) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *SaleResponseObjectSaleResponse) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *SaleResponseObjectSaleResponse) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetTerminalID

`func (o *SaleResponseObjectSaleResponse) GetTerminalID() string`

GetTerminalID returns the TerminalID field if non-nil, zero value otherwise.

### GetTerminalIDOk

`func (o *SaleResponseObjectSaleResponse) GetTerminalIDOk() (*string, bool)`

GetTerminalIDOk returns a tuple with the TerminalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalID

`func (o *SaleResponseObjectSaleResponse) SetTerminalID(v string)`

SetTerminalID sets TerminalID field to given value.

### HasTerminalID

`func (o *SaleResponseObjectSaleResponse) HasTerminalID() bool`

HasTerminalID returns a boolean if a field has been set.

### GetTerminalTrace

`func (o *SaleResponseObjectSaleResponse) GetTerminalTrace() int32`

GetTerminalTrace returns the TerminalTrace field if non-nil, zero value otherwise.

### GetTerminalTraceOk

`func (o *SaleResponseObjectSaleResponse) GetTerminalTraceOk() (*int32, bool)`

GetTerminalTraceOk returns a tuple with the TerminalTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalTrace

`func (o *SaleResponseObjectSaleResponse) SetTerminalTrace(v int32)`

SetTerminalTrace sets TerminalTrace field to given value.

### HasTerminalTrace

`func (o *SaleResponseObjectSaleResponse) HasTerminalTrace() bool`

HasTerminalTrace returns a boolean if a field has been set.

### GetSettlementBatchNumber

`func (o *SaleResponseObjectSaleResponse) GetSettlementBatchNumber() int32`

GetSettlementBatchNumber returns the SettlementBatchNumber field if non-nil, zero value otherwise.

### GetSettlementBatchNumberOk

`func (o *SaleResponseObjectSaleResponse) GetSettlementBatchNumberOk() (*int32, bool)`

GetSettlementBatchNumberOk returns a tuple with the SettlementBatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementBatchNumber

`func (o *SaleResponseObjectSaleResponse) SetSettlementBatchNumber(v int32)`

SetSettlementBatchNumber sets SettlementBatchNumber field to given value.

### HasSettlementBatchNumber

`func (o *SaleResponseObjectSaleResponse) HasSettlementBatchNumber() bool`

HasSettlementBatchNumber returns a boolean if a field has been set.

### GetCardReadMode

`func (o *SaleResponseObjectSaleResponse) GetCardReadMode() string`

GetCardReadMode returns the CardReadMode field if non-nil, zero value otherwise.

### GetCardReadModeOk

`func (o *SaleResponseObjectSaleResponse) GetCardReadModeOk() (*string, bool)`

GetCardReadModeOk returns a tuple with the CardReadMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardReadMode

`func (o *SaleResponseObjectSaleResponse) SetCardReadMode(v string)`

SetCardReadMode sets CardReadMode field to given value.

### HasCardReadMode

`func (o *SaleResponseObjectSaleResponse) HasCardReadMode() bool`

HasCardReadMode returns a boolean if a field has been set.

### GetCardReadModeDescription

`func (o *SaleResponseObjectSaleResponse) GetCardReadModeDescription() string`

GetCardReadModeDescription returns the CardReadModeDescription field if non-nil, zero value otherwise.

### GetCardReadModeDescriptionOk

`func (o *SaleResponseObjectSaleResponse) GetCardReadModeDescriptionOk() (*string, bool)`

GetCardReadModeDescriptionOk returns a tuple with the CardReadModeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardReadModeDescription

`func (o *SaleResponseObjectSaleResponse) SetCardReadModeDescription(v string)`

SetCardReadModeDescription sets CardReadModeDescription field to given value.

### HasCardReadModeDescription

`func (o *SaleResponseObjectSaleResponse) HasCardReadModeDescription() bool`

HasCardReadModeDescription returns a boolean if a field has been set.

### GetCardDescription

`func (o *SaleResponseObjectSaleResponse) GetCardDescription() string`

GetCardDescription returns the CardDescription field if non-nil, zero value otherwise.

### GetCardDescriptionOk

`func (o *SaleResponseObjectSaleResponse) GetCardDescriptionOk() (*string, bool)`

GetCardDescriptionOk returns a tuple with the CardDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDescription

`func (o *SaleResponseObjectSaleResponse) SetCardDescription(v string)`

SetCardDescription sets CardDescription field to given value.

### HasCardDescription

`func (o *SaleResponseObjectSaleResponse) HasCardDescription() bool`

HasCardDescription returns a boolean if a field has been set.

### GetCardTypeDescription

`func (o *SaleResponseObjectSaleResponse) GetCardTypeDescription() string`

GetCardTypeDescription returns the CardTypeDescription field if non-nil, zero value otherwise.

### GetCardTypeDescriptionOk

`func (o *SaleResponseObjectSaleResponse) GetCardTypeDescriptionOk() (*string, bool)`

GetCardTypeDescriptionOk returns a tuple with the CardTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardTypeDescription

`func (o *SaleResponseObjectSaleResponse) SetCardTypeDescription(v string)`

SetCardTypeDescription sets CardTypeDescription field to given value.

### HasCardTypeDescription

`func (o *SaleResponseObjectSaleResponse) HasCardTypeDescription() bool`

HasCardTypeDescription returns a boolean if a field has been set.

### GetCardCategory

`func (o *SaleResponseObjectSaleResponse) GetCardCategory() SaleResponseObjectSaleResponseCardCategory`

GetCardCategory returns the CardCategory field if non-nil, zero value otherwise.

### GetCardCategoryOk

`func (o *SaleResponseObjectSaleResponse) GetCardCategoryOk() (*SaleResponseObjectSaleResponseCardCategory, bool)`

GetCardCategoryOk returns a tuple with the CardCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCategory

`func (o *SaleResponseObjectSaleResponse) SetCardCategory(v SaleResponseObjectSaleResponseCardCategory)`

SetCardCategory sets CardCategory field to given value.

### HasCardCategory

`func (o *SaleResponseObjectSaleResponse) HasCardCategory() bool`

HasCardCategory returns a boolean if a field has been set.

### GetCardNumber

`func (o *SaleResponseObjectSaleResponse) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *SaleResponseObjectSaleResponse) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *SaleResponseObjectSaleResponse) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *SaleResponseObjectSaleResponse) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNumberMasked

`func (o *SaleResponseObjectSaleResponse) GetCardNumberMasked() string`

GetCardNumberMasked returns the CardNumberMasked field if non-nil, zero value otherwise.

### GetCardNumberMaskedOk

`func (o *SaleResponseObjectSaleResponse) GetCardNumberMaskedOk() (*string, bool)`

GetCardNumberMaskedOk returns a tuple with the CardNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberMasked

`func (o *SaleResponseObjectSaleResponse) SetCardNumberMasked(v string)`

SetCardNumberMasked sets CardNumberMasked field to given value.

### HasCardNumberMasked

`func (o *SaleResponseObjectSaleResponse) HasCardNumberMasked() bool`

HasCardNumberMasked returns a boolean if a field has been set.

### GetCardHashing

`func (o *SaleResponseObjectSaleResponse) GetCardHashing() string`

GetCardHashing returns the CardHashing field if non-nil, zero value otherwise.

### GetCardHashingOk

`func (o *SaleResponseObjectSaleResponse) GetCardHashingOk() (*string, bool)`

GetCardHashingOk returns a tuple with the CardHashing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHashing

`func (o *SaleResponseObjectSaleResponse) SetCardHashing(v string)`

SetCardHashing sets CardHashing field to given value.

### HasCardHashing

`func (o *SaleResponseObjectSaleResponse) HasCardHashing() bool`

HasCardHashing returns a boolean if a field has been set.

### GetCardExp

`func (o *SaleResponseObjectSaleResponse) GetCardExp() string`

GetCardExp returns the CardExp field if non-nil, zero value otherwise.

### GetCardExpOk

`func (o *SaleResponseObjectSaleResponse) GetCardExpOk() (*string, bool)`

GetCardExpOk returns a tuple with the CardExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExp

`func (o *SaleResponseObjectSaleResponse) SetCardExp(v string)`

SetCardExp sets CardExp field to given value.

### HasCardExp

`func (o *SaleResponseObjectSaleResponse) HasCardExp() bool`

HasCardExp returns a boolean if a field has been set.

### GetCardCryptogramResponse

`func (o *SaleResponseObjectSaleResponse) GetCardCryptogramResponse() string`

GetCardCryptogramResponse returns the CardCryptogramResponse field if non-nil, zero value otherwise.

### GetCardCryptogramResponseOk

`func (o *SaleResponseObjectSaleResponse) GetCardCryptogramResponseOk() (*string, bool)`

GetCardCryptogramResponseOk returns a tuple with the CardCryptogramResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCryptogramResponse

`func (o *SaleResponseObjectSaleResponse) SetCardCryptogramResponse(v string)`

SetCardCryptogramResponse sets CardCryptogramResponse field to given value.

### HasCardCryptogramResponse

`func (o *SaleResponseObjectSaleResponse) HasCardCryptogramResponse() bool`

HasCardCryptogramResponse returns a boolean if a field has been set.

### GetCardAppName

`func (o *SaleResponseObjectSaleResponse) GetCardAppName() string`

GetCardAppName returns the CardAppName field if non-nil, zero value otherwise.

### GetCardAppNameOk

`func (o *SaleResponseObjectSaleResponse) GetCardAppNameOk() (*string, bool)`

GetCardAppNameOk returns a tuple with the CardAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppName

`func (o *SaleResponseObjectSaleResponse) SetCardAppName(v string)`

SetCardAppName sets CardAppName field to given value.

### HasCardAppName

`func (o *SaleResponseObjectSaleResponse) HasCardAppName() bool`

HasCardAppName returns a boolean if a field has been set.

### GetCardAppIdentifier

`func (o *SaleResponseObjectSaleResponse) GetCardAppIdentifier() string`

GetCardAppIdentifier returns the CardAppIdentifier field if non-nil, zero value otherwise.

### GetCardAppIdentifierOk

`func (o *SaleResponseObjectSaleResponse) GetCardAppIdentifierOk() (*string, bool)`

GetCardAppIdentifierOk returns a tuple with the CardAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppIdentifier

`func (o *SaleResponseObjectSaleResponse) SetCardAppIdentifier(v string)`

SetCardAppIdentifier sets CardAppIdentifier field to given value.

### HasCardAppIdentifier

`func (o *SaleResponseObjectSaleResponse) HasCardAppIdentifier() bool`

HasCardAppIdentifier returns a boolean if a field has been set.

### GetCardAppLabel

`func (o *SaleResponseObjectSaleResponse) GetCardAppLabel() string`

GetCardAppLabel returns the CardAppLabel field if non-nil, zero value otherwise.

### GetCardAppLabelOk

`func (o *SaleResponseObjectSaleResponse) GetCardAppLabelOk() (*string, bool)`

GetCardAppLabelOk returns a tuple with the CardAppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppLabel

`func (o *SaleResponseObjectSaleResponse) SetCardAppLabel(v string)`

SetCardAppLabel sets CardAppLabel field to given value.

### HasCardAppLabel

`func (o *SaleResponseObjectSaleResponse) HasCardAppLabel() bool`

HasCardAppLabel returns a boolean if a field has been set.

### GetCardAuthRequestCryptogram

`func (o *SaleResponseObjectSaleResponse) GetCardAuthRequestCryptogram() string`

GetCardAuthRequestCryptogram returns the CardAuthRequestCryptogram field if non-nil, zero value otherwise.

### GetCardAuthRequestCryptogramOk

`func (o *SaleResponseObjectSaleResponse) GetCardAuthRequestCryptogramOk() (*string, bool)`

GetCardAuthRequestCryptogramOk returns a tuple with the CardAuthRequestCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthRequestCryptogram

`func (o *SaleResponseObjectSaleResponse) SetCardAuthRequestCryptogram(v string)`

SetCardAuthRequestCryptogram sets CardAuthRequestCryptogram field to given value.

### HasCardAuthRequestCryptogram

`func (o *SaleResponseObjectSaleResponse) HasCardAuthRequestCryptogram() bool`

HasCardAuthRequestCryptogram returns a boolean if a field has been set.

### GetCardAuthResponseCryptogram

`func (o *SaleResponseObjectSaleResponse) GetCardAuthResponseCryptogram() string`

GetCardAuthResponseCryptogram returns the CardAuthResponseCryptogram field if non-nil, zero value otherwise.

### GetCardAuthResponseCryptogramOk

`func (o *SaleResponseObjectSaleResponse) GetCardAuthResponseCryptogramOk() (*string, bool)`

GetCardAuthResponseCryptogramOk returns a tuple with the CardAuthResponseCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthResponseCryptogram

`func (o *SaleResponseObjectSaleResponse) SetCardAuthResponseCryptogram(v string)`

SetCardAuthResponseCryptogram sets CardAuthResponseCryptogram field to given value.

### HasCardAuthResponseCryptogram

`func (o *SaleResponseObjectSaleResponse) HasCardAuthResponseCryptogram() bool`

HasCardAuthResponseCryptogram returns a boolean if a field has been set.

### GetPayer

`func (o *SaleResponseObjectSaleResponse) GetPayer() SaleObjectSalePayer`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *SaleResponseObjectSaleResponse) GetPayerOk() (*SaleObjectSalePayer, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *SaleResponseObjectSaleResponse) SetPayer(v SaleObjectSalePayer)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *SaleResponseObjectSaleResponse) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetCustomer

`func (o *SaleResponseObjectSaleResponse) GetCustomer() SaleObjectSaleCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *SaleResponseObjectSaleResponse) GetCustomerOk() (*SaleObjectSaleCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *SaleResponseObjectSaleResponse) SetCustomer(v SaleObjectSaleCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *SaleResponseObjectSaleResponse) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetMerchantCategory

`func (o *SaleResponseObjectSaleResponse) GetMerchantCategory() SaleResponseObjectSaleResponseMerchantCategory`

GetMerchantCategory returns the MerchantCategory field if non-nil, zero value otherwise.

### GetMerchantCategoryOk

`func (o *SaleResponseObjectSaleResponse) GetMerchantCategoryOk() (*SaleResponseObjectSaleResponseMerchantCategory, bool)`

GetMerchantCategoryOk returns a tuple with the MerchantCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategory

`func (o *SaleResponseObjectSaleResponse) SetMerchantCategory(v SaleResponseObjectSaleResponseMerchantCategory)`

SetMerchantCategory sets MerchantCategory field to given value.

### HasMerchantCategory

`func (o *SaleResponseObjectSaleResponse) HasMerchantCategory() bool`

HasMerchantCategory returns a boolean if a field has been set.

### GetProducts

`func (o *SaleResponseObjectSaleResponse) GetProducts() []SaleResponseObjectSaleResponseProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *SaleResponseObjectSaleResponse) GetProductsOk() (*[]SaleResponseObjectSaleResponseProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *SaleResponseObjectSaleResponse) SetProducts(v []SaleResponseObjectSaleResponseProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *SaleResponseObjectSaleResponse) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *SaleResponseObjectSaleResponse) GetPaymentMethod() SaleResponseObjectSaleResponsePaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *SaleResponseObjectSaleResponse) GetPaymentMethodOk() (*SaleResponseObjectSaleResponsePaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *SaleResponseObjectSaleResponse) SetPaymentMethod(v SaleResponseObjectSaleResponsePaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *SaleResponseObjectSaleResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPlans

`func (o *SaleResponseObjectSaleResponse) GetPlans() SaleResponseObjectSaleResponsePlans`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *SaleResponseObjectSaleResponse) GetPlansOk() (*SaleResponseObjectSaleResponsePlans, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *SaleResponseObjectSaleResponse) SetPlans(v SaleResponseObjectSaleResponsePlans)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *SaleResponseObjectSaleResponse) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetPlanDescription

`func (o *SaleResponseObjectSaleResponse) GetPlanDescription() string`

GetPlanDescription returns the PlanDescription field if non-nil, zero value otherwise.

### GetPlanDescriptionOk

`func (o *SaleResponseObjectSaleResponse) GetPlanDescriptionOk() (*string, bool)`

GetPlanDescriptionOk returns a tuple with the PlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDescription

`func (o *SaleResponseObjectSaleResponse) SetPlanDescription(v string)`

SetPlanDescription sets PlanDescription field to given value.

### HasPlanDescription

`func (o *SaleResponseObjectSaleResponse) HasPlanDescription() bool`

HasPlanDescription returns a boolean if a field has been set.

### GetPlanConfigVersion

`func (o *SaleResponseObjectSaleResponse) GetPlanConfigVersion() int32`

GetPlanConfigVersion returns the PlanConfigVersion field if non-nil, zero value otherwise.

### GetPlanConfigVersionOk

`func (o *SaleResponseObjectSaleResponse) GetPlanConfigVersionOk() (*int32, bool)`

GetPlanConfigVersionOk returns a tuple with the PlanConfigVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanConfigVersion

`func (o *SaleResponseObjectSaleResponse) SetPlanConfigVersion(v int32)`

SetPlanConfigVersion sets PlanConfigVersion field to given value.

### HasPlanConfigVersion

`func (o *SaleResponseObjectSaleResponse) HasPlanConfigVersion() bool`

HasPlanConfigVersion returns a boolean if a field has been set.

### GetTickets

`func (o *SaleResponseObjectSaleResponse) GetTickets() []SaleResponseObjectSaleResponseTickets`

GetTickets returns the Tickets field if non-nil, zero value otherwise.

### GetTicketsOk

`func (o *SaleResponseObjectSaleResponse) GetTicketsOk() (*[]SaleResponseObjectSaleResponseTickets, bool)`

GetTicketsOk returns a tuple with the Tickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickets

`func (o *SaleResponseObjectSaleResponse) SetTickets(v []SaleResponseObjectSaleResponseTickets)`

SetTickets sets Tickets field to given value.

### HasTickets

`func (o *SaleResponseObjectSaleResponse) HasTickets() bool`

HasTickets returns a boolean if a field has been set.

### GetConfiguration

`func (o *SaleResponseObjectSaleResponse) GetConfiguration() SaleResponseObjectSaleResponseConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SaleResponseObjectSaleResponse) GetConfigurationOk() (*SaleResponseObjectSaleResponseConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SaleResponseObjectSaleResponse) SetConfiguration(v SaleResponseObjectSaleResponseConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *SaleResponseObjectSaleResponse) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *SaleResponseObjectSaleResponse) GetPaymentMethodId() int32`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *SaleResponseObjectSaleResponse) GetPaymentMethodIdOk() (*int32, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *SaleResponseObjectSaleResponse) SetPaymentMethodId(v int32)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *SaleResponseObjectSaleResponse) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetPaymentMethodDescription

`func (o *SaleResponseObjectSaleResponse) GetPaymentMethodDescription() string`

GetPaymentMethodDescription returns the PaymentMethodDescription field if non-nil, zero value otherwise.

### GetPaymentMethodDescriptionOk

`func (o *SaleResponseObjectSaleResponse) GetPaymentMethodDescriptionOk() (*string, bool)`

GetPaymentMethodDescriptionOk returns a tuple with the PaymentMethodDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodDescription

`func (o *SaleResponseObjectSaleResponse) SetPaymentMethodDescription(v string)`

SetPaymentMethodDescription sets PaymentMethodDescription field to given value.

### HasPaymentMethodDescription

`func (o *SaleResponseObjectSaleResponse) HasPaymentMethodDescription() bool`

HasPaymentMethodDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


