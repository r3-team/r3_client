# REI3 client
This repo hosts the currently unreleased [REI3](https://github.com/r3-team/r3) fat client - it is currently work-in-progress.

The fat client will be part of a future REI3 release. It will serve to extend the capabilities of REI3 on local devices - primarily local file handling.

## Currently planned functions
* Connect to REI3 instances via websocket and authentication token.
* Receive requests for local file handling from the REI3 web interface via websocket.
* React to changes to local files and upload new file versions on change.
* Remember cached files between sessions.
* Cleanup cached files when they are not used anymore.
* Auto-install itself in the user directory and to startup if desired.
