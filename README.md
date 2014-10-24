## Logviewer

[Logviewer](http://jijeshmohan.github.io/logviewer) is a realtime log monitoring in browser. Tool helps to monitor all the logs without connecting(ssh connections) to the server.


#### Install Logviewer

1. **Compile from source**
	
    Requires Go to compile from the source. 
    
    ```
go get github.com/jijeshmohan/logviewer
```
2. **Install from binary**

 Linux & Mac Binaries available
 
 [Download] (https://github.com/jijeshmohan/logviewer/releases/tag/v0.2)
 

#### Running

Modify the **config.json** file with logs to be monitored and run the logviewer command.

**Command line options are:**

```
-p port 				: Specify the webserver port ( default to 8080)
-c configfilename 	    : Configuration file name (default to config.json)
```

### Inspiration 

[Logio](http://logio.org/)
