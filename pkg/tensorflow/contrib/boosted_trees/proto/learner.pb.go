// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/contrib/boosted_trees/proto/learner.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	tensorflow/contrib/boosted_trees/proto/learner.proto
	tensorflow/contrib/boosted_trees/proto/quantiles.proto
	tensorflow/contrib/boosted_trees/proto/tree_config.proto

It has these top-level messages:
	TreeRegularizationConfig
	TreeConstraintsConfig
	LearningRateConfig
	LearningRateFixedConfig
	LearningRateLineSearchConfig
	AveragingConfig
	LearningRateDropoutDrivenConfig
	LearnerConfig
	QuantileConfig
	QuantileEntry
	QuantileSummaryState
	QuantileStreamState
	TreeNode
	TreeNodeMetadata
	Leaf
	Vector
	SparseVector
	DenseFloatBinarySplit
	SparseFloatBinarySplitDefaultLeft
	SparseFloatBinarySplitDefaultRight
	CategoricalIdBinarySplit
	CategoricalIdSetMembershipBinarySplit
	DecisionTreeConfig
	DecisionTreeMetadata
	GrowingMetadata
	DecisionTreeEnsembleConfig
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type LearnerConfig_PruningMode int32

const (
	LearnerConfig_PRE_PRUNE  LearnerConfig_PruningMode = 0
	LearnerConfig_POST_PRUNE LearnerConfig_PruningMode = 1
)

var LearnerConfig_PruningMode_name = map[int32]string{
	0: "PRE_PRUNE",
	1: "POST_PRUNE",
}
var LearnerConfig_PruningMode_value = map[string]int32{
	"PRE_PRUNE":  0,
	"POST_PRUNE": 1,
}

func (x LearnerConfig_PruningMode) String() string {
	return proto1.EnumName(LearnerConfig_PruningMode_name, int32(x))
}
func (LearnerConfig_PruningMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

type LearnerConfig_GrowingMode int32

const (
	LearnerConfig_WHOLE_TREE LearnerConfig_GrowingMode = 0
	// Layer by layer is only supported by the batch learner.
	LearnerConfig_LAYER_BY_LAYER LearnerConfig_GrowingMode = 1
)

var LearnerConfig_GrowingMode_name = map[int32]string{
	0: "WHOLE_TREE",
	1: "LAYER_BY_LAYER",
}
var LearnerConfig_GrowingMode_value = map[string]int32{
	"WHOLE_TREE":     0,
	"LAYER_BY_LAYER": 1,
}

func (x LearnerConfig_GrowingMode) String() string {
	return proto1.EnumName(LearnerConfig_GrowingMode_name, int32(x))
}
func (LearnerConfig_GrowingMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 1} }

type LearnerConfig_MultiClassStrategy int32

const (
	LearnerConfig_TREE_PER_CLASS   LearnerConfig_MultiClassStrategy = 0
	LearnerConfig_FULL_HESSIAN     LearnerConfig_MultiClassStrategy = 1
	LearnerConfig_DIAGONAL_HESSIAN LearnerConfig_MultiClassStrategy = 2
)

var LearnerConfig_MultiClassStrategy_name = map[int32]string{
	0: "TREE_PER_CLASS",
	1: "FULL_HESSIAN",
	2: "DIAGONAL_HESSIAN",
}
var LearnerConfig_MultiClassStrategy_value = map[string]int32{
	"TREE_PER_CLASS":   0,
	"FULL_HESSIAN":     1,
	"DIAGONAL_HESSIAN": 2,
}

func (x LearnerConfig_MultiClassStrategy) String() string {
	return proto1.EnumName(LearnerConfig_MultiClassStrategy_name, int32(x))
}
func (LearnerConfig_MultiClassStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 2}
}

// Tree regularization config.
type TreeRegularizationConfig struct {
	// Classic L1/L2.
	L1 float32 `protobuf:"fixed32,1,opt,name=l1" json:"l1,omitempty"`
	L2 float32 `protobuf:"fixed32,2,opt,name=l2" json:"l2,omitempty"`
	// Tree complexity penalizes overall model complexity effectively
	// limiting how deep the tree can grow in regions with small gain.
	TreeComplexity float32 `protobuf:"fixed32,3,opt,name=tree_complexity,json=treeComplexity" json:"tree_complexity,omitempty"`
}

func (m *TreeRegularizationConfig) Reset()                    { *m = TreeRegularizationConfig{} }
func (m *TreeRegularizationConfig) String() string            { return proto1.CompactTextString(m) }
func (*TreeRegularizationConfig) ProtoMessage()               {}
func (*TreeRegularizationConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TreeRegularizationConfig) GetL1() float32 {
	if m != nil {
		return m.L1
	}
	return 0
}

func (m *TreeRegularizationConfig) GetL2() float32 {
	if m != nil {
		return m.L2
	}
	return 0
}

func (m *TreeRegularizationConfig) GetTreeComplexity() float32 {
	if m != nil {
		return m.TreeComplexity
	}
	return 0
}

