package rule

import (
	"encoding/json"
	"fmt"
	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
	"regexp"
	"strconv"
)

type ruleData struct {
	Name string `json:"name"`
	Rule string `json:"rule"`
}

type Input struct {
	Ts         int64       `json:"ts"`
	Properties *Properties `json:"properties"`
}

type Properties struct {
	AddressNames1 float64 `json:"addressNames1"`
	AddressNames2 float64 `json:"addressNames2"`
	AddressNames3 float64 `json:"addressNames3"`
}

func Computing(buf []byte, path string) ([]byte, error) {
	input := GetData(buf)
	rule, err := GetRule(path)
	if err != nil {
		return nil, err
	}
	dataContext := context.NewDataContext()
	dataContext.Add("input", input)

	ruleBuilder := builder.NewRuleBuilder(dataContext)
	e1 := ruleBuilder.BuildRuleFromString(rule)
	if e1 != nil {
		panic(e1)
	}
	gengine := engine.NewGengine()
	e := gengine.Execute(ruleBuilder, true)
	if e != nil {
		panic(e)
	}
	return json.Marshal(input)
}

func GetRule(path string) (string, error) {
	rule := `
rule "test" "test"
begin
`
	ruleD := make([]ruleData, 3)
	//f, err := ioutil.ReadFile(path)
	//fmt.Println(string(f))
	//if err != nil {
	//	fmt.Println("err",err)
	//	return "", err
	//}
	s := `[{"name":"AddressNames1","rule":"AddressNames1 = AddressNames1 + 1"},{"name":"AddressNames2","rule":"AddressNames2 = AddressNames2 + 1"},{"name":"AddressNames3","rule":"AddressNames3 = AddressNames1 + AddressNames2"}]`
	err := json.Unmarshal([]byte(s), &ruleD)
	if err != nil {
		return "", err
	}
	names := make([]string, 0)
	for _, value := range ruleD {
		names = append(names, value.Name)
	}

	for _, value := range ruleD {
		if value.Rule != "" {
			rule = rule + ruleTmpComputer(value.Rule, names)
			rule = rule + "\n"
		}
	}
	return rule + "end", nil
}

func GetData(buf []byte) *Input {
	input := new(Input)
	err := json.Unmarshal(buf, input)
	if err != nil {
		fmt.Println(err)
	}
	return input
}

func ruleTmpComputer(rule string, names []string) string {
	for _, name := range names {
		nameTmp := "[" + name + "]" + "{" + strconv.Itoa(len(name)) + ",}"
		replace := "input.Properties." + name
		regIf := regexp.MustCompile(nameTmp)
		if regIf != nil {
			rule = regIf.ReplaceAllString(rule, replace)
		}
	}

	return rule
}
