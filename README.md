# go-p2p-server
Web P2P 服务器，依托(cdnbye)[httsp://www.cdnbye.com]的源码和(klink.tech)[https://klink.tech]页面而开发的服务器。

# 测试地址
- https://gop2p.cachecha.com/

## 接入
```
var hlsjsConfig = {
    debug: true,
    // Other hlsjsConfig options provided by hls.js
    p2pConfig: {
    	announce: "https://gop2p.cachecha.com",
        wsSignalerAddr: 'wss://gop2p.cachecha.com/ws',
        logLevel: 'debug',
        // Other p2pConfig options if applicable
    }
};
// Hls constructor is overriden by included bundle
var hls = new Hls(hlsjsConfig);
// Use hls just like the usual hls.js ...
hls.loadSource(contentUrl);
hls.attachMedia(video);
hls.on(Hls.Events.MANIFEST_PARSED,function() {
    video.play();
});
```

## 快速安装

```
curl -fsSL  https://raw.githubusercontent.com/midoks/go-p2p-server/master/scripts/install.sh | sh
```

## 快速开发
```
curl -fsSL  https://raw.githubusercontent.com/midoks/go-p2p-server/master/scripts/install_dev.sh | sh
```