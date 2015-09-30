2015 Sept 30:
--------------

transcript showing that it is currently impossible (bug?) for a Consumer
to receive exactly one message when there are two messages queued in a channel.

Protocol to reproduce:



~~~
shell#1: build, then start nsqd
$ make
go build -o caller caller.go
go build -o callee callee.go
$ nsqd --lookupd-tcp-address=127.0.0.1:4160 &

shell#2: start 1st caller
$ ./caller

shell#3: start 2nd caller
$ ./caller

shell#3: start the callee
$ ./callee
~~~

Output:
=======


shell#1:
--------
~~~
$ nsqd --lookupd-tcp-address=127.0.0.1:4160 &
[1] 3311
[nsqd] 2015/09/30 09:21:03.068501 nsqd v0.3.6 (built w/go1.5.1)
[nsqd] 2015/09/30 09:21:03.068569 ID: 842
[nsqd] 2015/09/30 09:21:03.068601 NSQ: persisting topic/channel metadata to nsqd.842.dat
$ [nsqd] 2015/09/30 09:21:03.114194 LOOKUP(127.0.0.1:4160): adding peer
[nsqd] 2015/09/30 09:21:03.114214 LOOKUP connecting to 127.0.0.1:4160
[nsqd] 2015/09/30 09:21:03.114234 HTTP: listening on [::]:4151
[nsqd] 2015/09/30 09:21:03.114304 TCP: listening on [::]:4150
[nsqlookupd] 2015/09/30 09:21:03.114481 TCP: new client(127.0.0.1:45332)
[nsqlookupd] 2015/09/30 09:21:03.114500 CLIENT(127.0.0.1:45332): desired protocol magic '  V1'
[nsqlookupd] 2015/09/30 09:21:03.114633 CLIENT(127.0.0.1:45332): IDENTIFY Address:i7 TCP:4150 HTTP:4151 Version:0.3.6
[nsqlookupd] 2015/09/30 09:21:03.114649 DB: client(127.0.0.1:45332) REGISTER category:client key: subkey:
[nsqd] 2015/09/30 09:21:03.114877 LOOKUPD(127.0.0.1:4160): peer info {TCPPort:4160 HTTPPort:4161 Version:0.3.6 BroadcastAddress:i7}
[nsqd] 2015/09/30 09:21:11.391908 TCP: new client(127.0.0.1:41868)
[nsqd] 2015/09/30 09:21:11.392025 CLIENT(127.0.0.1:41868): desired protocol magic '  V2'
[nsqd] 2015/09/30 09:21:11.392603 [127.0.0.1:41868] IDENTIFY: {ShortID:i7 LongID:i7 ClientID:i7 Hostname:i7 HeartbeatInterval:30000 OutputBufferSize:16384 OutputBufferTimeout:250 FeatureNegotiation:true TLSv1:false Deflate:false DeflateLevel:6 Snappy:false SampleRate:0 UserAgent:go-nsq/1.0.5 MsgTimeout:0}
[nsqd] 2015/09/30 09:21:11.393272 TOPIC(write_test): created
[nsqd] 2015/09/30 09:21:11.393340 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqd] 2015/09/30 09:21:11.393387 LOOKUPD(127.0.0.1:4160): topic REGISTER write_test
[nsqd] 2015/09/30 09:21:11.393389 CI: querying nsqlookupd http://i7:4161/channels?topic=write_test
[nsqlookupd] 2015/09/30 09:21:11.393559 DB: client(127.0.0.1:45332) REGISTER category:topic key:write_test subkey:
[nsqd] 2015/09/30 09:21:13.395455 TCP: new client(127.0.0.1:41870)
[nsqd] 2015/09/30 09:21:13.395485 CLIENT(127.0.0.1:41870): desired protocol magic '  V2'
[nsqd] 2015/09/30 09:21:13.395713 [127.0.0.1:41870] IDENTIFY: {ShortID:i7 LongID:i7 ClientID:i7 Hostname:i7 HeartbeatInterval:30000 OutputBufferSize:16384 OutputBufferTimeout:250 FeatureNegotiation:true TLSv1:false Deflate:false DeflateLevel:6 Snappy:false SampleRate:0 UserAgent:go-nsq/1.0.5 MsgTimeout:0}
[nsqd] 2015/09/30 09:21:13.448959 TOPIC(caller-call-number-7a50a6cfb60ce77fd54970972e456acc): created
[nsqd] 2015/09/30 09:21:13.449016 CI: querying nsqlookupd http://i7:4161/channels?topic=caller-call-number-7a50a6cfb60ce77fd54970972e456acc
[nsqd] 2015/09/30 09:21:13.449064 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqd] 2015/09/30 09:21:13.449113 LOOKUPD(127.0.0.1:4160): topic REGISTER caller-call-number-7a50a6cfb60ce77fd54970972e456acc
[nsqlookupd] 2015/09/30 09:21:13.449322 DB: client(127.0.0.1:45332) REGISTER category:topic key:caller-call-number-7a50a6cfb60ce77fd54970972e456acc subkey:
[nsqd] 2015/09/30 09:21:13.495095 PROTOCOL(V2): [127.0.0.1:41868] exiting ioloop
[nsqd] 2015/09/30 09:21:13.495207 PROTOCOL(V2): [127.0.0.1:41868] exiting messagePump
[nsqd] 2015/09/30 09:21:14.391308 TOPIC(caller-call-number-7a50a6cfb60ce77fd54970972e456acc): new channel(ch)
[nsqd] 2015/09/30 09:21:14.391445 LOOKUPD(127.0.0.1:4160): channel REGISTER caller-call-number-7a50a6cfb60ce77fd54970972e456acc ch
[nsqlookupd] 2015/09/30 09:21:14.391665 DB: client(127.0.0.1:45332) REGISTER category:channel key:caller-call-number-7a50a6cfb60ce77fd54970972e456acc subkey:ch
[nsqd] 2015/09/30 09:21:14.427787 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqd] 2015/09/30 09:21:18.114307 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:21:18.114501 CLIENT(127.0.0.1:45332): pinged (last ping 14.99985976s)
[nsqd] 2015/09/30 09:21:19.850204 TCP: new client(127.0.0.1:41873)
[nsqd] 2015/09/30 09:21:19.850236 CLIENT(127.0.0.1:41873): desired protocol magic '  V2'
[nsqd] 2015/09/30 09:21:19.850503 [127.0.0.1:41873] IDENTIFY: {ShortID:i7 LongID:i7 ClientID:i7 Hostname:i7 HeartbeatInterval:30000 OutputBufferSize:16384 OutputBufferTimeout:250 FeatureNegotiation:true TLSv1:false Deflate:false DeflateLevel:6 Snappy:false SampleRate:0 UserAgent:go-nsq/1.0.5 MsgTimeout:0}
[nsqd] 2015/09/30 09:21:19.852362 TCP: new client(127.0.0.1:41874)
[nsqd] 2015/09/30 09:21:19.852398 CLIENT(127.0.0.1:41874): desired protocol magic '  V2'
[nsqd] 2015/09/30 09:21:19.852670 [127.0.0.1:41874] IDENTIFY: {ShortID:i7 LongID:i7 ClientID:i7 Hostname:i7 HeartbeatInterval:30000 OutputBufferSize:16384 OutputBufferTimeout:250 FeatureNegotiation:true TLSv1:false Deflate:false DeflateLevel:6 Snappy:false SampleRate:0 UserAgent:go-nsq/1.0.5 MsgTimeout:0}
[nsqd] 2015/09/30 09:21:19.853078 TOPIC(caller-call-number-845002f5f63c29043cf0fd63a196cedc): created
[nsqd] 2015/09/30 09:21:19.853127 CI: querying nsqlookupd http://i7:4161/channels?topic=caller-call-number-845002f5f63c29043cf0fd63a196cedc
[nsqd] 2015/09/30 09:21:19.853158 LOOKUPD(127.0.0.1:4160): topic REGISTER caller-call-number-845002f5f63c29043cf0fd63a196cedc
[nsqd] 2015/09/30 09:21:19.853129 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqlookupd] 2015/09/30 09:21:19.853344 DB: client(127.0.0.1:45332) REGISTER category:topic key:caller-call-number-845002f5f63c29043cf0fd63a196cedc subkey:
[nsqd] 2015/09/30 09:21:19.952924 PROTOCOL(V2): [127.0.0.1:41873] exiting ioloop
[nsqd] 2015/09/30 09:21:19.953002 PROTOCOL(V2): [127.0.0.1:41873] exiting messagePump
[nsqd] 2015/09/30 09:21:21.853726 TOPIC(caller-call-number-845002f5f63c29043cf0fd63a196cedc): new channel(ch)
[nsqd] 2015/09/30 09:21:21.853789 LOOKUPD(127.0.0.1:4160): channel REGISTER caller-call-number-845002f5f63c29043cf0fd63a196cedc ch
[nsqlookupd] 2015/09/30 09:21:21.853986 DB: client(127.0.0.1:45332) REGISTER category:channel key:caller-call-number-845002f5f63c29043cf0fd63a196cedc subkey:ch
[nsqd] 2015/09/30 09:21:21.895827 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqd] 2015/09/30 09:21:29.209205 TCP: new client(127.0.0.1:41877)
[nsqd] 2015/09/30 09:21:29.209244 CLIENT(127.0.0.1:41877): desired protocol magic '  V2'
[nsqd] 2015/09/30 09:21:29.209622 [127.0.0.1:41877] IDENTIFY: {ShortID:i7 LongID:i7 ClientID:i7 Hostname:i7 HeartbeatInterval:30000 OutputBufferSize:16384 OutputBufferTimeout:250 FeatureNegotiation:true TLSv1:false Deflate:false DeflateLevel:6 Snappy:false SampleRate:0 UserAgent:go-nsq/1.0.5 MsgTimeout:0}
[nsqd] 2015/09/30 09:21:29.210107 TOPIC(write_test): new channel(ch)
[nsqd] 2015/09/30 09:21:29.210169 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqd] 2015/09/30 09:21:29.210332 LOOKUPD(127.0.0.1:4160): channel REGISTER write_test ch
[nsqlookupd] 2015/09/30 09:21:29.210582 DB: client(127.0.0.1:45332) REGISTER category:channel key:write_test subkey:ch
[nsqd] 2015/09/30 09:21:29.211513 TCP: new client(127.0.0.1:41878)
[nsqd] 2015/09/30 09:21:29.211546 CLIENT(127.0.0.1:41878): desired protocol magic '  V2'
[nsqd] 2015/09/30 09:21:29.211842 [127.0.0.1:41878] IDENTIFY: {ShortID:i7 LongID:i7 ClientID:i7 Hostname:i7 HeartbeatInterval:30000 OutputBufferSize:16384 OutputBufferTimeout:250 FeatureNegotiation:true TLSv1:false Deflate:false DeflateLevel:6 Snappy:false SampleRate:0 UserAgent:go-nsq/1.0.5 MsgTimeout:0}
[nsqd] 2015/09/30 09:21:29.254704 PROTOCOL(V2): [127.0.0.1:41877] exiting ioloop
[nsqd] 2015/09/30 09:21:29.254811 PROTOCOL(V2): [127.0.0.1:41877] exiting messagePump
[nsqd] 2015/09/30 09:21:29.254704 PROTOCOL(V2): [127.0.0.1:41878] exiting ioloop
[nsqd] 2015/09/30 09:21:29.254711 PROTOCOL(V2): [127.0.0.1:41870] exiting ioloop
[nsqd] 2015/09/30 09:21:29.254891 PROTOCOL(V2): [127.0.0.1:41878] exiting messagePump
[nsqd] 2015/09/30 09:21:29.255021 PROTOCOL(V2): [127.0.0.1:41870] exiting messagePump
[nsqd] 2015/09/30 09:21:33.114307 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:21:33.114506 CLIENT(127.0.0.1:45332): pinged (last ping 15.000004536s)
[nsqd] 2015/09/30 09:21:48.114306 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:21:48.114492 CLIENT(127.0.0.1:45332): pinged (last ping 14.999986706s)
[nsqd] 2015/09/30 09:22:03.114297 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:22:03.114500 CLIENT(127.0.0.1:45332): pinged (last ping 15.000004192s)
[nsqd] 2015/09/30 09:22:18.114298 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:22:18.114475 CLIENT(127.0.0.1:45332): pinged (last ping 14.999979025s)
[nsqd] 2015/09/30 09:22:33.114301 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:22:33.114509 CLIENT(127.0.0.1:45332): pinged (last ping 15.000036796s)
[nsqd] 2015/09/30 09:22:48.114302 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:22:48.114477 CLIENT(127.0.0.1:45332): pinged (last ping 14.999963924s)
[nsqd] 2015/09/30 09:23:03.114305 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:23:03.114481 CLIENT(127.0.0.1:45332): pinged (last ping 15.000005316s)
[nsqd] 2015/09/30 09:23:18.114298 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:23:18.114501 CLIENT(127.0.0.1:45332): pinged (last ping 15.000019287s)
[nsqd] 2015/09/30 09:23:33.114298 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:23:33.114470 CLIENT(127.0.0.1:45332): pinged (last ping 14.999963999s)
[nsqd] 2015/09/30 09:23:48.114298 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:23:48.114471 CLIENT(127.0.0.1:45332): pinged (last ping 15.000003498s)
[nsqd] 2015/09/30 09:24:03.114299 LOOKUPD(127.0.0.1:4160): sending heartbeat
...
~~~

