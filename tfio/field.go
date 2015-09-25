package tfio

//Field Types

type Field struct {
	FieldType   string   `json:"type"`
	Question    string   `json:"question"`
	Description string   `json:"description,omitempty"`
	Required    bool     `json:"required,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type Choice struct {
	Choice string `json:"label"`
}

type ShortTextField struct {
	Field
}

type LongTextField struct {
	Field
}

type MultipleChoiceField struct {
	Field
	Choices []Choice `json:"choices"`
}

type StatementField struct {
	Field
}

type DropdownField struct {
	Field
	Choices []Choice `json:"choices"`
}

type YesNoField struct {
	Field
}

type NumberField struct {
	Field
}

//Field Constructors

func ShortText() ShortTextField {
	f := ShortTextField{}
	f.FieldType = "short_text"
	return f
}

func LongText() LongTextField {
	f := LongTextField{}
	f.FieldType = "long_text"
	return f
}

func MultipleChoice() MultipleChoiceField {
	f := MultipleChoiceField{}
	f.FieldType = "multiple_choice"
	return f
}

func Statement() StatementField {
	f := StatementField{}
	f.FieldType = "statement"
	return f
}

func Dropdown() DropdownField {
	f := DropdownField{}
	f.FieldType = "dropdown"
	return f
}

func YesNo() YesNoField {
	f := YesNoField{}
	f.FieldType = "yes_no"
	return f
}

func Number() NumberField {
	f := NumberField{}
	f.FieldType = "number"
	return f
}
