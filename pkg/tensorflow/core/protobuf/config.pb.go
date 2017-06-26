// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/config.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow2 "github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/core/framework"
import tensorflow10 "github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/core/framework"
import tensorflow13 "github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/core/framework"

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
func (OptimizerOptions_Level) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

// Control the use of the compiler/jit.  Experimental.
type OptimizerOptions_GlobalJitLevel int32

const (
	OptimizerOptions_DEFAULT OptimizerOptions_GlobalJitLevel = 0
	OptimizerOptions_OFF     OptimizerOptions_GlobalJitLevel = -1
	// The following settings turn on compilation, with higher values being
	// more aggressive.  Higher values may reduce opportunities for parallelism
	// and may use more memory.  (At present, there is no distinction, but this
	// is expected to change.)
	OptimizerOptions_ON_1 OptimizerOptions_GlobalJitLevel = 1
	OptimizerOptions_ON_2 OptimizerOptions_GlobalJitLevel = 2
)

var OptimizerOptions_GlobalJitLevel_name = map[int32]string{
	0:  "DEFAULT",
	-1: "OFF",
	1:  "ON_1",
	2:  "ON_2",
}
var OptimizerOptions_GlobalJitLevel_value = map[string]int32{
	"DEFAULT": 0,
	"OFF":     -1,
	"ON_1":    1,
	"ON_2":    2,
}

func (x OptimizerOptions_GlobalJitLevel) String() string {
	return proto.EnumName(OptimizerOptions_GlobalJitLevel_name, int32(x))
}
func (OptimizerOptions_GlobalJitLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 1}
}

// TODO(pbar) Turn this into a TraceOptions proto which allows
// tracing to be controlled in a more orthogonal manner?
type RunOptions_TraceLevel int32

const (
	RunOptions_NO_TRACE       RunOptions_TraceLevel = 0
	RunOptions_SOFTWARE_TRACE RunOptions_TraceLevel = 1
	RunOptions_HARDWARE_TRACE RunOptions_TraceLevel = 2
	RunOptions_FULL_TRACE     RunOptions_TraceLevel = 3
)

var RunOptions_TraceLevel_name = map[int32]string{
	0: "NO_TRACE",
	1: "SOFTWARE_TRACE",
	2: "HARDWARE_TRACE",
	3: "FULL_TRACE",
}
var RunOptions_TraceLevel_value = map[string]int32{
	"NO_TRACE":       0,
	"SOFTWARE_TRACE": 1,
	"HARDWARE_TRACE": 2,
	"FULL_TRACE":     3,
}

func (x RunOptions_TraceLevel) String() string {
	return proto.EnumName(RunOptions_TraceLevel_name, int32(x))
}
func (RunOptions_TraceLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{6, 0} }

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
	// A comma-separated list of GPU ids that determines the 'visible'
	// to 'virtual' mapping of GPU devices.  For example, if TensorFlow
	// can see 8 GPU devices in the process, and one wanted to map
	// visible GPU devices 5 and 3 as "/gpu:0", and "/gpu:1", then one
	// would specify this field as "5,3".  This field is similar in
	// spirit to the CUDA_VISIBLE_DEVICES environment variable, except
	// it applies to the visible GPU devices in the process.
	//
	// NOTE: The GPU driver provides the process with the visible GPUs
	// in an order which is not guaranteed to have any correlation to
	// the *physical* GPU id in the machine.  This field is used for
	// remapping "visible" to "virtual", which means this operates only
	// after the process starts.  Users are required to use vendor
	// specific mechanisms (e.g., CUDA_VISIBLE_DEVICES) to control the
	// physical to visible device mapping prior to invoking TensorFlow.
	VisibleDeviceList string `protobuf:"bytes,5,opt,name=visible_device_list,json=visibleDeviceList" json:"visible_device_list,omitempty"`
	// In the event polling loop sleep this many microseconds between
	// PollEvents calls, when the queue is not empty.  If value is not
	// set or set to 0, gets set to a non-zero default.
	PollingActiveDelayUsecs int32 `protobuf:"varint,6,opt,name=polling_active_delay_usecs,json=pollingActiveDelayUsecs" json:"polling_active_delay_usecs,omitempty"`
	// In the event polling loop sleep this many millisconds between
	// PollEvents calls, when the queue is empty.  If value is not
	// set or set to 0, gets set to a non-zero default.
	PollingInactiveDelayMsecs int32 `protobuf:"varint,7,opt,name=polling_inactive_delay_msecs,json=pollingInactiveDelayMsecs" json:"polling_inactive_delay_msecs,omitempty"`
	// Force all tensors to be gpu_compatible. On a GPU-enabled TensorFlow,
	// enabling this option forces all CPU tensors to be allocated with Cuda
	// pinned memory. Normally, TensorFlow will infer which tensors should be
	// allocated as the pinned memory. But in case where the inference is
	// incomplete, this option can significantly speed up the cross-device memory
	// copy performance as long as it fits the memory.
	// Note that this option is not something that should be
	// enabled by default for unknown or very large models, since all Cuda pinned
	// memory is unpageable, having too much pinned memory might negatively impact
	// the overall host system performance.
	ForceGpuCompatible bool `protobuf:"varint,8,opt,name=force_gpu_compatible,json=forceGpuCompatible" json:"force_gpu_compatible,omitempty"`
}

func (m *GPUOptions) Reset()                    { *m = GPUOptions{} }
func (m *GPUOptions) String() string            { return proto.CompactTextString(m) }
func (*GPUOptions) ProtoMessage()               {}
func (*GPUOptions) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GPUOptions) GetPerProcessGpuMemoryFraction() float64 {
	if m != nil {
		return m.PerProcessGpuMemoryFraction
	}
	return 0
}

func (m *GPUOptions) GetAllocatorType() string {
	if m != nil {
		return m.AllocatorType
	}
	return ""
}

