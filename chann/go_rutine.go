package main

/* as we are learned about the thread in the os so go rutine is the same mean when we want to check the
website status so it will execute one then other and take some time to execute it so we use the concept of the
rutine in which we use the word go and it will make the different thread to execute and all the threads execute in the same
time parallel
......................Important point................
concurrency dosen't mean the parallelism
.....................Concurrency.....................
mean that we have the multiple threads working at the same time and one get blocked the other working thread is picked
and work on
in this rutine the main rutine is different and when we use the go then it will make the other child mean that when
go word called then go make the child of the other rutine
........................Channels.....................
channels is the medium of the channels between the child child thread mean when one child channel done their work then
it will inform the main thread that i have done my part of work
The channels is like the type struct and others
.............Communications between the channels .............
.............Important syntax are ...............
   channels <-5        .........mean that we are sending some value to the channels
   mynumber<-channels    ......mean some one is wating for the value to receive from the channels
   fmt.Println(<-channels) ..........mean printing the value form the channels

*/
