# kube-golang

Hi Team,

i created a small Command Line Interface for manging kube using client-go library

below are the compiling options
to compile the Program inside the project folder -> 
go install
go build -o kubetest .



to start you need to enter: ./kubetest -version="<version number>" -scale=<replicas> --kubeconfig=</path/to/file> 

which will create a initial deployment and will you the options below 

Please Select Options to continue
	get ->  To GET deployment,
	list ->  To List the Kube Deployments
	delete ->  To Delete the Current Kube Deployment
	exit -> To Exit the CLI 


<img width="1246" alt="image" src="https://user-images.githubusercontent.com/7764674/195016467-4c422b4c-2eb9-45a0-ba02-413291c3f5c0.png">

type list to list the kube deployments 
<img width="1246" alt="image" src="https://user-images.githubusercontent.com/7764674/195017370-d0e243bd-3d33-479d-8a7a-680ac5a0bb77.png">

type get to get the deployments -> will ask for entering the deployment name
<img width="1246" alt="image" src="https://user-images.githubusercontent.com/7764674/195017549-f2209c5e-62f9-4e47-baeb-a7828a56b5fc.png">

<img width="1246" alt="image" src="https://user-images.githubusercontent.com/7764674/195017676-42e0389c-0fc7-4645-bfa8-e1e9e67b23bd.png">

delete will delete the current deployment
<img width="1246" alt="image" src="https://user-images.githubusercontent.com/7764674/195017819-ddc05317-8ecb-4930-9e5a-9574348f72df.png">

exit will exit the cli

