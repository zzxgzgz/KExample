# KExample
 Simple web server that runs on VM set up by KubeEdge


## What it does: 

 It is a simple server that gets HTTP request for people's quote.

 People's name is a part of the URL. For example, if the request is: localhost:8080/Rio, the server looks for the quote for "Rio"; if the request is localhost:8080/BenjaminFranklin, the server looks for the quote for "Benjamin Franklin"; if the name is not found, it simply writes "Sorry, cannot find quote for 'requestName' "