package controllers

import (
	"context"

	. "github.com/onsi/ginkgo"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Example Controller", func() {
	data := map[string][]byte{
		"username": []byte("user"),
		"password": []byte("password"),
	}

	sec := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-namespace",
			Name:      "sec",
		},
		Data: data,
	}

	k8sClient.Create(context.Background(), sec)

})
