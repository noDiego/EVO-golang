# GetCardResponseObjectGetCardResponse

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
**DisplayResponseMessage** | Pointer to **[]string** | Información adicional/Mensaje promocional/Leyenda de respuesta a mostrar en pantalla en el ticket de la operación. Cada línea de este mensaje será un elemento dentro del array. | [optional] 
**TransmitionDateTime** | Pointer to **time.Time** | Fecha y hora de transmisión de la operación hacia el host - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**AuthResultCode** | Pointer to **string** | Código de Resultado retornado por el Host Adquirente. | [optional] 
**AuthHostProcessCode** | Pointer to **string** | Código de procesamiento de la operación enviada al host. Elemento 3 del protocolo de comunicaciones ISO 8583, formato de mensajes utilizado por los hosts para realizar operaciones financieras. | [optional] 
**AuthHostMsgType** | Pointer to **string** | Elemento 0 del protocolo de comunicaciones ISO 8583, formato de mensajes utilizado por los hosts para realizar operaciones financieras. El valor de este campo es el que devuelve el host en una respuesta a una petición. | [optional] 
**HostMsgType** | Pointer to **string** | Elemento 0 del protocolo de comunicaciones ISO 8583, formato de mensajes utilizado por los hosts para realizar operaciones financieras. El valor de este campo es el que se envio al host en el envio de una petición. | [optional] 
**AuthCode** | Pointer to **string** | Código de autorización retornado por el Host que resuelve la transacción. | [optional] 
**AuthDateTime** | Pointer to **time.Time** | Fecha y Hora de la transacción retornada por el Host que resuelve la Transacción - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**AuthRRN** | Pointer to **string** | Número de identificación de la transacción, utilizado por la mayoría de los hosts para realizar anulaciones y devoluciones. | [optional] 
**AuthenticationType** | Pointer to **string** | Tipo de autenticación | [optional] 
**AuthHostMessage** | Pointer to **string** | Mensaje Retornado por el Host Adquirente, normalmente asociado al valor de AuthResultCode | [optional] 
**AuthTicket** | Pointer to **int32** | Número Ticket  o Voucher Generado para la Plataforma. | [optional] 
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
**MerchantCategory** | Pointer to [**SaleResponseObjectSaleResponseMerchantCategory**](SaleResponseObjectSaleResponseMerchantCategory.md) |  | [optional] 
**PaymentMethod** | Pointer to [**SaleResponseObjectSaleResponsePaymentMethod**](SaleResponseObjectSaleResponsePaymentMethod.md) |  | [optional] 
**Plans** | Pointer to [**SaleResponseObjectSaleResponsePlans**](SaleResponseObjectSaleResponsePlans.md) |  | [optional] 
**Configuration** | Pointer to [**SaleResponseObjectSaleResponseConfiguration**](SaleResponseObjectSaleResponseConfiguration.md) |  | [optional] 
**Track1** | Pointer to **string** | Banda Número 1 de la tarjeta. Este dato sera necesario si el modo de ingreso fue por banda (deslizando la banda por el lector). | [optional] 
**Track2** | Pointer to **string** | Banda Número 2 de la tarjeta. Este dato sera necesario si el modo de ingreso fue por banda (deslizando la banda por el lector). | [optional] 
**SecurityCode** | Pointer to **string** | código de seguridad de la tarjeta. Este dato solo sera requerido si la parametrización de la marca a la que la tarjeta corresponde lo indique. | [optional] 
**CardCryptogram** | Pointer to **string** | Criptograma obtenido del Chip | [optional] 
**SettlementBatchNumber** | Pointer to **int32** | Para aquellos host que exista el concepto de Lote, es el número de lote al cual pertenece la transacción. | [optional] 
**TerminalID** | Pointer to **string** | Identificador de Terminal por el cual se envía la Transacción al Host. | [optional] 
**TerminalTrace** | Pointer to **int32** | Número de Trace/Secuencia que genera la plataforma para la transacción asociado al TerminalID. | [optional] 
**PaymentFacilitatorID** | Pointer to **string** | Identificador de facilitador de pagos o Payfac. | [optional] 
**MerchantID** | Pointer to **string** | Número de comercio utilizado para realizar la transacción. Este Número es asignado por el host, y parametrizado en la BD, relaciónado a cada uno de los planes disponibles. | [optional] 
**CardHolderName** | Pointer to **string** | Nombre del tarjetahabiente obtenido de la tarjeta. | [optional] 

## Methods

### NewGetCardResponseObjectGetCardResponse

`func NewGetCardResponseObjectGetCardResponse(responseCode int32, responseActions []string, responseMessage string, ) *GetCardResponseObjectGetCardResponse`

NewGetCardResponseObjectGetCardResponse instantiates a new GetCardResponseObjectGetCardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCardResponseObjectGetCardResponseWithDefaults

`func NewGetCardResponseObjectGetCardResponseWithDefaults() *GetCardResponseObjectGetCardResponse`

NewGetCardResponseObjectGetCardResponseWithDefaults instantiates a new GetCardResponseObjectGetCardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *GetCardResponseObjectGetCardResponse) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *GetCardResponseObjectGetCardResponse) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *GetCardResponseObjectGetCardResponse) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseActions

`func (o *GetCardResponseObjectGetCardResponse) GetResponseActions() []string`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *GetCardResponseObjectGetCardResponse) GetResponseActionsOk() (*[]string, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *GetCardResponseObjectGetCardResponse) SetResponseActions(v []string)`

SetResponseActions sets ResponseActions field to given value.


### GetResponseMessage

`func (o *GetCardResponseObjectGetCardResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *GetCardResponseObjectGetCardResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *GetCardResponseObjectGetCardResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetCompanyIdentification

`func (o *GetCardResponseObjectGetCardResponse) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *GetCardResponseObjectGetCardResponse) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *GetCardResponseObjectGetCardResponse) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *GetCardResponseObjectGetCardResponse) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetSystemIdentification

`func (o *GetCardResponseObjectGetCardResponse) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *GetCardResponseObjectGetCardResponse) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *GetCardResponseObjectGetCardResponse) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.

### HasSystemIdentification

`func (o *GetCardResponseObjectGetCardResponse) HasSystemIdentification() bool`

HasSystemIdentification returns a boolean if a field has been set.

### GetBranchIdentification

