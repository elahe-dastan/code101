package main

// Want to learn to use koanf because my gorgeous boyfriend is a contributor

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"log"
)

// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
var k = koanf.New(".")

type Elahe struct {
	Name      string `koanf:"name"`
	Age       string `koanf:"age"`
	Boyfriend struct {
		Kind string `koanf:"kind"`
	} `koanf:"boyfriend"`
}

func main() {
	// Load YAML config.
	// koanf.Provider is a generic interface that provides configuration , for example, from files, environment variables,...
	// koanf.Parser is a generic interface that takes raw bytes, parses, and returns a nested map[string]interface{} representation.
	f := file.Provider("Go/koanf/config.yml")

	if err := k.Load(f, yaml.Parser()); err != nil {
		log.Fatal(err)
	}

	var out Elahe
	if err := k.Unmarshal("elahe", &out); err != nil{
		log.Fatal(err)
	}

	fmt.Println(out)


	// Watching files for changes. The most important advantage of koanf over viper
	// Watch the file and get a callback on change. The callback can do whatever,
	// like re-load the configuration.
	//fuck(f, out)
	f.Watch(func(event interface{}, err error) {
		if err != nil {
			log.Fatal(err)
		}

		if err := k.Load(f, yaml.Parser()); err != nil {
			log.Fatal(err)
		}

		if err := k.Unmarshal("elahe", &out); err != nil{
			log.Fatal(err)
		}

		fmt.Println(out)

	})

	//block forever
	check()
}

func check() {
	//block forever
	<-make(chan bool)
}

// points
// relative paths are based on where your go.mod is
// the delimiter is used for addressing for example elahe.name="raha"
// if you want to unmarshal into a struct, its fields should be exported
