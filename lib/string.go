package jd

type jsonString string

var _ JsonNode = jsonString("")

func (s jsonString) Json(metadata ...Metadata) string {
	return renderJson(s.raw(metadata))
}

func (s jsonString) Yaml(metadata ...Metadata) string {
	return renderYaml(s.raw(metadata))
}

func (s jsonString) raw(_ []Metadata) interface{} {
	return string(s)
}

func (s1 jsonString) Equals(n JsonNode, metadata ...Metadata) bool {
	s2, ok := n.(jsonString)
	if !ok {
		return false
	}
	return s1 == s2
}

func (s jsonString) hashCode(metadata []Metadata) [8]byte {
	return hash([]byte(s))
}

func (s jsonString) Diff(n JsonNode, metadata ...Metadata) Diff {
	return s.diff(n, make(path, 0), metadata, getPatchStrategy(metadata))
}

func (s1 jsonString) diff(
	n JsonNode,
	path path,
	metadata []Metadata,
	strategy patchStrategy,
) Diff {
	return diff(s1, n, path, metadata, strategy)
}

func (s jsonString) Patch(d Diff) (JsonNode, error) {
	return patchAll(s, d)
}

func (s jsonString) patch(
	pathBehind, pathAhead path,
	oldValues, newValues []JsonNode,
	strategy patchStrategy,
) (JsonNode, error) {
	return patch(s, pathBehind, pathAhead, oldValues, newValues, strategy)
}
