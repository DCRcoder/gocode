package gocode

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Repository struct {
	path       string
	remotePath string
	lastSync   int64
}

func (repo Repository) Sync() {
	now := time.Now().Unix()
	log.Printf("Sync time: %d", now)
	if repo.lastSync == 0 {
		repo.lastSync = now
		repo.doSync()
	} else {
		if now > repo.lastSync {
			repo.lastSync = now + 1
			repo.doSync()
		}
	}
}

func (repo Repository) doSync() {
	log.Println("do syncing")
	ignoreContext := repo.ignore()
	args := []string{"-avz", "--exclude", ".git"}

	for _, c := range ignoreContext {
		args = append(args, "--exclude", c)
	}
	args = append(args, repo.path, repo.remotePath)

	cmd := exec.Command("rsync", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println(cmd)
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
	log.Println("sync finished")
}

func (repo Repository) ignore() []string {
	f, e := os.Open(repo.path + "/.gitignore")
	defer f.Close()
	if e != nil {
		log.Println("can`t find .gitignore file")
	}

	var excludeContext []string
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		liner := string(scanner.Bytes())
		if strings.HasPrefix(string(liner), "#") || liner == "" {
			continue
		}
		excludeContext = append(excludeContext, liner)
	}
	return excludeContext
}

func New(path string, remotePath string) Repository {
	return Repository{
		path:       path,
		remotePath: remotePath,
	}
}
