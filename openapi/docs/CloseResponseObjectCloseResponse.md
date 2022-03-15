# CloseResponseObjectCloseResponse

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

## Methods

### NewCloseResponseObjectCloseResponse

`func NewCloseResponseObjectCloseResponse(responseCode int32, responseActions []string, responseMessage string, ) *CloseResponseObjectCloseResponse`

NewCloseResponseObjectCloseResponse instantiates a new CloseResponseObjectCloseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloseResponseObjectCloseResponseWithDefaults

`func NewCloseResponseObjectCloseResponseWithDefaults() *CloseResponseObjectCloseResponse`

NewCloseResponseObjectCloseResponseWithDefaults instantiates a new CloseResponseObjectCloseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *CloseResponseObjectCloseResponse) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *CloseResponseObjectCloseResponse) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *CloseResponseObjectCloseResponse) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseActions

`func (o *CloseResponseObjectCloseResponse) GetResponseActions() []string`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *CloseResponseObjectCloseResponse) GetResponseActionsOk() (*[]string, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *CloseResponseObjectCloseResponse) SetResponseActions(v []string)`

SetResponseActions sets ResponseActions field to given value.


### GetResponseMessage

`func (o *CloseResponseObjectCloseResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *CloseResponseObjectCloseResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *CloseResponseObjectCloseResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetCompanyIdentification

`func (o *CloseResponseObjectCloseResponse) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *CloseResponseObjectCloseResponse) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *CloseResponseObjectCloseResponse) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *CloseResponseObjectCloseResponse) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetSystemIdentification

`func (o *CloseResponseObjectCloseResponse) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *CloseResponseObjectCloseResponse) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *CloseResponseObjectCloseResponse) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.

### HasSystemIdentification

`func (o *CloseResponseObjectCloseResponse) HasSystemIdentification() bool`

HasSystemIdentification returns a boolean if a field has been set.

### GetBranchIdentification

`func (o *CloseResponseObjectCloseResponse) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *CloseResponseObjectCloseResponse) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *CloseResponseObjectCloseResponse) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *CloseResponseObjectCloseResponse) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *CloseResponseObjectCloseResponse) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *CloseResponseObjectCloseResponse) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *CloseResponseObjectCloseResponse) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *CloseResponseObjectCloseResponse) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetForeignResponseCode

`func (o *CloseResponseObjectCloseResponse) GetForeignResponseCode() string`

GetForeignResponseCode returns the ForeignResponseCode field if non-nil, zero value otherwise.

### GetForeignResponseCodeOk

`func (o *CloseResponseObjectCloseResponse) GetForeignResponseCodeOk() (*string, bool)`

GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignResponseCode

`func (o *CloseResponseObjectCloseResponse) SetForeignResponseCode(v string)`

SetForeignResponseCode sets ForeignResponseCode field to given value.

### HasForeignResponseCode

`func (o *CloseResponseObjectCloseResponse) HasForeignResponseCode() bool`

HasForeignResponseCode returns a boolean if a field has been set.

### GetRequestType

`func (o *CloseResponseObjectCloseResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *CloseResponseObjectCloseResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *CloseResponseObjectCloseResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *CloseResponseObjectCloseResponse) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetServiceVersion

`func (o *CloseResponseObjectCloseResponse) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *CloseResponseObjectCloseResponse) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *CloseResponseObjectCloseResponse) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *CloseResponseObjectCloseResponse) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *CloseResponseObjectCloseResponse) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *CloseResponseObjectCloseResponse) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *CloseResponseObjectCloseResponse) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *CloseResponseObjectCloseResponse) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *CloseResponseObjectCloseResponse) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *CloseResponseObjectCloseResponse) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *CloseResponseObjectCloseResponse) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *CloseResponseObjectCloseResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetBlock

`func (o *CloseResponseObjectCloseResponse) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *CloseResponseObjectCloseResponse) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *CloseResponseObjectCloseResponse) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *CloseResponseObjectCloseResponse) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetRequestKey

`func (o *CloseResponseObjectCloseResponse) GetRequestKey() string`

GetRequestKey returns the RequestKey field if non-nil, zero value otherwise.

### GetRequestKeyOk

`func (o *CloseResponseObjectCloseResponse) GetRequestKeyOk() (*string, bool)`

GetRequestKeyOk returns a tuple with the RequestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestKey

`func (o *CloseResponseObjectCloseResponse) SetRequestKey(v string)`

SetRequestKey sets RequestKey field to given value.

### HasRequestKey

`func (o *CloseResponseObjectCloseResponse) HasRequestKey() bool`

HasRequestKey returns a boolean if a field has been set.

### GetWorkingKeys

`func (o *CloseResponseObjectCloseResponse) GetWorkingKeys() []SaleResponseObjectSaleResponseWorkingKeys`

GetWorkingKeys returns the WorkingKeys field if non-nil, zero value otherwise.

### GetWorkingKeysOk

`func (o *CloseResponseObjectCloseResponse) GetWorkingKeysOk() (*[]SaleResponseObjectSaleResponseWorkingKeys, bool)`

