package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 12:44:04
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 05);
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
    if input != "" {
    }
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
    `0,9 -> 5,9
    8,0 -> 0,8
    9,4 -> 3,4
    2,2 -> 2,1
    7,0 -> 7,4
    6,4 -> 2,0
    0,9 -> 2,9
    3,4 -> 1,4
    0,0 -> 8,8
    5,5 -> 8,2`,
};
var part1_test_output = []string{
    `5`,
};

type Line struct {x1,y1,x2,y2 int}

func get_point(raw string) (int, int) {
    xy_raw := strings.Split(raw, ",")
    x, _ := strconv.Atoi(xy_raw[0])
    y, _ := strconv.Atoi(xy_raw[1])
    return x, y
}

func get_lines(raws []string) []Line {
    lines := make([]Line, len(raws))
    for i, raw := range raws {
        points_raw := strings.Split(raw, " -> ")
        x1,y1 := get_point(points_raw[0])
        x2,y2 := get_point(points_raw[1])
        lines[i] = Line{x1,y1,x2,y2}
    }
    return lines
}

type Point struct {x,y int}

func get_points_for_line_ver_hos(line Line) []Point {
    var points = make([]Point, 0)
    //points = append(points, Point{line.x1, line.y1})
    //points = append(points, Point{line.x2, line.y2})
    // Handle two cases: either y1 == y2 or x1 == x2
    if line.x1 == line.x2 {
        y_min := utils.Min(line.y1, line.y2)
        y_max := utils.Max(line.y1, line.y2)
        for i := y_min; i <= y_max; i++ {
            points = append(points, Point{line.x1, i})
        }
    } else if line.y1 == line.y2 {
        x_min := utils.Min(line.x1, line.x2)
        x_max := utils.Max(line.x1, line.x2)
        for i := x_min; i <= x_max; i++ {
            points = append(points, Point{i, line.y1})
        }
    }
    return points
}

func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var lines = get_lines(inputs)
    var point_count_table = make(map[Point]int)
    for _, line := range lines {
        for _, point := range get_points_for_line_ver_hos(line) {
            point_count_table[point]++
        }
    }
    var overlap_count = get_overlaps(point_count_table)
    return strconv.Itoa(overlap_count);
}

var part2_test_input = part1_test_input
var part2_test_output = []string{
    `12`,
};

func get_overlaps(point_count_table map[Point]int) int {
    overlap_count := 0
    for _, value := range point_count_table {
        if value >= 2 {
            overlap_count++
        }
    }
    return overlap_count
}

func get_points_for_line_diag(line Line) []Point {
    var points = make([]Point, 0)
    // Handle three cases: either y1 == y2 or x1 == x2 or it's a straight diagonal line
    if line.x1 == line.x2 {
        y_min := utils.Min(line.y1, line.y2)
        y_max := utils.Max(line.y1, line.y2)
        for i := y_min; i <= y_max; i++ {
            points = append(points, Point{line.x1, i})
        }
    } else if line.y1 == line.y2 {
        x_min := utils.Min(line.x1, line.x2)
        x_max := utils.Max(line.x1, line.x2)
        for i := x_min; i <= x_max; i++ {
            points = append(points, Point{i, line.y1})
        }
    } else {
        // Assume it's a straight diagonal line
        if line.y1 < line.y2 && line.x1 < line.x2 {
            // e.g. (0,0) and (2,2)
            x := line.x1
            for y := line.y1; y <= line.y2; y++ {
                points = append(points, Point{x, y})
                x++
            }
        } else if line.y1 > line.y2 && line.x1 < line.x2 {
            // e.g. (2,0) and (0,2)
            x := line.x1
            for y := line.y1; y >= line.y2; y-- {
                points = append(points, Point{x, y})
                x++
            }
        } else if line.y1 < line.y2 && line.x1 > line.x2 {
            // e.g. (0,2) and (2,0)
            x := line.x1
            for y := line.y1; y <= line.y2; y++ {
                points = append(points, Point{x, y})
                x--
            }
        } else if line.y1 > line.y2 && line.x1 > line.x2 {
            // e.g. (2,2) and (0,0)
            x := line.x1
            for y := line.y1; y >= line.y2; y-- {
                points = append(points, Point{x, y})
                x--
            }
        } else {
            panic("unknown case!")
        }
    }
    return points
}


func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var lines = get_lines(inputs)
    //for _, line := range lines {
    //    fmt.Printf("Line %v has points %v\n", line, get_points_for_line_diag(line))
    //}
    var point_count_table = make(map[Point]int)
    for _, line := range lines {
        for _, point := range get_points_for_line_diag(line) {
            point_count_table[point]++
        }
    }
    var overlap_count = get_overlaps(point_count_table)

    return strconv.Itoa(overlap_count);
}
