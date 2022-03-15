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

// WalletRequestResponseObjectWalletRequestResponse struct for WalletRequestResponseObjectWalletRequestResponse
type WalletRequestResponseObjectWalletRequestResponse struct {
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
	// Tipo de Operación que se está requiriendo, solo necesario sobre formatos que no soportan elementos complejos o compuestos.
	AnswerType *string `json:"AnswerType,omitempty"`
	// Código de identificación, generado por Plataforma, de la operación realizada
	AnswerKey string `json:"AnswerKey"`
	// ID generado para la identificación por parte del Plataforma de la información generada en la ejecución de un GetCard o un Payment Method. Sera necesario para que un mensaje de Sale, Void o Payment Method identifique el contexto generado y lo utilice para esa operación.
	RequestKey *string `json:"RequestKey,omitempty"`
	// Versión del Servicio de la Plataforma que resolvió el requerimiento.
	ServerVersion *string `json:"ServerVersion,omitempty"`
	// Dirección IP del Server que atiende el requerimiento.
	ServerAddress *string `json:"ServerAddress,omitempty"`
	// Instancia de Server que atiende el requerimiento.
	ServerInstance *string `json:"ServerInstance,omitempty"`
	// Nombre del Nodo que atendió el requerimiento.
	ServerNodeName *string `json:"ServerNodeName,omitempty"`
	// Identificador Unívoco del Mensaje ( UUID v5 ).
	MessageID *string `json:"MessageID,omitempty"`
	// Versión del  Adaptador de Protocolo Entrante que atiende el Requerimiento.
	AdapterInputVersion *string `json:"AdapterInputVersion,omitempty"`
	// Dirección IP del Adaptador de Protocolo Entrante que atiende el requerimiento.
	AdapterInputAddress *string `json:"AdapterInputAddress,omitempty"`
	// Nombre del Nodo del Adaptador de Protocolo Entrante que atiende el requerimiento.
	AdapterInputNodeName *string `json:"AdapterInputNodeName,omitempty"`
	// Versión del  Adaptador de Protocolo Saliente que atiende el Requerimiento.
	AdapterOutputVersion *string `json:"AdapterOutputVersion,omitempty"`
	// Dirección IP  del  Adaptador de Protocolo Saliente que atiende el Requerimiento.
	AdapterOutputAddress *string `json:"AdapterOutputAddress,omitempty"`
	// Nombre del Nodo  del  Adaptador de Protocolo Saliente que atiende el Requerimiento.
	AdapterOutputNodeName *string `json:"AdapterOutputNodeName,omitempty"`
	// Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible.
	ServiceVersion *string `json:"ServiceVersion,omitempty"`
	// Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma.
	Sequence *string `json:"Sequence,omitempty"`
	// Datos asociados a la seguridad de la transacción o de elementos sensibles.
	Security []SaleObjectSaleSecurity `json:"Security,omitempty"`
	// Flag indicador de generación de reversa para la última operación reversable
	WasReversePrevious *int32 `json:"WasReversePrevious,omitempty"`
	// Flag indicador de cancelación o confirmación del último bloque de transacciones, previo al nuevo valor recibido
	WasClosedPreviousBlock *int32 `json:"WasClosedPreviousBlock,omitempty"`
	// ID que identifica a la operación que acaba de ser reversada.
	ReversedAnswerKey *string `json:"ReversedAnswerKey,omitempty"`
	// Secuencia de la transacción que fue reversada
	ReversedSequence *string `json:"ReversedSequence,omitempty"`
	// ID del bloque de transacciones que ha sido confirmado de forma automática (es decir, sin recibir un requerimiento de BlockClose). Este escenario se presentará si el Plataforma así se ha configurado para actuar bajo esa circunstancia.
	CommittedBlock *string `json:"CommittedBlock,omitempty"`
	// ID del bloque de transacciones que ha sido cancelado de forma automática (es decir, sin recibir un requerimiento de BlockClose). Este escenario se presentará si el Plataforma así se ha configurado para actuar bajo esa circunstancia.
	ReversedBlock *string `json:"ReversedBlock,omitempty"`
	// En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente.
	RequiredInformation []DebtPaymentObjectDebtPaymentRequiredInformation `json:"RequiredInformation,omitempty"`
	// En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente.
	AdditionalInformation []SaleResponseObjectSaleResponseAdditionalInformation `json:"AdditionalInformation,omitempty"`
	// Información adicional/Mensaje promocional/Leyenda de respuesta a mostrar en pantalla en el ticket de la operación. Cada línea de este mensaje será un elemento dentro del array.
	DisplayResponseMessage []string `json:"DisplayResponseMessage,omitempty"`
	// ID que identifica a un grupo de transacciones que serán confirmadas o canceladas
	Block *string `json:"Block,omitempty"`
	// Fecha y hora de transmision de la operación hacia el host - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14
	TransmitionDateTime *time.Time `json:"TransmitionDateTime,omitempty"`
	// ID de La transacción UNIVOCO para el Punto de venta
	TransactionIdentification *string `json:"TransactionIdentification,omitempty"`
	// Descripción del tipo de operación que se realizará
	TransactionDescription *string `json:"TransactionDescription,omitempty"`
	// Fecha y Hora de la transacción en la plataforma de integración - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14
	TrasactionDateTime *string `json:"TrasactionDateTime,omitempty"`
	// Monto del recargo impositivo aplicado al costo financiero que la transacción tiene
	TaxFinancialCostAmount *float32 `json:"TaxFinancialCostAmount,omitempty"`
	// Porcentaje de recargo impositivo a aplicar sobre el monto del costo financiero
	TaxFinancialCostPercentage *float32 `json:"TaxFinancialCostPercentage,omitempty"`
	// Monto del Costo financiero total generado en base al plan elegido
	FinancialCostAmount *float32 `json:"FinancialCostAmount,omitempty"`
	// Porcentaje del costo financiero a aplicar al monto de la transacción
	FinancialCostPercentage *float32 `json:"FinancialCostPercentage,omitempty"`
	// Monto libre de costos financerios e impuestos por el que la venta se realizó. El monto cobrado realmente no es este, dado que no incluye las tasas e impuestos
	RequestAmount *float32 `json:"RequestAmount,omitempty"`
	// Identificador Publico para ser identificar la transacción a realizar, normalmente es un código a Presentar
	WalletRequestToken *string `json:"WalletRequestToken,omitempty"`
	// Tipo de Identificador Público para ser identificar la transacción a realizar, normalmente es un código a Presentar. En el caso de JPG o PNG las mismas estaran codificadas en Base64
	WalletRequestTokenType *string `json:"WalletRequestTokenType,omitempty"`
	// Metodo de Pago utilizado en el Pago
	PaymentMethodIdentification *string `json:"PaymentMethodIdentification,omitempty"`
	// Metodo de Pago utilizado en el Pago expresado en la codificación que el aplicativo que se integra posea
	ForeignPaymentMethodIdentification *string `json:"ForeignPaymentMethodIdentification,omitempty"`
	// Emisor del Medio de Pago utilizado
	IssuerIdentification *string `json:"IssuerIdentification,omitempty"`
	// Emisor del Medio del Pago utilizado en el Pago expresado en la codificación que el aplicativo que se integra posea
	ForeignIssuerIdentification *string `json:"ForeignIssuerIdentification,omitempty"`
	// Número o Identificador de Credencia que se puede Mostrar o Imprimir ( Normalmente Número de Tarjeta enmascarado para el caso de Tarjetas de Credito o débito)
	DisplayCredentialIdentification *string `json:"DisplayCredentialIdentification,omitempty"`
	// Número Ticket  o Voucher Generado para la Plataforma.
	AuthTicket *int32 `json:"AuthTicket,omitempty"`
	// Número de identificación de la transacción, utilizado por la mayoría de los hosts para realizar anulaciones y devoluciones
	AuthRRN *string `json:"AuthRRN,omitempty"`
	// Identificador que genera el Host Adquirente para la Transacción en algunos podrá ser igual al AuthRRN
	IdentifierForTheAdquirer *string `json:"IdentifierForTheAdquirer,omitempty"`
	// Identificador que genera el Host o Plataforma que resuelve la transacción
	IdentifierForTheResolutor *string `json:"IdentifierForTheResolutor,omitempty"`
	// Tiempo para el proximo reintento de la operación <b>WalletRequest</b> expresado en milisegúndos, si esta presenta la acción <b>Retry</b>
	RetryTime *int32 `json:"RetryTime,omitempty"`
	MerchantCategory *SaleResponseObjectSaleResponseMerchantCategory `json:"MerchantCategory,omitempty"`
	Configuration *SaleResponseObjectSaleResponseConfiguration `json:"Configuration,omitempty"`
	// Elemento Compuesto que indica qué Tickets intervienen en la transacción y si deben ser digitalizados o impresos.
	Tickets []SaleResponseObjectSaleResponseTickets `json:"Tickets,omitempty"`
}

