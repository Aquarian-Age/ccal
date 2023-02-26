#!/bin/bash
set -v
for d in `ls`;do sshpass -p yourPassword scp -P remotePort -r $d root@1.5.4.2:/remotePath/;done

