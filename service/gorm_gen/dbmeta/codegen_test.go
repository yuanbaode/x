package dbmeta

import (
	"fmt"
	"testing"
)

func TestConfig_CreateContextForTableFile(t *testing.T) {

	c := &Config{
		SQLType:               "",
		SQLConnStr:            "",
		SQLDatabase:           "",
		Module:                "",
		ModelPackageName:      "",
		ModelFQPN:             "",
		AddJSONAnnotation:     false,
		AddGormAnnotation:     false,
		AddProtobufAnnotation: false,
		AddXMLAnnotation:      false,
		AddDBAnnotation:       false,
		UseGureguTypes:        false,
		JSONNameFormat:        "",
		XMLNameFormat:         "",
		ProtobufNameFormat:    "",
		DaoPackageName:        "",
		DaoFQPN:               "",
		APIPackageName:        "",
		APIFQPN:               "",
		GrpcPackageName:       "",
		GrpcFQPN:              "",
		Swagger:               nil,
		ServerPort:            0,
		ServerHost:            "",
		ServerScheme:          "",
		ServerListen:          "",
		Verbose:               false,
		OutDir:                "",
		Overwrite:             false,
		LineEndingCRLF:        false,
		CmdLine:               "",
		CmdLineWrapped:        "",
		CmdLineArgs:           nil,
		FileNamingTemplate:    "",
		ModelNamingTemplate:   "",
		FieldNamingTemplate:   "",
		ContextMap:            nil,
		TemplateLoader:        nil,
		TableInfos:            nil,
		FragmentsDir:          "",
		fragments:             nil,
	}
	m := ModelInfo{
		Index:           0,
		IndexPlus1:      0,
		PackageName:     "",
		StructName:      "",
		ShortStructName: "",
		TableName:       "",
		Fields:          nil,
		DBMeta:          nil,
		Instance:        nil,
		CodeFields:      nil,
	}
	r:=c.CreateContextForTableFile(&m)
	fmt.Println(r)

}

func Test_Replace(t *testing.T)  {
	a:=Replace("{{FmtFieldName (stringifyFirstChar .) }}","id")
	fmt.Println(a)
}