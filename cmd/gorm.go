package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yuanbaode/x/service/gorm_gen"
)

var gormCmd = &cobra.Command{
	Use:     "gorm",
	Short:   "gen gorm struct",
	Example: "",

	Run: gormC,
}

var (
	connstr  string
	database string
	model string
	sqlTable string
)

func init() {
	gormCmd.Flags().StringVarP(&connstr, "connstr", "c", "127.0.0.1:7777", "database connection string")
	gormCmd.Flags().StringVarP(&sqlTable, "table", "t", "", "table")
	gormCmd.Flags().StringVarP(&database, "database", "d", ".goc.kvstore", "Database to for connection")
	gormCmd.Flags().StringVarP(&model, "model", "m", "model", "package to model")

	rootCmd.AddCommand(gormCmd)
}

func gormC(cmd *cobra.Command, args []string) {
	gorm_gen.Gen(connstr, sqlTable,database,model )
}
