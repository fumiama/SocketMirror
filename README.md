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
$ go run main.go log_unix.go -c 1 
INFO [0001] beam max: 60000ns, min: 3000ns, avg: 7481ns
INFO [0001] see max: 146000ns, min: 1000ns, avg: 7244ns
INFO [0001] 1024/1024 succ/totl, speed: 113777.78 KB/s
```
### 2 clients
```bash
$ go run main.go log_unix.go -c 2
INFO [0001] beam max: 260000ns, min: 1000ns, avg: 6628ns
INFO [0002] beam max: 169000ns, min: 1000ns, avg: 7230ns
INFO [0001] see max: 1025000ns, min: 1000ns, avg: 10184ns
INFO [0001] 1024/1024 succ/totl, speed: 85333.33 KB/s
INFO [0002] see max: 3846000ns, min: 1000ns, avg: 13315ns
INFO [0002] 1024/1024 succ/totl, speed: 68266.67 KB/s
```
### 16 clients
```bash
$ go run main.go log_unix.go -c 16
INFO [0001] beam max: 7161000ns, min: 1000ns, avg: 53967ns
INFO [0016] beam max: 8828000ns, min: 1000ns, avg: 55760ns
INFO [0008] beam max: 26271000ns, min: 2000ns, avg: 79741ns
INFO [0004] beam max: 25361000ns, min: 1000ns, avg: 80464ns
INFO [0005] beam max: 27885000ns, min: 1000ns, avg: 82784ns
INFO [0002] beam max: 33236000ns, min: 1000ns, avg: 92153ns
INFO [0009] beam max: 32281000ns, min: 1000ns, avg: 94174ns
INFO [0010] beam max: 31622000ns, min: 2000ns, avg: 97208ns
INFO [0011] beam max: 31533000ns, min: 1000ns, avg: 99743ns
INFO [0015] beam max: 40255000ns, min: 2000ns, avg: 106580ns
INFO [0006] beam max: 35301000ns, min: 1000ns, avg: 109369ns
INFO [0003] beam max: 37556000ns, min: 1000ns, avg: 114556ns
INFO [0014] beam max: 29130000ns, min: 1000ns, avg: 115740ns
INFO [0013] beam max: 28519000ns, min: 1000ns, avg: 115935ns
INFO [0012] beam max: 43665000ns, min: 1000ns, avg: 120876ns
INFO [0001] see max: 31974000ns, min: 1000ns, avg: 126070ns
INFO [0001] 1024/1024 succ/totl, speed: 7816.79 KB/s
INFO [0008] see max: 36317000ns, min: 1000ns, avg: 129766ns
INFO [0008] 1024/1024 succ/totl, speed: 7585.19 KB/s
INFO [0004] see max: 38623000ns, min: 2000ns, avg: 133182ns
INFO [0004] 1024/1024 succ/totl, speed: 7420.29 KB/s
INFO [0016] see max: 43214000ns, min: 1000ns, avg: 136326ns
INFO [0016] 1024/1024 succ/totl, speed: 7211.27 KB/s
INFO [0005] see max: 34866000ns, min: 1000ns, avg: 137344ns
INFO [0005] 1024/1024 succ/totl, speed: 7211.27 KB/s
INFO [0007] beam max: 44927000ns, min: 1000ns, avg: 117091ns
INFO [0002] see max: 33814000ns, min: 2000ns, avg: 142317ns
INFO [0002] 1024/1024 succ/totl, speed: 6918.92 KB/s
INFO [0009] see max: 33241000ns, min: 2000ns, avg: 141835ns
INFO [0009] 1024/1024 succ/totl, speed: 6965.99 KB/s
INFO [0010] see max: 32862000ns, min: 2000ns, avg: 144799ns
INFO [0010] 1024/1024 succ/totl, speed: 6826.67 KB/s
INFO [0011] see max: 31368000ns, min: 2000ns, avg: 152928ns
INFO [0011] 1024/1024 succ/totl, speed: 6440.25 KB/s
INFO [0015] see max: 40166000ns, min: 1000ns, avg: 155754ns
INFO [0015] 1024/1024 succ/totl, speed: 6320.99 KB/s
INFO [0003] see max: 39549000ns, min: 2000ns, avg: 159208ns
INFO [0003] 1024/1024 succ/totl, speed: 6168.67 KB/s
INFO [0012] see max: 42214000ns, min: 2000ns, avg: 157546ns
INFO [0012] 1024/1024 succ/totl, speed: 6243.90 KB/s
INFO [0013] see max: 28259000ns, min: 1000ns, avg: 159937ns
INFO [0013] 1024/1024 succ/totl, speed: 6131.74 KB/s
INFO [0014] see max: 30461000ns, min: 2000ns, avg: 160645ns
INFO [0014] 1024/1024 succ/totl, speed: 6131.74 KB/s
INFO [0006] see max: 35183000ns, min: 1000ns, avg: 169056ns
INFO [0006] 1024/1024 succ/totl, speed: 5851.43 KB/s
INFO [0007] see max: 46509000ns, min: 2000ns, avg: 152061ns
INFO [0007] 1024/1024 succ/totl, speed: 6481.01 KB/s
```
### 64 clitens
```bash
$ go run main.go log_unix.go -c 64
INFO [0031] beam max: 37188000ns, min: 2000ns, avg: 182900ns
INFO [0030] beam max: 42241000ns, min: 1000ns, avg: 186544ns
INFO [0047] beam max: 55542000ns, min: 1000ns, avg: 189834ns
INFO [0009] beam max: 87534000ns, min: 1000ns, avg: 235896ns
INFO [0049] beam max: 68275000ns, min: 2000ns, avg: 256399ns
INFO [0043] beam max: 112447000ns, min: 1000ns, avg: 277849ns
INFO [0003] beam max: 72924000ns, min: 1000ns, avg: 278166ns
INFO [0011] beam max: 73900000ns, min: 1000ns, avg: 283889ns
INFO [0055] beam max: 114959000ns, min: 2000ns, avg: 304489ns
INFO [0033] beam max: 87846000ns, min: 1000ns, avg: 306791ns
INFO [0041] beam max: 99476000ns, min: 1000ns, avg: 333479ns
INFO [0042] beam max: 124490000ns, min: 2000ns, avg: 535670ns
INFO [0044] beam max: 133831000ns, min: 1000ns, avg: 545324ns
INFO [0010] beam max: 134561000ns, min: 2000ns, avg: 545850ns
INFO [0048] beam max: 136772000ns, min: 1000ns, avg: 547990ns
INFO [0052] beam max: 92226000ns, min: 2000ns, avg: 550473ns
INFO [0036] beam max: 94187000ns, min: 1000ns, avg: 553296ns
INFO [0053] beam max: 86200000ns, min: 2000ns, avg: 554791ns
INFO [0038] beam max: 140078000ns, min: 2000ns, avg: 557168ns
INFO [0013] beam max: 87000000ns, min: 2000ns, avg: 557413ns
INFO [0039] beam max: 131076000ns, min: 1000ns, avg: 558979ns
INFO [0050] beam max: 140499000ns, min: 2000ns, avg: 558837ns
INFO [0034] beam max: 110211000ns, min: 1000ns, avg: 558917ns
INFO [0008] beam max: 140937000ns, min: 1000ns, avg: 561116ns
INFO [0054] beam max: 76657000ns, min: 1000ns, avg: 561809ns
INFO [0064] beam max: 149343000ns, min: 2000ns, avg: 565130ns
INFO [0045] beam max: 140006000ns, min: 1000ns, avg: 572181ns
INFO [0058] beam max: 148821000ns, min: 1000ns, avg: 571833ns
INFO [0007] beam max: 149818000ns, min: 2000ns, avg: 572577ns
INFO [0006] beam max: 149531000ns, min: 1000ns, avg: 574657ns
INFO [0001] beam max: 146477000ns, min: 2000ns, avg: 584788ns
INFO [0046] beam max: 149694000ns, min: 1000ns, avg: 585994ns
INFO [0012] beam max: 93053000ns, min: 1000ns, avg: 589089ns
INFO [0004] beam max: 109219000ns, min: 2000ns, avg: 592615ns
INFO [0056] beam max: 147153000ns, min: 2000ns, avg: 589628ns
INFO [0059] beam max: 114140000ns, min: 1000ns, avg: 603981ns
INFO [0022] beam max: 135147000ns, min: 2000ns, avg: 604338ns
INFO [0030] see max: 73869000ns, min: 1000ns, avg: 619750ns
INFO [0030] 1024/1024 succ/totl, speed: 1607.54 KB/s
INFO [0040] beam max: 79147000ns, min: 2000ns, avg: 619744ns
INFO [0025] beam max: 140048000ns, min: 1000ns, avg: 620251ns
INFO [0035] beam max: 83532000ns, min: 2000ns, avg: 637767ns
INFO [0014] beam max: 130760000ns, min: 2000ns, avg: 653920ns
INFO [0037] beam max: 121437000ns, min: 1000ns, avg: 653507ns
INFO [0015] beam max: 119185000ns, min: 2000ns, avg: 654193ns
INFO [0009] see max: 131931000ns, min: 1000ns, avg: 663431ns
INFO [0009] 1024/1024 succ/totl, speed: 1503.67 KB/s
INFO [0049] see max: 125796000ns, min: 1000ns, avg: 668593ns
INFO [0049] 1024/1024 succ/totl, speed: 1492.71 KB/s
INFO [0031] see max: 111550000ns, min: 1000ns, avg: 667835ns
INFO [0031] 1024/1024 succ/totl, speed: 1492.71 KB/s
INFO [0062] beam max: 118464000ns, min: 1000ns, avg: 662963ns
INFO [0063] beam max: 113999000ns, min: 1000ns, avg: 668588ns
INFO [0018] beam max: 122088000ns, min: 2000ns, avg: 673336ns
INFO [0019] beam max: 130430000ns, min: 1000ns, avg: 676822ns
INFO [0057] beam max: 153908000ns, min: 2000ns, avg: 679166ns
INFO [0051] beam max: 149286000ns, min: 1000ns, avg: 683707ns
INFO [0020] beam max: 144997000ns, min: 1000ns, avg: 682708ns
INFO [0032] beam max: 153144000ns, min: 2000ns, avg: 688314ns
INFO [0002] beam max: 150864000ns, min: 1000ns, avg: 689957ns
INFO [0016] beam max: 112859000ns, min: 2000ns, avg: 692456ns
INFO [0021] beam max: 134827000ns, min: 1000ns, avg: 691786ns
INFO [0061] beam max: 120159000ns, min: 1000ns, avg: 702956ns
INFO [0024] beam max: 145852000ns, min: 1000ns, avg: 707585ns
INFO [0028] beam max: 143599000ns, min: 1000ns, avg: 696556ns
INFO [0027] beam max: 144509000ns, min: 2000ns, avg: 697747ns
INFO [0026] beam max: 144076000ns, min: 2000ns, avg: 700484ns
INFO [0060] beam max: 121303000ns, min: 2000ns, avg: 710484ns
INFO [0023] beam max: 144389000ns, min: 1000ns, avg: 712058ns
INFO [0033] see max: 109798000ns, min: 1000ns, avg: 724587ns
INFO [0033] 1024/1024 succ/totl, speed: 1378.20 KB/s
INFO [0055] see max: 139316000ns, min: 1000ns, avg: 725754ns
INFO [0055] 1024/1024 succ/totl, speed: 1374.50 KB/s
INFO [0047] see max: 131515000ns, min: 1000ns, avg: 726190ns
INFO [0047] 1024/1024 succ/totl, speed: 1374.50 KB/s
INFO [0029] beam max: 138345000ns, min: 2000ns, avg: 698177ns
INFO [0043] see max: 126768000ns, min: 1000ns, avg: 733999ns
INFO [0043] 1024/1024 succ/totl, speed: 1358.09 KB/s
INFO [0003] see max: 106981000ns, min: 1000ns, avg: 734111ns
INFO [0003] 1024/1024 succ/totl, speed: 1358.09 KB/s
INFO [0054] see max: 78727000ns, min: 1000ns, avg: 732621ns
INFO [0054] 1024/1024 succ/totl, speed: 1361.70 KB/s
INFO [0064] see max: 149331000ns, min: 1000ns, avg: 734904ns
INFO [0064] 1024/1024 succ/totl, speed: 1356.29 KB/s
INFO [0042] see max: 118091000ns, min: 1000ns, avg: 734678ns
INFO [0042] 1024/1024 succ/totl, speed: 1358.09 KB/s
INFO [0034] see max: 109744000ns, min: 1000ns, avg: 734261ns
INFO [0034] 1024/1024 succ/totl, speed: 1358.09 KB/s
INFO [0045] see max: 138993000ns, min: 2000ns, avg: 736272ns
INFO [0045] 1024/1024 succ/totl, speed: 1352.71 KB/s
INFO [0035] see max: 87283000ns, min: 1000ns, avg: 736247ns
INFO [0035] 1024/1024 succ/totl, speed: 1354.50 KB/s
INFO [0050] see max: 145682000ns, min: 2000ns, avg: 737625ns
INFO [0050] 1024/1024 succ/totl, speed: 1350.92 KB/s
INFO [0008] see max: 140507000ns, min: 1000ns, avg: 743475ns
INFO [0008] 1024/1024 succ/totl, speed: 1340.31 KB/s
INFO [0006] see max: 149594000ns, min: 2000ns, avg: 743456ns
INFO [0006] 1024/1024 succ/totl, speed: 1342.07 KB/s
INFO [0013] see max: 84603000ns, min: 1000ns, avg: 743812ns
INFO [0013] 1024/1024 succ/totl, speed: 1340.31 KB/s
INFO [0057] see max: 154330000ns, min: 1000ns, avg: 748564ns
INFO [0057] 1024/1024 succ/totl, speed: 1331.60 KB/s
INFO [0036] see max: 93373000ns, min: 1000ns, avg: 749031ns
INFO [0036] 1024/1024 succ/totl, speed: 1331.60 KB/s
INFO [0032] see max: 153192000ns, min: 1000ns, avg: 750166ns
INFO [0032] 1024/1024 succ/totl, speed: 1329.87 KB/s
INFO [0053] see max: 77314000ns, min: 2000ns, avg: 751488ns
INFO [0053] 1024/1024 succ/totl, speed: 1328.15 KB/s
INFO [0002] see max: 152738000ns, min: 2000ns, avg: 753428ns
INFO [0002] 1024/1024 succ/totl, speed: 1323.00 KB/s
INFO [0052] see max: 91435000ns, min: 2000ns, avg: 753113ns
INFO [0052] 1024/1024 succ/totl, speed: 1323.00 KB/s
INFO [0014] see max: 102825000ns, min: 1000ns, avg: 753715ns
INFO [0014] 1024/1024 succ/totl, speed: 1323.00 KB/s
INFO [0041] see max: 76787000ns, min: 1000ns, avg: 755765ns
INFO [0041] 1024/1024 succ/totl, speed: 1319.59 KB/s
INFO [0038] see max: 138743000ns, min: 2000ns, avg: 757900ns
INFO [0038] 1024/1024 succ/totl, speed: 1316.20 KB/s
INFO [0039] see max: 130395000ns, min: 2000ns, avg: 757713ns
INFO [0039] 1024/1024 succ/totl, speed: 1316.20 KB/s
INFO [0025] see max: 139892000ns, min: 1000ns, avg: 748614ns
INFO [0025] 1024/1024 succ/totl, speed: 1331.60 KB/s
INFO [0012] see max: 94129000ns, min: 1000ns, avg: 758599ns
INFO [0012] 1024/1024 succ/totl, speed: 1314.51 KB/s
INFO [0037] see max: 122789000ns, min: 2000ns, avg: 752620ns
INFO [0037] 1024/1024 succ/totl, speed: 1324.71 KB/s
INFO [0007] see max: 149807000ns, min: 1000ns, avg: 763545ns
INFO [0007] 1024/1024 succ/totl, speed: 1306.12 KB/s
INFO [0058] see max: 149137000ns, min: 2000ns, avg: 762250ns
INFO [0058] 1024/1024 succ/totl, speed: 1306.12 KB/s
INFO [0051] see max: 149585000ns, min: 1000ns, avg: 764835ns
INFO [0051] 1024/1024 succ/totl, speed: 1304.46 KB/s
INFO [0010] see max: 137973000ns, min: 2000ns, avg: 766040ns
INFO [0010] 1024/1024 succ/totl, speed: 1301.14 KB/s
INFO [0024] see max: 145865000ns, min: 1000ns, avg: 757960ns
INFO [0024] 1024/1024 succ/totl, speed: 1316.20 KB/s
INFO [0044] see max: 129739000ns, min: 1000ns, avg: 769332ns
INFO [0044] 1024/1024 succ/totl, speed: 1296.20 KB/s
INFO [0001] see max: 146599000ns, min: 2000ns, avg: 770353ns
INFO [0001] 1024/1024 succ/totl, speed: 1294.56 KB/s
INFO [0004] see max: 111608000ns, min: 1000ns, avg: 769993ns
INFO [0004] 1024/1024 succ/totl, speed: 1294.56 KB/s
INFO [0056] see max: 147086000ns, min: 1000ns, avg: 775590ns
INFO [0056] 1024/1024 succ/totl, speed: 1286.43 KB/s
INFO [0060] see max: 122410000ns, min: 1000ns, avg: 775435ns
INFO [0060] 1024/1024 succ/totl, speed: 1286.43 KB/s
INFO [0040] see max: 84407000ns, min: 2000ns, avg: 782478ns
INFO [0040] 1024/1024 succ/totl, speed: 1273.63 KB/s
INFO [0015] see max: 121071000ns, min: 2000ns, avg: 775167ns
INFO [0015] 1024/1024 succ/totl, speed: 1286.43 KB/s
INFO [0061] see max: 121192000ns, min: 1000ns, avg: 735114ns
INFO [0061] 1024/1024 succ/totl, speed: 1286.43 KB/s
INFO [0048] see max: 144104000ns, min: 2000ns, avg: 786148ns
INFO [0048] 1024/1024 succ/totl, speed: 1268.90 KB/s
INFO [0063] see max: 115423000ns, min: 2000ns, avg: 779117ns
INFO [0063] 1024/1024 succ/totl, speed: 1280.00 KB/s
INFO [0018] see max: 123249000ns, min: 2000ns, avg: 779822ns
INFO [0018] 1024/1024 succ/totl, speed: 1278.40 KB/s
INFO [0019] see max: 130903000ns, min: 2000ns, avg: 788788ns
INFO [0019] 1024/1024 succ/totl, speed: 1264.20 KB/s
INFO [0021] see max: 134743000ns, min: 1000ns, avg: 789694ns
INFO [0021] 1024/1024 succ/totl, speed: 1262.64 KB/s
INFO [0020] see max: 145433000ns, min: 2000ns, avg: 794361ns
INFO [0020] 1024/1024 succ/totl, speed: 1254.90 KB/s
INFO [0026] see max: 143486000ns, min: 1000ns, avg: 776143ns
INFO [0026] 1024/1024 succ/totl, speed: 1284.82 KB/s
INFO [0046] see max: 149719000ns, min: 1000ns, avg: 800805ns
INFO [0046] 1024/1024 succ/totl, speed: 1245.74 KB/s
INFO [0028] see max: 143601000ns, min: 1000ns, avg: 773111ns
INFO [0028] 1024/1024 succ/totl, speed: 1289.67 KB/s
INFO [0016] see max: 113485000ns, min: 1000ns, avg: 793376ns
INFO [0016] 1024/1024 succ/totl, speed: 1257.99 KB/s
INFO [0027] see max: 144513000ns, min: 2000ns, avg: 774719ns
INFO [0027] 1024/1024 succ/totl, speed: 1288.05 KB/s
INFO [0011] see max: 125465000ns, min: 1000ns, avg: 801448ns
INFO [0011] 1024/1024 succ/totl, speed: 1244.23 KB/s
INFO [0059] see max: 113011000ns, min: 1000ns, avg: 800336ns
INFO [0059] 1024/1024 succ/totl, speed: 1245.74 KB/s
INFO [0005] beam max: 105393000ns, min: 1000ns, avg: 798285ns
INFO [0017] beam max: 116842000ns, min: 2000ns, avg: 799027ns
INFO [0005] see max: 106295000ns, min: 2000ns, avg: 804079ns
INFO [0005] 1024/1024 succ/totl, speed: 1239.71 KB/s
INFO [0017] see max: 118647000ns, min: 1000ns, avg: 799914ns
INFO [0017] 1024/1024 succ/totl, speed: 1245.74 KB/s
INFO [0029] see max: 138417000ns, min: 1000ns, avg: 772489ns
INFO [0029] 1024/1024 succ/totl, speed: 1291.30 KB/s
INFO [0062] see max: 120671000ns, min: 1000ns, avg: 800584ns
INFO [0062] 1024/1024 succ/totl, speed: 1245.74 KB/s
INFO [0022] see max: 134952000ns, min: 1000ns, avg: 799019ns
INFO [0022] 1024/1024 succ/totl, speed: 1248.78 KB/s
INFO [0023] see max: 144094000ns, min: 1000ns, avg: 892574ns
INFO [0023] 1024/1024 succ/totl, speed: 1116.68 KB/s
```
