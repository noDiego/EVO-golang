# GetTransactionResponseObjectGetTransactionResponsePayer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identification** | Pointer to **string** | Identificador del Pagador | [optional] 
**IdentificationType** | Pointer to **string** | Tipo de documento del Pagador. CI, Cédula de identidad - PAS, Pasaporte, Documento nacional de identidad, Contrato - CONTRACT, Número de Cuenta - ACCOUNT_NUMBER | [optional] 
**Email** | Pointer to **string** | Email del Pagador | [optional] 
**DocumentType** | Pointer to **string** | Tipo de documento del Pagador. CI, Cédula de identidad - PAS, Pasaporte, Documento nacional de identidad, Contrato - CONTRACT, Número de Cuenta - ACCOUNT_NUMBER | [optional] 
**DocumentNumber** | Pointer to **string** | Número de Documento o identificación del Pagador | [optional] 
**FirstName** | Pointer to **string** | Primer Nombre del Pagador | [optional] 
**LastName** | Pointer to **string** | Apellido del Pagador | [optional] 
**MiddleName** | Pointer to **string** | Nombre/s del Medio | [optional] 
**AbbreviatedName** | Pointer to **string** | Nombre Abreviado | [optional] 
**Phone** | Pointer to **string** | Número de Teléfono | [optional] 
**ZipCode** | Pointer to **string** | Código Postal | [optional] 
**AddressStreet** | Pointer to **string** | Dirección, calle | [optional] 
**AddressNumber** | Pointer to **string** | Número esterior | [optional] 
**AddressInternal** | Pointer to **string** | Datos Adicionales de la dirección, apartamento, unidad , etc. | [optional] 
**AddressSuburb** | Pointer to **string** | Colonia, Barrio | [optional] 
**AddressDelegation** | Pointer to **string** | Delegación | [optional] 
**City** | Pointer to **string** | Código Postal | [optional] 
**State** | Pointer to **string** | Estado o Provincia | [optional] 
**Country** | Pointer to **string** | País | [optional] 
**CategoryCode** | Pointer to **int32** | Tipo o de categoría  del Comercio Pagador (ISO 18245) | [optional] 
**TaxCategoryType** | Pointer to **string** | Tipo de categoría tributaria del Pagador | [optional] 
**TaxIdentificationType** | Pointer to **string** | Identificador  Tributario ( RFC-Mexico, RUT-Chile, CUIT/CUIL-Argentina, etc) | [optional] 
**TaxIdentification** | Pointer to **string** | Identificador  Tributario | [optional] 
**NotifyURL** | Pointer to **string** | URL para notificación del Pagador | [optional] 
**Customer** | Pointer to [**SaleObjectSaleCustomer**](SaleObjectSaleCustomer.md) |  | [optional] 
**AmountToApply** | Pointer to **float32** | Importe o Monto de la Transacción a aplicar. | [optional] 
**AmountCharged** | Pointer to **float32** | Importe o Monto de la Transacción que efectivamente se cobro , si se envía en Void o Return en lugar de Amount, se genera un Ajuste si el Host lo soporta. | [optional] 
**CashbackAmount** | Pointer to **float32** | Monto del dinero en efectivo (cashback). | [optional] 
**TipAmount** | Pointer to **float32** | Importe o Monto de la Propina. | [optional] 
**Plan** | Pointer to **string** | Código/ID de Plan ( obtenido por la Operación PaymentMethod ) , en caso de ser enviado no se requiere en la Transacción el envío de CurrencyCode ni FacilityPayments | [optional] 
**CurrencyCode** | Pointer to **string** | código de Moneda - ISO 4217 &lt;https://en.wikipedia.org/wiki/ISO_4217 Se puede utilizar la Codificación Alfabética o Numérica &lt;br /&gt;   * Num   - Alpha - Description &lt;br /&gt;   * &#39;032&#39; - &#39;ARS&#39; - Pesos Argentinos &lt;br /&gt;   * &#39;152&#39; - &#39;CLP&#39; - Pesos Chilenos &lt;br/&gt;   * &#39;484&#39; - &#39;MXN&#39; - Pesos Mexicanos &lt;br/&gt;   * &#39;840&#39; - &#39;USD&#39; - dólares Americanos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;858&#39; - &#39;UYU&#39; - Pesos Uruguayos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;986&#39; - &#39;BRL&#39; - Real Brasileño | [optional] 
**FacilityPayments** | Pointer to **float32** | Cantidad de cuotas en las que sera realizada la transacción | [optional] 
**FacilityType** | Pointer to **string** | Tipo de Plan de Financiación | [optional] 
**TNA** | Pointer to **float32** | Se informará la tasa nominal anual, en casos en que el plan elegido para realizar la venta lo posea. Por ejemplo, el plan especial llamado Plan V de Prisma informara este valor, dado que se obtendra dinamicamente, consultandolo instante a instante | [optional] 
**TEM** | Pointer to **float32** | Se informará la tasa efectiva mensual, en casos en que el plan elegido para realizar la venta lo posea. Por ejemplo, el plan especial llamado Plan V de Prisma informara este valor, dado que se obtendra dinamicamente, consultandolo instante a instante | [optional] 
**TEA** | Pointer to **float32** | Tasa Efectiva anual. Este campo estara presente solo si el tipo de plan es dinamico, o si fue ingresado un valor en la base de datos | [optional] 
**CFT** | Pointer to **float32** | Costo Financiero Total. Este campo estara presente solo si el tipo de plan es dinamico, o si fue ingresado un valor en la base de datos         | [optional] 
**MerchantCategory** | Pointer to [**SaleResponseObjectSaleResponseMerchantCategory**](SaleResponseObjectSaleResponseMerchantCategory.md) |  | [optional] 
**Products** | Pointer to [**[]SaleResponseObjectSaleResponseProducts**](SaleResponseObjectSaleResponseProducts.md) | Detalle de Productos de la Operación. | [optional] 

