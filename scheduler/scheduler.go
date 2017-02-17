package scheduler

import (
  	"sync"
    "aboutcar/common"
    "aboutcar/config"
    "aboutcar/task"
    "log"
)
type scheduler struct {
  status       int          // 运行状态
	count        chan bool    // 总并发量计数
  tasks         []common.Task
  sync.RWMutex              // 全局读写锁
}

// 定义全局调度
var sdl = &scheduler{
	status: config.RUN,
	count:  make(chan bool, config.ThreadNum),
}

func Init() {
  sdl.tasks = []common.Task{}
  sdl.status = config.RUN
}

func AddTask(taskName, taskType, taskExecuteTime string, userName string) common.Task {
	newTask := task.NewCrawlTask(taskName, taskExecuteTime, userName)
	sdl.RLock()
	defer sdl.RUnlock()
	sdl.tasks = append(sdl.tasks, newTask)
  log.Println("v")
	return newTask
}

func RemoveTask(taskName string) bool {
	sdl.RLock()
	defer sdl.RUnlock()
  delIndex := -1
  for index, newTask := range sdl.tasks {
    if newTask.GetName() == taskName {
      delIndex = index
      break
    }
  }
  if delIndex != -1 {
    sdl.tasks = append(sdl.tasks[0:delIndex], sdl.tasks[delIndex+1:]...)
  }
  return true
}

func PauseRecover() {
	sdl.Lock()
	defer sdl.Unlock()
	switch sdl.status {
	case config.PAUSE:
		sdl.status = config.RUN
	case config.RUN:
		sdl.status = config.PAUSE
	}
}

// 终止任务
func Stop() {
	sdl.Lock()
	defer sdl.Unlock()
	sdl.status = config.STOP
	// 清空
	defer func() {
		recover()
	}()
	close(sdl.count)
	sdl.tasks = []common.Task{}
}

func (self *scheduler) checkStatus(s int) bool {
	self.RLock()
	b := self.status == s
	self.RUnlock()
	return b
}
