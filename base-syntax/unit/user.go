package unit

import (
	"fmt"
	"strconv"
)

type User struct {
	id int
	name string
	like string
}

func GetUserInfo(id int) User {
	return User{
		id: id,
		name: "leeprince",
		like: "golang",
	}
}

func GetUserName(id int) string {
	s := fmt.Sprint(id) // 整型转化字符串
	ss := strconv.Itoa(id) // 整型转化字符串
	return fmt.Sprintf("id: %s - id: %s; name: %s", s, ss, "prince")
}
