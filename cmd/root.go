package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Jeffail/gabs"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"k8s.io/apimachinery/pkg/util/yaml"
	sYaml "sigs.k8s.io/yaml"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "klabeler",
	Short: "Apply the current git hash to all k8s resources coming from STDIN",
	Long: `Apply the current git hash to all k8s resources coming from STDIN.
	
This is very useful if you want to track the desired state and whether some
resources in your cluster are maybe orphaned.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		bytes, err := ioutil.ReadAll(os.Stdin)

		splittedContents := strings.Split(string(bytes), "\n---")

		jsonObj := gabs.New()
		jsonObj.Array("output")

		for _, block := range splittedContents {

			// convert to json
			jsonData, err := yaml.ToJSON([]byte(block))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			jsonParsed, err := gabs.ParseJSON(jsonData)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			path := "."
			revision := "HEAD"

			r, err := git.PlainOpen(path)
			h, err := r.ResolveRevision(plumbing.Revision(revision))

			jsonParsed.Set(h.String(), "metadata", "labels", cmd.Flag("prefix").Value.String()+"git-hash")

			jsonObj.ArrayAppend(jsonParsed.Data(), "output")
		}
		rawYamlOutput, err := sYaml.JSONToYAML(jsonObj.Path("output").Bytes())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		yamlOutput := strings.TrimSpace(string(rawYamlOutput))
		if yamlOutput != "{}" && yamlOutput != "" {
			fmt.Println(yamlOutput)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.klabeler.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("prefix", "p", "klabeler.github.com/", "Label prefix")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".klabeler" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".klabeler")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
