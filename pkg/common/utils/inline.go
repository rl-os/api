package utils

// IfThen evaluates a condition, if true returns the parameters otherwise nil
func IfThen(condition bool, a interface{}) interface{} {
	if condition {
		return a
	}
	return nil
}

// IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// DefaultIfNil checks if the value is nil, if true returns the default value otherwise the original
func DefaultIfNil(value interface{}, defaultValue interface{}) interface{} {
	if value != nil {
		return value
	}
	return defaultValue
}

// FirstNonNil returns the first non nil parameter
func FirstNonNil(values ...interface{}) interface{} {
	for _, value := range values {
		if value != nil {
			return value
		}
	}
	return nil
}

// ContainsString returns true if a string is present in a iteratee.
func ContainsString(s []string, v string) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

// ContainsStrings returns true if a strings is present in a iteratee.
// noinspection GoUnusedExportedFunction
func ContainsStrings(s []string, v []string) bool {
	count := 0
	for _, v1 := range s {
		if ContainsString(v, v1) {
			count++
		}
	}
	return count == len(v)
}