// NewWalletRequestResponseObjectWalletRequestResponse instantiates a new WalletRequestResponseObjectWalletRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletRequestResponseObjectWalletRequestResponse(responseActions []string, responseMessage string, responseCode int32, answerKey string) *WalletRequestResponseObjectWalletRequestResponse {
	this := WalletRequestResponseObjectWalletRequestResponse{}
	this.ResponseActions = responseActions
	this.ResponseMessage = responseMessage
	this.ResponseCode = responseCode
	this.AnswerKey = answerKey
	return &this
}

// NewWalletRequestResponseObjectWalletRequestResponseWithDefaults instantiates a new WalletRequestResponseObjectWalletRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletRequestResponseObjectWalletRequestResponseWithDefaults() *WalletRequestResponseObjectWalletRequestResponse {
	this := WalletRequestResponseObjectWalletRequestResponse{}
	return &this
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequestType() string {
	if o == nil || o.RequestType == nil {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequestTypeOk() (*string, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasRequestType() bool {
	if o != nil && o.RequestType != nil {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetRequestType(v string) {
	o.RequestType = &v
}

// GetResponseActions returns the ResponseActions field value
func (o *WalletRequestResponseObjectWalletRequestResponse) GetResponseActions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResponseActions
}

// GetResponseActionsOk returns a tuple with the ResponseActions field value
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetResponseActionsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseActions, true
}

// SetResponseActions sets field value
func (o *WalletRequestResponseObjectWalletRequestResponse) SetResponseActions(v []string) {
	o.ResponseActions = v
}

// GetResponseMessage returns the ResponseMessage field value
func (o *WalletRequestResponseObjectWalletRequestResponse) GetResponseMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetResponseMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResponseMessage, true
}

// SetResponseMessage sets field value
func (o *WalletRequestResponseObjectWalletRequestResponse) SetResponseMessage(v string) {
	o.ResponseMessage = v
}

// GetSystemIdentification returns the SystemIdentification field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetSystemIdentification() string {
	if o == nil || o.SystemIdentification == nil {
		var ret string
		return ret
	}
	return *o.SystemIdentification
}

// GetSystemIdentificationOk returns a tuple with the SystemIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetSystemIdentificationOk() (*string, bool) {
	if o == nil || o.SystemIdentification == nil {
		return nil, false
	}
	return o.SystemIdentification, true
}

// HasSystemIdentification returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasSystemIdentification() bool {
	if o != nil && o.SystemIdentification != nil {
		return true
	}

	return false
}

// SetSystemIdentification gets a reference to the given string and assigns it to the SystemIdentification field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetSystemIdentification(v string) {
	o.SystemIdentification = &v
}

// GetCompanyIdentification returns the CompanyIdentification field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetCompanyIdentification() string {
	if o == nil || o.CompanyIdentification == nil {
		var ret string
		return ret
	}
	return *o.CompanyIdentification
}

// GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetCompanyIdentificationOk() (*string, bool) {
	if o == nil || o.CompanyIdentification == nil {
		return nil, false
	}
	return o.CompanyIdentification, true
}

// HasCompanyIdentification returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasCompanyIdentification() bool {
	if o != nil && o.CompanyIdentification != nil {
		return true
	}

	return false
}

// SetCompanyIdentification gets a reference to the given string and assigns it to the CompanyIdentification field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetCompanyIdentification(v string) {
	o.CompanyIdentification = &v
}

// GetBranchIdentification returns the BranchIdentification field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetBranchIdentification() string {
	if o == nil || o.BranchIdentification == nil {
		var ret string
		return ret
	}
	return *o.BranchIdentification
}

// GetBranchIdentificationOk returns a tuple with the BranchIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetBranchIdentificationOk() (*string, bool) {
	if o == nil || o.BranchIdentification == nil {
		return nil, false
	}
	return o.BranchIdentification, true
}

// HasBranchIdentification returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasBranchIdentification() bool {
	if o != nil && o.BranchIdentification != nil {
		return true
	}

	return false
}

// SetBranchIdentification gets a reference to the given string and assigns it to the BranchIdentification field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetBranchIdentification(v string) {
	o.BranchIdentification = &v
}

// GetPOSIdentification returns the POSIdentification field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetPOSIdentification() string {
	if o == nil || o.POSIdentification == nil {
		var ret string
		return ret
	}
	return *o.POSIdentification
}

// GetPOSIdentificationOk returns a tuple with the POSIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetPOSIdentificationOk() (*string, bool) {
	if o == nil || o.POSIdentification == nil {
		return nil, false
	}
	return o.POSIdentification, true
}

// HasPOSIdentification returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasPOSIdentification() bool {
	if o != nil && o.POSIdentification != nil {
		return true
	}

	return false
}

// SetPOSIdentification gets a reference to the given string and assigns it to the POSIdentification field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetPOSIdentification(v string) {
	o.POSIdentification = &v
}

// GetForeignResponseCode returns the ForeignResponseCode field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetForeignResponseCode() string {
	if o == nil || o.ForeignResponseCode == nil {
		var ret string
		return ret
	}
	return *o.ForeignResponseCode
}

// GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetForeignResponseCodeOk() (*string, bool) {
	if o == nil || o.ForeignResponseCode == nil {
		return nil, false
	}
	return o.ForeignResponseCode, true
}

// HasForeignResponseCode returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasForeignResponseCode() bool {
	if o != nil && o.ForeignResponseCode != nil {
		return true
	}

	return false
}

// SetForeignResponseCode gets a reference to the given string and assigns it to the ForeignResponseCode field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetForeignResponseCode(v string) {
	o.ForeignResponseCode = &v
}

// GetResponseCode returns the ResponseCode field value
func (o *WalletRequestResponseObjectWalletRequestResponse) GetResponseCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetResponseCodeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResponseCode, true
}

// SetResponseCode sets field value
func (o *WalletRequestResponseObjectWalletRequestResponse) SetResponseCode(v int32) {
	o.ResponseCode = v
}

// GetAnswerType returns the AnswerType field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAnswerType() string {
	if o == nil || o.AnswerType == nil {
		var ret string
		return ret
	}
	return *o.AnswerType
}

// GetAnswerTypeOk returns a tuple with the AnswerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAnswerTypeOk() (*string, bool) {
	if o == nil || o.AnswerType == nil {
		return nil, false
	}
	return o.AnswerType, true
}

// HasAnswerType returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasAnswerType() bool {
	if o != nil && o.AnswerType != nil {
		return true
	}

	return false
}

