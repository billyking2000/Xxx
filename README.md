## 免责声明

本项目只是本人个人学习开发，本人不保证任何可用性，也不对使用本软件造成的任何后果负责。

## 软件安装

### 一键安装

```
wget -N https://raw.githubusercontent.com/billyking2000/Xxx/master/install.sh && bash install.sh
```

### 手动安装

```
git clone https://github.com/billyking2000/Xxx
cd Xxx
go mod tidy
go build -o XrayR -ldflags "-s -w"
./XrayR --config config.yml
```



