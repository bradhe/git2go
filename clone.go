package git

/*
#cgo pkg-config: libgit2
#include <git2.h>
#include <git2/errors.h>
git_clone_options git_clone_options_init() {
  git_clone_options ret = GIT_CLONE_OPTIONS_INIT;
  return ret;
}
*/
import "C"

import (
  "unsafe"
)

type CloneOpts struct {
  // The name given to the "origin" remote. The default is "origin"
  RemoteName string

  // The URL to be used for pushing. The default is the fetch URL
  PushURL string

  // The fetch specification to be used for pushing.
  PushSpec string

  // The callback to be used if credentials are required furing the initial fetch.
  // CredAcquire func(cred *C.git_cred, url, usernameFromURL, string, allowedTypes uint, payload interface{}) int

  // A custom transport to be used for the initial fetch. NULL means it will be
  // autodetected from the URL.
  // Transport *C.git_transport

  // The nameo f hte branch to checkout. NULL means use the remote's HEAD.
  CheckoutBranch string

  // TODO: Add support for the following.

  // RemoteCallbacks
  // RemoteAutotag
}

func populateCloneOptions(ptr *C.git_clone_options, opts *CloneOpts) {

}

func Clone(url, path string, opts *CloneOpts) (*Repository, error) {
  var copts C.git_clone_options
  populateCloneOptions(&copts, opts)

  rep := new(Repository)

  curl := C.CString(url)
  defer C.free(unsafe.Pointer(curl))

  cpath := C.CString(path)
  defer C.free(unsafe.Pointer(cpath))

  ret := C.git_clone(&rep.ptr, curl, cpath, &copts)

  if ret < 0 {
    return nil, LastError()
  }

  return OpenRepository(url)
}
