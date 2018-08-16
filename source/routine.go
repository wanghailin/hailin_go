package source

import (
	// "fmt"
	"github.com/wanghailin/hailin_go/controller"

	// "hailinw/controller"
)

var regStruct map[string]interface{}

func controllerRegister() map[string]interface{}{
	regStruct = make(map[string]interface{})
	regStruct["User"] = controller.User{}
	return regStruct
}

func InitRoutine () (map[string]map[string]string, map[string]interface{}){
	 controllerRegister()
	 routines := make(map[string]map[string]string)
	 addRoutineMap(routines, "GET", "/user/info", "User", "Info")
	 addRoutineMap(routines, "POST", "/user/delete", "User", "Delete")
	 addRoutineMap(routines, "GET", "/user/name", "User", "Name")


	 controllerMap := controllerRegister()
	 return routines, controllerMap
}


func addRoutineMap(m map[string]map[string]string, httpMethod string, url string, controller string, action string) {
	mm, ok := m[url]
	if !ok {
		mm = make(map[string]string)
		mm["method"] = httpMethod
		mm["controller"] = controller
		mm["action"] = action
		m[url] = mm
	}
}


