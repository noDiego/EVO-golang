# EnableServiceResponseObjectEnableServiceResponse

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
**Configuration** | Pointer to [**SaleResponseObjectSaleResponseConfiguration**](SaleResponseObjectSaleResponseConfiguration.md) |  | [optional] 
**TrasactionDateTime** | Pointer to **string** | Fecha y Hora de la transacción en la plataforma de integración - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 

## Methods

### NewEnableServiceResponseObjectEnableServiceResponse

`func NewEnableServiceResponseObjectEnableServiceResponse(responseActions []string, responseMessage string, responseCode int32, answerKey string, ) *EnableServiceResponseObjectEnableServiceResponse`

NewEnableServiceResponseObjectEnableServiceResponse instantiates a new EnableServiceResponseObjectEnableServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableServiceResponseObjectEnableServiceResponseWithDefaults

`func NewEnableServiceResponseObjectEnableServiceResponseWithDefaults() *EnableServiceResponseObjectEnableServiceResponse`

NewEnableServiceResponseObjectEnableServiceResponseWithDefaults instantiates a new EnableServiceResponseObjectEnableServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestType

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetResponseActions

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetResponseActions() []string`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetResponseActionsOk() (*[]string, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetResponseActions(v []string)`

SetResponseActions sets ResponseActions field to given value.


### GetResponseMessage

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetSystemIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.

### HasSystemIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasSystemIdentification() bool`

HasSystemIdentification returns a boolean if a field has been set.

### GetCompanyIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetBranchIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetForeignResponseCode

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetForeignResponseCode() string`

GetForeignResponseCode returns the ForeignResponseCode field if non-nil, zero value otherwise.

### GetForeignResponseCodeOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetForeignResponseCodeOk() (*string, bool)`

GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignResponseCode

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetForeignResponseCode(v string)`

SetForeignResponseCode sets ForeignResponseCode field to given value.

### HasForeignResponseCode

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasForeignResponseCode() bool`

HasForeignResponseCode returns a boolean if a field has been set.

### GetResponseCode

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.


### GetAnswerType

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAnswerType() string`

GetAnswerType returns the AnswerType field if non-nil, zero value otherwise.

### GetAnswerTypeOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAnswerTypeOk() (*string, bool)`

GetAnswerTypeOk returns a tuple with the AnswerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerType

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetAnswerType(v string)`

SetAnswerType sets AnswerType field to given value.

### HasAnswerType

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasAnswerType() bool`

HasAnswerType returns a boolean if a field has been set.

### GetAnswerKey

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAnswerKey() string`

GetAnswerKey returns the AnswerKey field if non-nil, zero value otherwise.

### GetAnswerKeyOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAnswerKeyOk() (*string, bool)`

GetAnswerKeyOk returns a tuple with the AnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerKey

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetAnswerKey(v string)`

SetAnswerKey sets AnswerKey field to given value.


### GetRequestKey

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetRequestKey() string`

GetRequestKey returns the RequestKey field if non-nil, zero value otherwise.

### GetRequestKeyOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetRequestKeyOk() (*string, bool)`

GetRequestKeyOk returns a tuple with the RequestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestKey

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetRequestKey(v string)`

SetRequestKey sets RequestKey field to given value.

### HasRequestKey

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasRequestKey() bool`

HasRequestKey returns a boolean if a field has been set.

### GetServerVersion

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetServerAddress

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.

### GetServerInstance

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetServerInstance() string`

GetServerInstance returns the ServerInstance field if non-nil, zero value otherwise.

### GetServerInstanceOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetServerInstanceOk() (*string, bool)`

GetServerInstanceOk returns a tuple with the ServerInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstance

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetServerInstance(v string)`

SetServerInstance sets ServerInstance field to given value.

### HasServerInstance

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasServerInstance() bool`

HasServerInstance returns a boolean if a field has been set.

### GetServerNodeName

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetServerNodeName() string`

GetServerNodeName returns the ServerNodeName field if non-nil, zero value otherwise.

### GetServerNodeNameOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetServerNodeNameOk() (*string, bool)`

GetServerNodeNameOk returns a tuple with the ServerNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNodeName

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetServerNodeName(v string)`

SetServerNodeName sets ServerNodeName field to given value.

### HasServerNodeName

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasServerNodeName() bool`

HasServerNodeName returns a boolean if a field has been set.

### GetMessageID

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.

### HasMessageID

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasMessageID() bool`

HasMessageID returns a boolean if a field has been set.

### GetAdapterInputVersion

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAdapterInputVersion() string`

GetAdapterInputVersion returns the AdapterInputVersion field if non-nil, zero value otherwise.

### GetAdapterInputVersionOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAdapterInputVersionOk() (*string, bool)`

GetAdapterInputVersionOk returns a tuple with the AdapterInputVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputVersion

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetAdapterInputVersion(v string)`

SetAdapterInputVersion sets AdapterInputVersion field to given value.

### HasAdapterInputVersion

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasAdapterInputVersion() bool`

HasAdapterInputVersion returns a boolean if a field has been set.

