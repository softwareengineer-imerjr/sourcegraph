package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
	"k8s.io/utils/strings/slices"

	"github.com/sourcegraph/conc/pool"
	"github.com/sourcegraph/run"
)

// These commands are meant to be executed with a VERSION env var with a hypothetical stamped release version
// This type is used to assign the stamp version from VERSION
type stampVersionKey struct{}

// Register upgrade commands -- see README.md for more details.
func main() {
	app := &cli.App{
		Name:  "upgrade-test",
		Usage: "Upgrade test is a tool for testing the migrator services creation of upgrade paths and application of upgrade paths.\nWhen run relevant upgrade paths are tested for each version relevant to a given upgrade type, initializing Sourcegraph databases and frontend services for each version, and attempting to generate and apply an upgrade path to your current branches head.",
		Commands: []*cli.Command{
			{
				Name:    "all-types",
				Aliases: []string{"all"},
				Usage:   "Runs all upgrade test types\n\nRequires stamp-version for tryAutoUpgrade call.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "stamp-version",
						Aliases:  []string{"sv"},
						Usage:    "stamp-version is the version frontend:candidate and  migrator:candidate are set as. If the $VERSION env var is set this flag inherits that value.",
						EnvVars:  []string{"VERSION"},
						Required: true,
					},
				},
				Action: func(cCtx *cli.Context) error {
					ctx := context.WithValue(cCtx.Context, stampVersionKey{}, cCtx.String("stamp-version"))

					// check docker is running
					if err := run.Cmd(ctx, "docker", "ps").Run().Wait(); err != nil {
						fmt.Println("🚨 Error: could not connect to docker: ", err)
						os.Exit(1)
					}

					// Get init versions to use for initializing upgrade environments for tests
					latestMinorVersion, latestStableVersion, targetVersion, stdVersions, mvuVersions, autoVersions, err := getVersions(cCtx)
					if err != nil {
						fmt.Println("🚨 Error: failed to get test version ranges: ", err)
						os.Exit(1)
					}

					fmt.Println("Latest stable release version: ", latestStableVersion)
					fmt.Println("Latest minor version: ", latestMinorVersion)
					fmt.Println("Target version: ", targetVersion)
					fmt.Println("Standard Versions:", stdVersions)
					fmt.Println("Multiversion Versions:", mvuVersions)
					fmt.Println("Autoupgrade Versions:", autoVersions)

					// initialize test results
					var results TestResults

					// create array of all tests
					var versions []typeVersion
					for _, version := range stdVersions {
						versions = append(versions, typeVersion{
							Type:    "std",
							Version: version,
						})
					}
					for _, version := range mvuVersions {
						versions = append(versions, typeVersion{
							Type:    "mvu",
							Version: version,
						})
					}
					for _, version := range autoVersions {
						versions = append(versions, typeVersion{
							Type:    "auto",
							Version: version,
						})
					}

					// Run all test types
					testPool := pool.New().WithMaxGoroutines(10).WithErrors()
					for _, version := range versions {
						version := version
						if slices.Contains(knownBugVersions, version.Version.String()) {
							continue
						}

						switch version.Type {
						case "std":
							testPool.Go(func() error {
								fmt.Println("std: ", version.Version)
								start := time.Now()
								result := standardUpgradeTest(ctx, version.Version, targetVersion, latestStableVersion)
								result.Runtime = time.Since(start)
								results.AddStdTest(result)
								return nil
							})
						case "mvu":
							testPool.Go(func() error {
								fmt.Println("mvu: ", version.Version)
								start := time.Now()
								result := multiversionUpgradeTest(ctx, version.Version, targetVersion, latestStableVersion)
								result.Runtime = time.Since(start)
								results.AddMVUTest(result)
								return nil
							})
						case "auto":
							testPool.Go(func() error {
								fmt.Println("auto: ", version.Version)
								start := time.Now()
								result := autoUpgradeTest(ctx, version.Version, targetVersion, latestStableVersion)
								result.Runtime = time.Since(start)
								results.AddAutoTest(result)
								return nil
							})
						}
					}
					if err := testPool.Wait(); err != nil {
						log.Fatal(err)
					}

					// This is where we do the majority of our printing to stdout.
					results.OrderByVersion()
					results.PrintSimpleResults()

					return nil
				},
			},
			{
				Name:    "standard",
				Aliases: []string{"std"},
				Usage:   "Runs standard upgrade tests for all patch versions from the last minor version.\nEx: 5.1.x -> 5.2.x (head)",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "stamp-version",
						Aliases: []string{"sv"},
						Usage:   "stamp-version is the version frontend:candidate and  migrator:candidate are set as. If the $VERSION env var is set this flag inherits that value.",
						EnvVars: []string{"VERSION"},
					},
				},
				Action: func(cCtx *cli.Context) error {
					ctx := context.WithValue(cCtx.Context, stampVersionKey{}, cCtx.String("stamp-version"))

					// check docker is running
					if err := run.Cmd(ctx, "docker", "ps").Run().Wait(); err != nil {
						fmt.Println("🚨 Error: could not connect to docker: ", err)
						os.Exit(1)
					}

					// Get init versions to use for initializing upgrade environments for tests
					latestMinorVersion, latestStableVersion, targetVersion, stdVersions, _, _, err := getVersions(cCtx)
					if err != nil {
						fmt.Println("🚨 Error: failed to get test version ranges: ", err)
						os.Exit(1)
					}

					fmt.Println("Latest stable release version: ", latestStableVersion)
					fmt.Println("Latest minor version: ", latestMinorVersion)
					fmt.Println("Target version: ", targetVersion)
					fmt.Println("Standard Versions:", stdVersions)

					// initialize test results
					var results TestResults

					// Run Standard Upgrade Tests in goroutines. The current limit is set as 10 concurrent goroutines per test type (std, mvu, auto). This is to address
					// dynamic port allocation issues that occur in docker when creating many bridge networks, but tests begin to fail when a sufficient number of
					// goroutines are running on local machine. We may tune this in CI.
					// TODO this should likely be made an env var or something to make it easy to swamp out depending on the test box.
					stdTestPool := pool.New().WithMaxGoroutines(10).WithErrors()
					for _, version := range stdVersions {
						version := version
						if slices.Contains(knownBugVersions, version.String()) {
							continue
						}
						stdTestPool.Go(func() error {
							fmt.Println("std: ", version)
							start := time.Now()
							result := standardUpgradeTest(ctx, version, targetVersion, latestStableVersion)
							result.Runtime = time.Since(start)
							result.DisplayLog() // DEBUG
							results.AddStdTest(result)
							return nil
						})
					}
					if err := stdTestPool.Wait(); err != nil {
						log.Fatal(err)
					}

					// This is where we do the majority of our printing to stdout.
					results.OrderByVersion()
					results.PrintSimpleResults()

					return nil
				},
			},
			{
				Name:    "multiversion",
				Aliases: []string{"mvu"},
				Usage:   "Runs multiversion upgrade tests for all versions which would require a multiversion upgrade to reach your current repo head. i.e those versions more than a minor version behind the last minor release.\nEx: 3.4.1 -> 5.2.6",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "stamp-version",
						Aliases: []string{"sv"},
						Usage:   "stamp-version is the version frontend:candidate and  migrator:candidate are set as. If the $VERSION env var is set this flag inherits that value.",
						EnvVars: []string{"VERSION"},
					},
				},
				Action: func(cCtx *cli.Context) error {
					ctx := context.WithValue(cCtx.Context, stampVersionKey{}, cCtx.String("stamp-version"))

					// check docker is running
					if err := run.Cmd(ctx, "docker", "ps").Run().Wait(); err != nil {
						fmt.Println("🚨 Error: could not connect to docker: ", err)
						os.Exit(1)
					}

					// Get init versions to use for initializing upgrade environments for tests
					latestMinorVersion, latestStableVersion, targetVersion, _, mvuVersions, _, err := getVersions(cCtx)
					if err != nil {
						fmt.Println("🚨 Error: failed to get test version ranges: ", err)
						os.Exit(1)
					}

					fmt.Println("Latest stable release version: ", latestStableVersion)
					fmt.Println("Latest minor version: ", latestMinorVersion)
					fmt.Println("Target version: ", targetVersion)
					fmt.Println("MVU Versions:", mvuVersions)

					// initialize test results
					var results TestResults

					// Run MVU Upgrade Tests
					mvuTestPool := pool.New().WithMaxGoroutines(10).WithErrors()
					for _, version := range mvuVersions {
						version := version
						if slices.Contains(knownBugVersions, version.String()) {
							continue
						}
						mvuTestPool.Go(func() error {
							fmt.Println("mvu: ", version)
							start := time.Now()
							result := multiversionUpgradeTest(ctx, version, targetVersion, latestStableVersion)
							result.Runtime = time.Since(start)
							results.AddMVUTest(result)
							return nil
						})
					}
					if err := mvuTestPool.Wait(); err != nil {
						log.Fatal(err)
					}

					results.OrderByVersion()
					results.PrintSimpleResults()

					return nil
				},
			},
			{
				Name:    "autoupgrade",
				Aliases: []string{"auto"},
				Usage:   "Runs autoupgrade upgrade tests for all versions.\n\nRequires stamp-version for tryAutoUpgrade call.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "stamp-version",
						Aliases:  []string{"sv"},
						Usage:    "stamp-version is the version frontend:candidate and  migrator:candidate are set as. If the $VERSION env var is set this flag inherits that value.",
						EnvVars:  []string{"VERSION"},
						Required: true,
					},
				},
				Action: func(cCtx *cli.Context) error {
					ctx := context.WithValue(cCtx.Context, stampVersionKey{}, cCtx.String("stamp-version"))

					// check docker is running
					if err := run.Cmd(ctx, "docker", "ps").Run().Wait(); err != nil {
						fmt.Println("🚨 Error: could not connect to docker: ", err)
						os.Exit(1)
					}

					// Get init versions to use for initializing upgrade environments for tests
					latestMinorVersion, targetVersion, latestStableVersion, _, _, autoVersions, err := getVersions(cCtx)
					if err != nil {
						fmt.Println("🚨 Error: failed to get test version ranges: ", err)
						os.Exit(1)
					}

					fmt.Println("Latest stable release version: ", latestStableVersion)
					fmt.Println("Latest minor version: ", latestMinorVersion)
					fmt.Println("Auto Versions:", autoVersions)

					// initialize test results
					var results TestResults

					// Run Autoupgrade Tests
					autoTestPool := pool.New().WithMaxGoroutines(10).WithErrors()
					for _, version := range autoVersions {
						version := version
						if slices.Contains(knownBugVersions, version.String()) {
							continue
						}
						autoTestPool.Go(func() error {
							fmt.Println("auto: ", version)
							start := time.Now()
							result := autoUpgradeTest(ctx, version, targetVersion, latestStableVersion)
							result.Runtime = time.Since(start)
							results.AddAutoTest(result)
							result.DisplayLog() // DEBUG
							return nil
						})
					}
					if err := autoTestPool.Wait(); err != nil {
						log.Fatal(err)
					}

					results.OrderByVersion()
					results.PrintSimpleResults()

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
