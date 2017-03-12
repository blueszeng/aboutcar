package download

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	// "aboutcar/config"
)

type (
	Phantom struct {
		PhantomjsFile string            //Phantomjs完整文件名
		TempJsDir     string            //临时js存放目录
		jsFileMap     map[string]string //已存在的js文件
	}
	Response struct {
		Cookie string
		Body   string
	}
	Login struct {
		UserName string
		Password string
		Url      string
	}
	Cookie struct {
		// Domain string `json:"domain"`
		// Expires string  `json:"expires"`
		// Expiry int `json:"expiry"`
		// Httponly bool `json:"httponly"`
		Name string `json:"name"`
		// Path  string `json:"path"`
		// Secure  string `json:"secure"`
		Value string `json:"value"`
	}
	Cookielice struct {
		Cookies []Cookie `json:"Cookie"`
	}
)

func (self *Phantom) Login(login Login) (cookie string, err error) {
	var encoding = "utf-8"
	var args = []string{
		self.jsFileMap["login"],
		login.Url,
		encoding,
		login.UserName,
		login.Password,
	}
	cmd := exec.Command(self.PhantomjsFile, args...)
	var body io.Reader
	if body, err = cmd.StdoutPipe(); err != nil {
		log.Fatal(err)
	}
	if cmd.Start() != nil || body == nil {
		log.Fatal(err)
	}
	var b []byte
	b, err = ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}

	b = bytes.TrimSpace(b)
	len := bytes.Index(b, []byte{109, 115, 46, 10}) + 4
	b = b[len:]
	fmt.Println(string(b))
	cookies := Cookielice{}
	err = json.Unmarshal(b, &cookies)
	for _, ccookie := range cookies.Cookies {
		cookie += fmt.Sprintf("%s=%s;", ccookie.Name, ccookie.Value)
	}
	log.Println(cookies)
	return
}

func (self *Phantom) Downloads(req *http.Request) (resp *http.Response, err error) {
	encoding := "utf-8"
	var args []string
	url := fmt.Sprintf("%s", req.URL)
	body := fmt.Sprintf("%s", req.Body)
	switch req.Method {
	case "GET":
		args = []string{
			self.jsFileMap["get"],
			url,
			req.Header.Get("Cookie"),
			encoding,
			req.Header.Get("User-Agent"),
		}
	case "POST":
		args = []string{
			self.jsFileMap["post"],
			url,
			req.Header.Get("Cookie"),
			encoding,
			req.Header.Get("User-Agent"),
			body,
		}
	}
	fmt.Println(args)
	resp = new(http.Response)
	cmd := exec.Command(self.PhantomjsFile, args...)
	if resp.Body, err = cmd.StdoutPipe(); err != nil {
		log.Fatal(resp.Body)
	}
	if cmd.Start() != nil || resp.Body == nil {
		log.Fatal(err)
	}
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	retResp := Response{}

	s := `{"Cookie":"","Body":"<html><head><script language=\"javascript\"> \n "}`

	err = json.Unmarshal([]byte(s), &retResp)
	if err != nil {
		log.Fatal(err)
	}
	resp.Header = req.Header
	resp.Header.Set("Set-Cookie", retResp.Cookie)
	resp.Body = ioutil.NopCloser(strings.NewReader(retResp.Body))

	if err == nil {
		resp.StatusCode = http.StatusOK
		resp.Status = http.StatusText(http.StatusOK)
	} else {
		resp.StatusCode = http.StatusBadGateway
		resp.Status = http.StatusText(http.StatusBadGateway)
	}
	return
}

//销毁js临时文件
// func (self *Phantom) DestroyJsFiles() {
// 	p, _ := filepath.Split(self.TempJsDir)
// 	if p == "" {
// 		return
// 	}
// 	for _, filename := range self.jsFileMap {
// 		os.Remove(filename)
// 	}
// 	if len(WalkDir(p)) == 1 {
// 		os.Remove(p)
// 	}
// }

func (self *Phantom) createJsFile(fileName, jsCode string) {
	fullFileName := filepath.Join(self.TempJsDir, fileName)
	// 创建并写入文件
	f, _ := os.Create(fullFileName)
	f.Write([]byte(jsCode))
	f.Close()
	self.jsFileMap[fileName] = fullFileName
}