shell#2
-------
~~~
$ ./caller
caller starting.
2015/09/30 09:21:11 INF    1 (127.0.0.1:4150) connecting to nsqd
published 'main.MyMsg{Hello:"hello at 2015-09-30 09:21:11.391440446 -0700 PDT", ReplyTo:"caller-call-number-7a50a6cfb60ce77fd54970972e456acc", Reply:"", ReplyFrom:""}'
2015/09/30 09:21:13 INF    1 stopping
2015/09/30 09:21:13 ERR    1 (127.0.0.1:4150) IO error - EOF
2015/09/30 09:21:13 INF    1 exiting router
2015/09/30 09:21:13 INF    1 (127.0.0.1:4150) beginning close
2015/09/30 09:21:13 INF    1 (127.0.0.1:4150) readLoop exiting
2015/09/30 09:21:13 INF    1 (127.0.0.1:4150) breaking out of writeLoop
2015/09/30 09:21:13 INF    1 (127.0.0.1:4150) writeLoop exiting
2015/09/30 09:21:13 INF    2 [caller-call-number-7a50a6cfb60ce77fd54970972e456acc/ch] (127.0.0.1:4150) connecting to nsqd
2015/09/30 09:21:13 INF    1 (127.0.0.1:4150) finished draining, cleanup exiting
2015/09/30 09:21:13 INF    1 (127.0.0.1:4150) clean close complete
2015/09/30 09:21:29 caller: Got a message Body (as MyMsg): main.MyMsg{Hello:"hello at 2015-09-30 09:21:11.391440446 -0700 PDT", ReplyTo:"caller-call-number-7a50a6cfb60ce77fd54970972e456acc", Reply:"replying at 2015-09-30 09:21:29.210754198 -0700 PDT", ReplyFrom:"callee!"}
stats = &nsq.ConsumerStats{MessagesReceived:0x1, MessagesFinished:0x0, MessagesRequeued:0x0, Connections:1}
caller done with response='&main.MyMsg{Hello:"hello at 2015-09-30 09:21:11.391440446 -0700 PDT", ReplyTo:"caller-call-number-7a50a6cfb60ce77fd54970972e456acc", Reply:"replying at 2015-09-30 09:21:29.210754198 -0700 PDT", ReplyFrom:"callee!"}'.
$
~~~

