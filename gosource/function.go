package gosource

// Function describe the content to a scopre function.
type Function struct {
	name    string
	args    []Parameter
	rets    []Parameter
	content string
}

// NewFunction return a fresh function without arg and return element.
func NewFunction(name string) *Function {
	return &Function{
		name: name,
	}
}

// SetContent function.
func (f *Function) SetContent(content string) {
	f.content = content
}

// AddArg on the parameters function.
func (f *Function) AddArg(name, nameType string) {
	f.args = append(f.args, Parameter{
		Name: name,
		Type: nameType,
	})
}

// AddRet on the returns function.
func (f *Function) AddRet(name, nameType string) {
	f.rets = append(f.rets, Parameter{
		Name: name,
		Type: nameType,
	})
}
