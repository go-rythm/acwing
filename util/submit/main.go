package main

import (
	"bytes"
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	font    = "Font Name: Slant"
	comment = `/*
	 *        __         __
	 *       / /_  ___  / /___  ___  __________
	 *      / __ \/ _ \/ / __ \/ _ \/ ___/ ___/
	 *     / / / /  __/ / /_/ /  __/ /  (__  )
	 *    /_/ /_/\___/_/ .___/\___/_/  /____/
	 *                /_/
	 */`
)

var (
	re   = regexp.MustCompile(`^.+\.go$`)
	base = "./../.."
	s    = &strings.Builder{}
	num  string
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile | log.Ltime)
	flag.StringVar(&num, "n", "0001", "题号")
}

func main() {
	flag.Parse()

	// 写入原 solution
	globPath := filepath.Join(base, "problems", num, "*.go")
	solPath, err := filepath.Glob(globPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(solPath) != 1 {
		log.Fatal(solPath)
	}
	sol, err := os.ReadFile(solPath[0])
	if err != nil {
		log.Fatal(err.Error())
	}
	sol = bytes.ReplaceAll(sol, []byte("oj."), []byte{})
	s.Write(sol)

	// 写入 comment
	s.WriteString(comment + "\n\n")

	// 替换 oj 包的相关内容
	ojPath := filepath.Join(base, "/internal/oj")
	err = filepath.WalkDir(ojPath, func(fpath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !re.MatchString(d.Name()) {
			return nil
		}
		filepath := filepath.Join(ojPath, d.Name())
		dat, err := os.ReadFile(filepath)
		if err != nil {
			return err
		}
		lines := strings.Split(string(dat), "\n")
		for i, v := range lines {
			if strings.HasPrefix(v, "type") ||
				strings.HasPrefix(v, "func") ||
				strings.HasPrefix(v, "const") {
				lines = lines[i : len(lines)-1]
				break
			}
		}
		s.WriteString(strings.Join(lines, "\n"))
		s.WriteString("\n\n")
		return nil
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	// 生成可提交的代码
	submitPath := filepath.Join(base, "submit.go")
	sub, err := os.Create(submitPath)
	defer sub.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	sub.WriteString(s.String())
}
