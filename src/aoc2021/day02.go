package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 19:14:40
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 02);
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

type Submarine struct{ X, Y, aim int }

func pre_parse(raw string) (string, string) {
    var result = strings.Split(raw, " ")
    return result[0], result[1]
}

var part1_test_input = []string{
    `forward 5
    down 5
    forward 8
    up 3
    down 8
    forward 2`,
};
var part1_test_output = []string{
    `150`,
};
func part1(input string) string {
    var sub = Submarine{0,0,0}
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    for _, input := range inputs {
        var action, number_str = pre_parse(input)
        var number, _ = strconv.Atoi(number_str)
        switch action {
        case "forward":
            sub.X += number
        case "down":
            sub.Y += number
        case "up":
            sub.Y -= number
        }
    }

    var result = sub.X * sub.Y
    return strconv.Itoa(result);
}

var part2_test_input = part1_test_input 
var part2_test_output = []string{
    `900`,
};
func part2(input string) string {
    var sub = Submarine{0,0,0}
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    for _, input := range inputs {
        var action, number_str = pre_parse(input)
        var number, _ = strconv.Atoi(number_str)
        switch action {
        case "forward":
            sub.X += number
            sub.Y += number * sub.aim
        case "down":
            sub.aim += number
        case "up":
            sub.aim -= number
        }
    }

    var result = sub.X * sub.Y
    return strconv.Itoa(result);
}
