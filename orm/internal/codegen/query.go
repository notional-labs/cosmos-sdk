package codegen

import (
	"bytes"
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	ormv1 "cosmossdk.io/api/cosmos/orm/v1"

	"cosmossdk.io/orm/internal/fieldnames"
)

type queryProtoGen struct {
	*protogen.File
	imports map[string]bool
	svc     *writer
	msgs    *writer
	outFile *os.File
}

func (g queryProtoGen) gen() error {
	g.imports[g.Desc.Path()] = true

	g.svc.F("// %sService queries the state of the tables specified by %s.", g.queryServiceName(), g.Desc.Path())
	g.svc.F("service %sService {", g.queryServiceName())
	g.svc.Indent()
	for _, msg := range g.Messages {
		tableDesc := proto.GetExtension(msg.Desc.Options(), ormv1.E_Table).(*ormv1.TableDescriptor)
		if tableDesc != nil {
			err := g.genTableRPCMethods(msg, tableDesc)
			if err != nil {
				return err
			}
		}
		singletonDesc := proto.GetExtension(msg.Desc.Options(), ormv1.E_Singleton).(*ormv1.SingletonDescriptor)
		if singletonDesc != nil {
			err := g.genSingletonRPCMethods(msg)
			if err != nil {
				return err
			}
		}
	}
	g.svc.Dedent()
	g.svc.F("}")
	g.svc.F("")

	outBuf := newWriter()
	outBuf.F("// Code generated by protoc-gen-go-cosmos-orm-proto. DO NOT EDIT.")
	outBuf.F(`syntax = "proto3";`)
	outBuf.F("package %s;", g.Desc.Package())
	outBuf.F("")

	imports := maps.Keys(g.imports)
	slices.Sort(imports)
	for _, i := range imports {
		outBuf.F(`import "%s";`, i)
	}
	outBuf.F("")

	_, err := outBuf.Write(g.svc.Bytes())
	if err != nil {
		return err
	}

	_, err = outBuf.Write(g.msgs.Bytes())
	if err != nil {
		return err
	}

	_, err = g.outFile.Write(outBuf.Bytes())
	if err != nil {
		return err
	}

	return g.outFile.Close()
}

func (g queryProtoGen) genTableRPCMethods(msg *protogen.Message, desc *ormv1.TableDescriptor) error {
	name := msg.Desc.Name()
	g.svc.F("// Get queries the %s table by its primary key.", name)
	g.svc.F("rpc Get%s(Get%sRequest) returns (Get%sResponse) {}", name, name, name) // TODO grpc gateway

	g.startRequestType("Get%sRequest", name)
	g.msgs.Indent()
	primaryKeyFields := fieldnames.CommaSeparatedFieldNames(desc.PrimaryKey.Fields)
	fields := msg.Desc.Fields()
	for i, fieldName := range primaryKeyFields.Names() {
		field := fields.ByName(fieldName)
		if field == nil {
			return fmt.Errorf("can't find primary key field %s", fieldName)
		}
		g.msgs.F("// %s specifies the value of the %s field in the primary key.", fieldName, fieldName)
		g.msgs.F("%s %s = %d;", g.fieldType(field), fieldName, i+1)
	}
	g.msgs.Dedent()

	g.msgs.F("}")
	g.msgs.F("")
	g.startResponseType("Get%sResponse", name)
	g.msgs.Indent()
	g.msgs.F("// value is the response value.")
	g.msgs.F("%s value = 1;", name)
	g.msgs.Dedent()
	g.msgs.F("}")
	g.msgs.F("")

	for _, idx := range desc.Index {
		if !idx.Unique {
			continue
		}

		fieldsCamel := fieldsToCamelCase(idx.Fields)
		methodName := fmt.Sprintf("Get%sBy%s", name, fieldsCamel)
		g.svc.F("// %s queries the %s table by its %s index", methodName, name, fieldsCamel)
		g.svc.F("rpc %s(%sRequest) returns (%sResponse) {}", methodName, methodName, methodName) // TODO grpc gateway

		g.startRequestType("%sRequest", methodName)
		g.msgs.Indent()
		fieldNames := fieldnames.CommaSeparatedFieldNames(idx.Fields)
		for i, fieldName := range fieldNames.Names() {
			field := fields.ByName(fieldName)
			if field == nil {
				return fmt.Errorf("can't find unique index field %s", fieldName)
			}
			g.msgs.F("%s %s = %d;", g.fieldType(field), fieldName, i+1)
		}
		g.msgs.Dedent()

		g.msgs.F("}")
		g.msgs.F("")
		g.startResponseType("%sResponse", methodName)
		g.msgs.Indent()
		g.msgs.F("%s value = 1;", name)
		g.msgs.Dedent()
		g.msgs.F("}")
		g.msgs.F("")
	}

	g.imports["cosmos/base/query/v1beta1/pagination.proto"] = true
	g.svc.F("// List%s queries the %s table using prefix and range queries against defined indexes.", name, name)
	g.svc.F("rpc List%s(List%sRequest) returns (List%sResponse) {}", name, name, name) // TODO grpc gateway
	g.startRequestType("List%sRequest", name)
	g.msgs.Indent()
	g.msgs.F("// IndexKey specifies the value of an index key to use in prefix and range queries.")
	g.msgs.F("message IndexKey {")
	g.msgs.Indent()

	indexFields := []string{desc.PrimaryKey.Fields}
	// the primary key has field number 1
	fieldNums := []uint32{1}
	for _, index := range desc.Index {
		indexFields = append(indexFields, index.Fields)
		// index field numbers are their id + 1
		fieldNums = append(fieldNums, index.Id+1)
	}

	g.msgs.F("// key specifies the index key value.")
	g.msgs.F("oneof key {")
	g.msgs.Indent()
	for i, fields := range indexFields {
		fieldName := fieldsToSnakeCase(fields)
		typeName := fieldsToCamelCase(fields)
		g.msgs.F("// %s specifies the value of the %s index key to use in the query.", fieldName, typeName)
		g.msgs.F("%s %s = %d;", typeName, fieldName, fieldNums[i])
	}
	g.msgs.Dedent()
	g.msgs.F("}")

	for _, fieldNames := range indexFields {
		g.msgs.F("")
		g.msgs.F("message %s {", fieldsToCamelCase(fieldNames))
		g.msgs.Indent()
		for i, fieldName := range fieldnames.CommaSeparatedFieldNames(fieldNames).Names() {
			g.msgs.F("// %s is the value of the %s field in the index.", fieldName, fieldName)
			g.msgs.F("// It can be omitted to query for all valid values of that field in this segment of the index.")
			g.msgs.F("optional %s %s = %d;", g.fieldType(fields.ByName(fieldName)), fieldName, i+1)
		}
		g.msgs.Dedent()
		g.msgs.F("}")
	}

	g.msgs.Dedent()
	g.msgs.F("}")
	g.msgs.F("")
	g.msgs.F("// query specifies the type of query - either a prefix or range query.")
	g.msgs.F("oneof query {")
	g.msgs.Indent()
	g.msgs.F("// prefix_query specifies the index key value to use for the prefix query.")
	g.msgs.F("IndexKey prefix_query = 1;")
	g.msgs.F("// range_query specifies the index key from/to values to use for the range query.")
	g.msgs.F("RangeQuery range_query = 2;")
	g.msgs.Dedent()
	g.msgs.F("}")

	g.msgs.F("// pagination specifies optional pagination parameters.")
	g.msgs.F("cosmos.base.query.v1beta1.PageRequest pagination = 3;")
	g.msgs.F("")
	g.msgs.F("// RangeQuery specifies the from/to index keys for a range query.")
	g.msgs.F("message RangeQuery {")
	g.msgs.Indent()
	g.msgs.F("// from is the index key to use for the start of the range query.")
	g.msgs.F("// To query from the start of an index, specify an index key for that index with empty values.")
	g.msgs.F("IndexKey from = 1;")
	g.msgs.F("// to is the index key to use for the end of the range query.")
	g.msgs.F("// The index key type MUST be the same as the index key type used for from.")
	g.msgs.F("// To query from to the end of an index it can be omitted.")
	g.msgs.F("IndexKey to = 2;")
	g.msgs.Dedent()
	g.msgs.F("}")
	g.msgs.Dedent()
	g.msgs.F("}")
	g.msgs.F("")
	g.startResponseType("List%sResponse", name)
	g.msgs.Indent()
	g.msgs.F("// values are the results of the query.")
	g.msgs.F("repeated %s values = 1;", name)
	g.msgs.F("// pagination is the pagination response.")
	g.msgs.F("cosmos.base.query.v1beta1.PageResponse pagination = 2;")
	g.msgs.Dedent()
	g.msgs.F("}")
	g.msgs.F("")
	return nil
}

