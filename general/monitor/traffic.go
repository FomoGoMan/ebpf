package monitor

import (
	"ebpf_collector/types"
	"fmt"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target amd64 traffic ebpf.c
//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target arm64 traffic ebpf.c
//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target arm traffic ebpf.c

type Collector struct {
	objs  trafficObjects
	links map[string]link.Link
}

func NewCollector() (*Collector, error) {
	c := &Collector{
		links: make(map[string]link.Link),
	}
	if err := c.load(); err != nil {
		c.Close()
		return nil, err
	}
	return c, nil
}
func (c *Collector) load() error {
	if err := loadTrafficObjects(&c.objs, nil); err != nil {
		return fmt.Errorf("loading objects: %v", err)
	}

	rootCgroup := "/sys/fs/cgroup"

	// link Ingress
	if l, err := link.AttachCgroup(link.CgroupOptions{
		Path:    rootCgroup,
		Attach:  ebpf.AttachCGroupInetIngress,
		Program: c.objs.CgroupIngress,
	}); err != nil {
		return fmt.Errorf("attach ingress: %v", err)
	} else {
		c.links["cgroup_skb/ingress"] = l
	}

	// link Egress
	if l, err := link.AttachCgroup(link.CgroupOptions{
		Path:    rootCgroup,
		Attach:  ebpf.AttachCGroupInetEgress,
		Program: c.objs.CgroupEgress,
	}); err != nil {
		return fmt.Errorf("attach egress: %v", err)
	} else {
		c.links["cgroup_skb/egress"] = l
	}

	return nil
}

func (c *Collector) Close() {
	for _, link := range c.links {
		link.Close()
	}
	c.objs.Close()
}

func (c *Collector) Collect() (ingress, egress types.FlowCgroup) {
	ingress = make(types.FlowCgroup)
	egress = make(types.FlowCgroup)

	var key uint64
	var values []uint64

	iter := c.objs.trafficMaps.IngressStats.Iterate()
	for iter.Next(&key, &values) {
		total := uint64(0)
		for _, v := range values {
			total += v
		}
		ingress[key] = total
	}

	iter = c.objs.trafficMaps.EgressStats.Iterate()
	for iter.Next(&key, &values) {
		total := uint64(0)
		for _, v := range values {
			total += v
		}
		egress[key] = total
	}

	return ingress, egress
}