### GetAdapterInputAddress

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAdapterInputAddress() string`

GetAdapterInputAddress returns the AdapterInputAddress field if non-nil, zero value otherwise.

### GetAdapterInputAddressOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAdapterInputAddressOk() (*string, bool)`

GetAdapterInputAddressOk returns a tuple with the AdapterInputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputAddress

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetAdapterInputAddress(v string)`

SetAdapterInputAddress sets AdapterInputAddress field to given value.

### HasAdapterInputAddress

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasAdapterInputAddress() bool`

HasAdapterInputAddress returns a boolean if a field has been set.

### GetAdapterInputNodeName

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAdapterInputNodeName() string`

GetAdapterInputNodeName returns the AdapterInputNodeName field if non-nil, zero value otherwise.

### GetAdapterInputNodeNameOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAdapterInputNodeNameOk() (*string, bool)`

GetAdapterInputNodeNameOk returns a tuple with the AdapterInputNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputNodeName

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetAdapterInputNodeName(v string)`

SetAdapterInputNodeName sets AdapterInputNodeName field to given value.

### HasAdapterInputNodeName

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasAdapterInputNodeName() bool`

HasAdapterInputNodeName returns a boolean if a field has been set.

### GetAdapterOutputVersion

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAdapterOutputVersion() string`

GetAdapterOutputVersion returns the AdapterOutputVersion field if non-nil, zero value otherwise.

### GetAdapterOutputVersionOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAdapterOutputVersionOk() (*string, bool)`

GetAdapterOutputVersionOk returns a tuple with the AdapterOutputVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputVersion

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetAdapterOutputVersion(v string)`

SetAdapterOutputVersion sets AdapterOutputVersion field to given value.

### HasAdapterOutputVersion

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasAdapterOutputVersion() bool`

HasAdapterOutputVersion returns a boolean if a field has been set.

### GetAdapterOutputAddress

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAdapterOutputAddress() string`

GetAdapterOutputAddress returns the AdapterOutputAddress field if non-nil, zero value otherwise.

### GetAdapterOutputAddressOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAdapterOutputAddressOk() (*string, bool)`

GetAdapterOutputAddressOk returns a tuple with the AdapterOutputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputAddress

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetAdapterOutputAddress(v string)`

SetAdapterOutputAddress sets AdapterOutputAddress field to given value.

### HasAdapterOutputAddress

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasAdapterOutputAddress() bool`

HasAdapterOutputAddress returns a boolean if a field has been set.

### GetAdapterOutputNodeName

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAdapterOutputNodeName() string`

GetAdapterOutputNodeName returns the AdapterOutputNodeName field if non-nil, zero value otherwise.

### GetAdapterOutputNodeNameOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAdapterOutputNodeNameOk() (*string, bool)`

GetAdapterOutputNodeNameOk returns a tuple with the AdapterOutputNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputNodeName

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetAdapterOutputNodeName(v string)`

SetAdapterOutputNodeName sets AdapterOutputNodeName field to given value.

### HasAdapterOutputNodeName

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasAdapterOutputNodeName() bool`

HasAdapterOutputNodeName returns a boolean if a field has been set.

### GetServiceVersion

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetWasReversePrevious

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetWasReversePrevious() int32`

GetWasReversePrevious returns the WasReversePrevious field if non-nil, zero value otherwise.

### GetWasReversePreviousOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetWasReversePreviousOk() (*int32, bool)`

GetWasReversePreviousOk returns a tuple with the WasReversePrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasReversePrevious

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetWasReversePrevious(v int32)`

SetWasReversePrevious sets WasReversePrevious field to given value.

### HasWasReversePrevious

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasWasReversePrevious() bool`

HasWasReversePrevious returns a boolean if a field has been set.

### GetWasClosedPreviousBlock

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetWasClosedPreviousBlock() int32`

GetWasClosedPreviousBlock returns the WasClosedPreviousBlock field if non-nil, zero value otherwise.

### GetWasClosedPreviousBlockOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetWasClosedPreviousBlockOk() (*int32, bool)`

GetWasClosedPreviousBlockOk returns a tuple with the WasClosedPreviousBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasClosedPreviousBlock

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetWasClosedPreviousBlock(v int32)`

SetWasClosedPreviousBlock sets WasClosedPreviousBlock field to given value.

### HasWasClosedPreviousBlock

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasWasClosedPreviousBlock() bool`

HasWasClosedPreviousBlock returns a boolean if a field has been set.

### GetReversedAnswerKey

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetReversedAnswerKey() string`

GetReversedAnswerKey returns the ReversedAnswerKey field if non-nil, zero value otherwise.

### GetReversedAnswerKeyOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetReversedAnswerKeyOk() (*string, bool)`

GetReversedAnswerKeyOk returns a tuple with the ReversedAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedAnswerKey

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetReversedAnswerKey(v string)`

SetReversedAnswerKey sets ReversedAnswerKey field to given value.

### HasReversedAnswerKey

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasReversedAnswerKey() bool`

HasReversedAnswerKey returns a boolean if a field has been set.

