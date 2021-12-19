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
    input := Target{124, 174, -123, -86}
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
    p2 := part2(input);
    fmt.Printf("Part 2: %s\n", p2);
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
    p.stepX()
    p.stepY()
}

func (p *Probe) stepX() {
    p.x += p.vx
    // Drag
    if p.vx > 0 {
        p.vx--
    } else if p.vx < 0{
        p.vx++
    }
}

func (p *Probe) stepY() {
    p.y += p.vy

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
    //fmt.Printf("vx=%d vy=%d\n", vx, vy)
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

func getVXs(target Target) []int {
    vxs := make([]int, 0)
    lowest := getVx(target)
    for v := lowest; v <= target.xMax; v++ {
        p := Probe{0,0,v,0}
        for true {
            p.stepX()
            if p.x >= target.xMin && p.x <= target.xMax {
                vxs = append(vxs, v)
                break
            } else if p.x > target.xMax || p.vx == 0 {
                break
            }
        }
    }
    return vxs
}

func getVYs(target Target) []int {
    // We cheat and only consider the case where the target min y is negative
    vys := make([]int, 0)
    minVY := target.yMin
    maxVY := utils.Abs(target.yMin)
    for vy := minVY; vy <= maxVY; vy++ {
        p := Probe{0,0,0,vy}
        for true {
            p.stepY()
            if p.y >= target.yMin && p.y <= target.yMax {
                vys = append(vys, vy)
                break
            } else if p.y < target.yMin && p.vy < 0 {
                break
            }
        }
    }
    return vys
}

func part2(target Target) string {
    // var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    // var nums, _ = utils.StrToInt_array(inputs);
    vxs := getVXs(target)
    //fmt.Printf("vxs: %v\n", vxs)
    vys := getVYs(target)
    //fmt.Printf("vys: %v\n", vys)
    //fmt.Printf("%d\n", len(vxs) * len(vys))
    hitCount := 0
    for _, vx := range vxs {
        for _, vy := range vys {
            p := Probe{0,0,vx,vy}
            for true {
                p.step()
                if p.hit(target) {
                    hitCount++
                    break
                } else if p.wayOff(target) {
                    break
                }
            }
        }
    }

    return strconv.Itoa(hitCount);
}
