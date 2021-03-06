package libraries

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
    Send the Vault Token and retrieve  the secret from the  specified path
	Data can be pulled from data.file
*/

func GetSecret(x_vault_token string) map[string]interface{}{

	path := GetData("secretpath")
	env := GetConfig("environment")
	url := GetConfig("url_"+env)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url+path, nil)
	req.Header.Set("X-Vault-Token", x_vault_token)
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	var res map[string]interface{}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&res)
	client.CloseIdleConnections()
	return res
}

func GetData(key string) string{
	var value []string

	pwd, _ := os.Getwd()
	file, err := os.Open(pwd+"/data.file")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	for _, eachline := range txtlines {
		if strings.Contains(eachline, key){
			value = append(strings.Split(eachline, "="))
			break
		}
	}
	return value[1]
}

func  GetHealth(x_vault_token string) int {
	path := GetData("health_check_path")
	env := GetConfig("environment")
	url := GetConfig("url_"+env)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url+path, nil)
	req.Header.Set("X-Vault-Token", x_vault_token)
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	client.CloseIdleConnections()
	return resp.StatusCode
}