package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "sort"
);

/**
  * Start - 18:40:06
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 07);
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

const separator string = ",";

var part1_test_input = []string{
    `16,1,2,0,4,2,7,1,2,14`,
};
var part1_test_output = []string{
    `37`,
};

func sum_with_offset(addends []int, offset int) int {
    sum := 0
    for _, addend := range addends {
        sum += utils.Abs(addend - offset)
    }
    return sum
}

func get_new_best(old_best int, old_offset int, length int) {

}

func binarySearch(needle int, haystack []int) bool {
    return true
}

func iterate(nums []int, old_best int) int {
    length := len(nums)
    if length <= 1 {
        return 0
    }
    return iterate(nums[len(nums)/2:], old_best)
}

func findBest(nums []int) int {
    sort.Ints(nums)
    cand := 4294967295
    for i := nums[0]; i < nums[len(nums)-1]; i++ {
        new_cand := sum_with_offset(nums, i)
        if new_cand > cand {
            break
        } else {
            cand = new_cand
        }
    }
    return cand
}

func find_best(nums []int) int {
    sort.Ints(nums)
    cand := 4294967295
    work_nums := make([]int, len(nums))
    highest := nums[len(nums)-1]
    lowest := nums[0]
    copy(work_nums, nums)
    for true {
        if highest == lowest {
            // We're done
            new_cand := sum_with_offset(nums, highest)
            return utils.Max(cand, new_cand)
        } else {
            lowest = (highest + lowest)/2
            continue
        }
        length := len(work_nums)
        if length == 1 {
            new_cand := sum_with_offset(nums, work_nums[0])
            cand = utils.Min(cand, new_cand)
            break
        } else {
        }
        work_nums = work_nums[length/2:]
    }
    return cand
}

func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator))
    var nums, _ = utils.StrToInt_array(inputs)
    result := findBest(nums)
    for _, num := range nums {
        if num != 0 {
        }
    }

    return strconv.Itoa(result);
}

var part2_test_input = part1_test_input
var part2_test_output = []string{
    `168`,
};
func part2(input string) string {
    // var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    // var nums, _ = utils.StrToInt_array(inputs);

    // ...

    return "";
    // return strconv.Itoa(result);
}
