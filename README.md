## 一. all-in-one-bot

__Telegram机器人, 目前支持监控加密货币价格, ChatGPT, 自动抠图, Youtube视频/音频下载和剪切, Telegram贴纸Sticker下载, Telegram的gif图片下载, Bilibili视频下载, Douyin视频下载, 土狗币查询, 通用工具箱(base64,json格式化,时间戳转换)__

## 二. Usage

`安装`

```
bash -c "$(curl -L https://raw.githubusercontent.com/whereAreTheGuGuda/all-in-one-bot/master/install.sh)" @ install
```

> 注: 配置文件token必须添加,否则会启动失败, chatId不添加的情况下只能执行 /chatid 命令获取chatid, 获取到后添加到配置文件并重启服务(可以通过`其他`脚本输出`8` -> `2`进行添加)
* 在 all-in-one-bot.yml 添加你的 telegram token
`telegram 搜索用户 @BotFather 发送 /newbot 获取`
* 在 all-in-one-bot.yml 添加你的 telegram chatId
`添加token后启动应用,去你的bot发送 /chatid 即可获取`

`更新`

```
bash -c "$(curl -L https://raw.githubusercontent.com/whereAreTheGuGuda/all-in-one-bot/master/install.sh)" @ update
```

`卸载`

```
bash -c "$(curl -L https://raw.githubusercontent.com/whereAreTheGuGuda/all-in-one-bot/master/install.sh)" @ uninstall
```

`其他`

```
bash -c "$(curl -L https://raw.githubusercontent.com/whereAreTheGuGuda/all-in-one-bot/master/install.sh)" @
```

`操作`

```
// 启动
systemctl start aio
// 关闭
systemctl stop aio
// 自动启动
systemctl enable aio
// 状态
systemctl status aio
```

## 三. Use Case

* 功能太多太杂导致我自己使用都会有些混乱, 整理一些常用的组合用法
### 1. 如何找聪明钱包

__首先找到一个金狗(涨幅巨大),使用smart_addr_finder去找出早期购买且收益高的钱包地址, 我这边随便找一个作为示例__

`发送命令`

```
/smart_addr_finder
```

`发送参数`

* 50 1即最早的1-50号交易, 50 2 就是51-100号交易,以此类推

```
0x2890df158d76e584877a1d17a85fea3aeeb85aa6 50 10
```

__经过一分钟左右等待, 会自动分析出买卖利润>0且非空投的地址, 然后选择其中收益比较高的地址,然后使用wallet_tx_analyze分析他近期交易收益__

`发送命令`

```
/wallet_tx_analyze
```

`发送参数`

* 30为近期30条, 可以省略不写默认为30, 也可以根据情况扩大分析范围

```
0x1b63e884871aff9a6a55fdd30dbcb82d647d5f99 30
```

__观察钱包的交易总收益和情况,如果返回交易数特别少那大概率是狗庄的老鼠仓,参考价值不高可以省略, 如果交易数高为正常钱包的可能性更高,这个需要自行判断. 然后当判断改地址为聪明钱包时,使用wallet_tracking监控该地址之后的买入卖出操作__

`发送命令`

```
/wallet_tracking
```

### 2. 跟着聪明钱包买入后寻找出点

__聪明钱包不一定永远聪明, 土狗格局的下场就是深埋, 所以有时候聪明钱包在等多倍的时候, 我们应该选择翻倍出场或者是翻倍出本, 这样的操作可能会导致少赚但是更加安全. 所以买入后使用 add_meme_growth_monitor / add_meme_decline_monitor 进行价格监控, 自己选择出点而不是无脑相信聪明钱包__

`监控上涨命令`

```
/add_meme_growth_monitor
```

`监控下跌命令`

```
/add_meme_growth_monitor
```

`发送参数`

* 参数 eth/bsc 是该币的链, 2 为价格(usdt)

```
0x51187cab377ed5e1386042919a9c3d6b5ea402f0 eth 2.4
```

