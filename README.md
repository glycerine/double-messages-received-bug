2015 Sept 30:
--------------

transcript showing bug (or design?) that it is currently impossible for an nsq Consumer
to receive exactly one message when there are two messages queued in a channel.

Protocol to reproduce:



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

Output:
=======


shell#1:
--------
~~~
$ nsqd --lookupd-tcp-address=127.0.0.1:4160
[nsqd] 2015/09/30 09:36:41.036854 nsqd v0.3.6 (built w/go1.5.1)
[nsqd] 2015/09/30 09:36:41.036931 ID: 842
[nsqd] 2015/09/30 09:36:41.037054 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqd] 2015/09/30 09:36:41.091980 HTTP: listening on [::]:4151
[nsqd] 2015/09/30 09:36:41.092007 LOOKUP(127.0.0.1:4160): adding peer
[nsqd] 2015/09/30 09:36:41.092026 LOOKUP connecting to 127.0.0.1:4160
[nsqlookupd] 2015/09/30 09:36:41.092299 TCP: new client(127.0.0.1:45582)
[nsqd] 2015/09/30 09:36:41.092476 TCP: listening on [::]:4150
[nsqlookupd] 2015/09/30 09:36:41.092592 CLIENT(127.0.0.1:45582): desired protocol magic '  V1'
[nsqlookupd] 2015/09/30 09:36:41.092792 CLIENT(127.0.0.1:45582): IDENTIFY Address:i7 TCP:4150 HTTP:4151 Version:0.3.6
[nsqlookupd] 2015/09/30 09:36:41.092814 DB: client(127.0.0.1:45582) REGISTER category:client key: subkey:
[nsqd] 2015/09/30 09:36:41.093156 LOOKUPD(127.0.0.1:4160): peer info {TCPPort:4160 HTTPPort:4161 Version:0.3.6 BroadcastAddress:i7}

[nsqd] 2015/09/30 09:36:56.092166 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:36:56.092354 CLIENT(127.0.0.1:45582): pinged (last ping 14.999556386s)
[nsqd] 2015/09/30 09:37:05.766228 200 GET /stats (127.0.0.1:55932) 140.141µs
[nsqd] 2015/09/30 09:37:11.092111 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:37:11.092288 CLIENT(127.0.0.1:45582): pinged (last ping 14.999934667s)
[nsqd] 2015/09/30 09:37:26.092102 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:37:26.092305 CLIENT(127.0.0.1:45582): pinged (last ping 15.000017221s)
[nsqd] 2015/09/30 09:37:34.756743 TCP: new client(127.0.0.1:42121)
[nsqd] 2015/09/30 09:37:34.756778 CLIENT(127.0.0.1:42121): desired protocol magic '  V2'
[nsqd] 2015/09/30 09:37:34.757197 [127.0.0.1:42121] IDENTIFY: {ShortID:i7 LongID:i7 ClientID:i7 Hostname:i7 HeartbeatInterval:30000 OutputBufferSize:16384 OutputBufferTimeout:250 FeatureNegotiation:true TLSv1:false Deflate:false DeflateLevel:6 Snappy:false SampleRate:0 UserAgent:go-nsq/1.0.5 MsgTimeout:0}
[nsqd] 2015/09/30 09:37:34.757766 TOPIC(caller-call-number-c7b3b5c46547603faae7736b6131b36e): created
[nsqd] 2015/09/30 09:37:34.757829 CI: querying nsqlookupd http://i7:4161/channels?topic=caller-call-number-c7b3b5c46547603faae7736b6131b36e
[nsqd] 2015/09/30 09:37:34.757969 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqd] 2015/09/30 09:37:34.758003 LOOKUPD(127.0.0.1:4160): topic REGISTER caller-call-number-c7b3b5c46547603faae7736b6131b36e
[nsqlookupd] 2015/09/30 09:37:34.758189 DB: client(127.0.0.1:45582) REGISTER category:topic key:caller-call-number-c7b3b5c46547603faae7736b6131b36e subkey:
[nsqd] 2015/09/30 09:37:36.758463 TOPIC(caller-call-number-c7b3b5c46547603faae7736b6131b36e): new channel(ch)
[nsqd] 2015/09/30 09:37:36.758505 LOOKUPD(127.0.0.1:4160): channel REGISTER caller-call-number-c7b3b5c46547603faae7736b6131b36e ch
[nsqlookupd] 2015/09/30 09:37:36.758721 DB: client(127.0.0.1:45582) REGISTER category:channel key:caller-call-number-c7b3b5c46547603faae7736b6131b36e subkey:ch
[nsqd] 2015/09/30 09:37:36.796599 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqd] 2015/09/30 09:37:41.092103 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:37:41.092306 CLIENT(127.0.0.1:45582): pinged (last ping 15.000000968s)
[nsqd] 2015/09/30 09:37:56.092103 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:37:56.092275 CLIENT(127.0.0.1:45582): pinged (last ping 14.999969224s)
[nsqd] 2015/09/30 09:38:11.092109 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:38:11.092278 CLIENT(127.0.0.1:45582): pinged (last ping 15.000002949s)
[nsqd] 2015/09/30 09:38:26.092112 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:38:26.092307 CLIENT(127.0.0.1:45582): pinged (last ping 15.000029536s)
...

