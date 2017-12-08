package main

import (
	"strings"
	"strconv"
	"fmt"
	"os"
	"bufio"
	"io"
)

type Operation struct {
	Variable string
	Operation string
	Value int
}

type Condition struct {
	Variable string
	Operation string
	Value int
}

type Instruction struct {
	Op Operation
	Cond Condition
}

type Registry map[string]int

func (c Condition) Check(reg Registry) bool {
	variable, _ := reg[c.Variable]
	switch c.Operation {
	case ">":
		return variable > c.Value
	case "<":
		return variable < c.Value
	case "!=":
		return variable != c.Value
	case "==":
		return variable == c.Value
	case ">=":
		return variable >= c.Value
	case "<=":
		return variable <= c.Value
	}
	fmt.Printf("This should not happen! Condition is: %v\n", c)
	return false
}

func (op Operation) Perform(reg *Registry) {
	variable, _ := (*reg)[op.Variable]
	switch op.Operation {
	case "inc":
		(*reg)[op.Variable] = variable + op.Value
	case "dec":
		(*reg)[op.Variable] = variable - op.Value
	}
}

func (ins Instruction) Execute(reg *Registry) {
	if ins.Cond.Check(*reg) {
		ins.Op.Perform(reg)
	}
}

func ParseInstruction(row string) Instruction {
	chunks := strings.Split(strings.TrimSpace(row), " if ")
	return Instruction{
		ParseOperation(chunks[0]),
		ParseCondition(chunks[1]),
	}
}

func ParseOperation(row string) Operation {
	chunks := strings.Split(row, " ")
	value, _ := strconv.Atoi(chunks[2])
	return Operation{
		strings.TrimSpace(chunks[0]),
		strings.TrimSpace(chunks[1]),
		value,
	}
}

func ParseCondition(row string) Condition {
	chunks := strings.Split(row, " ")
	value, _ := strconv.Atoi(chunks[2])
	return Condition {
		strings.TrimSpace(chunks[0]),
		strings.TrimSpace(chunks[1]),
		value,
	}
}

func Part1(prog []Instruction) {
	reg := make(Registry)
	for _, inst := range prog {
		inst.Execute(&reg)
	}
	maxV := 0
	for _, v := range reg {
		if v > maxV {
			maxV  = v
		}
		fmt.Printf("Registry is %v\n", reg)
	}
	fmt.Printf("Maximum value is %v\n", maxV)
}

func Part2(prog []Instruction) {
	reg := make(Registry)
	maxV := 0
	for _, inst := range prog {
		inst.Execute(&reg)
		for _, v := range reg {
			if v > maxV {
				maxV  = v
			}
		}
	}
	fmt.Printf("Maximum value ever is %v\n", maxV)
}


func main() {
	program := make([]Instruction, 0, 0)
	file, _ := os.Open("day8/real_input")
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		inst := ParseInstruction(line)
		program = append(program, inst)
		if err != nil {
			if err != io.EOF {
				println(" > Failed!: %v\n", err)
			}
			break
		}
	}
	fmt.Println(program)
	Part2(program)
}
