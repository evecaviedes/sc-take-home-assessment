package folders

import (
	"github.com/gofrs/uuid"
)

func GetAllFolders(req *FetchFolderRequest, folderData []*Folder) (*FetchFolderResponse, error) {
	//declared but never used
	// var (
	// 	err error
	// 	f1  Folder
	// 	fs  []*Folder
	// )
	//Array of objects folder
	f := []Folder{}
	r, _ := FetchAllFoldersByOrgID(req.OrgID, folderData)
	//k is not the default index for a for loop, it should be _
	for _, v := range r {
		//Since it is a big amount of data it is better to user pointer in order to use the resources efficiently
		//It costs to pass the data more than just passing the address where the data is stored
		//append(f, *v) is assigned a value in the data memory location
		f = append(f, *v)
	}
	var fp []*Folder
	for _, v1 := range f {
		fp = append(fp, &v1)
	}
	var ffr *FetchFolderResponse
	//assigned value in the data memory location
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

func FetchAllFoldersByOrgID(orgID uuid.UUID, folderData []*Folder) ([]*Folder, error) {
	//This comes with sample from main
	folders := folderData

	resFolder := []*Folder{}
	for _, folder := range folders {
		//if folder in sample.json retrieve the orgID then it is added to the array
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
