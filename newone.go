package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"sync"
)

//func to read the API's body
func getRespo1() {
	// json data
	url:="http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	//read the API's body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	bs := []byte(body)
	var data map[string]interface{}

	//Unmarshaling the body
	err = json.Unmarshal(bs, &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Results: %v\n", data)
	fmt.Println("Name:",data["name"])
	fmt.Println("Character:",data["character"])

	//convert the character array to display value
	s:=reflect.ValueOf(data["character"])
	for i:=0;i<s.Len();i++ {
		map1 := s.Index(i)

		//convert to map[string]interface{}
		x := map1.Interface().(map[string]interface{})
		name := x["name"].(string)
		max_power := x["max_power"].(float64)
		//storring in globally declared map
		charac_map[name] = max_power
	}
}
func getRespo2() {
	// json data
	url:="http://www.mocky.io/v2/5ecfd6473200009dc1e3d64e"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	//read the API's body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	bs := []byte(body)
	var data map[string]interface{}

	//Unmarshaling the body
	err = json.Unmarshal(bs, &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Results: %v\n", data)
	fmt.Println("Name:",data["name"])
	fmt.Println("Character:",data["character"])

	//convert the character array to display value
	s:=reflect.ValueOf(data["character"])
	for i:=0;i<s.Len();i++ {
		map1 := s.Index(i)

		//convert to map[string]interface{}
		x := map1.Interface().(map[string]interface{})
		name := x["name"].(string)
		max_power := x["max_power"].(float64)
		//storring in globally declared map
		charac_map[name] = max_power
	}

}
/*func getRespo3() {
	// json data
	url:="http://www.mocky.io/v2/5ecfd630320000f1aee3d64d"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var data interface{} // Marvels
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Results: %v\n", data)
}*/

//gloally declared map
var charac_map = make(map[string]float64,15)

//main function
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	//goroutines to run APIs in parallel
	go func(){
		getRespo1()
		wg.Done()
	}()
	go func(){
		getRespo2()
		wg.Done()
	}()
	//API has error : "Anti heroes"
	/*go func(){
		getRespo3()
	}()*/
	wg.Wait()

	//Display content of the map in console
	fmt.Println("CONTENT OF MAP")
	for j, v := range charac_map {
		fmt.Println(j, v)
	}
	handleRequests()
}

//handle function for localhost
func handleRequests() {
	http.HandleFunc("/get_maxpower/", printcontent)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

//function to return the value of query request in encoded json
func printcontent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the POWERLEVEL:")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}

	//creating query request
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		log.Println("no key found")
		return
	}
	//storing map's key to query key
	key := keys[0]
	v := charac_map[key]
	json.Unmarshal(reqBody, &v)
	json.NewEncoder(w).Encode(v)
	log.Println("URL key ",key)

	//print message in console everytime localhost is hit
	fmt.Println("Endpoint Hit: homePage")
}