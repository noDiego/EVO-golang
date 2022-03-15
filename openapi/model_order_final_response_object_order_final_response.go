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
)

// OrderFinalResponseObjectOrderFinalResponse struct for OrderFinalResponseObjectOrderFinalResponse
type OrderFinalResponseObjectOrderFinalResponse struct {
	// Tipo de Operación que se requirió, contendrá el mismo valor que se recibió en el requerimiento, sobre formatos que no soportan elementos complejos o compuestos.
	RequestType *string `json:"RequestType,omitempty"`
	// Acciones a realizar por parte del POS y/o PINPAD en base al resultado de la operación que ha sido procesada. Cada uno de estos actions o acciones están concatenadas por comas. Los posibles actions son OK, Approve, Refuse, IssuerCall, Tickets, WithHold, GetCard, UseTerminalToAuthorize, ConfigurationError, SystemError, ResourceError, ProcessError, Completed.
	ResponseActions []string `json:"ResponseActions"`
	// Descripción del resultado del proceso del requerimiento recibido. Esta descripción es generada por la plataforma, no por el Host que termine resolviendo la transacción.
	ResponseMessage string `json:"ResponseMessage"`
	// ID que identifica el sistema desde donde proviene la petición.
	SystemIdentification *string `json:"SystemIdentification,omitempty"`
	// ID que identifica la companía desde donde proviene la petición.
	CompanyIdentification *string `json:"CompanyIdentification,omitempty"`
	// ID que identifica la sucursal desde donde proviene la petición. Esta sucursal pertenece a una determinada companía.
	BranchIdentification *string `json:"BranchIdentification,omitempty"`
	// ID que identifica la caja o punto de venta desde donde proviene la petición. Este punto de venta pertenece a una determinada sucursal y companía.
	POSIdentification *string `json:"POSIdentification,omitempty"`
	// Código de respuesta para el sistema externo, es decir, para la aplicación cliente que se comunica con el TEF.
	ForeignResponseCode *string `json:"ForeignResponseCode,omitempty"`
	// Código de Respuesta Interno de la plataforma, el POS debe actuar por lo que indican las acciones especificadas en ResponseActions y no por el código de respuesta informado en este campo o elemento, pero es una buena práctica que sea persistido por el mismo.
	ResponseCode int32 `json:"ResponseCode"`
	// Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible.
	ServiceVersion *string `json:"ServiceVersion,omitempty"`
	// Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma.
	Sequence *string `json:"Sequence,omitempty"`
	// Datos asociados a la seguridad de la transacción o de elementos sensibles.
	Security []SaleObjectSaleSecurity `json:"Security,omitempty"`
	// En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente.
	AdditionalInformation []SaleResponseObjectSaleResponseAdditionalInformation `json:"AdditionalInformation,omitempty"`
	// En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente.
	RequiredInformation []DebtPaymentObjectDebtPaymentRequiredInformation `json:"RequiredInformation,omitempty"`
	// ID de la operación a realizar, generado por el sistema externo
	TransactionIdentification *string `json:"TransactionIdentification,omitempty"`
	// Descripción del tipo de operación que se realizará
	TransactionDescription *string `json:"TransactionDescription,omitempty"`
	// .
	PaymentTransactionIdentification *string `json:"PaymentTransactionIdentification,omitempty"`
	// .
	HostResponseCode *string `json:"HostResponseCode,omitempty"`
	// .
	HostResponseMessage *string `json:"HostResponseMessage,omitempty"`
	// .
	HostResponseAction *string `json:"HostResponseAction,omitempty"`
	// .
	HostMerchantIdentification *string `json:"HostMerchantIdentification,omitempty"`
	// .
	HostTerminalIdentification *string `json:"HostTerminalIdentification,omitempty"`
	// .
	HostBatchNumber *string `json:"HostBatchNumber,omitempty"`
	// .
	HostTicketNumber *string `json:"HostTicketNumber,omitempty"`
	// .
	HostTraceNumber *string `json:"HostTraceNumber,omitempty"`
	// .
	HostRetrievalReferenceNumber *string `json:"HostRetrievalReferenceNumber,omitempty"`
	// .
	HostTransactionIdentification *string `json:"HostTransactionIdentification,omitempty"`
	// .
	HostAuthorizationCode *string `json:"HostAuthorizationCode,omitempty"`
	// Token asociado a la Credencial Enrolada
	CredentialToken *string `json:"CredentialToken,omitempty"`
	// Emisor del Token asociado a la credencial enrolada
	CredentialIssuerToken *string `json:"CredentialIssuerToken,omitempty"`
	// .
	PlanID *string `json:"PlanID,omitempty"`
	// .
	PaymentMethodId *string `json:"PaymentMethodId,omitempty"`
	// .
	HostId *string `json:"HostId,omitempty"`
	// Identificador de la transacción, utilizado solo por algunos hosts para realizar anulaciones y devoluciones
	AcquirerReferenceData *string `json:"AcquirerReferenceData,omitempty"`
	// Identificador de la transacción generado por Plataforma para ser enviado al Adquirente
	IdentifierForTheAcquirer *string `json:"IdentifierForTheAcquirer,omitempty"`
	Configuration *SaleResponseObjectSaleResponseConfiguration `json:"Configuration,omitempty"`
}

