# SaleObjectSaleSeller

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identification** | Pointer to **string** | Identificador del Vendedor | [optional] 
**IdentificationType** | Pointer to **string** | Tipo de documento del Vendedor. CI, Cédula de identidad - PAS, Pasaporte, Documento nacional de identidad, Contrato - CONTRACT, Número de Cuenta - ACCOUNT_NUMBER | [optional] 
**Email** | Pointer to **string** | Email del Vendedor | [optional] 
**DocumentType** | Pointer to **string** | Tipo de documento del Vendedor. CI, Cédula de identidad - PAS, Pasaporte, Documento nacional de identidad, Contrato - CONTRACT, Número de Cuenta - ACCOUNT_NUMBER | [optional] 
**DocumentNumber** | Pointer to **string** | Número de Documento o identificación del Vendedor | [optional] 
**FirstName** | Pointer to **string** | Primer Nombre del Vendedor | [optional] 
**LastName** | Pointer to **string** | Apellido del Vendedor | [optional] 
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
**CategoryCode** | Pointer to **int32** | Tipo o de categoría  del Comercio Vendedor (ISO 18245) | [optional] 
**TaxCategoryType** | Pointer to **string** | Tipo de categoría tributaria del Vendedor | [optional] 
**TaxIdentificationType** | Pointer to **string** | Identificador  Tributario ( RFC-Mexico, RUT-Chile, CUIT/CUIL-Argentina, etc) | [optional] 
**TaxIdentification** | Pointer to **string** | Identificador  Tributario | [optional] 
**NotifyURL** | Pointer to **string** | URL para notificación del Vendedor | [optional] 

## Methods

### NewSaleObjectSaleSeller

`func NewSaleObjectSaleSeller() *SaleObjectSaleSeller`

NewSaleObjectSaleSeller instantiates a new SaleObjectSaleSeller object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleObjectSaleSellerWithDefaults

`func NewSaleObjectSaleSellerWithDefaults() *SaleObjectSaleSeller`

NewSaleObjectSaleSellerWithDefaults instantiates a new SaleObjectSaleSeller object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentification

`func (o *SaleObjectSaleSeller) GetIdentification() string`

GetIdentification returns the Identification field if non-nil, zero value otherwise.

### GetIdentificationOk

`func (o *SaleObjectSaleSeller) GetIdentificationOk() (*string, bool)`

GetIdentificationOk returns a tuple with the Identification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentification

`func (o *SaleObjectSaleSeller) SetIdentification(v string)`

SetIdentification sets Identification field to given value.

### HasIdentification

`func (o *SaleObjectSaleSeller) HasIdentification() bool`

HasIdentification returns a boolean if a field has been set.

### GetIdentificationType

`func (o *SaleObjectSaleSeller) GetIdentificationType() string`

GetIdentificationType returns the IdentificationType field if non-nil, zero value otherwise.

### GetIdentificationTypeOk

`func (o *SaleObjectSaleSeller) GetIdentificationTypeOk() (*string, bool)`

GetIdentificationTypeOk returns a tuple with the IdentificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationType

`func (o *SaleObjectSaleSeller) SetIdentificationType(v string)`

SetIdentificationType sets IdentificationType field to given value.

### HasIdentificationType

`func (o *SaleObjectSaleSeller) HasIdentificationType() bool`

HasIdentificationType returns a boolean if a field has been set.

### GetEmail

`func (o *SaleObjectSaleSeller) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SaleObjectSaleSeller) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SaleObjectSaleSeller) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SaleObjectSaleSeller) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDocumentType

`func (o *SaleObjectSaleSeller) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *SaleObjectSaleSeller) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *SaleObjectSaleSeller) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *SaleObjectSaleSeller) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetDocumentNumber

`func (o *SaleObjectSaleSeller) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *SaleObjectSaleSeller) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *SaleObjectSaleSeller) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *SaleObjectSaleSeller) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### GetFirstName

`func (o *SaleObjectSaleSeller) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SaleObjectSaleSeller) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SaleObjectSaleSeller) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SaleObjectSaleSeller) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *SaleObjectSaleSeller) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SaleObjectSaleSeller) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SaleObjectSaleSeller) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SaleObjectSaleSeller) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMiddleName

`func (o *SaleObjectSaleSeller) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *SaleObjectSaleSeller) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *SaleObjectSaleSeller) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *SaleObjectSaleSeller) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetAbbreviatedName

`func (o *SaleObjectSaleSeller) GetAbbreviatedName() string`

GetAbbreviatedName returns the AbbreviatedName field if non-nil, zero value otherwise.

### GetAbbreviatedNameOk

`func (o *SaleObjectSaleSeller) GetAbbreviatedNameOk() (*string, bool)`

GetAbbreviatedNameOk returns a tuple with the AbbreviatedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviatedName

`func (o *SaleObjectSaleSeller) SetAbbreviatedName(v string)`

SetAbbreviatedName sets AbbreviatedName field to given value.

### HasAbbreviatedName

`func (o *SaleObjectSaleSeller) HasAbbreviatedName() bool`

HasAbbreviatedName returns a boolean if a field has been set.

### GetPhone

`func (o *SaleObjectSaleSeller) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SaleObjectSaleSeller) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SaleObjectSaleSeller) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SaleObjectSaleSeller) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetZipCode

`func (o *SaleObjectSaleSeller) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *SaleObjectSaleSeller) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *SaleObjectSaleSeller) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *SaleObjectSaleSeller) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetAddressStreet

`func (o *SaleObjectSaleSeller) GetAddressStreet() string`

GetAddressStreet returns the AddressStreet field if non-nil, zero value otherwise.

### GetAddressStreetOk

`func (o *SaleObjectSaleSeller) GetAddressStreetOk() (*string, bool)`

GetAddressStreetOk returns a tuple with the AddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStreet

`func (o *SaleObjectSaleSeller) SetAddressStreet(v string)`

SetAddressStreet sets AddressStreet field to given value.

### HasAddressStreet

`func (o *SaleObjectSaleSeller) HasAddressStreet() bool`

HasAddressStreet returns a boolean if a field has been set.

### GetAddressNumber

`func (o *SaleObjectSaleSeller) GetAddressNumber() string`

GetAddressNumber returns the AddressNumber field if non-nil, zero value otherwise.

### GetAddressNumberOk

`func (o *SaleObjectSaleSeller) GetAddressNumberOk() (*string, bool)`

GetAddressNumberOk returns a tuple with the AddressNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressNumber

`func (o *SaleObjectSaleSeller) SetAddressNumber(v string)`

SetAddressNumber sets AddressNumber field to given value.

### HasAddressNumber

`func (o *SaleObjectSaleSeller) HasAddressNumber() bool`

HasAddressNumber returns a boolean if a field has been set.

### GetAddressInternal

`func (o *SaleObjectSaleSeller) GetAddressInternal() string`

GetAddressInternal returns the AddressInternal field if non-nil, zero value otherwise.

### GetAddressInternalOk

`func (o *SaleObjectSaleSeller) GetAddressInternalOk() (*string, bool)`

GetAddressInternalOk returns a tuple with the AddressInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressInternal

`func (o *SaleObjectSaleSeller) SetAddressInternal(v string)`

SetAddressInternal sets AddressInternal field to given value.

### HasAddressInternal

`func (o *SaleObjectSaleSeller) HasAddressInternal() bool`

HasAddressInternal returns a boolean if a field has been set.

### GetAddressSuburb

`func (o *SaleObjectSaleSeller) GetAddressSuburb() string`

GetAddressSuburb returns the AddressSuburb field if non-nil, zero value otherwise.

### GetAddressSuburbOk

`func (o *SaleObjectSaleSeller) GetAddressSuburbOk() (*string, bool)`

GetAddressSuburbOk returns a tuple with the AddressSuburb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressSuburb

`func (o *SaleObjectSaleSeller) SetAddressSuburb(v string)`

SetAddressSuburb sets AddressSuburb field to given value.

### HasAddressSuburb

`func (o *SaleObjectSaleSeller) HasAddressSuburb() bool`

HasAddressSuburb returns a boolean if a field has been set.

### GetAddressDelegation

`func (o *SaleObjectSaleSeller) GetAddressDelegation() string`

GetAddressDelegation returns the AddressDelegation field if non-nil, zero value otherwise.

### GetAddressDelegationOk

`func (o *SaleObjectSaleSeller) GetAddressDelegationOk() (*string, bool)`

GetAddressDelegationOk returns a tuple with the AddressDelegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressDelegation

`func (o *SaleObjectSaleSeller) SetAddressDelegation(v string)`

SetAddressDelegation sets AddressDelegation field to given value.

### HasAddressDelegation

`func (o *SaleObjectSaleSeller) HasAddressDelegation() bool`

HasAddressDelegation returns a boolean if a field has been set.

### GetCity

`func (o *SaleObjectSaleSeller) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SaleObjectSaleSeller) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SaleObjectSaleSeller) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SaleObjectSaleSeller) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *SaleObjectSaleSeller) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SaleObjectSaleSeller) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SaleObjectSaleSeller) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SaleObjectSaleSeller) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *SaleObjectSaleSeller) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SaleObjectSaleSeller) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SaleObjectSaleSeller) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SaleObjectSaleSeller) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCategoryCode

`func (o *SaleObjectSaleSeller) GetCategoryCode() int32`

GetCategoryCode returns the CategoryCode field if non-nil, zero value otherwise.

### GetCategoryCodeOk

`func (o *SaleObjectSaleSeller) GetCategoryCodeOk() (*int32, bool)`

GetCategoryCodeOk returns a tuple with the CategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryCode

`func (o *SaleObjectSaleSeller) SetCategoryCode(v int32)`

SetCategoryCode sets CategoryCode field to given value.

### HasCategoryCode

`func (o *SaleObjectSaleSeller) HasCategoryCode() bool`

HasCategoryCode returns a boolean if a field has been set.

### GetTaxCategoryType

`func (o *SaleObjectSaleSeller) GetTaxCategoryType() string`

GetTaxCategoryType returns the TaxCategoryType field if non-nil, zero value otherwise.

### GetTaxCategoryTypeOk

`func (o *SaleObjectSaleSeller) GetTaxCategoryTypeOk() (*string, bool)`

GetTaxCategoryTypeOk returns a tuple with the TaxCategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCategoryType

`func (o *SaleObjectSaleSeller) SetTaxCategoryType(v string)`

SetTaxCategoryType sets TaxCategoryType field to given value.

### HasTaxCategoryType

`func (o *SaleObjectSaleSeller) HasTaxCategoryType() bool`

HasTaxCategoryType returns a boolean if a field has been set.

### GetTaxIdentificationType

`func (o *SaleObjectSaleSeller) GetTaxIdentificationType() string`

GetTaxIdentificationType returns the TaxIdentificationType field if non-nil, zero value otherwise.

### GetTaxIdentificationTypeOk

`func (o *SaleObjectSaleSeller) GetTaxIdentificationTypeOk() (*string, bool)`

GetTaxIdentificationTypeOk returns a tuple with the TaxIdentificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentificationType

`func (o *SaleObjectSaleSeller) SetTaxIdentificationType(v string)`

SetTaxIdentificationType sets TaxIdentificationType field to given value.

### HasTaxIdentificationType

`func (o *SaleObjectSaleSeller) HasTaxIdentificationType() bool`

HasTaxIdentificationType returns a boolean if a field has been set.

### GetTaxIdentification

`func (o *SaleObjectSaleSeller) GetTaxIdentification() string`

GetTaxIdentification returns the TaxIdentification field if non-nil, zero value otherwise.

### GetTaxIdentificationOk

`func (o *SaleObjectSaleSeller) GetTaxIdentificationOk() (*string, bool)`

GetTaxIdentificationOk returns a tuple with the TaxIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentification

`func (o *SaleObjectSaleSeller) SetTaxIdentification(v string)`

SetTaxIdentification sets TaxIdentification field to given value.

### HasTaxIdentification

`func (o *SaleObjectSaleSeller) HasTaxIdentification() bool`

HasTaxIdentification returns a boolean if a field has been set.

### GetNotifyURL

`func (o *SaleObjectSaleSeller) GetNotifyURL() string`

GetNotifyURL returns the NotifyURL field if non-nil, zero value otherwise.

### GetNotifyURLOk

`func (o *SaleObjectSaleSeller) GetNotifyURLOk() (*string, bool)`

GetNotifyURLOk returns a tuple with the NotifyURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyURL

`func (o *SaleObjectSaleSeller) SetNotifyURL(v string)`

SetNotifyURL sets NotifyURL field to given value.

### HasNotifyURL

`func (o *SaleObjectSaleSeller) HasNotifyURL() bool`

HasNotifyURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