`func (o *GetCardResponseObjectGetCardResponse) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *GetCardResponseObjectGetCardResponse) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *GetCardResponseObjectGetCardResponse) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *GetCardResponseObjectGetCardResponse) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *GetCardResponseObjectGetCardResponse) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *GetCardResponseObjectGetCardResponse) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *GetCardResponseObjectGetCardResponse) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *GetCardResponseObjectGetCardResponse) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetForeignResponseCode

`func (o *GetCardResponseObjectGetCardResponse) GetForeignResponseCode() string`

GetForeignResponseCode returns the ForeignResponseCode field if non-nil, zero value otherwise.

### GetForeignResponseCodeOk

`func (o *GetCardResponseObjectGetCardResponse) GetForeignResponseCodeOk() (*string, bool)`

GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignResponseCode

`func (o *GetCardResponseObjectGetCardResponse) SetForeignResponseCode(v string)`

SetForeignResponseCode sets ForeignResponseCode field to given value.

### HasForeignResponseCode

`func (o *GetCardResponseObjectGetCardResponse) HasForeignResponseCode() bool`

HasForeignResponseCode returns a boolean if a field has been set.

### GetRequestType

`func (o *GetCardResponseObjectGetCardResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *GetCardResponseObjectGetCardResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *GetCardResponseObjectGetCardResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *GetCardResponseObjectGetCardResponse) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetServiceVersion

`func (o *GetCardResponseObjectGetCardResponse) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *GetCardResponseObjectGetCardResponse) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *GetCardResponseObjectGetCardResponse) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *GetCardResponseObjectGetCardResponse) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *GetCardResponseObjectGetCardResponse) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetCardResponseObjectGetCardResponse) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetCardResponseObjectGetCardResponse) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *GetCardResponseObjectGetCardResponse) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *GetCardResponseObjectGetCardResponse) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *GetCardResponseObjectGetCardResponse) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *GetCardResponseObjectGetCardResponse) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *GetCardResponseObjectGetCardResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetBlock

`func (o *GetCardResponseObjectGetCardResponse) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *GetCardResponseObjectGetCardResponse) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *GetCardResponseObjectGetCardResponse) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *GetCardResponseObjectGetCardResponse) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetRequestKey

`func (o *GetCardResponseObjectGetCardResponse) GetRequestKey() string`

GetRequestKey returns the RequestKey field if non-nil, zero value otherwise.

### GetRequestKeyOk

`func (o *GetCardResponseObjectGetCardResponse) GetRequestKeyOk() (*string, bool)`

GetRequestKeyOk returns a tuple with the RequestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestKey

`func (o *GetCardResponseObjectGetCardResponse) SetRequestKey(v string)`

SetRequestKey sets RequestKey field to given value.

### HasRequestKey

`func (o *GetCardResponseObjectGetCardResponse) HasRequestKey() bool`

HasRequestKey returns a boolean if a field has been set.

### GetWorkingKeys

`func (o *GetCardResponseObjectGetCardResponse) GetWorkingKeys() []SaleResponseObjectSaleResponseWorkingKeys`

GetWorkingKeys returns the WorkingKeys field if non-nil, zero value otherwise.

### GetWorkingKeysOk

`func (o *GetCardResponseObjectGetCardResponse) GetWorkingKeysOk() (*[]SaleResponseObjectSaleResponseWorkingKeys, bool)`

GetWorkingKeysOk returns a tuple with the WorkingKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingKeys

`func (o *GetCardResponseObjectGetCardResponse) SetWorkingKeys(v []SaleResponseObjectSaleResponseWorkingKeys)`

SetWorkingKeys sets WorkingKeys field to given value.

### HasWorkingKeys

`func (o *GetCardResponseObjectGetCardResponse) HasWorkingKeys() bool`

HasWorkingKeys returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *GetCardResponseObjectGetCardResponse) GetRequiredInformation() []SaleResponseObjectSaleResponseRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *GetCardResponseObjectGetCardResponse) GetRequiredInformationOk() (*[]SaleResponseObjectSaleResponseRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *GetCardResponseObjectGetCardResponse) SetRequiredInformation(v []SaleResponseObjectSaleResponseRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *GetCardResponseObjectGetCardResponse) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetAnswerType

`func (o *GetCardResponseObjectGetCardResponse) GetAnswerType() string`

GetAnswerType returns the AnswerType field if non-nil, zero value otherwise.

### GetAnswerTypeOk

`func (o *GetCardResponseObjectGetCardResponse) GetAnswerTypeOk() (*string, bool)`

GetAnswerTypeOk returns a tuple with the AnswerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerType

`func (o *GetCardResponseObjectGetCardResponse) SetAnswerType(v string)`

SetAnswerType sets AnswerType field to given value.

### HasAnswerType

`func (o *GetCardResponseObjectGetCardResponse) HasAnswerType() bool`

HasAnswerType returns a boolean if a field has been set.

### GetAnswerKey

`func (o *GetCardResponseObjectGetCardResponse) GetAnswerKey() string`

GetAnswerKey returns the AnswerKey field if non-nil, zero value otherwise.

### GetAnswerKeyOk

`func (o *GetCardResponseObjectGetCardResponse) GetAnswerKeyOk() (*string, bool)`

GetAnswerKeyOk returns a tuple with the AnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerKey

`func (o *GetCardResponseObjectGetCardResponse) SetAnswerKey(v string)`

SetAnswerKey sets AnswerKey field to given value.

### HasAnswerKey

`func (o *GetCardResponseObjectGetCardResponse) HasAnswerKey() bool`

HasAnswerKey returns a boolean if a field has been set.

### GetPublicAnswerKey

`func (o *GetCardResponseObjectGetCardResponse) GetPublicAnswerKey() string`

GetPublicAnswerKey returns the PublicAnswerKey field if non-nil, zero value otherwise.

### GetPublicAnswerKeyOk

`func (o *GetCardResponseObjectGetCardResponse) GetPublicAnswerKeyOk() (*string, bool)`

GetPublicAnswerKeyOk returns a tuple with the PublicAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAnswerKey

`func (o *GetCardResponseObjectGetCardResponse) SetPublicAnswerKey(v string)`

SetPublicAnswerKey sets PublicAnswerKey field to given value.

### HasPublicAnswerKey

`func (o *GetCardResponseObjectGetCardResponse) HasPublicAnswerKey() bool`

HasPublicAnswerKey returns a boolean if a field has been set.

### GetWasReversePrevious

