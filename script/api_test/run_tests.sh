#
# 此脚本将会通过环境变量，设置 API 的基准 URL，并启动测试
# ./run_tests.sh 启动所有测试
# ./run_tests.sh <filename.py> 运行指定的测试
#

export API_BASE="http://devhost1.io:8080"

if [ ! -n "$1" ] ;then
    green *.py
else
    green $1
fi