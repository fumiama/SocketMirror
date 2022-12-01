# SocketMirror
Send random bytes to remote port and/or reflect all data to sender

## Usage
```bash
  -a string
        server addr (default "127.0.0.1:8888")
  -c uint
        client count
  -h    display this help
  -l uint
        batch size (length per write) (default 1024)
  -n string
        network (default "tcp")
  -r uint
        batch count (repeat times) (default 1024)
  -s    server mode
  -t uint
        connect timeout (s) (default 4)
```

## Example
### 1 client
```bash
$ go run main.go -c 1 -s 
INFO [serv] start
INFO [serv] accept: 127.0.0.1:52065
INFO [0001] beam max: 81000ns, min: 1000ns, avg: 6650ns
INFO [0001] see max: 115000ns, min: 1000ns, avg: 7981ns
INFO [0001] 1024/1024 succ/totl, speed: 116.51 B/s
INFO [serv] close: 127.0.0.1:52065
^CINFO [serv] terminate
```
### 2 clients
```bash
$ go run main.go -c 2 -s 
INFO [serv] start
INFO [serv] accept: 127.0.0.1:52069
INFO [serv] accept: 127.0.0.1:52070
INFO [0001] beam max: 142000ns, min: 1000ns, avg: 6511ns
INFO [0002] beam max: 230000ns, min: 1000ns, avg: 6970ns
INFO [0002] see max: 685000ns, min: 1000ns, avg: 11204ns
INFO [0002] 1024/1024 succ/totl, speed: 80.66 B/s
INFO [0001] see max: 581000ns, min: 1000ns, avg: 11414ns
INFO [0001] 1024/1024 succ/totl, speed: 80.66 B/s
INFO [serv] close: 127.0.0.1:52069
INFO [serv] close: 127.0.0.1:52070
^CINFO [serv] terminate
```
### 16 clients
```bash
$ go run main.go -c 16 -s 
INFO [serv] start
INFO [serv] accept: 127.0.0.1:52079
INFO [serv] accept: 127.0.0.1:52081
INFO [serv] accept: 127.0.0.1:52080
INFO [serv] accept: 127.0.0.1:52082
INFO [serv] accept: 127.0.0.1:52083
INFO [serv] accept: 127.0.0.1:52084
INFO [serv] accept: 127.0.0.1:52085
INFO [serv] accept: 127.0.0.1:52086
INFO [serv] accept: 127.0.0.1:52088
INFO [serv] accept: 127.0.0.1:52087
INFO [serv] accept: 127.0.0.1:52089
INFO [serv] accept: 127.0.0.1:52090
INFO [serv] accept: 127.0.0.1:52091
INFO [serv] accept: 127.0.0.1:52092
INFO [serv] accept: 127.0.0.1:52093
INFO [serv] accept: 127.0.0.1:52094
INFO [0001] beam max: 6351000ns, min: 1000ns, avg: 33535ns
INFO [0011] beam max: 7036000ns, min: 1000ns, avg: 33804ns
INFO [0003] beam max: 7425000ns, min: 1000ns, avg: 34208ns
INFO [0005] beam max: 6982000ns, min: 1000ns, avg: 40286ns
INFO [0013] beam max: 11218000ns, min: 1000ns, avg: 44741ns
INFO [0007] beam max: 10889000ns, min: 1000ns, avg: 45435ns
INFO [0006] beam max: 19072000ns, min: 1000ns, avg: 55626ns
INFO [0009] beam max: 16861000ns, min: 2000ns, avg: 57989ns
INFO [0012] beam max: 18309000ns, min: 1000ns, avg: 58515ns
INFO [0010] beam max: 14017000ns, min: 1000ns, avg: 59731ns
INFO [0008] beam max: 17997000ns, min: 1000ns, avg: 59722ns
INFO [0015] beam max: 21077000ns, min: 1000ns, avg: 62428ns
INFO [0004] beam max: 20567000ns, min: 1000ns, avg: 63575ns
INFO [0016] beam max: 16053000ns, min: 1000ns, avg: 63631ns
INFO [0014] beam max: 18837000ns, min: 1000ns, avg: 64478ns
INFO [0002] beam max: 21226000ns, min: 1000ns, avg: 65386ns
INFO [0001] see max: 17874000ns, min: 1000ns, avg: 73396ns
INFO [0001] 1024/1024 succ/totl, speed: 13.80 B/s
INFO [0005] see max: 13627000ns, min: 1000ns, avg: 74666ns
INFO [0005] 1024/1024 succ/totl, speed: 13.44 B/s
INFO [0013] see max: 19922000ns, min: 1000ns, avg: 77708ns
INFO [0013] 1024/1024 succ/totl, speed: 13.11 B/s
INFO [0011] see max: 20118000ns, min: 1000ns, avg: 79485ns
INFO [0011] 1024/1024 succ/totl, speed: 12.79 B/s
INFO [0003] see max: 21568000ns, min: 1000ns, avg: 79935ns
INFO [0003] 1024/1024 succ/totl, speed: 12.48 B/s
INFO [0007] see max: 19129000ns, min: 1000ns, avg: 80188ns
INFO [0007] 1024/1024 succ/totl, speed: 12.63 B/s
INFO [serv] close: 127.0.0.1:52079
INFO [0009] see max: 19397000ns, min: 1000ns, avg: 86088ns
INFO [0009] 1024/1024 succ/totl, speed: 11.78 B/s
INFO [0012] see max: 18683000ns, min: 2000ns, avg: 86269ns
INFO [0012] 1024/1024 succ/totl, speed: 11.78 B/s
INFO [serv] close: 127.0.0.1:52081
INFO [0008] see max: 17419000ns, min: 1000ns, avg: 87190ns
INFO [0008] 1024/1024 succ/totl, speed: 11.65 B/s
INFO [0010] see max: 13955000ns, min: 1000ns, avg: 85844ns
INFO [0010] 1024/1024 succ/totl, speed: 11.78 B/s
INFO [0006] see max: 12871000ns, min: 2000ns, avg: 87612ns
INFO [0006] 1024/1024 succ/totl, speed: 11.52 B/s
INFO [0015] see max: 21273000ns, min: 1000ns, avg: 88927ns
INFO [0015] 1024/1024 succ/totl, speed: 11.40 B/s
INFO [serv] close: 127.0.0.1:52091
INFO [0004] see max: 20048000ns, min: 1000ns, avg: 87566ns
INFO [0004] 1024/1024 succ/totl, speed: 11.52 B/s
INFO [serv] close: 127.0.0.1:52082
INFO [serv] close: 127.0.0.1:52080
INFO [serv] close: 127.0.0.1:52092
INFO [0016] see max: 23821000ns, min: 1000ns, avg: 91605ns
INFO [0016] 1024/1024 succ/totl, speed: 11.04 B/s
INFO [0014] see max: 24620000ns, min: 1000ns, avg: 92108ns
INFO [0014] 1024/1024 succ/totl, speed: 11.04 B/s
INFO [0002] see max: 27602000ns, min: 1000ns, avg: 93453ns
INFO [0002] 1024/1024 succ/totl, speed: 10.81 B/s
INFO [serv] close: 127.0.0.1:52090
INFO [serv] close: 127.0.0.1:52089
INFO [serv] close: 127.0.0.1:52088
INFO [serv] close: 127.0.0.1:52094
INFO [serv] close: 127.0.0.1:52087
INFO [serv] close: 127.0.0.1:52085
INFO [serv] close: 127.0.0.1:52093
INFO [serv] close: 127.0.0.1:52086
INFO [serv] close: 127.0.0.1:52084
INFO [serv] close: 127.0.0.1:52083
^CINFO [serv] terminate
```
### 64 clitens
```bash
$ go run main.go -c 64 -s 
INFO [serv] start
INFO [serv] accept: 127.0.0.1:52097
INFO [serv] accept: 127.0.0.1:52098
INFO [serv] accept: 127.0.0.1:52099
INFO [serv] accept: 127.0.0.1:52100
INFO [serv] accept: 127.0.0.1:52101
INFO [serv] accept: 127.0.0.1:52102
INFO [serv] accept: 127.0.0.1:52103
INFO [serv] accept: 127.0.0.1:52104
INFO [serv] accept: 127.0.0.1:52105
INFO [serv] accept: 127.0.0.1:52106
INFO [serv] accept: 127.0.0.1:52107
INFO [serv] accept: 127.0.0.1:52108
INFO [serv] accept: 127.0.0.1:52109
INFO [serv] accept: 127.0.0.1:52110
INFO [serv] accept: 127.0.0.1:52111
INFO [serv] accept: 127.0.0.1:52112
INFO [serv] accept: 127.0.0.1:52113
INFO [serv] accept: 127.0.0.1:52114
INFO [serv] accept: 127.0.0.1:52115
INFO [serv] accept: 127.0.0.1:52116
INFO [serv] accept: 127.0.0.1:52117
INFO [serv] accept: 127.0.0.1:52118
INFO [serv] accept: 127.0.0.1:52119
INFO [serv] accept: 127.0.0.1:52120
INFO [serv] accept: 127.0.0.1:52121
INFO [serv] accept: 127.0.0.1:52122
INFO [serv] accept: 127.0.0.1:52123
INFO [serv] accept: 127.0.0.1:52124
INFO [serv] accept: 127.0.0.1:52125
INFO [serv] accept: 127.0.0.1:52126
INFO [serv] accept: 127.0.0.1:52127
INFO [serv] accept: 127.0.0.1:52128
INFO [serv] accept: 127.0.0.1:52129
INFO [serv] accept: 127.0.0.1:52130
INFO [serv] accept: 127.0.0.1:52131
INFO [serv] accept: 127.0.0.1:52132
INFO [serv] accept: 127.0.0.1:52133
INFO [serv] accept: 127.0.0.1:52134
INFO [serv] accept: 127.0.0.1:52135
INFO [serv] accept: 127.0.0.1:52136
INFO [serv] accept: 127.0.0.1:52137
INFO [serv] accept: 127.0.0.1:52138
INFO [serv] accept: 127.0.0.1:52139
INFO [serv] accept: 127.0.0.1:52140
INFO [serv] accept: 127.0.0.1:52141
INFO [serv] accept: 127.0.0.1:52142
INFO [serv] accept: 127.0.0.1:52143
INFO [serv] accept: 127.0.0.1:52144
INFO [serv] accept: 127.0.0.1:52145
INFO [serv] accept: 127.0.0.1:52146
INFO [serv] accept: 127.0.0.1:52147
INFO [serv] accept: 127.0.0.1:52148
INFO [serv] accept: 127.0.0.1:52149
INFO [serv] accept: 127.0.0.1:52150
INFO [serv] accept: 127.0.0.1:52151
INFO [serv] accept: 127.0.0.1:52152
INFO [serv] accept: 127.0.0.1:52153
INFO [serv] accept: 127.0.0.1:52154
INFO [serv] accept: 127.0.0.1:52155
INFO [serv] accept: 127.0.0.1:52156
INFO [serv] accept: 127.0.0.1:52157
INFO [serv] accept: 127.0.0.1:52158
INFO [serv] accept: 127.0.0.1:52159
INFO [serv] accept: 127.0.0.1:52160
INFO [0001] beam max: 25085000ns, min: 2000ns, avg: 77621ns
INFO [0011] beam max: 30785000ns, min: 2000ns, avg: 109412ns
INFO [0009] beam max: 42294000ns, min: 1000ns, avg: 124561ns
INFO [0002] beam max: 70155000ns, min: 1000ns, avg: 164291ns
INFO [0025] beam max: 68533000ns, min: 1000ns, avg: 171348ns
INFO [0052] beam max: 68596000ns, min: 1000ns, avg: 172236ns
INFO [0053] beam max: 51584000ns, min: 1000ns, avg: 187089ns
INFO [0013] beam max: 73511000ns, min: 1000ns, avg: 192404ns
INFO [0007] beam max: 87223000ns, min: 1000ns, avg: 202881ns
INFO [0040] beam max: 80784000ns, min: 1000ns, avg: 215581ns
INFO [0014] beam max: 54692000ns, min: 1000ns, avg: 216950ns
INFO [0028] beam max: 58864000ns, min: 1000ns, avg: 272169ns
INFO [0004] beam max: 87795000ns, min: 1000ns, avg: 352485ns
INFO [0026] beam max: 93333000ns, min: 1000ns, avg: 358012ns
INFO [0039] beam max: 99185000ns, min: 1000ns, avg: 358472ns
INFO [0008] beam max: 86983000ns, min: 1000ns, avg: 360895ns
INFO [0038] beam max: 102204000ns, min: 1000ns, avg: 362749ns
INFO [0041] beam max: 64776000ns, min: 1000ns, avg: 363047ns
INFO [0031] beam max: 106251000ns, min: 1000ns, avg: 363321ns
INFO [0054] beam max: 57921000ns, min: 1000ns, avg: 363681ns
INFO [0015] beam max: 60037000ns, min: 1000ns, avg: 365023ns
INFO [0061] beam max: 103827000ns, min: 1000ns, avg: 366789ns
INFO [0032] beam max: 98784000ns, min: 1000ns, avg: 370187ns
INFO [0046] beam max: 92999000ns, min: 1000ns, avg: 374117ns
INFO [0019] beam max: 90544000ns, min: 1000ns, avg: 375835ns
INFO [0027] beam max: 80293000ns, min: 1000ns, avg: 377831ns
INFO [0020] beam max: 90796000ns, min: 1000ns, avg: 380837ns
INFO [0060] beam max: 85144000ns, min: 1000ns, avg: 379298ns
INFO [0036] beam max: 99840000ns, min: 1000ns, avg: 378957ns
INFO [0006] beam max: 105542000ns, min: 1000ns, avg: 380222ns
INFO [0012] beam max: 105410000ns, min: 1000ns, avg: 380134ns
INFO [0048] beam max: 100679000ns, min: 1000ns, avg: 380306ns
INFO [0018] beam max: 81925000ns, min: 1000ns, avg: 380267ns
INFO [0023] beam max: 97382000ns, min: 1000ns, avg: 381778ns
INFO [0011] see max: 95991000ns, min: 1000ns, avg: 388672ns
INFO [0011] 1024/1024 succ/totl, speed: 2.63 B/s
INFO [0056] beam max: 88873000ns, min: 1000ns, avg: 390975ns
INFO [0043] beam max: 93634000ns, min: 1000ns, avg: 395416ns
INFO [0009] see max: 105263000ns, min: 1000ns, avg: 410207ns
INFO [0009] 1024/1024 succ/totl, speed: 2.48 B/s
INFO [0016] beam max: 100365000ns, min: 1000ns, avg: 401173ns
INFO [serv] close: 127.0.0.1:52106
INFO [0030] beam max: 61372000ns, min: 1000ns, avg: 418475ns
INFO [0059] beam max: 62190000ns, min: 1000ns, avg: 425320ns
INFO [0062] beam max: 80471000ns, min: 1000ns, avg: 433701ns
INFO [0063] beam max: 92342000ns, min: 1000ns, avg: 436858ns
INFO [0034] beam max: 94282000ns, min: 1000ns, avg: 438970ns
INFO [0022] beam max: 97496000ns, min: 1000ns, avg: 441803ns
INFO [0005] beam max: 98574000ns, min: 1000ns, avg: 443867ns
INFO [0064] beam max: 93832000ns, min: 1000ns, avg: 443810ns
INFO [serv] close: 127.0.0.1:52099
INFO [0024] beam max: 100531000ns, min: 1000ns, avg: 448327ns
INFO [0049] beam max: 101558000ns, min: 1000ns, avg: 448698ns
INFO [0058] beam max: 106234000ns, min: 1000ns, avg: 456021ns
INFO [0010] beam max: 105759000ns, min: 1000ns, avg: 456403ns
INFO [0021] beam max: 87621000ns, min: 1000ns, avg: 459405ns
INFO [0051] beam max: 101922000ns, min: 1000ns, avg: 459848ns
INFO [0003] beam max: 102147000ns, min: 1000ns, avg: 460416ns
INFO [0033] beam max: 70541000ns, min: 1000ns, avg: 460834ns
INFO [0047] beam max: 71998000ns, min: 1000ns, avg: 461954ns
INFO [0042] beam max: 104577000ns, min: 1000ns, avg: 470958ns
INFO [0001] see max: 94221000ns, min: 1000ns, avg: 481074ns
INFO [0001] 1024/1024 succ/totl, speed: 2.12 B/s
INFO [0028] see max: 79509000ns, min: 1000ns, avg: 485651ns
INFO [0028] 1024/1024 succ/totl, speed: 2.10 B/s
INFO [0055] beam max: 98430000ns, min: 1000ns, avg: 479141ns
INFO [0002] see max: 90518000ns, min: 1000ns, avg: 487350ns
INFO [0002] 1024/1024 succ/totl, speed: 2.10 B/s
INFO [0025] see max: 88983000ns, min: 1000ns, avg: 487798ns
INFO [0025] 1024/1024 succ/totl, speed: 2.09 B/s
INFO [0029] beam max: 97381000ns, min: 1000ns, avg: 481743ns
INFO [0052] see max: 80405000ns, min: 1000ns, avg: 494515ns
INFO [0052] 1024/1024 succ/totl, speed: 2.06 B/s
INFO [0013] see max: 70365000ns, min: 1000ns, avg: 505919ns
INFO [0013] 1024/1024 succ/totl, speed: 2.02 B/s
INFO [0044] beam max: 73964000ns, min: 1000ns, avg: 506764ns
INFO [0014] see max: 91863000ns, min: 1000ns, avg: 513695ns
INFO [0014] 1024/1024 succ/totl, speed: 1.99 B/s
INFO [serv] close: 127.0.0.1:52108
INFO [0053] see max: 51882000ns, min: 1000ns, avg: 524438ns
INFO [0053] 1024/1024 succ/totl, speed: 1.95 B/s
INFO [serv] close: 127.0.0.1:52124
INFO [0035] beam max: 89837000ns, min: 1000ns, avg: 523154ns
INFO [serv] close: 127.0.0.1:52101
INFO [serv] close: 127.0.0.1:52115
INFO [serv] close: 127.0.0.1:52114
INFO [0037] beam max: 102437000ns, min: 1000ns, avg: 535389ns
INFO [0050] beam max: 106308000ns, min: 1000ns, avg: 537717ns
INFO [serv] close: 127.0.0.1:52098
INFO [serv] close: 127.0.0.1:52119
INFO [0043] see max: 98427000ns, min: 1000ns, avg: 552599ns
INFO [0043] 1024/1024 succ/totl, speed: 1.85 B/s
INFO [serv] close: 127.0.0.1:52118
INFO [0015] see max: 59250000ns, min: 1000ns, avg: 559255ns
INFO [0015] 1024/1024 succ/totl, speed: 1.83 B/s
INFO [0055] see max: 94350000ns, min: 1000ns, avg: 557822ns
INFO [0055] 1024/1024 succ/totl, speed: 1.83 B/s
INFO [0042] see max: 97643000ns, min: 1000ns, avg: 558122ns
INFO [0042] 1024/1024 succ/totl, speed: 1.83 B/s
INFO [0029] see max: 97409000ns, min: 1000ns, avg: 561120ns
INFO [0029] 1024/1024 succ/totl, speed: 1.82 B/s
INFO [0056] see max: 78050000ns, min: 1000ns, avg: 561671ns
INFO [0056] 1024/1024 succ/totl, speed: 1.82 B/s
INFO [0041] see max: 58761000ns, min: 1000ns, avg: 565628ns
INFO [0041] 1024/1024 succ/totl, speed: 1.81 B/s
INFO [0016] see max: 103861000ns, min: 1000ns, avg: 565910ns
INFO [0016] 1024/1024 succ/totl, speed: 1.80 B/s
INFO [0007] see max: 84152000ns, min: 1000ns, avg: 572649ns
INFO [0007] 1024/1024 succ/totl, speed: 1.78 B/s
INFO [0004] see max: 85625000ns, min: 1000ns, avg: 575756ns
INFO [0004] 1024/1024 succ/totl, speed: 1.77 B/s
INFO [0030] see max: 66733000ns, min: 1000ns, avg: 574741ns
INFO [0030] 1024/1024 succ/totl, speed: 1.78 B/s
INFO [0039] see max: 101197000ns, min: 1000ns, avg: 579200ns
INFO [0039] 1024/1024 succ/totl, speed: 1.76 B/s
INFO [0057] beam max: 274554000ns, min: 1000ns, avg: 574680ns
INFO [0026] see max: 103912000ns, min: 1000ns, avg: 579589ns
INFO [0026] 1024/1024 succ/totl, speed: 1.77 B/s
INFO [0008] see max: 104001000ns, min: 1000ns, avg: 580146ns
INFO [0008] 1024/1024 succ/totl, speed: 1.76 B/s
INFO [0054] see max: 57638000ns, min: 1000ns, avg: 581012ns
INFO [0054] 1024/1024 succ/totl, speed: 1.76 B/s
INFO [0060] see max: 85181000ns, min: 1000ns, avg: 578956ns
INFO [0060] 1024/1024 succ/totl, speed: 1.77 B/s
INFO [0018] see max: 82636000ns, min: 1000ns, avg: 580489ns
INFO [0018] 1024/1024 succ/totl, speed: 1.76 B/s
INFO [0044] see max: 73932000ns, min: 1000ns, avg: 580537ns
INFO [0044] 1024/1024 succ/totl, speed: 1.76 B/s
INFO [0019] see max: 93115000ns, min: 1000ns, avg: 580464ns
INFO [0019] 1024/1024 succ/totl, speed: 1.76 B/s
INFO [0038] see max: 102329000ns, min: 1000ns, avg: 584290ns
INFO [0038] 1024/1024 succ/totl, speed: 1.75 B/s
INFO [0031] see max: 110023000ns, min: 1000ns, avg: 582327ns
INFO [0031] 1024/1024 succ/totl, speed: 1.75 B/s
INFO [0046] see max: 95731000ns, min: 1000ns, avg: 582686ns
INFO [0046] 1024/1024 succ/totl, speed: 1.75 B/s
INFO [0061] see max: 104262000ns, min: 1000ns, avg: 583283ns
INFO [0061] 1024/1024 succ/totl, speed: 1.75 B/s
INFO [0032] see max: 101557000ns, min: 1000ns, avg: 584193ns
INFO [0032] 1024/1024 succ/totl, speed: 1.75 B/s
INFO [serv] close: 127.0.0.1:52129
INFO [0020] see max: 91306000ns, min: 1000ns, avg: 584593ns
INFO [0020] 1024/1024 succ/totl, speed: 1.75 B/s
INFO [0062] see max: 78305000ns, min: 1000ns, avg: 583737ns
INFO [0062] 1024/1024 succ/totl, speed: 1.75 B/s
INFO [0047] see max: 73810000ns, min: 1000ns, avg: 583455ns
INFO [0047] 1024/1024 succ/totl, speed: 1.75 B/s
INFO [0033] see max: 73814000ns, min: 1000ns, avg: 584791ns
INFO [0033] 1024/1024 succ/totl, speed: 1.74 B/s
INFO [0027] see max: 77345000ns, min: 1000ns, avg: 588718ns
INFO [0027] 1024/1024 succ/totl, speed: 1.74 B/s
INFO [serv] close: 127.0.0.1:52123
INFO [0063] see max: 91199000ns, min: 1000ns, avg: 586226ns
INFO [0063] 1024/1024 succ/totl, speed: 1.74 B/s
INFO [0021] see max: 87503000ns, min: 2000ns, avg: 587517ns
INFO [0021] 1024/1024 succ/totl, speed: 1.74 B/s
INFO [0034] see max: 94364000ns, min: 1000ns, avg: 588337ns
INFO [0034] 1024/1024 succ/totl, speed: 1.74 B/s
INFO [0022] see max: 97526000ns, min: 1000ns, avg: 589678ns
INFO [0022] 1024/1024 succ/totl, speed: 1.73 B/s
INFO [0005] see max: 105067000ns, min: 1000ns, avg: 589375ns
INFO [serv] close: 127.0.0.1:52126
INFO [0005] 1024/1024 succ/totl, speed: 1.73 B/s
INFO [serv] close: 127.0.0.1:52125
INFO [0035] see max: 89712000ns, min: 1000ns, avg: 590794ns
INFO [0035] 1024/1024 succ/totl, speed: 1.73 B/s
INFO [0064] see max: 93962000ns, min: 1000ns, avg: 591576ns
INFO [0064] 1024/1024 succ/totl, speed: 1.72 B/s
INFO [0023] see max: 97253000ns, min: 1000ns, avg: 591986ns
INFO [0023] 1024/1024 succ/totl, speed: 1.72 B/s
INFO [serv] close: 127.0.0.1:52128
INFO [0048] see max: 100665000ns, min: 1000ns, avg: 592715ns
INFO [0048] 1024/1024 succ/totl, speed: 1.72 B/s
INFO [serv] close: 127.0.0.1:52130
INFO [serv] close: 127.0.0.1:52121
INFO [0049] see max: 101416000ns, min: 1000ns, avg: 594446ns
INFO [0049] 1024/1024 succ/totl, speed: 1.72 B/s
INFO [0036] see max: 102826000ns, min: 1000ns, avg: 595456ns
INFO [0036] 1024/1024 succ/totl, speed: 1.71 B/s
INFO [serv] close: 127.0.0.1:52127
INFO [0024] see max: 100615000ns, min: 1000ns, avg: 594973ns
INFO [0024] 1024/1024 succ/totl, speed: 1.72 B/s
INFO [0012] see max: 105463000ns, min: 1000ns, avg: 600866ns
INFO [0012] 1024/1024 succ/totl, speed: 1.70 B/s
INFO [serv] close: 127.0.0.1:52107
INFO [0058] see max: 106241000ns, min: 1000ns, avg: 601498ns
INFO [0058] 1024/1024 succ/totl, speed: 1.70 B/s
INFO [0006] see max: 105822000ns, min: 1000ns, avg: 597886ns
INFO [0006] 1024/1024 succ/totl, speed: 1.71 B/s
INFO [0010] see max: 105880000ns, min: 1000ns, avg: 597345ns
INFO [0010] 1024/1024 succ/totl, speed: 1.71 B/s
INFO [0037] see max: 103917000ns, min: 1000ns, avg: 597651ns
INFO [0037] 1024/1024 succ/totl, speed: 1.71 B/s
INFO [serv] close: 127.0.0.1:52097
INFO [0050] see max: 103507000ns, min: 1000ns, avg: 598216ns
INFO [0050] 1024/1024 succ/totl, speed: 1.71 B/s
INFO [0051] see max: 100756000ns, min: 1000ns, avg: 606077ns
INFO [0051] 1024/1024 succ/totl, speed: 1.68 B/s
INFO [0003] see max: 102328000ns, min: 1000ns, avg: 600273ns
INFO [0003] 1024/1024 succ/totl, speed: 1.70 B/s
INFO [serv] close: 127.0.0.1:52132
INFO [serv] close: 127.0.0.1:52116
INFO [serv] close: 127.0.0.1:52113
INFO [serv] close: 127.0.0.1:52100
INFO [serv] close: 127.0.0.1:52122
INFO [serv] close: 127.0.0.1:52139
INFO [serv] close: 127.0.0.1:52136
INFO [serv] close: 127.0.0.1:52134
INFO [serv] close: 127.0.0.1:52140
INFO [serv] close: 127.0.0.1:52112
INFO [serv] close: 127.0.0.1:52137
INFO [serv] close: 127.0.0.1:52141
INFO [serv] close: 127.0.0.1:52143
INFO [serv] close: 127.0.0.1:52142
INFO [serv] close: 127.0.0.1:52144
INFO [serv] close: 127.0.0.1:52147
INFO [serv] close: 127.0.0.1:52146
INFO [serv] close: 127.0.0.1:52145
INFO [serv] close: 127.0.0.1:52120
INFO [serv] close: 127.0.0.1:52149
INFO [serv] close: 127.0.0.1:52148
INFO [serv] close: 127.0.0.1:52150
INFO [serv] close: 127.0.0.1:52104
INFO [serv] close: 127.0.0.1:52151
INFO [serv] close: 127.0.0.1:52152
INFO [serv] close: 127.0.0.1:52153
INFO [serv] close: 127.0.0.1:52154
INFO [serv] close: 127.0.0.1:52155
INFO [serv] close: 127.0.0.1:52156
INFO [serv] close: 127.0.0.1:52158
INFO [serv] close: 127.0.0.1:52157
INFO [serv] close: 127.0.0.1:52111
INFO [serv] close: 127.0.0.1:52110
INFO [serv] close: 127.0.0.1:52105
INFO [serv] close: 127.0.0.1:52103
INFO [serv] close: 127.0.0.1:52159
INFO [0040] see max: 82922000ns, min: 1000ns, avg: 622034ns
INFO [0040] 1024/1024 succ/totl, speed: 1.64 B/s
INFO [serv] close: 127.0.0.1:52160
INFO [serv] close: 127.0.0.1:52109
INFO [serv] close: 127.0.0.1:52102
INFO [0045] beam max: 371556000ns, min: 1000ns, avg: 621930ns
INFO [serv] close: 127.0.0.1:52117
INFO [0057] see max: 82003000ns, min: 1000ns, avg: 624325ns
INFO [0057] 1024/1024 succ/totl, speed: 1.63 B/s
INFO [serv] close: 127.0.0.1:52135
INFO [0017] beam max: 76009000ns, min: 1000ns, avg: 621469ns
INFO [0045] see max: 97873000ns, min: 1000ns, avg: 627034ns
INFO [0045] 1024/1024 succ/totl, speed: 1.63 B/s
INFO [serv] close: 127.0.0.1:52138
INFO [0017] see max: 76320000ns, min: 1000ns, avg: 627432ns
INFO [0017] 1024/1024 succ/totl, speed: 1.63 B/s
INFO [serv] close: 127.0.0.1:52131
INFO [0059] see max: 98142000ns, min: 1000ns, avg: 722076ns
INFO [0059] 1024/1024 succ/totl, speed: 1.42 B/s
INFO [serv] close: 127.0.0.1:52133
^CINFO [serv] terminate
```
