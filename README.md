# burpee

##Introduction
-----

Command line HTTP client that reads URLs sources from a text file and directs traffic to an HTTP Proxy. 
Ideal to automatically populate the Target tab in Burp Suite Pro.

##Download##
Binary packages for every supported operating system are availble [here](https://github.com/nano6yt3/burpee/releases/latest).

##Install##
Extract the archive, and optionally, install binary to $PATH.
```
$ tar -zxvf burpee*.tar.gz
$ cd burpee*
$ sudo cp burpee /usr/local/bin
```

##Usage##
```
 Usage: burpee -file <path to file with URLs> -proxy <http://localhost:8081>
 
  -file     Path to the file with URLs to visit.
  -proxy    Protocol://IP:Port (e.g. https://192.168.2.24:8081) - default http://localhost:8080.
```
