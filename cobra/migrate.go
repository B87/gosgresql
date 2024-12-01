package cobra

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	createtype = new(string)
)

func init() {
	DBCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(upCmd)
	migrateCmd.AddCommand(downCmd)
	migrateCmd.AddCommand(statusCmd)
	migrateCmd.AddCommand(createCmd)
	migrateCmd.AddCommand(resetCmd)

	createCmd.Flags().StringVarP(createtype, "type", "t", "sql", "Type of the migration")
}

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
	},
}

var upCmd = &cobra.Command{
	Use: "up",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := newDB()
		if err != nil {
			cmd.PrintErrf("Failed to connect to database: %v", err)
			return
		}
		defer db.Close()

		err = db.Migrator.Up()
		if err != nil {
			cmd.PrintErrf("Failed to migrate up: %v", err)
			return
		}
	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Migrate the database down",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := newDB()
		if err != nil {
			cmd.PrintErrf("Failed to connect to database: %v\n", err)
			os.Exit(1)
			return
		}
		defer db.Close()

		err = db.Migrator.Down()
		if err != nil {
			cmd.PrintErrf("Failed to migrate down: %v\n", err)
			os.Exit(1)
			return
		}
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show migration status",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new migration file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		db, err := newDB()
		if err != nil {
			cmd.PrintErrf("Failed to connect to database: %v\n", err)
			os.Exit(1)
			return
		}
		defer db.Close()

		err = db.Migrator.NewMigration(name, *createtype)
		if err != nil {
			cmd.PrintErrf("Failed to create migration: %v\n", err)
			os.Exit(1)
			return
		}
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the database (down + up)",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := newDB()
		if err != nil {
			cmd.PrintErrf("Failed to connect to database: %v\n", err)
			os.Exit(1)
			return
		}
		defer db.Close()

		err = db.Migrator.Down()
		if err != nil {
			cmd.PrintErrf("Failed to migrate down: %v\n", err)
			os.Exit(1)
			return
		}

		err = db.Migrator.Up()
		if err != nil {
			cmd.PrintErrf("Failed to migrate up: %v\n", err)
			os.Exit(1)
			return
		}
	},
}
