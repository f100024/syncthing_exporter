package collector

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/prometheus/common/promlog"
)

func TestNewDBStatusReport(t *testing.T) {

	jsonResponse, _ := os.ReadFile("fixtures/rest_db_status_response.json")

	ts := httptest.NewTLSServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, string(jsonResponse))
			},
		),
	)
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Errorf("url parse error: %s", err)
	}

	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)

	testToken := "12345"
	testFoldersid := []string{"aaaaa-bb11b"}

	expected := `
	# HELP syncthing_rest_db_status_errors Number of errors for current folder.
	# TYPE syncthing_rest_db_status_errors gauge
	syncthing_rest_db_status_errors{folderID="aaaaa-bb11b"} 0
	# HELP syncthing_rest_db_status_global_bytes Number of bytes globally.
	# TYPE syncthing_rest_db_status_global_bytes gauge
	syncthing_rest_db_status_global_bytes{folderID="aaaaa-bb11b"} 4.047955455069e+12
	# HELP syncthing_rest_db_status_global_deleted Number of bytes deleted.
	# TYPE syncthing_rest_db_status_global_deleted gauge
	syncthing_rest_db_status_global_deleted{folderID="aaaaa-bb11b"} 156112
	# HELP syncthing_rest_db_status_global_directories Number of directories globally.
	# TYPE syncthing_rest_db_status_global_directories gauge
	syncthing_rest_db_status_global_directories{folderID="aaaaa-bb11b"} 346372
	# HELP syncthing_rest_db_status_global_symlinks Number of symlinks globally.
	# TYPE syncthing_rest_db_status_global_symlinks gauge
	syncthing_rest_db_status_global_symlinks{folderID="aaaaa-bb11b"} 0
	# HELP syncthing_rest_db_status_global_total_items Number of total items globally.
	# TYPE syncthing_rest_db_status_global_total_items gauge
	syncthing_rest_db_status_global_total_items{folderID="aaaaa-bb11b"} 1.197108e+06
	# HELP syncthing_rest_db_status_ignore_patterns Is using ignore patterns.
	# TYPE syncthing_rest_db_status_ignore_patterns gauge
	syncthing_rest_db_status_ignore_patterns{folderID="aaaaa-bb11b"} 0
	# HELP syncthing_rest_db_status_insync_bytes Number of bytes currently in sync.
	# TYPE syncthing_rest_db_status_insync_bytes gauge
	syncthing_rest_db_status_insync_bytes{folderID="aaaaa-bb11b"} 4.047955455069e+12
	# HELP syncthing_rest_db_status_insync_files Number of files currently in sync.
	# TYPE syncthing_rest_db_status_insync_files gauge
	syncthing_rest_db_status_insync_files{folderID="aaaaa-bb11b"} 694624
	# HELP syncthing_rest_db_status_json_parse_failures Number of errors while parsing JSON.
	# TYPE syncthing_rest_db_status_json_parse_failures counter
	syncthing_rest_db_status_json_parse_failures 0
	# HELP syncthing_rest_db_status_local_bytes Number of bytes locally.
	# TYPE syncthing_rest_db_status_local_bytes gauge
	syncthing_rest_db_status_local_bytes{folderID="aaaaa-bb11b"} 4.047955455069e+12
	# HELP syncthing_rest_db_status_local_deleted Number of bytes deleted locally.
	# TYPE syncthing_rest_db_status_local_deleted gauge
	syncthing_rest_db_status_local_deleted{folderID="aaaaa-bb11b"} 318
	# HELP syncthing_rest_db_status_local_directories Number of local directories.
	# TYPE syncthing_rest_db_status_local_directories gauge
	syncthing_rest_db_status_local_directories{folderID="aaaaa-bb11b"} 346372
	# HELP syncthing_rest_db_status_local_symlinks Number of local symlinks
	# TYPE syncthing_rest_db_status_local_symlinks gauge
	syncthing_rest_db_status_local_symlinks{folderID="aaaaa-bb11b"} 0
	# HELP syncthing_rest_db_status_local_total_items Number of total items locally
	# TYPE syncthing_rest_db_status_local_total_items gauge
	syncthing_rest_db_status_local_total_items{folderID="aaaaa-bb11b"} 1.041314e+06
	# HELP syncthing_rest_db_status_need_bytes Number of bytes need for sync.
	# TYPE syncthing_rest_db_status_need_bytes gauge
	syncthing_rest_db_status_need_bytes{folderID="aaaaa-bb11b"} 0
	# HELP syncthing_rest_db_status_need_deletes Number of bytes need for deletes.
	# TYPE syncthing_rest_db_status_need_deletes gauge
	syncthing_rest_db_status_need_deletes{folderID="aaaaa-bb11b"} 0
	# HELP syncthing_rest_db_status_need_directories Number of directories for sync.
	# TYPE syncthing_rest_db_status_need_directories gauge
	syncthing_rest_db_status_need_directories{folderID="aaaaa-bb11b"} 0
	# HELP syncthing_rest_db_status_need_symlinks Number of symlinks need for sync.
	# TYPE syncthing_rest_db_status_need_symlinks gauge
	syncthing_rest_db_status_need_symlinks{folderID="aaaaa-bb11b"} 0
	# HELP syncthing_rest_db_status_need_total_items Number of total items need to sync.
	# TYPE syncthing_rest_db_status_need_total_items gauge
	syncthing_rest_db_status_need_total_items{folderID="aaaaa-bb11b"} 0
	# HELP syncthing_rest_db_status_pull_errors Number of pull errors.
	# TYPE syncthing_rest_db_status_pull_errors gauge
	syncthing_rest_db_status_pull_errors{folderID="aaaaa-bb11b"} 0
	# HELP syncthing_rest_db_status_sequence Total bytes received from remote device.
	# TYPE syncthing_rest_db_status_sequence gauge
	syncthing_rest_db_status_sequence{folderID="aaaaa-bb11b",state="idle",stateChanged="2077-02-06T00:02:00-07:00"} 1.213411e+06
	# HELP syncthing_rest_db_status_total_scrapes Current total Syncthings scrapes.
	# TYPE syncthing_rest_db_status_total_scrapes counter
	syncthing_rest_db_status_total_scrapes 1
	`

	err = testutil.CollectAndCompare(
		NewDBStatusReport(logger, HttpClient, u, &testToken, &testFoldersid),
		strings.NewReader(expected),
	)

	if err != nil {
		t.Errorf("NewDBStatusReport %s", err)
	}
}