// Tree constraints config.
type TreeConstraintsConfig struct {
	// Maximum depth of the trees.
	MaxTreeDepth uint32 `protobuf:"varint,1,opt,name=max_tree_depth,json=maxTreeDepth" json:"max_tree_depth,omitempty"`
	// Min hessian weight per node.
	MinNodeWeight float32 `protobuf:"fixed32,2,opt,name=min_node_weight,json=minNodeWeight" json:"min_node_weight,omitempty"`
}

func (m *TreeConstraintsConfig) Reset()                    { *m = TreeConstraintsConfig{} }
func (m *TreeConstraintsConfig) String() string            { return proto1.CompactTextString(m) }
func (*TreeConstraintsConfig) ProtoMessage()               {}
func (*TreeConstraintsConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TreeConstraintsConfig) GetMaxTreeDepth() uint32 {
	if m != nil {
		return m.MaxTreeDepth
	}
	return 0
}

func (m *TreeConstraintsConfig) GetMinNodeWeight() float32 {
	if m != nil {
		return m.MinNodeWeight
	}
	return 0
}

// LearningRateConfig describes all supported learning rate tuners.
type LearningRateConfig struct {
	// Types that are valid to be assigned to Tuner:
	//	*LearningRateConfig_Fixed
	//	*LearningRateConfig_Dropout
	//	*LearningRateConfig_LineSearch
	Tuner isLearningRateConfig_Tuner `protobuf_oneof:"tuner"`
}

func (m *LearningRateConfig) Reset()                    { *m = LearningRateConfig{} }
func (m *LearningRateConfig) String() string            { return proto1.CompactTextString(m) }
func (*LearningRateConfig) ProtoMessage()               {}
func (*LearningRateConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isLearningRateConfig_Tuner interface {
	isLearningRateConfig_Tuner()
}

type LearningRateConfig_Fixed struct {
	Fixed *LearningRateFixedConfig `protobuf:"bytes,1,opt,name=fixed,oneof"`
}
type LearningRateConfig_Dropout struct {
	Dropout *LearningRateDropoutDrivenConfig `protobuf:"bytes,2,opt,name=dropout,oneof"`
}
type LearningRateConfig_LineSearch struct {
	LineSearch *LearningRateLineSearchConfig `protobuf:"bytes,3,opt,name=line_search,json=lineSearch,oneof"`
}

func (*LearningRateConfig_Fixed) isLearningRateConfig_Tuner()      {}
func (*LearningRateConfig_Dropout) isLearningRateConfig_Tuner()    {}
func (*LearningRateConfig_LineSearch) isLearningRateConfig_Tuner() {}

func (m *LearningRateConfig) GetTuner() isLearningRateConfig_Tuner {
	if m != nil {
		return m.Tuner
	}
	return nil
}

func (m *LearningRateConfig) GetFixed() *LearningRateFixedConfig {
	if x, ok := m.GetTuner().(*LearningRateConfig_Fixed); ok {
		return x.Fixed
	}
	return nil
}

func (m *LearningRateConfig) GetDropout() *LearningRateDropoutDrivenConfig {
	if x, ok := m.GetTuner().(*LearningRateConfig_Dropout); ok {
		return x.Dropout
	}
	return nil
}

func (m *LearningRateConfig) GetLineSearch() *LearningRateLineSearchConfig {
	if x, ok := m.GetTuner().(*LearningRateConfig_LineSearch); ok {
		return x.LineSearch
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LearningRateConfig) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _LearningRateConfig_OneofMarshaler, _LearningRateConfig_OneofUnmarshaler, _LearningRateConfig_OneofSizer, []interface{}{
		(*LearningRateConfig_Fixed)(nil),
		(*LearningRateConfig_Dropout)(nil),
		(*LearningRateConfig_LineSearch)(nil),
	}
}

func _LearningRateConfig_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*LearningRateConfig)
	// tuner
	switch x := m.Tuner.(type) {
	case *LearningRateConfig_Fixed:
		b.EncodeVarint(1<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Fixed); err != nil {
			return err
		}
	case *LearningRateConfig_Dropout:
		b.EncodeVarint(2<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Dropout); err != nil {
			return err
		}
	case *LearningRateConfig_LineSearch:
		b.EncodeVarint(3<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.LineSearch); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LearningRateConfig.Tuner has unexpected type %T", x)
	}
	return nil
}

