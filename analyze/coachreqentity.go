package analyze

import (
	"aboutcar/common"
	. "aboutcar/rest/db"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type (
	Coach struct {
		TrainId string `xorm:"varchar(255) index not null unique 'train_id'"`
		Uuid    string `xorm:"varchar(255) index not null unique 'uuid'"`
		Name    string `xorm:"varchar(255)"`
		Phone   string `xorm:"varchar(12)"`
	}
	CoachReqEntity struct {
		Url    string
		Temp   map[string]string
		Coachs []Coach
	}
)

// 培训项目
func NewCoachReqEntity(url string, temp map[string]string) *CoachReqEntity {
	coachReq := &CoachReqEntity{
		Url:  url,
		Temp: temp,
	}
	return coachReq
}

func (self *CoachReqEntity) GetUrl() string {
	return self.Url
}

func (self *CoachReqEntity) Analyze(document *goquery.Document, task common.Task) {

	var trainId string
	var err bool
	if trainId, err = self.Temp["trainId"]; err != true {
		log.Fatalln("trainId not is null")
	}
	document.Find("#tab tr").Each(func(i int, contentSelectionTr *goquery.Selection) {
		uuid, _ := contentSelectionTr.Find("td:first-child input").Attr("id")
		name := contentSelectionTr.Find("td:nth-child(2)").Text()
		name = strings.TrimSpace(name)
		tempName := strings.Split(name, "：")[1]
		name = strings.TrimSuffix(tempName, "电话")
		phone := contentSelectionTr.Find("td:nth-child(2) a").Text()
		if strings.HasPrefix(uuid, "rad") {
			uuid = uuid[3:]
		}
		coach := Coach{
			TrainId: trainId,
			Uuid:    uuid,
			Name:    name,
			Phone:   phone,
		}
		self.Coachs = append(self.Coachs, coach)
	})
	log.Println(self.Coachs)
}
func (self *CoachReqEntity) SaveData() {
	var err error
	err = DB.Sync2(new(Coach))
	_, err = DB.Insert(&self.Coachs)
	log.Println(err)
}
