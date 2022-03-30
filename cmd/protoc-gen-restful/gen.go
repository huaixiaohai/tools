package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
)

const (
	jsonTag = "json"
	formTag = "form"
)

func generateFile(plugin *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	//if len(file.Services) == 0 {
	//	return nil
	//}
	fileName := file.GeneratedFilenamePrefix + ".pb.go"
	gf:=plugin.NewGeneratedFile(fileName, file.GoImportPath)
	gf.P("// Code generated by protoc-gen-go-grpc. DO NOT EDIT.")
	gf.P()
	gf.P("package ", file.GoPackageName)
	gf.P()

	generateFileContent(plugin, file, gf)

	return gf
}


// generateFileContent generates the gRPC service definitions, excluding the package statement.
func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}

	//g.P("const _ = ", ginPackage.Ident("SupportPackageIsVersion7")) // When changing, update version number above.
	g.P()

	for _, service := range file.Services {
		genInterface(g, service)
	}

	g.P()
	g.P()
	g.P()

	for _, message := range file.Messages {
		println(message.Desc.Name())
		genMessage(g, message)
	}
}

func genInterface(g *protogen.GeneratedFile, service *protogen.Service) {
	// Api interface.
	interfaceName := "I"+service.GoName
	g.P("// ", interfaceName, " 是 ", service.GoName, " 的interface")
	g.P("//")

	g.Annotate(interfaceName, service.Location)
	g.P("type ", interfaceName, " interface {")
	for _, method := range service.Methods {
		g.Annotate(interfaceName+"."+method.GoName, method.Location)
		g.P(method.Comments.Leading,
			genInterfaceFunc(g, method))
	}
	g.P("}")
	g.P()
}

func genInterfaceFunc(g *protogen.GeneratedFile, method *protogen.Method) string {
	reqArgs := []string{
		"ctx *"+g.QualifiedGoIdent(ginPackage.Ident("Context")),
		"req *"+g.QualifiedGoIdent(method.Input.GoIdent),
	}

	retArgs := []string{
		"*" + g.QualifiedGoIdent(method.Output.GoIdent),
		"error",
	}
	return fmt.Sprintf("%s(%s)(%s)", method.GoName, strings.Join(reqArgs, ", "), strings.Join(retArgs, ","))
}

func genMessage(g *protogen.GeneratedFile, message *protogen.Message) {
	g.P(message.Comments.Leading,"type ", message.GoIdent, " struct {")
	for _, field := range message.Fields {
		genMessageField(g, field)
	}
	g.P("}")
	g.P()

	// 处理message里面的messages
	for _, message = range message.Messages {
		genMessage(g, message)
	}
}


func genMessageField(g *protogen.GeneratedFile, field *protogen.Field) {
	goType, pointer := fieldGoType(g,field)
	if pointer {
		goType = "*"+goType
	}
	tags := structTags{
		{jsonTag,string(field.Desc.Name())},
		{formTag,string(field.Desc.Name())},
	}
	//tag := fmt.Sprintf("`json:\"%s\" form:\"%s\"`", field.Desc.Name(), field.Desc.Name())
	g.P(field.GoName, " ", goType, fieldTags(tags), field.Comments.Leading)
}

func fieldGoType(g *protogen.GeneratedFile, field *protogen.Field) (goType string, pointer bool) {
	if field.Desc.IsWeak() {
		return "struct{}", false
	}

	pointer = field.Desc.HasPresence()
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = g.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
		pointer = false // rely on nullability of slices for presence
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = "*" + g.QualifiedGoIdent(field.Message.GoIdent)
		pointer = false // pointer captured as part of the type
	}
	switch {
	case field.Desc.IsList():
		return "[]" + goType, false
	case field.Desc.IsMap():
		panic("message里面不允许map类型   "+ field.GoName)
	}
	return goType, pointer
}

type structTags [][2]string
func fieldTags(tags structTags) string {
	if len(tags) <= 0 {
		return ""
	}
	var str string
	for _, tag := range tags {
		str += fmt.Sprintf("%s:\"%s\" ", tag[0],tag[1])
	}
	return fmt.Sprintf("`%s`", str[0:len(str)-1])
}