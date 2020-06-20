package files_manager

import(
	"MinifyURL/shortener"
	"log"
	"os"
)

type fileManagerRepository struct{
	fileName string
	filePath string
}

func newFileManager(fileName string) *shortener.RedirectRepository{
	return &fileManagerRepository{fileName, "/redirects/" + fileName}
}

func (r *fileManagerRepository) Find(code string) (*Redirect, error)){

}

func (r *fileManagerRepository) Store(redirect *Redirect) error{
	file,err := os.Create(r.filePath)
	if err != nil{
		log.fatal("Failed to create file in repository")
		return err
	}
	
}