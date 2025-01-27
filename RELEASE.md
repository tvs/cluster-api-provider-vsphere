# Release Process

The Kubernetes cluster-api-provider-vsphere images are released on an as-needed basis. The process is as follows:

1. An issue is proposing a new release with a changelog since the last release.
1. All [OWNERS](OWNERS) must LGTM this release.
1. An OWNER creates a changelog using release tools.
1. An OWNER runs `git tag -s $VERSION` and inserts the changelog and pushes the tag with `git push $VERSION`.
1. The release issue is closed.
1. An announcement email is sent to `kubernetes-sig-cluster-lifecycle@googlegroups.com` with the subject `[ANNOUNCE] cluster-api-provider-vsphere $VERSION is released`.
