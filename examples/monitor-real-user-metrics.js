import pw from 'k6/x/playwright';

export default function () {
  pw.launch()
  pw.newPage()
  pw.goto("https://www.google.com/")
  pw.waitForSelector("input[title='Search']", {state: 'visible'})
  pw.type("input[title='Search']", "how to measure real user metrics with the xk6-playwright extension for k6?")

  //print out real user metrics of the google search page
  console.log(`First Paint: ${pw.firstPaint()}ms`)
  console.log(`First Contentful Paint: ${pw.firstContentfulPaint()}ms`)
  console.log(`Time to Minimally Interactive: ${pw.timeToMinimallyInteractive()}ms`)
  console.log(`First Input Delay: ${pw.firstInputDelay()}ms`)

  pw.kill()
}