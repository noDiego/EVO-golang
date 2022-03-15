package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"openapi"
	"os"
	"strings"
)

var client *openapi.APIClient

func main() {

	//asd
	configServer()

	//Seteando datos para llamada
	branchIdentification := "5"

	keepRequest := openapi.ApiKeepAlivePostRequest{}

	keepObject := openapi.KeepAliveObject{}
	keepData := openapi.BlockCreateObjectBlockCreate{
		CompanyIdentification: "1",
		SystemIdentification:  "0",
		BranchIdentification:  &branchIdentification,
		POSIdentification:     "1",
	}
	keepObject.SetKeepAlive(keepData)

	keepRequest.KeepAliveObject(keepObject)

	//Ejecutando llamada a keepalive

	_, response, err := client.PaymentApi.KeepAlivePost(context.Background()).KeepAliveObject(keepObject).Execute()

	if err != nil {
		println("Error al consultar KeepAlive")
		println(err.Error())
	}
	if response.StatusCode < 300 {
		fmt.Println("StatusCode:" + fmt.Sprint(response.StatusCode))

		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error en lectura de datos" + err.Error())
		}
		responseData := openapi.KeepAliveResponseObject{}

		json.Unmarshal(bodyBytes, &responseData)
		fmt.Println("ResponseCode: " + fmt.Sprint(responseData.KeepAliveResponse.GetResponseCode()))
		fmt.Println("ResponseMessage: " + responseData.KeepAliveResponse.GetResponseMessage())
	}
	return
}

func configServer() {

	conf := openapi.NewConfiguration()
	//Si vienen args
	if len(os.Args) > 1 && strings.Contains(os.Args[1], "--hostname=") {
		hostname := strings.Replace(os.Args[1], "--hostname=", "", 1)
		conf.Servers[0].URL = hostname
	}
	client = openapi.NewAPIClient(conf)
	println("Utilizando hostname: " + client.GetConfig().Servers[0].URL)
}
