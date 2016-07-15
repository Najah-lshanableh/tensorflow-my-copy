// Code generated by protoc-gen-go.
// source: tensorflow/core/protobuf/config.proto
// DO NOT EDIT!

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Optimization level
type OptimizerOptions_Level int32

const (
	// L1 is the default level.
	// Optimization performed at L1 :
	// 1. Common subexpression elimination
	// 2. Constant folding
	OptimizerOptions_L1 OptimizerOptions_Level = 0
	// No optimizations
	OptimizerOptions_L0 OptimizerOptions_Level = -1
)

var OptimizerOptions_Level_name = map[int32]string{
	0:  "L1",
	-1: "L0",
}
var OptimizerOptions_Level_value = map[string]int32{
	"L1": 0,
	"L0": -1,
}

func (x OptimizerOptions_Level) String() string {
	return proto.EnumName(OptimizerOptions_Level_name, int32(x))
}
func (OptimizerOptions_Level) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{1, 0} }

type RunOptions_TraceLevel int32

const (
	RunOptions_NO_TRACE   RunOptions_TraceLevel = 0
	RunOptions_FULL_TRACE RunOptions_TraceLevel = 1
)

var RunOptions_TraceLevel_name = map[int32]string{
	0: "NO_TRACE",
	1: "FULL_TRACE",
}
var RunOptions_TraceLevel_value = map[string]int32{
	"NO_TRACE":   0,
	"FULL_TRACE": 1,
}

func (x RunOptions_TraceLevel) String() string {
	return proto.EnumName(RunOptions_TraceLevel_name, int32(x))
}
func (RunOptions_TraceLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{4, 0} }

type GPUOptions struct {
	// A value between 0 and 1 that indicates what fraction of the
	// available GPU memory to pre-allocate for each process.  1 means
	// to pre-allocate all of the GPU memory, 0.5 means the process
	// allocates ~50% of the available GPU memory.
	PerProcessGpuMemoryFraction float64 `protobuf:"fixed64,1,opt,name=per_process_gpu_memory_fraction,json=perProcessGpuMemoryFraction" json:"per_process_gpu_memory_fraction,omitempty"`
	// The type of GPU allocation strategy to use.
	//
	// Allowed values:
	// "": The empty string (default) uses a system-chosen default
	//     which may change over time.
	//
	// "BFC": A "Best-fit with coalescing" algorithm, simplified from a
	//        version of dlmalloc.
	AllocatorType string `protobuf:"bytes,2,opt,name=allocator_type,json=allocatorType" json:"allocator_type,omitempty"`
	// Delay deletion of up to this many bytes to reduce the number of
	// interactions with gpu driver code.  If 0, the system chooses
	// a reasonable default (several MBs).
	DeferredDeletionBytes int64 `protobuf:"varint,3,opt,name=deferred_deletion_bytes,json=deferredDeletionBytes" json:"deferred_deletion_bytes,omitempty"`
	// If true, the allocator does not pre-allocate the entire specified
	// GPU memory region, instead starting small and growing as needed.
	AllowGrowth bool `protobuf:"varint,4,opt,name=allow_growth,json=allowGrowth" json:"allow_growth,omitempty"`
}

func (m *GPUOptions) Reset()                    { *m = GPUOptions{} }
func (m *GPUOptions) String() string            { return proto.CompactTextString(m) }
func (*GPUOptions) ProtoMessage()               {}
func (*GPUOptions) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

// Options passed to the graph optimizer
type OptimizerOptions struct {
	// If true, optimize the graph using common subexpression elimination.
	DoCommonSubexpressionElimination bool `protobuf:"varint,1,opt,name=do_common_subexpression_elimination,json=doCommonSubexpressionElimination" json:"do_common_subexpression_elimination,omitempty"`
	// If true, perform constant folding optimization on the graph.
	DoConstantFolding bool `protobuf:"varint,2,opt,name=do_constant_folding,json=doConstantFolding" json:"do_constant_folding,omitempty"`
	// If true, perform function inlining on the graph.
	DoFunctionInlining bool                   `protobuf:"varint,4,opt,name=do_function_inlining,json=doFunctionInlining" json:"do_function_inlining,omitempty"`
	OptLevel           OptimizerOptions_Level `protobuf:"varint,3,opt,name=opt_level,json=optLevel,enum=tensorflow.OptimizerOptions_Level" json:"opt_level,omitempty"`
}

