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


* Receiver is a goroutine which continously listens to the sender channel and prints the message on console.

* To terminate the communication, simple send the message "/exit" using the send API.     
  
Note: Multiple messages can be sent with multiple send API calls    
  
  
Example request in postman:
  
POST localhost:8080/send    
 ``` 
 Body raw application/JSON =>    
  
 {"content":"Example message"}
 ```
  



    
  