// SetAnswerType gets a reference to the given string and assigns it to the AnswerType field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetAnswerType(v string) {
	o.AnswerType = &v
}

// GetAnswerKey returns the AnswerKey field value
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAnswerKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnswerKey
}

// GetAnswerKeyOk returns a tuple with the AnswerKey field value
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAnswerKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AnswerKey, true
}

// SetAnswerKey sets field value
func (o *WalletRequestResponseObjectWalletRequestResponse) SetAnswerKey(v string) {
	o.AnswerKey = v
}

// GetRequestKey returns the RequestKey field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequestKey() string {
	if o == nil || o.RequestKey == nil {
		var ret string
		return ret
	}
	return *o.RequestKey
}

// GetRequestKeyOk returns a tuple with the RequestKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequestKeyOk() (*string, bool) {
	if o == nil || o.RequestKey == nil {
		return nil, false
	}
	return o.RequestKey, true
}

// HasRequestKey returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasRequestKey() bool {
	if o != nil && o.RequestKey != nil {
		return true
	}

	return false
}

// SetRequestKey gets a reference to the given string and assigns it to the RequestKey field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetRequestKey(v string) {
	o.RequestKey = &v
}

// GetServerVersion returns the ServerVersion field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerVersion() string {
	if o == nil || o.ServerVersion == nil {
		var ret string
		return ret
	}
	return *o.ServerVersion
}

// GetServerVersionOk returns a tuple with the ServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerVersionOk() (*string, bool) {
	if o == nil || o.ServerVersion == nil {
		return nil, false
	}
	return o.ServerVersion, true
}

// HasServerVersion returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasServerVersion() bool {
	if o != nil && o.ServerVersion != nil {
		return true
	}

	return false
}

// SetServerVersion gets a reference to the given string and assigns it to the ServerVersion field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetServerVersion(v string) {
	o.ServerVersion = &v
}

// GetServerAddress returns the ServerAddress field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerAddress() string {
	if o == nil || o.ServerAddress == nil {
		var ret string
		return ret
	}
	return *o.ServerAddress
}

// GetServerAddressOk returns a tuple with the ServerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerAddressOk() (*string, bool) {
	if o == nil || o.ServerAddress == nil {
		return nil, false
	}
	return o.ServerAddress, true
}

// HasServerAddress returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasServerAddress() bool {
	if o != nil && o.ServerAddress != nil {
		return true
	}

	return false
}

// SetServerAddress gets a reference to the given string and assigns it to the ServerAddress field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetServerAddress(v string) {
	o.ServerAddress = &v
}

// GetServerInstance returns the ServerInstance field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerInstance() string {
	if o == nil || o.ServerInstance == nil {
		var ret string
		return ret
	}
	return *o.ServerInstance
}

// GetServerInstanceOk returns a tuple with the ServerInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerInstanceOk() (*string, bool) {
	if o == nil || o.ServerInstance == nil {
		return nil, false
	}
	return o.ServerInstance, true
}

// HasServerInstance returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasServerInstance() bool {
	if o != nil && o.ServerInstance != nil {
		return true
	}

	return false
}

// SetServerInstance gets a reference to the given string and assigns it to the ServerInstance field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetServerInstance(v string) {
	o.ServerInstance = &v
}

// GetServerNodeName returns the ServerNodeName field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerNodeName() string {
	if o == nil || o.ServerNodeName == nil {
		var ret string
		return ret
	}
	return *o.ServerNodeName
}

// GetServerNodeNameOk returns a tuple with the ServerNodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetServerNodeNameOk() (*string, bool) {
	if o == nil || o.ServerNodeName == nil {
		return nil, false
	}
	return o.ServerNodeName, true
}

// HasServerNodeName returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasServerNodeName() bool {
	if o != nil && o.ServerNodeName != nil {
		return true
	}

	return false
}

// SetServerNodeName gets a reference to the given string and assigns it to the ServerNodeName field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetServerNodeName(v string) {
	o.ServerNodeName = &v
}

// GetMessageID returns the MessageID field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetMessageID() string {
	if o == nil || o.MessageID == nil {
		var ret string
		return ret
	}
	return *o.MessageID
}

// GetMessageIDOk returns a tuple with the MessageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetMessageIDOk() (*string, bool) {
	if o == nil || o.MessageID == nil {
		return nil, false
	}
	return o.MessageID, true
}

// HasMessageID returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasMessageID() bool {
	if o != nil && o.MessageID != nil {
		return true
	}

	return false
}

// SetMessageID gets a reference to the given string and assigns it to the MessageID field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetMessageID(v string) {
	o.MessageID = &v
}

// GetAdapterInputVersion returns the AdapterInputVersion field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterInputVersion() string {
	if o == nil || o.AdapterInputVersion == nil {
		var ret string
		return ret
	}
	return *o.AdapterInputVersion
}

// GetAdapterInputVersionOk returns a tuple with the AdapterInputVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterInputVersionOk() (*string, bool) {
	if o == nil || o.AdapterInputVersion == nil {
		return nil, false
	}
	return o.AdapterInputVersion, true
}

// HasAdapterInputVersion returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasAdapterInputVersion() bool {
	if o != nil && o.AdapterInputVersion != nil {
		return true
	}

	return false
}

// SetAdapterInputVersion gets a reference to the given string and assigns it to the AdapterInputVersion field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetAdapterInputVersion(v string) {
	o.AdapterInputVersion = &v
}

// GetAdapterInputAddress returns the AdapterInputAddress field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterInputAddress() string {
	if o == nil || o.AdapterInputAddress == nil {
		var ret string
		return ret
	}
	return *o.AdapterInputAddress
}

// GetAdapterInputAddressOk returns a tuple with the AdapterInputAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterInputAddressOk() (*string, bool) {
	if o == nil || o.AdapterInputAddress == nil {
		return nil, false
	}
	return o.AdapterInputAddress, true
}

// HasAdapterInputAddress returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasAdapterInputAddress() bool {
	if o != nil && o.AdapterInputAddress != nil {
		return true
	}

	return false
}

// SetAdapterInputAddress gets a reference to the given string and assigns it to the AdapterInputAddress field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetAdapterInputAddress(v string) {
	o.AdapterInputAddress = &v
}

// GetAdapterInputNodeName returns the AdapterInputNodeName field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterInputNodeName() string {
	if o == nil || o.AdapterInputNodeName == nil {
		var ret string
		return ret
	}
	return *o.AdapterInputNodeName
}

// GetAdapterInputNodeNameOk returns a tuple with the AdapterInputNodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterInputNodeNameOk() (*string, bool) {
	if o == nil || o.AdapterInputNodeName == nil {
		return nil, false
	}
	return o.AdapterInputNodeName, true
}

// HasAdapterInputNodeName returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasAdapterInputNodeName() bool {
	if o != nil && o.AdapterInputNodeName != nil {
		return true
	}

	return false
}

// SetAdapterInputNodeName gets a reference to the given string and assigns it to the AdapterInputNodeName field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetAdapterInputNodeName(v string) {
	o.AdapterInputNodeName = &v
}

// GetAdapterOutputVersion returns the AdapterOutputVersion field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterOutputVersion() string {
	if o == nil || o.AdapterOutputVersion == nil {
		var ret string
		return ret
	}
	return *o.AdapterOutputVersion
}

// GetAdapterOutputVersionOk returns a tuple with the AdapterOutputVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterOutputVersionOk() (*string, bool) {
	if o == nil || o.AdapterOutputVersion == nil {
		return nil, false
	}
	return o.AdapterOutputVersion, true
}

// HasAdapterOutputVersion returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasAdapterOutputVersion() bool {
	if o != nil && o.AdapterOutputVersion != nil {
		return true
	}

	return false
}

