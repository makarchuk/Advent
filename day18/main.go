package main

import (
	"strconv"
	"strings"
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
	case "mul": {
		reg := string(instr[4])
		sc.registers[reg] = sc.registers[reg] * sc.getSecondArg(instr)
		sc.position += 1
		return "MUL", sc.registers[reg], false
	}
	case "mod": {
		reg := string(instr[4])
		sc.registers[reg] = sc.registers[reg] % sc.getSecondArg(instr)
		sc.position += 1
		return "MOD", sc.registers[reg], false
	}
	case "snd": {
		reg := string(instr[4])
		sc.position += 1
		return "SND", sc.registers[reg], false
	}
	case "rcv": {
		reg := string(instr[4])
		val, err := sc.pop()
		if err != nil {
			return "RCV", val, true
		} else {
			sc.registers[reg] = val
			sc.position += 1
			return "RCV", val, false
		}
	}
	case "jgz": {
		reg := strings.Split(instr[4:], " ")[0]
		num, err := strconv.Atoi(reg)
		val := 0
		if err == nil {
			val = num
		} else {
			val = sc.registers[reg]
		}
		jump := sc.getSecondArg(instr)
		if val > 0 {
			sc.position += jump - 1
		}
		if reg == "f" {
			fmt.Printf("JUMP \"%v\" -> %v: %v. i=%v\n", reg, val, jump, sc.registers["i"])
		}
		sc.position += 1
		return "JGZ", jump, false
	}
	default: {
		fmt.Printf("Unknown instruction %v", instr)
	}
	}
	return "SNH", 0, true//Should not happen
}

func (sc *SoundCard) firstRcv() {
	for true {
		CMD, val, _ := sc.execute()
		if CMD=="SND" {
			sc.play(val)
		}
		if CMD=="RCV" && val != 0 {
			fmt.Printf("First non-zero RCV is %v\n", val)
			break
		}
	}
}

func (sc SoundCard) getSecondArg(instr string) int {
	stringValue := strings.Split(instr[4:], " ")[1]
	value, err := strconv.Atoi(stringValue)
	if err != nil {
		return sc.registers[stringValue]
	}
	return value
}

func testSoundCard() SoundCard {
	return SoundCard{
		make(map[string]int),
		[]string{
			"set a 1",
			"add a 2",
			"mul a a",
			"mod a 5",
			"snd a",
			"set a 0",
			"rcv a",
			"jgz a -1",
			"set a 1",
			"jgz a -2",
		},
		0,
		[]int{},
	}
}

func realSoundCardWithPid(val int) SoundCard {
	sc := realSoundCard()
	sc.registers["p"] = val
	return sc
}

func realSoundCard() SoundCard {
	return SoundCard{
		make(map[string]int),
		[]string{
			"set i 31",
			"set a 1",
			"mul p 17",
			"jgz p p",
			"mul a 2",
			"add i -1",
			"jgz i -2",
			"add a -1",
			"set i 127",
			"set p 464",
			"mul p 8505",
			"mod p a",
			"mul p 129749",
			"add p 12345",
			"mod p a",
			"set b p",
			"mod b 10000",
			"snd b",
			"add i -1",
			"jgz i -9",
			"jgz a 3",
			"rcv b",
			"jgz b -1",
			"set f 0",
			"set i 126",
			"rcv a",
			"rcv b",
			"set p a",
			"mul p -1",
			"add p b",
			"jgz p 4",
			"snd a",
			"set a b",
			"jgz 1 3",
			"snd b",
			"set f 1",
			"add i -1",
			"jgz i -11",
			"snd a",
			"jgz f -16",
			"jgz a -19",
		},
		0,
		[]int{},
	}
}

func part1(sc SoundCard) {
	sc.firstRcv()
}

func part2() {
	p1 := realSoundCardWithPid(0)
	p2 := realSoundCardWithPid(1)

	var p1SND, p1RCV, p2SND, p2RCV int
	for true {
		CMD1, val1, blocked1 := p1.execute()
		//fmt.Printf("P1 call %v %v %v. Position %v. Reg %v\n", CMD1, val1, blocked1, p1.position, p1.registers)
		if CMD1 == "SND" {
			p1SND += 1
			p2.play(val1)
			//fmt.Printf("P2 stack: %v\n", len(p2.played))
		}
		if CMD1 == "RCV" {
			//fmt.Printf("P1 stack: %v, blocked: %v\n", len(p1.played), blocked1)
			if !blocked1 {
				p1RCV += 1
			}
		}
		CMD2, val2, blocked2 := p2.execute()
		//fmt.Printf("P2 call %v %v %v. Position %v. Reg %v\n", CMD2, val2, blocked2, p2.position, p2.registers)
		if CMD2 == "SND" {
			p1.play(val2)
			//fmt.Printf("P1 stack: %v\n", len(p1.played))
			p2SND += 1
		}
		if CMD2 == "RCV" {
			//fmt.Printf("P2 stack: %v, blocked: %v\n", len(p2.played), blocked2)
			if !blocked2 {
				p2RCV += 1
			}
		}
		if blocked1 && blocked2 {
			fmt.Printf("Stats: %v/%v, %v/%v \n", p1SND, p2RCV, p2SND, p1RCV)
			fmt.Printf("P2 sent signal %v times", p2SND)
			break
		}
	}
}

func main() {
	part1(testSoundCard())
	part1(realSoundCard())
	part2()
}
