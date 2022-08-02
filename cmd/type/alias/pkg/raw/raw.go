package raw

type MyType struct {
	Age  int
	Name string
}

func New() MyType {
	return MyType{Age: 1, Name: "1y"}
}
