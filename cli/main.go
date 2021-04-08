package main

import (
	"github.com/pkg/errors"
	"github.com/prometheus/tsdb/fileutil"
	"github/bsc-task/cli/command"
	"github/bsc-task/log"
	"path/filepath"
)

func main() {
	instanceLock := lockDataDir()
	defer lockRelease(instanceLock)
	log.Init()
	command.Execute()
}

func lockDataDir() fileutil.Releaser {
	lock, _, err := fileutil.Flock(filepath.Join("LOCK"))
	if err != nil {
		panic(errors.Wrapf(err, "启动失败：程序正在执行或上次执行未正常结束").Error())
	}
	return lock
}

func lockRelease(instanceLock fileutil.Releaser) {
	err := instanceLock.Release()
	if err != nil {
		panic(err)
	}
}
