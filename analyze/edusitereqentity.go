package analyze
import (
   "github.com/PuerkitoBio/goquery"
   "aboutcar/common"
   . "aboutcar/rest/db"
   "log"
  //  "fmt"
   "strings"
)
type (
 Edusite struct {
    Uuid string `xorm:"varchar(255) index not null unique 'uuid'"`
    County string `xorm:"varchar(255)"`
    TrainName string `xorm:"varchar(255)"`
  }
  EduSiteReqEntity struct {
    Url string
    Temp map[string] string
    Edusites []Edusite

  }
)

// 地址
func NewEduSiteReqEntity(url string, temp map[string] string) *EduSiteReqEntity {
   eduSiteReq := &EduSiteReqEntity{
      Url: url,
      Temp: temp,
   }
   return eduSiteReq
}

func(self *EduSiteReqEntity) GetUrl() string {
  return self.Url
}

func(self *EduSiteReqEntity) Analyze(document *goquery.Document, task common.Task) {
    document.Find("#tab tr").Each(func(i int, contentSelectionTr *goquery.Selection) {
          uuid, _ := contentSelectionTr.Find("td:first-child input").Attr("id")
          county := contentSelectionTr.Find("td:nth-child(2)").Text()
          trainName := contentSelectionTr.Find("td:nth-child(3)").Text()
          if strings.HasPrefix(uuid, "rad") {
            uuid = uuid[3:]
          }
          data := Edusite{
            Uuid: uuid,
            County: county,
            TrainName: trainName,
          }
          self.Edusites = append(self.Edusites, data)
          // url := fmt.Sprintf("http://dsmis.jishunda.cn/WeiXin/Student/TrainItemList.aspx?CurrentPage=1&LoadAjaxData=LoadList&trainId=%s", uuid)
          // temp := map[string]string{
          //   "trainId" : uuid,
          // }
          // task.Push(NewCoachReqEntity(url, temp))
      })

}
func(self *EduSiteReqEntity) SaveData() {
  var err error
  err = DB.Sync2(new(Edusite))
  _, err = DB.Insert(&self.Edusites)
  log.Println(err)
}

// //教练
// func AnalyzeCoach(document *goquery.Document) (date *Data) {
//   date = &Data{}
//   document.Find("#tab tr").Each(func(i int, contentSelectionTr *goquery.Selection) {
//         id, _ := contentSelectionTr.Find("td:first-child input").Attr("id")
//         if strings.HasPrefix(id, "rad") {
//           id = id[3:]
//         }
//         date.Id = append(date.Id, id)
//     })
//   return
// }
//
// // 时段
// func AnalyzeSchedule(document *goquery.Document) (date *Data) {
//   date = &Data{}
//   document.Find("body div").Each(func(i int, contentSelectionDiv *goquery.Selection) {
//         id, _ := contentSelectionDiv.Attr("id")
//         date.Id = append(date.Id, id)
//     })
//   return
// }
//
//
// // 时段
// func AnalyzeHiddenMD5(document *goquery.Document) (date *Data) {
//   date = &Data{}
//   value, _ := document.Find(".aspNetHidden #__VIEWSTATE").Attr("value")
//   date.Id = append(date.Id, value)
//   value, _ = document.Find(".aspNetHidden #__EVENTVALIDATION").Attr("value")
//   date.Id = append(date.Id, value)
//   return
// }
