# README

## About

This is the official non-official project of me.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## TODO

Link in the settings to get api information, maybe even a dropdown with a guide so people are not confused
Chart data 
Chart zooming and panning
Value lines for the chart
PlayerList adds from right to left, bottom to top
If gamemode is changed and stats already exist, prompt user "are you sure" and then rerun getstats with new gamemode
Different themes (was using a ui library a good idea?)
Customization of chart theme?
Clean up some backend and work on sending errors to the front
Warning card can only be called once???