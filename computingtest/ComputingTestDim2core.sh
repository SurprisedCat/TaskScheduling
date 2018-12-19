#! /bin/bash
TESTDOCKER=`docker ps | grep cputestDim2core`
if test -z "$TESTDOCKER"
then
    echo "Need to execute the test docker"
    echo "docker run -d -ti --name \"cputestDim2core\" --cpuset=3,4 alpine:latest"
docker run -d -ti --name "cputestDim2core" --cpuset-cpus=3,4 alpine:latest
fi
docker cp computingtest cputestDim2core:/root/
docker exec cputestDim2core touch delayTest.csv
echo "sudo cpufreq-set -c 3,4 -g userspace"
sudo cpufreq-set -c 3,4 -g userspace
CPUFREQ=2000MHz
CPUS=2
sudo cpufreq-set -c 3,4 -f $CPUFREQ
docker update cputest --cpus=2
for DIM in 100 150 200 250 300 350 400 450 500 550 600 650 700 750 800
do
	echo "docker exec cputestDim2core /root/computingtest -d $DIM -i 300 -c $CPUS -z $CPUFREQ"
    docker exec cputestDim2core /root/computingtest -d $DIM -i 300 -c $CPUS -z $CPUFREQ
    sleep 3
done
echo "sudo cpufreq-set -c 3,4 -g ondemand"
sudo cpufreq-set -c 3,4 -g ondemand
docker cp cputestDim2core:delayTest.csv ./delayTestDim2core.csv
cat delayTestDim2core.csv
docker stop cputestDim2core
