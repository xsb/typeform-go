# typeform-go

Go library for [Typeform](http://www.typeform.com) APIs.

Implements:
- [Typeform I/O](http://docs.typeform.io/)
- [Typeform.com Data API](http://helpcenter.typeform.com/hc/en-us/articles/200071986-Data-API)

## Basic Examples

### Creating a simple typeform

```go
package main

import (
	"fmt"

	"github.com/xsb/typeform-go/tfio"
)

func main() {

	tfio.ApiToken = "<Your typeform.io API Token>"

	// Short Text field
	f1 := tfio.ShortText()
	f1.Question = "What's your name?"
	f1.Required = true

	// Multiple Choice field
	f2 := tfio.MultipleChoice()
	f2.Question = "Choose your favourite programming language"
	f2.Description = "It's Go, isn't it?"
	choices := []string{"Go", "Java", "Ruby", "JavaScript", "Python", "Perl"}
	for _, choice := range choices {
		f2.Choices = append(f2.Choices, tfio.Choice{choice})
	}

	// Create a typeform with the previous fields
	form := tfio.NewForm("My First Form")
	form.Fields = append(form.Fields, f1)
	form.Fields = append(form.Fields, f2)
	output, _ := form.CreateForm()
	fmt.Println(string(output))
}
```

### Getting submissions from the Data API

```go
package main

import (
	"fmt"

	"github.com/xsb/typeform-go/tfcom"
)

func main() {

	tfcom.ApiKey = "<Your typeform.com API Key>"
	formUid := "Form UID"

	b, err := tfcom.GetResponses(formUid)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
```
