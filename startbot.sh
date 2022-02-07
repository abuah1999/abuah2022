#!/bin/bash
git add .
git commit -m"m"
git push heroku master
pyinstaller --onefile abuahfish/abuahfish.py
chmod a+x /dist/abuahfish
python lichess-bot.py