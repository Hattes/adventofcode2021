package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 16:25:10
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 11);
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
    if false {
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
    //p2 := part2(input);
    //fmt.Printf("Part 2: %s\n", p2);
}

const separator string = "\n";

var part1_test_input = []string{
    `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`,
};
var part1_test_output = []string{
    `1656`,
};

func getOctopodes(rows []string) [][]Octopus {
    result := make([][]Octopus, 0)
    for _, row := range rows {
        numbers := strings.Split(row, "")
        octopodes := make([]Octopus, 0)
        for _, numberRaw := range numbers {
            number, _ := strconv.Atoi(numberRaw)
            octopodes = append(octopodes, Octopus{number, false})
        }
        result = append(result, octopodes)
    }
    return result
}

type Octopus struct {energy int
                     flashed bool}

func iterate(octopodes [][]Octopus, iterations int) int {
    // Run sim, return number of flashes
    flashes := 0
    for i := 0; i < iterations; i++ {
        for y := range octopodes {
            for x := range octopodes[y] {
                octopus := &(octopodes[y][x])
                if octopus.energy == 9 {
                    octopus.flashed = true
                    flashes++
                    octopus.energy = 0
                } else {
                    octopus.energy++
                }
            }
        }
        // Reset all flash flags
        for y := range octopodes {
            for x := range octopodes[y] {
                octopus := &(octopodes[y][x])
                octopus.flashed = false
            }
        }
    }
    return flashes
}
func part1(input string) string {
    var rows = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    octopodes := getOctopodes(rows)
    result := iterate(octopodes, 100)

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
