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
        var part2_result = part2(part2_test_input[i])
        if (part2_result != part2_test_output[i]) {
            success = false;
            fmt.Printf("part2 failed with input %s: result %s != expected %s \n",
                    part2_test_input[i],
                    part2_result,
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

var part2_test_input = part1_test_input
var part2_test_output = []string{
    `230`,
};
func get_template(bitcounts []BitCount) []int {
    template := make([]int, len(bitcounts))
    for i, bitcount := range bitcounts {
        if bitcount.one >= bitcount.zero {
            template[i] = 1
        } else {
            template[i] = 0
        }
    }
    return template
}
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var length = len(inputs[0])
    counts := make([]BitCount, length)
    for i := range counts {
        counts[i] = BitCount{0,0}
    }
    fmt.Printf("Byte length %d\n", length)
    // Loop to get all bit counts
    //for _, input := range inputs {
    //    for pos, char := range input {
    //        if char == '0' {
    //            counts[pos].zero++
    //        } else {
    //            counts[pos].one++
    //        }
    //    }
    //}

    var valids_oxy = make([]bool, len(inputs))  // A number is valid for oxy if it has the most common bit in all pos
    var valids_cot = make([]bool, len(inputs))  // Inversely, here a number should have the least common bit
    // Init 'valid' flags
    for i := range inputs {
        valids_oxy[i] = true
        valids_cot[i] = true
    }

    //// Get oxy
    // Loop first over length of an input, then over each input (i.e. as described in text)
    for i := 0; i < length; i++ {
        for j, input := range inputs {
            if !valids_oxy[j] {
                continue // Has already been marked ineligible for some earlier bit, so we shouldn't count it
            }
            var char = input[i]
            if char == '0' {
                counts[i].zero++
            } else {
                counts[i].one++
            }
        }
        for j, input := range inputs { // Loop again to use 'counts' data to set 'valids'
            if !valids_oxy[j] {
                continue // We should be able to skip it here too
            }
            var char = ([]rune(input))[i]
            var rightchar rune
            if counts[i].one >= counts[i].zero {
                rightchar = '1'
            } else {
                rightchar = '0'
            }
            if char != rightchar {
                valids_oxy[j] = false
            }
        }
        // Loop over valids to see if we only have one left
        var valid_count = 0
        for _, valid := range valids_oxy {
            if valid {
                valid_count++
            }
        }
        if valid_count < 2 {
            break // we're done
        }
    }

    var oxy int = 0
    for i, input := range inputs { // Now loop to find the one that is valid
        if valids_oxy[i] {
            var res, _ = strconv.ParseInt(input, 2, 32)
            oxy = int(res)
            break
        }
    }
    fmt.Printf("oxy %d\n", oxy)

    // Now we copy-paste for co2
    // Reset counts
    for i := range counts {
        counts[i] = BitCount{0,0}
    }
    // Loop first over length of an input, then over each input (i.e. as described in text)
    for i := 0; i < length; i++ {
        for j, input := range inputs {
            if !valids_cot[j] {
                continue // Has already been marked ineligible for some earlier bit, so we shouldn't count it
            }
            var char = input[i]
            if char == '0' {
                counts[i].zero++
            } else {
                counts[i].one++
            }
        }
        for j, input := range inputs { // Loop again to use 'counts' data to set 'valids'
            if !valids_cot[j] {
                continue // We should be able to skip it here too
            }
            var char = ([]rune(input))[i]
            var rightchar rune
            if counts[i].one >= counts[i].zero {
                rightchar = '0'
            } else {
                rightchar = '1'
            }
            if char != rightchar {
                valids_cot[j] = false
            }
        }
        // Loop over valids to see if we only have one left
        var valid_count = 0
        for _, valid := range valids_cot {
            if valid {
                valid_count++
            }
        }
        if valid_count < 2 {
            break // we're done
        }
    }
    var cot int = 0
    for i, input := range inputs { // Now loop to find the one that is valid
        if valids_cot[i] {
            var res, _ = strconv.ParseInt(input, 2, 32)
            cot = int(res)
            println("found one")
            break
        }
    }
    fmt.Printf("cot %d\n", cot)

    var result = oxy * cot
    return strconv.Itoa(result);
}
