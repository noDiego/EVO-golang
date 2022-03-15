/*
EVO Payment API

<h1 style=\"border:4px solid #004785;color:#004785; font-family: 'Open Sans', sans-serif; font-size: 32px;\">API de Pagos</h1> <br /><br /> <h1 style=\"border:4px solid #004785;color:#004785; font-family: 'Open Sans', sans-serif; font-size: 24px;\">Descripción del Servicio</h1> <p style=\"color:#004785;\"><b><u>Documentación en formato OpenAPI 3.0</b></u></p> <br/> Contrato especificado según especificaciones https://www.openapis.org/ y https://swagger.io/.<br /><br />  En el site https://editor.swagger.io/ se dispone de un  Viewer, Editor y  Generar de Código ( SDK ) para varios lenguajes de programación; incluyendo JAVA, C#, C++, Perl, Node.js, GO, PHP, Ruby y otros.<br/><br/> Para <b>ver</b> la documentación o <b>generar</b> código de la librería cliente o SDK  se deberá selecciónar en el menú horizontal  la opción <b>File</b>, en el menú vertical que se depliega la opción <b>Import File</b> y luego se deberá selecciónar el archivo del contrato deseado, ya sea  extensión <b>.json</b> o <b>.yaml</b>. <br/><br/> Además se puede generar el código de la librería cliente desde la línea de comandos a través de la herramienta  <b>CLI</b>  de  <b>OpenAPI Generator</b>. Esta presenta generadores de SDK en mayor variedad de lenguajes de programación.  En el site https://openapi-generator.tech/docs/installation se documenta cómo  <b>instalar</b> la herramienta CLI.<br/><br/> Los clientes generados contienen, adicionalmente al código,  la documentación de uso del mismo en <b>README.md</b>, como también en el subdirectorio <b>docs</b> toda la documentación del API o servicio y sus operaciones, con el detalle de los  campos o elementos y su dominio.<br /><br /><br /><br /> <h1 style=\"border:4px solid #004785;color:#004785; font-family: 'Open Sans', sans-serif; font-size: 24px;\">Notas a tener en cuenta para realizar la Integración</h1><br/> <p style=\"color:#004785;\"><b><u>Conceptos y/ Mecanismos relevantes Soportados por el Protocolo de Integración</u></b></p> <br/><br/> <span>&#8226;</span> <b>Interpretración de las Respuesta</b>,<br /><br/> El único campo que indica si la transacción fue aprobada, rechazada, o tienen algun error, es el elemento de las respuestas llamado <b>ResponseActions</b>, el  cual es un <b>ARRAY</b> de valores. Cada uno de estos indica una acción a realizar. Los elementos <b>ResponseCode</b>  y <b>ResponseMessage</b> son solamente informativos y por lo tanto no deben usarse para tomar acciones y los mismos pueden cambiar en base a la configuración de la Plataforma.<br/><br/> <span>&#8226;</span> <b>Bloque de transacciones</b>, permite Confirmar o Cancelar/Revertir todas las transacciones que forman parte de un bloque. <br/><br/> El POS puede definir un bloque o conjunto de transacciones simplemente indicando en todas ellas el mismo valor en el atributo/elemento/campo opcional llamado <b>Block</b>.<br/>  La operación <b>BlockCancel</b>, permite que el POS pueda solicitar a la plataforma la reversión y/o cancelación de todo el bloque de transacciones .<br/> La operación <b>BlockClose</b>, confirma todas las transacciones que forman parte del bloque especificado.<br/> Si el POS no posee un identificador unívoco de la transacción de venta, al momento de interactuar contra la plataforma podrá obtener uno con la operación  <b>BlockCreate</b>. Si el elemento o campo <b>Block</b>  existe y su contenido es Vacío o Nulo la plataforma realiza un <b>BlockCreate</b> automáticamente.<br /><br/> <span>&#8226;</span> <b>Reversas por Ruptura de Secuencia</b>. Evita la necesidad de persistir datos de la reversa y ahorra una transacción en el flujo.<br/>   El método llamado de ruptura de secuencia es utilizado para detectar los casos en los cuales el POS o Caja no pudo recibir una respuesta del mismo o no pudo procesarla adecuadamente. De esta forma permite a la misma reversar la transacción que no pudo procesar el POS o recibir la respuesta si fuese necesario.     En todo requerimiento el POS debe enviar el campo/elemento Sequence, con el valor recibido en el anterior requerimiento o vacío en el primero.    La plataforma  genera una nueva secuencia solamente cuando el requerimiento realizado es reversible o cuando se produce una ruptura.  Por lo tanto los comandos en los cuales la plataforma  genera una nueva secuencia son <b>Sale</b>, <b>Void</b>, <b>Authorize*</b>, <b>DebtPayment</b>, <b>VoidDebtPayment</b>, <b>Confirm</b>, <b>Close</b> y <b>Cancel</b>.    En caso de que la plataforma reverse el requerimiento previo retornará en la respuesta los siguiente campos o elementos.   <blockquote><b>WasReversePrevious</b>, con valor <b>1</b><br/>   <b>ReversedAnswerKey</b> conteniendo el <b>AnswerKey</b> de la transacción reversada<br/>   <b>ReversedSequence</b> conteniendo el <b>Sequence </b>de la transacción reversada</blockquote>    En caso de que la plataforma no reverse el requerimiento previo retornará los siguientes campos o elementos <blockquote><b>WasReversePrevious</b>, con valor <b>0</b></blockquote> <br/> <span>&#8226;</span> <b>Reversas Tradicionales</b>. El POS debe repetir el mismo requerimiento adicionando el atributo/elemento <b>IsReverse</b> con valor <b>1</b>.  Se debe tener en cuenta que en esta modalidad la plataforma no retorna los siguientes atributos/elementos.    <blockquote>   <b>WasReversePrevious</b><br/>   <b>ReversedAnswerKey</b><br/>   <b>ReversedSequence</b>   </blockquote>    <span>&#8226;</span> <b>Transacción Opcional de Confirmación</b>, ya que el mecanismo anterior permite que cada transacción Reverse o Confirme la anterior.<br/><br/> <span>&#8226;</span> <b>La Plataforma indica siempre las acciones que se deben realizar</b><br/><br/> <span>&#8226;</span><span>&#8226;</span> <b>Solicitar datos adicionales</b> ( <b>RequiredInformation</b> ), indicando no sólo cuáles son, sino también de qué tipo, valor  inicial, patrón de validación, si son mandatorios o no, qué Label se presenta al usuario, qué ayuda se presenta al usuario, etc.<br/> <span>&#8226;</span><span>&#8226;</span> <b>Mostrar Mensajes en Pantalla</b>. <span>&#8226;</span><span>&#8226;</span> <b>Imprimir Tickets</b>, ya sea en papel o capturar digitalmente el mismo, como así también el Layout de los mismos.<br/><br/><br/> <span>&#8226;</span> <b>Compresión de la trama</b> en base a codificación de los campos numéricos, string siempre de longitud variable, uso de sinónimos en los  campos, para que el programador programe usando los nombres largos y en el transporte se usen sus sinónimos cortos. <br/> <br/> <span>&#8226;</span> <b>Seguridad de los Datos Sensibles y de la Transaccion</b>, El elemento <b>Security</b> debera estar presente solo si los datos sensibles <b>CardNumber</b>, <b>ExpDate</b>, <b>PIN</b>, <b>Track1</b>, <b>Track2</b>, <b>SecurityCode</b> y  <b>CardCryptogram</b> deban ser envidos encriptados y por lo tanto este le elemento nos permite indicar el metodo de encriptacion utilizado y los datos adicionales que sean requeridos por la encriptacion. Si por ejemplo fuese el elemento PIN usando DUKPT y el resto de los datos sencible Track1, Track2 y SecurityCode, se deberian enviar  de la siguiente forma: </br>        \"Security” :  [         {           \"Type\": \"PIN\",           \"Values\":  [               {                    \"Name\": \"Method\",                   \"Value\": \"DUKPT\"               },               {                    \"Name\": \"KSN\",                   \"Value\": \"1234567890ABCDEF\"               },               {                    \"Name\": \"CRC32\",                   \"Value\": \"12345\"               },               {                    \"Name\": \"PlainTextLength\",                   \"Value\": \"4\"               },               {                    \"Name\": \"CipherCounter\",                   \"Value\": \"123\"               },               {                    \"Name\": \"ConsecutiveFailedCiphersCounter\",                   \"Value\": \"123\"               },               {                    \"Name\": \"Data\",                   \"Value\": \"01234567890123456\"               }           ]          },         {           \"Type\": \"SensitiveData\",           \"Values\":  [               {                    \"Name\": \"Method\",                   \"Value\": \"DUKPT-eGlobal\"               },               {                    \"Name\": \"KSN\",                   \"Value\": \"1234567890ABCDEF\"               },               {                    \"Name\": \"Track1CRC32\",                   \"Value\": \"12345\"               },               {                    \"Name\": \"Track2CRC32\",                   \"Value\": \"12345\"               },               {                    \"Name\": \"Track1Length\",                   \"Value\": \"79\"               },               {                    \"Name\": \"Track2Length\",                   \"Value\": \"37\"               },               {                    \"Name\": \"SecurityCodeLength\",                   \"Value\": \"3\"               },               {                    \"Name\": \"CipherCounter\",                   \"Value\": \"123\"               },               {                    \"Name\": \"ConsecutiveFailedCiphersCounter\",                   \"Value\": \"123\"               },               {                    \"Name\": \"Data\",                   \"Value\": \"1ahbcd23412345123412b213b1324b1234b2134b2134132b4123b23\"               }           ]          },         {           \"Type\": \"3DSecure\",           \"Values\":  [               {                    \"Name\": \"Method\",                   \"Value\": \"3DS-SNAP\"               },               {                                           \"Name\":  \"TransactionStatus\",                   \"Value\": \"SuccessfullyAuthenticated\"               },               {                                           \"Name\":  \"AuthenticationECI\",                   \"Value\": \"05\"               },               {                                           \"Name\":  \"IsChallengeMandated\",                   \"Value\": \"false\"               },               ...               {                                           \"Name\":  \"AcsReferenceNumber\",                   \"Value\": \"3DS_LOA_ACS_PPFU_020100_00009\"               },               {                    \"Name\":  \"ProcessedAsDataOnly\",                   \"Value\": \"false\"               }           ]          }               ] </br> Para el caso de DUKPT-eGlobal, <b>Track2</b>, <b>SecurityCode</b> y <b>Track1</b> se cifraran formando parte del mismo Bloque, El mismo se debera formar con el Track2 ( reemplazando el signo = por el digito D ) completandolo hasta los 38 digitos con el digito F, luego el  SecurityCode completandolo hasta 10 digitos y por ultimo el Track1 padeado completando el bloque  de los 208 digitos.  </br> Este elemento <b>Security</b> sera utilizado para enviar cualquier dato de autenticacion del pagador por ejemplo 3DSecure, para el caso de que el proveedor de la Autenticacion sea SNAP se deberan contener como valores todos los elementos definidos en el objeto <b>ThreeDSInformation</b>.     </br> Este mecanismo podra utilizarse en el futuro para encriptar otros datos que sean sensibles pero no del medio de pago, si no de las personas. </br> <h1 style=\"border:4px solid #004785;color:#004785; font-family: 'Open Sans', sans-serif; font-size: 24px;\">Log de Cambios</h1></br> <span>&#8226;</span> <b>Versión 5.6.1</b> <span>&#8226;</span><span>&#8226;</span> Se añade el campo <b>MerchantCategory</b> en las respuestas de todas las transacciones. Sólo se enviará en caso de que la categoría de la compañia exista.</br> <span>&#8226;</span> <b>Versión 5.6.0</b> <span>&#8226;</span><span>&#8226;</span> Los campos <b>ResponseCode</b>, <b>ResponseMessage</b> y <b>ResponseActions</b> son <b>obligatorios</b> en las respuestas de todas las transacciones.</br> <span>&#8226;</span> <b>Versión 5.5.7</b> <span>&#8226;</span><span>&#8226;</span> Se añade el elemento <b>Notification</b>. El mismo se encuentra dentro de <b>SaleResponse</b> y <b>AuthorizeSaleResponse</b>. Notificación a generar alertas vía e-mail.</br> <span>&#8226;</span> <b>Versión 5.5.6</b> <span>&#8226;</span><span>&#8226;</span> Se añaden los elementos <b>CardAppLabel</b>, <b>CardAuthRequestCryptogram</b> y <b>CardAuthResponseCryptogram</b>, para facilitar el analisis de los POS y ReadingDevices, el contenido de dichos elementos se encontraba en Tag de los elementos CardCryptogram y CardCryptogramResponse.</br> <span>&#8226;</span>  <b>Versión 5.5.5</b> <span>&#8226;</span><span>&#8226;</span> Se modifican los elementos <b>AuthorizeSale</b> y <b>AuthorizeSaleResponse</b> para su correcta documentación. Además, se añade el campo <b>ReadingDeviceOperatingFrom</b> el cual indica desde cuando se encuentra operativo o encendido el dispositivo</br> <span>&#8226;</span> <b>Versión 5.5.4</b> <span>&#8226;</span><span>&#8226;</span> Se renombra el atributo <b>ReasonReverse</b> a <b>ReverseReason</b>. Dicho campo permite notificar en las Reversas la razón por la cual fue necesario generarla.</br> <span>&#8226;</span> <b>Versión 5.5.3</b> <span>&#8226;</span><span>&#8226;</span> Se agregan atributos al elemento <b>Configuration</b> para la operación <b>PaymentMethod</b>. Por otra parte, se añade el mismo en todas las operaciones donde no se encontraba documentado. </br><b>• Versión 5.5.2</b> <span>&#8226;</span><span>&#8226;</span> Se Agrega el elemento <b>Payer</b> con los datos del Pagador. Originalmente hasta esta version se envian los mismos en el elemento <b>Customer</b>, pero desde ahora se permite que se informen personas ( fisicas y juridicas ) como cliente comprador y como pagador. Si el elemento <b>Payer</b> no esta presente se tomaran los datos del elemento <b>Customer</b>. Se da soporte al Tipo de Ticket Payer.</br> <span>&#8226;</span> <b>Versión 5.5.1</b> </br> <span>&#8226;</span><span>&#8226;</span> Se completa la documentacion de los Elementos <b>Seller</b> y <b>Customer</b>, agregandose los atributos <b>City</b> y  <b>AbbreviatedName</b>.<br/>   <span>&#8226;</span><span>&#8226;</span> Se unifica la definicion del Elemento  <b>Customer</b> .<br/>   <span>&#8226;</span><span>&#8226;</span> Se agrega el Elemento <b>PaymentFacilitatorID</b> para indicar el Identificador de Facilitador de pagos o Payfac.</br> <span>&#8226;</span> <b>Versión 5.5.0</b> </br> <span>&#8226;</span><span>&#8226;</span> El elemento <b>ResponseActions</b> y <b>PosOrDeviceAction</b> de todas las operaciones deja de ser una lista.<br/>  de elementos en un string y se convierte en un array de string. Cada valor de la lista anterior está representada por un elemento del array.<br/>   <span>&#8226;</span><span>&#8226;</span> Se agregan los campos <b>ForeignIdentifier</b>, <b>SmallImage</b> y <b>LargeImage</b> en el campo <b>Wallets</b> de la operación <b>WalletsResponse</b>.<br/> <span>&#8226;</span><span>&#8226;</span> En el campo <b>PaymentMethods</b> de la operación <b>PaymentMethodsResponse</b> se agregan las properties <b>Imag</b>, <b>SmallImage</b> y <b>LargeImage</b>. Además se adiciona el campo <b>ID</b> en <b>Category</b> y el campo <b>ForeignIdentifier</b> en <b>Type</b>. <br/> <span>&#8226;</span><span>&#8226;</span> Se agrupan los campos relacionados con los datos del cliente y del vendedor en dos únicos campos de tipo objeto denominados <b>Customer</b> y <b>Seller</b>, respectivamente.<br/> <span>&#8226;</span><span>&#8226;</span> El elemento Layout del campo <b>Tickets</b> se convierte en un array de objetos que contiene elementos que permiten describir, dar formato y codificar los datos a imprimir. <br/> <span>&#8226;</span><span>&#8226;</span> Se documenta la operación <b>OrderStatus</b>.</br> <span>&#8226;</span><span>&#8226;</span> Los campos que refieren a tiempo y fecha se convierten en formato date-time. </br> <span>&#8226;</span><span>&#8226;</span> Se agrega el campo <b>ForeignResponseCode</b> en todas las respuestas de las operaciones, como un código de para el sistema externo, es decir, para la aplicación cliente que se comunica con el TEF.</br> <span>&#8226;</span><span>&#8226;</span> Se adiciona el campo <b>CardGetMode</b> que permite indicar por cada elemento que contiene los datos sensibles, si están encriptados y también el algoritmo usado. En caso de no estar especificado se asume PLAIN.</br> <span>&#8226;</span><span>&#8226;</span> Se agrega el campo <b>OrigReference</b> en aquellas operaciones que pueden referenciar a una transacción previa, como <b>Void</b>, <b>Return</b> y <b>GetTransaction</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se cambia la estrutura de la respuesta de la Operacion <b>GetTransacion</b> por errores. </br> <span>&#8226;</span><span>&#8226;</span> Se agregan las acciones Ok, Error y Retry en los campos <b>ResponseActions</b>.</br> <span>&#8226;</span><span>&#8226;</span> En aquellas operaciones financieras en las que se especifica la tarjeta se agrega en el requerimiento el campo <b>Pin</b>y en la respuesta el campo <b>WorkingKeys</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se agrega el campo <b>Security</b> con el objetivo de indicar los datos sensibles de seguridad de una transacción tanto en los requerimientos como en las respuestas de las operaciones disponibles.</br> <span>&#8226;</span><span>&#8226;</span> Se agrega la operacion <b>KeysRenewal</b> Las claves podran ser retornadas en el elemento <b>Security</b> y en caso de obtener como accion de respuesta <b>KeysRenewal</b> se esta indicando que esta nueva operacion debe ser ejecutada.<br/>      <span>&#8226;</span><span>&#8226;</span> Se agrega la opcion <b>Signature</b>  .<br/>     <span>&#8226;</span><span>&#8226;</span> Se agrega el elemento  <b>CategoryCode</b> para especificar el MCC del Vendedor y/o del Cliente  .<br/>     <span>&#8226;</span><span>&#8226;</span> Se agregan los Elementos <b>MerchantID</b>, <b>TerminalID</b>, <b>TraceNumber</b> y <b>SettlementBatchNumber</b> En los requerimientos, en caso que los mismos contengan valor los mismos seran utilizados para enviar al Host Resolutor de la Transaccion.</br>  <span>&#8226;</span><span>&#8226;</span> Se agregan los valores para pagos recurrentes a  los Elementos  <b>CardReadMode</b> y  <b>CardReadModeDescription</b> <span>&#8226;</span> <b>Versión 5.4.0</b> </br> <span>&#8226;</span><span>&#8226;</span> Se cambia la dirección IP por el nombre.</br> <span>&#8226;</span><span>&#8226;</span> Se contemplan los Datos del <b>Vendedor/Seller</b> y del <b>Cliente/Customer</b> en las operaciones  <b>WalletRequest</b>, <b>Sale</b>, <b>AuthorizeSale</b>, <b>DebtPayment</b>,  <b>Deposit</b>,  <b>Settlement</b>,  <b>Capture</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se Agregan los elementos <b>POSGEO</b> y <b>ReadingDeviceGEO</b> para que el dispositivo de lectura y el Punto de venta Notifiquen su coordenadas georefenciales en el momento de que se realiza la transacción.</br> <span>&#8226;</span><span>&#8226;</span> Se unifica y amplía el elemento <b>RequiredInformation</b>  tanto en los requerimientos como en las respuestas</br>  <span>&#8226;</span><span>&#8226;</span> Se cambia el tipo el elemento <b>CurrencyCode</b> a string para permitir cualquieras de la notaciones posibles.</br> <span>&#8226;</span><span>&#8226;</span> Se cambia el  elemento <b>Currency</b> por <b>CurrencyCode</b>  en el elemento <b>Plans</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se contemplan del detalle ( elemento <b>Products</b> ) de la venta en las operaciones  <b>WalletRequest</b>, <b>Sale</b>, <b>Void</b>, <b>Return</b>, <b>AuthorizeSale</b>, <b>DebtPayment</b>,  <b>VoidDebtPayment</b>, <b>Deposit</b>,  <b>Settlement</b>,  <b>Capture</b>.</br> <span>&#8226;</span><span>&#8226;</span> Agregamos la operación <b>DebtInquiry</b> que actua como sinónimo de <b>BalanceInquery</b>, la cual podía ser usada para consulta de Saldo y también de deuda.</br> <span>&#8226;</span><span>&#8226;</span> Se corrigen los tipos de Datos de Varios campos <b>Amount</b> que en lugar de string debían ser number.</br> <span>&#8226;</span><span>&#8226;</span> Se agregan las operaciones <b>QueryCompanies</b> y <b>QueryLineOfBusiness</b> para la consulta de Rubros y Empresas que se pueden utilizar para pagar Servicios/Deuda/Facturas con la operación <b>DebtPayment</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se adiciona el elemento <b>Companies</b> en la Operacion <b>BalanceInquiry</b> para el caso de que existan mas de una Compania para el mismo codigo o identificador de la deuda o factura a pagar y adicionalmente se agrega para ese caso la posibilidad de especificar a que compania corrende el Pago en el elemento <b>DebtCompanyIdentification</b> en la operación <b>DebtPayment</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se adiciona el elemento <b>BaseAmonut</b> en los requerimientos de las operación <b>Return</b>, el elemento <b>Reference</b>  en las operaciones <b>Sale</b>, <b>AuthorizeSale</b>, <b>Void</b>, <b>Return</b>, <b>DebtPayment</b>, <b>VoidDebtPayment</b>, <b>GetTransaction</b> y <b>WalletRequest</b>.  Además, se agregan los elementos <b>TaxFinancialCostAmount</b>, <b>TaxFinancialCostPercentage</b>, <b>FinancialCostAmount</b>, <b>FinancialCostPercentage</b> y <b>RequestAmount</b>  en las respuestas de dichas operaciones.</br> <span>&#8226;</span><span>&#8226;</span> En cada plan que se devuelve a través del <b>PaymentMethodResponse</b> estarán presentes <b>TaxFinancialCostAmount</b>,  <b>TaxFinancialCostPercentage</b>, <b>FinancialCostAmount</b> y <b>FinancialCostPercentage</b>. </br> <span>&#8226;</span><span>&#8226;</span> Se agregan los elementos  <b>CardAppName</b> y <b>CardAppIdentifier</b> en las peticiones de las operaciones <b>Sale</b>, <b>AuthorizeSale</b>,  <b>Void</b>, <b>Return</b>, <b>PaymentMethods</b>, <b>GetCard</b>, <b>Validate</b>, <b>DebtInquiery</b>, <b>BalanceInquiry</b>, <b>DebtPayment</b> y <b>VoidDebtPayment</b>.  Además, se agregan en las respuestas de algunas de ellas.</br> <span>&#8226;</span><span>&#8226;</span> Se cambia la estructura del elemento <b>Tickets</b> de las respuestas donde el elemento <b>Action</b>  hace referencia a las acciones que debe ejecutar el punto de venta, el elemento <b>DeviceAction</b> a las acciones que debe ejecutar el dispositivo y <b>ExecutedAction</b> a las acciones  que ejecutó la plataforma para sus <b>Tickets</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se adicionan los elementos <b>POSOrDeviceAction</b>, <b>OperationMode</b> y <b>OperationModeDescription</b> a la operación <b>Configure</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se agrega el elemento <b>RemainderAmount</b> a la operación <b>GetBlockResponse</b> que hace referencia a la diferencia entre el monto total de la transacción y las devoluciones parciales realizadas.</br> <span>&#8226;</span><span>&#8226;</span> Se corrijen errores en la definición de varios campos, como <b>ReadingDeviceType</b> y <b>CardReadMode</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se reemplaza el campo <b>ApplicationIdentification</b> por <b>SystemIdentification</b> en las operaciones <b>EnableService</b>, <b>Wallets</b>, <b>QueryCompanies</b>,  <b>QueryLineOfBusiness</b> y sus respectivas respuestas. </br> <span>&#8226;</span><span>&#8226;</span> Se agregan el identifidor Tributario en <b>Sale</b>, <b>AuthorizeSale</b>, <b>Void</b>,  <b>Return</b>, <b>DebtPayment</b>, <b>VoidDebtPayment</b> y <b>Debtinquery</b> que permite transacciones con <b>Token</b> ( <b>CredentialToken</b> y  <b>CredentialIssuerToken</b> ). En estas operaciones se elimina de mandatorias al campo <b>BranchIdentification</b> y <b>POSIdentification</b><br/> <span>&#8226;</span><span>&#8226;</span> Se agrega la operación <b>Enrollment</b>, la cual permite transacciones con <b>Token</b> ( <b>CredentialToken</b> y  <b>CredentialIssuerToken</b> ) y pagos recurrentes.</br> <span>&#8226;</span><span>&#8226;</span> El campo <b>ResponseAction</b> deja de ser un enum y se convierte en string. Se indica en la descripción los posibles actions.</br> <span>&#8226;</span><span>&#8226;</span> Se agregan los campos <b>SellerIdentification</b> y <b>SellerIdentificationType</b> en aquellas operaciones en las que se especifican con los datos del vendedor.</br> <span>&#8226;</span><span>&#8226;</span> El campo <b>FacilityPayments</b> deja de ser mandatario en las operaciones <b>Enrollment</b> y <b>Sale</b>. </br> <span>&#8226;</span><span>&#8226;</span> Se elimina la posibilidad de envío en el header HTTP.<br/> <span>&#8226;</span><span>&#8226;</span> Se agregan los campos <b>CashbackAmount</b> y <b>TipAmount</b> en la operación <b>WalletRequest</b>.<br/> <span>&#8226;</span><span>&#8226;</span> Se adiciona en el campo <b>CardReadMode</b> la opción K de Token.<br/> <span>&#8226;</span><span>&#8226;</span> Se corrige el campo <b>Answertype</b> y se modifica por <b>AnswerType</b>.<br/> <span>&#8226;</span><span>&#8226;</span> Se agregan los campos referidos al vendedor en las operaciones <b>Void</b> y <b>Return</b>. <br/> <span>&#8226;</span><span>&#8226;</span> Se crea un primer nivel para cada operación de tipo objeto. <br/>  <span>&#8226;</span><span>&#8226;</span> Se crea el campo <b>InputTokens</b> como un array de objetos que contienen Name y Value como properties en las operaciones <b>Sale</b>, <b>AuthorizeSale</b>, <b>Void</b>, <b>Return</b>, <b>DebtPayment</b>, <b>VoidDebtPayment</b> y <b>DebtInquiry</b>.<br/> <span>&#8226;</span><span>&#8226;</span> Los elementos <b>Action</b>, <b>DeviceAction</b> y <b>ExecutedAction</b> del campo <b>Tickets</b> dejan de ser de tipo string y se convierten en arrays.<br/>     <span>&#8226;</span><span>&#8226;</span> Se agrega el elemento <b>AdditionalInformation</b> en las respuestas de todas las operaciones.<br/>     </br> <span>&#8226;</span> <b>Versión 5.3.0</b> Se amplía la definición de la Operación <b>Configure</b> permitiendo tanto en la respuesta como en el requerimiento los elementos <b>Operations</b>, <b>Tables</b> y <b>Files</b></br></br> Se agregan los elementos <b>VoidSupport</b>, <b>ReturnSupport</b>, <b>WalletUseInVoidTransaction</b> y <b>WalletUseInReturnTransaction</b> en las caracteristicas de un Wallet.<br/><br/> Se agrega el Valor <b>Display</b> en el elemento <b>ResponseActions</b> indicando que se debe mostrar en el Display del Dispositivo o Aplicativo el contenido del elemento <b>DisplayResponseMessage</b>.  En la respuesta de la operación  <b>BalanceInquery</b> se agregan los elementos <b>AmountAvailable</b> y <b>PointsAvailable</b> para indicar los saldos.</br> Se especifica en la documentación que el Cancel puede ser usado para Cancelar un Pago con Wallets en Curso.</br></br> Se agregan elementos en los Requerimientos y en las respuestas opcionales entre los POS* que permiten describir las características del punto de venta, los Device* que permiten especificar las características del Dispositivo de Lectura.<br/>   Se cambió el elemento <b>AnswerIdentification</b> por <b>AnswerKey</b>  para compatibilizar con el servicio de Pagos.<br/><br/>     Se agregaron <b>AccountNumber</b>, <b>AccountType</b> y <b>Balance</b> en las operaciones <b>BalanceInquiry</b> y <b>DebtPayment</b> .<br/><br/>     Se agregaron las Operaciones <b>Confirm</b> y  <b>Cancel</b>, donde la Operación <b>Confirm</b> es usada para confirmar un pago recibido por el POS o Aplicativo del comercio. Existen Wallets en los que la confirmación es automática y se indica en el Elemento  <b>AutoConfirm</b> de la respuesta del comando <b>Wallets</b>. La operación <b>Cancel</b> puede ser utilizada a partir de que la Plataforma retorne la acción <b>PaymentFlowIsCancelable</b> en la respuesta de una operación <b>WalletRequest</b>. El Wallet soporta Cancelación de Requerimiento lo cual está indicado con el Elemento <b>SupportRequestCancel</b> dentro de las propiedades de  los Wallets que son retornados por la Operación <b>Wallets</b>.<br/> Se agregó como carasterística de los Wallets también el elemento <b>SupportValidityOfTheRequest</b> que indica si en el primer requerimiento de la Operación <b>WalletRequest</b> se puede enviar el elemento <b>TransactionTimeout</b> que especifica el tiempo de vida de la intención de pago. Superado ese tiempo no se podrá pagar y el ciclo de reintento será detenido por la plataforma, indicado por las siguientes acciones: <b>Completed</b> y <b>Error</b>.<br/> Se agrega el Elemento <b>Tickets</b> en la respuesta de una Operación <b>WalletRequest</b>. El elemento estará presente si como acción está presente el Valor <b>Tickets</b>, indicando que los mismos deberán ser Impresos, capturados digitalmente, etc. según se indique. <br/><br/> Se permite en la Operación <b>PaymentMethod</b> la búsqueda por el Id o el ForeignIdentification<br/><br/> <span>&#8226;</span> <b>Versión 5.2.6</b> Se cambia el nombre del elemento <b>DateTime</b> por <b>TransactionDateTime</b> en la operación <b>WalletRequest</b>.<br/><br/> <span>&#8226;</span> <b>Versión 5.2.5</b> Se agregan en los Planes el atributo <b>POSOrDeviceActions</b> que permite indecarle al Dispositivo que debe solicitar  PIN para esa transacción y eso lo indica enviando la acción <b>RequestPIN</b>.<br/> <span>&#8226;</span><span>&#8226;</span> Se agrega el <b>ResponseActions</b> <b>Configure</b> que indica que se debe ejecutar una reconfiguración para obtener  parámetros nuevos ya que hay alguna actualización. <span>&#8226;</span><span>&#8226;</span> Se agregan las Operaciones <b>Wallets</b>, <b>WalletRequest</b> y <b>EnableService</b>, las mismas pueden formar parte  de un Block y forman parte de la ruptura de Secuencia.<br/><br/> <span>&#8226;</span> <b>Versión 5.2.4</b> Se agrega el identifidor Tributario en <b>OrderInitial</b>, que permite transacciones con <b>Token</b> ( <b>CredentialToken</b> y  <b>CredentialIssuerToken</b> ).<br/> <span>&#8226;</span><span>&#8226;</span> Se completa el <b>GetCardResponse</b> para que contenga los  elementos <b>PaymentMethod</b> y <b>Plans</b>.<br/> <span>&#8226;</span><span>&#8226;</span> Se completa el <b>PaymentMethodResponse</b> para que contenga los elementos <b>PaymentMethod</b> y <b>Plans</b>.<br/> <span>&#8226;</span><span>&#8226;</span> Se agrega en el <b>GetCard</b>: permite forzar un modo de lectura y permite solicitar los datos leídos al POS <b>CardGetMode</b>. <br/><br/> <span>&#8226;</span><span>&#8226;</span> Se permite el envío de datos del cliente <b>Customer*</b> en las operaciones Financieras.<br/><br/> <span>&#8226;</span> <b>Versión 5.2.3</b> Se cambian los valores posibles para <b>ResponseActionCancel</b> en las operaciones <b>GetBlock</b> y <b>GetTransaction</b>.<br/>   <br/> <span>&#8226;</span> <b>Versión 5.2.2</b> Se agrega el Atributo <b>ReasonReverse</b> que permite notificar en las Reversas la razón por la cual fue necesario  generarla, el atributo <b>ReasonSequenceBreak</b> que permite indicar la razón por la cual se produce la ruptura de secuencia que podrá generar una reversa si  fuese necesario, y el atributo </b>Reason</b> en la operación <b>Cancel</b>.<br/>   <br/> <span>&#8226;</span> <b>Versión 5.2.1</b> Se agrega el Atributo <b>IsReverse</b> en todos las operaciones reversables.<br/>   <br/><br/> <br/><br/> <br/><p style=\"color:Blue;\">&copy;2019-2021 EVO Payments Inc. All rights reserved.</p>The EVO Payments name, logo and related trademarks and service marks, owned<br /> by EVO Payments, are registered and/or used in the<br /> United States and many foreign countries. All other trademarks,<br /> service marks and trade names referenced in this site are the property<br /> of their respective owners.<br /> <br /> <br /> ANY USE, COPYING OR REPRODUCTION OF THE TRADEMARKS, LOGOS, INFORMATION,<br />  IMAGES OR DESIGNS CONTAINED IN THIS SITE IS STRICTLY<br />  PROHIBITED WITHOUT THE PRIOR WRITTEN PERMISSION OF EVO Payments Inc.<br /> <br /> 

API version: 5.6.1
Contact: integrations@evopayments.mx
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// BlockCreateObjectBlockCreate struct for BlockCreateObjectBlockCreate
type BlockCreateObjectBlockCreate struct {
	// Identificador de la Compañía para la Plataforma de Integración.
	CompanyIdentification string `json:"CompanyIdentification"`
	// Identificador de Aplicativo para la Plataforma de Integración que usa la Compañía especificada.
	SystemIdentification string `json:"SystemIdentification"`
	// Identificador de Aplicativo para la Plataforma de Integración que usa la Compañía especificada.
	BranchIdentification *string `json:"BranchIdentification,omitempty"`
	// Identificador del Punto de Venta/POS/Caja que pertenece a la sucursal y companía especificada.
	POSIdentification interface{} `json:"POSIdentification,omitempty"`
	// Tipo de Operación que se está requiriendo, solo necesario sobre formatos que no soportan elementos complejos o compuestos.
	RequestType *string `json:"RequestType,omitempty"`
	// Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible.
	ServiceVersion *string `json:"ServiceVersion,omitempty"`
	// Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma.
	Sequence *string `json:"Sequence,omitempty"`
	// Datos asociados a la seguridad de la transacción o de elementos sensibles.
	Security []SaleObjectSaleSecurity `json:"Security,omitempty"`
	// ID que identifica a un grupo de transacciones que serán confirmadas o canceladas
	Block *string `json:"Block,omitempty"`
	// Ticket Digitalizado ( Total o parte del mismo por ejemplo la Firma Digitalizada )    codificado en Base64.
	Ticket *string `json:"Ticket,omitempty"`
	// Identificador Unívoco de la Transacción que se quiere Referenciar a la cual pertenece el Ticket Digitalizado. El valor fue obtenido en el campo o elemento AnswerKey de la Respuesta de la transacción referenciada. Si firma fue capturada previamente y se envía en el requerimiento de la misma Operación Sale, Authorize*, Void, Return, Adjustment, DebtPayment o VoidDebtPayment no es necesario que se envíe este elemento o campo.
	TicketAnswerKey *string `json:"TicketAnswerKey,omitempty"`
	// Identificador del Requermiento obtenido en la respuesta de la operación WalletRequest.            
	RequestKey *string `json:"RequestKey,omitempty"`
	// Motivo por el cual se requiere romper la secuencia.
	ReasonSequenceBreak *string `json:"ReasonSequenceBreak,omitempty"`
	// Tipo de punto de venta.
	POSType *string `json:"POSType,omitempty"`
	// Versión del Aplicativo del punto de Venta.
	POSVersion *string `json:"POSVersion,omitempty"`
	// Dirección IP de la Caja o POS.
	POSAddress *string `json:"POSAddress,omitempty"`
	// Número de serie o identificador unívoco del punto de venta.
	POSSerial *string `json:"POSSerial,omitempty"`
	// Coordenadas Geográficas del aplicativo de punto de venta
	POSGEO *string `json:"POSGEO,omitempty"`
	// Versión del dispositivo.
	ReadingDeviceVersion *string `json:"ReadingDeviceVersion,omitempty"`
	// Tipo de dispositivo utilizado para ingresar los datos de la tarjeta. En función al dispositivo usado, serán realizadas ciertas verificaciones, por lo que ciertos datos serán requeridos. CustomerKeyboard, utilizado para ingreso manual de tarjeta a través de un portal web, por ejemplo - Keyboard, utilizado para ingreso manual de la tarjeta por parte del vendedor - MagneticStripReader, lector de banda de tarjetas por emulación de teclado, u otro valor configurado  que indentifique el dispositivo que se esta utilizando.
	ReadingDeviceType *string `json:"ReadingDeviceType,omitempty"`
	// Indica desde cuando se encuentra operativo o encendido el dispositivo
	ReadingDeviceOperatingFrom *time.Time `json:"ReadingDeviceOperatingFrom,omitempty"`
	// Número de serie o identificador unívoco del dispositivo.
	ReadingDeviceSerial *string `json:"ReadingDeviceSerial,omitempty"`
	// Coordenadas Geograficas del dispositivo de lectura
	ReadingDeviceGEO *string `json:"ReadingDeviceGEO,omitempty"`
	// Dirección IP o MAC Address del dispositivo.
	ReadingDeviceAddress *string `json:"ReadingDeviceAddress,omitempty"`
	// Identificador del usuario que está realizando la Transacción.
	UserID *string `json:"UserID,omitempty"`
	// Nombre del usuario que está realizando la Transacción.
	UserName *string `json:"UserName,omitempty"`
}

// NewBlockCreateObjectBlockCreate instantiates a new BlockCreateObjectBlockCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockCreateObjectBlockCreate(companyIdentification string, systemIdentification string) *BlockCreateObjectBlockCreate {
	this := BlockCreateObjectBlockCreate{}
	this.CompanyIdentification = companyIdentification
	this.SystemIdentification = systemIdentification
	return &this
}

// NewBlockCreateObjectBlockCreateWithDefaults instantiates a new BlockCreateObjectBlockCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockCreateObjectBlockCreateWithDefaults() *BlockCreateObjectBlockCreate {
	this := BlockCreateObjectBlockCreate{}
	return &this
}

// GetCompanyIdentification returns the CompanyIdentification field value
func (o *BlockCreateObjectBlockCreate) GetCompanyIdentification() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyIdentification
}

// GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field value
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetCompanyIdentificationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CompanyIdentification, true
}

// SetCompanyIdentification sets field value
func (o *BlockCreateObjectBlockCreate) SetCompanyIdentification(v string) {
	o.CompanyIdentification = v
}

// GetSystemIdentification returns the SystemIdentification field value
func (o *BlockCreateObjectBlockCreate) GetSystemIdentification() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemIdentification
}

// GetSystemIdentificationOk returns a tuple with the SystemIdentification field value
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetSystemIdentificationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SystemIdentification, true
}

// SetSystemIdentification sets field value
func (o *BlockCreateObjectBlockCreate) SetSystemIdentification(v string) {
	o.SystemIdentification = v
}

// GetBranchIdentification returns the BranchIdentification field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetBranchIdentification() string {
	if o == nil || o.BranchIdentification == nil {
		var ret string
		return ret
	}
	return *o.BranchIdentification
}

// GetBranchIdentificationOk returns a tuple with the BranchIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetBranchIdentificationOk() (*string, bool) {
	if o == nil || o.BranchIdentification == nil {
		return nil, false
	}
	return o.BranchIdentification, true
}

// HasBranchIdentification returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasBranchIdentification() bool {
	if o != nil && o.BranchIdentification != nil {
		return true
	}

	return false
}

// SetBranchIdentification gets a reference to the given string and assigns it to the BranchIdentification field.
func (o *BlockCreateObjectBlockCreate) SetBranchIdentification(v string) {
	o.BranchIdentification = &v
}

// GetPOSIdentification returns the POSIdentification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockCreateObjectBlockCreate) GetPOSIdentification() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.POSIdentification
}

// GetPOSIdentificationOk returns a tuple with the POSIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockCreateObjectBlockCreate) GetPOSIdentificationOk() (*interface{}, bool) {
	if o == nil || o.POSIdentification == nil {
		return nil, false
	}
	return &o.POSIdentification, true
}

// HasPOSIdentification returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasPOSIdentification() bool {
	if o != nil && o.POSIdentification != nil {
		return true
	}

	return false
}

// SetPOSIdentification gets a reference to the given interface{} and assigns it to the POSIdentification field.
func (o *BlockCreateObjectBlockCreate) SetPOSIdentification(v interface{}) {
	o.POSIdentification = v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetRequestType() string {
	if o == nil || o.RequestType == nil {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetRequestTypeOk() (*string, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasRequestType() bool {
	if o != nil && o.RequestType != nil {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *BlockCreateObjectBlockCreate) SetRequestType(v string) {
	o.RequestType = &v
}

// GetServiceVersion returns the ServiceVersion field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetServiceVersion() string {
	if o == nil || o.ServiceVersion == nil {
		var ret string
		return ret
	}
	return *o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetServiceVersionOk() (*string, bool) {
	if o == nil || o.ServiceVersion == nil {
		return nil, false
	}
	return o.ServiceVersion, true
}

// HasServiceVersion returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasServiceVersion() bool {
	if o != nil && o.ServiceVersion != nil {
		return true
	}

	return false
}

// SetServiceVersion gets a reference to the given string and assigns it to the ServiceVersion field.
func (o *BlockCreateObjectBlockCreate) SetServiceVersion(v string) {
	o.ServiceVersion = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetSequence() string {
	if o == nil || o.Sequence == nil {
		var ret string
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetSequenceOk() (*string, bool) {
	if o == nil || o.Sequence == nil {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasSequence() bool {
	if o != nil && o.Sequence != nil {
		return true
	}

	return false
}

// SetSequence gets a reference to the given string and assigns it to the Sequence field.
func (o *BlockCreateObjectBlockCreate) SetSequence(v string) {
	o.Sequence = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetSecurity() []SaleObjectSaleSecurity {
	if o == nil || o.Security == nil {
		var ret []SaleObjectSaleSecurity
		return ret
	}
	return o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetSecurityOk() ([]SaleObjectSaleSecurity, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given []SaleObjectSaleSecurity and assigns it to the Security field.
func (o *BlockCreateObjectBlockCreate) SetSecurity(v []SaleObjectSaleSecurity) {
	o.Security = v
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetBlock() string {
	if o == nil || o.Block == nil {
		var ret string
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetBlockOk() (*string, bool) {
	if o == nil || o.Block == nil {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasBlock() bool {
	if o != nil && o.Block != nil {
		return true
	}

	return false
}

// SetBlock gets a reference to the given string and assigns it to the Block field.
func (o *BlockCreateObjectBlockCreate) SetBlock(v string) {
	o.Block = &v
}

// GetTicket returns the Ticket field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetTicket() string {
	if o == nil || o.Ticket == nil {
		var ret string
		return ret
	}
	return *o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetTicketOk() (*string, bool) {
	if o == nil || o.Ticket == nil {
		return nil, false
	}
	return o.Ticket, true
}

// HasTicket returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasTicket() bool {
	if o != nil && o.Ticket != nil {
		return true
	}

	return false
}

// SetTicket gets a reference to the given string and assigns it to the Ticket field.
func (o *BlockCreateObjectBlockCreate) SetTicket(v string) {
	o.Ticket = &v
}

// GetTicketAnswerKey returns the TicketAnswerKey field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetTicketAnswerKey() string {
	if o == nil || o.TicketAnswerKey == nil {
		var ret string
		return ret
	}
	return *o.TicketAnswerKey
}

// GetTicketAnswerKeyOk returns a tuple with the TicketAnswerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetTicketAnswerKeyOk() (*string, bool) {
	if o == nil || o.TicketAnswerKey == nil {
		return nil, false
	}
	return o.TicketAnswerKey, true
}

// HasTicketAnswerKey returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasTicketAnswerKey() bool {
	if o != nil && o.TicketAnswerKey != nil {
		return true
	}

	return false
}

// SetTicketAnswerKey gets a reference to the given string and assigns it to the TicketAnswerKey field.
func (o *BlockCreateObjectBlockCreate) SetTicketAnswerKey(v string) {
	o.TicketAnswerKey = &v
}

// GetRequestKey returns the RequestKey field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetRequestKey() string {
	if o == nil || o.RequestKey == nil {
		var ret string
		return ret
	}
	return *o.RequestKey
}

// GetRequestKeyOk returns a tuple with the RequestKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetRequestKeyOk() (*string, bool) {
	if o == nil || o.RequestKey == nil {
		return nil, false
	}
	return o.RequestKey, true
}

// HasRequestKey returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasRequestKey() bool {
	if o != nil && o.RequestKey != nil {
		return true
	}

	return false
}

// SetRequestKey gets a reference to the given string and assigns it to the RequestKey field.
func (o *BlockCreateObjectBlockCreate) SetRequestKey(v string) {
	o.RequestKey = &v
}

// GetReasonSequenceBreak returns the ReasonSequenceBreak field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetReasonSequenceBreak() string {
	if o == nil || o.ReasonSequenceBreak == nil {
		var ret string
		return ret
	}
	return *o.ReasonSequenceBreak
}

// GetReasonSequenceBreakOk returns a tuple with the ReasonSequenceBreak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetReasonSequenceBreakOk() (*string, bool) {
	if o == nil || o.ReasonSequenceBreak == nil {
		return nil, false
	}
	return o.ReasonSequenceBreak, true
}

// HasReasonSequenceBreak returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasReasonSequenceBreak() bool {
	if o != nil && o.ReasonSequenceBreak != nil {
		return true
	}

	return false
}

// SetReasonSequenceBreak gets a reference to the given string and assigns it to the ReasonSequenceBreak field.
func (o *BlockCreateObjectBlockCreate) SetReasonSequenceBreak(v string) {
	o.ReasonSequenceBreak = &v
}

// GetPOSType returns the POSType field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetPOSType() string {
	if o == nil || o.POSType == nil {
		var ret string
		return ret
	}
	return *o.POSType
}

// GetPOSTypeOk returns a tuple with the POSType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetPOSTypeOk() (*string, bool) {
	if o == nil || o.POSType == nil {
		return nil, false
	}
	return o.POSType, true
}

// HasPOSType returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasPOSType() bool {
	if o != nil && o.POSType != nil {
		return true
	}

	return false
}

// SetPOSType gets a reference to the given string and assigns it to the POSType field.
func (o *BlockCreateObjectBlockCreate) SetPOSType(v string) {
	o.POSType = &v
}

// GetPOSVersion returns the POSVersion field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetPOSVersion() string {
	if o == nil || o.POSVersion == nil {
		var ret string
		return ret
	}
	return *o.POSVersion
}

// GetPOSVersionOk returns a tuple with the POSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetPOSVersionOk() (*string, bool) {
	if o == nil || o.POSVersion == nil {
		return nil, false
	}
	return o.POSVersion, true
}

// HasPOSVersion returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasPOSVersion() bool {
	if o != nil && o.POSVersion != nil {
		return true
	}

	return false
}

// SetPOSVersion gets a reference to the given string and assigns it to the POSVersion field.
func (o *BlockCreateObjectBlockCreate) SetPOSVersion(v string) {
	o.POSVersion = &v
}

// GetPOSAddress returns the POSAddress field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetPOSAddress() string {
	if o == nil || o.POSAddress == nil {
		var ret string
		return ret
	}
	return *o.POSAddress
}

// GetPOSAddressOk returns a tuple with the POSAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetPOSAddressOk() (*string, bool) {
	if o == nil || o.POSAddress == nil {
		return nil, false
	}
	return o.POSAddress, true
}

// HasPOSAddress returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasPOSAddress() bool {
	if o != nil && o.POSAddress != nil {
		return true
	}

	return false
}

// SetPOSAddress gets a reference to the given string and assigns it to the POSAddress field.
func (o *BlockCreateObjectBlockCreate) SetPOSAddress(v string) {
	o.POSAddress = &v
}

// GetPOSSerial returns the POSSerial field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetPOSSerial() string {
	if o == nil || o.POSSerial == nil {
		var ret string
		return ret
	}
	return *o.POSSerial
}

// GetPOSSerialOk returns a tuple with the POSSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetPOSSerialOk() (*string, bool) {
	if o == nil || o.POSSerial == nil {
		return nil, false
	}
	return o.POSSerial, true
}

// HasPOSSerial returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasPOSSerial() bool {
	if o != nil && o.POSSerial != nil {
		return true
	}

	return false
}

// SetPOSSerial gets a reference to the given string and assigns it to the POSSerial field.
func (o *BlockCreateObjectBlockCreate) SetPOSSerial(v string) {
	o.POSSerial = &v
}

// GetPOSGEO returns the POSGEO field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetPOSGEO() string {
	if o == nil || o.POSGEO == nil {
		var ret string
		return ret
	}
	return *o.POSGEO
}

// GetPOSGEOOk returns a tuple with the POSGEO field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetPOSGEOOk() (*string, bool) {
	if o == nil || o.POSGEO == nil {
		return nil, false
	}
	return o.POSGEO, true
}

// HasPOSGEO returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasPOSGEO() bool {
	if o != nil && o.POSGEO != nil {
		return true
	}

	return false
}

// SetPOSGEO gets a reference to the given string and assigns it to the POSGEO field.
func (o *BlockCreateObjectBlockCreate) SetPOSGEO(v string) {
	o.POSGEO = &v
}

// GetReadingDeviceVersion returns the ReadingDeviceVersion field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetReadingDeviceVersion() string {
	if o == nil || o.ReadingDeviceVersion == nil {
		var ret string
		return ret
	}
	return *o.ReadingDeviceVersion
}

// GetReadingDeviceVersionOk returns a tuple with the ReadingDeviceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetReadingDeviceVersionOk() (*string, bool) {
	if o == nil || o.ReadingDeviceVersion == nil {
		return nil, false
	}
	return o.ReadingDeviceVersion, true
}

// HasReadingDeviceVersion returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasReadingDeviceVersion() bool {
	if o != nil && o.ReadingDeviceVersion != nil {
		return true
	}

	return false
}

// SetReadingDeviceVersion gets a reference to the given string and assigns it to the ReadingDeviceVersion field.
func (o *BlockCreateObjectBlockCreate) SetReadingDeviceVersion(v string) {
	o.ReadingDeviceVersion = &v
}

// GetReadingDeviceType returns the ReadingDeviceType field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetReadingDeviceType() string {
	if o == nil || o.ReadingDeviceType == nil {
		var ret string
		return ret
	}
	return *o.ReadingDeviceType
}

// GetReadingDeviceTypeOk returns a tuple with the ReadingDeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetReadingDeviceTypeOk() (*string, bool) {
	if o == nil || o.ReadingDeviceType == nil {
		return nil, false
	}
	return o.ReadingDeviceType, true
}

// HasReadingDeviceType returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasReadingDeviceType() bool {
	if o != nil && o.ReadingDeviceType != nil {
		return true
	}

	return false
}

// SetReadingDeviceType gets a reference to the given string and assigns it to the ReadingDeviceType field.
func (o *BlockCreateObjectBlockCreate) SetReadingDeviceType(v string) {
	o.ReadingDeviceType = &v
}

// GetReadingDeviceOperatingFrom returns the ReadingDeviceOperatingFrom field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetReadingDeviceOperatingFrom() time.Time {
	if o == nil || o.ReadingDeviceOperatingFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.ReadingDeviceOperatingFrom
}

// GetReadingDeviceOperatingFromOk returns a tuple with the ReadingDeviceOperatingFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetReadingDeviceOperatingFromOk() (*time.Time, bool) {
	if o == nil || o.ReadingDeviceOperatingFrom == nil {
		return nil, false
	}
	return o.ReadingDeviceOperatingFrom, true
}

// HasReadingDeviceOperatingFrom returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasReadingDeviceOperatingFrom() bool {
	if o != nil && o.ReadingDeviceOperatingFrom != nil {
		return true
	}

	return false
}

// SetReadingDeviceOperatingFrom gets a reference to the given time.Time and assigns it to the ReadingDeviceOperatingFrom field.
func (o *BlockCreateObjectBlockCreate) SetReadingDeviceOperatingFrom(v time.Time) {
	o.ReadingDeviceOperatingFrom = &v
}

// GetReadingDeviceSerial returns the ReadingDeviceSerial field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetReadingDeviceSerial() string {
	if o == nil || o.ReadingDeviceSerial == nil {
		var ret string
		return ret
	}
	return *o.ReadingDeviceSerial
}

// GetReadingDeviceSerialOk returns a tuple with the ReadingDeviceSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetReadingDeviceSerialOk() (*string, bool) {
	if o == nil || o.ReadingDeviceSerial == nil {
		return nil, false
	}
	return o.ReadingDeviceSerial, true
}

// HasReadingDeviceSerial returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasReadingDeviceSerial() bool {
	if o != nil && o.ReadingDeviceSerial != nil {
		return true
	}

	return false
}

// SetReadingDeviceSerial gets a reference to the given string and assigns it to the ReadingDeviceSerial field.
func (o *BlockCreateObjectBlockCreate) SetReadingDeviceSerial(v string) {
	o.ReadingDeviceSerial = &v
}

// GetReadingDeviceGEO returns the ReadingDeviceGEO field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetReadingDeviceGEO() string {
	if o == nil || o.ReadingDeviceGEO == nil {
		var ret string
		return ret
	}
	return *o.ReadingDeviceGEO
}

// GetReadingDeviceGEOOk returns a tuple with the ReadingDeviceGEO field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetReadingDeviceGEOOk() (*string, bool) {
	if o == nil || o.ReadingDeviceGEO == nil {
		return nil, false
	}
	return o.ReadingDeviceGEO, true
}

// HasReadingDeviceGEO returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasReadingDeviceGEO() bool {
	if o != nil && o.ReadingDeviceGEO != nil {
		return true
	}

	return false
}

// SetReadingDeviceGEO gets a reference to the given string and assigns it to the ReadingDeviceGEO field.
func (o *BlockCreateObjectBlockCreate) SetReadingDeviceGEO(v string) {
	o.ReadingDeviceGEO = &v
}

// GetReadingDeviceAddress returns the ReadingDeviceAddress field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetReadingDeviceAddress() string {
	if o == nil || o.ReadingDeviceAddress == nil {
		var ret string
		return ret
	}
	return *o.ReadingDeviceAddress
}

// GetReadingDeviceAddressOk returns a tuple with the ReadingDeviceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetReadingDeviceAddressOk() (*string, bool) {
	if o == nil || o.ReadingDeviceAddress == nil {
		return nil, false
	}
	return o.ReadingDeviceAddress, true
}

// HasReadingDeviceAddress returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasReadingDeviceAddress() bool {
	if o != nil && o.ReadingDeviceAddress != nil {
		return true
	}

	return false
}

// SetReadingDeviceAddress gets a reference to the given string and assigns it to the ReadingDeviceAddress field.
func (o *BlockCreateObjectBlockCreate) SetReadingDeviceAddress(v string) {
	o.ReadingDeviceAddress = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetUserID() string {
	if o == nil || o.UserID == nil {
		var ret string
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetUserIDOk() (*string, bool) {
	if o == nil || o.UserID == nil {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasUserID() bool {
	if o != nil && o.UserID != nil {
		return true
	}

	return false
}

// SetUserID gets a reference to the given string and assigns it to the UserID field.
func (o *BlockCreateObjectBlockCreate) SetUserID(v string) {
	o.UserID = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *BlockCreateObjectBlockCreate) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockCreateObjectBlockCreate) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *BlockCreateObjectBlockCreate) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *BlockCreateObjectBlockCreate) SetUserName(v string) {
	o.UserName = &v
}

func (o BlockCreateObjectBlockCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["CompanyIdentification"] = o.CompanyIdentification
	}
	if true {
		toSerialize["SystemIdentification"] = o.SystemIdentification
	}
	if o.BranchIdentification != nil {
		toSerialize["BranchIdentification"] = o.BranchIdentification
	}
	if o.POSIdentification != nil {
		toSerialize["POSIdentification"] = o.POSIdentification
	}
	if o.RequestType != nil {
		toSerialize["RequestType"] = o.RequestType
	}
	if o.ServiceVersion != nil {
		toSerialize["ServiceVersion"] = o.ServiceVersion
	}
	if o.Sequence != nil {
		toSerialize["Sequence"] = o.Sequence
	}
	if o.Security != nil {
		toSerialize["Security"] = o.Security
	}
	if o.Block != nil {
		toSerialize["Block"] = o.Block
	}
	if o.Ticket != nil {
		toSerialize["Ticket"] = o.Ticket
	}
	if o.TicketAnswerKey != nil {
		toSerialize["TicketAnswerKey"] = o.TicketAnswerKey
	}
	if o.RequestKey != nil {
		toSerialize["RequestKey"] = o.RequestKey
	}
	if o.ReasonSequenceBreak != nil {
		toSerialize["ReasonSequenceBreak"] = o.ReasonSequenceBreak
	}
	if o.POSType != nil {
		toSerialize["POSType"] = o.POSType
	}
	if o.POSVersion != nil {
		toSerialize["POSVersion"] = o.POSVersion
	}
	if o.POSAddress != nil {
		toSerialize["POSAddress"] = o.POSAddress
	}
	if o.POSSerial != nil {
		toSerialize["POSSerial"] = o.POSSerial
	}
	if o.POSGEO != nil {
		toSerialize["POSGEO"] = o.POSGEO
	}
	if o.ReadingDeviceVersion != nil {
		toSerialize["ReadingDeviceVersion"] = o.ReadingDeviceVersion
	}
	if o.ReadingDeviceType != nil {
		toSerialize["ReadingDeviceType"] = o.ReadingDeviceType
	}
	if o.ReadingDeviceOperatingFrom != nil {
		toSerialize["ReadingDeviceOperatingFrom"] = o.ReadingDeviceOperatingFrom
	}
	if o.ReadingDeviceSerial != nil {
		toSerialize["ReadingDeviceSerial"] = o.ReadingDeviceSerial
	}
	if o.ReadingDeviceGEO != nil {
		toSerialize["ReadingDeviceGEO"] = o.ReadingDeviceGEO
	}
	if o.ReadingDeviceAddress != nil {
		toSerialize["ReadingDeviceAddress"] = o.ReadingDeviceAddress
	}
	if o.UserID != nil {
		toSerialize["UserID"] = o.UserID
	}
	if o.UserName != nil {
		toSerialize["UserName"] = o.UserName
	}
	return json.Marshal(toSerialize)
}

type NullableBlockCreateObjectBlockCreate struct {
	value *BlockCreateObjectBlockCreate
	isSet bool
}

func (v NullableBlockCreateObjectBlockCreate) Get() *BlockCreateObjectBlockCreate {
	return v.value
}

func (v *NullableBlockCreateObjectBlockCreate) Set(val *BlockCreateObjectBlockCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockCreateObjectBlockCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockCreateObjectBlockCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockCreateObjectBlockCreate(val *BlockCreateObjectBlockCreate) *NullableBlockCreateObjectBlockCreate {
	return &NullableBlockCreateObjectBlockCreate{value: val, isSet: true}
}

func (v NullableBlockCreateObjectBlockCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockCreateObjectBlockCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


