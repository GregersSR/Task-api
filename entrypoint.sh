envvars=()
DIR="/run/secrets"
for file in $(ls ${DIR})
do
    upper_file=$(echo ${file} | tr a-z A-Z)
    envvars+=( ${file}=$(cat ${DIR}/${file}) )
done
env ${envvars[@]} /build/taskapi "$@"