package main

import (
	"github.com/134130/protc-gen-grpc-rest/internal/genrestgrpc"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			_ = genrestgrpc.GenerateFile(gen, f)
		}
		return nil
	})
}
