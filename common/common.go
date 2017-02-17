package common

import "github.com/PuerkitoBio/goquery"
type (
  Task interface {
    GetName() string
    GetCookie() string
    Push(req ReqEntity)
    Pull() ReqEntity
    Start()
    Stop()
  }
  ReqEntity interface {
    GetUrl() string
    Analyze(*goquery.Document, Task)
    SaveData()
  }
)
