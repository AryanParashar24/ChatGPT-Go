package main

// Here in this project we'll be using viper and cobra where viper will be used for working with the environments and cobra will be used for working with the CLI
import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetResponse(client gpt3.Client, ctx context.Context, question string) {
	err := client.CompletionStreamWithEngine(ctx, gpt3.BabbageEngine, gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		MaxTokens:   gpt3.IntPtr(3000), // to decide how much of the tokens can the chatgpt provide abck in the repsonse to our request
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) { // to edit and manage with the choices we will get from the chatgpt after entering request
		fmt.Print(resp.Choices[0].Text) // out of all the choices we will get we want to access the 0th choice in the options
	})
	if err != nil { // in case we gets an error
		fmt.Println(err)
		os.Exit(13)
	}
	fmt.Printf("\n")
}

type NullWriter int

// NullWriter is a succeeding dummy function that does nothing and always succeeds.
// Here we are using it for the Logging purposes

func (NullWriter) Write([]byte) (int, error) {
	return 0, nil
}

func main() {
	log.SetOutput(new(NullWriter))       //
	viper.SetConfigFile(".env")          // to set the config file using the viper package
	viper.ReadInConfig()                 // to read the config file from the viper package
	apiKey := viper.GetString("API_KEY") //to get our chatgpt api key
	if apiKey == "" {
		panic("Missing API Key") // to check if our api key is missing in the file or not
	}

	ctx := context.Background()      //for working with the context
	client := gpt3.NewClient(apiKey) //for using the chatgpt client
	rootCmd := &cobra.Command{       // here Cobra will help you with the commands and arguments and will help us in building the CLI tool
		Use:   "chatgpt",                      // to use chatgpt
		Short: "Chat with chatgpt inn cosole", // description and text to ask & guide the user
		Run: func(cmd *cobra.Command, args []string) { //funciton to run and access the command and request been given to the chatgpt and then response in the similar manner
			scanner := bufio.NewScanner(os.Stdin) // to scan for the input
			quit := false                         // and set the quit to false to manage with the quit status across the place

			for !quit {
				fmt.Print("Say something ('quit' to end): ")
				if !scanner.Scan() {
					break
				}
				question := scanner.Text() // to assign the question whtever the user puts in
				switch question {
				case "quit": // switch case in order the user doesnt quit
					quit = true // here as we can see the user puts and sets the quit to True in case the user forgets to quit

				default:
					GetResponse(client, ctx, question) //its main functioning is to just call the chatgpt client and ask the questions been entered by the client
				}
			}
		},
	}
	rootCmd.Execute()
}