// SetAdapterOutputVersion gets a reference to the given string and assigns it to the AdapterOutputVersion field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetAdapterOutputVersion(v string) {
	o.AdapterOutputVersion = &v
}

// GetAdapterOutputAddress returns the AdapterOutputAddress field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterOutputAddress() string {
	if o == nil || o.AdapterOutputAddress == nil {
		var ret string
		return ret
	}
	return *o.AdapterOutputAddress
}

// GetAdapterOutputAddressOk returns a tuple with the AdapterOutputAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterOutputAddressOk() (*string, bool) {
	if o == nil || o.AdapterOutputAddress == nil {
		return nil, false
	}
	return o.AdapterOutputAddress, true
}

// HasAdapterOutputAddress returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasAdapterOutputAddress() bool {
	if o != nil && o.AdapterOutputAddress != nil {
		return true
	}

	return false
}

// SetAdapterOutputAddress gets a reference to the given string and assigns it to the AdapterOutputAddress field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetAdapterOutputAddress(v string) {
	o.AdapterOutputAddress = &v
}

// GetAdapterOutputNodeName returns the AdapterOutputNodeName field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterOutputNodeName() string {
	if o == nil || o.AdapterOutputNodeName == nil {
		var ret string
		return ret
	}
	return *o.AdapterOutputNodeName
}

// GetAdapterOutputNodeNameOk returns a tuple with the AdapterOutputNodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdapterOutputNodeNameOk() (*string, bool) {
	if o == nil || o.AdapterOutputNodeName == nil {
		return nil, false
	}
	return o.AdapterOutputNodeName, true
}

// HasAdapterOutputNodeName returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasAdapterOutputNodeName() bool {
	if o != nil && o.AdapterOutputNodeName != nil {
		return true
	}

	return false
}

// SetAdapterOutputNodeName gets a reference to the given string and assigns it to the AdapterOutputNodeName field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetAdapterOutputNodeName(v string) {
	o.AdapterOutputNodeName = &v
}

// GetServiceVersion returns the ServiceVersion field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetServiceVersion() string {
	if o == nil || o.ServiceVersion == nil {
		var ret string
		return ret
	}
	return *o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetServiceVersionOk() (*string, bool) {
	if o == nil || o.ServiceVersion == nil {
		return nil, false
	}
	return o.ServiceVersion, true
}

// HasServiceVersion returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasServiceVersion() bool {
	if o != nil && o.ServiceVersion != nil {
		return true
	}

	return false
}

// SetServiceVersion gets a reference to the given string and assigns it to the ServiceVersion field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetServiceVersion(v string) {
	o.ServiceVersion = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetSequence() string {
	if o == nil || o.Sequence == nil {
		var ret string
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetSequenceOk() (*string, bool) {
	if o == nil || o.Sequence == nil {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasSequence() bool {
	if o != nil && o.Sequence != nil {
		return true
	}

	return false
}

// SetSequence gets a reference to the given string and assigns it to the Sequence field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetSequence(v string) {
	o.Sequence = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetSecurity() []SaleObjectSaleSecurity {
	if o == nil || o.Security == nil {
		var ret []SaleObjectSaleSecurity
		return ret
	}
	return o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetSecurityOk() ([]SaleObjectSaleSecurity, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given []SaleObjectSaleSecurity and assigns it to the Security field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetSecurity(v []SaleObjectSaleSecurity) {
	o.Security = v
}

// GetWasReversePrevious returns the WasReversePrevious field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetWasReversePrevious() int32 {
	if o == nil || o.WasReversePrevious == nil {
		var ret int32
		return ret
	}
	return *o.WasReversePrevious
}

// GetWasReversePreviousOk returns a tuple with the WasReversePrevious field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetWasReversePreviousOk() (*int32, bool) {
	if o == nil || o.WasReversePrevious == nil {
		return nil, false
	}
	return o.WasReversePrevious, true
}

// HasWasReversePrevious returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasWasReversePrevious() bool {
	if o != nil && o.WasReversePrevious != nil {
		return true
	}

	return false
}

// SetWasReversePrevious gets a reference to the given int32 and assigns it to the WasReversePrevious field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetWasReversePrevious(v int32) {
	o.WasReversePrevious = &v
}

// GetWasClosedPreviousBlock returns the WasClosedPreviousBlock field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetWasClosedPreviousBlock() int32 {
	if o == nil || o.WasClosedPreviousBlock == nil {
		var ret int32
		return ret
	}
	return *o.WasClosedPreviousBlock
}

// GetWasClosedPreviousBlockOk returns a tuple with the WasClosedPreviousBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetWasClosedPreviousBlockOk() (*int32, bool) {
	if o == nil || o.WasClosedPreviousBlock == nil {
		return nil, false
	}
	return o.WasClosedPreviousBlock, true
}

// HasWasClosedPreviousBlock returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasWasClosedPreviousBlock() bool {
	if o != nil && o.WasClosedPreviousBlock != nil {
		return true
	}

	return false
}

// SetWasClosedPreviousBlock gets a reference to the given int32 and assigns it to the WasClosedPreviousBlock field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetWasClosedPreviousBlock(v int32) {
	o.WasClosedPreviousBlock = &v
}

// GetReversedAnswerKey returns the ReversedAnswerKey field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetReversedAnswerKey() string {
	if o == nil || o.ReversedAnswerKey == nil {
		var ret string
		return ret
	}
	return *o.ReversedAnswerKey
}

// GetReversedAnswerKeyOk returns a tuple with the ReversedAnswerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetReversedAnswerKeyOk() (*string, bool) {
	if o == nil || o.ReversedAnswerKey == nil {
		return nil, false
	}
	return o.ReversedAnswerKey, true
}

// HasReversedAnswerKey returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasReversedAnswerKey() bool {
	if o != nil && o.ReversedAnswerKey != nil {
		return true
	}

	return false
}

// SetReversedAnswerKey gets a reference to the given string and assigns it to the ReversedAnswerKey field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetReversedAnswerKey(v string) {
	o.ReversedAnswerKey = &v
}

// GetReversedSequence returns the ReversedSequence field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetReversedSequence() string {
	if o == nil || o.ReversedSequence == nil {
		var ret string
		return ret
	}
	return *o.ReversedSequence
}

// GetReversedSequenceOk returns a tuple with the ReversedSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetReversedSequenceOk() (*string, bool) {
	if o == nil || o.ReversedSequence == nil {
		return nil, false
	}
	return o.ReversedSequence, true
}

// HasReversedSequence returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasReversedSequence() bool {
	if o != nil && o.ReversedSequence != nil {
		return true
	}

	return false
}

// SetReversedSequence gets a reference to the given string and assigns it to the ReversedSequence field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetReversedSequence(v string) {
	o.ReversedSequence = &v
}

// GetCommittedBlock returns the CommittedBlock field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetCommittedBlock() string {
	if o == nil || o.CommittedBlock == nil {
		var ret string
		return ret
	}
	return *o.CommittedBlock
}

// GetCommittedBlockOk returns a tuple with the CommittedBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetCommittedBlockOk() (*string, bool) {
	if o == nil || o.CommittedBlock == nil {
		return nil, false
	}
	return o.CommittedBlock, true
}

// HasCommittedBlock returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasCommittedBlock() bool {
	if o != nil && o.CommittedBlock != nil {
		return true
	}

	return false
}

// SetCommittedBlock gets a reference to the given string and assigns it to the CommittedBlock field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetCommittedBlock(v string) {
	o.CommittedBlock = &v
}

// GetReversedBlock returns the ReversedBlock field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetReversedBlock() string {
	if o == nil || o.ReversedBlock == nil {
		var ret string
		return ret
	}
	return *o.ReversedBlock
}

