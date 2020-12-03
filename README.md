# Vhelp

Viper Config Helper, boot strapping of viper config and handles boiler plate

## Features

* Auto loads config in user home, current directory and /etc
* If configured, binds a watcher to handle config changes
* Auto calculates the binary name for the config dir ie if the binary is "example" it will look in:
  - /etc/example
  - $HOME/.example
  

Example usage

```
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
  //fetchs config instance which is a *viper.Viper
  conf,err := vhelp.Get("config")

  if(err!=nil){

  	log.Println(err)
  }

	log.Println(conf.ConfigFileUsed())

	log.Println(conf)
}
```
