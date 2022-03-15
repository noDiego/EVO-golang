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

// SaleResponseObjectSaleResponsePlans Plan utilizado en la transacción.
type SaleResponseObjectSaleResponsePlans struct {
	// Identificador interno formado por el TEF por una combinación de datos del plan especifico. Si el plan no posee un código de emisor, el último campo no estará presente (ni el _).
	Id *string `json:"Id,omitempty"`
	// código de el plan.
	PlanCode *string `json:"PlanCode,omitempty"`
	// ID externo para que el punto de venta pueda reconocer el plan.
	ForeignIdentifier *string `json:"ForeignIdentifier,omitempty"`
	// ID del emisor que se encuentra disponible para este plan.
	Issuer *string `json:"Issuer,omitempty"`
	// Descripción del plan.
	Description *string `json:"Description,omitempty"`
	// Cantidad de cuotas que permite este plan
	FacilityPayments *int32 `json:"FacilityPayments,omitempty"`
	// Tipo de Plan de Financiación
	FacilityType *string `json:"FacilityType,omitempty"`
	// código de Moneda - ISO 4217 <https://en.wikipedia.org/wiki/ISO_4217 Se puede utilizar la Codificación Alfabética o Numérica <br />   * Num   - Alpha - Description <br />   * '032' - 'ARS' - Pesos Argentinos <br />   * '152' - 'CLP' - Pesos Chilenos <br/>   * '484' - 'MXN' - Pesos Mexicanos <br/>   * '840' - 'USD' - dólares Americanos <br/>   * '878' - 'EUR' - Euros <br/>   * '858' - 'UYU' - Pesos Uruguayos <br/>   * '878' - 'EUR' - Euros <br/>   * '986' - 'BRL' - Real Brasileño
	CurrencyCode *string `json:"CurrencyCode,omitempty"`
	// Porcentaje de recargo completo para el monto de la operación.
	Rate *float32 `json:"Rate,omitempty"`
	// Monto debido al porcentaje de recargo.
	Amount *float32 `json:"Amount,omitempty"`
	// Monto máximo para operar de forma OFFLINE.
	OfflineAmountLimit *float32 `json:"OfflineAmountLimit,omitempty"`
	Deferral *SaleResponseObjectSaleResponsePlansDeferral `json:"Deferral,omitempty"`
	Cashback *SaleResponseObjectSaleResponsePlansCashback `json:"Cashback,omitempty"`
	// Se informará la tasa nominal anual, en casos en que el plan elegido para realizar la venta lo posea. Por ejemplo, el plan especial llamado Plan V de Prisma informará este valor, dado que se obtendrá dinámicamente, consultandolo instante a instante.
	TNA *float32 `json:"TNA,omitempty"`
	// Se informará la tasa efectiva mensual, en casos en que el plan elegido para realizar la venta lo posea. 
	TEM *float32 `json:"TEM,omitempty"`
	// Tasa Efectiva anual. Este campo estará presente solo si el tipo de plan es dinámico, o si fue ingresado un valor en la base de datos.
	TEA *float32 `json:"TEA,omitempty"`
	// Costo Financiero Total. Este campo estará presente solo si el tipo de plan es dinámico, o si fue ingresado un valor en la base de datos.        
	CFT *float32 `json:"CFT,omitempty"`
	// Monto final a pagar en cada una de las cuotas en las que se divida la compra
	ValueFacilityPayments *float32 `json:"ValueFacilityPayments,omitempty"`
	// Nombre del emisor de este plan. Estará presente solo si existe un solo emisor para todos los planes, o si el plan es dinámico
	IssuerName *string `json:"IssuerName,omitempty"`
	// Flag que indica si el plan es del tipo dinamico o no
	IsDynamic *bool `json:"IsDynamic,omitempty"`
	// Campo que le permite al punto de venta elegir entre un plan u otro
	Category *string `json:"Category,omitempty"`
	// Lista de Acciones que debe ejecutar el POS o el Dispositvo para el caso que este Plan sea seleccionado. Acciones para el Device <b>RequestPIN</b>
	POSOrDeviceActions []string `json:"POSOrDeviceActions,omitempty"`
}

