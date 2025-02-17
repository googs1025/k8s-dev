// FooSpec defines the desired state of Foo
type FooSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS -- desired state of cluster
}

// FooStatus defines the observed state of Foo.
// It should always be reconstructable from the state of the cluster and/or outside world.
type FooStatus struct {
	// INSERT ADDITIONAL STATUS FIELDS -- observed state of cluster
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is the Schema for the foos API
// +k8s:openapi-gen=true
type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FooSpec   `json:"spec,omitempty"`
	Status FooStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList contains a list of Foo
type FooList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Foo `json:"items"`
}
