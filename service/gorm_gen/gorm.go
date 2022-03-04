package gorm_gen

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/packr/v2"
	"github.com/jimsmart/schema"
	"github.com/yuanbaode/x/service/gorm_gen/dbmeta"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var baseTemplates *packr.Box

func Gen(sqlConnStr string, sqlTable, database, model string) {
	if sqlTable == "" && database == "" {
		log.Fatal("database is empty")
	}

	baseTemplates = packr.New("gen", "./template")

	conf := dbmeta.NewConfig(LoadTemplate, model)
	err := loadDefaultDBMappings(conf)
	if err != nil {
		log.Fatal(err)
	}
	db, err := initializeDB(sqlConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var dbTables []string
	// parse or read tables
	if sqlTable != "" {
		dbTables = strings.Split(sqlTable, ",")
	} else {
		schemaTables, err := schema.TableNames(db)
		if err != nil {
			log.Fatal(fmt.Sprintf("Error in fetching tables information from %s information schema from %s\n", "mysql", sqlConnStr))
		}
		for _, st := range schemaTables {
			if st[0] != database {
				continue
			}
			dbTables = append(dbTables, st[1]) // s[0] == sqlDatabase
		}
	}

	tableInfos := dbmeta.LoadTableInfo(db, dbTables, nil, conf)
	i := 1
	for tableName := range tableInfos {
		fmt.Printf("[%d] %s\n", i, tableName)
		i++
	}
	generate(conf, tableInfos)

	return
}

func initializeDB(sqlConnStr string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", sqlConnStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func LoadTemplate(filename string) (tpl *dbmeta.GenTemplate, err error) {
	baseName := filepath.Base(filename)
	// fmt.Printf("LoadTemplate: %s / %s\n", filename, baseName)
	content, err := baseTemplates.FindString(baseName)
	if err != nil {
		return nil, fmt.Errorf("%s not found internally", baseName)
	}
	tpl = &dbmeta.GenTemplate{Name: "internal://" + filename, Content: content}
	return tpl, nil
}

func loadDefaultDBMappings(conf *dbmeta.Config) error {
	var err error
	var content []byte
	content, err = baseTemplates.Find("mapping.json")
	if err != nil {
		return err
	}

	err = dbmeta.ProcessMappings("internal", content, conf.Verbose)
	if err != nil {
		return err
	}
	return nil
}

func generate(conf *dbmeta.Config, tableInfos map[string]*dbmeta.ModelInfo) error {
	var err error
	modelDir := filepath.Join(conf.OutDir, conf.ModelDir)
	overwrite := &conf.Overwrite
	err = os.MkdirAll(modelDir, 0777)
	if err != nil && !*overwrite {
		log.Fatal(fmt.Sprintf("unable to create modelDir: %s error: %v\n", modelDir, err))
	}

	var ModelTmpl *dbmeta.GenTemplate
	//var ModelBaseTmpl *dbmeta.GenTemplate
	if ModelTmpl, err = LoadTemplate("model.go.tmpl"); err != nil {
		log.Fatal(fmt.Sprintf("Error loading template %v\n", err))
		return err
	}
	//if ModelBaseTmpl, err = LoadTemplate("model_base.go.tmpl"); err != nil {
	//	log.Fatal(fmt.Sprintf("Error loading template %v\n", err))
	//	return err
	//}

	// generate go files for each table
	for tableName, tableInfo := range tableInfos {

		if len(tableInfo.Fields) == 0 {
			continue
		}

		modelInfo := conf.CreateContextForTableFile(tableInfo)

		modelFile := filepath.Join(modelDir, CreateGoSrcFileName(tableName))
		err = conf.WriteTemplate(ModelTmpl, modelInfo, modelFile)
		if err != nil {
			log.Fatal(fmt.Sprintf("Error writing file: %v\n", err))
		}
	}

	//data := map[string]interface{}{}

	//err = conf.WriteTemplate(ModelBaseTmpl, data, filepath.Join(modelDir, "model_base.go"))
	//if err != nil {
	//	log.Fatal(fmt.Sprintf("Error writing file: %v\n", err))
	//}

	//data = map[string]interface{}{
	//	"deps":        "go list -f '{{ join .Deps  \"\\n\"}}' .",
	//	"CommandLine": conf.CmdLine,
	//	"Config":      conf,
	//}

	GoFmt(conf.OutDir)
	return nil
}

// GoFmt exec gofmt for a code dir
func GoFmt(codeDir string) (string, error) {
	args := []string{"-s", "-d", "-w", "-l", codeDir}
	cmd := exec.Command("gofmt", args...)

	cmdLineArgs := strings.Join(args, " ")
	fmt.Printf("gofmt %s\n", cmdLineArgs)

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(fmt.Sprintf("error calling protoc: %T %v\n", err, err))
	}
	return string(stdoutStderr), nil
}

// CreateGoSrcFileName ensures name doesnt clash with go naming conventions like _test.go
func CreateGoSrcFileName(tableName string) string {
	name := tableName
	// name := inflection.Singular(tableName)

	if strings.HasSuffix(name, "_test") {
		name = name[0 : len(name)-5]
		name = name + "_tst"
	}
	return name + ".go"
}
