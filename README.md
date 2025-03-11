# ChatGPT-Go

ChatGPT-Go is a simple, easy-to-use Go-based application that interacts with OpenAI's ChatGPT API to generate human-like conversations. The project is designed to help developers integrate AI-powered chat functionalities into their Go applications using OpenAI's GPT model.

## Features:
  - **API Integration**: Connects to OpenAI's GPT-3.5 or GPT-4 API.
  - **Human-Like Conversations**: Generates responses based on user input in a natural, human-like manner.
  - **Simple Implementation**: The application is built in Go, making it easy to integrate and extend.
  - **Environment Configuration**: Uses `.env` files for easy configuration of API keys and other environment variables.

## Prerequisites:
Before you begin, ensure that you have met the following requirements:
  - **Go 1.18+**: Download and install Go from [https://golang.org/dl/](https://golang.org/dl/).
  - **OpenAI API Key**: You can obtain an API key by signing up at [https://beta.openai.com/signup/](https://beta.openai.com/signup/).
  - **Git**: Install Git from [https://git-scm.com/](https://git-scm.com/).

## Installation:

### Step 1: Clone the Repository:
Clone this repository to your local machine:
```bash
git clone https://github.com/AryanParashar24/ChatGPT-Go.git
cd ChatGPT-Go
```

### Step-2: Set Up Your Environment:
Create a .env file in the root of the project directory. Add the OpenAI API key like this:

```bash
OPENAI_API_KEY=your-api-key-here
```
Replace your-api-key-here with the actual API key you obtained from OpenAI.

### Step-3: Install Dependencies: 
Ensure that Go dependencies are installed by running:

``` bash
go mod tidy
```

### Step 4: Run the Application:
Once everything is set up, you can start the application by running:

``` bash
go run main.go
```
The program will prompt you for input, and ChatGPT will respond with generated replies.