// GetReversedBlockOk returns a tuple with the ReversedBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetReversedBlockOk() (*string, bool) {
	if o == nil || o.ReversedBlock == nil {
		return nil, false
	}
	return o.ReversedBlock, true
}

// HasReversedBlock returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasReversedBlock() bool {
	if o != nil && o.ReversedBlock != nil {
		return true
	}

	return false
}

// SetReversedBlock gets a reference to the given string and assigns it to the ReversedBlock field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetReversedBlock(v string) {
	o.ReversedBlock = &v
}

// GetRequiredInformation returns the RequiredInformation field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequiredInformation() []DebtPaymentObjectDebtPaymentRequiredInformation {
	if o == nil || o.RequiredInformation == nil {
		var ret []DebtPaymentObjectDebtPaymentRequiredInformation
		return ret
	}
	return o.RequiredInformation
}

// GetRequiredInformationOk returns a tuple with the RequiredInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequiredInformationOk() ([]DebtPaymentObjectDebtPaymentRequiredInformation, bool) {
	if o == nil || o.RequiredInformation == nil {
		return nil, false
	}
	return o.RequiredInformation, true
}

// HasRequiredInformation returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasRequiredInformation() bool {
	if o != nil && o.RequiredInformation != nil {
		return true
	}

	return false
}

// SetRequiredInformation gets a reference to the given []DebtPaymentObjectDebtPaymentRequiredInformation and assigns it to the RequiredInformation field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetRequiredInformation(v []DebtPaymentObjectDebtPaymentRequiredInformation) {
	o.RequiredInformation = v
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation {
	if o == nil || o.AdditionalInformation == nil {
		var ret []SaleResponseObjectSaleResponseAdditionalInformation
		return ret
	}
	return o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAdditionalInformationOk() ([]SaleResponseObjectSaleResponseAdditionalInformation, bool) {
	if o == nil || o.AdditionalInformation == nil {
		return nil, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasAdditionalInformation() bool {
	if o != nil && o.AdditionalInformation != nil {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given []SaleResponseObjectSaleResponseAdditionalInformation and assigns it to the AdditionalInformation field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation) {
	o.AdditionalInformation = v
}

// GetDisplayResponseMessage returns the DisplayResponseMessage field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetDisplayResponseMessage() []string {
	if o == nil || o.DisplayResponseMessage == nil {
		var ret []string
		return ret
	}
	return o.DisplayResponseMessage
}

// GetDisplayResponseMessageOk returns a tuple with the DisplayResponseMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetDisplayResponseMessageOk() ([]string, bool) {
	if o == nil || o.DisplayResponseMessage == nil {
		return nil, false
	}
	return o.DisplayResponseMessage, true
}

// HasDisplayResponseMessage returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasDisplayResponseMessage() bool {
	if o != nil && o.DisplayResponseMessage != nil {
		return true
	}

	return false
}

// SetDisplayResponseMessage gets a reference to the given []string and assigns it to the DisplayResponseMessage field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetDisplayResponseMessage(v []string) {
	o.DisplayResponseMessage = v
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetBlock() string {
	if o == nil || o.Block == nil {
		var ret string
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetBlockOk() (*string, bool) {
	if o == nil || o.Block == nil {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasBlock() bool {
	if o != nil && o.Block != nil {
		return true
	}

	return false
}

// SetBlock gets a reference to the given string and assigns it to the Block field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetBlock(v string) {
	o.Block = &v
}

// GetTransmitionDateTime returns the TransmitionDateTime field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetTransmitionDateTime() time.Time {
	if o == nil || o.TransmitionDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.TransmitionDateTime
}

// GetTransmitionDateTimeOk returns a tuple with the TransmitionDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetTransmitionDateTimeOk() (*time.Time, bool) {
	if o == nil || o.TransmitionDateTime == nil {
		return nil, false
	}
	return o.TransmitionDateTime, true
}

// HasTransmitionDateTime returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasTransmitionDateTime() bool {
	if o != nil && o.TransmitionDateTime != nil {
		return true
	}

	return false
}

// SetTransmitionDateTime gets a reference to the given time.Time and assigns it to the TransmitionDateTime field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetTransmitionDateTime(v time.Time) {
	o.TransmitionDateTime = &v
}

// GetTransactionIdentification returns the TransactionIdentification field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetTransactionIdentification() string {
	if o == nil || o.TransactionIdentification == nil {
		var ret string
		return ret
	}
	return *o.TransactionIdentification
}

// GetTransactionIdentificationOk returns a tuple with the TransactionIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetTransactionIdentificationOk() (*string, bool) {
	if o == nil || o.TransactionIdentification == nil {
		return nil, false
	}
	return o.TransactionIdentification, true
}

// HasTransactionIdentification returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasTransactionIdentification() bool {
	if o != nil && o.TransactionIdentification != nil {
		return true
	}

	return false
}

// SetTransactionIdentification gets a reference to the given string and assigns it to the TransactionIdentification field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetTransactionIdentification(v string) {
	o.TransactionIdentification = &v
}

// GetTransactionDescription returns the TransactionDescription field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetTransactionDescription() string {
	if o == nil || o.TransactionDescription == nil {
		var ret string
		return ret
	}
	return *o.TransactionDescription
}

// GetTransactionDescriptionOk returns a tuple with the TransactionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetTransactionDescriptionOk() (*string, bool) {
	if o == nil || o.TransactionDescription == nil {
		return nil, false
	}
	return o.TransactionDescription, true
}

// HasTransactionDescription returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasTransactionDescription() bool {
	if o != nil && o.TransactionDescription != nil {
		return true
	}

	return false
}

// SetTransactionDescription gets a reference to the given string and assigns it to the TransactionDescription field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetTransactionDescription(v string) {
	o.TransactionDescription = &v
}

// GetTrasactionDateTime returns the TrasactionDateTime field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetTrasactionDateTime() string {
	if o == nil || o.TrasactionDateTime == nil {
		var ret string
		return ret
	}
	return *o.TrasactionDateTime
}

// GetTrasactionDateTimeOk returns a tuple with the TrasactionDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetTrasactionDateTimeOk() (*string, bool) {
	if o == nil || o.TrasactionDateTime == nil {
		return nil, false
	}
	return o.TrasactionDateTime, true
}

// HasTrasactionDateTime returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasTrasactionDateTime() bool {
	if o != nil && o.TrasactionDateTime != nil {
		return true
	}

	return false
}

// SetTrasactionDateTime gets a reference to the given string and assigns it to the TrasactionDateTime field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetTrasactionDateTime(v string) {
	o.TrasactionDateTime = &v
}

// GetTaxFinancialCostAmount returns the TaxFinancialCostAmount field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetTaxFinancialCostAmount() float32 {
	if o == nil || o.TaxFinancialCostAmount == nil {
		var ret float32
		return ret
	}
	return *o.TaxFinancialCostAmount
}

// GetTaxFinancialCostAmountOk returns a tuple with the TaxFinancialCostAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetTaxFinancialCostAmountOk() (*float32, bool) {
	if o == nil || o.TaxFinancialCostAmount == nil {
		return nil, false
	}
	return o.TaxFinancialCostAmount, true
}

// HasTaxFinancialCostAmount returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasTaxFinancialCostAmount() bool {
	if o != nil && o.TaxFinancialCostAmount != nil {
		return true
	}

	return false
}

// SetTaxFinancialCostAmount gets a reference to the given float32 and assigns it to the TaxFinancialCostAmount field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetTaxFinancialCostAmount(v float32) {
	o.TaxFinancialCostAmount = &v
}

// GetTaxFinancialCostPercentage returns the TaxFinancialCostPercentage field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetTaxFinancialCostPercentage() float32 {
	if o == nil || o.TaxFinancialCostPercentage == nil {
		var ret float32
		return ret
	}
	return *o.TaxFinancialCostPercentage
}

