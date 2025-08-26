package kconsole

import (
	"encoding/json"

	"github.com/twmb/franz-go/pkg/sr"
)

// --- SchemaType ---
type SchemaType sr.SchemaType

const (
	TypeAvro     SchemaType = SchemaType(sr.TypeAvro)
	TypeProtobuf SchemaType = SchemaType(sr.TypeProtobuf)
	TypeJSON     SchemaType = SchemaType(sr.TypeJSON)
)

func (t SchemaType) MarshalJSON() ([]byte, error) {
	text, err := sr.SchemaType(t).MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

func (t SchemaType) String() string {
	return sr.SchemaType(t).String()
}

// --- CompatibilityLevel ---
type CompatibilityLevel sr.CompatibilityLevel

const (
	CompatNone               CompatibilityLevel = CompatibilityLevel(sr.CompatNone)
	CompatBackward           CompatibilityLevel = CompatibilityLevel(sr.CompatBackward)
	CompatBackwardTransitive CompatibilityLevel = CompatibilityLevel(sr.CompatBackwardTransitive)
	CompatForward            CompatibilityLevel = CompatibilityLevel(sr.CompatForward)
	CompatForwardTransitive  CompatibilityLevel = CompatibilityLevel(sr.CompatForwardTransitive)
	CompatFull               CompatibilityLevel = CompatibilityLevel(sr.CompatFull)
	CompatFullTransitive     CompatibilityLevel = CompatibilityLevel(sr.CompatFullTransitive)
)

func (l CompatibilityLevel) MarshalJSON() ([]byte, error) {
	text, err := sr.CompatibilityLevel(l).MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

func (l CompatibilityLevel) String() string {
	return sr.CompatibilityLevel(l).String()
}

// --- Mode ---
type Mode sr.Mode

const (
	ModeImport    Mode = Mode(sr.ModeImport)
	ModeReadOnly  Mode = Mode(sr.ModeReadOnly)
	ModeReadWrite Mode = Mode(sr.ModeReadWrite)
)

func (m Mode) MarshalJSON() ([]byte, error) {
	text, err := sr.Mode(m).MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

func (m Mode) String() string {
	return sr.Mode(m).String()
}

// --- SchemaRuleKind ---
type SchemaRuleKind sr.SchemaRuleKind

const (
	SchemaRuleKindTransform SchemaRuleKind = SchemaRuleKind(sr.SchemaRuleKindTransform)
	SchemaRuleKindCondition SchemaRuleKind = SchemaRuleKind(sr.SchemaRuleKindCondition)
)

func (k SchemaRuleKind) MarshalJSON() ([]byte, error) {
	text, err := sr.SchemaRuleKind(k).MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

func (k SchemaRuleKind) String() string {
	return sr.SchemaRuleKind(k).String()
}

// --- SchemaRuleMode ---
type SchemaRuleMode sr.SchemaRuleMode

const (
	SchemaRuleModeUpgrade   SchemaRuleMode = SchemaRuleMode(sr.SchemaRuleModeUpgrade)
	SchemaRuleModeDowngrade SchemaRuleMode = SchemaRuleMode(sr.SchemaRuleModeDowngrade)
	SchemaRuleModeUpdown    SchemaRuleMode = SchemaRuleMode(sr.SchemaRuleModeUpdown)
	SchemaRuleModeWrite     SchemaRuleMode = SchemaRuleMode(sr.SchemaRuleModeWrite)
	SchemaRuleModeRead      SchemaRuleMode = SchemaRuleMode(sr.SchemaRuleModeRead)
	SchemaRuleModeWriteRead SchemaRuleMode = SchemaRuleMode(sr.SchemaRuleModeWriteRead)
)

func (m SchemaRuleMode) MarshalJSON() ([]byte, error) {
	text, err := sr.SchemaRuleMode(m).MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

func (m SchemaRuleMode) String() string {
	return sr.SchemaRuleMode(m).String()
}
