package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "math"
);

/**
  * Start - 21:49:48
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 19);
    // fmt.Printf("Input: %s \n", input);

    var success = true;
    for i := range part1_test_input {
        var part1_result = part1(part1_test_input[i])
        if part1_result != part1_test_output[i] {
            success = false;
            fmt.Printf("Part 1 failed with input %s: result %s != expected %s \n",
                    part1_test_input[i][:100],
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
    `--- scanner 0 ---
404,-588,-901
528,-643,409
-838,591,734
390,-675,-793
-537,-823,-458
-485,-357,347
-345,-311,381
-661,-816,-575
-876,649,763
-618,-824,-621
553,345,-567
474,580,667
-447,-329,318
-584,868,-557
544,-627,-890
564,392,-477
455,729,728
-892,524,684
-689,845,-530
423,-701,434
7,-33,-71
630,319,-379
443,580,662
-789,900,-551
459,-707,401

--- scanner 1 ---
686,422,578
605,423,415
515,917,-361
-336,658,858
95,138,22
-476,619,847
-340,-569,-846
567,-361,727
-460,603,-452
669,-402,600
729,430,532
-500,-761,534
-322,571,750
-466,-666,-811
-429,-592,574
-355,545,-477
703,-491,-529
-328,-685,520
413,935,-424
-391,539,-444
586,-435,557
-364,-763,-893
807,-499,-711
755,-354,-619
553,889,-390

--- scanner 2 ---
649,640,665
682,-795,504
-784,533,-524
-644,584,-595
-588,-843,648
-30,6,44
-674,560,763
500,723,-460
609,671,-379
-555,-800,653
-675,-892,-343
697,-426,-610
578,704,681
493,664,-388
-671,-858,530
-667,343,800
571,-461,-707
-138,-166,112
-889,563,-600
646,-828,498
640,759,510
-630,509,768
-681,-892,-333
673,-379,-804
-742,-814,-386
577,-820,562

--- scanner 3 ---
-589,542,597
605,-692,669
-500,565,-823
-660,373,557
-458,-679,-417
-488,449,543
-626,468,-788
338,-750,-386
528,-832,-391
562,-778,733
-938,-730,414
543,643,-506
-524,371,-870
407,773,750
-104,29,83
378,-903,-323
-778,-728,485
426,699,580
-438,-605,-362
-469,-447,-387
509,732,623
647,635,-688
-868,-804,481
614,-800,639
595,780,-596

--- scanner 4 ---
727,592,562
-293,-554,779
441,611,-461
-714,465,-776
-743,427,-804
-660,-479,-426
832,-632,460
927,-485,-438
408,393,-506
466,436,-512
110,16,151
-258,-428,682
-393,719,612
-211,-452,876
808,-476,-593
-575,615,604
-485,667,467
-680,325,-822
-627,-443,-432
872,-547,-609
833,512,582
807,604,487
839,-516,451
891,-625,532
-652,-548,-490
30,-46,-14`,
};
var part1_test_output = []string{
    `79`,
};
func parsePoint(raw string) Point {
    nums := strings.Split(raw, ",")
    x, _ := strconv.Atoi(nums[0])
    y, _ := strconv.Atoi(nums[1])
    z, _ := strconv.Atoi(nums[2])
    return Point{x,y,z}
}
func parseScannerReport(raw string) []Point {
    lines := utils.Trim_array(strings.Split(strings.Trim(raw, "\n"), "\n"))
    lines = lines[1:]  // Get rid of the first line which just numbers the scanner
    report := make([]Point, len(lines))
    for i, line := range lines {
        report[i] = parsePoint(line)
    }
    return report
}
type Point struct {x, y, z int}

func dist(p1, p2 Point) float64 {
    return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) +
                     math.Pow(float64(p1.y-p2.y), 2) +
                     math.Pow(float64(p1.z-p2.z), 2))
}

func nearestNeighborIndex(point Point, others []Point) int {
    shortest := math.MaxFloat64
    var neighborIx int
    for i, other := range others {
        dist := dist(point, other)
        if dist < shortest {
            shortest = dist
            neighborIx = i
        }
    }
    return neighborIx
}

func nearestNeighborIndexes(pivots, others []Point) []int {
    // For each point in first list, find the index of its
    // nearest neighbor in the second list
    is := make([]int, len(pivots))
    for i, pivot := range pivots {
        is[i] = nearestNeighborIndex(pivot, others)
    }
    return is
}

func part1(input string) string {
    var scannersRaw = utils.Trim_array(strings.Split(strings.Trim(input, "\n\n"), "\n\n"));
    scanners := make([][]Point, len(scannersRaw))
    for i, scannerRaw := range scannersRaw {
        scanners[i] = parseScannerReport(scannerRaw)
    }
    fmt.Printf("%v\n", scanners)
    fmt.Printf("%v\n", nearestNeighborIndexes(scanners[0], scanners[1]))

    return "";
    // return strconv.Itoa(result);
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
