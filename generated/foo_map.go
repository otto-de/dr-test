package generated

type generateFn = func(amount int) *[]interface{}
func Generate(structName string) generateFn {
	return func(amount int) *[]interface{} {
		return &[]interface{}{}
	}
}
