# 9thAge Battle Simulator
Tool to help Math Hammer

## TODO

- Need to allow for charging
- Extra specalties for various races
- Mounts
- Multiple rounds of combat
- Unit type (chariot gets impact hits)
- Skirmishers is wider and need to be handeled
- Multi-unit combat (eg. 2 v 1)
- Flanks and Rears
- To wound modifiers
- Differentiate ageis and fortitude saves
- Reroll passes for things like to wound and armour saves

## Ways to run the app:
Note: the first 3 options will expose the app on `localhost:9000`

Method | Location | Command
--- | --- | ---
Go RUN | Local | `go run ./cmd/main.go`
Docker | Local - Docker| <ol> <li> `docker build -t duxbuse/9thAgeSimulator .`</li> <li>`docker run -it --rm -p 9000:9000 duxbuse/9thAgeSimulator`</li> An alternative to `docker run` is to use the compose file `docker-compose up`
Minikube | Local - Kubernetes| <ol> <li> `docker build -t duxbuse/9thAgeSimulator .`</li> <li>`docker stack deploy -c docker-compose.yml 9thAgeSimulator`</li>
Google Cloud Platform | Cloud | <ol> <li>`cd Step1.google-container-cluster.terraform/`</li> <li>`terraform init` </li><li>`terraform apply` -> `yes`\*\*</li><li>`cd ../Step2.google-cloudbuild-trigger.terraform/` </li><li>`terraform init` </li><li>`terraform apply` -> `yes` \*\* </li></ol>

### It should be noted that deploying to cloud has a few pre-requisites.

** This will require a `xxx.json` key file for a service account that you will need to create manually with `Editor` permissions. It will also need `Source Repository Administrator` roles.

```
provider "google" {
  project     = "${var.project}"
  credentials = "${file(">>9thAgeSimulator-220503-8497483a16e9.json<<")}"
  region      = "${var.region}"
  zone        = "${var.zone}"
}
```
You will likley need to run terraform a few times with errors telling you to enable various api's but that is ok.
You will also need to modify the permissions of the `cloudbuild.gserviceaccount.com` account to also allow `Kubernetes Engine Admin` permissions which will alow the cloudbuild to deploy images that succeed. As well as changing the `vars.tf` in step 1 to your project name.

An additional requirement is that you will need to mirror your github with this code to your new `Google Source repository` that Step2 will create for you. It is also important to set the container registry to public so that your build step can have unfettered access to your image. Otherwise there is additional steps required in setting permissions for this.

 ** **Note that untill the cloud build runs you wont have anything deployed.** ** You can run this trigger manuaully the first time if you want to see it working.

To view the app you will need to extract the `Endpoint` Ip address from the ingress service, which can be found in the kubernetes tab.

**What you get at the end:**
At the end of this you will have a CI pipe line that will run go tests before building go code and loading it into a lightweight alpine image and deploying to the cluster everytime there is a push to the master branch of your repository.

For new project the kubernetes deployment yaml needs to point to the new GCR created for your project so with will need to be updated if you spin this up yourself

## Debugging
It is helpfull to run dry build of the cloud builder
`cloud-build-local --dryrun=false --substitutions=REPO_NAME='test-repo',REVISION_ID='test-revision'  .`