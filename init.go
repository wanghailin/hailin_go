package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/wanghailin/hailin_go/source"
	"reflect"
	// "strings"
)

type HttpTest struct {
}

func main() {

	log.Fatal(http.ListenAndServe(":7777", HttpTest{}))
}

func (this HttpTest) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	routineMap , regStruct := source.InitRoutine()
	// fmt.Println(routineMap)
	// fmt.Println(regStruct)
	//获取访问的url
	urlString := req.URL.String()

	val, exist := routineMap[urlString]

	fmt.Println(val)

	if exist {
		controllerName :=val["controller"]
		actionName := val["action"]
		// fmt.Println(controllerName)
		// fmt.Println(actionName)

		controller := regStruct[controllerName]
		t := reflect.ValueOf(controller)
		// fmt.Println(t)
		methodValue := t.MethodByName(actionName)
		// fmt.Println(methodValue)
		args := make([]reflect.Value, 0)
		methodValue.Call(args)
	} else {
		fmt.Println("404")
	}


	fmt.Println()


}