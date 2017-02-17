package download

import (
  "net/http"
  "github.com/PuerkitoBio/goquery"
  "net/url"
  "io/ioutil"
  "strings"
  "log"
  // "fmt"
)

func Download(url, cookie string) (document *goquery.Document, err error) {
    client := &http.Client{}
    requst, err := http.NewRequest("GET", url, nil)
    if err != nil {
      log.Fatal(err)
    }
    requst.Header.Set("Cookie", cookie)
    requst.Header.Set("User-Agent",
      "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)")
    resp, err := client.Do(requst)
    if err != nil {
      log.Fatal(err)
    }
    defer resp.Body.Close()
    document, err = goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
      log.Fatal(err)
    }
    // data = analyze(document)
    return
}


// __VIEWSTATE	/wEPDwULLTE1MTc1Njc5MzJkZP+QaPM9triDAIEKYI6VoVALi9nopJJE1j5vbaEpv5sG
// __EVENTVALIDATION	/wEWDQK1zv26DgKP0OeuBgK9mbPjAwKA+7SWBAKi2b28AgLlicSCDgL5h+3dBgKct7iSDAL4p/7IAgKJ3pHHDQLypMGrDQKP3rHHDQKv3bT4B5mdIet/5ACKPZpHkluXYXRCgydVKoN/IFH0BZfBU3Z8
// hidEduSiteId	2F7FF2A1-C46D-446D-9793-140BBA157DFB
// hidTrainId	F66AB04A-A986-4D60-9016-D0D0C855322E
// hidCoachId	CAF16DBE-3D22-4121-858C-4DE9795BC7F2
// hidScheduleId	4953144；4953145
// hidFeesType	2
// hidPrice	0.00
// btnSave	提交
// hidEduSiteName	西乡
// hidTrainName	科目二训练
// hidCoachName	杨健
// hidTrainTime	2017-02-15 13:00-14:00
// hidOrderMinMinute	60

func Commimt(pUrl, cookie string) {
  v := url.Values{}
  v.Set("__VIEWSTATE", "/wEPDwULLTE1MTc1Njc5MzJkZP+QaPM9triDAIEKYI6VoVALi9nopJJE1j5vbaEpv5sG")
  v.Set("__EVENTVALIDATION", "/wEWDQK1zv26DgKP0OeuBgK9mbPjAwKA+7SWBAKi2b28AgLlicSCDgL5h+3dBgKct7iSDAL4p/7IAgKJ3pHHDQLypMGrDQKP3rHHDQKv3bT4B5mdIet/5ACKPZpHkluXYXRCgydVKoN/IFH0BZfBU3Z8")
  v.Set("hidEduSiteId", "2F7FF2A1-C46D-446D-9793-140BBA157DFB")
  v.Set("hidTrainId", "F66AB04A-A986-4D60-9016-D0D0C855322EB")
  v.Set("hidCoachId", "CAF16DBE-3D22-4121-858C-4DE9795BC7F2")
  v.Set("hidScheduleId", "4953144；4953145")
  // v.Set("hidFeesType", "2")
  // v.Set("hidPrice", "0.00")
  v.Set("btnSave", "提交")
  v.Set("hidEduSiteName", "西乡")
  v.Set("hidTrainName", "科目二训练")
  v.Set("hidCoachName", "杨健")
  v.Set("hidTrainTime", "2017-02-15 13:00-14:00")
  v.Set("hidOrderMinMinute", "60")
  body := ioutil.NopCloser(strings.NewReader(v.Encode()))
  client := &http.Client{}
  requst, err := http.NewRequest("POST", pUrl, body)
  if err != nil {
    log.Fatal(err)
  }
  requst.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  requst.Header.Set("Cookie", cookie)
  requst.Header.Set("User-Agent",
    "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)")
  resp, err := client.Do(requst)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
}
