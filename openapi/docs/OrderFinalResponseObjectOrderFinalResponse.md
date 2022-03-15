# OrderFinalResponseObjectOrderFinalResponse

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
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 
**AdditionalInformation** | Pointer to [**[]SaleResponseObjectSaleResponseAdditionalInformation**](SaleResponseObjectSaleResponseAdditionalInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**RequiredInformation** | Pointer to [**[]DebtPaymentObjectDebtPaymentRequiredInformation**](DebtPaymentObjectDebtPaymentRequiredInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**TransactionIdentification** | Pointer to **string** | ID de la operación a realizar, generado por el sistema externo | [optional] 
**TransactionDescription** | Pointer to **string** | Descripción del tipo de operación que se realizará | [optional] 
**PaymentTransactionIdentification** | Pointer to **string** | . | [optional] 
**HostResponseCode** | Pointer to **string** | . | [optional] 
**HostResponseMessage** | Pointer to **string** | . | [optional] 
**HostResponseAction** | Pointer to **string** | . | [optional] 
**HostMerchantIdentification** | Pointer to **string** | . | [optional] 
**HostTerminalIdentification** | Pointer to **string** | . | [optional] 
**HostBatchNumber** | Pointer to **string** | . | [optional] 
**HostTicketNumber** | Pointer to **string** | . | [optional] 
**HostTraceNumber** | Pointer to **string** | . | [optional] 
**HostRetrievalReferenceNumber** | Pointer to **string** | . | [optional] 
**HostTransactionIdentification** | Pointer to **string** | . | [optional] 
**HostAuthorizationCode** | Pointer to **string** | . | [optional] 
**CredentialToken** | Pointer to **string** | Token asociado a la Credencial Enrolada | [optional] 
**CredentialIssuerToken** | Pointer to **string** | Emisor del Token asociado a la credencial enrolada | [optional] 
**PlanID** | Pointer to **string** | . | [optional] 
**PaymentMethodId** | Pointer to **string** | . | [optional] 
**HostId** | Pointer to **string** | . | [optional] 
**AcquirerReferenceData** | Pointer to **string** | Identificador de la transacción, utilizado solo por algunos hosts para realizar anulaciones y devoluciones | [optional] 
**IdentifierForTheAcquirer** | Pointer to **string** | Identificador de la transacción generado por Plataforma para ser enviado al Adquirente | [optional] 
**Configuration** | Pointer to [**SaleResponseObjectSaleResponseConfiguration**](SaleResponseObjectSaleResponseConfiguration.md) |  | [optional] 

## Methods

### NewOrderFinalResponseObjectOrderFinalResponse

`func NewOrderFinalResponseObjectOrderFinalResponse(responseActions []string, responseMessage string, responseCode int32, ) *OrderFinalResponseObjectOrderFinalResponse`

NewOrderFinalResponseObjectOrderFinalResponse instantiates a new OrderFinalResponseObjectOrderFinalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFinalResponseObjectOrderFinalResponseWithDefaults

`func NewOrderFinalResponseObjectOrderFinalResponseWithDefaults() *OrderFinalResponseObjectOrderFinalResponse`

NewOrderFinalResponseObjectOrderFinalResponseWithDefaults instantiates a new OrderFinalResponseObjectOrderFinalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestType

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetResponseActions

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetResponseActions() []string`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetResponseActionsOk() (*[]string, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetResponseActions(v []string)`

SetResponseActions sets ResponseActions field to given value.


### GetResponseMessage

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetSystemIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.

### HasSystemIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasSystemIdentification() bool`

HasSystemIdentification returns a boolean if a field has been set.

### GetCompanyIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetBranchIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetForeignResponseCode

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetForeignResponseCode() string`

GetForeignResponseCode returns the ForeignResponseCode field if non-nil, zero value otherwise.

### GetForeignResponseCodeOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetForeignResponseCodeOk() (*string, bool)`

GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignResponseCode

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetForeignResponseCode(v string)`

SetForeignResponseCode sets ForeignResponseCode field to given value.

### HasForeignResponseCode

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasForeignResponseCode() bool`

HasForeignResponseCode returns a boolean if a field has been set.

### GetResponseCode

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.


### GetServiceVersion

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetAdditionalInformationOk() (*[]SaleResponseObjectSaleResponseAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetRequiredInformation() []DebtPaymentObjectDebtPaymentRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetRequiredInformationOk() (*[]DebtPaymentObjectDebtPaymentRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetRequiredInformation(v []DebtPaymentObjectDebtPaymentRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetTransactionIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetTransactionIdentification() string`

GetTransactionIdentification returns the TransactionIdentification field if non-nil, zero value otherwise.

### GetTransactionIdentificationOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetTransactionIdentificationOk() (*string, bool)`

GetTransactionIdentificationOk returns a tuple with the TransactionIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetTransactionIdentification(v string)`

SetTransactionIdentification sets TransactionIdentification field to given value.

### HasTransactionIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasTransactionIdentification() bool`

HasTransactionIdentification returns a boolean if a field has been set.

### GetTransactionDescription

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetTransactionDescription() string`

GetTransactionDescription returns the TransactionDescription field if non-nil, zero value otherwise.

### GetTransactionDescriptionOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetTransactionDescriptionOk() (*string, bool)`

GetTransactionDescriptionOk returns a tuple with the TransactionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDescription

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetTransactionDescription(v string)`

SetTransactionDescription sets TransactionDescription field to given value.

### HasTransactionDescription

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasTransactionDescription() bool`

HasTransactionDescription returns a boolean if a field has been set.

### GetPaymentTransactionIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetPaymentTransactionIdentification() string`

GetPaymentTransactionIdentification returns the PaymentTransactionIdentification field if non-nil, zero value otherwise.

### GetPaymentTransactionIdentificationOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetPaymentTransactionIdentificationOk() (*string, bool)`

GetPaymentTransactionIdentificationOk returns a tuple with the PaymentTransactionIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTransactionIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetPaymentTransactionIdentification(v string)`

SetPaymentTransactionIdentification sets PaymentTransactionIdentification field to given value.

### HasPaymentTransactionIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasPaymentTransactionIdentification() bool`

HasPaymentTransactionIdentification returns a boolean if a field has been set.

### GetHostResponseCode

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostResponseCode() string`

GetHostResponseCode returns the HostResponseCode field if non-nil, zero value otherwise.

### GetHostResponseCodeOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostResponseCodeOk() (*string, bool)`

GetHostResponseCodeOk returns a tuple with the HostResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostResponseCode

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostResponseCode(v string)`

SetHostResponseCode sets HostResponseCode field to given value.

### HasHostResponseCode

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostResponseCode() bool`

HasHostResponseCode returns a boolean if a field has been set.

### GetHostResponseMessage

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostResponseMessage() string`

GetHostResponseMessage returns the HostResponseMessage field if non-nil, zero value otherwise.

### GetHostResponseMessageOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostResponseMessageOk() (*string, bool)`

GetHostResponseMessageOk returns a tuple with the HostResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostResponseMessage

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostResponseMessage(v string)`

SetHostResponseMessage sets HostResponseMessage field to given value.

### HasHostResponseMessage

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostResponseMessage() bool`

HasHostResponseMessage returns a boolean if a field has been set.

### GetHostResponseAction

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostResponseAction() string`

GetHostResponseAction returns the HostResponseAction field if non-nil, zero value otherwise.

### GetHostResponseActionOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostResponseActionOk() (*string, bool)`

GetHostResponseActionOk returns a tuple with the HostResponseAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostResponseAction

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostResponseAction(v string)`

SetHostResponseAction sets HostResponseAction field to given value.

### HasHostResponseAction

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostResponseAction() bool`

HasHostResponseAction returns a boolean if a field has been set.

### GetHostMerchantIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostMerchantIdentification() string`

GetHostMerchantIdentification returns the HostMerchantIdentification field if non-nil, zero value otherwise.

### GetHostMerchantIdentificationOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostMerchantIdentificationOk() (*string, bool)`

GetHostMerchantIdentificationOk returns a tuple with the HostMerchantIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMerchantIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostMerchantIdentification(v string)`

SetHostMerchantIdentification sets HostMerchantIdentification field to given value.

### HasHostMerchantIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostMerchantIdentification() bool`

HasHostMerchantIdentification returns a boolean if a field has been set.

### GetHostTerminalIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTerminalIdentification() string`

GetHostTerminalIdentification returns the HostTerminalIdentification field if non-nil, zero value otherwise.

### GetHostTerminalIdentificationOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTerminalIdentificationOk() (*string, bool)`

GetHostTerminalIdentificationOk returns a tuple with the HostTerminalIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostTerminalIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostTerminalIdentification(v string)`

SetHostTerminalIdentification sets HostTerminalIdentification field to given value.

### HasHostTerminalIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostTerminalIdentification() bool`

HasHostTerminalIdentification returns a boolean if a field has been set.

### GetHostBatchNumber

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostBatchNumber() string`

GetHostBatchNumber returns the HostBatchNumber field if non-nil, zero value otherwise.

### GetHostBatchNumberOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostBatchNumberOk() (*string, bool)`

GetHostBatchNumberOk returns a tuple with the HostBatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostBatchNumber

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostBatchNumber(v string)`

SetHostBatchNumber sets HostBatchNumber field to given value.

### HasHostBatchNumber

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostBatchNumber() bool`

HasHostBatchNumber returns a boolean if a field has been set.

### GetHostTicketNumber

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTicketNumber() string`

GetHostTicketNumber returns the HostTicketNumber field if non-nil, zero value otherwise.

### GetHostTicketNumberOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTicketNumberOk() (*string, bool)`

GetHostTicketNumberOk returns a tuple with the HostTicketNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostTicketNumber

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostTicketNumber(v string)`

SetHostTicketNumber sets HostTicketNumber field to given value.

### HasHostTicketNumber

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostTicketNumber() bool`

HasHostTicketNumber returns a boolean if a field has been set.

### GetHostTraceNumber

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTraceNumber() string`

GetHostTraceNumber returns the HostTraceNumber field if non-nil, zero value otherwise.

### GetHostTraceNumberOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTraceNumberOk() (*string, bool)`

GetHostTraceNumberOk returns a tuple with the HostTraceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostTraceNumber

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostTraceNumber(v string)`

SetHostTraceNumber sets HostTraceNumber field to given value.

### HasHostTraceNumber

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostTraceNumber() bool`

HasHostTraceNumber returns a boolean if a field has been set.

### GetHostRetrievalReferenceNumber

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostRetrievalReferenceNumber() string`

GetHostRetrievalReferenceNumber returns the HostRetrievalReferenceNumber field if non-nil, zero value otherwise.

### GetHostRetrievalReferenceNumberOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostRetrievalReferenceNumberOk() (*string, bool)`

GetHostRetrievalReferenceNumberOk returns a tuple with the HostRetrievalReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostRetrievalReferenceNumber

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostRetrievalReferenceNumber(v string)`

SetHostRetrievalReferenceNumber sets HostRetrievalReferenceNumber field to given value.

### HasHostRetrievalReferenceNumber

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostRetrievalReferenceNumber() bool`

HasHostRetrievalReferenceNumber returns a boolean if a field has been set.

### GetHostTransactionIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTransactionIdentification() string`

GetHostTransactionIdentification returns the HostTransactionIdentification field if non-nil, zero value otherwise.

### GetHostTransactionIdentificationOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTransactionIdentificationOk() (*string, bool)`

GetHostTransactionIdentificationOk returns a tuple with the HostTransactionIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostTransactionIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostTransactionIdentification(v string)`

SetHostTransactionIdentification sets HostTransactionIdentification field to given value.

### HasHostTransactionIdentification

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostTransactionIdentification() bool`

HasHostTransactionIdentification returns a boolean if a field has been set.

### GetHostAuthorizationCode

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostAuthorizationCode() string`

GetHostAuthorizationCode returns the HostAuthorizationCode field if non-nil, zero value otherwise.

### GetHostAuthorizationCodeOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostAuthorizationCodeOk() (*string, bool)`

GetHostAuthorizationCodeOk returns a tuple with the HostAuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostAuthorizationCode

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostAuthorizationCode(v string)`

SetHostAuthorizationCode sets HostAuthorizationCode field to given value.

### HasHostAuthorizationCode

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostAuthorizationCode() bool`

HasHostAuthorizationCode returns a boolean if a field has been set.

### GetCredentialToken

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetCredentialToken() string`

GetCredentialToken returns the CredentialToken field if non-nil, zero value otherwise.

### GetCredentialTokenOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetCredentialTokenOk() (*string, bool)`

GetCredentialTokenOk returns a tuple with the CredentialToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialToken

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetCredentialToken(v string)`

SetCredentialToken sets CredentialToken field to given value.

### HasCredentialToken

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasCredentialToken() bool`

HasCredentialToken returns a boolean if a field has been set.

### GetCredentialIssuerToken

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetCredentialIssuerToken() string`

GetCredentialIssuerToken returns the CredentialIssuerToken field if non-nil, zero value otherwise.

### GetCredentialIssuerTokenOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetCredentialIssuerTokenOk() (*string, bool)`

GetCredentialIssuerTokenOk returns a tuple with the CredentialIssuerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIssuerToken

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetCredentialIssuerToken(v string)`

SetCredentialIssuerToken sets CredentialIssuerToken field to given value.

### HasCredentialIssuerToken

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasCredentialIssuerToken() bool`

HasCredentialIssuerToken returns a boolean if a field has been set.

### GetPlanID

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetPlanID() string`

GetPlanID returns the PlanID field if non-nil, zero value otherwise.

### GetPlanIDOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetPlanIDOk() (*string, bool)`

GetPlanIDOk returns a tuple with the PlanID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanID

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetPlanID(v string)`

SetPlanID sets PlanID field to given value.

### HasPlanID

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasPlanID() bool`

HasPlanID returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetHostId

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetAcquirerReferenceData

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetAcquirerReferenceData() string`

GetAcquirerReferenceData returns the AcquirerReferenceData field if non-nil, zero value otherwise.

### GetAcquirerReferenceDataOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetAcquirerReferenceDataOk() (*string, bool)`

GetAcquirerReferenceDataOk returns a tuple with the AcquirerReferenceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerReferenceData

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetAcquirerReferenceData(v string)`

SetAcquirerReferenceData sets AcquirerReferenceData field to given value.

### HasAcquirerReferenceData

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasAcquirerReferenceData() bool`

HasAcquirerReferenceData returns a boolean if a field has been set.

### GetIdentifierForTheAcquirer

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetIdentifierForTheAcquirer() string`

GetIdentifierForTheAcquirer returns the IdentifierForTheAcquirer field if non-nil, zero value otherwise.

### GetIdentifierForTheAcquirerOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetIdentifierForTheAcquirerOk() (*string, bool)`

GetIdentifierForTheAcquirerOk returns a tuple with the IdentifierForTheAcquirer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierForTheAcquirer

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetIdentifierForTheAcquirer(v string)`

SetIdentifierForTheAcquirer sets IdentifierForTheAcquirer field to given value.

### HasIdentifierForTheAcquirer

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasIdentifierForTheAcquirer() bool`

HasIdentifierForTheAcquirer returns a boolean if a field has been set.

### GetConfiguration

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetConfiguration() SaleResponseObjectSaleResponseConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *OrderFinalResponseObjectOrderFinalResponse) GetConfigurationOk() (*SaleResponseObjectSaleResponseConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *OrderFinalResponseObjectOrderFinalResponse) SetConfiguration(v SaleResponseObjectSaleResponseConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *OrderFinalResponseObjectOrderFinalResponse) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


