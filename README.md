# README

## About
osu-chart is an application that allows you to input the usernames of osu! players and view their stats over time from the osutrack API. The main use is comparing different players’ stats against each other. You can choose from various metrics such as pp, rank, accuracy, playcount, and more. 

I mainly made this because when I made YouTube videos, the charts on https://ameobea.me/osutrack/ were only available in light mode. Soon you will be able to customize the chart appearance, such as the color scheme, and the time range.

- Compare multiple players on the same chart
- Choose from different game modes and metrics
- Dark mode
- Responsive design
- Queue osu!track update when a player is entered

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

This project was made with Svelte
## Building

To build a redistributable, production mode package, use `wails build`.

## TODO
- Customize the chart appearance and settings
- Fix chart zoom and pan
- Different themes (was using a ui library a good idea?)