# Test K8s primitives

```bash
$ko apply -f ./cmd/leases/

$kubectl  get po
NAME     READY   STATUS    RESTARTS   AGE
leases   1/1     Running   0          56s

$kubectl get leases -ojson | grep holder
                "holderIdentity": "leases",
                "holderIdentity": "leases",
                "holderIdentity": "leases",
                "holderIdentity": "leases",
                "holderIdentity": "leases",
                "holderIdentity": "leases",
                "holderIdentity": "leases",
                "holderIdentity": "leases",
                "holderIdentity": "leases",
                "holderIdentity": "leases",
...

$ get leases -ojson | grep holder | grep leases  | wc -l
110


$kubectl delete po --all 

$get po
NAME     READY   STATUS        RESTARTS   AGE
leases   1/1     Terminating   0          94s


$kubectl get leases -ojson | grep holder 
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "leases",
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "leases",
...
$ kubectl get leases -ojson | grep holder | grep leases  | wc -l
19

```


