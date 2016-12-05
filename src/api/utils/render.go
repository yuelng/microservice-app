package utils

//import (
//	"reflect"
//	"fmt"
//	"time"
//)

//func RenderJson(c *gin.Context, code int, obj interface{})  {
//	c.Status(code)
//
//	value := reflect.ValueOf(obj)
//	ft := time.Now().Add(time.Second*3600)
//	reflect.ValueOf(&obj).Elem().FieldByName("StartAt").Set(reflect.ValueOf(ft))
//	fmt.Println("type of p:", value.Type())
//	fmt.Println("settability of p:" , value.CanSet())
//
//	for i := 0; i < value.NumField(); i++ {
//		if value.Field(i).Type().Name() == "Time" {
//			// value.Field(i).SetString("Sunset Strip")
//			fmt.Println("settability of p:" , value.CanSet())
//			//sliceValue := reflect.ValueOf([]int{1, 2, 3})
//			//value.FieldByName(typeOfT.Field(i).Name).Set(sliceValue)
//		}
//	}
//	if err := render.WriteJSON(c.Writer, obj); err != nil {
//		panic(err)
//	}
//}
