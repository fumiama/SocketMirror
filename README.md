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
0001 -> beam max: 86000ns, min: 1000ns, avg: 5469ns
0001 -> see max: 90000ns, min: 1000ns, avg: 7172ns
0001 -> 1024/1024 succ/totl, speed: 116.51 B/s
^Csignal: interrupt
```
### 16 clients
```bash
$ go run main.go -c 16 -s 
0004 -> beam max: 4582000ns, min: 1000ns, avg: 32251ns
0016 -> beam max: 5748000ns, min: 1000ns, avg: 32756ns
0003 -> beam max: 5881000ns, min: 1000ns, avg: 33494ns
0001 -> beam max: 4756000ns, min: 1000ns, avg: 34694ns
0014 -> beam max: 9422000ns, min: 1000ns, avg: 35479ns
0013 -> beam max: 9297000ns, min: 1000ns, avg: 36012ns
0015 -> beam max: 9427000ns, min: 1000ns, avg: 41104ns
0008 -> beam max: 14104000ns, min: 1000ns, avg: 41536ns
0009 -> beam max: 15841000ns, min: 1000ns, avg: 44399ns
0010 -> beam max: 13835000ns, min: 1000ns, avg: 46892ns
0011 -> beam max: 16868000ns, min: 1000ns, avg: 50151ns
0002 -> beam max: 17085000ns, min: 1000ns, avg: 51875ns
0005 -> beam max: 19507000ns, min: 1000ns, avg: 53383ns
0007 -> beam max: 21144000ns, min: 1000ns, avg: 55142ns
0003 -> see max: 17535000ns, min: 1000ns, avg: 61166ns
0003 -> 1024/1024 succ/totl, speed: 16.13 B/s
0016 -> see max: 14135000ns, min: 1000ns, avg: 66660ns
0016 -> 1024/1024 succ/totl, speed: 14.98 B/s
0001 -> see max: 13176000ns, min: 1000ns, avg: 66979ns
0001 -> 1024/1024 succ/totl, speed: 15.20 B/s
0006 -> beam max: 25635000ns, min: 1000ns, avg: 63666ns
0004 -> see max: 25146000ns, min: 1000ns, avg: 69759ns
0004 -> 1024/1024 succ/totl, speed: 14.36 B/s
0014 -> see max: 19035000ns, min: 1000ns, avg: 69695ns
0014 -> 1024/1024 succ/totl, speed: 14.36 B/s
0012 -> beam max: 17105000ns, min: 1000ns, avg: 63831ns
0013 -> see max: 14578000ns, min: 1000ns, avg: 72697ns
0013 -> 1024/1024 succ/totl, speed: 13.80 B/s
0015 -> see max: 14275000ns, min: 1000ns, avg: 73273ns
0015 -> 1024/1024 succ/totl, speed: 13.80 B/s
0009 -> see max: 20082000ns, min: 1000ns, avg: 76013ns
0009 -> 1024/1024 succ/totl, speed: 13.27 B/s
0010 -> see max: 15873000ns, min: 1000ns, avg: 78343ns
0010 -> 1024/1024 succ/totl, speed: 12.95 B/s
0011 -> see max: 19372000ns, min: 1000ns, avg: 78923ns
0011 -> 1024/1024 succ/totl, speed: 12.79 B/s
0002 -> see max: 18028000ns, min: 1000ns, avg: 85634ns
0002 -> 1024/1024 succ/totl, speed: 11.78 B/s
0008 -> see max: 16114000ns, min: 1000ns, avg: 85208ns
0008 -> 1024/1024 succ/totl, speed: 11.78 B/s
0005 -> see max: 19133000ns, min: 1000ns, avg: 86073ns
0005 -> 1024/1024 succ/totl, speed: 11.65 B/s
0007 -> see max: 23741000ns, min: 1000ns, avg: 87685ns
0007 -> 1024/1024 succ/totl, speed: 11.52 B/s
0006 -> see max: 25351000ns, min: 1000ns, avg: 89275ns
0006 -> 1024/1024 succ/totl, speed: 11.28 B/s
0012 -> see max: 20051000ns, min: 1000ns, avg: 85754ns
0012 -> 1024/1024 succ/totl, speed: 11.65 B/s
^Csignal: interrupt
```
### 64 clitens
```bash
$ go run main.go -c 64 -s 
0003 -> beam max: 29667000ns, min: 1000ns, avg: 101896ns
0064 -> beam max: 29414000ns, min: 1000ns, avg: 103788ns
0011 -> beam max: 36490000ns, min: 1000ns, avg: 109123ns
0004 -> beam max: 31938000ns, min: 1000ns, avg: 128744ns
0041 -> beam max: 31709000ns, min: 1000ns, avg: 149914ns
0019 -> beam max: 67266000ns, min: 1000ns, avg: 160933ns
0049 -> beam max: 36366000ns, min: 1000ns, avg: 170791ns
0051 -> beam max: 51872000ns, min: 1000ns, avg: 206817ns
0006 -> beam max: 61938000ns, min: 2000ns, avg: 221349ns
0044 -> beam max: 57328000ns, min: 1000ns, avg: 224810ns
0043 -> beam max: 51882000ns, min: 1000ns, avg: 305575ns
0034 -> beam max: 68411000ns, min: 1000ns, avg: 306165ns
0024 -> beam max: 46890000ns, min: 1000ns, avg: 311359ns
0007 -> beam max: 61936000ns, min: 1000ns, avg: 311817ns
0054 -> beam max: 79796000ns, min: 1000ns, avg: 314366ns
0033 -> beam max: 78811000ns, min: 1000ns, avg: 316100ns
0046 -> beam max: 79195000ns, min: 1000ns, avg: 318456ns
0037 -> beam max: 49029000ns, min: 1000ns, avg: 321234ns
0023 -> beam max: 81903000ns, min: 1000ns, avg: 327129ns
0013 -> beam max: 78673000ns, min: 1000ns, avg: 328734ns
0050 -> beam max: 74786000ns, min: 1000ns, avg: 330164ns
0035 -> beam max: 70077000ns, min: 1000ns, avg: 331571ns
0042 -> beam max: 70233000ns, min: 1000ns, avg: 332447ns
0005 -> beam max: 58996000ns, min: 1000ns, avg: 338054ns
0018 -> beam max: 52040000ns, min: 1000ns, avg: 348793ns
0063 -> beam max: 80578000ns, min: 1000ns, avg: 345681ns
0036 -> beam max: 63245000ns, min: 1000ns, avg: 358292ns
0025 -> beam max: 56849000ns, min: 1000ns, avg: 365402ns
0052 -> beam max: 42510000ns, min: 1000ns, avg: 364365ns
0012 -> beam max: 52832000ns, min: 1000ns, avg: 367025ns
0026 -> beam max: 62216000ns, min: 1000ns, avg: 369150ns
0045 -> beam max: 61750000ns, min: 1000ns, avg: 369556ns
0053 -> beam max: 64838000ns, min: 1000ns, avg: 370691ns
0038 -> beam max: 78348000ns, min: 1000ns, avg: 371801ns
0027 -> beam max: 79121000ns, min: 1000ns, avg: 369952ns
0008 -> beam max: 77941000ns, min: 1000ns, avg: 372970ns
0039 -> beam max: 48542000ns, min: 1000ns, avg: 380258ns
0028 -> beam max: 79157000ns, min: 1000ns, avg: 381541ns
0009 -> beam max: 78774000ns, min: 1000ns, avg: 382098ns
0047 -> beam max: 77872000ns, min: 1000ns, avg: 383130ns
0014 -> beam max: 53064000ns, min: 1000ns, avg: 382105ns
0055 -> beam max: 79360000ns, min: 1000ns, avg: 386027ns
0040 -> beam max: 57412000ns, min: 1000ns, avg: 387206ns
0048 -> beam max: 79294000ns, min: 1000ns, avg: 396267ns
0010 -> beam max: 78064000ns, min: 1000ns, avg: 399178ns
0015 -> beam max: 79136000ns, min: 1000ns, avg: 398747ns
0056 -> beam max: 81343000ns, min: 1000ns, avg: 406022ns
0020 -> beam max: 64068000ns, min: 1000ns, avg: 408491ns
0030 -> beam max: 62313000ns, min: 1000ns, avg: 408345ns
0031 -> beam max: 72794000ns, min: 1000ns, avg: 416310ns
0057 -> beam max: 71864000ns, min: 2000ns, avg: 416954ns
0029 -> beam max: 79829000ns, min: 1000ns, avg: 429810ns
0058 -> beam max: 93100000ns, min: 1000ns, avg: 429776ns
0016 -> beam max: 78470000ns, min: 1000ns, avg: 429907ns
0017 -> beam max: 96690000ns, min: 1000ns, avg: 433033ns
0002 -> beam max: 95657000ns, min: 1000ns, avg: 439057ns
0001 -> beam max: 97827000ns, min: 1000ns, avg: 439447ns
0032 -> beam max: 75267000ns, min: 1000ns, avg: 440222ns
0062 -> beam max: 93535000ns, min: 1000ns, avg: 455461ns
0022 -> beam max: 110249000ns, min: 1000ns, avg: 464425ns
0021 -> beam max: 111001000ns, min: 1000ns, avg: 465489ns
0011 -> see max: 118871000ns, min: 1000ns, avg: 488487ns
0011 -> 1024/1024 succ/totl, speed: 2.09 B/s
0064 -> see max: 122444000ns, min: 1000ns, avg: 499324ns
0064 -> 1024/1024 succ/totl, speed: 2.04 B/s
0003 -> see max: 120932000ns, min: 1000ns, avg: 502684ns
0003 -> 1024/1024 succ/totl, speed: 2.03 B/s
0019 -> see max: 121968000ns, min: 1000ns, avg: 506035ns
0019 -> 1024/1024 succ/totl, speed: 2.02 B/s
0044 -> see max: 92130000ns, min: 1000ns, avg: 511432ns
0044 -> 1024/1024 succ/totl, speed: 2.00 B/s
0049 -> see max: 60246000ns, min: 1000ns, avg: 513720ns
0049 -> 1024/1024 succ/totl, speed: 1.99 B/s
0034 -> see max: 105961000ns, min: 1000ns, avg: 516425ns
0034 -> 1024/1024 succ/totl, speed: 1.97 B/s
0024 -> see max: 62146000ns, min: 1000ns, avg: 517820ns
0024 -> 1024/1024 succ/totl, speed: 1.97 B/s
0033 -> see max: 116480000ns, min: 1000ns, avg: 532240ns
0033 -> 1024/1024 succ/totl, speed: 1.92 B/s
0004 -> see max: 80311000ns, min: 1000ns, avg: 533328ns
0004 -> 1024/1024 succ/totl, speed: 1.91 B/s
0050 -> see max: 99493000ns, min: 1000ns, avg: 534164ns
0050 -> 1024/1024 succ/totl, speed: 1.91 B/s
0041 -> see max: 67273000ns, min: 1000ns, avg: 537689ns
0041 -> 1024/1024 succ/totl, speed: 1.90 B/s
0023 -> see max: 112472000ns, min: 1000ns, avg: 541801ns
0023 -> 1024/1024 succ/totl, speed: 1.88 B/s
0010 -> see max: 98062000ns, min: 1000ns, avg: 540999ns
0010 -> 1024/1024 succ/totl, speed: 1.89 B/s
0035 -> see max: 92972000ns, min: 1000ns, avg: 550656ns
0035 -> 1024/1024 succ/totl, speed: 1.86 B/s
0042 -> see max: 95010000ns, min: 1000ns, avg: 551443ns
0042 -> 1024/1024 succ/totl, speed: 1.85 B/s
0059 -> beam max: 113021000ns, min: 1000ns, avg: 547960ns
0060 -> beam max: 116116000ns, min: 1000ns, avg: 550689ns
0005 -> see max: 80202000ns, min: 1000ns, avg: 559345ns
0005 -> 1024/1024 succ/totl, speed: 1.83 B/s
0061 -> beam max: 115454000ns, min: 1000ns, avg: 554614ns
0020 -> see max: 79514000ns, min: 1000ns, avg: 556341ns
0020 -> 1024/1024 succ/totl, speed: 1.83 B/s
0051 -> see max: 111213000ns, min: 1000ns, avg: 563039ns
0051 -> 1024/1024 succ/totl, speed: 1.81 B/s
0018 -> see max: 70089000ns, min: 1000ns, avg: 562583ns
0018 -> 1024/1024 succ/totl, speed: 1.81 B/s
0030 -> see max: 75435000ns, min: 1000ns, avg: 557880ns
0030 -> 1024/1024 succ/totl, speed: 1.83 B/s
0063 -> see max: 100392000ns, min: 1000ns, avg: 558100ns
0063 -> 1024/1024 succ/totl, speed: 1.83 B/s
0016 -> see max: 83447000ns, min: 1000ns, avg: 566511ns
0016 -> 1024/1024 succ/totl, speed: 1.80 B/s
0062 -> see max: 97250000ns, min: 1000ns, avg: 561841ns
0062 -> 1024/1024 succ/totl, speed: 1.82 B/s
0036 -> see max: 66416000ns, min: 1000ns, avg: 572165ns
0036 -> 1024/1024 succ/totl, speed: 1.78 B/s
0006 -> see max: 102442000ns, min: 1000ns, avg: 576533ns
0006 -> 1024/1024 succ/totl, speed: 1.77 B/s
0021 -> see max: 111611000ns, min: 1000ns, avg: 566706ns
0021 -> 1024/1024 succ/totl, speed: 1.80 B/s
0025 -> see max: 85424000ns, min: 1000ns, avg: 579358ns
0025 -> 1024/1024 succ/totl, speed: 1.76 B/s
0052 -> see max: 58374000ns, min: 1000ns, avg: 579929ns
0052 -> 1024/1024 succ/totl, speed: 1.76 B/s
0037 -> see max: 71424000ns, min: 1000ns, avg: 581242ns
0037 -> 1024/1024 succ/totl, speed: 1.76 B/s
0012 -> see max: 80775000ns, min: 1000ns, avg: 581700ns
0012 -> 1024/1024 succ/totl, speed: 1.76 B/s
0007 -> see max: 86386000ns, min: 1000ns, avg: 583466ns
0007 -> 1024/1024 succ/totl, speed: 1.75 B/s
0026 -> see max: 100221000ns, min: 1000ns, avg: 585842ns
0026 -> 1024/1024 succ/totl, speed: 1.74 B/s
0043 -> see max: 66714000ns, min: 1000ns, avg: 593476ns
0043 -> 1024/1024 succ/totl, speed: 1.72 B/s
0045 -> see max: 89354000ns, min: 1000ns, avg: 596085ns
0045 -> 1024/1024 succ/totl, speed: 1.71 B/s
0038 -> see max: 101756000ns, min: 1000ns, avg: 596207ns
0038 -> 1024/1024 succ/totl, speed: 1.71 B/s
0055 -> see max: 111601000ns, min: 1000ns, avg: 595841ns
0055 -> 1024/1024 succ/totl, speed: 1.71 B/s
0054 -> see max: 99047000ns, min: 1000ns, avg: 596304ns
0054 -> 1024/1024 succ/totl, speed: 1.71 B/s
0029 -> see max: 84289000ns, min: 1000ns, avg: 595361ns
0029 -> 1024/1024 succ/totl, speed: 1.72 B/s
0048 -> see max: 107921000ns, min: 1000ns, avg: 595498ns
0048 -> 1024/1024 succ/totl, speed: 1.72 B/s
0015 -> see max: 92341000ns, min: 1000ns, avg: 595827ns
0015 -> 1024/1024 succ/totl, speed: 1.71 B/s
0046 -> see max: 94828000ns, min: 1000ns, avg: 597403ns
0046 -> 1024/1024 succ/totl, speed: 1.71 B/s
0008 -> see max: 76804000ns, min: 1000ns, avg: 598713ns
0008 -> 1024/1024 succ/totl, speed: 1.71 B/s
0047 -> see max: 89142000ns, min: 1000ns, avg: 599315ns
0047 -> 1024/1024 succ/totl, speed: 1.70 B/s
0031 -> see max: 93005000ns, min: 1000ns, avg: 595987ns
0031 -> 1024/1024 succ/totl, speed: 1.71 B/s
0057 -> see max: 89159000ns, min: 1000ns, avg: 596048ns
0057 -> 1024/1024 succ/totl, speed: 1.71 B/s
0032 -> see max: 93038000ns, min: 1000ns, avg: 596549ns
0032 -> 1024/1024 succ/totl, speed: 1.71 B/s
0013 -> see max: 68624000ns, min: 1000ns, avg: 601335ns
0013 -> 1024/1024 succ/totl, speed: 1.70 B/s
0009 -> see max: 80549000ns, min: 1000ns, avg: 601666ns
0009 -> 1024/1024 succ/totl, speed: 1.70 B/s
0058 -> see max: 111324000ns, min: 1000ns, avg: 599693ns
0058 -> 1024/1024 succ/totl, speed: 1.70 B/s
0028 -> see max: 76634000ns, min: 1000ns, avg: 602690ns
0028 -> 1024/1024 succ/totl, speed: 1.69 B/s
0027 -> see max: 63003000ns, min: 1000ns, avg: 603162ns
0027 -> 1024/1024 succ/totl, speed: 1.69 B/s
0014 -> see max: 67037000ns, min: 1000ns, avg: 603986ns
0014 -> 1024/1024 succ/totl, speed: 1.69 B/s
0059 -> see max: 113015000ns, min: 1000ns, avg: 601005ns
0059 -> 1024/1024 succ/totl, speed: 1.70 B/s
0017 -> see max: 115282000ns, min: 1000ns, avg: 605733ns
0017 -> 1024/1024 succ/totl, speed: 1.69 B/s
0040 -> see max: 75363000ns, min: 1000ns, avg: 605166ns
0040 -> 1024/1024 succ/totl, speed: 1.69 B/s
0060 -> see max: 115851000ns, min: 1000ns, avg: 602280ns
0060 -> 1024/1024 succ/totl, speed: 1.70 B/s
0039 -> see max: 63637000ns, min: 1000ns, avg: 604841ns
0039 -> 1024/1024 succ/totl, speed: 1.69 B/s
0002 -> see max: 113309000ns, min: 1000ns, avg: 608207ns
0002 -> 1024/1024 succ/totl, speed: 1.68 B/s
0001 -> see max: 114241000ns, min: 1000ns, avg: 608921ns
0001 -> 1024/1024 succ/totl, speed: 1.68 B/s
0022 -> see max: 110201000ns, min: 1000ns, avg: 598102ns
0022 -> 1024/1024 succ/totl, speed: 1.71 B/s
0053 -> see max: 102941000ns, min: 1000ns, avg: 609823ns
0053 -> 1024/1024 succ/totl, speed: 1.68 B/s
0061 -> see max: 115475000ns, min: 1000ns, avg: 706398ns
0061 -> 1024/1024 succ/totl, speed: 1.45 B/s
0056 -> see max: 388860000ns, min: 1000ns, avg: 914374ns
0056 -> 1024/1024 succ/totl, speed: 1.12 B/s
^Csignal: interrupt
```
