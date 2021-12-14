package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
     "strconv"
);

/**
  * Start - 19:52:40
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 13);
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
    fmt.Printf("Part 2 minitest success: %t! \n", success);
    p2 := part2(input);
    fmt.Printf("Part 2: %s\n", p2);
}

const separator string = "\n";

var part1_test_input = []string{
    `6,10
    0,14
    9,10
    0,3
    10,4
    4,11
    6,0
    6,12
    4,1
    0,13
    10,12
    3,4
    3,0
    8,4
    1,10
    2,14
    8,10
    9,0

    fold along y=7
    fold along x=5`,
};
var part1_test_output = []string{
    `17`,
};

func parseDotCoords(dotRaw string) (int,int) {
    split := strings.Split(dotRaw, ",")
    x, _ := strconv.Atoi(split[0])
    y, _ := strconv.Atoi(split[1])
    return x, y
}

func getDotArray(dotsRaw []string) [][]bool {
    xs := make([]int, 0)
    ys := make([]int, 0)
    for _, dotRaw := range dotsRaw {
        x, y := parseDotCoords(dotRaw)
        xs = append(xs, x)
        ys = append(ys, y)
    }
    xMax := utils.ArrayMax(xs)
    yMax := utils.ArrayMax(ys)
    dotArray := make([][]bool, yMax+1)
    for i := range dotArray {
        dotArray[i] = make([]bool, xMax+1)
    }
    for i := range xs {
        y := ys[i]
        x := xs[i]
        dotArray[y][x] = true
    }
    return dotArray
}

type FoldSpec struct {orientationX bool
                      i int}

func getFoldSpec(foldRaw string) FoldSpec {
    split := strings.Split(foldRaw, "=")
    orientationX := split[0] == "fold along x"
    index, _ := strconv.Atoi(split[1])
    return FoldSpec{orientationX, index}
}

func getFoldSpecs(foldSpecsRaw []string) []FoldSpec{
    foldSpecs := make([]FoldSpec, len(foldSpecsRaw))
    for i, foldSpecRaw := range foldSpecsRaw {
        foldSpecs[i] = getFoldSpec(foldSpecRaw)
    }
    return foldSpecs
}

func printDots(dotArray [][]bool) {
    for y := range dotArray {
        for x := range dotArray[y] {
            if dotArray[y][x] {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }
        println("")
    }
}

func getDottedNumber(dotArray [][]bool) int {
    result := 0
    for y := range dotArray {
        for x := range dotArray[y] {
            if dotArray[y][x] {
                result++
            }
        }
    }
    return result
}

func transpose(slice [][]bool) [][]bool {
    xl := len(slice[0])
    yl := len(slice)
    result := make([][]bool, xl)
    for i := range result {
        result[i] = make([]bool, yl)
    }
    for i := 0; i < xl; i++ {
        for j := 0; j < yl; j++ {
            result[i][j] = slice[j][i]
        }
    }
    return result
}

func fold(dotArray [][]bool, foldSpec FoldSpec) [][]bool{
    // Transpose so that we can always fold along x
    if foldSpec.orientationX {
        dotArray = transpose(dotArray)
    }

    foldedLengthY := utils.Max(foldSpec.i, len(dotArray) - foldSpec.i - 1)
    //fmt.Printf("foldedLengthY %d\n", foldedLengthY)
    upperPartOffset := utils.Max(0, foldedLengthY - foldSpec.i)
    //fmt.Printf("upperPartOffset %d\n", upperPartOffset)
    //fmt.Printf("foldSpec.i %d\n", foldSpec.i)
    //fmt.Printf("foldSpec %v\n", foldSpec)
    yMax := len(dotArray) - 1

    folded := make([][]bool, foldedLengthY)
    foldedLengthX := len(dotArray[0])
    for i := range folded {
        folded[i] = make([]bool, foldedLengthX)
    }
    for y := range folded {
        for x := range folded[y] {
            //println(upperPartOffset)
            foldedOverDot := dotArray[yMax-y][x]
            folded[y][x] = dotArray[y + upperPartOffset][x] || foldedOverDot
        }
    }

    // Transpose back
    if foldSpec.orientationX {
        folded = transpose(folded)
    }
    return folded
}

func part1(input string) string {
    inputs := strings.Split(input, "\n\n")
    dotsRaw := utils.Trim_array(strings.Split(strings.Trim(inputs[0], separator), separator));
    foldSpecsRaw := utils.Trim_array(strings.Split(strings.Trim(inputs[1], separator), separator));
    dotArray := getDotArray(dotsRaw)
    foldSpecs := getFoldSpecs(foldSpecsRaw)
    //fmt.Printf("%v\n", foldSpecs)
    //printDots(dotArray)
    dotArray = fold(dotArray, foldSpecs[0])
    //printDots(dotArray)
    result := getDottedNumber(dotArray)

    return strconv.Itoa(result);
}

var part2_test_input = []string{
    ``,
};
var part2_test_output = []string{
    ``,
};
func part2(input string) string {
    inputs := strings.Split(input, "\n\n")
    dotsRaw := utils.Trim_array(strings.Split(strings.Trim(inputs[0], separator), separator));
    foldSpecsRaw := utils.Trim_array(strings.Split(strings.Trim(inputs[1], separator), separator));
    dotArray := getDotArray(dotsRaw)
    foldSpecs := getFoldSpecs(foldSpecsRaw)
    for _, foldSpec := range foldSpecs {
        dotArray = fold(dotArray, foldSpec)
    }
    printDots(dotArray)

    // There is no result to return here, unless I want to write
    // an algorithm to read letters from a dot matrix
    return "";
}