GetWorkingKeysOk returns a tuple with the WorkingKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingKeys

`func (o *CloseResponseObjectCloseResponse) SetWorkingKeys(v []SaleResponseObjectSaleResponseWorkingKeys)`

SetWorkingKeys sets WorkingKeys field to given value.

### HasWorkingKeys

`func (o *CloseResponseObjectCloseResponse) HasWorkingKeys() bool`

HasWorkingKeys returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *CloseResponseObjectCloseResponse) GetRequiredInformation() []SaleResponseObjectSaleResponseRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *CloseResponseObjectCloseResponse) GetRequiredInformationOk() (*[]SaleResponseObjectSaleResponseRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *CloseResponseObjectCloseResponse) SetRequiredInformation(v []SaleResponseObjectSaleResponseRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *CloseResponseObjectCloseResponse) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetAnswerType

`func (o *CloseResponseObjectCloseResponse) GetAnswerType() string`

GetAnswerType returns the AnswerType field if non-nil, zero value otherwise.

### GetAnswerTypeOk

`func (o *CloseResponseObjectCloseResponse) GetAnswerTypeOk() (*string, bool)`

GetAnswerTypeOk returns a tuple with the AnswerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerType

`func (o *CloseResponseObjectCloseResponse) SetAnswerType(v string)`

SetAnswerType sets AnswerType field to given value.

### HasAnswerType

`func (o *CloseResponseObjectCloseResponse) HasAnswerType() bool`

HasAnswerType returns a boolean if a field has been set.

### GetAnswerKey

`func (o *CloseResponseObjectCloseResponse) GetAnswerKey() string`

GetAnswerKey returns the AnswerKey field if non-nil, zero value otherwise.

### GetAnswerKeyOk

`func (o *CloseResponseObjectCloseResponse) GetAnswerKeyOk() (*string, bool)`

GetAnswerKeyOk returns a tuple with the AnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerKey

`func (o *CloseResponseObjectCloseResponse) SetAnswerKey(v string)`

SetAnswerKey sets AnswerKey field to given value.

### HasAnswerKey

`func (o *CloseResponseObjectCloseResponse) HasAnswerKey() bool`

HasAnswerKey returns a boolean if a field has been set.

### GetPublicAnswerKey

`func (o *CloseResponseObjectCloseResponse) GetPublicAnswerKey() string`

GetPublicAnswerKey returns the PublicAnswerKey field if non-nil, zero value otherwise.

### GetPublicAnswerKeyOk

`func (o *CloseResponseObjectCloseResponse) GetPublicAnswerKeyOk() (*string, bool)`

GetPublicAnswerKeyOk returns a tuple with the PublicAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAnswerKey

`func (o *CloseResponseObjectCloseResponse) SetPublicAnswerKey(v string)`

SetPublicAnswerKey sets PublicAnswerKey field to given value.

### HasPublicAnswerKey

`func (o *CloseResponseObjectCloseResponse) HasPublicAnswerKey() bool`

HasPublicAnswerKey returns a boolean if a field has been set.

### GetWasReversePrevious

`func (o *CloseResponseObjectCloseResponse) GetWasReversePrevious() int32`

GetWasReversePrevious returns the WasReversePrevious field if non-nil, zero value otherwise.

### GetWasReversePreviousOk

`func (o *CloseResponseObjectCloseResponse) GetWasReversePreviousOk() (*int32, bool)`

GetWasReversePreviousOk returns a tuple with the WasReversePrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasReversePrevious

`func (o *CloseResponseObjectCloseResponse) SetWasReversePrevious(v int32)`

SetWasReversePrevious sets WasReversePrevious field to given value.

### HasWasReversePrevious

`func (o *CloseResponseObjectCloseResponse) HasWasReversePrevious() bool`

HasWasReversePrevious returns a boolean if a field has been set.

### GetReversedAnswerKey

`func (o *CloseResponseObjectCloseResponse) GetReversedAnswerKey() string`

GetReversedAnswerKey returns the ReversedAnswerKey field if non-nil, zero value otherwise.

### GetReversedAnswerKeyOk

`func (o *CloseResponseObjectCloseResponse) GetReversedAnswerKeyOk() (*string, bool)`

GetReversedAnswerKeyOk returns a tuple with the ReversedAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedAnswerKey

`func (o *CloseResponseObjectCloseResponse) SetReversedAnswerKey(v string)`

SetReversedAnswerKey sets ReversedAnswerKey field to given value.

### HasReversedAnswerKey

`func (o *CloseResponseObjectCloseResponse) HasReversedAnswerKey() bool`

HasReversedAnswerKey returns a boolean if a field has been set.

### GetReversedSequence

`func (o *CloseResponseObjectCloseResponse) GetReversedSequence() string`

GetReversedSequence returns the ReversedSequence field if non-nil, zero value otherwise.

### GetReversedSequenceOk

`func (o *CloseResponseObjectCloseResponse) GetReversedSequenceOk() (*string, bool)`

GetReversedSequenceOk returns a tuple with the ReversedSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedSequence

