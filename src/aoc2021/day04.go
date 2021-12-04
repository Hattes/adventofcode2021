package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 12:57:15
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 04);
    // fmt.Printf("Input: %s \n", input);

    var success = true;
    for i := range part1_test_input {
        var part1_result = part1(part1_test_input[i])
        if part1_result != part1_test_output[i] {
            success = false;
            fmt.Printf("part1 failed with input %s: result %s != expected %s \n",
                    part1_test_input[i],
                    part1_result,
                    part1_test_output[i]);
            break;
        }
    }

    fmt.Printf("part1 minitest success: %t! \n", success);
    p1 := part1(input);
    fmt.Printf("part1: %s\n\n", p1);

    success = true;
    for i := range part2_test_input {
        var part2_result = part2(part2_test_input[i])
        if (part2_result != part2_test_output[i]) {
            success = false;
            fmt.Printf("part2 failed with input %s: result %s != expected %s \n",
                    part2_test_input[i],
                    part2_result,
                    part2_test_output[i]);
            break;
        }
    }
    fmt.Printf("part2 minitest success: %t! \n", success);
    p2 := part2(input);
    fmt.Printf("part2: %s\n", p2);
}

const separator string = "\n\n";

var part1_test_input = []string{
    `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`,
};
var part1_test_output = []string{
    `4512`,
};

type BingoSquare struct {
    number int
    marked bool
}

func parse_bingoboards(raws []string) [][][]BingoSquare {
    boards := make([][][]BingoSquare, len(raws))
    for i, raw := range raws {
        boards[i] = parse_bingoboard(raw)
    }
    return boards
}

func parse_bingoboard(raw string) [][]BingoSquare {
    var rows_raw = strings.Split(raw, "\n")
    var rows = make([][]BingoSquare, len(rows_raw))
    for i, row_raw := range rows_raw {
        var numbers_raw = utils.Trim_array(
            strings.Split(
                strings.Trim(
                    strings.Replace(row_raw, "  ", " ", -1), " "), " "));
        //fmt.Printf("Raw numbers %v\n", numbers_raw)
        var squares = make([]BingoSquare, len(numbers_raw))
        rows[i] = squares
        for j, number_raw := range numbers_raw {
            var number, _ = strconv.Atoi(number_raw)
            var square = BingoSquare{number, false}
            squares[j] = square
        }
    }
    return rows
}

func mark_and_check_bingo(board [][]BingoSquare, number int) bool {
    var hit = false
    var hit_x, hit_y int
    for i, row := range board {
        for j, square := range row {
            if number == square.number {
                hit = true
                hit_x = i
                hit_y = j
                board[i][j].marked = true
                //fmt.Printf("marked: %b", square.marked)
                break
            }
        }
    }
    var got_bingo = false
    if hit {
        var got_hor = true
        // Check horizontally and vertially from the hit square
        the_row := board[hit_x]
        for _, square := range the_row {
            if !square.marked {
                got_hor = false
                break
            }
        }
        var got_ver = true
        for _, row := range board {
            if !row[hit_y].marked {
                got_ver = false
                break
            }
        }
        got_bingo = got_hor || got_ver
    }
    return got_bingo
}

func get_unmarked_sum(board [][]BingoSquare) int {
    var sum = 0
    for _, row := range board {
        for _, square := range row {
            if !square.marked {
                sum += square.number
            }
        }
    }
    return sum
}

func print_board(board [][]BingoSquare) string {
    var builder strings.Builder
    for _, row := range board {
        for _, square := range row {
            var mark_x = " "
            if square.marked {
                mark_x = "x"
            }
            fmt.Fprintf(&builder, "%s%-3d", mark_x, square.number)
        }
        fmt.Fprintf(&builder, "\n")
    }
    return builder.String()
}

func parse_numbers(raw string) []int {
    var numbers_raw = utils.Trim_array(strings.Split(strings.Trim(raw, ","), ","));
    var numbers = make([]int, len(numbers_raw))
    for i, number_raw := range numbers_raw{
        var number, _ = strconv.Atoi(number_raw)
        numbers[i] = number
    }
    return numbers
}

func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    println(inputs[0])
    var numbers = parse_numbers(inputs[0])
    println(numbers[0])
    var boards = parse_bingoboards(inputs[1:])
    println(print_board(boards[0]))
    var got_winner = false
    var winning_board [][]BingoSquare
    var winning_score int
    for _, number := range numbers {
        //fmt.Printf("number is %d\n", number)
        for i := range boards {
            if mark_and_check_bingo(boards[i], number) {
                winning_board = boards[i]
                winning_score = get_unmarked_sum(boards[i]) * number
                got_winner = true
                break
            }
        }
        if got_winner {
            break
        }
    }
    println(print_board(winning_board))

    var result = winning_score
    return strconv.Itoa(result);
}

var part2_test_input = part1_test_input
var part2_test_output = []string{
    `1924`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    println(inputs[0])
    var numbers = parse_numbers(inputs[0])
    println(numbers[0])
    var boards = parse_bingoboards(inputs[1:])
    println(print_board(boards[0]))
    var winning_board [][]BingoSquare
    var winning_score int
    var boards_with_bingo = make([]bool, len(boards))
    for i := range boards_with_bingo {
        boards_with_bingo[i] = false
    }
    for _, number := range numbers {
        fmt.Printf("number is %d\n", number)
        for i := range boards {
            if boards_with_bingo[i] {
                continue
            }
            if mark_and_check_bingo(boards[i], number) {
                winning_board = boards[i]
                winning_score = get_unmarked_sum(boards[i]) * number
                boards_with_bingo[i] = true
            }
        }
    }
    println(print_board(winning_board))

    var result = winning_score
    return strconv.Itoa(result);
}
