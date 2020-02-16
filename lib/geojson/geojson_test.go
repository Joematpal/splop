package geojson

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoadGeoJson(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":200}`))
	}))
	defer ts.Close()

	type args struct {
		url      string
		filePath string
	}
	tf, err := ioutil.TempFile("/tmp", "")
	td := `{"type": "feature", "data":[
{"test": "test0"},
{"test": "test1"},
{"test": "test2"},
]}`
	defer tf.Close()

	if _, err := tf.WriteString(td); err != nil {
		t.Fatal(err)
	}

	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success:",
			args: args{
				url:      ts.URL,
				filePath: tf.Name(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadGeoJson(tt.args.url, tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("LoadGeoJson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
