package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Images struct {
	*metav1.TypeMeta  `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ImageSpec `json:"spec,omitempty"`
}

type ImageSpec struct {
	Images []Image `json:"images,omitempty"`
}

type Image struct {
	// Name is the full location for the image, can be referenced by tags or digests
	Name string `json:"name"`

	// Path is the path to the cosign public key used for verifying image signatures
	//Key string `json:"key,omitempty"`
	Key string `json:"key"`

	// Path is the path to the cosign public key used for verifying image signatures
	//Tlog string `json:"use-tlog-verify,omitempty"`
	Tlog bool `json:"use-tlog-verify"`

	// Platform of the image to be pulled.  If not specified, all platforms will be pulled.
	//Platform string `json:"key,omitempty"`
	Platform string `json:"platform"`
}
