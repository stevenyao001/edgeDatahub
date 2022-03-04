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
	MegnetStatus bool `json:"megnet_status"`
	Num          int  `json:"num"`
	Status       int  `json:"status"`
	Ia           int  `json:"ia"`
	Ep           int  `json:"ep"`
	InstantEp    int  `json:"instant_ep"`
}

type MiddleData struct {
	Ts         int64             `json:"ts"`
	Properties *MiddleProperties `json:"properties"`
}

type MiddleProperties struct {
	MegnetStatus bool `json:"megnet_status"`
	Num          int  `json:"num"`
	Status       int  `json:"status"`
	Ia           int  `json:"ia"`
	Ep           int  `json:"ep"`
	InstantEp    int  `json:"instant_ep"`
}

var middle *MiddleData

func InitMiddle() {
	middle = &MiddleData{
		Properties: &MiddleProperties{
			MegnetStatus: false,
			Num:          0,
			Status:       0,
			Ia:           0,
			Ep:           0,
			InstantEp:    0,
		},
	}
}

func Computing(buf []byte, path string) ([]byte, error) {
	input := GetData(buf)
	rule, err := GetRule(path)
	//fmt.Println(rule, "------", err)
	if err != nil {
		return nil, err
	}
	dataContext := context.NewDataContext()
	dataContext.Add("input", input)
	dataContext.Add("middle", middle)

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
	pushMiddleData(input)

	return json.Marshal(input)
}

func pushMiddleData(input *Input) {
	if middle.Properties.Num != input.Properties.Num {
		middle = &MiddleData{
			Ts: input.Ts,
			Properties: &MiddleProperties{
				MegnetStatus: input.Properties.MegnetStatus,
				Num:          middle.Properties.Num + 1,
				Status:       input.Properties.Status,
				Ia:           input.Properties.Ia,
				Ep:           input.Properties.Ep,
				InstantEp:    input.Properties.InstantEp,
			},
		}
	} else {
		middle = &MiddleData{
			Ts: input.Ts,
			Properties: &MiddleProperties{
				MegnetStatus: input.Properties.MegnetStatus,
				Num:          middle.Properties.Num,
				Status:       input.Properties.Status,
				Ia:           input.Properties.Ia,
				Ep:           input.Properties.Ep,
				InstantEp:    input.Properties.InstantEp,
			},
		}
	}
}

func GetRule(path string) (string, error) {
	rule := `
rule "test" "test"
begin
`
	ruleD := make([]ruleData, 5)
	//f, err := ioutil.ReadFile(path)
	//if err != nil {
	//	return "", err
	//}
	f := []byte(`[{"name":"MegnetStatus","rule":""},{"name":"Num","rule":"if MegnetStatus_last == false \u0026\u0026 MegnetStatus == true {\n\t Num = Num_last + 1\n} else {\n\t Num = Num_last\n}"},{"name":"Status","rule":"if Ia == 0 {\n\t Status = 0\n} else if Ia \u003e 40 {\n\t Status = 2\n} else {\n\t Status = 1\n}\n"},{"name":"Ia","rule":""},{"name":"Ep","rule":""},{"name":"InstantEp","rule":" InstantEp = Ep - Ep_last"}]`)
	err := json.Unmarshal(f, &ruleD)
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
		nameTmp2 := "[" + name + "_last]" + "{" + strconv.Itoa(len(name)+5) + ",}"
		replace2 := " middle.Properties." + name
		regIf2 := regexp.MustCompile(nameTmp2)
		if regIf2 != nil {
			rule = regIf2.ReplaceAllString(rule, replace2)
		}

		nameTmp := "[ ]+[" + name + "]" + "{" + strconv.Itoa(len(name)) + ",}"
		replace := " input.Properties." + name
		regIf := regexp.MustCompile(nameTmp)
		if regIf != nil {
			rule = regIf.ReplaceAllString(rule, replace)
		}
	}

	return rule
}
