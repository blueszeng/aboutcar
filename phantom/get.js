
var system = require('system');
var page = require('webpage').create();
var url = system.args[1];
var cookie = system.args[2];
var pageEncode = system.args[3];
var userAgent = system.args[4];

// phantom.libraryPath= "."
// phantom.injectJs('jquery-1.4.4.min.js')

page.onResourceRequested = function(requestData, request) {
    request.setHeader('Cookie', cookie)
    console.log("55555555")
};
phantom.outputEncoding = pageEncode;
page.settings.userAgent = userAgent;

page.open(url, function(status) {
  // console.log("23423423423")

    if (status !== 'success') {
        console.log('Unable to access network');
    } else {
       // 	var cookie = page.evaluate(function() {
        //     return true
        // });
        var resp = {
            "Cookie": page.cookies,
            "Body": page.content
        };
        console.log(JSON.stringify(resp));
    }
    phantom.exit();
});
