package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// we do define the structs in accordance with the config file , but how do we serialise/deserialise it?
// could do it manually by reading the file
// or use a package "cleanenv" , we are using those `` (struct tags) because of these
 
type HTTPServer struct{
	Addr string `yaml:"address" env-required:"true" env-default:"6969"` // default not good for production
 }




	
type Config struct{
	Env string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`// to use in other packages
 StoragePath string `yaml:"storage_path" env-required:"true"`
 HTTPServer `yaml:"http_server" `

}

func MustLoad() *Config{ // name to speak for itself
	// for the functions that are must(in any context of getting the project up and running)
	//  we don't keep error in return type as if they fail , we directly throw an error
	var configPath string
	configPath = os.Getenv(configPath)

	if configPath==""{
		flags:= flag.String("config","","path to teh configuration file")
		flag.Parse()

		configPath= *flags// dereferencing

		
	}

	if configPath==""{
		log.Fatal("config not mentioned") // for config not provided in either of those
	}	else{

	if _, err:= os.Stat(configPath);os.IsNotExist(err){
			log.Fatal("configFile does not exist")
	}
}  // khachhar kaam for error handling if config path is wrong or the file is missing

var cfg 	Config

err:= cleanenv.ReadConfig(configPath, &cfg) 
if(err!=nil){
	log.Fatal("cannot read config file due to ", err.Error())// this error function is that errorf() from the error interface
}
 
return &cfg // cfg would be exported

}
