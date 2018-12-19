#! /bin/bash
TESTDOCKER=`docker ps | grep cputest`
if test -z "$TESTDOCKER"
then
    echo "Need to execute the test docker"
    echo "docker run -d -ti --name \"cputest\" --cpuset=1 alpine:latest"
    docker run -d -ti --name "cputest" --cpuset-cpus=1 alpine:latest
fi
docker cp computingtest cputest:/root/
docker exec cputest touch delayTest.csv
echo "sudo cpufreq-set -c 1 -g userspace"
sudo cpufreq-set -c 1 -g userspace
for CPUFREQ in 1200MHz 1300MHz 1400MHz 1500MHz 1600MHz 1700MHz 1800MHz 1900MHz 2000MHz 2100MHz
do
    sudo cpufreq-set -c 1 -f $CPUFREQ
    for CPUS in 0.1 0.2 0.3 0.4 0.5 0.6 0.7 0.8 0.9 1.0
    do
        docker update cputest --cpus=$CPUS
        docker exec cputest /root/computingtest -d 200 -i 200 -c $CPUS -z $CPUFREQ
        sleep 3
        docker exec cputest /root/computingtest -d 300 -i 300 -c $CPUS -z $CPUFREQ
        sleep 3
    done
done
echo "sudo cpufreq-set -c 1 -g ondemand"
sudo cpufreq-set -c 1 -g ondemand
docker cp cputest:delayTest.csv ./delayTest.csv
cat delayTest.csv
docker stop cputest
