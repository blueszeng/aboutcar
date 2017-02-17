package task

import (
  "aboutcar/common"
  "aboutcar/config"
  "github.com/robfig/cron"
  "sync"
  "sync/atomic"
)

type CrawlTask struct {
    userName string
    taskName string
    taskExecuteTime string
    cookie string
    status     int          // 运行状态
    stop     chan struct{}
    cron     *cron.Cron
    resCount int32                       // 资源使用情况计数
    resChan  chan bool
    reqs   []common.ReqEntity
	  sync.Mutex
}

func NewCrawlTask(taskName, taskExecuteTime, userName string) *CrawlTask {
  crawlTask := &CrawlTask{
		userName:    userName,
		taskName:    taskName,
    status:      config.STOP,
    resCount:    20,
    resChan:     make(chan bool),
    stop:        make(chan struct{}),
    cron:        cron.New(),
		reqs:        []common.ReqEntity{},
	}
  return crawlTask
}


func (self *CrawlTask) checkStatus(s int) bool {
  self.Lock()
	b := self.status == s
  self.Unlock()
	return b
}


func (self *CrawlTask) GetName() string {
  return self.taskName
}
func (self *CrawlTask) GetCookie() string {
  return self.cookie
}

func (self *CrawlTask) Pull() (req common.ReqEntity) {
  self.Lock()
	defer self.Unlock()
	req = self.reqs[0]
  self.reqs = self.reqs[1:]
  return
}

// 添加请求到队列，并发安全
func (self *CrawlTask) Push(req common.ReqEntity) {
	self.Lock()
	defer self.Unlock()
	atomic.AddInt32(&self.resCount, -1)
  self.reqs = append(self.reqs, req)
  // 分配资源
  if self.resCount > 0 {
    self.resChan <- true
  }
}

func (self *CrawlTask) Use() (resChan chan bool) {
	return self.resChan
}

func (self *CrawlTask) Free() {
	atomic.AddInt32(&self.resCount, 1)
}

// func (self *CrawlTask) Len() int {
// 	self.Lock()
// 	defer self.Unlock()
// 	var l int
// 	for _, reqs := range self.reqs {
// 		l += len(reqs)
// 	}
// 	return l
// }