// NewSaleResponseObjectSaleResponsePlans instantiates a new SaleResponseObjectSaleResponsePlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaleResponseObjectSaleResponsePlans() *SaleResponseObjectSaleResponsePlans {
	this := SaleResponseObjectSaleResponsePlans{}
	return &this
}

// NewSaleResponseObjectSaleResponsePlansWithDefaults instantiates a new SaleResponseObjectSaleResponsePlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaleResponseObjectSaleResponsePlansWithDefaults() *SaleResponseObjectSaleResponsePlans {
	this := SaleResponseObjectSaleResponsePlans{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SaleResponseObjectSaleResponsePlans) SetId(v string) {
	o.Id = &v
}

// GetPlanCode returns the PlanCode field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetPlanCode() string {
	if o == nil || o.PlanCode == nil {
		var ret string
		return ret
	}
	return *o.PlanCode
}

// GetPlanCodeOk returns a tuple with the PlanCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetPlanCodeOk() (*string, bool) {
	if o == nil || o.PlanCode == nil {
		return nil, false
	}
	return o.PlanCode, true
}

// HasPlanCode returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasPlanCode() bool {
	if o != nil && o.PlanCode != nil {
		return true
	}

	return false
}

// SetPlanCode gets a reference to the given string and assigns it to the PlanCode field.
func (o *SaleResponseObjectSaleResponsePlans) SetPlanCode(v string) {
	o.PlanCode = &v
}

// GetForeignIdentifier returns the ForeignIdentifier field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetForeignIdentifier() string {
	if o == nil || o.ForeignIdentifier == nil {
		var ret string
		return ret
	}
	return *o.ForeignIdentifier
}

// GetForeignIdentifierOk returns a tuple with the ForeignIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetForeignIdentifierOk() (*string, bool) {
	if o == nil || o.ForeignIdentifier == nil {
		return nil, false
	}
	return o.ForeignIdentifier, true
}

// HasForeignIdentifier returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasForeignIdentifier() bool {
	if o != nil && o.ForeignIdentifier != nil {
		return true
	}

	return false
}

// SetForeignIdentifier gets a reference to the given string and assigns it to the ForeignIdentifier field.
func (o *SaleResponseObjectSaleResponsePlans) SetForeignIdentifier(v string) {
	o.ForeignIdentifier = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *SaleResponseObjectSaleResponsePlans) SetIssuer(v string) {
	o.Issuer = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SaleResponseObjectSaleResponsePlans) SetDescription(v string) {
	o.Description = &v
}

// GetFacilityPayments returns the FacilityPayments field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetFacilityPayments() int32 {
	if o == nil || o.FacilityPayments == nil {
		var ret int32
		return ret
	}
	return *o.FacilityPayments
}

// GetFacilityPaymentsOk returns a tuple with the FacilityPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetFacilityPaymentsOk() (*int32, bool) {
	if o == nil || o.FacilityPayments == nil {
		return nil, false
	}
	return o.FacilityPayments, true
}

// HasFacilityPayments returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasFacilityPayments() bool {
	if o != nil && o.FacilityPayments != nil {
		return true
	}

	return false
}

// SetFacilityPayments gets a reference to the given int32 and assigns it to the FacilityPayments field.
func (o *SaleResponseObjectSaleResponsePlans) SetFacilityPayments(v int32) {
	o.FacilityPayments = &v
}

// GetFacilityType returns the FacilityType field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetFacilityType() string {
	if o == nil || o.FacilityType == nil {
		var ret string
		return ret
	}
	return *o.FacilityType
}

// GetFacilityTypeOk returns a tuple with the FacilityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetFacilityTypeOk() (*string, bool) {
	if o == nil || o.FacilityType == nil {
		return nil, false
	}
	return o.FacilityType, true
}

// HasFacilityType returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasFacilityType() bool {
	if o != nil && o.FacilityType != nil {
		return true
	}

	return false
}

