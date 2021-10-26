package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

var (
	leftPage  = regexp.MustCompile(`<ahref="([^"]+)">&laquo;</a>`)
	rightPage = regexp.MustCompile(`<ahref="([^"]+)">&raquo;</a>`)
	community = regexp.MustCompile(`<aclass="community_question_title"href="([^"]+)">`)
	pageCode  = regexp.MustCompile(`<ahref="([^"]+)">AcWing`)

	coreURL = "https://www.acwing.com"
	d       = &detail{}
)

type activity struct {
	level     int
	startPage int
	endPage   int
	url       string
}

type Problem struct {
	Num     int
	Name    string
	Content string
	Link    string
	Level   int
	Rel     string
}

type detail struct {
	problems []*Problem
	mu       sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	// 106
	l1 := &activity{
		level:     1,
		startPage: 1,
		endPage:   6,
		url:       "https://www.acwing.com/activity/content/activity_person/content/4102/",
	}
	// 219
	l2 := &activity{
		level:     2,
		startPage: 1,
		endPage:   11,
		url:       "https://www.acwing.com/activity/content/activity_person/content/9226/",
	}
	// 133
	l3 := &activity{
		level:     3,
		startPage: 1,
		endPage:   7,
		url:       "https://www.acwing.com/activity/content/activity_person/content/33298/",
	}
	for _, a := range []*activity{
		l1,
		l2,
		l3,
	} {
		wg.Add(1)
		go func(a *activity) {
			defer wg.Done()
			a.run()
		}(a)
	}
	wg.Wait()

	sort.Slice(d.problems, func(i, j int) bool {
		if d.problems[i].Level < d.problems[j].Level {
			return true
		}
		if d.problems[i].Level > d.problems[j].Level {
			return false
		}
		return d.problems[i].Num < d.problems[j].Num
	})
	res, _ := json.Marshal(d.problems)
	probJson, _ := os.Create("problems.json")
	probJson.Write(res)
	log.Println(len(d.problems))

	f, _ := os.Create("README.md")
	f.WriteString("| # | 标题 | 题解 | 难度 | 知识点 |\n| - | - | - | - | - |\n")
	for _, v := range d.problems {
		var level string
		switch v.Level {
		case 1:
			level = "基础"
		case 2:
			level = "提高"
		case 3:
			level = "进阶"
		default:
			level = "未知"
		}
		one := fmt.Sprintf("| %04d | %s | [Go](problems/%04d-%s) | %s | %s |\n", v.Num, v.Name, v.Num, v.Name, level, "")
		f.WriteString(one)
	}
}

func parseURL(url string, reg *regexp.Regexp) []string {
	res := make([]string, 0)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	html = bytes.ReplaceAll(html, []byte(" "), []byte(""))
	html = bytes.ReplaceAll(html, []byte("\n"), []byte(""))
	submatch := reg.FindAllSubmatch(html, -1)
	for _, v := range submatch {
		res = append(res, string(v[1]))
	}
	return res
}

func newProblem(url string, level int) (p *Problem) {
	p = &Problem{
		Link:  url,
		Level: level,
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{Name: "sessionid", Value: "secret"})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(string(html))
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(markdown, "使用激活码") {
		p.Name = "查无此题"
		return
	}
	markdown = strings.ReplaceAll(markdown, "\n```\n\n", "```\n\n")
	markdown = strings.ReplaceAll(markdown, "#### ", "### ")
	slice := strings.SplitAfter(markdown, "\n")
	fullName := strings.Replace(slice[56], "\n", "", -1)
	title := strings.Split(fullName, "\\. ")
	p.Num, _ = strconv.Atoi(title[0])
	p.Name = title[1]
	p.Name = strings.Replace(p.Name, " ", "", -1)
	slice = slice[65:]
	content := strings.Join(slice, "")
	contentslice := strings.Split(content, "\n难度：\n")
	content = contentslice[0]
	p.Content = fmt.Sprintf("## [%s](%s)\n\n### 题目\n\n%s\n### 题解\n\n", fullName, url, content)
	return
}

func (p *Problem) save() {
	dir := fmt.Sprintf("problems/%04d-%s", p.Num, p.Name)
	os.MkdirAll(dir, 0700)
	codeDir := fmt.Sprintf("problems/%04d-%s/%04d-%s.go", p.Num, p.Name, p.Num, p.Name)
	fCode, _ := os.Create(codeDir)
	fCode.WriteString("package main\n\nfunc main() {\n\n}\n")
	fCode.Close()
	pDir := fmt.Sprintf("problems/%04d-%s/README.md", p.Num, p.Name)
	fP, _ := os.Create(pDir)
	fP.WriteString(p.Content)
	fP.Close()
}

func (this *activity) run() {
	var wg sync.WaitGroup
	for i := this.startPage; i <= this.endPage; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			u1 := this.url + fmt.Sprint(i)
			res := parseURL(u1, community)
			for _, v := range res {
				wg.Add(1)
				go func(v string) {
					defer wg.Done()
					final := parseURL(coreURL+v, pageCode)
					// log.Println(final[0])
					p := newProblem(final[0], this.level)
					p.save()
					p.Content = ""
					d.mu.Lock()
					d.problems = append(d.problems, p)
					d.mu.Unlock()
				}(v)
			}
		}(i)
	}
	wg.Wait()
}
