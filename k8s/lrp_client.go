package k8s

import (
	"code.cloudfoundry.org/eirini/k8s/stset"
	"code.cloudfoundry.org/eirini/opi"
	"code.cloudfoundry.org/lager"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type PodClient interface {
	GetAll() ([]corev1.Pod, error)
	GetByLRPIdentifier(opi.LRPIdentifier) ([]corev1.Pod, error)
	Delete(namespace, name string) error
}

type PodDisruptionBudgetClient interface {
	Update(namespace, name string, lrp *opi.LRP) error
	Delete(namespace string, name string) error
}

type StatefulSetClient interface {
	Create(namespace string, statefulSet *appsv1.StatefulSet) (*appsv1.StatefulSet, error)
	Update(namespace string, statefulSet *appsv1.StatefulSet) (*appsv1.StatefulSet, error)
	Delete(namespace string, name string) error
	GetBySourceType(sourceType string) ([]appsv1.StatefulSet, error)
	GetByLRPIdentifier(id opi.LRPIdentifier) ([]appsv1.StatefulSet, error)
}

type SecretsClient interface {
	Create(namespace string, secret *corev1.Secret) (*corev1.Secret, error)
	Delete(namespace string, name string) error
}

type EventsClient interface {
	GetByPod(pod corev1.Pod) ([]corev1.Event, error)
}

type LRPClient struct {
	stset.Desirer
	stset.Lister
	stset.Stopper
	stset.Updater
	stset.Getter
}

func NewLRPClient(
	logger lager.Logger,
	secrets SecretsClient,
	statefulSets StatefulSetClient,
	pods PodClient,
	pdbClient PodDisruptionBudgetClient,
	events EventsClient,
	lrpToStatefulSetConverter stset.LRPToStatefulSetConverter,
	statefulSetToLRPConverter stset.StatefulSetToLRPConverter,

) *LRPClient {
	return &LRPClient{
		Desirer: stset.NewDesirer(logger, secrets, statefulSets, lrpToStatefulSetConverter, pdbClient),
		Lister:  stset.NewLister(logger, statefulSets, statefulSetToLRPConverter),
		Stopper: stset.NewStopper(logger, statefulSets, statefulSets, pods, pdbClient, secrets),
		Updater: stset.NewUpdater(logger, statefulSets, statefulSets, pdbClient),
		Getter:  stset.NewGetter(logger, statefulSets, pods, events, statefulSetToLRPConverter),
	}
}
