### Chapter 6 , section 6.3

You can send request like :

```sh
curl 'http://localhost:8080/?x=1&y=26'
```

Which will return 
```json
{
"result": "27"
}
```
.If you send any request which not contains y or x parameter , it will returns you an error .
