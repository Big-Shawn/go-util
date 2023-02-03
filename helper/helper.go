package helper

func IF[A, B any](condition bool, t A, f B) any {
	if condition {
		return t
	} else {
		return f
	}
}
