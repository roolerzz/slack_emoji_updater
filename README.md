# slack_emoji_updater
This repository contains scraping code in Golang, that scrapes a website hosting the [slack emoji's](https://slackmojis.com/) and would upload them to a specifc user's slack workspace, provided authentication tokens for that workspace.

**Update:** This is not a WIP repo. I am not expecting any further development on this. I was able to successfully scrape the URL for the slack emoji's from the HTML content of a website, using an HTML parser, but I haven't implemented the part of uploading to the slack workspace, because of lack of API support from slack. So I ended up using a command line tool that was available to bulk upload the emojis. I ended up using [emojipacks](https://github.com/lambtron/emojipacks).
