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
[srv] start
[srv] accept: 127.0.0.1:50123
0001 -> beam max: 68000ns, min: 3000ns, avg: 6594ns
0001 -> see max: 104320000ns, min: 1000ns, avg: 108816ns
0001 -> 1024/1024 succ/totl, speed: 9.28 B/s
[srv] ERROR: read from 127.0.0.1:50123: EOF
[srv] close: 127.0.0.1:50123
^Csignal: interrupt
```
### 2 clients
```bash
$ go run main.go -c 2 -s 
[srv] start
[srv] accept: 127.0.0.1:50132
[srv] accept: 127.0.0.1:50133
0001 -> beam max: 179000ns, min: 1000ns, avg: 6760ns
0002 -> beam max: 171000ns, min: 1000ns, avg: 6952ns
0001 -> see max: 197000ns, min: 1000ns, avg: 11033ns
0001 -> 1024/1024 succ/totl, speed: 80.66 B/s
0002 -> see max: 480000ns, min: 1000ns, avg: 11297ns
0002 -> 1024/1024 succ/totl, speed: 80.66 B/s
[srv] ERROR: read from 127.0.0.1:50133: EOF
[srv] close: 127.0.0.1:50133
[srv] ERROR: read from 127.0.0.1:50132: EOF
[srv] close: 127.0.0.1:50132
^Csignal: interrupt
```
### 16 clients
```bash
$ go run main.go -c 16 -s 
[srv] start
[srv] accept: 127.0.0.1:50142
[srv] accept: 127.0.0.1:50143
[srv] accept: 127.0.0.1:50144
[srv] accept: 127.0.0.1:50145
[srv] accept: 127.0.0.1:50146
[srv] accept: 127.0.0.1:50147
[srv] accept: 127.0.0.1:50148
[srv] accept: 127.0.0.1:50149
[srv] accept: 127.0.0.1:50150
[srv] accept: 127.0.0.1:50151
[srv] accept: 127.0.0.1:50152
[srv] accept: 127.0.0.1:50153
[srv] accept: 127.0.0.1:50154
[srv] accept: 127.0.0.1:50155
[srv] accept: 127.0.0.1:50156
[srv] accept: 127.0.0.1:50157
0003 -> beam max: 9381000ns, min: 1000ns, avg: 40270ns
0001 -> beam max: 9383000ns, min: 1000ns, avg: 42046ns
0014 -> beam max: 8683000ns, min: 1000ns, avg: 44367ns
0009 -> beam max: 8860000ns, min: 1000ns, avg: 46335ns
0002 -> beam max: 10965000ns, min: 1000ns, avg: 49140ns
0005 -> beam max: 17798000ns, min: 1000ns, avg: 54858ns
0006 -> beam max: 11410000ns, min: 1000ns, avg: 57200ns
0013 -> beam max: 9618000ns, min: 1000ns, avg: 59171ns
0016 -> beam max: 11643000ns, min: 1000ns, avg: 61691ns
0008 -> beam max: 12875000ns, min: 1000ns, avg: 63643ns
0004 -> beam max: 23019000ns, min: 1000ns, avg: 65455ns
0007 -> beam max: 16039000ns, min: 1000ns, avg: 65698ns
0012 -> beam max: 17136000ns, min: 1000ns, avg: 66500ns
0010 -> beam max: 22858000ns, min: 1000ns, avg: 68531ns
0011 -> beam max: 11565000ns, min: 1000ns, avg: 73359ns
0001 -> see max: 13833000ns, min: 1000ns, avg: 75044ns
0001 -> 1024/1024 succ/totl, speed: 13.44 B/s
0014 -> see max: 20302000ns, min: 1000ns, avg: 76476ns
0014 -> 1024/1024 succ/totl, speed: 13.11 B/s
0003 -> see max: 24288000ns, min: 1000ns, avg: 78414ns
0003 -> 1024/1024 succ/totl, speed: 12.79 B/s
0002 -> see max: 20774000ns, min: 1000ns, avg: 79520ns
0002 -> 1024/1024 succ/totl, speed: 12.79 B/s
0015 -> beam max: 27679000ns, min: 1000ns, avg: 70354ns
0006 -> see max: 13379000ns, min: 1000ns, avg: 80461ns
0006 -> 1024/1024 succ/totl, speed: 12.63 B/s
0005 -> see max: 15717000ns, min: 1000ns, avg: 81780ns
0005 -> 1024/1024 succ/totl, speed: 12.34 B/s
[srv] ERROR: read from 127.0.0.1:50143: EOF
[srv] close: 127.0.0.1:50143
0009 -> see max: 15334000ns, min: 1000ns, avg: 82908ns
0009 -> 1024/1024 succ/totl, speed: 12.05 B/s
[srv] ERROR: read from 127.0.0.1:50147: EOF
[srv] close: 127.0.0.1:50147
[srv] ERROR: read from 127.0.0.1:50144: EOF
[srv] close: 127.0.0.1:50144
0016 -> see max: 18030000ns, min: 1000ns, avg: 89852ns
0016 -> 1024/1024 succ/totl, speed: 11.28 B/s
[srv] ERROR: read from 127.0.0.1:50146: EOF
[srv] close: 127.0.0.1:50146
0004 -> see max: 21791000ns, min: 1000ns, avg: 93691ns
0004 -> 1024/1024 succ/totl, speed: 10.81 B/s
0008 -> see max: 17852000ns, min: 1000ns, avg: 93841ns
0008 -> 1024/1024 succ/totl, speed: 10.81 B/s
[srv] ERROR: read from 127.0.0.1:50155: EOF
[srv] close: 127.0.0.1:50155
0013 -> see max: 21563000ns, min: 1000ns, avg: 95194ns
0013 -> 1024/1024 succ/totl, speed: 10.59 B/s
0007 -> see max: 18427000ns, min: 1000ns, avg: 95094ns
0007 -> 1024/1024 succ/totl, speed: 10.59 B/s
0012 -> see max: 18688000ns, min: 1000ns, avg: 95823ns
0012 -> 1024/1024 succ/totl, speed: 10.49 B/s
[srv] ERROR: read from 127.0.0.1:50154: EOF
[srv] close: 127.0.0.1:50154
[srv] ERROR: read from 127.0.0.1:50148: EOF
[srv] close: 127.0.0.1:50148
0011 -> see max: 17805000ns, min: 1000ns, avg: 96651ns
0011 -> 1024/1024 succ/totl, speed: 10.49 B/s
[srv] ERROR: read from 127.0.0.1:50142: EOF
[srv] close: 127.0.0.1:50142
[srv] ERROR: read from 127.0.0.1:50149: EOF
[srv] ERROR: read from 127.0.0.1:50152: EOF
[srv] close: 127.0.0.1:50149
[srv] close: 127.0.0.1:50152
[srv] ERROR: read from 127.0.0.1:50153: EOF
[srv] close: 127.0.0.1:50153
[srv] ERROR: read from 127.0.0.1:50150: EOF
[srv] close: 127.0.0.1:50150
[srv] ERROR: read from 127.0.0.1:50145: EOF
[srv] close: 127.0.0.1:50145
0015 -> see max: 29319000ns, min: 1000ns, avg: 93613ns
0015 -> 1024/1024 succ/totl, speed: 10.70 B/s
[srv] ERROR: read from 127.0.0.1:50151: EOF
[srv] close: 127.0.0.1:50151
0010 -> see max: 22669000ns, min: 1000ns, avg: 100963ns
0010 -> 1024/1024 succ/totl, speed: 9.99 B/s
[srv] ERROR: read from 127.0.0.1:50156: EOF
[srv] close: 127.0.0.1:50156
[srv] close: 127.0.0.1:50157
[srv] ERROR: read from 127.0.0.1:50157: EOF
^Csignal: interrupt
```
### 64 clitens
```bash
$ go run main.go -c 64 -s 
[srv] start
[srv] accept: 127.0.0.1:50166
[srv] accept: 127.0.0.1:50167
[srv] accept: 127.0.0.1:50165
[srv] accept: 127.0.0.1:50168
[srv] accept: 127.0.0.1:50169
[srv] accept: 127.0.0.1:50170
[srv] accept: 127.0.0.1:50171
[srv] accept: 127.0.0.1:50172
[srv] accept: 127.0.0.1:50173
[srv] accept: 127.0.0.1:50174
[srv] accept: 127.0.0.1:50175
[srv] accept: 127.0.0.1:50176
[srv] accept: 127.0.0.1:50177
[srv] accept: 127.0.0.1:50178
[srv] accept: 127.0.0.1:50179
[srv] accept: 127.0.0.1:50180
[srv] accept: 127.0.0.1:50181
[srv] accept: 127.0.0.1:50182
[srv] accept: 127.0.0.1:50183
[srv] accept: 127.0.0.1:50184
[srv] accept: 127.0.0.1:50185
[srv] accept: 127.0.0.1:50186
[srv] accept: 127.0.0.1:50187
[srv] accept: 127.0.0.1:50188
[srv] accept: 127.0.0.1:50189
[srv] accept: 127.0.0.1:50190
[srv] accept: 127.0.0.1:50191
[srv] accept: 127.0.0.1:50192
[srv] accept: 127.0.0.1:50193
[srv] accept: 127.0.0.1:50194
[srv] accept: 127.0.0.1:50195
[srv] accept: 127.0.0.1:50197
[srv] accept: 127.0.0.1:50196
[srv] accept: 127.0.0.1:50198
[srv] accept: 127.0.0.1:50199
[srv] accept: 127.0.0.1:50200
[srv] accept: 127.0.0.1:50201
[srv] accept: 127.0.0.1:50202
[srv] accept: 127.0.0.1:50203
[srv] accept: 127.0.0.1:50204
[srv] accept: 127.0.0.1:50205
[srv] accept: 127.0.0.1:50206
[srv] accept: 127.0.0.1:50207
[srv] accept: 127.0.0.1:50208
[srv] accept: 127.0.0.1:50209
[srv] accept: 127.0.0.1:50210
[srv] accept: 127.0.0.1:50211
[srv] accept: 127.0.0.1:50212
[srv] accept: 127.0.0.1:50213
[srv] accept: 127.0.0.1:50214
[srv] accept: 127.0.0.1:50215
[srv] accept: 127.0.0.1:50216
[srv] accept: 127.0.0.1:50217
[srv] accept: 127.0.0.1:50218
[srv] accept: 127.0.0.1:50219
[srv] accept: 127.0.0.1:50220
[srv] accept: 127.0.0.1:50221
[srv] accept: 127.0.0.1:50222
[srv] accept: 127.0.0.1:50223
[srv] accept: 127.0.0.1:50224
[srv] accept: 127.0.0.1:50225
[srv] accept: 127.0.0.1:50226
[srv] accept: 127.0.0.1:50227
[srv] accept: 127.0.0.1:50228
0002 -> beam max: 16378000ns, min: 1000ns, avg: 80856ns
0008 -> beam max: 16268000ns, min: 1000ns, avg: 82065ns
0012 -> beam max: 22308000ns, min: 1000ns, avg: 107791ns
0064 -> beam max: 40666000ns, min: 1000ns, avg: 109989ns
0014 -> beam max: 26111000ns, min: 1000ns, avg: 112076ns
0007 -> beam max: 58591000ns, min: 1000ns, avg: 142444ns
0015 -> beam max: 72047000ns, min: 1000ns, avg: 165064ns
0016 -> beam max: 71600000ns, min: 1000ns, avg: 167203ns
0004 -> beam max: 71265000ns, min: 1000ns, avg: 168146ns
0045 -> beam max: 75847000ns, min: 1000ns, avg: 168115ns
0005 -> beam max: 70432000ns, min: 1000ns, avg: 168866ns
0023 -> beam max: 70118000ns, min: 1000ns, avg: 182895ns
0046 -> beam max: 73613000ns, min: 1000ns, avg: 242709ns
0021 -> beam max: 43521000ns, min: 1000ns, avg: 320301ns
0001 -> beam max: 64932000ns, min: 1000ns, avg: 326474ns
0043 -> beam max: 74545000ns, min: 1000ns, avg: 337833ns
0022 -> beam max: 46798000ns, min: 1000ns, avg: 342206ns
0006 -> beam max: 74467000ns, min: 1000ns, avg: 342105ns
0011 -> beam max: 74489000ns, min: 1000ns, avg: 343555ns
0044 -> beam max: 68923000ns, min: 1000ns, avg: 344462ns
0017 -> beam max: 75909000ns, min: 1000ns, avg: 344474ns
0042 -> beam max: 74746000ns, min: 1000ns, avg: 346559ns
0003 -> beam max: 75192000ns, min: 1000ns, avg: 347342ns
0010 -> beam max: 75446000ns, min: 1000ns, avg: 351749ns
0019 -> beam max: 76612000ns, min: 1000ns, avg: 351688ns
0009 -> beam max: 76091000ns, min: 1000ns, avg: 352189ns
0054 -> beam max: 74927000ns, min: 1000ns, avg: 354279ns
0025 -> beam max: 72817000ns, min: 1000ns, avg: 355167ns
0050 -> beam max: 135689000ns, min: 1000ns, avg: 356144ns
0027 -> beam max: 141861000ns, min: 1000ns, avg: 356222ns
0020 -> beam max: 79577000ns, min: 1000ns, avg: 357191ns
0055 -> beam max: 77292000ns, min: 1000ns, avg: 356731ns
0064 -> see max: 42702000ns, min: 1000ns, avg: 365882ns
0064 -> 1024/1024 succ/totl, speed: 2.79 B/s
0056 -> beam max: 72652000ns, min: 1000ns, avg: 358352ns
0008 -> see max: 80133000ns, min: 1000ns, avg: 369229ns
0008 -> 1024/1024 succ/totl, speed: 2.77 B/s
0033 -> beam max: 81529000ns, min: 1000ns, avg: 360702ns
0041 -> beam max: 81880000ns, min: 1000ns, avg: 361972ns
0052 -> beam max: 70460000ns, min: 1000ns, avg: 364418ns
0018 -> beam max: 84916000ns, min: 2000ns, avg: 370362ns
0013 -> beam max: 84237000ns, min: 1000ns, avg: 376691ns
0002 -> see max: 94662000ns, min: 1000ns, avg: 399003ns
0002 -> 1024/1024 succ/totl, speed: 2.56 B/s
0007 -> see max: 76366000ns, min: 1000ns, avg: 414455ns
0007 -> 1024/1024 succ/totl, speed: 2.46 B/s
[srv] ERROR: read from 127.0.0.1:50167: EOF
[srv] close: 127.0.0.1:50167
0047 -> beam max: 97832000ns, min: 1000ns, avg: 420534ns
[srv] ERROR: read from 127.0.0.1:50170: EOF
[srv] close: 127.0.0.1:50170
0005 -> see max: 105053000ns, min: 1000ns, avg: 423236ns
0005 -> 1024/1024 succ/totl, speed: 2.41 B/s
0016 -> see max: 86522000ns, min: 1000ns, avg: 428571ns
0016 -> 1024/1024 succ/totl, speed: 2.38 B/s
0048 -> beam max: 96302000ns, min: 1000ns, avg: 430362ns
0024 -> beam max: 97316000ns, min: 1000ns, avg: 431284ns
0049 -> beam max: 101123000ns, min: 1000ns, avg: 441432ns
[srv] ERROR: read from 127.0.0.1:50168: EOF
[srv] close: 127.0.0.1:50168
0026 -> beam max: 91333000ns, min: 1000ns, avg: 441851ns
0053 -> beam max: 74052000ns, min: 1000ns, avg: 450980ns
[srv] ERROR: read from 127.0.0.1:50166: EOF
[srv] close: 127.0.0.1:50166
0030 -> beam max: 79543000ns, min: 1000ns, avg: 452243ns
[srv] ERROR: read from 127.0.0.1:50172: EOF
[srv] close: 127.0.0.1:50172
0057 -> beam max: 83184000ns, min: 1000ns, avg: 457628ns
0034 -> beam max: 87392000ns, min: 2000ns, avg: 460222ns
0028 -> beam max: 77835000ns, min: 1000ns, avg: 462169ns
[srv] ERROR: read from 127.0.0.1:50173: EOF
[srv] close: 127.0.0.1:50173
0058 -> beam max: 75872000ns, min: 1000ns, avg: 462088ns
0029 -> beam max: 77392000ns, min: 1000ns, avg: 465839ns
0035 -> beam max: 88999000ns, min: 2000ns, avg: 465485ns
0014 -> see max: 81579000ns, min: 1000ns, avg: 476970ns
0014 -> 1024/1024 succ/totl, speed: 2.14 B/s
0036 -> beam max: 78992000ns, min: 1000ns, avg: 465160ns
0059 -> beam max: 88939000ns, min: 2000ns, avg: 468291ns
0038 -> beam max: 77213000ns, min: 1000ns, avg: 466881ns
0061 -> beam max: 82135000ns, min: 1000ns, avg: 469512ns
0060 -> beam max: 86306000ns, min: 1000ns, avg: 471397ns
0037 -> beam max: 83577000ns, min: 1000ns, avg: 469666ns
0062 -> beam max: 86691000ns, min: 1000ns, avg: 471681ns
0039 -> beam max: 89650000ns, min: 2000ns, avg: 473708ns
0040 -> beam max: 88321000ns, min: 1000ns, avg: 473772ns
0063 -> beam max: 92308000ns, min: 1000ns, avg: 476582ns
0012 -> see max: 94300000ns, min: 1000ns, avg: 496808ns
0012 -> 1024/1024 succ/totl, speed: 2.06 B/s
0004 -> see max: 86330000ns, min: 1000ns, avg: 498542ns
0004 -> 1024/1024 succ/totl, speed: 2.05 B/s
0015 -> see max: 98436000ns, min: 1000ns, avg: 499812ns
0015 -> 1024/1024 succ/totl, speed: 2.04 B/s
0021 -> see max: 64068000ns, min: 1000ns, avg: 512660ns
0021 -> 1024/1024 succ/totl, speed: 1.99 B/s
[srv] ERROR: read from 127.0.0.1:50165: EOF
[srv] close: 127.0.0.1:50165
0001 -> see max: 91954000ns, min: 1000ns, avg: 515122ns
0001 -> 1024/1024 succ/totl, speed: 1.98 B/s
0043 -> see max: 98552000ns, min: 1000ns, avg: 532778ns
0043 -> 1024/1024 succ/totl, speed: 1.91 B/s
0006 -> see max: 100103000ns, min: 1000ns, avg: 534156ns
0006 -> 1024/1024 succ/totl, speed: 1.91 B/s
0022 -> see max: 54368000ns, min: 1000ns, avg: 536425ns
0022 -> 1024/1024 succ/totl, speed: 1.90 B/s
[srv] ERROR: read from 127.0.0.1:50174: EOF
[srv] close: 127.0.0.1:50174
[srv] ERROR: read from 127.0.0.1:50171: EOF
[srv] close: 127.0.0.1:50171
0011 -> see max: 99615000ns, min: 1000ns, avg: 538960ns
0011 -> 1024/1024 succ/totl, speed: 1.90 B/s
0017 -> see max: 101401000ns, min: 1000ns, avg: 536408ns
0017 -> 1024/1024 succ/totl, speed: 1.90 B/s
0044 -> see max: 91313000ns, min: 1000ns, avg: 539990ns
0044 -> 1024/1024 succ/totl, speed: 1.89 B/s
[srv] ERROR: read from 127.0.0.1:50169: EOF
[srv] close: 127.0.0.1:50169
0042 -> see max: 100047000ns, min: 1000ns, avg: 539508ns
0042 -> 1024/1024 succ/totl, speed: 1.89 B/s
0009 -> see max: 95532000ns, min: 1000ns, avg: 539248ns
0009 -> 1024/1024 succ/totl, speed: 1.89 B/s
0019 -> see max: 99586000ns, min: 1000ns, avg: 542639ns
0019 -> 1024/1024 succ/totl, speed: 1.88 B/s
0010 -> see max: 98757000ns, min: 1000ns, avg: 542889ns
0010 -> 1024/1024 succ/totl, speed: 1.88 B/s
0054 -> see max: 93021000ns, min: 1000ns, avg: 537522ns
0054 -> 1024/1024 succ/totl, speed: 1.90 B/s
0031 -> beam max: 152923000ns, min: 1000ns, avg: 538880ns
0020 -> see max: 94441000ns, min: 1000ns, avg: 545337ns
0020 -> 1024/1024 succ/totl, speed: 1.87 B/s
0032 -> beam max: 154726000ns, min: 1000ns, avg: 541203ns
0041 -> see max: 90719000ns, min: 1000ns, avg: 547235ns
0041 -> 1024/1024 succ/totl, speed: 1.87 B/s
0052 -> see max: 83245000ns, min: 1000ns, avg: 545666ns
0052 -> 1024/1024 succ/totl, speed: 1.87 B/s
0055 -> see max: 88592000ns, min: 1000ns, avg: 544020ns
0055 -> 1024/1024 succ/totl, speed: 1.88 B/s
0033 -> see max: 91663000ns, min: 1000ns, avg: 543767ns
0033 -> 1024/1024 succ/totl, speed: 1.88 B/s
0028 -> see max: 77832000ns, min: 1000ns, avg: 548204ns
0028 -> 1024/1024 succ/totl, speed: 1.86 B/s
[srv] ERROR: read from 127.0.0.1:50190: EOF
[srv] close: 127.0.0.1:50190
0030 -> see max: 78665000ns, min: 1000ns, avg: 551313ns
0030 -> 1024/1024 succ/totl, speed: 1.85 B/s
[srv] ERROR: read from 127.0.0.1:50176: EOF
[srv] close: 127.0.0.1:50176
0018 -> see max: 94835000ns, min: 1000ns, avg: 554762ns
0018 -> 1024/1024 succ/totl, speed: 1.84 B/s
0056 -> see max: 80656000ns, min: 1000ns, avg: 550851ns
0056 -> 1024/1024 succ/totl, speed: 1.86 B/s
0053 -> see max: 75679000ns, min: 1000ns, avg: 553580ns
0053 -> 1024/1024 succ/totl, speed: 1.85 B/s
0013 -> see max: 89378000ns, min: 1000ns, avg: 562540ns
0013 -> 1024/1024 succ/totl, speed: 1.82 B/s
0036 -> see max: 80073000ns, min: 1000ns, avg: 560153ns
0036 -> 1024/1024 succ/totl, speed: 1.82 B/s
0037 -> see max: 83482000ns, min: 1000ns, avg: 559881ns
0037 -> 1024/1024 succ/totl, speed: 1.82 B/s
[srv] ERROR: read from 127.0.0.1:50186: EOF
[srv] close: 127.0.0.1:50186
0045 -> see max: 95009000ns, min: 1000ns, avg: 571977ns
0045 -> 1024/1024 succ/totl, speed: 1.78 B/s
[srv] ERROR: read from 127.0.0.1:50178: EOF
[srv] close: 127.0.0.1:50178
[srv] ERROR: read from 127.0.0.1:50191: EOF
[srv] close: 127.0.0.1:50191
0038 -> see max: 82604000ns, min: 1000ns, avg: 562222ns
0038 -> 1024/1024 succ/totl, speed: 1.82 B/s
[srv] ERROR: read from 127.0.0.1:50185: EOF
[srv] close: 127.0.0.1:50185
[srv] ERROR: read from 127.0.0.1:50175: EOF
[srv] close: 127.0.0.1:50175
[srv] ERROR: read from 127.0.0.1:50189: EOF
[srv] close: 127.0.0.1:50189
0023 -> see max: 99683000ns, min: 1000ns, avg: 573955ns
0023 -> 1024/1024 succ/totl, speed: 1.78 B/s
0046 -> see max: 101100000ns, min: 1000ns, avg: 574392ns
0046 -> 1024/1024 succ/totl, speed: 1.78 B/s
[srv] ERROR: read from 127.0.0.1:50183: EOF
[srv] close: 127.0.0.1:50183
[srv] ERROR: read from 127.0.0.1:50182: EOF
[srv] close: 127.0.0.1:50182
[srv] ERROR: read from 127.0.0.1:50184: EOF
[srv] close: 127.0.0.1:50184
[srv] ERROR: read from 127.0.0.1:50188: EOF
[srv] close: 127.0.0.1:50188
[srv] ERROR: read from 127.0.0.1:50209: EOF
[srv] close: 127.0.0.1:50209
0047 -> see max: 97696000ns, min: 1000ns, avg: 575170ns
0047 -> 1024/1024 succ/totl, speed: 1.78 B/s
[srv] ERROR: read from 127.0.0.1:50187: EOF
[srv] close: 127.0.0.1:50187
[srv] ERROR: read from 127.0.0.1:50181: EOF
[srv] close: 127.0.0.1:50181
[srv] ERROR: read from 127.0.0.1:50205: EOF
[srv] close: 127.0.0.1:50205
[srv] ERROR: read from 127.0.0.1:50211: EOF
[srv] close: 127.0.0.1:50211
[srv] ERROR: read from 127.0.0.1:50214: EOF
[srv] close: 127.0.0.1:50214
[srv] ERROR: read from 127.0.0.1:50204: EOF
[srv] close: 127.0.0.1:50204
[srv] ERROR: read from 127.0.0.1:50208: EOF
[srv] close: 127.0.0.1:50208
[srv] ERROR: read from 127.0.0.1:50180: EOF
[srv] close: 127.0.0.1:50180
[srv] ERROR: read from 127.0.0.1:50213: EOF
[srv] close: 127.0.0.1:50213
0048 -> see max: 96251000ns, min: 1000ns, avg: 581601ns
0048 -> 1024/1024 succ/totl, speed: 1.75 B/s
[srv] ERROR: read from 127.0.0.1:50207: EOF
[srv] close: 127.0.0.1:50207
0024 -> see max: 97372000ns, min: 1000ns, avg: 583512ns
0024 -> 1024/1024 succ/totl, speed: 1.75 B/s
[srv] ERROR: read from 127.0.0.1:50179: EOF
[srv] close: 127.0.0.1:50179
0049 -> see max: 100371000ns, min: 1000ns, avg: 587445ns
0049 -> 1024/1024 succ/totl, speed: 1.74 B/s
0025 -> see max: 83868000ns, min: 1000ns, avg: 586073ns
0025 -> 1024/1024 succ/totl, speed: 1.74 B/s
[srv] ERROR: read from 127.0.0.1:50221: EOF
[srv] ERROR: read from 127.0.0.1:50219: EOF
[srv] close: 127.0.0.1:50221
[srv] close: 127.0.0.1:50219
0026 -> see max: 91931000ns, min: 1000ns, avg: 585305ns
0026 -> 1024/1024 succ/totl, speed: 1.74 B/s
[srv] ERROR: read from 127.0.0.1:50192: EOF
[srv] close: 127.0.0.1:50192
[srv] ERROR: read from 127.0.0.1:50223: EOF
[srv] close: 127.0.0.1:50223
0051 -> beam max: 338877000ns, min: 1000ns, avg: 586401ns
[srv] ERROR: read from 127.0.0.1:50193: EOF
[srv] close: 127.0.0.1:50193
[srv] ERROR: read from 127.0.0.1:50194: EOF
[srv] close: 127.0.0.1:50194
0003 -> see max: 100424000ns, min: 1000ns, avg: 590703ns
0003 -> 1024/1024 succ/totl, speed: 1.73 B/s
[srv] ERROR: read from 127.0.0.1:50195: EOF
[srv] close: 127.0.0.1:50195
0035 -> see max: 87762000ns, min: 1000ns, avg: 585204ns
0035 -> 1024/1024 succ/totl, speed: 1.74 B/s
0059 -> see max: 88780000ns, min: 1000ns, avg: 584768ns
0059 -> 1024/1024 succ/totl, speed: 1.75 B/s
0060 -> see max: 86261000ns, min: 1000ns, avg: 584616ns
0060 -> 1024/1024 succ/totl, speed: 1.75 B/s
[srv] ERROR: read from 127.0.0.1:50197: EOF
[srv] close: 127.0.0.1:50197
[srv] ERROR: read from 127.0.0.1:50196: EOF
[srv] close: 127.0.0.1:50196
0039 -> see max: 89778000ns, min: 1000ns, avg: 584099ns
0039 -> 1024/1024 succ/totl, speed: 1.75 B/s
0062 -> see max: 86787000ns, min: 1000ns, avg: 584339ns
0062 -> 1024/1024 succ/totl, speed: 1.75 B/s
0040 -> see max: 88369000ns, min: 1000ns, avg: 584062ns
0040 -> 1024/1024 succ/totl, speed: 1.75 B/s
0063 -> see max: 91737000ns, min: 1000ns, avg: 584326ns
0063 -> 1024/1024 succ/totl, speed: 1.75 B/s
[srv] ERROR: read from 127.0.0.1:50199: EOF
[srv] close: 127.0.0.1:50198
[srv] ERROR: read from 127.0.0.1:50198: EOF
[srv] close: 127.0.0.1:50199
[srv] ERROR: read from 127.0.0.1:50201: EOF
[srv] close: 127.0.0.1:50201
[srv] ERROR: read from 127.0.0.1:50177: EOF
[srv] close: 127.0.0.1:50177
[srv] ERROR: read from 127.0.0.1:50218: EOF
[srv] ERROR: read from 127.0.0.1:50220: EOF
[srv] close: 127.0.0.1:50218
[srv] close: 127.0.0.1:50220
[srv] ERROR: read from 127.0.0.1:50222: EOF
[srv] close: 127.0.0.1:50222
[srv] ERROR: read from 127.0.0.1:50226: EOF
[srv] close: 127.0.0.1:50226
[srv] ERROR: read from 127.0.0.1:50225: EOF
[srv] close: 127.0.0.1:50225
[srv] ERROR: read from 127.0.0.1:50227: EOF
[srv] close: 127.0.0.1:50227
[srv] ERROR: read from 127.0.0.1:50228: EOF
[srv] close: 127.0.0.1:50228
0050 -> see max: 85440000ns, min: 1000ns, avg: 595585ns
0050 -> 1024/1024 succ/totl, speed: 1.72 B/s
0027 -> see max: 88485000ns, min: 1000ns, avg: 596338ns
0027 -> 1024/1024 succ/totl, speed: 1.71 B/s
0031 -> see max: 95794000ns, min: 1000ns, avg: 595538ns
0031 -> 1024/1024 succ/totl, speed: 1.71 B/s
0032 -> see max: 94227000ns, min: 1000ns, avg: 595589ns
0032 -> 1024/1024 succ/totl, speed: 1.71 B/s
[srv] ERROR: read from 127.0.0.1:50200: EOF
[srv] close: 127.0.0.1:50202
[srv] close: 127.0.0.1:50200
[srv] ERROR: read from 127.0.0.1:50202: EOF
[srv] ERROR: read from 127.0.0.1:50210: EOF
[srv] close: 127.0.0.1:50210
[srv] ERROR: read from 127.0.0.1:50212: EOF
[srv] close: 127.0.0.1:50212
0051 -> see max: 105560000ns, min: 1000ns, avg: 601465ns
0051 -> 1024/1024 succ/totl, speed: 1.70 B/s
0058 -> see max: 76600000ns, min: 1000ns, avg: 596773ns
0058 -> 1024/1024 succ/totl, speed: 1.71 B/s
[srv] ERROR: read from 127.0.0.1:50217: EOF
[srv] ERROR: read from 127.0.0.1:50203: EOF
[srv] close: 127.0.0.1:50217
[srv] close: 127.0.0.1:50203
0057 -> see max: 83065000ns, min: 1000ns, avg: 597371ns
0057 -> 1024/1024 succ/totl, speed: 1.71 B/s
[srv] ERROR: read from 127.0.0.1:50215: EOF
[srv] close: 127.0.0.1:50215
0061 -> see max: 82062000ns, min: 1000ns, avg: 597210ns
0061 -> 1024/1024 succ/totl, speed: 1.71 B/s
[srv] ERROR: read from 127.0.0.1:50224: EOF
[srv] close: 127.0.0.1:50224
0034 -> see max: 100275000ns, min: 1000ns, avg: 692717ns
0034 -> 1024/1024 succ/totl, speed: 1.47 B/s
[srv] ERROR: read from 127.0.0.1:50216: EOF
[srv] close: 127.0.0.1:50216
0029 -> see max: 101247000ns, min: 1000ns, avg: 698663ns
0029 -> 1024/1024 succ/totl, speed: 1.46 B/s
[srv] ERROR: read from 127.0.0.1:50206: EOF
[srv] close: 127.0.0.1:50206
^Csignal: interrupt
```
