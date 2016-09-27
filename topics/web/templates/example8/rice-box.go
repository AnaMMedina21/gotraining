package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `app.js`,
		FileModTime: time.Unix(1475013082, 0),
		Content:     string([]byte{0x68, 0x31, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x28, 0x22, 0x68, 0x31, 0x22, 0x29, 0x5b, 0x30, 0x5d, 0x3b, 0xa, 0x68, 0x31, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x22, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x3b, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    `index.html`,
		FileModTime: time.Unix(1475013082, 0),
		Content:     string([]byte{0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x6d, 0x65, 0x74, 0x61, 0x20, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x3d, 0x22, 0x75, 0x74, 0x66, 0x2d, 0x38, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x55, 0x6c, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x20, 0x57, 0x65, 0x62, 0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x22, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x2e, 0x63, 0x73, 0x73, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x63, 0x73, 0x73, 0x22, 0x20, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x3d, 0x22, 0x61, 0x6c, 0x6c, 0x22, 0x20, 0x2f, 0x3e, 0xa, 0x3c, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x68, 0x31, 0x3e, 0x3c, 0x2f, 0x68, 0x31, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x20, 0x73, 0x72, 0x63, 0x3d, 0x22, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x6a, 0x73, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x6a, 0x61, 0x76, 0x61, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x20, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x3d, 0x22, 0x75, 0x74, 0x66, 0x2d, 0x38, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0xa, 0x3c, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    `styles.css`,
		FileModTime: time.Unix(1475013082, 0),
		Content:     string([]byte{0x68, 0x31, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x62, 0x6f, 0x6c, 0x64, 0x3b, 0xa, 0x20, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x62, 0x6c, 0x75, 0x65, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x20, 0x34, 0x38, 0x70, 0x78, 0x3b, 0xa, 0x7d, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1475013082, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // app.js
			file3, // index.html
			file4, // styles.css

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`./static`, &embedded.EmbeddedBox{
		Name: `./static`,
		Time: time.Unix(1475013082, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"app.js":     file2,
			"index.html": file3,
			"styles.css": file4,
		},
	})
}
