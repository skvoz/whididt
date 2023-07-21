# whididt
This CLI library allows you generat day(half day :) report for your chif.
With max setup you will generate report and send on messenger, in time what your need.

## Usage
Example genrate daylick report. Your get log from git an current day, and output on your screen. 
```
cd <your project pwd>
whididt stash-logs 
cd <your anohter project pwd>
whididt stash-logs 
<repeat n-times>
whididt pop-logs
```

## Special ops
whididt [OPTIONS] [COMMAND]

Options:
-c, --channel       set channel (slack, skype, DB)
-s, --schedult      generate report on time (and set it to channel)
-d, --diapazon      diapazon date(<n>d, <n>m, <n>y), default 1d(one day)
-D, --date          date start stash log, ( -D 01-01-2022 -d 3d) stash log from 01-01-2022 to 03-01-2022

Command:
-ls, --stash-logs   add logs to storage
-pl, --pop-logs     generate daylick report from logs


## Configurations 
Need for add channels, and CHAT-GPT for generate more intelligent report

Slack :
```TODO```
Skype
```TODO```
Telegram
```TODO```
CHAT-GPT
```TODO```
