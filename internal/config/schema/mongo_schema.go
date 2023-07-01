package schema

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type Package struct {
	ParentUserId string `bson:"parentUserId"`
	PackageName  string `bson:"packageName"`
}

type PackageVersion struct {
	ParentPackageId string `bson:"parentPackageId"`
	Version         string `bson:"version"`
	DownloadPath    string `bson:"downloadPath"`
}