// SetFacilityType gets a reference to the given string and assigns it to the FacilityType field.
func (o *SaleResponseObjectSaleResponsePlans) SetFacilityType(v string) {
	o.FacilityType = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *SaleResponseObjectSaleResponsePlans) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetRate() float32 {
	if o == nil || o.Rate == nil {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetRateOk() (*float32, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *SaleResponseObjectSaleResponsePlans) SetRate(v float32) {
	o.Rate = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *SaleResponseObjectSaleResponsePlans) SetAmount(v float32) {
	o.Amount = &v
}

// GetOfflineAmountLimit returns the OfflineAmountLimit field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetOfflineAmountLimit() float32 {
	if o == nil || o.OfflineAmountLimit == nil {
		var ret float32
		return ret
	}
	return *o.OfflineAmountLimit
}

// GetOfflineAmountLimitOk returns a tuple with the OfflineAmountLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetOfflineAmountLimitOk() (*float32, bool) {
	if o == nil || o.OfflineAmountLimit == nil {
		return nil, false
	}
	return o.OfflineAmountLimit, true
}

// HasOfflineAmountLimit returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasOfflineAmountLimit() bool {
	if o != nil && o.OfflineAmountLimit != nil {
		return true
	}

	return false
}

// SetOfflineAmountLimit gets a reference to the given float32 and assigns it to the OfflineAmountLimit field.
func (o *SaleResponseObjectSaleResponsePlans) SetOfflineAmountLimit(v float32) {
	o.OfflineAmountLimit = &v
}

// GetDeferral returns the Deferral field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetDeferral() SaleResponseObjectSaleResponsePlansDeferral {
	if o == nil || o.Deferral == nil {
		var ret SaleResponseObjectSaleResponsePlansDeferral
		return ret
	}
	return *o.Deferral
}

// GetDeferralOk returns a tuple with the Deferral field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetDeferralOk() (*SaleResponseObjectSaleResponsePlansDeferral, bool) {
	if o == nil || o.Deferral == nil {
		return nil, false
	}
	return o.Deferral, true
}

// HasDeferral returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasDeferral() bool {
	if o != nil && o.Deferral != nil {
		return true
	}

	return false
}

// SetDeferral gets a reference to the given SaleResponseObjectSaleResponsePlansDeferral and assigns it to the Deferral field.
func (o *SaleResponseObjectSaleResponsePlans) SetDeferral(v SaleResponseObjectSaleResponsePlansDeferral) {
	o.Deferral = &v
}

// GetCashback returns the Cashback field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetCashback() SaleResponseObjectSaleResponsePlansCashback {
	if o == nil || o.Cashback == nil {
		var ret SaleResponseObjectSaleResponsePlansCashback
		return ret
	}
	return *o.Cashback
}

// GetCashbackOk returns a tuple with the Cashback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetCashbackOk() (*SaleResponseObjectSaleResponsePlansCashback, bool) {
	if o == nil || o.Cashback == nil {
		return nil, false
	}
	return o.Cashback, true
}

// HasCashback returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasCashback() bool {
	if o != nil && o.Cashback != nil {
		return true
	}

	return false
}

// SetCashback gets a reference to the given SaleResponseObjectSaleResponsePlansCashback and assigns it to the Cashback field.
func (o *SaleResponseObjectSaleResponsePlans) SetCashback(v SaleResponseObjectSaleResponsePlansCashback) {
	o.Cashback = &v
}

// GetTNA returns the TNA field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetTNA() float32 {
	if o == nil || o.TNA == nil {
		var ret float32
		return ret
	}
	return *o.TNA
}

// GetTNAOk returns a tuple with the TNA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetTNAOk() (*float32, bool) {
	if o == nil || o.TNA == nil {
		return nil, false
	}
	return o.TNA, true
}

// HasTNA returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasTNA() bool {
	if o != nil && o.TNA != nil {
		return true
	}

	return false
}

// SetTNA gets a reference to the given float32 and assigns it to the TNA field.
func (o *SaleResponseObjectSaleResponsePlans) SetTNA(v float32) {
	o.TNA = &v
}

// GetTEM returns the TEM field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetTEM() float32 {
	if o == nil || o.TEM == nil {
		var ret float32
		return ret
	}
	return *o.TEM
}

// GetTEMOk returns a tuple with the TEM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetTEMOk() (*float32, bool) {
	if o == nil || o.TEM == nil {
		return nil, false
	}
	return o.TEM, true
}

// HasTEM returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasTEM() bool {
	if o != nil && o.TEM != nil {
		return true
	}

	return false
}

