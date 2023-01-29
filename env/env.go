package env

import "os"

func Loadenv(){
  os.Setenv("FOLDER", "./images/")
  os.Setenv("apikey", "?")
}
