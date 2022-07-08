package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/copier"
)

func main() {
	TestCopySameStructWithPointerField()
	pointerNoCopy()
	differentType()
}

type src struct {
	a int
	A string
	P *int
	T *time.Time
}

type User struct {
	Name     string
	Birthday *time.Time
	Nickname string
	Role     string
	Age      int32
	FakeAge  *int32
	Notes    []string
	flags    []byte
}

func TestCopySameStructWithPointerField() {
	var fakeAge int32 = 12
	var currentTime time.Time = time.Now()
	user := &User{Birthday: &currentTime, Name: "Jinzhu", Nickname: "jinzhu", Age: 18, FakeAge: &fakeAge, Role: "Admin", Notes: []string{"hello world", "welcome"}, flags: []byte{'x'}}
	newUser := &User{}
	copier.Copy(newUser, user)
	if user.Birthday == newUser.Birthday {
		panic("TestCopySameStructWithPointerField: copy Birthday failed since they need to have different address")
	}

	if user.FakeAge == newUser.FakeAge {
		panic("TestCopySameStructWithPointerField: copy FakeAge failed since they need to have different address")
	}
}

func pointerNoCopy() {
	// i := 1
	var i int = 12
	now := time.Now()
	s := src{
		a: 2,
		A: "hello",
		P: &i,
		T: &now,
	}

	var d src

	if err := copier.Copy(&d, &s); err != nil {
		panic(err)
	}

	if s.T == d.T {
		panic("time指针应该是不一样的")
	}

	if s.P == d.P {
		panic("两个地址应该不同")
	}
}

type d1 struct {
	A int
	B *time.Time
}

type d2 struct {
	A *time.Time
	B int
}

func differentType() {
	n := time.Now()
	v1 := d1{
		A: 100,
		B: &n,
	}
	var v2 d2
	copier.Copy(&v2, &v1)
	fmt.Printf("%+v", v2)
}
