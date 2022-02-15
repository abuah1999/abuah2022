#!/bin/bash
python -m nuitka --follow-imports abuahfish/abuahfish.py
ls
chmod a+x ./abuahfish.bin
python lichess-bot.py
