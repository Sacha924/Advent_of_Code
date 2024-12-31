package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction [2]int

var up = Direction{-1, 0}
var down = Direction{1, 0}
var left = Direction{0, -1}
var right = Direction{0, 1}

func main() {
	moves := "^>>^>v^<^^v<^^>^<v>><vvv>><<<>^^^>v<^>v><v^v>vv<^<<v^^v<^>v^^^<v^>>^<<>>>^<>v^v<<^v^<v^<><v><<v<vv^<v<^v^v<<<^v>^>>>^<^^v>v><v><<^<^^^v>v<<<<<><v><>>>^vv<v<<vvv>^>^<vv><v>><^<^<v><vv<<^<<<v^^vv<>^>>^>v<>>^^>>v>v<^v<v^<^^^>>>vvv^^v^vv><<^>^v<^v<v^v^^v^^><><<>vv^<<vvv<v><^>^>^<^<>>^v>vv^^v^>vv<<><>>v>><^^^<^^>>>v>^<^<>^^^<<><^<>^v><v^^^v^^<>vv>^vvvv^v<<^<^^>>v<><>vv^>^<>^v<^v<<>>^^^>>^^^<v<>^>>><vv>v^v^<vv<vvvvv^<<>v^v>^^<^>><^v^^<^v>>^<v^>^^>^^>>>^<<>^v<<^>^vv<^^^v^><vv^<^v^>^v^><<<<<^^^><<<<v^>^>><vv>^^>v>^<>^>^<><^>>v<<^^v^vvv^>v>v><><^^^>v^>>^<>^>v><^>>^>v>vv^v<^v<<^^>v^v<v>vv<^^>>^vv^v><^v>^^^v>>v<vvvv<>>>>v<>^v>>^^^v><^^^>v>^><>^>v<><v<>>^<^<^^vv><^<<^<v^>vv<>>vv>>v>>><<v<v>^<vv<<>v^v<<^^>><^^^<v<v^v^<vv^^v>><<<<<^^<^^<v<v<v<^>v^v^<v^<>^>vv>><^<<>>>^v^<<><^<^>^<<vv<^>v>v<><v<^<>vv>vv<v^^^<^><<v<<^v<>^^^><^><>^><v<^^v<v^^><>>^v^^^vv>^<><^v<><^^<>><>vvv^v>>>>v<<<>v<>vv>v<^vv<^><^v>>^^vv>^v^v^v^^>>^>^v^<^>^^v^<<><^^<^^^<^v^<<<<v^^^<<<>>v^^>v<vvv<v^v>^<>v^^v>^<<<><>v<v^^<<><v<<<><^<vvv\nvv>^^^v<<^v^^^>v^^<<^<>>^vv<<>v<^>>^^^<^vv<v^<^><^<>^<^v^^vv^vv^><v^v<<v^<^^<^vv>vv<vvv>^v<^vv>^>vvvvv^v^^>^><<v<>>><^<^v^><<><>>^^>><>^<><v^^<v>^>v<vvvv><<v^^v><>vv^^<^v<>vvvv>^^^^<^<^<^^^v>v<v^><>v^v><vvvvv^v^<vv><v>^>^>v<<^<vv^>v^>v><>^vv<v>v>>><^v>^<^<><v>v<<v>>v>>^<<>v^<<v<^<^v^^<>^v><vv>^^<<^<<<^>>v<v>^<>v^^<<v^<<v^v^>vv<>v<><<<>>v<>v<^><v^^vv<v<^^<^v>v^><><>v^v><><<^><<v<^><>^^^v><^^<>>v>>>>v<>^^<^>^>^<^^<v^<<^vv>v^^<v^^<v^><v<^<>v>v>^^<v>>v<v>v>v^v><^>>^>><<^<>^^^^><<>>v><v<>^<<v>v^v<>^^^<<^^vv><v^>v>><<^>^^>v>^<vv^>>^><^<>>v^^^^<>^v^v><^v>^v^<v>><vv>^^vv^<>>v^<<<^v>v>v^v^^<<<<vv<>vv^^><<>^>vvvv>v>v>>>^v<^<>v<^<v>>^>>^v<v<^>><v>v<<v>>^^v>><>vv>^>>v<<>>vv^vvv<^^v><><^<v^>vv^<v><vv^>^v<>><vv<<>v<^<v><^^^<v<><>^^v<v<^>>^><v^^^>>^v^<>vvv>^v^<><><^v<<>><>>^vvv><>^<<>>vv>vv<^>^vv>v^><^<<v><><v>v^^<>vv^><>^v>^><^^<>>^^^<<^^v^<^>>v^<<^v>>^>^^<^<>^<v>^>vv<v<>^><v>v<<^>^<><<>>>>>>>v>^v>><v^<<^<>^v<v<v^><^^vv<><v^>^<>vv<^v^vv^^^<v><^<v<^>>v>>^^>^<<^<>v^vv^vv><>>^<v><<<^>^><^v><<v<>v><>v<<\n<<<v^<^^v<^<<<><^^><^<^<^v>^<<^v^>^<>^^^^vvvv^>v>^^<v<<^>><^<^^>v^^<vvvvv>>>>^<>^>^>>^vv<^>><^><v>^>^>v>v>>><^vv>>^<<^^><^vv<^><<>^<vv<^^vv^^>>vv><v^<^v<v^vv<v^v>>^^<>>^^v><<vv<^><<<^^><<>>vv>>>>^<<>>><^>^<^<^<<<<vv<^v>>v>vv>^v<vv<>vvv<v<^>>v^>vv>^^>v>><v><v<>v^v<<><>^v>v<>vv^v<<>^>^v>^><v^^><>vv<^^>vv<^<>v^^>^v^>^v><v>>vvv^<^<^>v<^v^>v^^^^^v^^>>>><<^>^^v^><><v>v^<v<vv>^^<^>v<v<>>^<<<<^>v><>>vv>><>^^<<v>v>^^><vv<v>^v>><vv><^<><^^>^v<><vv^>><<<v<v<^<^<^^<v>v>vv>v<<^^v^<v>>^<v<>>v>^v<^v^vv<v<>>^><>v<v>v<<<v>^vvv><>^^vv>v>><>>^>^v^>^>>v>><>>>>>>>v<v<><v>>v^>vv<<^><^<v><^v<^v>^<v^>^^>>v^<^^>v<^v<>^^<v^>>>v^^>v><^<<^v^>^^v><>><^<^><><^v>^<<v><^v^><<^^<^<^<>><^^^><<vv^^v^vv><>vv<v><<<v<<>vvv^^><<^v<><>^><>v<^vv>v^v<<^^^^v>vv<^<v^<vv>^^^<<<<^<v><>><^v^><^>v>v<<<>^vvv<<<^<<v>v<v>v>v<<>v><vv<<^<>>v><<<<><^<><<<>^^<>^<vv>><>v^><><v<^vvv><v<>vvv><^<^v<v^v<v><^vvv>v^<<v^^>><<v^^<<^<>>><>>v<^><<><^<>>^^v<^v><>>>^^>><<>^<^v^<>><^v^><>>>v^v^v>><>>><^><^v<<^>^<<<^>>>v<><>v<<<^>^v><<^v>^<v>v^<>v^>>^^^^\n^<<^<<^<^><v>vv<<>^>^<vv>>^>^v<v>>>>>>><^<^>^<><><<^>^vv<^^v>^<^v^>^<^v>^v^v>v>vv><v^^^<>^^>v<>vv^^v<^<v^<^^^><<v<^<v^^v<^v<<v^<<>v<>^v^><v>^><v^v>vv<<^^vvv^^<<><v^>vv^<^<^>>vvv^vvvv<<vv>v^v>^><<v<>v^^>^<vv>^>v^<<vvvv^^<v<>v>v^<>>vv^><<<<vv<v^<>v<<^<^<<<^v^v<^<>^^v<<v^><><^><^><><>v^vv<<<>>vv>v<^<<^^>^<^>>><>v^vv^>vv^<>vvv^^v>v<^^^^^v>>><v><<^v>><v<^^^>>v<>v<vv^^vv><^><v<^v<^>>v<^v>>v^>^>>>v<<v>>><^>>^>^>>^><>>^>^><^<^>v^><^v><<<<<<^>^<>^<<vv>^<>vv<<vv<v<<<v^^<^>vv<^<<<^<>vv^^^<<<>>v^^<<<^^<>>>^<<v>^v^^>v^<vv<v^>vv^vv<^^^vvv>>v<vv>v^>v<vv<vv<<v<>v<<^<v><v<<^<<^v^^<v^v>>>vv<><<>>^><><>><<v>^>^v>>^<^>><<>v^v>^<>v>><v^^^<><>^<<v>>>^v>>v<v<<<v>^^^>><vvv<><^>v>^<vvv<^>^v<>^>v^<<<<<^v<^v><><v>vv^<vv<>^><>^><>^v<>>^<^>v^<v><<<^>^^^>^^v^>^<>><^v<vv<>^vv<^>v^>>>><<^<v<>^v<<>>^v<>v<^<<<><>>v^vv<v>^><v^^v^<<^^v<^vvvv<^>v><<>^<^><^vvv>^vv<>^>vv>v^>^^<>vv>>>>^vvv<v^<>>>>v^<<>v>v^^^^<>^>v<^^<<<<^<<<>v^>v^<<vv>>vvv<v^^^<<<<>>v<>v^>^>vv<v>^^<<v^^<^vv^<vv>v>><v<^<^^<^>^>>>><v<>>^>>><vv>vv^v<>^^^<<>^v^<\n<><^vvv^><<<vv<><<^<^v<>v^<vvv>^vv^^<^<^>>^><v<>>^>v<v<<<v^^v^v>>^>><^^vvv>v^<<v><<>><>><^<<v^v><v^><^>>^^<<<^>v^v>>vv^^^v^<^v><<v>>>>v>^^v^<>vvv^>><^>^^^>v<vv^<^^<>v<<<<v>>^^v>^v<<><^<<^v^vv^v^^v<>^v>^>v^v^vv^^>>v>v<<<v^>^>^v^<^^<<^<>^>>>>>>v>>v<^^<>>v^<>v<<^>^>^>^><^^v^<vv>^^^v>v^v<^>v<v^>^<><v^^<v^v^^>><v>v>^>v<^^>>^v^>v^>^^<v<>>v<^^v>v^>^<<v><^v>^v^<v><^^<^>^<<vv^v<v<^>>^<v<<vvv>^>^<>><^>^^>>^^v^v^>><><^><^<v^v^>v>^v<^<<<<v<^^vv><>v^>v<^^<><>^>>>^^^<^^v<^vv<vv<v><<^<^^v><v>><v<><>vv<v<^<^v>>>v<><^<^^^^<>>^><^^v<<^v>vv>>^>v<^^v<^^v<^<>^>v>>>>v^vv><>>>v<<^^><>>>v<v^v>^^v^v><v<^^<>^vv>^^<>>^<>>v>v^^><><v^v<v<>>><^^vvvv<<^v<<<^<<^<<^>v<^v<v<>^vv>>><<>v<^>v^v>><^v<>v<><<^<^v<^>>^^^<<<vvvv<><<<<v^<^>^v<<>v>><^vv<^^^<<<><>^<vv^<<<>vv^<v><v>v>^vv<v>vv>><>>vv>v<^<v^v^^>v<>>>v^>vv^^v<>>v><<^vv^>vv>v<<<vv<<^<<^<v<vv^v^v^v>><^><>>v^v^<^vv<v<<^^<v>><vv^v<vvvvv^<^<<<<^>^v><<vvvv<v>><^v>^<<v^^<^v<v><<v<v><><>^<^>^^v>v><>>^^<vv>>^>v>vv<v<<v^<v>vv^><v^v><>v>>>^^>><vv^><v^<>^<^vvvv^^<v^^<<>>^^^^<>^<\n>v<v>^<^vvv><^<^>^^<>>^v<v<^<>vv<<>><<v<^v^<^>v^<v>v>vv<<^>v<>>^v>v<v><>v^<^^^^>>^<<<vv^>^v<<^>^vv<>^v<vv<vv<><<>^>vv<>v<v^<v<>v<<<v<^vvv>>^<<<>>><v>>^v<<>vvv>>v>vv<vvv>>^v^>><^^>>vvvv^^vv^^<v>v^v^>vv^v<><^^^^><<^v>>>vv<<><<>vv^>>>vv>v<v>vvvv<vv^^><^>^v<^v>>>><v^<v^^v>^<><<<<<v>v><^>v^^v<v><v<v^>^<v^^<<^<v<<^^>v<vvvv>>>v<><vvvvv<>><<vv^>v^^>>>vv^^vv>>v<^vv<v>>>^^>><^^<><^^<v>vv^v<^^<v>v^^>><^v<v^v<><^v<<<<<<^^^>^<>v<^><<^>>vv^v><<v^<<^>v<v<^<<v>^v>>^^v^^^^v>><^^>v>>>^<v^^^^>>^<vvv>>^v>>v>^v^>>^<>><>v^<>>v>^vv^<><v<v<<<^>>>>v<<<<<<^<v>^<>^>v<^vv>v^v^>>><^v>v>v<<>^v>^<^<^<>v>><<>^v^^><v<^><<<v<<<v<>>>v>v<v<<^^v<vv^>>vv^<>v^<^>>v^^<^>>^<<v^vvv^<<><^<><><vvvv>^v><v>^vvvv<<<v^><vvv<^>^<<^<^v^v>^v>^>>^^^<<^^vv<>><>>^<<>>^>v>v<>^><<v^>vv^>>vv^vvv<^^^>><<>^><<>v^v><<^vv<^v^<v^><^vvv^<><>^^>>><>v<^<<v<>>v<>vv<>>^<>>><><^v^<^^v^v<^^<<>>>^>^<<v<^<<<v^v^>>>v^><^<v>^<^^<>vv>^^<>^^>>v<v>^v<>v<>^^<v^>vvvv^<<vv^>v<^v^v^<^vvv^^>>>>v><v<<^v<^<>><v><<v^v<vv<<vvv^v<<<v^>^<vv<^^<>>^vvv<><^<<vv>>^>v>>>>>v>^\n<><v>>^v^vv><v<v^^>^^>>^^v^^<>^^^>^^^^v^v^<>><>v><^<>>>^v<vv^>v^<>^^v>vv<<^<<vv<<<^<<vv<<v^^>^^^<>vv^^^<>v><^><<>v<>vv<v<v^<v>v<<<>v>^<<<^^v>^<^<>v>>><>>^v<<<^>v>><v^><vv><>^><^v>><>^v^>^<>>^><<^><>>^v<^>v<vvv><<>^^>>v<v>^vv>^^^<<v^^<v>v^>>^>^^v<<^^<<vvv>vv^>^vvv<><^>^v^^v<<<<^v><^><^vv^^v^^<^vv^v<v^>>^><^<^v^v<>^><>v^v<>>>^<v^v<v>^><v^^v^v<<<^v<>><<<>^><v<v<>>v<>^>^>><v^<v<>^^<v^<v<<>><><<vv^^<vv<<<>^^^v>^>v<^<^<^>>^><<>><v^<v>vv<^^^>>^<^^<^<<<>^<<vvv^>><v^^^>><<<>>^>vv><^^><><><<>>>>^<>>>>><<<^<>v<v><>v<>v^><v>vv^^^vvv>v><v>>^>vv>vvv<vv<^>v>vv>^<<v^^<><>^>^<<vv<^<^<^^^^^v<^v<<>vvv><>>^v^^^v<^>vv>^>vvv>>>>>^^^>^<vv>v<^<>v><<>>vv^^v>v<>^<>vv<<<>vvv>>^vv<<>^^>v^>>v><<^>^<vvv<^<<<<<^>vvv><<>vv>>v>vv>^<>>^><v><^><<><^>>>^<^^v^<<^^<><^>vv^>v><v^^>v><^^<<<>^>v^^<^v>^<^v^^><v>^<<^^>><v>>v^<^>^v^>vv<><^^><v<>v^<>^v>^<^>>v><^>><^^v<v>vvv<v^>>^^v<^^><^^^>><vv>^>^^v<><>v^<v>^^v^><<v^vv>^>v<vv>>v<^<>vv><<>vvv^^^>^<v>>vvv<<<v>^v>vv>><^<v<>^v<v<^>vv<v>^>^vv><v^^>><v^v>^v<^^>^<>^<>^>v^<^vv><<<^<<<v^\n<^<v^v>vv<^vvv<<<v>><^<>^v^<><>>>^<^v><v<v<^<<v>v><<<<^^^<<^<>^v^v<>v^<v<>^>><>v>><^^>>v><>>^>^<>v<v<>v>v<<v^>^>vv^<vv^>^><^><<v^v<<>^<vv<v>^^v<v<>^^<>^^>^v<v^^<v^v><<>>v^^^vvv<<v^^v<v>>^>>><v><^v>><vv>^>^<<^^v^<^<^>^^^>^v>>v<>>>v^^><<>>>>>^v<v>vv^v^><>v>>vv>^^v>><>v<v^v^><v>v^><v^<>^<v<>><>>><<^<v^v^v>><>><>>v>><v^><<><><^>v<<>>>v^<v^>^vv^vv>v<^^vv>v<v>^v>>^<v^^v^v<<<>vv<><<v<v^>>v^><v>^v><<^>>v^^>^^^>><vv<>>><^v^v^^>^v>v><^<^<>^^<v^v>^>><^v<<>v^v>v^^vv^^<v^>^vv>^v^>v^><<<^<v>^<v>>^>>v<v>v<vv<^>>v>>^^>^>v^v<>v^>>v^^vv>vv<^v>><>vv<^^><vv<vvv^<v>>vv^<<v>>>v><<^<^v>>^>vv>^>><vv^v>><v^v<><>>v<^<^^^^v^>v^>><><^<<>vvv>>v^<<^>>^<^>v<<^v>>>>^v><v<vv<>v>>v^<<^>>^<>v^^>^<<<<v^v<v<^^><^^>^>>>v^v>^<<<v^<v^>v>^>>>><>>v><>>v>><>v^v<v^vvvvv<<^><^^>^<<<^<^^><^>v>^>v<v^>>v^<>>>vv<^>>v<>><^<<><<>^<>^^<^^v^<v><<>v<v^<<v><<v>vv^<^<v^<v<<^^vv>^><<>v>v^v<^v^>>>^^><v<<>^vv>><v<^<<>>vv^^>v>v<<>v<^<v>>^<^>>^<<vvv^^>><<>^<^v^<v>v^<<^^><>^<^^v<v>vv^v<<<<><<<<<<v^^>>v>v>^^<>^<^<<vv<^<><v><>^>^^v>^^^<>^>^^>>v>v<>\n^>><^^^<v>v<><>^^<v<<<<^><>>^^<^>vv>^>>^>^v>v^vvvv>^<^>>^v><^v>>><<>><<>v^><>^>><^><v^v><<v^<vv^v>>vv<><<<^^<>>v^<<><>v>vv><<^><>^<v>v<>^<>^>^<v<vvv<>^vv<<vv><<^vvv>v^<^<^v>>><^>vv>>v<>>>>^^^^^<^<v<>v<v<v<<<><v>vv><<^v>v><vv<^>><^vvv><v<v>^^v<^^v^^>^^><^<<v<<>^>^v<<^>^v^v^>vv^v^<^><<>>^v><>>>><^v><>^^v<^vv><^^<<<v>v<v^v>v^>>^><><^^v>^vv>v^<>vvv^>vv<vv^v^v>v^<>vvv<>v><<^<v^^>>^v<<<<vv^>^^<>v<<><v^v<<v><^^>^>^<^>^v>>^<<^>^v>vv>^<<^^^<vv<^<v^^vv><^v>v^<v<v^vv^<<vv^>>v><v><<>^^><^>vv<^<>><>^v^^^><>v^<><v^^v<<<<<<><^^v>vvvv^^^<<>^v><^^^^^v^<^<v<^vv^>>v>^v^^<<><><<>>>>><^<>^v<v>v><^^<<^>^>^<^v<^v>>^v^^><^^v>^<>vv<><^<vvv<^<<v^>^>^v^^>>>^^v<>>^>>><^vv<^<<<>v^<v<<>v^v^^^<^v<v>^>^^<>v<>>vv>^><^v<>>^^v><<vv<^>>v<>>v><<^v^v>^>><<^vv<>v^>^<v^^><^>v^<>v^v<><^vv^^v>^<><>>v>^<^<^><v>v^>><<vv<v<v>><^^<>>v>>^^^v>>>v<<>^v^^v<<v^^^<>v<v>^<<><<^vv<<<>>^v<>>vv><><>^^v^vv>>^v><v^>v^><^>v^v<^><^<<^>^>^^v>v<><<v^<vv^>v>><v^>v<<<^>v^^<<><v^v>^<^^vv<vv>vv^>>>>v><^v^>><vvv^<^>vv^v<<^>vv<v<>^<^<vv^^>><^^<^v><^v<^\n>v>>>^v<>v<v<>>>^^<<vv^>^<<^v>>>^^^v^<^v>vv^><^v^^^^^v^<vv^v>vv^>v<<v>>>>^<v<^<<<<<>>^>^<>^v>^><<^>><><>><^><^^v^^^^>^<v>>v<^vv><^^^v>><^>>>><>^>^><vv^v>>v>v<^>v^vv^<^v>^<^^>^>>^><<<^>><^^v^vv^^>v><<^<^^<v^>^^vv^<>vv<vv^>>^<v>v^^>><v^vv<^<<><^<<<<>>vvv<<>v><v>>^^<v>>^>><>>v><<^v>^^<v>^>v><vv^>^^vv>v>><v><v^^v<v><vv^<vv<><>><^v^<><>^v<><<^^vvv<<>v>>^<>v<<>v<v>v<>vv^^vvv^v^>>>^^<<^<^><^vv^>v<<>vv<<^>>>v<vv<><^>^<^><<v>vvvv<^^>vv<<^<v<<^>^v<<>>^>^><vv<^>>>><v^<<^>>v<vvv>^><v^<<v^v^v<<^v>><>v^^><<v<>v>>>v^<^<^^^>>v<v^>^>><<v^^^>^v<>v>^<^v<^>^v<><<<>^>^v^v>v^vv<<>>>>^<<>vv<^>>>^^^^>^v>v<v<^v^>><>>v<v><>>vv^^^^^>^>>^^v>^>>^v><>><^<^>^^^vv^v>v<^v^<>>vv^<v^^^><^<<>>^>vv<<<v><^^v^<<v<v^^>^<v<>>^^^<vv^><<^<><>><v>><<v^<^v<^>v^v<vvv^^<<><v<vv<<^<>><^<^v^>^>><<<^>>^<<^<vv>v<^<v<v^><><vv^v>^<^<^^<<>><<^vv<^^<<v<v<>>v<><v<><<><><><<v>v>^>v^<^v^>><>v>v<<>v<>^>vv^v<<><>>^<><v>>>>^v^v>>^v>vv>v^^<v<^<<v^^<^>v<>^^<>^^>>><^v><<>^vv<^vv><v<^>vv>>>>v<vvv<^^v^<<^>v>v<>v^>v<<>>v<>><<v^>>v^^<<<^^v^<v><^vv>vvv>\n>v^v^v>>v<^^^<^<^>vv>^^><<>^>><v<<<<<v^>><>^v^^vv<vv^>v<>><^<v^vv>v>>>^^vv<<v^vv<<<<^>>^v^>^^^^<>^v^^><^^v>v><v>v><v>vv><^^<^<^v>^v<<v><v<^<<v><<<<^v<v^v<v<><><<>^<v^^v^>>>><^v^>^><<<^<<v^v<>>><v^^^><<<<^vv^v<^^>v^v>^v<<>>>>>>v<^>^>^^v^^>^<vvvv<<>><<>v>v<>v>vvvv<>v^>^^vv<^v>v^^v<^<v<^^<><vv<^<v^v^^<vv<^^^^v<^<<v^<<v<^<<<^<vv>><v>>>><>vv^<^v^<^<<vv<^^^<<v>vvv<v>v><v^v^<<^<v>v>^^<<>>vv^v<>v^>v<>^<^v><><>vv><<<<<<^v<>><><<>>>^>>><>vv>^>><<^^>><v<^^^<v^<>^>^>>>v^v>v^v^^>v><^v^<v>^^^^>^v^<<^>>^<>^<><><<v<^^<vvv^^v>^^>><>^>v<<vv<<>><<^<>^>>vv<^>>v<><<>^<<v^^v<<<<<^><^v^>vv^v^<^vv^>>^<<^^^vv<vv^^>v<>><<<^^<v^<^<^>^v^<v^<v^><v^^^^>vv<vv<<v><>v^^>^<^<^<^^^v>^>>>>^^<^v^<>vvv<<<v<>v>v<^^^><><^><><<<<^><v<^^>vv<>^<^<v>v><<v>>>>v<<<>v<^v<v>^><<v^^>v<v>><<<vv>v>v>v<vv^^<>^^v><><<v><^^<v^><><<vv^^v^>v>^>>v>^^^<<^><^v>v>>^^v^vv^^^^v^v<<><>><<vv^vv<^>>^>>^<^^<<<<>^^><>>v<^^^^v^<>v>>v^^v<v^^>v>>^v<>>^<<^v^v^^v>^vvv^<v^>^<>v^^^v>^<v>>v^vv^<^>>>><>>^<^<<^>v>vvvvv<v<<<><<<vv><>^^^v>^>^>><<<><^<v>>^<><v^^><\n>>><><^^^^v<>>v^^>^^^^v<>^<^v<>vvv^^>^v><vv><><>^^v^^v><<<<^v>>^>^<>v><^^v^v^>><v>v^>v<v>^v^<vvvvv<vv>v>><v^vv^^^v><>^v<>^^^>><vv<>><<<v<vvv^<<^v^v>^vv^<^>>v>^v<>v^v^><^>v^<vv^<>vv>>v^>^v>v<vv^^vv>>v<^<<^vv<<^^><^vv^>>><>^v<^<<v<>>^^v><^^v<^v^<v^><<<v<<>^<><^vvv<><<v^<v><>^>v^v>v>^^<<<>^^>><>v>^v<^<>v>v><<vv<>^^>>>^vvv<^v><vv>v^<vv^v<vv^^vv<^v^<>v><<v<v><v<vv^v^vv^^^^vvv^v>^^>v<^^v>v>>>v>>>>vvv>^<^>>^v>^<v<>>vv><^>v>^>^^><vvvv>>>>^<<<<>^^v^<^^^v<<>>^v^<v^^<<<>>^v>vv^^v>>>>^<v>^^<v>v<<>^v^v^><>>^<v>v>v^><vv<^v^^>^>v>^v^><^>v<^v^vvv^^^<<<vv<^>><v<>>>vvv<v>v>^>^>><<>vv^^vv>v^>><^^>^^v><v<>vvv<>vv^^^^<^<vv>>^^v^<^>v^><^vv>>>^<^<<^^<^<><<^v<^<<>^><>v><vvv^>^><^<vvvv^^<<v<v<v<>v>><^v<v>v^vv<<>^v>>v^^^v>v<^^^<>vv>v^vv^>v<^vv><^v^^>^^vv>^^>>v><><vv><v>v^v>v<vv><^^^vv<>vv<v^^^>^^>v>^<^><>^^>vvv^>v<^<v^v^<<v><^^^<v>v>v^>v<><>v>^>v>>v<^v<<><>>>>^>v<v^v<<<v><<>>v>><<v><v>^>>^><^v><^^v<>v<^^>v^v<<<><^>><>vvv>vv<^>v>><v>>>^^>v<<<>v^v>^^^v^vv>vvv>>^><<vv<^v<<<>v^^>>^vv^<><<^v>^^>vv>>>v^>v>^>^vv^^>>v>\nv>><v<^v<>^>>v<v>^>v<^>^<<>^vv<v><^^v^<^^<^<<^<^>v<><>v<<>v<>><><vv^^<v><^^><v^^<<^^<<>v>vvv>vv<>^v^<><^^v^>v>^^v<><vv<>>v><<vv><<>>>^^^<vv>>v<>^v^<vv^v>^^vvvv^^v^<^^><><><<^>^>v<<>^^v>vv><<<v^>^>^<^v<^v^vv^<vvv>>v<>^^<v^vvv<><<v><>>^v<^<>v>^^>v^<v>>^><>^^^>>>^>^<<^^>^v^>^vv><v<v<>v^<v<<^v^^^>^^<>>^>^^v>v<<<>>v<>v<^^>vv<^<vv^<^<v<<>^<>v>>><>vvv<>^>^<<>><>^>><<>^>>v>v>v<^v>^>>>^>v^^<^v>v<<><^<^^<^>^^>vv>>>>^^<^v>v^<<>^>>^>>>^>v<v^<<<vv^<^>^<<v<vv<><^>>>^^v>>^^^vv><v><v>^^<>>v^^>^>vvv<>v<^v<^<>^^^^v^>v^<v^>^^^<<^<v^v<><v>^v<v^^><<><<^>>v<>v>v>^<^^^<v>>^^v<>>>^vv><<>v^v><<>>v>>^>v>>><>v^>v<vv<^>^<^v><vv^><>v><><>^>vvv^vv^<v^vv<vv^>v^v><>^><^>vv<>>>v>^>^v<^<<v<^>^^>vv<><>>v<<><<>>v>vv^^>^^v^^>>>^>v>>>^^<<<^<<<v^<>>>v>><><^<<>v<<><v<>>v<v<<^<v^>>^>^<>^<<vv^vv^><v^<^<>>>vv>v^<^vv>>^v><>><>>>><>vv^v>vvv>>vv<<v<>v>v>><^vv^<><>^^<^<>^><>>v>v^v>v^>^>v^<^>v^v^^^<v>vv^<<<v>vv<v<^<^v<^^>><^>^v<<^v<vv<><^^v><^<>vv<>v>>^<^^>>vvvv^^^<>>v^>^^v<>v>vvv><v<<<vv^v^^v^^<^>^<vvv><<<<>^>^v^^>><><vvv<v<^<>^^>^\nv<v^v>^>^^vv<>v>^><v>^v<<>>^>v<^vv>^<v^>>>><<vvvvv><^v<vv<^<<vv<><<>^vv^>^^^>>>>>>v^<v>v><><<^<>v<>^<v<>><>><>v>^<vv<><>v<<v>^^<>>^v<^<v>><^<v^vv>>v^vv^>v<<<<^v^v><^<^<>v^>^><>^^vv<^vv<>^<v^><>^<^>^^<vv^>^>^^^<vvvv^vv<vv<<v>^><v>v^vv^vv^><vvv^<v^^v^^<v<v><>^v><v<v^<>>><^^v^^>^^v<><^^^><vv>><v<<^<v>>vv>^^><><>v>>^^>>vv>^v<v<>v<^><vv><<v^<vv^>^>^^><^>>^^>^vv^v^^v<>v^<<v^^v<<^v^v>^^v<^<v<<<^>><v>>^^v><>^v^>^v<^^^^<>v^^v>^<v>^<<>><>>^^><vvv<<><^vv^^^v><v>^v<<>v>^<<v><<>^<>v<vv^vv<v^<>v>v>^^<>v<><>v>^<<>vv>><^<^<vv^<v^^<<<^^^>vvv><^v><>v>^<^>v<^v<<><^<vvvv^<v<<<^vv>v<v<v><^>>vv>^>>^>>>><<v^v^^<^v<^>vv<><^^<>v<vv>v^>v<^<><<<>vvv<>^vv^v^>v<v^>vv<><v^>v><>^^>v^>>>vv>><^^>v^^^><v>>v^<>>^vvv>^^v>vv>v>^^^^vv>>>^<><^>^v^v><><v>^^v<>>v^vv^>v^vv^<v^v<^v^^>v<<>^<<<vvv><^<<<<^><<<<<v^^>v^>^>>^vv<><<>vv><><>^v^>^>^v^^v<v><vv<<<>^^v><^<^<<<><>vv><v<^>v<^>v><v>^>>^<v<^^^><<^>vv^>v^<^>>>^<<v>v<>>>^<^<vv^^<>><<v><vv><v<^v><^^^<<>vv>v><^^>v<<v^v>^vv><>vvv<<^<>vv><v^^>>^vvv<>^^<>>v<<>^vv>>>>v>^^vv>>v<^<^>v<^\nvvvv^<<^v<^vv^^><v>>^vvvvv^^^>^>^<>^><<<vv>>><^vv>>>vvv><^>v<^>^<vv^^^^^><v<>^>^>^<vv^>>v>vv<<>>^v<^^<><<v<><>><<><>>><v<v^<>>>v^<^>>>>>v^>^v>^v><^<>^v><>v^^<<><v^^^><<>^^vvvv><<v^>>^^<>v^>>^>^^<^><v<vvv^>vv^^><v^<<vv^>v<>>>>^>vv<>vv<^<v^v^^v<>^>^^>vv>^<<^^>>v>v>^^><<<v><<<<^^>>v<<vvv<>>^^vv<v>v^>>^vv><^<^>>^>>^<^>>^<>><^<<<<<^v<v^>^^vv>^vv>v<>>^v<^^^v^^v<><<<^^v>>v>>^<<><v>vvv^>^^<^>v^^>v>v^vv<^v^vv<^v>^>>^>v<>v^<vvvv>v><<v>><v>^^<^^<^vv>^vvv^v<v><<^v^>>>>><><v>^^^vv>^v^><<^>^^v<<<>v^<vvv>^<^^^><^v<^^><vvv>^>>v>vv>vv^^>>v><^v<>^v<^^^<<^>v^v>>v<^v><<>>^^v<v^^^v<^vv<v^^<><<<vvv<>^<v<vv^^^><<>><v<v^<<<v<<<<>>v>vv>vvvv<vvvv<>v<><>>>>^vvv<<vv><^^>>v^<<^><^v>><v<<><<<^v^>><vv<<>^<^^^v<>^<><^v>v>vv^^><^>>v<^^vv^v>>v<<^^<^>><vv^^<^^>>v^><^vv>v><vvv>^vv^<^><^<^>v>v>v^>>^vv<<^<>>>>v<^^<v><^v^<^vvv>v><>^^^^vv>>^>>v<^vvv>v<v^v>^^<v^>^v^v<<<<^<v<^<^>^^vv^v<^^v^>>^>v<><<><>^<v>>^<<^<>^<>><^vv><><<vv<<vv^v^<vv>^>v<^^><>v^<>v<^<>v>>^vv>v^v<^>v^><^^>>><<v><v>><>>v<v^>^v>v<v^><^>^^^<<<>v^><<v^^v>v^^^v^\n^>v<v<v>><>>^vv>>>>^>^>><vv<>v>>><^>^><vv^<<>^<vv^>v>>^<^v^<^v^v<>>>><><<>>^^^v^^>>>vv<<<v<<^vv^>^>^<><v^><>>>vv^^<<v<vv^>^vv<><v>^<^^>v>v><<^vvv>>><<v^v>>v<^vv^v<<v^<^^v>vvv<<>>v>><>>v<>^^<<>>><^<^v^^^v<^<^>v><<>v<<^>><^<^>><vv^^v<^v<><<^^<>^v<>v^v><>^>>>^><>^^>v<>^<<^v>v>><^vv<<>^><>^<v>^<vv<><<<<>^<><^v^v>^<<^><v<<>>^>v<v^^^v<^><^<v<^^<>^^v><>><>>>v>><>^>><v<^v<vv><<<>><><>v^<v><<<>>>>^<>>^<<v<<v<>>>^^<<>v<v<^^v<v>v>>v>>v>><^>vv^v><^v>v^v<<^<vv<v^<>^<>v>><^<^v>v>v^<^v>v^>v^><<>>^>vv<^<>>>^v<^v<v^vv><<^>^^<>>v>^^v^vv>^^^^>>vv<<<v<^^><>^v^^^^>>^>^v^v^v<v<^v^vvv<<><>^<<>>vv^>>>>v><<v<v<>><v^vvv^v<>>^v<<<v^^<<>><^<^v^v<v>>>vvvv^vv><<><>>^<^vvv>>><^v<>>^v^^^<>>^^><<v^v>>^<^>>v><<>vv>><^v^vv<^^<<>v^v><^^<^<^^><>^v^^>v>v>^^<><v^<^^v<><v>>vv^v^>^v><<<^v>v>>^>v>v<^<>>v><^^<^v^v^<<vv><^^>^>v<>^^<v>vv<>>>>^><<>>vv<^^<v>>>^v>v<<v^>^>>>^^^v>v^^>>^^^>^v><^>vv>vv<>^><v>^>vv^>^<<<<v^>v^>^^vv<<>><><<^^<^^v^vvv<<^<<v^^^<^v^v>>v<>v^v><^v^>^>v^v^v^<><<vv>><<<v^<v^vv><^<^<^<>>>>v<vv>^vvv^>^v<^^<v>>^v><v\n><v>>>^^>^v>vv^v^v^><^v<v<>vv^>>^>>vv><>vv^^>^><>^^v<v^><v><^^^v><>><<v<vvv<^><^><>^^v<>vv<v^<>^>^<^vv^^<^<>^>^v<vv^<<^^<<v>^>>vvv>^v<><vvv>>>>><<v^<^<<<^><^>^>v^^><<<v<v^<<v<>v><<>vvv<<<v>vv>vv^>^>><>>^>><v^<<v>^^><^v<<^<^>>^^>^vvv<<v>>^<>>^>^><v><^<>v<<<^><^^^>^^>>^<^<>>^^<<vvvv<>^>>^v<>^v^v<><v<vv^<^<^^>v<^>^<<<<<>^<^>><v^<v^^<<>^v^>^<<><><>^<<vv^<^v>v<><><v>^^v<v>v>>vv^>>><<><^^^><^<v^><^vvv<>v>>^<>>vv<v<<>^^><^^^^^^<vv>^><^^><>>>^<^^v<<v>v^<>>^v><^><v^^vv<>v>>^v^vv^>>^<^vv><><>^^<vv^^<vv^v^^^vv^<<^^^>v<vv><>^>^>vv<^^^^^^<vv^^<<>^v<v^<^>vv><>>><><>^<v>v^v^vvvvvv^^><<<<><^^^^>^<v<vv^v^>v^v<<>v<>^>>^<><^^>v<<v^<<<<^>v<>v<^<<v>v<vv^>^>>v<^^v<vv<^<^v^><>^<v^<<><<>^^>^>>^^v^v><^<>v^<>v<<><><v^>>v>vvvvv<>^<<><^v<<<><<<v>>^<^<<^><^>v>v<>^v<v>^<v^^^<<<vv^<v<^v^<>>>v<^^^^^><^v^v^<v^<v>v<v<>vv^vv<>vv<v>>vvv>><>^v>v^<vvvv>^><<v>vv^v<v>^v>v^><>^>v^>v>v<vvv<>v^>v<^>^vvv<<vvv^<v^>v^>><v^<<^>^><><v>^<>^v>v^vvvvvvv>v<>v^<>>>^^>v>vv^>v^>><<<><<vv^<>^vvv^>v^^>>>^>v<^v>^^<>v<^v^^><^v<vv>>^v<><vv^v<><\n^v><v<^v^^v^>v><<>v^>vv>v^vv^<^^><<^<<<<>>>><^^>v><vv<vv<<^>><>>>>v^^^>v^v^v^vv<^vvv^><<v>v<<vvv<v>>v<^><vv^>v<v>v^<<<v^>^v><^^>^>^vv>v<<<>v><>><>v^v^>vv<v^>v^vv>vvvv<>>^>^vv<<<^^^^v^<>^^<v>^><v^^><^v>v<v>>><^>>vv>^^>>>v<v<v<^^vv^>^<>v<><v^><v>>v<^v^>v^v><<>v^^^^^^^v^>^^^<^vv<<^>><<vvv>^>^>v<v<<>^<><<>>^>vv>vv>^>v<vv<^^>>^>^v>^v>><<><^^^^<v^<>>>^v^<v>^v<>>>v<v<v<<>vv^v<v<><v<><vv>>v>v<vv>^<>vvv<<>>v^<v><><^v><^^^vv^v<<<>vv<<vv^v<v^>v>^^v^>>><v^><<<<<^^<^^><<v^v>^v<^<v^<<><v<^<><<v>^v^<<^^<><v>><v^>^^<<<><>v<v>>vvv<<>>><^<v>>^^v>^vv^>v^<<^^^^>>>v^^<>><^<><vv^>>^vvv<^v<<>>^><^^v^<<^^^^<v>>><<<v^v<>>^v>><v<>>vvv^>><^vvv>v>^v^^v<>><^>><^>><<^>>v><>^>>^>v^>v<v<<<>v^vv<^vv>>v<v><^vv>v<^^>v>v>v^v<v^vvv^>v^<^v<v<>v^>>^<>v><>^><<>><v<>vv>^^<><^<><><<<>^>^^vv<><vvv>^>>><^v^<><<<v>v>v<^>>v<><^<^<^>v>^<vv<v^>^^^>^><<^>>>v<v<v<v>v>>vvv^>v>v<<^>^v^<^^^<<<<<^<^^><v>v^<v<>>>><v>v^<><v>^^<vv<vvv<vv^^<><><^><v<><vv^^^<^<<<>^vv<<^<v>^v>><v<vvv<^^v>^v^^v^vv^v^^><v<v^<>^<^v^>v<vv>v<v>^>vvv^>^<v<^^><>>v<<<v\n^>^v<<>^<v<v^^<v>>vvvv^<v^v><>v>>v^^v>>><>>v^<><^>v><<v^>^>>><>v>vv^<v^>>>^<^^<><><v<>v^^>v><v<<vvvvv<v><vv<<v<v^^v<>v<^v^<^^><^^<v^<v><<^<>^^vvvv>^v<>^^<^v<v^v><v>>vv><v^>^>v<^^<^^>>v^v^^^v>^v^>^>>^vv^v>^<v<<vv><<>^><^^^v>v^><>>v<v^<v^^v>v>><vv^<v><>v<>>vvv^><>><v<<vv^<<^<<^>>^^>v^v<<<>v>v<<v>^>v><><<^>><<^>^>v><v>v^<^v><<v<^>vv^<<vvv<vv<>^>>^<v><<>v^^>>>vv^>^<>^^<<^^><>^<>^^><<^^<><>><<>>^^^<^^>v<>^^>>^<><<<v^^v<v>^vv^>v<<^>v^>^<><<v<><<>><^^^>>^><^v^<^<vv>>>v<vvv^^^<^v>v>^^vv^<v<>>^vv<>><><^>v>^><><^v><v<v^v<<<<>vvv<<^>^^^v>v^>^>v>^>^^<>v<<^<^<^><vv<v^v<v<<>>^><<>v^^><<^<<<vv<>vv><^^>><>vv<^>^>>><<^v><<^>vvv<^>vv>>>^^vv<><<^><^><>^>v><<^>^^^<>v>v<vv>^><v>^>v^><>>^<v^><<<vv>^<^vv<<<^v^<>^<<>^>v>>><<^^v>>v<v^v<<<vv^v^v>>v^^>vv<v^vv^v<vv>v^<<^<^^^v>>^>^>>v^<v^>^^vv>v<^>>^<v^>^^v<^^>vv<^v><>>vv<>>^>v>vv^v^^^>^v^^>^<<^<<^>>v<v<<>vv<v<v<^^v<v^^v<>><<<v>^>>vv<v>^^^><vv^<<><>^<<v^<><><>^>^vvv>><<<v^<<v^^<<v<<^<v<^<<<><^vvv><<vvv<v<<<>^>^<>v<^^v><<^><^><>v><^<^<>vv^^<vv^v>v^^^<>><vv^>vvvv<v^\n^v<>vv>v<v<vv>>>v^^>>^vv>v>^v^^v><><<^<>>^v^>v<v><^v>^v<>^>>>>^>v>>v^>^<^^>v^^^vvv^v>^>v^^^^vv^^><v<v<>^<v^<vvv^<^>v^<>><^<<<><^^^v^v>v^<><<>^<>^vv<vv<^^^<><^^v><>v><v>v<<^<^v>^v>vv^^>vv<><<^>^<<v>>>v^<v<<>vv^<<v>v>><^<<v>><^<vv<v>>>>^<<><^<v<>^<><>vv><>v>v<><>^^<vv^<>>>^^vv^v<>vv<>vv<^v^>v<v^^v>^>^<^>^<^>v<><^v<>v^>>^<v<^v<v>>>v<vvv<^^<vv>>v<<>v<><<><<>^^<<^<v>^^>v^v>>v^vvv^<^>^v^^>^v>>^v^v^^^vv^^<<>^<<^<<v^vv>^>>^<<<^>vvvv<<>>>^<<>vvv^v>^>vv<<^>^>vv<vv>>v<v><vvv^^v^<<><><<^^^^^v<^<v^^^^vv<^>>><v>vvvvv^>^>>v<^>><>v<vv><^<<<v>v^<^>^^^<>>^^v><<>v<>v<^^<>>^><>>>^v^^v<<<<^<<<v^^v>^>v<>v>>^>v>^^<v>><<><<^>v^^^>^<^>v>v^>v<<^v<^<^v<><>>^>v^<<v>>^><<<^<<vvv<<vv><^>^>>>>>vv^v<v<^<v^<<<v<>v<vv>^>>>v>>^<<>>>>>><>vv>><vvv<><v<v^^>vv^<^>>v<>>^<>^v<>>v<>><^>^^^^<>v<v^>^^vvv>^^><^v>v^v<v<>>>>v><^^>^<><>>v^<^^^<^<<^v^v^vv>>v^<<^>^<>^^v^>^>v^^^>^^^<vvv<<<>v>><<>>^vv^^v^^><<>^>^>>^vv^v>^^<>^v^^^<>^v<^>v^^^><<v^^^>>^><<<^v<^v^v>v^>^<<><vv<<>^>^<v<^v<^><<^>v^<v<<>>><>^v^>v^^<^<vv><><><<^v<<<><v>>vv>^>^><"
	board := parseInput("./day15/input.txt") //part 1
	robotMoves := movesToDir(moves)
	robotX, robotY := getRobotStartPos(board)
	finalBoardState := Play(robotX, robotY, board, robotMoves) //part 1
	PrintMatrix(finalBoardState)                               //part 1
	fmt.Println(ComputeBoardScore(finalBoardState))            //part 1
}