shell#3
-------

~~~
$ ./caller
caller starting.
2015/09/30 09:21:19 INF    1 (127.0.0.1:4150) connecting to nsqd
published 'main.MyMsg{Hello:"hello at 2015-09-30 09:21:19.849700212 -0700 PDT", ReplyTo:"caller-call-number-845002f5f63c29043cf0fd63a196cedc", Reply:"", ReplyFrom:""}'
2015/09/30 09:21:19 INF    1 stopping
2015/09/30 09:21:19 INF    1 exiting router
2015/09/30 09:21:19 INF    2 [caller-call-number-845002f5f63c29043cf0fd63a196cedc/ch] (127.0.0.1:4150) connecting to nsqd
2015/09/30 09:21:19 ERR    1 (127.0.0.1:4150) IO error - EOF
2015/09/30 09:21:19 INF    1 (127.0.0.1:4150) beginning close
2015/09/30 09:21:19 INF    1 (127.0.0.1:4150) readLoop exiting
2015/09/30 09:21:19 INF    1 (127.0.0.1:4150) breaking out of writeLoop
2015/09/30 09:21:19 INF    1 (127.0.0.1:4150) writeLoop exiting
2015/09/30 09:21:19 INF    1 (127.0.0.1:4150) finished draining, cleanup exiting
2015/09/30 09:21:19 INF    1 (127.0.0.1:4150) clean close complete
<hangs forever here -- since the one callee has consumed both calls>
~~~

