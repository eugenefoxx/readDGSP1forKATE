#!/bin/bash
$[$1]
echo "Результат L2 - $1."
find /home/a20272/Code/github.com/eugenefoxx/readDGS/DGS_files/ExtractL2/ -name \*\.* -exec rm -rf {} \;
#find /mnt/dgs/00255C9D0E9D/ -name \*\.zip -exec unzip -o {} -d /home/a20272/Code/github.com/eugenefoxx/readDGS/DGS_files/Extract  \;
find /mnt/dgs/8CDF9DE0C303/ -name \*\.zip -exec unzip -o {} -d /home/a20272/Code/github.com/eugenefoxx/readDGS/DGS_files/ExtractL2  \;
grep -lrin $1 /home/a20272/Code/github.com/eugenefoxx/readDGS/DGS_files/ExtractL2/ > /home/a20272/Code/github.com/eugenefoxx/readDGS/DGS_files/ExtractL2/fileresult.txt
