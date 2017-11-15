check status host
=================
Learning Go - This example checks host status

### Build & Run
```
$ go build monit.go
```
Save the address of the hosts you want to monitor in `hosts.txt` and run:
```
$ echo "https://infoslack.pro" >> hosts.txt
$ ./monit
Monitoring...
Check https://infoslack.com
https://infoslack.com OK
Check https://infoslack.pro
https://infoslack.pro OK
^C
```
Press `ctrl+c` to stop.

### Logs
Logs are stored in status.log
```
$ tail status.log 
Nov 15 16:49:18 - https://infoslack.com - online: true
Nov 15 16:49:19 - https://infoslack.pro - online: true
Nov 15 16:49:23 - https://infoslack.com - online: true
Nov 15 16:49:23 - https://infoslack.pro - online: true
```