// NewOrderFinalResponseObjectOrderFinalResponse instantiates a new OrderFinalResponseObjectOrderFinalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderFinalResponseObjectOrderFinalResponse(responseActions []string, responseMessage string, responseCode int32) *OrderFinalResponseObjectOrderFinalResponse {
	this := OrderFinalResponseObjectOrderFinalResponse{}
	this.ResponseActions = responseActions
	this.ResponseMessage = responseMessage
	this.ResponseCode = responseCode
	return &this
}

// NewOrderFinalResponseObjectOrderFinalResponseWithDefaults instantiates a new OrderFinalResponseObjectOrderFinalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderFinalResponseObjectOrderFinalResponseWithDefaults() *OrderFinalResponseObjectOrderFinalResponse {
	this := OrderFinalResponseObjectOrderFinalResponse{}
	return &this
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetRequestType() string {
	if o == nil || o.RequestType == nil {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetRequestTypeOk() (*string, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasRequestType() bool {
	if o != nil && o.RequestType != nil {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetRequestType(v string) {
	o.RequestType = &v
}

// GetResponseActions returns the ResponseActions field value
func (o *OrderFinalResponseObjectOrderFinalResponse) GetResponseActions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResponseActions
}

// GetResponseActionsOk returns a tuple with the ResponseActions field value
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetResponseActionsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseActions, true
}

// SetResponseActions sets field value
func (o *OrderFinalResponseObjectOrderFinalResponse) SetResponseActions(v []string) {
	o.ResponseActions = v
}

// GetResponseMessage returns the ResponseMessage field value
func (o *OrderFinalResponseObjectOrderFinalResponse) GetResponseMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetResponseMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResponseMessage, true
}

// SetResponseMessage sets field value
func (o *OrderFinalResponseObjectOrderFinalResponse) SetResponseMessage(v string) {
	o.ResponseMessage = v
}

// GetSystemIdentification returns the SystemIdentification field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetSystemIdentification() string {
	if o == nil || o.SystemIdentification == nil {
		var ret string
		return ret
	}
	return *o.SystemIdentification
}

// GetSystemIdentificationOk returns a tuple with the SystemIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetSystemIdentificationOk() (*string, bool) {
	if o == nil || o.SystemIdentification == nil {
		return nil, false
	}
	return o.SystemIdentification, true
}

// HasSystemIdentification returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasSystemIdentification() bool {
	if o != nil && o.SystemIdentification != nil {
		return true
	}

	return false
}

// SetSystemIdentification gets a reference to the given string and assigns it to the SystemIdentification field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetSystemIdentification(v string) {
	o.SystemIdentification = &v
}

// GetCompanyIdentification returns the CompanyIdentification field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetCompanyIdentification() string {
	if o == nil || o.CompanyIdentification == nil {
		var ret string
		return ret
	}
	return *o.CompanyIdentification
}

// GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetCompanyIdentificationOk() (*string, bool) {
	if o == nil || o.CompanyIdentification == nil {
		return nil, false
	}
	return o.CompanyIdentification, true
}

// HasCompanyIdentification returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasCompanyIdentification() bool {
	if o != nil && o.CompanyIdentification != nil {
		return true
	}

	return false
}

// SetCompanyIdentification gets a reference to the given string and assigns it to the CompanyIdentification field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetCompanyIdentification(v string) {
	o.CompanyIdentification = &v
}

// GetBranchIdentification returns the BranchIdentification field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetBranchIdentification() string {
	if o == nil || o.BranchIdentification == nil {
		var ret string
		return ret
	}
	return *o.BranchIdentification
}

// GetBranchIdentificationOk returns a tuple with the BranchIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetBranchIdentificationOk() (*string, bool) {
	if o == nil || o.BranchIdentification == nil {
		return nil, false
	}
	return o.BranchIdentification, true
}