`func (o *CloseResponseObjectCloseResponse) SetReversedSequence(v string)`

SetReversedSequence sets ReversedSequence field to given value.

### HasReversedSequence

`func (o *CloseResponseObjectCloseResponse) HasReversedSequence() bool`

HasReversedSequence returns a boolean if a field has been set.

### GetCommittedBlock

`func (o *CloseResponseObjectCloseResponse) GetCommittedBlock() string`

GetCommittedBlock returns the CommittedBlock field if non-nil, zero value otherwise.

### GetCommittedBlockOk

`func (o *CloseResponseObjectCloseResponse) GetCommittedBlockOk() (*string, bool)`

GetCommittedBlockOk returns a tuple with the CommittedBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedBlock

`func (o *CloseResponseObjectCloseResponse) SetCommittedBlock(v string)`

SetCommittedBlock sets CommittedBlock field to given value.

### HasCommittedBlock

`func (o *CloseResponseObjectCloseResponse) HasCommittedBlock() bool`

HasCommittedBlock returns a boolean if a field has been set.

### GetReversedBlock

`func (o *CloseResponseObjectCloseResponse) GetReversedBlock() string`

GetReversedBlock returns the ReversedBlock field if non-nil, zero value otherwise.

### GetReversedBlockOk

`func (o *CloseResponseObjectCloseResponse) GetReversedBlockOk() (*string, bool)`

GetReversedBlockOk returns a tuple with the ReversedBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedBlock

`func (o *CloseResponseObjectCloseResponse) SetReversedBlock(v string)`

SetReversedBlock sets ReversedBlock field to given value.

### HasReversedBlock

`func (o *CloseResponseObjectCloseResponse) HasReversedBlock() bool`

HasReversedBlock returns a boolean if a field has been set.

### GetMessageID

`func (o *CloseResponseObjectCloseResponse) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *CloseResponseObjectCloseResponse) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *CloseResponseObjectCloseResponse) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.

### HasMessageID

`func (o *CloseResponseObjectCloseResponse) HasMessageID() bool`

HasMessageID returns a boolean if a field has been set.

### GetServerAddress

`func (o *CloseResponseObjectCloseResponse) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *CloseResponseObjectCloseResponse) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *CloseResponseObjectCloseResponse) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *CloseResponseObjectCloseResponse) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.

### GetServerInstance

`func (o *CloseResponseObjectCloseResponse) GetServerInstance() string`

GetServerInstance returns the ServerInstance field if non-nil, zero value otherwise.

### GetServerInstanceOk

`func (o *CloseResponseObjectCloseResponse) GetServerInstanceOk() (*string, bool)`

GetServerInstanceOk returns a tuple with the ServerInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstance

`func (o *CloseResponseObjectCloseResponse) SetServerInstance(v string)`

SetServerInstance sets ServerInstance field to given value.

### HasServerInstance

`func (o *CloseResponseObjectCloseResponse) HasServerInstance() bool`

HasServerInstance returns a boolean if a field has been set.

### GetServerNodeName

`func (o *CloseResponseObjectCloseResponse) GetServerNodeName() string`

GetServerNodeName returns the ServerNodeName field if non-nil, zero value otherwise.

### GetServerNodeNameOk

`func (o *CloseResponseObjectCloseResponse) GetServerNodeNameOk() (*string, bool)`

GetServerNodeNameOk returns a tuple with the ServerNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNodeName

`func (o *CloseResponseObjectCloseResponse) SetServerNodeName(v string)`

SetServerNodeName sets ServerNodeName field to given value.

### HasServerNodeName

`func (o *CloseResponseObjectCloseResponse) HasServerNodeName() bool`

HasServerNodeName returns a boolean if a field has been set.

### GetAdapterInputVersion

`func (o *CloseResponseObjectCloseResponse) GetAdapterInputVersion() string`

GetAdapterInputVersion returns the AdapterInputVersion field if non-nil, zero value otherwise.

### GetAdapterInputVersionOk

`func (o *CloseResponseObjectCloseResponse) GetAdapterInputVersionOk() (*string, bool)`

GetAdapterInputVersionOk returns a tuple with the AdapterInputVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputVersion

`func (o *CloseResponseObjectCloseResponse) SetAdapterInputVersion(v string)`

SetAdapterInputVersion sets AdapterInputVersion field to given value.

### HasAdapterInputVersion

`func (o *CloseResponseObjectCloseResponse) HasAdapterInputVersion() bool`

HasAdapterInputVersion returns a boolean if a field has been set.

### GetAdapterInputAddress

`func (o *CloseResponseObjectCloseResponse) GetAdapterInputAddress() string`

GetAdapterInputAddress returns the AdapterInputAddress field if non-nil, zero value otherwise.

### GetAdapterInputAddressOk

`func (o *CloseResponseObjectCloseResponse) GetAdapterInputAddressOk() (*string, bool)`

GetAdapterInputAddressOk returns a tuple with the AdapterInputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputAddress

`func (o *CloseResponseObjectCloseResponse) SetAdapterInputAddress(v string)`

