import pw from 'k6/x/playwright';

export default function () {
  try {
    pw.launch()
    pw.newPage()
    pw.goto('https://test.k6.io/', { waitUntil: 'networkidle' });
    pw.waitForSelector("input[name='login']", { state: 'visible' });
    pw.type("input[name='login']", "admin")

    //print out real user metrics of the google search page
    console.log(`First Paint: ${pw.firstPaint()}ms`)
    console.log(`First Contentful Paint: ${pw.firstContentfulPaint()}ms`)
    console.log(`Time to Minimally Interactive: ${pw.timeToMinimallyInteractive()}ms`)
    console.log(`First Input Delay: ${pw.firstInputDelay()}ms`)
  } catch (err) {
    console.log(err);
  } finally {
    pw.kill();
  }
}