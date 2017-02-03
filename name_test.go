package fvec

import (
	"github.com/wyndhblb/timeslab"
	"testing"
	"time"
)

func Test_Name_UniqueId(t *testing.T) {

	useT := time.Date(2008, 3, 2, 1, 23, 54, 0, time.UTC)
	vn := new(VName)
	vn.Key = "mighty_key"
	vn.Resolution = timeslab.Resolution_DAY
	vn.Time = useT.Unix()

	uid := vn.UniqueId()
	if uid != 16306963131860581163 {
		t.Fatalf("Wrong uid: wanted %d", uid)
	}

	if vn.UniqueIdString() != "3fw4vznzaml9n" {
		t.Fatalf("Wrong uid: wanted %s", vn.UniqueIdString())
	}

	vn.AddTagString("monkey", "funky")
	uid = vn.UniqueId()
	if uid != 9165573496101172186 {
		t.Fatalf("Wrong uid: wanted %d", uid)
	}

	if vn.UniqueIdString() != "1xmvwn40rsr8a" {
		t.Fatalf("Wrong uid: wanted %s", vn.UniqueIdString())
	}

}

func Test_Name_Slab(t *testing.T) {

	useT := time.Date(2008, 3, 2, 1, 23, 54, 0, time.UTC)
	vn := new(VName)
	vn.Key = "mighty_key"
	vn.Resolution = timeslab.Resolution_DAY
	vn.Time = useT.Unix()

	if vn.ToSlab(useT) != "20080302" {
		t.Fatalf("Wrong slab: %s", vn.ToSlab(useT))
	}
	if vn.TimeToSlab() != "20080302" {
		t.Fatalf("Wrong slab: %s", vn.TimeToSlab())
	}
}

func Test_Name_CassUpsertQuery(t *testing.T) {

	useT := time.Date(2008, 3, 2, 1, 23, 54, 0, time.UTC)
	vn := new(VName)
	vn.Key = "mighty_key"
	vn.Resolution = timeslab.Resolution_DAY
	vn.Time = useT.Unix()

	q := vn.CassUpsertQuery()

	if len(q.GetUpdateArgs()) != 1 {
		t.Fatalf("Not enough update args: %d", len(q.GetUpdateArgs()))
	}

	if len(q.GetWhereArgs()) != 4 {
		t.Fatalf("Not enough where args: %d", len(q.GetWhereArgs()))
	}

	if q.GetWhere() != "id=? AND slab=? AND ord=?" {
		t.Fatalf("Wrong where: %s", q.GetWhere())
	}

	if q.GetForUpdate() != "key=?" {
		t.Fatalf("Wrong where: %s", q.GetForUpdate())
	}

	// tags
	vn.Tags = make([]*Tag, 2)
	vn.Tags[0] = &Tag{Name: "tester", Value: "valll"}
	vn.Tags[1] = &Tag{Name: "tester2", Value: "valll2"}

	q = vn.CassUpsertQuery()

	if len(q.GetUpdateArgs()) != 2 {
		t.Fatalf("Not enough update args: %d", len(q.GetUpdateArgs()))
	}

	if len(q.GetWhereArgs()) != 4 {
		t.Fatalf("Not enough where args: %d", len(q.GetWhereArgs()))
	}

	if q.GetWhere() != "id=? AND slab=? AND ord=?" {
		t.Fatalf("Wrong where: %s", q.GetWhere())
	}

	if q.GetForUpdate() != "key=?, tags=?" {
		t.Fatalf("Wrong where: %s", q.GetForUpdate())
	}
}
