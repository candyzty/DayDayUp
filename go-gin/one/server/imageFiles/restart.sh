#!/bin/bash
#File: update or rollback server version
#Date: 2019-04-18
#Auth: gaopengcheng

source  /etc/profile



APP_NAME=takeaway
SERVER_PORT=6066
INSTALL_DIR=/opt/guming
CODE_HOME=/opt/guming/${APP_NAME}/target
File_localtion=${INSTALL_DIR}/$APP_NAME/file
DATE=$(date "+%Y-%m-%d-%H-%M-%S")
VERSION=$(ls ${File_localtion} |grep ${APP_NAME}.*jar)
BACK_FILE=$(ls ${INSTALL_DIR}/${APP_NAME} |grep "target-*"|uniq |sort  -rn|head  -n 1)
JAVA_JAR='/opt/guming/takeaway/target/takeaway.jar'
PARRM=$(ps -ef|grep  ${CODE_HOME}/${APP_NAME}.jar |grep -v grep|awk '{print $2}')
JVM_OPTS='-server
          -Xms1024M
          -Xmx1024M
          -Xss256k
          -XX:NewSize=140M
          -XX:MaxNewSize=140M
          -XX:SurvivorRatio=8
          -XX:CMSInitiatingOccupancyFraction=50
          -XX:+AggressiveOpts
          -XX:+UseBiasedLocking
          -XX:+DisableExplicitGC
          -XX:+UseParNewGC
          -XX:+UseConcMarkSweepGC
          -XX:MaxTenuringThreshold=15
          -XX:+CMSParallelRemarkEnabled
          -XX:+UseCMSCompactAtFullCollection
          -XX:LargePageSizeInBytes=128m
          -XX:+UseFastAccessorMethods
          -XX:+UseCMSInitiatingOccupancyOnly
          -Duser.timezone=Asia/Shanghai
          -XX:+PrintGCTimeStamps
          -XX:+PrintGCDetails
          -XX:+UseGCLogFileRotation
          -XX:GCLogFileSize=512M
          -Xloggc:/log/admin/gc.log
          -XX:+HeapDumpOnOutOfMemoryError
          -XX:HeapDumpPath=/log/admin/heap/
          -Djava.awt.headless=true'

#JVM_OPTS='-server -Xms375M -Xmx375M -Xss256k -XX:NewSize=140M -XX:MaxNewSize=140M -XX:SurvivorRatio=8 -XX:CMSInitiatingOccupancyFraction=50 -XX:+AggressiveOpts -XX:+UseBiasedLocking -XX:+DisableExplicitGC -XX:+UseParNewGC -XX:+UseConcMarkSweepGC -XX:MaxTenuringThreshold=15 -XX:+CMSParallelRemarkEnabled -XX:+UseCMSCompactAtFullCollection -XX:LargePageSizeInBytes=128m -XX:+UseFastAccessorMethods -XX:+UseCMSInitiatingOccupancyOnly -Duser.timezone=Asia/Shanghai -XX:+PrintGCTimeStamps -XX:+PrintGCDetails -XX:+UseGCLogFileRotation -XX:GCLogFileSize=512M -Xloggc:/log/admin/gc.log -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/log/admin/heap/ -Djava.awt.headless=true'

stop_server(){
    echo "----Stop server----"
    PARRM=$(ps -ef |grep java|grep ${JAVA_JAR}|grep -v grep|awk '{print $2}')
    count=$(ps -ef |grep java|grep ${JAVA_JAR}|grep -v grep|wc -l)
    if [ ${count}  != 0 ];then
        echo "Stop App :${APP_NAME}"
        kill ${PARRM}
       sleep 30
    else
        echo "App has stopped...."
    fi
}

start_server() {
    echo "----Start serever----"
    count=$(ps -ef |grep java|grep ${JAVA_JAR}|grep -v grep|wc -l)
    if [   ${count} != 0  ];then
        echo "App  running  ...."
    else
        nohup java -jar $JVM_OPTS ${CODE_HOME}/${APP_NAME}.jar --spring.profiles.active=pro > /dev/null  2>&1 &
        sleep  30
        echo "start app ...."
    fi
}
SERVER_CHECK(){
    #根据url 与端口检查来进行服务存活的判断  
    URL=127.0.0.1:${SERVER_PORT}/actuator/health
    HOST_IP=$(ip a |grep eth0 |grep inet |awk  -F /  '{print $1}'|awk '{print $2}')
    LOGS_DIR=/opt/Detection_url/${APP_NAME}/${APP_NAME}
    Status_code=$(curl -I -m 10 -o /dev/null -s -w %{http_code} ${URL})

    if [ ${Status_code} -eq 200 ] ;then
        echo   "$(date "+%Y-%m-%d-%H:%M:%S")   
项目名称：${APP_NAME}  
服务 IP : ${HOST_IP} 
URL 状态: ${Status_code}  
端口检查: ${SERVER_PORT}"  
    else
        echo  "$(date "+%Y-%m-%d-%H:%M:%S")   ${APP_NAME}    运行失败"  
    fi
}

Update(){
    stop_server
    echo "----Update server----"
    cd ${INSTALL_DIR}
    mv ${APP_NAME}/target    ${APP_NAME}/target-${DATE}
    mkdir   ${CODE_HOME}
    mv ${File_localtion}/${VERSION}   ${CODE_HOME}/${APP_NAME}.jar
    BACK_FILE=$(ls ${INSTALL_DIR}/${APP_NAME} |grep "target_*"|uniq |sort  -rn|head  -n 1)
    echo -e "\033[32m The backup file for this upgrade is:  ${BACK_FILE} \033[0m"
    start_server
}
Rollback(){
    stop_server
    echo "-----Rollback server----"
    BACK_FILE=$(ls ${INSTALL_DIR}/${APP_NAME} |grep "target-*"|uniq |sort  -rn|head  -n 1)
    #read -p "Please enter the date you want to return (2019-04-17-22-23-08):" Rollback_date
    cd ${INSTALL_DIR}/${APP_NAME}
    rm -rf  ${INSTALL_DIR}/${APP_NAME}/target
    mv ${BACK_FILE}/  target/
    start_server
}

case $1 in
    Update)
        Update
    ;;
    Rollback)
        Rollback
    ;;
    Restart)
        stop_server
        start_server
    ;;
    stop)
        stop_server
    ;;
    Check_Server)
        SERVER_CHECK
    ;;
    start)
        start_server
    ;;
    *)
        echo "Please enter parameters(Update|Rollback|Restart|Check_Server|start|stop)"
esac
