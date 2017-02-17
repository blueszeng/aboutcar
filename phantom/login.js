
console.log("sbsb")
var system = require('system');
var page = require('webpage').create();
var url = system.args[1];
var pageEncode = system.args[2];
var userName = system.args[3];
var passWord = system.args[4];
var userAgent = system.args[5];
var cookie = system.args[6];
var postdata = system.args[7];

url = "http://dsmis.jishunda.cn/WeiXin/Student/StuLogin.aspx?wx=gh_5151a34f78a2&openId=ojiehs0QIXElIHjPbGR0arPSGu6Q&back=%2fWeiXin%2fStudent%2fHome.aspx%3fcode%3d031JmrLB1nRvkh02cJJB11Y3LB1JmrLf%26state%3dgh_5151a34f78a2"
pageEncode = "utf-8"
console.log(url)


var waitFor = function (testFx, onReady, timeOutMillis) {
    var maxtimeOutMillis = timeOutMillis ? timeOutMillis : 3000, //< Default Max Timout is 3s
        start = new Date().getTime(),
        condition = false,
        interval = setInterval(function() {
            if ( (new Date().getTime() - start < maxtimeOutMillis) && !condition ) {
                // If not time-out yet and condition not yet fulfilled
                condition = (typeof(testFx) === "string" ? eval(testFx) : testFx()); //< defensive code
            } else {
                if(!condition) {
                    // If condition still not fulfilled (timeout but condition is 'false')
                    console.log("'waitFor()' timeout");
                    phantom.exit(1);
                } else {
                    // Condition fulfilled (timeout and/or condition is 'true')
                    console.log("'waitFor()' finished in " + (new Date().getTime() - start) + "ms.");
                    typeof(onReady) === "string" ? eval(onReady) : onReady(); //< Do what it's supposed to do once the condition is fulfilled
                    clearInterval(interval); //< Stop this interval
                }
            }
        }, 250); //< repeat check every 250ms
};
page.onResourceRequested = function(requestData, request) {
    request.setHeader('Cookie', cookie)
};
phantom.outputEncoding = pageEncode;
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
            return true;
          });
        }, function(err) {
          var resp = {
              "Cookie": page.cookies,
              "Body": page.content
          };
          console.log(JSON.stringify(resp));
          phantom.exit();
        })
    }
});
