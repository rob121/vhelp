package main

import(
	"github.com/rob121/vhelp"
	"log"
	"github.com/fsnotify/fsnotify"
)

func main(){


  vhelp.OnChange(func(e fsnotify.Event) {
	  log.Println("Config file changed:", e.Name)
  })
  //loads config.json
  vhelp.Load("config")
  //fetchs config instance
  conf,err := vhelp.Get("config")

  if(err!=nil){

  	log.Println(err)
  }

	log.Println(conf.ConfigFileUsed())

	log.Println(conf)
}