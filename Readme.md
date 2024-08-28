# Go OpenAI Helper

This project is a Go-based command-line tool that interacts with the OpenAI API to help developers with various tasks such as unit test generation, error handling, code completion, comment generation, code explanation, commit message generation, and code optimization.

## Features

- **Unit Test Generation**: Automatically generates unit test functions for your Go code.
- **Error Handling**: Enhances error handling in your Go code.
- **Code Completion**: Completes partial Go functions.
- **Comment Generation**: Adds comments to clarify your Go code.
- **Code Explanation**: Provides a brief explanation of your Go code.
- **Commit Message Generation**: Suggests descriptive commit messages.
- **Code Optimization**: Optimizes your Go code for better performance and readability.

## Prerequisites

- Go 1.16 or later
- An OpenAI API key

## Installation

### Clone the Repository

```bash
git clone https://github.com/yourusername/go-openai-helper.git
cd go-openai-helper
```

### Set Up the Environment
- Environment Variables: Create a .env file in the root directory of the project and add your OpenAI API key:

```bash

API_KEY=your_openai_api_key 
```
## Install Dependencies:

```bash

go mod tidy
```
### Running the Program
- To run the program, execute the following command in your terminal:
```bash
 go run main.go
```

### Available Options
- Once the program is running, you'll be prompted to choose an option from the following list:

- Unit Test Generation: Enter 1 to generate unit tests for the given Go code.
- Error Handling: Enter 2 to enhance the error handling in your Go code.
- Code Completion: Enter 3 to complete a partial Go function.
- Comment Generation: Enter 4 to add comments to your Go code.
- Code Explanation: Enter 5 to get a brief explanation of your Go code.
- Commit Message Generation: Enter 6 to generate a commit message for your Go code.
- Code Optimization: Enter 7 to optimize your Go code.