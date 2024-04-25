package v1alpha1

import corev1 "k8s.io/api/core/v1"

type RegionServerSpec struct {
	// +kubebuilder:validation:Optional
	Config *RegionConfigSpec `json:"config,omitempty"`

	// +kubebuilder:validation:Required
	RoleGroups map[string]RegonRoleGroupSpec `json:"roleGroups,omitempty"`

	PodDisruptionBudget *PodDisruptionBudgetSpec `json:"podDisruptionBudget,omitempty"`

	// +kubebuilder:validation:Optional
	CommandArgsOverrides []string `json:"commandArgsOverrides,omitempty"`

	// +kubebuilder:validation:Optional
	ConfigOverrides *RegionConfigOverrideSpec `json:"configOverrides,omitempty"`

	// +kubebuilder:validation:Optional
	EnvOverrides map[string]string `json:"envOverrides,omitempty"`

	// // +kubebuilder:validation:Optional
	// PodOverride *corev1.PodTemplateSpec `json:"podOverride,omitempty"`
}

type RegionConfigSpec struct {
	// +kubebuilder:validation:Optional
	Affinity *corev1.Affinity `json:"affinity"`

	// +kubebuilder:validation:Optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// +kubebuilder:validation:Optional
	Tolerations []corev1.Toleration `json:"tolerations"`

	// +kubebuilder:validation:Optional
	PodDisruptionBudget *PodDisruptionBudgetSpec `json:"podDisruptionBudget,omitempty"`

	// Use time.ParseDuration to parse the string
	// +kubebuilder:validation:Optional
	GracefulShutdownTimeout *string `json:"gracefulShutdownTimeout,omitempty"`

	// +kubebuilder:validation:Optional
	Logging *ContainerLoggingSpec `json:"logging,omitempty"`

	// +kubebuilder:validation:Optional
	Resources *ResourcesSpec `json:"resources,omitempty"`
}

type RegionConfigOverrideSpec struct{}

type RegonRoleGroupSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=1
	Replicas int32 `json:"replicas,omitempty"`

	// +kubebuilder:validation:Optional
	Config *MasterConfigSpec `json:"config,omitempty"`

	// +kubebuilder:validation:Required
	PodDisruptionBudget *PodDisruptionBudgetSpec `json:"podDisruptionBudget,omitempty"`

	// +kubebuilder:validation:Optional
	CommandArgsOverrides []string `json:"commandArgsOverrides,omitempty"`

	// +kubebuilder:validation:Optional
	ConfigOverrides *RegionConfigOverrideSpec `json:"configOverrides,omitempty"`

	// +kubebuilder:validation:Optional
	EnvOverrides map[string]string `json:"envOverrides,omitempty"`

	// // +kubebuilder:validation:Optional
	// PodOverride *corev1.PodTemplateSpec `json:"podOverride,omitempty"`
}
