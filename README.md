# whididt
This CLI library allows you generat day(half day :) report for your chif.
With max setup you will generate report and send on messenger, in time what your need.

## Setup
wget -C https://github.com/skvoz/whididt/whididt.v0.1b.tar.gz
tar -xzvf projects.tar.gz -C /opt/whididt.v0.1b
cd /opt/whididt.v0.1b && ln -s /usr/local/bin

## Usage
Example genrate daylick report. Your get log from git an current day, and output on your screen. 
```
cd <your project pwd>
whididt 
```

## Special ops
whididt [OPTIONS] 

Options:
-p --path          pwd your projects (ex: -p ~/porjects/1 ~/projects/2 )
-c --channel       set channel (slack, skype, DB)
-s --schedule      generate report on time (and set it to channel)
-S --start         date start log, (ex: -S 12-31-2022) get log from 12-31-2022 to 12-31-2022
-u --until         date until log,  (ex: -S 12-31-2022 -u 12-31-2022) get log from 12-31-2022 to 12-31-2022
-b --boss          boss name
-t --template      template name (add own tempalte)
-ai --ai           artificial intelligence, for more kindly report

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

##ROADMAP
v0.0 create MVP
v0.1b output simple report. working flags: -p, -S, -u, -b
v0.2b fix error form 0.1b, add unit tests
v0.3b work with flag -t, -c
v0.1 add flag -ai, add slack channel to flag -c (redy to use)

