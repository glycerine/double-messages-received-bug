2015 Sept 30:
--------------

Basic produce request, consume reply over nsq. What we learned:

1. One must use #ephemeral topics *and* #ephemeral channels to get auto-cleanup.

See also https://gist.github.com/glycerine/cbdd58e889b8805a7101
and the original https://gist.github.com/joshrotenberg/5a3acb44d3dbad884397

2. To avoid having a callee consume more than one message, per Jehiah's recommendations:

> On Wednesday, September 30, 2015 at 11:00:32 AM UTC-7, Jehiah wrote:
Typically you do all processing before calling Finish, and at that 
time you are ready to process the next message. (It's generally 
expected that consumers are long-lived, but that isn't a hard 
requirement) 

> if you want to gracefully shut down you have two options 

> 1) ChangeMaxInFlight(0) before calling .Finish() 
2) call consumer.Stop() and have your main thread wait on the stop 
channel - https://godoc.org/github.com/nsqio/go-nsq#Consumer.Stop 

>Typically you have a term signal handler that will call 
consumer.Stop() and your main goroutine will block on 
https://godoc.org/github.com/nsqio/go-nsq#Consumer.StopChan 
immediately after connecting to nsqd. 


~~~
shell#1: build, then start nsqd, verify no old messages.
$ make
go build -o caller caller.go
go build -o callee callee.go
$ nsqd --lookupd-tcp-address=127.0.0.1:4160 &
$ wget -O - http://127.0.0.1:4151/stats

shell#2: start 1st caller
$ ./caller

shell#3: start 2nd caller
$ ./caller

shell#3: start one callee
$ ./callee
~~~

The above used to result in the callee getting both jobs, and one of the callers never getting that lost job serviced. But not anymore with Consumer.Stop() being used.
