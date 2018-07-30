package charts

import "testing"

var testChart = &Chart{
	ID:   "test1",
	Opts: Opts{Title: "Test1", Units: "test"},
	Dims: Dims{
		{ID: "dim1"},
	},
	Vars: Vars{
		{ID: "var1"},
	},
}

func TestChart_Copy(t *testing.T) {

	c1 := &Chart{
		ID:   "test1",
		Opts: Opts{Title: "Test1"},
		Dims: Dims{
			{ID: "dim1"},
		},
	}

	c2 := c1.Copy()
	c2.ID = "test2"
	c2.Opts.Title = "Test2"
	c2.Dims[0].ID = "dim2"
	c2.Dims = append(c2.Dims, &Dim{ID: "dim3"})

	if c1.ID == c2.ID || c1.Opts == c2.Opts || c1.Dims[0].ID == c2.Dims[0].ID || len(c1.Dims) == len(c2.Dims) {
		t.Error("expected full copy, but got partial")
	}
}

func TestChart_AddDim(t *testing.T) {
	c := testChart.Copy()

	c.AddDim(&Dim{ID: "dim1"})

	if len(c.Dims) != 1 {
		t.Errorf("expected 1 dimensions, but got %d", len(c.Dims))
	}

	c.AddDim(&Dim{ID: "dim2"})

	if len(c.Dims) != 2 {
		t.Errorf("expected 2 dimensions, but got %d", len(c.Dims))
	}
}

func TestChart_AddVar(t *testing.T) {
	c := testChart.Copy()

	c.AddVar(&Var{ID: "var1"})

	if len(c.Vars) != 1 {
		t.Errorf("expected 1 variables, but got %d", len(c.Vars))
	}

	c.AddVar(&Var{ID: "var2"})

	if len(c.Vars) != 2 {
		t.Errorf("expected 2 variables, but got %d", len(c.Vars))
	}
}

func TestChart_DeleteDimByID(t *testing.T) {
	c := testChart.Copy()

	c.DeleteDimByID("dim1")

	if len(c.Dims) != 0 {
		t.Errorf("expected 0 dimensions, but got %d", len(c.Dims))
	}

}

func TestChart_GetDimByID(t *testing.T) {
	c := testChart.Copy()
	d := c.Dims[0]

	v := c.GetDimByID("dim1")

	if v != d {
		t.Errorf("expected %v, but got %v", d, v)
	}

	if _, ok := interface{}(v).(*Dim); !ok {
		t.Error("expected *Dim")
	}

	if v := c.GetDimByID("dim2"); v != nil {
		t.Errorf("expected nil, but got %v", v)
	}
}

func TestChart_LookupDimByID(t *testing.T) {
	c := testChart.Copy()
	d := c.Dims[0]

	v, ok := c.LookupDimByID("dim1")

	if v != d || !ok {
		t.Errorf("expected %v and true, but got %v and %v", d, v, ok)
	}

	if _, ok := interface{}(v).(*Dim); !ok {
		t.Error("expected *Dim")
	}

	if v, ok := c.LookupDimByID("dim2"); v != nil || ok {
		t.Errorf("expected nil and false, but got %v and %v", v, ok)
	}

}

func TestChart_Refresh(t *testing.T) {
	testChart.Refresh()
}

func Test_ChartType_Algorithm_Hidden(t *testing.T) {
	if Line.String() != "line" || Area.String() != "area" || Stacked.String() != "stacked" {
		t.Error("wrong chart type")
	}

	if Absolute.String() != "absolute" || Incremental.String() != "incremental" {
		t.Error("wrong dimension algorithm")
	}

	if PercentOfAbsolute.String() != "percentage-of-absolute-row" || PercentOfIncremental.String() != "percentage-of-incremental-row" {
		t.Error("wrong dimension algorithm")
	}

	if Hidden.String() != "hidden" || NotHidden.String() != "" {
		t.Error("wrong dimension hidden")
	}
}