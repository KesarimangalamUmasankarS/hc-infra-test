package main

import (
	"fmt"
	"hc-infra-test/libraries"
	"reflect"
)

/* Receive a token from Vault for the specified user - Axxxxxx
   Upon receiving the V-Vault-Token, send the token as a header to pull the secrets from Vault*/

func main() {

	var  x_vault_token string
	var secretval string
	secretlabel := libraries.GetData("secret")

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

	secret_JSON := libraries.GetSecret(x_vault_token)
	for _, jsonvalue := range secret_JSON {

		value := reflect.ValueOf(jsonvalue)
		switch value.Kind(){
		case reflect.Map:
			mapval := value.Interface().(map[string]interface{})
			if token_val, ok := mapval[secretlabel]; ok {
				secretval = reflect.ValueOf(token_val).Interface().(string)
				fmt.Println(secretval)
			}
		}
	}


}