`func (o *GetCardResponseObjectGetCardResponse) GetWasReversePrevious() int32`

GetWasReversePrevious returns the WasReversePrevious field if non-nil, zero value otherwise.

### GetWasReversePreviousOk

`func (o *GetCardResponseObjectGetCardResponse) GetWasReversePreviousOk() (*int32, bool)`

GetWasReversePreviousOk returns a tuple with the WasReversePrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasReversePrevious

`func (o *GetCardResponseObjectGetCardResponse) SetWasReversePrevious(v int32)`

SetWasReversePrevious sets WasReversePrevious field to given value.

### HasWasReversePrevious

`func (o *GetCardResponseObjectGetCardResponse) HasWasReversePrevious() bool`

HasWasReversePrevious returns a boolean if a field has been set.

### GetReversedAnswerKey

`func (o *GetCardResponseObjectGetCardResponse) GetReversedAnswerKey() string`

GetReversedAnswerKey returns the ReversedAnswerKey field if non-nil, zero value otherwise.

### GetReversedAnswerKeyOk

`func (o *GetCardResponseObjectGetCardResponse) GetReversedAnswerKeyOk() (*string, bool)`

GetReversedAnswerKeyOk returns a tuple with the ReversedAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedAnswerKey

`func (o *GetCardResponseObjectGetCardResponse) SetReversedAnswerKey(v string)`

SetReversedAnswerKey sets ReversedAnswerKey field to given value.

### HasReversedAnswerKey

`func (o *GetCardResponseObjectGetCardResponse) HasReversedAnswerKey() bool`

HasReversedAnswerKey returns a boolean if a field has been set.

### GetReversedSequence

`func (o *GetCardResponseObjectGetCardResponse) GetReversedSequence() string`

GetReversedSequence returns the ReversedSequence field if non-nil, zero value otherwise.

### GetReversedSequenceOk

`func (o *GetCardResponseObjectGetCardResponse) GetReversedSequenceOk() (*string, bool)`

GetReversedSequenceOk returns a tuple with the ReversedSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedSequence

`func (o *GetCardResponseObjectGetCardResponse) SetReversedSequence(v string)`

SetReversedSequence sets ReversedSequence field to given value.

### HasReversedSequence

`func (o *GetCardResponseObjectGetCardResponse) HasReversedSequence() bool`

HasReversedSequence returns a boolean if a field has been set.

### GetCommittedBlock

`func (o *GetCardResponseObjectGetCardResponse) GetCommittedBlock() string`

GetCommittedBlock returns the CommittedBlock field if non-nil, zero value otherwise.

### GetCommittedBlockOk

`func (o *GetCardResponseObjectGetCardResponse) GetCommittedBlockOk() (*string, bool)`

GetCommittedBlockOk returns a tuple with the CommittedBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedBlock

`func (o *GetCardResponseObjectGetCardResponse) SetCommittedBlock(v string)`

SetCommittedBlock sets CommittedBlock field to given value.

### HasCommittedBlock

`func (o *GetCardResponseObjectGetCardResponse) HasCommittedBlock() bool`

HasCommittedBlock returns a boolean if a field has been set.

### GetReversedBlock

`func (o *GetCardResponseObjectGetCardResponse) GetReversedBlock() string`

GetReversedBlock returns the ReversedBlock field if non-nil, zero value otherwise.

### GetReversedBlockOk

`func (o *GetCardResponseObjectGetCardResponse) GetReversedBlockOk() (*string, bool)`

GetReversedBlockOk returns a tuple with the ReversedBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedBlock

`func (o *GetCardResponseObjectGetCardResponse) SetReversedBlock(v string)`

SetReversedBlock sets ReversedBlock field to given value.

### HasReversedBlock

`func (o *GetCardResponseObjectGetCardResponse) HasReversedBlock() bool`

HasReversedBlock returns a boolean if a field has been set.

### GetMessageID

`func (o *GetCardResponseObjectGetCardResponse) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *GetCardResponseObjectGetCardResponse) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *GetCardResponseObjectGetCardResponse) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.

### HasMessageID

`func (o *GetCardResponseObjectGetCardResponse) HasMessageID() bool`

HasMessageID returns a boolean if a field has been set.

### GetServerAddress

`func (o *GetCardResponseObjectGetCardResponse) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *GetCardResponseObjectGetCardResponse) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *GetCardResponseObjectGetCardResponse) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *GetCardResponseObjectGetCardResponse) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.

### GetServerInstance

`func (o *GetCardResponseObjectGetCardResponse) GetServerInstance() string`

GetServerInstance returns the ServerInstance field if non-nil, zero value otherwise.

### GetServerInstanceOk

`func (o *GetCardResponseObjectGetCardResponse) GetServerInstanceOk() (*string, bool)`

GetServerInstanceOk returns a tuple with the ServerInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstance

`func (o *GetCardResponseObjectGetCardResponse) SetServerInstance(v string)`

SetServerInstance sets ServerInstance field to given value.

### HasServerInstance

`func (o *GetCardResponseObjectGetCardResponse) HasServerInstance() bool`

HasServerInstance returns a boolean if a field has been set.

### GetServerNodeName

`func (o *GetCardResponseObjectGetCardResponse) GetServerNodeName() string`

GetServerNodeName returns the ServerNodeName field if non-nil, zero value otherwise.

### GetServerNodeNameOk

`func (o *GetCardResponseObjectGetCardResponse) GetServerNodeNameOk() (*string, bool)`

GetServerNodeNameOk returns a tuple with the ServerNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNodeName

`func (o *GetCardResponseObjectGetCardResponse) SetServerNodeName(v string)`

SetServerNodeName sets ServerNodeName field to given value.

### HasServerNodeName

`func (o *GetCardResponseObjectGetCardResponse) HasServerNodeName() bool`

HasServerNodeName returns a boolean if a field has been set.

### GetAdapterInputVersion

`func (o *GetCardResponseObjectGetCardResponse) GetAdapterInputVersion() string`

GetAdapterInputVersion returns the AdapterInputVersion field if non-nil, zero value otherwise.

### GetAdapterInputVersionOk

`func (o *GetCardResponseObjectGetCardResponse) GetAdapterInputVersionOk() (*string, bool)`

GetAdapterInputVersionOk returns a tuple with the AdapterInputVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputVersion

`func (o *GetCardResponseObjectGetCardResponse) SetAdapterInputVersion(v string)`

SetAdapterInputVersion sets AdapterInputVersion field to given value.

