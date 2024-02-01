package mapper

import "encoding/json"

type userInfo struct {
	id       int
	password string
	email    string
}

// MarshalJSON 对于非导出字段，json 包无法获取，此时可通过实现 MarshalJSON 接口解决
func (u userInfo) MarshalJSON() ([]byte, error) {
	// map 是无序的
	/*return json.Marshal(map[string]interface{}{
		"id":       u.id,
		"password": u.password,
		"email":    u.email,
	})*/

	// slice/array 是有序的(下面采用 slice，元素类型 为 interface{} )
	return json.Marshal([]interface{}{u.id, u.password, u.email})
}

func Create() userInfo {
	u := userInfo{}
	return u
}

func (u userInfo) Set(id int, password string, email string) userInfo {
	u.id = id
	u.password = password
	u.email = email
	return u
}
