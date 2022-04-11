package main

import (
	"flag"
	"fmt"
	gen2 "github.com/huaixiaohai/tools/cmd/protoc-gen-restful/gen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

//import "google.golang.org/protobuf"

const (
	version = "1.2.0"

	//contextPackage = protogen.GoImportPath("context")
)
var requireUnimplemented *bool

// usage: protoc --restful_out=. ./msg/proto/*.proto
func main() {
	//////buf := make([]byte, 0)
	//////buf, _ = ioutil.ReadAll(os.Stdin)
	////////println(err.Error())
	//////os.WriteFile("a.txt",buf,0666)
	//buf, err := os.ReadFile("a.txt")
	//var err error
	//os.Stdin, err = os.Open("a.txt")
	//println(err)

	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-grpc %v\n", version)
		return
	}

	var flags flag.FlagSet
	requireUnimplemented = flags.Bool("require_unimplemented_servers", true, "set to false to match legacy behavior")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			gen2.GenerateFile(gen, f)
		}
		return nil
	})
}

