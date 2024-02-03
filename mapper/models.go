package mapper

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type userInfo struct {
	Id       int    /*字段必须导出，否则 reflect 包无法访问*/
	password string `label:"krxk"`
	email    string
}

// MarshalJSON 对于非导出字段，json 包无法获取，此时可通过实现 MarshalJSON 接口解决
func (u userInfo) MarshalJSON() ([]byte, error) {
	// map 是无序的
	/*return json.Marshal(map[string]interface{}{
		"Id":       u.Id,
		"password": u.password,
		"email":    u.email,
	})*/

	// slice/array 是有序的(下面采用 slice，元素类型 为 interface{} )
	return json.Marshal([]interface{}{u.Id, u.password, u.email})
}

func Create() userInfo {
	u := userInfo{}
	return u
}

func (u userInfo) Set(id int, password string, email string) userInfo {
	u.Id = id
	u.password = password
	u.email = email
	return reflectIntercept(u)
	//return u
}

// reflectIntercept 通过反射方式处理结构体
func reflectIntercept(val userInfo) userInfo {
	fmt.Println("reflectIntercept is called")
	t := reflect.TypeOf(val)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type not Struct")
		return val
	}
	filedNum := t.NumField()
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Ptr { // 只能通过指针进行写操作
		v = reflect.ValueOf(&val).Elem()
	}
	for i := 0; i < filedNum; i++ {
		fmt.Println(t.Field(i).Name, ": ", v.Field(i), " Tag: ", t.Field(i).Tag) /*读取值*/
		if i == 0 && t.Field(i).IsExported() {
			fmt.Println("设置字段: ", t.Field(i).Name)
			v.Field(i).Set(reflect.ValueOf(4)) /*修改值*/
		}
	}

	/*动态生成新结构体测试*/
	var structFields []reflect.StructField
	for i := 0; i < filedNum; i++ {
		field := t.Field(i)
		fmt.Println(i, field)
		structFields = append(structFields, field)
	}
	field := reflect.StructField{
		Name:      "Krxk",
		PkgPath:   "",
		Type:      reflect.TypeOf(""),
		Tag:       "",
		Offset:    0,
		Index:     nil,
		Anonymous: false,
	}
	structFields = append(structFields, field)

	structType := reflect.StructOf(structFields)
	iFace := reflect.New(structType) /* iFace 为生成的自定义结构体 */
	iFace.Elem().FieldByName("Krxk").Set(reflect.ValueOf("Hello"))
	iFaceNew := iFace.Interface() /* 转换为接口类型 */

	fmt.Println("新接口 ", iFaceNew)

	return val
}
