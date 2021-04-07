# BSC任务命令行工具

## 配置文件说明

**路径：conf/config.toml**

- bscSignUrl BSC签名的地址，如：http://127.0.0.1:8585
- bscSignUser 透传给BSC签名服务，如没要求可不传
- bscSignPassword 透传给BSC签名服务，如没要求可不传
- transaction
    - fromAddress 交易发起人的地址，如：0x603e01a2897abebad162f00fe86cfd02d5609c39
    - amount 转账金额，如：5000000000000000
    - fee 手续费,BSC签名服务会动态计算
    - toAddrFileName 收款人列表文件，请把文件于bsctaskcli程序放在同一目录，如：bsc_d_usb_abo1.csv。文件格式需要与BSC签名服务D类文件一致
    - contractAddress 合约地址，输入空字符串表示转账BSC

## Log文件

**路径：log/bsc-task.log**

存放程序运行的日志的文件

## 执行结果日志

如果收款人地址比较多，程序执行的时间就会比较长，可以通过生成的结果日志事实查看进度，该文件会直接创建到bsctaskcli程序同一目录

执行结果日志格式：[config.toAddrFileName]-[时间戳].log

## 运行

```bash
# 可使用以下命令直接执行
# 执行前请检查配置文件的各项参数

sh start.sh
```