## 四. Command List
### 1. 加密货币监控功能清单
- [x] tracking_lastest_tx 获取正在追踪的钱包最后一次交易的时间
- [x] analyze_addr_token_profit 分析钱包的指定加密货币总收益情况(钱包地址 加加密货币合约地址) 例:0x1c8075cfc18cd17f5fb7743fba811603b819234c 0x808a57ef754c18e1d2cea5d6cf30f00eeeaa1273
- [x] smart_addr_finder 分析高涨幅度币的地址收益来寻找聪明地址 例:  0x2890df158d76e584877a1d17a85fea3aeeb85aa6 50 1
- [x] list_wallet_tracking 列出正在追踪的聪明钱包地址
- [x] list_smart_addr_probe 列出正在探测的聪明钱包地址
- [x] dump_tracking_list dump追踪地址列表(建议每次准备重启服务的时候执行一次)
- [x] wallet_tx_analyze 分析钱包近n条交易的利润 例: 0xaA6a1993Ec0BC72dc44B8E18e1DCDeD11A69302E 30
- [x] wallet_tracking 追踪聪明钱包买卖动态 例: 0x7431931094e8BAe1ECAA7D0b57d2284e121F760e
- [x] stop_wallet_tracking 停止追踪聪明钱包买卖动态 例: 0x7431931094e8BAe1ECAA7D0b57d2284e121F760e
- [x] set_smart_addr_probe_itv 修改聪明地址探测频率 例: 15
- [x] dump_smart_addr_probe_list dump聪明地址的过滤合约(建议每次准备重启服务的时候执行一次)
- [x] smart_addr_tx 输入聪明地址(eth)和近n条交易 例: 0x6b75d8AF000000e20B7a7DDf000Ba900b4009A80 20
- [x] smart_addr_probe 监控聪明地址(eth)购买情况 例:  0x6b75d8AF000000e20B7a7DDf000Ba900b4009A80
- [x] delete_smart_addr_probe 输入关闭监控的聪明地址(eth) 例: 0x6b75d8AF000000e20B7a7DDf000Ba900b4009A80
- [x] add_kline_strategy_probe 探测连续3根一直走势的k线 例: btcusdt
- [x] delete_kline_strategy_probe 删除探测 例: btcusdt
- [x] get_meme 获取meme币信息 例: 0x6982508145454ce325ddbe47a25d4ec3d2311933 bsc(可选填)
- [x] add_meme_growth_monitor 添加加meme币高线监控 例: 0x6982508145454ce325ddbe47a25d4ec3d2311933 0.00000123 (单位USD)
- [x] add_meme_decline_monitor 添加加meme币低线监控 例: 0x6982508145454ce325ddbe47a25d4ec3d2311933 0.0000012 (单位USD)
- [x] meme_monitor_list 列出当前探测的meme币
- [x] delete_meme_monitor 删除meme币监控 例子: 0x6982508145454ce325ddbe47a25d4ec3d2311933 eth
- [x] list_kline_strategy_probe 列出当前探测的加密货币
- [x] add_crypto_growth_monitor 加密货币 提示价格 例: BNB 1110
- [x] add_crypto_decline_monitor 加密货币 提示价格 例: BNB 1110
- [x] get_crypto_price 加密货币[可选]
- [x] delete_crypto_minitor 加密货币(多个用逗号隔开) 例子: BNB,ARB
- [x] get_crypto_ufutures_price u本位合约[可选,默认BTCUSDT] 例子: ETHBTC

### 2. ChatGPT功能清单
- [x] chatgpt

### 3. VPS库存监控功能清单(已弃用)
- [ ] vps_monitor_supported_list 查看支持监控的网站
- [ ] vps_add_supported_list 添加支持监控的网站 例: url keyword name desc(有空格需要引号)
- [ ] add_vps_monitor url(必须是vps_monitor_supported_list有的,或者系统站点模版的商家)

### 4. 抠图功能
- [x] cutout (需要在配置文件添加apikey)

### 5. Telegram 信息获取
- [x] chatid

### 6. Cron 定时提醒
- [x] add_cron 每隔多久一次提醒,单位/秒 例: 15 提醒内容(必填)
- [x] delete_cron 删除 例: 1

### 7. 视频下载
- [x] youtube_download 下载ytb视频
- [x] youtube_audio_download 下载ytb音频
- [x] bilibili_download 下载bilibili视频
- [x] youtube_download_cut 下载ytb的视频并裁剪(需要安装ffmpeg)
- [x] youtube_audio_download_cut 下载ytb音频并裁剪(需要安装ffmpeg)
- [ ] twitter_download 下载twitter的视频
- [x] douyin_download 下载douyin的视频

