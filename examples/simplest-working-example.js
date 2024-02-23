import pw from 'k6/x/playwright';

export const options = {
  vus: 1,
};

export default function () {
  pw.launch()
  pw.newPage()
  pw.goto("https://www.github.com/", {waitUntil: 'networkidle'})
  pw.waitForSelector(".search-input", {state: 'visible'})
  pw.kill()
}