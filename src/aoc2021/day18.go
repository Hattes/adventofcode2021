package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 13:46:36
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 18);
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
    `[1,1]
[2,2]
[3,3]
[4,4]`,
    `[1,1]
[2,2]
[3,3]
[4,4]
[5,5]`,
    `[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
[6,6]`,
    `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`,
    `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`,
};
var part1_test_output = []string{
    `445`,
    `791`,
    `1137`,
    `3488`,
    `4140`,
};

func (sn *SnailN) isRegular() bool {
    if sn.left == nil {
        if sn.right != nil {
            panic("bad data")
        } else {
            return true
        }
    }
    return false
}

type SnailN struct {n int
                    left, right *SnailN}

func parseSnailNumbers(snailNRaws []string) []SnailN {
    snailNs := make([]SnailN, len(snailNRaws))
    for i := range snailNRaws {
        snailNs[i], _ = parseSnailNumber([]rune(snailNRaws[i]))
    }
    return snailNs
}

func take(raw []rune) (rune, []rune) {
    item := raw[0]
    return item, raw[1:]
}

func parseDigit(r rune) (int, bool) {
    dig := int(r - '0')
    if dig >= 0 && dig <= 9 {
        return dig, true
    } else {
        return -1, false
    }
}

func parseNumber(chars []rune) (int, []rune) {
    digits := make([]int, 0)
    var newChars []rune
    for true {
        var next rune
        next, newChars = take(chars)
        dig, ok := parseDigit(next)
        if ok {
            digits = append(digits, dig)
            chars = newChars
        } else {
            break
        }
    }
    total := 0
    for i := 0; i < len(digits); i++ {
        total += digits[i] * utils.IntPow(10, len(digits) - i - 1)
    }
    return total, chars
}

func (sn SnailN) String() string {
    if sn.isRegular() {
        return fmt.Sprintf("%d", sn.n)
    }
    return fmt.Sprintf("[%v,%v]", sn.left, sn.right)
}

func parseSnailNumber(chars []rune) (SnailN, []rune) {
    //time.Sleep(time.Second/2)
    origChars := chars
    first, chars := take(chars)
    if first != '[' {
        // Should be a 'regular' number
        n, chars := parseNumber(origChars)
        return SnailN{n, nil, nil}, chars
    }

    left, chars := parseSnailNumber(chars)
    comma, chars := take(chars)
    if comma != ',' {
        panic(fmt.Sprintf("Parse error: expected ',' but got '%s'", string(comma)))
    }
    right, chars := parseSnailNumber(chars)
    last, chars := take(chars)

    if last != ']' {
        panic(fmt.Sprintf("Parse error: expected ']' but got '%s'", string(last)))
    }

    return SnailN{0, &left, &right}, chars
}

func (sn *SnailN) magnitude() int {
    if sn.isRegular() {
        return sn.n
    }
    return 3 * sn.left.magnitude() + 2 * sn.right.magnitude()
}

func split(n int) (int, int) {
    return n/2, (n+1)/2
}

func (sn *SnailN) checkForExplosions() bool {
    exploded, _, _ := sn.checkForExplosionsHelper(0)
    return exploded
}

func (sn *SnailN) addToLeft(n int) {
    if sn.isRegular() {
        sn.n += n
    } else {
        sn.left.addToLeft(n)
    }
}

func (sn *SnailN) addToRight(n int) {
    if sn.isRegular() {
        sn.n += n
    } else {
        sn.right.addToRight(n)
    }
}

func (sn *SnailN) checkForExplosionsHelper(depth int) (bool, int, int) {
    if sn.isRegular() {
        return false, 0, 0
    }
    if depth == 3 {
        if !sn.left.isRegular() {
            // This is where the explosion happens
            rShrap := sn.left.right.n
            sn.right.addToLeft(rShrap)
            lShrap := sn.left.left.n
            sn.left = &(SnailN{0, nil, nil})
            return true, lShrap, 0
        } else if !sn.right.isRegular() {
            lShrap := sn.right.left.n
            sn.left.addToRight(lShrap)
            rShrap := sn.right.right.n
            sn.right = &(SnailN{0, nil, nil})
            return true, 0, rShrap
        } else {
            return false, 0, 0
        }
    }
    leftExpl, lShrap, rShrap := sn.left.checkForExplosionsHelper(depth+1)
    if leftExpl {
        if rShrap != 0 {
            sn.right.addToLeft(rShrap)
        }
        return true, lShrap, 0  // left shrapnel must be added at a higher level
    }
    rightExpl, lShrap, rShrap := sn.right.checkForExplosionsHelper(depth+1)
    if rightExpl {
        if lShrap != 0 {
            sn.left.addToRight(lShrap)
        }
        return true, 0, rShrap
    }
    return false, 0, 0
}

func halveRoundBothWays(n int) (int,int) {
    return (n / 2), ((n+1) / 2)
}

func (sn *SnailN) split() {
    n1, n2 := halveRoundBothWays(sn.n)
    sn.n = 0
    sn.left = &(SnailN{n1,nil,nil})
    sn.right = &(SnailN{n2,nil,nil})
}

func (sn *SnailN) checkForSplits() bool {
    if sn.isRegular() {
        if sn.n >= 10 {
            sn.split()
            return true
        }
        return false
    }
    return sn.left.checkForSplits() || sn.right.checkForSplits()
}

func (sn *SnailN) reduce() {
    count := 0
    for true {
        exploded := sn.checkForExplosions()
        if exploded {
            count++
            continue
        }
        splitted := sn.checkForSplits()
        if !splitted {
            break
        }
    }
}

func add(sn, other SnailN) SnailN {
    return SnailN{0, &sn, &other}
}

func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    sns := parseSnailNumbers(inputs)
    testStr := "[[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]],[7,[5,[[3,8],[1,4]]]]]"
    testSn, _ := parseSnailNumber([]rune(testStr))
    testSn.reduce()
    sn := sns[0]
    for i := 1; i < len(sns); i++ {
        sn = add(sn, sns[i])
        sn.reduce()
    }
    total := sn.magnitude()

    return strconv.Itoa(total);
}

func getPairs(sns []SnailN) [][]SnailN {
    pairs := make([][]SnailN, 0)
    for i := range sns {
        for j := i+1; j < len(sns); j++ {
            pairs = append(pairs, []SnailN{sns[i], sns[j]})
        }
    }
    return pairs
}

func (sn *SnailN) copy() *SnailN {
    if sn.isRegular() {
        return &SnailN{sn.n, nil, nil}
    } else {
        return &SnailN{0, sn.left.copy(), sn.right.copy()}
    }
}

var part2_test_input = []string{
    `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`,
};
var part2_test_output = []string{
    `3993`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    sns := parseSnailNumbers(inputs)
    pairs := getPairs(sns)
    maxMagn := 0

    for _, pair := range pairs {
        for _, is := range [][]int{{0,1},{1,0}} {
            first := pair[is[0]].copy()
            second := pair[is[1]].copy()
            sum := add(*first, *second)

            sum.reduce()
            magnitude := sum.magnitude()
            maxMagn = utils.Max(maxMagn, magnitude)
        }
    }
    return strconv.Itoa(maxMagn);
}