func (m *OptimizerOptions) Reset()                    { *m = OptimizerOptions{} }
func (m *OptimizerOptions) String() string            { return proto.CompactTextString(m) }
func (*OptimizerOptions) ProtoMessage()               {}
func (*OptimizerOptions) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

type GraphOptions struct {
	// If true, use control flow to schedule the activation of Recv nodes.
	// (Currently ignored.)
	EnableRecvScheduling bool `protobuf:"varint,2,opt,name=enable_recv_scheduling,json=enableRecvScheduling" json:"enable_recv_scheduling,omitempty"`
	// Options controlling how graph is optimized.
	OptimizerOptions *OptimizerOptions `protobuf:"bytes,3,opt,name=optimizer_options,json=optimizerOptions" json:"optimizer_options,omitempty"`
	// Build a cost model detailing the memory usage and performance of
	// each node of the graph.
	BuildCostModel bool `protobuf:"varint,4,opt,name=build_cost_model,json=buildCostModel" json:"build_cost_model,omitempty"`
}

func (m *GraphOptions) Reset()                    { *m = GraphOptions{} }
func (m *GraphOptions) String() string            { return proto.CompactTextString(m) }
func (*GraphOptions) ProtoMessage()               {}
func (*GraphOptions) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

func (m *GraphOptions) GetOptimizerOptions() *OptimizerOptions {
	if m != nil {
		return m.OptimizerOptions
	}
	return nil
}

