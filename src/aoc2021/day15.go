package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "github.com/beefsack/go-astar"
);

/**
  * Start - 12:13:06
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 15);
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
    `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`,
};
var part1_test_output = []string{
    `40`,
};

func getRiskLevels(rows []string) [][]int {
    riskLevels := make([][]int, len(rows))
    for y, row := range rows {
        riskLevels[y] = make([]int, len(row))
        for x, riskRaw := range strings.Split(row, "") {
            risk, _ := strconv.Atoi(riskRaw)
            riskLevels[y][x] = risk
        }
    }
    return riskLevels
}

func getRiskMap(riskLevels [][]int) RiskMap {
    riskMap := make(RiskMap)
    for y, row := range riskLevels {
        riskMap[y] = make(map[int]*Cell)
        for x, risk := range row {
            riskMap[y][x] = &Cell{risk, x, y, riskMap}
        }
    }
    return riskMap
}

func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    riskLevels := getRiskLevels(inputs)
    riskMap := getRiskMap(riskLevels)

    t1 := riskMap[0][0]
    t2 := riskMap[len(inputs)-1][len(inputs[0])-1]
    _, distance, found := astar.Path(t1, t2)
    if !found {
        println("Could not find path")
        return ""
    }

    return strconv.Itoa(int(distance));
}

var part2_test_input = part1_test_input
var part2_test_output = []string{
    `315`,
};

func replicateAndIncrease(riskLevels [][]int, times int) [][]int {
    extended := make([][]int, 0)
    for i := 0; i < times+1; i++ {
        yOffset := i*len(riskLevels)
        for j := 0; j < times+1; j++ {
            xOffset := j*len(riskLevels)
            //fmt.Printf("offsets are %d and %d\n", xOffset, yOffset)
            for y, row := range riskLevels {
                yIndex := yOffset + y
                if yIndex > len(extended) - 1 {
                    extended = append(extended, make([]int, len(riskLevels[y])*(times+1)))
                }
                for x, risk := range row {
                    elevated := ((risk + yOffset + xOffset - 1) % 9) + 1
                    extended[yOffset+y][xOffset+x] = elevated
                }
            }
        }
    }
    return extended
}

func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    riskLevels := getRiskLevels(inputs)
    riskLevels = replicateAndIncrease(riskLevels, 4)
    //fmt.Printf("%v\n", riskLevels)
    riskMap := getRiskMap(riskLevels)

    t1 := riskMap[0][0]
    t2 := riskMap[len(riskLevels)-1][len(riskLevels[0])-1]
    _, distance, found := astar.Path(t1, t2)
    if !found {
        println("Could not find path")
        return ""
    }

    return strconv.Itoa(int(distance));
}

type RiskMap map[int]map[int]*Cell

func (c *Cell) String() string {
    return fmt.Sprintf("cell risk:%d x:%d y:%d\n", c.risk, c.x, c.y)
}
func (rm RiskMap) Cell(x, y int) *Cell {
    if rm[y] == nil {
        return nil
    }
    return rm[y][x]
}

func (c *Cell) PathNeighbors() []astar.Pather {
    neighbors := []astar.Pather{}
    for _, offset := range [][]int{{-1,0},
                                   {1, 0},
                                   {0, -1},
                                   {0, 1},} {
        neighbor := c.riskMap.Cell(c.x + offset[0], c.y + offset[1])
        if neighbor != nil {
            neighbors = append(neighbors, neighbor)
        }
    }
    return neighbors
}

func (c *Cell) PathNeighborCost(to astar.Pather) float64 {
    toC := to.(*Cell)
    return float64(toC.risk)
}

func (c *Cell) PathEstimatedCost(to astar.Pather) float64 {
    toC := to.(*Cell)
    return float64(utils.Abs(c.x - toC.x) + utils.Abs(c.y - toC.y))
}

type Cell struct {
    risk int
    x,y int
    riskMap RiskMap
}
