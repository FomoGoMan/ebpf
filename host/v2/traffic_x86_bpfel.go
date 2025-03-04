// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

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
	FexitTcpSendmsg      *ebpf.ProgramSpec `ebpf:"fexit_tcp_sendmsg"`
	FexitTcpSendpage     *ebpf.ProgramSpec `ebpf:"fexit_tcp_sendpage"`
	FexitUdpSendmsg      *ebpf.ProgramSpec `ebpf:"fexit_udp_sendmsg"`
	FexitUdpv6Sendmsg    *ebpf.ProgramSpec `ebpf:"fexit_udpv6_sendmsg"`
	KprobeIp6Output      *ebpf.ProgramSpec `ebpf:"kprobe_ip6_output"`
	KprobeIpOutput       *ebpf.ProgramSpec `ebpf:"kprobe_ip_output"`
	KprobeSkbConsumeUdp  *ebpf.ProgramSpec `ebpf:"kprobe_skb_consume_udp"`
	KprobeTcpCleanupRbuf *ebpf.ProgramSpec `ebpf:"kprobe_tcp_cleanup_rbuf"`
}

// trafficMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type trafficMapSpecs struct {
	NetworkFlowMap *ebpf.MapSpec `ebpf:"network_flow_map"`
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
}

func (m *trafficMaps) Close() error {
	return _TrafficClose(
		m.NetworkFlowMap,
	)
}

// trafficPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadTrafficObjects or ebpf.CollectionSpec.LoadAndAssign.
type trafficPrograms struct {
	FexitTcpSendmsg      *ebpf.Program `ebpf:"fexit_tcp_sendmsg"`
	FexitTcpSendpage     *ebpf.Program `ebpf:"fexit_tcp_sendpage"`
	FexitUdpSendmsg      *ebpf.Program `ebpf:"fexit_udp_sendmsg"`
	FexitUdpv6Sendmsg    *ebpf.Program `ebpf:"fexit_udpv6_sendmsg"`
	KprobeIp6Output      *ebpf.Program `ebpf:"kprobe_ip6_output"`
	KprobeIpOutput       *ebpf.Program `ebpf:"kprobe_ip_output"`
	KprobeSkbConsumeUdp  *ebpf.Program `ebpf:"kprobe_skb_consume_udp"`
	KprobeTcpCleanupRbuf *ebpf.Program `ebpf:"kprobe_tcp_cleanup_rbuf"`
}

func (p *trafficPrograms) Close() error {
	return _TrafficClose(
		p.FexitTcpSendmsg,
		p.FexitTcpSendpage,
		p.FexitUdpSendmsg,
		p.FexitUdpv6Sendmsg,
		p.KprobeIp6Output,
		p.KprobeIpOutput,
		p.KprobeSkbConsumeUdp,
		p.KprobeTcpCleanupRbuf,
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
//go:embed traffic_x86_bpfel.o
var _TrafficBytes []byte
