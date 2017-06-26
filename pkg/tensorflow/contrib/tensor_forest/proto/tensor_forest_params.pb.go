// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/contrib/tensor_forest/proto/tensor_forest_params.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow_decision_trees "github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/contrib/decision_trees/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Leaf models specify what is returned at inference time, and how it is
// stored in the decision_trees.Leaf protos.
type LeafModelType int32

const (
	LeafModelType_MODEL_DENSE_CLASSIFICATION           LeafModelType = 0
	LeafModelType_MODEL_SPARSE_CLASSIFICATION          LeafModelType = 1
	LeafModelType_MODEL_REGRESSION                     LeafModelType = 2
	LeafModelType_MODEL_SPARSE_OR_DENSE_CLASSIFICATION LeafModelType = 3
)

var LeafModelType_name = map[int32]string{
	0: "MODEL_DENSE_CLASSIFICATION",
	1: "MODEL_SPARSE_CLASSIFICATION",
	2: "MODEL_REGRESSION",
	3: "MODEL_SPARSE_OR_DENSE_CLASSIFICATION",
}
var LeafModelType_value = map[string]int32{
	"MODEL_DENSE_CLASSIFICATION":           0,
	"MODEL_SPARSE_CLASSIFICATION":          1,
	"MODEL_REGRESSION":                     2,
	"MODEL_SPARSE_OR_DENSE_CLASSIFICATION": 3,
}