### 8. 贴纸和GIF下载
- [x] sticker_download 下载贴纸表情
- [x] gif_download 下载GIF(非贴纸)

### 9. 工具箱
- [x] base64_encode 进行base64加密
- [x] base64_decode 进行base64解密
- [x] timestamp_convert 时间戳转换为时间"2006-01-02 15:04:05"
- [x] time_convert 时间转换为时间戳"2006-01-02 15:04:05"
- [x] json_format 格式化json

## 五. 环境安装(可选)

* __Telegram 50M上传限制的解决思路__

1. 前往[Guide](https://tdlib.github.io/telegram-bot-api/build.html)根据自己的系统选择参数,根据他提供的命令执行安装 Local Telegram Api
2. 需要先去 https://my.telegram.org ，登录后，点API development tools可以看到你的api-id和api-hash
3. 执行以下命令,用上面的api-id和api-hash替换里面的<arg>
```
telegram-bot-api --api-id=<arg> --api-hash=<arg> --local -l /var/logs/tgserver.log -v 3
```
4. 通过golang执行该命令发送文件
```
curl -v -F chat_id="<chat_id>" -F video="file://<filepath>" -F supports_streaming=true -F caption="<filename>" http://localhost:8081/bot<token>/sendVideo
```

* __用到视频裁剪功能或者GIF下载功能需要安装 FFmpeg__

`Ubuntu或Debian`
```
sudo apt-get update
sudo apt-get install ffmpeg
```

`CentOS或RHEL`

```
sudo yum install epel-release
sudo yum install ffmpeg
```

`Fedora`

```
sudo dnf install ffmpeg
```

`Arch Linux`

```
sudo pacman -S ffmpeg
```

## 六. Telegram Commands

__通过 @BotFather /setcommands 发送添加__

* 由于功能不断添加 command列表过长命令难找,采用分组形式自行查询获取,建议只填加以下常用命令到command列表,有需要其他功能进行查询获取

`常用命令`
```
get_meme - 获取meme币信息 例: 0x6982508145454ce325ddbe47a25d4ec3d2311933 eth(可选填)
wallet_tx_analyze - 分析钱包近n条交易的利润 例: 0xaA6a1993Ec0BC72dc44B8E18e1DCDeD11A69302E 30
wallet_tracking - 追踪聪明钱包买卖动态 例: 0xaA6a1993Ec0BC72dc44B8E18e1DCDeD11A69302E
stop_wallet_tracking - 停止追踪聪明钱包买卖动态 例: 0xaA6a1993Ec0BC72dc44B8E18e1DCDeD11A69302E
smart_addr_finder - 分析高涨幅度币的地址收益来寻找聪明地址 例:  0x2890df158d76e584877a1d17a85fea3aeeb85aa6 50 1
analyze_addr_token_profit - 分析钱包的指定加密货币总收益情况(钱包地址 加加密货币合约地址) 例:0x1c8075cfc18cd17f5fb7743fba811603b819234c 0x808a57ef754c18e1d2cea5d6cf30f00eeeaa1273
chatgpt - chatgpt功能
youtube_audio_download_cut - 下载ytb音频并裁剪
cmd_list - 列出全部功能
crypto_cmd_list - 加密货币相关功能列表
video_cmd_list - 音视频下载处理功能列表
image_cmd_list - 图片处理/下载功能列表
utils_cmd_list - 工具类功能列表
list_cmd_list - 功能分类列表
```

`全部命令`

```
chatid - 查询chatid
tracking_lastest_tx - 获取正在追踪的钱包最后一次交易的时间
analyze_addr_token_profit - 分析钱包的指定加密货币总收益情况(钱包地址 加加密货币合约地址) 例:0x1c8075cfc18cd17f5fb7743fba811603b819234c 0x808a57ef754c18e1d2cea5d6cf30f00eeeaa1273
smart_addr_finder - 分析高涨幅度币的地址收益来寻找聪明地址 例:  0x2890df158d76e584877a1d17a85fea3aeeb85aa6 50 1
list_wallet_tracking - 列出正在追踪的聪明钱包地址
list_smart_addr_probe - 列出正在探测的聪明钱包地址
dump_tracking_list - dump追踪地址列表(建议每次准备重启服务的时候执行一次)
wallet_tx_analyze - 分析钱包近n条交易的利润 例: 0xaA6a1993Ec0BC72dc44B8E18e1DCDeD11A69302E 30
wallet_tracking - 追踪聪明钱包买卖动态 例: 0xaA6a1993Ec0BC72dc44B8E18e1DCDeD11A69302E
stop_wallet_tracking - 停止追踪聪明钱包买卖动态 例: 0xaA6a1993Ec0BC72dc44B8E18e1DCDeD11A69302E
set_smart_addr_probe_itv - 修改聪明地址探测频率 例: 15
smart_addr_tx - 输入聪明地址(eth)和近n条交易 例: 0x6b75d8AF000000e20B7a7DDf000Ba900b4009A80 50
dump_smart_addr_probe_list - dump聪明地址的过滤合约(建议每次准备重启服务的时候执行一次)
smart_addr_probe - 监控聪明地址(eth)购买情况 例:  0x6b75d8AF000000e20B7a7DDf000Ba900b4009A80
delete_smart_addr_probe - 输入关闭监控的聪明地址(eth) 例: 0x6b75d8AF000000e20B7a7DDf000Ba900b4009A80
add_kline_strategy_probe - 探测连续3根一直走势的k线 例: btcusdt
delete_kline_strategy_probe - 删除探测 例: btcusdt
get_meme - 获取meme币信息 例: 0x6982508145454ce325ddbe47a25d4ec3d2311933 eth(可选填)
add_meme_growth_monitor - 添加加meme币高线监控 例: 0x6982508145454ce325ddbe47a25d4ec3d2311933 eth 0.00000123 (单位USD)
add_meme_decline_monitor - 添加加meme币低线监控 例: 0x6982508145454ce325ddbe47a25d4ec3d2311933 bsc 0.0000012 (单位USD)
meme_monitor_list - 列出当前探测的meme币
delete_meme_monitor - 删除meme币监控 例: 0x6982508145454ce325ddbe47a25d4ec3d2311933 eth
list_kline_strategy_probe - 列出当前探测的加密货币
add_crypto_growth_monitor - 添加加密货币高线监控 例: BNB 1.1 (单位USD)
add_crypto_decline_monitor - 添加加密货币低线监控 例: BNB 1.1 (单位USD)
get_crypto_price - 查询当前价格(默认查询监控的加密货币) 例 : BNB
delete_crypto_minitor - 删除监控的加密货币 例: BNB,ARB
get_crypto_ufutures_price - 查询当前合约价格 例 : ETHUSDT
add_cron - 每隔多久一次提醒,单位/秒 例: 15 提醒内容(必填)
delete_cron - 删除 例: 1
chatgpt - chatgpt功能
cutout - 抠图功能
base64_encode - 进行base64加密
base64_decode - 进行base64解密
timestamp_convert - 时间戳转换为时间"2006-01-02 15:04:05"
time_convert - 时间转换为时间戳"2006-01-02 15:04:05"
json_format - 格式化json
youtube_download - 下载youtube的视频
youtube_audio_download - 下载ytb音频
youtube_download_cut - 下载youtube的视频并裁剪
youtube_audio_download_cut - 下载ytb音频并裁剪
bilibili_download - 下载bilibili的视频
douyin_download - 下载douyin的视频
sticker_download - 下载贴纸表情
gif_download - 下载GIF(非贴纸)
cmd_list - 列出全部功能
crypto_cmd_list - 加密货币相关功能列表
video_cmd_list - 音视频下载处理功能列表
image_cmd_list - 图片处理/下载功能列表
utils_cmd_list - 工具类功能列表
```

__待实现__

```
twitter_download - 下载twitter的视频

```

__弃用__

```
vps_monitor_supported_list - 查看支持监控的网站
add_vps_monitor - 添加VPS库存监控 例: URL(vps_monitor_supported_list里的)
vps_add_supported_list - 添加支持监控的网站 例: url keyword name desc(有空格需要引号)
```