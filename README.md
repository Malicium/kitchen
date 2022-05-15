# kitchen
Server



Quick Test
```
curl -d "cmd=node&input=function sayHello() { return 'Hello from Kitchen!' } sayHello()" -X POST http://localhost:8080/run-code/
```