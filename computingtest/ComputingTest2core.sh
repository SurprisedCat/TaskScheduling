#! /bin/bash
TESTDOCKER=`docker ps | grep cputest2core`
if test -z "$TESTDOCKER"
then
    echo "Need to execute the test docker for 2 cores"
    echo "docker run -d -ti --name \"cputest2core\" --cpuset=3,4 alpine:latest"
    docker run -d -ti --name "cputest2core" --cpuset-cpus=3,4 alpine:latest
fi
docker cp computingtest cputest2core:/root/
docker exec cputest2core touch delayTest2core.csv
echo "sudo cpufreq-set -c 3,4 -g userspace"
sudo cpufreq-set -c 3,4 -g userspace
for CPUFREQ in 1200MHz 1300MHz 1400MHz 1500MHz 1600MHz 1700MHz 1800MHz 1900MHz 2000MHz 2100MHz
do
    sudo cpufreq-set -c 3,4 -f $CPUFREQ
    for CPUS in 0.8 0.9 1.0 1.1 1.2 1.3 1.4 1.5 1.6 1.7 1.8 1.9 2.0
    do
        docker update cputest2core --cpus=$CPUS
        docker exec cputest2core /root/computingtest -d 200 -i 200 -c $CPUS -z $CPUFREQ
        sleep 3
        docker exec cputest2core /root/computingtest -d 300 -i 300 -c $CPUS -z $CPUFREQ
        sleep 3
    done
done
echo "sudo cpufreq-set -c 3,4 -g ondemand"
sudo cpufreq-set -c 3,4 -g ondemand
docker cp cputest2core:delayTest2core.csv .
cat delayTest2core.csv
docker stop cputest2core