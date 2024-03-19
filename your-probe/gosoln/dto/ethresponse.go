package dto

type EthLatestBlockResponse struct {
	Jsonrpc string `json:"jsonrpc,omitempty"`
	Result  struct {
		BaseFeePerGas         string   `json:"baseFeePerGas,omitempty"`
		BlobGasUsed           string   `json:"blobGasUsed,omitempty"`
		Difficulty            string   `json:"difficulty,omitempty"`
		ExcessBlobGas         string   `json:"excessBlobGas,omitempty"`
		ExtraData             string   `json:"extraData,omitempty"`
		GasLimit              string   `json:"gasLimit,omitempty"`
		GasUsed               string   `json:"gasUsed,omitempty"`
		Hash                  string   `json:"hash,omitempty"`
		LogsBloom             string   `json:"logsBloom,omitempty"`
		Miner                 string   `json:"miner,omitempty"`
		MixHash               string   `json:"mixHash,omitempty"`
		Nonce                 string   `json:"nonce,omitempty"`
		Number                string   `json:"number,omitempty"`
		ParentBeaconBlockRoot string   `json:"parentBeaconBlockRoot,omitempty"`
		ParentHash            string   `json:"parentHash,omitempty"`
		ReceiptsRoot          string   `json:"receiptsRoot,omitempty"`
		Sha3Uncles            string   `json:"sha3Uncles,omitempty"`
		Size                  string   `json:"size,omitempty"`
		StateRoot             string   `json:"stateRoot,omitempty"`
		Timestamp             string   `json:"timestamp,omitempty"`
		TotalDifficulty       string   `json:"totalDifficulty,omitempty"`
		Transactions          []string `json:"transactions,omitempty"`
		TransactionsRoot      string   `json:"transactionsRoot,omitempty"`
		Uncles                []any    `json:"uncles,omitempty"`
		Withdrawals           []struct {
			Address        string `json:"address,omitempty"`
			Amount         string `json:"amount,omitempty"`
			Index          string `json:"index,omitempty"`
			ValidatorIndex string `json:"validatorIndex,omitempty"`
		} `json:"withdrawals,omitempty"`
		WithdrawalsRoot string `json:"withdrawalsRoot,omitempty"`
	} `json:"result,omitempty"`
	ID int `json:"id,omitempty"`
}
