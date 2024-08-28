package helper

import (
	"fmt"
	"genAi-with-go/openAI"
	"os"
)

func Selection() {
	apiKey := os.Getenv("API_KEY")
	content, err := os.ReadFile("code.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	code := string(content)

	fmt.Println("How can i help you:??? \n Select : -> \n 1 - UnitTest \n 2 - ErrorHandling \n 3 - Code completion \n 4 - Comment generation \n 5 - Explain the Given code \n 6 - Generate a commit message for the given code \n 7 - Code Optimisation")
	var input string
	fmt.Scanln(&input)
	fmt.Println("UR INPUT IS", input)
	var qnsAndCode string

	switch input {
	case "1":
		qnsAndCode = UnitTest + code
	case "2":
		qnsAndCode = ErrorHandling + code
	case "3":
		qnsAndCode =CodeCompletion + code
	case "4":
		qnsAndCode = CommentGeneration + code

	case "5":
		qnsAndCode = Explanation + code

	case "6":
		qnsAndCode = GeneratingCommitMessage + code

	case "7":
		qnsAndCode = CodeOptimisation + code

	default:
		fmt.Println("Enter a valid number")

	}

	fmt.Println("YOUR QUESTION IS", qnsAndCode)
	answer, err := openAI.GetAnswerFromOpenAI(qnsAndCode, apiKey)
	if err != nil {
		fmt.Println(err)
	}
	answer = filteringIntoProperGOfile(answer, apiKey)
	if input == "1" {
		writingIntoFile("./a_test.go", answer)

	} else if input == "5" {
		writingIntoFile("./output.txt", answer)

	} else {

		writingIntoFile("./code.go", answer)

	}
}

//52463

func filteringIntoProperGOfile(answer, apiKey string) string {
	fmt.Println("entered filtering")
	fmt.Println(answer)
	qnsAndCode := "Make sure the code is properly executable and remove block comments(like ```go ) if present\n The code is: \n" + answer

	correctedAns, err := openAI.GetAnswerFromOpenAI(qnsAndCode, apiKey)
	if err != nil {
		fmt.Println(err)
	}
	return correctedAns
}

func writingIntoFile(filename, answer string) {
	if err := os.WriteFile(filename, []byte(answer), 0644); err != nil {
		fmt.Println(err)
	}

}