### HasAdapterInputVersion

`func (o *GetCardResponseObjectGetCardResponse) HasAdapterInputVersion() bool`

HasAdapterInputVersion returns a boolean if a field has been set.

### GetAdapterInputAddress

`func (o *GetCardResponseObjectGetCardResponse) GetAdapterInputAddress() string`

GetAdapterInputAddress returns the AdapterInputAddress field if non-nil, zero value otherwise.

### GetAdapterInputAddressOk

`func (o *GetCardResponseObjectGetCardResponse) GetAdapterInputAddressOk() (*string, bool)`

GetAdapterInputAddressOk returns a tuple with the AdapterInputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputAddress

`func (o *GetCardResponseObjectGetCardResponse) SetAdapterInputAddress(v string)`

SetAdapterInputAddress sets AdapterInputAddress field to given value.

### HasAdapterInputAddress

`func (o *GetCardResponseObjectGetCardResponse) HasAdapterInputAddress() bool`

HasAdapterInputAddress returns a boolean if a field has been set.

### GetAdapterInputNodeName

`func (o *GetCardResponseObjectGetCardResponse) GetAdapterInputNodeName() string`

GetAdapterInputNodeName returns the AdapterInputNodeName field if non-nil, zero value otherwise.

### GetAdapterInputNodeNameOk

`func (o *GetCardResponseObjectGetCardResponse) GetAdapterInputNodeNameOk() (*string, bool)`

GetAdapterInputNodeNameOk returns a tuple with the AdapterInputNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputNodeName

`func (o *GetCardResponseObjectGetCardResponse) SetAdapterInputNodeName(v string)`

SetAdapterInputNodeName sets AdapterInputNodeName field to given value.

### HasAdapterInputNodeName

`func (o *GetCardResponseObjectGetCardResponse) HasAdapterInputNodeName() bool`

HasAdapterInputNodeName returns a boolean if a field has been set.

### GetAdapterOutputVersion

`func (o *GetCardResponseObjectGetCardResponse) GetAdapterOutputVersion() string`

GetAdapterOutputVersion returns the AdapterOutputVersion field if non-nil, zero value otherwise.

### GetAdapterOutputVersionOk

`func (o *GetCardResponseObjectGetCardResponse) GetAdapterOutputVersionOk() (*string, bool)`

GetAdapterOutputVersionOk returns a tuple with the AdapterOutputVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputVersion

`func (o *GetCardResponseObjectGetCardResponse) SetAdapterOutputVersion(v string)`

SetAdapterOutputVersion sets AdapterOutputVersion field to given value.

### HasAdapterOutputVersion

`func (o *GetCardResponseObjectGetCardResponse) HasAdapterOutputVersion() bool`

HasAdapterOutputVersion returns a boolean if a field has been set.

### GetAdapterOutputAddress

`func (o *GetCardResponseObjectGetCardResponse) GetAdapterOutputAddress() string`

GetAdapterOutputAddress returns the AdapterOutputAddress field if non-nil, zero value otherwise.

### GetAdapterOutputAddressOk

`func (o *GetCardResponseObjectGetCardResponse) GetAdapterOutputAddressOk() (*string, bool)`

GetAdapterOutputAddressOk returns a tuple with the AdapterOutputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputAddress

`func (o *GetCardResponseObjectGetCardResponse) SetAdapterOutputAddress(v string)`

SetAdapterOutputAddress sets AdapterOutputAddress field to given value.

### HasAdapterOutputAddress

`func (o *GetCardResponseObjectGetCardResponse) HasAdapterOutputAddress() bool`

HasAdapterOutputAddress returns a boolean if a field has been set.

### GetAdapterOutputNodeName

`func (o *GetCardResponseObjectGetCardResponse) GetAdapterOutputNodeName() string`

GetAdapterOutputNodeName returns the AdapterOutputNodeName field if non-nil, zero value otherwise.

### GetAdapterOutputNodeNameOk

`func (o *GetCardResponseObjectGetCardResponse) GetAdapterOutputNodeNameOk() (*string, bool)`

GetAdapterOutputNodeNameOk returns a tuple with the AdapterOutputNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputNodeName

`func (o *GetCardResponseObjectGetCardResponse) SetAdapterOutputNodeName(v string)`

SetAdapterOutputNodeName sets AdapterOutputNodeName field to given value.

### HasAdapterOutputNodeName

`func (o *GetCardResponseObjectGetCardResponse) HasAdapterOutputNodeName() bool`

HasAdapterOutputNodeName returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *GetCardResponseObjectGetCardResponse) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *GetCardResponseObjectGetCardResponse) GetAdditionalInformationOk() (*[]SaleResponseObjectSaleResponseAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *GetCardResponseObjectGetCardResponse) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *GetCardResponseObjectGetCardResponse) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetDisplayResponseMessage

`func (o *GetCardResponseObjectGetCardResponse) GetDisplayResponseMessage() []string`

GetDisplayResponseMessage returns the DisplayResponseMessage field if non-nil, zero value otherwise.

### GetDisplayResponseMessageOk

`func (o *GetCardResponseObjectGetCardResponse) GetDisplayResponseMessageOk() (*[]string, bool)`

GetDisplayResponseMessageOk returns a tuple with the DisplayResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayResponseMessage

`func (o *GetCardResponseObjectGetCardResponse) SetDisplayResponseMessage(v []string)`

SetDisplayResponseMessage sets DisplayResponseMessage field to given value.

### HasDisplayResponseMessage

`func (o *GetCardResponseObjectGetCardResponse) HasDisplayResponseMessage() bool`

HasDisplayResponseMessage returns a boolean if a field has been set.

### GetTransmitionDateTime

`func (o *GetCardResponseObjectGetCardResponse) GetTransmitionDateTime() time.Time`

GetTransmitionDateTime returns the TransmitionDateTime field if non-nil, zero value otherwise.

### GetTransmitionDateTimeOk

`func (o *GetCardResponseObjectGetCardResponse) GetTransmitionDateTimeOk() (*time.Time, bool)`

GetTransmitionDateTimeOk returns a tuple with the TransmitionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitionDateTime

`func (o *GetCardResponseObjectGetCardResponse) SetTransmitionDateTime(v time.Time)`

SetTransmitionDateTime sets TransmitionDateTime field to given value.

### HasTransmitionDateTime

`func (o *GetCardResponseObjectGetCardResponse) HasTransmitionDateTime() bool`

