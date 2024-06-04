package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"ariga.io/atlas-go-sdk/atlasexec"
)

func main() {
	// Define the execution context, supplying a migration directory
	// and potentially an `atlas.hcl` configuration file using `atlasexec.WithHCL`.
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(
			os.DirFS("./app/cmd/migrate/migrations"),
		),
	)
	// atlasexec works on a temporary directory, so we need to close it
	defer func(workdir *atlasexec.WorkingDir) {
		err := workdir.Close()
		if err != nil {
			return
		}
	}(workdir)

	if err != nil {
		log.Panicf("failed to load working directory: %v", err)
	}

	// Initialize the client.
	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		log.Panicf("failed to initialize client: %v", err)
	}
	// Run `atlas migrate apply`
	hostPort := net.JoinHostPort(os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"))

	res, err := client.MigrateApply(context.Background(), &atlasexec.MigrateApplyParams{
		URL: fmt.Sprintf(
			"mysql://%s:%s@%s/%s",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			hostPort,
			os.Getenv("MYSQL_DATABASE"),
		),
		Env:             "",
		ConfigURL:       "",
		Context:         nil,
		DirURL:          "",
		AllowDirty:      false,
		RevisionsSchema: "",
		BaselineVersion: "",
		TxMode:          "",
		ExecOrder:       "",
		Amount:          0,
		DryRun:          false,
		Vars:            nil,
	})
	if err != nil {
		log.Panicf("failed to apply migrations: %v", err)
	}

	log.Printf("Applied %d migrations\n", len(res.Applied))
}
