package main

import (
	"fmt"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

func init() {
	multierror.ListFormatFunc = func(es []error) string {
		if len(es) == 1 {
			return fmt.Sprintf("出现1个错误:\n\t* %s\n\n", es[0])
		}

		points := make([]string, len(es))
		for i, err := range es {
			points[i] = fmt.Sprintf("* %s", err)
		}

		return fmt.Sprintf(
			"出现%d个错误:\n\t%s\n\n", len(es), strings.Join(points, "\n\t"))
	}
}

func A() error {
	return errors.WithMessage(B(), "A调用错误")
}
func B() error {
	return errors.WithMessage(C(), "B调用错误")
}
func C() error {
	return errors.New("C调用错误")
}
func D() (e error) {
	if err := B(); err != nil {
		e = multierror.Append(e, err)
	}
	if err := C(); err != nil {
		e = multierror.Append(e, err)
	}

	return
}
func E() (e error) {
	e = multierror.Append(e, nil)
	e = multierror.Append(e, B())
	e = multierror.Append(e, nil)
	e = multierror.Append(e, C())
	e = multierror.Append(e, nil)
	return
}

func main() {
	fmt.Printf("%v", A())
	println()
	fmt.Printf("%v", D())
	fmt.Printf("%v", E())

}
