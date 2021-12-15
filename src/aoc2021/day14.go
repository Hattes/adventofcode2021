package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 10:11:21
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 14);
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
    //p2 := part2(input);
    //fmt.Printf("Part 2: %s\n", p2);
}

const separator string = "\n";

var part1_test_input = []string{
    `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`,
};
var part1_test_output = []string{
    `1588`,
};

func getRules(rulesRaw []string) map[string]rune {
    rules := make(map[string]rune)
    for _, ruleRaw := range rulesRaw {
        rulePartsRaw := strings.Split(ruleRaw, " -> ")
        rules[rulePartsRaw[0]] = rune(rulePartsRaw[1][0])
    }
    return rules
}

func runInsertions(template string, rules map[string]rune, iterations int) string {
    current := template
    for i := 0; i < 10; i++ {
        var b strings.Builder
        for j := 0; j < len(current) - 1; j++ {
            b.WriteString(string(current[j]))
            pair := string(current[j]) + string(current[j+1])
            res, success := rules[pair]
            if success {
                b.WriteString(string(res))
            }
        }
        b.WriteString(string(current[len(current)-1]))
        current = b.String()
    }
    return current
}

func countElements(polymer string) (int, int) {
    counts := make(map[rune]int)
    for _, element := range polymer {
        counts[element]++
    }
    highest := 0
    lowest := utils.MaxInt
    for _, value := range counts {
        highest = utils.Max(highest, value)
        lowest = utils.Min(lowest, value)
    }
    return highest, lowest
}

func part1(input string) string {
    inputs := strings.Split(input, "\n\n")
    template := inputs[0]
    rulesRaw := utils.Trim_array(strings.Split(strings.Trim(inputs[1], separator), separator));

    rules := getRules(rulesRaw)

    sequence := runInsertions(template, rules, 10)
    //sequence := runInsertions(template, rules, 4)
    //fmt.Printf("%v\n", sequence)

    highest, lowest := countElements(sequence)
    result := highest - lowest
    return strconv.Itoa(result);
}

var part2_test_input = part1_test_input
var part2_test_output = []string{
    `2188189693529`,
};
func part2(input string) string {
    inputs := strings.Split(input, "\n\n")
    template := inputs[0]
    rulesRaw := utils.Trim_array(strings.Split(strings.Trim(inputs[1], separator), separator));

    rules := getRules(rulesRaw)

    sequence := runInsertions(template, rules, 40)

    highest, lowest := countElements(sequence)
    result := highest - lowest
    return strconv.Itoa(result);
}
