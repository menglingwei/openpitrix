#!/bin/bash

echo "Deploying k8s resource..."

[ -z `which kubectl` ] && echo "Deployed failed: You need to install 'kubectl' first." && exit 1

# Back to the root of the project
cd $(dirname $0)
cd ../..

DEFAULT_NAMESPACE=openpitrix-system

NAMESPACE=${DEFAULT_NAMESPACE}
VERSION=""
METADATA=0
DBCTRL=0
BASE=0
DASHBOARD=0
STORAGE=0
ALL=0
JOB_REPLICA=1
TASK_REPLICA=1
API_NODEPORT=""
PILOT_NODEPORT=""
DASHBOARD_NODEPORT=""

# use nodePort for api/pilot/dashboard service
# $cat ${NODEPORT_FILE}
# API_NODEPORT=31009
# PILOT_NODEPORT=31019
# DASHBOARD_NODEPORT=31029
NODEPORT_FILE="./config/node_port"
if [ -f ${NODEPORT_FILE} ];then
  API_NODEPORT=`cat ${NODEPORT_FILE} | grep API_NODEPORT | awk -F '=' '{print "nodePort: "$2}'`
  PILOT_NODEPORT=`cat ${NODEPORT_FILE} | grep PILOT_NODEPORT | awk -F '=' '{print "nodePort: "$2}'`
  DASHBOARD_NODEPORT=`cat ${NODEPORT_FILE} | grep DASHBOARD_NODEPORT | awk -F '=' '{print "nodePort: "$2}'`
fi

REQUESTS=100
LIMITS=500

usage() {
  echo "Usage:"
  echo "  deploy-k8s.sh [-n NAMESPACE] [-v VERSION] COMMAND"
  echo "Description:"
  echo "        -n NAMESPACE    : the namespace of kubernetes."
  echo "        -v VERSION      : the version to be deployed."
  echo "        -r REQUESTS     : the requests of container resources."
  echo "        -l LIMITS       : the limits of container resources."
  echo "        -j JOB REPLICA  : the job replica number."
  echo "        -t TASK REPLICA : the task replica number."
  echo "        -b              : base model will be applied."
  echo "        -m              : metadata will be applied."
  echo "        -d              : dbctrl will be applied."
  echo "        -s              : dashboard will be applied."
  echo "        -o              : storage will be applied."
  echo "        -a              : all of base/metadata/dbctrl/dashboard/storage will be applied."
  exit -1
}


while getopts n:v:r:l:j:t:hbdmsoa option
do
  case "${option}"
  in
  n) NAMESPACE=${OPTARG};;
  v) VERSION=${OPTARG};;
  r) REQUESTS=${OPTARG};;
  l) LIMITS=${OPTARG};;
  j) JOB_REPLICA=${OPTARG};;
  t) TASK_REPLICA=${OPTARG};;
  d) DBCTRL=1;;
  m) METADATA=1;;
  b) BASE=1;;
  s) DASHBOARD=1;;
  o) STORAGE=1;;
  a) ALL=1;;
  h) usage ;;
  *) usage ;;
  esac
done

if [ "${METADATA}" == "0" ] && \
   [ "${DBCTRL}" == "0" ] && \
   [ "${BASE}" == "0" ] && \
   [ "${DASHBOARD}" == "0" ] && \
   [ "${STORAGE}" == "0" ] && \
   [ "${ALL}" == "0" ]
then
  usage
fi

