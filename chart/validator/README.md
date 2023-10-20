
Validator
===========

validator monitors ValidationResults created by one or more validator plugins and uploads them to a configurable sink


## Configuration

The following table lists the configurable parameters of the Validator chart and their default values.

| Parameter                | Description             | Default        |
| ------------------------ | ----------------------- | -------------- |
| `controllerManager.kubeRbacProxy.args` |  | `["--secure-listen-address=0.0.0.0:8443", "--upstream=http://127.0.0.1:8080/", "--logtostderr=true", "--v=0"]` |
| `controllerManager.kubeRbacProxy.containerSecurityContext.allowPrivilegeEscalation` |  | `false` |
| `controllerManager.kubeRbacProxy.containerSecurityContext.capabilities.drop` |  | `["ALL"]` |
| `controllerManager.kubeRbacProxy.image.repository` |  | `"gcr.io/kubebuilder/kube-rbac-proxy"` |
| `controllerManager.kubeRbacProxy.image.tag` |  | `"v0.14.1"` |
| `controllerManager.kubeRbacProxy.resources.limits.cpu` |  | `"500m"` |
| `controllerManager.kubeRbacProxy.resources.limits.memory` |  | `"128Mi"` |
| `controllerManager.kubeRbacProxy.resources.requests.cpu` |  | `"5m"` |
| `controllerManager.kubeRbacProxy.resources.requests.memory` |  | `"64Mi"` |
| `controllerManager.manager.args` |  | `["--health-probe-bind-address=:8081", "--metrics-bind-address=127.0.0.1:8080", "--leader-elect"]` |
| `controllerManager.manager.containerSecurityContext.allowPrivilegeEscalation` |  | `false` |
| `controllerManager.manager.containerSecurityContext.capabilities.drop` |  | `["ALL"]` |
| `controllerManager.manager.image.repository` |  | `"quay.io/spectrocloud-labs/validator"` |
| `controllerManager.manager.image.tag` | x-release-please-version | `"v0.0.10"` |
| `controllerManager.manager.resources.limits.cpu` |  | `"500m"` |
| `controllerManager.manager.resources.limits.memory` |  | `"128Mi"` |
| `controllerManager.manager.resources.requests.cpu` |  | `"10m"` |
| `controllerManager.manager.resources.requests.memory` |  | `"64Mi"` |
| `controllerManager.replicas` |  | `1` |
| `controllerManager.serviceAccount.annotations` |  | `{}` |
| `kubernetesClusterDomain` |  | `"cluster.local"` |
| `metricsService.ports` |  | `[{"name": "https", "port": 8443, "protocol": "TCP", "targetPort": "https"}]` |
| `metricsService.type` |  | `"ClusterIP"` |
| `plugins` |  | `[{"chart": {"name": "validator-plugin-aws", "repository": "https://spectrocloud-labs.github.io/validator-plugin-aws", "version": "v0.0.2"}, "values": "controllerManager:\n  kubeRbacProxy:\n    args:\n    - --secure-listen-address=0.0.0.0:8443\n    - --upstream=http://127.0.0.1:8080/\n    - --logtostderr=true\n    - --v=0\n    containerSecurityContext:\n      allowPrivilegeEscalation: false\n      capabilities:\n        drop:\n        - ALL\n    image:\n      repository: gcr.io/kubebuilder/kube-rbac-proxy\n      tag: v0.14.1\n    resources:\n      limits:\n        cpu: 500m\n        memory: 128Mi\n      requests:\n        cpu: 5m\n        memory: 64Mi\n  manager:\n    args:\n    - --health-probe-bind-address=:8081\n    - --metrics-bind-address=127.0.0.1:8080\n    - --leader-elect\n    containerSecurityContext:\n      allowPrivilegeEscalation: false\n      capabilities:\n        drop:\n        - ALL\n    image:\n      repository: quay.io/spectrocloud-labs/validator-plugin-aws\n      tag: v0.0.2\n    resources:\n      limits:\n        cpu: 500m\n        memory: 128Mi\n      requests:\n        cpu: 10m\n        memory: 64Mi\n  replicas: 1\n  serviceAccount:\n    annotations: {}\nkubernetesClusterDomain: cluster.local\nmetricsService:\n  ports:\n  - name: https\n    port: 8443\n    protocol: TCP\n    targetPort: https\n  type: ClusterIP\nauth:\n  secretName: aws-creds\n  accessKeyId: \"\"\n  secretAccessKey: \"\""}]` |



---
_Documentation generated by [Frigate](https://frigate.readthedocs.io)._