func (m *GPUOptions) GetDeferredDeletionBytes() int64 {
	if m != nil {
		return m.DeferredDeletionBytes
	}
	return 0
}

func (m *GPUOptions) GetAllowGrowth() bool {
	if m != nil {
		return m.AllowGrowth
	}
	return false
}

func (m *GPUOptions) GetVisibleDeviceList() string {
	if m != nil {
		return m.VisibleDeviceList
	}
	return ""
}

func (m *GPUOptions) GetPollingActiveDelayUsecs() int32 {
	if m != nil {
		return m.PollingActiveDelayUsecs
	}
	return 0
}

func (m *GPUOptions) GetPollingInactiveDelayMsecs() int32 {
	if m != nil {
		return m.PollingInactiveDelayMsecs
	}
	return 0
}

func (m *GPUOptions) GetForceGpuCompatible() bool {
	if m != nil {
		return m.ForceGpuCompatible
	}
	return false
}

// Options passed to the graph optimizer
type OptimizerOptions struct {
	// If true, optimize the graph using common subexpression elimination.
	DoCommonSubexpressionElimination bool `protobuf:"varint,1,opt,name=do_common_subexpression_elimination,json=doCommonSubexpressionElimination" json:"do_common_subexpression_elimination,omitempty"`
	// If true, perform constant folding optimization on the graph.
	DoConstantFolding bool `protobuf:"varint,2,opt,name=do_constant_folding,json=doConstantFolding" json:"do_constant_folding,omitempty"`
	// If true, perform function inlining on the graph.
	DoFunctionInlining bool                            `protobuf:"varint,4,opt,name=do_function_inlining,json=doFunctionInlining" json:"do_function_inlining,omitempty"`
	OptLevel           OptimizerOptions_Level          `protobuf:"varint,3,opt,name=opt_level,json=optLevel,enum=tensorflow.OptimizerOptions_Level" json:"opt_level,omitempty"`
	GlobalJitLevel     OptimizerOptions_GlobalJitLevel `protobuf:"varint,5,opt,name=global_jit_level,json=globalJitLevel,enum=tensorflow.OptimizerOptions_GlobalJitLevel" json:"global_jit_level,omitempty"`
}

func (m *OptimizerOptions) Reset()                    { *m = OptimizerOptions{} }
func (m *OptimizerOptions) String() string            { return proto.CompactTextString(m) }
func (*OptimizerOptions) ProtoMessage()               {}
func (*OptimizerOptions) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *OptimizerOptions) GetDoCommonSubexpressionElimination() bool {
	if m != nil {
		return m.DoCommonSubexpressionElimination
	}
	return false
}

func (m *OptimizerOptions) GetDoConstantFolding() bool {
	if m != nil {
		return m.DoConstantFolding
	}
	return false
}

func (m *OptimizerOptions) GetDoFunctionInlining() bool {
	if m != nil {
		return m.DoFunctionInlining
	}
	return false
}

func (m *OptimizerOptions) GetOptLevel() OptimizerOptions_Level {
	if m != nil {
		return m.OptLevel
	}
	return OptimizerOptions_L1
}

func (m *OptimizerOptions) GetGlobalJitLevel() OptimizerOptions_GlobalJitLevel {
	if m != nil {
		return m.GlobalJitLevel
	}
	return OptimizerOptions_DEFAULT
}

type GraphOptions struct {
	// If true, use control flow to schedule the activation of Recv nodes.
	// (Currently ignored.)
	EnableRecvScheduling bool `protobuf:"varint,2,opt,name=enable_recv_scheduling,json=enableRecvScheduling" json:"enable_recv_scheduling,omitempty"`
	// Options controlling how graph is optimized.
	OptimizerOptions *OptimizerOptions `protobuf:"bytes,3,opt,name=optimizer_options,json=optimizerOptions" json:"optimizer_options,omitempty"`
	// The number of steps to run before returning a cost model detailing
	// the memory usage and performance of each node of the graph. 0 means
	// no cost model.
	BuildCostModel int64 `protobuf:"varint,4,opt,name=build_cost_model,json=buildCostModel" json:"build_cost_model,omitempty"`
	// The number of steps to skip before collecting statistics for the
	// cost model.
	BuildCostModelAfter int64 `protobuf:"varint,9,opt,name=build_cost_model_after,json=buildCostModelAfter" json:"build_cost_model_after,omitempty"`
	// Annotate each Node with Op output shape data, to the extent it can
	// be statically inferred.
	InferShapes bool `protobuf:"varint,5,opt,name=infer_shapes,json=inferShapes" json:"infer_shapes,omitempty"`
	// Only place the subgraphs that are run, rather than the entire graph.
	//
	// This is useful for interactive graph building, where one might
	// produce graphs that cannot be placed during the debugging
	// process.  In particular, it allows the client to continue work in
	// a session after adding a node to a graph whose placement
	// constraints are unsatisfiable.
	PlacePrunedGraph bool `protobuf:"varint,6,opt,name=place_pruned_graph,json=placePrunedGraph" json:"place_pruned_graph,omitempty"`
	// If true, transfer float values between processes as bfloat16.
	EnableBfloat16Sendrecv bool `protobuf:"varint,7,opt,name=enable_bfloat16_sendrecv,json=enableBfloat16Sendrecv" json:"enable_bfloat16_sendrecv,omitempty"`
	// If > 0, record a timeline every this many steps.
	// EXPERIMENTAL: This currently has no effect in MasterSession.
	TimelineStep int32 `protobuf:"varint,8,opt,name=timeline_step,json=timelineStep" json:"timeline_step,omitempty"`
	// Options that control the type and amount of graph rewriting.
	// Not currently configurable via the public Python API (i.e. there is no API
	// stability guarantee if you import RewriterConfig explicitly).
	RewriteOptions *RewriterConfig `protobuf:"bytes,10,opt,name=rewrite_options,json=rewriteOptions" json:"rewrite_options,omitempty"`
}

