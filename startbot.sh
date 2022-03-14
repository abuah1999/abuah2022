#!/bin/bash
go build .
ls
chmod a+x ./abuahfish_v2
python lichess-bot.py
