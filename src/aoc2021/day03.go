package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
     "strconv"
);

/**
  * Start - 20:13:47
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 03);
    // fmt.Printf("Input: %s \n", input);

    var success = true;
    for i := range part1_test_input {
        if (part1(part1_test_input[i]) != part1_test_output[i]) {
            success = false;
            fmt.Printf("part1 failed with input %s: result %s != expected %s \n",
                    part1_test_input[i],
                    part1(part1_test_input[i]),
                    part1_test_output[i]);
            break;
        }
    }

    fmt.Printf("part1 minitest success: %t! \n", success);
    p1 := part1(input);
    fmt.Printf("part1: %s\n\n", p1);

    success = true;
    for i := range part2_test_input {
        if (part2(part2_test_input[i]) != part2_test_output[i]) {
            success = false;
            fmt.Printf("part2 failed with input %s: result %s != expected %s \n",
                    part2_test_input[i],
                    part2(part2_test_input[i]),
                    part2_test_output[i]);
            break;
        }
    }
    fmt.Printf("part2 minitest success: %t! \n", success);
    p2 := part2(input);
    fmt.Printf("part2: %s\n", p2);
}

const separator string = "\n";

var part1_test_input = []string{
    `00100
    11110
    10110
    10111
    10101
    01111
    00111
    11100
    10000
    11001
    00010
    01010`,
};
var part1_test_output = []string{
    `198`,
};

func gamma_epsilon(bitcounts []BitCount) (int, int) {
    var gamma = 0
    var epsilon = 0
    var length = len(bitcounts) - 1
    for i, bitcount := range bitcounts {
        var n = length - i
        if bitcount.one > bitcount.zero {
            gamma += utils.IntPow(2, n)
        } else {
            epsilon += utils.IntPow(2, n)
        }
    }
    return gamma, epsilon
}

type BitCount struct {zero, one int}

func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var length = len(inputs[0])
    counts := make([]BitCount, length)
    for i := range counts {
        counts[i] = BitCount{0,0}
    }
    fmt.Printf("Byte length %d\n", length)
    for _, input := range inputs {
        for pos, char := range input {
            if char == '0' {
                counts[pos].zero++
            } else {
                counts[pos].one++
            }
        }
    }
    var gamma, epsilon = gamma_epsilon(counts)

    fmt.Printf("gamma   %d\n", gamma)
    fmt.Printf("epsilon %d\n", epsilon)
    var result = gamma * epsilon
    return strconv.Itoa(result);
}

var part2_test_input = []string{
    ``,
};
var part2_test_output = []string{
    ``,
};
func part2(input string) string {
    // var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    // var nums, _ = utils.StrToInt_array(inputs);

    // ...

    return "";
    // return strconv.Itoa(result);
}