func (x LeafModelType) String() string {
	return proto1.EnumName(LeafModelType_name, int32(x))
}
func (LeafModelType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Stats models generally specify information that is collected which is
// necessary to choose a split at a node. Specifically, they operate on
// a SplitCandidate::LeafStat proto.
type StatsModelType int32

const (
	StatsModelType_STATS_DENSE_GINI               StatsModelType = 0
	StatsModelType_STATS_SPARSE_GINI              StatsModelType = 1
	StatsModelType_STATS_LEAST_SQUARES_REGRESSION StatsModelType = 2
	StatsModelType_STATS_SPARSE_THEN_DENSE_GINI   StatsModelType = 3
)

var StatsModelType_name = map[int32]string{
	0: "STATS_DENSE_GINI",
	1: "STATS_SPARSE_GINI",
	2: "STATS_LEAST_SQUARES_REGRESSION",
	3: "STATS_SPARSE_THEN_DENSE_GINI",
}
var StatsModelType_value = map[string]int32{
	"STATS_DENSE_GINI":               0,
	"STATS_SPARSE_GINI":              1,
	"STATS_LEAST_SQUARES_REGRESSION": 2,
	"STATS_SPARSE_THEN_DENSE_GINI":   3,
}

func (x StatsModelType) String() string {
	return proto1.EnumName(StatsModelType_name, int32(x))
}
func (StatsModelType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// Allows selection of operations on the collection of split candidates.
// Basic infers right split stats from the leaf stats and each candidate's
// left stats.
type SplitCollectionType int32

const (
	SplitCollectionType_COLLECTION_BASIC        SplitCollectionType = 0
	SplitCollectionType_GRAPH_RUNNER_COLLECTION SplitCollectionType = 1
)

var SplitCollectionType_name = map[int32]string{
	0: "COLLECTION_BASIC",
	1: "GRAPH_RUNNER_COLLECTION",
}
var SplitCollectionType_value = map[string]int32{
	"COLLECTION_BASIC":        0,
	"GRAPH_RUNNER_COLLECTION": 1,
}

func (x SplitCollectionType) String() string {
	return proto1.EnumName(SplitCollectionType_name, int32(x))
}
func (SplitCollectionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// Pruning strategies define how candidates are pruned over time.
// SPLIT_PRUNE_HALF prunes the worst half of splits every prune_ever_samples,
// etc.  Note that prune_every_samples plays against the depth-dependent
// split_after_samples, so they should be set together.
type SplitPruningStrategyType int32

const (
	SplitPruningStrategyType_SPLIT_PRUNE_NONE       SplitPruningStrategyType = 0
	SplitPruningStrategyType_SPLIT_PRUNE_HALF       SplitPruningStrategyType = 1
	SplitPruningStrategyType_SPLIT_PRUNE_QUARTER    SplitPruningStrategyType = 2
	SplitPruningStrategyType_SPLIT_PRUNE_10_PERCENT SplitPruningStrategyType = 3
	// SPLIT_PRUNE_HOEFFDING prunes splits whose Gini impurity is worst than
	// the best split's by more than the Hoeffding bound.
	SplitPruningStrategyType_SPLIT_PRUNE_HOEFFDING SplitPruningStrategyType = 4
)

var SplitPruningStrategyType_name = map[int32]string{
	0: "SPLIT_PRUNE_NONE",
	1: "SPLIT_PRUNE_HALF",
	2: "SPLIT_PRUNE_QUARTER",
	3: "SPLIT_PRUNE_10_PERCENT",
	4: "SPLIT_PRUNE_HOEFFDING",
}
var SplitPruningStrategyType_value = map[string]int32{
	"SPLIT_PRUNE_NONE":       0,
	"SPLIT_PRUNE_HALF":       1,
	"SPLIT_PRUNE_QUARTER":    2,
	"SPLIT_PRUNE_10_PERCENT": 3,
	"SPLIT_PRUNE_HOEFFDING":  4,
}

func (x SplitPruningStrategyType) String() string {
	return proto1.EnumName(SplitPruningStrategyType_name, int32(x))
}
func (SplitPruningStrategyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// Finish strategies define when slots are considered finished.
// Basic requires at least split_after_samples, and doesn't allow slots to
// finish until the leaf has received more than one class. Hoeffding splits
// early after min_split_samples if one split is dominating the rest according
// to hoeffding bounds. Bootstrap does the same but compares gini's calculated
// with sampled smoothed counts.
type SplitFinishStrategyType int32

const (
	SplitFinishStrategyType_SPLIT_FINISH_BASIC              SplitFinishStrategyType = 0
	SplitFinishStrategyType_SPLIT_FINISH_DOMINATE_HOEFFDING SplitFinishStrategyType = 2
	SplitFinishStrategyType_SPLIT_FINISH_DOMINATE_BOOTSTRAP SplitFinishStrategyType = 3
)

var SplitFinishStrategyType_name = map[int32]string{
	0: "SPLIT_FINISH_BASIC",
	2: "SPLIT_FINISH_DOMINATE_HOEFFDING",
	3: "SPLIT_FINISH_DOMINATE_BOOTSTRAP",
}
var SplitFinishStrategyType_value = map[string]int32{
	"SPLIT_FINISH_BASIC":              0,
	"SPLIT_FINISH_DOMINATE_HOEFFDING": 2,
	"SPLIT_FINISH_DOMINATE_BOOTSTRAP": 3,
}

func (x SplitFinishStrategyType) String() string {
	return proto1.EnumName(SplitFinishStrategyType_name, int32(x))
}
func (SplitFinishStrategyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type SplitPruningConfig struct {
	PruneEverySamples *DepthDependentParam     `protobuf:"bytes,1,opt,name=prune_every_samples,json=pruneEverySamples" json:"prune_every_samples,omitempty"`
	Type              SplitPruningStrategyType `protobuf:"varint,2,opt,name=type,enum=tensorflow.tensorforest.SplitPruningStrategyType" json:"type,omitempty"`
}

func (m *SplitPruningConfig) Reset()                    { *m = SplitPruningConfig{} }
func (m *SplitPruningConfig) String() string            { return proto1.CompactTextString(m) }
func (*SplitPruningConfig) ProtoMessage()               {}
func (*SplitPruningConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SplitPruningConfig) GetPruneEverySamples() *DepthDependentParam {
	if m != nil {
		return m.PruneEverySamples
	}
	return nil
}

func (m *SplitPruningConfig) GetType() SplitPruningStrategyType {
	if m != nil {
		return m.Type
	}
	return SplitPruningStrategyType_SPLIT_PRUNE_NONE
}

type SplitFinishConfig struct {
	// Configure how often we check for finish, because some finish methods
	// are expensive to perform.
	CheckEverySteps *DepthDependentParam    `protobuf:"bytes,1,opt,name=check_every_steps,json=checkEverySteps" json:"check_every_steps,omitempty"`
	Type            SplitFinishStrategyType `protobuf:"varint,2,opt,name=type,enum=tensorflow.tensorforest.SplitFinishStrategyType" json:"type,omitempty"`
}

func (m *SplitFinishConfig) Reset()                    { *m = SplitFinishConfig{} }
func (m *SplitFinishConfig) String() string            { return proto1.CompactTextString(m) }
func (*SplitFinishConfig) ProtoMessage()               {}
func (*SplitFinishConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SplitFinishConfig) GetCheckEverySteps() *DepthDependentParam {
	if m != nil {
		return m.CheckEverySteps
	}
	return nil
}

func (m *SplitFinishConfig) GetType() SplitFinishStrategyType {
	if m != nil {
		return m.Type
	}
	return SplitFinishStrategyType_SPLIT_FINISH_BASIC
}

// A parameter that changes linearly with depth, with upper and lower bounds.
type LinearParam struct {
	Slope      float32 `protobuf:"fixed32,1,opt,name=slope" json:"slope,omitempty"`
	YIntercept float32 `protobuf:"fixed32,2,opt,name=y_intercept,json=yIntercept" json:"y_intercept,omitempty"`
	MinVal     float32 `protobuf:"fixed32,3,opt,name=min_val,json=minVal" json:"min_val,omitempty"`
	MaxVal     float32 `protobuf:"fixed32,4,opt,name=max_val,json=maxVal" json:"max_val,omitempty"`
}

func (m *LinearParam) Reset()                    { *m = LinearParam{} }
func (m *LinearParam) String() string            { return proto1.CompactTextString(m) }
func (*LinearParam) ProtoMessage()               {}
func (*LinearParam) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *LinearParam) GetSlope() float32 {
	if m != nil {
		return m.Slope
	}
	return 0
}

func (m *LinearParam) GetYIntercept() float32 {
	if m != nil {
		return m.YIntercept
	}
	return 0
}

func (m *LinearParam) GetMinVal() float32 {
	if m != nil {
		return m.MinVal
	}
	return 0
}

func (m *LinearParam) GetMaxVal() float32 {
	if m != nil {
		return m.MaxVal
	}
	return 0
}

// A parameter that changes expoentially with the form
//     f = c + mb^(k*d)
// where:
//  c: constant bias
//  b: base
//  m: multiplier
//  k: depth multiplier
//  d: depth
type ExponentialParam struct {
	Bias            float32 `protobuf:"fixed32,1,opt,name=bias" json:"bias,omitempty"`
	Base            float32 `protobuf:"fixed32,2,opt,name=base" json:"base,omitempty"`
	Multiplier      float32 `protobuf:"fixed32,3,opt,name=multiplier" json:"multiplier,omitempty"`
	DepthMultiplier float32 `protobuf:"fixed32,4,opt,name=depth_multiplier,json=depthMultiplier" json:"depth_multiplier,omitempty"`
}

func (m *ExponentialParam) Reset()                    { *m = ExponentialParam{} }
func (m *ExponentialParam) String() string            { return proto1.CompactTextString(m) }
func (*ExponentialParam) ProtoMessage()               {}
func (*ExponentialParam) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ExponentialParam) GetBias() float32 {
	if m != nil {
		return m.Bias
	}
	return 0
}

func (m *ExponentialParam) GetBase() float32 {
	if m != nil {
		return m.Base
	}
	return 0
}

func (m *ExponentialParam) GetMultiplier() float32 {
	if m != nil {
		return m.Multiplier
	}
	return 0
}

func (m *ExponentialParam) GetDepthMultiplier() float32 {
	if m != nil {
		return m.DepthMultiplier
	}
	return 0
}

// A parameter that is 'off' until depth >= a threshold, then is 'on'.
type ThresholdParam struct {
	OnValue   float32 `protobuf:"fixed32,1,opt,name=on_value,json=onValue" json:"on_value,omitempty"`
	OffValue  float32 `protobuf:"fixed32,2,opt,name=off_value,json=offValue" json:"off_value,omitempty"`
	Threshold float32 `protobuf:"fixed32,3,opt,name=threshold" json:"threshold,omitempty"`
}

func (m *ThresholdParam) Reset()                    { *m = ThresholdParam{} }
func (m *ThresholdParam) String() string            { return proto1.CompactTextString(m) }
func (*ThresholdParam) ProtoMessage()               {}
func (*ThresholdParam) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ThresholdParam) GetOnValue() float32 {
	if m != nil {
		return m.OnValue
	}
	return 0
}

func (m *ThresholdParam) GetOffValue() float32 {
	if m != nil {
		return m.OffValue
	}
	return 0
}

func (m *ThresholdParam) GetThreshold() float32 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

// A parameter that may change with node depth.
type DepthDependentParam struct {
	// Types that are valid to be assigned to ParamType:
	//	*DepthDependentParam_ConstantValue
	//	*DepthDependentParam_Linear
	//	*DepthDependentParam_Exponential
	//	*DepthDependentParam_Threshold
	ParamType isDepthDependentParam_ParamType `protobuf_oneof:"ParamType"`
}

func (m *DepthDependentParam) Reset()                    { *m = DepthDependentParam{} }
func (m *DepthDependentParam) String() string            { return proto1.CompactTextString(m) }
func (*DepthDependentParam) ProtoMessage()               {}
func (*DepthDependentParam) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type isDepthDependentParam_ParamType interface {
	isDepthDependentParam_ParamType()
}

type DepthDependentParam_ConstantValue struct {
	ConstantValue float32 `protobuf:"fixed32,1,opt,name=constant_value,json=constantValue,oneof"`
}
type DepthDependentParam_Linear struct {
	Linear *LinearParam `protobuf:"bytes,2,opt,name=linear,oneof"`
}
type DepthDependentParam_Exponential struct {
	Exponential *ExponentialParam `protobuf:"bytes,3,opt,name=exponential,oneof"`
}
type DepthDependentParam_Threshold struct {
	Threshold *ThresholdParam `protobuf:"bytes,4,opt,name=threshold,oneof"`
}

func (*DepthDependentParam_ConstantValue) isDepthDependentParam_ParamType() {}
func (*DepthDependentParam_Linear) isDepthDependentParam_ParamType()        {}
func (*DepthDependentParam_Exponential) isDepthDependentParam_ParamType()   {}
func (*DepthDependentParam_Threshold) isDepthDependentParam_ParamType()     {}

func (m *DepthDependentParam) GetParamType() isDepthDependentParam_ParamType {
	if m != nil {
		return m.ParamType
	}
	return nil
}

func (m *DepthDependentParam) GetConstantValue() float32 {
	if x, ok := m.GetParamType().(*DepthDependentParam_ConstantValue); ok {
		return x.ConstantValue
	}
	return 0
}

func (m *DepthDependentParam) GetLinear() *LinearParam {
	if x, ok := m.GetParamType().(*DepthDependentParam_Linear); ok {
		return x.Linear
	}
	return nil
}

func (m *DepthDependentParam) GetExponential() *ExponentialParam {
	if x, ok := m.GetParamType().(*DepthDependentParam_Exponential); ok {
		return x.Exponential
	}
	return nil
}

func (m *DepthDependentParam) GetThreshold() *ThresholdParam {
	if x, ok := m.GetParamType().(*DepthDependentParam_Threshold); ok {
		return x.Threshold
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DepthDependentParam) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _DepthDependentParam_OneofMarshaler, _DepthDependentParam_OneofUnmarshaler, _DepthDependentParam_OneofSizer, []interface{}{
		(*DepthDependentParam_ConstantValue)(nil),
		(*DepthDependentParam_Linear)(nil),
		(*DepthDependentParam_Exponential)(nil),
		(*DepthDependentParam_Threshold)(nil),
	}
}

func _DepthDependentParam_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*DepthDependentParam)
	// ParamType
	switch x := m.ParamType.(type) {
	case *DepthDependentParam_ConstantValue:
		b.EncodeVarint(1<<3 | proto1.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.ConstantValue)))
	case *DepthDependentParam_Linear:
		b.EncodeVarint(2<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Linear); err != nil {
			return err
		}
	case *DepthDependentParam_Exponential:
		b.EncodeVarint(3<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Exponential); err != nil {
			return err
		}
	case *DepthDependentParam_Threshold:
		b.EncodeVarint(4<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Threshold); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DepthDependentParam.ParamType has unexpected type %T", x)
	}
	return nil
}

