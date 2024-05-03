import { GasPrice } from "@cosmjs/stargate";

export interface Network {
  chainId: string;
  rpcEndpoint: string;
  prefix: string;
  denom: string;
  gasPrice: string;
  feeToken: string;
}

// TODO: move to a api server to host chain registry
const chainRegistry = {
  osmosis: {
    chainId: "osmosis-1",
    rpcEndpoint: "http://localhost:26657",
    prefix: "osmo",
    denom: "uosmo",
    gasPrice: "0.025uosmo",
    feeToken: "uosmo",
  },
  cosmos: {
    chainId: "cosmos-2",
    rpcEndpoint: "http://localhost:26653",
    prefix: "cosmos",
    denom: "stake",
    gasPrice: "0.025stake",
    feeToken: "stake",
  },
  juno: {
    chainId: "juno-0",
    rpcEndpoint: "http://localhost:26655",
    prefix: "juno",
    denom: "ujuno",
    gasPrice: "0.025ujuno",
    feeToken: "ujuno",
  },
};

export function getChainInfo(network: string) {
  // @ts-ignore
  let chainRegistryElement = chainRegistry[network];
  return chainRegistryElement;
}
