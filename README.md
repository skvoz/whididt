# whididt
This CLI library allows you generat day(half day :) report for your chif.
With max setup you will generate report and send on messenger, in time what your need.

## Usage
Example genrate daylick report. Your get log from git an current day, and output on your screen. 
```
cd <your project pwd>
whididt 
```

## Special ops
whididt [OPTIONS] 

Options:
-p, --path          pwd your projects (ex: -p ~/porjects/1 ~/projects/2 )
-c, --channel       set channel (slack, skype, DB)
-s, --schedule      generate report on time (and set it to channel)
-d, --diapazon      diapazon date(<n>d, <n>m, <n>y), default 1d(one day)
-S, --start         date start log, (ex: -S 12-31-2022) get log from 12-31-2022 to 12-31-2022
-u, --until         date until log,  (ex: -S 12-31-2022 -u 12-31-2022) get log from 12-31-2022 to 12-31-2022
-b --boss           boss name

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
