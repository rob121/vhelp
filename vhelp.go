package vhelp

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var cb func(e fsnotify.Event)
var watch bool
var configs map[string]*viper.Viper

func init (){

	watch = false
	configs  = make(map[string]*viper.Viper)
}

func OnChange(fn func(e fsnotify.Event)){
    watch = true
	cb = fn
}

func Load(file string){


	if(len(file)==0){
		file = "config"
	}

	runtime_viper := viper.New()

	exefull, _ := os.Executable()
	exe := filepath.Base(exefull)

	runtime_viper.SetConfigName(file)        // name of config file (without extension)
	runtime_viper.SetConfigType("json")          // REQUIRED if the config file does not have the extension in the name
	runtime_viper.AddConfigPath("/etc/" + exe)   // path to look for the config file in
	runtime_viper.AddConfigPath("$HOME/." + exe) // call multiple times to add many search paths
	runtime_viper.AddConfigPath(".")             // optionally look for config in the working directory
	err := runtime_viper.ReadInConfig()          // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if(watch==true) {
		runtime_viper.WatchConfig()
		runtime_viper.OnConfigChange(cb)
	}

	configs[file] = runtime_viper


}

func Get(name string) (*viper.Viper,error){

	var ok bool
	var runtime_viper *viper.Viper

	if runtime_viper, ok = configs[name]; !ok {


		 return nil,errors.New("Not Found")

	}

	return runtime_viper,nil


}