HasTransmitionDateTime returns a boolean if a field has been set.

### GetAuthResultCode

`func (o *GetCardResponseObjectGetCardResponse) GetAuthResultCode() string`

GetAuthResultCode returns the AuthResultCode field if non-nil, zero value otherwise.

### GetAuthResultCodeOk

`func (o *GetCardResponseObjectGetCardResponse) GetAuthResultCodeOk() (*string, bool)`

GetAuthResultCodeOk returns a tuple with the AuthResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResultCode

`func (o *GetCardResponseObjectGetCardResponse) SetAuthResultCode(v string)`

SetAuthResultCode sets AuthResultCode field to given value.

### HasAuthResultCode

`func (o *GetCardResponseObjectGetCardResponse) HasAuthResultCode() bool`

HasAuthResultCode returns a boolean if a field has been set.

### GetAuthHostProcessCode

`func (o *GetCardResponseObjectGetCardResponse) GetAuthHostProcessCode() string`

GetAuthHostProcessCode returns the AuthHostProcessCode field if non-nil, zero value otherwise.

### GetAuthHostProcessCodeOk

`func (o *GetCardResponseObjectGetCardResponse) GetAuthHostProcessCodeOk() (*string, bool)`

GetAuthHostProcessCodeOk returns a tuple with the AuthHostProcessCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHostProcessCode

`func (o *GetCardResponseObjectGetCardResponse) SetAuthHostProcessCode(v string)`

SetAuthHostProcessCode sets AuthHostProcessCode field to given value.

### HasAuthHostProcessCode

`func (o *GetCardResponseObjectGetCardResponse) HasAuthHostProcessCode() bool`

HasAuthHostProcessCode returns a boolean if a field has been set.

### GetAuthHostMsgType

`func (o *GetCardResponseObjectGetCardResponse) GetAuthHostMsgType() string`

GetAuthHostMsgType returns the AuthHostMsgType field if non-nil, zero value otherwise.

### GetAuthHostMsgTypeOk

`func (o *GetCardResponseObjectGetCardResponse) GetAuthHostMsgTypeOk() (*string, bool)`

GetAuthHostMsgTypeOk returns a tuple with the AuthHostMsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHostMsgType

`func (o *GetCardResponseObjectGetCardResponse) SetAuthHostMsgType(v string)`

SetAuthHostMsgType sets AuthHostMsgType field to given value.

### HasAuthHostMsgType

`func (o *GetCardResponseObjectGetCardResponse) HasAuthHostMsgType() bool`

HasAuthHostMsgType returns a boolean if a field has been set.

### GetHostMsgType

`func (o *GetCardResponseObjectGetCardResponse) GetHostMsgType() string`

GetHostMsgType returns the HostMsgType field if non-nil, zero value otherwise.

### GetHostMsgTypeOk

`func (o *GetCardResponseObjectGetCardResponse) GetHostMsgTypeOk() (*string, bool)`

GetHostMsgTypeOk returns a tuple with the HostMsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMsgType

`func (o *GetCardResponseObjectGetCardResponse) SetHostMsgType(v string)`

SetHostMsgType sets HostMsgType field to given value.

### HasHostMsgType

`func (o *GetCardResponseObjectGetCardResponse) HasHostMsgType() bool`

HasHostMsgType returns a boolean if a field has been set.

### GetAuthCode

`func (o *GetCardResponseObjectGetCardResponse) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *GetCardResponseObjectGetCardResponse) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *GetCardResponseObjectGetCardResponse) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *GetCardResponseObjectGetCardResponse) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetAuthDateTime

`func (o *GetCardResponseObjectGetCardResponse) GetAuthDateTime() time.Time`

GetAuthDateTime returns the AuthDateTime field if non-nil, zero value otherwise.

### GetAuthDateTimeOk

`func (o *GetCardResponseObjectGetCardResponse) GetAuthDateTimeOk() (*time.Time, bool)`

GetAuthDateTimeOk returns a tuple with the AuthDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDateTime

`func (o *GetCardResponseObjectGetCardResponse) SetAuthDateTime(v time.Time)`

SetAuthDateTime sets AuthDateTime field to given value.

### HasAuthDateTime

`func (o *GetCardResponseObjectGetCardResponse) HasAuthDateTime() bool`

HasAuthDateTime returns a boolean if a field has been set.

### GetAuthRRN

`func (o *GetCardResponseObjectGetCardResponse) GetAuthRRN() string`

GetAuthRRN returns the AuthRRN field if non-nil, zero value otherwise.

### GetAuthRRNOk

`func (o *GetCardResponseObjectGetCardResponse) GetAuthRRNOk() (*string, bool)`

GetAuthRRNOk returns a tuple with the AuthRRN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRRN

`func (o *GetCardResponseObjectGetCardResponse) SetAuthRRN(v string)`

SetAuthRRN sets AuthRRN field to given value.

### HasAuthRRN

`func (o *GetCardResponseObjectGetCardResponse) HasAuthRRN() bool`

HasAuthRRN returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *GetCardResponseObjectGetCardResponse) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *GetCardResponseObjectGetCardResponse) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *GetCardResponseObjectGetCardResponse) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *GetCardResponseObjectGetCardResponse) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetAuthHostMessage

`func (o *GetCardResponseObjectGetCardResponse) GetAuthHostMessage() string`

GetAuthHostMessage returns the AuthHostMessage field if non-nil, zero value otherwise.

### GetAuthHostMessageOk

`func (o *GetCardResponseObjectGetCardResponse) GetAuthHostMessageOk() (*string, bool)`

GetAuthHostMessageOk returns a tuple with the AuthHostMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHostMessage

`func (o *GetCardResponseObjectGetCardResponse) SetAuthHostMessage(v string)`

SetAuthHostMessage sets AuthHostMessage field to given value.

### HasAuthHostMessage

`func (o *GetCardResponseObjectGetCardResponse) HasAuthHostMessage() bool`

HasAuthHostMessage returns a boolean if a field has been set.

### GetAuthTicket

`func (o *GetCardResponseObjectGetCardResponse) GetAuthTicket() int32`

GetAuthTicket returns the AuthTicket field if non-nil, zero value otherwise.

### GetAuthTicketOk

`func (o *GetCardResponseObjectGetCardResponse) GetAuthTicketOk() (*int32, bool)`

GetAuthTicketOk returns a tuple with the AuthTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTicket

`func (o *GetCardResponseObjectGetCardResponse) SetAuthTicket(v int32)`

SetAuthTicket sets AuthTicket field to given value.

### HasAuthTicket

`func (o *GetCardResponseObjectGetCardResponse) HasAuthTicket() bool`

HasAuthTicket returns a boolean if a field has been set.

### GetCardReadMode

`func (o *GetCardResponseObjectGetCardResponse) GetCardReadMode() string`

GetCardReadMode returns the CardReadMode field if non-nil, zero value otherwise.

### GetCardReadModeOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardReadModeOk() (*string, bool)`

