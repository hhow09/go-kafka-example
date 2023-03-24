package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Env struct {
	OutDir    string
	LogFile   string
	RemoveRaw bool
}

func MustGetEnv(s string) string {
	v := os.Getenv(s)
	if v == "" {
		log.Fatalf("no %s", s)
	}
	return v
}

func getEnv() Env {
	outDir := MustGetEnv("OUT_DIR")
	logFile := MustGetEnv("LOG_FILE")
	removeRaw := MustGetEnv("REMOVE_RAW")

	return Env{OutDir: outDir, LogFile: logFile, RemoveRaw: removeRaw == "true"}
}

func getLogFiles(env Env) ([]string, error) {
	var files []string

	folder := env.OutDir

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.HasPrefix(info.Name(), env.LogFile) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func parseFile(file string) (int, int, int) {
	count, startTime, endTime := 0, -1, -1

	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			break
		}
		if len(strings.Split(line, " ")) < 2 {
			break
		}
		count += 1
		time, _ := strconv.Atoi(strings.Split(line, " ")[0])
		if startTime == -1 {
			startTime = time
		}
		endTime = time
	}
	return count, startTime, endTime
}

func main() {
	env := getEnv()
	files, err := getLogFiles(env)
	if err != nil {
		log.Fatal(err)
	}
	result, err := os.Create(env.OutDir + env.LogFile)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()
	count := 0
	max_duration := 0
	for _, file := range files {
		count_, startTime, endTime := parseFile(file)
		count += count_
		if endTime-startTime > max_duration {
			max_duration = endTime - startTime
		}
	}
	res := fmt.Sprintf("received %d messages. \ntotal time consumption %d seconds.", count, max_duration)
	result.WriteString(res)

	if env.RemoveRaw {
		for _, f := range files {
			os.Remove(f)
		}
	}
}
