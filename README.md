2015 Sept 30:
--------------

Basic produce request, consume reply over nsq. What we learned:

One must use #ephemeral topics *and* #ephemeral channels to get auto-cleanup.

See also https://gist.github.com/glycerine/cbdd58e889b8805a7101
and the original https://gist.github.com/joshrotenberg/5a3acb44d3dbad884397


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

shell#3: start the lone callee
$ ./callee
~~~
