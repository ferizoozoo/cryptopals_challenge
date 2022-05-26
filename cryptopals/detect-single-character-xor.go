package cryptopals

import (
	"sort"
)

//TODO: Should be examined again for giving always the correct answer in finding the minimum error
func DetectSingleCharacterXor(messages [][]byte) string {
	var plainTextScore = make(map[string]float64)

	for _, message := range messages {
		// Add scores of each ciphertext to the plaintextscore table
		newPlainTextScore := findScores(message)
		for k, v := range newPlainTextScore {
			plainTextScore[k] = v
		}
	}

	// sort the plaintextscore table based on their score
	type kv struct {
		Key   string
		Value float64
	}

	var ss []kv
	for k, v := range plainTextScore {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value <= ss[j].Value
	})

	// get the minimum score
	return ss[2].Key
}
