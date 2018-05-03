package gocode

import (
	"log"
	"os"
	"os/exec"
	"time"
)

type Repository struct {
	path       string
	remotePath string
	lastSync   int64
	ignoreRule string
}

func (repo Repository) Sync() {
	now := time.Now().Unix()
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
	cmd := exec.Command("rsync", "-av", "--delete", "--exclude=.git/", repo.path, repo.remotePath)
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

func New(path string, remotePath string) Repository {
	return Repository{
		path:       path,
		remotePath: remotePath,
	}
}

func ValidatePath(path string) bool {
	_, err := os.Stat(path + "/.git")
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
