package main

import (
	"bufio"
	"os"
	"encoding/json"
	"fmt"
	"strings"
)

type GenDef struct {
	Gentype string
	Namedas string
	Utype1  string
	Utype2  string
}

type Config struct {
	Name   string
	Redefs []GenDef
}

func loadConfig(fn string) *Config {
	config := &Config{}

	//filename is the path to the json config file
	file, err := os.Open(fn)
	if err != nil {
		panic("Unable to open config")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		panic(fmt.Sprintf("Unable to parse config: %s", err))
	}
	//config.Password = os.Getenv("Db_Password")
	return config
}

func emit(gd GenDef) {
	inFN := fmt.Sprintf("templates/%s.go", gd.Gentype)
	infile, err := os.Open(inFN)
	if err != nil {
		panic(fmt.Sprintf("unable to open template: %s", gd.Gentype))
	}
	defer infile.Close()

	println("opened", inFN)
	reader := bufio.NewReader(infile)

	read := func() (string,interface{}) {
		line,err := reader.ReadString('\n')
	    return strings.TrimSuffix( line, "\n" ),err
    }

	for {
		line, err := read()
		if err != nil {
			panic(err)
		}
		if line == "//code starts" {
			break
		}
	}

	for {
		line, err := read()
		if err != nil {
			break
		}
		line = strings.Replace(line, "Type1", gd.Utype1, 100)
		if gd.Utype2 != "" {
			line = strings.Replace(line, "Type2", gd.Utype2, 100)
		}
		println(line)
	}

}

func main() {
	argsWithoutProg := os.Args[1:]
	fn := argsWithoutProg[0]
	cfg := loadConfig(fn)

	outFN := fmt.Sprintf("%s.go", cfg.Name)

	fh, err := os.Create(outFN)

	if err != nil {
		panic(fmt.Sprintf("Can't open output file %s", outFN))
	}

	defer fh.Close()

	for _, gd := range cfg.Redefs {
		emit(gd)
	}
}
