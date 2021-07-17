package common

// DupCount return same items number in a list
func DupCount(list []string) map[string]int {

	duplicateFrequency := make(map[string]int)
	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map
		_, exist := duplicateFrequency[item]

		if exist {
			duplicateFrequency[item]++ // increase counter by 1 if already in the map
		} else {
			duplicateFrequency[item] = 1 // else start counting from 1
		}
	}
	return duplicateFrequency
}