// HasBranchIdentification returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasBranchIdentification() bool {
	if o != nil && o.BranchIdentification != nil {
		return true
	}

	return false
}

// SetBranchIdentification gets a reference to the given string and assigns it to the BranchIdentification field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetBranchIdentification(v string) {
	o.BranchIdentification = &v
}

// GetPOSIdentification returns the POSIdentification field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetPOSIdentification() string {
	if o == nil || o.POSIdentification == nil {
		var ret string
		return ret
	}
	return *o.POSIdentification
}

// GetPOSIdentificationOk returns a tuple with the POSIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetPOSIdentificationOk() (*string, bool) {
	if o == nil || o.POSIdentification == nil {
		return nil, false
	}
	return o.POSIdentification, true
}

// HasPOSIdentification returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasPOSIdentification() bool {
	if o != nil && o.POSIdentification != nil {
		return true
	}

	return false
}

// SetPOSIdentification gets a reference to the given string and assigns it to the POSIdentification field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetPOSIdentification(v string) {
	o.POSIdentification = &v
}

// GetForeignResponseCode returns the ForeignResponseCode field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetForeignResponseCode() string {
	if o == nil || o.ForeignResponseCode == nil {
		var ret string
		return ret
	}
	return *o.ForeignResponseCode
}

// GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetForeignResponseCodeOk() (*string, bool) {
	if o == nil || o.ForeignResponseCode == nil {
		return nil, false
	}
	return o.ForeignResponseCode, true
}

// HasForeignResponseCode returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasForeignResponseCode() bool {
	if o != nil && o.ForeignResponseCode != nil {
		return true
	}

	return false
}

// SetForeignResponseCode gets a reference to the given string and assigns it to the ForeignResponseCode field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetForeignResponseCode(v string) {
	o.ForeignResponseCode = &v
}

// GetResponseCode returns the ResponseCode field value
func (o *OrderFinalResponseObjectOrderFinalResponse) GetResponseCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetResponseCodeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResponseCode, true
}

// SetResponseCode sets field value
func (o *OrderFinalResponseObjectOrderFinalResponse) SetResponseCode(v int32) {
	o.ResponseCode = v
}

// GetServiceVersion returns the ServiceVersion field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetServiceVersion() string {
	if o == nil || o.ServiceVersion == nil {
		var ret string
		return ret
	}
	return *o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetServiceVersionOk() (*string, bool) {
	if o == nil || o.ServiceVersion == nil {
		return nil, false
	}
	return o.ServiceVersion, true
}

// HasServiceVersion returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasServiceVersion() bool {
	if o != nil && o.ServiceVersion != nil {
		return true
	}

	return false
}

