/*
   Vector Name Object methods

   The base objects structs are in vepr.proto and generated
*/

package fvec

import (
	"github.com/wyndhblb/go-utils/pools"
	"github.com/wyndhblb/timeslab"
	"sort"
	"strconv"
	"strings"
	"time"
)

// SetKey set the key name for the name
func (s *VName) SetKey(name string) {
	s.Key = name
	if name != s.Key {
		s.XUniqueId = 0
		s.XUniqueStr = ""
	}
}

// ToSlab takes a time and converts into the proper "slab" string for a given resolution
// min YYYYMMDDHHMM
// hour YYYYMMDDHH
// day YYYYMMDD
// month YYYYMM
// year YYYY
// all ALL
//
// MIN10 -> YYYYMMDDHHI10{min / 10}
// MIN30 -> YYYYMMDDHHI30{min / 30}
// MONTH2 -> YYYYM2{month / 2}
// MONTH3 -> YYYYM3{month / 3}
// MONTH6 -> YYYYM6{month / 6}
//
func (s *VName) ToSlab(t time.Time) string {
	return timeslab.ToSlab(s.Resolution, t)
}

// SetResolutionFromString will set the vector name resolution from a string
// (see ResolutionFromString for string to res mappings)
func (s *VName) SetResolutionFromString(res string) {
	s.Resolution = timeslab.ResolutionFromString(res)
}

// UniqueId take the various "parts" (keys, resolution, tags) and return a basic md5 hash of things
func (s *VName) UniqueId() uint64 {
	if s.XUniqueId > 0 {
		return s.XUniqueId
	}
	buf := pools.GetFnv64a()
	defer pools.PutFnv64a(buf)

	// all tags or just intrinsic tags
	buf.Write([]byte(s.Key + ":" + s.SortedTags().BaseString()))

	s.XUniqueId = buf.Sum64()
	return s.XUniqueId
}

// UniqueIdString base36 rep of the UniqueId
// keep it in the object as the computation can yield many GC things from the Fprintf above
func (s *VName) UniqueIdString() string {
	if s.XUniqueStr == "" {
		id := s.UniqueId()
		s.XUniqueStr = strconv.FormatUint(uint64(id), 36)
	}
	return s.XUniqueStr
}

// StringToUniqueId undo a base36 UnqueIdString to the uint64 version
func (s *VName) StringToUniqueId(inid string) uint64 {
	tInt, err := strconv.ParseUint(inid, 36, 64)
	if err != nil {
		return 0
	}
	return tInt
}

// SortedTags return an array of [ [name, val] ...] sorted by name
func (s *VName) SortedTags() Tags {
	if len(s.Tags) > 0 {
		sort.Sort(Tags(s.Tags))
	}
	return s.Tags
}

// MergeTags return an array of [ [name, val] ...] sorted by name
func (s *VName) MergeTags(tags Tags) Tags {
	var nTags Tags
	for _, tag := range tags {
		got := false
		for _, oTag := range s.Tags {
			if tag.Name == oTag.Name {
				nTags = append(nTags, &Tag{tag.Name, tag.Value})
				got = true
				break
			}
		}
		if !got {
			nTags = append(nTags, &Tag{Name: tag.Name, Value: tag.Value})
		}
	}
	s.Tags = nTags
	s.XUniqueStr = ""
	s.XUniqueId = 0
	return nTags
}

// IsBlank is the key set?
func (s *VName) IsBlank() bool {
	return len(s.Key) == 0
}

// StringName returns {key}.{tags}
// this is a full string compatible name of {key}.{name.value}.{name.value}
// with Names of the tags SORTED
func (s *VName) StringName() string {
	sTags := s.SortedTags()
	str := make([]string, 1+len(sTags))
	str[0] = s.Key
	for idx, tg := range sTags {
		str[idx+1] = tg.Join(".")
	}
	return strings.Join(str, ".")
}

// Copy clone a name
func (s *VName) Copy() *VName {
	n := new(VName)
	n.Key = s.Key
	n.XUniqueId = s.XUniqueId
	n.XUniqueStr = s.XUniqueStr
	var tgs Tags
	for _, t := range s.Tags {
		tgs = append(tgs, t)
	}
	n.Tags = tgs
	n.Resolution = s.Resolution
	n.Ttl = s.Ttl
	return n
}

// VNameSlice list of VNames
type VNameSlice []*VName

// Sort by Key string
func (p VNameSlice) Len() int           { return len(p) }
func (p VNameSlice) Less(i, j int) bool { return strings.Compare(p[i].Key, p[j].Key) < 0 }
func (p VNameSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