GetCardReadModeOk returns a tuple with the CardReadMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardReadMode

`func (o *GetCardResponseObjectGetCardResponse) SetCardReadMode(v string)`

SetCardReadMode sets CardReadMode field to given value.

### HasCardReadMode

`func (o *GetCardResponseObjectGetCardResponse) HasCardReadMode() bool`

HasCardReadMode returns a boolean if a field has been set.

### GetCardReadModeDescription

`func (o *GetCardResponseObjectGetCardResponse) GetCardReadModeDescription() string`

GetCardReadModeDescription returns the CardReadModeDescription field if non-nil, zero value otherwise.

### GetCardReadModeDescriptionOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardReadModeDescriptionOk() (*string, bool)`

GetCardReadModeDescriptionOk returns a tuple with the CardReadModeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardReadModeDescription

`func (o *GetCardResponseObjectGetCardResponse) SetCardReadModeDescription(v string)`

SetCardReadModeDescription sets CardReadModeDescription field to given value.

### HasCardReadModeDescription

`func (o *GetCardResponseObjectGetCardResponse) HasCardReadModeDescription() bool`

HasCardReadModeDescription returns a boolean if a field has been set.

### GetCardDescription

`func (o *GetCardResponseObjectGetCardResponse) GetCardDescription() string`

GetCardDescription returns the CardDescription field if non-nil, zero value otherwise.

### GetCardDescriptionOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardDescriptionOk() (*string, bool)`

GetCardDescriptionOk returns a tuple with the CardDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDescription

`func (o *GetCardResponseObjectGetCardResponse) SetCardDescription(v string)`

SetCardDescription sets CardDescription field to given value.

### HasCardDescription

`func (o *GetCardResponseObjectGetCardResponse) HasCardDescription() bool`

HasCardDescription returns a boolean if a field has been set.

### GetCardTypeDescription

`func (o *GetCardResponseObjectGetCardResponse) GetCardTypeDescription() string`

GetCardTypeDescription returns the CardTypeDescription field if non-nil, zero value otherwise.

### GetCardTypeDescriptionOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardTypeDescriptionOk() (*string, bool)`

GetCardTypeDescriptionOk returns a tuple with the CardTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardTypeDescription

`func (o *GetCardResponseObjectGetCardResponse) SetCardTypeDescription(v string)`

SetCardTypeDescription sets CardTypeDescription field to given value.

### HasCardTypeDescription

`func (o *GetCardResponseObjectGetCardResponse) HasCardTypeDescription() bool`

HasCardTypeDescription returns a boolean if a field has been set.

### GetCardCategory

`func (o *GetCardResponseObjectGetCardResponse) GetCardCategory() SaleResponseObjectSaleResponseCardCategory`

GetCardCategory returns the CardCategory field if non-nil, zero value otherwise.

### GetCardCategoryOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardCategoryOk() (*SaleResponseObjectSaleResponseCardCategory, bool)`

GetCardCategoryOk returns a tuple with the CardCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCategory

`func (o *GetCardResponseObjectGetCardResponse) SetCardCategory(v SaleResponseObjectSaleResponseCardCategory)`

SetCardCategory sets CardCategory field to given value.

### HasCardCategory

`func (o *GetCardResponseObjectGetCardResponse) HasCardCategory() bool`

HasCardCategory returns a boolean if a field has been set.

### GetCardNumber

`func (o *GetCardResponseObjectGetCardResponse) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *GetCardResponseObjectGetCardResponse) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *GetCardResponseObjectGetCardResponse) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNumberMasked

`func (o *GetCardResponseObjectGetCardResponse) GetCardNumberMasked() string`

GetCardNumberMasked returns the CardNumberMasked field if non-nil, zero value otherwise.

### GetCardNumberMaskedOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardNumberMaskedOk() (*string, bool)`

GetCardNumberMaskedOk returns a tuple with the CardNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberMasked

`func (o *GetCardResponseObjectGetCardResponse) SetCardNumberMasked(v string)`

SetCardNumberMasked sets CardNumberMasked field to given value.

### HasCardNumberMasked

`func (o *GetCardResponseObjectGetCardResponse) HasCardNumberMasked() bool`

HasCardNumberMasked returns a boolean if a field has been set.

### GetCardHashing

`func (o *GetCardResponseObjectGetCardResponse) GetCardHashing() string`

GetCardHashing returns the CardHashing field if non-nil, zero value otherwise.

### GetCardHashingOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardHashingOk() (*string, bool)`

GetCardHashingOk returns a tuple with the CardHashing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHashing

`func (o *GetCardResponseObjectGetCardResponse) SetCardHashing(v string)`

SetCardHashing sets CardHashing field to given value.

### HasCardHashing

`func (o *GetCardResponseObjectGetCardResponse) HasCardHashing() bool`

HasCardHashing returns a boolean if a field has been set.

### GetCardExp

`func (o *GetCardResponseObjectGetCardResponse) GetCardExp() string`

GetCardExp returns the CardExp field if non-nil, zero value otherwise.

### GetCardExpOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardExpOk() (*string, bool)`

GetCardExpOk returns a tuple with the CardExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExp

`func (o *GetCardResponseObjectGetCardResponse) SetCardExp(v string)`

SetCardExp sets CardExp field to given value.

### HasCardExp

`func (o *GetCardResponseObjectGetCardResponse) HasCardExp() bool`

HasCardExp returns a boolean if a field has been set.

### GetCardCryptogramResponse

`func (o *GetCardResponseObjectGetCardResponse) GetCardCryptogramResponse() string`

GetCardCryptogramResponse returns the CardCryptogramResponse field if non-nil, zero value otherwise.

### GetCardCryptogramResponseOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardCryptogramResponseOk() (*string, bool)`

