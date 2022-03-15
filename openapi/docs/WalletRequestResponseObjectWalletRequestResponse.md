# WalletRequestResponseObjectWalletRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestType** | Pointer to **string** | Tipo de Operación que se requirió, contendrá el mismo valor que se recibió en el requerimiento, sobre formatos que no soportan elementos complejos o compuestos. | [optional] 
**ResponseActions** | **[]string** | Acciones a realizar por parte del POS y/o PINPAD en base al resultado de la operación que ha sido procesada. Cada uno de estos actions o acciones están concatenadas por comas. Los posibles actions son OK, Approve, Refuse, IssuerCall, Tickets, WithHold, GetCard, UseTerminalToAuthorize, ConfigurationError, SystemError, ResourceError, ProcessError, Completed. | 
**ResponseMessage** | **string** | Descripción del resultado del proceso del requerimiento recibido. Esta descripción es generada por la plataforma, no por el Host que termine resolviendo la transacción. | 
**SystemIdentification** | Pointer to **string** | ID que identifica el sistema desde donde proviene la petición. | [optional] 
**CompanyIdentification** | Pointer to **string** | ID que identifica la companía desde donde proviene la petición. | [optional] 
**BranchIdentification** | Pointer to **string** | ID que identifica la sucursal desde donde proviene la petición. Esta sucursal pertenece a una determinada companía. | [optional] 
**POSIdentification** | Pointer to **string** | ID que identifica la caja o punto de venta desde donde proviene la petición. Este punto de venta pertenece a una determinada sucursal y companía. | [optional] 
**ForeignResponseCode** | Pointer to **string** | Código de respuesta para el sistema externo, es decir, para la aplicación cliente que se comunica con el TEF. | [optional] 
**ResponseCode** | **int32** | Código de Respuesta Interno de la plataforma, el POS debe actuar por lo que indican las acciones especificadas en ResponseActions y no por el código de respuesta informado en este campo o elemento, pero es una buena práctica que sea persistido por el mismo. | 
**AnswerType** | Pointer to **string** | Tipo de Operación que se está requiriendo, solo necesario sobre formatos que no soportan elementos complejos o compuestos. | [optional] 
**AnswerKey** | **string** | Código de identificación, generado por Plataforma, de la operación realizada | 
**RequestKey** | Pointer to **string** | ID generado para la identificación por parte del Plataforma de la información generada en la ejecución de un GetCard o un Payment Method. Sera necesario para que un mensaje de Sale, Void o Payment Method identifique el contexto generado y lo utilice para esa operación. | [optional] 
**ServerVersion** | Pointer to **string** | Versión del Servicio de la Plataforma que resolvió el requerimiento. | [optional] 
**ServerAddress** | Pointer to **string** | Dirección IP del Server que atiende el requerimiento. | [optional] 
**ServerInstance** | Pointer to **string** | Instancia de Server que atiende el requerimiento. | [optional] 
**ServerNodeName** | Pointer to **string** | Nombre del Nodo que atendió el requerimiento. | [optional] 
**MessageID** | Pointer to **string** | Identificador Unívoco del Mensaje ( UUID v5 ). | [optional] 
**AdapterInputVersion** | Pointer to **string** | Versión del  Adaptador de Protocolo Entrante que atiende el Requerimiento. | [optional] 
**AdapterInputAddress** | Pointer to **string** | Dirección IP del Adaptador de Protocolo Entrante que atiende el requerimiento. | [optional] 
**AdapterInputNodeName** | Pointer to **string** | Nombre del Nodo del Adaptador de Protocolo Entrante que atiende el requerimiento. | [optional] 
**AdapterOutputVersion** | Pointer to **string** | Versión del  Adaptador de Protocolo Saliente que atiende el Requerimiento. | [optional] 
**AdapterOutputAddress** | Pointer to **string** | Dirección IP  del  Adaptador de Protocolo Saliente que atiende el Requerimiento. | [optional] 
**AdapterOutputNodeName** | Pointer to **string** | Nombre del Nodo  del  Adaptador de Protocolo Saliente que atiende el Requerimiento. | [optional] 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 
**WasReversePrevious** | Pointer to **int32** | Flag indicador de generación de reversa para la última operación reversable | [optional] 
**WasClosedPreviousBlock** | Pointer to **int32** | Flag indicador de cancelación o confirmación del último bloque de transacciones, previo al nuevo valor recibido | [optional] 
**ReversedAnswerKey** | Pointer to **string** | ID que identifica a la operación que acaba de ser reversada. | [optional] 
**ReversedSequence** | Pointer to **string** | Secuencia de la transacción que fue reversada | [optional] 
**CommittedBlock** | Pointer to **string** | ID del bloque de transacciones que ha sido confirmado de forma automática (es decir, sin recibir un requerimiento de BlockClose). Este escenario se presentará si el Plataforma así se ha configurado para actuar bajo esa circunstancia. | [optional] 
**ReversedBlock** | Pointer to **string** | ID del bloque de transacciones que ha sido cancelado de forma automática (es decir, sin recibir un requerimiento de BlockClose). Este escenario se presentará si el Plataforma así se ha configurado para actuar bajo esa circunstancia. | [optional] 
**RequiredInformation** | Pointer to [**[]DebtPaymentObjectDebtPaymentRequiredInformation**](DebtPaymentObjectDebtPaymentRequiredInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**AdditionalInformation** | Pointer to [**[]SaleResponseObjectSaleResponseAdditionalInformation**](SaleResponseObjectSaleResponseAdditionalInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**DisplayResponseMessage** | Pointer to **[]string** | Información adicional/Mensaje promocional/Leyenda de respuesta a mostrar en pantalla en el ticket de la operación. Cada línea de este mensaje será un elemento dentro del array. | [optional] 
**Block** | Pointer to **string** | ID que identifica a un grupo de transacciones que serán confirmadas o canceladas | [optional] 
**TransmitionDateTime** | Pointer to **time.Time** | Fecha y hora de transmision de la operación hacia el host - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**TransactionIdentification** | Pointer to **string** | ID de La transacción UNIVOCO para el Punto de venta | [optional] 
**TransactionDescription** | Pointer to **string** | Descripción del tipo de operación que se realizará | [optional] 
**TrasactionDateTime** | Pointer to **string** | Fecha y Hora de la transacción en la plataforma de integración - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**TaxFinancialCostAmount** | Pointer to **float32** | Monto del recargo impositivo aplicado al costo financiero que la transacción tiene | [optional] 
**TaxFinancialCostPercentage** | Pointer to **float32** | Porcentaje de recargo impositivo a aplicar sobre el monto del costo financiero | [optional] 
**FinancialCostAmount** | Pointer to **float32** | Monto del Costo financiero total generado en base al plan elegido | [optional] 
**FinancialCostPercentage** | Pointer to **float32** | Porcentaje del costo financiero a aplicar al monto de la transacción | [optional] 
**RequestAmount** | Pointer to **float32** | Monto libre de costos financerios e impuestos por el que la venta se realizó. El monto cobrado realmente no es este, dado que no incluye las tasas e impuestos | [optional] 
**WalletRequestToken** | Pointer to **string** | Identificador Publico para ser identificar la transacción a realizar, normalmente es un código a Presentar | [optional] 
**WalletRequestTokenType** | Pointer to **string** | Tipo de Identificador Público para ser identificar la transacción a realizar, normalmente es un código a Presentar. En el caso de JPG o PNG las mismas estaran codificadas en Base64 | [optional] 
**PaymentMethodIdentification** | Pointer to **string** | Metodo de Pago utilizado en el Pago | [optional] 
**ForeignPaymentMethodIdentification** | Pointer to **string** | Metodo de Pago utilizado en el Pago expresado en la codificación que el aplicativo que se integra posea | [optional] 
**IssuerIdentification** | Pointer to **string** | Emisor del Medio de Pago utilizado | [optional] 
**ForeignIssuerIdentification** | Pointer to **string** | Emisor del Medio del Pago utilizado en el Pago expresado en la codificación que el aplicativo que se integra posea | [optional] 
**DisplayCredentialIdentification** | Pointer to **string** | Número o Identificador de Credencia que se puede Mostrar o Imprimir ( Normalmente Número de Tarjeta enmascarado para el caso de Tarjetas de Credito o débito) | [optional] 
**AuthTicket** | Pointer to **int32** | Número Ticket  o Voucher Generado para la Plataforma. | [optional] 
**AuthRRN** | Pointer to **string** | Número de identificación de la transacción, utilizado por la mayoría de los hosts para realizar anulaciones y devoluciones | [optional] 
**IdentifierForTheAdquirer** | Pointer to **string** | Identificador que genera el Host Adquirente para la Transacción en algunos podrá ser igual al AuthRRN | [optional] 
**IdentifierForTheResolutor** | Pointer to **string** | Identificador que genera el Host o Plataforma que resuelve la transacción | [optional] 
**RetryTime** | Pointer to **int32** | Tiempo para el proximo reintento de la operación &lt;b&gt;WalletRequest&lt;/b&gt; expresado en milisegúndos, si esta presenta la acción &lt;b&gt;Retry&lt;/b&gt; | [optional] 
**MerchantCategory** | Pointer to [**SaleResponseObjectSaleResponseMerchantCategory**](SaleResponseObjectSaleResponseMerchantCategory.md) |  | [optional] 
**Configuration** | Pointer to [**SaleResponseObjectSaleResponseConfiguration**](SaleResponseObjectSaleResponseConfiguration.md) |  | [optional] 
**Tickets** | Pointer to [**[]SaleResponseObjectSaleResponseTickets**](SaleResponseObjectSaleResponseTickets.md) | Elemento Compuesto que indica qué Tickets intervienen en la transacción y si deben ser digitalizados o impresos. | [optional] 

## Methods

### NewWalletRequestResponseObjectWalletRequestResponse

`func NewWalletRequestResponseObjectWalletRequestResponse(responseActions []string, responseMessage string, responseCode int32, answerKey string, ) *WalletRequestResponseObjectWalletRequestResponse`

NewWalletRequestResponseObjectWalletRequestResponse instantiates a new WalletRequestResponseObjectWalletRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletRequestResponseObjectWalletRequestResponseWithDefaults

`func NewWalletRequestResponseObjectWalletRequestResponseWithDefaults() *WalletRequestResponseObjectWalletRequestResponse`

NewWalletRequestResponseObjectWalletRequestResponseWithDefaults instantiates a new WalletRequestResponseObjectWalletRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestType

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetResponseActions

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetResponseActions() []string`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetResponseActionsOk() (*[]string, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetResponseActions(v []string)`

SetResponseActions sets ResponseActions field to given value.


### GetResponseMessage

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetSystemIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.

### HasSystemIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasSystemIdentification() bool`

HasSystemIdentification returns a boolean if a field has been set.

### GetCompanyIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetBranchIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetForeignResponseCode

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetForeignResponseCode() string`

GetForeignResponseCode returns the ForeignResponseCode field if non-nil, zero value otherwise.

### GetForeignResponseCodeOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetForeignResponseCodeOk() (*string, bool)`

GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignResponseCode

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetForeignResponseCode(v string)`

SetForeignResponseCode sets ForeignResponseCode field to given value.

### HasForeignResponseCode

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasForeignResponseCode() bool`

HasForeignResponseCode returns a boolean if a field has been set.

### GetResponseCode

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.


### GetAnswerType

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAnswerType() string`

GetAnswerType returns the AnswerType field if non-nil, zero value otherwise.

### GetAnswerTypeOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAnswerTypeOk() (*string, bool)`

GetAnswerTypeOk returns a tuple with the AnswerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerType

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetAnswerType(v string)`

SetAnswerType sets AnswerType field to given value.

### HasAnswerType

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasAnswerType() bool`

HasAnswerType returns a boolean if a field has been set.

### GetAnswerKey

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAnswerKey() string`

GetAnswerKey returns the AnswerKey field if non-nil, zero value otherwise.

### GetAnswerKeyOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAnswerKeyOk() (*string, bool)`

GetAnswerKeyOk returns a tuple with the AnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerKey

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetAnswerKey(v string)`

SetAnswerKey sets AnswerKey field to given value.


### GetRequestKey

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequestKey() string`

GetRequestKey returns the RequestKey field if non-nil, zero value otherwise.

### GetRequestKeyOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequestKeyOk() (*string, bool)`

GetRequestKeyOk returns a tuple with the RequestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestKey

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetRequestKey(v string)`

SetRequestKey sets RequestKey field to given value.

### HasRequestKey

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasRequestKey() bool`

HasRequestKey returns a boolean if a field has been set.

### GetServerVersion

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetServerAddress

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.

### GetServerInstance

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerInstance() string`

GetServerInstance returns the ServerInstance field if non-nil, zero value otherwise.

### GetServerInstanceOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerInstanceOk() (*string, bool)`

GetServerInstanceOk returns a tuple with the ServerInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstance

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetServerInstance(v string)`

SetServerInstance sets ServerInstance field to given value.

### HasServerInstance

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasServerInstance() bool`

HasServerInstance returns a boolean if a field has been set.

### GetServerNodeName

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerNodeName() string`

GetServerNodeName returns the ServerNodeName field if non-nil, zero value otherwise.

### GetServerNodeNameOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerNodeNameOk() (*string, bool)`

GetServerNodeNameOk returns a tuple with the ServerNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNodeName

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetServerNodeName(v string)`

SetServerNodeName sets ServerNodeName field to given value.

### HasServerNodeName

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasServerNodeName() bool`

HasServerNodeName returns a boolean if a field has been set.

### GetMessageID

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.

### HasMessageID

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasMessageID() bool`

HasMessageID returns a boolean if a field has been set.

### GetAdapterInputVersion

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterInputVersion() string`

GetAdapterInputVersion returns the AdapterInputVersion field if non-nil, zero value otherwise.

### GetAdapterInputVersionOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterInputVersionOk() (*string, bool)`

GetAdapterInputVersionOk returns a tuple with the AdapterInputVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputVersion

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetAdapterInputVersion(v string)`

SetAdapterInputVersion sets AdapterInputVersion field to given value.

### HasAdapterInputVersion

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasAdapterInputVersion() bool`

HasAdapterInputVersion returns a boolean if a field has been set.

### GetAdapterInputAddress

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterInputAddress() string`

GetAdapterInputAddress returns the AdapterInputAddress field if non-nil, zero value otherwise.

### GetAdapterInputAddressOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterInputAddressOk() (*string, bool)`

GetAdapterInputAddressOk returns a tuple with the AdapterInputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputAddress

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetAdapterInputAddress(v string)`

SetAdapterInputAddress sets AdapterInputAddress field to given value.

### HasAdapterInputAddress

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasAdapterInputAddress() bool`

HasAdapterInputAddress returns a boolean if a field has been set.

### GetAdapterInputNodeName

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterInputNodeName() string`

GetAdapterInputNodeName returns the AdapterInputNodeName field if non-nil, zero value otherwise.

### GetAdapterInputNodeNameOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterInputNodeNameOk() (*string, bool)`

GetAdapterInputNodeNameOk returns a tuple with the AdapterInputNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputNodeName

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetAdapterInputNodeName(v string)`

SetAdapterInputNodeName sets AdapterInputNodeName field to given value.

### HasAdapterInputNodeName

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasAdapterInputNodeName() bool`

HasAdapterInputNodeName returns a boolean if a field has been set.

### GetAdapterOutputVersion

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterOutputVersion() string`

GetAdapterOutputVersion returns the AdapterOutputVersion field if non-nil, zero value otherwise.

### GetAdapterOutputVersionOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterOutputVersionOk() (*string, bool)`

GetAdapterOutputVersionOk returns a tuple with the AdapterOutputVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputVersion

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetAdapterOutputVersion(v string)`

SetAdapterOutputVersion sets AdapterOutputVersion field to given value.

### HasAdapterOutputVersion

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasAdapterOutputVersion() bool`

HasAdapterOutputVersion returns a boolean if a field has been set.

### GetAdapterOutputAddress

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterOutputAddress() string`

GetAdapterOutputAddress returns the AdapterOutputAddress field if non-nil, zero value otherwise.

### GetAdapterOutputAddressOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterOutputAddressOk() (*string, bool)`

GetAdapterOutputAddressOk returns a tuple with the AdapterOutputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputAddress

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetAdapterOutputAddress(v string)`

SetAdapterOutputAddress sets AdapterOutputAddress field to given value.

### HasAdapterOutputAddress

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasAdapterOutputAddress() bool`

HasAdapterOutputAddress returns a boolean if a field has been set.

### GetAdapterOutputNodeName

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterOutputNodeName() string`

GetAdapterOutputNodeName returns the AdapterOutputNodeName field if non-nil, zero value otherwise.

### GetAdapterOutputNodeNameOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterOutputNodeNameOk() (*string, bool)`

GetAdapterOutputNodeNameOk returns a tuple with the AdapterOutputNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputNodeName

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetAdapterOutputNodeName(v string)`

SetAdapterOutputNodeName sets AdapterOutputNodeName field to given value.

### HasAdapterOutputNodeName

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasAdapterOutputNodeName() bool`

HasAdapterOutputNodeName returns a boolean if a field has been set.

### GetServiceVersion

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetWasReversePrevious

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetWasReversePrevious() int32`

GetWasReversePrevious returns the WasReversePrevious field if non-nil, zero value otherwise.

### GetWasReversePreviousOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetWasReversePreviousOk() (*int32, bool)`

GetWasReversePreviousOk returns a tuple with the WasReversePrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasReversePrevious

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetWasReversePrevious(v int32)`

SetWasReversePrevious sets WasReversePrevious field to given value.

### HasWasReversePrevious

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasWasReversePrevious() bool`

HasWasReversePrevious returns a boolean if a field has been set.

### GetWasClosedPreviousBlock

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetWasClosedPreviousBlock() int32`

GetWasClosedPreviousBlock returns the WasClosedPreviousBlock field if non-nil, zero value otherwise.

### GetWasClosedPreviousBlockOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetWasClosedPreviousBlockOk() (*int32, bool)`

GetWasClosedPreviousBlockOk returns a tuple with the WasClosedPreviousBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasClosedPreviousBlock

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetWasClosedPreviousBlock(v int32)`

SetWasClosedPreviousBlock sets WasClosedPreviousBlock field to given value.

### HasWasClosedPreviousBlock

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasWasClosedPreviousBlock() bool`

HasWasClosedPreviousBlock returns a boolean if a field has been set.

### GetReversedAnswerKey

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetReversedAnswerKey() string`

GetReversedAnswerKey returns the ReversedAnswerKey field if non-nil, zero value otherwise.

### GetReversedAnswerKeyOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetReversedAnswerKeyOk() (*string, bool)`

GetReversedAnswerKeyOk returns a tuple with the ReversedAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedAnswerKey

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetReversedAnswerKey(v string)`

SetReversedAnswerKey sets ReversedAnswerKey field to given value.

### HasReversedAnswerKey

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasReversedAnswerKey() bool`

HasReversedAnswerKey returns a boolean if a field has been set.

### GetReversedSequence

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetReversedSequence() string`

GetReversedSequence returns the ReversedSequence field if non-nil, zero value otherwise.

### GetReversedSequenceOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetReversedSequenceOk() (*string, bool)`

GetReversedSequenceOk returns a tuple with the ReversedSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedSequence

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetReversedSequence(v string)`

SetReversedSequence sets ReversedSequence field to given value.

### HasReversedSequence

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasReversedSequence() bool`

HasReversedSequence returns a boolean if a field has been set.

### GetCommittedBlock

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetCommittedBlock() string`

GetCommittedBlock returns the CommittedBlock field if non-nil, zero value otherwise.

### GetCommittedBlockOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetCommittedBlockOk() (*string, bool)`

GetCommittedBlockOk returns a tuple with the CommittedBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedBlock

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetCommittedBlock(v string)`

SetCommittedBlock sets CommittedBlock field to given value.

### HasCommittedBlock

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasCommittedBlock() bool`

HasCommittedBlock returns a boolean if a field has been set.

### GetReversedBlock

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetReversedBlock() string`

GetReversedBlock returns the ReversedBlock field if non-nil, zero value otherwise.

### GetReversedBlockOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetReversedBlockOk() (*string, bool)`

GetReversedBlockOk returns a tuple with the ReversedBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedBlock

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetReversedBlock(v string)`

SetReversedBlock sets ReversedBlock field to given value.

### HasReversedBlock

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasReversedBlock() bool`

HasReversedBlock returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequiredInformation() []DebtPaymentObjectDebtPaymentRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequiredInformationOk() (*[]DebtPaymentObjectDebtPaymentRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetRequiredInformation(v []DebtPaymentObjectDebtPaymentRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdditionalInformationOk() (*[]SaleResponseObjectSaleResponseAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetDisplayResponseMessage

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetDisplayResponseMessage() []string`

GetDisplayResponseMessage returns the DisplayResponseMessage field if non-nil, zero value otherwise.

### GetDisplayResponseMessageOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetDisplayResponseMessageOk() (*[]string, bool)`

GetDisplayResponseMessageOk returns a tuple with the DisplayResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayResponseMessage

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetDisplayResponseMessage(v []string)`

SetDisplayResponseMessage sets DisplayResponseMessage field to given value.

### HasDisplayResponseMessage

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasDisplayResponseMessage() bool`

HasDisplayResponseMessage returns a boolean if a field has been set.

### GetBlock

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetTransmitionDateTime

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetTransmitionDateTime() time.Time`

GetTransmitionDateTime returns the TransmitionDateTime field if non-nil, zero value otherwise.

### GetTransmitionDateTimeOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetTransmitionDateTimeOk() (*time.Time, bool)`

GetTransmitionDateTimeOk returns a tuple with the TransmitionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitionDateTime

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetTransmitionDateTime(v time.Time)`

SetTransmitionDateTime sets TransmitionDateTime field to given value.

### HasTransmitionDateTime

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasTransmitionDateTime() bool`

HasTransmitionDateTime returns a boolean if a field has been set.

### GetTransactionIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetTransactionIdentification() string`

GetTransactionIdentification returns the TransactionIdentification field if non-nil, zero value otherwise.

### GetTransactionIdentificationOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetTransactionIdentificationOk() (*string, bool)`

GetTransactionIdentificationOk returns a tuple with the TransactionIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetTransactionIdentification(v string)`

SetTransactionIdentification sets TransactionIdentification field to given value.

### HasTransactionIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasTransactionIdentification() bool`

HasTransactionIdentification returns a boolean if a field has been set.

### GetTransactionDescription

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetTransactionDescription() string`

GetTransactionDescription returns the TransactionDescription field if non-nil, zero value otherwise.

### GetTransactionDescriptionOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetTransactionDescriptionOk() (*string, bool)`

GetTransactionDescriptionOk returns a tuple with the TransactionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDescription

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetTransactionDescription(v string)`

SetTransactionDescription sets TransactionDescription field to given value.

### HasTransactionDescription

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasTransactionDescription() bool`

HasTransactionDescription returns a boolean if a field has been set.

### GetTrasactionDateTime

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetTrasactionDateTime() string`

GetTrasactionDateTime returns the TrasactionDateTime field if non-nil, zero value otherwise.

### GetTrasactionDateTimeOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetTrasactionDateTimeOk() (*string, bool)`

GetTrasactionDateTimeOk returns a tuple with the TrasactionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrasactionDateTime

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetTrasactionDateTime(v string)`

SetTrasactionDateTime sets TrasactionDateTime field to given value.

### HasTrasactionDateTime

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasTrasactionDateTime() bool`

HasTrasactionDateTime returns a boolean if a field has been set.

### GetTaxFinancialCostAmount

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetTaxFinancialCostAmount() float32`

GetTaxFinancialCostAmount returns the TaxFinancialCostAmount field if non-nil, zero value otherwise.

### GetTaxFinancialCostAmountOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetTaxFinancialCostAmountOk() (*float32, bool)`

GetTaxFinancialCostAmountOk returns a tuple with the TaxFinancialCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFinancialCostAmount

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetTaxFinancialCostAmount(v float32)`

SetTaxFinancialCostAmount sets TaxFinancialCostAmount field to given value.

### HasTaxFinancialCostAmount

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasTaxFinancialCostAmount() bool`

HasTaxFinancialCostAmount returns a boolean if a field has been set.

### GetTaxFinancialCostPercentage

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetTaxFinancialCostPercentage() float32`

GetTaxFinancialCostPercentage returns the TaxFinancialCostPercentage field if non-nil, zero value otherwise.

### GetTaxFinancialCostPercentageOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetTaxFinancialCostPercentageOk() (*float32, bool)`

GetTaxFinancialCostPercentageOk returns a tuple with the TaxFinancialCostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFinancialCostPercentage

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetTaxFinancialCostPercentage(v float32)`

SetTaxFinancialCostPercentage sets TaxFinancialCostPercentage field to given value.

### HasTaxFinancialCostPercentage

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasTaxFinancialCostPercentage() bool`

HasTaxFinancialCostPercentage returns a boolean if a field has been set.

### GetFinancialCostAmount

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetFinancialCostAmount() float32`

GetFinancialCostAmount returns the FinancialCostAmount field if non-nil, zero value otherwise.

### GetFinancialCostAmountOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetFinancialCostAmountOk() (*float32, bool)`

GetFinancialCostAmountOk returns a tuple with the FinancialCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialCostAmount

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetFinancialCostAmount(v float32)`

SetFinancialCostAmount sets FinancialCostAmount field to given value.

### HasFinancialCostAmount

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasFinancialCostAmount() bool`

HasFinancialCostAmount returns a boolean if a field has been set.

### GetFinancialCostPercentage

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetFinancialCostPercentage() float32`

GetFinancialCostPercentage returns the FinancialCostPercentage field if non-nil, zero value otherwise.

### GetFinancialCostPercentageOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetFinancialCostPercentageOk() (*float32, bool)`

GetFinancialCostPercentageOk returns a tuple with the FinancialCostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialCostPercentage

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetFinancialCostPercentage(v float32)`

SetFinancialCostPercentage sets FinancialCostPercentage field to given value.

### HasFinancialCostPercentage

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasFinancialCostPercentage() bool`

HasFinancialCostPercentage returns a boolean if a field has been set.

### GetRequestAmount

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequestAmount() float32`

GetRequestAmount returns the RequestAmount field if non-nil, zero value otherwise.

### GetRequestAmountOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequestAmountOk() (*float32, bool)`

GetRequestAmountOk returns a tuple with the RequestAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAmount

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetRequestAmount(v float32)`

SetRequestAmount sets RequestAmount field to given value.

### HasRequestAmount

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasRequestAmount() bool`

HasRequestAmount returns a boolean if a field has been set.

### GetWalletRequestToken

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetWalletRequestToken() string`

GetWalletRequestToken returns the WalletRequestToken field if non-nil, zero value otherwise.

### GetWalletRequestTokenOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetWalletRequestTokenOk() (*string, bool)`

GetWalletRequestTokenOk returns a tuple with the WalletRequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletRequestToken

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetWalletRequestToken(v string)`

SetWalletRequestToken sets WalletRequestToken field to given value.

### HasWalletRequestToken

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasWalletRequestToken() bool`

HasWalletRequestToken returns a boolean if a field has been set.

### GetWalletRequestTokenType

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetWalletRequestTokenType() string`

GetWalletRequestTokenType returns the WalletRequestTokenType field if non-nil, zero value otherwise.

### GetWalletRequestTokenTypeOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetWalletRequestTokenTypeOk() (*string, bool)`

GetWalletRequestTokenTypeOk returns a tuple with the WalletRequestTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletRequestTokenType

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetWalletRequestTokenType(v string)`

SetWalletRequestTokenType sets WalletRequestTokenType field to given value.

### HasWalletRequestTokenType

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasWalletRequestTokenType() bool`

HasWalletRequestTokenType returns a boolean if a field has been set.

### GetPaymentMethodIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetPaymentMethodIdentification() string`

GetPaymentMethodIdentification returns the PaymentMethodIdentification field if non-nil, zero value otherwise.

### GetPaymentMethodIdentificationOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetPaymentMethodIdentificationOk() (*string, bool)`

GetPaymentMethodIdentificationOk returns a tuple with the PaymentMethodIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetPaymentMethodIdentification(v string)`

SetPaymentMethodIdentification sets PaymentMethodIdentification field to given value.

### HasPaymentMethodIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasPaymentMethodIdentification() bool`

HasPaymentMethodIdentification returns a boolean if a field has been set.

### GetForeignPaymentMethodIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetForeignPaymentMethodIdentification() string`

GetForeignPaymentMethodIdentification returns the ForeignPaymentMethodIdentification field if non-nil, zero value otherwise.

### GetForeignPaymentMethodIdentificationOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetForeignPaymentMethodIdentificationOk() (*string, bool)`

GetForeignPaymentMethodIdentificationOk returns a tuple with the ForeignPaymentMethodIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignPaymentMethodIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetForeignPaymentMethodIdentification(v string)`

SetForeignPaymentMethodIdentification sets ForeignPaymentMethodIdentification field to given value.

### HasForeignPaymentMethodIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasForeignPaymentMethodIdentification() bool`

HasForeignPaymentMethodIdentification returns a boolean if a field has been set.

### GetIssuerIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetIssuerIdentification() string`

GetIssuerIdentification returns the IssuerIdentification field if non-nil, zero value otherwise.

### GetIssuerIdentificationOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetIssuerIdentificationOk() (*string, bool)`

GetIssuerIdentificationOk returns a tuple with the IssuerIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetIssuerIdentification(v string)`

SetIssuerIdentification sets IssuerIdentification field to given value.

### HasIssuerIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasIssuerIdentification() bool`

HasIssuerIdentification returns a boolean if a field has been set.

### GetForeignIssuerIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetForeignIssuerIdentification() string`

GetForeignIssuerIdentification returns the ForeignIssuerIdentification field if non-nil, zero value otherwise.

### GetForeignIssuerIdentificationOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetForeignIssuerIdentificationOk() (*string, bool)`

GetForeignIssuerIdentificationOk returns a tuple with the ForeignIssuerIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignIssuerIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetForeignIssuerIdentification(v string)`

SetForeignIssuerIdentification sets ForeignIssuerIdentification field to given value.

### HasForeignIssuerIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasForeignIssuerIdentification() bool`

HasForeignIssuerIdentification returns a boolean if a field has been set.

### GetDisplayCredentialIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetDisplayCredentialIdentification() string`

GetDisplayCredentialIdentification returns the DisplayCredentialIdentification field if non-nil, zero value otherwise.

### GetDisplayCredentialIdentificationOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetDisplayCredentialIdentificationOk() (*string, bool)`

GetDisplayCredentialIdentificationOk returns a tuple with the DisplayCredentialIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCredentialIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetDisplayCredentialIdentification(v string)`

SetDisplayCredentialIdentification sets DisplayCredentialIdentification field to given value.

### HasDisplayCredentialIdentification

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasDisplayCredentialIdentification() bool`

HasDisplayCredentialIdentification returns a boolean if a field has been set.

### GetAuthTicket

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAuthTicket() int32`

GetAuthTicket returns the AuthTicket field if non-nil, zero value otherwise.

### GetAuthTicketOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAuthTicketOk() (*int32, bool)`

GetAuthTicketOk returns a tuple with the AuthTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTicket

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetAuthTicket(v int32)`

SetAuthTicket sets AuthTicket field to given value.

### HasAuthTicket

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasAuthTicket() bool`

HasAuthTicket returns a boolean if a field has been set.

### GetAuthRRN

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAuthRRN() string`

GetAuthRRN returns the AuthRRN field if non-nil, zero value otherwise.

### GetAuthRRNOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetAuthRRNOk() (*string, bool)`

GetAuthRRNOk returns a tuple with the AuthRRN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRRN

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetAuthRRN(v string)`

SetAuthRRN sets AuthRRN field to given value.

### HasAuthRRN

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasAuthRRN() bool`

HasAuthRRN returns a boolean if a field has been set.

### GetIdentifierForTheAdquirer

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetIdentifierForTheAdquirer() string`

GetIdentifierForTheAdquirer returns the IdentifierForTheAdquirer field if non-nil, zero value otherwise.

### GetIdentifierForTheAdquirerOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetIdentifierForTheAdquirerOk() (*string, bool)`

GetIdentifierForTheAdquirerOk returns a tuple with the IdentifierForTheAdquirer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierForTheAdquirer

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetIdentifierForTheAdquirer(v string)`

SetIdentifierForTheAdquirer sets IdentifierForTheAdquirer field to given value.

### HasIdentifierForTheAdquirer

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasIdentifierForTheAdquirer() bool`

HasIdentifierForTheAdquirer returns a boolean if a field has been set.

### GetIdentifierForTheResolutor

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetIdentifierForTheResolutor() string`

GetIdentifierForTheResolutor returns the IdentifierForTheResolutor field if non-nil, zero value otherwise.

### GetIdentifierForTheResolutorOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetIdentifierForTheResolutorOk() (*string, bool)`

GetIdentifierForTheResolutorOk returns a tuple with the IdentifierForTheResolutor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierForTheResolutor

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetIdentifierForTheResolutor(v string)`

SetIdentifierForTheResolutor sets IdentifierForTheResolutor field to given value.

### HasIdentifierForTheResolutor

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasIdentifierForTheResolutor() bool`

HasIdentifierForTheResolutor returns a boolean if a field has been set.

### GetRetryTime

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetRetryTime() int32`

GetRetryTime returns the RetryTime field if non-nil, zero value otherwise.

### GetRetryTimeOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetRetryTimeOk() (*int32, bool)`

GetRetryTimeOk returns a tuple with the RetryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryTime

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetRetryTime(v int32)`

SetRetryTime sets RetryTime field to given value.

### HasRetryTime

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasRetryTime() bool`

HasRetryTime returns a boolean if a field has been set.

### GetMerchantCategory

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetMerchantCategory() SaleResponseObjectSaleResponseMerchantCategory`

GetMerchantCategory returns the MerchantCategory field if non-nil, zero value otherwise.

### GetMerchantCategoryOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetMerchantCategoryOk() (*SaleResponseObjectSaleResponseMerchantCategory, bool)`

GetMerchantCategoryOk returns a tuple with the MerchantCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategory

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetMerchantCategory(v SaleResponseObjectSaleResponseMerchantCategory)`

SetMerchantCategory sets MerchantCategory field to given value.

### HasMerchantCategory

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasMerchantCategory() bool`

HasMerchantCategory returns a boolean if a field has been set.

### GetConfiguration

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetConfiguration() SaleResponseObjectSaleResponseConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetConfigurationOk() (*SaleResponseObjectSaleResponseConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetConfiguration(v SaleResponseObjectSaleResponseConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTickets

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetTickets() []SaleResponseObjectSaleResponseTickets`

GetTickets returns the Tickets field if non-nil, zero value otherwise.

### GetTicketsOk

`func (o *WalletRequestResponseObjectWalletRequestResponse) GetTicketsOk() (*[]SaleResponseObjectSaleResponseTickets, bool)`

GetTicketsOk returns a tuple with the Tickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickets

`func (o *WalletRequestResponseObjectWalletRequestResponse) SetTickets(v []SaleResponseObjectSaleResponseTickets)`

SetTickets sets Tickets field to given value.

### HasTickets

`func (o *WalletRequestResponseObjectWalletRequestResponse) HasTickets() bool`

HasTickets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


