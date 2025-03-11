package main

// So here in this project we will be entereing the code that we haev in the input file with code and then ask the gpt to review our code and then give the possible output OF TH LIBRARIES THAT OUR CODE MIGHT BE USING
import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
		panic("Missing API Key")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	const inputFile = "./input_with_code.txt"
	fileBytes, err := os.ReadFile(inputFile) // here we have some  code in our file which will be processed to get the output.
	if err != nil {
		log.Fatal("fatal to read file: %v", err)
	}

	msgPrefix := "give me a lis of libraries that are used in the code \n```python\n"
	msgSuffix := "\n```"
	msg := msgPrefix + string(fileBytes) + msgSuffix // now here it will pass the code not the fil along with the message prefix and message suffix

	outputBuilder := strings.Builder{} // here in the code it will be similar to string builder in the Java whihc is a a class that represents a mutable sequence of characters, or string, that can be modified at any time.
	err = client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			msg,
		},
		MaxTokens:   gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) { // here this function will take response which is a completion request
		outputBuilder.WriteString(resp.Choices[0].Text)
	})
	if err != nil {
		log.Fatal(err)
	}
	output := strings.TrimSpace(outputBuilder.String())         // because now here we want to start printing the output and in the output we need to trim the space in the outputBuilder as everything is written in the outputBuilder
	const outputFile = "./output.txt"                           // then we creat eout output file which is a constant in the output.txt
	err = os.WriteFile(outputFile, []byte(output), os.ModePerm) // now we are going to write in file using our outputFile then output and os.ModePerm
	// os.ModePerm is a constant in the os module of Python's standard library. It represents the mode used to create a file or directory with the maximum possible permissions.
	if err != nil {
		log.Fatal("failed to read file: %v", err)
	}
}