func _LearningRateConfig_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*LearningRateConfig)
	switch tag {
	case 1: // tuner.fixed
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(LearningRateFixedConfig)
		err := b.DecodeMessage(msg)
		m.Tuner = &LearningRateConfig_Fixed{msg}
		return true, err
	case 2: // tuner.dropout
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(LearningRateDropoutDrivenConfig)
		err := b.DecodeMessage(msg)
		m.Tuner = &LearningRateConfig_Dropout{msg}
		return true, err
	case 3: // tuner.line_search
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(LearningRateLineSearchConfig)
		err := b.DecodeMessage(msg)
		m.Tuner = &LearningRateConfig_LineSearch{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LearningRateConfig_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*LearningRateConfig)
	// tuner
	switch x := m.Tuner.(type) {
	case *LearningRateConfig_Fixed:
		s := proto1.Size(x.Fixed)
		n += proto1.SizeVarint(1<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *LearningRateConfig_Dropout:
		s := proto1.Size(x.Dropout)
		n += proto1.SizeVarint(2<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *LearningRateConfig_LineSearch:
		s := proto1.Size(x.LineSearch)
		n += proto1.SizeVarint(3<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Config for a fixed learning rate.
type LearningRateFixedConfig struct {
	LearningRate float32 `protobuf:"fixed32,1,opt,name=learning_rate,json=learningRate" json:"learning_rate,omitempty"`
}

func (m *LearningRateFixedConfig) Reset()                    { *m = LearningRateFixedConfig{} }
func (m *LearningRateFixedConfig) String() string            { return proto1.CompactTextString(m) }
func (*LearningRateFixedConfig) ProtoMessage()               {}
func (*LearningRateFixedConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LearningRateFixedConfig) GetLearningRate() float32 {
	if m != nil {
		return m.LearningRate
	}
	return 0
}

// Config for a tuned learning rate.
type LearningRateLineSearchConfig struct {
	// Max learning rate. Must be strictly positive.
	MaxLearningRate float32 `protobuf:"fixed32,1,opt,name=max_learning_rate,json=maxLearningRate" json:"max_learning_rate,omitempty"`
	// Number of learning rate values to consider between [0, max_learning_rate).
	NumSteps int32 `protobuf:"varint,2,opt,name=num_steps,json=numSteps" json:"num_steps,omitempty"`
}

func (m *LearningRateLineSearchConfig) Reset()                    { *m = LearningRateLineSearchConfig{} }
func (m *LearningRateLineSearchConfig) String() string            { return proto1.CompactTextString(m) }
func (*LearningRateLineSearchConfig) ProtoMessage()               {}
func (*LearningRateLineSearchConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LearningRateLineSearchConfig) GetMaxLearningRate() float32 {
	if m != nil {
		return m.MaxLearningRate
	}
	return 0
}

func (m *LearningRateLineSearchConfig) GetNumSteps() int32 {
	if m != nil {
		return m.NumSteps
	}
	return 0
}

// When we have a sequence of trees 1, 2, 3 ... n, these essentially represent
// weights updates in functional space, and thus we can use averaging of weight
// updates to achieve better performance. For example, we can say that our final
// ensemble will be an average of ensembles of tree 1, and ensemble of tree 1
// and tree 2 etc .. ensemble of all trees.
// Note that this averaging will apply ONLY DURING PREDICTION. The training
// stays the same.
type AveragingConfig struct {
	// Types that are valid to be assigned to Config:
	//	*AveragingConfig_AverageLastNTrees
	//	*AveragingConfig_AverageLastPercentTrees
	Config isAveragingConfig_Config `protobuf_oneof:"config"`
}

func (m *AveragingConfig) Reset()                    { *m = AveragingConfig{} }
func (m *AveragingConfig) String() string            { return proto1.CompactTextString(m) }
func (*AveragingConfig) ProtoMessage()               {}
func (*AveragingConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type isAveragingConfig_Config interface {
	isAveragingConfig_Config()
}

type AveragingConfig_AverageLastNTrees struct {
	AverageLastNTrees float32 `protobuf:"fixed32,1,opt,name=average_last_n_trees,json=averageLastNTrees,oneof"`
}
type AveragingConfig_AverageLastPercentTrees struct {
	AverageLastPercentTrees float32 `protobuf:"fixed32,2,opt,name=average_last_percent_trees,json=averageLastPercentTrees,oneof"`
}

func (*AveragingConfig_AverageLastNTrees) isAveragingConfig_Config()       {}
func (*AveragingConfig_AverageLastPercentTrees) isAveragingConfig_Config() {}

func (m *AveragingConfig) GetConfig() isAveragingConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *AveragingConfig) GetAverageLastNTrees() float32 {
	if x, ok := m.GetConfig().(*AveragingConfig_AverageLastNTrees); ok {
		return x.AverageLastNTrees
	}
	return 0
}

func (m *AveragingConfig) GetAverageLastPercentTrees() float32 {
	if x, ok := m.GetConfig().(*AveragingConfig_AverageLastPercentTrees); ok {
		return x.AverageLastPercentTrees
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AveragingConfig) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _AveragingConfig_OneofMarshaler, _AveragingConfig_OneofUnmarshaler, _AveragingConfig_OneofSizer, []interface{}{
		(*AveragingConfig_AverageLastNTrees)(nil),
		(*AveragingConfig_AverageLastPercentTrees)(nil),
	}
}

func _AveragingConfig_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*AveragingConfig)
	// config
	switch x := m.Config.(type) {
	case *AveragingConfig_AverageLastNTrees:
		b.EncodeVarint(1<<3 | proto1.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.AverageLastNTrees)))
	case *AveragingConfig_AverageLastPercentTrees:
		b.EncodeVarint(2<<3 | proto1.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.AverageLastPercentTrees)))
	case nil:
	default:
		return fmt.Errorf("AveragingConfig.Config has unexpected type %T", x)
	}
	return nil
}