// SetServiceVersion gets a reference to the given string and assigns it to the ServiceVersion field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetServiceVersion(v string) {
	o.ServiceVersion = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetSequence() string {
	if o == nil || o.Sequence == nil {
		var ret string
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetSequenceOk() (*string, bool) {
	if o == nil || o.Sequence == nil {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasSequence() bool {
	if o != nil && o.Sequence != nil {
		return true
	}

	return false
}

// SetSequence gets a reference to the given string and assigns it to the Sequence field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetSequence(v string) {
	o.Sequence = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetSecurity() []SaleObjectSaleSecurity {
	if o == nil || o.Security == nil {
		var ret []SaleObjectSaleSecurity
		return ret
	}
	return o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetSecurityOk() ([]SaleObjectSaleSecurity, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given []SaleObjectSaleSecurity and assigns it to the Security field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetSecurity(v []SaleObjectSaleSecurity) {
	o.Security = v
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation {
	if o == nil || o.AdditionalInformation == nil {
		var ret []SaleResponseObjectSaleResponseAdditionalInformation
		return ret
	}
	return o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetAdditionalInformationOk() ([]SaleResponseObjectSaleResponseAdditionalInformation, bool) {
	if o == nil || o.AdditionalInformation == nil {
		return nil, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasAdditionalInformation() bool {
	if o != nil && o.AdditionalInformation != nil {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given []SaleResponseObjectSaleResponseAdditionalInformation and assigns it to the AdditionalInformation field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation) {
	o.AdditionalInformation = v
}

// GetRequiredInformation returns the RequiredInformation field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetRequiredInformation() []DebtPaymentObjectDebtPaymentRequiredInformation {
	if o == nil || o.RequiredInformation == nil {
		var ret []DebtPaymentObjectDebtPaymentRequiredInformation
		return ret
	}
	return o.RequiredInformation
}

// GetRequiredInformationOk returns a tuple with the RequiredInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetRequiredInformationOk() ([]DebtPaymentObjectDebtPaymentRequiredInformation, bool) {
	if o == nil || o.RequiredInformation == nil {
		return nil, false
	}
	return o.RequiredInformation, true
}

// HasRequiredInformation returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasRequiredInformation() bool {
	if o != nil && o.RequiredInformation != nil {
		return true
	}

	return false
}

// SetRequiredInformation gets a reference to the given []DebtPaymentObjectDebtPaymentRequiredInformation and assigns it to the RequiredInformation field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetRequiredInformation(v []DebtPaymentObjectDebtPaymentRequiredInformation) {
	o.RequiredInformation = v
}

// GetTransactionIdentification returns the TransactionIdentification field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetTransactionIdentification() string {
	if o == nil || o.TransactionIdentification == nil {
		var ret string
		return ret
	}
	return *o.TransactionIdentification
}

// GetTransactionIdentificationOk returns a tuple with the TransactionIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetTransactionIdentificationOk() (*string, bool) {
	if o == nil || o.TransactionIdentification == nil {
		return nil, false
	}
	return o.TransactionIdentification, true
}

// HasTransactionIdentification returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasTransactionIdentification() bool {
	if o != nil && o.TransactionIdentification != nil {
		return true
	}

	return false
}

// SetTransactionIdentification gets a reference to the given string and assigns it to the TransactionIdentification field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetTransactionIdentification(v string) {
	o.TransactionIdentification = &v
}

// GetTransactionDescription returns the TransactionDescription field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetTransactionDescription() string {
	if o == nil || o.TransactionDescription == nil {
		var ret string
		return ret
	}
	return *o.TransactionDescription
}

// GetTransactionDescriptionOk returns a tuple with the TransactionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetTransactionDescriptionOk() (*string, bool) {
	if o == nil || o.TransactionDescription == nil {
		return nil, false
	}
	return o.TransactionDescription, true
}

// HasTransactionDescription returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasTransactionDescription() bool {
	if o != nil && o.TransactionDescription != nil {
		return true
	}

	return false
}

// SetTransactionDescription gets a reference to the given string and assigns it to the TransactionDescription field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetTransactionDescription(v string) {
	o.TransactionDescription = &v
}

// GetPaymentTransactionIdentification returns the PaymentTransactionIdentification field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetPaymentTransactionIdentification() string {
	if o == nil || o.PaymentTransactionIdentification == nil {
		var ret string
		return ret
	}
	return *o.PaymentTransactionIdentification
}

// GetPaymentTransactionIdentificationOk returns a tuple with the PaymentTransactionIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetPaymentTransactionIdentificationOk() (*string, bool) {
	if o == nil || o.PaymentTransactionIdentification == nil {
		return nil, false
	}
	return o.PaymentTransactionIdentification, true
}

// HasPaymentTransactionIdentification returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasPaymentTransactionIdentification() bool {
	if o != nil && o.PaymentTransactionIdentification != nil {
		return true
	}

	return false
}

// SetPaymentTransactionIdentification gets a reference to the given string and assigns it to the PaymentTransactionIdentification field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetPaymentTransactionIdentification(v string) {
	o.PaymentTransactionIdentification = &v
}

// GetHostResponseCode returns the HostResponseCode field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostResponseCode() string {
	if o == nil || o.HostResponseCode == nil {
		var ret string
		return ret
	}
	return *o.HostResponseCode
}

// GetHostResponseCodeOk returns a tuple with the HostResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostResponseCodeOk() (*string, bool) {
	if o == nil || o.HostResponseCode == nil {
		return nil, false
	}
	return o.HostResponseCode, true
}

// HasHostResponseCode returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostResponseCode() bool {
	if o != nil && o.HostResponseCode != nil {
		return true
	}

	return false
}

// SetHostResponseCode gets a reference to the given string and assigns it to the HostResponseCode field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostResponseCode(v string) {
	o.HostResponseCode = &v
}

// GetHostResponseMessage returns the HostResponseMessage field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostResponseMessage() string {
	if o == nil || o.HostResponseMessage == nil {
		var ret string
		return ret
	}
	return *o.HostResponseMessage
}

// GetHostResponseMessageOk returns a tuple with the HostResponseMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostResponseMessageOk() (*string, bool) {
	if o == nil || o.HostResponseMessage == nil {
		return nil, false
	}
	return o.HostResponseMessage, true
}

