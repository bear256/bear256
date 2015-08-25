package main

import (
	"fmt"
	"github.com/miguel-branco/goconfig"
	"log"
	"net/http"
)

func readCfg(cfg string) {
	c, err := goconfig.ReadConfigFile(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	//Site
	charset, _ := c.GetString("site", "charset")
	title, _ := c.GetString("site", "title")
	author, _ := c.GetString("site", "author")
	icon, _ := c.GetString("site", "icon")
	column, _ := c.GetString("site", "column")
	root, _ := c.GetString("site", "root")
	fmt.Println(charset, title, author, icon, column, root)
}

func main() {
	fmt.Println("TaiChi Panda SiteGenerater")
	readCfg("../../template/site.cfg")
	http.ListenAndServe(":8000", http.FileServer(http.Dir("../../")))
}
