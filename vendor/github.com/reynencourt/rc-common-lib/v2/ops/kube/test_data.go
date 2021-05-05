package kube

var (
	TestNodeList = `
{
  "apiVersion": "v1",
  "items": [
    {
      "apiVersion": "v1",
      "kind": "Node",
      "metadata": {
        "annotations": {
          "kubeadm.alpha.kubernetes.io/cri-socket": "/var/run/dockershim.sock",
          "node.alpha.kubernetes.io/ttl": "0",
          "volumes.kubernetes.io/controller-managed-attach-detach": "true"
        },
        "creationTimestamp": "2020-03-13T13:36:08Z",
        "labels": {
          "beta.kubernetes.io/arch": "amd64",
          "beta.kubernetes.io/os": "linux",
          "kubernetes.io/hostname": "velanefktest-k8sm-0",
          "node-role.kubernetes.io/master": ""
        },
        "name": "velanefktest-k8sm-0",
        "resourceVersion": "2614217",
        "selfLink": "/api/v1/nodes/velanefktest-k8sm-0",
        "uid": "987f9d16-652f-11ea-bfce-02a6ca74ebae"
      },
      "spec": {
        "podCIDR": "10.233.64.0/24",
        "taints": [
          {
            "effect": "NoSchedule",
            "key": "node-role.kubernetes.io/master"
          }
        ]
      },
      "status": {
        "addresses": [
          {
            "address": "192.168.1.77",
            "type": "InternalIP"
          },
          {
            "address": "velanefktest-k8sm-0",
            "type": "Hostname"
          }
        ],
        "allocatable": {
          "cpu": "1800m",
          "ephemeral-storage": "46779129369",
          "hugepages-2Mi": "0",
          "memory": "7564556Ki",
          "pods": "110"
        },
        "capacity": {
          "cpu": "2",
          "ephemeral-storage": "50758604Ki",
          "hugepages-2Mi": "0",
          "memory": "8166956Ki",
          "pods": "110"
        },
        "conditions": [
          {
            "lastHeartbeatTime": "2020-03-25T14:59:46Z",
            "lastTransitionTime": "2020-03-13T13:36:00Z",
            "message": "kubelet has sufficient memory available",
            "reason": "KubeletHasSufficientMemory",
            "status": "False",
            "type": "MemoryPressure"
          },
          {
            "lastHeartbeatTime": "2020-03-25T14:59:46Z",
            "lastTransitionTime": "2020-03-13T13:36:00Z",
            "message": "kubelet has no disk pressure",
            "reason": "KubeletHasNoDiskPressure",
            "status": "False",
            "type": "DiskPressure"
          },
          {
            "lastHeartbeatTime": "2020-03-25T14:59:46Z",
            "lastTransitionTime": "2020-03-13T13:36:00Z",
            "message": "kubelet has sufficient PID available",
            "reason": "KubeletHasSufficientPID",
            "status": "False",
            "type": "PIDPressure"
          },
          {
            "lastHeartbeatTime": "2020-03-25T14:59:46Z",
            "lastTransitionTime": "2020-03-13T13:39:45Z",
            "message": "kubelet is posting ready status. AppArmor enabled",
            "reason": "KubeletReady",
            "status": "True",
            "type": "Ready"
          }
        ],
        "daemonEndpoints": {
          "kubeletEndpoint": {
            "Port": 10250
          }
        },
        "images": [
          {
            "names": [
              "staging.repo.rcplatform.io/public/falcosecurity/falco@sha256:d20cac066b63283d855fef956a4032cf4a560af93483098f602c44d23331a999",
              "staging.repo.rcplatform.io/public/falcosecurity/falco:0.18.0"
            ],
            "sizeBytes": 733727436
          },
          {
            "names": [
              "quay.io/calico/node@sha256:a35541153f7695b38afada46843c64a2c546548cd8c171f402621736c6cf3f0b",
              "quay.io/calico/node:v3.1.3"
            ],
            "sizeBytes": 248202699
          },
          {
            "names": [
              "gcr.io/google-containers/etcd@sha256:905d7ca17fd02bc24c0eba9a062753aba15db3e31422390bc3238eb762339b20",
              "gcr.io/google-containers/etcd:3.2.24"
            ],
            "sizeBytes": 219655340
          },
          {
            "names": [
              "gcr.io/google-containers/kube-apiserver@sha256:b5d21891ea745562d467f2220cfbbd1ee3724c219b0d3c6c4d28a527c3d07915",
              "gcr.io/google-containers/kube-apiserver:v1.13.11"
            ],
            "sizeBytes": 181121790
          },
          {
            "names": [
              "gcr.io/google-containers/kube-controller-manager@sha256:19850e8d15523efdd2ea6df34ff57b9c9ba18b6c2c6b753253a0265258c7a9c0",
              "gcr.io/google-containers/kube-controller-manager:v1.13.11"
            ],
            "sizeBytes": 146382238
          },
          {
            "names": [
              "gcr.io/google_containers/kubernetes-dashboard-amd64@sha256:0ae6b69432e78069c5ce2bcde0fe409c5c4d6f0f4d9cd50a17974fea38898747",
              "gcr.io/google_containers/kubernetes-dashboard-amd64:v1.10.1"
            ],
            "sizeBytes": 121711221
          },
          {
            "names": [
              "gcr.io/google-containers/kube-proxy@sha256:a966ec39a9a6c5864299386216e10325728586b2cfafbcacb2a56c31766153c3",
              "gcr.io/google-containers/kube-proxy:v1.13.11"
            ],
            "sizeBytes": 80270012
          },
          {
            "names": [
              "gcr.io/google-containers/kube-scheduler@sha256:38328d70aed541947713665639e964c5db33d1f725dc6e18ecba8fab19b903ec",
              "gcr.io/google-containers/kube-scheduler:v1.13.11"
            ],
            "sizeBytes": 79621886
          },
          {
            "names": [
              "quay.io/calico/cni@sha256:ed172c28bc193bb09bce6be6ed7dc6bfc85118d55e61d263cee8bbb0fd464a9d",
              "quay.io/calico/cni:v3.1.3"
            ],
            "sizeBytes": 68849270
          },
          {
            "names": [
              "quay.io/calico/kube-controllers@sha256:16582e5add4849c18b4add4ccc63284a44c3093dbaf8e6e8fe0ae49c2bec6135",
              "quay.io/calico/kube-controllers:v3.1.3"
            ],
            "sizeBytes": 54985151
          },
          {
            "names": [
              "quay.io/coreos/flannel@sha256:6ecef07be35e5e861016ee557f986f89ad8244df47198de379a1bf4e580185df",
              "quay.io/coreos/flannel:v0.10.0"
            ],
            "sizeBytes": 44598861
          },
          {
            "names": [
              "quay.io/calico/ctl@sha256:59f633d1d1e0af150f79c247ab36abcb8e85f9b1351e9d106710ea462bf38233",
              "quay.io/calico/ctl:v3.1.3"
            ],
            "sizeBytes": 42750665
          },
          {
            "names": [
              "coredns/coredns@sha256:81936728011c0df9404cb70b95c17bbc8af922ec9a70d0561a5d01fefa6ffa51",
              "gcr.io/google-containers/coredns@sha256:81936728011c0df9404cb70b95c17bbc8af922ec9a70d0561a5d01fefa6ffa51",
              "coredns/coredns:1.2.6",
              "gcr.io/google-containers/coredns:1.2.6"
            ],
            "sizeBytes": 40017418
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/prometheus/node-exporter@sha256:b630fb29d99b3483c73a2a7db5fc01a967392a3d7ad754c8eccf9f4a67e7ee31",
              "staging.repo.rcplatform.io/public/prometheus/node-exporter:v0.18.1"
            ],
            "sizeBytes": 22933477
          },
          {
            "names": [
              "gcr.io/google-containers/pause@sha256:f78411e19d84a252e53bff71a4407a5686c46983a2c2eeed83929b888179acea",
              "gcr.io/google_containers/pause-amd64@sha256:59eec8837a4d942cc19a52b8c09ea75121acc38114a2c68b98983ce9356b8610",
              "gcr.io/google-containers/pause:3.1",
              "gcr.io/google_containers/pause-amd64:3.1"
            ],
            "sizeBytes": 742472
          }
        ],
        "nodeInfo": {
          "architecture": "amd64",
          "bootID": "8ec6b9cc-d18b-49ab-a4df-cbba51e4ac50",
          "containerRuntimeVersion": "docker://18.6.1",
          "kernelVersion": "4.15.0-1021-aws",
          "kubeProxyVersion": "v1.13.11",
          "kubeletVersion": "v1.13.11",
          "machineID": "c4f0562cd4264d5bb0bf459c016c263b",
          "operatingSystem": "linux",
          "osImage": "Ubuntu 18.04.1 LTS",
          "systemUUID": "EC22FF18-48C7-5E6A-3489-5F5D53F79328"
        }
      }
    },
    {
      "apiVersion": "v1",
      "kind": "Node",
      "metadata": {
        "annotations": {
          "csi.volume.kubernetes.io/nodeid": "{\"rc-system.cephfs.csi.ceph.com\":\"velanefktest-k8sw-0\",\"rc-system.rbd.csi.ceph.com\":\"velanefktest-k8sw-0\"}",
          "kubeadm.alpha.kubernetes.io/cri-socket": "/var/run/dockershim.sock",
          "node.alpha.kubernetes.io/ttl": "0",
          "volumes.kubernetes.io/controller-managed-attach-detach": "true"
        },
        "creationTimestamp": "2020-03-13T13:39:11Z",
        "labels": {
          "beta.kubernetes.io/arch": "amd64",
          "beta.kubernetes.io/os": "linux",
          "kubernetes.io/hostname": "velanefktest-k8sw-0",
          "node-role.kubernetes.io/node": ""
        },
        "name": "velanefktest-k8sw-0",
        "resourceVersion": "2614223",
        "selfLink": "/api/v1/nodes/velanefktest-k8sw-0",
        "uid": "057497f8-6530-11ea-bfce-02a6ca74ebae"
      },
      "spec": {
        "podCIDR": "10.233.67.0/24"
      },
      "status": {
        "addresses": [
          {
            "address": "192.168.1.136",
            "type": "InternalIP"
          },
          {
            "address": "velanefktest-k8sw-0",
            "type": "Hostname"
          }
        ],
        "allocatable": {
          "cpu": "1900m",
          "ephemeral-storage": "46779129369",
          "hugepages-2Mi": "0",
          "memory": "7814556Ki",
          "pods": "110"
        },
        "capacity": {
          "cpu": "2",
          "ephemeral-storage": "50758604Ki",
          "hugepages-2Mi": "0",
          "memory": "8166956Ki",
          "pods": "110"
        },
        "conditions": [
          {
            "lastHeartbeatTime": "2020-03-25T14:59:49Z",
            "lastTransitionTime": "2020-03-13T13:39:11Z",
            "message": "kubelet has sufficient memory available",
            "reason": "KubeletHasSufficientMemory",
            "status": "False",
            "type": "MemoryPressure"
          },
          {
            "lastHeartbeatTime": "2020-03-25T14:59:49Z",
            "lastTransitionTime": "2020-03-13T13:39:11Z",
            "message": "kubelet has no disk pressure",
            "reason": "KubeletHasNoDiskPressure",
            "status": "False",
            "type": "DiskPressure"
          },
          {
            "lastHeartbeatTime": "2020-03-25T14:59:49Z",
            "lastTransitionTime": "2020-03-13T13:39:11Z",
            "message": "kubelet has sufficient PID available",
            "reason": "KubeletHasSufficientPID",
            "status": "False",
            "type": "PIDPressure"
          },
          {
            "lastHeartbeatTime": "2020-03-25T14:59:49Z",
            "lastTransitionTime": "2020-03-13T13:41:37Z",
            "message": "kubelet is posting ready status. AppArmor enabled",
            "reason": "KubeletReady",
            "status": "True",
            "type": "Ready"
          }
        ],
        "daemonEndpoints": {
          "kubeletEndpoint": {
            "Port": 10250
          }
        },
        "images": [
          {
            "names": [
              "quay.io/cephcsi/cephcsi@sha256:c1aa8bbca88cd4f42574b58eb7b3bbd11ca3a28c6101d52b1448d0c27f6e3a1d",
              "quay.io/cephcsi/cephcsi:v1.2.2"
            ],
            "sizeBytes": 983917738
          },
          {
            "names": [
              "rook/ceph@sha256:48b1596a801301715fa8d021c37cc1e1cd949877f3e73f9dc9a8d9dea24cafe7",
              "rook/ceph:v1.2.2"
            ],
            "sizeBytes": 911190029
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/ceph/ceph@sha256:24092f24ab0aaa05bd828e2a64825fa294bedf876060da145f8c30e13b9909f6",
              "staging.repo.rcplatform.io/public/ceph/ceph:v14.2.6"
            ],
            "sizeBytes": 831047730
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/falcosecurity/falco@sha256:d20cac066b63283d855fef956a4032cf4a560af93483098f602c44d23331a999",
              "staging.repo.rcplatform.io/public/falcosecurity/falco:0.18.0"
            ],
            "sizeBytes": 733727436
          },
          {
            "names": [
              "quay.io/calico/node@sha256:a35541153f7695b38afada46843c64a2c546548cd8c171f402621736c6cf3f0b",
              "quay.io/calico/node:v3.1.3"
            ],
            "sizeBytes": 248202699
          },
          {
            "names": [
              "nginx@sha256:b1d09e9718890e6ebbbd2bc319ef1611559e30ce1b6f56b2e3b479d9da51dc35",
              "nginx:1.13"
            ],
            "sizeBytes": 108958610
          },
          {
            "names": [
              "gcr.io/google-containers/kube-proxy@sha256:a966ec39a9a6c5864299386216e10325728586b2cfafbcacb2a56c31766153c3",
              "gcr.io/google-containers/kube-proxy:v1.13.11"
            ],
            "sizeBytes": 80270012
          },
          {
            "names": [
              "traefik@sha256:6348d14975b683a2783cc8dbaa76efe5367a1d67bd3e9e2ee812da7da012224b",
              "traefik:1.7.9"
            ],
            "sizeBytes": 70200455
          },
          {
            "names": [
              "quay.io/calico/cni@sha256:ed172c28bc193bb09bce6be6ed7dc6bfc85118d55e61d263cee8bbb0fd464a9d",
              "quay.io/calico/cni:v3.1.3"
            ],
            "sizeBytes": 68849270
          },
          {
            "names": [
              "fluent/fluent-bit@sha256:3cb640ccc160a6ac98763f8c095bf65bdfcc259a4a308681a33f8a23fc72635f",
              "fluent/fluent-bit:1.3.5-debug"
            ],
            "sizeBytes": 56023444
          },
          {
            "names": [
              "quay.io/calico/kube-controllers@sha256:16582e5add4849c18b4add4ccc63284a44c3093dbaf8e6e8fe0ae49c2bec6135",
              "quay.io/calico/kube-controllers:v3.1.3"
            ],
            "sizeBytes": 54985151
          },
          {
            "names": [
              "quay.io/k8scsi/csi-provisioner@sha256:3a14f801f330d5eacee11f544d2c2c7cc4a733835e25b59887053358108cea69",
              "quay.io/k8scsi/csi-provisioner:v1.4.0"
            ],
            "sizeBytes": 54456157
          },
          {
            "names": [
              "jettech/kube-webhook-certgen@sha256:58fde0ddd7a0d1bf1483fed53e363144ae8741d8a2d6c129422e8b1b9aa0543c",
              "jettech/kube-webhook-certgen:v1.0.0"
            ],
            "sizeBytes": 51532105
          },
          {
            "names": [
              "quay.io/k8scsi/csi-attacher@sha256:26fccd7a99d973845df1193b46ebdcc6ab8dc5f6e6be319750c471fce1742d13",
              "quay.io/k8scsi/csi-attacher:v1.2.0"
            ],
            "sizeBytes": 46226754
          },
          {
            "names": [
              "gcr.io/google_containers/cluster-proportional-autoscaler-amd64@sha256:4fd37c5b29a38b02c408c56254bd1a3a76f3e236610bc7a8382500bbf9ecfc76",
              "gcr.io/google_containers/cluster-proportional-autoscaler-amd64:1.3.0"
            ],
            "sizeBytes": 45844959
          },
          {
            "names": [
              "quay.io/coreos/flannel@sha256:6ecef07be35e5e861016ee557f986f89ad8244df47198de379a1bf4e580185df",
              "quay.io/coreos/flannel:v0.10.0"
            ],
            "sizeBytes": 44598861
          },
          {
            "names": [
              "quay.io/calico/ctl@sha256:59f633d1d1e0af150f79c247ab36abcb8e85f9b1351e9d106710ea462bf38233",
              "quay.io/calico/ctl:v3.1.3"
            ],
            "sizeBytes": 42750665
          },
          {
            "names": [
              "coredns/coredns@sha256:81936728011c0df9404cb70b95c17bbc8af922ec9a70d0561a5d01fefa6ffa51",
              "coredns/coredns:1.2.6"
            ],
            "sizeBytes": 40017418
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/prometheus/node-exporter@sha256:b630fb29d99b3483c73a2a7db5fc01a967392a3d7ad754c8eccf9f4a67e7ee31",
              "staging.repo.rcplatform.io/public/prometheus/node-exporter:v0.18.1"
            ],
            "sizeBytes": 22933477
          },
          {
            "names": [
              "singaravelan21/auction_biding@sha256:6374bec91b4124d261ab0b1ad6cd5c564ceb17fb1d1df52755996b68cf77ab22",
              "singaravelan21/auction_biding:v1.0.7"
            ],
            "sizeBytes": 19134269
          },
          {
            "names": [
              "singaravelan21/auction_biding@sha256:e9d46a9b3b051e851f97df0c94d51b97004103438b232d0bdac7265959b92bc3",
              "singaravelan21/auction_biding:v1.0.10"
            ],
            "sizeBytes": 19119780
          },
          {
            "names": [
              "singaravelan21/auction_biding@sha256:7ce7c3594abb549febf5b958017d4660cd902a20ed576e47a174bcb5c583bd42",
              "singaravelan21/auction_biding:v1.0.9"
            ],
            "sizeBytes": 19119780
          },
          {
            "names": [
              "quay.io/k8scsi/csi-node-driver-registrar@sha256:13daf82fb99e951a4bff8ae5fc7c17c3a8fe7130be6400990d8f6076c32d4599",
              "quay.io/k8scsi/csi-node-driver-registrar:v1.1.0"
            ],
            "sizeBytes": 15815995
          },
          {
            "names": [
              "gcr.io/google_containers/pause-amd64@sha256:59eec8837a4d942cc19a52b8c09ea75121acc38114a2c68b98983ce9356b8610",
              "gcr.io/google_containers/pause-amd64:3.1"
            ],
            "sizeBytes": 742472
          }
        ],
        "nodeInfo": {
          "architecture": "amd64",
          "bootID": "5fce754d-a69c-4a7f-a6bc-2cea25391b61",
          "containerRuntimeVersion": "docker://18.6.1",
          "kernelVersion": "4.15.0-1021-aws",
          "kubeProxyVersion": "v1.13.11",
          "kubeletVersion": "v1.13.11",
          "machineID": "99eefc2cefb4485f9d3f8883cb8f1a61",
          "operatingSystem": "linux",
          "osImage": "Ubuntu 18.04.1 LTS",
          "systemUUID": "EC2EAFF3-C8CB-AE5F-65C1-D1AE8DE377E5"
        }
      }
    },
    {
      "apiVersion": "v1",
      "kind": "Node",
      "metadata": {
        "annotations": {
          "csi.volume.kubernetes.io/nodeid": "{\"rc-system.cephfs.csi.ceph.com\":\"velanefktest-k8sw-1\",\"rc-system.rbd.csi.ceph.com\":\"velanefktest-k8sw-1\"}",
          "kubeadm.alpha.kubernetes.io/cri-socket": "/var/run/dockershim.sock",
          "node.alpha.kubernetes.io/ttl": "0",
          "volumes.kubernetes.io/controller-managed-attach-detach": "true"
        },
        "creationTimestamp": "2020-03-13T13:39:11Z",
        "labels": {
          "beta.kubernetes.io/arch": "amd64",
          "beta.kubernetes.io/os": "linux",
          "kubernetes.io/hostname": "velanefktest-k8sw-1",
          "node-role.kubernetes.io/node": ""
        },
        "name": "velanefktest-k8sw-1",
        "resourceVersion": "2614222",
        "selfLink": "/api/v1/nodes/velanefktest-k8sw-1",
        "uid": "05712b0e-6530-11ea-bfce-02a6ca74ebae"
      },
      "spec": {
        "podCIDR": "10.233.66.0/24"
      },
      "status": {
        "addresses": [
          {
            "address": "192.168.1.150",
            "type": "InternalIP"
          },
          {
            "address": "velanefktest-k8sw-1",
            "type": "Hostname"
          }
        ],
        "allocatable": {
          "cpu": "1900m",
          "ephemeral-storage": "46779129369",
          "hugepages-2Mi": "0",
          "memory": "7814556Ki",
          "pods": "110"
        },
        "capacity": {
          "cpu": "2",
          "ephemeral-storage": "50758604Ki",
          "hugepages-2Mi": "0",
          "memory": "8166956Ki",
          "pods": "110"
        },
        "conditions": [
          {
            "lastHeartbeatTime": "2020-03-25T14:59:48Z",
            "lastTransitionTime": "2020-03-13T13:39:11Z",
            "message": "kubelet has sufficient memory available",
            "reason": "KubeletHasSufficientMemory",
            "status": "False",
            "type": "MemoryPressure"
          },
          {
            "lastHeartbeatTime": "2020-03-25T14:59:48Z",
            "lastTransitionTime": "2020-03-13T13:39:11Z",
            "message": "kubelet has no disk pressure",
            "reason": "KubeletHasNoDiskPressure",
            "status": "False",
            "type": "DiskPressure"
          },
          {
            "lastHeartbeatTime": "2020-03-25T14:59:48Z",
            "lastTransitionTime": "2020-03-13T13:39:11Z",
            "message": "kubelet has sufficient PID available",
            "reason": "KubeletHasSufficientPID",
            "status": "False",
            "type": "PIDPressure"
          },
          {
            "lastHeartbeatTime": "2020-03-25T14:59:48Z",
            "lastTransitionTime": "2020-03-13T13:41:37Z",
            "message": "kubelet is posting ready status. AppArmor enabled",
            "reason": "KubeletReady",
            "status": "True",
            "type": "Ready"
          }
        ],
        "daemonEndpoints": {
          "kubeletEndpoint": {
            "Port": 10250
          }
        },
        "images": [
          {
            "names": [
              "quay.io/cephcsi/cephcsi@sha256:c1aa8bbca88cd4f42574b58eb7b3bbd11ca3a28c6101d52b1448d0c27f6e3a1d",
              "quay.io/cephcsi/cephcsi:v1.2.2"
            ],
            "sizeBytes": 983917738
          },
          {
            "names": [
              "rook/ceph@sha256:48b1596a801301715fa8d021c37cc1e1cd949877f3e73f9dc9a8d9dea24cafe7",
              "rook/ceph:v1.2.2"
            ],
            "sizeBytes": 911190029
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/ceph/ceph@sha256:24092f24ab0aaa05bd828e2a64825fa294bedf876060da145f8c30e13b9909f6",
              "staging.repo.rcplatform.io/public/ceph/ceph:v14.2.6"
            ],
            "sizeBytes": 831047730
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/falcosecurity/falco@sha256:d20cac066b63283d855fef956a4032cf4a560af93483098f602c44d23331a999",
              "staging.repo.rcplatform.io/public/falcosecurity/falco:0.18.0"
            ],
            "sizeBytes": 733727436
          },
          {
            "names": [
              "quay.io/calico/node@sha256:a35541153f7695b38afada46843c64a2c546548cd8c171f402621736c6cf3f0b",
              "quay.io/calico/node:v3.1.3"
            ],
            "sizeBytes": 248202699
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/appscode/stash@sha256:8a9225d72450c6cd68c330f2fb96ce5d390f46b7cb178f90e16ad7f05fdfd32f",
              "staging.repo.rcplatform.io/public/appscode/stash:v0.9.0-rc.2"
            ],
            "sizeBytes": 160782271
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/prometheus/prometheus@sha256:c906d6ecc3586d3fd00e2aab7a13c961d13f938dab5e7338c9973e15d5e7b62b",
              "staging.repo.rcplatform.io/public/prometheus/prometheus:v2.13.1"
            ],
            "sizeBytes": 128568244
          },
          {
            "names": [
              "gcr.io/google_containers/kubernetes-dashboard-amd64@sha256:0ae6b69432e78069c5ce2bcde0fe409c5c4d6f0f4d9cd50a17974fea38898747",
              "gcr.io/google_containers/kubernetes-dashboard-amd64:v1.10.1"
            ],
            "sizeBytes": 121711221
          },
          {
            "names": [
              "nginx@sha256:b1d09e9718890e6ebbbd2bc319ef1611559e30ce1b6f56b2e3b479d9da51dc35",
              "nginx:1.13"
            ],
            "sizeBytes": 108958610
          },
          {
            "names": [
              "gcr.io/google-containers/kube-proxy@sha256:a966ec39a9a6c5864299386216e10325728586b2cfafbcacb2a56c31766153c3",
              "gcr.io/google-containers/kube-proxy:v1.13.11"
            ],
            "sizeBytes": 80270012
          },
          {
            "names": [
              "traefik@sha256:6348d14975b683a2783cc8dbaa76efe5367a1d67bd3e9e2ee812da7da012224b",
              "traefik:1.7.9"
            ],
            "sizeBytes": 70200455
          },
          {
            "names": [
              "quay.io/calico/cni@sha256:ed172c28bc193bb09bce6be6ed7dc6bfc85118d55e61d263cee8bbb0fd464a9d",
              "quay.io/calico/cni:v3.1.3"
            ],
            "sizeBytes": 68849270
          },
          {
            "names": [
              "quay.io/thanos/thanos@sha256:fccf337c57b0a98ca5960760c49fcc5d33d16816bc83f4b53d01906420e89199",
              "quay.io/thanos/thanos:v0.9.0-rc.0"
            ],
            "sizeBytes": 65399149
          },
          {
            "names": [
              "fluent/fluent-bit@sha256:3cb640ccc160a6ac98763f8c095bf65bdfcc259a4a308681a33f8a23fc72635f",
              "fluent/fluent-bit:1.3.5-debug"
            ],
            "sizeBytes": 56023444
          },
          {
            "names": [
              "quay.io/calico/kube-controllers@sha256:16582e5add4849c18b4add4ccc63284a44c3093dbaf8e6e8fe0ae49c2bec6135",
              "quay.io/calico/kube-controllers:v3.1.3"
            ],
            "sizeBytes": 54985151
          },
          {
            "names": [
              "jettech/kube-webhook-certgen@sha256:58fde0ddd7a0d1bf1483fed53e363144ae8741d8a2d6c129422e8b1b9aa0543c",
              "jettech/kube-webhook-certgen:v1.0.0"
            ],
            "sizeBytes": 51532105
          },
          {
            "names": [
              "gcr.io/google_containers/cluster-proportional-autoscaler-amd64@sha256:4fd37c5b29a38b02c408c56254bd1a3a76f3e236610bc7a8382500bbf9ecfc76",
              "gcr.io/google_containers/cluster-proportional-autoscaler-amd64:1.3.0"
            ],
            "sizeBytes": 45844959
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/coreos/prometheus-operator@sha256:9bdf69fdf2b9083f3467f6ebdbd781dce9c3bf843f532f2ab397c4cb881b2388",
              "staging.repo.rcplatform.io/public/coreos/prometheus-operator:v0.34.0"
            ],
            "sizeBytes": 45519458
          },
          {
            "names": [
              "quay.io/coreos/flannel@sha256:6ecef07be35e5e861016ee557f986f89ad8244df47198de379a1bf4e580185df",
              "quay.io/coreos/flannel:v0.10.0"
            ],
            "sizeBytes": 44598861
          },
          {
            "names": [
              "quay.io/calico/ctl@sha256:59f633d1d1e0af150f79c247ab36abcb8e85f9b1351e9d106710ea462bf38233",
              "quay.io/calico/ctl:v3.1.3"
            ],
            "sizeBytes": 42750665
          },
          {
            "names": [
              "coredns/coredns@sha256:81936728011c0df9404cb70b95c17bbc8af922ec9a70d0561a5d01fefa6ffa51",
              "coredns/coredns:1.2.6"
            ],
            "sizeBytes": 40017418
          },
          {
            "names": [
              "quay.io/coreos/kube-state-metrics@sha256:f75c3e5c5c7f65846ddd6883d6187b38f77721a3938f241c9e5d0ebe7beb8e19",
              "quay.io/coreos/kube-state-metrics:v1.8.0"
            ],
            "sizeBytes": 33616576
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/squareup/ghostunnel@sha256:80674e1ec91cae428ca0d92d2fe7f23e8c76b1d60c32d6ab4c8c5605609a7b1c",
              "staging.repo.rcplatform.io/public/squareup/ghostunnel:v1.4.1"
            ],
            "sizeBytes": 27416604
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/prometheus/node-exporter@sha256:b630fb29d99b3483c73a2a7db5fc01a967392a3d7ad754c8eccf9f4a67e7ee31",
              "staging.repo.rcplatform.io/public/prometheus/node-exporter:v0.18.1"
            ],
            "sizeBytes": 22933477
          },
          {
            "names": [
              "singaravelan21/auction_biding@sha256:6374bec91b4124d261ab0b1ad6cd5c564ceb17fb1d1df52755996b68cf77ab22",
              "singaravelan21/auction_biding:v1.0.7"
            ],
            "sizeBytes": 19134269
          },
          {
            "names": [
              "singaravelan21/auction_biding@sha256:7ce7c3594abb549febf5b958017d4660cd902a20ed576e47a174bcb5c583bd42",
              "singaravelan21/auction_biding:v1.0.9"
            ],
            "sizeBytes": 19119780
          },
          {
            "names": [
              "singaravelan21/auction_biding@sha256:e9d46a9b3b051e851f97df0c94d51b97004103438b232d0bdac7265959b92bc3",
              "singaravelan21/auction_biding:v1.0.10"
            ],
            "sizeBytes": 19119780
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/prom/pushgateway@sha256:4c9fa582c167f6fec89c6ebbd454ec1b5dba3ee0e1e21118fac6737c06861f92",
              "staging.repo.rcplatform.io/public/prom/pushgateway:v0.5.2"
            ],
            "sizeBytes": 16442555
          },
          {
            "names": [
              "quay.io/k8scsi/csi-node-driver-registrar@sha256:13daf82fb99e951a4bff8ae5fc7c17c3a8fe7130be6400990d8f6076c32d4599",
              "quay.io/k8scsi/csi-node-driver-registrar:v1.1.0"
            ],
            "sizeBytes": 15815995
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/coreos/prometheus-config-reloader@sha256:d408f9a644fd65f342e175dd388d69f731cf1791887305447cac809f7d810d68",
              "staging.repo.rcplatform.io/public/coreos/prometheus-config-reloader:v0.34.0"
            ],
            "sizeBytes": 10152002
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/coreos/configmap-reload@sha256:50c53db55ece9a6e1a7274e497f308bcc24164bdb4c0885524037c1b8e4e758d",
              "staging.repo.rcplatform.io/public/coreos/configmap-reload:v0.0.1"
            ],
            "sizeBytes": 4785056
          },
          {
            "names": [
              "gcr.io/google_containers/pause-amd64@sha256:59eec8837a4d942cc19a52b8c09ea75121acc38114a2c68b98983ce9356b8610",
              "gcr.io/google_containers/pause-amd64:3.1"
            ],
            "sizeBytes": 742472
          }
        ],
        "nodeInfo": {
          "architecture": "amd64",
          "bootID": "cc7440f0-ef1e-4098-8d4e-c6ba29ec06b4",
          "containerRuntimeVersion": "docker://18.6.1",
          "kernelVersion": "4.15.0-1021-aws",
          "kubeProxyVersion": "v1.13.11",
          "kubeletVersion": "v1.13.11",
          "machineID": "1b8d43df0bd24c9c946436462142d0e1",
          "operatingSystem": "linux",
          "osImage": "Ubuntu 18.04.1 LTS",
          "systemUUID": "EC25828F-DBF1-36CD-3DF3-4402787146AF"
        },
        "volumesAttached": [
          {
            "devicePath": "csi-05cbbd15172c1830c5c3c7283bf0341a77855d54c6db4fc8e7a30a3c7c2e2ee7",
            "name": "kubernetes.io/csi/rc-system.rbd.csi.ceph.com^0001-0009-rc-system-0000000000000001-c0d57a9c-6531-11ea-a68d-4a025565c306"
          }
        ],
        "volumesInUse": [
          "kubernetes.io/csi/rc-system.rbd.csi.ceph.com^0001-0009-rc-system-0000000000000001-c0d57a9c-6531-11ea-a68d-4a025565c306"
        ]
      }
    },
    {
      "apiVersion": "v1",
      "kind": "Node",
      "metadata": {
        "annotations": {
          "csi.volume.kubernetes.io/nodeid": "{\"rc-system.cephfs.csi.ceph.com\":\"velanefktest-k8sw-2\",\"rc-system.rbd.csi.ceph.com\":\"velanefktest-k8sw-2\"}",
          "kubeadm.alpha.kubernetes.io/cri-socket": "/var/run/dockershim.sock",
          "node.alpha.kubernetes.io/ttl": "0",
          "volumes.kubernetes.io/controller-managed-attach-detach": "true"
        },
        "creationTimestamp": "2020-03-13T13:39:11Z",
        "labels": {
          "beta.kubernetes.io/arch": "amd64",
          "beta.kubernetes.io/os": "linux",
          "kubernetes.io/hostname": "velanefktest-k8sw-2",
          "node-role.kubernetes.io/node": ""
        },
        "name": "velanefktest-k8sw-2",
        "resourceVersion": "2614228",
        "selfLink": "/api/v1/nodes/velanefktest-k8sw-2",
        "uid": "056d30f7-6530-11ea-bfce-02a6ca74ebae"
      },
      "spec": {
        "podCIDR": "10.233.65.0/24"
      },
      "status": {
        "addresses": [
          {
            "address": "192.168.1.43",
            "type": "InternalIP"
          },
          {
            "address": "velanefktest-k8sw-2",
            "type": "Hostname"
          }
        ],
        "allocatable": {
          "cpu": "1900m",
          "ephemeral-storage": "46779129369",
          "hugepages-2Mi": "0",
          "memory": "7814556Ki",
          "pods": "110"
        },
        "capacity": {
          "cpu": "2",
          "ephemeral-storage": "50758604Ki",
          "hugepages-2Mi": "0",
          "memory": "8166956Ki",
          "pods": "110"
        },
        "conditions": [
          {
            "lastHeartbeatTime": "2020-03-25T14:59:51Z",
            "lastTransitionTime": "2020-03-13T13:39:11Z",
            "message": "kubelet has sufficient memory available",
            "reason": "KubeletHasSufficientMemory",
            "status": "False",
            "type": "MemoryPressure"
          },
          {
            "lastHeartbeatTime": "2020-03-25T14:59:51Z",
            "lastTransitionTime": "2020-03-13T13:39:11Z",
            "message": "kubelet has no disk pressure",
            "reason": "KubeletHasNoDiskPressure",
            "status": "False",
            "type": "DiskPressure"
          },
          {
            "lastHeartbeatTime": "2020-03-25T14:59:51Z",
            "lastTransitionTime": "2020-03-13T13:39:11Z",
            "message": "kubelet has sufficient PID available",
            "reason": "KubeletHasSufficientPID",
            "status": "False",
            "type": "PIDPressure"
          },
          {
            "lastHeartbeatTime": "2020-03-25T14:59:51Z",
            "lastTransitionTime": "2020-03-13T13:41:38Z",
            "message": "kubelet is posting ready status. AppArmor enabled",
            "reason": "KubeletReady",
            "status": "True",
            "type": "Ready"
          }
        ],
        "daemonEndpoints": {
          "kubeletEndpoint": {
            "Port": 10250
          }
        },
        "images": [
          {
            "names": [
              "quay.io/cephcsi/cephcsi@sha256:c1aa8bbca88cd4f42574b58eb7b3bbd11ca3a28c6101d52b1448d0c27f6e3a1d",
              "quay.io/cephcsi/cephcsi:v1.2.2"
            ],
            "sizeBytes": 983917738
          },
          {
            "names": [
              "rook/ceph@sha256:48b1596a801301715fa8d021c37cc1e1cd949877f3e73f9dc9a8d9dea24cafe7",
              "rook/ceph:v1.2.2"
            ],
            "sizeBytes": 911190029
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/ceph/ceph@sha256:24092f24ab0aaa05bd828e2a64825fa294bedf876060da145f8c30e13b9909f6",
              "staging.repo.rcplatform.io/public/ceph/ceph:v14.2.6"
            ],
            "sizeBytes": 831047730
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/falcosecurity/falco@sha256:d20cac066b63283d855fef956a4032cf4a560af93483098f602c44d23331a999",
              "staging.repo.rcplatform.io/public/falcosecurity/falco:0.18.0"
            ],
            "sizeBytes": 733727436
          },
          {
            "names": [
              "quay.io/calico/node@sha256:a35541153f7695b38afada46843c64a2c546548cd8c171f402621736c6cf3f0b",
              "quay.io/calico/node:v3.1.3"
            ],
            "sizeBytes": 248202699
          },
          {
            "names": [
              "nginx@sha256:b1d09e9718890e6ebbbd2bc319ef1611559e30ce1b6f56b2e3b479d9da51dc35",
              "nginx:1.13"
            ],
            "sizeBytes": 108958610
          },
          {
            "names": [
              "gcr.io/google-containers/kube-proxy@sha256:a966ec39a9a6c5864299386216e10325728586b2cfafbcacb2a56c31766153c3",
              "gcr.io/google-containers/kube-proxy:v1.13.11"
            ],
            "sizeBytes": 80270012
          },
          {
            "names": [
              "traefik@sha256:6348d14975b683a2783cc8dbaa76efe5367a1d67bd3e9e2ee812da7da012224b",
              "traefik:1.7.9"
            ],
            "sizeBytes": 70200455
          },
          {
            "names": [
              "quay.io/calico/cni@sha256:ed172c28bc193bb09bce6be6ed7dc6bfc85118d55e61d263cee8bbb0fd464a9d",
              "quay.io/calico/cni:v3.1.3"
            ],
            "sizeBytes": 68849270
          },
          {
            "names": [
              "fluent/fluent-bit@sha256:3cb640ccc160a6ac98763f8c095bf65bdfcc259a4a308681a33f8a23fc72635f",
              "fluent/fluent-bit:1.3.5-debug"
            ],
            "sizeBytes": 56023444
          },
          {
            "names": [
              "quay.io/calico/kube-controllers@sha256:16582e5add4849c18b4add4ccc63284a44c3093dbaf8e6e8fe0ae49c2bec6135",
              "quay.io/calico/kube-controllers:v3.1.3"
            ],
            "sizeBytes": 54985151
          },
          {
            "names": [
              "quay.io/k8scsi/csi-provisioner@sha256:3a14f801f330d5eacee11f544d2c2c7cc4a733835e25b59887053358108cea69",
              "quay.io/k8scsi/csi-provisioner:v1.4.0"
            ],
            "sizeBytes": 54456157
          },
          {
            "names": [
              "quay.io/k8scsi/csi-snapshotter@sha256:fbef4d21ec79c9c5aced2e0e0396b9272e69ffe4cc718c3602ae6a466f682d41",
              "quay.io/k8scsi/csi-snapshotter:v1.2.2"
            ],
            "sizeBytes": 47647695
          },
          {
            "names": [
              "quay.io/k8scsi/csi-attacher@sha256:26fccd7a99d973845df1193b46ebdcc6ab8dc5f6e6be319750c471fce1742d13",
              "quay.io/k8scsi/csi-attacher:v1.2.0"
            ],
            "sizeBytes": 46226754
          },
          {
            "names": [
              "gcr.io/google_containers/cluster-proportional-autoscaler-amd64@sha256:4fd37c5b29a38b02c408c56254bd1a3a76f3e236610bc7a8382500bbf9ecfc76",
              "gcr.io/google_containers/cluster-proportional-autoscaler-amd64:1.3.0"
            ],
            "sizeBytes": 45844959
          },
          {
            "names": [
              "quay.io/coreos/flannel@sha256:6ecef07be35e5e861016ee557f986f89ad8244df47198de379a1bf4e580185df",
              "quay.io/coreos/flannel:v0.10.0"
            ],
            "sizeBytes": 44598861
          },
          {
            "names": [
              "quay.io/calico/ctl@sha256:59f633d1d1e0af150f79c247ab36abcb8e85f9b1351e9d106710ea462bf38233",
              "quay.io/calico/ctl:v3.1.3"
            ],
            "sizeBytes": 42750665
          },
          {
            "names": [
              "coredns/coredns@sha256:81936728011c0df9404cb70b95c17bbc8af922ec9a70d0561a5d01fefa6ffa51",
              "gcr.io/google-containers/coredns@sha256:81936728011c0df9404cb70b95c17bbc8af922ec9a70d0561a5d01fefa6ffa51",
              "coredns/coredns:1.2.6",
              "gcr.io/google-containers/coredns:1.2.6"
            ],
            "sizeBytes": 40017418
          },
          {
            "names": [
              "staging.repo.rcplatform.io/public/prometheus/node-exporter@sha256:b630fb29d99b3483c73a2a7db5fc01a967392a3d7ad754c8eccf9f4a67e7ee31",
              "staging.repo.rcplatform.io/public/prometheus/node-exporter:v0.18.1"
            ],
            "sizeBytes": 22933477
          },
          {
            "names": [
              "singaravelan21/auction_biding@sha256:6374bec91b4124d261ab0b1ad6cd5c564ceb17fb1d1df52755996b68cf77ab22",
              "singaravelan21/auction_biding:v1.0.7"
            ],
            "sizeBytes": 19134269
          },
          {
            "names": [
              "singaravelan21/auction_biding@sha256:7ce7c3594abb549febf5b958017d4660cd902a20ed576e47a174bcb5c583bd42",
              "singaravelan21/auction_biding:v1.0.9"
            ],
            "sizeBytes": 19119780
          },
          {
            "names": [
              "quay.io/k8scsi/csi-node-driver-registrar@sha256:13daf82fb99e951a4bff8ae5fc7c17c3a8fe7130be6400990d8f6076c32d4599",
              "quay.io/k8scsi/csi-node-driver-registrar:v1.1.0"
            ],
            "sizeBytes": 15815995
          },
          {
            "names": [
              "gcr.io/google_containers/pause-amd64@sha256:59eec8837a4d942cc19a52b8c09ea75121acc38114a2c68b98983ce9356b8610",
              "gcr.io/google_containers/pause-amd64:3.1"
            ],
            "sizeBytes": 742472
          }
        ],
        "nodeInfo": {
          "architecture": "amd64",
          "bootID": "9cc8f8c9-a9c3-497b-96cb-f04228c4b6ee",
          "containerRuntimeVersion": "docker://18.6.1",
          "kernelVersion": "4.15.0-1021-aws",
          "kubeProxyVersion": "v1.13.11",
          "kubeletVersion": "v1.13.11",
          "machineID": "d127cf3d6ada4945adc9690322de3562",
          "operatingSystem": "linux",
          "osImage": "Ubuntu 18.04.1 LTS",
          "systemUUID": "EC2FD197-34F2-9829-3A77-10B66C68CD46"
        }
      }
    }
  ],
  "kind": "List",
  "metadata": {
    "resourceVersion": "",
    "selfLink": ""
  }
}`
)
