package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "errors"
    "sort"
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

var part2Scores = map[rune]int{'(': 1,
                               '[': 2,
                               '{': 3,
                               '<': 4}

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

func findIncomplete(line string) (stack, error) { // Could have been called 'hasWrong'
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
            return state, errors.New("Corrupted")
        }
    }
    return state, nil
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

var part2_test_input = part1_test_input
var part2_test_output = []string{
    `288957`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));

    scores := make([]int, 0)
    for _, input := range inputs {
        rest, err := findIncomplete(input)
        if err == nil {
            score := 0
            for i := len(rest) - 1; i >= 0; i-- {
                char := rest[i]
                score = score * 5
                score += part2Scores[char]
            }
            scores = append(scores, score)
        }
    }
    sort.Ints(scores)
    result := scores[len(scores)/2]
    return strconv.Itoa(result);
}