// Session configuration parameters.
// The system picks an appropriate values for fields that are not set.
type ConfigProto struct {
	// Map from device type name (e.g., "CPU" or "GPU" ) to maximum
	// number of devices of that type to use.  If a particular device
	// type is not found in the map, the system picks an appropriate
	// number.
	DeviceCount map[string]int32 `protobuf:"bytes,1,rep,name=device_count,json=deviceCount" json:"device_count,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// The execution of an individual op (for some op types) can be
	// parallelized on a pool of intra_op_parallelism_threads.
	// 0 means the system picks an appropriate number.
	IntraOpParallelismThreads int32 `protobuf:"varint,2,opt,name=intra_op_parallelism_threads,json=intraOpParallelismThreads" json:"intra_op_parallelism_threads,omitempty"`
	// Nodes that perform blocking operations are enqueued on a pool of
	// inter_op_parallelism_threads available in each process.
	//
	// 0 means the system picks an appropriate number.
	//
	// Note that the first Session created in the process sets the
	// number of threads for all future sessions unless use_per_session_threads is
	// true.
	InterOpParallelismThreads int32 `protobuf:"varint,5,opt,name=inter_op_parallelism_threads,json=interOpParallelismThreads" json:"inter_op_parallelism_threads,omitempty"`
	// If true, use a new set of threads for this session rather than the global
	// pool of threads. Only supported by direct sessions.
	//
	// If false, use the global threads created by the first session.
	UsePerSessionThreads bool `protobuf:"varint,9,opt,name=use_per_session_threads,json=usePerSessionThreads" json:"use_per_session_threads,omitempty"`
	// Assignment of Nodes to Devices is recomputed every placement_period
	// steps until the system warms up (at which point the recomputation
	// typically slows down automatically).
	PlacementPeriod int32 `protobuf:"varint,3,opt,name=placement_period,json=placementPeriod" json:"placement_period,omitempty"`
	// When any filters are present sessions will ignore all devices which do not
	// match the filters. Each filter can be partially specified, e.g. "/job:ps"
	// "/job:worker/replica:3", etc.
	DeviceFilters []string `protobuf:"bytes,4,rep,name=device_filters,json=deviceFilters" json:"device_filters,omitempty"`
	// Options that apply to all GPUs.
	GpuOptions *GPUOptions `protobuf:"bytes,6,opt,name=gpu_options,json=gpuOptions" json:"gpu_options,omitempty"`
	// Whether soft placement is allowed. If allow_soft_placement is true,
	// an op will be placed on CPU if
	//   1. there's no GPU implementation for the OP
	// or
	//   2. no GPU devices are known or registered
	// or
	//   3. need to co-locate with reftype input(s) which are from CPU.
	AllowSoftPlacement bool `protobuf:"varint,7,opt,name=allow_soft_placement,json=allowSoftPlacement" json:"allow_soft_placement,omitempty"`
	// Whether device placements should be logged.
	LogDevicePlacement bool `protobuf:"varint,8,opt,name=log_device_placement,json=logDevicePlacement" json:"log_device_placement,omitempty"`
	// Options that apply to all graphs.
	GraphOptions *GraphOptions `protobuf:"bytes,10,opt,name=graph_options,json=graphOptions" json:"graph_options,omitempty"`
	// Global timeout for all blocking operations in this session.  If non-zero,
	// and not overridden on a per-operation basis, this value will be used as the
	// deadline for all blocking operations.
	OperationTimeoutInMs int64 `protobuf:"varint,11,opt,name=operation_timeout_in_ms,json=operationTimeoutInMs" json:"operation_timeout_in_ms,omitempty"`
}

func (m *ConfigProto) Reset()                    { *m = ConfigProto{} }
func (m *ConfigProto) String() string            { return proto.CompactTextString(m) }
func (*ConfigProto) ProtoMessage()               {}
func (*ConfigProto) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{3} }

func (m *ConfigProto) GetDeviceCount() map[string]int32 {
	if m != nil {
		return m.DeviceCount
	}
	return nil
}

func (m *ConfigProto) GetGpuOptions() *GPUOptions {
	if m != nil {
		return m.GpuOptions
	}
	return nil
}

func (m *ConfigProto) GetGraphOptions() *GraphOptions {
	if m != nil {
		return m.GraphOptions
	}
	return nil
}

// EXPERIMENTAL. Options for a single Run() call.
type RunOptions struct {
	TraceLevel RunOptions_TraceLevel `protobuf:"varint,1,opt,name=trace_level,json=traceLevel,enum=tensorflow.RunOptions_TraceLevel" json:"trace_level,omitempty"`
	// Time to wait for operation to complete in milliseconds.
	TimeoutInMs int64 `protobuf:"varint,2,opt,name=timeout_in_ms,json=timeoutInMs" json:"timeout_in_ms,omitempty"`
}

func (m *RunOptions) Reset()                    { *m = RunOptions{} }
func (m *RunOptions) String() string            { return proto.CompactTextString(m) }
func (*RunOptions) ProtoMessage()               {}
func (*RunOptions) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{4} }

// EXPERIMENTAL. Metadata output (i.e., non-Tensor) for a single Run() call.
type RunMetadata struct {
	// Statistics traced for this step. Populated if tracing is turned on via the
	// "RunOptions" proto.
	// EXPERIMENTAL: The format and set of events may change in future versions.
	StepStats *StepStats `protobuf:"bytes,1,opt,name=step_stats,json=stepStats" json:"step_stats,omitempty"`
}

func (m *RunMetadata) Reset()                    { *m = RunMetadata{} }
func (m *RunMetadata) String() string            { return proto.CompactTextString(m) }
func (*RunMetadata) ProtoMessage()               {}
func (*RunMetadata) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{5} }

func (m *RunMetadata) GetStepStats() *StepStats {
	if m != nil {
		return m.StepStats
	}
	return nil
}

func init() {
	proto.RegisterType((*GPUOptions)(nil), "tensorflow.GPUOptions")
	proto.RegisterType((*OptimizerOptions)(nil), "tensorflow.OptimizerOptions")
	proto.RegisterType((*GraphOptions)(nil), "tensorflow.GraphOptions")
	proto.RegisterType((*ConfigProto)(nil), "tensorflow.ConfigProto")
	proto.RegisterType((*RunOptions)(nil), "tensorflow.RunOptions")
	proto.RegisterType((*RunMetadata)(nil), "tensorflow.RunMetadata")
	proto.RegisterEnum("tensorflow.OptimizerOptions_Level", OptimizerOptions_Level_name, OptimizerOptions_Level_value)
	proto.RegisterEnum("tensorflow.RunOptions_TraceLevel", RunOptions_TraceLevel_name, RunOptions_TraceLevel_value)
}

var fileDescriptor16 = []byte{
	// 929 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0xe9, 0xa6, 0x24, 0xc7, 0x69, 0x37, 0x6b, 0xba, 0xbb, 0x61, 0x59, 0x89, 0xae, 0x51,
	0xa5, 0xd2, 0x8b, 0x74, 0x29, 0xbf, 0x42, 0x02, 0x44, 0xd3, 0xa6, 0x2a, 0x34, 0x34, 0x9a, 0x64,
	0xaf, 0x47, 0x4e, 0x3c, 0x49, 0xad, 0xda, 0x1e, 0x6b, 0x66, 0xdc, 0x12, 0x9e, 0x86, 0x2b, 0x9e,
	0x81, 0x07, 0xe1, 0x19, 0x78, 0x0d, 0x38, 0x33, 0xe3, 0xbf, 0xad, 0x56, 0xd0, 0x9b, 0x8e, 0xcf,
	0xf7, 0x9d, 0x99, 0xef, 0xfc, 0x06, 0x0e, 0x14, 0x4b, 0x25, 0x17, 0xab, 0x98, 0xdf, 0x1f, 0x2f,
	0xb9, 0x60, 0xc7, 0x99, 0xe0, 0x8a, 0x2f, 0xf2, 0x15, 0x7e, 0xa5, 0xab, 0x68, 0x3d, 0x34, 0xdf,
	0x1e, 0xd4, 0xb4, 0x17, 0x47, 0x0f, 0x5d, 0x56, 0x22, 0x48, 0xd8, 0x3d, 0x17, 0xb7, 0xc7, 0x52,
	0xb1, 0x8c, 0x4a, 0x15, 0x28, 0x69, 0xfd, 0xfc, 0xbf, 0x1c, 0x80, 0x8b, 0xe9, 0x9b, 0xeb, 0x4c,
	0x45, 0x3c, 0x95, 0xde, 0x19, 0x7c, 0x9c, 0x31, 0x41, 0x11, 0x5b, 0x32, 0x29, 0xe9, 0x3a, 0xcb,
	0x69, 0xc2, 0x12, 0x2e, 0x36, 0x14, 0xef, 0x58, 0x6a, 0xce, 0xc0, 0xd9, 0x77, 0x0e, 0x1d, 0xf2,
	0x11, 0xd2, 0xa6, 0x96, 0x75, 0x91, 0xe5, 0x13, 0xc3, 0x19, 0x17, 0x14, 0xef, 0x00, 0x76, 0x83,
	0x38, 0xe6, 0xcb, 0x40, 0x71, 0x41, 0xd5, 0x26, 0x63, 0x83, 0x16, 0x3a, 0x75, 0xc9, 0x4e, 0x65,
	0x9d, 0xa3, 0xd1, 0xfb, 0x0a, 0x9e, 0x87, 0x6c, 0xc5, 0x84, 0x60, 0x21, 0x0d, 0x59, 0xcc, 0xb4,
	0x2f, 0x5d, 0x6c, 0x14, 0x93, 0x83, 0x2d, 0xe4, 0x6f, 0x91, 0xa7, 0x25, 0x7c, 0x56, 0xa0, 0xa7,
	0x1a, 0xf4, 0x5e, 0x41, 0x4f, 0x5f, 0x74, 0x4f, 0xd7, 0x82, 0xdf, 0xab, 0x9b, 0xc1, 0x23, 0x24,
	0x77, 0x88, 0x6b, 0x6c, 0x17, 0xc6, 0xe4, 0xff, 0xd1, 0x82, 0xbe, 0x8e, 0x29, 0x89, 0x7e, 0x63,
	0xa2, 0x0c, 0x6e, 0x02, 0x9f, 0x84, 0x9c, 0x2e, 0x79, 0x92, 0xe0, 0x3b, 0x32, 0x5f, 0xb0, 0x5f,
	0x33, 0x81, 0xfa, 0xf5, 0xab, 0x2c, 0x46, 0x6a, 0x1a, 0x54, 0x01, 0x76, 0xc8, 0x7e, 0xc8, 0x47,
	0x86, 0x39, 0x6b, 0x12, 0xcf, 0x6b, 0x9e, 0x37, 0x84, 0x0f, 0xcc, 0x75, 0x29, 0xe6, 0x33, 0x55,
	0x74, 0xc5, 0xe3, 0x30, 0x4a, 0xd7, 0x26, 0xd4, 0x0e, 0x79, 0xa2, 0xdd, 0x2d, 0x32, 0xb6, 0x80,
	0xf7, 0x1a, 0xf6, 0x90, 0xbf, 0xca, 0x53, 0x93, 0x24, 0x1a, 0xa5, 0x71, 0x94, 0x6a, 0x07, 0x2b,
	0xdf, 0x0b, 0xf9, 0xb8, 0x80, 0x2e, 0x0b, 0xc4, 0xfb, 0x01, 0xba, 0x3c, 0x53, 0x34, 0x66, 0x77,
	0x2c, 0x36, 0x29, 0xd9, 0x3d, 0xf1, 0x87, 0x75, 0x71, 0x87, 0x0f, 0x23, 0x1c, 0x5e, 0x69, 0x26,
	0xe9, 0xa0, 0x93, 0x39, 0xf9, 0xfb, 0xd0, 0x36, 0x07, 0x6f, 0x1b, 0x5a, 0x57, 0x9f, 0xf5, 0xdf,
	0xf3, 0x1e, 0xe3, 0xff, 0xd7, 0xfd, 0x7f, 0xca, 0x3f, 0xc7, 0xff, 0xdb, 0x81, 0xde, 0x85, 0x08,
	0xb2, 0x9b, 0x32, 0x49, 0x5f, 0xc0, 0x33, 0x96, 0x06, 0x8b, 0x98, 0x51, 0xc1, 0x96, 0x77, 0x54,
	0x2e, 0x6f, 0x58, 0x98, 0xc7, 0x75, 0x60, 0x7b, 0x16, 0x25, 0x08, 0xce, 0x2a, 0xcc, 0xbb, 0x84,
	0x27, 0xbc, 0x14, 0x43, 0xb9, 0xbd, 0xca, 0x28, 0x76, 0x4f, 0x5e, 0xfe, 0x97, 0x62, 0xd2, 0xe7,
	0x0f, 0xab, 0x74, 0x08, 0xfd, 0x45, 0x1e, 0xc5, 0x21, 0x66, 0x56, 0x2a, 0x9a, 0x70, 0xec, 0x8c,
	0x22, 0x45, 0xbb, 0xc6, 0x3e, 0x42, 0xf3, 0x44, 0x5b, 0x7f, 0x7a, 0xd4, 0x71, 0xfa, 0x2d, 0x72,
	0x20, 0x6f, 0xa3, 0xec, 0x7f, 0xab, 0xea, 0xff, 0xd9, 0x06, 0x77, 0x64, 0x26, 0x66, 0x6a, 0x06,
	0xe6, 0x67, 0xe8, 0x85, 0xec, 0x2e, 0x5a, 0x32, 0x74, 0xcd, 0x53, 0x85, 0x55, 0xdf, 0x42, 0xb1,
	0x87, 0x4d, 0xb1, 0x0d, 0xfa, 0xf0, 0xcc, 0x70, 0x47, 0x9a, 0x7a, 0x9e, 0x2a, 0xb1, 0x21, 0x6e,
	0x58, 0x5b, 0xb0, 0x50, 0x2f, 0x23, 0xb4, 0x06, 0x18, 0x3a, 0xcd, 0x02, 0x81, 0x9d, 0x88, 0x4f,
	0xcb, 0x84, 0xaa, 0x1b, 0xc1, 0x82, 0x50, 0x9a, 0xd4, 0xb5, 0xc9, 0x87, 0x86, 0x73, 0x9d, 0x4d,
	0x6b, 0xc6, 0xdc, 0x12, 0x8a, 0x0b, 0x4c, 0xee, 0xde, 0x79, 0x41, 0xbb, 0xba, 0x40, 0x27, 0xea,
	0x1d, 0x17, 0x7c, 0x09, 0xcf, 0x73, 0xc9, 0xa8, 0x1e, 0x5e, 0x59, 0x44, 0x5f, 0xfa, 0x76, 0x6d,
	0xdd, 0x10, 0x9e, 0x32, 0x31, 0xb3, 0x60, 0xe9, 0xf6, 0x29, 0xf4, 0xb3, 0x38, 0x58, 0xb2, 0x84,
	0x61, 0x07, 0xa3, 0x73, 0xc4, 0x43, 0x53, 0xb6, 0x36, 0x79, 0x5c, 0xd9, 0xa7, 0xc6, 0xac, 0x87,
	0xba, 0x48, 0xd8, 0x2a, 0x8a, 0x51, 0x86, 0xc4, 0xaa, 0x6c, 0xe9, 0xa1, 0xb6, 0xd6, 0xb1, 0x35,
	0x7a, 0x5f, 0x83, 0xab, 0xb7, 0x46, 0xd9, 0x03, 0xdb, 0xa6, 0x07, 0x9e, 0x35, 0xd3, 0x5a, 0xaf,
	0x1b, 0x02, 0x48, 0x2d, 0xeb, 0x8e, 0xe3, 0x61, 0xa7, 0x5a, 0xf2, 0x15, 0x6a, 0x29, 0x5f, 0x1f,
	0xbc, 0x6f, 0xc7, 0xc3, 0x60, 0x33, 0x84, 0xa6, 0x25, 0xa2, 0x3d, 0x62, 0xbe, 0xa6, 0x85, 0xaa,
	0xda, 0xa3, 0x63, 0x3d, 0x10, 0xb3, 0x55, 0xab, 0x3d, 0xbe, 0x83, 0x9d, 0xb5, 0x6e, 0xf6, 0x4a,
	0x1e, 0x18, 0x79, 0x83, 0xb7, 0xe4, 0x35, 0xa6, 0x81, 0xf4, 0xd6, 0xcd, 0xd9, 0xc0, 0x24, 0x73,
	0x4c, 0x92, 0x69, 0x28, 0x8a, 0x6d, 0xcb, 0x78, 0xae, 0x70, 0x8e, 0x69, 0x22, 0x07, 0xae, 0x59,
	0x58, 0x7b, 0x15, 0x3c, 0xb7, 0xe8, 0x65, 0x3a, 0x91, 0x2f, 0xbe, 0x87, 0xfe, 0xc3, 0xf6, 0xf1,
	0xfa, 0xb0, 0x75, 0xcb, 0x36, 0x66, 0xd7, 0x74, 0x89, 0x3e, 0x7a, 0x7b, 0xd0, 0xbe, 0x0b, 0xe2,
	0x9c, 0x15, 0xcd, 0x62, 0x3f, 0xbe, 0x6d, 0x7d, 0xe3, 0xf8, 0xbf, 0xe3, 0x8e, 0x26, 0x79, 0x5a,
	0xaa, 0x38, 0x05, 0x17, 0xdb, 0x08, 0x23, 0xb6, 0x7b, 0xc1, 0x31, 0x7b, 0xe1, 0x55, 0x33, 0x84,
	0x9a, 0x3c, 0x9c, 0x6b, 0xa6, 0x5d, 0x0b, 0xa0, 0xaa, 0xb3, 0xe7, 0xc3, 0xce, 0xdb, 0xfa, 0x5b,
	0x46, 0xbf, 0xab, 0x6a, 0xd9, 0xfe, 0x11, 0x40, 0xed, 0xed, 0xf5, 0xa0, 0xf3, 0xcb, 0x35, 0x9d,
	0x93, 0x1f, 0x47, 0xe7, 0xb8, 0x47, 0x76, 0x01, 0xc6, 0x6f, 0xae, 0xae, 0x8a, 0x6f, 0xc7, 0x1f,
	0x81, 0x8b, 0x8f, 0x4e, 0x98, 0x0a, 0xc2, 0x40, 0x05, 0xb8, 0x44, 0xa0, 0xfe, 0xa5, 0x31, 0x0a,
	0xdd, 0x93, 0xa7, 0x4d, 0x85, 0x33, 0x44, 0x67, 0x1a, 0x24, 0x5d, 0x59, 0x1e, 0x4f, 0x8f, 0x60,
	0xc0, 0xc5, 0xba, 0x49, 0xab, 0x7e, 0xb8, 0x4e, 0x7b, 0x8d, 0x61, 0x94, 0x53, 0x67, 0xb1, 0x6d,
	0x7e, 0xbe, 0x3e, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xfa, 0xef, 0xec, 0x1f, 0x07, 0x00,
	0x00,
}