package factory

import (
	"fmt"
)

type Factory struct {
	JsonConfigPath string
}

func (self *Factory) Initialize() {
	fmt.Print("", "", self.JsonConfigPath+"config.json")

}