IMAGE_PULL_POLICY="IfNotPresent"
DASHBOARD_IMAGE="openpitrix/dashboard"
if [ "${VERSION}" == "" ];then
  VERSION=$(curl -L -s https://api.github.com/repos/openpitrix/openpitrix/releases/latest | grep tag_name | sed "s/ *\"tag_name\": *\"\(.*\)\",*/\1/")
fi

if [ "${VERSION}" == "dev" ];then
  IMAGE="openpitrix/openpitrix-dev:latest"
  METADATA_IMAGE="openpitrix/openpitrix-dev:metadata"
  FLYWAY_IMAGE="openpitrix/openpitrix-dev:flyway"
  IMAGE_PULL_POLICY="Always"
elif [ "$VERSION" == "latest" ];then
  IMAGE="openpitrix/openpitrix:latest"
  METADATA_IMAGE="openpitrix/openpitrix:metadata"
  FLYWAY_IMAGE="openpitrix/openpitrix:flyway"
  IMAGE_PULL_POLICY="Always"
else
  IMAGE="openpitrix/openpitrix:$VERSION"
  METADATA_IMAGE="openpitrix/openpitrix:metadata-$VERSION"
  FLYWAY_IMAGE="openpitrix/openpitrix:flyway-$VERSION"
fi

replace() {
  sed -e "s!\${NAMESPACE}!${NAMESPACE}!g" \
	  -e "s!\${IMAGE}!${IMAGE}!g" \
	  -e "s!\${DASHBOARD_IMAGE}!${DASHBOARD_IMAGE}!g" \
	  -e "s!\${METADATA_IMAGE}!${METADATA_IMAGE}!g" \
	  -e "s!\${FLYWAY_IMAGE}!${FLYWAY_IMAGE}!g" \
	  -e "s!\${REQUESTS}!${REQUESTS}!g" \
	  -e "s!\${LIMITS}!${LIMITS}!g" \
	  -e "s!\${JOB_REPLICA}!${JOB_REPLICA}!g" \
	  -e "s!\${TASK_REPLICA}!${TASK_REPLICA}!g" \
	  -e "s!\${VERSION}!${VERSION}!g" \
	  -e "s!\${IMAGE_PULL_POLICY}!${IMAGE_PULL_POLICY}!g" \
	  -e "s!\${API_NODEPORT}!${API_NODEPORT}!g" \
	  -e "s!\${PILOT_NODEPORT}!${PILOT_NODEPORT}!g" \
	  -e "s!\${DASHBOARD_NODEPORT}!${DASHBOARD_NODEPORT}!g" \
	  $1
}

[ -z `which make` ] && echo "Deployed failed: You need to install 'make' first." && exit 1

kubectl get ns ${NAMESPACE}
if [ $? != 0 ];then
  kubectl create namespace ${NAMESPACE}
fi

if [ "${STORAGE}" == "1" ] || [ "${ALL}" == "1" ];then
  kubectl create secret generic mysql-pass --from-file=./kubernetes/password.txt -n ${NAMESPACE}
  for FILE in `ls ./kubernetes/db/ | grep -v "\-job.yaml$"`;do
    replace ./kubernetes/db/${FILE} | kubectl apply -f -
  done

  for FILE in `ls ./kubernetes/etcd/`;do
    replace ./kubernetes/etcd/${FILE} | kubectl apply -f -
  done

  for FILE in `ls ./kubernetes/minio/`;do
    replace ./kubernetes/minio/${FILE} | kubectl apply -f -
  done
fi
if [ "${DBCTRL}" == "1" ] || [ "${ALL}" == "1" ];then
  for FILE in `ls ./kubernetes/ctrl`;do
    replace ./kubernetes/ctrl/${FILE} | kubectl delete -f - --ignore-not-found=true
    replace ./kubernetes/ctrl/${FILE} | kubectl apply -f -
  done

  for FILE in `ls ./kubernetes/db/ | grep "\-job.yaml$"`;do
    replace ./kubernetes/db/${FILE} | kubectl delete -f - --ignore-not-found=true
    replace ./kubernetes/db/${FILE} | kubectl apply -f -
  done
fi
if [ "${BASE}" == "1" ] || [ "${ALL}" == "1" ];then
  cd ./kubernetes/iam-config
  make
  cd ../..

  kubectl create secret generic iam-secret-key --from-file=./kubernetes/iam-config/secret-key.txt -n ${NAMESPACE}
  for FILE in `ls ./kubernetes/openpitrix/ | grep "^openpitrix-"`;do
    replace ./kubernetes/openpitrix/${FILE} | kubectl delete -f - --ignore-not-found=true
    replace ./kubernetes/openpitrix/${FILE} | kubectl apply -f -
  done
fi
if [ "${METADATA}" == "1" ] || [ "${ALL}" == "1" ];then
  ./kubernetes/scripts/generate-certs.sh -n ${NAMESPACE}
  if [ $? -ne 0 ]; then
	echo "Deploy failed."
	exit 1
  fi

  for FILE in `ls ./kubernetes/openpitrix/metadata/`;do
    replace ./kubernetes/openpitrix/metadata/${FILE} | kubectl delete -f - --ignore-not-found=true
    replace ./kubernetes/openpitrix/metadata/${FILE} | kubectl apply -f -
  done
fi
if [ "${DASHBOARD}" == "1" ] || [ "${ALL}" == "1" ];then
  for FILE in `ls ./kubernetes/openpitrix/dashboard/`;do
    replace ./kubernetes/openpitrix/dashboard/${FILE} | kubectl delete -f - --ignore-not-found=true
    replace ./kubernetes/openpitrix/dashboard/${FILE} | kubectl apply -f -
  done
fi

echo "Deployed successfully"