func (m *GraphOptions) Reset()                    { *m = GraphOptions{} }
func (m *GraphOptions) String() string            { return proto.CompactTextString(m) }
func (*GraphOptions) ProtoMessage()               {}
func (*GraphOptions) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GraphOptions) GetEnableRecvScheduling() bool {
	if m != nil {
		return m.EnableRecvScheduling
	}
	return false
}

func (m *GraphOptions) GetOptimizerOptions() *OptimizerOptions {
	if m != nil {
		return m.OptimizerOptions
	}
	return nil
}

func (m *GraphOptions) GetBuildCostModel() int64 {
	if m != nil {
		return m.BuildCostModel
	}
	return 0
}

func (m *GraphOptions) GetBuildCostModelAfter() int64 {
	if m != nil {
		return m.BuildCostModelAfter
	}
	return 0
}

func (m *GraphOptions) GetInferShapes() bool {
	if m != nil {
		return m.InferShapes
	}
	return false
}

func (m *GraphOptions) GetPlacePrunedGraph() bool {
	if m != nil {
		return m.PlacePrunedGraph
	}
	return false
}

func (m *GraphOptions) GetEnableBfloat16Sendrecv() bool {
	if m != nil {
		return m.EnableBfloat16Sendrecv
	}
	return false
}

func (m *GraphOptions) GetTimelineStep() int32 {
	if m != nil {
		return m.TimelineStep
	}
	return 0
}

func (m *GraphOptions) GetRewriteOptions() *RewriterConfig {
	if m != nil {
		return m.RewriteOptions
	}
	return nil
}

type ThreadPoolOptionProto struct {
	// The number of threads in the pool.
	//
	// 0 means the system picks a value based on where this option proto is used
	// (see the declaration of the specific field for more info).
	NumThreads int32 `protobuf:"varint,1,opt,name=num_threads,json=numThreads" json:"num_threads,omitempty"`
}

func (m *ThreadPoolOptionProto) Reset()                    { *m = ThreadPoolOptionProto{} }
func (m *ThreadPoolOptionProto) String() string            { return proto.CompactTextString(m) }
func (*ThreadPoolOptionProto) ProtoMessage()               {}
func (*ThreadPoolOptionProto) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ThreadPoolOptionProto) GetNumThreads() int32 {
	if m != nil {
		return m.NumThreads
	}
	return 0
}

type RPCOptions struct {
	// If true, always use RPC to contact the session target.
	//
	// If false (the default option), TensorFlow may use an optimized
	// transport for client-master communication that avoids the RPC
	// stack. This option is primarily for used testing the RPC stack.
	UseRpcForInprocessMaster bool `protobuf:"varint,1,opt,name=use_rpc_for_inprocess_master,json=useRpcForInprocessMaster" json:"use_rpc_for_inprocess_master,omitempty"`
}

func (m *RPCOptions) Reset()                    { *m = RPCOptions{} }
func (m *RPCOptions) String() string            { return proto.CompactTextString(m) }
func (*RPCOptions) ProtoMessage()               {}
func (*RPCOptions) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *RPCOptions) GetUseRpcForInprocessMaster() bool {
	if m != nil {
		return m.UseRpcForInprocessMaster
	}
	return false
}

// Session configuration parameters.
// The system picks appropriate values for fields that are not set.
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
	// true or session_inter_op_thread_pool is configured.
	InterOpParallelismThreads int32 `protobuf:"varint,5,opt,name=inter_op_parallelism_threads,json=interOpParallelismThreads" json:"inter_op_parallelism_threads,omitempty"`
	// If true, use a new set of threads for this session rather than the global
	// pool of threads. Only supported by direct sessions.
	//
	// If false, use the global threads created by the first session, or the
	// per-session thread pools configured by session_inter_op_thread_pool.
	//
	// This option is deprecated. The same effect can be achieved by setting
	// session_inter_op_thread_pool to have one element, whose num_threads equals
	// inter_op_parallelism_threads.
	UsePerSessionThreads bool `protobuf:"varint,9,opt,name=use_per_session_threads,json=usePerSessionThreads" json:"use_per_session_threads,omitempty"`
	// This option is experimental - it may be replaced with a different mechanism
	// in the future. The intended use is for when some session invocations need
	// to run in a background pool limited to a small number of threads.
	//
	// Configures session thread pools. If this is configured, then RunOptions for
	// a Run call can select the thread pool to use.
	//
	// If a pool's num_threads is 0, then inter_op_parallelism_threads is used.
	SessionInterOpThreadPool []*ThreadPoolOptionProto `protobuf:"bytes,12,rep,name=session_inter_op_thread_pool,json=sessionInterOpThreadPool" json:"session_inter_op_thread_pool,omitempty"`
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
	// Options that apply when this session uses the distributed runtime.
	RpcOptions *RPCOptions `protobuf:"bytes,13,opt,name=rpc_options,json=rpcOptions" json:"rpc_options,omitempty"`
	// Optional list of all workers to use in this session.
	ClusterDef *ClusterDef `protobuf:"bytes,14,opt,name=cluster_def,json=clusterDef" json:"cluster_def,omitempty"`
}

func (m *ConfigProto) Reset()                    { *m = ConfigProto{} }
func (m *ConfigProto) String() string            { return proto.CompactTextString(m) }
func (*ConfigProto) ProtoMessage()               {}
func (*ConfigProto) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ConfigProto) GetDeviceCount() map[string]int32 {
	if m != nil {
		return m.DeviceCount
	}
	return nil
}

func (m *ConfigProto) GetIntraOpParallelismThreads() int32 {
	if m != nil {
		return m.IntraOpParallelismThreads
	}
	return 0
}

