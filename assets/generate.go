// generate.go does the preprocessing for build, includes:
//   1. Making copies of assets.
//   2. Generating consts for assets from the giving YAML file.
//
// Copies:
//   - ./img/ -> camera/assets/
//   - ./map/ -> base/assets/
//
// Generates: (uses the file passed by os.Args[1] as input)
//   - base/objtype_data.go
//   - camera/img_data.go
//   - base/flag_data.go

//go:generate rm -rf ../camera/assets
//go:generate cp -r img/ ../camera/assets/
//go:generate rm -rf ../base/assets
//go:generate cp -r map/ ../base/assets/

//go:generate go run generate.go data.yaml
//go:generate go fmt ../base/objtype_data.go

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

func templateData(d []*ObjectData, tmpl, dst string) error {
	tpl, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}

	_ = os.Remove(dst)
	f, err := os.Create(dst)
	if err != nil {
		return err
	}

	if err := tpl.Execute(f, d); err != nil {
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
	src := os.Args[1]

	objTmpl := "objtype_data.go.tmpl"
	objDst := "../base/objtype_data.go"

	imgTmpl := "img_data.go.tmpl"
	imgDst := "../camera/img_data.go"

	flagTmpl := "flag_data.go.tmpl"
	flagDst := "../base/flag_data.go"

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

	if err := templateData(d, objTmpl, objDst); err != nil {
		log.Fatal(err)
	}
	if err := templateData(d, imgTmpl, imgDst); err != nil {
		log.Fatal(err)
	}
	if err := templateData(d, flagTmpl, flagDst); err != nil {
		log.Fatal(err)
	}
}
