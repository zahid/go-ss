// phantomjs rasterize.js https://google.com

"use strict";

var page = require('webpage').create();
var system = require('system');

if (typeof system.args[1] === 'undefined') {
    system.stderr.writeLine('No URL given. Use via `phantomjs rasterize.js URL`.')
    phantom.exit(1);
}

var config = {
    url: system.args[1],
    viewportSize: { width: 1024, height: 768 },
    zoomFactor: 0.25,
    wait: 2000
};

page.open(config.url, function (status) {
    page.zoomFactor = config.zoomFactor;
    page.viewportSize = config.viewport;
    if (status !== 'success') {
        system.stderr.writeLine('Something went wrong opening %s', config.url);
        phantom.exit(1);
    } else {
        window.setTimeout(function () {
            system.stdout.write(page.renderBase64('PNG'));
            phantom.exit();
        }, config.wait);
    }
});