GetCardCryptogramResponseOk returns a tuple with the CardCryptogramResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCryptogramResponse

`func (o *GetCardResponseObjectGetCardResponse) SetCardCryptogramResponse(v string)`

SetCardCryptogramResponse sets CardCryptogramResponse field to given value.

### HasCardCryptogramResponse

`func (o *GetCardResponseObjectGetCardResponse) HasCardCryptogramResponse() bool`

HasCardCryptogramResponse returns a boolean if a field has been set.

### GetCardAppName

`func (o *GetCardResponseObjectGetCardResponse) GetCardAppName() string`

GetCardAppName returns the CardAppName field if non-nil, zero value otherwise.

### GetCardAppNameOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardAppNameOk() (*string, bool)`

GetCardAppNameOk returns a tuple with the CardAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppName

`func (o *GetCardResponseObjectGetCardResponse) SetCardAppName(v string)`

SetCardAppName sets CardAppName field to given value.

### HasCardAppName

`func (o *GetCardResponseObjectGetCardResponse) HasCardAppName() bool`

HasCardAppName returns a boolean if a field has been set.

### GetCardAppIdentifier

`func (o *GetCardResponseObjectGetCardResponse) GetCardAppIdentifier() string`

GetCardAppIdentifier returns the CardAppIdentifier field if non-nil, zero value otherwise.

### GetCardAppIdentifierOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardAppIdentifierOk() (*string, bool)`

GetCardAppIdentifierOk returns a tuple with the CardAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppIdentifier

`func (o *GetCardResponseObjectGetCardResponse) SetCardAppIdentifier(v string)`

SetCardAppIdentifier sets CardAppIdentifier field to given value.

### HasCardAppIdentifier

`func (o *GetCardResponseObjectGetCardResponse) HasCardAppIdentifier() bool`

HasCardAppIdentifier returns a boolean if a field has been set.

### GetCardAppLabel

`func (o *GetCardResponseObjectGetCardResponse) GetCardAppLabel() string`

GetCardAppLabel returns the CardAppLabel field if non-nil, zero value otherwise.

### GetCardAppLabelOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardAppLabelOk() (*string, bool)`

GetCardAppLabelOk returns a tuple with the CardAppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppLabel

`func (o *GetCardResponseObjectGetCardResponse) SetCardAppLabel(v string)`

SetCardAppLabel sets CardAppLabel field to given value.

### HasCardAppLabel

`func (o *GetCardResponseObjectGetCardResponse) HasCardAppLabel() bool`

HasCardAppLabel returns a boolean if a field has been set.

### GetCardAuthRequestCryptogram

`func (o *GetCardResponseObjectGetCardResponse) GetCardAuthRequestCryptogram() string`

GetCardAuthRequestCryptogram returns the CardAuthRequestCryptogram field if non-nil, zero value otherwise.

### GetCardAuthRequestCryptogramOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardAuthRequestCryptogramOk() (*string, bool)`

GetCardAuthRequestCryptogramOk returns a tuple with the CardAuthRequestCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthRequestCryptogram

`func (o *GetCardResponseObjectGetCardResponse) SetCardAuthRequestCryptogram(v string)`

SetCardAuthRequestCryptogram sets CardAuthRequestCryptogram field to given value.

### HasCardAuthRequestCryptogram

`func (o *GetCardResponseObjectGetCardResponse) HasCardAuthRequestCryptogram() bool`

HasCardAuthRequestCryptogram returns a boolean if a field has been set.

### GetCardAuthResponseCryptogram

`func (o *GetCardResponseObjectGetCardResponse) GetCardAuthResponseCryptogram() string`

GetCardAuthResponseCryptogram returns the CardAuthResponseCryptogram field if non-nil, zero value otherwise.

### GetCardAuthResponseCryptogramOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardAuthResponseCryptogramOk() (*string, bool)`

GetCardAuthResponseCryptogramOk returns a tuple with the CardAuthResponseCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthResponseCryptogram

`func (o *GetCardResponseObjectGetCardResponse) SetCardAuthResponseCryptogram(v string)`

SetCardAuthResponseCryptogram sets CardAuthResponseCryptogram field to given value.

### HasCardAuthResponseCryptogram

`func (o *GetCardResponseObjectGetCardResponse) HasCardAuthResponseCryptogram() bool`

HasCardAuthResponseCryptogram returns a boolean if a field has been set.

### GetMerchantCategory

`func (o *GetCardResponseObjectGetCardResponse) GetMerchantCategory() SaleResponseObjectSaleResponseMerchantCategory`

GetMerchantCategory returns the MerchantCategory field if non-nil, zero value otherwise.

### GetMerchantCategoryOk

`func (o *GetCardResponseObjectGetCardResponse) GetMerchantCategoryOk() (*SaleResponseObjectSaleResponseMerchantCategory, bool)`

GetMerchantCategoryOk returns a tuple with the MerchantCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategory

`func (o *GetCardResponseObjectGetCardResponse) SetMerchantCategory(v SaleResponseObjectSaleResponseMerchantCategory)`

SetMerchantCategory sets MerchantCategory field to given value.

### HasMerchantCategory

`func (o *GetCardResponseObjectGetCardResponse) HasMerchantCategory() bool`

HasMerchantCategory returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *GetCardResponseObjectGetCardResponse) GetPaymentMethod() SaleResponseObjectSaleResponsePaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *GetCardResponseObjectGetCardResponse) GetPaymentMethodOk() (*SaleResponseObjectSaleResponsePaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *GetCardResponseObjectGetCardResponse) SetPaymentMethod(v SaleResponseObjectSaleResponsePaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *GetCardResponseObjectGetCardResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPlans

`func (o *GetCardResponseObjectGetCardResponse) GetPlans() SaleResponseObjectSaleResponsePlans`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *GetCardResponseObjectGetCardResponse) GetPlansOk() (*SaleResponseObjectSaleResponsePlans, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *GetCardResponseObjectGetCardResponse) SetPlans(v SaleResponseObjectSaleResponsePlans)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *GetCardResponseObjectGetCardResponse) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetCardResponseObjectGetCardResponse) GetConfiguration() SaleResponseObjectSaleResponseConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetCardResponseObjectGetCardResponse) GetConfigurationOk() (*SaleResponseObjectSaleResponseConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetCardResponseObjectGetCardResponse) SetConfiguration(v SaleResponseObjectSaleResponseConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetCardResponseObjectGetCardResponse) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTrack1

