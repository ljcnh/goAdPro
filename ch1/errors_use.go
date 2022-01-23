package main

import (
	"fmt"
	"github.com/ljcnh/goAdPro/ch1/errors"
	"io/ioutil"
)

func loadConfig() error {
	_, err := ioutil.ReadFile("/path/to/file")
	if err != nil {
		return errors.Wrap(err, "read failed")
	}
	return nil
	// ...
}

func setup() error {
	err := loadConfig()
	if err != nil {
		return errors.Wrap(err, "invalid config")
	}
	return nil
	// ...
}

func main() {
	err := setup()
	//if  err != nil {
	//	log.Fatal(err)
	//}
	for i, e := range err.(errors.Error).Wraped() {
		fmt.Printf("wraped(%d): %v\n", i, e)
	}
	// ...
}
