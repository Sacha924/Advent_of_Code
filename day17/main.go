package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

////go:embed sample2.txt
//var inputTest2 string

const (
	A = 0
	B = 1
	C = 2
)

type Machine struct {
	ip        int
	registers [3]int
	program   []int
}

type Output struct {
	out []int
}

func (m *Machine) Register(n int) int {
	if n < 0 || n >= len(m.registers) {
		return 0
	}
	return m.registers[n]
}

func (m *Machine) SetRegister(n int, v int) {
	if n < 0 || n >= len(m.registers) {
		return
	}
	m.registers[n] = v
}

func (m *Machine) Ip() int {
	return m.ip
}

func (m *Machine) Run(a, b, c int) Output {
	m.ip = 0
	m.registers[A] = a
	m.registers[B] = b
	m.registers[C] = c
	var output = Output{}

	for m.ip >= 0 && m.ip < len(m.program) {
		var inst = m.program[m.ip]
		var operand = m.program[m.ip+1]

		var combo = func(v int) int {
			if v == 7 {
				panic("not implemented")
			}
			if v >= 0 && v <= 3 {
				return v
			}
			return m.Register(v - 4)
		}

		var dv = func(reg int) {
			var num = m.Register(A)
			var den = 1 << combo(operand)
			var res = num / den
			//fmt.Printf("%d-dv operand: %d num: %d deb: %d res: %d\n", reg, operand, num, den, res)
			m.SetRegister(reg, res)
		}

		switch inst {
		case 0: // adv
			dv(A)
		case 1: // bxl
			m.SetRegister(B, m.Register(B)^operand)
		case 2: // bst
			var res = combo(operand) % 8
			m.SetRegister(B, res)
		case 3: // jnz
			if m.Register(A) != 0 {
				m.ip = operand
				m.ip -= 2
			}
		case 4: // bxc
			m.SetRegister(B, m.Register(B)^m.Register(C))
		case 5: // out
			output.out = append(output.out, combo(operand)%8)
		case 6: //bdv
			dv(B)
		case 7: // cdv
			dv(C)
		}

		m.ip += 2
	}
	return output
}

func CreateMachine(inst []int) *Machine {
	return &Machine{program: inst}
}

func parse() ([]int, int, int, int) {
	registers := "Register A: 729\nRegister B: 0\nRegister C: 0"
	program := "Program: 0,1,5,4,3,0"
	var a, b, c int
	fmt.Sscanf(registers, "Register A: %d\nRegister B: %d\nRegister C: %d", &a, &b, &c)
	var _, after, _ = strings.Cut(program, " ")
	var inst []int
	for _, e := range strings.Split(after, ",") {
		var v, _ = strconv.Atoi(e)
		inst = append(inst, v)
	}
	return inst, a, b, c
}

func Part1() string {
	var inst, a, b, c = parse()
	var m = CreateMachine(inst)
	var o = m.Run(a, b, c)
	var res string
	for i, e := range o.out {
		if i > 0 {
			res = res + ","
		}
		res = res + strconv.Itoa(e)
	}
	return res
}

func solveFast(m *Machine, a int, index int, expected []int) int {
	//fmt.Println("index", index, "start", start)
	if index < 0 {
		return a
	}
	for i := 0; i < 8; i++ {
		var nextA = 8*a + i
		if o := m.Run(nextA, 0, 0); slices.Equal(o.out, expected[index:]) {
			if v := solveFast(m, nextA, index-1, expected); v >= 0 {
				return v
			}
		}
	}
	return -1
}

func Part2() int {
	var inst, _, _, _ = parse()
	var m = CreateMachine(inst)
	return solveFast(m, 0, len(inst)-1, inst)
}

func main() {
	start := time.Now()
	fmt.Println("part1: ", Part1())
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2())
	fmt.Println(time.Since(start))
}
