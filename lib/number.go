package jd

import (
	"bytes"
	"encoding/binary"
)

type jsonNumber float64

var _ JsonNode = jsonNumber(0)

func (n jsonNumber) Json(metadata ...Metadata) string {
	return renderJson(n.raw(metadata))
}

func (n jsonNumber) Yaml(metadata ...Metadata) string {
	return renderYaml(n.raw(metadata))
}

func (n jsonNumber) raw(_ []Metadata) interface{} {
	return float64(n)
}

func (n1 jsonNumber) Equals(node JsonNode, metadata ...Metadata) bool {
	n2, ok := node.(jsonNumber)
	if !ok {
		return false
	}
	if n1 != n2 {
		return false
	}
	return true
}

func (n jsonNumber) hashCode(metadata []Metadata) [8]byte {
	a := make([]byte, 0, 8)
	b := bytes.NewBuffer(a)
	binary.Write(b, binary.LittleEndian, n)
	return hash(b.Bytes())
}

func (n jsonNumber) Diff(node JsonNode, metadata ...Metadata) Diff {
	return n.diff(node, make(path, 0), metadata, getPatchStrategy(metadata))
}

func (n jsonNumber) diff(
	node JsonNode,
	path path,
	metadata []Metadata,
	strategy patchStrategy,
) Diff {
	return diff(n, node, path, metadata, strategy)
}

func (n jsonNumber) Patch(d Diff) (JsonNode, error) {
	return patchAll(n, d)
}

func (n jsonNumber) patch(
	pathBehind, pathAhead path,
	oldValues, newValues []JsonNode,
	strategy patchStrategy,
) (JsonNode, error) {
	return patch(n, pathBehind, pathAhead, oldValues, newValues, strategy)
}