SetAdapterInputAddress sets AdapterInputAddress field to given value.

### HasAdapterInputAddress

`func (o *CloseResponseObjectCloseResponse) HasAdapterInputAddress() bool`

HasAdapterInputAddress returns a boolean if a field has been set.

### GetAdapterInputNodeName

`func (o *CloseResponseObjectCloseResponse) GetAdapterInputNodeName() string`

GetAdapterInputNodeName returns the AdapterInputNodeName field if non-nil, zero value otherwise.

### GetAdapterInputNodeNameOk

`func (o *CloseResponseObjectCloseResponse) GetAdapterInputNodeNameOk() (*string, bool)`

GetAdapterInputNodeNameOk returns a tuple with the AdapterInputNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputNodeName

`func (o *CloseResponseObjectCloseResponse) SetAdapterInputNodeName(v string)`

SetAdapterInputNodeName sets AdapterInputNodeName field to given value.

### HasAdapterInputNodeName

`func (o *CloseResponseObjectCloseResponse) HasAdapterInputNodeName() bool`

HasAdapterInputNodeName returns a boolean if a field has been set.

### GetAdapterOutputVersion

`func (o *CloseResponseObjectCloseResponse) GetAdapterOutputVersion() string`

GetAdapterOutputVersion returns the AdapterOutputVersion field if non-nil, zero value otherwise.

### GetAdapterOutputVersionOk

`func (o *CloseResponseObjectCloseResponse) GetAdapterOutputVersionOk() (*string, bool)`

GetAdapterOutputVersionOk returns a tuple with the AdapterOutputVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputVersion

`func (o *CloseResponseObjectCloseResponse) SetAdapterOutputVersion(v string)`

SetAdapterOutputVersion sets AdapterOutputVersion field to given value.

### HasAdapterOutputVersion

`func (o *CloseResponseObjectCloseResponse) HasAdapterOutputVersion() bool`

HasAdapterOutputVersion returns a boolean if a field has been set.

### GetAdapterOutputAddress

`func (o *CloseResponseObjectCloseResponse) GetAdapterOutputAddress() string`

GetAdapterOutputAddress returns the AdapterOutputAddress field if non-nil, zero value otherwise.

### GetAdapterOutputAddressOk

`func (o *CloseResponseObjectCloseResponse) GetAdapterOutputAddressOk() (*string, bool)`

GetAdapterOutputAddressOk returns a tuple with the AdapterOutputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputAddress

`func (o *CloseResponseObjectCloseResponse) SetAdapterOutputAddress(v string)`

SetAdapterOutputAddress sets AdapterOutputAddress field to given value.

### HasAdapterOutputAddress

`func (o *CloseResponseObjectCloseResponse) HasAdapterOutputAddress() bool`

HasAdapterOutputAddress returns a boolean if a field has been set.

### GetAdapterOutputNodeName

`func (o *CloseResponseObjectCloseResponse) GetAdapterOutputNodeName() string`

GetAdapterOutputNodeName returns the AdapterOutputNodeName field if non-nil, zero value otherwise.

### GetAdapterOutputNodeNameOk

`func (o *CloseResponseObjectCloseResponse) GetAdapterOutputNodeNameOk() (*string, bool)`

GetAdapterOutputNodeNameOk returns a tuple with the AdapterOutputNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputNodeName

`func (o *CloseResponseObjectCloseResponse) SetAdapterOutputNodeName(v string)`

SetAdapterOutputNodeName sets AdapterOutputNodeName field to given value.

### HasAdapterOutputNodeName

`func (o *CloseResponseObjectCloseResponse) HasAdapterOutputNodeName() bool`

HasAdapterOutputNodeName returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *CloseResponseObjectCloseResponse) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *CloseResponseObjectCloseResponse) GetAdditionalInformationOk() (*[]SaleResponseObjectSaleResponseAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *CloseResponseObjectCloseResponse) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *CloseResponseObjectCloseResponse) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetDisplayResponseMessage

`func (o *CloseResponseObjectCloseResponse) GetDisplayResponseMessage() []string`

GetDisplayResponseMessage returns the DisplayResponseMessage field if non-nil, zero value otherwise.

### GetDisplayResponseMessageOk

`func (o *CloseResponseObjectCloseResponse) GetDisplayResponseMessageOk() (*[]string, bool)`

GetDisplayResponseMessageOk returns a tuple with the DisplayResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayResponseMessage

`func (o *CloseResponseObjectCloseResponse) SetDisplayResponseMessage(v []string)`

SetDisplayResponseMessage sets DisplayResponseMessage field to given value.

### HasDisplayResponseMessage

`func (o *CloseResponseObjectCloseResponse) HasDisplayResponseMessage() bool`

HasDisplayResponseMessage returns a boolean if a field has been set.

### GetTransmitionDateTime

`func (o *CloseResponseObjectCloseResponse) GetTransmitionDateTime() time.Time`

GetTransmitionDateTime returns the TransmitionDateTime field if non-nil, zero value otherwise.

### GetTransmitionDateTimeOk

