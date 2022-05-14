package main

import (
	"fmt"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	ruleLogger "github.com/hyperjumptech/grule-rule-engine/logger"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/sirupsen/logrus"
)

const (
	rule2 = `
rule AgeNameCheck "test" {
    when
			d.Age > 10 && d.CheckDuration()
    then
			d.Age = d.Age - 1;
			Forget("d.CheckDuration()");
}
`
)

type RuleDuration struct {
	Age int
}

func (r RuleDuration) CheckDuration() bool {
	fmt.Println("=============== 最后一个检查")
	fmt.Println(r.Age)
	return true
}

func main() {
	ruleLogger.Log.Logger.SetLevel(logrus.DebugLevel)
	dataContext := ast.NewDataContext()
	d := &RuleDuration{Age: 10}
	err := dataContext.Add("d", d)
	if err != nil {
		panic(err)
	}

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	if err = ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(rule2))); err != nil {
		panic(err)
	}
	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	eng1 := &engine.GruleEngine{MaxCycle: 1000}
	if err = eng1.Execute(dataContext, kb); err != nil {
		panic(err)
	}
}
