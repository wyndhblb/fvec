/*
   Tags and various helper methods

   A Tag is just a tuple of name and value both as strings
*/

package vepr

import (
	"fmt"
	"io"
	"strings"
)

// Join makes a {name}{sep}{value} string
func (t *Tag) Join(sep string) string {
	return t.Name + sep + t.Value
}

type Tags []*Tag

func (p Tags) Len() int           { return len(p) }
func (p Tags) Less(i, j int) bool { return strings.Compare(p[i].Name, p[j].Name) < 0 }
func (p Tags) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// TagsFromString make a tag array from a string input inder a few conditions
// tag=val.tag=val.tag=val
// or
// tag_is_val.tag_is_val
// or
// tag=val,tag=val
// or
// tag_is_val,tag_is_val
// or
// tag=val tag=val
// or
// tag_is_val tag_is_val
// or
// tag:val,tag:val
func TagsFromString(key string) Tags {

	var parse_tsg []string
	if strings.Contains(key, ".") {
		parse_tsg = strings.Split(key, ".")
	} else if strings.Contains(key, ",") {
		parse_tsg = strings.Split(key, ",")
	} else {
		parse_tsg = strings.Split(key, " ")
	}
	return TagsFromArray(parse_tsg)
}

// TagsFromArray make Tags list from an array of {name}{sep}{value} strings
// follows the same rules as the TagsFromString for individual strings allowed
func TagsFromArray(keys []string) Tags {

	outs := make(Tags, 0)

	for _, tgs := range keys {
		spls := strings.Split(tgs, "=")
		if len(spls) < 2 {
			// try "_is_"
			spls = strings.Split(tgs, "_is_")
			if len(spls) < 2 {
				// try ":"
				spls = strings.Split(tgs, ":")
				if len(spls) < 2 {
					continue
				}
			}
		}
		outs = append(outs, &Tag{Name: spls[0], Value: spls[1]})
	}
	return outs
}

// String dump as {name}={value} {name}={value} ...
func (s Tags) String() string {
	return s.BaseString()
}

// StringList dump as [{name}={value}, {name}={value}] ...
// NOTE: not thread safe
func (s Tags) StringList() (out []string) {
	for _, t := range s {
		out = append(out, fmt.Sprintf("%s=%s", t.Name, t.Value))
	}
	return out
}

// BaseString default string representation of a tag
func (s Tags) BaseString() string {
	return s.ToStringSep("=", " ")
}

// IsEmpty have any data?
func (s Tags) IsEmpty() bool {
	return len(s) == 0
}

// Tags cast to a []*Tag
func (s Tags) Tags() []*Tag {
	return s
}

// Merge merge two tag sets
// NOTE: the Incoming tags will overwrite any ones in the based tags set if they are the same name
// NOTE: not thread safe
func (s Tags) Merge(tags Tags) Tags {
	if len(tags) == 0 {
		return s
	}
	if len(s) == 0 {
		return tags
	}
	n_tags := s
	for _, tag := range tags {
		got := false
		for idx, o_tag := range n_tags {
			if tag.Name == o_tag.Name {
				n_tags[idx].Value = tag.Value
				got = true
				break
			}
		}
		if !got {
			n_tags = append(n_tags, &Tag{Name: tag.Name, Value: tag.Value})
		}
	}
	return n_tags
}

// HasAllTags given an input set of tags, see if the incoming set matches all the name/value pairs
// in the current list (basically does s intersect with the incoming)
// NOTE not thread safe
func (s Tags) HasAllTags(tags Tags) bool {
	if len(tags) == 0 {
		return true
	}
	if len(s) == 0 && len(tags) > 0 {
		return false
	}
	have_ct := 0
	for _, tag := range tags {
		for _, o_tag := range s {
			if tag.Name == o_tag.Name && tag.Value == o_tag.Value {
				have_ct++
			}
		}
	}
	return have_ct == len(tags)
}

// ToStringSep make a string of the form {name}{tagsep}{value}{wordsep}{name}{tagsep}{value}...
// NOTE not thread safe
func (s Tags) ToStringSep(wordsep string, tagsep string) string {
	str := make([]string, len(s))
	for idx, tag := range s {
		str[idx] = tag.Join(wordsep)
	}
	return strings.Join(str, tagsep)
}

// WriteBytes write tags to a byte buffer
// NOTE not thread safe
func (s Tags) WriteBytes(buf io.Writer, wordsep []byte, tagsep []byte) {

	l := len(s)
	for idx, tag := range s {
		buf.Write([]byte(tag.Name))
		buf.Write(wordsep)
		buf.Write([]byte(tag.Value))
		if idx < l-1 {
			buf.Write(tagsep)
		}
	}
}

// Find find tag by name
// NOTE not thread safe
func (s Tags) Find(name string) string {
	for _, tag := range s {
		if tag.Name == name {
			return tag.Value
		}
	}
	return ""
}

// Set add name/value (or overwrite if name is already there)
// NOTE not thread safe
func (s Tags) Set(name string, val string) Tags {
	for idx, tag := range s {
		if tag.Name == name {
			s[idx].Value = val
			return s
		}
	}
	s = append(s, &Tag{Name: name, Value: val})
	return s
}