shell#4 : the Consumer gets both messages, instead of just one. The caller in shell#3 thus never gets a reply.
-------

~~~
$ ./callee
2015/09/30 09:21:29 INF    1 [write_test/ch] (127.0.0.1:4150) connecting to nsqd
2015/09/30 09:21:29 callee.ConsumeRequest(): Got a message Body (as MyMsg): main.MyMsg{Hello:"hello at 2015-09-30 09:21:11.391440446 -0700 PDT", ReplyTo:"caller-call-number-7a50a6cfb60ce77fd54970972e456acc", Reply:"", ReplyFrom:""}
stats = &nsq.ConsumerStats{MessagesReceived:0x1, MessagesFinished:0x0, MessagesRequeued:0x0, Connections:1}
2015/09/30 09:21:29 callee.ConsumeRequest(): Got a message Body (as MyMsg): main.MyMsg{Hello:"hello at 2015-09-30 09:21:19.849700212 -0700 PDT", ReplyTo:"caller-call-number-845002f5f63c29043cf0fd63a196cedc", Reply:"", ReplyFrom:""}
2015/09/30 09:21:29 INF    2 (127.0.0.1:4150) connecting to nsqd
callee.ProduceReply(): replied-to: 'caller-call-number-7a50a6cfb60ce77fd54970972e456acc' with value '&main.MyMsg{Hello:"hello at 2015-09-30 09:21:11.391440446 -0700 PDT", ReplyTo:"caller-call-number-7a50a6cfb60ce77fd54970972e456acc", Reply:"replying at 2015-09-30 09:21:29.210754198 -0700 PDT", ReplyFrom:"callee!"}'
2015/09/30 09:21:29 INF    2 stopping
$
~~~
