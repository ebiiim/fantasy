//go:generate rm -rf ../img/assets
//go:generate cp -r img ../img/assets
//go:generate go run generate.go

package main

import (
	"io"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

type ObjectData struct {
	Name    string   `yaml:"name"`
	Img     string   `yaml:"img"`
	Flags   []string `yaml:"flags"`
	FlagStr string   `yaml:"-"`
}

func templateData(d []*ObjectData, tmpl string, dst io.Writer) error {
	tpl, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}
	if err := tpl.Execute(dst, d); err != nil {
		return err
	}
	return nil
}

func parse(r io.Reader) ([]*ObjectData, error) {
	var d []*ObjectData
	if err := yaml.NewDecoder(r).Decode(&d); err != nil {
		return nil, err
	}
	return d, nil
}

func mergeFlags(d []*ObjectData) {
	defaultFlag := "None"
	for _, od := range d {
		flagStr := ""
		for idx, f := range od.Flags {
			flagStr += f
			if idx != len(od.Flags)-1 {
				flagStr += " | "
			}
		}
		if flagStr == "" {
			flagStr = defaultFlag
		}
		od.FlagStr = flagStr
	}
}

func main() {
	src := "data.yaml"
	dstFlag := "../flag/flag_data.go"
	dstImg := "../img/img_data.go"

	f, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	d, err := parse(f)
	if err != nil {
		log.Fatal(err)
	}
	mergeFlags(d)
	// for _, od := range d {
	// 	fmt.Printf("%+v\n", od)
	// }

	_ = os.Remove(dstImg)
	fImg, err := os.Create(dstImg)
	if err != nil {
		log.Fatal(err)
	}
	_ = os.Remove(dstFlag)
	fFlag, err := os.Create(dstFlag)
	if err != nil {
		log.Fatal(err)
	}
	if err := templateData(d, "flag_data.go.tmpl", fFlag); err != nil {
		log.Fatal(err)
	}
	if err := templateData(d, "img_data.go.tmpl", fImg); err != nil {
		log.Fatal(err)
	}
}
