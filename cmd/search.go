package cmd

import (
	"fmt"

	"github.com/anandpathak/candra/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var valueFlag string
var tagFlag string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called", valueFlag, tagFlag)
		var config = aws.Config{
			Region:      aws.String(viper.GetString("region")),
			Credentials: credentials.NewStaticCredentials(viper.GetString("accessKey"), viper.GetString("secretKey"), ""),
		}
		servers := utils.GetServersList(config, tagFlag, valueFlag)
		for i, server := range servers {
			fmt.Printf("%d. %s\n", i, server.Name)
		}
		var query = utils.Query{
			Name:         "selectedServer",
			Question:     "which server you want to select",
			DefaultValue: "0",
			AnswerType:   "int",
		}
		query.Prompt()
		i := query.InType().(int)
		fmt.Println(servers[i])
		var keyname = viper.GetString("keyFileLocation") + "/" + servers[i].PemKey + ".pem"
		//utils.Login(keyname, "ubuntu", servers[i].PublicIP, "22")
		utils.Commando(keyname, servers[i].PublicIP, "ubuntu")

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")
	searchCmd.PersistentFlags().StringVarP(&valueFlag, "value", "v", "*", "add the search keywords")
	searchCmd.PersistentFlags().StringVarP(&tagFlag, "tag", "t", "tag:Name", "search by ? ")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
