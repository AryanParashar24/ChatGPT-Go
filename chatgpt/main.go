// Now here we will be using a lot of libraaries liek as godotenv to handle the .env file & then the other one is
// the chatgpt client by the company called pull request which helps in building chatgpt client for GoLang

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // to inport to godotenv package and then load my environment buildings

	apiKey := os.Getenv("API_KEY") //to import the apikey form the env file
	if apiKey == "" {
		log.Fatalln("Missing API Key") // here it will make sure if the api key is not missing in the applciation
	}

	ctx := context.Background()      //to work with APIs and clients who make sure we are using contexts
	client := gpt3.NewClient(apiKey) // now client is going to use gpt library which goves us access to the function called NewClient

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{ //here we will capture the response we got form the request and then handle any errors we might get
		Prompt:    []string{"The first thing that you should know about golang is"},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	}) // So here in the clien we will add the Completion request to take the context and then provide the response to the context.

	if err != nil {
		log.Fatalln(err) // if there is an error then it will log the error
	}
	fmt.Println(resp.Choices[0].Text)

}