// SetTEM gets a reference to the given float32 and assigns it to the TEM field.
func (o *SaleResponseObjectSaleResponsePlans) SetTEM(v float32) {
	o.TEM = &v
}

// GetTEA returns the TEA field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetTEA() float32 {
	if o == nil || o.TEA == nil {
		var ret float32
		return ret
	}
	return *o.TEA
}

// GetTEAOk returns a tuple with the TEA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetTEAOk() (*float32, bool) {
	if o == nil || o.TEA == nil {
		return nil, false
	}
	return o.TEA, true
}

// HasTEA returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasTEA() bool {
	if o != nil && o.TEA != nil {
		return true
	}

	return false
}

// SetTEA gets a reference to the given float32 and assigns it to the TEA field.
func (o *SaleResponseObjectSaleResponsePlans) SetTEA(v float32) {
	o.TEA = &v
}

// GetCFT returns the CFT field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetCFT() float32 {
	if o == nil || o.CFT == nil {
		var ret float32
		return ret
	}
	return *o.CFT
}

// GetCFTOk returns a tuple with the CFT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetCFTOk() (*float32, bool) {
	if o == nil || o.CFT == nil {
		return nil, false
	}
	return o.CFT, true
}

// HasCFT returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasCFT() bool {
	if o != nil && o.CFT != nil {
		return true
	}

	return false
}

// SetCFT gets a reference to the given float32 and assigns it to the CFT field.
func (o *SaleResponseObjectSaleResponsePlans) SetCFT(v float32) {
	o.CFT = &v
}

// GetValueFacilityPayments returns the ValueFacilityPayments field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetValueFacilityPayments() float32 {
	if o == nil || o.ValueFacilityPayments == nil {
		var ret float32
		return ret
	}
	return *o.ValueFacilityPayments
}

// GetValueFacilityPaymentsOk returns a tuple with the ValueFacilityPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetValueFacilityPaymentsOk() (*float32, bool) {
	if o == nil || o.ValueFacilityPayments == nil {
		return nil, false
	}
	return o.ValueFacilityPayments, true
}

// HasValueFacilityPayments returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasValueFacilityPayments() bool {
	if o != nil && o.ValueFacilityPayments != nil {
		return true
	}

	return false
}

// SetValueFacilityPayments gets a reference to the given float32 and assigns it to the ValueFacilityPayments field.
func (o *SaleResponseObjectSaleResponsePlans) SetValueFacilityPayments(v float32) {
	o.ValueFacilityPayments = &v
}

// GetIssuerName returns the IssuerName field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetIssuerName() string {
	if o == nil || o.IssuerName == nil {
		var ret string
		return ret
	}
	return *o.IssuerName
}

// GetIssuerNameOk returns a tuple with the IssuerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetIssuerNameOk() (*string, bool) {
	if o == nil || o.IssuerName == nil {
		return nil, false
	}
	return o.IssuerName, true
}

// HasIssuerName returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasIssuerName() bool {
	if o != nil && o.IssuerName != nil {
		return true
	}

	return false
}

// SetIssuerName gets a reference to the given string and assigns it to the IssuerName field.
func (o *SaleResponseObjectSaleResponsePlans) SetIssuerName(v string) {
	o.IssuerName = &v
}

// GetIsDynamic returns the IsDynamic field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetIsDynamic() bool {
	if o == nil || o.IsDynamic == nil {
		var ret bool
		return ret
	}
	return *o.IsDynamic
}

// GetIsDynamicOk returns a tuple with the IsDynamic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetIsDynamicOk() (*bool, bool) {
	if o == nil || o.IsDynamic == nil {
		return nil, false
	}
	return o.IsDynamic, true
}

// HasIsDynamic returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasIsDynamic() bool {
	if o != nil && o.IsDynamic != nil {
		return true
	}

	return false
}

// SetIsDynamic gets a reference to the given bool and assigns it to the IsDynamic field.
func (o *SaleResponseObjectSaleResponsePlans) SetIsDynamic(v bool) {
	o.IsDynamic = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SaleResponseObjectSaleResponsePlans) SetCategory(v string) {
	o.Category = &v
}

