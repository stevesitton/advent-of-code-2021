package day14

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func ExtendedPolymerization(input string, steps int) int {
	inputFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	r, _ := regexp.Compile(" -> ")
	template := ""
	codes := make(map[string]string)
	counter := 1
	for scanner.Scan() {
		entry := scanner.Text()
		if counter == 1 {
			template = entry
		} else if counter > 2 {
			codeSplit := r.Split(entry, 3)
			codes[codeSplit[0]] = codeSplit[1]
		}
		counter++
	}

	for i := 0; i < steps; i++ {
		template = RunTemplate(template, codes)
	}

	low, high := GetLowHigh(template)
	//fmt.Println(high - low)
	return high - low
}

func RunTemplate(template string, codes map[string]string) string {
	output := ""
	for i := 0; i+1 < len(template); i++ {
		output += string(template[i]) + codes[template[i:i+2]]
		if i+2 == len(template) {
			output += string(template[len(template)-1])
		}
	}
	return output
}

func GetLowHigh(template string) (int, int) {
	low, high := 999, 0
	counts := make(map[string]int)
	for _, code := range strings.Split(template, "") {
		counts[code]++
	}
	for _, value := range counts {
		if value < low {
			low = value
		}
		if value > high {
			high = value
		}
	}
	return low, high
}

func ExtendedPolymerization_Part2(input string, steps int) int {
	inputFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	r, _ := regexp.Compile(" -> ")
	template := ""
	codes := make(map[string]string)
	counter := 1
	for scanner.Scan() {
		entry := scanner.Text()
		if counter == 1 {
			template = entry
		} else if counter > 2 {
			codeSplit := r.Split(entry, 3)
			codes[codeSplit[0]] = codeSplit[1]
		}
		counter++
	}

	// count of top level codes
	counts := make(map[string]int)
	for _, code := range template {
		counts[string(code)]++
	}
	pairCounts := make(map[string]int)

	for i := 0; i < len(template)-1; i++ {
		pairCounts[string(template[i])+string(template[i+1:i+2])]++
	}
	for i := 0; i < steps; i++ {
		pairCounts = Pairs(pairCounts, codes, counts)
	}

	low, high := 0, 0
	for _, value := range counts {
		if value < low || low == 0 {
			low = value
		}
		if value > high {
			high = value
		}
	}
	//fmt.Println(high - low)
	return high - low
}

func Pairs(pairCounts map[string]int, codes map[string]string, counts map[string]int) map[string]int {
	newPairCounts := make(map[string]int)

	for pair, count := range pairCounts {
		newCode := codes[pair]
		newPair1, newPair2 := pair[:1]+newCode, newCode+pair[1:]
		newPairCounts[newPair1] += count
		newPairCounts[newPair2] += count
		counts[newCode] += count
	}
	return newPairCounts
}

/*
FAILED ATTEMPTS - didn't perform

func RecurseTemplate(template string, steps int, codes map[string]string, counts map[string]int) {
	if steps == 0 {
		//fmt.Println("Steps 0: ", template)
	} else {
		for i := 0; i < len(template)-1; i++ {
			pair := template[i : i+2]
			newCode := codes[pair]
			counts[newCode]++
			RecurseTemplate(pair[:1]+newCode+pair[1:], steps-1, codes, counts)
		}
	}
}

func RecursePairs(subtemplate string, steps int, totalSteps int, codes map[string]string, counts map[string]int) {
	if steps == 0 {
		counts[subtemplate[:1]]++
	} else {
		template := string(subtemplate[:1]) + codes[subtemplate] + string(subtemplate[1:])
		RecursePairs(template[:2], steps-1, totalSteps, codes, counts)
		RecursePairs(template[1:], steps-1, totalSteps, codes, counts)
	}
}

func RunTemplate2(template string, codes map[string]string) string {
	sliceLength := 500000
	//output := make([]string, 0, sliceLength)
	sliceS := make([]string, 0, sliceLength)
	sliceM := make([]string, 0, sliceLength*10)
	sliceL := make([]string, 0, len(template)*3)
	templateLength := len(template)
	for i := 0; i+1 < templateLength; i++ {
		sliceS = append(sliceS, string(template[i]), codes[template[i:i+2]])

		if len(sliceS) >= sliceLength {
			sliceM = append(sliceM, strings.Join(sliceS, ""))
			sliceS = make([]string, 0, sliceLength)
			if len(sliceM) >= sliceLength*10 {
				sliceL = append(sliceL, strings.Join(sliceM, ""))
				sliceM = make([]string, 0, sliceLength*10)
			}
		}
		if i+2 == templateLength {
			sliceL = append(sliceL, strings.Join(sliceS, ""), strings.Join(sliceM, ""),
				string(template[templateLength-1]))
		}
	}
	return strings.Join(sliceL, "")
}
*/