// GetTaxFinancialCostPercentageOk returns a tuple with the TaxFinancialCostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetTaxFinancialCostPercentageOk() (*float32, bool) {
	if o == nil || o.TaxFinancialCostPercentage == nil {
		return nil, false
	}
	return o.TaxFinancialCostPercentage, true
}

// HasTaxFinancialCostPercentage returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasTaxFinancialCostPercentage() bool {
	if o != nil && o.TaxFinancialCostPercentage != nil {
		return true
	}

	return false
}

// SetTaxFinancialCostPercentage gets a reference to the given float32 and assigns it to the TaxFinancialCostPercentage field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetTaxFinancialCostPercentage(v float32) {
	o.TaxFinancialCostPercentage = &v
}

// GetFinancialCostAmount returns the FinancialCostAmount field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetFinancialCostAmount() float32 {
	if o == nil || o.FinancialCostAmount == nil {
		var ret float32
		return ret
	}
	return *o.FinancialCostAmount
}

// GetFinancialCostAmountOk returns a tuple with the FinancialCostAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetFinancialCostAmountOk() (*float32, bool) {
	if o == nil || o.FinancialCostAmount == nil {
		return nil, false
	}
	return o.FinancialCostAmount, true
}

// HasFinancialCostAmount returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasFinancialCostAmount() bool {
	if o != nil && o.FinancialCostAmount != nil {
		return true
	}

	return false
}

// SetFinancialCostAmount gets a reference to the given float32 and assigns it to the FinancialCostAmount field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetFinancialCostAmount(v float32) {
	o.FinancialCostAmount = &v
}

// GetFinancialCostPercentage returns the FinancialCostPercentage field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetFinancialCostPercentage() float32 {
	if o == nil || o.FinancialCostPercentage == nil {
		var ret float32
		return ret
	}
	return *o.FinancialCostPercentage
}

// GetFinancialCostPercentageOk returns a tuple with the FinancialCostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetFinancialCostPercentageOk() (*float32, bool) {
	if o == nil || o.FinancialCostPercentage == nil {
		return nil, false
	}
	return o.FinancialCostPercentage, true
}

// HasFinancialCostPercentage returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasFinancialCostPercentage() bool {
	if o != nil && o.FinancialCostPercentage != nil {
		return true
	}

	return false
}

// SetFinancialCostPercentage gets a reference to the given float32 and assigns it to the FinancialCostPercentage field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetFinancialCostPercentage(v float32) {
	o.FinancialCostPercentage = &v
}

// GetRequestAmount returns the RequestAmount field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequestAmount() float32 {
	if o == nil || o.RequestAmount == nil {
		var ret float32
		return ret
	}
	return *o.RequestAmount
}

// GetRequestAmountOk returns a tuple with the RequestAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetRequestAmountOk() (*float32, bool) {
	if o == nil || o.RequestAmount == nil {
		return nil, false
	}
	return o.RequestAmount, true
}

// HasRequestAmount returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasRequestAmount() bool {
	if o != nil && o.RequestAmount != nil {
		return true
	}

	return false
}

// SetRequestAmount gets a reference to the given float32 and assigns it to the RequestAmount field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetRequestAmount(v float32) {
	o.RequestAmount = &v
}

// GetWalletRequestToken returns the WalletRequestToken field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetWalletRequestToken() string {
	if o == nil || o.WalletRequestToken == nil {
		var ret string
		return ret
	}
	return *o.WalletRequestToken
}

// GetWalletRequestTokenOk returns a tuple with the WalletRequestToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetWalletRequestTokenOk() (*string, bool) {
	if o == nil || o.WalletRequestToken == nil {
		return nil, false
	}
	return o.WalletRequestToken, true
}

// HasWalletRequestToken returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasWalletRequestToken() bool {
	if o != nil && o.WalletRequestToken != nil {
		return true
	}

	return false
}

// SetWalletRequestToken gets a reference to the given string and assigns it to the WalletRequestToken field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetWalletRequestToken(v string) {
	o.WalletRequestToken = &v
}

// GetWalletRequestTokenType returns the WalletRequestTokenType field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetWalletRequestTokenType() string {
	if o == nil || o.WalletRequestTokenType == nil {
		var ret string
		return ret
	}
	return *o.WalletRequestTokenType
}

// GetWalletRequestTokenTypeOk returns a tuple with the WalletRequestTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetWalletRequestTokenTypeOk() (*string, bool) {
	if o == nil || o.WalletRequestTokenType == nil {
		return nil, false
	}
	return o.WalletRequestTokenType, true
}

// HasWalletRequestTokenType returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasWalletRequestTokenType() bool {
	if o != nil && o.WalletRequestTokenType != nil {
		return true
	}

	return false
}

// SetWalletRequestTokenType gets a reference to the given string and assigns it to the WalletRequestTokenType field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetWalletRequestTokenType(v string) {
	o.WalletRequestTokenType = &v
}

// GetPaymentMethodIdentification returns the PaymentMethodIdentification field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetPaymentMethodIdentification() string {
	if o == nil || o.PaymentMethodIdentification == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodIdentification
}

// GetPaymentMethodIdentificationOk returns a tuple with the PaymentMethodIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetPaymentMethodIdentificationOk() (*string, bool) {
	if o == nil || o.PaymentMethodIdentification == nil {
		return nil, false
	}
	return o.PaymentMethodIdentification, true
}

// HasPaymentMethodIdentification returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasPaymentMethodIdentification() bool {
	if o != nil && o.PaymentMethodIdentification != nil {
		return true
	}

	return false
}

// SetPaymentMethodIdentification gets a reference to the given string and assigns it to the PaymentMethodIdentification field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetPaymentMethodIdentification(v string) {
	o.PaymentMethodIdentification = &v
}

// GetForeignPaymentMethodIdentification returns the ForeignPaymentMethodIdentification field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetForeignPaymentMethodIdentification() string {
	if o == nil || o.ForeignPaymentMethodIdentification == nil {
		var ret string
		return ret
	}
	return *o.ForeignPaymentMethodIdentification
}

// GetForeignPaymentMethodIdentificationOk returns a tuple with the ForeignPaymentMethodIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetForeignPaymentMethodIdentificationOk() (*string, bool) {
	if o == nil || o.ForeignPaymentMethodIdentification == nil {
		return nil, false
	}
	return o.ForeignPaymentMethodIdentification, true
}

// HasForeignPaymentMethodIdentification returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasForeignPaymentMethodIdentification() bool {
	if o != nil && o.ForeignPaymentMethodIdentification != nil {
		return true
	}

	return false
}

// SetForeignPaymentMethodIdentification gets a reference to the given string and assigns it to the ForeignPaymentMethodIdentification field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetForeignPaymentMethodIdentification(v string) {
	o.ForeignPaymentMethodIdentification = &v
}

// GetIssuerIdentification returns the IssuerIdentification field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetIssuerIdentification() string {
	if o == nil || o.IssuerIdentification == nil {
		var ret string
		return ret
	}
	return *o.IssuerIdentification
}

// GetIssuerIdentificationOk returns a tuple with the IssuerIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetIssuerIdentificationOk() (*string, bool) {
	if o == nil || o.IssuerIdentification == nil {
		return nil, false
	}
	return o.IssuerIdentification, true
}

// HasIssuerIdentification returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasIssuerIdentification() bool {
	if o != nil && o.IssuerIdentification != nil {
		return true
	}

	return false
}

// SetIssuerIdentification gets a reference to the given string and assigns it to the IssuerIdentification field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetIssuerIdentification(v string) {
	o.IssuerIdentification = &v
}

