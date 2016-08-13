package mnemosyne

import (
	"reflect"
	"testing"

	"github.com/piotrkowalczuk/mnemosyne/mnemosyned"
	"golang.org/x/net/context"
)

func TestMnemosyne_Start(t *testing.T) {
	ctx := context.Background()
	subjectID := "1"
	subjectClient := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36"
	bag := map[string]string{"username": "johnsnow@gmail.com"}

	addr, closer := mnemosyned.TestDaemon(t, mnemosyned.TestDaemonOpts{
		StoragePostgresAddress: testPostgresAddress,
	})
	defer closer.Close()

	m, err := New(MnemosyneOpts{
		Addresses: []string{addr.String()},
	})
	defer m.Close()

	if err != nil {
		t.Fatalf("unexpected client initialization error: %s", err.Error())
	}
	res, err := m.Start(ctx, subjectID, subjectClient, bag)
	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if res.AccessToken == "" {
		t.Error("access token cannot be nil")
	}
	if len(res.AccessToken) == 0 {
		t.Error("access token should not be empty")
	}
	if res.ExpireAt.Seconds == 0 || res.ExpireAt.Nanos == 0 {
		t.Error("expire at should not be zero value")
	}
	if res.SubjectId != subjectID {
		t.Errorf("wrong subject id, expected %s but got %s", res.SubjectId, subjectID)
	}
	if res.SubjectClient != subjectClient {
		t.Errorf("wrong subject client, expected %s but got %s", res.SubjectClient, subjectClient)
	}
	if !reflect.DeepEqual(res.Bag, bag) {
		t.Errorf("wrong bag, expected:\n%s\nbut got:\n %s", res.Bag, bag)
	}

}