func (g queryProtoGen) genSingletonRPCMethods(msg *protogen.Message) error {
	name := msg.Desc.Name()
	g.svc.F("// Get%s queries the %s singleton.", name, name)
	g.svc.F("rpc Get%s(Get%sRequest) returns (Get%sResponse) {}", name, name, name) // TODO grpc gateway
	g.startRequestType("Get%sRequest", name)
	g.msgs.F("}")
	g.msgs.F("")
	g.startRequestType("Get%sResponse", name)
	g.msgs.Indent()
	g.msgs.F("%s value = 1;", name)
	g.msgs.Dedent()
	g.msgs.F("}")
	g.msgs.F("")
	return nil
}

func (g queryProtoGen) startRequestType(format string, args ...any) {
	g.startRequestResponseType("request", format, args...)
}

func (g queryProtoGen) startResponseType(format string, args ...any) {
	g.startRequestResponseType("response", format, args...)
}

func (g queryProtoGen) startRequestResponseType(typ, format string, args ...any) {
	msgTypeName := fmt.Sprintf(format, args...)
	g.msgs.F("// %s is the %s/%s %s type.", msgTypeName, g.queryServiceName(), msgTypeName, typ)
	g.msgs.F("message %s {", msgTypeName)
}

func (g queryProtoGen) queryServiceName() string {
	return fmt.Sprintf("%sQuery", strcase.ToCamel(fileShortName(g.File)))
}

func (g queryProtoGen) fieldType(descriptor protoreflect.FieldDescriptor) string {
	if descriptor.Kind() == protoreflect.MessageKind {
		message := descriptor.Message()
		g.imports[message.ParentFile().Path()] = true
		return string(message.FullName())
	}

	return descriptor.Kind().String()
}

type writer struct {
	*bytes.Buffer
	indent    int
	indentStr string
}

func newWriter() *writer {
	return &writer{
		Buffer: &bytes.Buffer{},
	}
}

func (w *writer) F(format string, args ...interface{}) {
	_, err := w.Write([]byte(w.indentStr))
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintf(w, format, args...)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintln(w)
	if err != nil {
		panic(err)
	}
}

func (w *writer) Indent() {
	w.indent++
	w.updateIndent()
}

func (w *writer) updateIndent() {
	w.indentStr = ""
	for i := 0; i < w.indent; i++ {
		w.indentStr += "  "
	}
}

func (w *writer) Dedent() {
	w.indent--
	w.updateIndent()
}
