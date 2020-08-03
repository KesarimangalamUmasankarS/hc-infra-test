package libraries

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func VaultConnection() map[string]interface{}{

	username := getConfig("username")
	env := getConfig("environment")
	url := getConfig("url_"+env)
	path := "auth/ldap/aeth/login"


	data := []byte(`{"password": ""}`)
	datavar := bytes.NewBuffer(data)
	client := &http.Client{}
	resp, err := client.Post(url+path+"/"+username, "application/json", datavar)

	if err != nil {
		panic(err)
	}
	var res map[string]interface{}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&res)
	client.CloseIdleConnections()
	return res
}

func getConfig(key string) string{
	var value []string

	pwd, _ := os.Getwd()
	file, err := os.Open(pwd+"/config.file")
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