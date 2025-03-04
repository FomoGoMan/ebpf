// Code generated by bpf2go; DO NOT EDIT.
//go:build arm

package traffic

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadTraffic returns the embedded CollectionSpec for traffic.
func loadTraffic() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_TrafficBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load traffic: %w", err)
	}

	return spec, err
}

// loadTrafficObjects loads traffic and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*trafficObjects
//	*trafficPrograms
//	*trafficMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadTrafficObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadTraffic()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// trafficSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type trafficSpecs struct {
	trafficProgramSpecs
	trafficMapSpecs
}

// trafficSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type trafficProgramSpecs struct {
	KprobeIp6Output       *ebpf.ProgramSpec `ebpf:"kprobe_ip6_output"`
	KprobeIpOutput        *ebpf.ProgramSpec `ebpf:"kprobe_ip_output"`
	KprobeSkbConsumeUdp   *ebpf.ProgramSpec `ebpf:"kprobe_skb_consume_udp"`
	KprobeTcpCleanupRbuf  *ebpf.ProgramSpec `ebpf:"kprobe_tcp_cleanup_rbuf"`
	KprobeTcpSendmsg      *ebpf.ProgramSpec `ebpf:"kprobe_tcp_sendmsg"`
	KprobeTcpSendpage     *ebpf.ProgramSpec `ebpf:"kprobe_tcp_sendpage"`
	KprobeUdpSendmsg      *ebpf.ProgramSpec `ebpf:"kprobe_udp_sendmsg"`
	KprobeUdpv6Sendmsg    *ebpf.ProgramSpec `ebpf:"kprobe_udpv6_sendmsg"`
	KretprobeTcpSendmsg   *ebpf.ProgramSpec `ebpf:"kretprobe_tcp_sendmsg"`
	KretprobeTcpSendpage  *ebpf.ProgramSpec `ebpf:"kretprobe_tcp_sendpage"`
	KretprobeUdpSendmsg   *ebpf.ProgramSpec `ebpf:"kretprobe_udp_sendmsg"`
	KretprobeUdpv6Sendmsg *ebpf.ProgramSpec `ebpf:"kretprobe_udpv6_sendmsg"`
}

// trafficMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type trafficMapSpecs struct {
	NetworkFlowMap *ebpf.MapSpec `ebpf:"network_flow_map"`
	TcpSendmsgMap  *ebpf.MapSpec `ebpf:"tcp_sendmsg_map"`
	UdpSendmsgMap  *ebpf.MapSpec `ebpf:"udp_sendmsg_map"`
}

// trafficObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadTrafficObjects or ebpf.CollectionSpec.LoadAndAssign.
type trafficObjects struct {
	trafficPrograms
	trafficMaps
}

func (o *trafficObjects) Close() error {
	return _TrafficClose(
		&o.trafficPrograms,
		&o.trafficMaps,
	)
}

// trafficMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadTrafficObjects or ebpf.CollectionSpec.LoadAndAssign.
type trafficMaps struct {
	NetworkFlowMap *ebpf.Map `ebpf:"network_flow_map"`
	TcpSendmsgMap  *ebpf.Map `ebpf:"tcp_sendmsg_map"`
	UdpSendmsgMap  *ebpf.Map `ebpf:"udp_sendmsg_map"`
}

func (m *trafficMaps) Close() error {
	return _TrafficClose(
		m.NetworkFlowMap,
		m.TcpSendmsgMap,
		m.UdpSendmsgMap,
	)
}

// trafficPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadTrafficObjects or ebpf.CollectionSpec.LoadAndAssign.
type trafficPrograms struct {
	KprobeIp6Output       *ebpf.Program `ebpf:"kprobe_ip6_output"`
	KprobeIpOutput        *ebpf.Program `ebpf:"kprobe_ip_output"`
	KprobeSkbConsumeUdp   *ebpf.Program `ebpf:"kprobe_skb_consume_udp"`
	KprobeTcpCleanupRbuf  *ebpf.Program `ebpf:"kprobe_tcp_cleanup_rbuf"`
	KprobeTcpSendmsg      *ebpf.Program `ebpf:"kprobe_tcp_sendmsg"`
	KprobeTcpSendpage     *ebpf.Program `ebpf:"kprobe_tcp_sendpage"`
	KprobeUdpSendmsg      *ebpf.Program `ebpf:"kprobe_udp_sendmsg"`
	KprobeUdpv6Sendmsg    *ebpf.Program `ebpf:"kprobe_udpv6_sendmsg"`
	KretprobeTcpSendmsg   *ebpf.Program `ebpf:"kretprobe_tcp_sendmsg"`
	KretprobeTcpSendpage  *ebpf.Program `ebpf:"kretprobe_tcp_sendpage"`
	KretprobeUdpSendmsg   *ebpf.Program `ebpf:"kretprobe_udp_sendmsg"`
	KretprobeUdpv6Sendmsg *ebpf.Program `ebpf:"kretprobe_udpv6_sendmsg"`
}

func (p *trafficPrograms) Close() error {
	return _TrafficClose(
		p.KprobeIp6Output,
		p.KprobeIpOutput,
		p.KprobeSkbConsumeUdp,
		p.KprobeTcpCleanupRbuf,
		p.KprobeTcpSendmsg,
		p.KprobeTcpSendpage,
		p.KprobeUdpSendmsg,
		p.KprobeUdpv6Sendmsg,
		p.KretprobeTcpSendmsg,
		p.KretprobeTcpSendpage,
		p.KretprobeUdpSendmsg,
		p.KretprobeUdpv6Sendmsg,
	)
}

func _TrafficClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed traffic_arm_bpfel.o
var _TrafficBytes []byte