func _AveragingConfig_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*AveragingConfig)
	switch tag {
	case 1: // config.average_last_n_trees
		if wire != proto1.WireFixed32 {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Config = &AveragingConfig_AverageLastNTrees{math.Float32frombits(uint32(x))}
		return true, err
	case 2: // config.average_last_percent_trees
		if wire != proto1.WireFixed32 {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Config = &AveragingConfig_AverageLastPercentTrees{math.Float32frombits(uint32(x))}
		return true, err
	default:
		return false, nil
	}
}

func _AveragingConfig_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*AveragingConfig)
	// config
	switch x := m.Config.(type) {
	case *AveragingConfig_AverageLastNTrees:
		n += proto1.SizeVarint(1<<3 | proto1.WireFixed32)
		n += 4
	case *AveragingConfig_AverageLastPercentTrees:
		n += proto1.SizeVarint(2<<3 | proto1.WireFixed32)
		n += 4
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LearningRateDropoutDrivenConfig struct {
	// Probability of dropping each tree in an existing so far ensemble.
	DropoutProbability float32 `protobuf:"fixed32,1,opt,name=dropout_probability,json=dropoutProbability" json:"dropout_probability,omitempty"`
	// When trees are built after dropout happen, they don't "advance" to the
	// optimal solution, they just rearrange the path. However you can still
	// choose to skip dropout periodically, to allow a new tree that "advances"
	// to be added.
	// For example, if running for 200 steps with probability of dropout 1/100,
	// you would expect the dropout to start happening for sure for all iterations
	// after 100. However you can add probability_of_skipping_dropout of 0.1, this
	// way iterations 100-200 will include approx 90 iterations of dropout and 10
	// iterations of normal steps.Set it to 0 if you want just keep building
	// the refinement trees after dropout kicks in.
	ProbabilityOfSkippingDropout float32 `protobuf:"fixed32,2,opt,name=probability_of_skipping_dropout,json=probabilityOfSkippingDropout" json:"probability_of_skipping_dropout,omitempty"`
	// Between 0 and 1.
	LearningRate float32 `protobuf:"fixed32,3,opt,name=learning_rate,json=learningRate" json:"learning_rate,omitempty"`
}

func (m *LearningRateDropoutDrivenConfig) Reset()                    { *m = LearningRateDropoutDrivenConfig{} }
func (m *LearningRateDropoutDrivenConfig) String() string            { return proto1.CompactTextString(m) }
func (*LearningRateDropoutDrivenConfig) ProtoMessage()               {}
func (*LearningRateDropoutDrivenConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LearningRateDropoutDrivenConfig) GetDropoutProbability() float32 {
	if m != nil {
		return m.DropoutProbability
	}
	return 0
}

func (m *LearningRateDropoutDrivenConfig) GetProbabilityOfSkippingDropout() float32 {
	if m != nil {
		return m.ProbabilityOfSkippingDropout
	}
	return 0
}

func (m *LearningRateDropoutDrivenConfig) GetLearningRate() float32 {
	if m != nil {
		return m.LearningRate
	}
	return 0
}

type LearnerConfig struct {
	// Number of classes.
	NumClasses uint32 `protobuf:"varint,1,opt,name=num_classes,json=numClasses" json:"num_classes,omitempty"`
	// Fraction of features to consider in each tree sampled randomly
	// from all available features.
	//
	// Types that are valid to be assigned to FeatureFraction:
	//	*LearnerConfig_FeatureFractionPerTree
	//	*LearnerConfig_FeatureFractionPerLevel
	FeatureFraction isLearnerConfig_FeatureFraction `protobuf_oneof:"feature_fraction"`
	// Regularization.
	Regularization *TreeRegularizationConfig `protobuf:"bytes,4,opt,name=regularization" json:"regularization,omitempty"`
	// Constraints.
	Constraints *TreeConstraintsConfig `protobuf:"bytes,5,opt,name=constraints" json:"constraints,omitempty"`
	// Pruning.
	PruningMode LearnerConfig_PruningMode `protobuf:"varint,8,opt,name=pruning_mode,json=pruningMode,enum=tensorflow.boosted_trees.learner.LearnerConfig_PruningMode" json:"pruning_mode,omitempty"`
	// Growing Mode.
	GrowingMode LearnerConfig_GrowingMode `protobuf:"varint,9,opt,name=growing_mode,json=growingMode,enum=tensorflow.boosted_trees.learner.LearnerConfig_GrowingMode" json:"growing_mode,omitempty"`
	// Learning rate.
	LearningRateTuner *LearningRateConfig `protobuf:"bytes,6,opt,name=learning_rate_tuner,json=learningRateTuner" json:"learning_rate_tuner,omitempty"`
	// Multi-class strategy.
	MultiClassStrategy LearnerConfig_MultiClassStrategy `protobuf:"varint,10,opt,name=multi_class_strategy,json=multiClassStrategy,enum=tensorflow.boosted_trees.learner.LearnerConfig_MultiClassStrategy" json:"multi_class_strategy,omitempty"`
	// If you want to average the ensembles (for regularization), provide the
	// config below.
	AveragingConfig *AveragingConfig `protobuf:"bytes,11,opt,name=averaging_config,json=averagingConfig" json:"averaging_config,omitempty"`
}

func (m *LearnerConfig) Reset()                    { *m = LearnerConfig{} }
func (m *LearnerConfig) String() string            { return proto1.CompactTextString(m) }
func (*LearnerConfig) ProtoMessage()               {}
func (*LearnerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isLearnerConfig_FeatureFraction interface {
	isLearnerConfig_FeatureFraction()
}

type LearnerConfig_FeatureFractionPerTree struct {
	FeatureFractionPerTree float32 `protobuf:"fixed32,2,opt,name=feature_fraction_per_tree,json=featureFractionPerTree,oneof"`
}
type LearnerConfig_FeatureFractionPerLevel struct {
	FeatureFractionPerLevel float32 `protobuf:"fixed32,3,opt,name=feature_fraction_per_level,json=featureFractionPerLevel,oneof"`
}

func (*LearnerConfig_FeatureFractionPerTree) isLearnerConfig_FeatureFraction()  {}
func (*LearnerConfig_FeatureFractionPerLevel) isLearnerConfig_FeatureFraction() {}

func (m *LearnerConfig) GetFeatureFraction() isLearnerConfig_FeatureFraction {
	if m != nil {
		return m.FeatureFraction
	}
	return nil
}

func (m *LearnerConfig) GetNumClasses() uint32 {
	if m != nil {
		return m.NumClasses
	}
	return 0
}

func (m *LearnerConfig) GetFeatureFractionPerTree() float32 {
	if x, ok := m.GetFeatureFraction().(*LearnerConfig_FeatureFractionPerTree); ok {
		return x.FeatureFractionPerTree
	}
	return 0
}

func (m *LearnerConfig) GetFeatureFractionPerLevel() float32 {
	if x, ok := m.GetFeatureFraction().(*LearnerConfig_FeatureFractionPerLevel); ok {
		return x.FeatureFractionPerLevel
	}
	return 0
}

func (m *LearnerConfig) GetRegularization() *TreeRegularizationConfig {
	if m != nil {
		return m.Regularization
	}
	return nil
}

func (m *LearnerConfig) GetConstraints() *TreeConstraintsConfig {
	if m != nil {
		return m.Constraints
	}
	return nil
}

func (m *LearnerConfig) GetPruningMode() LearnerConfig_PruningMode {
	if m != nil {
		return m.PruningMode
	}
	return LearnerConfig_PRE_PRUNE
}

func (m *LearnerConfig) GetGrowingMode() LearnerConfig_GrowingMode {
	if m != nil {
		return m.GrowingMode
	}
	return LearnerConfig_WHOLE_TREE
}

func (m *LearnerConfig) GetLearningRateTuner() *LearningRateConfig {
	if m != nil {
		return m.LearningRateTuner
	}
	return nil
}

func (m *LearnerConfig) GetMultiClassStrategy() LearnerConfig_MultiClassStrategy {
	if m != nil {
		return m.MultiClassStrategy
	}
	return LearnerConfig_TREE_PER_CLASS
}

func (m *LearnerConfig) GetAveragingConfig() *AveragingConfig {
	if m != nil {
		return m.AveragingConfig
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LearnerConfig) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _LearnerConfig_OneofMarshaler, _LearnerConfig_OneofUnmarshaler, _LearnerConfig_OneofSizer, []interface{}{
		(*LearnerConfig_FeatureFractionPerTree)(nil),
		(*LearnerConfig_FeatureFractionPerLevel)(nil),
	}
}

func _LearnerConfig_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*LearnerConfig)
	// feature_fraction
	switch x := m.FeatureFraction.(type) {
	case *LearnerConfig_FeatureFractionPerTree:
		b.EncodeVarint(2<<3 | proto1.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.FeatureFractionPerTree)))
	case *LearnerConfig_FeatureFractionPerLevel:
		b.EncodeVarint(3<<3 | proto1.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.FeatureFractionPerLevel)))
	case nil:
	default:
		return fmt.Errorf("LearnerConfig.FeatureFraction has unexpected type %T", x)
	}
	return nil
}

