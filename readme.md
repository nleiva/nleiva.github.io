## Setup

### Docker

- Build the image

```bash
docker build -t webimage -f ./redirect/Dockerfile .
```

- Run it

```bash
docker run -t -p 80:80 -p 443:443 --rm --name website webimage
```

### TLS

Runs in the host with some delay at the begining

```bash
nicolasleiva at go-instance in ~/go/src/github.com/nleiva/nleiva.github.io on master [!?]
$ sudo go run redirect/main.go
2018/06/04 22:17:15 http: TLS handshake error from 173.38.117.85:61279: context deadline exceeded
2018/06/04 22:17:18 http: TLS handshake error from 173.38.117.85:65228: EOF
2018/06/04 22:17:21 http: TLS handshake error from 173.38.117.85:61297: acme/autocert: missing certificate
..
2018/06/04 22:18:02 http: TLS handshake error from 173.38.117.85:61312: EOF
2018/06/04 22:18:02 http: TLS handshake error from 173.38.117.85:40761: acme/autocert: missing certificate
2018/06/04 22:18:02 http: TLS handshake error from 173.38.117.85:61314: acme/autocert: missing certificate
2018/06/04 22:18:06 http: TLS handshake error from 173.38.117.85:56914: acme/autocert: missing certificate
2018/06/04 22:18:06 http: TLS handshake error from 173.38.117.85:62687: acme/autocert: missing certificate
2018/06/04 22:18:08 http: TLS handshake error from 173.38.117.85:61317: acme/autocert: missing certificate
2018/06/04 22:18:08 http: TLS handshake error from 173.38.117.85:61318: acme/autocert: missing certificate
```

Still running into issues with the Alpine Container

```bash
nicolasleiva at go-instance in ~/go/src/github.com/nleiva/nleiva.github.io on master [!?]
$ docker run -t -p 80:80 -p 443:443 --rm --name website webimage
2018/06/04 22:22:45 http: TLS handshake error from 23.27.152.202:19669: tls: client offered an unsupported, maximum protocol version of 301
2018/06/04 22:23:00 http: TLS handshake error from 173.38.117.85:61342: EOF
2018/06/04 22:26:55 http: TLS handshake error from 173.38.117.85:61336: context deadline exceeded
2018/06/04 22:26:55 http: TLS handshake error from 173.38.117.85:61366: acme/autocert: missing certificate
2018/06/04 22:26:55 http: TLS handshake error from 173.38.117.85:61341: acme/autocert: missing certificate
2018/06/04 22:26:55 http: TLS handshake error from 173.38.117.85:61348: acme/autocert: missing certificate
2018/06/04 22:26:55 http: TLS handshake error from 173.38.117.85:61350: acme/autocert: missing certificate
2018/06/04 22:26:55 http: TLS handshake error from 173.38.117.85:61359: acme/autocert: missing certificate
2018/06/04 22:26:55 http: TLS handshake error from 173.38.117.85:61372: acme/autocert: missing certificate
...
2018/06/04 22:29:19 http: TLS handshake error from 173.38.117.85:61402: acme/autocert: missing certificate
2018/06/04 22:29:21 http: TLS handshake error from 173.38.117.85:61403: acme/autocert: missing certificate
2018/06/04 22:29:21 http: TLS handshake error from 173.38.117.85:61404: acme/autocert: missing certificate
```

### Certbot

```bash
lsb_release -a
sudo apt-get install software-properties-common
sudo add-apt-repository ppa:certbot/certbot
sudo apt-get update
sudo apt-get install certbot
sudo certbot certonly
```

```bash
nicolasleiva at go-instance in ~/go/src/github.com/nleiva/nleiva.github.io on master [!?]
$ sudo certbot certonly
Saving debug log to /var/log/letsencrypt/letsencrypt.log

How would you like to authenticate with the ACME CA?
-------------------------------------------------------------------------------
1: Spin up a temporary webserver (standalone)
2: Place files in webroot directory (webroot)
-------------------------------------------------------------------------------
Select the appropriate number [1-2] then [enter] (press 'c' to cancel): 2
Plugins selected: Authenticator webroot, Installer None
Starting new HTTPS connection (1): acme-v01.api.letsencrypt.org
Please enter in your domain name(s) (comma and/or space separated)  (Enter 'c'
to cancel): www.nleiva.com, nleiva.com
Obtaining a new certificate
Performing the following challenges:
http-01 challenge for www.nleiva.com
http-01 challenge for nleiva.com
Input the webroot for www.nleiva.com: (Enter 'c' to cancel): /home/nicolasleiva/go/src/github.com/nleiva/nleiva.github.io

Select the webroot for nleiva.com:
-------------------------------------------------------------------------------
1: Enter a new webroot
2: /home/nicolasleiva/go/src/github.com/nleiva/nleiva.github.io
-------------------------------------------------------------------------------
Select the appropriate number [1-2] then [enter] (press 'c' to cancel): 2
Waiting for verification...
Cleaning up challenges

IMPORTANT NOTES:
 - Congratulations! Your certificate and chain have been saved at:
   /etc/letsencrypt/live/www.nleiva.com/fullchain.pem
   Your key file has been saved at:
   /etc/letsencrypt/live/www.nleiva.com/privkey.pem
   Your cert will expire on 2018-09-02. To obtain a new or tweaked
   version of this certificate in the future, simply run certbot
   again. To non-interactively renew *all* of your certificates, run
   "certbot renew"
 - If you like Certbot, please consider supporting our work by:

   Donating to ISRG / Let's Encrypt:   https://letsencrypt.org/donate
   Donating to EFF:                    https://eff.org/donate-le
```

### Resources

- [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert)
- [Building Docker Images for Static Go Binaries](https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07)
- [Contributors App](https://github.com/kelseyhightower/contributors)
- [Install Certificates in Alpine Image to establish Secured Communication (SSL/TLS)](https://hackernoon.com/alpine-docker-image-with-secured-communication-ssl-tls-go-restful-api-128eb6b54f1f)
- [certbot](https://certbot.eff.org/lets-encrypt/ubuntutrusty-other)