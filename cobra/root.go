package cobra

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/b87/db-kit/database"
)

var (
	host       *string
	port       *int
	user       *string
	password   *string
	db         *string
	migrations *string
	backups    *string
)

func newDB() (*database.DB, error) {
	return database.New(database.Config{
		Host:          *host,
		Port:          *port,
		User:          *user,
		Password:      *password,
		DBName:        *db,
		MigrationsDir: *migrations,
		BackupsDir:    *backups,
	})
}

var DBCmd = &cobra.Command{
	Use: "db",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
	},
}

func init() {
	host = DBCmd.PersistentFlags().String("host", "localhost", "postgres host")
	port = DBCmd.PersistentFlags().Int("port", 5432, "postgres port")
	user = DBCmd.PersistentFlags().String("user", "postgres", "postgres user")
	password = DBCmd.PersistentFlags().String("password", "postgres", "postgres password")
	db = DBCmd.PersistentFlags().String("db", "postgres", "postgres database")
	migrations = DBCmd.PersistentFlags().String("migrations", "./tmp/migrations", "directory to store migrations")
	backups = DBCmd.PersistentFlags().String("backups", "./tmp/backups", "directory to store backups")

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := DBCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
