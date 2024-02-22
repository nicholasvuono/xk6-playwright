import pw from 'k6/x/playwright';

export default function () {
  pw.connect("http://localhost:9222")
  pw.goto("https://www.github.com/", {waitUntil: 'networkidle'})
  pw.waitForSelector(".search-input", {state: 'visible'})
  pw.kill()
}