package tfio

type Field struct {
	FieldType   string   `json:"type"`
	Question    string   `json:"question"`
	Description string   `json:"description,omitempty"`
	Required    bool     `json:"required,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type FieldDescriber interface {
	ToJson() []byte
}

type Choice struct {
	Choice string `json:"label"`
}

//Field: Short Text

type ShortTextField struct {
	Field
}

func ShortText() ShortTextField {
	f := ShortTextField{}
	f.FieldType = "short_text"
	return f
}

//Field: Long Text

type LongTextField struct {
	Field
}

func LongText() LongTextField {
	f := LongTextField{}
	f.FieldType = "long_text"
	return f
}

//Field: Multiple Choice

type MultipleChoiceField struct {
	Field
	Choices []Choice `json:"choices"`
}

func MultipleChoice() MultipleChoiceField {
	mc := MultipleChoiceField{}
	mc.FieldType = "multiple_choice"
	return mc
}

//Field: Statement

type StatementField struct {
	Field
}

func Statement() StatementField {
	st := StatementField{}
	st.FieldType = "statement"
	return st
}

//Field: Dropdown

type DropdownField struct {
	Field
	Choices []Choice `json:"choices"`
}

func Dropdown() DropdownField {
	dr := DropdownField{}
	dr.FieldType = "dropdown"
	return dr
}

//Field: Yes/No

type YesNoField struct {
	Field
}

func YesNo() YesNoField {
	yn := YesNoField{}
	yn.FieldType = "yes_no"
	return yn
}

//Field: Number

type NumberField struct {
	Field
}

func Number() NumberField {
	num := NumberField{}
	num.FieldType = "number"
	return num
}
