# REI3 client
This repo hosts the REI3 fat client. It is packaged and released with official versions of [REI3](https://github.com/r3-team/r3).

# Basic setup
* Connects to one or more REI3 instances via websocket.
* Auto-installs itself in the user directory and to startup if desired.

# Features

## Direct file editing from inside a REI3 browser session
* Receives requests for local file handling from a REI3 browser session (if same login/computer).
* Reacts to changes to managed files and uploads new file versions on change.
* Remembers cached files between sessions.
* Cleans up cached files when they are not used anymore.

## Handles client events, defined by REI3 applications
* Receives client events from connected REI3 instances.
* Listens to keyboard inputs if global hotkey events are used and specifically enabled by the user in their personal settings.
* Executes requested keystrokes from a REI3 browser session (if same login/computer).
* Sends basic information (hostname, username, current window title) and clipboard content to a REI3 instance or open browser session  (if same login/computer), if requested by client event such as a global hotkey.
