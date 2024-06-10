# Test K8s primitives

## Waiting before shutdown
```bash
$ko apply -f ./cmd/leases/

$kubectl get po
NAME     READY   STATUS    RESTARTS   AGE
leases   1/1     Running   0          12s
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

$ kubectl delete po --all 

$kubectl get po
NAME     READY   STATUS        RESTARTS   AGE
leases   1/1     Terminating   0          21s

$kubectl get leases -ojson | grep holder
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "",
                "holderIdentity": "",
```

## No wait time

```bash
Following the above and the output is not removing the holder identity.

$ kubectl get leases -ojson | grep holder
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
                
 

```
