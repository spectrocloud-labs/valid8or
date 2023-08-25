
Valid8or
===========

valid8or monitors ValidationResults created by one or more valid8or plugins and uploads them to a configurable sink


## Configuration

The following table lists the configurable parameters of the Valid8or chart and their default values.

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
| `controllerManager.manager.image.repository` |  | `"quay.io/spectrocloud-labs/valid8or"` |
| `controllerManager.manager.image.tag` |  | `"latest"` |
| `controllerManager.manager.resources.limits.cpu` |  | `"500m"` |
| `controllerManager.manager.resources.limits.memory` |  | `"128Mi"` |
| `controllerManager.manager.resources.requests.cpu` |  | `"10m"` |
| `controllerManager.manager.resources.requests.memory` |  | `"64Mi"` |
| `controllerManager.replicas` |  | `1` |
| `controllerManager.serviceAccount.annotations` |  | `{}` |
| `kubernetesClusterDomain` |  | `"cluster.local"` |
| `metricsService.ports` |  | `[{"name": "https", "port": 8443, "protocol": "TCP", "targetPort": "https"}]` |
| `metricsService.type` |  | `"ClusterIP"` |
| `valid8orPluginAws.enabled` |  | `false` |
| `valid8orPluginAws.auth.accessKey` |  | `""` |
| `valid8orPluginAws.auth.secretKey` |  | `""` |
| `valid8orPluginAws.validator.name` |  | `"awsvalidator"` |
| `valid8orPluginAws.validator.auth.secretName` |  | `"aws-creds"` |
| `valid8orPluginAws.validator.iamRules` |  | `[{"iamRole": "<iam_role_name>", "statements": [{"actions": ["cognito-sync:ListDatasets"], "resource": "*"}]}]` |
| `valid8orPluginAws.validator.tagRules` |  | `[{"key": "kubernetes.io/role/elb", "expectedValue": "1", "region": "us-east-2", "resourceType": "subnet", "arns": ["<arn_1>"]}]` |



---
_Documentation generated by [Frigate](https://frigate.readthedocs.io)._
