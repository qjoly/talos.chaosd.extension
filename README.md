## Talos Chaosd extension

This extension is a [Talos](https://talos.dev) extension that provides a chaosd (chaos-mesh host daemon). This daemon is responsible for running chaos experiments on the host machine (outside of the Kubernetes cluster).

Be aware that this extension is a work in progress and is not yet fully functional (e.g. it does not yet support the authentication through mTLS).

## Usage

Install the extension:

```bash
talosctl upgrade -i upgrade --force -i ghcr.io/qjoly/talos.chaosd.extension/installer:v1.9.4-amd64
```

Get the logs of the chaosd daemon:
```bash
talosctl logs ext-chaosd
```

Expected output:
```bash
talosctl  -n 192.168.1.164 -e 192.168.1.6 --talosconfig talosconfig logs ext-chaosd -f
192.168.1.164: Chaosd Server Version: version.Info{GitVersion:"v0.0.0-master+$Format:%h$", GitCommit:"$Format:%H$", BuildDate:"2025-03-04T07:03:34Z", GoVersion:"go1.20.14", Compiler:"gc", Platform:"linux/amd64"}
192.168.1.164: [2025/03/04 07:09:12.649 +00:00] [INFO] [cron.go:183] ["Starting Scheduler"]
192.168.1.164: [2025/03/04 07:09:12.650 +00:00] [INFO] [server.go:71] ["starting HTTP server"] [address=0.0.0.0:31767]
192.168.1.164: [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
192.168.1.164:
192.168.1.164: [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
192.168.1.164:  - using env:	export GIN_MODE=release
192.168.1.164:  - using code:	gin.SetMode(gin.ReleaseMode)
192.168.1.164:
192.168.1.164: [GIN-debug] GET    /api/system/health        --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).healthcheck-fm (4 handlers)
192.168.1.164: [GIN-debug] GET    /api/system/version       --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).version-fm (4 handlers)
192.168.1.164: [GIN-debug] GET    /api/swagger/*any         --> github.com/chaos-mesh/chaosd/pkg/swaggerserver.Handler.func1 (4 handlers)
192.168.1.164: [GIN-debug] POST   /api/attack/process       --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).createProcessAttack-fm (4 handlers)
192.168.1.164: [GIN-debug] POST   /api/attack/stress        --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).createStressAttack-fm (4 handlers)
192.168.1.164: [GIN-debug] POST   /api/attack/network       --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).createNetworkAttack-fm (4 handlers)
192.168.1.164: [GIN-debug] POST   /api/attack/disk          --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).createDiskAttack-fm (4 handlers)
192.168.1.164: [GIN-debug] POST   /api/attack/clock         --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).createClockAttack-fm (4 handlers)
192.168.1.164: [GIN-debug] POST   /api/attack/jvm           --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).createJVMAttack-fm (4 handlers)
192.168.1.164: [GIN-debug] POST   /api/attack/kafka         --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).createKafkaAttack-fm (4 handlers)
192.168.1.164: [GIN-debug] POST   /api/attack/vm            --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).createVMAttack-fm (4 handlers)
192.168.1.164: [GIN-debug] POST   /api/attack/redis         --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).createRedisAttack-fm (4 handlers)
192.168.1.164: [GIN-debug] POST   /api/attack/user_defined  --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).createUserDefinedAttack-fm (4 handlers)
192.168.1.164: [GIN-debug] DELETE /api/attack/:uid          --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).recoverAttack-fm (4 handlers)
192.168.1.164: [GIN-debug] GET    /api/experiments/         --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).listExperiments-fm (4 handlers)
192.168.1.164: [GIN-debug] GET    /api/experiments/:uid/runs --> github.com/chaos-mesh/chaosd/pkg/server/httpserver.(*HttpServer).listExperimentRuns-fm (4 handlers)
192.168.1.164: [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
192.168.1.164: Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
192.168.1.164: [GIN-debug] Listening and serving HTTP on 0.0.0.0:31767
```

Using the chaos-mesh stack, you can now create chaos experiments that will be executed on the host machine.