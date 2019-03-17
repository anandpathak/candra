package cmd

import (
	"errors"
	"fmt"
	"strings"

	utils "github.com/anandpathak/aws-ssh/utils"
	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "use add or list to add config or view config",
	Long: `Dude ! you need to pass add or view. something like

aws-ssh config add
aws-ssh config view

	`,
}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the current configuration",
	Long:  `Not sure what's your config? try list command to list down saved config`,
	Run: func(cmd *cobra.Command, args []string) {
		keys := viper.AllKeys()
		for _, value := range keys {
			fmt.Printf("%s : %s\n", value, viper.GetString(value))
		}

	},
}
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new config the current configuration",
	Long: `I can allow you to configure your application again
	
Set --secretAccessKey: Aws Secret keys
Set --accessKeyId: Aws Access Keys
Set --pem__filepath: Location where all the pemFile is stored
Set --default_region: AWS default region`,
	Run: func(cmd *cobra.Command, args []string) {
		err := setConfig()
		utils.Check(err)

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(listCmd, addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// take user input and set Configuration
func setConfig() error {
	var inputs = []utils.Query{
		utils.Query{
			Name:         "accessKey",
			Question:     "what's AWS AccessKey? ",
			DefaultValue: viper.GetString("accessKey"),
			AnswerType:   "string",
		},
		utils.Query{
			Name:         "secretKey",
			Question:     "what's AWS SecretKey?",
			DefaultValue: viper.GetString("secretKey"),
			AnswerType:   "string",
		},
		utils.Query{
			Name:         "region",
			Question:     "what's AWS Region?",
			DefaultValue: viper.GetString("region"),
			AnswerType:   "string",
		},
		utils.Query{
			Name:         "keyFileLocation",
			Question:     "where is the pemFiles stored ? ",
			DefaultValue: viper.GetString("keyFileLocation"),
			AnswerType:   "string",
		},
	}

	// check if the user provided input is valid
	isValid := func(answer string) error {
		if len(answer) < 0 {
			return errors.New("does not seems a valid input")
		}
		return nil
	}
	filter := func(answer string) (error, string) {
		return nil, strings.TrimSpace(answer)
	}
	for _, input := range inputs {
		input.Prompt()
		input.Filter(filter)
		err := input.Validate(isValid)
		utils.Check(err)
		viper.Set(input.Name, input.InType())
	}
	home, err := homedir.Dir()
	utils.Check(err)
	err2 := viper.WriteConfigAs(home + "/.aws-ssh.json")
	return err2
}
