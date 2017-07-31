# Slack Channel Inviter

> Invite all of #channelA to #channelB (skipping those already in #channelB)

I had a need to shutdown and migrate a large channel into another. This little CLI tool did the job. It may be of use to others.

### Usage

```golang
env SLACK_TOKEN=xoxp-<your_token> go run main.go channelA channelB
```
### Output

```bash
2017/07/31 09:39:49 Invited userA (U0A3DXXX) to #channelB (C0KKV4XXX)
2017/07/31 09:39:50 Invited userB (U0A3DXXX) to #channelB (C0KKV4XXX)
...
...
```
