Still under development

##### Duet Timelapse Service
---
As 3dprinter owner with duetWifi board I want to create awesome timelapse

### What is it?
---
It's small service write in go that listen to telnet and execute command and also can serve latest image

### What I need?
---
- DuetWifi board with telnet open
`M586 P2 S1   ; Enable Telnet`
- Rasberry pi zero
- Pi camera
- Genrated gcode file with:
  - "print start" command sent on telnet
    `M118 P4 S"print start"` 
  this will execute scripts/print_start.sh
  create timelapse folder & delete old timelapse files 

   - On layer change
    `M118 P4 S"layer changed [layer_num]"`
    scripts/snapshot.sh
    taking image & updating latest image to serve

   - Print finish 
   `M118 P4 S"print finish"` 
    create tar gzip from timelapse folder with timestamp

##### Creating video from timelapse folder
The second part relay on windows machine with:
- ffmpeg install
- Git bash 

### Make it work!
---
Connect to your pi zero using ssh
and run installtion script
`wget -O - https://raw.githubusercontent.com/doryosef/dts/master/install.sh | sudo sh`

Edit the configuration file `/opt/dts/conf.ini`.

`[telnet]` section host paramter to your duetwifi ip

execute the following commands
```
sudo systemctl daemon-reload
sudo systemctl start dts
sudo systemctl enable dts.service
sudo systemctl status dts
```
Now print some stuffs


On your windows pc download `timelapse.sh` file and edit the host to your pi zero ip
open git bash and execute `$./timelapse.sh`

It will:
- Connect to your pi zero using sftp (read below tip to use ssh key connection without password) 
- Download tmielapse tar gzip files
- Delete the tmielapse tar gzip files from server 
- On local it will extract the files 
- Create movie using ffmpeg 
- Clean temp directory

#### Tips
----
##### Use ssh key to connect to your zero pi
```
On windows open git bash and execute command
	$ ssh-keygen -t rsa
Then execute command
	$ ssh-copy-id pi@remote_host
It will try to connect and will automatically copy the public key to the machine
Now try to login
	$ ssh pi@remote_host
```