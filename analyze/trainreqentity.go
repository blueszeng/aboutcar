package analyze

import (
   "github.com/PuerkitoBio/goquery"
    . "aboutcar/rest/db"
   "aboutcar/common"
   "log"
   "strings"
)
type (
  Train struct {
    Uuid string `xorm:"varchar(255) index not null unique 'uuid'"`
    ProjectName  string `xorm:"varchar(255)"`
    Price string `xorm:"varchar(255)"`
  }
  TrainReqEntity struct {
    Url string
    Temp map[string] string
    Trains []Train
  }
)


// 培训项目
func NewTrainReqEntity(url string, temp map[string] string) *TrainReqEntity {
   trainReq := &TrainReqEntity{
      Url: url,
      Temp: temp,
   }
   return trainReq
}

func(self *TrainReqEntity) GetUrl() string {
  return self.Url
}

func(self *TrainReqEntity) Analyze(document *goquery.Document, task common.Task) {
    document.Find("#tab tr").Each(func(i int, contentSelectionTr *goquery.Selection) {
          uuid, _ := contentSelectionTr.Find("td:first-child input").Attr("id")
          projectName := contentSelectionTr.Find("td:nth-child(2)").Text()
          price := contentSelectionTr.Find("td:nth-child(3)").Text()
          if strings.HasPrefix(uuid, "rad") {
            uuid = uuid[3:]
          }
          train := Train{
            Uuid: uuid,
            ProjectName: projectName,
            Price: price,
          }
          self.Trains = append(self.Trains, train)
      })

      // url := "http://dsmis.jishunda.cn/WeiXin/Student/CoachList.aspx?CurrentPage=1&LoadAjaxData=LoadList"
      // task.Push(NewCoachReqEntity(url))

}
func(self *TrainReqEntity) SaveData() {
  var err error
  err = DB.Sync2(new(Train))
  _, err = DB.Insert(&self.Trains)
  log.Println(err)
}
