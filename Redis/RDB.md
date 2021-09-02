## 通过rdbtools分析RDB文件找出bigkeys top100

### 准备工作
1. 安装rdbtools
```shell script
pip install rdbtools
```
或者手动下载安装
```shell script
wget https://pypi.python.org/packages/source/p/pip/pip-1.5.4.tar.gz
tar -zxvf pip-1.5.4.tar.gz
cd pip-1.5.4
python setup.py install
pip install rdbtools
```

2. 可以再安装一个python-lzf，解析速度比较快（可选）
```shell script
pip install python-lzf
```

### 分析文件
```shell script
rdb -c memory test.rdb -l 100 -f ./dump_memory.csv
```