func _DepthDependentParam_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*DepthDependentParam)
	switch tag {
	case 1: // ParamType.constant_value
		if wire != proto1.WireFixed32 {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.ParamType = &DepthDependentParam_ConstantValue{math.Float32frombits(uint32(x))}
		return true, err
	case 2: // ParamType.linear
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(LinearParam)
		err := b.DecodeMessage(msg)
		m.ParamType = &DepthDependentParam_Linear{msg}
		return true, err
	case 3: // ParamType.exponential
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(ExponentialParam)
		err := b.DecodeMessage(msg)
		m.ParamType = &DepthDependentParam_Exponential{msg}
		return true, err
	case 4: // ParamType.threshold
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(ThresholdParam)
		err := b.DecodeMessage(msg)
		m.ParamType = &DepthDependentParam_Threshold{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DepthDependentParam_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*DepthDependentParam)
	// ParamType
	switch x := m.ParamType.(type) {
	case *DepthDependentParam_ConstantValue:
		n += proto1.SizeVarint(1<<3 | proto1.WireFixed32)
		n += 4
	case *DepthDependentParam_Linear:
		s := proto1.Size(x.Linear)
		n += proto1.SizeVarint(2<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *DepthDependentParam_Exponential:
		s := proto1.Size(x.Exponential)
		n += proto1.SizeVarint(3<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *DepthDependentParam_Threshold:
		s := proto1.Size(x.Threshold)
		n += proto1.SizeVarint(4<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TensorForestParams struct {
	// ------------ Types that control training subsystems ------ //
	LeafType       LeafModelType       `protobuf:"varint,1,opt,name=leaf_type,json=leafType,enum=tensorflow.tensorforest.LeafModelType" json:"leaf_type,omitempty"`
	StatsType      StatsModelType      `protobuf:"varint,2,opt,name=stats_type,json=statsType,enum=tensorflow.tensorforest.StatsModelType" json:"stats_type,omitempty"`
	CollectionType SplitCollectionType `protobuf:"varint,3,opt,name=collection_type,json=collectionType,enum=tensorflow.tensorforest.SplitCollectionType" json:"collection_type,omitempty"`
	PruningType    *SplitPruningConfig `protobuf:"bytes,4,opt,name=pruning_type,json=pruningType" json:"pruning_type,omitempty"`
	FinishType     *SplitFinishConfig  `protobuf:"bytes,5,opt,name=finish_type,json=finishType" json:"finish_type,omitempty"`
	// --------- Parameters that can't change by definition --------------- //
	NumTrees           int32                                         `protobuf:"varint,6,opt,name=num_trees,json=numTrees" json:"num_trees,omitempty"`
	MaxNodes           int32                                         `protobuf:"varint,7,opt,name=max_nodes,json=maxNodes" json:"max_nodes,omitempty"`
	NumFeatures        int32                                         `protobuf:"varint,21,opt,name=num_features,json=numFeatures" json:"num_features,omitempty"`
	InequalityTestType tensorflow_decision_trees.InequalityTest_Type `protobuf:"varint,19,opt,name=inequality_test_type,json=inequalityTestType,enum=tensorflow.decision_trees.InequalityTest_Type" json:"inequality_test_type,omitempty"`
	// Some booleans controlling execution
	IsRegression          bool `protobuf:"varint,8,opt,name=is_regression,json=isRegression" json:"is_regression,omitempty"`
	DropFinalClass        bool `protobuf:"varint,9,opt,name=drop_final_class,json=dropFinalClass" json:"drop_final_class,omitempty"`
	CollateExamples       bool `protobuf:"varint,10,opt,name=collate_examples,json=collateExamples" json:"collate_examples,omitempty"`
	CheckpointStats       bool `protobuf:"varint,11,opt,name=checkpoint_stats,json=checkpointStats" json:"checkpoint_stats,omitempty"`
	UseRunningStatsMethod bool `protobuf:"varint,20,opt,name=use_running_stats_method,json=useRunningStatsMethod" json:"use_running_stats_method,omitempty"`
	// Number of classes (classification) or targets (regression)
	NumOutputs int32 `protobuf:"varint,12,opt,name=num_outputs,json=numOutputs" json:"num_outputs,omitempty"`
	// --------- Parameters that could be depth-dependent --------------- //
	NumSplitsToConsider *DepthDependentParam `protobuf:"bytes,13,opt,name=num_splits_to_consider,json=numSplitsToConsider" json:"num_splits_to_consider,omitempty"`
	SplitAfterSamples   *DepthDependentParam `protobuf:"bytes,14,opt,name=split_after_samples,json=splitAfterSamples" json:"split_after_samples,omitempty"`
	DominateFraction    *DepthDependentParam `protobuf:"bytes,15,opt,name=dominate_fraction,json=dominateFraction" json:"dominate_fraction,omitempty"`
	MinSplitSamples     *DepthDependentParam `protobuf:"bytes,18,opt,name=min_split_samples,json=minSplitSamples" json:"min_split_samples,omitempty"`
	// --------- Parameters for experimental features ---------------------- //
	GraphDir          string `protobuf:"bytes,16,opt,name=graph_dir,json=graphDir" json:"graph_dir,omitempty"`
	NumSelectFeatures int32  `protobuf:"varint,17,opt,name=num_select_features,json=numSelectFeatures" json:"num_select_features,omitempty"`
}

func (m *TensorForestParams) Reset()                    { *m = TensorForestParams{} }
func (m *TensorForestParams) String() string            { return proto1.CompactTextString(m) }
func (*TensorForestParams) ProtoMessage()               {}
func (*TensorForestParams) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *TensorForestParams) GetLeafType() LeafModelType {
	if m != nil {
		return m.LeafType
	}
	return LeafModelType_MODEL_DENSE_CLASSIFICATION
}

func (m *TensorForestParams) GetStatsType() StatsModelType {
	if m != nil {
		return m.StatsType
	}
	return StatsModelType_STATS_DENSE_GINI
}

func (m *TensorForestParams) GetCollectionType() SplitCollectionType {
	if m != nil {
		return m.CollectionType
	}
	return SplitCollectionType_COLLECTION_BASIC
}

func (m *TensorForestParams) GetPruningType() *SplitPruningConfig {
	if m != nil {
		return m.PruningType
	}
	return nil
}

func (m *TensorForestParams) GetFinishType() *SplitFinishConfig {
	if m != nil {
		return m.FinishType
	}
	return nil
}

func (m *TensorForestParams) GetNumTrees() int32 {
	if m != nil {
		return m.NumTrees
	}
	return 0
}

func (m *TensorForestParams) GetMaxNodes() int32 {
	if m != nil {
		return m.MaxNodes
	}
	return 0
}

func (m *TensorForestParams) GetNumFeatures() int32 {
	if m != nil {
		return m.NumFeatures
	}
	return 0
}

func (m *TensorForestParams) GetInequalityTestType() tensorflow_decision_trees.InequalityTest_Type {
	if m != nil {
		return m.InequalityTestType
	}
	return tensorflow_decision_trees.InequalityTest_LESS_OR_EQUAL
}

func (m *TensorForestParams) GetIsRegression() bool {
	if m != nil {
		return m.IsRegression
	}
	return false
}

func (m *TensorForestParams) GetDropFinalClass() bool {
	if m != nil {
		return m.DropFinalClass
	}
	return false
}

func (m *TensorForestParams) GetCollateExamples() bool {
	if m != nil {
		return m.CollateExamples
	}
	return false
}

func (m *TensorForestParams) GetCheckpointStats() bool {
	if m != nil {
		return m.CheckpointStats
	}
	return false
}

func (m *TensorForestParams) GetUseRunningStatsMethod() bool {
	if m != nil {
		return m.UseRunningStatsMethod
	}
	return false
}

func (m *TensorForestParams) GetNumOutputs() int32 {
	if m != nil {
		return m.NumOutputs
	}
	return 0
}

func (m *TensorForestParams) GetNumSplitsToConsider() *DepthDependentParam {
	if m != nil {
		return m.NumSplitsToConsider
	}
	return nil
}

func (m *TensorForestParams) GetSplitAfterSamples() *DepthDependentParam {
	if m != nil {
		return m.SplitAfterSamples
	}
	return nil
}

func (m *TensorForestParams) GetDominateFraction() *DepthDependentParam {
	if m != nil {
		return m.DominateFraction
	}
	return nil
}

func (m *TensorForestParams) GetMinSplitSamples() *DepthDependentParam {
	if m != nil {
		return m.MinSplitSamples
	}
	return nil
}

func (m *TensorForestParams) GetGraphDir() string {
	if m != nil {
		return m.GraphDir
	}
	return ""
}

func (m *TensorForestParams) GetNumSelectFeatures() int32 {
	if m != nil {
		return m.NumSelectFeatures
	}
	return 0
}

func init() {
	proto1.RegisterType((*SplitPruningConfig)(nil), "tensorflow.tensorforest.SplitPruningConfig")
	proto1.RegisterType((*SplitFinishConfig)(nil), "tensorflow.tensorforest.SplitFinishConfig")
	proto1.RegisterType((*LinearParam)(nil), "tensorflow.tensorforest.LinearParam")
	proto1.RegisterType((*ExponentialParam)(nil), "tensorflow.tensorforest.ExponentialParam")
	proto1.RegisterType((*ThresholdParam)(nil), "tensorflow.tensorforest.ThresholdParam")
	proto1.RegisterType((*DepthDependentParam)(nil), "tensorflow.tensorforest.DepthDependentParam")
	proto1.RegisterType((*TensorForestParams)(nil), "tensorflow.tensorforest.TensorForestParams")
	proto1.RegisterEnum("tensorflow.tensorforest.LeafModelType", LeafModelType_name, LeafModelType_value)
	proto1.RegisterEnum("tensorflow.tensorforest.StatsModelType", StatsModelType_name, StatsModelType_value)
	proto1.RegisterEnum("tensorflow.tensorforest.SplitCollectionType", SplitCollectionType_name, SplitCollectionType_value)
	proto1.RegisterEnum("tensorflow.tensorforest.SplitPruningStrategyType", SplitPruningStrategyType_name, SplitPruningStrategyType_value)
	proto1.RegisterEnum("tensorflow.tensorforest.SplitFinishStrategyType", SplitFinishStrategyType_name, SplitFinishStrategyType_value)
}

func init() {
	proto1.RegisterFile("tensorflow/contrib/tensor_forest/proto/tensor_forest_params.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 1378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0x5f, 0x53, 0xdb, 0x48,
	0x12, 0xb7, 0x80, 0x10, 0xdc, 0xe6, 0x8f, 0x18, 0x20, 0x28, 0x21, 0x95, 0x70, 0x24, 0x75, 0x21,
	0x5c, 0x62, 0x12, 0xee, 0xe1, 0xde, 0xae, 0xce, 0xd8, 0x32, 0x76, 0x9d, 0xb1, 0x1d, 0x49, 0xa4,
	0xee, 0xae, 0xae, 0x6a, 0x56, 0xd8, 0x63, 0x7b, 0x2a, 0xd2, 0x48, 0x2b, 0x8d, 0x58, 0x5c, 0xfb,
	0xb8, 0xb5, 0xef, 0xfb, 0x94, 0x2f, 0xb0, 0x1f, 0x61, 0x9f, 0xf6, 0xdb, 0x6d, 0x4d, 0x8f, 0x0c,
	0x76, 0x16, 0xa7, 0x52, 0x3c, 0xa1, 0xf9, 0x75, 0xf7, 0x6f, 0x7e, 0xd3, 0xd3, 0xd3, 0x6d, 0xa0,
	0x22, 0x99, 0x48, 0xa3, 0x64, 0x10, 0x44, 0x3f, 0x1c, 0xf7, 0x22, 0x21, 0x13, 0x7e, 0x79, 0xac,
	0x21, 0x3a, 0x88, 0x12, 0x96, 0xca, 0xe3, 0x38, 0x89, 0x64, 0x34, 0x8b, 0xd1, 0xd8, 0x4f, 0xfc,
	0x30, 0x2d, 0xa3, 0x89, 0xec, 0xde, 0x52, 0x94, 0xf3, 0x4f, 0xf4, 0x7a, 0xf2, 0xaf, 0x3b, 0xb8,
	0xfb, 0xac, 0xc7, 0x53, 0x1e, 0x09, 0x2a, 0x13, 0xc6, 0xd2, 0x9c, 0x7c, 0xc8, 0x04, 0x4b, 0x78,
	0x0f, 0x31, 0x1a, 0x46, 0x7d, 0x16, 0x68, 0xea, 0x83, 0xdf, 0x0d, 0x20, 0x6e, 0x1c, 0x70, 0xd9,
	0x4d, 0x32, 0xc1, 0xc5, 0xb0, 0x1a, 0x89, 0x01, 0x1f, 0x92, 0xff, 0xc3, 0x56, 0x9c, 0x64, 0x82,
	0x51, 0x76, 0xc5, 0x92, 0x31, 0x4d, 0xfd, 0x30, 0x0e, 0x58, 0x6a, 0x19, 0xfb, 0xc6, 0x61, 0xe9,
	0xe4, 0x4d, 0x79, 0x8e, 0x9e, 0x72, 0x8d, 0xc5, 0x72, 0x54, 0x63, 0x31, 0x13, 0x7d, 0x26, 0x64,
	0x57, 0x9d, 0xc1, 0xd9, 0x44, 0x22, 0x5b, 0xf1, 0xb8, 0x9a, 0x86, 0xd8, 0xb0, 0x24, 0xc7, 0x31,
	0xb3, 0x16, 0xf6, 0x8d, 0xc3, 0xf5, 0x93, 0xf7, 0x73, 0xe9, 0xa6, 0x85, 0xb9, 0x32, 0xf1, 0x25,
	0x1b, 0x8e, 0xbd, 0x71, 0xcc, 0x1c, 0x0c, 0x3f, 0xf8, 0xcd, 0x80, 0x4d, 0x74, 0xa9, 0x73, 0xc1,
	0xd3, 0x51, 0x2e, 0xfd, 0x3f, 0xb0, 0xd9, 0x1b, 0xb1, 0xde, 0xa7, 0x89, 0x74, 0xc9, 0xe2, 0xfb,
	0x09, 0xdf, 0x40, 0x1a, 0x2d, 0x5c, 0x91, 0x90, 0xda, 0x8c, 0xec, 0x77, 0x5f, 0x97, 0xad, 0x35,
	0xdd, 0xa1, 0xfa, 0x0a, 0x4a, 0x2d, 0x2e, 0x98, 0x9f, 0xe0, 0x2e, 0x64, 0x1b, 0x1e, 0xa4, 0x41,
	0x14, 0x33, 0x94, 0xb8, 0xe0, 0xe8, 0x05, 0x79, 0x0e, 0xa5, 0x31, 0xe5, 0x42, 0xb2, 0xa4, 0xc7,
	0x62, 0x89, 0x3b, 0x2e, 0x38, 0x30, 0x6e, 0x4e, 0x10, 0xb2, 0x0b, 0x0f, 0x43, 0x2e, 0xe8, 0x95,
	0x1f, 0x58, 0x8b, 0x68, 0x5c, 0x0e, 0xb9, 0xf8, 0xe8, 0x07, 0x68, 0xf0, 0xaf, 0xd1, 0xb0, 0x94,
	0x1b, 0xfc, 0xeb, 0x8f, 0x7e, 0x70, 0xf0, 0xb3, 0x01, 0xa6, 0x7d, 0x1d, 0x47, 0x82, 0x09, 0xc9,
	0xfd, 0x40, 0xef, 0x4e, 0x60, 0xe9, 0x92, 0xfb, 0x69, 0xbe, 0x39, 0x7e, 0x23, 0xe6, 0xa7, 0x2c,
	0xdf, 0x14, 0xbf, 0xc9, 0x33, 0x80, 0x30, 0x0b, 0x24, 0x8f, 0x03, 0xce, 0x92, 0x7c, 0xc7, 0x29,
	0x84, 0xbc, 0x06, 0xb3, 0xaf, 0x52, 0x48, 0xa7, 0xbc, 0xf4, 0xf6, 0x1b, 0x88, 0x9f, 0xdf, 0xc0,
	0x07, 0x03, 0x58, 0xf7, 0x46, 0x09, 0x4b, 0x47, 0x51, 0xd0, 0xd7, 0x22, 0x1e, 0xc3, 0x4a, 0x84,
	0x47, 0xc9, 0x26, 0x59, 0x78, 0x18, 0xa9, 0xb3, 0x64, 0x8c, 0xec, 0x41, 0x31, 0x1a, 0x0c, 0x72,
	0x9b, 0x16, 0xb4, 0x12, 0x0d, 0x06, 0xda, 0xf8, 0x14, 0x8a, 0x72, 0xc2, 0x94, 0x6b, 0xba, 0x05,
	0x0e, 0x7e, 0x5d, 0x80, 0xad, 0x3b, 0xae, 0x95, 0xbc, 0x82, 0xf5, 0x5e, 0x24, 0x52, 0xe9, 0x0b,
	0x39, 0xbd, 0x67, 0xa3, 0xe0, 0xac, 0x4d, 0x70, 0x4d, 0xff, 0x4f, 0x58, 0x0e, 0xf0, 0xa2, 0x70,
	0xe3, 0xd2, 0xc9, 0xcb, 0xb9, 0x17, 0x3e, 0x75, 0x9f, 0x8d, 0x82, 0x93, 0x47, 0x91, 0x73, 0x28,
	0xb1, 0xdb, 0x7c, 0xa3, 0xc0, 0xd2, 0xc9, 0xeb, 0xb9, 0x24, 0x5f, 0xde, 0x4d, 0xa3, 0xe0, 0x4c,
	0xc7, 0x93, 0xb3, 0xe9, 0xd3, 0x2e, 0x21, 0xd9, 0xab, 0xb9, 0x64, 0xb3, 0x19, 0x6e, 0x14, 0xa6,
	0x12, 0x73, 0x5a, 0x82, 0x22, 0xa2, 0xaa, 0x26, 0x0f, 0x3e, 0x03, 0x10, 0x0f, 0x23, 0xeb, 0x18,
	0x89, 0x96, 0x94, 0x54, 0xa1, 0x18, 0x30, 0x7f, 0x40, 0xb1, 0xde, 0x0d, 0xac, 0xf7, 0xbf, 0xce,
	0x3f, 0x3e, 0xf3, 0x07, 0xe7, 0xaa, 0xa7, 0x60, 0x95, 0xaf, 0xa8, 0x40, 0xf5, 0x45, 0xea, 0x00,
	0xa9, 0xf4, 0x65, 0x4a, 0xa7, 0x5e, 0xcd, 0x7c, 0xc9, 0xae, 0x72, 0xbd, 0xa5, 0x29, 0x62, 0x28,
	0xf2, 0x5c, 0xc0, 0x46, 0x2f, 0x0a, 0x02, 0xd6, 0x93, 0xd8, 0xd6, 0x14, 0xd9, 0x22, 0x92, 0xbd,
	0xf9, 0xfa, 0x13, 0xac, 0xde, 0x04, 0x21, 0xe3, 0x7a, 0x6f, 0x66, 0x4d, 0xda, 0xb0, 0x1a, 0xeb,
	0xde, 0xa2, 0x39, 0x75, 0x4e, 0xff, 0xf6, 0x4d, 0xdd, 0x48, 0xf7, 0x1a, 0xa7, 0x94, 0x13, 0x20,
	0xdf, 0xbf, 0xa1, 0x34, 0xc0, 0x47, 0xaf, 0xe9, 0x1e, 0x20, 0xdd, 0xd1, 0xb7, 0x74, 0x89, 0x9c,
	0x0d, 0x74, 0x38, 0x92, 0xed, 0x41, 0x51, 0x64, 0xa1, 0xee, 0xe1, 0xd6, 0xf2, 0xbe, 0x71, 0xf8,
	0xc0, 0x59, 0x11, 0x59, 0xe8, 0xa9, 0xb5, 0x32, 0xaa, 0x37, 0x2e, 0xa2, 0x3e, 0x4b, 0xad, 0x87,
	0xda, 0x18, 0xfa, 0xd7, 0x6d, 0xb5, 0x26, 0x7f, 0x81, 0x55, 0x15, 0x39, 0x60, 0xbe, 0xcc, 0x12,
	0x96, 0x5a, 0x3b, 0x68, 0x2f, 0x89, 0x2c, 0xac, 0xe7, 0x10, 0xf9, 0x0e, 0xb6, 0xb9, 0x60, 0xdf,
	0x67, 0x7e, 0xc0, 0xe5, 0x98, 0x4a, 0x35, 0x6f, 0x50, 0xf2, 0x16, 0x66, 0xb5, 0x3c, 0x2d, 0x79,
	0x76, 0x9a, 0x94, 0x9b, 0x37, 0x61, 0x1e, 0x96, 0x99, 0xca, 0x2b, 0xe1, 0x33, 0x20, 0xca, 0x7f,
	0x01, 0x6b, 0x3c, 0xa5, 0x09, 0x1b, 0x26, 0x2c, 0x55, 0xe1, 0xd6, 0xca, 0xbe, 0x71, 0xb8, 0xe2,
	0xac, 0xf2, 0xd4, 0xb9, 0xc1, 0xc8, 0x21, 0x98, 0xfd, 0x24, 0x8a, 0xe9, 0x80, 0x0b, 0x3f, 0xa0,
	0xbd, 0xc0, 0x4f, 0x53, 0xab, 0x88, 0x7e, 0xeb, 0x0a, 0xaf, 0x2b, 0xb8, 0xaa, 0x50, 0xd5, 0x5e,
	0xd4, 0xe5, 0xf9, 0x92, 0x51, 0x76, 0x9d, 0xcf, 0x22, 0x40, 0xcf, 0x8d, 0x1c, 0xb7, 0x73, 0x18,
	0x5d, 0x55, 0xdf, 0x8e, 0x23, 0x2e, 0x24, 0xc5, 0x22, 0xb2, 0x4a, 0xb9, 0xeb, 0x0d, 0x8e, 0xb5,
	0x46, 0xfe, 0x01, 0x56, 0x96, 0x32, 0x9a, 0x64, 0x02, 0x8b, 0x40, 0xd7, 0x6a, 0xc8, 0xe4, 0x28,
	0xea, 0x5b, 0xdb, 0x18, 0xb2, 0x93, 0xa5, 0xcc, 0xd1, 0x66, 0x5d, 0x9e, 0x68, 0x54, 0xdd, 0x59,
	0xa5, 0x38, 0xca, 0x64, 0x9c, 0xc9, 0xd4, 0x5a, 0xc5, 0x0c, 0x83, 0xc8, 0xc2, 0x8e, 0x46, 0x88,
	0x0f, 0x8f, 0x94, 0x43, 0xaa, 0xae, 0x38, 0xa5, 0x32, 0xa2, 0xaa, 0xb3, 0xf0, 0x3e, 0x4b, 0xac,
	0xb5, 0x7b, 0x0c, 0xa2, 0x2d, 0x91, 0x85, 0x58, 0x2d, 0xa9, 0x17, 0x55, 0x73, 0x22, 0x35, 0xa1,
	0x91, 0x9e, 0xfa, 0x03, 0xc9, 0x92, 0x9b, 0x09, 0xbd, 0x7e, 0x9f, 0x09, 0x8d, 0x44, 0x15, 0xc5,
	0x33, 0x99, 0xd0, 0xff, 0x85, 0xcd, 0x7e, 0x14, 0x72, 0xa1, 0x32, 0x3e, 0x48, 0x7c, 0x7c, 0x34,
	0xd6, 0xc6, 0x3d, 0xb8, 0xcd, 0x09, 0x4d, 0x3d, 0x67, 0x51, 0xf3, 0x59, 0x4d, 0x2e, 0x2d, 0x7e,
	0x22, 0x9b, 0xdc, 0x67, 0x3e, 0x87, 0x5c, 0x60, 0x5a, 0x26, 0xa2, 0xf7, 0xa0, 0x38, 0x4c, 0xfc,
	0x78, 0x44, 0xfb, 0x3c, 0xb1, 0xcc, 0x7d, 0xe3, 0xb0, 0xe8, 0xac, 0x20, 0x50, 0xe3, 0x09, 0x29,
	0xc3, 0x16, 0x5e, 0x09, 0x53, 0x2d, 0xe0, 0xf6, 0x75, 0x6c, 0xe2, 0xdd, 0x6d, 0xaa, 0x0c, 0xa3,
	0x65, 0xf2, 0x46, 0x8e, 0x7e, 0x31, 0x60, 0x6d, 0xa6, 0xb1, 0x91, 0x67, 0xf0, 0xe4, 0xbc, 0x53,
	0xb3, 0x5b, 0xb4, 0x66, 0xb7, 0x5d, 0x9b, 0x56, 0x5b, 0x15, 0xd7, 0x6d, 0xd6, 0x9b, 0xd5, 0x8a,
	0xd7, 0xec, 0xb4, 0xcd, 0x02, 0x79, 0x0e, 0x7b, 0xda, 0xee, 0x76, 0x2b, 0xce, 0x9f, 0x1d, 0x0c,
	0xb2, 0x0d, 0xa6, 0x76, 0x70, 0xec, 0x33, 0xc7, 0x76, 0x5d, 0x85, 0x2e, 0x90, 0x43, 0x78, 0x39,
	0x13, 0xd6, 0x71, 0xee, 0xde, 0x60, 0xf1, 0xe8, 0x27, 0x03, 0xd6, 0x67, 0xbb, 0xa4, 0xa2, 0x74,
	0xbd, 0x8a, 0xe7, 0xe6, 0x21, 0x67, 0xcd, 0x76, 0xd3, 0x2c, 0x90, 0x1d, 0xd8, 0xd4, 0x68, 0x4e,
	0x89, 0xb0, 0x41, 0x0e, 0xe0, 0x99, 0x86, 0x5b, 0x76, 0xc5, 0xf5, 0xa8, 0xfb, 0xe1, 0xa2, 0xe2,
	0xd8, 0xee, 0xac, 0x9a, 0x7d, 0x78, 0x3a, 0x13, 0xea, 0x35, 0xec, 0xf6, 0x34, 0xf9, 0xe2, 0x51,
	0x03, 0xb6, 0xee, 0xe8, 0xae, 0x4a, 0x49, 0xb5, 0xd3, 0x6a, 0xd9, 0x55, 0x25, 0x96, 0x9e, 0x56,
	0xdc, 0x66, 0xd5, 0x2c, 0x90, 0x3d, 0xd8, 0x3d, 0x73, 0x2a, 0xdd, 0x06, 0x75, 0x2e, 0xda, 0x6d,
	0xdb, 0xa1, 0xb7, 0x2e, 0xa6, 0x71, 0xf4, 0xd9, 0x00, 0x6b, 0xde, 0x4f, 0x3c, 0x3c, 0x59, 0xb7,
	0xd5, 0xf4, 0x68, 0xd7, 0xb9, 0x68, 0xdb, 0xb4, 0xdd, 0x69, 0xdb, 0x66, 0xe1, 0x4b, 0xb4, 0x51,
	0x69, 0xd5, 0x4d, 0x83, 0xec, 0xc2, 0xd6, 0x34, 0xaa, 0xce, 0xe5, 0xd9, 0x8e, 0xb9, 0x40, 0x9e,
	0xc0, 0xa3, 0x69, 0xc3, 0xfb, 0x77, 0xb4, 0x6b, 0x3b, 0x55, 0xbb, 0xed, 0x99, 0x8b, 0xe4, 0x31,
	0xec, 0xcc, 0x50, 0x75, 0xec, 0x7a, 0xbd, 0xd6, 0x6c, 0x9f, 0x99, 0x4b, 0x47, 0x3f, 0xc2, 0xee,
	0x9c, 0xdf, 0x70, 0xe4, 0x11, 0x10, 0x1d, 0x55, 0x6f, 0xb6, 0x9b, 0x6e, 0xe3, 0xe6, 0xa0, 0x2f,
	0xe0, 0xf9, 0x0c, 0x5e, 0xeb, 0x9c, 0x37, 0xdb, 0x15, 0x6f, 0x9a, 0x77, 0x61, 0xbe, 0xd3, 0x69,
	0xa7, 0xe3, 0xb9, 0x9e, 0x53, 0xe9, 0x9a, 0x8b, 0xa7, 0xee, 0xff, 0x3e, 0x0c, 0xb9, 0x1c, 0x65,
	0x97, 0xe5, 0x5e, 0x14, 0x1e, 0xf7, 0x13, 0x36, 0xfe, 0x74, 0x7c, 0xfb, 0x2c, 0xde, 0xa6, 0x2c,
	0xb9, 0xe2, 0x62, 0xf8, 0x76, 0x18, 0x1d, 0xc7, 0x9f, 0x86, 0xc7, 0xdf, 0xf6, 0xcf, 0xc5, 0xe5,
	0x32, 0xfe, 0xf9, 0xfb, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x71, 0xb0, 0xfb, 0x5e, 0x8d, 0x0c,
	0x00, 0x00,
}
