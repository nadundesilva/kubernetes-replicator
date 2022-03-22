/*
 * Copyright (c) 2022, Nadun De Silva. All Rights Reserved.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *   http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package kubernetes

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
)

func (c *client) SecretInformer() cache.SharedIndexInformer {
	return c.resourceInformerFactory.Core().V1().Secrets().Informer()
}

func (c *client) CreateSecret(ctx context.Context, namespace string, secret *corev1.Secret) (*corev1.Secret, error) {
	return c.clientset.CoreV1().Secrets(namespace).Create(ctx, secret, defaultCreateOptions)
}

func (c *client) GetSecret(ctx context.Context, namespace, name string) (*corev1.Secret, error) {
	return c.clientset.CoreV1().Secrets(namespace).Get(ctx, name, defaultGetOptions)
}