func _LearnerConfig_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*LearnerConfig)
	switch tag {
	case 2: // feature_fraction.feature_fraction_per_tree
		if wire != proto1.WireFixed32 {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.FeatureFraction = &LearnerConfig_FeatureFractionPerTree{math.Float32frombits(uint32(x))}
		return true, err
	case 3: // feature_fraction.feature_fraction_per_level
		if wire != proto1.WireFixed32 {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.FeatureFraction = &LearnerConfig_FeatureFractionPerLevel{math.Float32frombits(uint32(x))}
		return true, err
	default:
		return false, nil
	}
}

func _LearnerConfig_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*LearnerConfig)
	// feature_fraction
	switch x := m.FeatureFraction.(type) {
	case *LearnerConfig_FeatureFractionPerTree:
		n += proto1.SizeVarint(2<<3 | proto1.WireFixed32)
		n += 4
	case *LearnerConfig_FeatureFractionPerLevel:
		n += proto1.SizeVarint(3<<3 | proto1.WireFixed32)
		n += 4
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto1.RegisterType((*TreeRegularizationConfig)(nil), "tensorflow.boosted_trees.learner.TreeRegularizationConfig")
	proto1.RegisterType((*TreeConstraintsConfig)(nil), "tensorflow.boosted_trees.learner.TreeConstraintsConfig")
	proto1.RegisterType((*LearningRateConfig)(nil), "tensorflow.boosted_trees.learner.LearningRateConfig")
	proto1.RegisterType((*LearningRateFixedConfig)(nil), "tensorflow.boosted_trees.learner.LearningRateFixedConfig")
	proto1.RegisterType((*LearningRateLineSearchConfig)(nil), "tensorflow.boosted_trees.learner.LearningRateLineSearchConfig")
	proto1.RegisterType((*AveragingConfig)(nil), "tensorflow.boosted_trees.learner.AveragingConfig")
	proto1.RegisterType((*LearningRateDropoutDrivenConfig)(nil), "tensorflow.boosted_trees.learner.LearningRateDropoutDrivenConfig")
	proto1.RegisterType((*LearnerConfig)(nil), "tensorflow.boosted_trees.learner.LearnerConfig")
	proto1.RegisterEnum("tensorflow.boosted_trees.learner.LearnerConfig_PruningMode", LearnerConfig_PruningMode_name, LearnerConfig_PruningMode_value)
	proto1.RegisterEnum("tensorflow.boosted_trees.learner.LearnerConfig_GrowingMode", LearnerConfig_GrowingMode_name, LearnerConfig_GrowingMode_value)
	proto1.RegisterEnum("tensorflow.boosted_trees.learner.LearnerConfig_MultiClassStrategy", LearnerConfig_MultiClassStrategy_name, LearnerConfig_MultiClassStrategy_value)
}

