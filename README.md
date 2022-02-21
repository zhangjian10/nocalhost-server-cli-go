<p align="center">
  <a href="https://github.com/actions/typescript-action/actions"><img alt="typescript-action status" src="https://github.com/actions/typescript-action/workflows/build-test/badge.svg"></a>
</p>

# Nocalhost server cli go

This action handles [nocalhost](https://nocalhost.dev/zh-CN/docs/introduction) server side related operations

Now supports devspace vcluster create and delete

# Usage

## login

```shell
nh-server-cli login -u $email -p $password -h $hostname
```

## vcluster

### create

```shell
nh-server-cli devspace create --id $cluster_id
```

output:

```shell
ID="1"
KUBECONFIG=""
```

### delete

```shell
nh-server-cli devspace delete --id $devspace_id
```
