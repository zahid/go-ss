**Not done**

- [ ] validate incoming requests
    - [ ] Accepts only a GET request and send an HTTP 422 if not
    - [ ] Checks that URL is ??valid?? and send an HTTP 422 if not
    - [ ] Check for exit codes and send an HTTO 5xx if not exit code 0
    - [ ] Return 201 if it worked

- [ ] Unit tests
- [ ] Integration tests

- [ ] Configure "profiles" (and let them be choosable via the api)
    Make some combinations of browser/device/viewport combinations to offer

- [ ] Benchmark
    - [ ] How long for 1 capture?
        - what is the device, browser, os, viewport etc?
    - [ ] How many captures can we handle at once?



**Done**

These should be added to the changelog
- [x] Make a screenshot interface
    - [x] Can plug in different instances of screenshot takers
- [x] Simplify rasterize.js phantomjs program
- [x] Take a screenshot
- [x] Make an HTTP API for taking screenshots

