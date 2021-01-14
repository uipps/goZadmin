package main

import (
	"github.com/zserge/lorca"
	"log"
)

func main() {
	ui, err := lorca.New("https://www.baidu.com", "", 1024, 768)
	if err != nil {
		log.Fatal(err)
	}
	ui.Eval(`
        //alert('1');
	`)
	defer ui.Close()
	<-ui.Done()
}
