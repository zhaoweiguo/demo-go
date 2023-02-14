

# 说明
- MacOS 的网络安全策略导致的。在 MacOS 上，默认禁用了组播数据的发送和接收。
- 可以在 Terminal 中输入以下命令以允许组播数据的发送和接收
```
sudo sysctl -w net.inet.ip.mforwarding=1
- 如果要永久开启该功能，请将以下命令写入 /etc/sysctl.conf 文件：
net.inet.ip.mforwarding=1
```
- 上面这个方法在最新版的macOS已经不可行了
```
要执行多播转发，可以尝试其他方法，例如在网络代理设备上进行配置。
```

