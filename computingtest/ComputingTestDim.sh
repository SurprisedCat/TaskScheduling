#! /bin/bash
TESTDOCKER=`docker ps | grep cputestDim`
if test -z "$TESTDOCKER"
then
    echo "Need to execute the test docker"
    echo "docker run -d -ti --name \"cputestDim\" --cpuset=1 alpine:latest"
    docker run -d -ti --name "cputestDim" --cpuset-cpus=1 alpine:latest
fi
docker cp computingtest cputestDim:/root/
docker exec cputestDim touch delayTest.csv
echo "sudo cpufreq-set -c 1 -g userspace"
sudo cpufreq-set -c 1 -g userspace
CPUFREQ=2000MHz
CPUS=1
sudo cpufreq-set -c 1 -f $CPUFREQ
docker update cputest --cpus=1
for DIM in 100 150 200 250 300 350 400 450 500 550 600 650 700 750 800
do
    echo "docker exec cputestDim /root/computingtest -d $DIM -i 300 -c $CPUS -z $CPUFREQ"
	docker exec cputestDim /root/computingtest -d $DIM -i 300 -c $CPUS -z $CPUFREQ
    sleep 3
done
echo "sudo cpufreq-set -c 1 -g ondemand"
sudo cpufreq-set -c 1 -g ondemand
docker cp cputestDim:delayTest.csv ./delayTestDim.csv
cat delayTestDim.csv
docker stop cputestDim
