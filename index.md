## Setup

### Docker

```bash
docker build -t webimage -f ./redirect/Dockerfile .
docker run -t -p 80:80 --rm --name website webimage
```

### TLS

TODO