# clipboard_share
A self-hosted clipboard share tool across your devices

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
or you can get the msg throw your browser: http://127.0.0.1:9090/your_key