## Methods

### NewGetTransactionResponseObjectGetTransactionResponsePayer

`func NewGetTransactionResponseObjectGetTransactionResponsePayer() *GetTransactionResponseObjectGetTransactionResponsePayer`

NewGetTransactionResponseObjectGetTransactionResponsePayer instantiates a new GetTransactionResponseObjectGetTransactionResponsePayer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionResponseObjectGetTransactionResponsePayerWithDefaults

`func NewGetTransactionResponseObjectGetTransactionResponsePayerWithDefaults() *GetTransactionResponseObjectGetTransactionResponsePayer`

NewGetTransactionResponseObjectGetTransactionResponsePayerWithDefaults instantiates a new GetTransactionResponseObjectGetTransactionResponsePayer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentification

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetIdentification() string`

GetIdentification returns the Identification field if non-nil, zero value otherwise.

### GetIdentificationOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetIdentificationOk() (*string, bool)`

GetIdentificationOk returns a tuple with the Identification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentification

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetIdentification(v string)`

SetIdentification sets Identification field to given value.

### HasIdentification

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasIdentification() bool`

HasIdentification returns a boolean if a field has been set.

### GetIdentificationType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetIdentificationType() string`

GetIdentificationType returns the IdentificationType field if non-nil, zero value otherwise.

### GetIdentificationTypeOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetIdentificationTypeOk() (*string, bool)`

GetIdentificationTypeOk returns a tuple with the IdentificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetIdentificationType(v string)`

SetIdentificationType sets IdentificationType field to given value.

### HasIdentificationType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasIdentificationType() bool`

HasIdentificationType returns a boolean if a field has been set.

### GetEmail

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDocumentType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetDocumentNumber

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### GetFirstName

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMiddleName

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetAbbreviatedName

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAbbreviatedName() string`

GetAbbreviatedName returns the AbbreviatedName field if non-nil, zero value otherwise.

### GetAbbreviatedNameOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAbbreviatedNameOk() (*string, bool)`

GetAbbreviatedNameOk returns a tuple with the AbbreviatedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviatedName

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetAbbreviatedName(v string)`

SetAbbreviatedName sets AbbreviatedName field to given value.

### HasAbbreviatedName

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasAbbreviatedName() bool`

HasAbbreviatedName returns a boolean if a field has been set.

### GetPhone

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetZipCode

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetAddressStreet

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAddressStreet() string`

GetAddressStreet returns the AddressStreet field if non-nil, zero value otherwise.

### GetAddressStreetOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAddressStreetOk() (*string, bool)`

GetAddressStreetOk returns a tuple with the AddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStreet

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetAddressStreet(v string)`

SetAddressStreet sets AddressStreet field to given value.

### HasAddressStreet

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasAddressStreet() bool`

HasAddressStreet returns a boolean if a field has been set.

### GetAddressNumber

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAddressNumber() string`

GetAddressNumber returns the AddressNumber field if non-nil, zero value otherwise.

### GetAddressNumberOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAddressNumberOk() (*string, bool)`

GetAddressNumberOk returns a tuple with the AddressNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressNumber

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetAddressNumber(v string)`

SetAddressNumber sets AddressNumber field to given value.

### HasAddressNumber

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasAddressNumber() bool`

HasAddressNumber returns a boolean if a field has been set.

### GetAddressInternal

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAddressInternal() string`

GetAddressInternal returns the AddressInternal field if non-nil, zero value otherwise.

### GetAddressInternalOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAddressInternalOk() (*string, bool)`

GetAddressInternalOk returns a tuple with the AddressInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressInternal

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetAddressInternal(v string)`

SetAddressInternal sets AddressInternal field to given value.

### HasAddressInternal

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasAddressInternal() bool`

