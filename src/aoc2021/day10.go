package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 00:39:01
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 10);
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
    p1 := part1(input);
    fmt.Printf("Part 1: %s\n\n", p1);

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

const separator string = "\n";

var part1_test_input = []string{
    `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`,
};
var part1_test_output = []string{
    `26397`,
};

type stack []rune

func (s stack) Push(v rune) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, rune) {
    // FIXME: What do we do if the stack is empty, though?

    l := len(s)
    return  s[:l-1], s[l-1]
}

var scores = map[rune]int{')': 3,
                          ']': 57,
                          '}': 1197,
                          '>': 25137}

func getWrongs(line string) []rune {
    result := make([]rune, 0)
    state := make(stack, 0)
    for _, char := range line {
        var expected int32
        if char == '(' || char == '[' || char == '{' || char == '<' {
            state = state.Push(char)
            continue
        } else if char == ')' {
            expected = '('
        } else if char == ']' {
            expected = '['
        } else if char == '}' {
            expected = '{'
        } else if char == '>' {
            expected = '<'
        }
        var last rune
        state, last = state.Pop()
        if last != expected {
            result = append(result, char)
        }
    }
    return result
}

func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    result := 0
    for _, line := range inputs {
        wrongs := getWrongs(line)
        for _, wrong := range wrongs {
            result += scores[wrong]
        }
    }

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