// GetForeignIssuerIdentification returns the ForeignIssuerIdentification field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetForeignIssuerIdentification() string {
	if o == nil || o.ForeignIssuerIdentification == nil {
		var ret string
		return ret
	}
	return *o.ForeignIssuerIdentification
}

// GetForeignIssuerIdentificationOk returns a tuple with the ForeignIssuerIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetForeignIssuerIdentificationOk() (*string, bool) {
	if o == nil || o.ForeignIssuerIdentification == nil {
		return nil, false
	}
	return o.ForeignIssuerIdentification, true
}

// HasForeignIssuerIdentification returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasForeignIssuerIdentification() bool {
	if o != nil && o.ForeignIssuerIdentification != nil {
		return true
	}

	return false
}

// SetForeignIssuerIdentification gets a reference to the given string and assigns it to the ForeignIssuerIdentification field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetForeignIssuerIdentification(v string) {
	o.ForeignIssuerIdentification = &v
}

// GetDisplayCredentialIdentification returns the DisplayCredentialIdentification field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetDisplayCredentialIdentification() string {
	if o == nil || o.DisplayCredentialIdentification == nil {
		var ret string
		return ret
	}
	return *o.DisplayCredentialIdentification
}

// GetDisplayCredentialIdentificationOk returns a tuple with the DisplayCredentialIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetDisplayCredentialIdentificationOk() (*string, bool) {
	if o == nil || o.DisplayCredentialIdentification == nil {
		return nil, false
	}
	return o.DisplayCredentialIdentification, true
}

// HasDisplayCredentialIdentification returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasDisplayCredentialIdentification() bool {
	if o != nil && o.DisplayCredentialIdentification != nil {
		return true
	}

	return false
}

// SetDisplayCredentialIdentification gets a reference to the given string and assigns it to the DisplayCredentialIdentification field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetDisplayCredentialIdentification(v string) {
	o.DisplayCredentialIdentification = &v
}

// GetAuthTicket returns the AuthTicket field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAuthTicket() int32 {
	if o == nil || o.AuthTicket == nil {
		var ret int32
		return ret
	}
	return *o.AuthTicket
}

// GetAuthTicketOk returns a tuple with the AuthTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAuthTicketOk() (*int32, bool) {
	if o == nil || o.AuthTicket == nil {
		return nil, false
	}
	return o.AuthTicket, true
}

// HasAuthTicket returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasAuthTicket() bool {
	if o != nil && o.AuthTicket != nil {
		return true
	}

	return false
}

// SetAuthTicket gets a reference to the given int32 and assigns it to the AuthTicket field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetAuthTicket(v int32) {
	o.AuthTicket = &v
}

// GetAuthRRN returns the AuthRRN field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAuthRRN() string {
	if o == nil || o.AuthRRN == nil {
		var ret string
		return ret
	}
	return *o.AuthRRN
}

// GetAuthRRNOk returns a tuple with the AuthRRN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetAuthRRNOk() (*string, bool) {
	if o == nil || o.AuthRRN == nil {
		return nil, false
	}
	return o.AuthRRN, true
}

// HasAuthRRN returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasAuthRRN() bool {
	if o != nil && o.AuthRRN != nil {
		return true
	}

	return false
}

// SetAuthRRN gets a reference to the given string and assigns it to the AuthRRN field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetAuthRRN(v string) {
	o.AuthRRN = &v
}

// GetIdentifierForTheAdquirer returns the IdentifierForTheAdquirer field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetIdentifierForTheAdquirer() string {
	if o == nil || o.IdentifierForTheAdquirer == nil {
		var ret string
		return ret
	}
	return *o.IdentifierForTheAdquirer
}

// GetIdentifierForTheAdquirerOk returns a tuple with the IdentifierForTheAdquirer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetIdentifierForTheAdquirerOk() (*string, bool) {
	if o == nil || o.IdentifierForTheAdquirer == nil {
		return nil, false
	}
	return o.IdentifierForTheAdquirer, true
}

// HasIdentifierForTheAdquirer returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasIdentifierForTheAdquirer() bool {
	if o != nil && o.IdentifierForTheAdquirer != nil {
		return true
	}

	return false
}

// SetIdentifierForTheAdquirer gets a reference to the given string and assigns it to the IdentifierForTheAdquirer field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetIdentifierForTheAdquirer(v string) {
	o.IdentifierForTheAdquirer = &v
}

// GetIdentifierForTheResolutor returns the IdentifierForTheResolutor field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetIdentifierForTheResolutor() string {
	if o == nil || o.IdentifierForTheResolutor == nil {
		var ret string
		return ret
	}
	return *o.IdentifierForTheResolutor
}

// GetIdentifierForTheResolutorOk returns a tuple with the IdentifierForTheResolutor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetIdentifierForTheResolutorOk() (*string, bool) {
	if o == nil || o.IdentifierForTheResolutor == nil {
		return nil, false
	}
	return o.IdentifierForTheResolutor, true
}

// HasIdentifierForTheResolutor returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasIdentifierForTheResolutor() bool {
	if o != nil && o.IdentifierForTheResolutor != nil {
		return true
	}

	return false
}

// SetIdentifierForTheResolutor gets a reference to the given string and assigns it to the IdentifierForTheResolutor field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetIdentifierForTheResolutor(v string) {
	o.IdentifierForTheResolutor = &v
}

// GetRetryTime returns the RetryTime field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetRetryTime() int32 {
	if o == nil || o.RetryTime == nil {
		var ret int32
		return ret
	}
	return *o.RetryTime
}

// GetRetryTimeOk returns a tuple with the RetryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetRetryTimeOk() (*int32, bool) {
	if o == nil || o.RetryTime == nil {
		return nil, false
	}
	return o.RetryTime, true
}

// HasRetryTime returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasRetryTime() bool {
	if o != nil && o.RetryTime != nil {
		return true
	}

	return false
}

// SetRetryTime gets a reference to the given int32 and assigns it to the RetryTime field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetRetryTime(v int32) {
	o.RetryTime = &v
}

// GetMerchantCategory returns the MerchantCategory field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetMerchantCategory() SaleResponseObjectSaleResponseMerchantCategory {
	if o == nil || o.MerchantCategory == nil {
		var ret SaleResponseObjectSaleResponseMerchantCategory
		return ret
	}
	return *o.MerchantCategory
}

// GetMerchantCategoryOk returns a tuple with the MerchantCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetMerchantCategoryOk() (*SaleResponseObjectSaleResponseMerchantCategory, bool) {
	if o == nil || o.MerchantCategory == nil {
		return nil, false
	}
	return o.MerchantCategory, true
}

// HasMerchantCategory returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasMerchantCategory() bool {
	if o != nil && o.MerchantCategory != nil {
		return true
	}

	return false
}