// GetPOSOrDeviceActions returns the POSOrDeviceActions field value if set, zero value otherwise.
func (o *SaleResponseObjectSaleResponsePlans) GetPOSOrDeviceActions() []string {
	if o == nil || o.POSOrDeviceActions == nil {
		var ret []string
		return ret
	}
	return o.POSOrDeviceActions
}

// GetPOSOrDeviceActionsOk returns a tuple with the POSOrDeviceActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleResponseObjectSaleResponsePlans) GetPOSOrDeviceActionsOk() ([]string, bool) {
	if o == nil || o.POSOrDeviceActions == nil {
		return nil, false
	}
	return o.POSOrDeviceActions, true
}

// HasPOSOrDeviceActions returns a boolean if a field has been set.
func (o *SaleResponseObjectSaleResponsePlans) HasPOSOrDeviceActions() bool {
	if o != nil && o.POSOrDeviceActions != nil {
		return true
	}

	return false
}

// SetPOSOrDeviceActions gets a reference to the given []string and assigns it to the POSOrDeviceActions field.
func (o *SaleResponseObjectSaleResponsePlans) SetPOSOrDeviceActions(v []string) {
	o.POSOrDeviceActions = v
}

func (o SaleResponseObjectSaleResponsePlans) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.PlanCode != nil {
		toSerialize["PlanCode"] = o.PlanCode
	}
	if o.ForeignIdentifier != nil {
		toSerialize["ForeignIdentifier"] = o.ForeignIdentifier
	}
	if o.Issuer != nil {
		toSerialize["Issuer"] = o.Issuer
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.FacilityPayments != nil {
		toSerialize["FacilityPayments"] = o.FacilityPayments
	}
	if o.FacilityType != nil {
		toSerialize["FacilityType"] = o.FacilityType
	}
	if o.CurrencyCode != nil {
		toSerialize["CurrencyCode"] = o.CurrencyCode
	}
	if o.Rate != nil {
		toSerialize["Rate"] = o.Rate
	}
	if o.Amount != nil {
		toSerialize["Amount"] = o.Amount
	}
	if o.OfflineAmountLimit != nil {
		toSerialize["OfflineAmountLimit"] = o.OfflineAmountLimit
	}
	if o.Deferral != nil {
		toSerialize["Deferral"] = o.Deferral
	}
	if o.Cashback != nil {
		toSerialize["Cashback"] = o.Cashback
	}
	if o.TNA != nil {
		toSerialize["TNA"] = o.TNA
	}
	if o.TEM != nil {
		toSerialize["TEM"] = o.TEM
	}
	if o.TEA != nil {
		toSerialize["TEA"] = o.TEA
	}
	if o.CFT != nil {
		toSerialize["CFT"] = o.CFT
	}
	if o.ValueFacilityPayments != nil {
		toSerialize["ValueFacilityPayments"] = o.ValueFacilityPayments
	}
	if o.IssuerName != nil {
		toSerialize["IssuerName"] = o.IssuerName
	}
	if o.IsDynamic != nil {
		toSerialize["IsDynamic"] = o.IsDynamic
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.POSOrDeviceActions != nil {
		toSerialize["POSOrDeviceActions"] = o.POSOrDeviceActions
	}
	return json.Marshal(toSerialize)
}

type NullableSaleResponseObjectSaleResponsePlans struct {
	value *SaleResponseObjectSaleResponsePlans
	isSet bool
}

func (v NullableSaleResponseObjectSaleResponsePlans) Get() *SaleResponseObjectSaleResponsePlans {
	return v.value
}

func (v *NullableSaleResponseObjectSaleResponsePlans) Set(val *SaleResponseObjectSaleResponsePlans) {
	v.value = val
	v.isSet = true
}

func (v NullableSaleResponseObjectSaleResponsePlans) IsSet() bool {
	return v.isSet
}

func (v *NullableSaleResponseObjectSaleResponsePlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaleResponseObjectSaleResponsePlans(val *SaleResponseObjectSaleResponsePlans) *NullableSaleResponseObjectSaleResponsePlans {
	return &NullableSaleResponseObjectSaleResponsePlans{value: val, isSet: true}
}

func (v NullableSaleResponseObjectSaleResponsePlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaleResponseObjectSaleResponsePlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


