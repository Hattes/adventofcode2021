package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
     "strconv"
);

/**
  * Start - 18:12:54
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 06);
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

const separator string = ",";

var part1_test_input = []string{
    `3,4,3,1,2`,
};
var part1_test_output = []string{
    `5934`,
};

func iterate(fish_counts []int, iterations int) ([]int, int, int) {
    //for i := 0; i < iterations; i++ {
    //    for _, fish_count := range fish_counts {
    //        fish_count.countdown--
    //        if fish_count.countdown == 0 {

    //        }
    //    }
    //}
    //return fish_counts
    newborn := 0
    halfgrown := 0
    for i := 0; i < iterations; i++ {
        offset := i % 7
        temp_newborn := newborn
        newborn = fish_counts[offset]
        fish_counts[offset] += halfgrown
        halfgrown = temp_newborn
    }
    return fish_counts, newborn, halfgrown
}

func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var nums, _ = utils.StrToInt_array(inputs);

    return part1_helper(nums, 80)
}

type FishCount struct{countdown,size int}

func part1_helper(nums []int, iterations int) string {
    fish_counts := make([]int, 9)  // One counter for each start-offset (which input can be regarded as)
    //fish_counts := make([]FishCount, 9)
    // Initialization
    //for i, num := range nums {
    //    fish_counts[i].countdown = num
    //    fish_counts[i].countup++
    //}
    for _, num := range nums {
        fish_counts[num]++
    }

    // Iterations
    fish_counts, newborn, halfgrown := iterate(fish_counts, iterations)

    result := utils.Sum(fish_counts) + newborn + halfgrown
    //result := 0
    //for _, fish_count := range fish_counts {
    //    result += fish_count.countup
    //}
    return strconv.Itoa(result);
}

var part2_test_input = part1_test_input
var part2_test_output = []string{
    `26984457539`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var nums, _ = utils.StrToInt_array(inputs);
    return part1_helper(nums, 256)
}
