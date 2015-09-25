# typeform-go

Go library for Typeform APIs

## Example

```go
package main

import (
	"github.com/xsb/typeform-go/tfio"
)

func main() {

	tfio.ApiToken = "<Your typeform.io API Token>"

	//Create Short Field
	f1 := tfio.ShortText()
	f1.Question = "What's your name?"
	f1.Required = true

	//Create Long Field
	f2 := tfio.LongText()
	f2.Question = "Where do you live?"
	f2.Description = "In the Internet is not a valid answer ;)"

	//Create Statement
	f3 := tfio.Statement()
	f3.Question = "This is a statement"
	f3.Description = "We are not expecting an answer for it"

	//Create Multiple Choice
	f4 := tfio.MultipleChoice()
	f4.Question = "Choose your favourite fruit"
	choice1 := tfio.Choice{"banana"}
	choice2 := tfio.Choice{"orange"}
	choice3 := tfio.Choice{"pineapple"}
	choice4 := tfio.Choice{"apple"}
	f4.Choices = []tfio.Choice{choice1, choice2, choice3}

	//Create Dropdown
	f5 := tfio.Dropdown()
	f5.Question = "Choose a Programming Language"
	choice1 = tfio.Choice{"Go"}
	choice2 = tfio.Choice{"Ruby"}
	choice3 = tfio.Choice{"Perl"}
	choice4 = tfio.Choice{"JavaScript"}
	f5.Choices = []tfio.Choice{choice1, choice2, choice3, choice4}

	//Create Yes No Field
	f6 := tfio.YesNo()
	f6.Question = "Are you sure?"

	//Create Number Field
	f7 := tfio.Number()
	f7.Question = "How old are you?"

	//Make a Description of a Form
	fd := tfio.FormDescription{}
	fd.Title = "My first form"
	fd.Fields = append(fd.Fields, f1)
	fd.Fields = append(fd.Fields, f2)
	fd.Fields = append(fd.Fields, f3)
	fd.Fields = append(fd.Fields, f4)
	fd.Fields = append(fd.Fields, f5)
	fd.Fields = append(fd.Fields, f6)
	fd.Fields = append(fd.Fields, f7)

	//Create a Form
	err := fd.CreateForm()
	if err != nil {
		panic(err)
	}
}
```