func (m *ConfigProto) GetInterOpParallelismThreads() int32 {
	if m != nil {
		return m.InterOpParallelismThreads
	}
	return 0
}

func (m *ConfigProto) GetUsePerSessionThreads() bool {
	if m != nil {
		return m.UsePerSessionThreads
	}
	return false
}

func (m *ConfigProto) GetSessionInterOpThreadPool() []*ThreadPoolOptionProto {
	if m != nil {
		return m.SessionInterOpThreadPool
	}
	return nil
}

func (m *ConfigProto) GetPlacementPeriod() int32 {
	if m != nil {
		return m.PlacementPeriod
	}
	return 0
}

func (m *ConfigProto) GetDeviceFilters() []string {
	if m != nil {
		return m.DeviceFilters
	}
	return nil
}

func (m *ConfigProto) GetGpuOptions() *GPUOptions {
	if m != nil {
		return m.GpuOptions
	}
	return nil
}

func (m *ConfigProto) GetAllowSoftPlacement() bool {
	if m != nil {
		return m.AllowSoftPlacement
	}
	return false
}

func (m *ConfigProto) GetLogDevicePlacement() bool {
	if m != nil {
		return m.LogDevicePlacement
	}
	return false
}

func (m *ConfigProto) GetGraphOptions() *GraphOptions {
	if m != nil {
		return m.GraphOptions
	}
	return nil
}

func (m *ConfigProto) GetOperationTimeoutInMs() int64 {
	if m != nil {
		return m.OperationTimeoutInMs
	}
	return 0
}

func (m *ConfigProto) GetRpcOptions() *RPCOptions {
	if m != nil {
		return m.RpcOptions
	}
	return nil
}

func (m *ConfigProto) GetClusterDef() *ClusterDef {
	if m != nil {
		return m.ClusterDef
	}
	return nil
}

// Options for a single Run() call.
type RunOptions struct {
	TraceLevel RunOptions_TraceLevel `protobuf:"varint,1,opt,name=trace_level,json=traceLevel,enum=tensorflow.RunOptions_TraceLevel" json:"trace_level,omitempty"`
	// Time to wait for operation to complete in milliseconds.
	TimeoutInMs int64 `protobuf:"varint,2,opt,name=timeout_in_ms,json=timeoutInMs" json:"timeout_in_ms,omitempty"`
	// The thread pool to use, if session_inter_op_thread_pool is configured.
	InterOpThreadPool int32 `protobuf:"varint,3,opt,name=inter_op_thread_pool,json=interOpThreadPool" json:"inter_op_thread_pool,omitempty"`
	// Whether the partition graph(s) executed by the executor(s) should be
	// outputted via RunMetadata.
	OutputPartitionGraphs bool `protobuf:"varint,5,opt,name=output_partition_graphs,json=outputPartitionGraphs" json:"output_partition_graphs,omitempty"`
	// EXPERIMENTAL.  Options used to initialize DebuggerState, if enabled.
	DebugOptions *DebugOptions `protobuf:"bytes,6,opt,name=debug_options,json=debugOptions" json:"debug_options,omitempty"`
}

func (m *RunOptions) Reset()                    { *m = RunOptions{} }
func (m *RunOptions) String() string            { return proto.CompactTextString(m) }
func (*RunOptions) ProtoMessage()               {}
func (*RunOptions) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *RunOptions) GetTraceLevel() RunOptions_TraceLevel {
	if m != nil {
		return m.TraceLevel
	}
	return RunOptions_NO_TRACE
}

func (m *RunOptions) GetTimeoutInMs() int64 {
	if m != nil {
		return m.TimeoutInMs
	}
	return 0
}

func (m *RunOptions) GetInterOpThreadPool() int32 {
	if m != nil {
		return m.InterOpThreadPool
	}
	return 0
}

func (m *RunOptions) GetOutputPartitionGraphs() bool {
	if m != nil {
		return m.OutputPartitionGraphs
	}
	return false
}

func (m *RunOptions) GetDebugOptions() *DebugOptions {
	if m != nil {
		return m.DebugOptions
	}
	return nil
}

// Metadata output (i.e., non-Tensor) for a single Run() call.
type RunMetadata struct {
	// Statistics traced for this step. Populated if tracing is turned on via the
	// "RunOptions" proto.
	// EXPERIMENTAL: The format and set of events may change in future versions.
	StepStats *tensorflow13.StepStats `protobuf:"bytes,1,opt,name=step_stats,json=stepStats" json:"step_stats,omitempty"`
	// The cost graph for the computation defined by the run call.
	CostGraph *tensorflow2.CostGraphDef `protobuf:"bytes,2,opt,name=cost_graph,json=costGraph" json:"cost_graph,omitempty"`
	// Graphs of the partitions executed by executors.
	PartitionGraphs []*tensorflow10.GraphDef `protobuf:"bytes,3,rep,name=partition_graphs,json=partitionGraphs" json:"partition_graphs,omitempty"`
}

func (m *RunMetadata) Reset()                    { *m = RunMetadata{} }
func (m *RunMetadata) String() string            { return proto.CompactTextString(m) }
func (*RunMetadata) ProtoMessage()               {}
func (*RunMetadata) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *RunMetadata) GetStepStats() *tensorflow13.StepStats {
	if m != nil {
		return m.StepStats
	}
	return nil
}

func (m *RunMetadata) GetCostGraph() *tensorflow2.CostGraphDef {
	if m != nil {
		return m.CostGraph
	}
	return nil
}

func (m *RunMetadata) GetPartitionGraphs() []*tensorflow10.GraphDef {
	if m != nil {
		return m.PartitionGraphs
	}
	return nil
}

