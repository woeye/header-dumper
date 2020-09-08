# header-dumper

A super simple little tool for dumping HTTP header variables of an incoming request.

You might want to use the docker image like so:

```
docker run --rm -it -p 8000:8000 woeye/header-dumper
```
Then make a request with your HTTP client, for example:

```
http http://localhost:8000
```
