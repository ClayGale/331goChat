# 331goChat
simple golang based chatroom app for cosc331

Step 1:
create digital ocean server 

Step 2:
Install relevant software
-Tried to use docker-machine, but gave errors when trying to make a swarm.
  -could not recognize virtual box even though it was downloaded
  -fiddled with that for a bit and realized we needed to take a different path
  -tried to use digital ocean api in the mean time had some issues
-Decided to use docker containers instead
-wrote the go script
-made html for site chatroom system
-tired running html with docker, kept giving running screen
-it just hung in an infaninte loop, had to force stop

Step 3:
Pull repo
Step n:
Create the database tables with
```
Create table messages ( ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, message varchar(255));
Create table users ( name varchar(20), colour varchar(10));
```

conclusion- could not get to the databases because we could not make it past step 2
we thought that creating a simple project with go, and deploying it with docker would be simple to do
