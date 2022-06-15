//go:build ignore
// +build ignore

package main

func main() {
}

// // Prepare some data to insert into the template.
// type Recipient struct {
// 	Name, Gift string
// 	Attended   bool
// }

// var recipients = []Recipient{
// 	{"Aunt Mildred", "bone china tea set", true},
// 	{"Uncle John", "moleskin pants", false},
// 	{"Cousin Rodney", "", false},
// }

// type Variant struct {
// 	// Name is the variant name: should be unique among variants.
// 	Name string

// 	// Path is the file path into which the generator will emit the code for this
// 	// variant.
// 	Path string

// 	// Package is the package this code will be emitted into.
// 	Package string

// 	// Imports is the imports needed for this package.
// 	Imports string

// 	// FuncSuffix is appended to all function names in this variant's code. All
// 	// suffixes should be unique within a package.
// 	FuncSuffix string

// 	// DataType is the type of the data parameter of functions in this variant's
// 	// code.
// 	DataType string

// 	// TypeParam is the optional type parameter for the function.
// 	TypeParam string

// 	// ExtraParam is an extra parameter to pass to the function. Should begin with
// 	// ", " to separate from other params.
// 	ExtraParam string

// 	// ExtraArg is an extra argument to pass to calls between functions; typically
// 	// it invokes ExtraParam. Should begin with ", " to separate from other args.
// 	ExtraArg string

// 	// Funcs is a map of functions used from within the template. The following
// 	// functions are expected to exist:
// 	//
// 	//    Less (name, i, j):
// 	//      emits a comparison expression that checks if the value `name` at
// 	//      index `i` is smaller than at index `j`.
// 	//
// 	//    Swap (name, i, j):
// 	//      emits a statement that performs a data swap between elements `i` and
// 	//      `j` of the value `name`.
// 	Funcs template.FuncMap
// }

// // generate generates the code for variant `v` into a file named by `v.Path`.
// func generate(v *Variant) {
// 	// Parse templateCode anew for each variant because Parse requires Funcs to be
// 	// registered, and it helps type-check the funcs.
// 	tmpl, err := template.New("gen").Funcs(v.Funcs).Parse(templateCode)
// 	if err != nil {
// 		log.Fatal("template Parse:", err)
// 	}

// 	var out bytes.Buffer
// 	err = tmpl.Execute(&out, v)
// 	if err != nil {
// 		log.Fatal("template Execute:", err)
// 	}

// 	formatted, err := format.Source(out.Bytes())
// 	if err != nil {
// 		log.Fatal("format:", err)
// 	}

// 	if err := os.WriteFile(v.Path, formatted, 0644); err != nil {
// 		log.Fatal("WriteFile:", err)
// 	}
// }
