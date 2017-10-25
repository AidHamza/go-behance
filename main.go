package main

import (
	"net/http"
	"fmt"
	"time"
	"reflect"
	"io/ioutil"
	"encoding/json"
)

type Stats struct {
	Appreciations int `json:"appreciations"`
	Views int `json:"views"`
	Comments int `json:"comments"`
}

type Project struct {  
    Id int `json:"id"`
    Name string `json:"name"`
    URL string `json:"url"`
    Stat Stats `json:"stats"`
    Description string `json:"description"`
    Modules interface{} `json:"modules"`

}

type ProjectResult struct {
	Project Project `json:"project"`
}


//api_key
const BEHANCE_API = "https://www.behance.net/v2"
const ENDPOINT_PROJECTS = "/projects/57915943";
const ENDPOINT_COLLECTIONS = "/collections";

// Client is the http client used by the Fetch function. It can be customized as needed.
var Client = &http.Client{
	Timeout: 20 * time.Second,
}

func main() {
	var response *http.Response
	var err error
	var resultFinal ProjectResult

	clientId := ""

	response, err = Client.Get(BEHANCE_API + ENDPOINT_PROJECTS + "?api_key=" + clientId)
	defer func() {
		_ = response.Body.Close()
	}()

	if err != nil {
		fmt.Errorf("Error retreived from Behance API %+v", err)
	}

	if response.StatusCode != 200 {
		fmt.Errorf("http error: %s", response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("unable to read response body: %s", err)
	}

	err = json.Unmarshal(body, &resultFinal)
	if err != nil {
		fmt.Errorf("unable to Unmarshal JSON: %s", err)
	}

	kind := reflect.TypeOf(resultFinal.Project.Modules).Kind()
	fmt.Printf("Kind : %v\n", kind)
	fmt.Printf("Data : %v", resultFinal.Project.Modules)
	//s := reflect.ValueOf(resultFinal.Project.Modules)

	/*for moduleKey, moduleContent := range resultFinal.Project.Modules { 
    	fmt.Printf("Module ID [%s], Module Value : [%v]\n", moduleKey, moduleContent)
	}*/

}