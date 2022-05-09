### cobra Learn
1. [x] https://github.com/spf13/cobra-cli/blob/main/README.md

1.Installing
Using Cobra is easy. First, use go get to install the latest version of the library.
```shell
go get -u github.com/spf13/cobra@latest
```

Next, include Cobra in your application:
```shell
import "github.com/spf13/cobra"
```

Install cli:
```shell
go install github.com/spf13/cobra-cli@latest
```

2.Usage:

```shell
cd $HOME/code/myapp
cobra-cli init
go run main.go
```
Optional flags:
You can provide it your author name with the --author flag. e.g. cobra-cli init --author "Steve Francia spf@spf13.com"
You can provide a license to use with --license e.g. cobra-cli init --license apache
Use the --viper flag to automatically setup viper
Viper is a companion to Cobra intended to provide easy handling of environment variables and config files and seamlessly connecting them to the application flags.
Add commands to a project
Once a cobra application is initialized you can continue to use the Cobra generator to add additional commands to your application. The command to do this is cobra-cli add.

Let's say you created an app and you wanted the following commands for it:

```
app serve
app config
app config create
```
In your project directory (where your main.go file is) you would run the following:
```shell
cobra-cli add serve
cobra-cli add config
cobra-cli add create -p 'configCmd'
```
`cobra-cli` add supports all the same optional flags as `cobra-cli init` does (described above).

You'll notice that this final command has a `-p` flag. This is used to assign a parent command to the newly added command. In this case, we want to assign the "create" command to the "config" command. All commands have a default parent of rootCmd if not specified.

By default cobra-cli will append Cmd to the name provided and uses this name for the internal variable name. When specifying a parent, be sure to match the variable name used in the code.

Note: Use camelCase (not snake_case/kebab-case) for command names. Otherwise, you will encounter errors. For example, cobra-cli add add-user is incorrect, but cobra-cli add addUser is valid.