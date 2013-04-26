package git

import "C"

import (
  "unsafe"
)

func Clone(url, path string) (*Repository, error) {
  curl := C.CString(url)
  defer C.free(unsafe.Pointer(curl))

  cpath := C.CString(path)
  defer C.free(unsafe.Pointer(cpath))

  ret := C.git_clone(curl, cpath)

  if ret < 0 {
    return nil, LastError()
  }

  return OpenRepository(url)
}
