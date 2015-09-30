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
$  nsqd --lookupd-tcp-address=127.0.0.1:4160
[nsqd] 2015/09/30 09:51:22.815042 nsqd v0.3.6 (built w/go1.5.1)
[nsqd] 2015/09/30 09:51:22.815128 ID: 842
[nsqd] 2015/09/30 09:51:22.815174 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqd] 2015/09/30 09:51:22.860048 LOOKUP(127.0.0.1:4160): adding peer
[nsqd] 2015/09/30 09:51:22.860079 LOOKUP connecting to 127.0.0.1:4160
[nsqd] 2015/09/30 09:51:22.860406 TCP: listening on [::]:4150
[nsqd] 2015/09/30 09:51:22.860457 HTTP: listening on [::]:4151
[nsqlookupd] 2015/09/30 09:51:22.860663 TCP: new client(127.0.0.1:45816)
[nsqlookupd] 2015/09/30 09:51:22.860693 CLIENT(127.0.0.1:45816): desired protocol magic '  V1'
[nsqlookupd] 2015/09/30 09:51:22.860769 CLIENT(127.0.0.1:45816): IDENTIFY Address:i7 TCP:4150 HTTP:4151 Version:0.3.6
[nsqlookupd] 2015/09/30 09:51:22.860785 DB: client(127.0.0.1:45816) REGISTER category:client key: subkey:
[nsqd] 2015/09/30 09:51:22.861144 LOOKUPD(127.0.0.1:4160): peer info {TCPPort:4160 HTTPPort:4161 Version:0.3.6 BroadcastAddress:i7}
[nsqd] 2015/09/30 09:51:26.452143 200 GET /stats (127.0.0.1:56164) 130.314µs
[nsqd] 2015/09/30 09:51:37.860195 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:51:37.860366 CLIENT(127.0.0.1:45816): pinged (last ping 14.999588725s)
[nsqd] 2015/09/30 09:51:38.301164 TCP: new client(127.0.0.1:42360)
[nsqd] 2015/09/30 09:51:38.301183 CLIENT(127.0.0.1:42360): desired protocol magic '  V2'
[nsqd] 2015/09/30 09:51:38.301417 [127.0.0.1:42360] IDENTIFY: {ShortID:i7 LongID:i7 ClientID:i7 Hostname:i7 HeartbeatInterval:30000 OutputBufferSize:16384 OutputBufferTimeout:250 FeatureNegotiation:true TLSv1:false Deflate:false DeflateLevel:6 Snappy:false SampleRate:0 UserAgent:go-nsq/1.0.5 MsgTimeout:0}
[nsqd] 2015/09/30 09:51:38.301742 TOPIC(write_test): created
[nsqd] 2015/09/30 09:51:38.301765 CI: querying nsqlookupd http://i7:4161/channels?topic=write_test
[nsqd] 2015/09/30 09:51:38.302034 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqd] 2015/09/30 09:51:38.302053 LOOKUPD(127.0.0.1:4160): topic REGISTER write_test
[nsqlookupd] 2015/09/30 09:51:38.302165 DB: client(127.0.0.1:45816) REGISTER category:topic key:write_test subkey:
[nsqd] 2015/09/30 09:51:40.303503 TCP: new client(127.0.0.1:42362)
[nsqd] 2015/09/30 09:51:40.303533 CLIENT(127.0.0.1:42362): desired protocol magic '  V2'
[nsqd] 2015/09/30 09:51:40.303843 [127.0.0.1:42362] IDENTIFY: {ShortID:i7 LongID:i7 ClientID:i7 Hostname:i7 HeartbeatInterval:30000 OutputBufferSize:16384 OutputBufferTimeout:250 FeatureNegotiation:true TLSv1:false Deflate:false DeflateLevel:6 Snappy:false SampleRate:0 UserAgent:go-nsq/1.0.5 MsgTimeout:0}
[nsqd] 2015/09/30 09:51:40.339732 TOPIC(caller-call-number-9b955616adb00c399f4e3bb75e32dd7f): created
[nsqd] 2015/09/30 09:51:40.339779 CI: querying nsqlookupd http://i7:4161/channels?topic=caller-call-number-9b955616adb00c399f4e3bb75e32dd7f
[nsqd] 2015/09/30 09:51:40.339794 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqd] 2015/09/30 09:51:40.339836 LOOKUPD(127.0.0.1:4160): topic REGISTER caller-call-number-9b955616adb00c399f4e3bb75e32dd7f
[nsqlookupd] 2015/09/30 09:51:40.340009 DB: client(127.0.0.1:45816) REGISTER category:topic key:caller-call-number-9b955616adb00c399f4e3bb75e32dd7f subkey:
[nsqd] 2015/09/30 09:51:40.403082 PROTOCOL(V2): [127.0.0.1:42360] exiting ioloop
[nsqd] 2015/09/30 09:51:40.403246 PROTOCOL(V2): [127.0.0.1:42360] exiting messagePump
[nsqd] 2015/09/30 09:51:41.299180 TOPIC(caller-call-number-9b955616adb00c399f4e3bb75e32dd7f): new channel(ch)
[nsqd] 2015/09/30 09:51:41.299544 LOOKUPD(127.0.0.1:4160): channel REGISTER caller-call-number-9b955616adb00c399f4e3bb75e32dd7f ch
[nsqlookupd] 2015/09/30 09:51:41.299754 DB: client(127.0.0.1:45816) REGISTER category:channel key:caller-call-number-9b955616adb00c399f4e3bb75e32dd7f subkey:ch
[nsqd] 2015/09/30 09:51:41.363409 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqd] 2015/09/30 09:51:44.780450 TCP: new client(127.0.0.1:42364)
[nsqd] 2015/09/30 09:51:44.780526 CLIENT(127.0.0.1:42364): desired protocol magic '  V2'
[nsqd] 2015/09/30 09:51:44.780922 [127.0.0.1:42364] IDENTIFY: {ShortID:i7 LongID:i7 ClientID:i7 Hostname:i7 HeartbeatInterval:30000 OutputBufferSize:16384 OutputBufferTimeout:250 FeatureNegotiation:true TLSv1:false Deflate:false DeflateLevel:6 Snappy:false SampleRate:0 UserAgent:go-nsq/1.0.5 MsgTimeout:0}
[nsqd] 2015/09/30 09:51:44.782254 TCP: new client(127.0.0.1:42365)
[nsqd] 2015/09/30 09:51:44.782297 CLIENT(127.0.0.1:42365): desired protocol magic '  V2'
[nsqd] 2015/09/30 09:51:44.782588 [127.0.0.1:42365] IDENTIFY: {ShortID:i7 LongID:i7 ClientID:i7 Hostname:i7 HeartbeatInterval:30000 OutputBufferSize:16384 OutputBufferTimeout:250 FeatureNegotiation:true TLSv1:false Deflate:false DeflateLevel:6 Snappy:false SampleRate:0 UserAgent:go-nsq/1.0.5 MsgTimeout:0}
[nsqd] 2015/09/30 09:51:44.782927 TOPIC(caller-call-number-0dea18ffca4a1ad50b25f915cc039803): created
[nsqd] 2015/09/30 09:51:44.782969 CI: querying nsqlookupd http://i7:4161/channels?topic=caller-call-number-0dea18ffca4a1ad50b25f915cc039803
[nsqd] 2015/09/30 09:51:44.783006 LOOKUPD(127.0.0.1:4160): topic REGISTER caller-call-number-0dea18ffca4a1ad50b25f915cc039803
[nsqd] 2015/09/30 09:51:44.782975 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqlookupd] 2015/09/30 09:51:44.783200 DB: client(127.0.0.1:45816) REGISTER category:topic key:caller-call-number-0dea18ffca4a1ad50b25f915cc039803 subkey:
[nsqd] 2015/09/30 09:51:44.882055 PROTOCOL(V2): [127.0.0.1:42364] exiting ioloop
[nsqd] 2015/09/30 09:51:44.882195 PROTOCOL(V2): [127.0.0.1:42364] exiting messagePump
[nsqd] 2015/09/30 09:51:46.783576 TOPIC(caller-call-number-0dea18ffca4a1ad50b25f915cc039803): new channel(ch)
[nsqd] 2015/09/30 09:51:46.783700 LOOKUPD(127.0.0.1:4160): channel REGISTER caller-call-number-0dea18ffca4a1ad50b25f915cc039803 ch
[nsqlookupd] 2015/09/30 09:51:46.783963 DB: client(127.0.0.1:45816) REGISTER category:channel key:caller-call-number-0dea18ffca4a1ad50b25f915cc039803 subkey:ch
[nsqd] 2015/09/30 09:51:46.831189 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqd] 2015/09/30 09:51:52.860198 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:51:52.860386 CLIENT(127.0.0.1:45816): pinged (last ping 15.000018332s)
[nsqd] 2015/09/30 09:52:07.443939 200 GET /stats (127.0.0.1:56186) 239.866µs
[nsqd] 2015/09/30 09:52:07.860157 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:52:07.860369 CLIENT(127.0.0.1:45816): pinged (last ping 14.999985465s)
[nsqd] 2015/09/30 09:52:14.609058 TCP: new client(127.0.0.1:42373)
[nsqd] 2015/09/30 09:52:14.609091 CLIENT(127.0.0.1:42373): desired protocol magic '  V2'
[nsqd] 2015/09/30 09:52:14.609495 [127.0.0.1:42373] IDENTIFY: {ShortID:i7 LongID:i7 ClientID:i7 Hostname:i7 HeartbeatInterval:30000 OutputBufferSize:16384 OutputBufferTimeout:250 FeatureNegotiation:true TLSv1:false Deflate:false DeflateLevel:6 Snappy:false SampleRate:0 UserAgent:go-nsq/1.0.5 MsgTimeout:0}
[nsqd] 2015/09/30 09:52:14.610020 TOPIC(write_test): new channel(ch)
[nsqd] 2015/09/30 09:52:14.610075 NSQ: persisting topic/channel metadata to nsqd.842.dat
[nsqd] 2015/09/30 09:52:14.610373 LOOKUPD(127.0.0.1:4160): channel REGISTER write_test ch
[nsqlookupd] 2015/09/30 09:52:14.610570 DB: client(127.0.0.1:45816) REGISTER category:channel key:write_test subkey:ch
[nsqd] 2015/09/30 09:52:14.611283 TCP: new client(127.0.0.1:42374)
[nsqd] 2015/09/30 09:52:14.611312 CLIENT(127.0.0.1:42374): desired protocol magic '  V2'
[nsqd] 2015/09/30 09:52:14.611610 [127.0.0.1:42374] IDENTIFY: {ShortID:i7 LongID:i7 ClientID:i7 Hostname:i7 HeartbeatInterval:30000 OutputBufferSize:16384 OutputBufferTimeout:250 FeatureNegotiation:true TLSv1:false Deflate:false DeflateLevel:6 Snappy:false SampleRate:0 UserAgent:go-nsq/1.0.5 MsgTimeout:0}
[nsqd] 2015/09/30 09:52:14.657893 PROTOCOL(V2): [127.0.0.1:42373] exiting ioloop
[nsqd] 2015/09/30 09:52:14.657997 PROTOCOL(V2): [127.0.0.1:42373] exiting messagePump
[nsqd] 2015/09/30 09:52:14.657895 PROTOCOL(V2): [127.0.0.1:42374] exiting ioloop
[nsqd] 2015/09/30 09:52:14.657902 PROTOCOL(V2): [127.0.0.1:42362] exiting ioloop
[nsqd] 2015/09/30 09:52:14.658141 PROTOCOL(V2): [127.0.0.1:42374] exiting messagePump
[nsqd] 2015/09/30 09:52:14.658236 PROTOCOL(V2): [127.0.0.1:42362] exiting messagePump
[nsqd] 2015/09/30 09:52:19.307115 200 GET /stats (127.0.0.1:56191) 192.122µs
[nsqd] 2015/09/30 09:52:22.860184 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:52:22.860388 CLIENT(127.0.0.1:45816): pinged (last ping 15.000019237s)
[nsqd] 2015/09/30 09:52:37.860174 LOOKUPD(127.0.0.1:4160): sending heartbeat
[nsqlookupd] 2015/09/30 09:52:37.860372 CLIENT(127.0.0.1:45816): pinged (last ping 14.99998356s)
...

