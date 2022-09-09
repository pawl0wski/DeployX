package prompts

func convertValToString(val interface{}) string {
	valAsString, ok := val.(string)
	if !ok {
		panic("error while validating value")
	}
	return valAsString
}