func NewPhantom(phantomjsFile, tempJsDir string) *Phantom {
	phantom := &Phantom{
		PhantomjsFile: phantomjsFile,
		TempJsDir:     tempJsDir,
		jsFileMap:     make(map[string]string),
	}
	if !filepath.IsAbs(phantom.PhantomjsFile) {
		phantom.PhantomjsFile, _ = filepath.Abs(phantom.PhantomjsFile)
	}
	if !filepath.IsAbs(phantom.TempJsDir) {
		phantom.TempJsDir, _ = filepath.Abs(phantom.TempJsDir)
	}
	// 创建/打开目录
	err := os.MkdirAll(phantom.TempJsDir, 0777)
	if err != nil {
		log.Printf("[E] Surfer: %v\n", err)
		return phantom
	}
	phantom.createJsFile("get", getJs)
	phantom.createJsFile("post", postJs)
	phantom.createJsFile("login", loginJs)
	return phantom
}

/*
* GET method
* system.args[0] == get.js
* system.args[1] == url
* system.args[2] == cookie
* system.args[3] == pageEncode
* system.args[4] == userAgent
 */

const getJs string = `
var system = require('system');
var page = require('webpage').create();
var url = system.args[1];
var cookie = system.args[2];
var pageEncode = system.args[3];
var userAgent = system.args[4];
page.onResourceRequested = function(requestData, request) {
    request.setHeader('Cookie', cookie)
};
phantom.outputEncoding = pageEncode;
page.settings.userAgent = userAgent;
page.open(url, function(status) {
    if (status !== 'success') {
        console.log('Unable to access network');
    } else {
       	var cookie = page.evaluate(function(s) {
            return document.cookie;
        });
        var resp = {
            "Cookie": cookie,
            "Body": page.content
        };
        console.log(JSON.stringify(resp));
    }
    phantom.exit();
});
`

/*
* POST method
* system.args[0] == post.js
* system.args[1] == url
* system.args[2] == cookie
* system.args[3] == pageEncode
* system.args[4] == userAgent
* system.args[5] == postdata
 */
const postJs string = `
var system = require('system');
var page = require('webpage').create();
var url = system.args[1];
var cookie = system.args[2];
var pageEncode = system.args[3];
var userAgent = system.args[4];
var postdata = system.args[5];
page.onResourceRequested = function(requestData, request) {
    request.setHeader('Cookie', cookie)
};
phantom.outputEncoding = pageEncode;
page.settings.userAgent = userAgent;
page.open(url, 'post', postdata, function(status) {
    if (status !== 'success') {
        console.log('Unable to access network');
    } else {
        var cookie = page.evaluate(function(s) {
            return document.cookie;
        });
        var resp = {
            "Cookie": cookie,
            "Body": page.content
        };
        console.log(JSON.stringify(resp));
    }
    phantom.exit();
});
`

const loginJs string = `
var system = require('system');
var page = require('webpage').create();
var url = system.args[1];
var pageEncode = system.args[2];
var userName = system.args[3];
var passWord = system.args[4];
var userAgent = system.args[5];
var cookie = system.args[6];
var postdata = system.args[7];
var waitFor = function (testFx, onReady, timeOutMillis) {
    var maxtimeOutMillis = timeOutMillis ? timeOutMillis : 4000,
        start = new Date().getTime(),
        condition = false,
        interval = setInterval(function() {
            if ( (new Date().getTime() - start < maxtimeOutMillis) && !condition ) {
                condition = (typeof(testFx) === "string" ? eval(testFx) : testFx());
            } else {
                if(!condition) {
                    console.log("'waitFor()' timeout");
                    phantom.exit(1);
                } else {
                    console.log("'waitFor()' finished in " + (new Date().getTime() - start) + "ms.");
                    typeof(onReady) === "string" ? eval(onReady) : onReady();
                    clearInterval(interval);
                }
            }
        }, 250);
};
page.onResourceRequested = function(requestData, request) {
    request.setHeader('Cookie', cookie)
};
phantom.outputEncoding = "utf-8";
page.settings.userAgent = userAgent;
page.open(url, 'post', postdata, function(status) {
    if (status !== 'success') {
        console.log('Unable to access network');
    } else {
        waitFor(function() {
          return page.evaluate(function() {
            $("#txtUserName")[0].value = "430521199110034261"
            $("#txtPassword")[0].value = "034261"
            $("#btnBinding").click()
            return document.cookie;
          });
        }, function(err) {

          var cookies = page.cookies
          for(var i in cookies) {
              delete cookies[i].path
              delete cookies[i].domain
              delete cookies[i].expires
              delete cookies[i].expiry
              delete cookies[i].httponly
              delete cookies[i].secure
              // delete cookies[i].value
              // delete cookies[i].name

          }
          var resp = {
              "Cookie": cookies
          };
          page.render('./ee.png')
          console.log(JSON.stringify(resp));
          phantom.exit();
        });
    }
});
`