# on separate shell, issue check to see that no old messages are hanging about:
$ wget -O - http://127.0.0.1:4151/stats
--2015-09-30 09:51:26--  http://127.0.0.1:4151/stats
Connecting to 127.0.0.1:4151... connected.
HTTP request sent, awaiting response... 200 OK
Length: 98 [text/plain]
Saving to: `STDOUT'

 0% [                                       ] 0           --.-K/s              nsqd v0.3.6 (built w/go1.5.1)
start_time 2015-09-30T09:51:22-07:00
uptime 3.637047471s

NO_TOPICS
100%[======================================>] 98          --.-K/s   in 0s      

2015-09-30 09:51:26 (14.1 MB/s) - written to stdout [98/98]

$
~~~

shell#2: start first caller
-------
~~~
$ ./caller
caller starting.
2015/09/30 09:51:38 INF    1 (127.0.0.1:4150) connecting to nsqd
published 'main.MyMsg{Hello:"hello at 2015-09-30 09:51:38.300817114 -0700 PDT", ReplyTo:"caller-call-number-9b955616adb00c399f4e3bb75e32dd7f", Reply:"", ReplyFrom:""}'
2015/09/30 09:51:40 INF    1 stopping
2015/09/30 09:51:40 ERR    1 (127.0.0.1:4150) IO error - EOF
2015/09/30 09:51:40 INF    1 exiting router
2015/09/30 09:51:40 INF    1 (127.0.0.1:4150) beginning close
2015/09/30 09:51:40 INF    1 (127.0.0.1:4150) readLoop exiting
2015/09/30 09:51:40 INF    1 (127.0.0.1:4150) breaking out of writeLoop
2015/09/30 09:51:40 INF    1 (127.0.0.1:4150) writeLoop exiting
2015/09/30 09:51:40 INF    2 [caller-call-number-9b955616adb00c399f4e3bb75e32dd7f/ch] (127.0.0.1:4150) connecting to nsqd
2015/09/30 09:51:40 INF    1 (127.0.0.1:4150) finished draining, cleanup exiting
2015/09/30 09:51:40 INF    1 (127.0.0.1:4150) clean close complete
2015/09/30 09:52:14 caller: Got a message Body (as MyMsg): main.MyMsg{Hello:"hello at 2015-09-30 09:51:38.300817114 -0700 PDT", ReplyTo:"caller-call-number-9b955616adb00c399f4e3bb75e32dd7f", Reply:"replying at 2015-09-30 09:52:14.610584567 -0700 PDT", ReplyFrom:"callee!"}
stats = &nsq.ConsumerStats{MessagesReceived:0x1, MessagesFinished:0x0, MessagesRequeued:0x0, Connections:1}
caller done with response='&main.MyMsg{Hello:"hello at 2015-09-30 09:51:38.300817114 -0700 PDT", ReplyTo:"caller-call-number-9b955616adb00c399f4e3bb75e32dd7f", Reply:"replying at 2015-09-30 09:52:14.610584567 -0700 PDT", ReplyFrom:"callee!"}'.
$  # comment: finished fine, got reply from callee once the callee was started.

~~~

shell#3: start 2nd caller
-------

~~~
$ ./caller
caller starting.
2015/09/30 09:51:44 INF    1 (127.0.0.1:4150) connecting to nsqd
published 'main.MyMsg{Hello:"hello at 2015-09-30 09:51:44.779999658 -0700 PDT", ReplyTo:"caller-call-number-0dea18ffca4a1ad50b25f915cc039803", Reply:"", ReplyFrom:""}'
2015/09/30 09:51:44 INF    1 stopping
2015/09/30 09:51:44 INF    1 exiting router
2015/09/30 09:51:44 ERR    1 (127.0.0.1:4150) IO error - EOF
2015/09/30 09:51:44 INF    1 (127.0.0.1:4150) beginning close
2015/09/30 09:51:44 INF    1 (127.0.0.1:4150) readLoop exiting
2015/09/30 09:51:44 INF    1 (127.0.0.1:4150) breaking out of writeLoop
2015/09/30 09:51:44 INF    1 (127.0.0.1:4150) writeLoop exiting
2015/09/30 09:51:44 INF    2 [caller-call-number-0dea18ffca4a1ad50b25f915cc039803/ch] (127.0.0.1:4150) connecting to nsqd
2015/09/30 09:51:44 INF    1 (127.0.0.1:4150) finished draining, cleanup exiting
2015/09/30 09:51:44 INF    1 (127.0.0.1:4150) clean close complete
<hangs forever here -- since the one callee has consumed both calls>
~~~

shell#4 : start one callee. Observe the bug: the callee Consumer gets both messages, instead of just one. The caller in shell#3 thus never gets a reply.
-------

 also check the stats before and after:

~~~
$ wget -O - http://127.0.0.1:4151/stats
--2015-09-30 09:52:07--  http://127.0.0.1:4151/stats
Connecting to 127.0.0.1:4151... connected.
HTTP request sent, awaiting response... 200 OK
Length: 909 [text/plain]
Saving to: `STDOUT'

 0% [                                       ] 0           --.-K/s              nsqd v0.3.6 (built w/go1.5.1)