HasAddressInternal returns a boolean if a field has been set.

### GetAddressSuburb

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAddressSuburb() string`

GetAddressSuburb returns the AddressSuburb field if non-nil, zero value otherwise.

### GetAddressSuburbOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAddressSuburbOk() (*string, bool)`

GetAddressSuburbOk returns a tuple with the AddressSuburb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressSuburb

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetAddressSuburb(v string)`

SetAddressSuburb sets AddressSuburb field to given value.

### HasAddressSuburb

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasAddressSuburb() bool`

HasAddressSuburb returns a boolean if a field has been set.

### GetAddressDelegation

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAddressDelegation() string`

GetAddressDelegation returns the AddressDelegation field if non-nil, zero value otherwise.

### GetAddressDelegationOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAddressDelegationOk() (*string, bool)`

GetAddressDelegationOk returns a tuple with the AddressDelegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressDelegation

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetAddressDelegation(v string)`

SetAddressDelegation sets AddressDelegation field to given value.

### HasAddressDelegation

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasAddressDelegation() bool`

HasAddressDelegation returns a boolean if a field has been set.

### GetCity

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCategoryCode

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetCategoryCode() int32`

GetCategoryCode returns the CategoryCode field if non-nil, zero value otherwise.

### GetCategoryCodeOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetCategoryCodeOk() (*int32, bool)`

GetCategoryCodeOk returns a tuple with the CategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryCode

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetCategoryCode(v int32)`

SetCategoryCode sets CategoryCode field to given value.

### HasCategoryCode

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasCategoryCode() bool`

HasCategoryCode returns a boolean if a field has been set.

### GetTaxCategoryType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetTaxCategoryType() string`

GetTaxCategoryType returns the TaxCategoryType field if non-nil, zero value otherwise.

### GetTaxCategoryTypeOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetTaxCategoryTypeOk() (*string, bool)`

GetTaxCategoryTypeOk returns a tuple with the TaxCategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCategoryType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetTaxCategoryType(v string)`

SetTaxCategoryType sets TaxCategoryType field to given value.

### HasTaxCategoryType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasTaxCategoryType() bool`

HasTaxCategoryType returns a boolean if a field has been set.

### GetTaxIdentificationType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetTaxIdentificationType() string`

GetTaxIdentificationType returns the TaxIdentificationType field if non-nil, zero value otherwise.

### GetTaxIdentificationTypeOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetTaxIdentificationTypeOk() (*string, bool)`

GetTaxIdentificationTypeOk returns a tuple with the TaxIdentificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentificationType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetTaxIdentificationType(v string)`

SetTaxIdentificationType sets TaxIdentificationType field to given value.

### HasTaxIdentificationType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasTaxIdentificationType() bool`

HasTaxIdentificationType returns a boolean if a field has been set.

### GetTaxIdentification

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetTaxIdentification() string`

GetTaxIdentification returns the TaxIdentification field if non-nil, zero value otherwise.

### GetTaxIdentificationOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetTaxIdentificationOk() (*string, bool)`

GetTaxIdentificationOk returns a tuple with the TaxIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentification

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetTaxIdentification(v string)`

SetTaxIdentification sets TaxIdentification field to given value.

### HasTaxIdentification

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasTaxIdentification() bool`

HasTaxIdentification returns a boolean if a field has been set.

### GetNotifyURL

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetNotifyURL() string`

GetNotifyURL returns the NotifyURL field if non-nil, zero value otherwise.

### GetNotifyURLOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetNotifyURLOk() (*string, bool)`

GetNotifyURLOk returns a tuple with the NotifyURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyURL

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetNotifyURL(v string)`

SetNotifyURL sets NotifyURL field to given value.

### HasNotifyURL

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasNotifyURL() bool`

HasNotifyURL returns a boolean if a field has been set.

