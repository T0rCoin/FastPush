# Telegram bot high-performance message push

## how to work?


* Receive push command ->
* set push data package -> 
* set uid to redis queue -> 
* trigger push signal -> 
* push server operation -> 
* message push server out of queue -> 
* user receives message

### You need to edit ```conf/conf.ini``` to configure this push server

### We set Redis's No. 13 database as the database for storing Telegram _Uid_ by default, you can modify it on line _43_ of ```src/Redis/Redis.go```
