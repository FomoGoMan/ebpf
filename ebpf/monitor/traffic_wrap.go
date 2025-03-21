package monitor

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	general "github.com/FomoGoMan/container-netprobe/interface"
	helperCg "github.com/FomoGoMan/container-netprobe/pkg/cgroup"
	helper "github.com/FomoGoMan/container-netprobe/pkg/container"
	"github.com/cilium/ebpf/rlimit"
)

var _ general.Collector = (*ContainerEbpfMonitor)(nil)
var _ general.CGroupInfoGetter = (*ContainerEbpfMonitor)(nil)
var _ general.PidInfoGetter = (*ContainerEbpfMonitor)(nil)
var _ general.SuspiciousDetector = (*ContainerEbpfMonitor)(nil)

type ContainerEbpfMonitor struct {
	monitor    *EBPFCollector
	cgroupPath string
	cgroupId   uint64
	pid        int
}

func NewEbpfCollector(containerID string) (*ContainerEbpfMonitor, error) {
	cgroupPath, err := helper.GetContainerInfo(containerID)
	if err != nil {
		return nil, err
	}
	fmt.Printf("CGroup Path: %v\n", cgroupPath)

	cgroupID, err := helper.GetCgroupID(cgroupPath)
	if err != nil {
		log.Printf("fail to get CGroup ID : %v", err)
		return nil, err
	}
	log.Printf("CGroup ID: %d\n", cgroupID)

	pid, err := helper.GetPid(containerID)
	if err != nil {
		return nil, err
	}

	monitor, err := NewCollector()
	if err != nil {
		return nil, err
	}
	return &ContainerEbpfMonitor{
		monitor:    monitor,
		cgroupPath: cgroupPath,
		cgroupId:   cgroupID,
		pid:        pid,
	}, nil
}

func (c *ContainerEbpfMonitor) EnableSuspiciousDetect() (suspicious chan int, err error) {
	pidWhiteList := []int{c.pid}
	suspicious = make(chan int, 1)

	go func() {
		for {
			time.Sleep(5 * time.Second)

			pidsGot, err := helperCg.GetPidOfCgroup(filepath.Join(c.GetCgroupPath(), "cgroup.procs"))
			if err != nil {
				log.Printf("GetPidOfCgroup Error: %v\n", err)
				continue
			}

			// TODO: may be allow of pid parent pid belongs to
			for _, whitePid := range pidWhiteList {
				for _, pid := range pidsGot {
					if pid == whitePid {
						continue
					}
					suspicious <- pid
					return
				}
			}
		}
	}()

	return
}

func (c *ContainerEbpfMonitor) CollectTotal() (in, out uint64) {
	return c.monitor.CollectTotal(c.cgroupId)
}

func (c *ContainerEbpfMonitor) SetUp() error {
	// Remove resource limits for kernels <5.11.
	if err := rlimit.RemoveMemlock(); err != nil {
		return fmt.Errorf("EbpfMonitor Removing memlock: %v", err)
	}
	return c.monitor.load()
}

func (c *ContainerEbpfMonitor) Cleanup() {
	c.monitor.Close()
}

func (c *ContainerEbpfMonitor) GetCgroupPath() string {
	return c.cgroupPath
}

func (c *ContainerEbpfMonitor) GetPid() int {
	return c.pid
}
