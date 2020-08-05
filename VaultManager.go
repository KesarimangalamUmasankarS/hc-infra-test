package main

import (
	"fmt"
	"hc-infra-test/libraries"
	"reflect"
)

/* Receive a token from Vault for the specified user - Axxxxxx  */

func main() {

	var  x_vault_token string
	var secretval string
	secretlabel := libraries.GetData("label")
	path := libraries.GetData("secretpath")
	env := libraries.GetConfig("environment")
	url := libraries.GetConfig("url_"+env)

	XVaultToken_JSON  := libraries.VaultConnection()
	for _, jsonvalue := range XVaultToken_JSON {

		value := reflect.ValueOf(jsonvalue)
		switch value.Kind(){
		case reflect.Map:
			mapval := value.Interface().(map[string]interface{})
			if token_val, ok := mapval["client_token"]; ok {
				x_vault_token = reflect.ValueOf(token_val).Interface().(string)
			}
		}
	}

	/* Upon receiving the Vault-Token, send the token as a header to perform health check on the Vault system.
	   Status code 200 is considered active. Status code 429 is also considered active if the Vault system is behind a load balancer */


	health_check_status_code := libraries.GetHealth(x_vault_token)
	if(health_check_status_code != 200 && health_check_status_code != 429) {
		switch health_check_status_code{
		default:
			fmt.Println("Some issue with Vault and the status code returned is ", health_check_status_code)
		}
	} else {
		fmt.Println("Vault "+ env + " environment is active and the status code returned is ", health_check_status_code)
	}

	/* Vault token is also sent to retrieve secrets from the path specified in data.file */

	secret_JSON := libraries.GetSecret(x_vault_token)
	for _, jsonvalue := range secret_JSON {

		value := reflect.ValueOf(jsonvalue)
		switch value.Kind(){
		case reflect.Map:
			mapval := value.Interface().(map[string]interface{})
			if token_val, ok := mapval[secretlabel]; ok {
				secretval = reflect.ValueOf(token_val).Interface().(string)
				fmt.Println("\""+ secretlabel + "\" retrieved from " + url + path + " is "+secretval)
			}
		}
	}


}