// HasHostResponseMessage returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostResponseMessage() bool {
	if o != nil && o.HostResponseMessage != nil {
		return true
	}

	return false
}

// SetHostResponseMessage gets a reference to the given string and assigns it to the HostResponseMessage field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostResponseMessage(v string) {
	o.HostResponseMessage = &v
}

// GetHostResponseAction returns the HostResponseAction field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostResponseAction() string {
	if o == nil || o.HostResponseAction == nil {
		var ret string
		return ret
	}
	return *o.HostResponseAction
}

// GetHostResponseActionOk returns a tuple with the HostResponseAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostResponseActionOk() (*string, bool) {
	if o == nil || o.HostResponseAction == nil {
		return nil, false
	}
	return o.HostResponseAction, true
}

// HasHostResponseAction returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostResponseAction() bool {
	if o != nil && o.HostResponseAction != nil {
		return true
	}

	return false
}

// SetHostResponseAction gets a reference to the given string and assigns it to the HostResponseAction field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostResponseAction(v string) {
	o.HostResponseAction = &v
}

// GetHostMerchantIdentification returns the HostMerchantIdentification field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostMerchantIdentification() string {
	if o == nil || o.HostMerchantIdentification == nil {
		var ret string
		return ret
	}
	return *o.HostMerchantIdentification
}

// GetHostMerchantIdentificationOk returns a tuple with the HostMerchantIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostMerchantIdentificationOk() (*string, bool) {
	if o == nil || o.HostMerchantIdentification == nil {
		return nil, false
	}
	return o.HostMerchantIdentification, true
}

// HasHostMerchantIdentification returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostMerchantIdentification() bool {
	if o != nil && o.HostMerchantIdentification != nil {
		return true
	}

	return false
}

// SetHostMerchantIdentification gets a reference to the given string and assigns it to the HostMerchantIdentification field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostMerchantIdentification(v string) {
	o.HostMerchantIdentification = &v
}

// GetHostTerminalIdentification returns the HostTerminalIdentification field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTerminalIdentification() string {
	if o == nil || o.HostTerminalIdentification == nil {
		var ret string
		return ret
	}
	return *o.HostTerminalIdentification
}

// GetHostTerminalIdentificationOk returns a tuple with the HostTerminalIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTerminalIdentificationOk() (*string, bool) {
	if o == nil || o.HostTerminalIdentification == nil {
		return nil, false
	}
	return o.HostTerminalIdentification, true
}

// HasHostTerminalIdentification returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostTerminalIdentification() bool {
	if o != nil && o.HostTerminalIdentification != nil {
		return true
	}

	return false
}

// SetHostTerminalIdentification gets a reference to the given string and assigns it to the HostTerminalIdentification field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostTerminalIdentification(v string) {
	o.HostTerminalIdentification = &v
}

// GetHostBatchNumber returns the HostBatchNumber field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostBatchNumber() string {
	if o == nil || o.HostBatchNumber == nil {
		var ret string
		return ret
	}
	return *o.HostBatchNumber
}

// GetHostBatchNumberOk returns a tuple with the HostBatchNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostBatchNumberOk() (*string, bool) {
	if o == nil || o.HostBatchNumber == nil {
		return nil, false
	}
	return o.HostBatchNumber, true
}

// HasHostBatchNumber returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostBatchNumber() bool {
	if o != nil && o.HostBatchNumber != nil {
		return true
	}

	return false
}

// SetHostBatchNumber gets a reference to the given string and assigns it to the HostBatchNumber field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostBatchNumber(v string) {
	o.HostBatchNumber = &v
}

// GetHostTicketNumber returns the HostTicketNumber field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTicketNumber() string {
	if o == nil || o.HostTicketNumber == nil {
		var ret string
		return ret
	}
	return *o.HostTicketNumber
}

// GetHostTicketNumberOk returns a tuple with the HostTicketNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTicketNumberOk() (*string, bool) {
	if o == nil || o.HostTicketNumber == nil {
		return nil, false
	}
	return o.HostTicketNumber, true
}

// HasHostTicketNumber returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostTicketNumber() bool {
	if o != nil && o.HostTicketNumber != nil {
		return true
	}

	return false
}

// SetHostTicketNumber gets a reference to the given string and assigns it to the HostTicketNumber field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostTicketNumber(v string) {
	o.HostTicketNumber = &v
}

// GetHostTraceNumber returns the HostTraceNumber field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTraceNumber() string {
	if o == nil || o.HostTraceNumber == nil {
		var ret string
		return ret
	}
	return *o.HostTraceNumber
}

