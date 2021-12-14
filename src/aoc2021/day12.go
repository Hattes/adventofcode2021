package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 21:02:53
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 12);
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
    `start-A
start-b
A-c
A-b
b-d
A-end
b-end`,
};
var part1_test_output = []string{
    `10`,
};

type Cave struct {id string
                  big bool
                  adj []string}

func getCaves(inputs []string) map[string]Cave {
    caveMap := make(map[string]Cave)
    for _, input := range inputs {
        twoCaves := strings.Split(input, "-")
        cave1Id := twoCaves[0]
        cave2Id := twoCaves[1]
        cave1, ok := caveMap[cave1Id]
        if !ok {
            cave1 = Cave{cave1Id, utils.IsUpper(cave1Id), make([]string, 0)}
        }
        cave2, ok := caveMap[cave2Id]
        if !ok {
            cave2 = Cave{cave2Id, utils.IsUpper(cave2Id), make([]string, 0)}
        }
        cave1.adj = append(cave1.adj, cave2.id)
        cave2.adj = append(cave2.adj, cave1.id)
        //println(input)
        //fmt.Printf("%s is adj to %s\n", cave1Id, cave2Id)
        //fmt.Printf("%s is adj to %s\n", cave1.id, cave2.id)
        caveMap[cave1Id] = cave1
        caveMap[cave2Id] = cave2
    }
    return caveMap
}

func canVisit(cave Cave, visited map[string]bool, usedDouble bool) (bool, bool) {
    //return cave.big || !visited[cave.id], true
    if cave.id == "start" {
        return false, usedDouble
    }
    if cave.big {
        return true, usedDouble
    }
    if visited[cave.id] {
        return !usedDouble, true
    }
    return true, usedDouble
}

func getPathsHelper(caveMap map[string]Cave,
                    visited map[string]bool,
                    path []string,
                    id string,
                    usedDouble bool) [][]string {
    paths := make([][]string, 0)
    if id == "end" {
        paths = append(paths, path)
        return paths
    }
    visited[id] = true
    for _, caveId := range caveMap[id].adj {
        //fmt.Printf("cave id %s\n", caveId)
        canVisitNow, usedDoubleNow := canVisit(caveMap[caveId], visited, usedDouble)
        if !canVisitNow {
            continue
        } else {
            newVisited := make(map[string]bool)
            for k,v := range visited {
                newVisited[k] = v
            }
            path = append(path, caveId)
            newPath := make([]string, len(path))
            copy(newPath, path)
            paths = append(paths, getPathsHelper(caveMap, newVisited, newPath, caveId, usedDoubleNow)...)
        }
    }
    return paths
}

func getPaths(caveMap map[string]Cave) [][]string {
    visited := make(map[string]bool)
    path := []string{"start"}
    paths := getPathsHelper(caveMap, visited, path, "start", true)
    return paths
}

func getPaths2(caveMap map[string]Cave) [][]string {
    visited := make(map[string]bool)
    path := []string{"start"}
    paths := getPathsHelper(caveMap, visited, path, "start", false)
    return paths
}

func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    caves := getCaves(inputs)
    //fmt.Printf("%v\n", caves)
    paths := getPaths(caves)
    //fmt.Printf("%v\n", paths)
    result := len(paths)
    return strconv.Itoa(result);
}

var part2_test_input = part1_test_input
var part2_test_output = []string{
    `36`,
};

func printPaths(paths [][]string) {
    for _, path := range paths {
        println(strings.Join(path, ","))
    }
}

func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    caves := getCaves(inputs)
    paths := getPaths2(caves)
    //printPaths(paths)

    result := len(paths)
    return strconv.Itoa(result);
}
