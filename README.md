# cryptodevops

Repository for simple DevOps tasks.

## Directory Structure

- `.github/workflows`: A simple build and deploy pipeline that reaccts upon commits to the main branch. It builds the image described in the `/docker` directory, pushes it to GCR and deploys it as a `StatefulSet` to a GKE cluster.
- `converter`: A simple Go application that converts a CSV file to JSON.
- `docker`: Contains a simple Dockerfile that creates a container image of the Litecoin project
- `manifests`: Contains the manifest for a Kubernetes StatefulSet, leveraging [GKE's functionality](https://cloud.google.com/kubernetes-engine/docs/concepts/statefulset) for automating persistent disk storage.
- `terraform`: Creates a 3 node GKE cluster which is used to deploy the `StatefulSet`. It also ontains a module that defines a service account, a custom role with permissions `iam.roles.list` and a project member for assigning said role to the service account that was just created.
