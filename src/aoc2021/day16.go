package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 20:05:24
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 16);
    // fmt.Printf("Input: %s \n", input);

    var success = true;
    for i := range part1_test_input {
        var part1_result = part1(part1_test_input[i])
        if part1_result != part1_test_output[i] {
            success = false;
            fmt.Printf("Part 1 failed with input %s: result %s != expected %s \n",
                    part1_test_input[i],
                    part1_result,
                    part1_test_output[i]);
            break;
        }
    }

    fmt.Printf("Part 1 minitest success: %t! \n", success);
    if true {
        p1 := part1(input);
        fmt.Printf("Part 1: %s\n\n", p1);
    }

    success = true;
    for i := range part2_test_input {
        var part2_result = part2(part2_test_input[i])
        if (part2_result != part2_test_output[i]) {
            success = false;
            fmt.Printf("Part 2 failed with input %s: result %s != expected %s \n",
                    part2_test_input[i],
                    part2_result,
                    part2_test_output[i]);
            break;
        }
    }
    fmt.Printf("Part 2 minitest success: %t! \n", success);
    p2 := part2(input);
    fmt.Printf("Part 2: %s\n", p2);
}

const separator string = "";

var part1_test_input = []string{
    `D2FE28`,
    `8A004A801A8002F478`,
    `620080001611562C8802118E34`,
    `C0015000016115A2E0802F182340`,
    `A0016C880162017C3686B18A3D4780`,
};
var part1_test_output = []string{
    `6`,
    `16`,
    `12`,
    `23`,
    `31`,
};

func nibblesToBits(nibbles []int) []int {
    bits := make([]int, len(nibbles)*4)
    for i, nibble := range nibbles {
        for k := 0; k < 4; k++ {
            j := i*4 + k
            bit := nibble & utils.IntPow(2, 3-k)
            if bit != 0 {
                bits[j] = 1
            } else {
                bits[j] = 0
            }
        }
    }
    return bits
}

func parseHexs(hexsRaw []string) []int {
    hexs := make([]int, len(hexsRaw))
    for i, hexRaw := range hexsRaw {
        parsed, _ := strconv.ParseInt(hexRaw, 16, 64)
        hexs[i] = int(parsed)
    }
    return hexs
}

func take(bits []int, length int) (int, []int) {
    chunk := bits[:length]
    return read(chunk), bits[length:]
}

func read(bits []int) int {
    result := 0
    for i := len(bits) - 1; i >= 0; i-- {
        if bits[len(bits)-i-1] == 1 {
            result += utils.IntPow(2, i)
        }
    }
    return result
}

func getLiteral(bits []int) (int, []int) {
    litBits := make([]int, 0)
    for true {
        var zeroIfLast int
        zeroIfLast, bits = take(bits, 1)
        litNibble := bits[:4]
        bits = bits[4:]
        litBits = append(litBits, litNibble...)
        if zeroIfLast == 0 {
            break
        }
    }
    return read(litBits), bits
}

func getValue(bits []int) (int, []int) {
    if len(bits) < 6 {
        return 0, []int{}
    }

    // Get header
    _, bits = take(bits, 3) // Get rid of version
    var typeID int
    typeID, bits = take(bits, 3)

    if typeID == 4 {
        var literal int
        literal, bits = getLiteral(bits)
        return literal, bits
    } else {
        var lengthTypeID int
        lengthTypeID, bits = take(bits, 1)
        values := make([]int, 0)
        if lengthTypeID == 0 {
            var length int
            length, bits = take(bits, 15)
            innerBits := bits[:length]
            bits = bits[length:]
            for true {
                var value int
                value, innerBits = getValue(innerBits)
                values = append(values, value)
                if len(innerBits) == 0 {
                    break
                }
            }
        } else {
            var number int
            number, bits = take(bits, 11)
            for i := 0; i < number; i++ {
                var value int
                value, bits = getValue(bits)
                values = append(values, value)
            }
        }
        switch typeID {
        case 0:
            return utils.Sum(values), bits
        case 1:
            return prod(values), bits
        case 2:
            return utils.ArrayMin(values), bits
        case 3:
            return utils.ArrayMax(values), bits
        case 5:
            if values[0] > values[1] {
                return 1, bits
            } else {
                return 0, bits
            }
        case 6:
            if values[0] < values[1] {
                return 1, bits
            } else {
                return 0, bits
            }
        case 7:
            if values[0] == values[1] {
                return 1, bits
            } else {
                return 0, bits
            }
        }
        return 0, bits
    }
}

func prod(nums []int) int {
    result := 1
    for _, num := range nums {
        result *= num
    }
    return result
}

func countVersions(bits []int) (int, []int) {
    if len(bits) < 6 {
        return 0, []int{}
    }

    // Get header
    var version int
    version, bits = take(bits, 3)
    //fmt.Printf("bits: %v\n", bits)
    //fmt.Printf("version: %d\n", version)
    var typeID int
    typeID, bits = take(bits, 3)
    //fmt.Printf("bits: %v\n", bits)
    //fmt.Printf("type: %d\n", typeID)

    versionsCount := version
    if typeID != 4 {
        var lengthTypeID int
        lengthTypeID, bits = take(bits, 1)
        if lengthTypeID == 0 {
            var length int
            length, bits = take(bits, 15)
            innerBits := bits[:length]
            bits = bits[length:]
            for true {
                var versionsCountInner int
                versionsCountInner, innerBits = countVersions(innerBits)
                versionsCount += versionsCountInner
                if len(innerBits) == 0 {
                    break
                }
            }
        } else {
            var number int
            number, bits = take(bits, 11)
            //fmt.Printf("number: %d\n", number)
            for i := 0; i < number; i++ {
                var versionsCountInner int
                versionsCountInner, bits = countVersions(bits)
                versionsCount += versionsCountInner
            }
        }
    } else {
        _, bits = getLiteral(bits)
    }
    return versionsCount, bits
}

func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    //fmt.Printf("inputs: %v\n", inputs)
    var nibbles = parseHexs(inputs);
    //fmt.Printf("nibbles: %v\n", nibbles)
    bits := nibblesToBits(nibbles)
    //fmt.Printf("bits: %v\n", bits)

    versionCount, _ := countVersions(bits)

    return strconv.Itoa(versionCount);
}

var part2_test_input = []string{
    `D2FE28`,
    `C200B40A82`,
    `04005AC33890`,
    `880086C3E88112`,
    `CE00C43D881120`,
    `D8005AC2A8F0`,
    `F600BC2D8F`,
    `9C005AC2F8F0`,
    `9C0141080250320F1802104A08`,
}
var part2_test_output = []string{
    `2021`,
    `3`,
    `54`,
    `7`,
    `9`,
    `1`,
    `0`,
    `0`,
    `1`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var nibbles = parseHexs(inputs);
    bits := nibblesToBits(nibbles)

    value, _ := getValue(bits)

    return strconv.Itoa(value);
}
