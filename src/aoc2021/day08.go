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
  * Start - 18:51:41
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 8);
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
    if true {
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
    p2 := part2(input);
    fmt.Printf("Part 2: %s\n", p2);
}

const separator string = "\n";

var part1_test_input = []string{
    `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`,
};
var part1_test_output = []string{
    `26`,
};

type WireSolution [g+1][Bottom+1]bool

type Note struct { patterns, output [][]Wire }
type Wire int
type Segment int
const (
    a Wire = iota
    b
    c
    d
    e
    f
    g
)
const (
    Top Segment = iota
    TopLeft
    TopRight
    Middle
    BottomLeft
    BottomRight
    Bottom
)

func getSolutionMap(solution WireSolution) (map[Wire]Segment, error) {
    solutionMap := make(map[Wire]Segment)
    var err error
    for i, flags := range solution {
        var wire Wire = Wire(i)
        foundOne := false
        for j, flag := range flags {
            var segment = Segment(j)
            if flag {
                if foundOne {
                    err = errors.New("Solution not finalized!")
                    break
                } else {
                    foundOne = true
                    solutionMap[wire] = segment
                }
            }
        }
        if err != nil {
            break
        }
    }
    return solutionMap, err
}

func determineNumber(solutionMap map[Wire]Segment, wires []Wire) int {
    length := len(wires)
    if length == 2 {
        return 1
    } else if length == 4 {
        return 4
    } else if length == 3 {
        return 7
    } else if length == 7 {
        return 8
    } else {
        segmentFlags := make([]bool, Bottom+1)
        for _, wire := range wires {
            segmentFlags[solutionMap[wire]] = true
        }
        fmt.Printf("%v\n", segmentFlags)
        if segmentFlags[TopLeft] { // 0,5,6 or 9
            if !segmentFlags[Middle] {
                return 0
            } else if !segmentFlags[BottomLeft] {
                return 5
            } else if segmentFlags[TopRight] {
                return 9
            } else {
                return 6
            }
        } else { // 2 or 3
            if segmentFlags[BottomLeft] {
                return 2
            } else {
                return 3
            }
        }

    }
}

func determineValue(solution WireSolution, wireGroups [][]Wire) int {
    solutionMap, err := getSolutionMap(solution)
    if err != nil {
        panic("derp")
    }
    total := 0
    length := len(wireGroups)
    for i, wires := range wireGroups {
        number := determineNumber(solutionMap, wires)
        total += number * utils.IntPow(10, length - i)
    }
    return total
}

func parse(input string) []Note {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    notes := make([]Note, 0)
    for _, line := range lines {
        notes = append(notes, parseNote(line))
    }
    return notes
}

func parseNote(raw string) Note {
    note := utils.Trim_array(strings.Split(strings.Trim(raw, " | "), " | "));
    patterns := parseWireGroups(note[0])
    output := parseWireGroups(note[1])
    return Note{ patterns, output }
}

func parseWireGroups(raw string) [][]Wire {
    raws := utils.Trim_array(strings.Split(strings.Trim(raw, " "), " "));
    segmentGroups := make([][]Wire, len(raws))
    for i, rawWireGroup := range raws {
        segmentGroups[i] = make([]Wire, len(rawWireGroup))
        for j, char := range rawWireGroup {
            switch char {
            case 'a':
                segmentGroups[i][j] = a
            case 'b':
                segmentGroups[i][j] = b
            case 'c':
                segmentGroups[i][j] = c
            case 'd':
                segmentGroups[i][j] = d
            case 'e':
                segmentGroups[i][j] = e
            case 'f':
                segmentGroups[i][j] = f
            case 'g':
                segmentGroups[i][j] = g
            }
        }
    }
    return segmentGroups
}

func getSimplesN(segments [][]Wire) int {
    // Number of segments that are 'simple' to identify
    result := 0
    for _, segment := range segments {
        if isSimple(segment) {
            result++
        }
    }
    return result
}

func isSimple(segment []Wire) bool {
    length := len(segment)
    // Display numbers 1, 4, 7 and 8 use a unique number of display segments
    return length == 2 || length == 4 || length == 3 || length == 7
}

func part1(input string) string {
    notes := parse(input)
    result := 0
    for _, note := range notes {
        result += getSimplesN(note.output)
    }

    return strconv.Itoa(result);
}

var part2_test_input = part1_test_input
var part2_test_output = []string{
    `61229`,
};

func test() {
    solutionMap := map[Wire]Segment{
        d: Top,
        e: TopLeft,
        a: TopRight,
        f: Middle,
        g: BottomLeft,
        b: BottomRight,
        c: Bottom,
    }
    five := []Wire{c,d,f,e,b}
    three := []Wire{f,c,a,d,b}
    fmt.Printf("Five: %d\n", determineNumber(solutionMap, five))
    fmt.Printf("Three: %d\n", determineNumber(solutionMap, three))
}

type WireGroups [][]Wire

func (wg WireGroups) Len() int { return len(wg)}
func (wg WireGroups) Swap(i, j int) { wg[i], wg[j] = wg[j], wg[i] }
func (wg WireGroups) Less(i, j int) bool {
    return len(wg[i]) < len(wg[j])
}

func contains(s []Wire, e Wire) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func findSolution(wireGroups WireGroups) WireSolution {
    sort.Sort(wireGroups)
    oneGroup := wireGroups[0]
    sevenGroup := wireGroups[1]
    solutionMap := make(map[Wire]Segment)
    var solutionHelp WireSolution
    for _, wire := range oneGroup {
        if !contains(sevenGroup, wire) {
            // This is the top element
            solutionMap[wire] = Top
        }
    }
    var result WireSolution
    return result
}

func findSolution2(wireGroups [][]Wire) WireSolution {
    var solution WireSolution
    for _, wireGroup := range wireGroups {
        length := len(wireGroup)
        // For each possible number or set of numbers, go through and exclude the mappings that 
        // are impossible given the data
        if length == 2 { // 1
            excludeOther := false
            //var exclusion Segment
            for wire := range wireGroup {
                solution[wire][Top] = true // Exclude
                solution[wire][TopLeft] = true
                solution[wire][Middle] = true
                solution[wire][BottomLeft] = true
                solution[wire][Bottom] = true
                if excludeOther {
                    //solution[wire][segment] = true
                }
                // Check if this one has already been excluded from any of the 1-segments
                if solution[wire][TopRight] {
                    excludeOther = true
                    //exclusion = BottomRight
                } else if solution[wire][BottomRight] {
                }
            }
        } else if length == 3 { // 7
            for wire := range wireGroup {
                solution[wire][TopLeft] = true
                solution[wire][Middle] = true
                solution[wire][BottomLeft] = true
                solution[wire][Bottom] = true
            }
        } else if length == 4 { // 4
            for wire := range wireGroup {
                solution[wire][Top] = true // Exclude
                solution[wire][BottomLeft] = true
                solution[wire][Bottom] = true
            }
        } else if length == 5 { // 2,3 or 5
        } else if length == 6 { // 2,3 or 5
        } else if length == 7 { // 8
            // This gives us nothing
        }
    }
    return solution
}

func part2(input string) string {
    notes := parse(input)

    for _, note := range notes {
        findSolution2(note.patterns)
    }
    return "";
    // return strconv.Itoa(result);
}
