import "./styles/app.css";
import App from "./App.svelte";

async function startApp() {
  // Wails injects its bridge on / and /index.html, but a history-route refresh
  // (such as /chart) goes straight to Vite's fallback HTML without those scripts.
  if (!window['go']) {
    for (const src of ['/wails/ipc.js', '/wails/runtime.js']) {
      await new Promise<void>((resolve, reject) => {
        const script = document.createElement('script');
        script.src = src;
        script.onload = () => resolve();
        script.onerror = () => reject(new Error('Could not load the application bridge.'));
        document.head.appendChild(script);
      });
    }
  }
  return new App({ target: document.getElementById('app') });
}

export default startApp().catch(() => {
  document.getElementById('app').textContent = 'Could not connect to the application. Please reopen osu!Chart.';
});