### GetCustomer

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetCustomer() SaleObjectSaleCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetCustomerOk() (*SaleObjectSaleCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetCustomer(v SaleObjectSaleCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetAmountToApply

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAmountToApply() float32`

GetAmountToApply returns the AmountToApply field if non-nil, zero value otherwise.

### GetAmountToApplyOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAmountToApplyOk() (*float32, bool)`

GetAmountToApplyOk returns a tuple with the AmountToApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountToApply

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetAmountToApply(v float32)`

SetAmountToApply sets AmountToApply field to given value.

### HasAmountToApply

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasAmountToApply() bool`

HasAmountToApply returns a boolean if a field has been set.

### GetAmountCharged

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAmountCharged() float32`

GetAmountCharged returns the AmountCharged field if non-nil, zero value otherwise.

### GetAmountChargedOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetAmountChargedOk() (*float32, bool)`

GetAmountChargedOk returns a tuple with the AmountCharged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCharged

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetAmountCharged(v float32)`

SetAmountCharged sets AmountCharged field to given value.

### HasAmountCharged

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasAmountCharged() bool`

HasAmountCharged returns a boolean if a field has been set.

### GetCashbackAmount

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetCashbackAmount() float32`

GetCashbackAmount returns the CashbackAmount field if non-nil, zero value otherwise.

### GetCashbackAmountOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetCashbackAmountOk() (*float32, bool)`

GetCashbackAmountOk returns a tuple with the CashbackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackAmount

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetCashbackAmount(v float32)`

SetCashbackAmount sets CashbackAmount field to given value.

### HasCashbackAmount

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasCashbackAmount() bool`

HasCashbackAmount returns a boolean if a field has been set.

### GetTipAmount

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetTipAmount() float32`

GetTipAmount returns the TipAmount field if non-nil, zero value otherwise.

### GetTipAmountOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetTipAmountOk() (*float32, bool)`

GetTipAmountOk returns a tuple with the TipAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipAmount

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetTipAmount(v float32)`

SetTipAmount sets TipAmount field to given value.

### HasTipAmount

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasTipAmount() bool`

HasTipAmount returns a boolean if a field has been set.

### GetPlan

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetFacilityPayments

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetFacilityPayments() float32`

GetFacilityPayments returns the FacilityPayments field if non-nil, zero value otherwise.

### GetFacilityPaymentsOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetFacilityPaymentsOk() (*float32, bool)`

GetFacilityPaymentsOk returns a tuple with the FacilityPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityPayments

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetFacilityPayments(v float32)`

SetFacilityPayments sets FacilityPayments field to given value.

### HasFacilityPayments

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasFacilityPayments() bool`

HasFacilityPayments returns a boolean if a field has been set.

### GetFacilityType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetFacilityType() string`

GetFacilityType returns the FacilityType field if non-nil, zero value otherwise.

### GetFacilityTypeOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetFacilityTypeOk() (*string, bool)`

GetFacilityTypeOk returns a tuple with the FacilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetFacilityType(v string)`

SetFacilityType sets FacilityType field to given value.

### HasFacilityType

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasFacilityType() bool`

HasFacilityType returns a boolean if a field has been set.

### GetTNA

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetTNA() float32`

GetTNA returns the TNA field if non-nil, zero value otherwise.

### GetTNAOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetTNAOk() (*float32, bool)`

GetTNAOk returns a tuple with the TNA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTNA

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetTNA(v float32)`

SetTNA sets TNA field to given value.

### HasTNA

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasTNA() bool`

HasTNA returns a boolean if a field has been set.

### GetTEM

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetTEM() float32`

GetTEM returns the TEM field if non-nil, zero value otherwise.

### GetTEMOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetTEMOk() (*float32, bool)`

GetTEMOk returns a tuple with the TEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEM

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetTEM(v float32)`

SetTEM sets TEM field to given value.

### HasTEM

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasTEM() bool`

HasTEM returns a boolean if a field has been set.

### GetTEA

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetTEA() float32`

GetTEA returns the TEA field if non-nil, zero value otherwise.

### GetTEAOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetTEAOk() (*float32, bool)`

GetTEAOk returns a tuple with the TEA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEA

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetTEA(v float32)`

SetTEA sets TEA field to given value.

### HasTEA

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasTEA() bool`

HasTEA returns a boolean if a field has been set.

### GetCFT

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetCFT() float32`

GetCFT returns the CFT field if non-nil, zero value otherwise.

### GetCFTOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetCFTOk() (*float32, bool)`

GetCFTOk returns a tuple with the CFT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCFT

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetCFT(v float32)`

SetCFT sets CFT field to given value.

### HasCFT

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasCFT() bool`

HasCFT returns a boolean if a field has been set.

### GetMerchantCategory

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetMerchantCategory() SaleResponseObjectSaleResponseMerchantCategory`

GetMerchantCategory returns the MerchantCategory field if non-nil, zero value otherwise.

### GetMerchantCategoryOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetMerchantCategoryOk() (*SaleResponseObjectSaleResponseMerchantCategory, bool)`

GetMerchantCategoryOk returns a tuple with the MerchantCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategory

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetMerchantCategory(v SaleResponseObjectSaleResponseMerchantCategory)`

SetMerchantCategory sets MerchantCategory field to given value.

### HasMerchantCategory

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasMerchantCategory() bool`

HasMerchantCategory returns a boolean if a field has been set.

### GetProducts

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetProducts() []SaleResponseObjectSaleResponseProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) GetProductsOk() (*[]SaleResponseObjectSaleResponseProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) SetProducts(v []SaleResponseObjectSaleResponseProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *GetTransactionResponseObjectGetTransactionResponsePayer) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


