package main

import "testing"

func Test_applyReplacement(t *testing.T) {
	type args struct {
		fileContent string
		ranges      []codeRange
		replacement string
	}
	tests := []struct {
		name        string
		args        args
		wantNewCode string
		wantErr     bool
	}{
		{
			name: "replace same size symbol",
			args: args{
				replacement: "NewType",
				// ranges are unsorted on purpose, the tested code should handle the sorting
				ranges:      testRanges(),
				fileContent: twoTypesOneLine,
			},
			wantNewCode: twoTypesOneLineReplaced,
			wantErr:     false,
		},
		{
			name: "replace longer size symbol",
			args: args{
				replacement: "NewTypeLonger",
				// ranges are unsorted on purpose, the tested code should handle the sorting
				ranges:      testRanges(),
				fileContent: twoTypesOneLine,
			},
			wantNewCode: twoTypesOneLineReplacedWithLonger,
			wantErr:     false,
		},
		{
			name: "replace shorter size symbol",
			args: args{
				replacement: "New",
				// ranges are unsorted on purpose, the tested code should handle the sorting
				ranges:      testRanges(),
				fileContent: twoTypesOneLine,
			},
			wantNewCode: twoTypesOneLineReplacedWithShorter,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewCode, err := applyReplacement(tt.args.fileContent, tt.args.ranges, tt.args.replacement)
			if (err != nil) != tt.wantErr {
				t.Errorf("applyReplacement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNewCode != tt.wantNewCode {
				t.Errorf("applyReplacement() gotNewCode = %v, want %v", gotNewCode, tt.wantNewCode)
			}
		})
	}
}

func testRanges() []codeRange {
	return []codeRange{
		// replacement of 'b' parameter
		{
			start: codeLocation{
				line:      3,
				character: 22,
			},
			end: codeLocation{
				line:      3,
				character: 29,
			},
		},
		// replacement of 'a' parameter
		{
			start: codeLocation{
				line:      3,
				character: 11,
			},
			end: codeLocation{
				line:      3,
				character: 18,
			},
		},
	}
}

const twoTypesOneLine = `
package main

func foo(a OldType, b OldType) {
	println("hello")
}

type OldType = string
type NewType = string
type New = string
`

const twoTypesOneLineReplaced = `
package main

func foo(a NewType, b NewType) {
	println("hello")
}

type OldType = string
type NewType = string
type New = string
`

const twoTypesOneLineReplacedWithLonger = `
package main

func foo(a NewTypeLonger, b NewTypeLonger) {
	println("hello")
}

type OldType = string
type NewType = string
type New = string
`

const twoTypesOneLineReplacedWithShorter = `
package main

func foo(a New, b New) {
	println("hello")
}

type OldType = string
type NewType = string
type New = string
`

func Test_writeReplacement(t *testing.T) {
	type args struct {
		ranges      map[string][]codeRange
		replacement string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "replaces stuff in a single file",
			args: args{
				ranges: map[string][]codeRange{
					"cmd/renamer/sample_test_file.go": testRanges(),
				},
				replacement: "ReplacedStuff",
			},
		},
		//TODO more files
		//TODO err case
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeReplacement(tt.args.ranges, tt.args.replacement); (err != nil) != tt.wantErr {
				t.Errorf("writeReplacement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
