package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type problem struct {
	num      string
	name     string
	link     string
	level    string
	preNum   string
	preKnowl string
	curKnowl string
}

var cnt int

func main() {
	rootReadme := filepath.Join("./../..", "README.md")
	dat, err := os.ReadFile(rootReadme)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(dat), "\n")
	for i := range lines {
		lines[i] = parseLine(lines[i])
	}
	output := strings.Join(lines, "\n")
	output = strings.ReplaceAll(output, "probCnt", strconv.Itoa(cnt))
	err = os.WriteFile(rootReadme, []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("目前已完成 %d 道题目\n", cnt)
}

func parseLine(s string) string {
	a := strings.Split(s, "|")
	for i := range a {
		a[i] = strings.TrimSpace(a[i])
	}
	if len(a) < 5 || len(a[1]) != 4 {
		return s
	}
	readmePath := filepath.Join("./../..", "problems", a[1], "README.md")
	dat, err := os.ReadFile(readmePath)
	if err != nil {
		return s
	}
	p := &problem{
		num:   a[1],
		name:  a[2],
		link:  a[3],
		level: a[4],
	}
	lines := strings.Split(string(dat), "\n")
	for _, v := range lines {
		if strings.HasPrefix(v, "前置题目：") {
			p.preNum = strings.TrimPrefix(v, "前置题目：")
		}
		if strings.HasPrefix(v, "前置知识：") {
			p.preKnowl = strings.TrimPrefix(v, "前置知识：")
		}
		if strings.HasPrefix(v, "本题知识：") {
			p.curKnowl = strings.TrimPrefix(v, "本题知识：")
		}
	}
	if p.curKnowl != "" {
		cnt++
	}
	res := fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |", p.num, p.name, p.link, p.level, p.preNum, p.preKnowl, p.curKnowl)
	return res
}
