# Sender Receiver Framework
Simple Sender Receiver framework in Go

Requirements:
* Postman tool or any alternate API testing tool  
  
Working: 
* Server is hosted in localhost port :8080. 
* API's are provided below,

&nbsp;&nbsp;* POST /send

```
raw application/JSON body  
  
request body example = {"content" : "Hi. This is my first message"}
```
  
&nbsp;&nbsp;* GET /receive    
  
Note: Multiple messages can be sent with multiple send API calls. Receive API will return all the messages to be received.    
  
  
Example request in postman:
  
1. POST localhost:8080/send    
 ``` 
 Body raw application/JSON =>    
  
 {"content":"Example message"}
 ```
  
2. GET localhost:8080/receive



    
  
