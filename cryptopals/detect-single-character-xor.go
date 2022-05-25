package cryptopals

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func DetectSingleCharacterXor() {
	var plainTextScore = make(map[string]float64)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()

		if message == "" {
			break
		}

		// Add scores of each ciphertext to the plaintextscore table
		newPlainTextScore := findScores([]byte(message))
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
	fmt.Printf("%s", ss[0].Key)
}