`func (o *CloseResponseObjectCloseResponse) GetTransmitionDateTimeOk() (*time.Time, bool)`

GetTransmitionDateTimeOk returns a tuple with the TransmitionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitionDateTime

`func (o *CloseResponseObjectCloseResponse) SetTransmitionDateTime(v time.Time)`

SetTransmitionDateTime sets TransmitionDateTime field to given value.

### HasTransmitionDateTime

`func (o *CloseResponseObjectCloseResponse) HasTransmitionDateTime() bool`

HasTransmitionDateTime returns a boolean if a field has been set.

### GetAuthResultCode

`func (o *CloseResponseObjectCloseResponse) GetAuthResultCode() string`

GetAuthResultCode returns the AuthResultCode field if non-nil, zero value otherwise.

### GetAuthResultCodeOk

`func (o *CloseResponseObjectCloseResponse) GetAuthResultCodeOk() (*string, bool)`

GetAuthResultCodeOk returns a tuple with the AuthResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResultCode

`func (o *CloseResponseObjectCloseResponse) SetAuthResultCode(v string)`

SetAuthResultCode sets AuthResultCode field to given value.

### HasAuthResultCode

`func (o *CloseResponseObjectCloseResponse) HasAuthResultCode() bool`

HasAuthResultCode returns a boolean if a field has been set.

### GetAuthHostProcessCode

`func (o *CloseResponseObjectCloseResponse) GetAuthHostProcessCode() string`

GetAuthHostProcessCode returns the AuthHostProcessCode field if non-nil, zero value otherwise.

### GetAuthHostProcessCodeOk

`func (o *CloseResponseObjectCloseResponse) GetAuthHostProcessCodeOk() (*string, bool)`

GetAuthHostProcessCodeOk returns a tuple with the AuthHostProcessCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHostProcessCode

`func (o *CloseResponseObjectCloseResponse) SetAuthHostProcessCode(v string)`

SetAuthHostProcessCode sets AuthHostProcessCode field to given value.

### HasAuthHostProcessCode

`func (o *CloseResponseObjectCloseResponse) HasAuthHostProcessCode() bool`

HasAuthHostProcessCode returns a boolean if a field has been set.

### GetAuthHostMsgType

`func (o *CloseResponseObjectCloseResponse) GetAuthHostMsgType() string`

GetAuthHostMsgType returns the AuthHostMsgType field if non-nil, zero value otherwise.

### GetAuthHostMsgTypeOk

`func (o *CloseResponseObjectCloseResponse) GetAuthHostMsgTypeOk() (*string, bool)`

GetAuthHostMsgTypeOk returns a tuple with the AuthHostMsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHostMsgType

`func (o *CloseResponseObjectCloseResponse) SetAuthHostMsgType(v string)`

SetAuthHostMsgType sets AuthHostMsgType field to given value.

### HasAuthHostMsgType

`func (o *CloseResponseObjectCloseResponse) HasAuthHostMsgType() bool`

HasAuthHostMsgType returns a boolean if a field has been set.

### GetHostMsgType

`func (o *CloseResponseObjectCloseResponse) GetHostMsgType() string`

GetHostMsgType returns the HostMsgType field if non-nil, zero value otherwise.

### GetHostMsgTypeOk

`func (o *CloseResponseObjectCloseResponse) GetHostMsgTypeOk() (*string, bool)`

GetHostMsgTypeOk returns a tuple with the HostMsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMsgType

`func (o *CloseResponseObjectCloseResponse) SetHostMsgType(v string)`

SetHostMsgType sets HostMsgType field to given value.

### HasHostMsgType

`func (o *CloseResponseObjectCloseResponse) HasHostMsgType() bool`

HasHostMsgType returns a boolean if a field has been set.

### GetAuthCode

