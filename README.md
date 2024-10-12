# SSLTool

一款SSL证书检测工具，支持在线/离线检测SSL证书的有效期。

## 使用方法

### 在线检测

```shell
./ssl-tool -i www.example.com:443
```

### 离线检测

```shell
./ssl-tool -f /path/to/cert.pem
```

## 额外说明

- 本地检测时，证书文件的格式必须是PEM格式或CRT格式。
- 本工具仅支持检测X.509证书。
- 本工具的在线检测实际上可遍历证书链，有需要者可以自行查看源码。