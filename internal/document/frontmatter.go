package document

import (
	"bytes"
	"encoding/json"
	stderrors "errors"

	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/stateful/runme/internal/ulid"
	"github.com/stateful/runme/internal/version"
)

var ErrFrontmatterInvalid = stderrors.New("invalid frontmatter")

const (
	frontmatterFormatYAML = "yaml"
	frontmatterFormatJSON = "json"
	frontmatterFormatTOML = "toml"
)

type RunmeMetadata struct {
	ID      string `yaml:"id,omitempty" json:"id,omitempty" toml:"id,omitempty"`
	Version string `yaml:"version,omitempty" json:"version,omitempty" toml:"version,omitempty"`
}

type Frontmatter struct {
	Runme       RunmeMetadata `yaml:"runme,omitempty"`
	Shell       string        `yaml:"shell"`
	Cwd         string        `yaml:"cwd"`
	SkipPrompts bool          `yaml:"skipPrompts,omitempty"`

	format string
	raw    string // using string to be able to compare using ==
}

func newFrontmatter() *Frontmatter {
	return &Frontmatter{
		Runme: RunmeMetadata{
			ID:      ulid.GenerateID(),
			Version: version.BaseVersion(),
		},

		format: frontmatterFormatYAML,
	}
}

// Marshal returns a marshaled frontmatter including triple-dashed lines.
// If the identity is required, but Frontmatter is nil, a new one is created.
func (f *Frontmatter) Marshal(requireIdentity bool) ([]byte, error) {
	if f == nil {
		if !requireIdentity {
			return nil, nil
		}
		f = newFrontmatter()
	}
	return f.marshal(requireIdentity)
}

func (f *Frontmatter) marshal(requireIdentity bool) ([]byte, error) {
	if requireIdentity {
		f.ensureID()
	}

	switch f.format {
	case frontmatterFormatYAML:
		m := make(map[string]interface{})

		if err := yaml.Unmarshal([]byte(f.raw), &m); err != nil {
			return nil, errors.WithStack(err)
		}

		if requireIdentity {
			m["runme"] = f.Runme
		}

		var buf bytes.Buffer
		encoder := yaml.NewEncoder(&buf)
		encoder.SetIndent(2)
		if err := encoder.Encode(m); err != nil {
			return nil, errors.WithStack(err)
		}
		if err := encoder.Close(); err != nil {
			return nil, errors.WithStack(err)
		}

		return append(append([]byte("---\n"), buf.Bytes()...), []byte("---")...), nil

	case frontmatterFormatJSON:
		m := make(map[string]interface{})

		if err := json.Unmarshal([]byte(f.raw), &m); err != nil {
			return nil, errors.WithStack(err)
		}

		if requireIdentity {
			m["runme"] = f.Runme
		}

		data, err := json.Marshal(m)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return append(append([]byte("---\n"), data...), []byte("---")...), nil

	case frontmatterFormatTOML:
		m := make(map[string]interface{})

		if err := toml.Unmarshal([]byte(f.raw), &m); err != nil {
			return nil, errors.WithStack(err)
		}

		if requireIdentity {
			m["runme"] = f.Runme
		}

		data, err := toml.Marshal(m)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return append(append([]byte("+++\n"), data...), []byte("+++")...), nil

	default:
		panic("invariant: Frontmatter created with invalid format")
	}
}

func (f *Frontmatter) ensureID() {
	if !ulid.ValidID(f.Runme.ID) {
		f.Runme.ID = ulid.GenerateID()
	}

	baseVersion := version.BaseVersion()
	if baseVersion == "v99.9" || baseVersion == "v0.0" {
		return
	}
	f.Runme.Version = baseVersion
}

// TODO(adamb): Frontmatter can return (nil, nil) which indicates that
// there is no error, but also that there is no FrontMatter. It is not
// the best API. Consider adding HasFrontmatter() and returning an error
// if there is none from this method.
func (d *Document) Frontmatter() (*Frontmatter, error) {
	d.splitSource()

	if d.splitSourceErr != nil {
		return nil, d.splitSourceErr
	}

	d.parseFrontmatter()

	if d.parseFrontmatterErr != nil {
		return nil, d.parseFrontmatterErr
	}

	return d.frontmatter, nil
}

func (d *Document) parseFrontmatter() {
	d.onceParseFrontmatter.Do(func() {
		d.frontmatter, d.parseFrontmatterErr = parseFrontmatter(d.frontmatterRaw)
	})
}

// TODO(adamb): it should be removed when the complete refactoring of the project is finished.
func ParseFrontmatter(raw []byte) (*Frontmatter, error) {
	return parseFrontmatter(raw)
}

func parseFrontmatter(raw []byte) (*Frontmatter, error) {
	if len(raw) == 0 {
		return nil, nil
	}

	// We know that frontmatter is not empty,
	// so d.frontmatter won't be nil ever.
	// However, it may still be invalid and
	// this detail will be in d.parseFrontmatterErr.
	var f Frontmatter

	lines := bytes.Split(raw, []byte{'\n'})

	if len(lines) < 2 || !bytes.Equal(bytes.TrimSpace(lines[0]), bytes.TrimSpace(lines[len(lines)-1])) {
		return nil, errors.WithStack(ErrFrontmatterInvalid)
	}

	raw = bytes.Join(lines[1:len(lines)-1], []byte{'\n'})

	// TODO(adamb): discuss how to approach this in the most sensible way.
	// It can always return to the initial idea of returning all errors,
	// but the client will be left with the same problem.
	parsers := []func([]byte, any) error{
		yaml.Unmarshal,
		json.Unmarshal,
		toml.Unmarshal,
	}
	parsersNames := []string{
		frontmatterFormatYAML,
		frontmatterFormatJSON,
		frontmatterFormatTOML,
	}
	errorsCount := 0

	var firstError error

	for idx, parser := range parsers {
		err := parser(raw, &f)
		if err != nil {
			errorsCount++

			if firstError == nil {
				firstError = errors.Wrap(err, "failed to parse frontmatter content")
			}
		} else {
			f.format = parsersNames[idx]
			f.raw = string(raw)
			break
		}
	}

	// If all parsers returned errors, select the first one.
	if errorsCount == len(parsers) {
		return nil, firstError
	}

	return &f, nil
}
