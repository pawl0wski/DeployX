package prompts

func convertValToString(val interface{}) string {
	valAsString, ok := val.(string)
	if !ok {
		panic("Can't validate value")
	}
	return valAsString
}
