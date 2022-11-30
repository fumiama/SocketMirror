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
0001 -> beam max: 33000ns, min: 1000ns, avg: 5098ns
0001 -> see max: 88000ns, min: 1000ns, avg: 6464ns
0001 -> 1024/1024 succ/totl, speed: 131.07 B/s
^Csignal: interrupt
```
### 10 clients
```bash
$ go run main.go -c 10 -s 
0006 -> beam max: 2684000ns, min: 1000ns, avg: 24568ns
0009 -> beam max: 5195000ns, min: 1000ns, avg: 24863ns
0005 -> beam max: 5349000ns, min: 1000ns, avg: 25237ns
0008 -> beam max: 5980000ns, min: 1000ns, avg: 25811ns
0004 -> beam max: 6049000ns, min: 1000ns, avg: 30324ns
0007 -> beam max: 10628000ns, min: 1000ns, avg: 35058ns
0003 -> beam max: 11998000ns, min: 1000ns, avg: 42560ns
0002 -> beam max: 12519000ns, min: 1000ns, avg: 43179ns
0010 -> beam max: 12603000ns, min: 1000ns, avg: 43410ns
0004 -> see max: 6528000ns, min: 1000ns, avg: 51128ns
0004 -> 529/1024 succ/totl, speed: 19.42 B/s
0009 -> see max: 9102000ns, min: 1000ns, avg: 51369ns
0009 -> 552/1024 succ/totl, speed: 19.42 B/s
0006 -> see max: 5403000ns, min: 1000ns, avg: 51821ns
0006 -> 372/1024 succ/totl, speed: 19.07 B/s
0007 -> see max: 10369000ns, min: 1000ns, avg: 52113ns
0007 -> 557/1024 succ/totl, speed: 19.07 B/s
0005 -> see max: 8730000ns, min: 1000ns, avg: 54338ns
0005 -> 444/1024 succ/totl, speed: 18.40 B/s
0008 -> see max: 9091000ns, min: 1000ns, avg: 54355ns
0008 -> 545/1024 succ/totl, speed: 18.40 B/s
0003 -> see max: 11895000ns, min: 1000ns, avg: 57611ns
0003 -> 345/1024 succ/totl, speed: 17.19 B/s
0002 -> see max: 13106000ns, min: 1000ns, avg: 56289ns
0002 -> 287/1024 succ/totl, speed: 17.77 B/s
0010 -> see max: 12634000ns, min: 1000ns, avg: 56375ns
0010 -> 326/1024 succ/totl, speed: 17.77 B/s
0001 -> beam max: 23363000ns, min: 1000ns, avg: 59882ns
0001 -> see max: 15255000ns, min: 1000ns, avg: 60074ns
0001 -> 644/1024 succ/totl, speed: 16.64 B/s
^Csignal: interrupt
```
### 64 clitens
```bash
$ go run main.go -c 64 -s 
0001 -> beam max: 14682000ns, min: 1000ns, avg: 86771ns
0002 -> beam max: 36235000ns, min: 1000ns, avg: 132656ns
0013 -> beam max: 41494000ns, min: 1000ns, avg: 135347ns
0033 -> beam max: 30257000ns, min: 1000ns, avg: 140781ns
0003 -> beam max: 77510000ns, min: 1000ns, avg: 166496ns
0048 -> beam max: 56195000ns, min: 1000ns, avg: 172379ns
0022 -> beam max: 45633000ns, min: 1000ns, avg: 172704ns
0031 -> beam max: 49682000ns, min: 1000ns, avg: 174806ns
0032 -> beam max: 63546000ns, min: 1000ns, avg: 179052ns
0039 -> beam max: 57762000ns, min: 1000ns, avg: 181775ns
0004 -> beam max: 58247000ns, min: 1000ns, avg: 183431ns
0016 -> beam max: 57396000ns, min: 1000ns, avg: 315688ns
0006 -> beam max: 87473000ns, min: 1000ns, avg: 317913ns
0064 -> beam max: 77143000ns, min: 1000ns, avg: 318394ns
0010 -> beam max: 76939000ns, min: 1000ns, avg: 318707ns
0012 -> beam max: 80304000ns, min: 1000ns, avg: 319589ns
0008 -> beam max: 80197000ns, min: 1000ns, avg: 321088ns
0041 -> beam max: 52908000ns, min: 1000ns, avg: 323661ns
0029 -> beam max: 81810000ns, min: 1000ns, avg: 328235ns
0017 -> beam max: 46529000ns, min: 1000ns, avg: 328761ns
0007 -> beam max: 82227000ns, min: 1000ns, avg: 328892ns
0047 -> beam max: 83100000ns, min: 1000ns, avg: 329859ns
0037 -> beam max: 42829000ns, min: 1000ns, avg: 332366ns
0015 -> beam max: 70855000ns, min: 1000ns, avg: 332442ns
0050 -> beam max: 84833000ns, min: 1000ns, avg: 341846ns
0059 -> beam max: 76463000ns, min: 1000ns, avg: 341911ns
0024 -> beam max: 47168000ns, min: 1000ns, avg: 352522ns
0046 -> beam max: 80506000ns, min: 1000ns, avg: 352670ns
0051 -> beam max: 56376000ns, min: 1000ns, avg: 362643ns
0025 -> beam max: 61810000ns, min: 1000ns, avg: 366380ns
0053 -> beam max: 60418000ns, min: 1000ns, avg: 374306ns
0034 -> beam max: 77301000ns, min: 1000ns, avg: 375575ns
0018 -> beam max: 69377000ns, min: 1000ns, avg: 374696ns
0042 -> beam max: 70180000ns, min: 1000ns, avg: 376232ns
0054 -> beam max: 69452000ns, min: 1000ns, avg: 377533ns
0035 -> beam max: 71405000ns, min: 1000ns, avg: 381917ns
0019 -> beam max: 73373000ns, min: 1000ns, avg: 383612ns
0043 -> beam max: 77326000ns, min: 1000ns, avg: 388769ns
0055 -> beam max: 69322000ns, min: 1000ns, avg: 386700ns
0036 -> beam max: 70243000ns, min: 1000ns, avg: 389031ns
0001 -> see max: 36888000ns, min: 1000ns, avg: 394669ns
0001 -> 604/1024 succ/totl, speed: 2.58 B/s
0044 -> beam max: 70005000ns, min: 1000ns, avg: 391901ns
0057 -> beam max: 65013000ns, min: 1000ns, avg: 402239ns
0058 -> beam max: 69858000ns, min: 1000ns, avg: 408994ns
0052 -> beam max: 48483000ns, min: 1000ns, avg: 416834ns
0005 -> beam max: 83511000ns, min: 1000ns, avg: 417119ns
0002 -> see max: 39481000ns, min: 1000ns, avg: 422936ns
0002 -> 636/1024 succ/totl, speed: 2.41 B/s
0009 -> beam max: 74685000ns, min: 1000ns, avg: 418995ns
0011 -> beam max: 84048000ns, min: 1000ns, avg: 420294ns
0020 -> beam max: 84388000ns, min: 1000ns, avg: 420960ns
0038 -> beam max: 80102000ns, min: 1000ns, avg: 421985ns
0021 -> beam max: 85243000ns, min: 1000ns, avg: 422694ns
0026 -> beam max: 64284000ns, min: 1000ns, avg: 425742ns
0045 -> beam max: 60427000ns, min: 1000ns, avg: 426979ns
0030 -> beam max: 82018000ns, min: 1000ns, avg: 429291ns
0023 -> beam max: 83642000ns, min: 1000ns, avg: 433210ns
0014 -> beam max: 85056000ns, min: 1000ns, avg: 435753ns
0049 -> beam max: 85219000ns, min: 1000ns, avg: 436473ns
0060 -> beam max: 85242000ns, min: 1000ns, avg: 439773ns
0062 -> beam max: 87012000ns, min: 1000ns, avg: 445036ns
0063 -> beam max: 87327000ns, min: 1000ns, avg: 445619ns
0040 -> beam max: 55365000ns, min: 1000ns, avg: 455751ns
0003 -> see max: 87218000ns, min: 1000ns, avg: 473833ns
0003 -> 839/1024 succ/totl, speed: 2.15 B/s
0048 -> see max: 83643000ns, min: 1000ns, avg: 482628ns
0004 -> see max: 83410000ns, min: 1000ns, avg: 482599ns
0048 -> 356/1024 succ/totl, speed: 2.12 B/s
0004 -> 353/1024 succ/totl, speed: 2.12 B/s
0056 -> beam max: 56760000ns, min: 1000ns, avg: 476085ns
0027 -> beam max: 55356000ns, min: 1000ns, avg: 476758ns
0013 -> see max: 85488000ns, min: 1000ns, avg: 488441ns
0013 -> 615/1024 succ/totl, speed: 2.09 B/s
0022 -> see max: 45107000ns, min: 1000ns, avg: 490085ns
0022 -> 473/1024 succ/totl, speed: 2.08 B/s
0033 -> see max: 85567000ns, min: 1000ns, avg: 490239ns
0033 -> 407/1024 succ/totl, speed: 2.08 B/s
0028 -> beam max: 83237000ns, min: 1000ns, avg: 493925ns
0031 -> see max: 85056000ns, min: 1000ns, avg: 504903ns
0031 -> 617/1024 succ/totl, speed: 2.02 B/s
0039 -> see max: 87517000ns, min: 1000ns, avg: 508285ns
0039 -> 569/1024 succ/totl, speed: 2.01 B/s
0061 -> beam max: 85998000ns, min: 1000ns, avg: 499727ns
0032 -> see max: 62893000ns, min: 1000ns, avg: 511575ns
0032 -> 443/1024 succ/totl, speed: 2.00 B/s
0024 -> see max: 47904000ns, min: 1000ns, avg: 519071ns
0024 -> 611/1024 succ/totl, speed: 1.97 B/s
0043 -> see max: 76069000ns, min: 1000ns, avg: 518501ns
0043 -> 566/1024 succ/totl, speed: 1.97 B/s
0016 -> see max: 49891000ns, min: 1000ns, avg: 519596ns
0016 -> 402/1024 succ/totl, speed: 1.97 B/s
0034 -> see max: 72407000ns, min: 1000ns, avg: 519303ns
0034 -> 651/1024 succ/totl, speed: 1.96 B/s
0025 -> see max: 60259000ns, min: 1000ns, avg: 520085ns
0025 -> 552/1024 succ/totl, speed: 1.96 B/s
0051 -> see max: 55583000ns, min: 1000ns, avg: 520442ns
0051 -> 575/1024 succ/totl, speed: 1.96 B/s
0036 -> see max: 59156000ns, min: 1000ns, avg: 519112ns
0036 -> 496/1024 succ/totl, speed: 1.97 B/s
0044 -> see max: 66037000ns, min: 1000ns, avg: 519450ns
0044 -> 494/1024 succ/totl, speed: 1.97 B/s
0064 -> see max: 77121000ns, min: 1000ns, avg: 522657ns
0064 -> 356/1024 succ/totl, speed: 1.95 B/s
0010 -> see max: 76912000ns, min: 1000ns, avg: 522350ns
0010 -> 325/1024 succ/totl, speed: 1.96 B/s
0053 -> see max: 62914000ns, min: 1000ns, avg: 521995ns
0053 -> 477/1024 succ/totl, speed: 1.96 B/s
0006 -> see max: 84368000ns, min: 1000ns, avg: 523990ns
0006 -> 699/1024 succ/totl, speed: 1.95 B/s
0019 -> see max: 71605000ns, min: 1000ns, avg: 521715ns
0019 -> 392/1024 succ/totl, speed: 1.96 B/s
0012 -> see max: 80399000ns, min: 1000ns, avg: 524020ns
0012 -> 447/1024 succ/totl, speed: 1.95 B/s
0008 -> see max: 80544000ns, min: 2000ns, avg: 524100ns
0008 -> 433/1024 succ/totl, speed: 1.95 B/s
0035 -> see max: 64102000ns, min: 1000ns, avg: 522644ns
0035 -> 683/1024 succ/totl, speed: 1.96 B/s
0041 -> see max: 54117000ns, min: 1000ns, avg: 523514ns
0041 -> 619/1024 succ/totl, speed: 1.95 B/s
0057 -> see max: 64823000ns, min: 1000ns, avg: 519689ns
0057 -> 615/1024 succ/totl, speed: 1.97 B/s
0026 -> see max: 64250000ns, min: 1000ns, avg: 519698ns
0026 -> 667/1024 succ/totl, speed: 1.97 B/s
0029 -> see max: 81827000ns, min: 1000ns, avg: 527189ns
0029 -> 422/1024 succ/totl, speed: 1.94 B/s
0017 -> see max: 47105000ns, min: 1000ns, avg: 526092ns
0017 -> 589/1024 succ/totl, speed: 1.94 B/s
0054 -> see max: 56266000ns, min: 1000ns, avg: 525223ns
0054 -> 533/1024 succ/totl, speed: 1.94 B/s
0058 -> see max: 69916000ns, min: 1000ns, avg: 522414ns
0058 -> 446/1024 succ/totl, speed: 1.95 B/s
0037 -> see max: 45439000ns, min: 1000ns, avg: 528262ns
0037 -> 510/1024 succ/totl, speed: 1.93 B/s
0007 -> see max: 82259000ns, min: 1000ns, avg: 529420ns
0042 -> see max: 52295000ns, min: 1000ns, avg: 528143ns
0042 -> 504/1024 succ/totl, speed: 1.93 B/s
0007 -> 385/1024 succ/totl, speed: 1.93 B/s
0047 -> see max: 83137000ns, min: 1000ns, avg: 530704ns
0047 -> 456/1024 succ/totl, speed: 1.92 B/s
0055 -> see max: 54684000ns, min: 1000ns, avg: 528012ns
0055 -> 609/1024 succ/totl, speed: 1.93 B/s
0052 -> see max: 48410000ns, min: 1000ns, avg: 529167ns
0052 -> 550/1024 succ/totl, speed: 1.93 B/s
0005 -> see max: 83511000ns, min: 1000ns, avg: 530553ns
0005 -> 293/1024 succ/totl, speed: 1.92 B/s
0015 -> see max: 70116000ns, min: 1000ns, avg: 530387ns
0015 -> 516/1024 succ/totl, speed: 1.92 B/s
0059 -> see max: 76846000ns, min: 1000ns, avg: 525014ns
0018 -> see max: 47011000ns, min: 1000ns, avg: 528425ns
0018 -> 584/1024 succ/totl, speed: 1.93 B/s
0059 -> 476/1024 succ/totl, speed: 1.95 B/s
0009 -> see max: 74356000ns, min: 1000ns, avg: 530715ns
0009 -> 423/1024 succ/totl, speed: 1.92 B/s
0011 -> see max: 84015000ns, min: 1000ns, avg: 530647ns
0011 -> 326/1024 succ/totl, speed: 1.93 B/s
0040 -> see max: 55228000ns, min: 1000ns, avg: 529445ns
0040 -> 657/1024 succ/totl, speed: 1.93 B/s
0020 -> see max: 84448000ns, min: 1000ns, avg: 531613ns
0020 -> 404/1024 succ/totl, speed: 1.92 B/s
0038 -> see max: 79861000ns, min: 1000ns, avg: 531066ns
0038 -> 574/1024 succ/totl, speed: 1.92 B/s
0021 -> see max: 85131000ns, min: 1000ns, avg: 528377ns
0021 -> 493/1024 succ/totl, speed: 1.93 B/s
0050 -> see max: 84840000ns, min: 1000ns, avg: 531170ns
0050 -> 450/1024 succ/totl, speed: 1.92 B/s
0028 -> see max: 83213000ns, min: 1000ns, avg: 525645ns
0028 -> 707/1024 succ/totl, speed: 1.94 B/s
0030 -> see max: 82025000ns, min: 1000ns, avg: 529514ns
0030 -> 634/1024 succ/totl, speed: 1.93 B/s
0023 -> see max: 83595000ns, min: 1000ns, avg: 530341ns
0023 -> 605/1024 succ/totl, speed: 1.92 B/s
0014 -> see max: 85045000ns, min: 1000ns, avg: 532505ns
0014 -> 550/1024 succ/totl, speed: 1.92 B/s
0046 -> see max: 84948000ns, min: 1000ns, avg: 532416ns
0046 -> 525/1024 succ/totl, speed: 1.92 B/s
0049 -> see max: 85045000ns, min: 1000ns, avg: 533318ns
0049 -> 686/1024 succ/totl, speed: 1.92 B/s
0060 -> see max: 85824000ns, min: 1000ns, avg: 526000ns
0060 -> 802/1024 succ/totl, speed: 1.94 B/s
0062 -> see max: 88132000ns, min: 1000ns, avg: 526119ns
0062 -> 791/1024 succ/totl, speed: 1.94 B/s
0063 -> see max: 87569000ns, min: 1000ns, avg: 526227ns
0063 -> 901/1024 succ/totl, speed: 1.94 B/s
0056 -> see max: 101381000ns, min: 1000ns, avg: 618336ns
0056 -> 537/1024 succ/totl, speed: 1.65 B/s
0027 -> see max: 101204000ns, min: 1000ns, avg: 625924ns
0027 -> 665/1024 succ/totl, speed: 1.63 B/s
0061 -> see max: 96634000ns, min: 1000ns, avg: 624208ns
0061 -> 890/1024 succ/totl, speed: 1.64 B/s
0045 -> see max: 101304000ns, min: 1000ns, avg: 634986ns
0045 -> 755/1024 succ/totl, speed: 1.61 B/s
^Csignal: interrupt
```
