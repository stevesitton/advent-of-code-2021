package day16

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PacketDecoder(input string, isFile bool) int {
	if isFile {
		data, _ := os.ReadFile(input)
		input = string(data)
	}

	type packet struct {
		version, typeLengthId, noSubPackets, typeId int
		remaining                                   string
	}

	var sb strings.Builder
	for _, hex := range input {
		sb.WriteString(hexToBin[string(hex)])
	}
	b := sb.String()

	process := func(b string) packet {
		p := packet{}
		version, _ := strconv.ParseInt(b[0:3], 2, 64)
		b = b[3:]
		p.version = int(version)
		typeId, _ := strconv.ParseInt(b[0:3], 2, 64)
		p.typeId = int(typeId)
		b = b[3:]
		if p.typeId == 4 {
			for {
				value := b[:5]
				b = b[5:]
				if value[0] == '0' {
					break
				}
			}
			p.remaining = b
		} else {
			p.typeLengthId, _ = strconv.Atoi(b[0:1])
			b = b[1:]
			packetSize := 15
			if p.typeLengthId == 1 {
				packetSize = 11
			}
			subPacketSize, _ := strconv.ParseInt(b[0:packetSize], 2, 64)
			p.noSubPackets = int(subPacketSize)
			p.remaining = b[packetSize:]
		}
		//fmt.Println("Version: ", p.version)
		return p
	}

	versionTotal := 0
	var queue Stack
	queue.Push(b)
	for {
		if queue.Len() == 0 {
			break
		}
		packetBin, _ := queue.Pop()
		if len(packetBin) < 10 {
			break
		}
		p := process(packetBin)

		versionTotal += p.version
		queue.Push(p.remaining)
	}

	fmt.Println(versionTotal)
	return versionTotal
}

/*
Didn't complete Part 2
*/
func PacketDecoder_Part2(input string, isFile bool) int {
	if isFile {
		data, _ := os.ReadFile(input)
		input = string(data)
	}

	type packet struct {
		version, typeLengthId, noSubPackets, typeId, literal int
		remaining                                            string
	}

	var sb strings.Builder
	for _, hex := range input {
		sb.WriteString(hexToBin[string(hex)])
	}
	b := sb.String()

	process := func(b string) packet {
		p := packet{}
		version, _ := strconv.ParseInt(b[0:3], 2, 64)
		b = b[3:]
		p.version = int(version)
		typeId, _ := strconv.ParseInt(b[0:3], 2, 64)
		p.typeId = int(typeId)
		b = b[3:]
		if p.typeId == 4 {
			valueStr := ""
			for {
				value := b[:5]
				valueStr += value[1:]
				b = b[5:]
				if value[0] == '0' {
					break
				}
			}
			literal, _ := strconv.ParseInt(valueStr, 2, 64)
			fmt.Println("Value: ", valueStr, literal)
			p.literal = int(literal)
			p.remaining = b
		} else {
			p.typeLengthId, _ = strconv.Atoi(b[0:1])
			b = b[1:]
			packetSize := 15
			if p.typeLengthId == 1 {
				packetSize = 11
			}
			subPacketSize, _ := strconv.ParseInt(b[0:packetSize], 2, 64)
			p.noSubPackets = int(subPacketSize)
			p.remaining = b[packetSize:]
		}
		//fmt.Println("Version: ", p.version)
		return p
	}

	versionTotal := 0
	packets := make([]packet, 0)
	counter := 0
	var queue Stack
	queue.Push(b)
	for {
		if queue.Len() == 0 {
			break
		}
		packetBin, _ := queue.Pop()
		if len(packetBin) < 10 {
			break
		}
		p := process(packetBin)
		fmt.Println("Type: ", p.typeId)
		packets = append(packets, p)
		counter++

		versionTotal += p.version
		queue.Push(p.remaining)
	}

	fmt.Println("Packets: ", len(packets))
	output := 0
	typeId := packets[0].typeId
	literals := make([]int, 0)
	for _, p := range packets {
		if p.typeId == 4 {
			literals = append(literals, p.literal)
		}
	}
	switch typeId {
	case 0:
		fmt.Println("Case 0")
		for _, val := range literals {
			output += val
		}
	case 1:
		fmt.Println("Case 1")
		output = literals[0]
		for i := 1; i < len(literals); i++ {
			output = output * literals[i]
		}
	case 2:
		fmt.Println("Case 2")
		output = 777777
		for _, val := range literals {
			if val < output {
				output = val
			}
		}
	case 3:
		fmt.Println("Case 3")
		for _, val := range literals {
			if val > output {
				output = val
			}
		}
	case 5:
		fmt.Println("Case 5")
		if literals[0] > literals[1] {
			output = 1
		}
	case 6:
		fmt.Println("Case 6")
		if literals[0] < literals[1] {
			output = 1
		}
	case 7:
		fmt.Println("Case 7")
		if literals[0] == literals[1] {
			output = 1
		}
	}

	fmt.Println(output)
	return output
}

/*
620080001611562C8802118E34 represents an operator packet (version 3) which contains two sub-packets;
each sub-packet is an operator packet that contains two literal values. This packet has a version
sum of 12.
01100010000000001000000000000000000101100001000101010110001011001000100000000010000100011000111000110100
VVVTTTILLLLLLLLLLL
3  0   2
                  00000000000000000101100001000101010110001011001000100000000010000100011000111000110100
                  VVVTTTILLLLLLLLLLLLLLLAAAAAAAAAAAAAAAAAAAAAA
				  0  0   22             VVVTTTNNNNNVVVTTTNNNNN
                                        0  4       5  4
                                                              001000100000000010000100011000111000110100
															  VVVTTTILLLLLLLLLLLVVVTTTNNNNNVVVTTTNNNNN
															  1  0   2          0  4       3  4

										---
01010010001001000000000
BBBBBBBBBBBBBBBB
VVVTTT
2  4
1000100100
 0001 0100
 20

8A004A801A8002F478
100010100000000001001010100000000001101010000000000000101111010001111000
VVVTTTILLLLLLLLLLLAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
4  2   1
                  001010100000000001101010000000000000101111010001111000
				  VVVTTTILLLLLLLLLLL
				  1  2   1
				                    101010000000000000101111010001111000
									VVVTTTILLLLLLLLLLLLLLLAAAAAAAAAAA
									5  2   11             VVVTTT
									                      6  4
*/
var hexToBin = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

// Stack implementation
type Stack []string

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Push(point string) {
	*s = append(*s, point)
}

func (s *Stack) Pop() (string, bool) {
	if s.Len() == 0 {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
