module openi.pcl.ac.cn/openiml/openiml/device-plugin

go 1.18

require (
	github.com/fsnotify/fsnotify v1.6.0
	github.com/jessevdk/go-flags v1.4.0
	google.golang.org/grpc v1.57.2
	k8s.io/klog/v2 v2.2.0
	k8s.io/kubelet v0.19.0
	openi.pcl.ac.cn/openiml/openiml/common v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-logr/logr v0.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	go-eflib v1.4.10 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	huawei.com/npu-exporter/v6 v6.0.0-RC2.b002.fix // indirect
	k8s.io/apimachinery v0.19.0 // indirect
)

replace (
	go-eflib => ../thirdparty/go-eflib
	huawei.com/npu-exporter/v6 => gitee.com/lh120407/ascend-npu-exporter/v6 v6.0.0-RC2.b002.fix
	openi.pcl.ac.cn/openiml/openiml/common => ../common
)
