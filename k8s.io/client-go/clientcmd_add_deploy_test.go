package client_go

import (
	"fmt"
	appV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"testing"
)

func TestCreateDeployment(t *testing.T) {
	clientset := getClientset()

	cases := []struct {
		name      string
		namespace string
	}{
		{"ttt", "pvt"},
	}

	for _, c := range cases {
		ns := c.namespace
		name := c.name
		labels := map[string]string{"app": "ttt"}
		deployment := getDeployment(name, ns, labels)
		log.Println(deployment.Name)
		client := clientset.AppsV1().Deployments(ns)
		result, err := client.Create(deployment)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result.GetObjectMeta().GetName())
		t.Log(result)

	}
}

func getDeployment(name, namespace string, labels map[string]string) *appV1.Deployment {
	sshVolume := getVolume()
	deployObjectMeta := metav1.ObjectMeta{
		Name:      name,
		Namespace: namespace,
		Labels:    labels,
		Annotations: map[string]string{
			"refCount": "1",
		},
	}
	args := []string{}
	//args := []string{"--debug"}
	image := "alpine"
	podSpec := v1.PodSpec{
		Containers: []v1.Container{
			container(image, args),
		},
		Volumes: []v1.Volume{
			sshVolume,
		},
	}
	podObjectMeta := metav1.ObjectMeta{
		Labels: labels,
	}
	template := v1.PodTemplateSpec{
		ObjectMeta: podObjectMeta,
		Spec:       podSpec,
	}
	selector := &metav1.LabelSelector{
		MatchLabels: labels,
	}
	deploySpec := appV1.DeploymentSpec{
		Selector: selector,
		Template: template,
	}
	deploy := &appV1.Deployment{
		ObjectMeta: deployObjectMeta,
		Spec:       deploySpec,
	}
	return deploy
}

func getVolume() v1.Volume {
	sshVolume := v1.Volume{
		Name: "ssh-public-key",
		VolumeSource: v1.VolumeSource{
			ConfigMap: &v1.ConfigMapVolumeSource{
				LocalObjectReference: v1.LocalObjectReference{
					Name: "volume-name",
				},
				Items: []v1.KeyToPath{
					{
						Key:  "authorized",
						Path: "authorized_keys",
					},
				},
			},
		},
	}
	return sshVolume
}

func container(image string, args []string) v1.Container {
	return v1.Container{
		Name:            "standalone",
		Image:           image,
		ImagePullPolicy: "Always",
		Command:         []string{"printenv"},
		Args:            args,
		VolumeMounts: []v1.VolumeMount{
			{
				Name:      "ssh-public-key",
				MountPath: fmt.Sprintf("/root/%s", SSHAuthKey),
			},
		},
	}
}