func parseInput(filename string) [][]rune {
	board := make([][]rune, 0)
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		board = append(board, []rune(scanner.Text()))
	}
	return board
}

func movesToDir(moves string) []Direction {
	dirs := make([]Direction, 0)
	for _, move := range strings.Split(moves, "") {
		switch move {
		case "^":
			dirs = append(dirs, up)
		case "v":
			dirs = append(dirs, down)
		case "<":
			dirs = append(dirs, left)
		case ">":
			dirs = append(dirs, right)
		}
	}
	return dirs
}

func getRobotStartPos(board [][]rune) (int, int) {
	for i, row := range board {
		for j := range row {
			if board[i][j] == '@' {
				return i, j
			}
		}
	}
	return -1, -1
}

func Play(robotPosX, robotPosY int, board [][]rune, robotMoves []Direction) [][]rune {
	for _, dir := range robotMoves {
		// new Pos for robot if he can move to it, based on if there is a wall that block the dir or boxes he can/can't move
		// also board is surround by walls wo we can't go out of range, we will meet a wall before
		newX, newY := robotPosX+dir[0], robotPosY+dir[1]
		switch board[newX][newY] {
		case '#':
			continue
		case '.':
			board[robotPosX][robotPosY] = '.'
			robotPosX, robotPosY = newX, newY
			board[robotPosX][robotPosY] = '@'
		case 'O':
			// idea : check if at the end of the steak of 0 (if there is many) there is a wall
			// if not, robot take the pos of the first box (box = 0), and first box takes the place of the free space found
			// so instead of moving the full line of boxes we are just moving one box but we get same results
			nextPosAfterBoxX, nextPosAfterBoxY := newX+dir[0], newY+dir[1]
			for board[nextPosAfterBoxX][nextPosAfterBoxY] == 'O' {
				nextPosAfterBoxX, nextPosAfterBoxY = nextPosAfterBoxX+dir[0], nextPosAfterBoxY+dir[1]
			}
			if board[nextPosAfterBoxX][nextPosAfterBoxY] == '#' { // wall, can't move, do nothing
				continue
			} else { // can move, free space
				board[robotPosX][robotPosY] = '.'
				robotPosX, robotPosY = newX, newY
				board[robotPosX][robotPosY] = '@'
				board[nextPosAfterBoxX][nextPosAfterBoxY] = 'O'
			}
		}
	}
	return board
}

func ComputeBoardScore(finalBoardState [][]rune) int {
	var res int
	for i, row := range finalBoardState {
		for j, col := range row {
			if col == 'O' {
				res += i*100 + j
			}
		}
	}
	return res
}

// helper
func PrintMatrix(matrix [][]rune) {
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%c ", col)
		}
		fmt.Println()
	}
}

func parseInput2(filename string) [][]rune {
	board := make([][]rune, 0)
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "#", "##")
		line = strings.ReplaceAll(line, "O", "[]")
		line = strings.ReplaceAll(line, ".", "..")
		line = strings.ReplaceAll(line, "@", "@.")
		board = append(board, []rune(line))
	}
	return board
}