func init() {
	proto1.RegisterFile("tensorflow/contrib/boosted_trees/proto/learner.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 946 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x6f, 0x6f, 0xdb, 0xb6,
	0x13, 0x8e, 0xdc, 0x5f, 0xd2, 0xe4, 0x14, 0x3b, 0x0e, 0x9b, 0xdf, 0xaa, 0x6d, 0x05, 0x12, 0x78,
	0xc3, 0x56, 0x0c, 0xab, 0x8d, 0x78, 0x05, 0x86, 0xad, 0x58, 0x01, 0x3b, 0x71, 0x9a, 0x02, 0x6a,
	0xe2, 0xca, 0xe9, 0x8a, 0x0c, 0xdb, 0x08, 0x5a, 0xa2, 0x15, 0x22, 0x12, 0x29, 0x50, 0x54, 0xfe,
	0xec, 0x3b, 0xec, 0x9b, 0xec, 0xd5, 0x3e, 0xc0, 0x3e, 0xd7, 0x5e, 0x0e, 0xa4, 0x18, 0x5b, 0xf9,
	0xd3, 0x25, 0xd9, 0x3b, 0xdd, 0xf1, 0x9e, 0xe7, 0x78, 0xcf, 0x1d, 0x0f, 0x82, 0xe7, 0x8a, 0xf2,
	0x5c, 0xc8, 0x49, 0x22, 0x4e, 0x3b, 0xa1, 0xe0, 0x4a, 0xb2, 0x71, 0x67, 0x2c, 0x44, 0xae, 0x68,
	0x84, 0x95, 0xa4, 0x34, 0xef, 0x64, 0x52, 0x28, 0xd1, 0x49, 0x28, 0x91, 0x9c, 0xca, 0xb6, 0xb1,
	0xd0, 0xc6, 0x0c, 0xd5, 0xbe, 0x14, 0xdd, 0xb6, 0x71, 0xad, 0x10, 0xbc, 0x03, 0x49, 0x69, 0x40,
	0xe3, 0x22, 0x21, 0x92, 0xfd, 0x46, 0x14, 0x13, 0x7c, 0x4b, 0xf0, 0x09, 0x8b, 0x51, 0x03, 0x6a,
	0xc9, 0xa6, 0xe7, 0x6c, 0x38, 0x4f, 0x6b, 0x41, 0x2d, 0xd9, 0x34, 0x76, 0xd7, 0xab, 0x59, 0xbb,
	0x8b, 0xbe, 0x84, 0x15, 0x4d, 0x86, 0x43, 0x91, 0x66, 0x09, 0x3d, 0x63, 0xea, 0xdc, 0x7b, 0x60,
	0x0e, 0x1b, 0xda, 0xbd, 0x35, 0xf5, 0xb6, 0x28, 0xfc, 0xff, 0xc0, 0x78, 0x78, 0xae, 0x24, 0x61,
	0x5c, 0xe5, 0x36, 0xc3, 0xe7, 0xd0, 0x48, 0xc9, 0x99, 0xb9, 0x12, 0x8e, 0x68, 0xa6, 0x8e, 0x4c,
	0xb6, 0x7a, 0xb0, 0x9c, 0x92, 0x33, 0x8d, 0xd8, 0xd6, 0x3e, 0xf4, 0x05, 0xac, 0xa4, 0x8c, 0x63,
	0x2e, 0x22, 0x8a, 0x4f, 0x29, 0x8b, 0x8f, 0x94, 0xbd, 0x44, 0x3d, 0x65, 0x7c, 0x4f, 0x44, 0xf4,
	0xbd, 0x71, 0xb6, 0xfe, 0xac, 0x01, 0xf2, 0x75, 0x5d, 0x8c, 0xc7, 0x01, 0x51, 0xd4, 0x26, 0x79,
	0x0b, 0xf3, 0x13, 0x76, 0x46, 0x23, 0xc3, 0xed, 0x76, 0xbf, 0x6b, 0xdf, 0x26, 0x4a, 0xbb, 0x4a,
	0xb2, 0xa3, 0xa1, 0x25, 0xd3, 0xee, 0x5c, 0x50, 0x32, 0xa1, 0x5f, 0xe0, 0x61, 0x24, 0x45, 0x26,
	0x8a, 0xf2, 0x26, 0x6e, 0xb7, 0x77, 0x3f, 0xd2, 0xed, 0x12, 0xbc, 0x2d, 0xd9, 0x09, 0xe5, 0x53,
	0xf2, 0x0b, 0x4e, 0x44, 0xc0, 0x4d, 0x18, 0xa7, 0x38, 0xa7, 0x44, 0x86, 0x47, 0x46, 0x54, 0xb7,
	0xfb, 0xf2, 0x7e, 0x29, 0x7c, 0xc6, 0xe9, 0xc8, 0xe0, 0xa7, 0xfc, 0x90, 0x4c, 0x7d, 0xfd, 0x87,
	0x30, 0xaf, 0x0a, 0x3d, 0x00, 0x2f, 0xe1, 0xf1, 0x07, 0xca, 0x45, 0x9f, 0x41, 0x3d, 0xb1, 0x47,
	0x58, 0x12, 0x45, 0xed, 0x28, 0x2c, 0x27, 0x95, 0xf8, 0x56, 0x0c, 0x4f, 0xfe, 0x2d, 0x2d, 0xfa,
	0x0a, 0x56, 0x75, 0x8b, 0x6f, 0x22, 0x5a, 0x49, 0xc9, 0x59, 0x15, 0x8b, 0x3e, 0x85, 0x25, 0x5e,
	0xa4, 0x38, 0x57, 0x34, 0xcb, 0x8d, 0xb0, 0xf3, 0xc1, 0x22, 0x2f, 0xd2, 0x91, 0xb6, 0x5b, 0xbf,
	0x3b, 0xb0, 0xd2, 0x3b, 0xa1, 0x92, 0xc4, 0x8c, 0xc7, 0x96, 0x7c, 0x13, 0xd6, 0x88, 0x71, 0x51,
	0x9c, 0x90, 0x5c, 0x61, 0x5e, 0x0a, 0x52, 0xf2, 0xef, 0xce, 0x05, 0xab, 0xf6, 0xd4, 0x27, 0xb9,
	0xda, 0xd3, 0x03, 0x95, 0xa3, 0x1f, 0xe0, 0x93, 0x4b, 0x90, 0x8c, 0xca, 0x90, 0x72, 0x65, 0x81,
	0x35, 0x0b, 0x7c, 0x5c, 0x01, 0x0e, 0xcb, 0x08, 0x03, 0xef, 0x2f, 0xc2, 0x42, 0x68, 0x72, 0xb7,
	0xfe, 0x72, 0x60, 0xfd, 0x96, 0x9e, 0xa2, 0x0e, 0x3c, 0xb2, 0x3d, 0xc5, 0x99, 0x14, 0x63, 0x32,
	0x66, 0x89, 0x7e, 0x25, 0x65, 0xf9, 0xc8, 0x1e, 0x0d, 0x67, 0x27, 0x68, 0x00, 0xeb, 0x95, 0x40,
	0x2c, 0x26, 0x38, 0x3f, 0x66, 0x59, 0xa6, 0x85, 0xab, 0x0e, 0x5c, 0x2d, 0x78, 0x52, 0x09, 0xdb,
	0x9f, 0x8c, 0x6c, 0x90, 0xbd, 0xc3, 0xf5, 0xce, 0x3d, 0xb8, 0xa1, 0x73, 0x7f, 0x2c, 0x42, 0xdd,
	0x2f, 0x27, 0xc7, 0x5e, 0x77, 0x1d, 0x5c, 0xad, 0x7f, 0x98, 0x90, 0x3c, 0xb7, 0x2a, 0xd6, 0x03,
	0xe0, 0x45, 0xba, 0x55, 0x7a, 0xd0, 0x0b, 0xf8, 0x78, 0x42, 0x89, 0x2a, 0x24, 0xc5, 0x13, 0x49,
	0x42, 0xbd, 0x2b, 0xb4, 0x80, 0x46, 0xbc, 0xa9, 0x76, 0x1f, 0xd9, 0x90, 0x1d, 0x1b, 0x31, 0xa4,
	0x52, 0x6b, 0xa7, 0x95, 0xbf, 0x11, 0x9c, 0xd0, 0x13, 0x9a, 0x94, 0x37, 0xd4, 0xca, 0x5f, 0x47,
	0xfb, 0x3a, 0x00, 0x8d, 0xa1, 0x21, 0x2f, 0x6d, 0x29, 0xef, 0x7f, 0xe6, 0x5d, 0x7c, 0x7f, 0xfb,
	0xbb, 0xf8, 0xd0, 0x86, 0x0b, 0xae, 0x30, 0xa2, 0x43, 0x70, 0xc3, 0xd9, 0x92, 0xf2, 0xe6, 0x4d,
	0x82, 0x6f, 0xef, 0x96, 0xe0, 0xda, 0x76, 0x0b, 0xaa, 0x5c, 0xe8, 0x57, 0x58, 0xce, 0x64, 0x61,
	0x3a, 0x92, 0x8a, 0x88, 0x7a, 0x8b, 0x1b, 0xce, 0xd3, 0x46, 0xf7, 0xc5, 0x1d, 0x1f, 0xf5, 0x45,
	0x8b, 0xda, 0xc3, 0x92, 0xe3, 0x8d, 0x88, 0x68, 0xe0, 0x66, 0x33, 0x43, 0xf3, 0xc7, 0x52, 0x9c,
	0x4e, 0xf9, 0x97, 0xfe, 0x1b, 0xff, 0xab, 0x92, 0xa3, 0xe4, 0x8f, 0x67, 0x06, 0x8a, 0xe0, 0xd1,
	0xa5, 0x91, 0xc2, 0x66, 0x7d, 0x78, 0x0b, 0x46, 0xa2, 0xe7, 0xf7, 0xdb, 0x4d, 0x56, 0x9f, 0xd5,
	0xea, 0x38, 0x1e, 0x68, 0x3a, 0xa4, 0x60, 0x2d, 0x2d, 0x12, 0xc5, 0xca, 0x19, 0xc4, 0x5a, 0x3d,
	0x45, 0xe3, 0x73, 0x0f, 0x4c, 0x35, 0xfd, 0xfb, 0x56, 0xf3, 0x46, 0x73, 0x99, 0xe1, 0x1d, 0x59,
	0xa6, 0x00, 0xa5, 0xd7, 0x7c, 0xe8, 0x67, 0x68, 0x92, 0x8b, 0xcd, 0x82, 0xcb, 0xe7, 0xed, 0xb9,
	0xa6, 0xb0, 0xcd, 0xdb, 0x33, 0x5e, 0xd9, 0x49, 0xc1, 0x0a, 0xb9, 0xec, 0x68, 0x7d, 0x0d, 0x6e,
	0xa5, 0x6b, 0xa8, 0x0e, 0x4b, 0xc3, 0x60, 0x80, 0x87, 0xc1, 0xbb, 0xbd, 0x41, 0x73, 0x0e, 0x35,
	0x00, 0x86, 0xfb, 0xa3, 0x03, 0x6b, 0x3b, 0xad, 0x4d, 0x70, 0x2b, 0x3d, 0xd0, 0xc7, 0xef, 0x77,
	0xf7, 0xfd, 0x01, 0x3e, 0x08, 0x06, 0x3a, 0x1c, 0x41, 0xc3, 0xef, 0x1d, 0x0e, 0x02, 0xdc, 0x3f,
	0xc4, 0xe6, 0xa3, 0xe9, 0xb4, 0x86, 0x80, 0xae, 0x17, 0xaa, 0x23, 0x35, 0x06, 0x0f, 0x07, 0x01,
	0xde, 0xf2, 0x7b, 0xa3, 0x51, 0x73, 0x0e, 0x35, 0x61, 0x79, 0xe7, 0x9d, 0xef, 0xe3, 0xdd, 0xc1,
	0x68, 0xf4, 0xba, 0xb7, 0xd7, 0x74, 0xd0, 0x1a, 0x34, 0xb7, 0x5f, 0xf7, 0x5e, 0xed, 0xef, 0xf5,
	0x66, 0xde, 0x5a, 0x1f, 0x41, 0xf3, 0xea, 0x53, 0xed, 0xff, 0xf8, 0xd3, 0xdb, 0x98, 0xa9, 0xa3,
	0x62, 0xdc, 0x0e, 0x45, 0xda, 0x89, 0x24, 0x3d, 0x3f, 0xee, 0xcc, 0xc4, 0x79, 0x96, 0x53, 0x79,
	0xc2, 0x78, 0xfc, 0x2c, 0x16, 0x9d, 0xec, 0x38, 0xee, 0xdc, 0xed, 0x77, 0xe5, 0x6f, 0xc7, 0x19,
	0x2f, 0x98, 0xaf, 0x6f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x61, 0x5c, 0xc6, 0xf7, 0xe2, 0x08,
	0x00, 0x00,
}