### ok-short
Short link service built on golang

### Build On
* gin
* gorm
* redis
* mysql
* nats
* swagger

[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

### Feature
* Long links -> Short link
* Jump to the original url by visiting the short link
* Query the original link through the short link

### Prepare the environment
1. redis
2. mysql
3. nats

Set up the above environment and modify the configs/config.yaml

### Compilation
```shell script
scripts/build.sh # Use scripts/build.bat on Windows
```
3. Access swagger page: http://ip:port/swagger/index.html

### License
Released under the [MIT License](https://github.com/Joeyscat/ok-short/blob/master/LICENSE)