<br>
<div align="center">
   <img src="images/xk6_logo.PNG" width="400" alt="pdq"/><br>
   <h1><b>xk6 playwright</b></h1><br>
   <p>k6 extension that adds support for browser automation and end-to-end web testing using <a href="https://github.com/mxschmitt/playwright-go" target="_blank">playwright-go</a></p>
   <p>Special thanks to all the contributors over at <a href="https://github.com/grafana/k6/graphs/contributors" target="_blank">k6</a> and <a href="https://github.com/mxschmitt/playwright-go/graphs/contributors" target="_blank">playwright-go</a>
   <p>Here's to open source!</p>
   
   <a href="https://github.com/nicholasvuono/xk6-playwright/releases"><img alt="GitHub license" src="https://img.shields.io/badge/release-v0.1.0-blue"></a>
   <a href="https://goreportcard.com/badge/github.com/nicholasvuono/xk6-playwright)"><img src="https://goreportcard.com/badge/github.com/nicholasvuono/xk6-playwright" alt="Go Report Card"></a>
   <a href="https://github.com/nicholasvuono/xk6-playwright/blob/main/LICENSE"><img alt="GitHub license" src="https://img.shields.io/github/license/nicholasvuono/xk6-playwright?color=red"></a>
</div>



----

### Simplest Working Example

```JavaScript
import pw from 'k6/x/playwright';

export default function () {
  pw.launch()
  pw.newPage()
  pw.goto("https://www.google.com/", {waitUntil: 'networkidle'})
  pw.waitForSelector("//html/body/div[1]/div[2]", {state: 'visible'})
  pw.kill()
}
```
