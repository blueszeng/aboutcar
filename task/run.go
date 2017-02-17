package task

import (
  "aboutcar/analyze"
  "aboutcar/download"
  "aboutcar/config"
  "log"
  // "time"
)


func (self *CrawlTask) login() {
  login := download.Login{
    Url: "http://dsmis.jishunda.cn/WeiXin/Student/StuLogin.aspx?wx=gh_5151a34f78a2&openId=ojiehs0QIXElIHjPbGR0arPSGu6Q&back=%2fWeiXin%2fStudent%2fHome.aspx%3fcode%3d031JmrLB1nRvkh02cJJB11Y3LB1JmrLf%26state%3dgh_5151a34f78a2",
    UserName: "430521199110034261",
    Password: "034261",
  }
  phantom := download.NewPhantom("/usr/local/bin/phantomjs", "phantom")
  cookie, _ := phantom.Login(login)
  self.cookie = cookie

}
func (self *CrawlTask) init() {
  url := "http://dsmis.jishunda.cn/WeiXin/Student/EduSiteList.aspx?CurrentPage=1&LoadAjaxData=LoadList"
  self.Push(analyze.NewEduSiteReqEntity(url, map[string]string{}))
}
func (self *CrawlTask) Run() {
  log.Println("login")
  self.login()
  go self.init()
  for {
    select {
      case <- self.Use():
          log.Println("task...=--->")
          req := self.Pull()
          go func() {
              document, _ := download.Download(req.GetUrl(), self.cookie)
              req.Analyze(document, self)
              req.SaveData()
              self.Free()
          }()
      case <- self.stop:
        log.Println("stop2")
        // self.cron.Stop()
        return
    }
  }

}

func (self *CrawlTask) Start() {
  // log.Println("v")
  if self.checkStatus(config.RUN) {
    return
  }
  self.status = config.RUN
  // self.cron.AddFunc("1 * * * * * ", func() {
    go self.Run()
  // })
  // self.cron.Start()
}
func (self *CrawlTask) Stop() {
  if self.checkStatus(config.STOP) {
    return
  }
  self.Lock()
	defer self.Unlock()
  log.Println("send-->")
  self.stop <- struct{}{}
	self.status = config.STOP
}
