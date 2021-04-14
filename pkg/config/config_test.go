package config

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const testYamlConfig = `
data:
  key_1: true
  key_2: 'some text'

arr:
  - item_1
  - item_2
`

func TestNew(t *testing.T) {
	tmpDir := os.TempDir()
	os.Mkdir(tmpDir, os.ModeDir)

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
		before  func()
		after   func()
	}{
		{
			name: "load config without file path",
			args: args{
				path: "",
			},
			want:    map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "load config from invalid file path",
			args: args{
				path: "someindalidpath",
			},
			want:    map[string]interface{}{},
			wantErr: true,
		},
		{
			name: "load config from yaml file",
			args: args{
				path: tmpDir + "/config.yaml",
			},
			want: map[string]interface{}{
				"data": map[string]interface{}{
					"key_1": true,
					"key_2": "some text",
				},
				"arr": []interface{}{"item_1", "item_2"},
			},
			wantErr: true,
			before: func() {
				ioutil.WriteFile(tmpDir+"/config.yaml", []byte(testYamlConfig), 0644)
			},
			after: func() {
				os.Remove(tmpDir + "/config.yaml")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before()
			}
			if tt.after != nil {
				defer tt.after()
			}

			inst, err := New(tt.args.path)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				}

				return
			}

			inst.Debug()

			var data map[string]interface{}
			inst.Unmarshal(&data)

			assert.Equal(t, data, tt.want)
		})
	}
}
