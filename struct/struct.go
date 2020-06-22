package _struct

import "fmt"

func main() {
	stu := &User{
		name: "Tom",
	}
	//stu2 := new(User)
	msg := stu.hello("Jack")
	fmt.Println(msg) // hello Jack, I am Tom
}

type User struct {
	name string
	age  int
}

func (user *User) hello(name string) string {
	return fmt.Sprintf("hello %s, I am %s", user, user.name)

}