// GetHostTraceNumberOk returns a tuple with the HostTraceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTraceNumberOk() (*string, bool) {
	if o == nil || o.HostTraceNumber == nil {
		return nil, false
	}
	return o.HostTraceNumber, true
}

// HasHostTraceNumber returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostTraceNumber() bool {
	if o != nil && o.HostTraceNumber != nil {
		return true
	}

	return false
}

// SetHostTraceNumber gets a reference to the given string and assigns it to the HostTraceNumber field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostTraceNumber(v string) {
	o.HostTraceNumber = &v
}

// GetHostRetrievalReferenceNumber returns the HostRetrievalReferenceNumber field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostRetrievalReferenceNumber() string {
	if o == nil || o.HostRetrievalReferenceNumber == nil {
		var ret string
		return ret
	}
	return *o.HostRetrievalReferenceNumber
}

// GetHostRetrievalReferenceNumberOk returns a tuple with the HostRetrievalReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostRetrievalReferenceNumberOk() (*string, bool) {
	if o == nil || o.HostRetrievalReferenceNumber == nil {
		return nil, false
	}
	return o.HostRetrievalReferenceNumber, true
}

// HasHostRetrievalReferenceNumber returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostRetrievalReferenceNumber() bool {
	if o != nil && o.HostRetrievalReferenceNumber != nil {
		return true
	}

	return false
}

// SetHostRetrievalReferenceNumber gets a reference to the given string and assigns it to the HostRetrievalReferenceNumber field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostRetrievalReferenceNumber(v string) {
	o.HostRetrievalReferenceNumber = &v
}

// GetHostTransactionIdentification returns the HostTransactionIdentification field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTransactionIdentification() string {
	if o == nil || o.HostTransactionIdentification == nil {
		var ret string
		return ret
	}
	return *o.HostTransactionIdentification
}

// GetHostTransactionIdentificationOk returns a tuple with the HostTransactionIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostTransactionIdentificationOk() (*string, bool) {
	if o == nil || o.HostTransactionIdentification == nil {
		return nil, false
	}
	return o.HostTransactionIdentification, true
}

// HasHostTransactionIdentification returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostTransactionIdentification() bool {
	if o != nil && o.HostTransactionIdentification != nil {
		return true
	}

	return false
}

// SetHostTransactionIdentification gets a reference to the given string and assigns it to the HostTransactionIdentification field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostTransactionIdentification(v string) {
	o.HostTransactionIdentification = &v
}

// GetHostAuthorizationCode returns the HostAuthorizationCode field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostAuthorizationCode() string {
	if o == nil || o.HostAuthorizationCode == nil {
		var ret string
		return ret
	}
	return *o.HostAuthorizationCode
}

// GetHostAuthorizationCodeOk returns a tuple with the HostAuthorizationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostAuthorizationCodeOk() (*string, bool) {
	if o == nil || o.HostAuthorizationCode == nil {
		return nil, false
	}
	return o.HostAuthorizationCode, true
}

// HasHostAuthorizationCode returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostAuthorizationCode() bool {
	if o != nil && o.HostAuthorizationCode != nil {
		return true
	}

	return false
}

// SetHostAuthorizationCode gets a reference to the given string and assigns it to the HostAuthorizationCode field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostAuthorizationCode(v string) {
	o.HostAuthorizationCode = &v
}

// GetCredentialToken returns the CredentialToken field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetCredentialToken() string {
	if o == nil || o.CredentialToken == nil {
		var ret string
		return ret
	}
	return *o.CredentialToken
}

// GetCredentialTokenOk returns a tuple with the CredentialToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetCredentialTokenOk() (*string, bool) {
	if o == nil || o.CredentialToken == nil {
		return nil, false
	}
	return o.CredentialToken, true
}

// HasCredentialToken returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasCredentialToken() bool {
	if o != nil && o.CredentialToken != nil {
		return true
	}

	return false
}

// SetCredentialToken gets a reference to the given string and assigns it to the CredentialToken field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetCredentialToken(v string) {
	o.CredentialToken = &v
}

// GetCredentialIssuerToken returns the CredentialIssuerToken field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetCredentialIssuerToken() string {
	if o == nil || o.CredentialIssuerToken == nil {
		var ret string
		return ret
	}
	return *o.CredentialIssuerToken
}

