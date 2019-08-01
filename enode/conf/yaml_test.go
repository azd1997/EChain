/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/19 21:16
* @Description: The file is for
***********************************************************************/

package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"testing"
)

func prepareDataForTest() string {
	var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

	return data
}

func TestYaml(t *testing.T) {

	data := prepareDataForTest()

	// Note: struct fields must be public in order for unmarshal to
	// correctly populate the data.
	type T struct {
		A string
		B struct {
			RenamedC int   `yaml:"c"`
			D        []int `yaml:",flow"`
		}
	}

	t1 := T{}

	err := yaml.Unmarshal([]byte(data), &t1)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t1)

	d, err := yaml.Marshal(&t1)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
