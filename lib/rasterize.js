// phantomjs rasterize.js https://google.com out/1.png

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
    wait: 1000
};

page.open(config.url, function (status) {
    page.zoomFactor = config.zoomFactor;
    page.viewportSize = config.viewport;
    if (status !== 'success') {
        phantom.exit(1);
    } else {
        window.setTimeout(function () {
            system.stdout.write(page.renderBase64());
            phantom.exit();
        }, config.wait);
    }
});