start_time 2015-09-30T09:51:22-07:00
uptime 44.628783869s

Health: OK

   [caller-call-number-0dea18ffca4a1ad50b25f915cc039803] depth: 0     be-depth: 0     msgs: 0        e2e%: 
      [ch                       ] depth: 0     be-depth: 0     inflt: 0    def: 0    re-q: 0     timeout: 0     msgs: 0        e2e%: 
        [V2 i7:42365             ] state: 3 inflt: 0    rdy: 1    fin: 0        re-q: 0        msgs: 0        connected: 23s

   [caller-call-number-9b955616adb00c399f4e3bb75e32dd7f] depth: 0     be-depth: 0     msgs: 0        e2e%: 
      [ch                       ] depth: 0     be-depth: 0     inflt: 0    def: 0    re-q: 0     timeout: 0     msgs: 0        e2e%: 
        [V2 i7:42362             ] state: 3 inflt: 0    rdy: 1    fin: 0        re-q: 0        msgs: 0        connected: 27s

   [write_test     ] depth: 2     be-depth: 0     msgs: 2        e2e%: 
100%[======================================>] 909         --.-K/s   in 0s      

2015-09-30 09:52:07 (131 MB/s) - written to stdout [909/909]

$ ./callee
2015/09/30 09:52:14 INF    1 [write_test/ch] (127.0.0.1:4150) connecting to nsqd
2015/09/30 09:52:14 callee.ConsumeRequest(): Got a message Body (as MyMsg): main.MyMsg{Hello:"hello at 2015-09-30 09:51:38.300817114 -0700 PDT", ReplyTo:"caller-call-number-9b955616adb00c399f4e3bb75e32dd7f", Reply:"", ReplyFrom:""}
stats = &nsq.ConsumerStats{MessagesReceived:0x1, MessagesFinished:0x0, MessagesRequeued:0x0, Connections:1}
2015/09/30 09:52:14 callee.ConsumeRequest(): Got a message Body (as MyMsg): main.MyMsg{Hello:"hello at 2015-09-30 09:51:44.779999658 -0700 PDT", ReplyTo:"caller-call-number-0dea18ffca4a1ad50b25f915cc039803", Reply:"", ReplyFrom:""}
2015/09/30 09:52:14 INF    2 (127.0.0.1:4150) connecting to nsqd
callee.ProduceReply(): replied-to: 'caller-call-number-9b955616adb00c399f4e3bb75e32dd7f' with value '&main.MyMsg{Hello:"hello at 2015-09-30 09:51:38.300817114 -0700 PDT", ReplyTo:"caller-call-number-9b955616adb00c399f4e3bb75e32dd7f", Reply:"replying at 2015-09-30 09:52:14.610584567 -0700 PDT", ReplyFrom:"callee!"}'
2015/09/30 09:52:14 INF    2 stopping
2015/09/30 09:52:14 INF    2 exiting router
$ wget -O - http://127.0.0.1:4151/stats
--2015-09-30 09:52:19--  http://127.0.0.1:4151/stats
Connecting to 127.0.0.1:4151... connected.
HTTP request sent, awaiting response... 200 OK
Length: 918 [text/plain]
Saving to: `STDOUT'

 0% [                                       ] 0           --.-K/s              nsqd v0.3.6 (built w/go1.5.1)
start_time 2015-09-30T09:51:22-07:00
uptime 56.491967971s

Health: OK

   [caller-call-number-0dea18ffca4a1ad50b25f915cc039803] depth: 0     be-depth: 0     msgs: 0        e2e%: 
      [ch                       ] depth: 0     be-depth: 0     inflt: 0    def: 0    re-q: 0     timeout: 0     msgs: 0        e2e%: 
        [V2 i7:42365             ] state: 3 inflt: 0    rdy: 1    fin: 0        re-q: 0        msgs: 0        connected: 35s

   [caller-call-number-9b955616adb00c399f4e3bb75e32dd7f] depth: 0     be-depth: 0     msgs: 1        e2e%: 
      [ch                       ] depth: 0     be-depth: 0     inflt: 1    def: 0    re-q: 0     timeout: 0     msgs: 1        e2e%: 

   [write_test     ] depth: 0     be-depth: 0     msgs: 2        e2e%: 
      [ch                       ] depth: 0     be-depth: 0     inflt: 0    def: 0    re-q: 0     timeout: 0     msgs: 2        e2e%: 
100%[======================================>] 918         --.-K/s   in 0s      

2015-09-30 09:52:19 (137 MB/s) - written to stdout [918/918]
$
~~~
