## Logviewer

[Logviewer](https://github.com/jijeshmohan/logviewer) is a realtime log monitoring in browser. Tool helps to monitor all the logs without connecting to the server.


#### Install Logviewer

1. **Compile from source**
	
    Requires Go to compile from the source. 
    
    ```
go get github.com/jijeshmohan/logviewer
```
2. **Install from binary**

  [Linux 64bit](https://www.dropbox.com/s/jze8uukli1h053m/logviewer_linux_x64.tar.gz)

#### Running

Modify the **config.json** file with logs to be monitored and run the logviewer command.

**Command line options are:**

```
-p port 				: Specify the webserver port ( default to 8080)
-c configfilename 	    : Configuration file name (default to config.json)
```

