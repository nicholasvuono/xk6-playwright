import pw from 'k6/x/playwright';



export default function () {
  try {
    pw.connect('http://localhost:9222');
    pw.goto('https://test.k6.io/my_messages.php', { waitUntil: 'networkidle' });
    pw.waitForSelector("input[name='login']", { state: 'visible' });
  } catch (err) {
    console.log(err);
  } finally {
    pw.kill();
  }
}
