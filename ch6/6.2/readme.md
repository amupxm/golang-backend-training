### Chapter 6 , section 6.2



```sh
curl 'http://localhost:8080/?name=YOURNAME' 
```
### Chapter 5 , section 2

You can send request like :

```sh
curl 'http://localhost:8080/?name=YOURNAME'
```

Which will return 
```json
{
"message": "Hello YOURNAME"
}
```
.If you send any request which not contains name parameter , it will returns you an error .
