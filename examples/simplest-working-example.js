import pw from 'k6/x/playwright';

export default function () {
  pw.launch()
  pw.newPage()
  pw.goto("https://www.google.com/", {waitUntil: 'networkidle'})
  pw.waitForSelector("input[title='Search']", {state: 'visible'})
  pw.kill()
}