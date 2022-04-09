#!/bin/bash
$[$1]
echo "Результат L1 - $1."
find /home/a20272/Code/github.com/eugenefoxx/readDGSP1forKATE/DGS_files/ExtractL1/ -name \*\.* -exec rm -rf {} \;
#find /mnt/dgs/00255C9D0E9D/ -name \*\.zip -exec unzip -o {} -d /home/a20272/Code/github.com/eugenefoxx/readDGS/DGS_files/ExtractL1  \;
find /mnt/dgs/00255C9D0E9D/ -name \*\.zip -mmin -10080 -exec unzip -o {} -d /home/a20272/Code/github.com/eugenefoxx/readDGSP1forKATE/DGS_files/ExtractL1  \;
#find /mnt/dgs/8CDF9DE0C303/ -name \*\.zip -exec unzip -o {} -d /home/a20272/Code/github.com/eugenefoxx/readDGS/DGS_files/Extract  \;
grep -lrin $1 /home/a20272/Code/github.com/eugenefoxx/readDGSP1forKATE/DGS_files/ExtractL1/ > /home/a20272/Code/github.com/eugenefoxx/readDGSP1forKATE/DGS_files/ExtractL1/fileresult.txt

