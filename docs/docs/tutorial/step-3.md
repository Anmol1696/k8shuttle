# **Step 3:** Spin up Starship

In this step, we will spin up couple of cosmos chains and relayers between them.

By the end of this tutorial you should be able to have osmosis and gaia chain running on your machine.

## Setup Starship Helm charts
We use helm as the package manager for starship.

Run
```bash
helm repo add starship https://cosmology-tech.github.io/starship/
helm repo update
helm search repo starship/devnet --version 0.1.26
```

## Define the infrastructure
We will now define the infrastructure for our starship, specify the chains and relayers run
between them.

We will run:
* Osmosis chain (single validator)
* Cosmos hub chain (single validator)
* Hermes relayer between osmosis and cosmos chain
* Ping Pub explorer for the chains

Create a file `starship.yaml` with the following content
```yaml
chains:
  - name: osmosis-1
    type: osmosis
    numValidators: 1
    ports:
      rest: 1313
      rpc: 26653
  - name: gaia-1
    type: cosmos
    numValidators: 1
    ports:
      rest: 1317
      rpc: 26657

relayers:
  - name: osmos-gaia
    type: hermes
    replicas: 1
    chains:
      - osmosis-1
      - gaia-1

explorer:
  enabled: true
  ports:
    rest: 8080
```

Above configuration would use around 3 CPUs and 4GB of RAM.
If you are constanstrained on resources, checkout the next step for `tiny-starship.yaml`

Documentaion has more details on the configuration options and how to reduce the resource usage.
For the tutorial we will keep it simple.

## Spin up the infrastructure
Spin up the infrastructure with
```bash
helm install -f starship.yaml tutorial starship/devnet --version 0.1.27 --wait --timeout 20m

# Where
# -f starship.yaml: use the starship.yaml file as the configuration
# tutorial: name of the helm release
# starship/devnet: helm chart to use
# --version 0.1.27: version of the helm chart to use
# --wait: wait for the infrastructure to be ready
# --timeout 20m: timeout after 20 minutes
```

This will take some time to spin up the infrastructure, you can check the status in another terminal with
```bash
kubectl get pods
# OR, to watch the pods
kubectl get pods -w
```

You would need to wait for the pods to be in `Running` state. 
This can take upto 2-5 minutes, depending on the underlying machine.

## Connect to the cluster
Once the pods are in `Running` state, you can port forward all the nodes to local host with
```bash
# Download the port-forward script
curl -Ls https://raw.githubusercontent.com/cosmology-tech/starship/main/scripts/port-forward.sh > port-forward.sh
# Make it executable
chmod +x port-forward.sh
# Run it with the config file
./port-forward.sh --config=starship.yaml
```

This will forward the `ports` of the nodes in the `starship.yaml` to your local machine.

## Interact with the chains
You can then interact with the chain on localhost at
* Osmosis: http://localhost:26653/status
* Cosmos: http://localhost:26657/status

And open up the explorer at http://localhost:8080
