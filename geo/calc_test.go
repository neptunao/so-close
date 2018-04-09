package geo

import (
	"reflect"
	"testing"

	"github.com/neptunao/so-close/data"
)

type TextIterator struct {
	pos  int
	text [][]string
	err  error
}

func (itr *TextIterator) setErr(err error) {
	itr.err = err
}

func (itr *TextIterator) Next() (interface{}, bool) {
	if itr.pos >= len(itr.text) {
		return nil, false
	}
	res := itr.text[itr.pos]
	itr.pos++
	return res, true
}

func (itr *TextIterator) Err() error {
	return itr.err
}

func (itr *TextIterator) Close() error {
	return nil
}

func Test_geoHeapTop(t *testing.T) {
	type args struct {
		h     *FixedSizeHeap
		count int
	}
	tests := []struct {
		name string
		args args
		want []Coord
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := geoHeapTop(tt.args.h, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("geoHeapTop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcTopPoints(t *testing.T) {
	type args struct {
		center      Coord
		resultCount int
		itr         data.Iterator
	}
	tests := []struct {
		name    string
		args    args
		wantMin []Coord
		wantMax []Coord
		wantErr bool
	}{
		{
			name: "resultCount is greater then data in iterator",
			args: args{
				resultCount: 1,
				itr:         &TextIterator{},
			},
			wantErr: true,
		},
		{
			name: "top 1 test with one record",
			args: args{
				resultCount: 1,
				center:      Coord{"", 50, 50},
				itr: &TextIterator{
					text: [][]string{
						[]string{"", "51", "51"},
					},
				},
			},
			wantMin: []Coord{Coord{"", 51, 51}},
			wantMax: []Coord{Coord{"", 51, 51}},
			wantErr: false,
		},
		{
			name: "top 1 test with two record",
			args: args{
				resultCount: 1,
				center:      Coord{"", 50, 50},
				itr: &TextIterator{
					text: [][]string{
						[]string{"", "51", "51"},
						[]string{"", "49.999", "49.999"},
					},
				},
			},
			wantMin: []Coord{Coord{"", 49.999, 49.999}},
			wantMax: []Coord{Coord{"", 51, 51}},
			wantErr: false,
		},
		{
			// reference data from http://www.gpsvisualizer.com
			name: "top 3 complex test",
			args: args{
				resultCount: 3,
				center:      Coord{"center", 52.32161250000001, 4.953189800000001},
				itr: &TextIterator{
					text: [][]string{
						[]string{"505868", "52.09791479999999808", "5.11686619999999959"},
						[]string{"381769", "52.2934316", "4.9934547"},
						[]string{"419117", "48.8653063", "2.3794788"},
						[]string{"23928", "50.8651879", "5.707368199999999"},
						[]string{"1049729", "52.37165040000000005", "4.90306019999999965"},
						[]string{"90872", "52.06214569999999", "4.235672099999999"},
						[]string{"492800", "52.0624873", "5.273607"},
						[]string{"41238", "51.9056776", "4.454951299999999"},
						[]string{"488611", "51.89783940000000229", "4.51041999999999987"},
						[]string{"636618", "50.853433", "5.6841425"},
						[]string{"22310", "51.9214966", "4.506075"},
					},
				},
			},
			wantMin: []Coord{
				Coord{"381769", 52.2934316, 4.9934547},
				Coord{"1049729", 52.37165040000000005, 4.90306019999999965},
				Coord{"505868", 52.09791479999999808, 5.11686619999999959},
			},
			wantMax: []Coord{
				Coord{"419117", 48.8653063, 2.3794788},
				Coord{"636618", 50.853433, 5.6841425},
				Coord{"23928", 50.8651879, 5.707368199999999},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax, err := CalcTopPoints(tt.args.center, tt.args.resultCount, tt.args.itr)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcTopPoints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMin, tt.wantMin) {
				t.Errorf("CalcTopPoints() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if !reflect.DeepEqual(gotMax, tt.wantMax) {
				t.Errorf("CalcTopPoints() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
