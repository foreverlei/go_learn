/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"learn/cobra_viper/global"
	"os"
)

var (
	cfgFile string
	author  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra_viper",
	Short: "cobra_viper --config a.yaml  -t=false ",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		tog, err := cmd.Flags().GetBool("toggle")
		if err != nil {
			fmt.Println("get bool err:", err)
		}
		fmt.Println("Root run!", tog, cfgFile, author)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra_viper.yaml)")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $pwd/.cobra_viper.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// usage: -t=true
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	fmt.Println("root init:")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		//home, err := os.UserHomeDir()
		//home, err := os.Getwd()
		//cobra.CheckErr(err)

		// Search config in home directory with name ".cobra_viper" (without extension).
		//viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("cobra_viper")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := viper.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(global.Config.Name, global.Config.Author)
	})
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic("Fatal Unmarshal config file")
	}
	fmt.Println(global.Config.Name, global.Config.Author)
	select {}
}