// SetMerchantCategory gets a reference to the given SaleResponseObjectSaleResponseMerchantCategory and assigns it to the MerchantCategory field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetMerchantCategory(v SaleResponseObjectSaleResponseMerchantCategory) {
	o.MerchantCategory = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetConfiguration() SaleResponseObjectSaleResponseConfiguration {
	if o == nil || o.Configuration == nil {
		var ret SaleResponseObjectSaleResponseConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetConfigurationOk() (*SaleResponseObjectSaleResponseConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given SaleResponseObjectSaleResponseConfiguration and assigns it to the Configuration field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetConfiguration(v SaleResponseObjectSaleResponseConfiguration) {
	o.Configuration = &v
}

// GetTickets returns the Tickets field value if set, zero value otherwise.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetTickets() []SaleResponseObjectSaleResponseTickets {
	if o == nil || o.Tickets == nil {
		var ret []SaleResponseObjectSaleResponseTickets
		return ret
	}
	return o.Tickets
}

// GetTicketsOk returns a tuple with the Tickets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) GetTicketsOk() ([]SaleResponseObjectSaleResponseTickets, bool) {
	if o == nil || o.Tickets == nil {
		return nil, false
	}
	return o.Tickets, true
}

// HasTickets returns a boolean if a field has been set.
func (o *WalletRequestResponseObjectWalletRequestResponse) HasTickets() bool {
	if o != nil && o.Tickets != nil {
		return true
	}

	return false
}

// SetTickets gets a reference to the given []SaleResponseObjectSaleResponseTickets and assigns it to the Tickets field.
func (o *WalletRequestResponseObjectWalletRequestResponse) SetTickets(v []SaleResponseObjectSaleResponseTickets) {
	o.Tickets = v
}

func (o WalletRequestResponseObjectWalletRequestResponse) MarshalJSON() ([]byte, error) {
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
	if o.AnswerType != nil {
		toSerialize["AnswerType"] = o.AnswerType
	}
	if true {
		toSerialize["AnswerKey"] = o.AnswerKey
	}
	if o.RequestKey != nil {
		toSerialize["RequestKey"] = o.RequestKey
	}
	if o.ServerVersion != nil {
		toSerialize["ServerVersion"] = o.ServerVersion
	}
	if o.ServerAddress != nil {
		toSerialize["ServerAddress"] = o.ServerAddress
	}
	if o.ServerInstance != nil {
		toSerialize["ServerInstance"] = o.ServerInstance
	}
	if o.ServerNodeName != nil {
		toSerialize["ServerNodeName"] = o.ServerNodeName
	}
	if o.MessageID != nil {
		toSerialize["MessageID"] = o.MessageID
	}
	if o.AdapterInputVersion != nil {
		toSerialize["AdapterInputVersion"] = o.AdapterInputVersion
	}
	if o.AdapterInputAddress != nil {
		toSerialize["AdapterInputAddress"] = o.AdapterInputAddress
	}
	if o.AdapterInputNodeName != nil {
		toSerialize["AdapterInputNodeName"] = o.AdapterInputNodeName
	}
	if o.AdapterOutputVersion != nil {
		toSerialize["AdapterOutputVersion"] = o.AdapterOutputVersion
	}
	if o.AdapterOutputAddress != nil {
		toSerialize["AdapterOutputAddress"] = o.AdapterOutputAddress
	}
	if o.AdapterOutputNodeName != nil {
		toSerialize["AdapterOutputNodeName"] = o.AdapterOutputNodeName
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
	if o.WasReversePrevious != nil {
		toSerialize["WasReversePrevious"] = o.WasReversePrevious
	}
	if o.WasClosedPreviousBlock != nil {
		toSerialize["WasClosedPreviousBlock"] = o.WasClosedPreviousBlock
	}
	if o.ReversedAnswerKey != nil {
		toSerialize["ReversedAnswerKey"] = o.ReversedAnswerKey
	}
	if o.ReversedSequence != nil {
		toSerialize["ReversedSequence"] = o.ReversedSequence
	}
	if o.CommittedBlock != nil {
		toSerialize["CommittedBlock"] = o.CommittedBlock
	}
	if o.ReversedBlock != nil {
		toSerialize["ReversedBlock"] = o.ReversedBlock
	}
	if o.RequiredInformation != nil {
		toSerialize["RequiredInformation"] = o.RequiredInformation
	}
	if o.AdditionalInformation != nil {
		toSerialize["AdditionalInformation"] = o.AdditionalInformation
	}
	if o.DisplayResponseMessage != nil {
		toSerialize["DisplayResponseMessage"] = o.DisplayResponseMessage
	}
	if o.Block != nil {
		toSerialize["Block"] = o.Block
	}
	if o.TransmitionDateTime != nil {
		toSerialize["TransmitionDateTime"] = o.TransmitionDateTime
	}
	if o.TransactionIdentification != nil {
		toSerialize["TransactionIdentification"] = o.TransactionIdentification
	}
	if o.TransactionDescription != nil {
		toSerialize["TransactionDescription"] = o.TransactionDescription
	}
	if o.TrasactionDateTime != nil {
		toSerialize["TrasactionDateTime"] = o.TrasactionDateTime
	}
	if o.TaxFinancialCostAmount != nil {
		toSerialize["TaxFinancialCostAmount"] = o.TaxFinancialCostAmount
	}
	if o.TaxFinancialCostPercentage != nil {
		toSerialize["TaxFinancialCostPercentage"] = o.TaxFinancialCostPercentage
	}
	if o.FinancialCostAmount != nil {
		toSerialize["FinancialCostAmount"] = o.FinancialCostAmount
	}
	if o.FinancialCostPercentage != nil {
		toSerialize["FinancialCostPercentage"] = o.FinancialCostPercentage
	}
	if o.RequestAmount != nil {
		toSerialize["RequestAmount"] = o.RequestAmount
	}
	if o.WalletRequestToken != nil {
		toSerialize["WalletRequestToken"] = o.WalletRequestToken
	}
	if o.WalletRequestTokenType != nil {
		toSerialize["WalletRequestTokenType"] = o.WalletRequestTokenType
	}
	if o.PaymentMethodIdentification != nil {
		toSerialize["PaymentMethodIdentification"] = o.PaymentMethodIdentification
	}
	if o.ForeignPaymentMethodIdentification != nil {
		toSerialize["ForeignPaymentMethodIdentification"] = o.ForeignPaymentMethodIdentification
	}
	if o.IssuerIdentification != nil {
		toSerialize["IssuerIdentification"] = o.IssuerIdentification
	}
	if o.ForeignIssuerIdentification != nil {
		toSerialize["ForeignIssuerIdentification"] = o.ForeignIssuerIdentification
	}
	if o.DisplayCredentialIdentification != nil {
		toSerialize["DisplayCredentialIdentification"] = o.DisplayCredentialIdentification
	}
	if o.AuthTicket != nil {
		toSerialize["AuthTicket"] = o.AuthTicket
	}
	if o.AuthRRN != nil {
		toSerialize["AuthRRN"] = o.AuthRRN
	}
	if o.IdentifierForTheAdquirer != nil {
		toSerialize["IdentifierForTheAdquirer"] = o.IdentifierForTheAdquirer
	}
	if o.IdentifierForTheResolutor != nil {
		toSerialize["IdentifierForTheResolutor"] = o.IdentifierForTheResolutor
	}
	if o.RetryTime != nil {
		toSerialize["RetryTime"] = o.RetryTime
	}
	if o.MerchantCategory != nil {
		toSerialize["MerchantCategory"] = o.MerchantCategory
	}
	if o.Configuration != nil {
		toSerialize["Configuration"] = o.Configuration
	}
	if o.Tickets != nil {
		toSerialize["Tickets"] = o.Tickets
	}
	return json.Marshal(toSerialize)
}

type NullableWalletRequestResponseObjectWalletRequestResponse struct {
	value *WalletRequestResponseObjectWalletRequestResponse
	isSet bool
}

func (v NullableWalletRequestResponseObjectWalletRequestResponse) Get() *WalletRequestResponseObjectWalletRequestResponse {
	return v.value
}

func (v *NullableWalletRequestResponseObjectWalletRequestResponse) Set(val *WalletRequestResponseObjectWalletRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletRequestResponseObjectWalletRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletRequestResponseObjectWalletRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletRequestResponseObjectWalletRequestResponse(val *WalletRequestResponseObjectWalletRequestResponse) *NullableWalletRequestResponseObjectWalletRequestResponse {
	return &NullableWalletRequestResponseObjectWalletRequestResponse{value: val, isSet: true}
}

func (v NullableWalletRequestResponseObjectWalletRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletRequestResponseObjectWalletRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


