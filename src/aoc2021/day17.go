package main;

import (
    "aoc/libs/utils"
    "fmt"
    //"strings"
    "strconv"
    //"time"
);

/**
  * Start - 21:53:19
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var success = true;
    for i := range part1_test_input {
        var part1_result = part1(part1_test_input[i])
        if part1_result != part1_test_output[i] {
            success = false;
            fmt.Printf("Part 1 failed with input %v: result %s != expected %s \n",
                    part1_test_input[i],
                    part1_result,
                    part1_test_output[i]);
            break;
        }
    }
    input := Target{124, 174, -123, -85}
    fmt.Printf("Part 1 minitest success: %t! \n", success);
        p1 := part1(input)
        fmt.Printf("Part 1: %s\n\n", p1);

    success = true;
    for i := range part2_test_input {
        var part2_result = part2(part2_test_input[i])
        if (part2_result != part2_test_output[i]) {
            success = false;
            fmt.Printf("Part 2 failed with input %v: result %s != expected %s \n",
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

type Target struct {xMin, xMax, yMin, yMax int}

var part1_test_input = []Target{
    Target{20,30,-10,-5},
};
var part1_test_output = []string{
    `45`,
};

type Probe struct {x, y, vx, vy int}


func (p *Probe) step() {

    p.x += p.vx
    p.y += p.vy

    // Drag
    if p.vx > 0 {
        p.vx--
    } else if p.vx < 0{
        p.vx++
    }
    // Gravity
    p.vy--

}

func (p *Probe) hit(target Target) bool {
    return (p.x >= target.xMin && p.x <= target.xMax &&
            p.y >= target.yMin && p.y <= target.yMax)
}

func (p *Probe) wayOff(target Target) bool {
    // Indicate that the probe will never hit the target

    // It's fallen below target
    if p.y < target.yMin {
        return true
    }

    // It has gone beyond the target in a positive direction
    if p.vx >= 0 && p.x > target.xMax {
        return true
    }
    // It has gone beyond the target in a negative direction
    if p.vx <= 0 && p.x < target.xMin {
        return true
    }
    return false
}

func maxVY(target Target) int {
    // We want to just exactly hit the target, so there should be only one step
    // between hitting the y=0 plane and hitting the target
    // then we should test this with the given vx and see whether we have time
    // to actually reverse our y speed
    if target.yMin > 0 {
        return target.yMin
    }
    return utils.Abs(target.yMin) - 1
}

func (p *Probe) stepX() {
    p.x += p.vx
    if p.vx > 0 {
        p.vx--
    } else if p.vx < 0{
        p.vx++
    }
}

func getVx(target Target) int {
    // Lowest vx to use and still hit the target x position (given some y velocity)
    // This gives longest possible time to hit a high y position
    vx := target.xMax / 3
    foundV := 0
    for true {
        p := Probe{0,0,vx,0}
        found := false
        //fmt.Printf("Trying with vx=%d\n", vx)
        //time.Sleep(1 * time.Second)
        for true {
            p.stepX()
            if p.vx == 0 && p.x < target.xMin {
                // Last one was good
                found = true
                break
            }
            if p.x <= target.xMax && p.x >= target.xMin {
                //println("hit")
                foundV = vx
                vx--
                break
            }
        }
        if found {
            break
        }
    }
    return foundV
}

func findVelocityForHighest(target Target) (int, int) {
    vx := getVx(target)
    vy := maxVY(target)
    fmt.Printf("vx=%d vy=%d\n", vx, vy)
    return vx, vy
}

func part1(target Target) string {
    //var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    // var nums, _ = utils.StrToInt_array(inputs);
    vx, vy := findVelocityForHighest(target)
    p := Probe{0,0,vx,vy}
    maxY := 0
    for true {
        p.step()
        maxY = utils.Max(maxY, p.y)
        if p.hit(target) {
            break
        }
    }
    result := maxY
    return strconv.Itoa(result);
}

var part2_test_input = []Target{
    Target{20,30,-10,-5},
};
var part2_test_output = []string{
    `112`,
};
func part2(target Target) string {
    // var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    // var nums, _ = utils.StrToInt_array(inputs);

    // ...

    return "";
    // return strconv.Itoa(result);
}