`func (o *GetCardResponseObjectGetCardResponse) GetTrack1() string`

GetTrack1 returns the Track1 field if non-nil, zero value otherwise.

### GetTrack1Ok

`func (o *GetCardResponseObjectGetCardResponse) GetTrack1Ok() (*string, bool)`

GetTrack1Ok returns a tuple with the Track1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack1

`func (o *GetCardResponseObjectGetCardResponse) SetTrack1(v string)`

SetTrack1 sets Track1 field to given value.

### HasTrack1

`func (o *GetCardResponseObjectGetCardResponse) HasTrack1() bool`

HasTrack1 returns a boolean if a field has been set.

### GetTrack2

`func (o *GetCardResponseObjectGetCardResponse) GetTrack2() string`

GetTrack2 returns the Track2 field if non-nil, zero value otherwise.

### GetTrack2Ok

`func (o *GetCardResponseObjectGetCardResponse) GetTrack2Ok() (*string, bool)`

GetTrack2Ok returns a tuple with the Track2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack2

`func (o *GetCardResponseObjectGetCardResponse) SetTrack2(v string)`

SetTrack2 sets Track2 field to given value.

### HasTrack2

`func (o *GetCardResponseObjectGetCardResponse) HasTrack2() bool`

HasTrack2 returns a boolean if a field has been set.

### GetSecurityCode

`func (o *GetCardResponseObjectGetCardResponse) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *GetCardResponseObjectGetCardResponse) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *GetCardResponseObjectGetCardResponse) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *GetCardResponseObjectGetCardResponse) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### GetCardCryptogram

`func (o *GetCardResponseObjectGetCardResponse) GetCardCryptogram() string`

GetCardCryptogram returns the CardCryptogram field if non-nil, zero value otherwise.

### GetCardCryptogramOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardCryptogramOk() (*string, bool)`

GetCardCryptogramOk returns a tuple with the CardCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCryptogram

`func (o *GetCardResponseObjectGetCardResponse) SetCardCryptogram(v string)`

SetCardCryptogram sets CardCryptogram field to given value.

### HasCardCryptogram

`func (o *GetCardResponseObjectGetCardResponse) HasCardCryptogram() bool`

HasCardCryptogram returns a boolean if a field has been set.

### GetSettlementBatchNumber

`func (o *GetCardResponseObjectGetCardResponse) GetSettlementBatchNumber() int32`

GetSettlementBatchNumber returns the SettlementBatchNumber field if non-nil, zero value otherwise.

### GetSettlementBatchNumberOk

`func (o *GetCardResponseObjectGetCardResponse) GetSettlementBatchNumberOk() (*int32, bool)`

GetSettlementBatchNumberOk returns a tuple with the SettlementBatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementBatchNumber

`func (o *GetCardResponseObjectGetCardResponse) SetSettlementBatchNumber(v int32)`

SetSettlementBatchNumber sets SettlementBatchNumber field to given value.

### HasSettlementBatchNumber

`func (o *GetCardResponseObjectGetCardResponse) HasSettlementBatchNumber() bool`

HasSettlementBatchNumber returns a boolean if a field has been set.

### GetTerminalID

`func (o *GetCardResponseObjectGetCardResponse) GetTerminalID() string`

GetTerminalID returns the TerminalID field if non-nil, zero value otherwise.

### GetTerminalIDOk

`func (o *GetCardResponseObjectGetCardResponse) GetTerminalIDOk() (*string, bool)`

GetTerminalIDOk returns a tuple with the TerminalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalID

`func (o *GetCardResponseObjectGetCardResponse) SetTerminalID(v string)`

SetTerminalID sets TerminalID field to given value.

### HasTerminalID

`func (o *GetCardResponseObjectGetCardResponse) HasTerminalID() bool`

HasTerminalID returns a boolean if a field has been set.

### GetTerminalTrace

`func (o *GetCardResponseObjectGetCardResponse) GetTerminalTrace() int32`

GetTerminalTrace returns the TerminalTrace field if non-nil, zero value otherwise.

### GetTerminalTraceOk

`func (o *GetCardResponseObjectGetCardResponse) GetTerminalTraceOk() (*int32, bool)`

GetTerminalTraceOk returns a tuple with the TerminalTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalTrace

`func (o *GetCardResponseObjectGetCardResponse) SetTerminalTrace(v int32)`

SetTerminalTrace sets TerminalTrace field to given value.

### HasTerminalTrace

`func (o *GetCardResponseObjectGetCardResponse) HasTerminalTrace() bool`

HasTerminalTrace returns a boolean if a field has been set.

### GetPaymentFacilitatorID

`func (o *GetCardResponseObjectGetCardResponse) GetPaymentFacilitatorID() string`

GetPaymentFacilitatorID returns the PaymentFacilitatorID field if non-nil, zero value otherwise.

### GetPaymentFacilitatorIDOk

`func (o *GetCardResponseObjectGetCardResponse) GetPaymentFacilitatorIDOk() (*string, bool)`

GetPaymentFacilitatorIDOk returns a tuple with the PaymentFacilitatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFacilitatorID

`func (o *GetCardResponseObjectGetCardResponse) SetPaymentFacilitatorID(v string)`

SetPaymentFacilitatorID sets PaymentFacilitatorID field to given value.

### HasPaymentFacilitatorID

`func (o *GetCardResponseObjectGetCardResponse) HasPaymentFacilitatorID() bool`

HasPaymentFacilitatorID returns a boolean if a field has been set.

### GetMerchantID

`func (o *GetCardResponseObjectGetCardResponse) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *GetCardResponseObjectGetCardResponse) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *GetCardResponseObjectGetCardResponse) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *GetCardResponseObjectGetCardResponse) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetCardHolderName

`func (o *GetCardResponseObjectGetCardResponse) GetCardHolderName() string`

GetCardHolderName returns the CardHolderName field if non-nil, zero value otherwise.

### GetCardHolderNameOk

`func (o *GetCardResponseObjectGetCardResponse) GetCardHolderNameOk() (*string, bool)`

GetCardHolderNameOk returns a tuple with the CardHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderName

`func (o *GetCardResponseObjectGetCardResponse) SetCardHolderName(v string)`

SetCardHolderName sets CardHolderName field to given value.

### HasCardHolderName

`func (o *GetCardResponseObjectGetCardResponse) HasCardHolderName() bool`

HasCardHolderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