`func (o *CloseResponseObjectCloseResponse) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *CloseResponseObjectCloseResponse) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *CloseResponseObjectCloseResponse) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *CloseResponseObjectCloseResponse) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetAuthDateTime

`func (o *CloseResponseObjectCloseResponse) GetAuthDateTime() time.Time`

GetAuthDateTime returns the AuthDateTime field if non-nil, zero value otherwise.

### GetAuthDateTimeOk

`func (o *CloseResponseObjectCloseResponse) GetAuthDateTimeOk() (*time.Time, bool)`

GetAuthDateTimeOk returns a tuple with the AuthDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDateTime

`func (o *CloseResponseObjectCloseResponse) SetAuthDateTime(v time.Time)`

SetAuthDateTime sets AuthDateTime field to given value.

### HasAuthDateTime

`func (o *CloseResponseObjectCloseResponse) HasAuthDateTime() bool`

HasAuthDateTime returns a boolean if a field has been set.

### GetAuthRRN

`func (o *CloseResponseObjectCloseResponse) GetAuthRRN() string`

GetAuthRRN returns the AuthRRN field if non-nil, zero value otherwise.

### GetAuthRRNOk

`func (o *CloseResponseObjectCloseResponse) GetAuthRRNOk() (*string, bool)`

GetAuthRRNOk returns a tuple with the AuthRRN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRRN

`func (o *CloseResponseObjectCloseResponse) SetAuthRRN(v string)`

SetAuthRRN sets AuthRRN field to given value.

### HasAuthRRN

`func (o *CloseResponseObjectCloseResponse) HasAuthRRN() bool`

HasAuthRRN returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *CloseResponseObjectCloseResponse) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *CloseResponseObjectCloseResponse) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *CloseResponseObjectCloseResponse) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *CloseResponseObjectCloseResponse) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetAuthHostMessage

`func (o *CloseResponseObjectCloseResponse) GetAuthHostMessage() string`

GetAuthHostMessage returns the AuthHostMessage field if non-nil, zero value otherwise.

### GetAuthHostMessageOk

`func (o *CloseResponseObjectCloseResponse) GetAuthHostMessageOk() (*string, bool)`

GetAuthHostMessageOk returns a tuple with the AuthHostMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHostMessage

`func (o *CloseResponseObjectCloseResponse) SetAuthHostMessage(v string)`

SetAuthHostMessage sets AuthHostMessage field to given value.

### HasAuthHostMessage

`func (o *CloseResponseObjectCloseResponse) HasAuthHostMessage() bool`

HasAuthHostMessage returns a boolean if a field has been set.

### GetAuthTicket

`func (o *CloseResponseObjectCloseResponse) GetAuthTicket() int32`

GetAuthTicket returns the AuthTicket field if non-nil, zero value otherwise.

### GetAuthTicketOk

`func (o *CloseResponseObjectCloseResponse) GetAuthTicketOk() (*int32, bool)`

GetAuthTicketOk returns a tuple with the AuthTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTicket

`func (o *CloseResponseObjectCloseResponse) SetAuthTicket(v int32)`

SetAuthTicket sets AuthTicket field to given value.

### HasAuthTicket

`func (o *CloseResponseObjectCloseResponse) HasAuthTicket() bool`

HasAuthTicket returns a boolean if a field has been set.

### GetCardReadMode

`func (o *CloseResponseObjectCloseResponse) GetCardReadMode() string`

GetCardReadMode returns the CardReadMode field if non-nil, zero value otherwise.

### GetCardReadModeOk

`func (o *CloseResponseObjectCloseResponse) GetCardReadModeOk() (*string, bool)`

GetCardReadModeOk returns a tuple with the CardReadMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardReadMode

`func (o *CloseResponseObjectCloseResponse) SetCardReadMode(v string)`

SetCardReadMode sets CardReadMode field to given value.

### HasCardReadMode

`func (o *CloseResponseObjectCloseResponse) HasCardReadMode() bool`

HasCardReadMode returns a boolean if a field has been set.

### GetCardReadModeDescription

`func (o *CloseResponseObjectCloseResponse) GetCardReadModeDescription() string`

GetCardReadModeDescription returns the CardReadModeDescription field if non-nil, zero value otherwise.

### GetCardReadModeDescriptionOk

`func (o *CloseResponseObjectCloseResponse) GetCardReadModeDescriptionOk() (*string, bool)`

GetCardReadModeDescriptionOk returns a tuple with the CardReadModeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardReadModeDescription

`func (o *CloseResponseObjectCloseResponse) SetCardReadModeDescription(v string)`

SetCardReadModeDescription sets CardReadModeDescription field to given value.

### HasCardReadModeDescription

`func (o *CloseResponseObjectCloseResponse) HasCardReadModeDescription() bool`

HasCardReadModeDescription returns a boolean if a field has been set.

### GetCardDescription

`func (o *CloseResponseObjectCloseResponse) GetCardDescription() string`

GetCardDescription returns the CardDescription field if non-nil, zero value otherwise.

### GetCardDescriptionOk

`func (o *CloseResponseObjectCloseResponse) GetCardDescriptionOk() (*string, bool)`

GetCardDescriptionOk returns a tuple with the CardDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDescription

`func (o *CloseResponseObjectCloseResponse) SetCardDescription(v string)`

SetCardDescription sets CardDescription field to given value.

### HasCardDescription

`func (o *CloseResponseObjectCloseResponse) HasCardDescription() bool`

HasCardDescription returns a boolean if a field has been set.

### GetCardTypeDescription

`func (o *CloseResponseObjectCloseResponse) GetCardTypeDescription() string`

GetCardTypeDescription returns the CardTypeDescription field if non-nil, zero value otherwise.

### GetCardTypeDescriptionOk

`func (o *CloseResponseObjectCloseResponse) GetCardTypeDescriptionOk() (*string, bool)`

GetCardTypeDescriptionOk returns a tuple with the CardTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardTypeDescription

`func (o *CloseResponseObjectCloseResponse) SetCardTypeDescription(v string)`

SetCardTypeDescription sets CardTypeDescription field to given value.

### HasCardTypeDescription

`func (o *CloseResponseObjectCloseResponse) HasCardTypeDescription() bool`

HasCardTypeDescription returns a boolean if a field has been set.

### GetCardCategory

`func (o *CloseResponseObjectCloseResponse) GetCardCategory() SaleResponseObjectSaleResponseCardCategory`

GetCardCategory returns the CardCategory field if non-nil, zero value otherwise.

### GetCardCategoryOk

`func (o *CloseResponseObjectCloseResponse) GetCardCategoryOk() (*SaleResponseObjectSaleResponseCardCategory, bool)`

GetCardCategoryOk returns a tuple with the CardCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCategory

`func (o *CloseResponseObjectCloseResponse) SetCardCategory(v SaleResponseObjectSaleResponseCardCategory)`

SetCardCategory sets CardCategory field to given value.

### HasCardCategory

`func (o *CloseResponseObjectCloseResponse) HasCardCategory() bool`

HasCardCategory returns a boolean if a field has been set.

### GetCardNumber

`func (o *CloseResponseObjectCloseResponse) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CloseResponseObjectCloseResponse) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CloseResponseObjectCloseResponse) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *CloseResponseObjectCloseResponse) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNumberMasked

