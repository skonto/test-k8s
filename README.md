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

$kubectl get leases -ojson | grep holder | grep leases  | wc -l
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

$ kubectl get po
No resources found in default namespace.

$ kubectl get leases -ojson | grep holder | grep leases  | wc -l
19

```
Logs show that object may be modified before the controller has a chance to update it again:


```
I0610 15:07:13.947116       1 leases.go:57] leader lost: leases
I0610 15:07:14.085148       1 leases.go:57] leader lost: leases
I0610 15:07:14.287422       1 leases.go:57] leader lost: leases
I0610 15:07:14.486824       1 leases.go:57] leader lost: leases
I0610 15:07:14.687426       1 leases.go:57] leader lost: leases
I0610 15:07:14.952600       1 leases.go:57] leader lost: leases
E0610 15:07:15.076863       1 leaderelection.go:308] Failed to release lock: Operation cannot be fulfilled on leases.coordination.k8s.io "lease-70": the object has been modified; please apply your changes to the latest version and try again
I0610 15:07:15.077238       1 leases.go:57] leader lost: leases
I0610 15:07:15.354260       1 leases.go:57] leader lost: leases
I0610 15:07:15.498681       1 leases.go:57] leader lost: leases
I0610 15:07:15.687356       1 leases.go:57] leader lost: leases
I0610 15:07:15.949535       1 leases.go:57] leader lost: leases
I0610 15:07:16.086689       1 leases.go:57] leader lost: leases
I0610 15:07:16.362943       1 leases.go:57] leader lost: leases
```


