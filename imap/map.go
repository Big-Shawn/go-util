package imap

const CaseUpper = 1
const CaseLower = 2

func ChangeKeyCase[V any](maps map[string]V, caseOption int) map[string]V {
	newMap := make(map[string]V)
	var fn func(key string) string
	switch caseOption {
	case CaseUpper:
		fn = changeKeyCaseUpper
	case CaseLower:
		fn = changeKeyCaseLower
	}
	for key, val := range maps {
		newKey := fn(key)
		newMap[newKey] = val
	}
	return newMap
}

func changeKeyCaseLower(key string) string {
	newStr := ""
	for _, v := range key {
		if v >= 'A' && v <= 'Z' {
			v = v - 'A' + 'a'
		}
		newStr += string(v)
	}
	return newStr
}

func changeKeyCaseUpper(key string) string {
	newStr := ""
	for _, v := range key {
		if v >= 'a' && v <= 'z' {
			v = v - 'a' + 'A'
		}
		newStr += string(v)
	}
	return newStr
}
