package master

import (
	"context"

	hbasev1alph1 "github.com/zncdatadev/hbase-operator/api/v1alpha1"
	"github.com/zncdatadev/hbase-operator/internal/controller/common"
	"github.com/zncdatadev/operator-go/pkg/client"
	"github.com/zncdatadev/operator-go/pkg/reconciler"
	"github.com/zncdatadev/operator-go/pkg/util"
	ctrl "sigs.k8s.io/controller-runtime"

	commonsv1alpha1 "github.com/zncdatadev/operator-go/pkg/apis/commons/v1alpha1"
)

var (
	logger = ctrl.Log.WithName("controller").WithName("master")
)

var _ reconciler.RoleReconciler = &Reconciler{}

type Reconciler struct {
	reconciler.BaseRoleReconciler[*hbasev1alph1.MasterSpec]
	ClusterConfig *hbasev1alph1.ClusterConfigSpec
	Image         *util.Image
}

func NewReconciler(
	client *client.Client,
	roleInfo reconciler.RoleInfo,
	clusterOperation *commonsv1alpha1.ClusterOperationSpec,
	clusterConfig *hbasev1alph1.ClusterConfigSpec,
	image *util.Image,
	spec *hbasev1alph1.MasterSpec,
) *Reconciler {

	return &Reconciler{
		BaseRoleReconciler: *reconciler.NewBaseRoleReconciler(
			client,
			roleInfo,
			clusterOperation,
			spec,
		),
		Image:         image,
		ClusterConfig: clusterConfig,
	}
}

func (r *Reconciler) RegisterResources(ctx context.Context) error {
	for name, roleGroup := range r.Spec.RoleGroups {
		mergedRoleGroup := r.MergeRoleGroupSpec(&roleGroup)

		info := reconciler.RoleGroupInfo{
			RoleInfo:      r.RoleInfo,
			RoleGroupName: name,
		}

		reconcilers, err := r.RegisterResourceWithRoleGroup(ctx, info, mergedRoleGroup)
		if err != nil {
			return err
		}

		for _, reconciler := range reconcilers {
			r.AddResource(reconciler)
			logger.Info("registered resource", "role", r.GetName(), "roleGroup", name, "reconciler", reconciler.GetName())
		}

	}
	return nil
}

func (r *Reconciler) RegisterResourceWithRoleGroup(_ context.Context, info reconciler.RoleGroupInfo, roleGroupSpec any) ([]reconciler.Reconciler, error) {

	var reconcilers []reconciler.Reconciler

	stopped := false

	if r.ClusterOperation != nil && r.ClusterOperation.Stopped {
		stopped = true
	}

	statefulSetReconciler, err := NewStatefulSetReconciler(
		r.Client,
		r.ClusterConfig,
		info,
		Ports,
		r.Image,
		stopped,
		roleGroupSpec.(*hbasev1alph1.MasterRoleGroupSpec),
	)

	if err != nil {
		return nil, err
	}

	reconcilers = append(reconcilers, statefulSetReconciler)

	serviceReconciler := common.NewServiceReconciler(
		r.GetClient(),
		Ports,
		info,
	)

	reconcilers = append(reconcilers, serviceReconciler)

	configMapReconciler := common.NewConfigMapReconciler(
		r.GetClient(),
		r.ClusterConfig,
		info,
		roleGroupSpec.(*hbasev1alph1.MasterRoleGroupSpec),
	)
	reconcilers = append(reconcilers, configMapReconciler)

	return reconcilers, nil
}
