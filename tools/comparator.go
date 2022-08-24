package tools

// Comparator return [-1, 0, 1]:
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
type Comparator func(a, b interface{}) int

func GetCmp(e interface{}) (cmp Comparator) {
	if e == nil {
		return nil
	}
	switch e.(type) {
	case int:
		return intCmp
	case int64:
		return int64Cmp
	case float32:
		return float32Cmp
	case float64:
		return float64Cmp
	}
	return nil
}

func intCmp(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int) > b.(int) {
		return 1
	} else if a.(int) < b.(int) {
		return -1
	}
	return 0
}

func int64Cmp(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int64) > b.(int64) {
		return 1
	} else if a.(int64) < b.(int64) {
		return -1
	}
	return 0
}

func float32Cmp(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(float32) > b.(float32) {
		return 1
	} else if a.(float32) < b.(float32) {
		return -1
	}
	return 0
}

func float64Cmp(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(float64) > b.(float64) {
		return 1
	} else if a.(float64) < b.(float64) {
		return -1
	}
	return 0
}
