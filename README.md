## Logviewer

[Logviewer](http://jijeshmohan.github.io/logviewer) is a realtime log monitoring in browser. Tool helps to monitor all the logs without connecting(ssh connections) to the server.


#### Install Logviewer

1. **Compile from source**
	
    Requires Go to compile from the source. 
    
    ```
go get github.com/jijeshmohan/logviewer
```
2. **Install from binary**

  Linux [32bit](https://www.dropbox.com/s/5unjcsd0amty38o/logviewer_linux_x86.zip) | [64bit](https://www.dropbox.com/s/uqqf3yz55fna6md/logviewer_linux_x64.zip)
  
  Mac [64bit](https://www.dropbox.com/s/iz2z0h4p7ueh82h/logviewer_mac_x64.zip)

#### Running

Modify the **config.json** file with logs to be monitored and run the logviewer command.

**Command line options are:**

```
-p port 				: Specify the webserver port ( default to 8080)
-c configfilename 	    : Configuration file name (default to config.json)
```

### Inspiration 

[Logio](http://logio.org/)
