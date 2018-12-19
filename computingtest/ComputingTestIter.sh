#! /bin/bash
TESTDOCKER=`docker ps | grep cputestIter`
if test -z "$TESTDOCKER"
then
    echo "Need to execute the test docker"
    echo "docker run -d -ti --name \"cputestIter\" --cpuset=1 alpine:latest"
docker run -d -ti --name "cputestIter" --cpuset-cpus=1 alpine:latest
fi
docker cp computingtest cputestIter:/root/
docker exec cputestIter touch delayTest.csv
echo "sudo cpufreq-set -c 1 -g userspace"
sudo cpufreq-set -c 1 -g userspace
CPUFREQ=2000MHz
CPUS=1
sudo cpufreq-set -c 1 -f $CPUFREQ
docker update cputest --cpus=1
for ITER in 100 150 200 250 300 350 400 450 500 550 600 650 700 750 800 850 900 950 1000
do
	echo "docker exec cputestIter /root/computingtest -d 300 -i $ITER -c $CPUS -z $CPUFREQ"
    docker exec cputestIter /root/computingtest -d 300 -i $ITER -c $CPUS -z $CPUFREQ
    sleep 3
done
echo "sudo cpufreq-set -c 1 -g ondemand"
sudo cpufreq-set -c 1 -g ondemand
docker cp cputestIter:delayTest.csv ./delayTestIter.csv
cat delayTestIter.csv
docker stop cputestIter
