package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/golangdaddy/leap/sdk/cloudfunc"

	"github.com/dsoprea/go-jpeg-image-structure"
	"github.com/dsoprea/go-png-image-structure"
	"github.com/kr/pretty"
)

func (app *App) Upload{{uppercase .Object.Name}}(w http.ResponseWriter, r *http.Request, parent *Internals, user *User) {

	log.Println("PARSING FORM")
	if err := r.ParseMultipartForm(300 << 20); err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("file")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	buf := bytes.NewBuffer(nil)
	// Copy the uploaded file to the created file on the filesystem
	if n, err := io.Copy(buf, file); err != nil {
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return
	} else {
		log.Println("copy: wrote", n, "bytes")
	}

	newFileObjects := []*{{uppercase .Object.Name}}{}

	if !strings.Contains(strings.ToLower(handler.Filename), "zip") {

		obj, err := app.newUploadObject{{uppercase .Object.Name}}(parent, user, 0, handler.Filename, buf.Bytes())
		if err != nil {
			log.Println(err)
			return
		}

		newFileObjects = append(newFileObjects, obj)

	} else {

		log.Println("HANDLING ZIP")

		// Open the zip archive from the buffer
		zipReader, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
		if err != nil {
			err = fmt.Errorf("Error opening zip archive: %v", err)
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return 
		}

		// Extract each file from the zip archive
		for n, zipFile := range zipReader.File {
	
			extractedContent, err := readZipFile{{uppercase .Object.Name}}(zipFile)
			if err != nil {
				log.Println(err)
				continue
			}
	
			obj, err := app.newUploadObject{{uppercase .Object.Name}}(parent, user, n, zipFile.Name, extractedContent)
			if err != nil {
				log.Println(err)
				continue
			}

			newFileObjects = append(newFileObjects, obj)
		}
	}

	// make the documents proper
	for _, obj := range newFileObjects {
		if err := app.CreateDocument{{uppercase .Object.Name}}(parent, obj); err != nil {
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return		
		}
	}

	return
}

func (app *App) newUploadObject{{uppercase .Object.Name}}(parent *Internals, user *User, n int, name string, b []byte) (*{{uppercase .Object.Name}}, error) {
	fields := Fields{{uppercase .Object.Name}}{}
	object := user.New{{uppercase .Object.Name}}(parent, fields)
	object.Meta.Name = name
	object.Meta.Context.Order = n
	// generate a new URI
	uri := object.Meta.NewURI()
	log.Println(name, "URI", uri)

	// check if it is an image
	if err := checkImage{{uppercase .Object.Name}}(b); err != nil {
		fmt.Errorf("skipping file that cannot be decoded: %s", name)
		return nil, err
	} else {
		object.Meta.Media.Image = true
	}

	// determine image format
	if jpegstructure.NewJpegMediaParser().LooksLikeFormat(b) {
		parser := jpegstructure.NewJpegMediaParser()
		mctx, err := parser.ParseBytes(b)
		if err != nil {
			return nil, err
		}
		pretty.Println(mctx)
		_, b, err  := mctx.Exif()
		if err != nil {
			return nil, err
		}
		object.Meta.Media.Format = "JPEG"
	} else {
		if pngstructure.NewPngMediaParser().LooksLikeFormat(b) {
			object.Meta.Media.Format = "PNG"
		}
	}

	if err := app.write{{titlecase .Object.Name}}File(CONST_BUCKET_UPLOADS, uri, b); err != nil {
		return nil, err
	}

	return object, nil
}

// assert file is an image because of .Object.Options.Image
func checkImage{{uppercase .Object.Name}}(fileBytes []byte) error {
	_, _, err := image.Decode(bytes.NewBuffer(fileBytes))
	return err
}

func readZipFile{{uppercase .Object.Name}}(zipFile *zip.File) ([]byte, error) {
	// Open the file from the zip archive
	zipFileReader, err := zipFile.Open()
	if err != nil {
		return nil, fmt.Errorf("Error opening zip file entry: %v", err)
	}
	defer zipFileReader.Close()

	// Read the content of the file from the zip archive
	var extractedContent bytes.Buffer
	if _, err := io.Copy(&extractedContent, zipFileReader); err != nil {
		return nil, fmt.Errorf("Error reading zip file entry content: %v", err)
	}

	return extractedContent.Bytes(), nil
}

func (app *App) write{{titlecase .Object.Name}}File(bucketName, objectName string, content []byte) error {
	writer := app.GCPClients.GCS().Bucket(bucketName).Object(objectName).NewWriter(app.Context())
	//writer.ObjectAttrs.CacheControl = "no-store"
	defer writer.Close()
	n, err := writer.Write(content)
	fmt.Printf("wrote %s %d bytes to bucket: %s \n", objectName, n, bucketName)
	return err
}
