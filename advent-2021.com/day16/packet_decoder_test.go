package day16

import (
	"testing"
)

func TestPacketDecoder(t *testing.T) {
	// if PacketDecoder("8A004A801A8002F478", false) != 16 {
	// 	t.Fatal()
	// }
	// if PacketDecoder("620080001611562C8802118E34", false) != 12 {
	// 	t.Fatal()
	// }
	// if PacketDecoder("C0015000016115A2E0802F182340", false) != 23 {
	// 	t.Fatal()
	// }
	// if PacketDecoder("A0016C880162017C3686B18A3D4780", false) != 31 {
	// 	t.Fatal()
	// }
	if PacketDecoder("input.txt", true) != 860 {
		t.Fatal()
	}
}

func TestPacketDecoder_Part2(t *testing.T) {

	// C200B40A82 finds the sum of 1 and 2, resulting in the value 3.
	// if PacketDecoder_Part2("C200B40A82", false) != 3 {
	// 	t.Fatal()
	// }
	// 04005AC33890 finds the product of 6 and 9, resulting in the value 54.
	// if PacketDecoder_Part2("04005AC33890", false) != 54 {
	// 	t.Fatal()
	// }
	// 880086C3E88112 finds the minimum of 7, 8, and 9, resulting in the value 7.
	// if PacketDecoder_Part2("880086C3E88112", false) != 7 {
	// 	t.Fatal()
	// }
	// CE00C43D881120 finds the maximum of 7, 8, and 9, resulting in the value 9.
	// if PacketDecoder_Part2("CE00C43D881120", false) != 9 {
	// 	t.Fatal()
	// }
	// D8005AC2A8F0 produces 1, because 5 is less than 15.
	// if PacketDecoder_Part2("D8005AC2A8F0", false) != 1 {
	// 	t.Fatal()
	// }
	// F600BC2D8F produces 0, because 5 is not greater than 15.
	// if PacketDecoder_Part2("F600BC2D8F", false) != 0 {
	// 	t.Fatal()
	// }
	// 9C005AC2F8F0 produces 0, because 5 is not equal to 15.
	// if PacketDecoder_Part2("9C005AC2F8F0", false) != 0 {
	// 	t.Fatal()
	// }
	// 9C0141080250320F1802104A08 produces 1, because 1 + 3 = 2 * 2.
	// if PacketDecoder_Part2("9C0141080250320F1802104A08", false) != 1 {
	// 	t.Fatal()
	// }
	if PacketDecoder_Part2("input.txt", true) != 860 {
		t.Fatal()
	}
}
