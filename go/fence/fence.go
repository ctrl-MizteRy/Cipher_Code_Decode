package fence

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var dict = make(map[string]byte)

func GetMessage(dictionary map[string]byte, text string,
	key string, cod string) string {

	cod = strings.ToLower(cod)
	dict = dictionary

	if cod == "code" {
		return encodeCipher(text, key)
	} else if cod == "decode" {
		return decodeCipher(text, key)
	} else {
		return "We could not do the command: " + cod + ", only code or decode"
	}
}

func encodeCipher(input string, key string) string {
	steps, err := strconv.Atoi(key)
	if err != nil {
		panic(fmt.Sprintf("Could not convert %s into a number of step", key))
	}

	text := []byte(input)
	var msg = []byte{}
	text_len := len(text)
	for i := range steps {
		pos := i
		step := i
		direction := 1
		for pos < text_len {
			if step == i {
				msg = append(msg, text[pos])
			}

			if step == 0 {
				direction = 1
			} else if step == steps-1 {
				direction = -1
			}

			step += direction
			pos++
		}
	}

	return string(msg[:])
}

func decodeCipher(input string, key string) string {
	if key == "" {
		return decodeWithoutKey(input)
	} else {
		steps, err := strconv.Atoi(key)
		if err != nil {
			panic(fmt.Sprintf("Cannot convert %s into integer for step", key))
		}
		if steps <= 1 || steps >= len(input) {
			return input
		}
		text := []byte(input)
		text_len := len(text)
		var fences = make([][]string, steps)
		for i := range steps {
			fences[i] = make([]string, text_len)
			for j := 0; j < text_len; j++ {
				fences[i] = append(fences[i], "")
			}
		}
		text_step := 0
		bigStep := 2 * (steps - 2)
		for i := range len(fences) {
			step := i
			step1 := bigStep - (2*(i) - 1)
			step2 := bigStep - step1
			firstStep := true
			secondStep := false
			for step < text_len {
				fences[i][step] = string(text[text_step])
				text_step++
				if text_step == text_len {
					break
				}
				if firstStep {
					secondStep = true
					firstStep = false
					if step1 > 0 {
						step += step1 + 1
					} else {
						step += step2 + 1
					}
				} else if secondStep {
					firstStep = true
					secondStep = false
					if step2 > 0 {
						step += step2 + 1
					} else {
						step += step1 + 1
					}
				}
			}
		}

		var message = []string{}
		for col := range text_len {
			for row := range steps {
				if fences[row][col] != "" {
					message = append(message, fences[row][col])
				}
			}
		}
		return strings.Join(message, "")
	}
}

func decodeWithoutKey(input string) string {
	ans := make(map[int]string)
	for i := range len(input) {
		message := decodeCipher(input, strconv.Itoa(i))
		if ok, percent := checkDict(message); ok {
			if len(ans) == 0 {
				ans[0] = message
				ans[1] = strconv.FormatFloat(percent, 'f', -1, 64)
			} else {
				val, _ := strconv.ParseFloat(ans[1], 64)
				if percent >= val {
					ans[0] = message
					ans[1] = strconv.FormatFloat(percent, 'f', -1, 64)
				}
			}
		}
	}
	if len(ans) != 0 {
		return ans[0]
	} else {
		return "We could not decode this message :("
	}
}

func checkDict(input string) (bool, float64) {
	words := strings.Split(input, " ")
	count := 0
	for _, word := range words {
		regex := regexp.MustCompile("[!.?\":'()\\[\\]{}]+")
		newWord := regex.ReplaceAllString(word, "")
		_, ok := dict[strings.ToLower(newWord)]
		if ok {
			count++
		}
	}
	percentage := float64(count) / float64(len(words))
	if len(words) < 5 && percentage >= 0.5 {
		return true, percentage
	} else if percentage >= 0.7 {
		return true, percentage
	} else {
		return false, 0.0
	}
}
