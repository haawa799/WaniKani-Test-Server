# WaniKani-Test-Server

I made this small server for testing WaniKani API responses.
Here's a detailed tutorial on how to run small server with Go and Martini, that I used 

https://www.digitalocean.com/community/tutorials/how-to-use-martini-to-serve-go-applications-behind-an-nginx-server-on-ubuntu

# Usage 

You can use following endpoints:
```
http://haawa.org/test/wanikani/version/type
```

Where version is a number from 0 to 2 (data samples), and type right now can be either ```level-progression``` or ```study-queue```.

Eg.: ```http://haawa.org/test/wanikani/0/study-queue```


# Important

This is under active development, so don't rely on it too much :)
