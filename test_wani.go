package main


import (
	"github.com/codegangsta/martini"
    "io/ioutil"
)

func main() {

    server := martini.Classic()
    server.Get("/test/wanikani/:version/:type", func(args martini.Params) string {

    	version := args["version"]
    	typ := args["type"]

    	fileName := version + "_" + typ + ".json"

    	contents,_ := ioutil.ReadFile(fileName)
        return string(contents)

    })

    server.Run()

}
