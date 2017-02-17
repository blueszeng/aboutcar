package main
import (
  "net/http"
  "log"
  "fmt"

  "aboutcar/scheduler"
  "github.com/urfave/negroni"
  "github.com/julienschmidt/httprouter"
)



func AddUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    v := scheduler.AddTask("tttd", "", "", "")
    v.Start()
    fmt.Fprint(w, "sfsdfsdfsdf")
}

func main() {

  router := httprouter.New()
  router.GET("/user/addUser/",AddUser)
  n := negroni.Classic()
  n.UseHandler(router)

  log.Fatal(http.ListenAndServe(":8080", n))


  // login.Login("", "")
  // login := download.Login{
  //   Url: "http://dsmis.jishunda.cn/WeiXin/Student/StuLogin.aspx?wx=gh_5151a34f78a2&openId=ojiehs0QIXElIHjPbGR0arPSGu6Q&back=%2fWeiXin%2fStudent%2fHome.aspx%3fcode%3d031JmrLB1nRvkh02cJJB11Y3LB1JmrLf%26state%3dgh_5151a34f78a2",
  //   UserName: "430521199110034261",
  //   Password: "034261",
  // }
  // phantom := download.NewPhantom("/usr/local/bin/phantomjs", "phantom")
  // cookiea, _ := phantom.Logins(login)
  // fmt.Println(cookiea)
  // req, _ := http.NewRequest("GET", "http://dsmis.jishunda.cn/WeiXin/Student/TrainItemList.aspx?CurrentPage=1&LoadAjaxData=LoadList&id=0.6626654246027404&pageSize=10", nil)
  // req.Header.Set("Set-Cookie", cookie)
  // req.Header.Set("Cookie", "LoginCookie=OpenId=ojiehs0QIXElIHjPbGR0arPSGu6Q&Wx=gh_5151a34f78a2&UserName=430521199110034261&Password=&RememberPwd=0;ASP.NET_SessionId=r2ncnxatn3owbs12hmjzog1d;")

  // url := "http://dsmis.jishunda.cn/WeiXin/Student/EduSiteList.aspx?CurrentPage=1&LoadAjaxData=LoadList&id=0.1484064094300539&pageSize=10"
  // cookie := "LoginCookie=OpenId=ojiehs0QIXElIHjPbGR0arPSGu6Q&Wx=gh_5151a34f78a2&UserName=430521199110034261&Password=&RememberPwd=0;ASP.NET_SessionId=fptx0rownofcu4fn2w51jcdj;"
  // doc, _ := download.Download(url, cookie, analyze.AnalyzeEduSite)
  // // eduSiteId=2F7FF2A1-C46D-446D-9793-140BBA157DFB  //地区
  // fmt.Println(doc)

  // url = "http://dsmis.jishunda.cn/WeiXin/Student/TrainItemList.aspx?CurrentPage=1&LoadAjaxData=LoadList&id=0.6626654246027404&pageSize=10"
  // doc, _ = download.Download(url, cookie, analyze.AnalyzeTrain)
  // fmt.Println(doc)
  //
  // // eduSiteId 地区   //  trainId 项目   // coachId  教练
  // url = "http://dsmis.jishunda.cn/WeiXin/Student/CoachList.aspx?CurrentPage=1&LoadAjaxData=LoadList&
  // &trainId=F66AB04A-A986-4D60-9016-D0D0C855322E&id=0.08414783112968893&pageSize=10"
  // doc, _ = download.Download(url, cookie, analyze.AnalyzeCoach)
  // fmt.Println(doc)
  //
  //
  // url = "http://dsmis.jishunda.cn/WeiXin/Student/ScheduleList.aspx?&LoadAjaxData=LoadDate&coachId=CAF16DBE-3D22-4121-858C-4DE9795BC7F2&id=0.5313873193696194"
  // doc, _ = download.Download(url, cookie, analyze.AnalyzeSchedule)
  // fmt.Println(doc)
  //
  //
  // url = "http://dsmis.jishunda.cn/WeiXin/Student/AddOrder.aspx"
  // doc, _ = download.Download(url, cookie, analyze.AnalyzeHiddenMD5)
  // fmt.Println(doc)

 // date=2017-02-15 日期
// http://dsmis.jishunda.cn/WeiXin/Student/ScheduleList.aspx?&LoadAjaxData=LoadTime&coachId=CAF16DBE-3D22-4121-858C-4DE9795BC7F2&date=2017-02-15&id=0.6976076563891189
  // postUrl := "http://dsmis.jishunda.cn/WeiXin/Student/AddOrder.aspx"
  // download.Commimt(postUrl, cookie)



}