`func (o *CloseResponseObjectCloseResponse) GetCardNumberMasked() string`

GetCardNumberMasked returns the CardNumberMasked field if non-nil, zero value otherwise.

### GetCardNumberMaskedOk

`func (o *CloseResponseObjectCloseResponse) GetCardNumberMaskedOk() (*string, bool)`

GetCardNumberMaskedOk returns a tuple with the CardNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberMasked

`func (o *CloseResponseObjectCloseResponse) SetCardNumberMasked(v string)`

SetCardNumberMasked sets CardNumberMasked field to given value.

### HasCardNumberMasked

`func (o *CloseResponseObjectCloseResponse) HasCardNumberMasked() bool`

HasCardNumberMasked returns a boolean if a field has been set.

### GetCardHashing

`func (o *CloseResponseObjectCloseResponse) GetCardHashing() string`

GetCardHashing returns the CardHashing field if non-nil, zero value otherwise.

### GetCardHashingOk

`func (o *CloseResponseObjectCloseResponse) GetCardHashingOk() (*string, bool)`

GetCardHashingOk returns a tuple with the CardHashing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHashing

`func (o *CloseResponseObjectCloseResponse) SetCardHashing(v string)`

SetCardHashing sets CardHashing field to given value.

### HasCardHashing

`func (o *CloseResponseObjectCloseResponse) HasCardHashing() bool`

HasCardHashing returns a boolean if a field has been set.

### GetCardExp

`func (o *CloseResponseObjectCloseResponse) GetCardExp() string`

GetCardExp returns the CardExp field if non-nil, zero value otherwise.

### GetCardExpOk

`func (o *CloseResponseObjectCloseResponse) GetCardExpOk() (*string, bool)`

GetCardExpOk returns a tuple with the CardExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExp

`func (o *CloseResponseObjectCloseResponse) SetCardExp(v string)`

SetCardExp sets CardExp field to given value.

### HasCardExp

`func (o *CloseResponseObjectCloseResponse) HasCardExp() bool`

HasCardExp returns a boolean if a field has been set.

### GetCardCryptogramResponse

`func (o *CloseResponseObjectCloseResponse) GetCardCryptogramResponse() string`

GetCardCryptogramResponse returns the CardCryptogramResponse field if non-nil, zero value otherwise.

### GetCardCryptogramResponseOk

`func (o *CloseResponseObjectCloseResponse) GetCardCryptogramResponseOk() (*string, bool)`

GetCardCryptogramResponseOk returns a tuple with the CardCryptogramResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCryptogramResponse

`func (o *CloseResponseObjectCloseResponse) SetCardCryptogramResponse(v string)`

SetCardCryptogramResponse sets CardCryptogramResponse field to given value.

### HasCardCryptogramResponse

`func (o *CloseResponseObjectCloseResponse) HasCardCryptogramResponse() bool`

HasCardCryptogramResponse returns a boolean if a field has been set.

### GetCardAppName

`func (o *CloseResponseObjectCloseResponse) GetCardAppName() string`

GetCardAppName returns the CardAppName field if non-nil, zero value otherwise.

### GetCardAppNameOk

`func (o *CloseResponseObjectCloseResponse) GetCardAppNameOk() (*string, bool)`

GetCardAppNameOk returns a tuple with the CardAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppName

`func (o *CloseResponseObjectCloseResponse) SetCardAppName(v string)`

SetCardAppName sets CardAppName field to given value.

### HasCardAppName

`func (o *CloseResponseObjectCloseResponse) HasCardAppName() bool`

HasCardAppName returns a boolean if a field has been set.

### GetCardAppIdentifier

`func (o *CloseResponseObjectCloseResponse) GetCardAppIdentifier() string`

GetCardAppIdentifier returns the CardAppIdentifier field if non-nil, zero value otherwise.

### GetCardAppIdentifierOk

`func (o *CloseResponseObjectCloseResponse) GetCardAppIdentifierOk() (*string, bool)`

GetCardAppIdentifierOk returns a tuple with the CardAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppIdentifier

