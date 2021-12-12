# address-server

This is a simple address server, that echo your ip and port once you 
visit it at `/echo-address`. If you want to setup a firewall rule to let your
ip throught it, this may be helpful.

Usage:

- Assuming that you are using a powershell terminal, just run `build.ps1` to
  build both linux and window (amd64) executable.
- If you want to change the listen port or address, you should change the code
  and rebuild it.
