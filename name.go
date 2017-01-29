/*
   Vector Name Object methods

   The base objects structs are in vepr.proto and generated
*/

package vepr

import (
	"fmt"
	"github.com/wyndhblb/go-utils/pools"
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
	switch s.Resolution {
	case Resolution_MIN:
		return t.Format("200601021504")
	case Resolution_MIN10:
		m := t.Minute() / 10
		return t.Format("2006010215") + "I10" + strconv.Itoa(m)
	case Resolution_MIN30:
		m := t.Minute() / 30
		return t.Format("2006010215") + "I30" + strconv.Itoa(m)
	case Resolution_HOUR:
		return t.Format("2006010215")
	case Resolution_DAY:
		return t.Format("20060102")
	case Resolution_WEEK:
		ynum, wnum := t.ISOWeek()
		return fmt.Sprintf("%04d%02d", ynum, wnum)
	case Resolution_MONTH:
		return t.Format("200601")
	case Resolution_MONTH2:
		m := (int(t.Month()) / 2)
		return t.Format("2006") + "M2" + strconv.Itoa(m)
	case Resolution_MONTH3:
		m := (int(t.Month()) / 3)
		return t.Format("2006") + "M3" + strconv.Itoa(m)
	case Resolution_MONTH6:
		m := (int(t.Month()) / 6)
		return t.Format("2006") + "M6" + strconv.Itoa(m)
	case Resolution_YEAR:
		return t.Format("2006")
	case Resolution_ALL:
		return "ALL"
	default:
		return t.Format("2006010215")
	}
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
