package main

import (
	"strings"
	"strconv"
	"fmt"
)

type SoundCard struct {
	registers map[string]int
	instructions []string
	position int
	played []int
}

func (sc *SoundCard) pop() (int, error) {
	if len(sc.played)  == 0 {
		return 0, fmt.Errorf("Nothing to pop")
	}
	res := sc.played[0]
	sc.played = sc.played[1:]
	return res, nil
}

func (sc *SoundCard) play(freq int) {
	sc.played = append(sc.played, freq)
}

//Returns command, register, value, blocked
func (sc *SoundCard) execute() (string, int, bool) {
	if sc.position < 0 || sc.position >= len(sc.instructions) {
		return "OOR", 0, true//Out of Range
	}
	instr := sc.instructions[sc.position]
	switch instr[0:3] {
	case "set": {
		reg := string(instr[4])
		sc.registers[reg] = sc.getSecondArg(instr)
		sc.position += 1
		return "SET", sc.registers[reg], false
	}
	case "add": {
		reg := string(instr[4])
		sc.registers[reg] = sc.registers[reg] + sc.getSecondArg(instr)
		sc.position += 1
		return "ADD", sc.registers[reg], false
	}
	case "sub": {
		reg := string(instr[4])
		sc.registers[reg] = sc.registers[reg] - sc.getSecondArg(instr)
		sc.position += 1
		return "SUB", sc.registers[reg], false
	}
	case "mul": {
		reg := string(instr[4])
		sc.registers[reg] = sc.registers[reg] * sc.getSecondArg(instr)
		sc.position += 1
		return "MUL", sc.registers[reg], false
	}
	case "jnz": {
		reg := strings.Split(instr[4:], " ")[0]
		num, err := strconv.Atoi(reg)
		val := 0
		if err == nil {
			val = num
		} else {
			val = sc.registers[reg]
		}
		jump := sc.getSecondArg(instr)
		if val != 0 {
			sc.position += jump - 1
		}
		sc.position += 1
		return "JNZ", jump, false
	}
	default: {
		fmt.Printf("Unknown instruction %v\n", instr)
	}
	}
	return "SNH", 0, true//Should not happen
}

func (sc SoundCard) getSecondArg(instr string) int {
	stringValue := strings.Split(instr[4:], " ")[1]
	value, err := strconv.Atoi(stringValue)
	if err != nil {
		return sc.registers[stringValue]
	}
	return value
}


func realSoundCard() SoundCard {
	return SoundCard{
		make(map[string]int),
		[]string{
			"set b 81",
			"set c b",
			"jnz a 2",
			"jnz 1 5",
			"mul b 100",
			"sub b -100000",
			"set c b",
			"sub c -17000",
			"set f 1",
			"set d 2",
			"set e 2",
			"set g d",
			"mul g e",
			"sub g b",
			"jnz g 2",
			"set f 0",
			"sub e -1",
			"set g e",
			"sub g b",
			"jnz g -8",
			"sub d -1",
			"set g d",
			"sub g b",
			"jnz g -13",
			"jnz f 2",
			"sub h -1",
			"set g b",
			"sub g c",
			"jnz g 2",
			"jnz 1 3",
			"sub b -17",
			"jnz 1 -23",
		},
		0,
		[]int{},
	}
}
func realSoundCardWithA(a int) SoundCard {
	return SoundCard{
		map[string]int{"a": a},
		[]string{
			"set b 81",
			"set c b",
			"jnz a 2",
			"jnz 1 5",
			"mul b 100",
			"sub b -100000",
			"set c b",
			"sub c -17000",
			"set f 1",
			"set d 2",
			"set e 2",
			"set g d",
			"mul g e",
			"sub g b",
			"jnz g 2",
			"set f 0",
			"sub e -1",
			"set g e",
			"sub g b",
			"jnz g -8",
			"sub d -1",
			"set g d",
			"sub g b",
			"jnz g -13",
			"jnz f 2",
			"sub h -1",
			"set g b",
			"sub g c",
			"jnz g 2",
			"jnz 1 3",
			"sub b -17",
			"jnz 1 -23",
		},
		0,
		[]int{},
	}
}

func part1(p SoundCard) {
	var mulCount int
	for true {
		CMD, _, blocked := p.execute()
		if CMD == "MUL" {
			mulCount += 1
		}
		if blocked {
			break
		}
	}
	fmt.Printf("%v MUL calls\n", mulCount)
}

func part2(p SoundCard) {
	for i:=0; i<1000; i++ {
		fmt.Println(p.instructions[p.position])
		CMD, _, blocked := p.execute()
		if CMD == "JNZ" && p.position == 11 {
			fmt.Printf("a: %v, c: %v, f: %v, e: %v, g: %v \n", p.registers["a"], p.registers["c"], p.registers["f"], p.registers["e"], p.registers["g"])
		}
		//fmt.Printf("CMD: %v, position: %v\n", CMD, p.position)
		if blocked {
			break
		}
	}
	fmt.Printf("Final h values is \n", p.registers["h"])
}

func main() {
	part1(realSoundCard())
	part2(realSoundCardWithA(1))
}