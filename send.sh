#!/bin/bash

SMPP_USER=root SMPP_PASSWD=secret

sms --addr 194.176.111.242:8018 --user DOSCredo --passwd='cReD0!x' query 8989 538adcd0
sms --addr 194.176.111.242:8018 --user DOSCredo1 --passwd='NrNcgvkC' query 6791 538adcd0

sms --addr 194.176.111.242:8018 --user DOSCredo --passwd='cReD0!x' send --encoding ucs2 --register 8989 996771977377 "Я твой Дед!"
sms --addr 194.176.111.242:8018 --user DOSCredo --passwd='cReD0!x' send --encoding ucs2 --register 8989 996771977377 "Мен сиздин Дед!"
sms --addr 194.176.111.242:8018 --user DOSCredo --passwd='cReD0!x' send --encoding ucs2 --register 8989 996771977377 "Мен сиздин Чоң-Ата! [Ѳө Ңң Үү]"
sms --addr 194.176.111.242:8018 --user DOSCredo --passwd='cReD0!x' send --encoding ucs2 --register 8989 996771977377 "I'm your Grandfather, Luke!"

sms --addr 194.176.111.242:8018 --user DOSCredo1 --passwd='NrNcgvkC' send --encoding ucs2 --register 6791 996771977377 "Я твой Дед!"
sms --addr 194.176.111.242:8018 --user DOSCredo1 --passwd='NrNcgvkC' send --encoding ucs2 --register 6791 996771977377 "Мен сиздин Дед!"
sms --addr 194.176.111.242:8018 --user DOSCredo1 --passwd='NrNcgvkC' send --encoding ucs2 --register 6791 996771977377 "Мен сиздин Чоң-Ата! [Ѳө Ңң Үү]"
sms --addr 194.176.111.242:8018 --user DOSCredo1 --passwd='NrNcgvkC' send --encoding ucs2 --register 6791 996771977377 "I'm your Grandfather, Luke!"

sms --addr 194.176.111.242:8018 --user DOSCredo --passwd='cReD0!x' send --encoding ucs2 --register 8989 996771977377 "Я твой Дед! Мен сиздин Дед! Мен сиздин Чоң-Ата! [Ѳө Ңң Үү] I'm your Grandfather, Luke!"
sms --addr 194.176.111.242:8018 --user DOSCredo1 --passwd='NrNcgvkC' send --encoding ucs2 --register 6791 996771977377 "Я твой Дед! Мен сиздин Дед! Мен сиздин Чоң-Ата! [Ѳө Ңң Үү] I'm your Grandfather, Luke!"

echo -ne 'Я твой Дед!\nМен сиздин Дед!\nМен сиздин Чоң-Ата! [Ѳө Ңң Үү]\nI\x27m your Grandfather, Luke!\n' | ./sms --addr 194.176.111.242:8018 --user DOSCredo --passwd='cReD0!x' send --encoding ucs2 --register 8989 996771977377 "$(cat)"
echo -ne 'Я твой Дед!\nМен сиздин Дед!\nМен сиздин Чоң-Ата! [Ѳө Ңң Үү]\nI\x27m your Grandfather, Luke!\n' | ./sms --addr 194.176.111.242:8018 --user DOSCredo1 --passwd='NrNcgvkC' send --encoding ucs2 --register 6791 996771977377 "$(cat)"

