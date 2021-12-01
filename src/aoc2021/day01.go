package main;

import (
    "fmt"
    "strings"
    "aoc/libs/utils"
    "strconv"
);

/**
  * Start - 19:53:13
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 01);
    //fmt.Printf("Input: %s \n", input);

    var success = true;
    for i := range part1_test_input {
        if (part1(part1_test_input[i]) != part1_test_output[i]) {
            success = false;
            fmt.Printf("part1 failed with input %s: result %s != expected %s \n"
                    ,part1_test_input[i],
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
    `199
    200
    208
    210
    200
    207
    240
    269
    260
    263`,
};
var part1_test_output = []string{
    `7`,
};

func parse_inputs(inputs []string) ([]int, error) {
    return utils.StrToInt_array(inputs);
}

func part1(input string) string {
    var raw_inputs = strings.Split(strings.Trim(input, separator), separator);
    var inputs, err = parse_inputs(raw_inputs)
    if err != nil {
        println(err.Error())
    }
    fmt.Println(len(inputs))
    var increases = 0
    //fmt.Printf("Inputs %s\n", inputs)
    for i := 0; i < len(inputs) - 1; i++ {
        old := inputs[i]
        new := inputs[i+1]
        if new > old {
            //fmt.Printf("%d is greater than %d", new, old)
            increases++
        }
    }
    // ...
    println(increases)

    return strconv.Itoa(increases);
}

var part2_test_input = part1_test_input
var part2_test_output = []string{
    `5`,
};
func part2(input string) string {
    var raw_inputs = strings.Split(strings.Trim(input, separator), separator);
    var inputs, err = parse_inputs(raw_inputs)
    if err != nil {
        println(err.Error())
    }
    fmt.Println(len(inputs))
    var increases = 0
    //fmt.Printf("Inputs %s\n", inputs)
    for i := 0; i < len(inputs) - 3; i++ {
        old := inputs[i] + inputs[i+1] + inputs[i+2]
        new := inputs[i+1] + inputs[i+2] + inputs[i+3]
        if new > old {
            //fmt.Printf("%d is greater than %d\n", new, old)
            increases++
        }
    }
    // ...
    println(increases)

    return strconv.Itoa(increases);
}
