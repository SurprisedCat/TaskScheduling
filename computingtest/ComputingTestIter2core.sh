#! /bin/bash
TESTDOCKER=`docker ps | grep cputestIter2core`
if test -z "$TESTDOCKER"
then
    echo "Need to execute the test docker"
    echo "docker run -d -ti --name \"cputestIter2core\" --cpuset=3,4 alpine:latest"
docker run -d -ti --name "cputestIter2core" --cpuset-cpus=3,4 alpine:latest
fi
docker cp computingtest cputestIter2core:/root/
docker exec cputestIter2core touch delayTest.csv
echo "sudo cpufreq-set -c 3,4 -g userspace"
sudo cpufreq-set -c 3,4 -g userspace
CPUFREQ=2000MHz
CPUS=2
sudo cpufreq-set -c 3,4 -f $CPUFREQ
docker update cputest --cpus=2
for ITER in 100 150 200 250 300 350 400 450 500 550 600 650 700 750 800 850 900 950 1000
do
	echo "docker exec cputestIter2core /root/computingtest -d 300 -i $ITER -c $CPUS -z $CPUFREQ"
    docker exec cputestIter2core /root/computingtest -d 300 -i $ITER -c $CPUS -z $CPUFREQ
    sleep 3
done
echo "sudo cpufreq-set -c 3,4 -g ondemand"
sudo cpufreq-set -c 3,4 -g ondemand
docker cp cputestIter2core:delayTest.csv ./delayTestIter2core.csv
cat delayTestIter2core.csv
docker stop cputestIter2core
