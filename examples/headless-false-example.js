import pw from 'k6/x/playwright';

export default function () {
  pw.launch({headless: false})
  pw.newPage()
  pw.goto("https://www.github.com/", { waitUntil: 'networkidle'})
  pw.waitForSelector(".search-input", {state: 'visible'})
  pw.kill()
}