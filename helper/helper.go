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
		qnsAndCode = "Generate a unit test function from the given code.Make sure the code is executable and remove the block comments if present.Also return only the testfile without the given function. Dont return the function again, return only the test function is needed \n" + code
	case "2":
		qnsAndCode = "Please generate a corrected Go file based on the provided code snippet. Ensure that any errors or issues in the original code are addressed, and provide comments or annotations where necessary to explain the changes made.Make sure the code is executable and remove the block comments if present.\n" + code
	case "3":
		qnsAndCode = "Please generate a Go file based on the provided code snippet. Ensure to complete the given function without main function  and provide comments or annotations where necessary to explain the changes made to complete the given function.The output should be a fully functional Go file with complete function.Make sure the code is executable and remove the block comments if present. \n" + code
	case "4":
		qnsAndCode = "Generate a Go file from the given code also add some comments to enhance clarity. Ensure that the returning should be a go file.Make sure the code is executable and remove the block comments if present.\n" + code

	case "5":
		qnsAndCode = "Go through the given code and explain it briefly\n" + code

	case "6":
		qnsAndCode = "Time to commit! Provide a descriptive message summarizing the changes you've made. What improvements or updates did you introduce in this commit?\n" + code

	case "7":
		qnsAndCode = "Please generate a corrected Go file based on the provided code snippet. Ensure that the code follows proper code optimisation, and provide comments or annotations where necessary to explain the changes made. The output should be a fully functional Go file with proper optimised code.Make sure the code is executable and remove the block comments if present.\n" + code

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
