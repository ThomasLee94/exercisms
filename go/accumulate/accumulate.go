package accumulate

func Accumulate(collection []string, operation func(string) string) (result []string) {
	for _, element := range collection {
		anonymousFuncResult := operation(element)
		result = append(result, anonymousFuncResult)
	}
	return result
}
