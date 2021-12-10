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

func subset(first, second []Wire) bool {
    set := make(map[Wire]Wire)
    for _, value := range second {
        set[value] += 1
    }

    for _, value := range first {
        if count, found := set[value]; !found {
            return false
        } else if count < 1 {
            return false
        } else {
            set[value] = count - 1
        }
    }

    return true
}

func otherIndexes(index int) (int, int) {
    if index == 0 {
        return 1,2
    } else if index == 1 {
        return 0,2
    } else {
        return 0,1
    }
}

func union(a, b []Wire) []Wire{
      m := make(map[Wire]bool)

      for _, item := range a {
              m[item] = true
      }

      for _, item := range b {
              if _, ok := m[item]; !ok {
                      a = append(a, item)
              }
      }
      return a
}


func findSolution(wireGroups WireGroups) map[Wire]Segment {
    allWires := []Wire{a,b,c,d,e,f,g}
    sort.Sort(wireGroups)
    oneGroup := wireGroups[0]
    fmt.Printf("One is   %v\n", oneGroup)
    sevenGroup := wireGroups[1]
    fmt.Printf("Seven is %v\n", sevenGroup)
    solutionMap := make(map[Wire]Segment)
    solutionMapInv := make(map[Segment]Wire)
    for _, wire := range oneGroup {
        if !contains(sevenGroup, wire) {
            // This is the top element
            solutionMap[wire] = Top
            solutionMapInv[Top] = wire
        }
    }
    fourGroup := wireGroups[2]
    fmt.Printf("Four  is %v\n", fourGroup)

    twoThreeFiveGroups := wireGroups[3:6]
    var twoFiveGroups [][]Wire
    var threeGroup []Wire
    for i, cand := range twoThreeFiveGroups {
        other1, other2 := otherIndexes(i)
        group1 := twoThreeFiveGroups[other1]
        group2 := twoThreeFiveGroups[other2]
        if subset(cand, union(group1, group2)) {
            // This is the three
            threeGroup = wireGroups[i + 3]
            twoFiveGroups = append(twoFiveGroups, group1)
            twoFiveGroups = append(twoFiveGroups, group2)
            break
        }
    }
    fmt.Printf("Three is %v\n", threeGroup)

    for _, wire := range allWires {
        if contains(threeGroup, wire) && contains(fourGroup, wire) && !contains(oneGroup, wire) {
            solutionMap[wire] = Middle
            solutionMapInv[Middle] = wire
            break
        }
    }
    var zeroGroup []Wire
    var sixCandGroups [][]Wire
    for _, group := range wireGroups {
        middleWire := solutionMapInv[Middle]
        if len(group) == 6 {
            if !contains(group, middleWire) {
                zeroGroup = group
            } else {
                sixCandGroups = append(sixCandGroups, group)
            }
        }
    }
    fmt.Printf("Zero is  %v\n", zeroGroup)

    var sixGroup []Wire
    var nineGroup []Wire
    if subset(sixCandGroups[0], union(threeGroup, fourGroup)) {
        sixGroup, nineGroup = sixCandGroups[1], sixCandGroups[0]
    } else {
        sixGroup, nineGroup = sixCandGroups[0], sixCandGroups[1]
    }
    fmt.Printf("Six is   %v\n", sixGroup)
    fmt.Printf("Nine is  %v\n", nineGroup)
    var twoGroup []Wire
    var fiveGroup []Wire
    if subset(twoFiveGroups[0], union(threeGroup, sixGroup)) {
        twoGroup, fiveGroup = twoFiveGroups[0], twoFiveGroups[1]
    } else {
        twoGroup, fiveGroup = twoFiveGroups[1], twoFiveGroups[0]
    }
    fmt.Printf("Two is   %v\n", twoGroup)
    fmt.Printf("Five is  %v\n", fiveGroup)

    var eightGroup []Wire
    for _, group := range wireGroups {
        if len(group) == 7 {
            eightGroup = group
            break
        }
    }
    fmt.Printf("Eight is %v\n", eightGroup)

    for _, wire := range allWires {
        if !contains(zeroGroup, wire) {
            solutionMap[wire] = Middle
        } else if !contains(nineGroup, wire) {
            solutionMap[wire] = BottomLeft
        } else if !contains(sixGroup, wire) {
            solutionMap[wire] = TopRight
        } else if contains(fiveGroup, wire) && !contains(twoGroup, wire) {
            solutionMap[wire] = BottomRight
        } else if contains(sevenGroup, wire) && !contains(oneGroup, wire) {
            solutionMap[wire] = Top
        } else if contains(fourGroup, wire) && !contains(oneGroup, wire) {
            solutionMap[wire] = TopLeft
        } else if contains(twoGroup, wire) && !contains(oneGroup, wire) {
            solutionMap[wire] = Bottom
        } else {
            fmt.Printf("%v\n", solutionMap)
            fmt.Printf("current wire %v\n", wire)
            //panic("Something is wrong!")
            fmt.Println("Something is wrong!")
        }
    }
    return solutionMap
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

var mytestinput = "cf acf bcdf acdfg abcefg abdefg abcdfg acdeg abdfg abcdefg | cf cf cf cf"

func getSegmentFlags(solutionMap map[Wire]Segment, wires []Wire) map[Segment]bool {
    result := make(map[Segment]bool)
    for _, wire := range wires {
        result[solutionMap[wire]] = true
    }
    return result
}

func getNumber(solutionMap map[Wire]Segment, wires []Wire) int {
    if len(wires) == 2 {
        return 1
    } else if len(wires) == 3 {
        return 7
    } else if len(wires) == 4 {
        return 4
    } else if len(wires) == 7 {
        return 8
    } else if len(wires) == 5 {
        segments := getSegmentFlags(solutionMap, wires)
        if segments[TopRight] {
            if segments[BottomLeft] {
                return 2
            } else {
                return 3
            }
        } else {
            return 5
        }
    } else {
        segments := getSegmentFlags(solutionMap, wires)
        if segments[TopRight] {
            if segments[Middle] {
                return 9
            } else {
                return 0
            }
        } else {
            return 6
        }
    }
}

func part2(input string) string {
    notes := parse(input)

    total := 0
    for _, note := range notes {
        solutionMap := findSolution(note.patterns)
        length := len(note.output)
        for i, output := range note.output {
            number := getNumber(solutionMap, output)
            println(number)
            total += utils.IntPow(10, length - i - 1) * number
        }
    }
    result := total
    return strconv.Itoa(result);
}
