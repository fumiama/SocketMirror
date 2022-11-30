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
0001 -> beam max: 13000ns, min: 3000ns, avg: 6239ns
0001 -> see max: 197000ns, min: 1000ns, avg: 6092ns
0001 -> 1024/1024 succ/totl, speed: 131.07 B/s
^Csignal: interrup
```
### 10 clients
```bash
$ go run main.go -c 10 -s 
0003 -> beam max: 5152000ns, min: 1000ns, avg: 22658ns
0010 -> beam max: 4998000ns, min: 1000ns, avg: 23085ns
0008 -> beam max: 4915000ns, min: 1000ns, avg: 23097ns
0005 -> beam max: 4737000ns, min: 1000ns, avg: 23023ns
0009 -> beam max: 4874000ns, min: 1000ns, avg: 23150ns
0004 -> beam max: 7197000ns, min: 1000ns, avg: 29473ns
0001 -> beam max: 6575000ns, min: 1000ns, avg: 30664ns
0006 -> beam max: 10550000ns, min: 1000ns, avg: 36394ns
0007 -> beam max: 11110000ns, min: 1000ns, avg: 37282ns
0002 -> beam max: 9169000ns, min: 1000ns, avg: 40517ns
0009 -> see max: 9744000ns, min: 1000ns, avg: 45606ns
0009 -> 80/1024 succ/totl, speed: 21.40 B/s
0001 -> see max: 9187000ns, min: 1000ns, avg: 45706ns
0001 -> 66/1024 succ/totl, speed: 21.85 B/s
0003 -> see max: 9864000ns, min: 1000ns, avg: 46166ns
0003 -> 76/1024 succ/totl, speed: 21.40 B/s
0008 -> see max: 11670000ns, min: 1000ns, avg: 46694ns
0008 -> 68/1024 succ/totl, speed: 21.40 B/s
0004 -> see max: 9053000ns, min: 1000ns, avg: 47000ns
0004 -> 62/1024 succ/totl, speed: 20.97 B/s
0005 -> see max: 8821000ns, min: 1000ns, avg: 47694ns
0005 -> 92/1024 succ/totl, speed: 20.97 B/s
0010 -> see max: 8946000ns, min: 1000ns, avg: 47808ns
0010 -> 75/1024 succ/totl, speed: 20.56 B/s
0006 -> see max: 11561000ns, min: 1000ns, avg: 52102ns
0006 -> 200/1024 succ/totl, speed: 19.42 B/s
0007 -> see max: 12031000ns, min: 1000ns, avg: 52417ns
0007 -> 102/1024 succ/totl, speed: 19.07 B/s
0002 -> see max: 10132000ns, min: 1000ns, avg: 50840ns
0002 -> 116/1024 succ/totl, speed: 19.78 B/s
^Csignal: interrupt
```
### 64 clitens
```bash
$ go run main.go -c 64 -s 
0021 -> beam max: 9457000ns, min: 1000ns, avg: 63116ns
0012 -> beam max: 10993000ns, min: 1000ns, avg: 63576ns
0023 -> beam max: 28234000ns, min: 1000ns, avg: 96542ns
0009 -> beam max: 37837000ns, min: 1000ns, avg: 103391ns
0014 -> beam max: 33803000ns, min: 1000ns, avg: 104684ns
0011 -> beam max: 44920000ns, min: 1000ns, avg: 113033ns
0025 -> beam max: 63361000ns, min: 1000ns, avg: 156197ns
0002 -> beam max: 63223000ns, min: 1000ns, avg: 156017ns
0018 -> beam max: 36897000ns, min: 1000ns, avg: 156575ns
0003 -> beam max: 54265000ns, min: 1000ns, avg: 161769ns
0022 -> beam max: 64612000ns, min: 1000ns, avg: 162008ns
0010 -> beam max: 49291000ns, min: 1000ns, avg: 162610ns
0026 -> beam max: 57262000ns, min: 1000ns, avg: 168803ns
0024 -> beam max: 64849000ns, min: 1000ns, avg: 170715ns
0019 -> beam max: 48710000ns, min: 1000ns, avg: 171054ns
0027 -> beam max: 67404000ns, min: 1000ns, avg: 171169ns
0015 -> beam max: 68681000ns, min: 1000ns, avg: 171633ns
0004 -> beam max: 69670000ns, min: 1000ns, avg: 173960ns
0017 -> beam max: 68090000ns, min: 1000ns, avg: 174417ns
0013 -> beam max: 68647000ns, min: 1000ns, avg: 176049ns
0008 -> beam max: 69439000ns, min: 1000ns, avg: 176188ns
0028 -> beam max: 70260000ns, min: 1000ns, avg: 296237ns
0007 -> beam max: 77468000ns, min: 1000ns, avg: 297517ns
0016 -> beam max: 73251000ns, min: 1000ns, avg: 298701ns
0020 -> beam max: 75527000ns, min: 1000ns, avg: 300106ns
0001 -> beam max: 70897000ns, min: 1000ns, avg: 301930ns
0023 -> see max: 39259000ns, min: 1000ns, avg: 307565ns
0023 -> 84/1024 succ/totl, speed: 3.32 B/s
0030 -> beam max: 63227000ns, min: 1000ns, avg: 308924ns
0064 -> beam max: 73082000ns, min: 1000ns, avg: 308469ns
0031 -> beam max: 67648000ns, min: 1000ns, avg: 309228ns
0033 -> beam max: 74379000ns, min: 1000ns, avg: 309123ns
0021 -> see max: 67025000ns, min: 1000ns, avg: 316029ns
0021 -> 121/1024 succ/totl, speed: 3.23 B/s
0034 -> beam max: 74696000ns, min: 1000ns, avg: 309109ns
0029 -> beam max: 58115000ns, min: 1000ns, avg: 314076ns
0035 -> beam max: 77782000ns, min: 1000ns, avg: 314268ns
0012 -> see max: 74557000ns, min: 1000ns, avg: 324016ns
0012 -> 119/1024 succ/totl, speed: 3.15 B/s
0011 -> see max: 62284000ns, min: 1000ns, avg: 324436ns
0011 -> 76/1024 succ/totl, speed: 3.15 B/s
0036 -> beam max: 75890000ns, min: 1000ns, avg: 315592ns
0039 -> beam max: 75491000ns, min: 1000ns, avg: 312379ns
0006 -> beam max: 75620000ns, min: 1000ns, avg: 313003ns
0040 -> beam max: 75175000ns, min: 1000ns, avg: 314201ns
0005 -> beam max: 75351000ns, min: 1000ns, avg: 315011ns
0041 -> beam max: 75745000ns, min: 1000ns, avg: 315926ns
0042 -> beam max: 75201000ns, min: 1000ns, avg: 315827ns
0043 -> beam max: 76107000ns, min: 1000ns, avg: 318055ns
0044 -> beam max: 74860000ns, min: 1000ns, avg: 317673ns
0045 -> beam max: 62836000ns, min: 1000ns, avg: 312187ns
0047 -> beam max: 73490000ns, min: 1000ns, avg: 324882ns
0057 -> beam max: 73389000ns, min: 1000ns, avg: 330270ns
0009 -> see max: 74786000ns, min: 1000ns, avg: 367550ns
0009 -> 107/1024 succ/totl, speed: 2.78 B/s
0014 -> see max: 50996000ns, min: 1000ns, avg: 368262ns
0014 -> 83/1024 succ/totl, speed: 2.77 B/s
0032 -> beam max: 65121000ns, min: 1000ns, avg: 370943ns
0037 -> beam max: 118644000ns, min: 1000ns, avg: 379599ns
0048 -> beam max: 74331000ns, min: 1000ns, avg: 375532ns
0018 -> see max: 39749000ns, min: 1000ns, avg: 395242ns
0018 -> 81/1024 succ/totl, speed: 2.58 B/s
0049 -> beam max: 77393000ns, min: 1000ns, avg: 377926ns
0051 -> beam max: 71013000ns, min: 1000ns, avg: 376396ns
0050 -> beam max: 77240000ns, min: 1000ns, avg: 380443ns
0052 -> beam max: 72225000ns, min: 1000ns, avg: 377566ns
0053 -> beam max: 68124000ns, min: 1000ns, avg: 374996ns
0055 -> beam max: 74817000ns, min: 1000ns, avg: 379433ns
0054 -> beam max: 71740000ns, min: 1000ns, avg: 377781ns
0022 -> see max: 64594000ns, min: 1000ns, avg: 401889ns
0022 -> 91/1024 succ/totl, speed: 2.54 B/s
0056 -> beam max: 74784000ns, min: 1000ns, avg: 380312ns
0002 -> see max: 63682000ns, min: 1000ns, avg: 403785ns
0002 -> 92/1024 succ/totl, speed: 2.53 B/s
0010 -> see max: 49017000ns, min: 1000ns, avg: 403897ns
0010 -> 37/1024 succ/totl, speed: 2.53 B/s
0058 -> beam max: 67137000ns, min: 1000ns, avg: 379761ns
0059 -> beam max: 68237000ns, min: 1000ns, avg: 381216ns
0060 -> beam max: 71396000ns, min: 1000ns, avg: 383563ns
0061 -> beam max: 69749000ns, min: 1000ns, avg: 382926ns
0062 -> beam max: 76708000ns, min: 1000ns, avg: 386464ns
0063 -> beam max: 77539000ns, min: 1000ns, avg: 387065ns
0024 -> see max: 63959000ns, min: 1000ns, avg: 413764ns
0024 -> 62/1024 succ/totl, speed: 2.47 B/s
0019 -> see max: 48581000ns, min: 1000ns, avg: 412958ns
0019 -> 71/1024 succ/totl, speed: 2.47 B/s
0008 -> see max: 66830000ns, min: 1000ns, avg: 415596ns
0008 -> 40/1024 succ/totl, speed: 2.46 B/s
0013 -> see max: 68907000ns, min: 1000ns, avg: 417007ns
0013 -> 46/1024 succ/totl, speed: 2.45 B/s
0015 -> see max: 73183000ns, min: 1000ns, avg: 415944ns
0015 -> 69/1024 succ/totl, speed: 2.45 B/s
0017 -> see max: 70619000ns, min: 1000ns, avg: 416708ns
0017 -> 58/1024 succ/totl, speed: 2.45 B/s
0004 -> see max: 69129000ns, min: 1000ns, avg: 417461ns
0004 -> 58/1024 succ/totl, speed: 2.45 B/s
0027 -> see max: 67282000ns, min: 1000ns, avg: 416087ns
0027 -> 73/1024 succ/totl, speed: 2.45 B/s
0016 -> see max: 72483000ns, min: 1000ns, avg: 427497ns
0016 -> 36/1024 succ/totl, speed: 2.39 B/s
0007 -> see max: 76622000ns, min: 1000ns, avg: 425125ns
0007 -> 52/1024 succ/totl, speed: 2.40 B/s
0020 -> see max: 75247000ns, min: 1000ns, avg: 423404ns
0020 -> 57/1024 succ/totl, speed: 2.41 B/s
0001 -> see max: 68327000ns, min: 1000ns, avg: 428178ns
0001 -> 197/1024 succ/totl, speed: 2.38 B/s
0064 -> see max: 71178000ns, min: 1000ns, avg: 432330ns
0064 -> 62/1024 succ/totl, speed: 2.36 B/s
0035 -> see max: 77618000ns, min: 1000ns, avg: 434142ns
0035 -> 283/1024 succ/totl, speed: 2.35 B/s
0003 -> see max: 54556000ns, min: 1000ns, avg: 442657ns
0003 -> 83/1024 succ/totl, speed: 2.31 B/s
0025 -> see max: 63511000ns, min: 1000ns, avg: 444310ns
0025 -> 430/1024 succ/totl, speed: 2.30 B/s
0039 -> see max: 75461000ns, min: 1000ns, avg: 431179ns
0039 -> 63/1024 succ/totl, speed: 2.37 B/s
0040 -> see max: 75180000ns, min: 1000ns, avg: 431489ns
0040 -> 66/1024 succ/totl, speed: 2.36 B/s
0005 -> see max: 75345000ns, min: 1000ns, avg: 432384ns
0005 -> 81/1024 succ/totl, speed: 2.36 B/s
0006 -> see max: 75589000ns, min: 1000ns, avg: 433302ns
0006 -> 120/1024 succ/totl, speed: 2.36 B/s
0042 -> see max: 75164000ns, min: 1000ns, avg: 432123ns
0042 -> 480/1024 succ/totl, speed: 2.36 B/s
0045 -> see max: 63947000ns, min: 1000ns, avg: 430136ns
0045 -> 491/1024 succ/totl, speed: 2.37 B/s
0047 -> see max: 73502000ns, min: 1000ns, avg: 433135ns
0047 -> 618/1024 succ/totl, speed: 2.36 B/s
0026 -> see max: 57212000ns, min: 1000ns, avg: 450224ns
0026 -> 150/1024 succ/totl, speed: 2.26 B/s
0048 -> see max: 74306000ns, min: 1000ns, avg: 432064ns
0048 -> 646/1024 succ/totl, speed: 2.36 B/s
0051 -> see max: 71026000ns, min: 1000ns, avg: 433220ns
0051 -> 716/1024 succ/totl, speed: 2.36 B/s
0050 -> see max: 77263000ns, min: 1000ns, avg: 437815ns
0050 -> 752/1024 succ/totl, speed: 2.33 B/s
0055 -> see max: 74860000ns, min: 1000ns, avg: 438755ns
0055 -> 633/1024 succ/totl, speed: 2.33 B/s
0058 -> see max: 67074000ns, min: 1000ns, avg: 434850ns
0058 -> 581/1024 succ/totl, speed: 2.35 B/s
0052 -> see max: 72214000ns, min: 1000ns, avg: 441124ns
0052 -> 553/1024 succ/totl, speed: 2.31 B/s
0057 -> see max: 73359000ns, min: 1000ns, avg: 438248ns
0057 -> 700/1024 succ/totl, speed: 2.33 B/s
0061 -> see max: 69746000ns, min: 1000ns, avg: 437852ns
0061 -> 585/1024 succ/totl, speed: 2.33 B/s
0062 -> see max: 76645000ns, min: 1000ns, avg: 439306ns
0062 -> 697/1024 succ/totl, speed: 2.32 B/s
0060 -> see max: 71000000ns, min: 1000ns, avg: 440053ns
0060 -> 681/1024 succ/totl, speed: 2.32 B/s
0028 -> see max: 69221000ns, min: 1000ns, avg: 464143ns
0028 -> 43/1024 succ/totl, speed: 2.20 B/s
0029 -> see max: 53987000ns, min: 1000ns, avg: 463815ns
0029 -> 86/1024 succ/totl, speed: 2.20 B/s
0030 -> see max: 62493000ns, min: 1000ns, avg: 465112ns
0030 -> 57/1024 succ/totl, speed: 2.20 B/s
0031 -> see max: 66181000ns, min: 1000ns, avg: 464571ns
0031 -> 61/1024 succ/totl, speed: 2.20 B/s
0032 -> see max: 64314000ns, min: 1000ns, avg: 465845ns
0032 -> 75/1024 succ/totl, speed: 2.19 B/s
0033 -> see max: 72862000ns, min: 1000ns, avg: 466067ns
0033 -> 89/1024 succ/totl, speed: 2.19 B/s
0034 -> see max: 74366000ns, min: 1000ns, avg: 466135ns
0034 -> 91/1024 succ/totl, speed: 2.19 B/s
0036 -> see max: 75757000ns, min: 1000ns, avg: 464107ns
0036 -> 291/1024 succ/totl, speed: 2.20 B/s
0041 -> see max: 75723000ns, min: 1000ns, avg: 457249ns
0041 -> 364/1024 succ/totl, speed: 2.23 B/s
0037 -> see max: 74932000ns, min: 1000ns, avg: 464283ns
0037 -> 285/1024 succ/totl, speed: 2.20 B/s
0056 -> see max: 75078000ns, min: 1000ns, avg: 452244ns
0056 -> 664/1024 succ/totl, speed: 2.26 B/s
0038 -> beam max: 156394000ns, min: 1000ns, avg: 464697ns
0046 -> beam max: 71201000ns, min: 1000ns, avg: 458367ns
0038 -> see max: 75579000ns, min: 1000ns, avg: 466067ns
0038 -> 376/1024 succ/totl, speed: 2.19 B/s
0059 -> see max: 68433000ns, min: 1000ns, avg: 453404ns
0059 -> 589/1024 succ/totl, speed: 2.25 B/s
0046 -> see max: 70682000ns, min: 1000ns, avg: 463064ns
0046 -> 136/1024 succ/totl, speed: 2.20 B/s
0043 -> see max: 76222000ns, min: 1000ns, avg: 465020ns
0043 -> 570/1024 succ/totl, speed: 2.19 B/s
0053 -> see max: 70328000ns, min: 1000ns, avg: 459364ns
0053 -> 620/1024 succ/totl, speed: 2.22 B/s
0049 -> see max: 77327000ns, min: 1000ns, avg: 463475ns
0049 -> 667/1024 succ/totl, speed: 2.20 B/s
0054 -> see max: 100185000ns, min: 1000ns, avg: 555161ns
0054 -> 559/1024 succ/totl, speed: 1.84 B/s
0063 -> see max: 100213000ns, min: 1000ns, avg: 551921ns
0063 -> 692/1024 succ/totl, speed: 1.85 B/s
0044 -> see max: 99698000ns, min: 1000ns, avg: 560025ns
0044 -> 527/1024 succ/totl, speed: 1.82 B/s
^Csignal: interrupt
```
