package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"io"
)

type Tree struct {
	Name string
	Value int
	Children []Tree
}

func (t Tree) Weight() int {
	sum := 0
	for _, child := range t.Children {
		sum += child.Weight()
	}
	return sum + t.Value
}

func (t Tree) isBalanced() bool {
	if len(t.Children) == 0 {
		return true
	}
	values := make(map[int]bool)
	for _, child := range t.Children {
		values[child.Weight()] = true
	}
	return len(values) == 1
}

func (t Tree) findUnbalanced() Tree {
	if len(t.Children) == 0 {
		return Tree{}
	} else {
		for _, child := range t.Children {
			if !child.isBalanced() {
				return child.findUnbalanced()
			}
		}
		return t
	}
}

func getLines() []string {
	lines := make([]string, 0, 0)
	file, _ := os.Open("day7/real_input")
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		lines = append(lines, line)
		if err != nil {
			if err != io.EOF {
				println(" > Failed!: %v\n", err)
			}
			break
		}
	}
	return lines
}

func treeFromString(line string) Tree {
	chunks := strings.Split(line, "->")
	t := Tree{}
	left := chunks[0]
	left = strings.TrimSpace(left)
	pieces := strings.Split(left, " ")
	t.Name = pieces[0]
	value, _ := strconv.Atoi(strings.Trim(pieces[1], "()"))
	t.Value = value
	if len(chunks) > 1 {
		right := strings.TrimSpace(chunks[1])
		for _, piece := range strings.Split(right, ", ") {
			child := Tree{piece, 0, []Tree{}}
			t.Children = append(t.Children, child)
		}
	}
	return t

}

func findRoot(nodes []Tree)  Tree {
	children := make([]string, 0, 0)
	for _, node := range nodes {
		for _, child := range node.Children {
			children = append(children, child.Name)
		}
	}
	for _, node := range nodes {
		isChild := false
		for _, name := range children {
			if node.Name == name {
				isChild = true
				break
			}
		}
		if !isChild {
			return node
		}
	}
	return Tree{}
}

func composeTree(nodes []Tree, root *Tree){
	setChildren(nodes, root)
	for _, child := range root.Children {
		composeTree(nodes, &child)
	}
 }

func setChildren(nodes []Tree, root *Tree) {
	for i, fakeChild := range root.Children {
		for _, node := range nodes {
			if fakeChild.Name == node.Name {
				root.Children[i] = node
				break
			}
		}
	}
}


func main() {
	nodes := make([]Tree, 0, 0)
	for _, line := range getLines() {
		t := treeFromString(line)
		nodes = append(nodes, t)
	}
	fmt.Println(nodes)
	root := findRoot(nodes)
	fmt.Printf("Part1: Root is %v \n", root)
	composeTree(nodes, &root)
	fmt.Printf("Composed Tree: %v \n", root)
	unbalanced := root.findUnbalanced()
	fmt.Printf("Unbalanced Tree: %v \n", unbalanced)
	for _, child := range unbalanced.Children {
		fmt.Printf("%v -> %v/%v\n", child.Name, child.Value, child.Weight())
	}
}
