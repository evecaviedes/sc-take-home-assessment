package folders_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
		}

		expected := folders.Folder{
			Id:      uuid.FromStringOrNil(folders.DefaultOrgID),
			Name:    "Evelyn",
			OrgId:   uuid.FromStringOrNil(folders.DefaultOrgID),
			Deleted: false,
		}
		expectedFolder := []*folders.Folder{}
		expectedFolder = append(expectedFolder, &expected)

		out, _ := folders.GetAllFolders(req, expectedFolder)

		if !reflect.DeepEqual(getObject(expectedFolder), getObject(out.Folders)) {
			t.Fatalf("the expected result %v is not be equal to what we obtained %v", expectedFolder, out)
		}
	})
}

func getObject(b interface{}) string {
	s, _ := json.MarshalIndent(b, "", "\t")
	return string(s)
}