# on separate shell, issue check to see that no old messages are hanging about:
$ wget -O - http://127.0.0.1:4151/stats
--2015-09-30 09:37:05--  http://127.0.0.1:4151/stats
Connecting to 127.0.0.1:4151... connected.
HTTP request sent, awaiting response... 200 OK
Length: 99 [text/plain]
Saving to: `STDOUT'

 0% [                                       ] 0           --.-K/s              nsqd v0.3.6 (built w/go1.5.1)
start_time 2015-09-30T09:36:41-07:00
uptime 24.729326856s

NO_TOPICS
100%[======================================>] 99          --.-K/s   in 0s      

2015-09-30 09:37:05 (16.1 MB/s) - written to stdout [99/99]

$
~~~

shell#2: start first caller
-------
~~~
$ ./caller
caller starting.
2015/09/30 09:41:15 INF    1 (127.0.0.1:4150) connecting to nsqd
published 'main.MyMsg{Hello:"hello at 2015-09-30 09:41:15.403974272 -0700 PDT", ReplyTo:"caller-call-number-1e8d3e1411ecfac823ffb8c52573f035", Reply:"", ReplyFrom:""}'
2015/09/30 09:41:17 INF    1 stopping
2015/09/30 09:41:17 ERR    1 (127.0.0.1:4150) IO error - EOF
2015/09/30 09:41:17 INF    1 exiting router
2015/09/30 09:41:17 INF    1 (127.0.0.1:4150) beginning close
2015/09/30 09:41:17 INF    1 (127.0.0.1:4150) readLoop exiting
2015/09/30 09:41:17 INF    1 (127.0.0.1:4150) breaking out of writeLoop
2015/09/30 09:41:17 INF    1 (127.0.0.1:4150) writeLoop exiting
2015/09/30 09:41:17 INF    2 [caller-call-number-1e8d3e1411ecfac823ffb8c52573f035/ch] (127.0.0.1:4150) connecting to nsqd
2015/09/30 09:41:17 INF    1 (127.0.0.1:4150) finished draining, cleanup exiting
2015/09/30 09:41:17 INF    1 (127.0.0.1:4150) clean close complete
2015/09/30 09:41:33 caller: Got a message Body (as MyMsg): main.MyMsg{Hello:"hello at 2015-09-30 09:41:15.403974272 -0700 PDT", ReplyTo:"caller-call-number-1e8d3e1411ecfac823ffb8c52573f035", Reply:"replying at 2015-09-30 09:41:33.010813598 -0700 PDT", ReplyFrom:"callee!"}
stats = &nsq.ConsumerStats{MessagesReceived:0x1, MessagesFinished:0x0, MessagesRequeued:0x0, Connections:1}
caller done with response='&main.MyMsg{Hello:"hello at 2015-09-30 09:41:15.403974272 -0700 PDT", ReplyTo:"caller-call-number-1e8d3e1411ecfac823ffb8c52573f035", Reply:"replying at 2015-09-30 09:41:33.010813598 -0700 PDT", ReplyFrom:"callee!"}'.
$
~~~

shell#3: start 2nd caller
-------

~~~
$ ./caller
caller starting.
2015/09/30 09:41:24 INF    1 (127.0.0.1:4150) connecting to nsqd
published 'main.MyMsg{Hello:"hello at 2015-09-30 09:41:24.270684415 -0700 PDT", ReplyTo:"caller-call-number-1d984bd497893c94b124337fcf10778a", Reply:"", ReplyFrom:""}'
2015/09/30 09:41:24 INF    1 stopping
2015/09/30 09:41:24 ERR    1 (127.0.0.1:4150) IO error - EOF
2015/09/30 09:41:24 INF    1 exiting router
2015/09/30 09:41:24 INF    1 (127.0.0.1:4150) beginning close
2015/09/30 09:41:24 INF    1 (127.0.0.1:4150) readLoop exiting
2015/09/30 09:41:24 INF    1 (127.0.0.1:4150) breaking out of writeLoop
2015/09/30 09:41:24 INF    1 (127.0.0.1:4150) writeLoop exiting
2015/09/30 09:41:24 INF    2 [caller-call-number-1d984bd497893c94b124337fcf10778a/ch] (127.0.0.1:4150) connecting to nsqd
2015/09/30 09:41:24 INF    1 (127.0.0.1:4150) finished draining, cleanup exiting
2015/09/30 09:41:24 INF    1 (127.0.0.1:4150) clean close complete

<hangs forever here -- since the one callee has consumed both calls>
~~~

shell#4 : start one callee. The callee Consumer gets both messages, instead of just one. The caller in shell#3 thus never gets a reply.
-------

~~~
$ ./callee
2015/09/30 09:41:33 INF    1 [write_test/ch] (127.0.0.1:4150) connecting to nsqd
2015/09/30 09:41:33 callee.ConsumeRequest(): Got a message Body (as MyMsg): main.MyMsg{Hello:"hello at 2015-09-30 09:41:15.403974272 -0700 PDT", ReplyTo:"caller-call-number-1e8d3e1411ecfac823ffb8c52573f035", Reply:"", ReplyFrom:""}
stats = &nsq.ConsumerStats{MessagesReceived:0x1, MessagesFinished:0x0, MessagesRequeued:0x0, Connections:1}
2015/09/30 09:41:33 callee.ConsumeRequest(): Got a message Body (as MyMsg): main.MyMsg{Hello:"hello at 2015-09-30 09:41:24.270684415 -0700 PDT", ReplyTo:"caller-call-number-1d984bd497893c94b124337fcf10778a", Reply:"", ReplyFrom:""}
2015/09/30 09:41:33 INF    2 (127.0.0.1:4150) connecting to nsqd
callee.ProduceReply(): replied-to: 'caller-call-number-1e8d3e1411ecfac823ffb8c52573f035' with value '&main.MyMsg{Hello:"hello at 2015-09-30 09:41:15.403974272 -0700 PDT", ReplyTo:"caller-call-number-1e8d3e1411ecfac823ffb8c52573f035", Reply:"replying at 2015-09-30 09:41:33.010813598 -0700 PDT", ReplyFrom:"callee!"}'
2015/09/30 09:41:33 INF    2 stopping
2015/09/30 09:41:33 INF    2 exiting router
$
$ wget -O - http://127.0.0.1:4151/stats
--2015-09-30 09:42:01--  http://127.0.0.1:4151/stats
Connecting to 127.0.0.1:4151... connected.
HTTP request sent, awaiting response... 200 OK
Length: 919 [text/plain]
Saving to: `STDOUT'

 0% [                                       ] 0           --.-K/s              nsqd v0.3.6 (built w/go1.5.1)
start_time 2015-09-30T09:41:00-07:00
uptime 1m1.503297924s

Health: OK

   [caller-call-number-1d984bd497893c94b124337fcf10778a] depth: 0     be-depth: 0     msgs: 0        e2e%: 
      [ch                       ] depth: 0     be-depth: 0     inflt: 0    def: 0    re-q: 0     timeout: 0     msgs: 0        e2e%: 
        [V2 i7:42171             ] state: 3 inflt: 0    rdy: 1    fin: 0        re-q: 0        msgs: 0        connected: 37s

   [caller-call-number-1e8d3e1411ecfac823ffb8c52573f035] depth: 0     be-depth: 0     msgs: 1        e2e%: 
      [ch                       ] depth: 0     be-depth: 0     inflt: 0    def: 0    re-q: 0     timeout: 0     msgs: 1        e2e%: 

   [write_test     ] depth: 0     be-depth: 0     msgs: 2        e2e%: 
      [ch                       ] depth: 0     be-depth: 0     inflt: 0    def: 0    re-q: 0     timeout: 0     msgs: 2        e2e%: 
100%[======================================>] 919         --.-K/s   in 0s      

2015-09-30 09:42:01 (135 MB/s) - written to stdout [919/919]

$
~~~
