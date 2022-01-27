package setting

import (
	"encoding/json"
	"fmt"
	"os"
)

// setting ...
type setting struct {
	DGSExtractL1 string
	DGSExtractL2 string
	Bash         string
	Files        string
}

var Conf setting

func init() {

	file, e := os.Open("/home/a20272/Code/github.com/eugenefoxx/readDGSP1forKATE/setting/setting.cfg")
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось открыть файл конфигурации.")
	}

	defer file.Close()

	stat, e := file.Stat()
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось прочитать информацию о файле конфигурации.")
	}

	readByte := make([]byte, stat.Size())

	_, e = file.Read(readByte)
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось прочитать информацию о файле конфигурации.")
	}

	e = json.Unmarshal(readByte, &Conf)
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось прочитать информацию о файле конфигурации.")
	}
}
