# clipboard_share
A self-hosted clipboard share tool across your devices
## use bowser UI
push: http://127.0.0.1:9090/idxq

pull: http://127.0.0.1:9090/your_key
## use api
### push
```python
import requests
import json

headers = {
    'Content-Type': 'application/json',
}
data = {
    "key": "your_key",
    "type": "your_text",
    "value": "your_content"
}
data = json.dumps(data)

response = requests.post('http://127.0.0.1:9090/postq',headers=headers, data=data)
```
### pull
```python
requests.get('http://127.0.0.1:9090/your_key').text
```

## additional features
> Highly customizable, maybe only suitable for me

When pushing messages, it can push msg to other platforms like ntfy, weixin through http request.
All you need to do is specify the url.
- use startup parameter(recommended)
  clipboardshare -port 29090 -ntfyAddr  https://ntfy.sh/your_topic -WXAddr http://192.168.1.55:3002/webhook/msg/v2
- env parameter
  PORT NTFYADDR WXADDR

for more information about weixin push, you can refer https://github.com/danni-cool/wechatbot-webhook