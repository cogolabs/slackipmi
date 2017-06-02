[![Build Status](https://travis-ci.org/cogolabs/slackipmi.svg?branch=master)](https://travis-ci.org/cogolabs/slackipmi)
[![Docker Build Status](https://img.shields.io/docker/build/cogolabs/slackipmi.svg)](https://hub.docker.com/r/cogolabs/slackipmi/)
[![Coverage Status](https://coveralls.io/repos/github/cogolabs/slackipmi/badge.svg?branch=master)](https://coveralls.io/github/cogolabs/slackipmi?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/cogolabs/slackipmi)](https://goreportcard.com/report/github.com/cogolabs/slackipmi)
# slackipmi
reboot servers via IPMI over Slack!
```
~]# docker run cogolabs/slackipmi -h
Usage of /go/bin/slackipmi:
  -base-url string
    	base URL (Slack-accessible) (default "https://slack.colofoo.net")
  -http string
    	 (default ":80")
  -oauth-client-id string
    	Slack-provided client ID
  -oauth-client-secret string
    	Slack-provided client secret
  -slack-team string
    	ignore requests outside your Slack Team (default "myslackorg")
  -slack-token string
    	token verifies reqs are actually coming from Slack (default "V4hafFbeT1doasdfkXeE4f")
```
![image](screenshots/slack-buttons.png)
![image](screenshots/slack-command.png)
