package main

import (
	"fmt"
	"hc-infra-test/libraries"
)

/* Receive a token from Vault for the specified user - Axxxxxx
   Upon receiving the V-Vault-Token, send the token as a header to pull the secrets from Vault*/

func main() {

	XVaultToken := libraries.VaultConnection()
	fmt.Println(XVaultToken)

}