### GetReversedSequence

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetReversedSequence() string`

GetReversedSequence returns the ReversedSequence field if non-nil, zero value otherwise.

### GetReversedSequenceOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetReversedSequenceOk() (*string, bool)`

GetReversedSequenceOk returns a tuple with the ReversedSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedSequence

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetReversedSequence(v string)`

SetReversedSequence sets ReversedSequence field to given value.

### HasReversedSequence

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasReversedSequence() bool`

HasReversedSequence returns a boolean if a field has been set.

### GetCommittedBlock

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetCommittedBlock() string`

GetCommittedBlock returns the CommittedBlock field if non-nil, zero value otherwise.

### GetCommittedBlockOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetCommittedBlockOk() (*string, bool)`

GetCommittedBlockOk returns a tuple with the CommittedBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedBlock

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetCommittedBlock(v string)`

SetCommittedBlock sets CommittedBlock field to given value.

### HasCommittedBlock

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasCommittedBlock() bool`

HasCommittedBlock returns a boolean if a field has been set.

### GetReversedBlock

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetReversedBlock() string`

GetReversedBlock returns the ReversedBlock field if non-nil, zero value otherwise.

### GetReversedBlockOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetReversedBlockOk() (*string, bool)`

GetReversedBlockOk returns a tuple with the ReversedBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedBlock

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetReversedBlock(v string)`

SetReversedBlock sets ReversedBlock field to given value.

### HasReversedBlock

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasReversedBlock() bool`

HasReversedBlock returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetRequiredInformation() []DebtPaymentObjectDebtPaymentRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetRequiredInformationOk() (*[]DebtPaymentObjectDebtPaymentRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetRequiredInformation(v []DebtPaymentObjectDebtPaymentRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetAdditionalInformationOk() (*[]SaleResponseObjectSaleResponseAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetDisplayResponseMessage

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetDisplayResponseMessage() []string`

GetDisplayResponseMessage returns the DisplayResponseMessage field if non-nil, zero value otherwise.

### GetDisplayResponseMessageOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetDisplayResponseMessageOk() (*[]string, bool)`

GetDisplayResponseMessageOk returns a tuple with the DisplayResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayResponseMessage

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetDisplayResponseMessage(v []string)`

SetDisplayResponseMessage sets DisplayResponseMessage field to given value.

### HasDisplayResponseMessage

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasDisplayResponseMessage() bool`

HasDisplayResponseMessage returns a boolean if a field has been set.

### GetBlock

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetTransmitionDateTime

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetTransmitionDateTime() time.Time`

GetTransmitionDateTime returns the TransmitionDateTime field if non-nil, zero value otherwise.

### GetTransmitionDateTimeOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetTransmitionDateTimeOk() (*time.Time, bool)`

GetTransmitionDateTimeOk returns a tuple with the TransmitionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitionDateTime

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetTransmitionDateTime(v time.Time)`

SetTransmitionDateTime sets TransmitionDateTime field to given value.

### HasTransmitionDateTime

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasTransmitionDateTime() bool`

HasTransmitionDateTime returns a boolean if a field has been set.

### GetTransactionIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetTransactionIdentification() string`

GetTransactionIdentification returns the TransactionIdentification field if non-nil, zero value otherwise.

### GetTransactionIdentificationOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetTransactionIdentificationOk() (*string, bool)`

GetTransactionIdentificationOk returns a tuple with the TransactionIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetTransactionIdentification(v string)`

SetTransactionIdentification sets TransactionIdentification field to given value.

### HasTransactionIdentification

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasTransactionIdentification() bool`

HasTransactionIdentification returns a boolean if a field has been set.

### GetTransactionDescription

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetTransactionDescription() string`

GetTransactionDescription returns the TransactionDescription field if non-nil, zero value otherwise.

### GetTransactionDescriptionOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetTransactionDescriptionOk() (*string, bool)`

GetTransactionDescriptionOk returns a tuple with the TransactionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDescription

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetTransactionDescription(v string)`

SetTransactionDescription sets TransactionDescription field to given value.

### HasTransactionDescription

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasTransactionDescription() bool`

HasTransactionDescription returns a boolean if a field has been set.

### GetConfiguration

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetConfiguration() SaleResponseObjectSaleResponseConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetConfigurationOk() (*SaleResponseObjectSaleResponseConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetConfiguration(v SaleResponseObjectSaleResponseConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTrasactionDateTime

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetTrasactionDateTime() string`

GetTrasactionDateTime returns the TrasactionDateTime field if non-nil, zero value otherwise.

### GetTrasactionDateTimeOk

`func (o *EnableServiceResponseObjectEnableServiceResponse) GetTrasactionDateTimeOk() (*string, bool)`

GetTrasactionDateTimeOk returns a tuple with the TrasactionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrasactionDateTime

`func (o *EnableServiceResponseObjectEnableServiceResponse) SetTrasactionDateTime(v string)`

SetTrasactionDateTime sets TrasactionDateTime field to given value.

### HasTrasactionDateTime

`func (o *EnableServiceResponseObjectEnableServiceResponse) HasTrasactionDateTime() bool`

HasTrasactionDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