`func (o *CloseResponseObjectCloseResponse) SetCardAppIdentifier(v string)`

SetCardAppIdentifier sets CardAppIdentifier field to given value.

### HasCardAppIdentifier

`func (o *CloseResponseObjectCloseResponse) HasCardAppIdentifier() bool`

HasCardAppIdentifier returns a boolean if a field has been set.

### GetCardAppLabel

`func (o *CloseResponseObjectCloseResponse) GetCardAppLabel() string`

GetCardAppLabel returns the CardAppLabel field if non-nil, zero value otherwise.

### GetCardAppLabelOk

`func (o *CloseResponseObjectCloseResponse) GetCardAppLabelOk() (*string, bool)`

GetCardAppLabelOk returns a tuple with the CardAppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppLabel

`func (o *CloseResponseObjectCloseResponse) SetCardAppLabel(v string)`

SetCardAppLabel sets CardAppLabel field to given value.

### HasCardAppLabel

`func (o *CloseResponseObjectCloseResponse) HasCardAppLabel() bool`

HasCardAppLabel returns a boolean if a field has been set.

### GetCardAuthRequestCryptogram

`func (o *CloseResponseObjectCloseResponse) GetCardAuthRequestCryptogram() string`

GetCardAuthRequestCryptogram returns the CardAuthRequestCryptogram field if non-nil, zero value otherwise.

### GetCardAuthRequestCryptogramOk

`func (o *CloseResponseObjectCloseResponse) GetCardAuthRequestCryptogramOk() (*string, bool)`

GetCardAuthRequestCryptogramOk returns a tuple with the CardAuthRequestCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthRequestCryptogram

`func (o *CloseResponseObjectCloseResponse) SetCardAuthRequestCryptogram(v string)`

SetCardAuthRequestCryptogram sets CardAuthRequestCryptogram field to given value.

### HasCardAuthRequestCryptogram

`func (o *CloseResponseObjectCloseResponse) HasCardAuthRequestCryptogram() bool`

HasCardAuthRequestCryptogram returns a boolean if a field has been set.

### GetCardAuthResponseCryptogram

`func (o *CloseResponseObjectCloseResponse) GetCardAuthResponseCryptogram() string`

GetCardAuthResponseCryptogram returns the CardAuthResponseCryptogram field if non-nil, zero value otherwise.

### GetCardAuthResponseCryptogramOk

`func (o *CloseResponseObjectCloseResponse) GetCardAuthResponseCryptogramOk() (*string, bool)`

GetCardAuthResponseCryptogramOk returns a tuple with the CardAuthResponseCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthResponseCryptogram

`func (o *CloseResponseObjectCloseResponse) SetCardAuthResponseCryptogram(v string)`

SetCardAuthResponseCryptogram sets CardAuthResponseCryptogram field to given value.

### HasCardAuthResponseCryptogram

`func (o *CloseResponseObjectCloseResponse) HasCardAuthResponseCryptogram() bool`

HasCardAuthResponseCryptogram returns a boolean if a field has been set.

### GetMerchantCategory

`func (o *CloseResponseObjectCloseResponse) GetMerchantCategory() SaleResponseObjectSaleResponseMerchantCategory`

GetMerchantCategory returns the MerchantCategory field if non-nil, zero value otherwise.

### GetMerchantCategoryOk

`func (o *CloseResponseObjectCloseResponse) GetMerchantCategoryOk() (*SaleResponseObjectSaleResponseMerchantCategory, bool)`

GetMerchantCategoryOk returns a tuple with the MerchantCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategory

`func (o *CloseResponseObjectCloseResponse) SetMerchantCategory(v SaleResponseObjectSaleResponseMerchantCategory)`

SetMerchantCategory sets MerchantCategory field to given value.

### HasMerchantCategory

`func (o *CloseResponseObjectCloseResponse) HasMerchantCategory() bool`

HasMerchantCategory returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CloseResponseObjectCloseResponse) GetPaymentMethod() SaleResponseObjectSaleResponsePaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CloseResponseObjectCloseResponse) GetPaymentMethodOk() (*SaleResponseObjectSaleResponsePaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CloseResponseObjectCloseResponse) SetPaymentMethod(v SaleResponseObjectSaleResponsePaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CloseResponseObjectCloseResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPlans

`func (o *CloseResponseObjectCloseResponse) GetPlans() SaleResponseObjectSaleResponsePlans`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *CloseResponseObjectCloseResponse) GetPlansOk() (*SaleResponseObjectSaleResponsePlans, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *CloseResponseObjectCloseResponse) SetPlans(v SaleResponseObjectSaleResponsePlans)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *CloseResponseObjectCloseResponse) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetConfiguration

`func (o *CloseResponseObjectCloseResponse) GetConfiguration() SaleResponseObjectSaleResponseConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CloseResponseObjectCloseResponse) GetConfigurationOk() (*SaleResponseObjectSaleResponseConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CloseResponseObjectCloseResponse) SetConfiguration(v SaleResponseObjectSaleResponseConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CloseResponseObjectCloseResponse) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


