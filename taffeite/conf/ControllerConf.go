package conf

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"os"
)

type ControllerConf struct {
	HomePageUrl string
}

func NewControllerConf() *ControllerConf {
	data, err := os.ReadFile("resources/conf.yaml")
	if err != nil {
		panic(err)
	}

	var config map[string]interface{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	homePagePath, ok := config["home-page-path"].(string)
	if !ok {
		fmt.Println("Property not found")
	} else {
		fmt.Println("Value:", homePagePath)
	}

	return &ControllerConf{
		HomePageUrl: homePagePath,
	}
}
