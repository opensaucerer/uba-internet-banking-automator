# Installation

Xvfb is a virtual framebuffer X server for UNIX-like operating systems. It is used to run graphical applications on a system that does not have a display server. It is also used to run graphical applications on a remote machine.

```bash
sudo apt-get install xvfb openjdk-11-jre
```

Install Java dependencies and drivers

```bash
cd binaries
go run init.go --alsologtostderr  --download_browsers --download_latest
```
