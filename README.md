# REI3 client
This repo hosts the REI3 fat client. It is packaged and released with official versions of [REI3](https://github.com/r3-team/r3).

## Features
* Connect to one or many REI3 instances via websocket and authentication token.
* Receive requests for local file handling from REI3 via websocket.
* React to changes to local files and upload new file versions on change.
* Remember cached files between sessions.
* Cleanup cached files when they are not used anymore.
* Auto-install itself in the user directory and to startup if desired.
