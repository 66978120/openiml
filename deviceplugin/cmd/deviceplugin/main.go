package main

import (
	"openi.pcl.ac.cn/openiml/openiml/device-plugin/internal/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"k8s.io/klog/v2"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

func main() {

	options := server.ParseFlags()

	klog.Info("Starting FS watcher.")
	watcher, err := startFSWatcher(pluginapi.DevicePluginPath)
	if err != nil {
		klog.Infof("Failed to created FS watcher. err: %v", err)
		os.Exit(1)
	}
	defer watcher.Close()

	klog.Info("Starting OS watcher.")
	sigs := startOSWatcher(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	var devicePlugin *server.Server

restart:
	if devicePlugin != nil {
		devicePlugin.Stop()
	}
	startErr := make(chan struct{})
	devicePlugin, err = server.NewServer(options)
	if err != nil {
		panic("Failed to create device plugin: " + err.Error())
	}
	if err := devicePlugin.Serve(); err != nil {
		klog.Infof("serve device plugin err: %v, restarting.", err)
		close(startErr)
		goto events
	}

events:
	for {
		select {
		case <-startErr:
			goto restart
		case event := <-watcher.Events:
			if event.Name == pluginapi.KubeletSocket && event.Op&fsnotify.Create == fsnotify.Create {
				klog.Infof("inotify: %s created, restarting.", pluginapi.KubeletSocket)
				goto restart
			}
		case err := <-watcher.Errors:
			klog.Infof("inotify err: %v", err)
		case s := <-sigs:
			switch s {
			case syscall.SIGHUP:
				klog.Info("Received SIGHUP, restarting.")
				goto restart
			default:
				klog.Infof("Received signal %v, shutting down.", s)
				devicePlugin.Stop()
				break events
			}
		}
	}
}

func startFSWatcher(files ...string) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		err = watcher.Add(f)
		if err != nil {
			watcher.Close()
			return nil, err
		}
	}

	return watcher, nil
}

func startOSWatcher(sigs ...os.Signal) chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, sigs...)

	return sigChan
}
