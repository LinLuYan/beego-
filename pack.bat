@echo off
set packName=ezonNew
chcp 65001
del /F /A /Q  %packName%
del /F /A /Q  %packName%.tar.gz
del /F /A /Q  lastupdate.tmp
del /F /A /Q  project.log

bee pack -be GOOS=linux -exr="^.*?\.(txt|bat|sh|log|lock|yaml)$" -exp=".babelrc:.postcssrc.js:index.html:package.json:package-lock.json:models:swagger:vendor:.idea:utils:ble_proto:node_modules:dist:config:src:static_origin:files/logs:files/ota"
echo 打包完成