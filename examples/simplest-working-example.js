import pw from 'k6/x/playwright';

export default function () {
  try {
    pw.launch();
    pw.newPage();
    pw.goto('https://test.k6.io/', { waitUntil: 'networkidle' });
    pw.waitForSelector("input[name='login']", { state: 'visible' });
  } catch (err) {
    console.log(err);
  } finally {
    pw.kill();
  }
}
