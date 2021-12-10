package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "sort"
);

/**
  * Start - 21:18:17
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 9);
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
    `2199943210
3987894921
9856789892
8767896789
9899965678`,
};
var part1_test_output = []string{
    `15`,
};

func getAtIndex(points [][]int, x, y int) int {
    // Return a huge number as dummy if out of bounds
    if y < 0 || x < 0 || y >= len(points) || x >= len(points[y]) {
        return 256
    } else {
        return points[y][x]
    }
}
func getLeft(points [][]int, x, y int) int {
    return getAtIndex(points, x - 1, y)
}
func getRight(points [][]int, x, y int) int {
    return getAtIndex(points, x + 1, y)
}
func getTop(points [][]int, x, y int) int {
    return getAtIndex(points, x,     y - 1)
}
func getBottom(points [][]int, x, y int) int {
    return getAtIndex(points, x    , y + 1)
}

func isLow(points [][]int, x, y int) bool {
    point := points[y][x]
    return point < getLeft(points, x, y) &&
           point < getRight(points, x, y) &&
           point < getTop(points, x, y) &&
           point < getBottom(points, x, y)
}

func findLowPoints(points [][]int) []int {
    result := make([]int, 0)
    for y, row := range points {
        for x, point := range row {
            if isLow(points, x, y) {
                result = append(result, point)
            }
        }
    }
    return result
}

func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    points := make([][]int, 0)
    for _, rawRow := range inputs {
        var row = utils.Trim_array(strings.Split(rawRow, ""));
        var nums, _ = utils.StrToInt_array(row);
        points = append(points, nums)
    }
    lowPoints := findLowPoints(points)
    result := utils.Sum(lowPoints) + len(lowPoints)
    return strconv.Itoa(result);
}

var part2_test_input = part1_test_input
var part2_test_output = []string{
    `1134`,
};

type Point struct{x,y int}

func findBasinsHelper(points [][]int, basinIds *[][]int, x, y, currentBasin int) {
    if x < 0 || y < 0 || y >= len(points) || x >= len(points[y]) {
        return
    }
    if points[y][x] == 9 {
        return
    }
    oldBasin := (*basinIds)[y][x]
    if oldBasin == -1 {
        (*basinIds)[y][x] = currentBasin
        findBasinsHelper(points, basinIds, x-1, y,   currentBasin)
        findBasinsHelper(points, basinIds, x,   y-1, currentBasin)
        findBasinsHelper(points, basinIds, x+1, y,   currentBasin)
        findBasinsHelper(points, basinIds, x,   y+1, currentBasin)
    }
    return
}

func findBasins(points [][]int) PointLists {
    basinIds := make([][]int, len(points)) // For every point, which basin is it part of?
    for i := range basinIds {
        basinIds[i] = make([]int, len(points[i]))
        for j := range basinIds[i] {
            basinIds[i][j] = -1
        }
    }

    currentBasin := 0
    foundOne := true
    for foundOne {
        foundOne = false
        for y, row := range points {
            for x, point := range row {
                oldBasin := basinIds[y][x]
                if oldBasin == -1 && point != 9 {
                    findBasinsHelper(points, &basinIds, x, y, currentBasin)
                    currentBasin++
                    foundOne = true
                    break
                }
            }
            if foundOne {
                break
            }
        }
    }

    basins := make([][]Point, currentBasin+1)
    for y := range basins {
        subList := make([]Point, 0)
        basins[y] = subList
    }
    for y, row := range basinIds {
        for x, point := range row {
            if point != -1 {
                basins[point] = append(basins[point], Point{x,y})
            }
        }
    }
    return basins
}

type PointLists [][]Point

func (pl PointLists) Len() int { return len(pl)}
func (pl PointLists) Swap(i, j int) { pl[i], pl[j] = pl[j], pl[i] }
func (pl PointLists) Less(i, j int) bool {
    return len(pl[i]) < len(pl[j])
}

func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    points := make([][]int, 0)
    for _, rawRow := range inputs {
        var row = utils.Trim_array(strings.Split(rawRow, ""));
        var nums, _ = utils.StrToInt_array(row);
        points = append(points, nums)
    }
    basins := findBasins(points)
    sort.Sort(basins)
    result := 1
    for _, basin := range basins[len(basins)-3:len(basins)] {
        result = result * len(basin)
    }
    return strconv.Itoa(result);
}
