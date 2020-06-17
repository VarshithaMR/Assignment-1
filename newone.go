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

func getRespo1() {
	// json data
	url:="http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	bs := []byte(body)
	var data map[string]interface{}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Results: %v\n", data)
	fmt.Println("Name:",data["name"])
	fmt.Println("Character:",data["character"])
	s:=reflect.ValueOf(data["character"])
	for i:=0;i<s.Len();i++ {
		map1 := s.Index(i)
		x := map1.Interface().(map[string]interface{})
		name := x["name"].(string)
		max_power := x["max_power"].(float64)
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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	bs := []byte(body)
	var data map[string]interface{}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Results: %v\n", data)
	fmt.Println("Name:",data["name"])
	fmt.Println("Character:",data["character"])
	s:=reflect.ValueOf(data["character"])
	for i:=0;i<s.Len();i++ {
		map1 := s.Index(i)
		x := map1.Interface().(map[string]interface{})
		name := x["name"].(string)
		max_power := x["max_power"].(float64)
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
var charac_map = make(map[string]float64,15)
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
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
	fmt.Println("CONTENT OF MAP")
	for j, v := range charac_map {
		fmt.Println(j, v)
	}
	handleRequests()
}
func handleRequests() {
	http.HandleFunc("/get_maxpower/", printcontent)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func printcontent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the POWERLEVEL:")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		log.Println("no key found")
		return
	}
	key := keys[0]
	v := charac_map[key]
	json.Unmarshal(reqBody, &v)
	json.NewEncoder(w).Encode(v)
	log.Println("URL key " + string(key))

	fmt.Println("Endpoint Hit: homePage")
}