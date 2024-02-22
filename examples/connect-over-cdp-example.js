import pw from 'k6/x/playwright';

export default function () {
  pw.connect("http://localhost:9222")
  pw.goto("https://www.google.com/", {waitUntil: 'networkidle'})
  pw.waitForSelector("input[title='Search']", {state: 'visible'})
  pw.kill()
}