func init() {
	proto.RegisterType((*GPUOptions)(nil), "tensorflow.GPUOptions")
	proto.RegisterType((*OptimizerOptions)(nil), "tensorflow.OptimizerOptions")
	proto.RegisterType((*GraphOptions)(nil), "tensorflow.GraphOptions")
	proto.RegisterType((*ThreadPoolOptionProto)(nil), "tensorflow.ThreadPoolOptionProto")
	proto.RegisterType((*RPCOptions)(nil), "tensorflow.RPCOptions")
	proto.RegisterType((*ConfigProto)(nil), "tensorflow.ConfigProto")
	proto.RegisterType((*RunOptions)(nil), "tensorflow.RunOptions")
	proto.RegisterType((*RunMetadata)(nil), "tensorflow.RunMetadata")
	proto.RegisterEnum("tensorflow.OptimizerOptions_Level", OptimizerOptions_Level_name, OptimizerOptions_Level_value)
	proto.RegisterEnum("tensorflow.OptimizerOptions_GlobalJitLevel", OptimizerOptions_GlobalJitLevel_name, OptimizerOptions_GlobalJitLevel_value)
	proto.RegisterEnum("tensorflow.RunOptions_TraceLevel", RunOptions_TraceLevel_name, RunOptions_TraceLevel_value)
}

