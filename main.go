package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"

	rest "github.com/google/go-github/v29/github"
	graphql "github.com/shurcooL/githubv4"
)

var (
	// options
	help     bool
	cfg      string
	hostname string
	token    string

	// -----

	httpClient    *http.Client
	restClient    *rest.Client
	graphqlClient *graphql.Client

	ctx = context.Background()

	blue  = color.FgBlue.Render
	green = color.FgGreen.Render
	red   = color.FgRed.Render

	bold   = color.OpBold.Render
	dimmed = color.OpFuzzy.Render

	//  -----

	output string
)

func init() {
	// flags
	pflag.BoolVar(&help, "help", false, "print this help")
	pflag.StringVarP(&cfg, "config", "c", "", "path to the YAML config file (defaults to $HOME/)")
	pflag.StringVarP(&hostname, "hostname", "h", "", "hostname")
	pflag.StringVarP(&token, "token", "t", "", "personal access token")
	// TODO
	pflag.Parse()

	// read config
	viper.SetConfigName(".ghe-migration-info")
	viper.SetConfigType("yml")

	if cfg != "" {
		viper.AddConfigPath(cfg)
	} else {
		viper.AddConfigPath("$HOME")
	}

	if err := viper.ReadInConfig(); err != nil && cfg != "" {
		printHelpOnError(
			fmt.Sprintf("config file .ghe-migration-info not found in %s", cfg),
		)
	}
	viper.BindPFlags(pflag.CommandLine)

	// assign values
	help = viper.GetBool("help")
	hostname = viper.GetString("hostname")
	token = viper.GetString("token")

	// validate
	validateFlags()

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	httpClient = oauth2.NewClient(ctx, src)

	graphqlURL := fmt.Sprintf("https://%s/api/graphql", hostname)
	graphqlClient = graphql.NewEnterpriseClient(graphqlURL, httpClient)

	restURL := fmt.Sprintf("https://%s/api/v3", hostname)
	restClient, _ = rest.NewEnterpriseClient(restURL, restURL, httpClient)
}

func main() {
	output = getAdminStats()
	output += "\n---------------------------\n"
	output += getTotalDiskUsage()

	fmt.Printf("%s", output)
}

// helpers -------------------------------------------------------------------------------------------------------------

func validateFlags() {
	if help {
		printHelp()
		os.Exit(0)
	}

	if hostname == "" {
		printHelpOnError("hostname missing")
	}

	if hostname == "github.com" {
		printHelpOnError("github.com is not supported")
	}

	if token == "" {
		printHelpOnError("token missing")
	}
}

func printHelp() {
	fmt.Println(`USAGE:
  ghe-get-all-owners [OPTIONS]

OPTIONS:`)
	pflag.PrintDefaults()
	fmt.Println(`
EXAMPLE:
  $ ghe-get-all-owners -h github.example.com -t AA123...`)
	fmt.Println()
}

func printHelpOnError(s string) {
	printHelp()
	errorAndExit(errors.New(s))
}

func errorAndExit(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(2)
}
