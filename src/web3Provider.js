import Web3Modal from "web3modal";
import WalletConnectProvider from "@walletconnect/web3-provider";

const providerOptions = {
  walletconnect: {
    package: WalletConnectProvider,
    options: {
      infuraId: "TODO",
    }
  },
};

const web3Modal = new Web3Modal({
  network: "rinkeby",
  cacheProvider: true,
  providerOptions,
  theme: "dark"
});

export default web3Modal