// GetCredentialIssuerTokenOk returns a tuple with the CredentialIssuerToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetCredentialIssuerTokenOk() (*string, bool) {
	if o == nil || o.CredentialIssuerToken == nil {
		return nil, false
	}
	return o.CredentialIssuerToken, true
}

// HasCredentialIssuerToken returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasCredentialIssuerToken() bool {
	if o != nil && o.CredentialIssuerToken != nil {
		return true
	}

	return false
}

// SetCredentialIssuerToken gets a reference to the given string and assigns it to the CredentialIssuerToken field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetCredentialIssuerToken(v string) {
	o.CredentialIssuerToken = &v
}

// GetPlanID returns the PlanID field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetPlanID() string {
	if o == nil || o.PlanID == nil {
		var ret string
		return ret
	}
	return *o.PlanID
}

// GetPlanIDOk returns a tuple with the PlanID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetPlanIDOk() (*string, bool) {
	if o == nil || o.PlanID == nil {
		return nil, false
	}
	return o.PlanID, true
}

// HasPlanID returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasPlanID() bool {
	if o != nil && o.PlanID != nil {
		return true
	}

	return false
}

// SetPlanID gets a reference to the given string and assigns it to the PlanID field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetPlanID(v string) {
	o.PlanID = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetPaymentMethodId() string {
	if o == nil || o.PaymentMethodId == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || o.PaymentMethodId == nil {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasPaymentMethodId() bool {
	if o != nil && o.PaymentMethodId != nil {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostId() string {
	if o == nil || o.HostId == nil {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetHostIdOk() (*string, bool) {
	if o == nil || o.HostId == nil {
		return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasHostId() bool {
	if o != nil && o.HostId != nil {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetHostId(v string) {
	o.HostId = &v
}

// GetAcquirerReferenceData returns the AcquirerReferenceData field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetAcquirerReferenceData() string {
	if o == nil || o.AcquirerReferenceData == nil {
		var ret string
		return ret
	}
	return *o.AcquirerReferenceData
}

// GetAcquirerReferenceDataOk returns a tuple with the AcquirerReferenceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetAcquirerReferenceDataOk() (*string, bool) {
	if o == nil || o.AcquirerReferenceData == nil {
		return nil, false
	}
	return o.AcquirerReferenceData, true
}

// HasAcquirerReferenceData returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasAcquirerReferenceData() bool {
	if o != nil && o.AcquirerReferenceData != nil {
		return true
	}

	return false
}

// SetAcquirerReferenceData gets a reference to the given string and assigns it to the AcquirerReferenceData field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetAcquirerReferenceData(v string) {
	o.AcquirerReferenceData = &v
}

// GetIdentifierForTheAcquirer returns the IdentifierForTheAcquirer field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetIdentifierForTheAcquirer() string {
	if o == nil || o.IdentifierForTheAcquirer == nil {
		var ret string
		return ret
	}
	return *o.IdentifierForTheAcquirer
}

// GetIdentifierForTheAcquirerOk returns a tuple with the IdentifierForTheAcquirer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetIdentifierForTheAcquirerOk() (*string, bool) {
	if o == nil || o.IdentifierForTheAcquirer == nil {
		return nil, false
	}
	return o.IdentifierForTheAcquirer, true
}

// HasIdentifierForTheAcquirer returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasIdentifierForTheAcquirer() bool {
	if o != nil && o.IdentifierForTheAcquirer != nil {
		return true
	}

	return false
}

// SetIdentifierForTheAcquirer gets a reference to the given string and assigns it to the IdentifierForTheAcquirer field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetIdentifierForTheAcquirer(v string) {
	o.IdentifierForTheAcquirer = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetConfiguration() SaleResponseObjectSaleResponseConfiguration {
	if o == nil || o.Configuration == nil {
		var ret SaleResponseObjectSaleResponseConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) GetConfigurationOk() (*SaleResponseObjectSaleResponseConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *OrderFinalResponseObjectOrderFinalResponse) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given SaleResponseObjectSaleResponseConfiguration and assigns it to the Configuration field.
func (o *OrderFinalResponseObjectOrderFinalResponse) SetConfiguration(v SaleResponseObjectSaleResponseConfiguration) {
	o.Configuration = &v
}

func (o OrderFinalResponseObjectOrderFinalResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestType != nil {
		toSerialize["RequestType"] = o.RequestType
	}
	if true {
		toSerialize["ResponseActions"] = o.ResponseActions
	}
	if true {
		toSerialize["ResponseMessage"] = o.ResponseMessage
	}
	if o.SystemIdentification != nil {
		toSerialize["SystemIdentification"] = o.SystemIdentification
	}
	if o.CompanyIdentification != nil {
		toSerialize["CompanyIdentification"] = o.CompanyIdentification
	}
	if o.BranchIdentification != nil {
		toSerialize["BranchIdentification"] = o.BranchIdentification
	}
	if o.POSIdentification != nil {
		toSerialize["POSIdentification"] = o.POSIdentification
	}
	if o.ForeignResponseCode != nil {
		toSerialize["ForeignResponseCode"] = o.ForeignResponseCode
	}
	if true {
		toSerialize["ResponseCode"] = o.ResponseCode
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
	if o.AdditionalInformation != nil {
		toSerialize["AdditionalInformation"] = o.AdditionalInformation
	}
	if o.RequiredInformation != nil {
		toSerialize["RequiredInformation"] = o.RequiredInformation
	}
	if o.TransactionIdentification != nil {
		toSerialize["TransactionIdentification"] = o.TransactionIdentification
	}
	if o.TransactionDescription != nil {
		toSerialize["TransactionDescription"] = o.TransactionDescription
	}
	if o.PaymentTransactionIdentification != nil {
		toSerialize["PaymentTransactionIdentification"] = o.PaymentTransactionIdentification
	}
	if o.HostResponseCode != nil {
		toSerialize["HostResponseCode"] = o.HostResponseCode
	}
	if o.HostResponseMessage != nil {
		toSerialize["HostResponseMessage"] = o.HostResponseMessage
	}
	if o.HostResponseAction != nil {
		toSerialize["HostResponseAction"] = o.HostResponseAction
	}
	if o.HostMerchantIdentification != nil {
		toSerialize["HostMerchantIdentification"] = o.HostMerchantIdentification
	}
	if o.HostTerminalIdentification != nil {
		toSerialize["HostTerminalIdentification"] = o.HostTerminalIdentification
	}
	if o.HostBatchNumber != nil {
		toSerialize["HostBatchNumber"] = o.HostBatchNumber
	}
	if o.HostTicketNumber != nil {
		toSerialize["HostTicketNumber"] = o.HostTicketNumber
	}
	if o.HostTraceNumber != nil {
		toSerialize["HostTraceNumber"] = o.HostTraceNumber
	}
	if o.HostRetrievalReferenceNumber != nil {
		toSerialize["HostRetrievalReferenceNumber"] = o.HostRetrievalReferenceNumber
	}
	if o.HostTransactionIdentification != nil {
		toSerialize["HostTransactionIdentification"] = o.HostTransactionIdentification
	}
	if o.HostAuthorizationCode != nil {
		toSerialize["HostAuthorizationCode"] = o.HostAuthorizationCode
	}
	if o.CredentialToken != nil {
		toSerialize["CredentialToken"] = o.CredentialToken
	}
	if o.CredentialIssuerToken != nil {
		toSerialize["CredentialIssuerToken"] = o.CredentialIssuerToken
	}
	if o.PlanID != nil {
		toSerialize["PlanID"] = o.PlanID
	}
	if o.PaymentMethodId != nil {
		toSerialize["PaymentMethodId"] = o.PaymentMethodId
	}
	if o.HostId != nil {
		toSerialize["HostId"] = o.HostId
	}
	if o.AcquirerReferenceData != nil {
		toSerialize["AcquirerReferenceData"] = o.AcquirerReferenceData
	}
	if o.IdentifierForTheAcquirer != nil {
		toSerialize["IdentifierForTheAcquirer"] = o.IdentifierForTheAcquirer
	}
	if o.Configuration != nil {
		toSerialize["Configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableOrderFinalResponseObjectOrderFinalResponse struct {
	value *OrderFinalResponseObjectOrderFinalResponse
	isSet bool
}

func (v NullableOrderFinalResponseObjectOrderFinalResponse) Get() *OrderFinalResponseObjectOrderFinalResponse {
	return v.value
}

func (v *NullableOrderFinalResponseObjectOrderFinalResponse) Set(val *OrderFinalResponseObjectOrderFinalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderFinalResponseObjectOrderFinalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderFinalResponseObjectOrderFinalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderFinalResponseObjectOrderFinalResponse(val *OrderFinalResponseObjectOrderFinalResponse) *NullableOrderFinalResponseObjectOrderFinalResponse {
	return &NullableOrderFinalResponseObjectOrderFinalResponse{value: val, isSet: true}
}

func (v NullableOrderFinalResponseObjectOrderFinalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderFinalResponseObjectOrderFinalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


