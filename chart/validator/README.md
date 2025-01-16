
Validator
===========

Monitor results created by validator plugins and upload them to a configurable sink


## Configuration

The following table lists the configurable parameters of the Validator chart and their default values.

| Parameter                | Description             | Default        |
| ------------------------ | ----------------------- | -------------- |
| `controllerManager.manager.args` |  | `["--health-probe-bind-address=:8081", "--metrics-bind-address=:8443", "--leader-elect"]` |
| `controllerManager.manager.containerSecurityContext.allowPrivilegeEscalation` |  | `false` |
| `controllerManager.manager.containerSecurityContext.capabilities.drop` |  | `["ALL"]` |
| `controllerManager.manager.image.repository` |  | `"quay.io/validator-labs/validator"` |
| `controllerManager.manager.image.tag` | x-release-please-version | `"v0.1.17"` |
| `controllerManager.manager.resources.limits.cpu` |  | `"500m"` |
| `controllerManager.manager.resources.limits.memory` |  | `"512Mi"` |
| `controllerManager.manager.resources.requests.cpu` |  | `"10m"` |
| `controllerManager.manager.resources.requests.memory` |  | `"64Mi"` |
| `controllerManager.manager.sinkWebhookTimeout` |  | `"30s"` |
| `controllerManager.replicas` |  | `1` |
| `controllerManager.serviceAccount.annotations` |  | `{}` |
| `kubernetesClusterDomain` |  | `"cluster.local"` |
| `metricsService.ports` |  | `[{"name": "https", "port": 8443, "protocol": "TCP", "targetPort": 8443}]` |
| `metricsService.type` |  | `"ClusterIP"` |
| `env` |  | `[]` |
| `proxy.enabled` |  | `false` |
| `proxy.image` |  | `"quay.io/validator-labs/validator-certs-init:1.0.0"` |
| `proxy.secretName` |  | `"proxy-cert"` |
| `proxy.createSecret` |  | `false` |
| `proxy.caCert` |  | `"-----BEGIN CERTIFICATE-----\n<your certificate content here>\n-----END CERTIFICATE-----\n"` |
| `sink` |  | `{}` |
| `cleanup.image` |  | `"gcr.io/spectro-images-public/release/spectro-cleanup:1.2.0"` |
| `cleanup.grpcServerEnabled` |  | `true` |
| `cleanup.hostname` |  | `"validator-cleanup-service"` |
| `cleanup.port` |  | `3006` |
| `pluginSecrets.aws` | Don't forget to delete these curly braces if you're specifying credentials here! | `{}` |
| `pluginSecrets.azure` | Don't forget to delete these curly braces if you're specifying credentials here! | `{}` |
| `pluginSecrets.maas` | Don't forget to delete these curly braces if you're specifying credentials here! | `{}` |
| `pluginSecrets.network.auth` | Don't forget to delete these square brackets if you're specifying credentials here! | `[]` |
| `pluginSecrets.oci.auth` | Don't forget to delete these square brackets if you're specifying credentials here! | `[]` |
| `pluginSecrets.oci.pubKeys` | Don't forget to delete these square brackets if you're specifying public keys here! | `[]` |
| `pluginSecrets.vSphere` | Don't forget to delete these curly braces if you're specifying credentials here! | `{}` |
| `helmConfig.registry` |  | `"https://validator-labs.github.io"` |
| `plugins` |  | `[{"chart": {"name": "validator-plugin-azure", "repository": "validator-plugin-azure", "version": "v0.0.26"}, "values": "controllerManager:\n  manager:\n    args:\n      - --health-probe-bind-address=:8081\n      - --metrics-bind-address=:8443\n      - --leader-elect\n    containerSecurityContext:\n      allowPrivilegeEscalation: false\n      capabilities:\n        drop:\n          - ALL\n    image:\n      repository: quay.io/validator-labs/validator-plugin-azure\n      tag: v0.0.26\n    resources:\n      limits:\n        cpu: 500m\n        memory: 128Mi\n      requests:\n        cpu: 10m\n        memory: 64Mi\n    # Optionally specify a volumeMount to mount a volume containing a private key\n    # to leverage Azure Service principal with certificate authentication.\n    volumeMounts: []\n  replicas: 1\n  serviceAccount:\n    annotations: {}\n  # Optionally specify a volume containing a private key to leverage Azure Service\n  # principal with certificate authentication.\n  volumes: []\n  # Optionally specify additional labels to use for the controller-manager Pods.\n  podLabels: {}\nkubernetesClusterDomain: cluster.local\nmetricsService:\n  ports:\n    - name: https\n      port: 8443\n      protocol: TCP\n      targetPort: 8443\n  type: ClusterIP\nauth:\n  # Override the service account used by Azure validator (optional, could be used for WorkloadIdentityCredentials on AKS)\n  # WARNING: the chosen service account must include all RBAC privileges found in templates/manager-rbac.yaml\n  serviceAccountName: \"\"\n# Optionally specify the Azure environment to use. Defaults to \"AzureCloud\" for public Azure cloud.\n# Other acceptable values are \"AzureUSGovernment\" and \"AzureChinaCloud\".\nazureEnvironment: \"AzureCloud\""}, {"chart": {"name": "validator-plugin-oci", "repository": "validator-plugin-oci", "version": "v0.3.5"}, "values": "controllerManager:\n  manager:\n    args:\n      - --health-probe-bind-address=:8081\n      - --metrics-bind-address=:8443\n      - --leader-elect\n    containerSecurityContext:\n      allowPrivilegeEscalation: false\n      capabilities:\n        drop:\n          - ALL\n    image:\n      repository: quay.io/validator-labs/validator-plugin-oci\n      tag: v0.3.5\n    resources:\n      limits:\n        cpu: 500m\n        memory: 128Mi\n      requests:\n        cpu: 10m\n        memory: 64Mi\n  replicas: 1\n  serviceAccount:\n    annotations: {}\nkubernetesClusterDomain: cluster.local\nmetricsService:\n  ports:\n    - name: https\n      port: 8443\n      protocol: TCP\n      targetPort: 8443\n  type: ClusterIP"}, {"chart": {"name": "validator-plugin-kubescape", "repository": "validator-plugin-kubescape", "version": "v0.0.6"}, "values": "controllerManager:\n  manager:\n    args:\n      - --health-probe-bind-address=:8081\n      - --metrics-bind-address=:8443\n      - --leader-elect\n    containerSecurityContext:\n      allowPrivilegeEscalation: false\n      capabilities:\n        drop:\n          - ALL\n    image:\n      repository: quay.io/validator-labs/validator-plugin-kubescape\n      tag: v0.0.6\n    resources:\n      limits:\n        cpu: 500m\n        memory: 128Mi\n      requests:\n        cpu: 10m\n        memory: 64Mi\n    # Optionally specify a volumeMount to mount a volume containing a private key\n    # to leverage Azure Service principal with certificate authentication.\n    volumeMounts: []\n  replicas: 1\n  serviceAccount:\n    annotations: {}\n  # Optionally specify a volume containing a private key to leverage Azure Service\n  # principal with certificate authentication.\n  volumes: []\n  # Optionally specify additional labels to use for the controller-manager Pods.\n  podLabels: {}\nkubernetesClusterDomain: cluster.local\nmetricsService:\n  ports:\n    - name: https\n      port: 8443\n      protocol: TCP\n      targetPort: 8443\n  type: ClusterIP"}, {"chart": {"name": "validator-plugin-aws", "repository": "validator-plugin-aws", "version": "v0.1.12"}, "values": "controllerManager:\n  manager:\n    args:\n      - --health-probe-bind-address=:8081\n      - --metrics-bind-address=:8443\n      - --leader-elect\n    containerSecurityContext:\n      allowPrivilegeEscalation: false\n      capabilities:\n        drop:\n          - ALL\n    image:\n      repository: quay.io/validator-labs/validator-plugin-aws\n      tag: v0.1.12\n    resources:\n      limits:\n        cpu: 500m\n        memory: 128Mi\n      requests:\n        cpu: 10m\n        memory: 64Mi\n  replicas: 1\n  serviceAccount:\n    annotations: {}\nkubernetesClusterDomain: cluster.local\nmetricsService:\n  ports:\n    - name: https\n      port: 8443\n      protocol: TCP\n      targetPort: 8443\n  type: ClusterIP\nauth:\n  # Override the service account used by AWS validator (optional, could be used for IAM roles for Service Accounts)\n  # WARNING: the chosen service account must have the same RBAC privileges as seen in templates/manager-rbac.yaml\n  serviceAccountName: \"\""}, {"chart": {"name": "validator-plugin-network", "repository": "validator-plugin-network", "version": "v0.1.2"}, "values": "controllerManager:\n  manager:\n    args:\n      - --health-probe-bind-address=:8081\n      - --metrics-bind-address=:8443\n      - --leader-elect\n    containerSecurityContext:\n      allowPrivilegeEscalation: true\n      capabilities:\n        add:\n          - NET_RAW\n        drop:\n          - ALL\n    image:\n      repository: quay.io/validator-labs/validator-plugin-network\n      tag: v0.1.2\n    resources:\n      limits:\n        cpu: 500m\n        memory: 128Mi\n      requests:\n        cpu: 10m\n        memory: 64Mi\n  replicas: 1\n  serviceAccount:\n    annotations: {}\nkubernetesClusterDomain: cluster.local\nmetricsService:\n  ports:\n    - name: https\n      port: 8443\n      protocol: TCP\n      targetPort: 8443\n  type: ClusterIP"}, {"chart": {"name": "validator-plugin-maas", "repository": "validator-plugin-maas", "version": "v0.0.13"}, "values": "controllerManager:\n  manager:\n    args:\n      - --metrics-bind-address=:8443\n      - --health-probe-bind-address=:8081\n      - --leader-elect\n    containerSecurityContext:\n      allowPrivilegeEscalation: false\n      capabilities:\n        drop:\n          - ALL\n    image:\n      repository: quay.io/validator-labs/validator-plugin-maas\n      tag: v0.0.13\n    resources:\n      limits:\n        cpu: 500m\n        memory: 128Mi\n      requests:\n        cpu: 10m\n        memory: 64Mi\n  replicas: 1\n  serviceAccount:\n    annotations: {}\nkubernetesClusterDomain: cluster.local\nmetricsService:\n  ports:\n    - name: https\n      port: 8443\n      protocol: TCP\n      targetPort: 8443\n  type: ClusterIP"}, {"chart": {"name": "validator-plugin-vsphere", "repository": "validator-plugin-vsphere", "version": "v0.1.6"}, "values": "controllerManager:\n  manager:\n    args:\n      - --health-probe-bind-address=:8081\n      - --metrics-bind-address=:8443\n      - --leader-elect\n    containerSecurityContext:\n      allowPrivilegeEscalation: false\n      capabilities:\n        drop:\n          - ALL\n    image:\n      repository: quay.io/validator-labs/validator-plugin-vsphere\n      tag: v0.1.6\n    resources:\n      limits:\n        cpu: 500m\n        memory: 128Mi\n      requests:\n        cpu: 10m\n        memory: 64Mi\n  replicas: 1\n  serviceAccount:\n    annotations: {}\nkubernetesClusterDomain: cluster.local\nmetricsService:\n  ports:\n    - name: https\n      port: 8443\n      protocol: TCP\n      targetPort: 8443\n  type: ClusterIP"}]` |



---
_Documentation generated by [Frigate](https://frigate.readthedocs.io)._

