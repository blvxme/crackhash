package seqgen

import "fmt"

type SequenceGenerator struct {
	alphabet         string
	maxWordLen       int
	currentWordIndex int64
	finalWordIndex   int64
}

func New(alphabet string, maxWordLen int, offset int64, nWords int64) (sg *SequenceGenerator) {
	sg = &SequenceGenerator{
		alphabet:         alphabet,
		maxWordLen:       maxWordLen,
		currentWordIndex: offset,
		finalWordIndex:   offset + nWords - 1,
	}

	return
}

func (sg *SequenceGenerator) Next() (nextWord string, err error) {
	if sg.currentWordIndex > sg.finalWordIndex {
		err = fmt.Errorf("word out of range")
		return
	}

	alphabetLen := len(sg.alphabet)
	currentWordLen, currentSum := 1, int64(0)
	for currentWordLen < sg.maxWordLen {
		currentSum += pow(alphabetLen, currentWordLen)
		if currentSum > sg.currentWordIndex {
			break
		}
		currentWordLen++
	}

	currentWordLocalIndex := sg.currentWordIndex - (currentSum - pow(alphabetLen, currentWordLen))

	for i := 0; i < currentWordLen; i++ {
		nextWord = string(sg.alphabet[currentWordLocalIndex%int64(alphabetLen)]) + nextWord
		currentWordLocalIndex /= int64(alphabetLen)
	}

	sg.currentWordIndex++

	return
}

func pow(a, b int) (result int64) {
	result = 1
	for i := 1; i <= b; i++ {
		result *= int64(a)
	}

	return
}
