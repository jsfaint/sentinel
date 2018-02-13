# WanKeYun Sentinel

This is a client for [WanKeYun](http://www.onethingcloud.com/us/site/index.html) which is inspired by [imt-wanke-client](https://github.com/Immortalt/imt-wanke-client)

Sentinel has no GUI interface, it is configed with json file, and only pushing the message via [servchan](http://sc.ftqq.com)

# Supported Platform

1. x86 Windows/Linux/macOS
2. mips/arm Linux

# Features

Sentinel has three main features:
1. Monitor the device status, refresh data every 30s
2. Report last day's summary at 9:00 of the next day
3. Withdraw the coin at 10:10 of the specific day (Monday ~ Friday)

# Configuration

Please fill the `config.json` before starting sentinel
```json
{
    "token": "",
    "draw_day": 1,
    "accounts": [
        {"phone": "", "pass": ""}
    ]
}
```

Fill the token with servchan token in `token` filed and account info in `accounts` field

# Donation

If you like this utility, please consider donating me via this address `0x8815d91b73e0c8dd0a1ee3d917cbfdb20635a6e8`

Thanks
