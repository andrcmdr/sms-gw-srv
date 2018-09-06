#!/bin/bash

sudo mkdir -v -p /var/log/smsgw/

sudo chown -v -R andrcmdr:andrcmdr /var/log/smsgw/

./sms.srv >> /var/log/smsgw/smsgw.log 2>&1 & disown

