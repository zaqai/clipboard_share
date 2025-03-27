# clipboard_share

A self-hosted clipboard sharing tool across your devices

## Usage

Download the latest pre-built binaries from the [Releases Page](https://github.com/zaqai/clipboard_share/releases).
run the executable file

### Docker Deployment
<https://hub.docker.com/r/zhou29/clipboardshare>
`docker run --name clipboardshare --restart=unless-stopped -d -p 9090:9090 -v /root/data/docker_data/clipboardshare:/app/nutsdb zhou29/clipboardshare`

## Browser-based UI

index: <http://127.0.0.1:9090/idxq>

push: <http://127.0.0.1:9090/postq>
pull: <http://127.0.0.1:9090/your_key>

## API Integration (Windows Notification)

You can set a shortcut keys using [AutoHotkey](https://www.autohotkey.com) to run the following python scripts push or pull clipboard content.

### push

```python
import requests
import pyperclip
import json
from win10toast import ToastNotifier

content = pyperclip.paste()
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
ToastNotifier().show_toast(title="已推送",msg=content,duration=3,threaded=True)
```

### pull

```python
import requests
import pyperclip
from win10toast import ToastNotifier

headers = {
    'User-Agent': 'Apipost client Runtime/+https://www.apipost.cn/',
}

response = requests.get('http://127.0.0.1:9090/your_key',  headers=headers)
data=response.text
pyperclip.copy(data)
ToastNotifier().show_toast(title = "已复制", msg = data, duration = 3, threaded = False)
```

## more features

> Highly customizable - may require technical expertise to adapt for different use cases

When sending messages, the tool can forward content to other platforms like ntfy and WeChat through HTTP requests. To enable this feature:

- Use startup parameters (recommended):
  clipboardshare -port 29090 -ntfyAddr  <https://ntfy.sh/your_topic> -WXAddr <http://192.168.1.55:3002/webhook/msg/v2>
- Environment variables:
  PORT NTFYADDR WXADDR

For more information about WeChat integration, refer to: <https://github.com/danni-cool/wechatbot-webhook>

## Automation Pipeline

If you are interested in github actions, you can refer to <https://github.com/zaqai/clipboard_share/blob/master/.github/workflows/build-go-binary.yml>

when I draft a new release, the workflow will automatically build the cross platform binaries and upload them to the release page. It can also push the latest image(amd64&arm64) to dockerhub.