func init() { proto.RegisterFile("tensorflow/core/protobuf/config.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x57, 0xdd, 0x6e, 0xdb, 0xc8,
	0x15, 0x0e, 0x25, 0x2b, 0x91, 0x8f, 0x1c, 0x45, 0x99, 0x75, 0x62, 0xd6, 0x0d, 0xb0, 0x5a, 0x6d,
	0x53, 0xb8, 0x3f, 0x2b, 0x25, 0xde, 0xdd, 0x24, 0x68, 0xd1, 0x0d, 0x6c, 0xd9, 0x72, 0x9d, 0xda,
	0x6b, 0x61, 0x64, 0xa3, 0x40, 0x6f, 0x06, 0x14, 0x39, 0xa4, 0xa7, 0xa6, 0x38, 0x83, 0x99, 0xa1,
	0x5d, 0xf7, 0x1d, 0x0a, 0xf4, 0x61, 0x7a, 0xdb, 0x9b, 0x3e, 0x54, 0xd1, 0xbb, 0x16, 0x33, 0x43,
	0x8a, 0xf4, 0x5f, 0xea, 0x2b, 0xf2, 0x7c, 0xdf, 0x39, 0x9c, 0xf9, 0xce, 0x9f, 0x05, 0xaf, 0x35,
	0xcd, 0x14, 0x97, 0x71, 0xca, 0xaf, 0x46, 0x21, 0x97, 0x74, 0x24, 0x24, 0xd7, 0x7c, 0x9e, 0xc7,
	0xa3, 0x90, 0x67, 0x31, 0x4b, 0x86, 0xf6, 0x1d, 0x41, 0x45, 0xdb, 0xfc, 0xe5, 0x6d, 0x97, 0x58,
	0x06, 0x0b, 0x7a, 0xc5, 0xe5, 0xc5, 0x28, 0xe4, 0x4a, 0x93, 0x44, 0x06, 0xe2, 0xdc, 0xf9, 0x6d,
	0xbe, 0x7e, 0x98, 0x5b, 0xa7, 0x7d, 0x26, 0xa4, 0xd2, 0x54, 0x10, 0xa5, 0x03, 0xad, 0x0a, 0xee,
	0xcf, 0x1e, 0x3c, 0x71, 0x44, 0xe7, 0x79, 0x71, 0xe0, 0xcd, 0x9f, 0x3f, 0x7c, 0xaf, 0x34, 0x57,
	0x9a, 0xca, 0x82, 0x37, 0x7c, 0x90, 0x27, 0xe9, 0x95, 0x64, 0x9a, 0x4a, 0x52, 0x17, 0x62, 0xf0,
	0x8f, 0x26, 0xc0, 0xc1, 0xf4, 0xec, 0x44, 0x68, 0xc6, 0x33, 0x85, 0xf6, 0xe0, 0x4b, 0x41, 0x25,
	0x11, 0x92, 0x87, 0x54, 0x29, 0x92, 0x88, 0x9c, 0x2c, 0xe8, 0x82, 0xcb, 0x6b, 0x12, 0xcb, 0x20,
	0x34, 0x1c, 0xdf, 0xeb, 0x7b, 0x5b, 0x1e, 0xfe, 0xa9, 0xa0, 0x72, 0xea, 0x58, 0x07, 0x22, 0x3f,
	0xb6, 0x9c, 0x49, 0x41, 0x41, 0xaf, 0xa1, 0x1b, 0xa4, 0x29, 0x0f, 0x03, 0xcd, 0x25, 0xd1, 0xd7,
	0x82, 0xfa, 0x8d, 0xbe, 0xb7, 0xb5, 0x8a, 0x9f, 0x2e, 0xad, 0xa7, 0xd7, 0x82, 0xa2, 0x77, 0xb0,
	0x11, 0xd1, 0x98, 0x4a, 0x49, 0x23, 0x12, 0xd1, 0x94, 0x1a, 0x5f, 0x32, 0xbf, 0xd6, 0x54, 0xf9,
	0xcd, 0xbe, 0xb7, 0xd5, 0xc4, 0x2f, 0x4a, 0x78, 0xaf, 0x40, 0x77, 0x0d, 0x88, 0xbe, 0x82, 0x35,
	0x13, 0xe8, 0x8a, 0x24, 0x92, 0x5f, 0xe9, 0x73, 0x7f, 0xa5, 0xef, 0x6d, 0xb5, 0x71, 0xc7, 0xda,
	0x0e, 0xac, 0x09, 0x0d, 0xe1, 0x8b, 0x4b, 0xa6, 0xd8, 0x3c, 0xa5, 0x24, 0xa2, 0x97, 0x2c, 0xa4,
	0x24, 0x65, 0x4a, 0xfb, 0x2d, 0x7b, 0x8c, 0xe7, 0x05, 0xb4, 0x67, 0x91, 0x23, 0xa6, 0x34, 0xfa,
	0x2d, 0x6c, 0x0a, 0x9e, 0xa6, 0x2c, 0x4b, 0x88, 0xb9, 0xc3, 0xa5, 0x71, 0x4b, 0x83, 0x6b, 0x92,
	0x2b, 0x1a, 0x2a, 0xff, 0x71, 0xdf, 0xdb, 0x6a, 0xe1, 0x8d, 0x82, 0xb1, 0x63, 0x09, 0x7b, 0x06,
	0x3f, 0x33, 0x30, 0xfa, 0x08, 0xaf, 0x4a, 0x67, 0x96, 0xdd, 0x70, 0x5f, 0x58, 0xf7, 0x27, 0xd6,
	0xfd, 0x27, 0x05, 0xe7, 0xb0, 0xa0, 0xd8, 0x00, 0xc7, 0x36, 0xc0, 0x1b, 0x58, 0x8f, 0xb9, 0x0c,
	0xa9, 0xd5, 0x3b, 0xe4, 0x0b, 0x11, 0x68, 0x73, 0x3e, 0xbf, 0x6d, 0x2f, 0x86, 0x2c, 0x76, 0x20,
	0xf2, 0xf1, 0x12, 0x19, 0xfc, 0xab, 0x09, 0x3d, 0x93, 0xb3, 0x05, 0xfb, 0x2b, 0x95, 0x65, 0xf2,
	0x8e, 0xe1, 0xeb, 0x88, 0x1b, 0xff, 0x05, 0xcf, 0x88, 0xca, 0xe7, 0xf4, 0x2f, 0x42, 0x52, 0xa5,
	0x8c, 0xaa, 0x34, 0x65, 0x0b, 0x96, 0x05, 0xcb, 0x04, 0xb6, 0x71, 0x3f, 0xe2, 0x63, 0xcb, 0x9c,
	0xd5, 0x89, 0xfb, 0x15, 0xcf, 0x68, 0x68, 0xc3, 0x65, 0x4a, 0x07, 0x99, 0x26, 0x31, 0x4f, 0x23,
	0x96, 0x25, 0x36, 0x95, 0x6d, 0xfc, 0xdc, 0xb8, 0x3b, 0x64, 0xe2, 0x00, 0x73, 0x8b, 0x88, 0x93,
	0x38, 0xcf, 0x6c, 0x11, 0x10, 0x96, 0xa5, 0x2c, 0x33, 0x0e, 0x2e, 0x3d, 0x28, 0xe2, 0x93, 0x02,
	0x3a, 0x2c, 0x10, 0xf4, 0x11, 0x56, 0xb9, 0xd0, 0x24, 0xa5, 0x97, 0x34, 0xb5, 0x29, 0xef, 0x6e,
	0x0f, 0x6a, 0x05, 0x3c, 0xbc, 0x7d, 0xc3, 0xe1, 0x91, 0x61, 0xe2, 0x36, 0x17, 0xda, 0x3e, 0xa1,
	0x33, 0xe8, 0x25, 0x29, 0x9f, 0x07, 0x29, 0xf9, 0x33, 0x2b, 0xe3, 0xb4, 0x6c, 0x9c, 0x5f, 0x7d,
	0x36, 0xce, 0x81, 0x75, 0xfa, 0xc4, 0x5c, 0x18, 0xdc, 0x4d, 0x6e, 0xbc, 0x0f, 0xfa, 0xd0, 0x72,
	0xf1, 0x1f, 0x43, 0xe3, 0xe8, 0x6d, 0xef, 0x11, 0x7a, 0x06, 0x8d, 0xa3, 0x37, 0xbd, 0xff, 0x96,
	0x7f, 0xde, 0x60, 0x0c, 0xdd, 0x9b, 0x31, 0x50, 0x07, 0x9e, 0xec, 0xed, 0x4f, 0x76, 0xce, 0x8e,
	0x4e, 0x7b, 0x8f, 0x50, 0x0f, 0x9a, 0x27, 0x93, 0x49, 0xcd, 0x01, 0xb5, 0x61, 0xe5, 0xe4, 0x47,
	0xf2, 0xb6, 0x57, 0x3e, 0x6d, 0xf7, 0x1a, 0x83, 0xbf, 0xaf, 0xc0, 0xda, 0x81, 0x99, 0x1a, 0x65,
	0x02, 0xbf, 0x83, 0x97, 0x34, 0x0b, 0x4c, 0xd1, 0x4a, 0x1a, 0x5e, 0x12, 0x15, 0x9e, 0xd3, 0x28,
	0x4f, 0x2b, 0xd1, 0xd7, 0x1d, 0x8a, 0x69, 0x78, 0x39, 0x5b, 0x62, 0xe8, 0x10, 0x9e, 0xf3, 0xf2,
	0x82, 0x84, 0xbb, 0x50, 0x56, 0xcd, 0xce, 0xf6, 0xab, 0xcf, 0xa9, 0x80, 0x7b, 0xfc, 0x76, 0x05,
	0x6d, 0x41, 0x6f, 0x9e, 0xb3, 0x34, 0x22, 0x76, 0xf0, 0x2d, 0x78, 0x44, 0x53, 0x9b, 0xbe, 0x26,
	0xee, 0x5a, 0xfb, 0x98, 0x2b, 0x7d, 0x6c, 0xac, 0xe8, 0x5b, 0x78, 0x79, 0x9b, 0x49, 0x82, 0x58,
	0x53, 0xe9, 0xaf, 0x5a, 0xfe, 0x17, 0x37, 0xf9, 0x3b, 0x06, 0x32, 0x8d, 0xcb, 0xb2, 0x98, 0x4a,
	0xa2, 0xce, 0x03, 0x41, 0x95, 0x4d, 0x55, 0x1b, 0x77, 0xac, 0x6d, 0x66, 0x4d, 0xe8, 0xd7, 0x80,
	0x44, 0x1a, 0x84, 0x94, 0x08, 0x99, 0x67, 0x34, 0x72, 0xc3, 0xd7, 0x36, 0x60, 0x1b, 0xf7, 0x2c,
	0x32, 0xb5, 0x80, 0xd5, 0x0d, 0x7d, 0x00, 0xbf, 0x10, 0x6c, 0x1e, 0xa7, 0x3c, 0xd0, 0x6f, 0xdf,
	0x11, 0x45, 0xb3, 0xc8, 0xa8, 0x67, 0xbb, 0xae, 0x8d, 0x0b, 0x41, 0x77, 0x0b, 0x78, 0x56, 0xa0,
	0xe8, 0x6b, 0x78, 0xaa, 0xd9, 0x82, 0xa6, 0x2c, 0xa3, 0xc4, 0x8c, 0x64, 0xdb, 0x6b, 0x2d, 0xbc,
	0x56, 0x1a, 0x67, 0x9a, 0x0a, 0x34, 0x86, 0x67, 0xc5, 0xd4, 0x5c, 0xea, 0x0a, 0x56, 0xd7, 0xfa,
	0x98, 0x1d, 0xe2, 0x62, 0xb0, 0x8e, 0xed, 0x5c, 0xc5, 0xdd, 0xc2, 0xa5, 0xd0, 0xf4, 0xd3, 0x4a,
	0xdb, 0xeb, 0x35, 0xf0, 0x6b, 0x75, 0xc1, 0xc4, 0xff, 0xed, 0xcd, 0xc1, 0x07, 0x78, 0x71, 0x7a,
	0x2e, 0x69, 0x10, 0x4d, 0x39, 0x4f, 0x5d, 0x84, 0xa9, 0x5d, 0x58, 0x5f, 0x42, 0x27, 0xcb, 0x17,
	0x44, 0x5b, 0x50, 0xd9, 0x1e, 0x6e, 0x61, 0xc8, 0xf2, 0x85, 0xa3, 0xab, 0xc1, 0x11, 0x00, 0x9e,
	0x8e, 0xcb, 0x44, 0xfe, 0x00, 0xaf, 0x72, 0x45, 0x89, 0x14, 0x21, 0x89, 0xb9, 0x24, 0x2c, 0x2b,
	0x27, 0xfa, 0x22, 0x30, 0xcb, 0xa2, 0x98, 0x01, 0x7e, 0xae, 0x28, 0x16, 0xe1, 0x84, 0xcb, 0xc3,
	0x92, 0x70, 0x6c, 0xf1, 0xc1, 0xdf, 0x9e, 0x40, 0xc7, 0xdd, 0xc7, 0x7d, 0xfe, 0x0f, 0xb0, 0x56,
	0xcc, 0xd1, 0x90, 0xe7, 0x99, 0xf6, 0xbd, 0x7e, 0x73, 0xab, 0xb3, 0xbd, 0x55, 0x97, 0xa1, 0x46,
	0x1f, 0xba, 0xc9, 0x3a, 0x36, 0xd4, 0xfd, 0x4c, 0xcb, 0x6b, 0xdc, 0x89, 0x2a, 0x8b, 0x99, 0x97,
	0x2c, 0xd3, 0x32, 0x20, 0x5c, 0x10, 0x11, 0xc8, 0x20, 0x4d, 0x69, 0xca, 0x54, 0x75, 0xb9, 0x86,
	0x9b, 0x97, 0x96, 0x73, 0x22, 0xa6, 0x15, 0xa3, 0xb8, 0x6b, 0x11, 0xc0, 0x56, 0xfb, 0xbd, 0x01,
	0x5a, 0xcb, 0x00, 0xa6, 0xb4, 0xef, 0x09, 0xf0, 0x3d, 0x6c, 0x18, 0x79, 0xcc, 0xaa, 0x53, 0x45,
	0x16, 0x4a, 0xdf, 0x55, 0xd7, 0x69, 0xb9, 0xa2, 0x53, 0x2a, 0x67, 0x0e, 0x2c, 0xdd, 0x02, 0x78,
	0x55, 0xd2, 0x97, 0xdf, 0x77, 0x7e, 0x44, 0x70, 0x9e, 0xfa, 0x6b, 0x56, 0x95, 0xaf, 0xea, 0xaa,
	0xdc, 0x9b, 0x4d, 0xec, 0x17, 0x61, 0x0e, 0xdd, 0x09, 0x2b, 0x12, 0xfa, 0x05, 0xb8, 0x2a, 0x5f,
	0xd0, 0x4c, 0x9b, 0xf3, 0x31, 0x1e, 0xd9, 0x5e, 0x6e, 0xe1, 0x67, 0x4b, 0xfb, 0xd4, 0x9a, 0xcd,
	0x96, 0x2d, 0x72, 0x12, 0xb3, 0x54, 0x53, 0xa9, 0xfc, 0x95, 0x7e, 0xd3, 0x6c, 0x59, 0x67, 0x9d,
	0x38, 0x23, 0x7a, 0x0f, 0x1d, 0xb3, 0x56, 0xca, 0x02, 0x7e, 0x6c, 0x0b, 0xf8, 0x65, 0xfd, 0x8c,
	0xd5, 0xfe, 0xc7, 0x90, 0x88, 0xbc, 0xac, 0xa1, 0x37, 0xb0, 0xee, 0xd6, 0xac, 0xe2, 0xb1, 0x26,
	0xcb, 0xaf, 0x17, 0x8d, 0x85, 0x2c, 0x36, 0xe3, 0xb1, 0x9e, 0x96, 0x88, 0xf1, 0x48, 0x79, 0x52,
	0x6e, 0xdc, 0xca, 0xa3, 0xd8, 0x63, 0x29, 0x4f, 0x5c, 0x61, 0x54, 0x1e, 0xbf, 0x83, 0xa7, 0xb6,
	0xc3, 0x6f, 0xf5, 0x97, 0x7f, 0xe3, 0x78, 0xb5, 0x11, 0x89, 0xd7, 0x92, 0xfa, 0xc0, 0xfc, 0x1e,
	0x36, 0xb8, 0xa0, 0xd2, 0xf6, 0x0e, 0x31, 0xad, 0xcb, 0x73, 0x4d, 0x58, 0x46, 0x16, 0xca, 0xef,
	0xd8, 0x31, 0xb4, 0xbe, 0x84, 0x4f, 0x1d, 0x7a, 0x98, 0x1d, 0x5b, 0x49, 0x4c, 0x67, 0x94, 0xdf,
	0x7c, 0x7a, 0x57, 0x92, 0xaa, 0x95, 0x30, 0x48, 0x11, 0x96, 0xdf, 0x7b, 0x0f, 0x9d, 0xe2, 0xdf,
	0x2d, 0x12, 0xd1, 0xd8, 0xef, 0xde, 0x75, 0x1c, 0x3b, 0x78, 0x8f, 0xc6, 0x18, 0xc2, 0xe5, 0xf3,
	0xe6, 0x0f, 0xd0, 0xbb, 0xdd, 0x13, 0x66, 0x49, 0x5c, 0xd0, 0x6b, 0xdb, 0x8a, 0xab, 0xd8, 0x3c,
	0xa2, 0x75, 0x68, 0x5d, 0x06, 0x69, 0x4e, 0x8b, 0x0e, 0x70, 0x2f, 0xbf, 0x69, 0x7c, 0xf0, 0x06,
	0xff, 0x6e, 0x00, 0xe0, 0x3c, 0x2b, 0xcf, 0xb1, 0x0b, 0x1d, 0x2d, 0xcd, 0x94, 0x74, 0x2b, 0xcf,
	0xb3, 0x2b, 0xef, 0x46, 0xdd, 0x55, 0xe4, 0xe1, 0xa9, 0x61, 0xba, 0x45, 0x07, 0x7a, 0xf9, 0x8c,
	0x06, 0x6e, 0x02, 0x56, 0x8a, 0x35, 0xac, 0x62, 0x1d, 0x5d, 0x13, 0x6a, 0x04, 0xeb, 0xf7, 0x16,
	0xba, 0xab, 0xc8, 0xe7, 0xec, 0x4e, 0xf9, 0xbe, 0x83, 0x0d, 0x9e, 0x6b, 0x91, 0x6b, 0xd3, 0x97,
	0x9a, 0xd9, 0xbc, 0xd8, 0x8c, 0x95, 0xc3, 0xfe, 0x85, 0x83, 0xa7, 0x25, 0x6a, 0x93, 0xab, 0x4c,
	0x1d, 0xd8, 0xff, 0x76, 0x6f, 0x95, 0xe9, 0x8d, 0x3a, 0xd8, 0x33, 0x84, 0x65, 0x1d, 0x44, 0xb5,
	0xb7, 0x01, 0x06, 0xa8, 0x6e, 0x89, 0xd6, 0xa0, 0xfd, 0xe3, 0x09, 0x39, 0xc5, 0x3b, 0xe3, 0xfd,
	0xde, 0x23, 0x84, 0xa0, 0x3b, 0x3b, 0x99, 0x9c, 0xfe, 0x71, 0x07, 0xef, 0x17, 0x36, 0xcf, 0xd8,
	0x7e, 0xbf, 0x83, 0xf7, 0x6a, 0xb6, 0x06, 0xea, 0x02, 0x4c, 0xce, 0x8e, 0x8e, 0x8a, 0xf7, 0xe6,
	0xa7, 0x95, 0xf6, 0x4a, 0xaf, 0x35, 0xf8, 0xa7, 0x07, 0x1d, 0x9c, 0x67, 0xc7, 0x54, 0x07, 0x51,
	0xa0, 0x03, 0xf4, 0x1d, 0x40, 0xf5, 0x1f, 0xbc, 0x15, 0xbe, 0xb3, 0xfd, 0xa2, 0x7e, 0x4a, 0xb3,
	0x38, 0x66, 0x06, 0xc4, 0xab, 0xaa, 0x7c, 0x44, 0xef, 0x01, 0xaa, 0x9f, 0x12, 0x56, 0xe8, 0x5b,
	0x77, 0x33, 0x8b, 0xd2, 0x4a, 0x61, 0x0a, 0x67, 0x35, 0x2c, 0xdf, 0xd0, 0x47, 0xe8, 0xdd, 0x11,
	0xb2, 0x69, 0xa7, 0xcc, 0xfa, 0x9d, 0x16, 0x31, 0xae, 0xcf, 0xc4, 0x4d, 0x61, 0x77, 0x25, 0xf8,
	0x5c, 0x26, 0x75, 0xee, 0xf2, 0xa7, 0xc8, 0xee, 0x5a, 0x6d, 0x64, 0xab, 0xa9, 0xf7, 0xa7, 0x71,
	0xc2, 0xf4, 0x79, 0x3e, 0x1f, 0x86, 0x7c, 0x31, 0x8a, 0x24, 0xbd, 0xbe, 0x18, 0x55, 0x6e, 0xdf,
	0x28, 0x2a, 0x2f, 0x59, 0x96, 0x7c, 0x93, 0xf0, 0x91, 0xb8, 0x48, 0x46, 0x0f, 0xfd, 0xcc, 0xf8,
	0x8f, 0xe7, 0xcd, 0x1f, 0xdb, 0x97, 0x6f, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x82, 0xec,
	0x61, 0x8c, 0x0d, 0x00, 0x00,
}
