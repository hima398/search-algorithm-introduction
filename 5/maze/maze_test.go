package maze

import "testing"

func TestAlternateMazeState_String(t *testing.T) {
	type fields struct {
		points     [Height][Width]int
		turn       int
		characters []Character
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"odd",
			fields{[Height][Width]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0}},
				1,
				[]Character{
					{1, 2, 0},
					{1, 0, 0}}},
			"\nturn:\t1\nscore(0)\t0\t\ty: 1 x: 0\nscore(1)\t0\t\ty: 1 x: 2\n...\nA.B\n...\n"},
		{"even",
			fields{[Height][Width]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0}},
				2,
				[]Character{
					{1, 0, 0},
					{1, 2, 0}}},
			"\nturn:\t2\nscore(0)\t0\t\ty: 1 x: 0\nscore(1)\t0\t\ty: 1 x: 2\n...\nA.B\n...\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AlternateMazeState{
				points:     tt.fields.points,
				turn:       tt.fields.turn,
				characters: tt.fields.characters,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("AlternateMazeState.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
